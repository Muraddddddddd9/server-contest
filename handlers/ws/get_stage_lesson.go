package ws

import (
	"contest/constants"
	"log"
	"time"

	"github.com/gofiber/contrib/websocket"
)

func GetStageLesson(c *websocket.Conn) {
	defer func() {
		c.Close()
		log.Println(constants.SuccCloseWS)
	}()

	session := c.Cookies(constants.SessionKey)
	if session == "" {
		log.Println(constants.ErrUserNotFound)
		return
	}

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	var lastSendStage int

	for range ticker.C {
		if lastSendStage == constants.NowStageLesson {
			continue
		}

		if err := c.WriteJSON(map[string]int{"lesson": constants.NowStageLesson}); err != nil {
			log.Println(constants.ErrInternalServer, err)
			return
		}

		lastSendStage = constants.NowStageLesson
	}
}
