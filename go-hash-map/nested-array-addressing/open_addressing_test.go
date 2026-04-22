package nestedarrayaddressing

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"testing"
	"time"
)

type Metrics struct {
	putLatencies []int64
	getLatencies []int64
}

func (m *Metrics) RecordSet(t time.Duration) {
	m.putLatencies = append(m.putLatencies, t.Nanoseconds())
}

func (m *Metrics) RecordGet(t time.Duration) {
	m.getLatencies = append(m.getLatencies, t.Nanoseconds())
}

func percentile(data []int64, p float64) int64 {
	sort.Slice(data, func(i, j int) bool {
		return data[i] < data[j]
	})

	index := int(float64(len(data))*p) - 1
	if index < 0 {
		index = 0
	}
	return data[index]
}

func (m *Metrics) Report() {
	fmt.Println("PUT:")
	fmt.Println("p50:", percentile(m.putLatencies, 0.50))
	fmt.Println("p90:", percentile(m.putLatencies, 0.90))
	fmt.Println("p99:", percentile(m.putLatencies, 0.99))

	fmt.Println("GET:")
	fmt.Println("p99:", percentile(m.getLatencies, 0.99))
}

var MAP_LEN int = 2000

func TestLoad(t *testing.T) {
	m := Metrics{
		getLatencies: make([]int64, MAP_LEN),
		putLatencies: make([]int64, MAP_LEN),
	}
	hm := newMap(10)
	for i := range MAP_LEN {
		start := time.Now()
		hm.set("hm"+strconv.Itoa(i), rand.Int())
		m.RecordSet(time.Since(start))
	}
	for i := range MAP_LEN {
		start := time.Now()
		hm.get("hm" + strconv.Itoa(i))
		m.RecordGet(time.Since(start))
	}
	m.Report()
}
