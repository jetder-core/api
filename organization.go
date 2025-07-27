package api

import (
	"context"
	"strings"

	"github.com/moonrhythm/validator"
)

type Organization interface {
	Create(ctx context.Context, m *OrganizationCreate) (*Empty, error)
	List(ctx context.Context, _ *Empty) (*OrganizationListResult, error)
	Get(ctx context.Context, m *OrganizationGet) (*OrganizationItem, error)
	Update(ctx context.Context, m *OrganizationUpdate) (*Empty, error)
	Delete(ctx context.Context, m *OrganizationDelete) (*Empty, error)
}

type OrganizationItem struct {
	ID   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
}

type OrganizationListResult struct {
	Items []*OrganizationItem `json:"items" yaml:"items"`
}

type OrganizationGet struct {
	ID string `json:"id" yaml:"id"`
}

func (m *OrganizationGet) Valid() error {
	m.ID = strings.TrimSpace(m.ID)
	m.ID = strings.ToLower(m.ID)

	v := validator.New()
	v.Must(m.ID != "", "id required")

	return WrapValidate(v)
}

type OrganizationCreate struct {
	ID   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
}

func (m *OrganizationCreate) Valid() error {
	m.ID = strings.TrimSpace(m.ID)
	m.ID = strings.ToLower(m.ID)
	m.Name = strings.TrimSpace(m.Name)

	v := validator.New()
	v.Must(m.ID != "", "id required")
	v.Must(m.Name != "", "name required")

	return WrapValidate(v)
}

type OrganizationUpdate struct {
	ID   string `json:"id" yaml:"id"`
	Name string `json:"name" yaml:"name"`
}

func (m *OrganizationUpdate) Valid() error {
	m.ID = strings.TrimSpace(m.ID)
	m.ID = strings.ToLower(m.ID)
	m.Name = strings.TrimSpace(m.Name)

	v := validator.New()
	v.Must(m.ID != "", "id required")
	v.Must(m.Name != "", "name required")

	return WrapValidate(v)
}

type OrganizationDelete struct {
	ID string `json:"id" yaml:"id"`
}

func (m *OrganizationDelete) Valid() error {
	m.ID = strings.TrimSpace(m.ID)
	m.ID = strings.ToLower(m.ID)

	v := validator.New()
	v.Must(m.ID != "", "id required")

	return WrapValidate(v)
}
