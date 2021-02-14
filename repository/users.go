package repository

import (
	"go-fiber-auth-api/db"
	"go-fiber-auth-api/models"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const UsersCollection = "users"

type UsersRepository interface {
	Save(user *models.User) error
	Update(user *models.User) error
	GetById(id string) (user *models.User, err error)
	GetByEmail(email string) (user *models.User, err error)
	GetAll() (users []*models.User, err error)
	Delete(id string) error
}

type usersRepository struct {
	c *mgo.Collection
}

func NewUsersRepository(conn db.Connection) UsersRepository {
	return &usersRepository{conn.DB().C(UsersCollection)}
}

func (r *usersRepository) Save(user *models.User) error {
	return r.c.Insert(user)
}

func (r *usersRepository) Update(user *models.User) error {
	return r.c.UpdateId(user.Id, user)
}

func (r *usersRepository) GetById(id string) (user *models.User, err error) {
	err = r.c.FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

func (r *usersRepository) GetByEmail(email string) (user *models.User, err error) {
	err = r.c.Find(bson.M{"email": email}).One(&user)
	return user, err
}

func (r *usersRepository) GetAll() (users []*models.User, err error) {
	err = r.c.Find(bson.M{}).All(&users)
	return users, err
}

func (r *usersRepository) Delete(id string) error {
	return r.c.RemoveId(bson.ObjectIdHex(id))
}
