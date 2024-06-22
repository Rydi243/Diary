package app

import (
	"fmt"
	"net/http"
	"Diary/internal/config"
)

// Приветствие админа и печать ежедневника всего
func HelloAdminFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello, admin!")
	PrintStoremaps(w, config.Stormaps)

}