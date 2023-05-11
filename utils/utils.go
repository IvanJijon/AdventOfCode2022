package utils

func RemoveElementFromSlice[T comparable, L []*T](e *T, slice L) L {
	for k, v := range slice {
		if v == e {
			slice = append(slice[:k], slice[k+1:]...)
		}
	}
	return slice
}
