type Pair struct {
    time int
    website string
}

func mostVisitedPattern(username []string, timestamp []int, website []string) []string {
    userVisitMap := make(map[string][]Pair)
    for i := 0; i < len(username); i++ {
        if _, ok := userVisitMap[username[i]]; !ok {
            userVisitMap[username[i]] = []Pair{}
        }
        userVisitMap[username[i]] = append(userVisitMap[username[i]], Pair{timestamp[i], website[i]})
    }
    
    cand := ""
    counter := make(map[string]int)
    for _, pairs := range userVisitMap {
        sort.Slice(pairs, func(i, j int) bool {
            return pairs[i].time < pairs[j].time
        })
        set := make(map[string]bool)
        for i := 0; i < len(pairs)-2; i++ {
            for j := i+1; j < len(pairs)-1; j++ {
                for k := j+1; k < len(pairs); k++ {
                    str := pairs[i].website + " " + pairs[j].website + " " + pairs[k].website
                    if _, ok := set[str]; ok {
                        continue
                    }
                    set[str] = true
                    counter[str]++
                    if cand == "" || counter[str] > counter[cand] || (counter[str] == counter[cand] && str < cand) {
                        cand = str
                    }
                }
            }
        }
    }
    
    sequences := strings.Split(cand, " ")
    res := make([]string, 0)
    for _, sequence := range sequences {
        res = append(res, sequence)
    }
    return res
}
