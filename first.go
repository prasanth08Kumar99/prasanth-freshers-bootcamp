package main
import (
	"fmt"
)
type Matrix struct {
	Rows int
	Columns int
	Elements [][]int
}
func (matrix *Matrix) GetRows() int {
	return matrix.Rows
}
func (matrix *Matrix) GetColumns() int {
	return matrix.Columns
}
func (matrix *Matrix) SetElement(i, j, x int) {
	matrix.Elements[i][j] = x
}
func (matrix *Matrix) AddMatrix(addMatrix *Matrix) {
	if (len(matrix.Elements) != len(addMatrix.Elements)) || (len(matrix.Elements[0]) != len(addMatrix.Elements[0])) {
		panic("Unsupported matrix addition")
	}
	for i:=0 ; i<len(matrix.Elements) ; i++ {
		for j:=0 ; j<len(matrix.Elements[0]) ; j++ {
			matrix.Elements[i][j] = matrix.Elements[i][j] + addMatrix.Elements[i][j]
		}
	}
}
func (matrix *Matrix) print() {
	fmt.Println(matrix.Elements)
}
func main() {
	rows, columns := 3, 3
	s := &Matrix{
		Rows:    rows,
		Columns:    columns,
		Elements: make([][]int, rows),
	}
	for i:=range s.Elements {
		s.Elements[i]=make([]int, columns)
	}
	s1 := &Matrix{
		Rows:    rows,
		Columns:    columns,
		Elements: make([][]int, rows),
	}
	for i:=range s1.Elements {
		s1.Elements[i]=make([]int, columns)
	}
	for i:=0 ; i<rows ; i++ {
		for j:=0; j<columns ; j++ {
			s.SetElement(i, j, i+j)
			s1.SetElement(i, j, i-j)
		}
	}
	fmt.Println("Number of row : ", s.GetRows())
	fmt.Println("Number of columns : ", s.GetColumns())
	s.print()
	s1.print()
	s.AddMatrix(s1)
	s.print()
}