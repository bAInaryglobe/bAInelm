package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// SearchEnginesFile is the path to the file containing the list of search engines
const SearchEnginesFile = "search_engines.txt"

// RequestsFolder is the path to the folder containing search term files
const RequestsFolder = "REQUESTS"

// ScraperAPIKey is your ScraperAPI key
const ScraperAPIKey = "YOUR_SCRAPER_API_KEY"

// UserAgents contains a list of user agent strings to rotate

var UserAgents = []string{
	// Windows
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Firefox/89.0 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Edge/91.0.864.59 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:89.0) Gecko/20100101 Firefox/89.0",
	"Mozilla/5.0 (Windows NT 6.3; Win64; x64; Trident/7.0; rv:11.0) like Gecko",
	"Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:89.0) Gecko/20100101 Firefox/89.0",
	"Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:78.0) Gecko/20100101 Firefox/78.0",
	"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; rv:11.0) like Gecko",
	"Mozilla/5.0 (Windows NT 6.1; Trident/7.0; rv:11.0) like Gecko",
	"Mozilla/5.0 (Windows NT 6.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Firefox/89.0 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Edge/91.0.864.59 Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.0; Win64; x64; rv:89.0) Gecko/20100101 Firefox/89.0",
	"Mozilla/5.0 (Windows NT 6.0; Win64; x64; rv:78.0) Gecko/20100101 Firefox/78.0",
	"Mozilla/5.0 (Windows NT 6.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.128 Safari/537.36",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) Gecko/20100101 Firefox/52.0",
	"Mozilla/5.0 (Windows NT 5.1; rv:78.0) Gecko/20100101 Firefox/78.0",
	"Mozilla/5.0 (Windows NT 5.1; rv:60.0) Gecko/20100101 Firefox/60.0",
	"Mozilla/5.0 (Windows NT 5.1; rv:59.0) Gecko/20100101 Firefox/59.0",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) Gecko/20100101 Firefox/51.0",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) Gecko/20100101 Firefox/50.0",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) Gecko/20100101 Firefox/49.0",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) Gecko/20100101 Firefox/48.0",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) Gecko/20100101 Firefox/47.0",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) Gecko/20100101 Firefox/46.0",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) Gecko/20100101 Firefox/45.0",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) Gecko/20100101 Firefox/44.0",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) Gecko/20100101 Firefox/43.0",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) Gecko/20100101 Firefox/42.0",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) Gecko/20100101 Firefox/41.0",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) Gecko/20100101 Firefox/40.0",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) Gecko/20100101 Firefox/39.0",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) Gecko/20100101 Firefox/38.0",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) Gecko/20100101 Firefox/37.0",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) Gecko/20100101 Firefox/36.0",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) Gecko/20100101 Firefox/35.0",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) Gecko/20100101 Firefox/34.0",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) Gecko/20100101 Firefox/33.0",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) Gecko/20100101 Firefox/32.0",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) Gecko/20100101 Firefox/31.0",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) Gecko/20100101 Firefox/30.0",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) Gecko/20100101 Firefox/29.0",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) Gecko/20100101 Firefox/28.0",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) Gecko/20100101 Firefox/27.0",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) Gecko/20100101 Firefox/26.0",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) Gecko/20100101 Firefox/25.0",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) Gecko/20100101 Firefox/24.0",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) Gecko/20100101 Firefox/23.0",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) Gecko/20100101 Firefox/22.0",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) Gecko/20100101 Firefox/21.0",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) Gecko/20100101 Firefox/20.0",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) Gecko/20100101 Firefox/19.0",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) Gecko/20100101 Firefox/18.0",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) Gecko/20100101 Firefox/17.0",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) Gecko/20100101 Firefox/16.0",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) Gecko/20100101 Firefox/15.0",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) Gecko/20100101 Firefox/14.0",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) Gecko/20100101 Firefox/13.0",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) Gecko/20100101 Firefox/12.0",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) Gecko/20100101 Firefox/11.0",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) Gecko/20100101 Firefox/10.0",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) Gecko/20100101 Firefox/9.0",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) Gecko/20100101 Firefox/8.0",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) Gecko/20100101 Firefox/7.0",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) Gecko/20100101 Firefox/6.0",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) Gecko/20100101 Firefox/5.0",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) Gecko/20100101 Firefox/4.0",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) Gecko/20100101 Firefox/3.6",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) Gecko/20100101 Firefox/3.5",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) Gecko/20100101 Firefox/3.0",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) AppleWebKit/536.11 (KHTML, like Gecko) Chrome/20.0.1132.57 Safari/536.11",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) AppleWebKit/536.6 (KHTML, like Gecko) Chrome/20.0.1092.0 Safari/536.6",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) AppleWebKit/536.5 (KHTML, like Gecko) Chrome/19.0.1084.56 Safari/536.5",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1063.0 Safari/536.3",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1063.0 Safari/536.3",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1062.0 Safari/536.3",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1062.0 Safari/536.3",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.1 Safari/536.3",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.0 Safari/536.3",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.0 Safari/536.3",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.0 Safari/536.3",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.0 Safari/536.3",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.0 Safari/536.3",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.0 Safari/536.3",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.0 Safari/536.3",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.0 Safari/536.3",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.0 Safari/536.3",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.0 Safari/536.3",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.0 Safari/536.3",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.0 Safari/536.3",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.0 Safari/536.3",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.0 Safari/536.3",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.0 Safari/536.3",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.0 Safari/536.3",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.0 Safari/536.3",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.0 Safari/536.3",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.0 Safari/536.3",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.0 Safari/536.3",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.0 Safari/536.3",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.0 Safari/536.3",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.0 Safari/536.3",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.0 Safari/536.3",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.0 Safari/536.3",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.0 Safari/536.3",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.0 Safari/536.3",
	"Mozilla/5.0 (Windows NT 5.1; rv:52.0) AppleWebKit/536.3 (KHTML, like Gecko) Chrome/19.0.1061.0 Safari/536.3",
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.61 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.97 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.120 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.90 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.100 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.157 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.131 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.121 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.119 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.96 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.81 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.76 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.69 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.53 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.51 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.40 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.80 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.77 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.31 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.20 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.10 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.110 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.80 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.77 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.31 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.20 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.10 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.110 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.110 Safari/537.36",
	// Mac
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36",
	// Linux
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36",
	"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:83.0) Gecko/20100101 Firefox/83.0",
	// Android
	"Mozilla/5.0 (Linux; Android 10; SM-G975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Mobile Safari/537.36",
	"Mozilla/5.0 (Linux; Android 9; SM-G960U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.73 Mobile Safari/537.36",

	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36",
}

