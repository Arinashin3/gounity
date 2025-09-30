package gounity

import (
	"log"
	"testing"
)

func TestName(t *testing.T) {
	client := NewClient("https://10.77.77.221", "admin", "Passw0rd1!", true)
	client.GetBasicSystemInfo()
	client.GetSystemCapacity()
	log.Printf("%v", client)

}
