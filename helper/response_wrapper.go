package helper

// ResponseWrapper is the structure for the response.
type ResponseWrapper struct {
	Success *bool       `json:"success"` // Pointer to bool for optional success
	Message *string     `json:"message"` // Pointer to string for optional message
	Data    interface{} `json:"data"`    // Data can be nil or any type
}

// ResponseSuccess is a helper function to send JSON responses with optional fields.
func ResponseSuccess(res ResponseWrapper) ResponseWrapper {
	// Set default value for Success if it's nil
	if res.Success == nil {
		defaultSuccess := true
		res.Success = &defaultSuccess
	}

	// Set default value for Message if it's nil
	if res.Message == nil {
		defaultMessage := "Request was successful"
		res.Message = &defaultMessage
	}

	// Send the JSON response
	return res
}
