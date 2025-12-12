package mathutil

type number interface {
	~int | ~int8 | ~int16 | ~int32 | ~float32 | ~float64 | ~byte
}

func Add[T number](a, b T) T {
	return a + b
}
