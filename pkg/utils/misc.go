package utils

//ParseLinuxTemp - Convert linux temp variables into degrees
func ParseLinuxTemp(linuxTemp int) int {
	return linuxTemp / 1000
}
