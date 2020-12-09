package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/micro/micro/v3/service/logger"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"

	board_entities "github.com/ygpark2/mboard/service/board/proto/entities"
)

// boardRepository interface
type BoardRepository interface {
	Exist(model *board_entities.BoardORM) bool
	List(limit, page int, sort string, model *board_entities.BoardORM) (total int64, boards []*board_entities.BoardORM, err error)
	Get(id string) (*board_entities.BoardORM, error)
	Create(model *board_entities.BoardORM) error
	Update(id string, model *board_entities.BoardORM) error
	Delete(model *board_entities.BoardORM) error
}

// boardRepository struct
type boardRepository struct {
	db *gorm.DB
}

// NewBoardRepository returns an instance of `BoardRepository`.
func NewBoardRepository(db *gorm.DB) BoardRepository {
	return &boardRepository{
		db: db,
	}
}

// Exist
func (repo *boardRepository) Exist(model *board_entities.BoardORM) bool {
	logger.Infof("Received boardRepository.Exist request %v", *model)
	var count int64
	if model.Username != nil && len(*model.Username) > 0 {
		repo.db.Model(&board_entities.BoardORM{}).Where("username = ?", model.Username).Count(&count)
		if count > 0 {
			return true
		}
	}
	if len(model.Id.String()) > 0 {
		repo.db.Model(&board_entities.BoardORM{}).Where("id = ?", model.Id.String()).Count(&count)
		if count > 0 {
			return true
		}
	}
	if model.Email != "" {
		repo.db.Model(&board_entities.BoardORM{}).Where("email = ?", model.Email).Count(&count)
		if count > 0 {
			return true
		}
	}
	return false
}

// List
func (repo *boardRepository) List(limit, page int, sort string, model *board_entities.BoardORM) (total int64, boards []*board_entities.BoardORM, err error) {
	db := repo.db

	if limit == 0 {
		limit = 10
	}
	var offset int
	if page > 1 {
		offset = (page - 1) * limit
	} else {
		offset = 0
	}
	if sort == "" {
		sort = "created_at desc"
	}

	if model.Username != nil && len(*model.Username) > 0 {
		db = db.Where("username like ?", "%"+*model.Username+"%")
	}
	if model.FirstName != "" {
		db = db.Where("first_name like ?", "%"+model.FirstName+"%")
	}
	if model.LastName != "" {
		db = db.Where("last_name like ?", "%"+model.LastName+"%")
	}
	if model.Email != "" {
		db = db.Where("email like ?", "%"+model.Email+"%")
	}
	// enable auto preloading for `Profile`
	if err = db.Set("gorm:auto_preload", true).Order(sort).Limit(limit).Offset(offset).Find(&boards).Count(&total).Error; err != nil {
		logger.WithError(err).Error("Error in boardRepository.List")
		return
	}
	return
}

// Find by ID
func (repo *boardRepository) Get(id string) (user *board_entities.BoardORM, err error) {
	u2, err := uuid.FromString(id)
	if err != nil {
		return
	}
	user = &board_entities.BoardORM{Id: u2}
	// enable auto preloading for `Profile`
	if err = repo.db.Set("gorm:auto_preload", true).First(user).Error; err != nil && err != gorm.ErrRecordNotFound {
		logger.WithError(err).Error("Error in boardRepository.Get")
	}
	return
}

// Create
func (repo *boardRepository) Create(model *board_entities.BoardORM) error {
	if exist := repo.Exist(model); exist {
		return errors.New("board already exist")
	}
	// if err := repo.db.Set("gorm:association_autoupdate", false).Create(model).Error; err != nil {
	if err := repo.db.Create(model).Error; err != nil {
		logger.WithError(err).Error("Error in boardRepository.Create")
		return err
	}
	return nil
}

// Update TODO: Translation
func (repo *boardRepository) Update(id string, model *board_entities.BoardORM) error {
	u2, err := uuid.FromString(id)
	if err != nil {
		return err
	}
	user := &board_entities.BoardORM{
		Id: u2,
	}
	// result := repo.db.Set("gorm:association_autoupdate", false).Save(model)
	result := repo.db.Model(user).Updates(model)
	if err := result.Error; err != nil {
		logger.WithError(err).Error("Error in boardRepository.Update")
		return err
	}
	if rowsAffected := result.RowsAffected; rowsAffected == 0 {
		logger.Errorf("Error in boardRepository.Update, rowsAffected: %v", rowsAffected)
		return errors.New("no records updated, No match was found")
	}
	return nil
}

// Delete
func (repo *boardRepository) Delete(model *board_entities.BoardORM) error {
	result := repo.db.Delete(model)
	if err := result.Error; err != nil {
		logger.WithError(err).Error("Error in boardRepository.Delete")
		return err
	}
	if rowsAffected := result.RowsAffected; rowsAffected == 0 {
		logger.Errorf("Error in boardRepository.Delete, rowsAffected: %v", rowsAffected)
		return errors.New("no records deleted, No match was found")
	}
	return nil
}
