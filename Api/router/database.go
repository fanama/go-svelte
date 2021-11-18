package router

import (
	"encoding/json"

	"github.com/fanama/go-svelte/Api/database"
	"github.com/gofiber/fiber/v2"
)

func InsertData(c *fiber.Ctx) (err error) {
	body := c.Body()
	db := c.Params("database")

	var data map[string]string

	err = json.Unmarshal(body, &data)

	if err != nil {
		return c.Status(400).JSON(err)
	}

	err = database.Write(db, data)

	if err != nil {
		return c.Status(400).JSON(err)
	}

	return c.JSON("success!!")
}

func GetData(c *fiber.Ctx) error {
	db := c.Params("database")

	result, err := database.Read(db)

	if err != nil {
		return c.Status(400).JSON(err)
	}

	return c.JSON(result)
}

func DeleteData(c *fiber.Ctx) error {
	db := c.Params("database")

	result, err := database.Read(db)

	if err != nil {
		return c.Status(400).JSON(err)
	}

	return c.JSON(result)
}
