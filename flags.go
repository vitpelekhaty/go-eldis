package eldis

// Flag тип флага выполнения методов API
type Flag byte

const (
	// CompressedResponse использовать сжатие тела ответа API
	CompressedResponse Flag = iota
	// UseCompressedResponseFlagInHeader использовать параметр compressed-response в заголовке запроса
	UseCompressedResponseFlagInHeader
)

func flagExists(flag Flag, flags ...Flag) bool {
	if len(flags) > 0 {
		for _, f := range flags {
			if f == flag {
				return true
			}
		}
	}

	return false
}
