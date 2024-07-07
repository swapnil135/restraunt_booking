package main

type IRestrauntManager interface{
	RegisterRestraunt(restr *Restraunt)
	SearchRestraunt(specs []RestrauntSpecification) ([]*Restraunt, error)
	BookTable(restId, date string, startTime, numPeople int) error
}

type RestrauntManger struct {
	storage RestrauntStorage
}

func (r *RestrauntManger) RegisterRestraunt(restr *Restraunt) {
	r.storage.AddRestraunt(restr)
}

func (r *RestrauntManger) SearchRestraunt(specs []RestrauntSpecification) ([]*Restraunt, error){
	return r.storage.SearchRestraunt(specs)
}

func (r *RestrauntManger) BookTable(restId , date string, startTime, numPeople int) error{
	restr, err := r.storage.GetRestrauntById(restId)
	if err != nil {
		return err
	}

	err = restr.BookSlot(date, startTime, numPeople)
	if err != nil {
		return err
	}

	return nil
}