package slices

func expandSlice(slice []int, additions ...int) []int {
	neededCapacity := len(slice) + len(additions)

	if mustGrow := (neededCapacity > cap(slice)); mustGrow {
		grownSlice := make([]int, len(slice), (neededCapacity+1)*2)
		copy(grownSlice, slice)
		slice = grownSlice
	}

	sliceTail := slice[len(slice):cap(slice)]
	for i, v := range additions {
		sliceTail[i] = v
	}
	return slice[:neededCapacity]
}
