package api

import (
	"contest/constants"
	"contest/utils"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type SendQuestionsStruct struct {
	ID       string   `json:"id"`
	Question string   `json:"question"`
	Answers  []string `json:"answers"`
	Socer    uint64   `json:"socer"`
}

func GetQuestions(c *fiber.Ctx) error {
	questionsJSON, err := json.Marshal(constants.Questions)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": constants.ErrInternalServer,
		})
	}

	var sendQuestions []SendQuestionsStruct
	err = json.Unmarshal(questionsJSON, &sendQuestions)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": constants.ErrInternalServer,
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"questions": sendQuestions,
	})
}

func GetQuestionsTeam(c *fiber.Ctx) error {
	session := c.Cookies(constants.SessionKey)
	if session == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": constants.ErrUserExit,
		})
	}

	userName := strings.Split(session, ":")
	user := utils.GetUserData(userName[0])
	if user == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": constants.ErrUserNotFound,
		})
	}

	if !user.TeamLeader && user.Status == constants.StudentStatus {
		var leader string
		for k := range constants.Users {
			if constants.Users[k].Team == user.Team && constants.Users[k].TeamLeader {
				leader = k
			}
		}

		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
			"questions": nil,
			"message":   fmt.Sprintf("Ваша команда №%v. Ваш лидер: '%v'", user.Team, leader),
		})
	}

	questionsJSON, err := json.Marshal(constants.QuestionsTeam)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": constants.ErrInternalServer,
		})
	}

	var sendQuestions []SendQuestionsStruct
	err = json.Unmarshal(questionsJSON, &sendQuestions)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": constants.ErrInternalServer,
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"questions": sendQuestions,
	})
}
