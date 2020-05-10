func canVisitAllRooms(rooms [][]int) bool {
    if rooms == nil || len(rooms) == 0 {
        return false
    }
    foundKeys := make(map[int]bool)
    dfs(rooms, 0, foundKeys)
    return len(foundKeys) == len(rooms)
}

func dfs(rooms [][]int, key int, foundKeys map[int]bool) {
    foundKeys[key] = true
    for _, k := range rooms[key] {
        if _, ok := foundKeys[k]; !ok {
            dfs(rooms, k, foundKeys)
        }
    }
}

