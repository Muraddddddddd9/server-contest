package ws

import (
	"contest/constants"
	"contest/utils"
	"encoding/json"
	"log"
	"sort"
	"time"

	"github.com/gofiber/contrib/websocket"
)

type SendUserDataStruct struct {
	Username   string `json:"username"`
	BimCoin    uint64 `json:"bim_coin"`
	Team       int    `json:"team"`
	TeamLeader bool   `json:"team_leader"`
}

func GetUsers(c *websocket.Conn) {
	defer func() {
		c.Close()
		log.Println(constants.SuccCloseWS)
	}()

	session := c.Cookies(constants.SessionKey)
	user, errFoundUser := utils.GetUserDataSession(session)
	if errFoundUser != "" && user == nil {
		log.Println(errFoundUser)
		return
	}

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	var lastSendUser []SendUserDataStruct

	for range ticker.C {
		usersIDs := constants.UserToID
		users := constants.Users
		var sendUser []SendUserDataStruct

		for k := range usersIDs {
			id := usersIDs[k]
			user := users[id]
			if user.Status != constants.TeacherStatus {
				sendUser = append(sendUser, SendUserDataStruct{
					Username:   user.Name,
					BimCoin:    user.BimCoin,
					Team:       user.Team,
					TeamLeader: user.TeamLeader,
				})
			}
		}

		sort.Slice(sendUser, func(i, j int) bool {
			numI := sendUser[i].Username
			numJ := sendUser[j].Username
			return numI < numJ
		})

		sort.Slice(sendUser, func(i, j int) bool {
			return sendUser[i].BimCoin > sendUser[j].BimCoin
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
