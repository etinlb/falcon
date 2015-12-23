package network

import (
	"encoding/json"
	// "github.com/etinlb/falcon/core_lib"
)

type Message struct {
	Event string           `json:"event"`
	Data  *json.RawMessage `json:"data"` // how data is parsed depends on the event
	// SequenceNum int             `json:"seq_num"`
}

type Events struct {
	Events []Message `json:"events"`
}
