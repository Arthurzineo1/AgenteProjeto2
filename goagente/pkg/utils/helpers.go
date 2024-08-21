package utils

func BytesToGigabytes(bytes uint64) float64 {
	const bytesInGigabyte = 1024 * 1024 * 1024
	return float64(bytes) / bytesInGigabyte
}
