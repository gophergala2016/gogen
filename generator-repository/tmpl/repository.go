package repositorytmpl

var (
	// MongoRepositoryTemplate is base template for mongo repository
	MongoRepositoryTemplate = `
	package {{.PackageName}}

  type {{.Name}}Repository struct {
  	*mgo.Collection
  }

  func New{{.Name}}Repository(col *mgo.Collection) *{{.Name}}Repository {
  	return &{{.Name}}Repository{Collection: col}
  }

  func (repo *{{.Name}}Repository) Create(model *{{.Name}}Model) (error) {
  	return repo.Insert(model)
  }

  func (repo *{{.Name}}Repository) Read(selection *{{.Name}}Model) ([]{{.Name}}Model, error) {
  	models := []{{.Name}}Model{}

  	return models, repo.Find(selection).All(&models)
  }

  func (repo *{{.Name}}Repository) ReadOne(selection *{{.Name}}Model) (*{{.Name}}Model, error) {
  	model := &{{.Name}}Model{}

  	return model, repo.Find(selection).One(model)
  }

  func (repo *{{.Name}}Repository) ReadID(id bson.ObjectId) (*{{.Name}}Model, error) {
  	model := &{{.Name}}Model{}

  	return model, repo.FindId(id).One(model)
  }

  func (repo *{{.Name}}Repository) Update(selection *{{.Name}}Model, update *{{.Name}}Model) (error) {
  	return repo.Update(selection, update)
  }

  func (repo *{{.Name}}Repository) UpdateID(update *{{.Name}}Model) (error) {
  	return repo.UpdateId(update.ID, update)
  }

  func (repo *{{.Name}}Repository) Delete(selection *{{.Name}}Model) (error) {
  	return repo.Remove(selection)
  }

  func (repo *{{.Name}}Repository) DeleteID(id bson.ObjectId) (error) {
  	return repo.RemoveId(id)
  }`
)
