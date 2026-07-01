package main

import "fmt"


type ChargeableDevice interface {
	Charge() string
}

type Phone struct {
	BatteryLevel float64
}

func (p Phone) Charge() string {
	if p.BatteryLevel < 20 {
		return "Charging phone with 20W fast charger..."
	} else {
		return "Phone is sufficient charged"
	}
}

func DeviceChargeStatus(device ChargeableDevice) {
	fmt.Println(device.Charge())
}


type Laptop struct {
	BatteryLevel float64
}

func (l Laptop) Charge() string {
	if l.BatteryLevel < 20 {
		return "Charging laptop with 65W fast charger..."
	} else {
		return "Laptop is sufficient charged"
	}
}


func main() {

	phone := Phone{BatteryLevel: 10}
	DeviceChargeStatus(phone)

	laptop := Laptop{BatteryLevel:22}
	DeviceChargeStatus(laptop)
	
}

