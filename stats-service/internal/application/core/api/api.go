package api

import (
	"bytes"
	"fmt"

	"github.com/vedantwankhade/tsrgen/stats-service/internal/application/core/domain"
)

type application struct{}

func NewApplication() *application {
	return &application{}
}

func (a *application) GetHTML(issues []*domain.Issue) (string, error) {
	var html bytes.Buffer
	html.WriteString("<table><tr>")
	for _, header := range []string{"ID", "Key", "Testable"} {
		html.WriteString(fmt.Sprintf("<th>%s</th>", header))
	}
	html.WriteString("</tr>")
	for _, issue := range issues {
		html.WriteString(fmt.Sprintf(`<tr><td>%s</td><td>%s</td><td>%s</td></tr>`, issue.Id, issue.Key, issue.Fields.TestType))
	}
	html.WriteString("</table>")
	return html.String(), nil
}
