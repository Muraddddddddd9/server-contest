package ws

import (
	"contest/constants"
	"contest/utils"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gofiber/contrib/websocket"
)

var (
	minute = 45
	second = 0
)

func GetTimeLesson(c *websocket.Conn) {
	defer func() {
		c.Close()
		log.Println(constants.SuccCloseWS)
	}()

	session := c.Cookies(constants.SessionKey)
	if session == "" {
		log.Println(constants.ErrUserNotFound)
		return
	}

	userName := strings.Split(session, ":")

	user := utils.GetUserData(userName[0])
	if user == nil {
		log.Println(constants.ErrUserNotFound)
		return
	}

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		if constants.FlagTime {
			if user.Status == constants.TeacherStatus {
				if minute > 0 || second > 0 {
					if second == 0 {
						minute--
						second = 59
					} else {
						second--
					}
				}
			}
		}

		if err := c.WriteJSON(map[string]any{"time": fmt.Sprintf("%02d:%02d", minute, second), "flag": constants.FlagTime}); err != nil {
			log.Println(constants.ErrInternalServer, err)
			return
		}
	}
}

func GetTimeLessonFunc() (*int, *int) {
	return &minute, &second
}
