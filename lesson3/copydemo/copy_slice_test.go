package copydemo

import "testing"

func testGetLast(t *testing.T, f func([]int) []int) {
	result := make([][]int, 0)
	for k := 0; k < 100; k++ {
		origin := generateWithCap(128 * 1024)
		result = append(result, f(origin))
	}
	printMem(t)
	_ = result
}

func TestGetLastBySlice(t *testing.T) {
	testGetLast(t, GetLastBySlice)
}

func TestGetLastByCopy(t *testing.T) {
	testGetLast(t, GetLastByCopy)
}
