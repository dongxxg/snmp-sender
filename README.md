# snmp-sender

Overview
- Go-based SNMP trap sender and syslog bridge for ESight-like alarms. It can construct SNMP traps from JSON definitions and send them to an SNMP manager. It also provides a lightweight syslog forwarder that converts syslog messages into SNMP traps.

make pipeline-pack platform=arm64
cd dist