package services

import (
	"assignment1/models"
	"assignment1/repositories"
	"fmt"
)

func Checking(name string) bool {
	for _, n := range repositories.Names {
		if name == n {
			return true
		}
	}
	return false
}

func DataProcess(name string) {
	var person = models.Bio{}
	names := repositories.Names
	addresses := repositories.Addresses
	jobs := repositories.Jobs
	reasons := repositories.Reasons
	keys := repositories.Keys
	result := repositories.Result
	mapResult := repositories.MapResult

	for i, n := range names {
		if name == n {
			person = models.Bio{
				ID:      i,
				Name:    names[i],
				Address: addresses[i],
				Job:     jobs[i],
				Reason:  person.SetReason(name, reasons[i]),
			}
			for _, k := range keys {
				switch k {
				case "ID":
					result = append(result, fmt.Sprintf("%v : %v", k, person.ID))
					mapResult[k] = person.ID
				case "name":
					result = append(result, fmt.Sprintf("%v : %v", k, person.Name))
					mapResult[k] = person.Name
				case "address":
					result = append(result, fmt.Sprintf("%v : %v", k, person.Address))
					mapResult[k] = person.Address
				case "job":
					result = append(result, fmt.Sprintf("%v : %v", k, person.Job))
					mapResult[k] = person.Job
				case "reason":
					result = append(result, fmt.Sprintf("%v : %v", k, person.Reason))
					mapResult[k] = person.Reason
				}
			}
			break
		}
	}
	for _, v := range result {
		fmt.Printf("%v\n", v)
	}
	fmt.Printf("Print from struct: Bio%+v\n", person)
	fmt.Printf("Print from mapEmptyInterface: %+v\n", mapResult)

}
