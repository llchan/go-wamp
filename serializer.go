package wamp

type Serializer interface {
	Serialize(*Message) ([]byte, error)
	Deserialize([]byte, *Message) error
}
