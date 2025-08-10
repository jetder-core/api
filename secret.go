package api

import (
	"context"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/moonrhythm/validator"
)

type Secret interface {
	Create(ctx context.Context, m *SecretCreate) (*Empty, error)
	Get(ctx context.Context, m *SecretGet) (*SecretItem, error)
	List(ctx context.Context, m *SecretList) (*SecretListResult, error)
	Delete(ctx context.Context, m *SecretDelete) (*Empty, error)
}

type SecretCreate struct {
	Project string `json:"project" yaml:"project"`
	Name    string `json:"name" yaml:"name"`
	Value   string `json:"value" yaml:"value"`
}

func (m *SecretCreate) Valid() error {
	m.Name = strings.TrimSpace(m.Name)

	v := validator.New()

	v.Must(m.Project != "", "project required")
	v.Must(ReValidName.MatchString(m.Name), "name invalid "+ReValidNameStr)
	{
		cnt := utf8.RuneCountInString(m.Name)
		v.Mustf(cnt >= MinNameLength && cnt <= MaxNameLength, "name must have length between %d-%d characters", MinNameLength, MaxNameLength)
	}

	return WrapValidate(v)
}

type SecretGet struct {
	Project string `json:"project" yaml:"project"`
	Name    string `json:"name" yaml:"name"`
}

func (m *SecretGet) Valid() error {
	v := validator.New()

	v.Must(m.Project != "", "project required")
	v.Must(m.Name != "", "name required")

	return WrapValidate(v)
}

type SecretList struct {
	Project string `json:"project" yaml:"project"`
}

func (m *SecretList) Valid() error {
	v := validator.New()

	v.Must(m.Project != "", "project required")

	return WrapValidate(v)
}

type SecretListResult struct {
	Items []*SecretItem `json:"items" yaml:"items"`
}

type SecretItem struct {
	Name      string    `json:"name" yaml:"name"`
	Value     string    `json:"value" yaml:"value"`
	CreatedAt time.Time `json:"createdAt" yaml:"createdAt"`
	CreatedBy string    `json:"createdBy" yaml:"createdBy"`
}

type SecretDelete struct {
	Project string `json:"project" yaml:"project"`
	Name    string `json:"name" yaml:"name"`
}

func (m *SecretDelete) Valid() error {
	v := validator.New()

	v.Must(m.Project != "", "project required")
	v.Must(m.Name != "", "name required")

	return WrapValidate(v)
}
