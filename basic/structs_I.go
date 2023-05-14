package basic

func StructI(){
	type Country struct{
		Poblation int64
		Lenguage string
		Capital string
		Name string
	}

	var ARG Country

	ARG.Poblation = 48583921
	ARG.Capital = "Buenos Aires"
	ARG.Lenguage = "espanol"
	ARG.Name = "Argentina"
}