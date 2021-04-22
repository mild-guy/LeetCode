/*
554. Brick Wall
There is a brick wall in front of you. The wall is rectangular and has several rows of bricks. The bricks have the same height but different width. You want to draw a vertical line from the top to the bottom and cross the least bricks.
The brick wall is represented by a list of rows. Each row is a list of integers representing the width of each brick in this row from left to right.
If your line go through the edge of a brick, then the brick is not considered as crossed. You need to find out how to draw the line to cross the least bricks and return the number of crossed bricks.
You cannot draw a line just along one of the two vertical edges of the wall, in which case the line will obviously cross no bricks.
*/
func LeastBricks(wall [][]int) int {
	ret := 0
	hashTable := map[int]int{}

	for i := range wall {
		curWidth := 0
		for j := 0; j < len(wall[i])-1; j++ {
			curWidth += wall[i][j]
			hashTable[curWidth]++
		}
	}
	for _, v := range hashTable {
		if v > ret {
			ret = v
		}
	}
	return len(wall) - ret
}

/*
Success
Runtime: 28 ms, faster than 60.87% of Go online submissions for Brick Wall.
Memory Usage: 7.1 MB, less than 65.22% of Go online submissions for Brick Wall.
*/
