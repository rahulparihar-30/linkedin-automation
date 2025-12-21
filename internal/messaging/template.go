package messaging

import (
	"strings"
)

func RenderTemplate(template string, firstName string) string {
	res := strings.ReplaceAll(template, "{{firstName}}", firstName)
	return res
}
