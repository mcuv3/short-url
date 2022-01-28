package validation

import "regexp"

var (
	urlRegex, _ = regexp.Compile(`https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`)
)

func ValidateURL(url string) bool {
	return urlRegex.MatchString(url)
}
