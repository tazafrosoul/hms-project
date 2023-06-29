package userservice

import (
	i "hms-project/common/interfaces"
	s "hms-project/common/structs"
)

type UserService struct {

	//repository interface
	repo i.Repo
}

func NewUserService(repo i.Repo) *UserService {
	return &UserService{}
}

func (us *UserService) AddUser(req s.AddUserReq) (s.AddUserRes, error) {
	//get request
	//TODO validation here

	hashedPw := req.Password //TODO ----- (1) Hashing password...

	//format it into user struct
	user := s.User{
		FullName:       req.FullName,
		Username:       req.Username,
		Email:          req.Email,
		Avatar:         req.Avatar,
		HashedPassword: hashedPw,
		Role:           req.Role,
		CreatedBy:      req.By, //TODO find a way to send a request with the active user ID
	}

	//call repository method
	res, err := us.repo.Create(user)

	if err != nil {
		return s.AddUserRes{}, err
	}
	return s.AddUserRes{
		ID:       res,
		FullName: req.FullName,
		Username: req.Username,
		Email:    req.Email,
		Avatar:   req.Avatar,
		Role:     req.Role,
	}, nil
}
