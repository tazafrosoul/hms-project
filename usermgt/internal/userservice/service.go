package userservice

import (
	i "hms-project/common/interfaces"
	s "hms-project/common/structs"
	"log"
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

	hashedPw := req.Password //TODO Hash password in the http server, dont transport it naked
	log.Printf("hashing password %v", hashedPw)

	//format it into user struct
	user := s.User{
		FullName:       req.FullName,
		Username:       req.Username,
		Email:          req.Email,
		Avatar:         req.Avatar,
		HashedPassword: hashedPw,
		Role:           req.Role,
		CreatedBy:      req.By,
	}

	log.Print("callin inmemo create")
	res, err := us.repo.Create(user)

	log.Print("returned successfully")

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
