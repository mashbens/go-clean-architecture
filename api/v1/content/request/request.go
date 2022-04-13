package request

type Request struct {
	ID    int    `json:"id" param:"id"`
	Title string `json:"title"`
}
