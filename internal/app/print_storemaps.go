package app

import (
	"net/http"
	"fmt"
)

// Печать ежедневника
func PrintStoremaps(w http.ResponseWriter, stormaps map[string]string) {
	for key, value := range stormaps {
		fmt.Fprintf(w, "%s : %v\n", key, value)
	}
}
