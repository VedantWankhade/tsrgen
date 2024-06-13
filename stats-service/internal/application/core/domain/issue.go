package domain

type Issue struct {
	Id     string      `json:"id"`
	Key    string      `json:"key"`
	Fields IssueFields `json:"fields"`
}

type IssueFields struct {
	TestType string `json:"customfield_10034"`
}
