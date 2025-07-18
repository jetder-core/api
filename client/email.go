package client

import (
	"context"

	"github.com/jetder-core/api"
)

type emailClient struct {
	inv invoker
}

func (c emailClient) Send(ctx context.Context, m *api.EmailSend) (*api.Empty, error) {
	var res api.Empty
	err := c.inv.invoke(ctx, "email.send", m, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (c emailClient) List(ctx context.Context, m *api.EmailList) (*api.EmailListResult, error) {
	var res api.EmailListResult
	err := c.inv.invoke(ctx, "email.list", m, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
