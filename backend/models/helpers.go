package models

func GetStringFromInterface(v interface{}) string {
	if v == nil {
		return ""
	}

	return v.(string)
}
