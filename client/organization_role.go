package client

import (
	"context"

	"github.com/jetder-core/api"
)

type organizationRoleClient struct {
	inv invoker
}

func (c organizationRoleClient) Create(ctx context.Context, m *api.OrganizationRoleCreate) (*api.Empty, error) {
	var res api.Empty
	err := c.inv.invoke(ctx, "organizationRole.create", m, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (c organizationRoleClient) Get(ctx context.Context, m *api.OrganizationRoleGet) (*api.OrganizationRoleGetResult, error) {
	var res api.OrganizationRoleGetResult
	err := c.inv.invoke(ctx, "organizationRole.get", m, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (c organizationRoleClient) List(ctx context.Context, m *api.OrganizationRoleList) (*api.OrganizationRoleListResult, error) {
	var res api.OrganizationRoleListResult
	err := c.inv.invoke(ctx, "organizationRole.list", m, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (c organizationRoleClient) Delete(ctx context.Context, m *api.OrganizationRoleDelete) (*api.Empty, error) {
	var res api.Empty
	err := c.inv.invoke(ctx, "organizationRole.delete", m, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (c organizationRoleClient) Grant(ctx context.Context, m *api.OrganizationRoleGrant) (*api.Empty, error) {
	var res api.Empty
	err := c.inv.invoke(ctx, "organizationRole.grant", m, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (c organizationRoleClient) Revoke(ctx context.Context, m *api.OrganizationRoleRevoke) (*api.Empty, error) {
	var res api.Empty
	err := c.inv.invoke(ctx, "organizationRole.revoke", m, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (c organizationRoleClient) Users(ctx context.Context, m *api.OrganizationRoleUsers) (*api.OrganizationRoleUsersResult, error) {
	var res api.OrganizationRoleUsersResult
	err := c.inv.invoke(ctx, "organizationRole.users", m, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (c organizationRoleClient) Bind(ctx context.Context, m *api.OrganizationRoleBind) (*api.Empty, error) {
	var res api.Empty
	err := c.inv.invoke(ctx, "organizationRole.bind", m, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (c organizationRoleClient) Permissions(ctx context.Context, m *api.Empty) ([]string, error) {
	var res []string
	err := c.inv.invoke(ctx, "organizationRole.permissions", m, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
