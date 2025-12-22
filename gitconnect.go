package api

import (
	"context"
	"time"
)

type GitConnect interface {
	Bind(ctx context.Context, m *GitConnectBind) (*GitConnectBindResult, error)
	BindStatus(ctx context.Context) (*GitConnectBindStatus, error)
	GetRepos(ctx context.Context, m *GitConnectGetRepos) (*GitConnectGetReposResult, error)
	Create(ctx context.Context, m *GitConnectCreate) (*Empty, error)
	Update(ctx context.Context, m *GitConnectUpdate) (*Empty, error)
	List(ctx context.Context, m *GitConnectList) (*GitConnectListResult, error)
	Get(ctx context.Context, m *GitConnectGet) (*GitConnectItem, error)
	Delete(ctx context.Context, m *GitConnectDelete) (*Empty, error)
	Builds(ctx context.Context, m *GitConnectBuilds) (*GitConnectBuildsResult, error)
	Build(ctx context.Context, m *GitConnectBuild) (*GitConnectItem, error)
	BuildLogs(ctx context.Context, m *GitConnectBuildLogs) (*GitConnectBuildLogsResult, error)
}

type GitConnectBind struct {
	Provider string `json:"provider" yaml:"provider"`
	Token    string `json:"token" yaml:"token"`
}

type GitConnectBindResult struct {
	Name string `json:"name" yaml:"name"`
}

type GitConnectBindStatus struct {
	Github    *GitConnectStatusProvider `json:"github" yaml:"github"`
	Gitlab    *GitConnectStatusProvider `json:"gitlab" yaml:"gitlab"`
	Bitbucket *GitConnectStatusProvider `json:"bitbucket" yaml:"bitbucket"`
}

type GitConnectGetRepos struct {
	Provider string `json:"provider" yaml:"provider"`
}

type GitConnectGetReposResult struct {
	Items []*GitConnectRepo `json:"items" yaml:"items"`
}

type GitConnectRepo struct {
	Name string `json:"name" yaml:"name"`
}

type GitConnectStatusProvider struct {
	Name string `json:"name" yaml:"name"`
}

type GitConnectCreate struct {
	Project          string              `json:"project" yaml:"project"`
	Location         string              `json:"location" yaml:"location"`
	Name             string              `json:"name" yaml:"name"`
	Provider         string              `json:"provider" yaml:"provider"`
	Origin           string              `json:"origin" yaml:"origin"`
	ProductionBranch string              `json:"productionBranch" yaml:"productionBranch"`
	PreviewBranch    string              `json:"previewBranch" yaml:"previewBranch"`
	Type             DeploymentType      `json:"type" yaml:"type"`
	Env              map[string]string   `json:"env" yaml:"env"`
	EnvGroups        []string            `json:"envGroups" yaml:"envGroups"`
	Resources        *DeploymentResource `json:"resources" yaml:"resources"`
}

type GitConnectUpdate struct {
	Project          string              `json:"project" yaml:"project"`
	Name             string              `json:"name" yaml:"name"`
	ProductionBranch string              `json:"productionBranch" yaml:"productionBranch"`
	PreviewBranch    string              `json:"previewBranch" yaml:"previewBranch"`
	Env              map[string]string   `json:"env" yaml:"env"`
	EnvGroups        []string            `json:"envGroups" yaml:"envGroups"`
	Resources        *DeploymentResource `json:"resources" yaml:"resources"`
}

type GitConnectList struct {
	Project string `json:"project" yaml:"project"`
}

type GitConnectItem struct {
	Location         string              `json:"location" yaml:"location"`
	Name             string              `json:"name" yaml:"name"`
	Origin           string              `json:"origin" yaml:"origin"`
	ProductionBranch string              `json:"productionBranch" yaml:"productionBranch"`
	PreviewBranch    string              `json:"previewBranch" yaml:"previewBranch"`
	Type             DeploymentType      `json:"type" yaml:"type"`
	Env              map[string]string   `json:"env" yaml:"env"`
	EnvGroups        []string            `json:"envGroups" yaml:"envGroups"`
	Resources        *DeploymentResource `json:"resources" yaml:"resources"`
	CreatedAt        time.Time           `json:"createdAt" yaml:"createdAt"`
}

type GitConnectListResult struct {
	Items []*GitConnectItem `json:"items" yaml:"items"`
}

type GitConnectGet struct {
	Project string `json:"project" yaml:"project"`
	Name    string `json:"name" yaml:"name"`
}

type GitConnectDelete struct {
	Project string `json:"project" yaml:"project"`
	Name    string `json:"name" yaml:"name"`
}

type GitConnectBuilds struct {
	Project string `json:"project" yaml:"project"`
	Name    string `json:"name" yaml:"name"`
}

type GitConnectBuildsResult struct {
	Items []*GitConnectBuildItem `json:"items" yaml:"items"`
}

type GitConnectBuildItem struct {
	ID              int64     `json:"id,string" yaml:"id"`
	Commit          string    `json:"commit" yaml:"commit"`
	Branch          string    `json:"branch" yaml:"branch"`
	Production      bool      `json:"production" yaml:"production"`
	CreatedAt       time.Time `json:"createdAt" yaml:"createdAt"`
	BuildStartedAt  time.Time `json:"buildStartedAt" yaml:"buildStartedAt"`
	BuildFinishedAt time.Time `json:"buildFinishedAt" yaml:"buildFinishedAt"`
}

type GitConnectBuild struct {
	Project string `json:"project" yaml:"project"`
	Name    string `json:"name" yaml:"name"`
	BuildID int64  `json:"buildId,string" yaml:"buildId"`
}

type GitConnectBuildLogs struct {
	Project string `json:"project" yaml:"project"`
	Name    string `json:"name" yaml:"name"`
	BuildID int64  `json:"buildId,string" yaml:"buildId"`
}

type GitConnectBuildLogsResult struct {
	Text string `json:"text" yaml:"text"`
}
