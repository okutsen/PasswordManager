package schemabuilder

import (
	"github.com/okutsen/PasswordManager/schema/apischema"
	"github.com/okutsen/PasswordManager/schema/dbschema"
)

func BuildAPIUserFromDBUser(user *dbschema.User) apischema.User {
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

func BuildAPIUsersFromDBUsers(users []dbschema.User) []apischema.User {
	usersAPI := make([]apischema.User, len(users))
	for i, v := range users {
		usersAPI[i] = BuildAPIUserFromDBUser(&v)
	}

	return usersAPI
}

func BuildDBUserFromAPIUser(user *apischema.User) dbschema.User {
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
