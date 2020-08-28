func removeComments(source []string) []string {
    res := make([]string, 0)
    sb := strings.Builder{}
    inBlock := false
    
    for _, line := range source {
        if !inBlock {
            sb.Reset()
        }
        for i := 0; i < len(line); i++ {
            if !inBlock && i < len(line)-1 && line[i] == '/' && line[i+1] == '*' {
                inBlock = true
                i++
            } else if inBlock && i < len(line)-1 && line[i] == '*' && line[i+1] == '/' {
                inBlock = false
                i++
            } else if !inBlock && i < len(line)-1 && line[i] == '/' && line[i+1] == '/' {
                break
            } else if !inBlock {
                sb.WriteByte(line[i])
            }
        }
        if !inBlock && sb.Len() > 0 {
            res = append(res, sb.String())
        }
    }
    return res
}
