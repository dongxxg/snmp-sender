package internal

import (
	"fmt"
	"github.com/gosnmp/gosnmp"
	"strconv"
	"strings"
)

// EsightRealtimeAlarmFieldDef describes a table5 field mapping.
// Table5 OIDs are "base OID" without instance suffix; real reported OID appends ".0".
type EsightRealtimeAlarmFieldDef struct {
	BaseOID  string
	SnmpType  gosnmp.Asn1BER
}

// EsightRealtimeAlarmFieldDefs maps table5 field names to their base OID and SNMP type.
// If you put these field names into trap.json's "oid" field, main.go will auto-build the varbind.
var EsightRealtimeAlarmFieldDefs = map[string]EsightRealtimeAlarmFieldDef{
	"iMAPNorthboundAlarmCSN":                         {BaseOID: "1.3.6.1.4.1.2011.2.15.2.4.3.3.1", SnmpType: gosnmp.OctetString},
	"iMAPNorthboundAlarmCategory":                   {BaseOID: "1.3.6.1.4.1.2011.2.15.2.4.3.3.2", SnmpType: gosnmp.OctetString},
	"iMAPNorthboundAlarmOccurTime":                  {BaseOID: "1.3.6.1.4.1.2011.2.15.2.4.3.3.3", SnmpType: gosnmp.OctetString},
	"iMAPNorthboundAlarmMOName":                    {BaseOID: "1.3.6.1.4.1.2011.2.15.2.4.3.3.4", SnmpType: gosnmp.OctetString},
	"iMAPNorthboundAlarmProductID":                {BaseOID: "1.3.6.1.4.1.2011.2.15.2.4.3.3.5", SnmpType: gosnmp.Integer},
	"iMAPNorthboundAlarmNEType":                   {BaseOID: "1.3.6.1.4.1.2011.2.15.2.4.3.3.6", SnmpType: gosnmp.OctetString},
	"iMAPNorthboundAlarmNEDevID":                  {BaseOID: "1.3.6.1.4.1.2011.2.15.2.4.3.3.7", SnmpType: gosnmp.OctetString},
	"iMAPNorthboundAlarmDevCsn":                   {BaseOID: "1.3.6.1.4.1.2011.2.15.2.4.3.3.8", SnmpType: gosnmp.OctetString},
	"iMAPNorthboundAlarmID":                       {BaseOID: "1.3.6.1.4.1.2011.2.15.2.4.3.3.9", SnmpType: gosnmp.Integer},
	"iMAPNorthboundAlarmType":                     {BaseOID: "1.3.6.1.4.1.2011.2.15.2.4.3.3.10", SnmpType: gosnmp.Integer},
	"iMAPNorthboundAlarmLevel":                    {BaseOID: "1.3.6.1.4.1.2011.2.15.2.4.3.3.11", SnmpType: gosnmp.Integer},
	"iMAPNorthboundAlarmRestore":                  {BaseOID: "1.3.6.1.4.1.2011.2.15.2.4.3.3.12", SnmpType: gosnmp.Integer},
	"iMAPNorthboundAlarmConfirm":                  {BaseOID: "1.3.6.1.4.1.2011.2.15.2.4.3.3.13", SnmpType: gosnmp.Integer},
	"iMAPNorthboundAlarmAckTime":                 {BaseOID: "1.3.6.1.4.1.2011.2.15.2.4.3.3.14", SnmpType: gosnmp.OctetString},
	"iMAPNorthboundAlarmRestoreTime":             {BaseOID: "1.3.6.1.4.1.2011.2.15.2.4.3.3.15", SnmpType: gosnmp.OctetString},
	"iMAPNorthboundAlarmOperator":               {BaseOID: "1.3.6.1.4.1.2011.2.15.2.4.3.3.16", SnmpType: gosnmp.OctetString},
	"iMAPNorthboundAlarmParams1":                {BaseOID: "1.3.6.1.4.1.2011.2.15.2.4.3.3.17", SnmpType: gosnmp.Integer},
	"iMAPNorthboundAlarmParams2":                {BaseOID: "1.3.6.1.4.1.2011.2.15.2.4.3.3.18", SnmpType: gosnmp.Integer},
	"iMAPNorthboundAlarmParams3":                {BaseOID: "1.3.6.1.4.1.2011.2.15.2.4.3.3.19", SnmpType: gosnmp.Integer},
	"iMAPNorthboundAlarmParams4":                {BaseOID: "1.3.6.1.4.1.2011.2.15.2.4.3.3.20", SnmpType: gosnmp.Integer},
	"iMAPNorthboundAlarmParams5":                {BaseOID: "1.3.6.1.4.1.2011.2.15.2.4.3.3.21", SnmpType: gosnmp.Integer},
	"iMAPNorthboundAlarmParams6":                {BaseOID: "1.3.6.1.4.1.2011.2.15.2.4.3.3.22", SnmpType: gosnmp.Integer},
	"iMAPNorthboundAlarmParams7":                {BaseOID: "1.3.6.1.4.1.2011.2.15.2.4.3.3.23", SnmpType: gosnmp.Integer},
	"iMAPNorthboundAlarmParams8":                {BaseOID: "1.3.6.1.4.1.2011.2.15.2.4.3.3.24", SnmpType: gosnmp.Integer},
	"iMAPNorthboundAlarmParams9":                {BaseOID: "1.3.6.1.4.1.2011.2.15.2.4.3.3.25", SnmpType: gosnmp.Integer},
	"iMAPNorthboundAlarmParams10":               {BaseOID: "1.3.6.1.4.1.2011.2.15.2.4.3.3.26", SnmpType: gosnmp.Integer},
	"iMAPNorthboundAlarmExtendInfo":            {BaseOID: "1.3.6.1.4.1.2011.2.15.2.4.3.3.27", SnmpType: gosnmp.OctetString},
	"iMAPNorthboundAlarmProbablecause":        {BaseOID: "1.3.6.1.4.1.2011.2.15.2.4.3.3.28", SnmpType: gosnmp.OctetString},
	"iMAPNorthboundAlarmProposedrepairactions": {BaseOID: "1.3.6.1.4.1.2011.2.15.2.4.3.3.29", SnmpType: gosnmp.OctetString},
	"iMAPNorthboundAlarmSpecificproblems":     {BaseOID: "1.3.6.1.4.1.2011.2.15.2.4.3.3.30", SnmpType: gosnmp.OctetString},
	"iMAPNorthboundAlarmExtendProductItem1":  {BaseOID: "1.3.6.1.4.1.2011.2.15.2.4.3.3.31", SnmpType: gosnmp.OctetString},
	"iMAPNorthboundAlarmExtendProductItem2":  {BaseOID: "1.3.6.1.4.1.2011.2.15.2.4.3.3.32", SnmpType: gosnmp.OctetString},
	"iMAPNorthboundAlarmExtendProductItem3":  {BaseOID: "1.3.6.1.4.1.2011.2.15.2.4.3.3.33", SnmpType: gosnmp.OctetString},
	"iMAPNorthboundAlarmExtendProductItem4":  {BaseOID: "1.3.6.1.4.1.2011.2.15.2.4.3.3.34", SnmpType: gosnmp.OctetString},
	"iMAPNorthboundAlarmExtendProductItem5":  {BaseOID: "1.3.6.1.4.1.2011.2.15.2.4.3.3.35", SnmpType: gosnmp.OctetString},
	"iMAPNorthboundAlarmExtendProductItem6":  {BaseOID: "1.3.6.1.4.1.2011.2.15.2.4.3.3.36", SnmpType: gosnmp.OctetString},
	"iMAPNorthboundAlarmExtendProductItem7":  {BaseOID: "1.3.6.1.4.1.2011.2.15.2.4.3.3.37", SnmpType: gosnmp.OctetString},
	"iMAPNorthboundAlarmExtendProductItem8":  {BaseOID: "1.3.6.1.4.1.2011.2.15.2.4.3.3.38", SnmpType: gosnmp.OctetString},
	"iMAPNorthboundAlarmExtendProductItem9":  {BaseOID: "1.3.6.1.4.1.2011.2.15.2.4.3.3.39", SnmpType: gosnmp.OctetString},
	"iMAPNorthboundAlarmExtendProductItem10": {BaseOID: "1.3.6.1.4.1.2011.2.15.2.4.3.3.40", SnmpType: gosnmp.OctetString},
	"iMAPNorthboundAlarmExtendProductItem11": {BaseOID: "1.3.6.1.4.1.2011.2.15.2.4.3.3.41", SnmpType: gosnmp.OctetString},
	"iMAPNorthboundAlarmExtendProductItem12": {BaseOID: "1.3.6.1.4.1.2011.2.15.2.4.3.3.42", SnmpType: gosnmp.OctetString},
	"iMAPNorthboundAlarmExtendProductItem13": {BaseOID: "1.3.6.1.4.1.2011.2.15.2.4.3.3.43", SnmpType: gosnmp.OctetString},
	"iMAPNorthboundAlarmExtendProductItem14": {BaseOID: "1.3.6.1.4.1.2011.2.15.2.4.3.3.44", SnmpType: gosnmp.OctetString},
	"iMAPNorthboundAlarmExtendProductItem15": {BaseOID: "1.3.6.1.4.1.2011.2.15.2.4.3.3.45", SnmpType: gosnmp.OctetString},
	"iMAPNorthboundAlarmClearOperator":        {BaseOID: "1.3.6.1.4.1.2011.2.15.2.4.3.3.46", SnmpType: gosnmp.OctetString},
	"iMAPNorthboundAlarmObjectInstanceType":  {BaseOID: "1.3.6.1.4.1.2011.2.15.2.4.3.3.47", SnmpType: gosnmp.OctetString},
	"iMAPNorthboundAlarmClearCategory":      {BaseOID: "1.3.6.1.4.1.2011.2.15.2.4.3.3.48", SnmpType: gosnmp.OctetString},
	"iMAPNorthboundAlarmClearType":          {BaseOID: "1.3.6.1.4.1.2011.2.15.2.4.3.3.49", SnmpType: gosnmp.OctetString},
	"iMAPNorthboundAlarmServiceAffectFlag": {BaseOID: "1.3.6.1.4.1.2011.2.15.2.4.3.3.50", SnmpType: gosnmp.OctetString},
	"iMAPNorthboundAlarmAdditionalInfo":    {BaseOID: "1.3.6.1.4.1.2011.2.15.2.4.3.3.51", SnmpType: gosnmp.OctetString},
}

