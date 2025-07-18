package client

import (
	"context"

	"github.com/jetder-core/api"
)

type pullSecretClient struct {
	inv invoker
}

func (c pullSecretClient) Create(ctx context.Context, m *api.PullSecretCreate) (*api.Empty, error) {
	var res api.Empty
	err := c.inv.invoke(ctx, "pullsecret.create", m, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (c pullSecretClient) Get(ctx context.Context, m *api.PullSecretGet) (*api.PullSecretItem, error) {
	var res api.PullSecretItem
	err := c.inv.invoke(ctx, "pullsecret.get", m, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (c pullSecretClient) Delete(ctx context.Context, m *api.PullSecretDelete) (*api.Empty, error) {
	var res api.Empty
	err := c.inv.invoke(ctx, "pullsecret.delete", m, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (c pullSecretClient) List(ctx context.Context, m *api.PullSecretList) (*api.PullSecretListResult, error) {
	var res api.PullSecretListResult
	err := c.inv.invoke(ctx, "pullsecret.list", m, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
