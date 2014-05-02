package wamp

import (
    "io"
)

type Transport interface {
    io.Closer
    Send(Message) error
    Recv() (Message, error)
}
