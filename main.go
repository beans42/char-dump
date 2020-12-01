package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gofiber/fiber"
)

const (
	port         = ":80"
	databaseFile = "./database.json"
)

var pastes map[string]string = make(map[string]string)

func randomHexString() (out string) {
	bytes := make([]byte, 20)
	rand.Read(bytes)
	for _, v := range bytes {
		out += fmt.Sprintf("%02x", v)
	}
	return
}

func getPaste(c *fiber.Ctx) error { // /:id
	text, ok := pastes[c.Params("id")]
	if ok {
		return c.SendString(text)
	}
	return c.SendString("invalid id")
}

func publishPaste(c *fiber.Ctx) error {
	var query struct {
		Text string `query:"text"`
	}

	if err := c.BodyParser(&query); err != nil {
		return err
	}

	randomString := randomHexString()
	pastes[randomString] = query.Text
	bytes, _ := json.Marshal(pastes)
	defer ioutil.WriteFile(databaseFile, bytes, 0644)
	return c.Redirect("/" + randomString)
}

func main() {
	app := fiber.New()

	content, err := ioutil.ReadFile(databaseFile)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(content, &pastes)
	if err != nil {
		log.Fatal(err)
	}

	app.Static("/", "./static")
	app.Get("/:id", getPaste)
	app.Post("/", publishPaste)

	app.Listen(port)
}
