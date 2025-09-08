package api

import (
	"context"
	"encoding/json"
	"time"

	"github.com/moonrhythm/validator"
)

type ActivityLog interface {
	User(ctx context.Context, m *ActivityLogUserList) (*ActivityLogResult, error)
	Project(ctx context.Context, m *ActivityLogProjectList) (*ActivityLogResult, error)
	Organization(ctx context.Context, m *ActivityLogOrganizationList) (*ActivityLogResult, error)
}

const (
	MaxActivityLogLimit = 1000
)

type ActivityLogUserList struct {
	Limit  int64 `json:"limit" yaml:"limit"`
	Offset int64 `json:"offset" yaml:"offset"`
}

func (m *ActivityLogUserList) Valid() error {
	v := validator.New()

	v.Must(m.Limit > 0, "limit required")
	v.Must(m.Limit <= MaxActivityLogLimit, "limit must be less than or equal to "+string(MaxActivityLogLimit))
	v.Must(m.Offset >= 0, "offset must be greater than or equal to 0")

	return WrapValidate(v)
}

type ActivityLogProjectList struct {
	Project string `json:"project" yaml:"project"`
	Limit   int64  `json:"limit" yaml:"limit"`
	Offset  int64  `json:"offset" yaml:"offset"`
}

func (m *ActivityLogProjectList) Valid() error {
	v := validator.New()

	v.Must(m.Limit > 0, "limit required")
	v.Must(m.Limit <= MaxActivityLogLimit, "limit must be less than or equal to "+string(MaxActivityLogLimit))
	v.Must(m.Offset >= 0, "offset must be greater than or equal to 0")
	v.Must(m.Project != "", "project required")

	return WrapValidate(v)
}

type ActivityLogOrganizationList struct {
	OrgID  int64 `json:"orgId,string" yaml:"orgId"`
	Limit  int64 `json:"limit" yaml:"limit"`
	Offset int64 `json:"offset" yaml:"offset"`
}

func (m *ActivityLogOrganizationList) Valid() error {
	v := validator.New()

	v.Must(m.Limit > 0, "limit required")
	v.Must(m.Limit <= MaxActivityLogLimit, "limit must be less than or equal to "+string(MaxActivityLogLimit))
	v.Must(m.Offset >= 0, "offset must be greater than or equal to 0")
	v.Must(m.OrgID > 0, "orgId required")

	return WrapValidate(v)
}

type ActivityLogResult struct {
	Items []*ActivityLogItem `json:"items" yaml:"items"`
}

type ActivityLogItem struct {
	ID        int64           `json:"id,string" yaml:"id"`
	Email     string          `json:"email" yaml:"email"`
	Action    string          `json:"action" yaml:"action"`
	Params    json.RawMessage `json:"params" yaml:"params"`
	CreatedAt time.Time       `json:"createdAt" yaml:"createdAt"`
}
