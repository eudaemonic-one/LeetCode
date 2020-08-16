func countOfAtoms(formula string) string {
    N := len(formula)
    stack := make([](map[string]int), 0)
    stack = append(stack, make(map[string]int))
    
    for i := 0; i < N; {
        if formula[i] == '(' {
            stack = append(stack, make(map[string]int))
            i++
        } else if formula[i] == ')' {
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            i++
            digitStart := i
            multiples := 1
            for i < N && isDigit(formula[i]) {
                i++
            }
            if i > digitStart {
                multiples, _ = strconv.Atoi(formula[digitStart:i])
            }
            for atom, val := range top {
                stack[len(stack)-1][atom] += val * multiples
            }
        } else {
            atomStart := i
            i++
            for i < N && isLowerCase(formula[i]) {
                i++
            }
            atomName := formula[atomStart:i]
            digitStart := i
            for i < N && isDigit(formula[i]) {
                i++
            }
            multiples := 1
            if i > digitStart {
                multiples, _ = strconv.Atoi(formula[digitStart:i])
            }
            stack[len(stack)-1][atomName] += multiples
        }
    }
    
    if len(stack) > 0 {
        return mapToString(stack[len(stack)-1])
    }
    
    return ""
}

func isLowerCase(ch byte) bool {
    return 'a' <= ch && ch <= 'z'
}

func isDigit(ch byte) bool {
    return '0' <= ch && ch <= '9'
}

type pair struct {
    atomName string
    count int
}

func mapToString(table map[string]int) string {
    pairs := []pair{}
    for atomName, multiples := range table {
        pairs = append(pairs, pair{atomName, multiples})
    }
    
    sort.SliceStable(pairs, func (i, j int) bool {
        if pairs[i].atomName == pairs[j].atomName {
            return pairs[i].count < pairs[j].count
        }
        return pairs[i].atomName < pairs[j].atomName
    })
    
    var sb strings.Builder
    for _, pair := range pairs {
        sb.WriteString(pair.atomName)
        if pair.count > 1 {
            sb.WriteString(strconv.Itoa(pair.count))
        }
    }
    return sb.String()
}
