// NOTE: just a simple, naive implementation for now
package wamp

import (
	"errors"
	"regexp"
)

var ErrNilMessage = errors.New("nil message")
var ErrEmptyMessage = errors.New("empty message")
var ErrBadMessageType = errors.New("bad message type")
var ErrBadMessageLength = errors.New("bad message length")
var ErrBadPartsLength = errors.New("bad parts length")
var ErrIdOverflow = errors.New("id overflow")
var ErrBadUri = errors.New("bad uri")
var ErrNilDict = errors.New("nil dict")
var ErrNilList = errors.New("nil list")

type Message interface {
	Type() MessageType
	Parts() []MessagePart
}

func NewMessage(t MessageType) (m Message, err error) {
	switch t {
		case HELLO: m = &HelloMessage{}
		case WELCOME: m = &WelcomeMessage{}
		case ABORT: m = &AbortMessage{}
		// case CHALLENGE: m = &ChallengeMessage{}
		// case AUTHENTICATE: m = &AuthenticateMessage{}
		case GOODBYE: m = &GoodbyeMessage{}
		// case HEARTBEAT: m = &HeartbeatMessage{}
		case ERROR: m = &ErrorMessage{}
		case PUBLISH: m = &PublishMessage{}
		case PUBLISHED: m = &PublishedMessage{}
		case SUBSCRIBE: m = &SubscribeMessage{}
		case SUBSCRIBED: m = &SubscribedMessage{}
		case UNSUBSCRIBE: m = &UnsubscribeMessage{}
		case UNSUBSCRIBED: m = &UnsubscribedMessage{}
		case EVENT: m = &EventMessage{}
		case CALL: m = &CallMessage{}
		// case CANCEL: m = &CancelMessage{}
		case RESULT: m = &ResultMessage{}
		case REGISTER: m = &RegisterMessage{}
		case REGISTERED: m = &RegisteredMessage{}
		case UNREGISTER: m = &UnregisterMessage{}
		case UNREGISTERED: m = &UnregisteredMessage{}
		case INVOCATION: m = &InvocationMessage{}
		// case INTERRUPT: m = &InterruptMessage{}
		case YIELD: m = &YieldMessage{}
		default:
	}
	return m, err
}

// TODO: do this more elegantly
type VariadicMessage interface {
	NumOptional() int
}

// MessagePart is a wrapper interface to add some compile-time safety and
// validation to our messages
type MessagePart interface {
	Value() interface{}
	Validate() error
}

type MessageType uint64
const (
	UNKNOWN      MessageType = 0
	HELLO        MessageType = 1
	WELCOME      MessageType = 2
	ABORT        MessageType = 3
	// CHALLENGE    MessageType = 4
	// AUTHENTICATE MessageType = 5
	GOODBYE      MessageType = 6
	// HEARTBEAT    MessageType = 7
	ERROR        MessageType = 8
	PUBLISH      MessageType = 16
	PUBLISHED    MessageType = 17
	SUBSCRIBE    MessageType = 32
	SUBSCRIBED   MessageType = 33
	UNSUBSCRIBE  MessageType = 34
	UNSUBSCRIBED MessageType = 35
	EVENT        MessageType = 36
	CALL         MessageType = 48
	// CANCEL       MessageType = 49
	RESULT       MessageType = 50
	REGISTER     MessageType = 64
	REGISTERED   MessageType = 65
	UNREGISTER   MessageType = 66
	UNREGISTERED MessageType = 67
	INVOCATION   MessageType = 68
	// INTERRUPT    MessageType = 69
	YIELD        MessageType = 70
)
func (m MessageType) Value() interface{} { return m }
func (m MessageType) Validate() error {
	switch m {
		case HELLO: return nil
		case WELCOME: return nil
		case ABORT: return nil
		// case CHALLENGE: return nil
		// case AUTHENTICATE: return nil
		case GOODBYE: return nil
		// case HEARTBEAT: return nil
		case ERROR: return nil
		case PUBLISH: return nil
		case PUBLISHED: return nil
		case SUBSCRIBE: return nil
		case SUBSCRIBED: return nil
		case UNSUBSCRIBE: return nil
		case UNSUBSCRIBED: return nil
		case EVENT: return nil
		case CALL: return nil
		// case CANCEL: return nil
		case RESULT: return nil
		case REGISTER: return nil
		case REGISTERED: return nil
		case UNREGISTER: return nil
		case UNREGISTERED: return nil
		case INVOCATION: return nil
		// case INTERRUPT: return nil
		case YIELD: return nil
		default: return ErrBadMessageType
	}
}

