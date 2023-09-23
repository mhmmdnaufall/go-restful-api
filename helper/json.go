package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody[T any](request *http.Request, result T) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	PanicIfError(err)

}

func WriteToResponseBody[T any](writer http.ResponseWriter, response T) {
	writer.Header().Add("Content-Type", "application/json")

	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	PanicIfError(err)

}
