package main

import "fmt"

func main() {

	fmt.Println("No. 6")
	m5 := matrixNo6()
	displayMatrix(m5)

	fmt.Println("\nNo. 7")
	m5 = matrixNo7()
	displayMatrix(m5)

	fmt.Println("\nNo. 8")
	n := 7
	m7 := matrixNo8(n)
	displayMatrixNo8(m7)

	fmt.Println("\nNo. 9")
	m8 := matrixNo9()
	displayMatrix(m8)
	fmt.Println("\nNo. 10")
	matrixNo10()
}

func matrixNo6() [][]int {
	matrix := make([][]int, 5)
	for i := 0; i < 5; i++ {
		matrix[i] = make([]int, 5)
		for j := 0; j < 5; j++ {
			if i == j {
				matrix[i][j] = i + 1
			} else if j > i {
				matrix[i][j] = 10
			} else {
				matrix[i][j] = 20
			}
		}
	}
	return matrix
}

func matrixNo7() [][]int {
	matrix := make([][]int, 5)
	for i := 0; i < 5; i++ {
		matrix[i] = make([]int, 5)
		for j := 0; j < 5; j++ {
			if i == j {
				matrix[i][j] = 5 - i
			} else if j > i {
				matrix[i][j] = 20
			} else {
				matrix[i][j] = 10
			}
		}
	}
	return matrix
}

func matrixNo8(n int) [][]int {
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 || i == n-1 || j == n-1 {
				matrix[i][j] = i + j
			}
		}
	}
	return matrix
}

func displayMatrixNo8(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i > 0 && i < n-1 && j > 0 && j < n-1 {
				fmt.Printf("    ")
			} else {
				fmt.Printf("%4d", matrix[i][j])
			}
		}
		fmt.Println()
	}
}

func matrixNo9() [][]int {
	matrix := make([][]int, 8)
	for i := 0; i < 8; i++ {
		matrix[i] = make([]int, 8)
	}

	for i := 0; i < 7; i++ {
		for j := 0; j < 7; j++ {
			matrix[i][j] = i + j
			matrix[i][7] += matrix[i][j]
			matrix[7][j] += matrix[i][j]
			matrix[7][7] += matrix[i][j]
		}
	}
	return matrix
}

func displayMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%4d", matrix[i][j])
		}
		fmt.Println()
	}
}
func matrixNo10() {
	jawaban := [][]string{
		{"A", "B", "A", "C", "C", "D", "E", "E", "A", "D"}, //Student-0
		{"D", "B", "A", "B", "C", "A", "E", "E", "A", "D"}, //Student-1
		{"E", "D", "D", "A", "C", "B", "E", "E", "A", "D"}, //Student-2
		{"C", "B", "A", "E", "D", "C", "E", "E", "A", "D"}, //Student-3
		{"A", "B", "D", "C", "C", "D", "E", "E", "A", "D"}, //Student-4
		{"B", "B", "E", "C", "C", "D", "E", "E", "A", "D"}, //Student-5
		{"B", "B", "A", "C", "C", "D", "E", "E", "A", "D"}, //Student-6
		{"E", "B", "E", "C", "C", "D", "E", "E", "A", "D"}, //Student-7
	}
	kunci := []string{"D", "B", "D", "C", "C", "D", "A", "E", "A", "D"}
	for i := 0; i < len(jawaban); i++ {
		benar := 0
		for j := 0; j < len(kunci); j++ {
			if jawaban[i][j] == kunci[j] {
				benar++
			}
		}
		fmt.Printf("Jawaban Siswa %d yang benar : %d\n", i, benar)
	}
}
