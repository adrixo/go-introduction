package practice

func getChunkSize(slice []int) (offset int, chunkSize int) {
	return len(slice) % 4, len(slice) / 4
}
