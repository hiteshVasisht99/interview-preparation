package basicdsa

func StringToInt(str string) int {
	result := 0
	for i := 0; i < len(str); i++ {
		bytes := str[i]
		result = result*10 + int(bytes-'0')
	}
	return result
}
