package ports

type APIPort interface {
	ConfluenceClientPort
	JiraClientPort
	StatsClientPort
}
