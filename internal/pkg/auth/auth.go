package auth

import (
	"Diary/internal/config"
)

// Аутентификация
func AuthCheck(login string, password string) bool {
	return login == config.AdminLog && password == config.AdminPas
}
