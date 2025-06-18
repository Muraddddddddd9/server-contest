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
	minuteTeamTest = 1
	secondTeamTest = 0
)

func GetTimeTeam(c *websocket.Conn) {
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
		if constants.FlagTimeTeamTest {
			if user.Status == constants.TeacherStatus {
				if minuteTeamTest > 0 || secondTeamTest > 0 {
					if secondTeamTest == 0 {
						minuteTeamTest--
						secondTeamTest = 59
					} else {
						secondTeamTest--
					}
				}
			}
		}

		if err := c.WriteJSON(map[string]any{"time": fmt.Sprintf("%02d:%02d", minuteTeamTest, secondTeamTest), "flag": constants.FlagTimeTeamTest}); err != nil {
			log.Println(constants.ErrInternalServer, err)
			return
		}
	}
}

func GetTimeTeamFunc() (int, int) {
	return minuteTeamTest, secondTeamTest
}
