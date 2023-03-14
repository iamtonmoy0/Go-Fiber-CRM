package routes

import "github.com/gofiber/fiber"

func main(app *fiber.App) {
	app.Get(GetLeads)
	app.Post(NewLead)
	app.Delete(DeleteLead)
}
