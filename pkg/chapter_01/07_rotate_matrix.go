package chapter_01

// RotateMatrix Given an image represented by an NxN matrix, where each pixel in the image is 4 bytes, write a method
// to rotate the image by 90 degrees. Can you do this in place?
func RotateMatrix(matrix *[][]int) {
	l := len(*matrix)
	for i := 0; i < l/2; i++ {
		for j := i; j < l-i-1; j++ {
			temp := (*matrix)[i][j]
			(*matrix)[i][j] = (*matrix)[l-1-j][i]
			(*matrix)[l-1-j][i] = (*matrix)[l-1-i][l-1-j]
			(*matrix)[l-1-i][l-1-j] = (*matrix)[j][l-1-i]
			(*matrix)[j][l-1-i] = temp
		}
	}
}
