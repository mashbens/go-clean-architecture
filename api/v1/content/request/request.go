package request

type CreateContentRequest struct {
	ID          int    `json:"id" param:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
