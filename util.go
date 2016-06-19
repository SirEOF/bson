package bson

func assert(a bool, msg string) {
	if !a {
		panic(msg)
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
