package main

import (
	"TaskManager/api"
	"TaskManager/task"
	"fmt"
)

func main() {
	list := task.NewTasksList()
	httpHandlers := api.NewHTTPhandler(&list)
	httpServer := api.NewHTTPServer(httpHandlers)

	if err := httpServer.StartServer(); err != nil {
		fmt.Println(err)
	}

}
