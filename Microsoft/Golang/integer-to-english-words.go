func numberToWords(num int) string {
    if num == 0 {
        return "Zero"
    }
    
    billion := num / 1000000000
    million := (num - billion * 1000000000) / 1000000
    thousand := (num - billion * 1000000000 - million * 1000000) / 1000
    rest := num - billion * 1000000000 - million * 1000000 - thousand * 1000

    var sb strings.Builder
    
    if billion != 0 {
        sb.WriteString(threeDigitsToStr(billion) + " Billion ")
    }
    
    if million != 0 {
        sb.WriteString(threeDigitsToStr(million) + " Million ")
    }
    
    if thousand != 0 {
        sb.WriteString(threeDigitsToStr(thousand) + " Thousand ")
    }
    
    if rest != 0 {
        sb.WriteString(threeDigitsToStr(rest))
    }
    
    return strings.TrimSpace(sb.String())
}

func digitToStr(num int) string {
    switch num {
        case 0:
            return ""
        case 1:
            return "One"
        case 2:
            return "Two"
        case 3:
            return "Three"
        case 4:
            return "Four"
        case 5:
            return "Five"
        case 6:
            return "Six"
        case 7:
            return "Seven"
        case 8:
            return "Eight"
        case 9:
            return "Nine"
    }
    return ""
}

func twoDigitsLessThan20ToStr(num int) string {
    switch num {
        case 10:
            return "Ten"
        case 11:
            return "Eleven"
        case 12:
            return "Twelve"
        case 13:
            return "Thirteen"
        case 14:
            return "Fourteen"
        case 15:
            return "Fifteen"
        case 16:
            return "Sixteen"
        case 17:
            return "Seventeen"
        case 18:
            return "Eighteen"
        case 19:
            return "Nineteen"
    }
    return ""
}

func tensToStr(num int) string {
    switch num {
        case 2:
            return "Twenty"
        case 3:
            return "Thirty"
        case 4:
            return "Forty"
        case 5:
            return "Fifty"
        case 6:
            return "Sixty"
        case 7:
            return "Seventy"
        case 8:
            return "Eighty"
        case 9:
            return "Ninety"
    }
    return ""
}

func twoDigitsToStr(num int) string {
    if num == 0 {
        return ""
    } else if num < 10 {
        return digitToStr(num)
    } else if num < 20 {
        return twoDigitsLessThan20ToStr(num)
    } else {
        ten := num / 10
        rest := num - ten * 10
        if rest != 0 {
            return tensToStr(ten) + " " + digitToStr(rest)
        } else {
            return tensToStr(ten)
        }
    }
}

func threeDigitsToStr(num int) string {
    if num == 0 {
        return ""
    }
    hundred := num / 100
    rest := num - hundred * 100
    if hundred * rest != 0 {
        return digitToStr(hundred) + " Hundred " + twoDigitsToStr(rest)
    } else if hundred == 0 && rest != 0 {
        return twoDigitsToStr(rest)
    } else if hundred != 0 && rest == 0 {
        return digitToStr(hundred) + " Hundred"
    }
    return ""
}
