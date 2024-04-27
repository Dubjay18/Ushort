package dto

// UrlCreationRequest represents the request body for creating a short URL
type UrlCreationRequest struct {
	// The long URL to be shortened
	LongUrl string `json:"long_url"binding:"required"`
}

type UrlCreationResponse struct {
	Message  string `json:"message"`
	ShortUrl string `json:"short_url"`
}
