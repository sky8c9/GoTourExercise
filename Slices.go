//Exercise: Slices
package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	mat := make([][]uint8, dy)
	
	for i:= range mat {
		mat[i] = make([]uint8, dx)
		for j:=0;j<dx;j++ {
			mat[i][j] = uint8(i^j) 
		}
	}
	
	return mat
}

func main() {
	pic.Show(Pic)
}
