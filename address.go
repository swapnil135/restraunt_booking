package main

type Address struct {
	street string
	locality string
	city string
}

func NewAddress(street string, locality string, city string) *Address {
	return &Address{
		street: street,
		locality: locality,
		city: city,
	}
}

func (a *Address) GetCity() string {
	if(a == nil) {return ""}
	return a.city
}

func (a *Address) GetLocality() string {
	if(a == nil) {return ""}
	return a.locality
}

func (a *Address) Getstreet() string {
	if(a == nil) {return ""}
	return a.street
}