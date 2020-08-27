func validIPAddress(IP string) string {
    ipv4Fields := strings.Split(IP, ".")
    if len(ipv4Fields) == 4 {
        for _, field := range ipv4Fields {
            if len(field) > 1 && field[0] == '0' {
                return "Neither"
            }
            value, err := strconv.Atoi(field)
            if err != nil || value < 0 || value > 255 {
                return "Neither"
            }
        }
        return "IPv4"
    }
    
    ipv6Fields := strings.Split(IP, ":")
    if len(ipv6Fields) == 8 {
        for _, field := range ipv6Fields {
            if len(field) < 1 || len(field) > 4 {
                return "Neither"
            }
            for i := 0; i < len(field); i++ {
                if ch := field[i]; !(('0' <= ch && ch <= '9') || ('a' <= ch && ch <= 'f') || ('A' <= ch && ch <= 'F')) {
                    return "Neither"
                }
            }
        }
        return "IPv6"
    }
    
    return "Neither"
}
