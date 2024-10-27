package router

import (
	"encoding/json"
	"net/http"
	"scibat/src/common/domain/query"
	queryResponse "scibat/src/competition/infrastructure/cqrs/response/query"
)

func handler(command query.Command) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		data := command.Execute()
		err := json.NewEncoder(writer).Encode(queryResponse.New(data))
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
	}
}
