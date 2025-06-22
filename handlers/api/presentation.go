package api

import (
	"contest/constants"
	"contest/utils"

	"github.com/gofiber/fiber/v2"
)

type RedactPresentationSturct struct {
	ID string `json:"id"`
}

func RedactPresentation(c *fiber.Ctx) error {
	session := c.Cookies(constants.SessionKey)
	var dataPresentation RedactPresentationSturct

	if err := c.BodyParser(&dataPresentation); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": constants.ErrInputValue,
		})
	}

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

	constants.IdPresentation = dataPresentation.ID
	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": constants.SuccUpdateIdPresentation,
	})
}

func GetPresentation(c *fiber.Ctx) error {
	session := c.Cookies(constants.SessionKey)

	user, errFoundUser := utils.GetUserDataSession(session)
	if errFoundUser == "" && user == nil {
		utils.CleatCookieForExit(c)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": constants.ErrUserNotFound,
			"redirect": "/",
		})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"id": constants.IdPresentation,
	})
}
