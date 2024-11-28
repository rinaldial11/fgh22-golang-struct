package structexample

import "fmt"

type nameStruct struct {
	name map[string]string
}
type profile struct {
	results []nameStruct
}

func DataBio() {
	data := profile{
		results: []nameStruct{
			{
				name: map[string]string{
					"firstname": "Ji eun",
					"lastname":  "Kim",
				},
			},
			{
				name: map[string]string{
					"firstname": "Tae ri",
					"lastname":  "Kim",
				},
			},
			{
				name: map[string]string{
					"firstname": "Suzy",
					"lastname":  "Bae",
				},
			},
			{
				name: map[string]string{
					"firstname": "Da mi",
					"lastname":  "Kim",
				},
			},
			{
				name: map[string]string{
					"firstname": "Yoo jung",
					"lastname":  "Kim",
				},
			},
			{
				name: map[string]string{
					"firstname": "Fazztrack",
					"lastname":  "Ahay",
				},
			},
		},
	}

	fmt.Println(data.results[5].name["firstname"])
	// fmt.Println(data.results[0].name["lastname"], data.results[0].name["firstname"])
}
