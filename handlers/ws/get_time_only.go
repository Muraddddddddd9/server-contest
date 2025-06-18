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
	minuteOnlyTest = 1
	secondOnlyTest = 0
)

func GetTimeOnly(c *websocket.Conn) {
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
		if constants.FlagTimeOnlyTest {
			if user.Status == constants.TeacherStatus {
				if minuteOnlyTest > 0 || secondOnlyTest > 0 {
					if secondOnlyTest == 0 {
						minuteOnlyTest--
						secondOnlyTest = 59
					} else {
						secondOnlyTest--
					}
				}
			}
		}

		if err := c.WriteJSON(map[string]any{"time": fmt.Sprintf("%02d:%02d", minuteOnlyTest, secondOnlyTest), "flag": constants.FlagTimeOnlyTest}); err != nil {
			log.Println(constants.ErrInternalServer, err)
			return
		}
	}
}

func GetTimeOnlyFunc() (int, int) {
	return minuteOnlyTest, secondOnlyTest
}
