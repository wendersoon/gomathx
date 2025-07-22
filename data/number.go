package data

// Restrictive generic interface for generic types
type Number interface {
	~int | ~float64 | ~int32 | ~int64 | ~int8 | ~int16 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32
}
