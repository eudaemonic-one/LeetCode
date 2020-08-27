var days = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

func dayOfTheWeek(day int, month int, year int) string {
    if month < 3 {
        month += 12
        year -= 1
    }
    c := year / 100
    year = year % 100
    w := (c/4 - 2*c + year + year/4 + 13*(month+1)/5 + day - 1) % 7
    return days[(w+7)%7]
}
