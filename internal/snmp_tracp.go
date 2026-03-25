package internal

import (
	"fmt"
	"github.com/gosnmp/gosnmp"
	"log"
	"time"
)

var Num int

type Oid struct {
	Oid   string `json:"oid,omitempty"`
	Value string `json:"value,omitempty"`
	Type  string `json:"type,omitempty"` // str,int,
}

type Trap struct {
	TrapOid string `json:"trap_oid,omitempty"`
	Oids    []Oid  `json:"oids,omitempty"`
}

func Send(sDatas []gosnmp.SnmpTrap, sleep time.Duration) {
	dft := *gosnmp.Default
	err := dft.Connect()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dft.Conn.Close()
	for _, t := range sDatas {
		_, err = dft.SendTrap(t)
		time.Sleep(sleep)
		if err != nil {
			log.Fatalf("SendTrap() err: %v\n", err)
		}
	}
	if err != nil {
		log.Fatal(err)
	}
}