package stringManipulation

func returnName() string {
	return "Vivek Kumar Verma"
}

func GiveName() string {
	return returnName()
}

func Concate(str1 string, str2 string) string {
	return str1 + str2
}

func ReveseString(str string) string {
	runeString := []rune(str)
	for i, j := 0, len(runeString)-1; i < j; i, j = i+1, j-1 {
		runeString[i], runeString[j] = runeString[j], runeString[i]
	}
	return string(runeString)
}
