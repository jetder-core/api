package client

import (
	"context"

	"github.com/jetder-core/api"
)

type roleClient struct {
	inv invoker
}

func (c roleClient) Create(ctx context.Context, m *api.RoleCreate) (*api.Empty, error) {
	var res api.Empty
	err := c.inv.invoke(ctx, "role.create", m, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (c roleClient) Get(ctx context.Context, m *api.RoleGet) (*api.RoleGetResult, error) {
	var res api.RoleGetResult
	err := c.inv.invoke(ctx, "role.get", m, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (c roleClient) List(ctx context.Context, m *api.RoleList) (*api.RoleListResult, error) {
	var res api.RoleListResult
	err := c.inv.invoke(ctx, "role.list", m, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (c roleClient) Delete(ctx context.Context, m *api.RoleDelete) (*api.Empty, error) {
	var res api.Empty
	err := c.inv.invoke(ctx, "role.delete", m, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (c roleClient) Grant(ctx context.Context, m *api.RoleGrant) (*api.Empty, error) {
	var res api.Empty
	err := c.inv.invoke(ctx, "role.grant", m, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (c roleClient) Revoke(ctx context.Context, m *api.RoleRevoke) (*api.Empty, error) {
	var res api.Empty
	err := c.inv.invoke(ctx, "role.revoke", m, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (c roleClient) Users(ctx context.Context, m *api.RoleUsers) (*api.RoleUsersResult, error) {
	var res api.RoleUsersResult
	err := c.inv.invoke(ctx, "role.users", m, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (c roleClient) Bind(ctx context.Context, m *api.RoleBind) (*api.Empty, error) {
	var res api.Empty
	err := c.inv.invoke(ctx, "role.bind", m, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (c roleClient) Permissions(ctx context.Context, m *api.Empty) ([]string, error) {
	var res []string
	err := c.inv.invoke(ctx, "role.permissions", m, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
