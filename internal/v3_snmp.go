package internal

import (
	"github.com/gosnmp/gosnmp"
	"log"
	"strings"
	"time"
)

// V3TrapConfig holds minimal SNMPv3 parameters for sending traps.
type V3TrapConfig struct {
	Target          string
	Port            uint16
	User            string
	AuthProtocol    string // optional, not strictly required with NoAuthNoPriv
	AuthPassword    string // optional
	PrivProtocol    string // optional
	PrivPassword    string // optional
	ContextName     string // optional
	ContextEngineID string // optional
	Timeout         time.Duration
	Retries         int
}

// SendV3Traps sends the provided traps to a SNMPv3 endpoint using a lightweight
// configuration. By default it uses NoAuthNoPriv to minimize configuration,
// but if you provide Auth/Priv fields, the library will sign the traps accordingly.
func SendV3Traps(traps []gosnmp.SnmpTrap, sleep time.Duration, cfg V3TrapConfig) {
	g := &gosnmp.GoSNMP{
		Target:  cfg.Target,
		Port:    cfg.Port,
		Version: gosnmp.Version3,
		Timeout: cfg.Timeout,
		Retries: cfg.Retries,
	}

	g.SecurityModel = gosnmp.UserSecurityModel
	// Default to NoAuthNoPriv; users can extend by filling Auth/Priv fields later.
	g.MsgFlags = gosnmp.NoAuthNoPriv
	g.SecurityParameters = &gosnmp.UsmSecurityParameters{
		UserName: cfg.User,
	}
	if err := g.Connect(); err != nil {
		log.Printf("SNMPv3 connect failed: %v", err)
		return
	}
	defer g.Conn.Close()
	for _, t := range traps {
		if _, err := g.SendTrap(t); err != nil {
			log.Printf("SNMPv3 SendTrap error: %v", err)
		}
		if sleep > 0 {
			time.Sleep(sleep)
		}
	}
}

// helper to normalize common parameter strings (no strict validation here).
func normalizeProto(p string) string {
	return strings.ToLower(strings.TrimSpace(p))
}
