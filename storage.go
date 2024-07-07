package main

import (
	"errors"
)

type RestrauntStorage interface {
	AddRestraunt(restr *Restraunt)
	GetRestrauntById(id string) (*Restraunt, error)
	SearchRestraunt(specs []RestrauntSpecification) ([]*Restraunt, error)
}

var _ RestrauntStorage = &RestrauntStorageSlice{}

type RestrauntStorageSlice struct {
	restraunts []*Restraunt
}

func (r *RestrauntStorageSlice) AddRestraunt(restr *Restraunt) {
	r.restraunts = append(r.restraunts, restr)
}

func (r *RestrauntStorageSlice) GetRestrauntById(id string) (*Restraunt, error) {
	for _, r := range r.restraunts {
		if r.GetId() == id {
			return r, nil
		}
	}

	return nil, errors.New("record not found")
}

func (r *RestrauntStorageSlice) SearchRestraunt(specs []RestrauntSpecification) ([]*Restraunt, error) {
	var result []*Restraunt

	andSpecs := AndSpecifications{
		specs: specs,
	}

	for _, restraunt := range r.restraunts {
		if andSpecs.IsSatisfiedByAll(restraunt){
			result = append(result,  restraunt)
		}		
	}

	if len(result) == 0{
		return nil, errors.New("record not found")
	}

	return result, nil
}