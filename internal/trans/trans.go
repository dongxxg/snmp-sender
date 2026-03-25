/**
  @Author: dongxx
  @Since: 2026/3/25 17:59
  @Desc: //TODO
**/

package trans

import (
	"github.com/gosnmp/gosnmp"
	"strings"
)

// ──────────────────────────────────────────
// 枚举转换
// ──────────────────────────────────────────

func SecurityLevel(level string) gosnmp.SnmpV3MsgFlags {
	switch strings.ToLower(level) {
	case "authpriv":
		return gosnmp.AuthPriv
	case "authnopriv":
		return gosnmp.AuthNoPriv
	default:
		return gosnmp.NoAuthNoPriv
	}
}

func AuthProtocol(proto string) gosnmp.SnmpV3AuthProtocol {
	switch strings.ToUpper(proto) {
	case "SHA":
		return gosnmp.SHA
	case "SHA224":
		return gosnmp.SHA224
	case "SHA256":
		return gosnmp.SHA256
	case "SHA384":
		return gosnmp.SHA384
	case "SHA512":
		return gosnmp.SHA512
	default:
		return gosnmp.MD5
	}
}

func PrivProtocol(proto string) gosnmp.SnmpV3PrivProtocol {
	switch strings.ToUpper(proto) {
	case "AES":
		return gosnmp.AES
	case "AES192":
		return gosnmp.AES192
	case "AES256":
		return gosnmp.AES256
	case "DES":
		return gosnmp.DES
	default:
		return gosnmp.AES
	}
}