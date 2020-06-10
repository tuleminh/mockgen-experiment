package main

import "errors"

func main() {

}

//
// Service
//

func NewService(r Repository) *Service {
	return &Service{
		r: r,
	}
}

type Service struct {
	r Repository
}

type CreateSomethingRequest struct {
	Num int `json:"num"`
}

func (s *Service) CreateSomething(request *CreateSomethingRequest) error {
	insertedID, err := s.r.Insert(request.Num)
	if err != nil {
		return errors.New("inserting failed")
	}
	if insertedID <= 0 {
		return errors.New("invalid inserted ID")
	}
	return nil
}

type Repository interface {
	Insert(num int) (insertedID int, err error)
}
