package gounity

import (
	"log"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	c := NewClient("https://10.77.77.221", "admin", "Passw0rd1!", true)
	var filters []string
	filters = append(filters, "isHistoricalAvailable EQ true")
	filters = append(filters, "isRealtimeAvailable EQ true")
	c.GetBasicSystemInfoInstances()
	c.GetLunInstances([]string{}, nil)
	c.GetEventInstances([]string{}, nil)
	data, err := c.PostMetricRealTimeQueryInstances([]string{"sp.*.cpu.summary.busyTicks"}, 1*time.Minute)
	if err != nil {
		t.Error(err)
		return
	}
	log.Printf("%v", data)

}
