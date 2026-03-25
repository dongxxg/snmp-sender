package internal

type Oid struct {
	Oid   string `json:"oid,omitempty"`
	Value string `json:"value,omitempty"`
	Type  string `json:"type,omitempty"` // str,int,
}

type Trap struct {
	TrapOid string `json:"trap_oid,omitempty"`
	Oids    []Oid  `json:"oids,omitempty"`
}
