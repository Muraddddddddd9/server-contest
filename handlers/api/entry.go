package api

import (
	"contest/constants"
	"contest/utils"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
)

type EntryStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Entry(c *fiber.Ctx) error {
	var entryData EntryStruct
	if err := c.BodyParser(&entryData); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": constants.ErrInputValue,
		})
	}

	user := utils.GetUserData(entryData.Username)
	if user == nil || user.Password != entryData.Password {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": constants.ErrUserNotFound,
		})
	}

	if user.StatusEntry {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": constants.ErrAlreadyEntry,
		})
	}

	user.StatusEntry = true

	sessionID := fmt.Sprintf("%v:%v", entryData.Username, encryptcookie.GenerateKey())
	c.Cookie(&fiber.Cookie{
		Name:     constants.SessionKey,
		Value:    sessionID,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Lax",
		Path:     "/",
	})

	c.Cookie(&fiber.Cookie{
		Name:     constants.StatusKey,
		Value:    user.Status,
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: false,
		Secure:   true,
		SameSite: "Lax",
		Path:     "/",
	})

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message":  constants.SuccEntry,
		"redirect": "/lesson",
	})
}
