package utils

func AdjustKey(key []byte) []byte {
	switch {
	case len(key) > 32:
		return key[:32]
	case len(key) > 24:
		return key[:24]
	case len(key) > 16:
		return key[:16]
	default:
		padded := make([]byte, 16)
		copy(padded, key)
		return padded
	}
}
