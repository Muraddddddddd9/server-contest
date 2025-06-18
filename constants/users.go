package constants

type UserStruct struct {
	Password    string `json:"password"`
	Status      string `json:"status"`
	StatusEntry bool   `json:"status_entry"`
	BimCoin     uint64 `json:"bim_coin"`
	Team        uint8  `json:"team"`
	TeamLeader  bool   `json:"team_leader"`
}

var Users = map[string]*UserStruct{
	"Teacher": {
		Password:    "BIM_LOCAL123",
		StatusEntry: false,
		Status:      "Teacher",
		Team:        0,
	},
	"Student_1": {
		Password:    "123",
		StatusEntry: false,
		Status:      "Student",
		BimCoin:     0,
		Team:        1,
		TeamLeader:  true,
	},
	"Student_2": {
		Password:    "123",
		StatusEntry: false,
		Status:      "Student",
		BimCoin:     0,
		Team:        1,
		TeamLeader:  false,
	},
	"Student_3": {
		Password:    "123",
		StatusEntry: false,
		Status:      "Student",
		BimCoin:     0,
		Team:        1,
		TeamLeader:  false,
	},
	"Student_4": {
		Password:    "123",
		StatusEntry: false,
		Status:      "Student",
		BimCoin:     0,
		Team:        1,
		TeamLeader:  false,
	},
	"Student_5": {
		Password:    "123",
		StatusEntry: false,
		Status:      "Student",
		BimCoin:     0,
		Team:        1,
		TeamLeader:  false,
	},
	"Student_6": {
		Password:    "123",
		StatusEntry: false,
		Status:      "Student",
		BimCoin:     0,
		Team:        2,
		TeamLeader:  true,
	},
	"Student_7": {
		Password:    "123",
		StatusEntry: false,
		Status:      "Student",
		BimCoin:     0,
		Team:        2,
		TeamLeader:  false,
	},
	"Student_8": {
		Password:    "123",
		StatusEntry: false,
		Status:      "Student",
		BimCoin:     0,
		Team:        2,
		TeamLeader:  false,
	},
	"Student_9": {
		Password:    "123",
		StatusEntry: false,
		Status:      "Student",
		BimCoin:     0,
		Team:        2,
		TeamLeader:  false,
	},
	"Student_10": {
		Password:    "123",
		StatusEntry: false,
		Status:      "Student",
		BimCoin:     0,
		Team:        2,
		TeamLeader:  false,
	},
}
