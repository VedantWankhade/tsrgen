package clients

import (
	"context"

	"github.com/vedantwankhade/tsrgen/gateway-service/internal/application/core/domain"
	"github.com/vedantwankhade/tsrgen/gateway-service/services/jira"
	"github.com/vedantwankhade/tsrgen/gateway-service/services/stats"
	"google.golang.org/grpc"
)

type statsClient struct {
	Client stats.StatsClient
}

func NewStatsClient(conn *grpc.ClientConn) *statsClient {
	return &statsClient{Client: stats.NewStatsClient(conn)}
}

func (c *statsClient) GetHTML(issues []*domain.Issue) (string, error) {
	var statsIssues []*jira.Issue
	for _, issue := range issues {
		statsIssues = append(statsIssues, &jira.Issue{
			Id:       issue.Id,
			Key:      issue.Key,
			TestType: issue.Fields.TestType,
		})
	}
	statsReq := &stats.StatsReq{
		Issues: statsIssues,
	}
	html, err := c.Client.GetHTML(context.Background(), statsReq)
	if err != nil {
		return "", err
	}
	return html.Html, nil
}
