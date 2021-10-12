package store

import (
	"github.com/gamberooni/go-cats/model"
	"gorm.io/gorm"
)

// cat store is a wrapper around the gorm db
type CatStore struct {
	db *gorm.DB
}

// return store instance to interact with db
func NewCatStore(db *gorm.DB) *CatStore {
	return &CatStore{
		db: db,
	}
}

func (cs *CatStore) DeleteCatById(id int) error {
	cat := model.Cat{}                 // initialize an empty variable of Cat type
	err := cs.db.First(&cat, id).Error // get cat by id from db
	if err != nil {
		return err
	}
	cs.db.Delete(&cat) // delete the cat with the specified id
	return nil
}

func (cs *CatStore) UpdateCatById(id int, c *model.Cat) (*model.Cat, error) {
	cat := model.Cat{}
	err := cs.db.First(&cat, id).Error // get cat by id from db

	if err != nil {
		return nil, err
	}

	// update the cat with the specified id from db with the new values
	cs.db.Model(&cat).Updates(model.Cat{Name: c.Name, Breed: c.Breed})

	return &cat, nil
}

func (cs *CatStore) GetCatById(id int) (*model.Cat, error) {
	cat := model.Cat{}
	err := cs.db.First(&cat, id).Error

	if err != nil {
		return nil, err
	}

	return &cat, nil
}

func (cs *CatStore) GetAllCats() ([]model.Cat, error) {
	cats := []model.Cat{}
	result := cs.db.Find(&cats)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil // return nil if error is record not found - not raised as error
		}

		return nil, result.Error
	}
	return cats, nil
}

func (cs *CatStore) AddCat(c *model.Cat) error {
	result := cs.db.Create(&c)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
