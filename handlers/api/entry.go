package api

import (
	"contest/constants"
	"fmt"
	"strings"
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

	username := strings.TrimSpace(entryData.Username)
	password := strings.TrimSpace(entryData.Password)

	userId := constants.UserToID[username]
	user := constants.Users[userId]
	if user == nil {
		newUserData := constants.UserStruct{
			Name:        username,
			Password:    password,
			Status:      constants.StudentStatus,
			StatusEntry: false,
			BimCoin:     0,
			Team:        len(constants.Users) % 2,
			TeamLeader:  false,
		}
		if !constants.OneTeamLeader {
			newUserData.TeamLeader = true
			constants.OneTeamLeader = true
		} else if !constants.TwoTeamLeader {
			newUserData.TeamLeader = true
			constants.TwoTeamLeader = true
		}

		fmt.Println(constants.OneTeamLeader)

		userId = fmt.Sprint(len(constants.Users))
		constants.UserToID[username] = userId
		constants.Users[userId] = &newUserData
		user = &newUserData
	} else {
		if user.Password != password {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": constants.ErrUserNotFound,
			})
		}
	}

	if user.StatusEntry {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": constants.ErrAlreadyEntry,
		})
	}

	user.StatusEntry = true

	sessionID := fmt.Sprintf("%v:%v", userId, encryptcookie.GenerateKey())
	c.Cookie(&fiber.Cookie{
		Name:     constants.SessionKey,
		Value:    sessionID,
		Expires:  time.Now().Add(50 * time.Minute),
		HTTPOnly: true,
		Secure:   false,
		SameSite: "Lax",
		Path:     "/",
	})

	c.Cookie(&fiber.Cookie{
		Name:     constants.StatusKey,
		Value:    user.Status,
		Expires:  time.Now().Add(50 * time.Minute),
		HTTPOnly: false,
		Secure:   false,
		SameSite: "Lax",
		Path:     "/",
	})

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message":  constants.SuccEntry,
		"redirect": "/lesson",
	})
}
