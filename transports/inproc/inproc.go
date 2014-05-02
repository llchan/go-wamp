package inproc

import (
	"github.com/llchan/go-wamp"
)

type Transport struct {
	sendbuf chan wamp.Message
	recvbuf chan wamp.Message
}

func NewTransportPair(capacity int) (*Transport, *Transport) {
	if capacity <= 0 {
		capacity = 1
	}
	buf1 := make(chan wamp.Message, capacity)
	buf2 := make(chan wamp.Message, capacity)
	t1 := &Transport{
		sendbuf: buf1,
		recvbuf: buf2,
	}
	t2 := &Transport{
		sendbuf: buf2,
		recvbuf: buf1,
	}
	return t1, t2
}

func (t *Transport) Close() error {
	close(t.sendbuf)
	return nil
}

func (t *Transport) Send(m wamp.Message) error {
	t.sendbuf <- m
	return nil
}

func (t *Transport) Recv() (wamp.Message, error) {
	return <-t.recvbuf, nil
}
