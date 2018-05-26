package pm25

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func GetData(location string) (int, error) {
	rp, e := http.Get("http://www.baidu.com/s?wd=" + location + "空气质量")
	if e != nil {
		return -1, e
	}
	b, e := ioutil.ReadAll(rp.Body)
	if e != nil {
		return -1, e
	}
	key := "op_pm25_top_column op_pm25_top_aqi"
	str := string(b)
	index := strings.Index(str, key)
	if index < 0 {
		return index, errors.New("couldn't find key")
	}
	mindex := index + len(key) + 2
	number := trimNumber(str[mindex : mindex+3])
	fmt.Println("number :", number)
	n, e := strconv.Atoi(number)
	if e != nil {
		fmt.Println(e)
		return -1, e
	}
	return n, nil
}
func trimNumber(number string) string {
	for i := 0; i < len(number); i++ {
		if number[i:i+1] == "<" {
			return number[:i]
		}
	}
	return number
}
