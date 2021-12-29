package main

import (
	"log"
	"time"

	_routes "disspace/app/routes"
	_mongoDriver "disspace/drivers/mongoDB"

	_threadUseCase "disspace/business/threads"
	_threadController "disspace/controllers/threads"
	_threadRepository "disspace/drivers/databases/threads"

	_categoryUseCase "disspace/business/categories"
	_categoryController "disspace/controllers/categories"
	_categoryRepository "disspace/drivers/databases/categories"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("app/config.json")
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

}

func main() {
	config := _mongoDriver.Config{
		Username: viper.GetString("database.username"),
		Password: viper.GetString("database.password"),
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		Name:     viper.GetString("database.name"),
	}
	db, _ := config.ConnectDB()

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second

	threadRepository := _threadRepository.NewMongoDBThreadRepository(db)
	threadUseCase := _threadUseCase.NewThreadUseCase(threadRepository, timeoutContext)
	threadController := _threadController.NewThreadController(threadUseCase)

	categoriesRepository := _categoryRepository.NewMongoDBCategoriesRepository(db)
	categoryUseCase := _categoryUseCase.NewCategoriesUseCase(categoriesRepository)
	categoryController := _categoryController.NewCategoriesController(categoryUseCase)

	routesInit := _routes.ControllerList{
		ThreadController:     *threadController,
		CategoriesController: *categoryController,
	}

	routesInit.RouteRegister(e)
	log.Fatal(e.Start(viper.GetString("server.address")))
}
