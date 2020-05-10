// Approach 2 BFS

func canVisitAllRooms(rooms [][]int) bool {
    if rooms == nil || len(rooms) == 0 {
        return false
    }
    
    foundKeys := make(map[int]bool)
    foundKeys[0] = true
    
    queue := list.New()
    queue.PushBack(0)
    
    for queue.Len() > 0 {
        key := queue.Front()
        queue.Remove(key)
        for _, k := range rooms[key.Value.(int)] {
            if _, ok := foundKeys[k]; ok {
                continue
            }
            foundKeys[k] = true
            queue.PushBack(k)
        }
    }
    return len(foundKeys) == len(rooms)
}

// Approach 1 DFS

// func canVisitAllRooms(rooms [][]int) bool {
//     if rooms == nil || len(rooms) == 0 {
//         return false
//     }
//     foundKeys := make(map[int]bool)
//     dfs(rooms, 0, foundKeys)
//     return len(foundKeys) == len(rooms)
// }

// func dfs(rooms [][]int, key int, foundKeys map[int]bool) {
//     foundKeys[key] = true
//     for _, k := range rooms[key] {
//         if _, ok := foundKeys[k]; !ok {
//             dfs(rooms, k, foundKeys)
//         }
//     }
// }