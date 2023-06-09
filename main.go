package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/iamtonmoy0/Go-Fiber-CRM/database"
	"github.com/iamtonmoy0/Go-Fiber-CRM/lead"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// db
func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Connection Opened to database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database migrated")

}

// routes
func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func main() {
	app := fiber.New()
	setupRoutes(app)
	app.Listen(3000)
	defer database.DBConn.Close()

}
