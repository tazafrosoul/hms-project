package userservice

import (
	i "github.com/tazafrosoul/hms-project/common/interfaces"
	s "github.com/tazafrosoul/hms-project/common/structs"
)

type UserService struct {

	//repository interface
	repo i.Repo
}

func NewUserService(repo i.Repo) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (us *UserService) AddUser(req s.AddUserReq) (s.AddUserRes, error) {
	//get request
	//TODO 2nd validation here

	//format it into user struct
	user := s.User{
		FullName:       req.FullName,
		Username:       req.Username,
		Avatar:         req.Avatar,
		HashedPassword: req.Password,
		Role:           req.Role,
		CreatedBy:      req.By,
	}

	res, err := us.repo.Create(user)

	if err != nil {
		return s.AddUserRes{}, err
	}
	return s.AddUserRes{
		ID:       res,
		FullName: req.FullName,
		Username: req.Username,
		Avatar:   req.Avatar,
		Role:     req.Role,
	}, nil
}
