package rpc

import (
	"context"

	"opencsg.com/csghub-server/api/httpbase"
	"opencsg.com/csghub-server/common/types"
)

type ModerationSvcClient interface {
	PassTextCheck(ctx context.Context, scenario, text string) (*CheckResult, error)
	PassImageCheck(ctx context.Context, scenario, ossBucketName, ossObjectName string) (*CheckResult, error)
	PassStreamCheck(ctx context.Context, scenario, text, id string) (*CheckResult, error)
	SubmitRepoCheck(ctx context.Context, repoType types.RepositoryType, namespace, name string) error
}

type CheckResult struct {
	IsSensitive bool   `json:"is_sensitive"`
	Reason      string `json:"reason"`
}

type ModerationSvcHttpClient struct {
	hc *HttpClient
}

func NewModerationSvcHttpClient(endpoint string, opts ...RequestOption) ModerationSvcClient {
	return &ModerationSvcHttpClient{
		hc: NewHttpClient(endpoint, opts...),
	}
}

func (c *ModerationSvcHttpClient) PassTextCheck(ctx context.Context, scenario, text string) (*CheckResult, error) {
	type CheckRequest struct {
		Scenario string `json:"scenario"`
		Text     string `json:"text"`
	}

	req := &CheckRequest{
		Scenario: scenario,
		Text:     text,
	}
	const path = "/api/v1/text"
	var resp httpbase.R
	resp.Data = &CheckResult{}
	err := c.hc.Post(ctx, path, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Data.(*CheckResult), nil
}

func (c *ModerationSvcHttpClient) PassStreamCheck(ctx context.Context, scenario, text, id string) (*CheckResult, error) {
	type ServiceParameters struct {
		Content   string `json:"content"`
		SessionId string `json:"sessionId"`
	}
	type CheckRequest struct {
		Service           string            `json:"Service"`
		ServiceParameters ServiceParameters `json:"ServiceParameters"`
	}
	req := &CheckRequest{
		Service: scenario,
		ServiceParameters: ServiceParameters{
			Content:   text,
			SessionId: id,
		},
	}
	const path = "/api/v1/stream"
	var resp httpbase.R
	resp.Data = &CheckResult{}
	err := c.hc.Post(ctx, path, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Data.(*CheckResult), nil
}

func (c *ModerationSvcHttpClient) PassImageCheck(ctx context.Context, scenario, ossBucketName, ossObjectName string) (*CheckResult, error) {
	type CheckRequest struct {
		Scenario      string `json:"scenario"`
		OssBucketName string `json:"oss_bucket_name"`
		OssObjectName string `json:"oss_object_name"`
	}

	req := &CheckRequest{
		Scenario:      scenario,
		OssBucketName: ossBucketName,
		OssObjectName: ossObjectName,
	}
	var resp httpbase.R
	resp.Data = &CheckResult{}
	const path = "/api/v1/image"
	err := c.hc.Post(ctx, path, req, &resp)
	if err != nil {
		return nil, err
	}
	return resp.Data.(*CheckResult), nil
}

func (c *ModerationSvcHttpClient) SubmitRepoCheck(ctx context.Context, repoType types.RepositoryType, namespace, name string) error {
	type CheckRequest struct {
		RepoType  types.RepositoryType `json:"repo_type"`
		Namespace string               `json:"namespace"`
		Name      string               `json:"name"`
	}

	req := &CheckRequest{
		RepoType:  repoType,
		Namespace: namespace,
		Name:      name,
	}
	const path = "/api/v1/repo"
	var resp httpbase.R
	return c.hc.Post(ctx, path, req, &resp)
}
