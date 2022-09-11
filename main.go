package main

import (
	"fmt"

	"github.com/caiosev/crm/db"
	"github.com/caiosev/crm/lead"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Routes(app *fiber.App) {
	app.Get("/api/Leads", lead.GetLeads)
	app.Get("/api/Lead/:id", lead.GetLead)
	app.Post("/api/Lead", lead.NewLead)
	app.Put("/api/Lead/:id", lead.UpdateLead)
	app.Delete("/api/Lead/:id", lead.DeleteLead)
	app.Get("/api/NextStep/:id/:step", lead.NextStep)

}

func initDb() {
	var err error
	db.DBCon, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic("Erro ao conectar com o banco")
	}
	fmt.Println("Conectado ao banco")
	db.DBCon.AutoMigrate(&lead.Lead{})
	fmt.Println("Banco Migrado")
}

func main() {
	app := fiber.New()
	initDb()
	Routes(app)
	app.Listen(":3000")
	defer db.DBCon.Close()
}
