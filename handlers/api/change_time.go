package api

import (
	"contest/constants"
	"contest/handlers/ws"
	"contest/utils"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type ChangeTimeStruct struct {
	ChageTime string `json:"change_time"`
}

type RedactTimeStuct struct {
	ChageTime string `json:"change_time"`
	NewTime   string `json:"new_time"`
}

func ChangeTime(c *fiber.Ctx) error {
	session := c.Cookies(constants.SessionKey)
	var changeTime ChangeTimeStruct

	if err := c.BodyParser(&changeTime); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": constants.ErrInputValue,
		})
	}

	user, errFoundUser := utils.GetUserDataSession(session)
	if errFoundUser == "" && user == nil {
		utils.CleatCookieForExit(c)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":  constants.ErrUserNotFound,
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

	switch changeTime.ChageTime {
	case "lesson":
		_, _, flag := ws.TimeLesson.GetDataTime()
		ws.TimeLesson.Flag = !*flag
	case "only":
		_, _, flag := ws.TimeOnly.GetDataTime()
		ws.TimeOnly.Flag = !*flag
	case "team":
		_, _, flag := ws.TimeTeam.GetDataTime()
		ws.TimeTeam.Flag = !*flag
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": constants.SuccChangeTime,
	})
}

func RedactTime(c *fiber.Ctx) error {
	session := c.Cookies(constants.SessionKey)
	var redactTime RedactTimeStuct

	if err := c.BodyParser(&redactTime); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": constants.ErrInputValue,
		})
	}

	user, errFoundUser := utils.GetUserDataSession(session)
	if errFoundUser == "" && user == nil {
		utils.CleatCookieForExit(c)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message":  constants.ErrUserNotFound,
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

	timeSplit := strings.Split(redactTime.NewTime, ":")
	minute, _ := strconv.Atoi(timeSplit[0])
	second, _ := strconv.Atoi(timeSplit[1])

	switch redactTime.ChageTime {
	case "only":
		ws.TimeOnly.Flag = false
		ws.TimeOnly.Minute = minute
		ws.TimeOnly.Second = second
	case "team":
		ws.TimeTeam.Flag = false
		ws.TimeTeam.Minute = minute
		ws.TimeTeam.Second = second

	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": constants.SuccChangeTime,
	})
}
