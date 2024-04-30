package main

func rotate(matrix [][]int) {
	matrixLen := len(matrix)
	i, j := 0, matrixLen-1
	for i < j {
		clen := len(matrix[0])
		for i2 := 0; i2 < clen; i2++ {
			matrix[i][i2], matrix[j][i2] = matrix[j][i2], matrix[i][i2]
		}
		i++
		j--
	}

	for i = 1; i < matrixLen; i++ {
		for j = 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
