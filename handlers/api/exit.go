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
		utils.CleatCookieForExit(c)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": constants.ErrUserExit,
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

	user.StatusEntry = false

	c.Cookie(&fiber.Cookie{
		Name:     constants.SessionKey,
		Value:    "",
		MaxAge:   -1,
		HTTPOnly: true,
		Secure:   false,
		SameSite: "Lax",
		Path:     "/",
	})

	c.Cookie(&fiber.Cookie{
		Name:     constants.StatusKey,
		Value:    "",
		MaxAge:   -1,
		HTTPOnly: false,
		Secure:   false,
		SameSite: "Lax",
		Path:     "/",
	})

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message":  constants.SuccExit,
		"redirect": "/",
	})
}
