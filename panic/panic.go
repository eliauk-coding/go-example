package main

func main() {
	triggerPanic(make([]string, 1, 2), "煎鱼", 3)
}

func triggerPanic(slice []string, str string, i int) {
	panic("trigger panic")
}
