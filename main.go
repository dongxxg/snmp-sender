/**
  @Author: dongxx
  @Since: 2026/3/25 16:57
  @Desc: //TODO
**/

package main

import (
	"encoding/json"
	"fmt"
	"github.com/gosnmp/gosnmp"
	"io"
	"os"
	"snmp-sender/internal"
	"strconv"
	"strings"
	"sync"
	"time"
	"unitechs.com/unios-dice/uni-base/core/log"
)

func main() {
	log.Println("Starting")
	open, err := os.Open("conf.json")
	if err != nil {
		log.Println(err)
		return
	}
	defer open.Close()
	all, err := io.ReadAll(open)
	if err != nil {
		return
	}
	var conf internal.Conf
	_ = json.Unmarshal(all, &conf)
	gosnmp.Default.Target = conf.Target
	//gosnmp.Default.Transport = conf.Transport
	gosnmp.Default.Port = conf.Port
	gosnmp.Default.Community = conf.Community
	gosnmp.Default.Version = gosnmp.Version2c
	open2, err := os.Open("trap.json")
	if err != nil {
		return
	}
	defer open2.Close()
	all, err = io.ReadAll(open2)
	if err != nil {
		return
	}
	var traps []internal.Trap
	err = json.Unmarshal(all, &traps)
	if err != nil {
		log.Fatal(err)
	}
	sleep := time.Millisecond * time.Duration(conf.Sleep)
	ipFile, err := os.Open("ip.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ipFile.Close()
	all, err = io.ReadAll(ipFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	var sDatas []gosnmp.SnmpTrap
	ips := strings.Split(string(all), "\n")
	var ipCursor int
	for _, trap := range traps {
		var oids []gosnmp.SnmpPDU
		oids = append(oids, gosnmp.SnmpPDU{
			Name:  ".1.3.6.1.6.3.1.1.4.1.0",
			Type:  gosnmp.ObjectIdentifier,
			Value: ".1.3.6.1.6.3.1.1.5.1",
		})
		for _, oid := range trap.Oids {
			if oid.Value == "$IP" {
				oids = append(oids, gosnmp.SnmpPDU{
					Name:  oid.Oid,
					Value: ips[ipCursor],
					Type:  gosnmp.OctetString,
				})
				ipCursor += 1
				if ipCursor == len(ips) {
					ipCursor = 0
				}
			} else if oid.Value == "$NUM++" {
				oids = append(oids, gosnmp.SnmpPDU{
					Name:  oid.Oid,
					Value: fmt.Sprintf("xx%d", internal.Num),
					Type:  gosnmp.OctetString,
				})
				internal.Num += 1
			} else {
				if oid.Type == "str" {
					oids = append(oids, gosnmp.SnmpPDU{
						Name:  oid.Oid,
						Value: oid.Value,
						Type:  gosnmp.OctetString,
					})
				} else if oid.Type == "int" {
					val, _ := strconv.Atoi(oid.Value)
					oids = append(oids, gosnmp.SnmpPDU{
						Name:  oid.Oid,
						Value: val,
						Type:  gosnmp.Integer,
					})
				}
			}
		}

		t := gosnmp.SnmpTrap{
			Variables: oids,
			//Enterprise: ".1.3.6.1.6.3.1.1.5.1",
		}
		sDatas = append(sDatas, t)
	}
	for i := 0; i < conf.Repeats; i++ {
		fmt.Println("new:", i)
		t := time.Now()
		group := sync.WaitGroup{}
		for j := 0; j < conf.Threads; j++ {
			group.Add(1)
			go func() {
				internal.Send(sDatas, sleep)
				group.Done()
			}()
		}
		group.Wait()
		fmt.Println("finish:", i, " use time:", time.Now().Sub(t).String())
	}
}
