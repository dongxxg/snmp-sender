/**
  @Author: dongxx
  @Since: 2026/3/25 16:57
  @Desc: //TODO
**/

package main

import (
	"encoding/json"
	"github.com/gosnmp/gosnmp"
	"io"
	"os"
	"snmp-sender/internal"
	"snmp-sender/internal/client"
	"strings"
	"time"
	"unitechs.com/unios-dice/uni-base/core/log"
)

func normalizeOID(oid string) string {
	oid = strings.TrimSpace(oid)
	if oid == "" {
		return ""
	}
	if strings.HasPrefix(oid, ".") {
		return oid
	}
	return "." + oid
}

func main() {
	loadTraps := func(path string) ([]internal.Trap, error) {
		open, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		defer open.Close()
		raw, err := io.ReadAll(open)
		if err != nil {
			return nil, err
		}
		var traps []internal.Trap
		if err := json.Unmarshal(raw, &traps); err != nil {
			return nil, err
		}
		return traps, nil
	}

	buildTrap := func(t internal.Trap, start time.Time) (gosnmp.SnmpTrap, error) {
		ticks := uint32(time.Since(start).Milliseconds() / 10) // SNMP TimeTicks: 1/100s

		vars := []gosnmp.SnmpPDU{
			{
				Name:  ".1.3.6.1.2.1.1.3.0", // sysUpTime.0
				Type:  gosnmp.TimeTicks,
				Value: ticks,
			},
			{
				Name:  ".1.3.6.1.6.3.1.1.4.1.0", // snmpTrapOID.0
				Type:  gosnmp.ObjectIdentifier,
				Value: normalizeOID(t.TrapOid),
			},
		}

		for _, f := range t.Oids {
			// 兼容两种 trap.json 写法：
			// 1) oid 填表5字段名（EsightRealtimeAlarmFieldDefs 的 key，如 iMAPNorthboundAlarmCSN）
			// 2) oid 填 base OID（如 1.3.6.1.4.1.2011.2...3）
			pdu, ok, err := internal.BuildEsightRealtimeAlarmPDU(f.Oid, f.Value)
			if err != nil {
				return gosnmp.SnmpTrap{}, err
			}
			if !ok {
				pdu2, ok2, err2 := internal.BuildEsightRealtimeAlarmPDUFromOID(f.Oid, f.Value)
				if err2 != nil {
					return gosnmp.SnmpTrap{}, err2
				}
				if !ok2 {
					log.Printf("unknown table5 field/oid: %s, skip", f.Oid)
					continue
				}
				pdu = pdu2
			}
			vars = append(vars, pdu)
		}

		return gosnmp.SnmpTrap{Variables: vars}, nil
	}

	log.Println("Starting")

	// 使用 SNMPv3 配置，通过配置文件读取并构造 GoSNMP 客户端
	// gosnmp.Default 的设置移除，改为通过 client.Build() 获取配置好的客户端

	// 兼容两种放置方式：根目录 `trap.json` 或 `internal/trap.json`
	traps, err := loadTraps("trap.json")
	if err != nil {
		traps, err = loadTraps("internal/trap.json")
	}
	if err != nil {
		log.Fatal(err)
	}
	if len(traps) == 0 {
		log.Fatal("trap.json contains no traps")
	}

	start := time.Now()

	// 构造并连接一个基于配置的 SNMPv3 客户端
	g, err := client.Build()
	if err != nil {
		log.Fatal(err)
	}
	if err := g.Connect(); err != nil {
		log.Fatal(err)
	}
	defer g.Conn.Close()

	idx := 0
	for {
		t := traps[idx]
		snmpTrap, err := buildTrap(t, start)
		if err != nil {
			log.Printf("build trap failed: %v", err)
		} else {
			if _, err := g.SendTrap(snmpTrap); err != nil {
				log.Printf("SendTrap failed: %v", err)
			}
		}

		idx++
		if idx >= len(traps) {
			idx = 0
		}
		time.Sleep(1 * time.Second)
	}
}
