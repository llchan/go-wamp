// NOTE: just a simple, naive implementation for now
package wamp

import (
	"errors"
	"regexp"
)

var ErrEmptyMessage = errors.New("empty message")
var ErrBadMessageType = errors.New("bad message type")
var ErrIdOverflow = errors.New("id overflow")
var ErrBadUri = errors.New("bad uri")

type Message struct {
	Type  MessageType
	Parts []MessagePart
}

type MessageType uint64

const (
	UNKNOWN      MessageType = 0
	HELLO        MessageType = 1
	WELCOME      MessageType = 2
	ABORT        MessageType = 3
	CHALLENGE    MessageType = 4
	AUTHENTICATE MessageType = 5
	GOODBYE      MessageType = 6
	HEARTBEAT    MessageType = 7
	ERROR        MessageType = 8
	PUBLISH      MessageType = 16
	PUBLISHED    MessageType = 17
	SUBSCRIBE    MessageType = 32
	SUBSCRIBED   MessageType = 33
	UNSUBSCRIBE  MessageType = 34
	UNSUBSCRIBED MessageType = 35
	EVENT        MessageType = 36
	CALL         MessageType = 48
	CANCEL       MessageType = 49
	RESULT       MessageType = 50
	REGISTER     MessageType = 64
	REGISTERED   MessageType = 65
	UNREGISTER   MessageType = 66
	UNREGISTERED MessageType = 67
	INVOCATION   MessageType = 68
	INTERRUPT    MessageType = 69
	YIELD        MessageType = 70
)

// XXX: This is a rather silly interface that really only
// helps us catch type-related errors at compile time.
type MessagePart interface {
	Value() interface{}
	Validate() error
}

// TODO: investigate whether we need to copy these before returning pointer
type Integer uint64
func (i Integer) Value() interface{} { return i }
func (i Integer) Validate() error { return nil }

type String string
func (s String) Value() interface{} { return s }
func (s String) Validate() error { return nil }

type Bool bool
func (b Bool) Value() interface{} { return b }
func (b Bool) Validate() error { return nil }

type Id uint64
const MaxId Id = 9007199254740992
func (id Id) Value() interface{} { return id }
func (id Id) Validate() error {
	if id > MaxId {
		return ErrIdOverflow
	}
	return nil
}

type Uri string
var UriRegexp = regexp.MustCompile(`^([^\s\.#]+)(\.([^\s\.#]+))*$`)
func (u Uri) Value() interface{} { return u }
func (u Uri) Validate() error {
	if !UriRegexp.MatchString(string(u)) {
		return ErrBadUri
	}
	return nil
}

type Dict map[string]MessagePart
func (d Dict) Value() interface{} { return d }
func (d Dict) Validate() error {
	for _, v := range d {
		if v == nil {
			continue
		}
		if err := v.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type List []MessagePart
func (l List) Value() interface{} { return l }
func (l List) Validate() error {
	for _, v := range l {
		if v == nil {
			continue
		}
		if err := v.Validate(); err != nil {
			return err
		}
	}
	return nil
}
