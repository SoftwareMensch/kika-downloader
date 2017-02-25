package utils

import "regexp"

// IsUUID4 validate if string is a valid uuid version 4
func IsUUID4(in string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")

	return r.MatchString(in)
}
