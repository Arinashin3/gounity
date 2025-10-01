package gounity

import (
	"log"
	"testing"
)

func TestName(t *testing.T) {
	client := NewClient("https://10.77.77.221", "admin", "Passw0rd1!", true)
	bsi, err := client.GetBasicSystemInfoInstances()
	log.Printf("%v%v%v", client, bsi, err)
	lun, err := client.GetLunInstances([]string{"id", "sizeTotal"}, "")
	for _, entry := range lun.Entries {
		entry.Content.SizeTotal.ToMiB()
	}
	log.Printf("%v%v", lun, err)

	// System Test
	//system, err := client.GetSystemInstances([]string{"id", "health", "name", "model", "serialNumber"}, "")
	//if err != nil {
	//	log.Printf("%v", err)
	//	os.Exit(1)
	//}
	//for _, entry := range system.Entries {
	//	log.Printf("%s", entry.Content.Name)
	//}

	// SystemCapacity Test
	//systemCapacity, err := client.GetSystemCapacityInstances([]string{"id", "sizeFree", "sizeTotal", "sizeSubscribed", "totalLogicalSize", "dataReductionRatio", "sizeUsed", "tiers"}, "")
	//if err != nil {
	//	log.Printf("%v", err)
	//	os.Exit(1)
	//}
	//for _, entry := range systemCapacity.Entries {
	//	log.Printf("%v", entry.Content.Id)
	//	log.Printf("Size Total: %.2f", entry.Content.SizeTotal.ToGiB())
	//	log.Printf("Size Used:  %.2f", entry.Content.SizeUsed.ToGiB())
	//	log.Printf("Size Subscribed:  %.2f", entry.Content.SizeSubscribed.ToGiB())
	//	log.Printf("Total Logical Size:  %.2f", entry.Content.TotalLogicalSize.ToGiB())
	//	log.Printf("%v", entry.Content.Tiers)
	//	for _, tier := range entry.Content.Tiers {
	//		log.Printf("%v", tier.TierType.String())
	//	}
	//}

	// StorageResource Test
	//sr, err := client.GetStorageResourceInstances([]string{"id", "name", "description", "type", "sizeTotal", "health", "dataReductionSizeSaved", "dataReductionPercent", "dataReductionRatio"}, "")
	//for _, entry := range sr.Entries {
	//	log.Printf("%v%v%v", sr, entry, err)
	//}

}
