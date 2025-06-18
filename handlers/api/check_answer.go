package api

import (
	"contest/constants"
	"contest/handlers/ws"
	"contest/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type AnswerStruct struct {
	ID     string `json:"id"`
	Answer string `json:"answer"`
}

type GetAnswer struct {
	AnswerUser []AnswerStruct `json:"answer"`
}

var userOnlyAnswer = map[string]bool{}
var userTeamAnswer = map[uint8]bool{}

func CheckAnswer(c *fiber.Ctx) error {
	min, sec := ws.GetTimeOnlyFunc()

	if !constants.FlagTimeOnlyTest {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": constants.ErrTestNotStart,
		})
	}

	if min == 0 && sec == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": constants.ErrTestEnd,
		})
	}

	session := c.Cookies(constants.SessionKey)
	if session == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": constants.ErrUserExit,
		})
	}

	var getAnswerUser GetAnswer
	var err error
	if err = c.BodyParser(&getAnswerUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": constants.ErrInternalServer,
		})
	}

	if len(getAnswerUser.AnswerUser) != len(constants.Questions) {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": constants.ErrNoFullAnser,
		})
	}

	userName := strings.Split(session, ":")
	user := utils.GetUserData(userName[0])
	if user == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": constants.ErrUserNotFound,
		})
	}

	if userOnlyAnswer[userName[0]] {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": constants.ErrAlreadyReplied,
		})
	}

	question := constants.Questions
	for _, q := range question {
		for _, a := range getAnswerUser.AnswerUser {
			if q.ID == a.ID && q.AnswerTrue == a.Answer {
				user.BimCoin += q.Socer
				break
			}
		}
	}

	userOnlyAnswer[userName[0]] = true

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": constants.SuccGetAnswer,
	})
}

func CheckAnswerTeam(c *fiber.Ctx) error {
	min, sec := ws.GetTimeTeamFunc()

	if !constants.FlagTimeTeamTest {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": constants.ErrTestNotStart,
		})
	}

	if min == 0 && sec == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": constants.ErrTestEnd,
		})
	}

	if constants.NowStageLesson != 5 {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": constants.ErrInternalServer,
		})
	}

	session := c.Cookies(constants.SessionKey)
	if session == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": constants.ErrUserExit,
		})
	}

	var getAnswerTeam GetAnswer
	if err := c.BodyParser(&getAnswerTeam); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": constants.ErrInternalServer,
		})
	}

	if len(getAnswerTeam.AnswerUser) != len(constants.QuestionsTeam) {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": constants.ErrNoFullAnser,
		})
	}

	userName := strings.Split(session, ":")
	user := utils.GetUserData(userName[0])
	if user == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": constants.ErrUserNotFound,
		})
	}

	if userTeamAnswer[user.Team] {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": constants.ErrAlreadyReplied,
		})
	}

	if !user.TeamLeader {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": constants.ErrNotLeader,
		})
	}

	users := constants.Users

	question := constants.QuestionsTeam
	for _, q := range question {
		for _, a := range getAnswerTeam.AnswerUser {
			if q.ID == a.ID && q.AnswerTrue == a.Answer {
				for key := range users {
					usr := utils.GetUserData(key)
					if usr.Team == user.Team {
						usr.BimCoin += q.Socer
					}
				}
				break
			}
		}
	}

	userTeamAnswer[user.Team] = true

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": constants.SuccGetAnswer,
	})
}
