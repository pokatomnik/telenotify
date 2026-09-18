package notify

import (
	"context"

	"github.com/pokatomnik/telenotify/internal/entities/notifreq"
)

type Transport interface {
	SendMessage(ctx context.Context, msg notifreq.NotificationRequest) error
}

type Notify struct {
	transport Transport
}

func New(transport Transport) Notify {
	return Notify{
		transport: transport,
	}
}

func (n Notify) Notify(ctx context.Context, msg notifreq.NotificationRequest) error {
	return n.transport.SendMessage(ctx, msg)
}