// TODO: investigate whether we need to copy these before returning pointer
type Integer uint64
func (i *Integer) Value() interface{} { return *i }
func (i *Integer) Validate() error { return nil }

type String string
func (s *String) Value() interface{} { return *s }
func (s *String) Validate() error { return nil }

type Bool bool
func (b *Bool) Value() interface{} { return *b }
func (b *Bool) Validate() error { return nil }

type Id uint64
const MaxId Id = 9007199254740992
func (id *Id) Value() interface{} { return *id }
func (id *Id) Validate() error {
	if *id > MaxId {
		return ErrIdOverflow
	}
	return nil
}

type Uri string
var UriRegexp = regexp.MustCompile(`^([^\s\.#]+)(\.([^\s\.#]+))*$`)
func (u *Uri) Value() interface{} { return *u }
func (u *Uri) Validate() error {
	if !UriRegexp.MatchString(string(*u)) {
		return ErrBadUri
	}
	return nil
}

type Dict map[string]interface{}
func (d *Dict) Value() interface{} { return *d }
func (d *Dict) Validate() error {
	if d == nil {
		return ErrNilDict
	}
	return nil
}

type List []interface{}
func (l *List) Value() interface{} { return *l }
func (l *List) Validate() error {
	if l == nil {
		return ErrNilList
	}
	return nil
}

type ArgsKwargs struct {
	Args        List
	Kwargs      Dict
}

type HelloMessage struct {
	Realm   Uri
	Details Dict
}
func (h *HelloMessage) Type() MessageType { return HELLO }
func (h *HelloMessage) Parts() []MessagePart {
	return []MessagePart{HELLO, &h.Realm, &h.Details}
}

type WelcomeMessage struct {
	Session Id
	Details Dict
}
func (w *WelcomeMessage) Type() MessageType { return WELCOME }
func (w *WelcomeMessage) Parts() []MessagePart {
	return []MessagePart{WELCOME, &w.Session, &w.Details}
}

type AbortMessage struct {
	Details Dict
	Reason  Uri
}
func (a *AbortMessage) Type() MessageType { return ABORT }
func (a *AbortMessage) Parts() []MessagePart {
	return []MessagePart{ABORT, &a.Details, &a.Reason}
}

type GoodbyeMessage struct {
	Details Dict
	Reason  Uri
}
func (g *GoodbyeMessage) Type() MessageType { return GOODBYE }
func (g *GoodbyeMessage) Parts() []MessagePart {
	return []MessagePart{GOODBYE, &g.Details, &g.Reason}
}

type ErrorMessage struct {
	RequestType MessageType
	Request     Id
	Details     Dict
	Error       Uri
	ArgsKwargs
}
func (e *ErrorMessage) Type() MessageType { return ERROR }
func (e *ErrorMessage) Parts() []MessagePart {
	return []MessagePart{
		ERROR,
		&e.RequestType,
		&e.Request,
		&e.Details,
		&e.Error,
		&e.Args,
		&e.Kwargs,
	}
}
func (e *ErrorMessage) NumOptional() int { return 2 }

type PublishMessage struct {
	Request Id
	Options Dict
	Topic   Uri
	ArgsKwargs
}
func (p *PublishMessage) Type() MessageType { return PUBLISH }
func (p *PublishMessage) Parts() []MessagePart {
	return []MessagePart{
		PUBLISH,
		&p.Request,
		&p.Options,
		&p.Topic,
		&p.Args,
		&p.Kwargs,
	}
}
func (p *PublishMessage) NumOptional() int { return 2 }

type PublishedMessage struct {
	Request     Id
	Publication Id
}
func (p *PublishedMessage) Type() MessageType { return PUBLISHED }
func (p *PublishedMessage) Parts() []MessagePart {
	return []MessagePart{PUBLISHED, &p.Request, &p.Publication}
}

type SubscribeMessage struct {
	Request Id
	Options Dict
	Topic   Uri
}
func (s *SubscribeMessage) Type() MessageType { return SUBSCRIBE }
func (s *SubscribeMessage) Parts() []MessagePart {
	return []MessagePart{SUBSCRIBE, &s.Request, &s.Options, &s.Topic}
}

