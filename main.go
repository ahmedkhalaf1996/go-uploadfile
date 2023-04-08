package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	// load templates
	engine := html.New("./views", ".html")

	// create fiber app
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Serve static assets
	app.Static("/", "./public")

	// Routing
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})

	// upalod routing
	app.Post("/", func(c *fiber.Ctx) error {
		file, err := c.FormFile("upload")

		if err != nil {
			return err
		}

		c.SaveFile(file, "public/"+file.Filename)

		return c.Render("index", fiber.Map{})
	})

	// Start app
	app.Listen(":3000")
}
