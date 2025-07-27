package client

import (
	"context"

	"github.com/jetder-core/api"
)

type organizationClient struct {
	inv invoker
}

func (c organizationClient) Create(ctx context.Context, m *api.OrganizationCreate) (*api.Empty, error) {
	var res api.Empty
	err := c.inv.invoke(ctx, "organization.create", m, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (c organizationClient) Update(ctx context.Context, m *api.OrganizationUpdate) (*api.Empty, error) {
	var res api.Empty
	err := c.inv.invoke(ctx, "organization.update", m, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (c organizationClient) List(ctx context.Context, m *api.Empty) (*api.OrganizationListResult, error) {
	var res api.OrganizationListResult
	err := c.inv.invoke(ctx, "organization.list", m, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (c organizationClient) Get(ctx context.Context, m *api.OrganizationGet) (*api.OrganizationItem, error) {
	var res api.OrganizationItem
	err := c.inv.invoke(ctx, "organization.get", m, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (c organizationClient) Delete(ctx context.Context, m *api.OrganizationDelete) (*api.Empty, error) {
	var res api.Empty
	err := c.inv.invoke(ctx, "organization.delete", m, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
