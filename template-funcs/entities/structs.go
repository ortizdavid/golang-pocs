package entities

type Customer struct {
	FirstName	string
	MiddleName	string
	LastName	string
	Gender		string
	Balance		float64
}

// customer List sample
var CustomerList = []Customer{
	{
		FirstName: "John",
		MiddleName: "B.",
		LastName: "Doe",
		Gender: "M",
		Balance: 1900000.86,
	},
	{
		FirstName: "Anna",
		MiddleName: "Nick.",
		LastName: "Smith",
		Gender: "f",
		Balance: 340000.99,
	},
	{
		FirstName: "Mary",
		MiddleName: "",
		LastName: "Joseph",
		Gender: "",
		Balance: 3500,
	},
	{
		FirstName: "Peter",
		MiddleName: "D. T.",
		LastName: "Jackson",
		Gender: "m",
		Balance: 7430922.35,
	},
}