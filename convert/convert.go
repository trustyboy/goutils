package convert

import (
	"github.com/Chain-Zhang/pinyin"
	"strconv"
	"strings"
	"time"
)

func StrToInt64(numStr string) int64 {
	num, err := strconv.ParseInt(numStr, 10, 64)
	if err != nil {
		return 0
	}
	return num
}

func StrToInt(numStr string) int {
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return 0
	}
	return num
}

func StrToFloat64(numStr string) float64 {
	num, err := strconv.ParseFloat(numStr, 32)
	if err != nil {
		return 0
	}
	return num
}

func TimeStrToUnix(timeStr string) int64 {
	times, _ := time.Parse("2006-01-02 15:04:05", timeStr)
	return times.Unix()
}

func ConvertToPY(hanZhi string) (string, error) {
	pyStr, err := pinyin.New(hanZhi).Split("-").Convert()
	if err != nil {
		return "", err
	}
	arrStr := strings.Split(pyStr, "-")
	ret := ""
	for _, value := range arrStr {
		if len(value) > 0 {
			ret = ret + string(value[0])
		}
	}
	return ret, nil
}


