package iteration

// Repeat task a char intput and repeat times,
// then return repeated several times char
func Repeat(char string, times int) (result string) {
	for i := 0; i < times; i++ {
		result += char
	}
	return
}
