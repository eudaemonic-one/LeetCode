type DisjointSet struct {
    parent []int
    size []int
}

func NewDisjointSet(n int) DisjointSet {
    s := DisjointSet{make([]int, n), make([]int, n)}
    for i := 0; i < n; i++ {
        s.parent[i] = i
        s.size[i] = 1
    }
    return s
}

func (s DisjointSet) find(x int) int {
    for s.parent[x] != x {
        x = s.parent[x]
    }
    return x
}

func (s DisjointSet) union(x, y int) {
    rootX := s.find(x)
    rootY := s.find(y)
    if rootX == rootY {
        return
    }
    if s.size[rootX] < s.size[rootY] {
        s.parent[rootX] = rootY
        s.size[rootY] += s.size[rootX]
    } else {
        s.parent[rootY] = rootX
        s.size[rootX] += s.size[rootY]
    }
}

func solve(board [][]byte)  {
    m := len(board)
    if m == 0 {
        return
    }
    
    n := len(board[0])
    if n == 0 {
        return
    }
    
    dummy := m*n
    ds := NewDisjointSet(dummy+1)
    
    for i := 0; i < m; i++ {
        if board[i][0] == 'O' {
            ds.union(i*n, dummy)
        }
        if board[i][n-1] == 'O' {
            ds.union(i*n+n-1, dummy)
        }
    }
    
    for j := 1; j < n-1; j++ {
        if board[0][j] == 'O' {
            ds.union(j, dummy)
        }
        if board[m-1][j] == 'O' {
            ds.union((m-1)*n+j, dummy)
        }
    }
    
    for i := 1; i < m-1; i++ {
        for j := 1; j < n-1; j++ {
            if board[i][j] == 'O' {
                if board[i-1][j] == 'O' {
                    ds.union((i-1)*n+j, i*n+j)
                }
                if board[i+1][j] == 'O' {
                    ds.union((i+1)*n+j, i*n+j)
                }
                if board[i][j-1] == 'O' {
                    ds.union(i*n+j, i*n+j-1)
                }
                if board[i][j+1] == 'O' {
                    ds.union(i*n+j, i*n+j+1)
                }
            }
        }
    }
    
    for i := 1; i < m-1; i++ {
        for j := 1; j < n-1; j++ {
            if ds.find(dummy) != ds.find(i*n+j) {
                board[i][j] = 'X'
            }
        }
    }
}
