package repositorytmpl

var (
	// MongoRepositoryTemplate is base template for mongo repository
	MongoRepositoryTemplate = `
	package {{.PackageName}}

  type {{.Model.Name}}Repository struct {
  	*mgo.Collection
  }

  func New{{.Model.Name}}Repository(col *mgo.Collection) *{{.Model.Name}}Repository {
  	return &{{.Model.Name}}Repository{Collection: col}
  }

  func (repo *{{.Model.Name}}Repository) Create(model *{{.Model.ImportName}}) (error) {
  	return repo.Insert(model)
  }

  func (repo *{{.Model.Name}}Repository) Read(selection *{{.Model.ImportName}}) ([]{{.Model.ImportName}}, error) {
  	models := []{{.Model.ImportName}}{}

  	return models, repo.Find(selection).All(&models)
  }

  func (repo *{{.Model.Name}}Repository) ReadOne(selection *{{.Model.ImportName}}) (*{{.Model.ImportName}}, error) {
  	model := &{{.Model.ImportName}}{}

  	return model, repo.Find(selection).One(model)
  }

  func (repo *{{.Model.Name}}Repository) ReadID(id bson.ObjectId) (*{{.Model.ImportName}}, error) {
  	model := &{{.Model.ImportName}}{}

  	return model, repo.FindId(id).One(model)
  }

  func (repo *{{.Model.Name}}Repository) Update(selection *{{.Model.ImportName}}, update *{{.Model.ImportName}}) (error) {
  	return repo.Update(selection, update)
  }

  func (repo *{{.Model.Name}}Repository) UpdateID(update *{{.Model.ImportName}}) (error) {
  	return repo.UpdateId(update.ID, update)
  }

  func (repo *{{.Model.Name}}Repository) Delete(selection *{{.Model.ImportName}}) (error) {
  	return repo.Remove(selection)
  }

  func (repo *{{.Model.Name}}Repository) DeleteID(id bson.ObjectId) (error) {
  	return repo.RemoveId(id)
  }`
)
