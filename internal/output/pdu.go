/**
  @Author: dongxx
  @Since: 2026/3/25 17:53
  @Desc: //TODO
**/

package output

import (
	"github.com/gosnmp/gosnmp"
	"unitechs.com/unios-dice/uni-base/core/log"
)

// ──────────────────────────────────────────
// 输出
// ──────────────────────────────────────────

func PrintPDUs(pdus []gosnmp.SnmpPDU) {
	for _, pdu := range pdus {
		PrintPDU(pdu)
	}
}

func PrintPDU(pdu gosnmp.SnmpPDU) {
	switch pdu.Type {
	case gosnmp.OctetString:
		// gosnmp may store OctetString as either []byte or string depending on how we build the PDU.
		switch v := pdu.Value.(type) {
		case []byte:
			log.Printf("OID: %-40s  Type: OctetString  Value: %s\n", pdu.Name, string(v))
		case string:
			log.Printf("OID: %-40s  Type: OctetString  Value: %s\n", pdu.Name, v)
		default:
			log.Printf("OID: %-40s  Type: OctetString  Value: %v\n", pdu.Name, pdu.Value)
		}
	case gosnmp.Integer:
		log.Printf("OID: %-40s  Type: Integer      Value: %d\n", pdu.Name, gosnmp.ToBigInt(pdu.Value))
	case gosnmp.Counter32:
		log.Printf("OID: %-40s  Type: Counter32    Value: %d\n", pdu.Name, pdu.Value.(uint))
	case gosnmp.Counter64:
		log.Printf("OID: %-40s  Type: Counter64    Value: %d\n", pdu.Name, pdu.Value.(uint64))
	case gosnmp.Gauge32:
		log.Printf("OID: %-40s  Type: Gauge32      Value: %d\n", pdu.Name, pdu.Value.(uint))
	case gosnmp.TimeTicks:
		log.Printf("OID: %-40s  Type: TimeTicks    Value: %d\n", pdu.Name, pdu.Value.(uint32))
	case gosnmp.IPAddress:
		log.Printf("OID: %-40s  Type: IPAddress    Value: %s\n", pdu.Name, pdu.Value.(string))
	case gosnmp.ObjectIdentifier:
		log.Printf("OID: %-40s  Type: OID          Value: %s\n", pdu.Name, pdu.Value.(string))
	default:
		log.Printf("OID: %-40s  Type: %v        Value: %v\n", pdu.Name, pdu.Type, pdu.Value)
	}
}