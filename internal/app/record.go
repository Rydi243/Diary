package app

// Запись в ежедневник
func Record(date string, affairs string, stormaps map[string]string) {
	stormaps[date] = affairs
}
