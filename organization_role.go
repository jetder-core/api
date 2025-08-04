package api

import (
	"context"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/asaskevich/govalidator"
	"github.com/moonrhythm/validator"
)

var orgPermissions = []string{
	"*",
	"role.*",
	"role.create",
	"role.list",
	"role.get",
	"role.delete",
	"role.bind",
}

func OrganizationPermissions() []string {
	xs := make([]string, len(orgPermissions))
	copy(xs, orgPermissions)
	return xs
}

type OrganizationRole interface {
	Create(ctx context.Context, m *OrganizationRoleCreate) (*Empty, error)
	List(ctx context.Context, m *OrganizationRoleList) (*OrganizationRoleListResult, error)
	Get(ctx context.Context, m *OrganizationRoleGet) (*OrganizationRoleGetResult, error)
	Delete(ctx context.Context, m *OrganizationRoleDelete) (*Empty, error)
	Grant(ctx context.Context, m *OrganizationRoleGrant) (*Empty, error)
	Revoke(ctx context.Context, m *OrganizationRoleRevoke) (*Empty, error)
	Users(ctx context.Context, m *OrganizationRoleUsers) (*OrganizationRoleUsersResult, error)
	Bind(ctx context.Context, m *OrganizationRoleBind) (*Empty, error)
	Permissions(ctx context.Context, _ *Empty) ([]string, error)
}

type OrganizationRoleCreate struct {
	Organization string   `json:"organization" yaml:"organization"` // organization id
	Role         string   `json:"role" yaml:"role"`                 // role sid
	Name         string   `json:"name" yaml:"name"`                 // role name (free text)
	Permissions  []string `json:"permissions" yaml:"permissions"`
}

func (m *OrganizationRoleCreate) Valid() error {
	m.Role = strings.TrimSpace(m.Role)
	m.Name = strings.TrimSpace(m.Name)

	v := validator.New()

	v.Must(m.Organization != "", "organization required")
	v.Must(m.Role != "owner", "not allow to edit owner role")
	v.Must(ReValidSID.MatchString(m.Role), "role invalid")
	{
		cnt := utf8.RuneCountInString(m.Role)
		v.Must(cnt >= 3 && cnt <= 20, "role must have length between 3-20 characters")
	}
	v.Must(utf8.ValidString(m.Name), "name invalid")
	{
		cnt := utf8.RuneCountInString(m.Name)
		v.Must(cnt >= 3 && cnt <= 64, "name must have length between 3-64 characters")
	}

	return WrapValidate(v)
}

type OrganizationRoleGet struct {
	Organization string `json:"organization" yaml:"organization"` // organization id
	Role         string `json:"role" yaml:"role"`                 // role sid
}

type OrganizationRoleGetResult struct {
	Role         string    `json:"role" yaml:"role"`                 // role sid
	Organization string    `json:"organization" yaml:"organization"` // organization id
	Name         string    `json:"name" yaml:"name"`                 // role name
	Permissions  []string  `json:"permissions" yaml:"permissions"`
	CreatedAt    time.Time `json:"createdAt" yaml:"createdAt"`
}

func (m *OrganizationRoleGetResult) Table() [][]string {
	table := [][]string{
		{"ROLE", "NAME", "AGE"},
		{
			m.Role,
			m.Name,
			age(m.CreatedAt),
		},
	}
	return table
}

type OrganizationRoleList struct {
	Organization string `json:"organization" yaml:"organization"` // organization id
}

type OrganizationRoleListResult struct {
	Organization string          `json:"organization" yaml:"organization"` // organization id
	Items        []*RoleListItem `json:"items" yaml:"items"`
}

func (m *OrganizationRoleListResult) Table() [][]string {
	table := [][]string{
		{"ROLE", "NAME", "AGE"},
	}
	for _, x := range m.Items {
		table = append(table, []string{
			x.Role,
			x.Name,
			age(x.CreatedAt),
		})
	}
	return table
}

type OrganizationRoleListItem struct {
	Role        string    `json:"role" yaml:"role"` // role sid
	Name        string    `json:"name" yaml:"name"` // role name
	Permissions []string  `json:"permissions" yaml:"permissions"`
	CreatedAt   time.Time `json:"createdAt" yaml:"createdAt"`
	CreatedBy   string    `json:"createdBy" yaml:"createdBy"`
}

type OrganizationRoleDelete struct {
	Organization string `json:"organization" yaml:"organization"` // organization id
	Role         string `json:"role" yaml:"role"`
}

func (m *OrganizationRoleDelete) Valid() error {
	v := validator.New()

	v.Must(m.Organization != "", "organization required")
	v.Must(m.Role != "", "role required")

	return WrapValidate(v)
}

type OrganizationRoleGrant struct {
	Organization string `json:"organization" yaml:"organization"` // organization id
	Role         string `json:"role" yaml:"role"`                 // role sid
	Email        string `json:"email" yaml:"email"`               // user email
}

func (m *OrganizationRoleGrant) Valid() error {
	m.Email = strings.TrimSpace(m.Email)

	v := validator.New()

	v.Must(m.Organization != "", "organization required")
	v.Must(ReValidSID.MatchString(m.Role), "role invalid")
	cnt := utf8.RuneCountInString(m.Role)
	v.Must(cnt >= 6 && cnt <= 20, "role must have length between 6-20 characters")
	v.Must(m.Email != "", "email required")
	v.Must(govalidator.IsEmail(m.Email), "email invalid")

	return WrapValidate(v)
}

type OrganizationRoleRevoke struct {
	Organization string `json:"organization" yaml:"organization"` // organization id
	Role         string `json:"role" yaml:"role"`                 // role sid
	Email        string `json:"email" yaml:"email"`               // user email
}

func (m *OrganizationRoleRevoke) Valid() error {
	m.Email = strings.TrimSpace(m.Email)

	v := validator.New()

	v.Must(m.Organization != "", "organization required")
	v.Must(ReValidSID.MatchString(m.Role), "role invalid")
	cnt := utf8.RuneCountInString(m.Role)
	v.Must(cnt >= 6 && cnt <= 20, "role must have length between 6-20 characters")
	v.Must(m.Email != "", "email required")
	v.Must(govalidator.IsEmail(m.Email), "email invalid")

	return WrapValidate(v)
}

type OrganizationRoleUsers struct {
	Organization string `json:"organization" yaml:"organization"` // organization id
}

func (m *OrganizationRoleUsers) Valid() error {
	v := validator.New()

	v.Must(m.Organization != "", "organization required")

	return WrapValidate(v)
}

type OrganizationRoleUsersResult struct {
	Organization string           `json:"organization" yaml:"organization"` // organization id
	Items        []*RoleUsersItem `json:"items" yaml:"items"`
	Users        []*RoleUsersItem `json:"users" yaml:"users"`
}

func (m *OrganizationRoleUsersResult) Table() [][]string {
	table := [][]string{
		{"EMAIL", "ROLE"},
	}
	for _, u := range m.Items {
		for _, r := range u.Roles {
			table = append(table, []string{
				u.Email,
				r,
			})
		}
	}
	return table
}

type OrganizationRoleUsersItem struct {
	Email string   `json:"email" yaml:"email"`
	Roles []string `json:"roles" yaml:"roles"`
}

type OrganizationRoleBind struct {
	Organization string   `json:"organization" yaml:"organization"` // organization id
	Email        string   `json:"email" yaml:"email"`
	Roles        []string `json:"roles" yaml:"roles"`
}
