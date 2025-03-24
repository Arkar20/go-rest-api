package helper

// APIResponse creates a standardized JSON response
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"` // Omits "data" if nil
}
