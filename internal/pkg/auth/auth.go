package auth


// Аутентификация
func AuthCheck(login string, password string) bool {
	return login == "admin" && password == "admin"
}
