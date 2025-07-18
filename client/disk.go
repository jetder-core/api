package client

import (
	"context"

	"github.com/jetder-core/api"
)

type diskClient struct {
	inv invoker
}

func (c diskClient) Create(ctx context.Context, m *api.DiskCreate) (*api.Empty, error) {
	var res api.Empty
	err := c.inv.invoke(ctx, "disk.create", m, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (c diskClient) Get(ctx context.Context, m *api.DiskGet) (*api.DiskItem, error) {
	var res api.DiskItem
	err := c.inv.invoke(ctx, "disk.get", m, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (c diskClient) Update(ctx context.Context, m *api.DiskUpdate) (*api.Empty, error) {
	var res api.Empty
	err := c.inv.invoke(ctx, "disk.update", m, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (c diskClient) List(ctx context.Context, m *api.DiskList) (*api.DiskListResult, error) {
	var res api.DiskListResult
	err := c.inv.invoke(ctx, "disk.list", m, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (c diskClient) Delete(ctx context.Context, m *api.DiskDelete) (*api.Empty, error) {
	var res api.Empty
	err := c.inv.invoke(ctx, "disk.delete", m, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (c diskClient) Metrics(ctx context.Context, m *api.DiskMetrics) (*api.DiskMetricsResult, error) {
	var res api.DiskMetricsResult
	err := c.inv.invoke(ctx, "disk.metrics", m, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
