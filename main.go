package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/iamtonmoy0/Go-Fiber-CRM/database"
	"github.com/iamtonmoy0/Go-Fiber-CRM/lead"
	"github.com/jinzhu/gorm"
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
	app.Get(GetLeads)
	app.Get(GetLead)
	app.Post(NewLead)
	app.Delete(DeleteLead)
}

func main() {
	app := fiber.New()
	setupRoutes(app)

	app.Listen(3000)
	defer database.DBConn.Close()

}
