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
		// "username": "Teacher",
		// "password": "BIM_LOCAL123"
		return api.Entry(c)
	})

	app.Post("/api/exit", func(c *fiber.Ctx) error {
		// Cookie,
		return api.Exit(c)
	})

	app.Post("/api/change_lesson", func(c *fiber.Ctx) error {
		// "action": "next | prev"
		// Cookie
		return api.StageLessonChange(c)
	})

	app.Get("/api/questions", func(c *fiber.Ctx) error {
		return api.GetQuestions(c)
	})

	app.Get("/api/questions/team", func(c *fiber.Ctx) error {
		return api.GetQuestionsTeam(c)
	})

	app.Post("/api/check_answer", func(c *fiber.Ctx) error {
		// "answer": [
		//     {
		//         "id": "Questions_1",
		//         "answer": "Норм"
		//     },
		//     {
		//         "id": "Questions_2",
		//         "answer": "Как как"
		//     }
		// ]
		return api.CheckAnswer(c)
	})

	app.Post("/api/check_answer/team", func(c *fiber.Ctx) error {
		// "answer": [
		//     {
		//         "id": "Questions_Team_1",
		//         "answer": "ОГООГО"
		//     },
		//     {
		//         "id": "Questions_Team_2",
		//         "answer": "Как"
		//     }
		// ]
		// "team": 1
		return api.CheckAnswerTeam(c)
	})

	app.Post("/api/change_time", func(c *fiber.Ctx) error {
		return api.ChangeTime(c)
	})

	app.Get("/ws/stage_lesson", websocket.New(func(c *websocket.Conn) {
		ws.GetStageLesson(c)
	}))

	app.Get("/ws/get_time", websocket.New(func(c *websocket.Conn) {
		ws.GetTimeLesson(c)
	}))

	app.Get("/ws/get_users", websocket.New(func(c *websocket.Conn) {
		ws.GetUsers(c)
	}))

	app.Get("/ws/get_time_only", websocket.New(func(c *websocket.Conn) {
		ws.GetTimeOnly(c)
	}))

	app.Get("/ws/get_time_team", websocket.New(func(c *websocket.Conn) {
		ws.GetTimeTeam(c)
	}))

	app.Listen(":8080")
}
