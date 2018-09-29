package allocate

func AllocateSlice(size, loop int) {
	slice := make([]string, size)
	value := "test"
	for count := 1; count < loop; count++ {
		slice = append(slice, value)
	}
}
