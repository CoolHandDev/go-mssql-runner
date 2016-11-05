package progress

import "time"

//Prefix applies a basic time prefix to a message
func Prefix(s string) string {
	return time.Now().Format("2006-01-02 T15:04:05") + ":" + s
}
