package router

import (
	"encoding/json"
	"net/http"
	"scibat/src/common/domain/event"
	commandResponse "scibat/src/competition/infrastructure/cqrs/response/command"
)

func handler(command event.Command) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		command.Execute()
		err := json.NewEncoder(writer).Encode(commandResponse.New())
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
			return
		}
	}
}
