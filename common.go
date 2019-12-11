package gohtparser

func GetName(s []string) (string, []string) {
	var slice []string
	if len(s) == 0 {
		return "", slice
	}
	return s[0], s[1:]
}
