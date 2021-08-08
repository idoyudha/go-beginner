package main

import "fmt"

func main() {
	address1 := Address{"Jember", "East Java", "Indonesia"}
	address2 := &address1 // pointer to address 1 with add '&', without '&' it just copied the values of address1
	// address3 := &address1

	address2.City = "Jakarta" // address1 change to Jakarta
	address2.Province = "DKI Jakarta"

	// address2 = &Address{"Bandung", "West Java", "Indonesia"} // set address2 to new Address (different with address1)
	*address2 = Address{"Bandung", "West Java", "Indonesia"} // set all to new Address (address1 follow this changes)

	/**
	{Jakarta DKI Jakarta Indonesia}
	&{Bandung West Java Indonesia}
	*/

	address2 = &Address{"Bandung", "West Java", "Indonesia"}
	/**
	{Bandung West Java Indonesia}
	&{Bandung West Java Indonesia}
	*/

	fmt.Println(address1)
	fmt.Println(address2)
	// fmt.Println(address3)

	address4 := new(Address)
	address4.City = "Palembang"
	fmt.Println(address4) // Palembang

	alamat := Address {
		City: "Tokyo",
		Province: "Middle",
		Country: "",
	}
	ChangeCountryToJapan(&alamat)
	fmt.Println(alamat)
}

type Address struct {
	City, Province, Country string
}


func ChangeCountryToJapan(address *Address) {
	address.Country = "Japan"
}