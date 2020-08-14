func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
    if len(image) == 0 || len(image[0]) == 0 {
        return image
    }
    dfs(image, sr, sc, image[sr][sc], newColor)
    return image
}

func dfs(image [][]int, r, c int, oldColor, newColor int) {
    if r < 0 || r >= len(image) || c < 0 || c >= len(image[0]) || image[r][c] != oldColor {
        return
    }
    image[r][c] = -1
    dfs(image, r, c-1, oldColor, newColor)
    dfs(image, r, c+1, oldColor, newColor)
    dfs(image, r-1, c, oldColor, newColor)
    dfs(image, r+1, c, oldColor, newColor)
    image[r][c] = newColor
}
