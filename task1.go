package main

import "fmt"


func main(){
	assignment1()
}

func assignment1() {
	familytoDeviceTypesMap := make(map[string][]string)

	familytoDeviceTypesMap["sros"] = []string{"7750 SR", "7705 SAR", "7210 SAS", "dcwbx", "dctor", "dcvsa", "7705 SAR Hm"}
	familytoDeviceTypesMap["aos"] = []string{"10k", "6250", "5350"}
	familytoDeviceTypesMap["mpr"] = []string{"9500", "9500e"}

	fmt.Println(familytoDeviceTypesMap)

	//Iterating and printing

	for key, val := range familytoDeviceTypesMap {
		fmt.Printf("Device Family: %v\n\t\t", key)
		for _, data := range val {
			fmt.Printf("Device Type: %v, ", data)
		}
		fmt.Println()
	}

	//checking whether the key is present
	val, ok := familytoDeviceTypesMap["sdn"]

	if !ok {
		fmt.Println("The device family is not found")
		familytoDeviceTypesMap["sdn"] = []string{"sdn"}
	} else {
		fmt.Println("Value is:", val)
	}

	//checking a device type present in device family

	srosdevices, ok := familytoDeviceTypesMap["sros"]
	if !ok {
		fmt.Println("The device family not found")
	} else {
		var found bool = false
		for _, devicetype := range srosdevices {
			if devicetype == "7210 SAS" {
				found = true
				break
			}
		}
		if found {
			fmt.Println("Device Found!")
		} else {
			fmt.Println("Device Not Found!")
		}
	}

	//removing a depricated devices from the list of mpr devices
	mprdevices, ok := familytoDeviceTypesMap["mpr"]
	input := "9500e"
	if !ok {
		fmt.Println("MPR devices are not found")
	} else {
		for index, devices := range mprdevices {
			//var filterdValues []string
			if devices == input {
				fmt.Println("Removing the device")
				mprdevices = append(mprdevices[:index], mprdevices[index+1:]...)
				break
			}
		}
	}

	fmt.Println(mprdevices)
}
