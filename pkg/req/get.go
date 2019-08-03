package req

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/girishramnani/uptimer/pkg/types"
)

func GetAllUrls(urlList []string) chan types.ServiceResp {
	resps := make(chan types.ServiceResp, len(urlList))
	var wg sync.WaitGroup
	wg.Add(len(urlList))

	for _, url := range urlList {
		go func(url string) {
			defer wg.Done()
			resp, err := GetDataAndStatus(url)
			if err != nil {
				log.Println("Error while getting", url, ":", err)
				resps <- types.ServiceResp{
					URL:      url,
					RespCode: 0,
					Data:     err.Error(),
				}
				return
			}
			resps <- *resp
		}(url)
	}

	go func() {
		// we wait and close
		wg.Wait()
		close(resps)
	}()

	return resps

}

func GetDataAndStatus(url string) (*types.ServiceResp, error) {
	t := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	next := time.Now()

	return &types.ServiceResp{
		URL:                 url,
		Data:                strings.TrimSpace(string(body)),
		RespCode:            resp.StatusCode,
		TimeTakenNanoSecond: int(next.UnixNano() - t.UnixNano()),
	}, nil
}
