import "math/rand"

func kClosest(points [][]int, K int) [][]int {
    res := make([][]int, 0)
    if len(points) == 0 || K == 0 {
        return res
    }
    sort(&points, 0, len(points)-1, K)
    return points[:K]
}

func sort(points *[][]int, i, j, K int) {
    if i >= j {
        return
    }
    pi := i + rand.Intn(j-i)
    swap(points, i, pi)
    
    mid := partition(points, i, j)
    leftLength := mid - i + 1
    if K < leftLength {
        sort(points, i, mid-1, K)
    } else if K > leftLength {
        sort(points, mid+1, j, K-leftLength)
    }
}

func partition(points *[][]int, i, j int) int {
    oi := i
    pivot := distanceToOrigin((*points)[i])
    i++
    
    for {
        for i < j && distanceToOrigin((*points)[i]) <= pivot {
            i++
        }
        for i <= j && distanceToOrigin((*points)[j]) >= pivot {
            j--
        }
        if i >= j {
            break
        }
        swap(points, i, j)
    }
    swap(points, oi, j)
    return j
}

func swap(points *[][]int, i, j int) {
    (*points)[i][0], (*points)[i][1], (*points)[j][0], (*points)[j][1] = (*points)[j][0], (*points)[j][1], (*points)[i][0], (*points)[i][1]
}

func distanceToOrigin(point []int) int {
    return point[0]*point[0] + point[1]*point[1]
}
