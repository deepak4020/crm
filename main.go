package main

import (
	"fmt"

	"github.com/deepak4020/go-crm/database"
	"github.com/deepak4020/go-crm/lead"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead:id", lead.GetLead)
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead:id", lead.DeleteLead)

}

func initDatabase() {

	var err error
	database.DBConn, err = gorm.Open("sqlite3", "lead.db")
	if err != nil {
		panic("failed to conncet")

	}
	fmt.Println("successful")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("db transfer ")
}

func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(3000)
	defer database.DBConn.Close()
}
