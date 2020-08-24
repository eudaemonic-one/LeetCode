func kClosest(points [][]int, K int) [][]int {
    if len(points) == 0 || K == 0 {
        return make([][]int, 0)
    }
    mySort(points, 0, len(points)-1, K)
    return points[:K]
}

func mySort(points [][]int, lo, hi, K int) {
    if lo >= hi {
        return
    }
    mid := partition(points, lo, hi)
    leftLength := mid - lo + 1
    
    if K < leftLength {
        mySort(points, lo, mid-1, K)
    } else if K > leftLength {
        mySort(points, mid+1, hi, K-leftLength)
    }
}

func partition(points [][]int, i, j int) int {
    randIdx := i + rand.Intn(j-i)
    swap(points, i, randIdx)
    pivot := i
    pivotDist := distanceToOrigin(points[pivot])
    i++
    
    for {
        for i < j && distanceToOrigin(points[i]) <= pivotDist {
            i++
        }
        for i <= j && distanceToOrigin(points[j]) >= pivotDist {
            j--
        }
        if i >= j {
            break
        }
        swap(points, i, j)
    }
    
    swap(points, pivot, j)
    
    return j
}

func distanceToOrigin(point []int) int {
    return point[0]*point[0] + point[1]*point[1]
}

func swap(points [][]int, i, j int) {
    points[i][0], points[i][1], points[j][0], points[j][1] = points[j][0], points[j][1], points[i][0], points[i][1]
}

