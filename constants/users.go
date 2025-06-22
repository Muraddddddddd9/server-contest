package constants

type UserStruct struct {
	Name        string `json:"name"`
	Password    string `json:"password"`
	Status      string `json:"status"`
	StatusEntry bool   `json:"status_entry"`
	BimCoin     uint64 `json:"bim_coin"`
	Team        int    `json:"team"`
	TeamLeader  bool   `json:"team_leader"`
}

var UserToID = map[string]string{
	"Учитель": "0",
}

var Users = map[string]*UserStruct{
	"0": {
		Name:        "Учитель",
		Password:    "BIM_LOCAL123",
		StatusEntry: false,
		Status:      TeacherStatus,
		Team:        0,
	},
}
