package shared

func Assert(condition bool, msg string) {
	if !condition {
		panic("Assertion failed: " + msg)
	}
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
