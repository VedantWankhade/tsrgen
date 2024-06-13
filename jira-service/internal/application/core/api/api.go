package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	gearhttp "github.com/vedantwankhade/gopher-gear/net/http"
	"github.com/vedantwankhade/tsrgen/jira-service/internal/application/core/domain"
)

type application struct{}

func NewApplication() *application {
	return &application{}
}

type jsonRes struct {
	Issues     []*domain.Issue `json:"issues"`
	Total      int             `json:"total"`
	MaxResults int             `json:"maxResults"`
}

func (a *application) GetIssuesWithJQL(jql, jiraInstance, jiraUsername, jiraToken string) ([]*domain.Issue, error) {
	url := fmt.Sprintf("https://%s/rest/api/3/search", jiraInstance)
	headers := make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Authorization"] = "Basic " + gearhttp.GetBearerToken(jiraUsername, jiraToken)
	headers["Content-Type"] = "application/json"

	body := fmt.Sprintf(`{
        "expand": [
            "names"
          ],
          "fields": [
            "status",
            "customfield_10034"
          ],
          "fieldsByKeys": false,
          "jql": "%s",
          "maxResults": 15,
          "startAt": 0
    }`, jql)

	resBody, err := gearhttp.MakeRequest(http.MethodPost, url, headers, strings.NewReader(body), nil)
	if err != nil {
		return nil, fmt.Errorf("error getting issues: %w", err)
	}
	defer resBody.Close()
	var issuesRes jsonRes
	err = json.NewDecoder(resBody).Decode(&issuesRes)
	if err != nil {
		return nil, fmt.Errorf("error decoding response body: %w", err)
	}
	return issuesRes.Issues, nil
}

func (a *application) GetIssuesTested(issues []*domain.Issue) []*domain.Issue {
	var issuesFiltered []*domain.Issue
	for _, issue := range issues {
		if issue.Fields.TestType != "" {
			issuesFiltered = append(issuesFiltered, issue)
		}
	}
	return issuesFiltered
}

func (a *application) GetIssuesNotTested(issues []*domain.Issue) []*domain.Issue {
	var issuesFiltered []*domain.Issue
	for _, issue := range issues {
		if issue.Fields.TestType == "" {
			issuesFiltered = append(issuesFiltered, issue)
		}
	}
	return issuesFiltered
}
