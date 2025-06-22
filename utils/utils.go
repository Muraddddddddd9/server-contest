package utils

import (
	"contest/constants"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetUserData(key string) (*constants.UserStruct, string) {
	user, ok := constants.Users[key]
	if !ok {
		return nil, ""
	}

	return user, key
}

func GetUserDataSession(session string) (*constants.UserStruct, string) {
	if session == "" {
		return nil, constants.ErrUserNotFound
	}

	userName := strings.Split(session, ":")
	user, _ := GetUserData(userName[0])
	if user == nil {
		return nil, constants.ErrUserNotFound
	}

	return user, ""
}

func CleatCookieForExit(c *fiber.Ctx) {
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
}
