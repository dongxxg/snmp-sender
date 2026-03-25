/**
  @Author: dongxx
  @Since: 2026/3/25 17:52
  @Desc: //TODO
**/

package client

import (
	"github.com/gosnmp/gosnmp"
	"snmp-sender/internal/trans"
	"time"
	"unitechs.com/unios-dice/uni-base/core/config"
)

// ──────────────────────────────────────────
// 构建客户端
// ──────────────────────────────────────────

func Build() (*gosnmp.GoSNMP, error) {
	params := &gosnmp.GoSNMP{
		Target:        config.GetString("Service.Target"),
		Port:          uint16(config.GetInt("Service.Port")),
		Version:       gosnmp.Version3,
		Timeout:       time.Duration(config.GetInt("Service.Timeout")) * time.Second,
		Retries:       config.GetInt("Service.Retries"),
		SecurityModel: gosnmp.UserSecurityModel,
		ContextName:   config.GetString("Service.ContextName"),
		MsgFlags:      trans.SecurityLevel(config.GetString("SNMP.SecurityLevel")),
		SecurityParameters: &gosnmp.UsmSecurityParameters{
			UserName:                 config.GetString("SNMP.UserName"),
			AuthenticationProtocol:   trans.AuthProtocol(config.GetString("SNMP.AuthProtocol")),
			AuthenticationPassphrase: config.GetString("SNMP.AuthPassword"),
			PrivacyProtocol:          trans.PrivProtocol(config.GetString("SNMP.PrivProtocol")),
			PrivacyPassphrase:        config.GetString("SNMP.PrivPassword"),
		},
	}
	return params, nil
}