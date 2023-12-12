package models

import "github.com/gofrs/uuid"

func GetStringFromInterface(v interface{}) string {
	if v == nil {
		return ""
	}

	return v.(string)
}

//Generates new uuid
func GenerateUUID() uuid.UUID {
	uid, _ := uuid.NewV4()
	return uid
}
