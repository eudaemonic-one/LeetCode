type Codec struct {
    tinyUrlMap map[string]string
}


func Constructor() Codec {
    return Codec{make(map[string]string)}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
    h := sha256.New()
	h.Write([]byte(longUrl))
    sha256 := string(h.Sum(nil))
    this.tinyUrlMap[sha256] = longUrl
    return "http://tinyurl.com/" + sha256
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
    return this.tinyUrlMap[strings.ReplaceAll(shortUrl, "http://tinyurl.com/", "")]
}


/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */

