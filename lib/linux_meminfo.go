package linuxMeminfo

import "fmt"

var meminfoDict = []map[string]string{
	{"MemTotal": "memtotal"},
	{"MemFree": "memfree"},
	{"MemAvailable": "memavailable"},
	{"Buffers": "buffers"},
	{"Cached": "cached"},
	{"SwapCached": "swapcached"},
	{"Active": "active"},
	{"Inactive": "inactive"},
	{"Active(anon)": "active_anon"},
	{"Inactive(anon)": "inactive_anon"},
	{"Active(file)": "active_anon"},
	{"Inactive(file)": "inactive_file"},
}

type MemInfo struct {
	memtotal      float64
	memfree       float64
	memavailable  float64
	buffers       float64
	cached        float64
	swapcached    float64
	active        float64
	inactive      float64
	active_anon   float64
	inactive_anon float64
	active_file   float64
	inactive_file float64
}

func (m *MemInfo) parseMetrics() {
	for k, v := range meminfoDict {
		fmt.Println(k, v)
	}
}

func Do() {
	m := new(MemInfo)
	m.parseMetrics()
	fmt.Println("test")
}
