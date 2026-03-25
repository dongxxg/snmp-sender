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
		var v uint32
		switch val := pdu.Value.(type) {
		case uint:
			v = uint32(val)
		case uint32:
			v = val
		case int:
			v = uint32(val)
		default:
			v = 0
		}
		log.Printf("OID: %-40s  Type: Counter32    Value: %d\n", pdu.Name, v)
	case gosnmp.Counter64:
		var v uint64
		switch val := pdu.Value.(type) {
		case uint64:
			v = val
		case uint:
			v = uint64(val)
		case int:
			v = uint64(val)
		default:
			v = 0
		}
		log.Printf("OID: %-40s  Type: Counter64    Value: %d\n", pdu.Name, v)
	case gosnmp.Gauge32:
		var v uint32
		switch val := pdu.Value.(type) {
		case uint:
			v = uint32(val)
		case uint32:
			v = val
		case int:
			v = uint32(val)
		default:
			v = 0
		}
		log.Printf("OID: %-40s  Type: Gauge32      Value: %d\n", pdu.Name, v)
	case gosnmp.TimeTicks:
		var v uint32
		switch val := pdu.Value.(type) {
		case uint32:
			v = val
		case uint:
			v = uint32(val)
		case int:
			v = uint32(val)
		default:
			v = 0
		}
		log.Printf("OID: %-40s  Type: TimeTicks    Value: %d\n", pdu.Name, v)
	case gosnmp.IPAddress:
		log.Printf("OID: %-40s  Type: IPAddress    Value: %s\n", pdu.Name, pdu.Value.(string))
	case gosnmp.ObjectIdentifier:
		log.Printf("OID: %-40s  Type: OID          Value: %s\n", pdu.Name, pdu.Value.(string))
	default:
		log.Printf("OID: %-40s  Type: %v        Value: %v\n", pdu.Name, pdu.Type, pdu.Value)
	}
}
