from curses.ascii import SO


class Solution(object):
    def floodFill(self, image, sr, sc, newColor):
        """
        :type image: List[List[int]]
        :type sr: int
        :type sc: int
        :type newColor: int
        :rtype: List[List[int]]
        """
        visited = set()
        oldColor = image[sr][sc]
        self.replaceColor(image, sr, sc, oldColor, newColor, visited)
        
        return image
    
    def replaceColor(self, image, row, col, oldColor, newColor, visited):
        rowInBounds = 0 <= row < len(image)
        colInBounds = 0 <= col < len(image[0])
        if not rowInBounds or not colInBounds:
            return
        
        if image[row][col] != oldColor:
            return
        
        pos = (row, col)
        print("current position", pos)
        if pos in visited:
            return
        visited.add(pos)
        
        image[row][col] = newColor
        directions = [(-1, 0), (0, 1), (1, 0), (0, -1)]
        for rowOffset, colOffset in directions:
            neighborRow = row + rowOffset
            neighborCol = col + colOffset
            self.replaceColor(image, neighborRow, neighborCol, oldColor, newColor, visited)


if __name__ == "__main__":
    solution = Solution()
    image = [
        [1, 1, 1, 0, 1],
        [1, 0, 1, 1, 0],
        [1, 1, 1, 0, 1],
        [1, 1, 0, 1, 1],
    ]
    sr = 2
    sc = 2
    newColor = 2
    print("flood fill result", solution.floodFill(image, sr, sc, newColor))

"""
output:
('current position', (2, 2))
('current position', (1, 2))
('current position', (0, 2))
('current position', (0, 1))
('current position', (0, 0))
('current position', (1, 0))
('current position', (2, 0))
('current position', (2, 1))
('current position', (3, 1))
('current position', (3, 0))
('current position', (1, 3))
('flood fill result', [[2, 2, 2, 0, 1], [2, 0, 2, 2, 0], [2, 2, 2, 0, 1], [2, 2, 0, 1, 1]])
"""