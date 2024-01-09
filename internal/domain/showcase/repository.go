package showcase

import (
	"github.com/fiqrikm18/wikifood/core/internal/exception"
	"gorm.io/gorm"
)

type IShowcaseRepository interface {
	CreateShowcase(Showcase Showcase) (*Showcase, error)
	UpdateShowcase(Showcase Showcase, id string) (*Showcase, error)
	DeleteShowcase(id string) error
}

type ShowcaseRepository struct{ DB *gorm.DB }

func NewShowcaseRepository(db *gorm.DB) ShowcaseRepository {
	return ShowcaseRepository{DB: db}
}

func (repo *ShowcaseRepository) CreateShowcase(showcase Showcase) (*Showcase, error) {
	dbTx := repo.DB.Create(&Showcase{})
	if err := dbTx.Error; err != nil {
		return nil, err
	}

	err := dbTx.Row().Scan(&showcase)
	if err != nil {
		return nil, err
	}

	return &showcase, nil
}

func (repo *ShowcaseRepository) UpdateShowcase(showcase Showcase, id int) (*Showcase, error) {
	return nil, &exception.NotImplementedException{}
}

func (repo *ShowcaseRepository) DeleteShowcase(id string) error {
	dbTx := repo.DB.Where("id = ?", id).Delete(&Showcase{})
	if err := dbTx.Error; err != nil {
		return err
	}

	return nil
}
