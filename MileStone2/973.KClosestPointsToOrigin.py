import heapq
class Solution:
    def kClosest(self, points: List[List[int]], k: int) -> List[List[int]]:
        """
        :type points: List[List[int]]
        :type k: int
        :rtype: List[List[int]]
        """
        heap = []
        for(x,y) in points:
            dist = -(x*x  + y*y)
            if len(heap) == k:
                heapq.heappushpop(heap, (dist, x ,y))
            else: 
                heapq.heappush(heap, (dist, x,y))

        result = [(x,y) for (dist, x,y) in heap]
        return result 