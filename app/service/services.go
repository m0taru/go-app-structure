package service

import (
	"context"
	
	"github.com/google/uuid"

	"github.com/m0taru/go-app-structure/app/model"
)

// UserService is a service for users
type UserService interface {
	GetUserByEmail(context.Context, string) (*model.User, error)
	
	Get(context.Context, uuid.UUID) (*model.User, error)
	Create(context.Context, *model.User) (*model.User, error)
	Update(context.Context, *model.User) (*model.User, error)
	Delete(context.Context, uuid.UUID) error
	List(context.Context) (model.Users, error)

}