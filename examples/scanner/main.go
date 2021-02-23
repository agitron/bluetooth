package main

import (
	"encoding/json"
	"fmt"

	"github.com/lacendarko/bluetooth"
)

var adapter = bluetooth.DefaultAdapter

func main() {
	// Enable BLE interface.
	must("enable BLE stack", adapter.Enable())

	// Start scanning.
	println("scanning...123")
	err := adapter.Scan(func(adapter *bluetooth.Adapter, device bluetooth.ScanResult) {
		println("found device:", device.Address.String(), device.RSSI, device.LocalName())
		b, err := json.Marshal(device.GetManufacturerData())
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(b))
	})

	must("start scan", err)
}

func must(action string, err error) {
	if err != nil {
		panic("failed to " + action + ": " + err.Error())
	}
}
