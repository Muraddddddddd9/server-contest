package api

import (
	"contest/constants"
	"contest/utils"

	"github.com/gofiber/fiber/v2"
)

func ClearData(c *fiber.Ctx) error {
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

	if user.Status != constants.TeacherStatus {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": constants.ErrInternalServer,
		})
	}

	constants.NowStageLesson = 1

	constants.UserOnlyAnswer = make(map[string]bool)
	constants.UserTeamAnswer = make(map[int]bool)

	constants.UserToID = make(map[string]string)
	constants.UserToID["Учитель"] = "0"

	constants.Users = make(map[string]*constants.UserStruct)
	constants.Users["0"] = &constants.UserStruct{
		Name:        "Учитель",
		Password:    "BIM_LOCAL123",
		StatusEntry: true,
		Status:      constants.TeacherStatus,
		Team:        0,
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": constants.SuccClearData,
	})
}
