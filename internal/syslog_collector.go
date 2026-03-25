package internal

import (
	"fmt"
	"github.com/gosnmp/gosnmp"
	"net"
	"strings"
	"time"
)

// StartSyslogForwarder runs a simple UDP syslog listener and forwards received
// messages as SNMPv3 traps. This is a minimal bridge intended for demo/testing.
// listenAddr should be in the form "host:port". Example: ":5140" or "0.0.0.0:10514".
// v3cfg provides the SNMPv3 destination configuration.
func StartSyslogForwarder(listenAddr string, v3cfg V3TrapConfig) error {
	addr, err := net.ResolveUDPAddr("udp", listenAddr)
	if err != nil {
		return fmt.Errorf("resolve udp addr: %w", err)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		return fmt.Errorf("listen udp: %w", err)
	}
	// Non-blocking loop: caller may choose to run this in a goroutine.
	defer conn.Close()

	buf := make([]byte, 65535)
	for {
		n, remote, err := conn.ReadFromUDP(buf)
		if err != nil {
			// If we can't read, skip this iteration and continue listening
			// to avoid crashing the bridge.
			fmt.Printf("syslog read error from %v: %v\n", remote, err)
			time.Sleep(100 * time.Millisecond)
			continue
		}
		msg := strings.TrimSpace(string(buf[:n]))
		if msg == "" {
			continue
		}
		// Build a trap containing the syslog message as the first variable.
		// This uses a simple, stable OID for demonstration purposes.
		trap := gosnmp.SnmpTrap{
			Variables: []gosnmp.SnmpPDU{
				{
					Name:  ".1.3.6.1.2.1.1.4.0", // sysContact - repurposed for the log message
					Type:  gosnmp.OctetString,
					Value: msg,
				},
			},
		}
		// Forward via SNMPv3; sleep not used here to keep responsiveness.
		SendV3Traps([]gosnmp.SnmpTrap{trap}, 0, v3cfg)
	}
}
