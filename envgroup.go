package api

import (
	"context"
	"unicode/utf8"

	"github.com/moonrhythm/validator"
)

type EnvGroup interface {
	Create(ctx context.Context, m *EnvGroupCreate) (*Empty, error)
	Get(ctx context.Context, m *EnvGroupGet) (*EnvGroupItem, error)
	List(ctx context.Context, m *EnvGroupList) (*EnvGroupListResult, error)
	Delete(ctx context.Context, m *EnvGroupDelete) (*Empty, error)
}

type EnvGroupCreate struct {
	Project string            `json:"project" yaml:"project"`
	Name    string            `json:"name" yaml:"name"`
	Env     map[string]string `json:"env" yaml:"env"`
}

func (m *EnvGroupCreate) Valid() error {
	v := validator.New()

	v.Must(m.Project != "", "project required")
	v.Must(ReValidName.MatchString(m.Name), "name invalid "+ReValidNameStr)
	{
		cnt := utf8.RuneCountInString(m.Name)
		v.Mustf(cnt >= MinNameLength && cnt <= MaxNameLength, "name must have length between %d-%d characters", MinNameLength, MaxNameLength)
	}

	return WrapValidate(v)
}

type EnvGroupItem struct {
	Name string            `json:"name" yaml:"name"`
	Env  map[string]string `json:"env" yaml:"env"`
}

type EnvGroupGet struct {
	Project string `json:"project" yaml:"project"`
	Name    string `json:"name" yaml:"name"`
}

func (m *EnvGroupGet) Valid() error {
	v := validator.New()
	v.Must(m.Project != "", "project required")
	v.Must(ReValidName.MatchString(m.Name), "name invalid "+ReValidNameStr)
	return WrapValidate(v)
}

type EnvGroupListResult struct {
	Items []*EnvGroupItem `json:"items" yaml:"items"`
}

type EnvGroupList struct {
	Project string `json:"project" yaml:"project"`
}

func (m *EnvGroupList) Valid() error {
	v := validator.New()
	v.Must(m.Project != "", "project required")
	return WrapValidate(v)
}

type EnvGroupDelete struct {
	Project string `json:"project" yaml:"project"`
	Name    string `json:"name" yaml:"name"`
}

func (m *EnvGroupDelete) Valid() error {
	v := validator.New()
	v.Must(m.Project != "", "project required")
	return WrapValidate(v)
}
