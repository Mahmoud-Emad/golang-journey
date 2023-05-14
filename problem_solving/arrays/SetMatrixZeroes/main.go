package main

import "fmt"

func setZeroes(matrix [][]int) {
	m := len(matrix)    // O(1)
	n := len(matrix[0]) // O(1)

	rows := make([]bool, m) // O(m)
	cols := make([]bool, n) // O(n)

	for i := 0; i < m; i++ { // O(m)
		for j := 0; j < n; j++ { // O(n)
			if matrix[i][j] == 0 { // O(1)
				rows[i] = true // O(1)
				cols[j] = true // O(1)
			}
		}
		// O(n)
	} // O(m * n)

	for i := 0; i < m; i++ { // O(m)
		for j := 0; j < n; j++ { // O(n)
			if rows[i] || cols[j] { // O(1)
				matrix[i][j] = 0 // O(1)
			}
		}
	} // O(m * n) => n2s

	// Space O(m+n)
	// Time O(m*n)

	// Time => O(m*n + m*n + m + n + 1)
}

func main() {
	matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(matrix)
	fmt.Println(matrix)
	matrix = [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(matrix)
	fmt.Println(matrix)
}

/* Input: matrix = [
	[1, 1, 1, 1]
	[1, 0, 1, 0]
	[1, 1, 1, 0]
]
*/

/*
i,j = 0,0
i,j = 0,1
i,j = 0,2
i,j = 1,0
i,j = 1,1  => [[1,0,1], [0,0,0], [1,0,1]]
i,j = 1,2 => [[1,0,0], [0,0,0], [1,0,0]] => Wrong solution, this'll set all the matrix = 0 based on the first 0 index

Input: matrix = [[1,1,1],[1,0,1],[1,1,1]]

for i := 0; i < m; i++ { O(m)
for j := 0; j < n; j++ { O(n)
if matrix[i][j] == 0 {
Set the row to zero  | n elemnts = O(n) => columns
Set the column to zero | m elemnts = O(m) => rows
}
O(m+n)
} O((m+n) * n)
} // O(m * n)

m = 5
n = 5

O(m+n) = 11
O((m+n) * n)

O((m + n) * (n * m))
O((m + m) * m2)
*/