type SubscribedMessage struct {
	Request      Id
	Subscription Id
}
func (s *SubscribedMessage) Type() MessageType { return SUBSCRIBED }
func (s *SubscribedMessage) Parts() []MessagePart {
	return []MessagePart{SUBSCRIBED, &s.Request, &s.Subscription}
}

type UnsubscribeMessage struct {
	Request      Id
	Subscription Id
}
func (u *UnsubscribeMessage) Type() MessageType { return UNSUBSCRIBE }
func (u *UnsubscribeMessage) Parts() []MessagePart {
	return []MessagePart{UNSUBSCRIBE, &u.Request, &u.Subscription}
}

type UnsubscribedMessage struct {
	Request Id
}
func (u *UnsubscribedMessage) Type() MessageType { return UNSUBSCRIBED }
func (u *UnsubscribedMessage) Parts() []MessagePart {
	return []MessagePart{UNSUBSCRIBED, &u.Request}
}

type EventMessage struct {
	Subscription Id
	Publication  Id
	ArgsKwargs
}
func (e *EventMessage) Type() MessageType { return EVENT }
func (e *EventMessage) Parts() []MessagePart {
	return []MessagePart{
		EVENT,
		&e.Subscription,
		&e.Publication,
		&e.Args,
		&e.Kwargs,
	}
}
func (e *EventMessage) NumOptional() int { return 2 }

type CallMessage struct {
	Request   Id
	Options   Dict
	Procedure Uri
	ArgsKwargs
}
func (c *CallMessage) Type() MessageType { return CALL }
func (c *CallMessage) Parts() []MessagePart {
	return []MessagePart{
		CALL,
		&c.Request,
		&c.Options,
		&c.Procedure,
		&c.Args,
		&c.Kwargs,
	}
}
func (c *CallMessage) NumOptional() int { return 2 }

type ResultMessage struct {
	Request Id
	Details Dict
	ArgsKwargs
}
func (r *ResultMessage) Type() MessageType { return RESULT }
func (r *ResultMessage) Parts() []MessagePart {
	return []MessagePart{RESULT, &r.Request, &r.Details, &r.Args, &r.Kwargs}
}
func (r *ResultMessage) NumOptional() int { return 2 }

type RegisterMessage struct {
	Request   Id
	Options   Dict
	Procedure Uri
}
func (r *RegisterMessage) Type() MessageType { return REGISTER }
func (r *RegisterMessage) Parts() []MessagePart {
	return []MessagePart{REGISTER, &r.Request, &r.Options, &r.Procedure}
}

type RegisteredMessage struct {
	Request      Id
	Registration Id
}
func (r *RegisteredMessage) Type() MessageType { return REGISTERED }
func (r *RegisteredMessage) Parts() []MessagePart {
	return []MessagePart{REGISTERED, &r.Request, &r.Registration}
}

type UnregisterMessage struct {
	Request      Id
	Registration Id
}
func (u *UnregisterMessage) Type() MessageType { return UNREGISTER }
func (u *UnregisterMessage) Parts() []MessagePart {
	return []MessagePart{UNREGISTER, &u.Request, &u.Registration}
}

type UnregisteredMessage struct {
	Request Id
}
func (u *UnregisteredMessage) Type() MessageType { return UNREGISTERED }
func (u *UnregisteredMessage) Parts() []MessagePart {
	return []MessagePart{UNREGISTERED, &u.Request}
}

type InvocationMessage struct {
	Request      Id
	Registration Id
	Details      Dict
	ArgsKwargs
}
func (i *InvocationMessage) Type() MessageType { return INVOCATION }
func (i *InvocationMessage) Parts() []MessagePart {
	return []MessagePart{
		INVOCATION,
		&i.Request,
		&i.Registration,
		&i.Details,
		&i.Args,
		&i.Kwargs,
	}
}
func (i *InvocationMessage) NumOptional() int { return 2 }

type YieldMessage struct {
	Request Id
	Options Dict
	ArgsKwargs
}
func (y *YieldMessage) Type() MessageType { return YIELD }
func (y *YieldMessage) Parts() []MessagePart {
	return []MessagePart{YIELD, &y.Request, &y.Options, &y.Args, &y.Kwargs}
}
func (y *YieldMessage) NumOptional() int { return 2 }
