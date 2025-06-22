package api

import (
	"contest/constants"
	"contest/handlers/ws"
	"contest/utils"
	"fmt"
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

func CheckUserOnlyAnswer(username string, user *constants.UserStruct, getAnswer GetAnswer) error {
	if constants.UserOnlyAnswer[username] {
		return fmt.Errorf("%v", constants.ErrAlreadyReplied)
	}

	question := constants.Questions

	for _, q := range question {
		for _, a := range getAnswer.AnswerUser {
			if q.ID == a.ID && q.AnswerTrue == a.Answer {
				user.BimCoin += q.Socer
				break
			}
		}
	}

	constants.UserOnlyAnswer[username] = true

	return nil
}

func CheckUserTeamAnswer(user *constants.UserStruct, getAnswer GetAnswer) error {
	if constants.UserTeamAnswer[user.Team] {
		return fmt.Errorf("%v", constants.ErrAlreadyReplied)
	}

	if !user.TeamLeader {
		return fmt.Errorf("%v", constants.ErrNotLeader)
	}

	users := constants.Users
	question := constants.QuestionsTeam

	for _, q := range question {
		for _, a := range getAnswer.AnswerUser {
			if q.ID == a.ID && q.AnswerTrue == a.Answer {
				for key := range users {
					usr, _ := utils.GetUserData(key)
					if usr.Team == user.Team {
						usr.BimCoin += q.Socer
					}
				}
				break
			}
		}
	}

	constants.UserTeamAnswer[user.Team] = true

	return nil
}

func CheckAnswer(c *fiber.Ctx, test string, time *ws.TimeData) error {
	min, sec, flag := time.GetDataTime()

	if !*flag {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": constants.ErrTestNotStart,
		})
	}

	if *min == 0 && *sec == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": constants.ErrTestEnd,
		})
	}

	session := c.Cookies(constants.SessionKey)
	user, errFoundUser := utils.GetUserDataSession(session)
	if errFoundUser == "" && user == nil {
		utils.CleatCookieForExit(c)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": constants.ErrUserNotFound,
			"redirect": "/",
		})
	}

	if !user.StatusEntry {
		utils.CleatCookieForExit(c)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":  constants.ErrUserNotEntry,
			"redirect": "/",
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

	switch test {
	case "only":
		userName := strings.Split(session, ":")
		err = CheckUserOnlyAnswer(userName[0], user, getAnswerUser)
	case "team":
		err = CheckUserTeamAnswer(user, getAnswerUser)
	}

	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": constants.SuccGetAnswer,
	})
}

func CheckAnswerOnly(c *fiber.Ctx) error {
	return CheckAnswer(c, "only", ws.TimeOnly)
}

func CheckAnswerTeam(c *fiber.Ctx) error {
	return CheckAnswer(c, "team", ws.TimeTeam)
}
