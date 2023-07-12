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
