package client

import (
	"context"

	"github.com/jetder-core/api"
)

type secretClient struct {
	inv invoker
}

func (c secretClient) Create(ctx context.Context, m *api.SecretCreate) (*api.Empty, error) {
	var res api.Empty
	err := c.inv.invoke(ctx, "secret.create", m, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (c secretClient) Get(ctx context.Context, m *api.SecretGet) (*api.SecretItem, error) {
	var res api.SecretItem
	err := c.inv.invoke(ctx, "secret.get", m, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (c secretClient) List(ctx context.Context, m *api.SecretList) (*api.SecretListResult, error) {
	var res api.SecretListResult
	err := c.inv.invoke(ctx, "secret.list", m, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (c secretClient) Delete(ctx context.Context, m *api.SecretDelete) (*api.Empty, error) {
	var res api.Empty
	err := c.inv.invoke(ctx, "secret.delete", m, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
