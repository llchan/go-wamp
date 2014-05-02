package jsonserializer

import (
	"github.com/llchan/go-wamp"
	"encoding/json"
)

type Serializer struct {
}

func (*Serializer) Serialize(m wamp.Message) ([]byte, error) {
	if m == nil {
		return nil, wamp.ErrNilMessage
	}
	return json.Marshal(m.Parts())
}

func (*Serializer) Deserialize(p []byte) (wamp.Message, error) {
	var raw []json.RawMessage
	if err := json.Unmarshal(p, &raw); err != nil {
		return nil, err
	}
	if len(raw) == 0 {
		return nil, wamp.ErrEmptyMessage
	}
	var t wamp.MessageType
	if err := json.Unmarshal(raw[0], &t); err != nil {
		return nil, err
	}
	m, err := wamp.NewMessage(t)
	if err != nil {
		return nil, err
	}
	numopt := 0
	vm, ok := m.(wamp.VariadicMessage)
	if ok {
		numopt = vm.NumOptional()
	}
	parts := m.Parts()
	if len(parts) == 0 {
		return nil, wamp.ErrBadPartsLength
	}
	N := len(raw) - 1
	M := len(parts) - 1
	if N < M - numopt || N > M {
		return nil, wamp.ErrBadMessageLength
	}
	for i := 0; i < N; i++ {
		if err := json.Unmarshal(raw[i + 1], parts[i + 1]); err != nil {
			return nil, err
		}
	}
	return m, nil
}
