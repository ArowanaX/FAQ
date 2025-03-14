package validator

func IsValidString(st string) bool { return len(st) > 3 && len(st) < 255 }
