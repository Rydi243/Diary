package app



// Запись в ежедневник
func Record(date string, affairs string, stormaps map[string]string) {
	stormaps[date] = affairs
	//fmt.Println(time.Now(), "Новая запись в ежедневнике =>", date, affairs)
}