package ports

import "github.com/vedantwankhade/tsrgen/jira-service/internal/application/core/domain"

type APIPort interface {
	GetIssuesWithJQL(jql, jiraInstance, jiraUsername, jiraToken string) ([]*domain.Issue, error)
}
