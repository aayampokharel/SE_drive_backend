package functions

func StringToBool(str string) bool {
	switch str {
	case "yes", "1", "true", "t", "y":
		return true
	default:
		return false
	}

}
