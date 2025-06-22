package main

import (
	"contest/handlers/api"
	"contest/handlers/ws"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	cfg, err := LoadConfig()
	if err != nil {
		panic(err)
	}

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     cfg.ORIGIN_URL,
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST",
	}))

	app.Post("/api/entry", func(c *fiber.Ctx) error {
		return api.Entry(c)
	})

	app.Post("/api/exit", func(c *fiber.Ctx) error {
		return api.Exit(c)
	})

	app.Post("/api/change_lesson", func(c *fiber.Ctx) error {
		return api.StageLessonChange(c)
	})

	app.Get("/api/questions", func(c *fiber.Ctx) error {
		return api.GetQuestions(c)
	})

	app.Get("/api/questions/team", func(c *fiber.Ctx) error {
		return api.GetQuestionsTeam(c)
	})

	app.Post("/api/check_answer", func(c *fiber.Ctx) error {
		return api.CheckAnswerOnly(c)
	})

	app.Post("/api/check_answer/team", func(c *fiber.Ctx) error {
		return api.CheckAnswerTeam(c)
	})

	app.Post("/api/change_time", func(c *fiber.Ctx) error {
		return api.ChangeTime(c)
	})

	app.Post("/api/redact_time", func(c *fiber.Ctx) error {
		return api.RedactTime(c)
	})

	app.Post("/api/redact_presentation", func(c *fiber.Ctx) error {
		return api.RedactPresentation(c)
	})

	app.Get("/api/get_presentation", func(c *fiber.Ctx) error {
		return api.GetPresentation(c)
	})

	app.Post("/api/clear", func(c *fiber.Ctx) error {
		return api.ClearData(c)
	})

	app.Get("/ws/stage_lesson", websocket.New(func(c *websocket.Conn) {
		ws.GetStageLesson(c)
	}))

	app.Get("/ws/get_users", websocket.New(func(c *websocket.Conn) {
		ws.GetUsers(c)
	}))

	app.Get("/ws/time_lesson", websocket.New(func(c *websocket.Conn) {
		ws.GetLessonTime(c)
	}))

	app.Get("/ws/time_only", websocket.New(func(c *websocket.Conn) {
		ws.GetOnlyTime(c)
	}))

	app.Get("/ws/time_team", websocket.New(func(c *websocket.Conn) {
		ws.GetTeamTime(c)
	}))

	app.Listen(":8080")
}
