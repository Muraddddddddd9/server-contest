package api

import (
	"contest/constants"
	"contest/handlers/ws"
	"contest/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type ActionLesson struct {
	Action int `json:"action"`
}

func StageLessonChange(c *fiber.Ctx) error {
	session := c.Cookies(constants.SessionKey)
	var action ActionLesson

	if err := c.BodyParser(&action); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": constants.ErrInputValue,
		})
	}

	if session == "" {
		utils.CleatCookieForExit(c)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": constants.ErrUserNotFound,
			"redirect": "/",
		})
	}

	userName := strings.Split(session, ":")
	user, _ := utils.GetUserData(userName[0])
	if user == nil {
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

	if user.Status != constants.TeacherStatus {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": constants.ErrInternalServer,
		})
	}

	constants.NowStageLesson = action.Action

	_, _, flagOnly := ws.TimeOnly.GetDataTime()
	if *flagOnly {
		*flagOnly = false
	}

	_, _, flagTeam := ws.TimeTeam.GetDataTime()
	if *flagTeam {
		*flagTeam = false
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": constants.SuccUpdateLessonStage,
	})
}
