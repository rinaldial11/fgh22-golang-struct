package main

import (
	"fmt"
	"strconv"
)

type bio struct {
	name        string
	dateOfBirth string
}

func (profile bio) dateConv(date string) string {

	monthForm := [12]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"august",
		"September",
		"October",
		"November",
		"December",
	}

	resDay := ""

	if string(date[0]) == "0" {
		resDay += string(date[1])
	} else {
		resDay += string(date[0]) + string(date[1])
	}

	resMonth := ""

	if string(date[3]) == "0" {
		for i := 0; i < len(monthForm); i++ {
			monthIndex := "0" + strconv.Itoa(i+1)

			if string(date[3])+string(date[4]) == monthIndex {
				resMonth += monthForm[i]
				break
			}

		}
		return resDay + " " + resMonth + " " + string(date[6]) + string(date[7]) + string(date[8]) + string(date[9])
	}

	for y := 9; y < len(monthForm); y++ {
		monthIndexS := strconv.Itoa(y + 1)
		// fmt.Println(monthIndexS)
		if string(date[3])+string(date[4]) == monthIndexS {
			resMonth += monthForm[y]
			break
		}
	}
	return resDay + " " + resMonth + " " + string(date[6]) + string(date[7]) + string(date[8]) + string(date[9])
}

func main() {
	profile := bio{
		name:        "Kim Nal di",
		dateOfBirth: "01/12/1997",
	}

	fmt.Println(profile.name)
	fmt.Println(profile.dateOfBirth)
	fmt.Println(profile.dateConv(profile.dateOfBirth))
}
