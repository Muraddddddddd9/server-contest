package api

import (
	"contest/constants"
	"contest/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type ChangeTimeStruct struct {
	ChageTime string `json:"change_time"`
}

func ChangeTime(c *fiber.Ctx) error {
	session := c.Cookies(constants.SessionKey)
	var changeTime ChangeTimeStruct

	if err := c.BodyParser(&changeTime); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": constants.ErrInputValue,
		})
	}

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

	if user.Status != constants.TeacherStatus {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": constants.ErrInternalServer,
		})
	}

	switch changeTime.ChageTime {
	case "lesson":
		constants.FlagTime = !constants.FlagTime
	case "only":
		constants.FlagTimeOnlyTest = !constants.FlagTimeOnlyTest
	case "team":
		constants.FlagTimeTeamTest = !constants.FlagTimeTeamTest
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": constants.SuccChangeTime,
	})
}