// Reverse map for when trap.json uses base OID instead of field name key.
var esightRealtimeAlarmFieldDefsByBaseOID map[string]EsightRealtimeAlarmFieldDef

func init() {
	esightRealtimeAlarmFieldDefsByBaseOID = make(map[string]EsightRealtimeAlarmFieldDef, len(EsightRealtimeAlarmFieldDefs))
	for k, v := range EsightRealtimeAlarmFieldDefs {
		_ = k
		esightRealtimeAlarmFieldDefsByBaseOID[v.BaseOID] = v
	}
}

// BuildEsightRealtimeAlarmPDU builds a varbind for table5 field name.
// Returns ok=false if fieldName is not in table5 mapping.
func BuildEsightRealtimeAlarmPDU(fieldName string, value string) (pdu gosnmp.SnmpPDU, ok bool, err error) {
	def, ok := EsightRealtimeAlarmFieldDefs[fieldName]
	if !ok {
		return gosnmp.SnmpPDU{}, false, nil
	}
	oid := "." + strings.TrimPrefix(def.BaseOID, ".") + ".0"

	switch def.SnmpType {
	case gosnmp.Integer:
		v := strings.TrimSpace(value)
		if v == "" {
			v = "0"
		}
		i, convErr := strconv.Atoi(v)
		if convErr != nil {
			return gosnmp.SnmpPDU{}, true, fmt.Errorf("parse integer32 %s=%q: %w", fieldName, value, convErr)
		}
		return gosnmp.SnmpPDU{
			Name:  oid,
			Type:  gosnmp.Integer,
			Value: i,
		}, true, nil
	default:
		// OCTET STRING: send bytes; empty => zero-length bytes.
		return gosnmp.SnmpPDU{
			Name:  oid,
			Type:  def.SnmpType,
			Value: []byte(value),
		}, true, nil
	}
}

