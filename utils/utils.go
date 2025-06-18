package utils

import (
	"contest/constants"
)

func GetUserData(key string) *constants.UserStruct {
	user, ok := constants.Users[key]
	if !ok {
		return nil
	}

	return user
}