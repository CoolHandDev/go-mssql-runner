package config

//TODO: Create a connection struct

//GetCnString returns a sql server connection string
func GetCnString(u, p, s, d string) string {
	return "user id=" + u +
		";password=" + p +
		";server=" + s +
		";database=" + d
}
