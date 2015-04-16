package routes

import "strings"

// tokens = path splitted by ["/", "_"]
var tokenNormalizeLookup = map[string]string{
	"id":    "ID",
	"api":   "API",
	"oauth": "OAuth",
}

func RouteNameForPath(path string) string {
	params := []string{}
	names := []string{}

	path = strings.TrimSpace(path)
	if path == "/" {
		return "Root"
	}

	parts := strings.Split(path, "/")
	for _, token := range parts {
		token = strings.TrimSpace(token)
		if isParam(token) {
			params = append(params, cleanAndNormalize(token))
		} else {
			names = append(names, cleanAndNormalize(token))
		}
	}

	out := strings.Join(names, "")
	if len(params) > 0 {
		out += "For" + strings.Join(params, "And")
	}
	return out
}

func isParam(token string) bool {
	return strings.HasPrefix(token, ":")
}

func cleanAndNormalize(token string) string {
	cleaned := ""
	token = strings.Replace(token, ":", "", -1)

	parts := strings.Split(token, "_")
	for _, p := range parts {
		replacement, ok := tokenNormalizeLookup[p]
		if ok {
			cleaned += replacement
		} else {
			cleaned += strings.Title(p)
		}
	}

	return cleaned
}
