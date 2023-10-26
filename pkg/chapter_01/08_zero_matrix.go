package chapter_01

// ZeroMatrix Write an algorithm such that if an element in an MxN matrix is 0, its entire row and column are set to 0
func ZeroMatrix(matrix [][]int) [][]int {
	x, y := 0, 0
	found := false
	for i, v1 := range matrix {
		for j, v2 := range v1 {
			if v2 == 0 {
				found = true
				x, y = i, j
				break
			}
		}
	}

	if found {
		for i := 0; i < len(matrix); i++ {
			matrix[i][y] = 0
			if i == x {
				for k := range matrix[i] {
					matrix[i][k] = 0
				}
			}
		}
	}

	return matrix
}
