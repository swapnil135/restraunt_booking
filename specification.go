package main

type RestrauntSpecification interface{
	IsRestrauntValid(restraunt *Restraunt) bool
}

type RestrauntSpecificationCity struct {
	city string
}

func NewRestrauntSpecificationCity(city string) *RestrauntSpecificationCity{
	return &RestrauntSpecificationCity{city: city}
}

func (r *RestrauntSpecificationCity) IsRestrauntValid(restraunt *Restraunt) bool{
	return restraunt.GetAddress().GetCity() == r.city
}


type RestrauntSpecificationCuisine struct {
	cuisuin Cuisine
}

func NewRestrauntSpecificationCuisine(cuisine Cuisine) *RestrauntSpecificationCuisine{
	return &RestrauntSpecificationCuisine{cuisuin: cuisine}
}

func (r *RestrauntSpecificationCuisine) IsRestrauntValid(restraunt *Restraunt) bool{
	return restraunt.GetCuisine() == r.cuisuin
}

type AndSpecifications struct {
	specs []RestrauntSpecification
}

func (s *AndSpecifications) IsSatisfiedByAll(restraunt *Restraunt) bool {
	for _, spec := range s.specs {
		if !spec.IsRestrauntValid(restraunt){
			return false
		}
	}

	return true;
}