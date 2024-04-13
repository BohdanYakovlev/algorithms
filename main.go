package main

import "fmt"

func binarySearch(mas []int, data int, left int, right int) int {
	if len(mas) == 0 {
		return -2
	}
	if left == right {
		if mas[left] == data {
			return left
		}
		return -1
	}
	if mas[(left+right)/2] < data {
		return binarySearch(mas, data, (left+right)/2+1, right)
	}
	if mas[(left+right)/2] > data {
		return binarySearch(mas, data, left, (left+right)/2-1)
	}
	return (left + right) / 2
}

func dfs(matrix [][]int, masIndex []int, point int) {
	masIndex[point] = 1
	for i := 0; i < len(masIndex); i++ {
		if matrix[i][point] == 1 && masIndex[i] == 0 {
			fmt.Println(i)
			dfs(matrix, masIndex, i)
		}
	}
}

func bfs(matrix [][]int, masIndex []int, point int) {
	var temp = []int{point}
	masIndex[point] = 1
	for i := 0; i < len(temp); i++ {
		for j := 0; j < len(masIndex); j++ {
			if matrix[j][temp[i]] == 1 && masIndex[j] == 0 {
				masIndex[j] = 1
				fmt.Println(j)
				temp = append(temp, j)
			}
		}
	}
}

func main() {

}
