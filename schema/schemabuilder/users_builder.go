package schemabuilder

import (
	"github.com/okutsen/PasswordManager/schema/apischema"
	"github.com/okutsen/PasswordManager/schema/controllerSchema"
	"github.com/okutsen/PasswordManager/schema/dbschema"
)

func BuildControllerUserFromDBUser(user *dbschema.User) controllerSchema.User {
	return controllerSchema.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Login:     user.Login,
		Password:  user.Password,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func BuildControllerUsersFromDBUsers(users []dbschema.User) []controllerSchema.User {
	usersController := make([]controllerSchema.User, len(users))
	for i, v := range users {
		usersController[i] = BuildControllerUserFromDBUser(&v)
	}

	return usersController
}

func BuildAPIUserFromControllerUser(user *controllerSchema.User) apischema.User {
	return apischema.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Login:     user.Login,
		Password:  user.Password,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func BuildAPIUsersFromControllerUsers(users []controllerSchema.User) []apischema.User {
	usersController := make([]apischema.User, len(users))
	for i, v := range users {
		usersController[i] = BuildAPIUserFromControllerUser(&v)
	}

	return usersController
}

func BuildControllerUserFromAPIUser(user *apischema.User) controllerSchema.User {
	return controllerSchema.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Login:     user.Login,
		Password:  user.Password,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func BuildDBUserFromControllerUser(user *controllerSchema.User) dbschema.User {
	return dbschema.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Login:     user.Login,
		Password:  user.Password,
		Phone:     user.Phone,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
