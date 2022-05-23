package controller

import (
	"github.com/google/uuid"

	"github.com/okutsen/PasswordManager/internal/log"
	"github.com/okutsen/PasswordManager/schema/apischema"
	"github.com/okutsen/PasswordManager/schema/dbschema"
	"github.com/okutsen/PasswordManager/schema/schemabuilder"
)

type Repository interface {
	AllRecords() ([]dbschema.Record, error)
	RecordByID(id uuid.UUID) (*dbschema.Record, error)
	CreateRecord(record *dbschema.Record) (*dbschema.Record, error)
	UpdateRecord(record *dbschema.Record) (*dbschema.Record, error)
	DeleteRecord(id uuid.UUID) (*dbschema.Record, error)

	AllUsers() ([]dbschema.User, error)
	UserByID(id uuid.UUID) (*dbschema.User, error)
	CreateUser(record *dbschema.User) (*dbschema.User, error)
	UpdateUser(record *dbschema.User) (*dbschema.User, error)
	DeleteUser(id uuid.UUID) (*dbschema.User, error)
}

type Controller struct {
	repo Repository
	log  log.Logger
}

func NewController(logger log.Logger, ctrl Repository) *Controller {
	return &Controller{
		log:  logger.WithFields(log.Fields{"service": "Controller"}),
		repo: ctrl,
	}
}

func (c *Controller) AllRecords() ([]apischema.Record, error) {
	getRecords, err := c.repo.AllRecords()
	if err != nil {
		return nil, err
	}

	recordsAPI := schemabuilder.BuildAPIRecordsFromDBRecords(getRecords)
	return recordsAPI, err
}

func (c *Controller) Record(id uuid.UUID) (*apischema.Record, error) {
	getRecord, err := c.repo.RecordByID(id) // TODO: pass uuid
	if err != nil {
		return nil, err
	}

	recordAPI := schemabuilder.BuildAPIRecordFromDBRecord(getRecord)
	return &recordAPI, err
}

// TODO: return specific errors to identify on api 404 Not found, 409 Conflict(if exists)
func (c *Controller) CreateRecord(record *apischema.Record) (*apischema.Record, error) {
	dbRecord := schemabuilder.BuildDBRecordFromAPIRecord(record)
	createdDBRecord, err := c.repo.CreateRecord(&dbRecord)
	if err != nil {
		return nil, err
	}

	createdAPIRecord := schemabuilder.BuildAPIRecordFromDBRecord(createdDBRecord)
	return &createdAPIRecord, err
}

// 200, 204(if no changes?), 404
func (c *Controller) UpdateRecord(id uuid.UUID, record *apischema.Record) (*apischema.Record, error) {
	dbRecord := schemabuilder.BuildDBRecordFromAPIRecord(record)
	dbRecord.ID = id

	updatedRecord, err := c.repo.UpdateRecord(&dbRecord)
	if err != nil {
		return nil, err
	}

	updatedApiRecord := schemabuilder.BuildAPIRecordFromDBRecord(updatedRecord)
	return &updatedApiRecord, err
}

// 200, 404
func (c *Controller) DeleteRecord(id uuid.UUID) (*apischema.Record, error) {
	dbRecord, err := c.repo.DeleteRecord(id)
	if err != nil {
		return nil, err
	}

	record := schemabuilder.BuildAPIRecordFromDBRecord(dbRecord)

	return &record, err
}

func (c *Controller) AllUsers() ([]apischema.User, error) {
	getUsers, err := c.repo.AllUsers()
	if err != nil {
		return nil, err
	}

	usersAPI := schemabuilder.BuildAPIUsersFromDBUsers(getUsers)
	return usersAPI, err
}

func (c *Controller) User(id uuid.UUID) (*apischema.User, error) {
	getUser, err := c.repo.UserByID(id) // TODO: pass uuid
	if err != nil {
		return nil, err
	}

	user := schemabuilder.BuildAPIUserFromDBUser(getUser)
	return &user, err
}

// TODO: return specific errors to identify on api 404 Not found, 409 Conflict(if exists)

func (c *Controller) CreateUser(user *apischema.User) (*apischema.User, error) {
	dbUser := schemabuilder.BuildDBUserFromAPIUser(user)
	createdDBUser, err := c.repo.CreateUser(&dbUser)
	if err != nil {
		return nil, err
	}

	createAPIUser := schemabuilder.BuildAPIUserFromDBUser(createdDBUser)
	return &createAPIUser, err
}

// 200, 204(if no changes?), 404
func (c *Controller) UpdateUser(id uuid.UUID, user *apischema.User) (*apischema.User, error) {
	dbUser := schemabuilder.BuildDBUserFromAPIUser(user)
	dbUser.ID = id

	updatedUser, err := c.repo.UpdateUser(&dbUser)
	if err != nil {
		return nil, err
	}

	updatedAPIUser := schemabuilder.BuildAPIUserFromDBUser(updatedUser)
	return &updatedAPIUser, err
}

// 200, 404
func (c *Controller) DeleteUser(id uuid.UUID) (*apischema.User, error) {
	dbUser, err := c.repo.DeleteUser(id)
	if err != nil {
		return nil, err
	}

	user := schemabuilder.BuildAPIUserFromDBUser(dbUser)

	return &user, err
}
