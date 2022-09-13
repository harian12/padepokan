package response

func ResponseSuccess(data interface{}, message string) interface{} {
	response := map[string]interface{}{
		"status":  true,
		"data":    data,
		"error":   false,
		"message": message,
	}
	return response
}

func ResponseError(message string) interface{} {
	response := map[string]interface{}{
		"status":  false,
		"data":    []string{},
		"error":   true,
		"message": message,
	}
	return response
}
