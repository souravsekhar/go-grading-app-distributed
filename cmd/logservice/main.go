package main

import (
	"context"
	"fmt"
	stlog "log"

	"github.com/souravsekhar/go-grading-app-distributed/log"
	"github.com/souravsekhar/go-grading-app-distributed/registry"
	"github.com/souravsekhar/go-grading-app-distributed/service"
)

func main() {
	log.Run("./app.log")
	host, port := "localhost", "4000"
	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)

	var r registry.Registration
	r.ServiceName = registry.LogService
	r.ServiceURL = serviceAddress

	ctx, err := service.Start(
		context.Background(),
		r,
		host,
		port,
		log.RegisterHandlers)
	if err != nil {
		stlog.Fatal(err)
	}
	<-ctx.Done()
}
