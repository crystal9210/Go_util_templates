package sample

// io.Writerと同じ;構造体ではなくインターフェース
type writer interface {
	Write([]byte) (int, error)
}
