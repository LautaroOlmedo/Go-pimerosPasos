package basic

type employed struct{
	name string
	age int
	salary float32
	WorkStation string
	Nationality string
}

func CreateEmployed(name string, age int, salary float32, station string, nationality string) employed{
	myEmployed := employed{}

	myEmployed.name = name
	myEmployed.age = age
	myEmployed.Nationality = nationality
	myEmployed.salary = salary
	myEmployed.WorkStation = station

	return myEmployed
}

func (e *employed) SetSalary(salary float32) {
	e.salary = salary
}

func (e *employed) GetSalary() float32{
	return e.salary
}