package repository

import (
	"github.com/gophergala2016/gogen/examples/repository/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserRepository struct {
	*mgo.Collection
}

func NewUserRepository(col *mgo.Collection) *UserRepository {
	return &UserRepository{Collection: col}
}

func (repo *UserRepository) Create(model *model.User) error {
	return repo.Insert(model)
}

func (repo *UserRepository) Read(selection *model.User) ([]model.User, error) {
	models := []model.User{}

	return models, repo.Find(selection).All(&models)
}

func (repo *UserRepository) ReadOne(selection *model.User) (*model.User, error) {
	model := &model.User{}

	return model, repo.Find(selection).One(model)
}

func (repo *UserRepository) ReadID(id bson.ObjectId) (*model.User, error) {
	model := &model.User{}

	return model, repo.FindId(id).One(model)
}

func (repo *UserRepository) Update(selection *model.User, update *model.User) error {
	return repo.Update(selection, update)
}

func (repo *UserRepository) UpdateID(update *model.User) error {
	return repo.UpdateId(update.ID, update)
}

func (repo *UserRepository) Delete(selection *model.User) error {
	return repo.Remove(selection)
}

func (repo *UserRepository) DeleteID(id bson.ObjectId) error {
	return repo.RemoveId(id)
}
