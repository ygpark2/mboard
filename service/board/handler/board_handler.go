package handler

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/auth"
	"github.com/micro/micro/v3/service/errors"
	"github.com/rs/zerolog/log"
	uuid "github.com/satori/go.uuid"
	"github.com/thoas/go-funk"

	boardPB "github.com/ygpark2/mboard/service/board/proto/board"
	board_entities "github.com/ygpark2/mboard/service/board/proto/entities"
	"github.com/ygpark2/mboard/service/board/repository"
	myErrors "github.com/ygpark2/mboard/shared/errors"
)

// boardHandler struct
type boardHandler struct {
	boardRepository repository.BoardRepository
	event           *service.Event
}

// NewBoardHandler returns an instance of `BoardServiceHandler`.
func NewBoardHandler(repo repository.BoardRepository, eve *service.Event) boardPB.BoardServiceHandler {
	return &boardHandler{
		boardRepository: repo,
		event:           eve,
	}
}

func (h *boardHandler) Exist(ctx context.Context, req *boardPB.ExistRequest, rsp *boardPB.ExistResponse) error {
	log.Info().Msg("Received boardHandler.Exist request")
	/*
			google.protobuf.StringValue title = 7 [(gorm.field).tag = { size: 255 not_null: true }];
		    google.protobuf.StringValue mobile_title = 8 [(gorm.field).tag = { size: 255 not_null: true }];
		    google.protobuf.UInt32Value order = 9 [(gorm.field).tag = { not_null: true }];
		    google.protobuf.BoolValue search = 10 [(gorm.field).tag = { not_null: true }];
	*/
	model := board_entities.BoardORM{}
	model.Id = uuid.FromStringOrNil(req.Id.GetValue())
	title := req.title.GetValue()
	model.title = &title
	mobile_title := req.mobile_title.GetValue()
	model.mobile_title = &mobile_title
	model.order = req.order.GetValue()
	model.search = req.search.GetValue()

	exists := h.boardRepository.Exist(&model)
	log.Info().Msgf("user exists? %t", exists)
	rsp.Result = exists
	return nil
}

func (h *boardHandler) List(ctx context.Context, req *boardPB.ListRequest, rsp *boardPB.ListResponse) error {
	log.Info().Msg("Received boardHandler.List request")
	model := board_entities.BoardORM{}
	username := req.Username.GetValue()
	model.Username = &username
	model.FirstName = req.FirstName.GetValue()
	model.LastName = req.LastName.GetValue()
	model.Email = req.Email.GetValue()

	total, boards, err := h.boardRepository.List(int(req.Limit.GetValue()), int(req.Page.GetValue()), req.Sort.GetValue(), &model)
	if err != nil {
		return errors.NotFound("mkit.service.board.list", "Error %v", err.Error())
	}
	rsp.Total = total

	// newBoards := make([]*accountPB.User, len(boards))
	// for index, board := range boards {
	// 	tmpBoard, _ := board.ToPB(ctx)
	// 	newBoards[index] = &tmpBoard
	// 	// *newBoards[index], _ = board.ToPB(ctx) ???
	// }
	newBoards := funk.Map(boards, func(board *board_entities.BoardORM) *board_entities.BoardORM {
		tmpUser, _ := board.ToPB(ctx)
		return &tmpBoard
	}).([]*board_entities.BoardORM)

	rsp.Results = newBoards
	return nil
}

func (h *boardHandler) Get(ctx context.Context, req *boardPB.GetRequest, rsp *boardPB.GetResponse) error {
	log.Info().Msg("Received boardHandler.Get request")

	id := req.Id.GetValue()
	if id == "" {
		return myErrors.ValidationError("mkit.service.account.user.get", "validation error: Missing Id")
	}
	user, err := h.boardRepository.Get(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			rsp.Result = nil
			return nil
		}
		return myErrors.AppError(myErrors.DBE, err)
	}

	tempUser, _ := user.ToPB(ctx)
	rsp.Result = &tempUser

	return nil
}

func (h *boardHandler) Create(ctx context.Context, req *boardPB.CreateRequest, rsp *boardPB.CreateResponse) error {
	log.Info().Msg("Received boardHandler.Create request")

	model := account_entities.UserORM{}
	username := req.Username.GetValue()
	model.Username = &username
	model.FirstName = req.FirstName.GetValue()
	model.LastName = req.LastName.GetValue()
	model.Email = req.Email.GetValue()

	if err := h.boardRepository.Create(&model); err != nil {
		return myErrors.AppError(myErrors.DBE, err)
	}

	// send email (TODO: async `go h.Event.Publish(...)`)
	/*
		if err := events.Publish(ctx, &emailerPB.Message{To: req.Email.GetValue()}); err != nil {
			log.Error().Err(err).Msg("Received Event.Publish request error")
			return myErrors.AppError(myErrors.PSE, err)
		}
	*/

	return nil
}

func (h *boardHandler) Update(ctx context.Context, req *boardPB.UpdateRequest, rsp *boardPB.UpdateResponse) error {
	log.Info().Msg("Received boardHandler.Update request")
	// Identify the user
	acc, ok := auth.AccountFromContext(ctx)
	if !ok {
		return errors.Unauthorized("mkit.service.account.user.update", "A valid auth token is required")
	}
	log.Info().Msgf("Caller Account: %v", acc)

	id := req.Id.GetValue()
	if id == "" {
		return myErrors.ValidationError("mkit.service.account.user.update", "validation error: Missing Id")
	}

	model := account_entities.UserORM{}
	username := req.Username.GetValue()
	model.Username = &username
	model.FirstName = req.FirstName.GetValue()
	model.LastName = req.LastName.GetValue()
	model.Email = req.Email.GetValue()

	if err := h.boardRepository.Update(id, &model); err != nil {
		return myErrors.AppError(myErrors.DBE, err)
	}

	return nil
}

func (h *boardHandler) Delete(ctx context.Context, req *boardPB.DeleteRequest, rsp *boardPB.DeleteResponse) error {
	log.Info().Msg("Received boardHandler.Delete request")

	id := req.Id.GetValue()
	if id == "" {
		return myErrors.ValidationError("mkit.service.account.user.update", "validation error: Missing Id")
	}

	model := account_entities.UserORM{}
	model.Id = uuid.FromStringOrNil(id)

	if err := h.boardRepository.Delete(&model); err != nil {
		return myErrors.AppError(myErrors.DBE, err)
	}

	return nil
}
