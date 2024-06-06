package domain

type Page struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  struct {
		HTML map[string]string `json:"storage"`
	} `json:"body"`
}
