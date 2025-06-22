package ws

import (
	"contest/constants"
	"contest/utils"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/contrib/websocket"
)

type TimeData struct {
	Minute int
	Second int
	Flag   bool
}

func CreateNewTime(minute, second int, flag bool) *TimeData {
	return &TimeData{Minute: minute, Second: second, Flag: flag}
}

func (t *TimeData) GetDataTime() (*int, *int, *bool) {
	return &t.Minute, &t.Second, &t.Flag
}

func (t *TimeData) CountdownTime(user *constants.UserStruct) interface{} {
	if t.Flag && user.Status == constants.TeacherStatus {
		if t.Minute > 0 || t.Second > 0 {
			if t.Second == 0 {
				t.Minute--
				t.Second = 59
			} else {
				t.Second--
			}
		}
	}

	data := map[string]interface{}{
		"time": fmt.Sprintf("%02d:%02d", t.Minute, t.Second),
		"flag": t.Flag,
	}

	return data
}

func GetTime(c *websocket.Conn, timeData *TimeData) {
	defer func() {
		c.Close()
		log.Println(constants.SuccCloseWS)
	}()

	session := c.Cookies(constants.SessionKey)
	user, errFoundUser := utils.GetUserDataSession(session)
	if errFoundUser != "" && user == nil {
		err := c.WriteMessage(1, []byte(errFoundUser))
		if err != nil {
			log.Printf("%v\n", errFoundUser)
		}
		return
	}

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		data := timeData.CountdownTime(user)

		if err := c.WriteJSON(data); err != nil {
			log.Println(constants.ErrInternalServer, err)
			return
		}
	}
}

var TimeLesson = CreateNewTime(45, 0, false)
var TimeOnly = CreateNewTime(2, 0, false)
var TimeTeam = CreateNewTime(1, 0, false)

func GetLessonTime(c *websocket.Conn) {
	GetTime(c, TimeLesson)
}

func GetOnlyTime(c *websocket.Conn) {
	GetTime(c, TimeOnly)
}

func GetTeamTime(c *websocket.Conn) {
	GetTime(c, TimeTeam)
}
