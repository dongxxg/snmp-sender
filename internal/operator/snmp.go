/**
  @Author: dongxx
  @Since: 2026/3/25 17:52
  @Desc: //TODO
**/

package operator

import (
	"github.com/gosnmp/gosnmp"
	"snmp-sender/internal/output"
	"unitechs.com/unios-dice/uni-base/core/log"
)

// ──────────────────────────────────────────
// SNMP 操作
// ──────────────────────────────────────────

func doGet(client *gosnmp.GoSNMP, oids []string) {
	log.Println("=== GET ===")
	result, err := client.Get(oids)
	if err != nil {
		log.Fatalf("GET 失败: %v", err)
	}
	output.PrintPDUs(result.Variables)
}

func DoWalk(client *gosnmp.GoSNMP, oid string) {
	log.Println("=== WALK ===")
	err := client.Walk(oid, func(pdu gosnmp.SnmpPDU) error {
		output.PrintPDU(pdu)
		return nil
	})
	if err != nil {
		log.Errorf("WALK 失败: %v", err)
	}
}

func DoSet(client *gosnmp.GoSNMP, oid string, value string) {
	log.Println("=== SET ===")
	pdus := []gosnmp.SnmpPDU{
		{
			Name:  oid,
			Type:  gosnmp.OctetString,
			Value: value,
		},
	}
	result, err := client.Set(pdus)
	if err != nil {
		log.Errorf("SET 失败: %v", err)
		return
	}
	log.Printf("SET 响应错误状态: %v\n", result.Error)
	output.PrintPDUs(result.Variables)
}