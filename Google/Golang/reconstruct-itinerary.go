type Tickets [][]string

func (tickets Tickets) Len() int {
    return len(tickets)
}

func (tickets Tickets) Less(i, j int) bool {
    if tickets[i][0] == tickets[j][0] {
        return tickets[i][1] < tickets[j][1]
    }
    return tickets[i][0] < tickets[j][0]
}

func (tickets Tickets) Swap(i, j int) {
    tickets[i], tickets[j] = tickets[j], tickets[i]
}

func findItinerary(tickets [][]string) []string {
    itinerary := make([]string, 0)
    targets := make(map[string][]string)
    sort.Sort(sort.Reverse(Tickets(tickets)))
    for _, ticket := range tickets {
        departure, arrival := ticket[0], ticket[1]
        targets[departure] = append(targets[departure], arrival)
    }
    dfs(targets, "JFK", &itinerary)
    for i := 0; i < len(itinerary)/2; i++ {
        itinerary[i], itinerary[len(itinerary)-i-1] = itinerary[len(itinerary)-i-1], itinerary[i]
    }
    return itinerary
}

func dfs(targets map[string][]string, airport string, itinerary *[]string) {
    for len(targets[airport]) > 0 {
        destinations := targets[airport]
        destination := destinations[len(destinations)-1]
        targets[airport] = targets[airport][:len(destinations)-1]
        dfs(targets, destination, itinerary)
    }
    *itinerary = append(*itinerary, airport)
}
