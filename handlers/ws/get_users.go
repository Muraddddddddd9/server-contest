package ws

import (
	"contest/constants"
	"encoding/json"
	"log"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/contrib/websocket"
)

type SendUserDataStruct struct {
	Username   string `json:"username"`
	BimCoin    uint64 `json:"bim_coin"`
	Team       uint8  `json:"team"`
	TeamLeader bool   `json:"team_leader"`
}

func GetUsers(c *websocket.Conn) {
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

	var lastSendUser []SendUserDataStruct

	for range ticker.C {
		users := constants.Users
		var sendUser []SendUserDataStruct

		for k := range users {
			if k != "Teacher" {
				sendUser = append(sendUser, SendUserDataStruct{
					Username:   k,
					BimCoin:    users[k].BimCoin,
					Team:       users[k].Team,
					TeamLeader: users[k].TeamLeader,
				})
			}
		}

		sort.Slice(sendUser, func(i, j int) bool {
			numI := extractNumber(sendUser[i].Username)
			numJ := extractNumber(sendUser[j].Username)
			return numI < numJ
		})

		if usersEqual(lastSendUser, sendUser) {
			continue
		}

		data, err := json.Marshal(sendUser)
		if err != nil {
			log.Println(constants.ErrInternalServer)
			return
		}

		if err := c.WriteMessage(1, data); err != nil {
			log.Println(constants.ErrInternalServer, err)
			return
		}

		lastSendUser = sendUser
	}
}

func usersEqual(a, b []SendUserDataStruct) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func extractNumber(username string) int {
	parts := strings.Split(username, "_")
	if len(parts) != 2 {
		return 0
	}
	num, _ := strconv.Atoi(parts[1])
	return num
}
