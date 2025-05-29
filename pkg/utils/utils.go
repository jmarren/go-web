package utils

func PanicE(e error) {
	if e != nil {
		panic(e)
	}
}
