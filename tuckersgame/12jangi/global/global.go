package global

const (
	ScreenWidth  = 480
	ScreenHeight = 362
)

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
