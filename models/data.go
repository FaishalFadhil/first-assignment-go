package models

import (
	"fmt"
	// "sort"
)

var keys = []string{"ID", "name", "address", "job", "reason"}
var names = []string{"Thomas", "Ridwan", "Tiara", "Udin", "David"}
var addresses = []string{"Bandung", "Jakarta", "Tangerang", "Bogor", "Sukabumi"}
var jobs = []string{"Programmer", "Developer", "Researcher", "Diver", "Photographer"}
var reasons = []string{"know golang", "upgrade skill", "curious about golang", "learn golang other than dive", "...I don't know, i got this class for free"}
var result []string

func Checking(name string) bool {
	for _, n := range names {
		if name == n {
			return true
		}
	}
	return false
}

func DataProcess(name string) {
	var person = Bio{}
	for i, n := range names {
		if name == n {
			person = Bio{
				ID:      i,
				name:    names[i],
				address: addresses[i],
				job:     jobs[i],
				reason:  person.setReason(name, reasons[i]),
			}
			for _, k := range keys {
				switch k {
				case "ID":
					result = append(result, fmt.Sprintf("%v : %v", k, person.ID))
				case "name":
					result = append(result, fmt.Sprintf("%v : %v", k, person.name))
				case "address":
					result = append(result, fmt.Sprintf("%v : %v", k, person.address))
				case "job":
					result = append(result, fmt.Sprintf("%v : %v", k, person.job))
				case "reason":
					result = append(result, fmt.Sprintf("%v : %v", k, person.reason))
				}
			}
			break
		}
	}
	for _, v := range result {
		fmt.Printf("%v\n", v)
	}
	fmt.Printf("Print from struct: Bio%+v", person)

}
