package types

import (
	"fmt"
	"time"
)

type ServiceResp struct {
	URL                 string
	Data                string
	RespCode            int
	TimeTakenNanoSecond int
	InsertedAt          time.Time
}

func (sr ServiceResp) String() string {
	return fmt.Sprintf("ServiceResp<%s %s %d>", sr.URL, sr.Data, sr.RespCode)
}
