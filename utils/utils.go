package utils

import (
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
	"time"
)

func GetContent(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Status error: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Read body: %v", err)
	}

	return data, nil
}

func ValInSlice(list []string, a string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func KeyInSlice(slice []string, val string) int {
	for i, item := range slice {
		if item == val {
			return i
		}
	}
	return -1
}

func KeyInMap(data map[string]interface{}, key string) bool {
	for item, _ := range data {
		if item == key {
			return true
		}
	}
	return false
}

func KeyInStringMap(data map[string]string, key string) bool {
	for item, _ := range data {
		if item == key {
			return true
		}
	}
	return false
}

func GetFormateTime() string {
	t := time.Now()
	day := strconv.Itoa(t.Day())
	mon := string(int(t.Month()))
	hour := strconv.Itoa(t.Hour())
	min := strconv.Itoa(t.Minute())
	//week := t.Weekday().String()[0:3]

	return mon + "-" + day + "-" + hour + ":" + min
}

func FmtDuration(d time.Duration) string {
	d = d.Round(time.Minute)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	return fmt.Sprintf("%02d:%02d", h, m)
}

func FormatedDuration(dur time.Duration) string {
	m := math.Round(dur.Minutes())
	h := math.Floor(m / 60)
	if h == 0 {
		return fmt.Sprintf("%02dm", int(m))
	}
	d := int(math.Floor(h / 24))
	if d == 0 {
		return fmt.Sprintf("%02dh", int(h))
	}
	return fmt.Sprintf("%02dd", int(d))
}
