package regex

import "regexp"

func Match() {
	u :=
	if match, err := regexp.MatchString(u, url); err == nil && match {
		return true
	}
}
