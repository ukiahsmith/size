package size

const (
	KiB int64 = 1 << ((iota + 1) * 10) // Kibibyte
	MiB                                // Mebibyte
	GiB                                // Gibibyte
	TiB                                // Tebibyte
	PiB                                // Pebibyte
	EiB                                // Exbiyte
)
