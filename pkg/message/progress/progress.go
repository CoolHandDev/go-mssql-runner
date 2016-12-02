package progress

import (
	"strings"
	"time"
)

//Prefix applies a basic time prefix to a message
func Prefix(msg ...string) string {
	return time.Now().Format("2006-01-02 15:04:05") + "| " + strings.Join(msg, " ")
}
