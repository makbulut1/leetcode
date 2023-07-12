# Intuition
<!-- Describe your first thoughts on how to solve this problem. -->
The code aims to find the shortest path in a binary matrix from the top-left corner to the bottom-right corner.

# Approach
<!-- Describe your approach to solving the problem. -->
1. Check if the top-left or bottom-right cell in the grid is blocked (marked as 1). If either of them is blocked, return -1, indicating that there is no valid path.
2. Initialize a queue to perform a breadth-first search (BFS) traversal.
3. Add the top-left cell (0, 0) to the queue as the starting point.
4. Define the eight possible directions to explore (diagonals included) using the directions array.
5. Mark the top-left cell as visited by setting its value to 1.
6. While the queue is not empty, perform the following steps:
- Dequeue the front element from the queue and assign its coordinates to row and col.
- Check if the dequeued cell is the bottom-right cell. If it is, return the value of that cell as it represents the shortest path distance.
- Iterate through all possible directions:
    - Calculate the coordinates of the new cell by adding the direction offsets to the current cell's coordinates.
    - Check if the new cell is within the grid boundaries and is unvisited (cell value is 0).
- If the conditions are met, enqueue the new cell and mark it as visited by updating its value to the distance of the current cell plus one.
7. If the queue is empty and the bottom-right cell was not reached, it means there is no valid path. Return -1.
# Complexity
- Time complexity: The code has a time complexity of O(N^2), where N represents the length of the grid's side. In the worst case, it may visit all cells in the grid.
<!-- Add your time complexity here, e.g. $$O(n)$$ -->

- Space complexity: The code has a space complexity of O(N^2) because the queue can store at most all the cells in the grid in the worst case.
<!-- Add your space complexity here, e.g. $$O(n)$$ -->

# Code
```
class Solution {
public:
    int shortestPathBinaryMatrix(vector<vector<int>>& grid) {
        int n = grid.size();
    if (grid[0][0] == 1 || grid[n-1][n-1] == 1) {
        return -1;
    }

    vector<vector<int>> directions = {{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}};
    grid[0][0] = 1;

    queue<vector<int>> q;
    q.push({0, 0});

    while (!q.empty()) {
        vector<int> curr = q.front();
        q.pop();
        int row = curr[0];
        int col = curr[1];

        if (row == n-1 && col == n-1) {
            return grid[row][col];
        }

        for (auto dir : directions) {
            int newRow = row + dir[0];
            int newCol = col + dir[1];
            if (newRow >= 0 && newRow < n && newCol >= 0 && newCol < n && grid[newRow][newCol] == 0) {
                q.push({newRow, newCol});
                grid[newRow][newCol] = grid[row][col] + 1;
            }
        }
    }

    return -1;
    }
};

```
