package bytesize

const (
	Kibibyte int64 = 1 << ((iota + 1) * 10)
	Mebibyte
	Gibibyte
	Tebibyte
	Pebibyte
	Exbiyte
)
