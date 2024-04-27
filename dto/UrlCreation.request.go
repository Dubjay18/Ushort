package dto

// UrlCreationRequest represents the request body for creating a short URL
type UrlCreationRequest struct {
	// The long URL to be shortened
	LongUrl string `json:"long_url"binding:"required"`
	// The user ID of the user who created the short URL
	//UserId string `json:"user_id"binding:"required"`
}
