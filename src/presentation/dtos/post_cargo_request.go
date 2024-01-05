package dtos

type PostCargoRequest struct {
	Name        string `json:"cargo_name"`
	ContentType string `json:"cargo_content_type"`
}
