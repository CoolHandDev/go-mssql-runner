package config

//MssqlCn represents the connection string
type MssqlCn struct {
	UserName  string
	Password  string
	Server    string
	Database  string
	Port      string
	CnTimeout string
	AppName   string
}

//GetCnString returns a sql server connection string
func GetCnString(c MssqlCn) string {
	return "user id=" + c.UserName +
		";password=" + c.Password +
		";server=" + c.Server +
		";database=" + c.Database +
		";port=" + c.Port +
		";connection timeout=" + c.CnTimeout +
		";app name=" + c.AppName

}
