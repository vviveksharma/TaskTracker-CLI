package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/vviveksharma/taskTracker-CLI/db"
	"github.com/vviveksharma/taskTracker-CLI/handlers"
	"github.com/vviveksharma/taskTracker-CLI/services"
)

func Server() {
	app := fiber.New()
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
	}))

	//Intilising the database
	go func() {
		db, err := db.NewDbRequest()
		if err != nil {
			log.Fatalln("error in creating a DB request")
			return
		}
		_, err = db.InitDB()
		if err != nil {
			log.Fatalln("error in starting the DataBase: ", err)
			return
		}
	}()
	TaskService, err := services.NewTodoServiceRequest()
	if err != nil {
		log.Println("task service instance starting failure: " + err.Error())
		return
	}

	handlers := handlers.NewHandler().Task(TaskService)
	Routes(app, handlers)

	err = app.Listen(":8000")
	if err != nil {
		log.Print("error in starting the server:", err)
	}
}
