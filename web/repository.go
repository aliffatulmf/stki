package web

import (
	"aliffatulmf/stki/model"
	"fmt"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) Create(model *model.Corpus) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(model).Error; err != nil {
			return fmt.Errorf("RepositoryCreate: %w", err)
		}
		return nil
	})
}

func (r *Repository) Finds() ([]model.Corpus, error) {
	var result []model.Corpus
	db := r.DB
	if err := db.Find(&result).Error; err != nil {
		return result, fmt.Errorf("RepositoryFinds: %w", err)
	}

	return result, nil
}

func (r *Repository) Find(conds ...string) (model.Corpus, error) {
	var result model.Corpus
	db := r.DB
	if err := db.First(&result, conds[0], conds[1:]).Error; err != nil {
		return result, fmt.Errorf("RepositoryFind: %w", err)
	}

	return result, nil
}
