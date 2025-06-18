package api

import (
	"contest/constants"
	"contest/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type ActionLesson struct {
	Action string `json:"action"`
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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": constants.ErrUserNotFound,
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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": constants.ErrInternalServer,
		})
	}

	switch strings.ToLower(action.Action) {
	case "next":
		if constants.NowStageLesson == 7 {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"message": constants.ErrStageLessonEnd,
			})
		}
		constants.NowStageLesson++
	case "prev":
		if constants.NowStageLesson == 1 {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"message": constants.ErrStageLessonEnd,
			})
		}
		constants.NowStageLesson--
	}

	if constants.FlagTimeOnlyTest {
		constants.FlagTimeOnlyTest = false
	}
	
	if constants.FlagTimeTeamTest {
		constants.FlagTimeTeamTest = false
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
		"message": constants.SuccUpdateLessonStage,
	})
}