// BuildEsightRealtimeAlarmPDUFromOID builds a varbind for table5 base OID.
// trap.json can provide either base OID (without ".0") or full instance OID (with ".0").
// Returns ok=false if oid doesn't match table5 base OIDs.
func BuildEsightRealtimeAlarmPDUFromOID(oid string, value string) (pdu gosnmp.SnmpPDU, ok bool, err error) {
	oid = strings.TrimSpace(oid)
	oid = strings.TrimPrefix(oid, ".")
	oid = strings.TrimSuffix(oid, ".0")

	def, ok := esightRealtimeAlarmFieldDefsByBaseOID[oid]
	if !ok {
		return gosnmp.SnmpPDU{}, false, nil
	}

	fullOID := "." + def.BaseOID + ".0"

	switch def.SnmpType {
	case gosnmp.Integer:
		v := strings.TrimSpace(value)
		if v == "" {
			v = "0"
		}
		i, convErr := strconv.Atoi(v)
		if convErr != nil {
			return gosnmp.SnmpPDU{}, true, fmt.Errorf("parse integer32 oid=%s value=%q: %w", oid, value, convErr)
		}
		return gosnmp.SnmpPDU{
			Name:  fullOID,
			Type:  gosnmp.Integer,
			Value: i,
		}, true, nil
	default:
		return gosnmp.SnmpPDU{
			Name:  fullOID,
			Type:  def.SnmpType,
			Value: []byte(value),
		}, true, nil
	}
}

