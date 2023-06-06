package authentication_service

import (
	"fmt"
	"ngmi_server/internal/models/user_model"
	authenticationRepository "ngmi_server/internal/repositories/authentication_repository"
	"ngmi_server/pkg/log"
)

type Service interface {
	Register(userRegisterReq user_model.RegisterReq) (user_model.Resp, error)
	Login(userLoginReq user_model.LoginReq) (user_model.Resp, error)
}

type service struct {
	repo   authenticationRepository.Repository
	logger log.Logger
}

func NewService(repo authenticationRepository.Repository, logger log.Logger) Service {
	return service{repo: repo, logger: logger}
}

func (s service) Register(userRegisterReq user_model.RegisterReq) (user_model.Resp, error) {
	if result, err := s.repo.Register(userRegisterReq); err != nil {
		return user_model.Resp{}, err
	} else {
		if id, err := result.LastInsertId(); err != nil {
			return user_model.Resp{}, fmt.Errorf("an error has occurred while trying to get the last registeration id: %v", err)
		} else {
			return user_model.Resp{
				Id:           id,
				FirstName:    userRegisterReq.FirstName,
				LastName:     userRegisterReq.LastName,
				Username:     userRegisterReq.Username,
				EmailAddress: userRegisterReq.EmailAddress,
				BirthDate:    userRegisterReq.BirthDate,
				Gender:       userRegisterReq.Gender,
			}, nil
		}
	}
}

func (s service) Login(userLoginReq user_model.LoginReq) (user_model.Resp, error) {
	//if u, err := s.repo.Login(userLoginReq); err != nil {
	//	return user_model.Resp{}, err
	//} else {
	//	return u, nil
	//}
	return user_model.Resp{}, nil
}
