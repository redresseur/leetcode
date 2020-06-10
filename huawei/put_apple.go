package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
题目描述
把M个同样的苹果放在N个同样的盘子里，允许有的盘子空着不放，问共有多少种不同的分法？（用K表示）5，1，1和1，5，1 是同一种分法。

输入
每个用例包含二个整数M和N。0<=m<=10，1<=n<=10。

样例输入
7 3

样例输出
8
*/

func dp(m, n int) [][]int {
	matrix := make([][]int, m+1)
	for i, _ := range matrix {
		matrix[i] = make([]int, n+1)
	}
	return matrix
}

func Count(m, n int) int {
	dp := dp(m, n)
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if j == 1 || i <= 1 { // 只有一个苹果或 0 个苹果则只有一种摆法，只有一个果篮也只有一种摆法
				dp[i][j] = 1
				continue
			}

			if j-1 >= 1 {
				if i-j >= 0 {
					dp[i][j] = dp[i][j-1] + dp[i-j][j] // count(m,n) = count(m, n - 1) + count(m - n, n)
				} else {
					dp[i][j] = dp[i][i]
				}
			}
		}
	}

	return dp[m][n]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		ns, _, err := reader.ReadLine()
		if err != nil {
			return
		}

		arrs := strings.Split(string(ns), " ")
		if len(arrs) >= 2 {
			m, err := strconv.Atoi(arrs[0])
			if err != nil {
				continue
			}

			n, err := strconv.Atoi(arrs[1])
			if err != nil {
				continue
			}

			fmt.Println(Count(m, n))
		}
	}
}
