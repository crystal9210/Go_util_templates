package someprocess

func Run(val int) int {
	return call(val)
}

var call = func(val int) int {
	return val
}
