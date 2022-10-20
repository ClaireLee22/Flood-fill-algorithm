from collections import deque
class Solution(object):
    def floodFill(self, image, sr, sc, newColor):
        """
        :type image: List[List[int]]
        :type sr: int
        :type sc: int
        :type newColor: int
        :rtype: List[List[int]]
        """
        oldColor = image[sr][sc]
        if oldColor == newColor:
            return image
        
        visited = set([(sr, sc)])
        queue = deque([(sr, sc)])

        directions = [(-1, 0), (0, 1), (1, 0), (0, -1)]
        while len(queue) > 0:
            print("queue", queue)
            currentRow, currentCol = queue.popleft()
            image[currentRow][currentCol] = newColor
            for rowOffset, colOffset in directions:
                neighborRow = currentRow + rowOffset
                neighborCol = currentCol + colOffset
                pos = (neighborRow, neighborCol)
                if self.isInBounds(image, neighborRow, neighborCol) and pos not in visited and image[neighborRow][neighborCol] == oldColor:
                    visited.add(pos)
                    queue.append((neighborRow, neighborCol))
        
        return image
    
    def isInBounds(self, image, row, col):
        return 0 <= row < len(image) and 0 <= col < len(image[0])
        

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
('queue', deque([(2, 2)]))
('queue', deque([(1, 2), (2, 1)]))
('queue', deque([(2, 1), (0, 2), (1, 3)]))
('queue', deque([(0, 2), (1, 3), (3, 1), (2, 0)]))
('queue', deque([(1, 3), (3, 1), (2, 0), (0, 1)]))
('queue', deque([(3, 1), (2, 0), (0, 1)]))
('queue', deque([(2, 0), (0, 1), (3, 0)]))
('queue', deque([(0, 1), (3, 0), (1, 0)]))
('queue', deque([(3, 0), (1, 0), (0, 0)]))
('queue', deque([(1, 0), (0, 0)]))
('queue', deque([(0, 0)]))
('flood fill result', [[2, 2, 2, 0, 1], [2, 0, 2, 2, 0], [2, 2, 2, 0, 1], [2, 2, 0, 1, 1]])
"""
