package controller

import (
	"fmt"
	"strconv"

	"github.com/DangPhuongTay/travelblog-golang/database"
	"github.com/DangPhuongTay/travelblog-golang/models"
	"github.com/gofiber/fiber/v2"
)

func DetailAbout(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var aboutDL models.About
	database.DB.Where("id=?", id).First(&aboutDL)
	return c.JSON(fiber.Map{
		"data": aboutDL,
	})

}

func UpdateAbout(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	about := models.About{
		Id: uint(id),
	}
	if err := c.BodyParser(&about); err != nil {
		fmt.Println("Unable to parse body")
	}
	database.DB.Model(&about).Updates(about)
	return c.JSON(fiber.Map{
		"message": "about update successfully",
	})
}
