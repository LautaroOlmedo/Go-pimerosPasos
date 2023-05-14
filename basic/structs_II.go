package basic

import "fmt"

type Country struct{
	Poblation int64
	Lenguage string
	Capital string
	Name string
}

func StructII(){
	var countries = make([]Country, 0)
	var option string = "n"
	// var country Country // ---> Opcion 1 para cargar valores en Slice
	var countryName string
	var countryPoblation int64
	var countryLenguage string
	var countryCapital string

	for {
		fmt.Println("Do you want a storage a new country? (y/n)")
		fmt.Scanln(&option)
		if(option != "y"){
			fmt.Println("Exit. Bye.")
			break
		}else{
			fmt.Println("Enter the country name")
			fmt.Scanln(&countryName)

		    fmt.Printf("Enter the capital of %v \n", countryName)
		    fmt.Scanln(&countryCapital)

		    fmt.Printf("Enter the lenguaje spoken %v \n", countryName)
		    fmt.Scanln(&countryLenguage)

			fmt.Printf("How many people lives in %v ? \n", countryName)
		    fmt.Scanln(&countryPoblation)

			// country.Name = countryName              |
			// country.Capital = countryCapital        |
			// country.Lenguage = countryLenguage      | ---> Opcion 1 para cargar valores en Slice
			// country.Poblation = countryPoblation    |
			// countries = append(countries, country)  |

			countries = append(countries, Country{Name: countryName,  // |
				Capital: countryCapital,                              // | ---> Opcion 2 para cargar valores en Slice
				Lenguage: countryLenguage,                            // |
				Poblation: countryPoblation})                         // | 

			//fmt.Print(countries, "\n") // ---> Opcion 1 para mostrar

			for k, country := range countries { // ---> Opcion 2 para mostrar
				fmt.Printf("---------- COUNTRY %v ---------- \n", k+1)
				fmt.Println("Country name: ", country.Name)
				fmt.Println("Country capital: ", country.Capital)
				fmt.Println("Country lenguage: ", country.Lenguage)
				fmt.Println("Country poblation: ", country.Poblation)
				fmt.Println("--------------------")
			}
		}
	}
}