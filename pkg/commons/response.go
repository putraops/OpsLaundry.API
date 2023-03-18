package commons

// type StandartResult struct {
// 	Success bool        `json:"success"`
// 	Message string      `json:"message"`
// 	Data    interface{} `json:"data"`
// }

// type EmptyObj struct{}
// type EmptyArray []struct{}

// func Result(success bool, message string, data interface{}) StandartResult {
// 	res := StandartResult{
// 		Success: success,
// 		Message: message,
// 		Data:    data,
// 	}
// 	return res
// }

func ErrorResponse(message string) interface{} {
	return map[string]interface{}{"error": message}
}
