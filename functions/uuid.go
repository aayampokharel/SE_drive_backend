package functions

import "github.com/google/uuid"

func GenerateUUID() string { //@returns a new and unique UUID
	return uuid.New().String()
	//returning format : 8 digits - 4 -4 -4 -12
}
