package day5

func isVectorHorizontalOrVertical(vector Vector) bool {
	result := vector.start.x == vector.end.x || vector.start.y == vector.end.y
	return result
}

func removeDiagonals(input []Vector) []Vector {
	var result []Vector
	for _, currentVector := range input {
		if isVectorHorizontalOrVertical(currentVector) {
			result = append(result, currentVector)
		}
	}
	return result
}
