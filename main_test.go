package gounity

import (
	"log"
	"testing"
)

func TestName(t *testing.T) {
	c := NewClient("https://10.77.77.221", "admin", "Passw0rd1!", true)
	var filters []string
	filters = append(filters, "isHistoricalAvailable EQ true")
	filters = append(filters, "isRealtimeAvailable EQ true")
	c.GetBasicSystemInfoInstances()
	d, err := c.GetMetricValueInstances("sp.*.physical.disk.*.responseTime")
	//data, err := c.PostMetricRealTimeQueryInstances([]string{"sp.*.cpu.summary.busyTicks"}, 1*time.Minute)
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
	log.Printf("%v", d)
	log.Printf("%v", err)

}
