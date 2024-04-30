package main

func uniquePathsIII(grid [][]int) int {
	cnt:=0
	l,r:=-1,-1
	for i:=0;i<len(grid);i++{
		for j:=0;j<len(grid[0]);j++{
			if grid[i][j]!=-1{
				cnt++
			}
			if grid[i][j]==1{
				grid[i][j]=0
				l,r=i,j
			}
		}
	}
	res:=0
	dfs(l,r,grid,cnt,&res,1)
	return res
}

func dfs(i,j int, grid [][]int,cnt int,res *int,c int){
	if i<0||j<0||i>=len(grid)||j>=len(grid[0])||grid[i][j]==-1{
		return
	}
	if grid[i][j]==2{
		if cnt==c{
			*res+=1
		}
		return
	}
	grid[i][j]=-1 // used
	dfs(i+1,j,grid,cnt,res,c+1)
	dfs(i-1,j,grid,cnt,res,c+1)
	dfs(i,j+1,grid,cnt,res,c+1)
	dfs(i,j-1,grid,cnt,res,c+1)
	grid[i][j]=0

}
