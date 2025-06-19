package api

import (
	"contest/constants"
	"contest/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Exit(c *fiber.Ctx) error {
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

	user.StatusEntry = false

	c.Cookie(&fiber.Cookie{
		Name:     constants.SessionKey,
		Value:    "",
		MaxAge:   -1,
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Lax",
		Path:     "/",
	})

	c.Cookie(&fiber.Cookie{
		Name:     constants.StatusKey,
		Value:    "",
		MaxAge:   -1,
		HTTPOnly: false,
		Secure:   true,
		SameSite: "Lax",
		Path:     "/",
	})

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message":  constants.SuccExit,
		"redirect": "/",
	})
}
