package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

/*
it seems the prometheus client channel copy is the root cause of GC there.
Hence in this part of code.
we should check memory information with mock code.
0. basic data as a struct with int value.
*/
type Basic struct {
	Value int
}

func newDescforTest() *prometheus.Desc {
	return prometheus.NewDesc(
		"dummy",
		"just a dummy",
		[]string{"instance"}, nil,
	)
}

/*
1. we do samilar with prometheus client.
1.1 create a channel
1.2 convert basic data into prometheus type and feed into channel
1.3 copy channel
1.4 read from channel and close channel
1.5 loop from 1.2 to 1.4 for 4 times as data refresh
*/
func ChannelBased() {
	MemStats := &runtime.MemStats{}
	runtime.ReadMemStats(MemStats)
	for k := 0; k < 3; k++ {
		theChan := make(chan prometheus.Metric, 1000)
		for i := 0; i < 1000; i++ {
			instance := &Basic{
				Value: i,
			}
			theChan <- prometheus.MustNewConstMetric(
				newDescforTest(),
				prometheus.GaugeValue,
				float64(instance.Value),
				"dummy",
			)
		}
		newChan := theChan
		close(theChan)
		for j := 0; j < 1000; j++ {
			test := <-newChan
			if test == nil {
				fmt.Errorf("empty")
				os.Exit(1)
			}
		}
	}
	after := &runtime.MemStats{}
	runtime.ReadMemStats(after)
	fmt.Println("Channel")
	fmt.Println(after.TotalAlloc - MemStats.TotalAlloc)
	fmt.Println(after.HeapAlloc - MemStats.HeapAlloc)
	fmt.Println(after.Mallocs - MemStats.Mallocs)
	fmt.Println(after.NextGC - MemStats.NextGC)
	fmt.Println(after.GCSys - MemStats.GCSys)
	fmt.Println(after.NumGC - MemStats.NumGC)
}

/*
comparing with just pointer memory usage
as this is the basic case for all data in the memeory which alloc can not avoid.
2. use a memory based hash map
2.1 make a prometheus struct based basic data struct
*/
func BaseLine() {
	MemStats := &runtime.MemStats{}
	runtime.ReadMemStats(MemStats)
	//fmt.Println(MemStats.TotalAlloc)
	table := make(map[string]prometheus.Gauge)
	for i := 0; i < 1000; i++ {
		instance := &Basic{
			Value: i,
		}
		table[strconv.Itoa(i)] = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Namespace: "Dummy",
				Help:      "just a dummy",
			},
		)
		table[strconv.Itoa(i)].Set(float64(instance.Value))
	}
	for j := 1; j < 3; j++ {
		for i := 0; i < 1000; i++ {
			table[strconv.Itoa(i)].Set(float64(i))
		}
	}
	after := &runtime.MemStats{}
	runtime.ReadMemStats(after)
	fmt.Println("baseline")
	fmt.Println(after.TotalAlloc - MemStats.TotalAlloc)
	fmt.Println(after.HeapAlloc - MemStats.HeapAlloc)
	fmt.Println(after.Mallocs - MemStats.Mallocs)
	fmt.Println(after.NextGC - MemStats.NextGC)
	fmt.Println(after.GCSys - MemStats.GCSys)
	fmt.Println(after.NumGC - MemStats.NumGC)
}

func main() {
	//for i := 0; i < 10; i++ {
	debug.FreeOSMemory()
	BaseLine()
	//}
	//debug.FreeOSMemory()
	//for i := 0; i < 10; i++ {
	debug.FreeOSMemory()
	ChannelBased()
	//}
}
