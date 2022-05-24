package controller

import (
	"github.com/google/uuid"

	"github.com/okutsen/PasswordManager/internal/log"
	"github.com/okutsen/PasswordManager/schema/controllerSchema"
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
	CreateUser(user *dbschema.User) (*dbschema.User, error)
	UpdateUser(user *dbschema.User) (*dbschema.User, error)
	DeleteUser(id uuid.UUID) (*dbschema.User, error)
}

type Controller struct {
	repo Repository
	log  log.Logger
}

func New(logger log.Logger, ctrl Repository) *Controller {
	return &Controller{
		log:  logger.WithFields(log.Fields{"service": "Controller"}),
		repo: ctrl,
	}
}

func (c *Controller) AllRecords() ([]controllerSchema.Record, error) {
	getDBRecords, err := c.repo.AllRecords()
	if err != nil {
		return nil, err
	}
	records := schemabuilder.BuildControllerRecordsFromDBRecords(getDBRecords)

	return records, nil
}

func (c *Controller) Record(id uuid.UUID) (*controllerSchema.Record, error) {
	DBRecord, err := c.repo.RecordByID(id) // TODO: pass uuid
	if err != nil {
		return nil, err
	}

	record := schemabuilder.BuildControllerRecordFromDBRecord(DBRecord)
	return &record, nil
}

// TODO: return specific errors to identify on api 404 Not found, 409 Conflict(if exists)
func (c *Controller) CreateRecord(record *controllerSchema.Record) (*controllerSchema.Record, error) {
	dbRecord := schemabuilder.BuildDBRecordFromControllerRecord(record)
	createRecord, err := c.repo.CreateRecord(&dbRecord)
	if err != nil {
		return nil, err
	}

	createdRecord := schemabuilder.BuildControllerRecordFromDBRecord(createRecord)
	return &createdRecord, nil
}

// 200, 204(if no changes?), 404
func (c *Controller) UpdateRecord(id uuid.UUID, record *controllerSchema.Record) (*controllerSchema.Record, error) {
	dbRecord := schemabuilder.BuildDBRecordFromControllerRecord(record)
	dbRecord.ID = id

	updateRecord, err := c.repo.UpdateRecord(&dbRecord)
	if err != nil {
		return nil, err
	}

	updatedRecord := schemabuilder.BuildControllerRecordFromDBRecord(updateRecord)
	return &updatedRecord, nil
}

// 200, 404
func (c *Controller) DeleteRecord(id uuid.UUID) (*controllerSchema.Record, error) {
	dbRecord, err := c.repo.DeleteRecord(id)
	if err != nil {
		return nil, err
	}

	record := schemabuilder.BuildControllerRecordFromDBRecord(dbRecord)

	return &record, nil
}

func (c *Controller) AllUsers() ([]controllerSchema.User, error) {
	getUsers, err := c.repo.AllUsers()
	if err != nil {
		return nil, err
	}

	users := schemabuilder.BuildControllerUsersFromDBUsers(getUsers)
	return users, nil
}

func (c *Controller) User(id uuid.UUID) (*controllerSchema.User, error) {
	getUser, err := c.repo.UserByID(id) // TODO: pass uuid
	if err != nil {
		return nil, err
	}

	user := schemabuilder.BuildControllerUserFromDBUser(getUser)
	return &user, nil
}

// TODO: return specific errors to identify on api 404 Not found, 409 Conflict(if exists)

func (c *Controller) CreateUser(user *controllerSchema.User) (*controllerSchema.User, error) {
	dbUser := schemabuilder.BuildDBUserFromControllerUser(user)
	createdDBUser, err := c.repo.CreateUser(&dbUser)
	if err != nil {
		return nil, err
	}

	createdUser := schemabuilder.BuildControllerUserFromDBUser(createdDBUser)
	return &createdUser, nil
}

// 200, 204(if no changes?), 404
func (c *Controller) UpdateUser(id uuid.UUID, user *controllerSchema.User) (*controllerSchema.User, error) {
	dbUser := schemabuilder.BuildDBUserFromControllerUser(user)
	dbUser.ID = id

	updateUser, err := c.repo.UpdateUser(&dbUser)
	if err != nil {
		return nil, err
	}

	updatedUser := schemabuilder.BuildControllerUserFromDBUser(updateUser)
	return &updatedUser, nil
}

// 200, 404
func (c *Controller) DeleteUser(id uuid.UUID) (*controllerSchema.User, error) {
	dbUser, err := c.repo.DeleteUser(id)
	if err != nil {
		return nil, err
	}

	user := schemabuilder.BuildControllerUserFromDBUser(dbUser)

	return &user, nil
}
