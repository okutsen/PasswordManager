package controller

import (
	"github.com/okutsen/PasswordManager/internal/log"
	"github.com/okutsen/PasswordManager/schema/apischema"
	"github.com/okutsen/PasswordManager/schema/dbschema"
	"github.com/okutsen/PasswordManager/schema/schemabuilder"
)

type UsersRepo interface {
	AllUsers() ([]dbschema.User, error)
	UserByID(id uint64) (*dbschema.User, error)
	CreateUser(record *dbschema.User) (*dbschema.User, error)
	UpdateUser(record *dbschema.User) (*dbschema.User, error)
	DeleteUser(id uint64) error
}

type UsersController struct {
	users UsersRepo
	log   log.Logger
}

func NewUsers(logger log.Logger, repo UsersRepo) *UsersController {
	return &UsersController{
		log:   logger.WithFields(log.Fields{"service": "Controller"}),
		users: repo,
	}
}

func (c *UsersController) AllUsers() ([]apischema.User, error) {
	getUsers, err := c.users.AllUsers()
	if err != nil {
		return nil, err
	}

	usersAPI := schemabuilder.BuildAPIUsersFromDBUsers(getUsers)
	return usersAPI, err
}

func (c *UsersController) User(id uint64) (*apischema.User, error) {
	getUser, err := c.users.UserByID(id) // TODO: pass uuid
	if err != nil {
		return nil, err
	}

	recordAPI := schemabuilder.BuildAPIUserFromDBUser(getUser)
	return &recordAPI, err
}

// TODO: return specific errors to identify on api 404 Not found, 409 Conflict(if exists)

func (c *UsersController) CreateUser(user *apischema.User) (*apischema.User, error) {
	dbUser := schemabuilder.BuildDBUserFromAPIUser(user)
	createdDBUser, err := c.users.CreateUser(&dbUser)
	if err != nil {
		return nil, err
	}

	createAPIUser := schemabuilder.BuildAPIUserFromDBUser(createdDBUser)
	return &createAPIUser, err
}

// 200, 204(if no changes?), 404
func (c *UsersController) UpdateUser(id uint64, user *apischema.User) (*apischema.User, error) {
	dbUser := schemabuilder.BuildDBUserFromAPIUser(user)
	dbUser.ID = id

	updatedUser, err := c.users.UpdateUser(&dbUser)
	if err != nil {
		return nil, err
	}

	updatedAPIUser := schemabuilder.BuildAPIUserFromDBUser(updatedUser)
	return &updatedAPIUser, err
}

// 200, 404
func (c *UsersController) DeleteUser(id uint64) error {
	return c.users.DeleteUser(id)
}
