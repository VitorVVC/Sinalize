package repositories

import (
	"SinalizeB/models"
	"gorm.io/gorm"
)

// TranslateRepository interface define os métodos para o repositório
type TranslateRepository interface {
	Create(translate *models.Translate) error
	FindByWord(word string) (*models.Translate, error)
}

// translateRepositoryImpl é a implementação do TranslateRepository
type translateRepositoryImpl struct {
	db *gorm.DB
}

// NewTranslateRepository cria uma nova instância do repositório de traduções
func NewTranslateRepository(db *gorm.DB) TranslateRepository {
	return &translateRepositoryImpl{db: db}
}

func (r *translateRepositoryImpl) Create(translate *models.Translate) error {
	return r.db.Create(translate).Error
}
func (r *translateRepositoryImpl) FindByWord(word string) (*models.Translate, error) {
	var translate models.Translate
	if err := r.db.Where("word = ?", word).First(&translate).Error; err != nil {
		return nil, err
	}
	return &translate, nil
}
