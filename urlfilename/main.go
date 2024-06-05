package main

import (
	"fmt"
	"log"
	"net/url"
	"path/filepath"
	"sync"
)

func main() {
	pool := sync.Pool{
		New: func() interface{} {
			b := make([]int, 10)

			return &b
		},
	}

	bPtr := pool.Get().(*[]int)
	buf := *bPtr
	buf = buf[0:10]
	buf[0] = 7

	fmt.Printf("%#v", buf)
	fmt.Printf("%#v", len(buf))

	*bPtr = buf
	pool.Put(bPtr)

	bPtr = pool.Get().(*[]int)
	buf = *bPtr
	buf = buf[0:10]
	// buf[0] = 7

	fmt.Printf("%#v", buf)
	fmt.Printf("%#v", len(buf))

	*bPtr = buf
	pool.Put(bPtr)

	// bPtr = pool.Get().(*[]int)
	// buf = *bPtr

	// fmt.Printf("%#v", buf)
	// fmt.Printf("%#v", len(buf))
	// buf = append(buf, 123)

	// *bPtr = buf
	// pool.Put(bPtr)

	// url := "https://prd-main-s3-app-documents.s3.eu-central-1.amazonaws.com/customer/stock_export/00982350-b87b-11eb-af9d-891e063ef659/report_1666690688?response-content-disposition=inline&X-Amz-Security-Token=IQoJb3JpZ2luX2VjEBQaDmFwLXNvdXRoZWFzdC0xIkcwRQIgfxtDiFJ2NZLTXHB5QEANcA4YmzpIoZ3PyjxTD2Cnc0ECIQCie0pLgj%2Fc%2FJVfelX5yUa004eh2SMPWzJpt4%2B5uppoLir3Awjt%2F%2F%2F%2F%2F%2F%2F%2F%2F%2F8BEAIaDDIzNzU0ODcxODQwNSIM22r5OBfPS1YRv2m2KssDjIulo3Y86F1zo5HWqrLjkR%2FQzQdzqonwvVYqX4yDKlWTWjR%2B1Zmxflj%2Fd5oUrTm5Ezdluh2XOpqAcxZ414xgS8m%2FmMI7979rXlA%2FBoZYDD%2BoaNXDqWyyjbQhSbs39CBk0PENEQeGkypBBNMgDcKFfp60eqUxALQJxjcAdsWZWFCCiIxHczMf2UVVsVqz5pTUjXKl7S9Nm%2FtrIgNq3xIgTpE0Zc4SM4%2FTVwboGvjZR2Hb5eo9ptkwqGAQFuZWhTjYuANTyvwUpApx8LtFi%2BAqvqqeuh4IeeW1RvUGpquL2hZJfxAUUibnmfWS4v6dd1L5SU3fVMN0Nx4Rl2g7dfbO%2BpT%2Fks3JNOztKRMbftHSMPVzq5yc1DLkDAibP7i6VIxzPvFU1DeVfeXAMl%2BjbROvW3DrF9idwqV%2FhAKsbyCnuhhsMuSkHOGNAnKK7yU2E0qXA7h3yInuk9Lr5DrUcSjpfFsMDMtKuxmZoM9hCFUWzlGk9LQm%2Bo0hpcdBjcVBAB8st1eCFQg5RKjZ1OaDexDpZiPM68cFYSBc13s70qApySrY32At3Vwlg%2B6PxLyxf2gnKPj2YE3P1PZOMI3S9%2FWqQSZ%2Br7SbmRH706M4MPPUx64GOpQCde5ohhKlMlGL6NkOnxmA9nuDO1tk%2BLTesNm3nVNT46eJfWo4Z2TP3TNhfseRISJkx%2F6OMo2GnfzyaVqTm8tfBWrXukhh6ZOfOZDE97ZAwqy2KykagnvgMZKy3blXglR1oRC008mb0DeixtoT6TvaMvx1utNSe%2FkbJJJLAAiK4lw1doVGlzxPsKmvt7Tj7h4nDy%2BCcM374nWbFhuMk%2FJQ%2FSCm67T56d%2FfoeObzVLjg6YTAQxOpwe3enRbtS4WMqQ13HFBvOENfm643c%2BTYk2uk5uYC592yCgY%2FWD65foso5U%2B2UeVioihbCBP6ynr%2BiPZRTuZ6hcb2eMUEVigqKvigg98BTroK2mJ9cGgBZuR9snoDjPQ&X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20240218T113335Z&X-Amz-SignedHeaders=host&X-Amz-Expires=300&X-Amz-Credential=ASIATOTYBZFCQFYPE7UF%2F20240218%2Feu-central-1%2Fs3%2Faws4_request&X-Amz-Signature=bff290b2fe1708edd708d4c359746e9d55561fe5cf0a1ac4199d04bb2192078e"

	// fmt.Println(FilenameFromUrl(url))
}

func FilenameFromUrl(urlstr string) string {
	u, err := url.Parse(urlstr)
	if err != nil {
		log.Fatal("Error due to parsing url: ", err)
	}
	x, _ := url.QueryUnescape(u.EscapedPath())
	return filepath.Base(x)
}
