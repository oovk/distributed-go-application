package main

import (
	"context"
	"distributedgoapp/log"
	"distributedgoapp/service"
	"fmt"
	stlog "log"
)

func main() {
	log.Run("./app.log") //write all log entries to app.log

	host, port := "localhost", "4000"

	ctx, err := service.Start(context.Background(), "log service", host, port, log.RegisterHandlers)
	if err != nil {
		stlog.Fatal(err)
	}
	<-ctx.Done()
	fmt.Println("Shutting down the service")
}
