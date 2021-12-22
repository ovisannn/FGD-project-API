package main

import (
	"context"
	"log"
	"time"

	_routes "disspace/app/routes"
	_mongoDriver "disspace/drivers/mongoDB"

	_threadUseCase "disspace/business/threads"
	_threadController "disspace/controllers/threads"
	_threadRepository "disspace/drivers/databases/threads"

	"github.com/labstack/echo/v4"
)

func main() {

	var ctx = context.Background()
	db, _ := _mongoDriver.ConnectDB(ctx)

	e := echo.New()
	timeoutContext := time.Duration(20 * time.Second)

	threadRepository := _threadRepository.NewMongoDBThreadRepository(db)
	threadUseCase := _threadUseCase.NewThreadUseCase(threadRepository, timeoutContext)
	threadController := _threadController.NewThreadController(threadUseCase)

	routesInit := _routes.ControllerList{
		ThreadController: *threadController,
	}

	routesInit.RouteRegister(e)
	log.Fatal(e.Start(":8080"))
}
