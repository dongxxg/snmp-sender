snmp-sender

Overview
- Go-based SNMP trap sender and syslog bridge for ESight-like alarms. It can construct SNMP traps from JSON definitions and send them to an SNMP manager. It also provides a lightweight syslog forwarder that converts syslog messages into SNMP traps.

What it does
- Builds SNMP traps from trap.json (or internal/trap.json)
- Supports SNMPv2c by default and SNMPv3 with configurable security (via internal/v3_snmp.go and internal/trans/*.go)
- Exposes a small bridge to forward UDP syslog messages as traps (StartSyslogForwarder)

How to run
- Prereqs: Go toolchain, network access to the SNMP manager
- Build: go build ./...
- Run: ./snmp-sender
- Optional: set trap.json or provide trap definitions under trap.json or internal/trap.json

Configuration and traps
- The app reads configuration from uni-base/core/config (as used by Build() in internal/client/snmp.go)
- A sample trap JSON is provided at internal/trap.json.sample

What to improve (next steps)
- Add tests for trap building and SNMP v3 configuration
- Add CI (go vet, golangci-lint, unit tests)
- Provide a clearer example trap.json with both field-name and base-oid variants
- Add a small Makefile to simplify builds and runs
