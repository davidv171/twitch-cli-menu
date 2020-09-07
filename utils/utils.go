package utils

// Return possible video heights, based on maximum height given
// We hardcode the heights because we KISS
func Qualities(maxq int) []int {
	// All possible qualities, twitch doesn't yet support 1440p and 2160, we future proof
	allq := []int{160, 360, 480, 720, 1080, 1440, 2160}
	//Available qualities, to offer the user
	avq := make([]int, 0)
	for _, v := range allq {
		if maxq >= v {
			avq = append(avq, v)
		}
	}
	return avq
}

