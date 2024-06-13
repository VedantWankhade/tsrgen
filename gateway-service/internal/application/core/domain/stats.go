package domain

type StatReq struct {
	Issues []*Issue `json:"issues"`
}
