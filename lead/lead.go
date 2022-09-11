package lead

import (
	"fmt"
	"net/mail"
	"strconv"

	"github.com/caiosev/crm/db"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type LeadORString interface {
	string | Lead
}

type Lead struct {
	gorm.Model
	Nome    string `json:"nome"`
	Email   string `json:"email"`
	Empresa int    `json:"empresa"`
	Tel     string `json:"tel"`
	Etapa   int    `json:"etapa"`
}

func GetLeads(c *fiber.Ctx) {
	db := db.DBCon
	var leads []Lead
	db.Find(&leads)
	c.JSON(leads)
}

func GetLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := db.DBCon
	var lead Lead
	db.Find(&lead, id)
	c.JSON(lead)

}

func NewLead(c *fiber.Ctx) {
	db := db.DBCon
	lead := new(Lead)
	var err []string
	if err := c.BodyParser(lead); err != nil {
		c.Status(500).SendString(err.Error())
		return
	}
	validateName(lead.Nome, c, &err)
	validateEmail(lead.Email, c, &err)
	validateEmpresa(lead.Empresa, c, &err)
	fmt.Println(err)
	if len(err) != 0 {
		c.Status(500)
		c.JSON(map[string]interface{}{
			"res": err[0],
		})
	} else {
		lead.Etapa = 1
		db.Create(&lead)
		c.JSON(lead)
		c.JSON(map[string]interface{}{
			"res":  "criado com sucesso",
			"lead": lead,
		})
		c.Status(200)
	}
}

func validateEmpresa(code int, c *fiber.Ctx, err *[]string) {
	if code > 10 || code <= 0 {
		*err = append(*err, "Codigo da empresa Invalido")
	}
}

func validateName(name string, c *fiber.Ctx, err *[]string) {
	if len(name) < 3 {
		*err = append(*err, "Nome Invalido")
	}
}

func validateEmail(email string, c *fiber.Ctx, err *[]string) {
	_, errEmail := mail.ParseAddress(email)
	if errEmail != nil {
		*err = append(*err, "Email Invalido")
	}
}

func UpdateLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := db.DBCon
	var lead Lead
	db.First(&lead, id)
	c.BodyParser(&lead)
	db.Save(&lead)
	c.JSON(map[string]string{
		"res": "Atualizado com sucesso",
	})
	c.Status(200)
}

func DeleteLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := db.DBCon
	var lead Lead
	db.First(&lead, id)
	db.Delete(&lead)
	c.JSON(map[string]string{
		"res": "Deletado com sucesso",
	})
	c.Status(200)
}

func NextStep(c *fiber.Ctx) {
	id := c.Params("id")
	step := c.Params("step")
	db := db.DBCon
	var lead Lead
	db.First(&lead, id)
	val, _ := strconv.Atoi(step)
	lead.Etapa = val
	db.Save(&lead)
	c.JSON(map[string]string{
		"res": "etapa atualizada com sucesso",
	})
	c.Status(200)

}
