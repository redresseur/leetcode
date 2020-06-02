package dtgh

/*
已知一个 NxN 的国际象棋棋盘，棋盘的行号和列号都是从 0 开始。即最左上角的格子记为 (0, 0)，最右下角的记为 (N-1, N-1)。

现有一个 “马”（也译作 “骑士”）位于 (r, c) ，并打算进行 K 次移动。

如下图所示，国际象棋的 “马” 每一步先沿水平或垂直方向移动 2 个格子，然后向与之相垂直的方向再移动 1 个格子，共有 8 个可选的位置。

现在 “马” 每一步都从可选的位置（包括棋盘外部的）中独立随机地选择一个进行移动，直到移动了 K 次或跳到了棋盘外面。

求移动结束后，“马” 仍留在棋盘上的概率。

输入: 3, 2, 0, 0
输出: 0.0625
解释:
输入的数据依次为 N, K, r, c
第 1 步时，有且只有 2 种走法令 “马” 可以留在棋盘上（跳到（1,2）或（2,1））。对于以上的两种情况，各自在第2步均有且只有2种走法令 “马” 仍然留在棋盘上。
所以 “马” 在结束后仍在棋盘上的概率为 0.0625。

*/

func knightProbability1(N int, K int, r int, c int, dp [][]map[int]float64) float64 {
	next := [][2]int{{1, 2}, {2, 1}, {1, -2}, {-2, 1},
		{-1, 2}, {2, -1}, {-2, -1}, {-1, -2}}
	res := float64(0)
	if K <= 0 {
		if r <= N-1 && r >= 0 && c <= N-1 && c >= 0 {
			res = 1
		}
		return res
	}

	for _, l := range next {
		r1, c1 := r+l[0], c+l[1]
		if r1 > N-1 || r1 < 0 || c1 > N-1 || c1 < 0 {
			continue
		}
		if exist, ok := dp[r1][c1][K-1]; ok {
			//fmt.Printf("%d %d #%d %f ", r1, c1, K-1, dp[r1][c1][K-1])
			res += exist
		} else {
			exist := knightProbability1(N, K-1, r1, c1, dp)
			dp[r1][c1][K-1] = exist
			res += exist
		}
	}

	return res / 8
}

// 时间复杂度: O(k*N^2); 空间复杂度 O(N^2)
func knightProbability(N int, K int, r int, c int) float64 {
	dp := make([][]map[int]float64, N)
	for i, _ := range dp {
		dp[i] = make([]map[int]float64, N)
		for j, _ := range dp[i] {
			// dp[i][j] = make(map[int]float64, K)
			dp[i][j] = map[int]float64{}
		}
	}

	res := knightProbability1(N, K, r, c, dp)
	return res
}

func newBatch(N int) [][]float64 {
	dp := make([][]float64, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]float64, N)
	}

	return dp
}

// 时间复杂度: O(k*N^2); 空间复杂度 O(N^2)
func knightProbabilityIteration(N int, K int, sr int, sc int) float64 {
	dp := newBatch(N)
	next := [][2]int{{1, 2}, {2, 1}, {1, -2}, {-2, 1},
		{-1, 2}, {2, -1}, {-2, -1}, {-1, -2}}

	if N <= 0 || K <= 0 || sr < 0 || sr > N-1 || sc < 0 || sc > N-1 {
		return 0
	}

	dp[sr][sc] = 1

	for ; K > 0; K-- {
		dp1 := newBatch(N)
		// 枚举所有位置成为下一步落脚点的概率
		for c := 0; c < N; c++ {
			for r := 0; r < N; r++ {
				for _, d := range next{
					dr, dc := r + d[0], c + d[1]
					if  N - 1 < dr || dr < 0 || dc < 0 || dc > N - 1{
						continue
					}

					dp1[dr][dc] += dp[r][c]/8
				}
			}
		}
		dp = dp1
	}

	ans := float64(0)
	for _, row := range dp{
		for _, a := range row{
			ans += a
		}
	}

	return ans
}
