package model

type ErrorResponse struct {
	Error string `json:"error"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

type DocumentUpdate struct {
	Title   string `json:"title"`
	Owner   string `json:"owner"`
	Payload string `json:"payload"`
}

type DocumentCreate struct {
	ID uint `json:"id"`
}

type FileUpload struct {
	ID uint `json:"id"`
}

type GetDocuments struct {
	Items []Document `json:"items"`
	Total uint       `json:"total"`
}