func longestConsecutive(nums []int) int {
    parent = make(map[int]int)
    size = make(map[int]int)
    for _, num := range nums {
        if _, ok := parent[num]; !ok {
            parent[num] = num
            size[num] = 1
        }
        if _, ok := parent[num-1]; ok {
            union(parent, size, num, num-1)
        }
        if _, ok := parent[num+1]; ok {
            union(parent, size, num, num+1)
        }
    }
    res := 0
    for _, x := range size {
        if x > res {
            res = x
        }
    }
    return res
}

var parent map[int]int
var size map[int]int

func find(parent map[int]int, x int) int {
    for x != parent[x] {
        parent[x] = parent[parent[x]]
        x = parent[x]
    }
    return x
}

func union(parent, size map[int]int, x, y int) {
    rootX := find(parent, x)
    rootY := find(parent, y)
    if rootX != rootY {
        if size[rootX] < size[rootY] {
            parent[rootX] = rootY
            size[rootY] += size[rootX]
        } else {
            parent[rootY] = rootX
            size[rootX] += size[rootY]
        }
    }
}

