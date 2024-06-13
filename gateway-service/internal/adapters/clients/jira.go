package clients

import (
	"context"

	"github.com/vedantwankhade/tsrgen/gateway-service/internal/application/core/domain"
	"github.com/vedantwankhade/tsrgen/gateway-service/services/jira"
	"google.golang.org/grpc"
)

type jiraClient struct {
	Client jira.JiraClient
}

func NewJiraClient(conn *grpc.ClientConn) *jiraClient {
	return &jiraClient{Client: jira.NewJiraClient(conn)}
}

func (c *jiraClient) GetIssues(jql, jiraInstance, username, token string) ([]*domain.Issue, error) {
	issuesReq := jira.IssueReq{
		Jql:          jql,
		JiraInstance: jiraInstance,
		JiraUsername: username,
		JiraToken:    token,
	}
	issuesRes, err := c.Client.GetIssues(context.Background(), &issuesReq)
	if err != nil {
		return nil, err
	}
	var issues []*domain.Issue
	for _, issue := range issuesRes.Issues {
		issues = append(issues, &domain.Issue{
			Id:  issue.Id,
			Key: issue.Key,
			Fields: domain.IssueFields{
				TestType: issue.TestType,
			},
		})
	}
	return issues, nil
}
