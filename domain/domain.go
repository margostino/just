package domain

type JobPosition struct {
	Title    string `json:"title"`
	Link     string `json:"link"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Datetime string `json:"datetime"`
}
