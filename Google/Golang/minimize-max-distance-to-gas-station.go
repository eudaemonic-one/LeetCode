func minmaxGasDist(stations []int, K int) float64 {
    l, r := 0.0, float64(stations[len(stations)-1]-stations[0])
    m := 0.0
    for r-l > 1e-6 {
        m = (l+r) / 2.0
        cnt := 0
        for i := 0; i < len(stations)-1; i++ {
            cnt += int(float64(stations[i+1]-stations[i]) / m)
        }
        if cnt <= K {
            r = m
        } else {
            l = m
        }
    }
    return l
}
