package registry

import (
	board_entities "github.com/ygpark2/mboard/service/board/proto/entities"
	"github.com/ygpark2/mboard/service/board/repository"
	"github.com/ygpark2/mboard/shared/database"
	configPB "github.com/ygpark2/mboard/shared/proto/config"
)

func buildBoardRepository(cfg configPB.Configuration) (interface{}, error) {
	// db := database.GetDatabaseConnection(*cfg.Database).(*gorm.DB)

	db, err := database.GetDatabaseConnection(*cfg.Database)
	if err != nil {
		db.AutoMigrate(&board_entities.BoardORM{})
	}

	return repository.NewBoardRepository(db), nil
}
