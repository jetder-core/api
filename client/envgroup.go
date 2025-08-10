package client

import (
	"context"

	"github.com/jetder-core/api"
)

type envGroupClient struct {
	inv invoker
}

func (c envGroupClient) Create(ctx context.Context, m *api.EnvGroupCreate) (*api.Empty, error) {
	var res api.Empty
	err := c.inv.invoke(ctx, "envgroup.create", m, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (c envGroupClient) Get(ctx context.Context, m *api.EnvGroupGet) (*api.EnvGroupItem, error) {
	var res api.EnvGroupItem
	err := c.inv.invoke(ctx, "envgroup.get", m, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (c envGroupClient) List(ctx context.Context, m *api.EnvGroupList) (*api.EnvGroupListResult, error) {
	var res api.EnvGroupListResult
	err := c.inv.invoke(ctx, "envgroup.list", m, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (c envGroupClient) Delete(ctx context.Context, m *api.EnvGroupDelete) (*api.Empty, error) {
	var res api.Empty
	err := c.inv.invoke(ctx, "envgroup.delete", m, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
