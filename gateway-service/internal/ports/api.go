package ports

type APIPort interface {
	ConfluenceClientPort
	JiraClientPort
	StatsClientPort
	DBClientPort
	Generate(atlassianInstance, atlassianUsername, atlassianToken, jql, title, parentId, spaceId string) (string, error)
}
