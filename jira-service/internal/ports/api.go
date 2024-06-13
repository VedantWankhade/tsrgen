package ports

import "github.com/vedantwankhade/tsrgen/jira-service/internal/application/core/domain"

type APIPort interface {
	GetIssuesWithJQL(jql, jiraInstance, jiraUsername, jiraToken string) ([]*domain.Issue, error)
	GetIssuesTested(issues []*domain.Issue) []*domain.Issue
	GetIssuesNotTested(issues []*domain.Issue) []*domain.Issue
}