// ScraperAPIEndpoint returns the URL for ScraperAPI requests
func ScraperAPIEndpoint(searchEngine, searchTerm, userAgent string) string {
	return fmt.Sprintf("http://api.scraperapi.com/?api_key=%s&url=%s&q=%s", ScraperAPIKey, searchEngine, searchTerm)
}

// RandomUserAgent returns a randomly selected user agent string
func RandomUserAgent() string {
	rand.Seed(time.Now().UnixNano())
	return UserAgents[rand.Intn(len(UserAgents))]
}

// ReadSearchEngines reads the list of search engines from the file
func ReadSearchEngines() ([]string, error) {
	file, err := os.Open(SearchEnginesFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var searchEngines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		searchEngines = append(searchEngines, scanner.Text())
	}

	return searchEngines, scanner.Err()
}

// ReadSearchTerm reads the first non-empty search term file and empties it
func ReadSearchTerm() (string, error) {
	files, err := ioutil.ReadDir(RequestsFolder)
	if err != nil {
		return "", err
	}

	for _, file := range files {
		if !file.IsDir() && file.Size() > 0 {
			filePath := filepath.Join(RequestsFolder, file.Name())
			content, err := ioutil.ReadFile(filePath)
			if err != nil {
				return "", err
			}

			// Empty the search term file
			err = ioutil.WriteFile(filePath, []byte{}, 0644)
			if err != nil {
				log.Println(err)
			}

			return string(content), nil
		}
	}

	return "", fmt.Errorf("no non-empty search term file found")
}

// PerformSearch performs the search using ScraperAPI
func PerformSearch(searchEngine, searchTerm, userAgent string) {
	url := ScraperAPIEndpoint(searchEngine, searchTerm, userAgent)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Set("User-Agent", userAgent)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	// Process the search results here
	fmt.Println(string(body))
}

func main() {
	searchEngines, err := ReadSearchEngines()
	if err != nil {
		log.Fatal(err)
	}

	for {
		searchTerm, err := ReadSearchTerm()
		if err != nil {
			log.Fatal(err)
		}

		// Select 10 random search engines
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(searchEngines), func(i, j int) {
			searchEngines[i], searchEngines[j] = searchEngines[j], searchEngines[i]
		})
		selectedSearchEngines := searchEngines[:10]

		// Perform search for each selected search engine
		for _, searchEngine := range selectedSearchEngines {
			userAgent := RandomUserAgent()
			PerformSearch(searchEngine, searchTerm, userAgent)
		}
	}
}
