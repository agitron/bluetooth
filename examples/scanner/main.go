package main

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"tinygo.org/x/bluetooth"
)

var adapter = bluetooth.DefaultAdapter

const manufacturerKey = 0x944

func main() {
	// Enable BLE interface.
	must("enable BLE stack", adapter.Enable())

	// Start scanning.
	println("scanning...")
	err := adapter.Scan(func(adapter *bluetooth.Adapter, device bluetooth.ScanResult) {
		println("found device:", device.Address.String(), device.RSSI, device.LocalName())
		md := device.AdvertisementPayload.GetManufacturerData(manufacturerKey)

		byteReader := bytes.NewReader(md)

		byteStorage := make([]byte, 2)
		n, _ := byteReader.ReadAt(byteStorage, 2)

		bat := byteStorage[:n]

		if len(bat) > 0 {
			fmt.Println(bat)
			data := binary.BigEndian.Uint16(bat)
			fmt.Println(data)
		}

		// fmt.Printf("%d \n", )

	})

	must("start scan", err)
}

func must(action string, err error) {
	if err != nil {
		panic("failed to " + action + ": " + err.Error())
	}
}
