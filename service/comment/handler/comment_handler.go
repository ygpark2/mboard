package handler

import (
	"context"

	"github.com/micro/micro/v3/service/logger"

	comment_entities "github.com/ygpark2/mboard/service/comment/proto/entities"
	commentPB "github.com/ygpark2/mboard/service/comment/proto/comment"
	"github.com/ygpark2/mboard/service/comment/repository"
)

type CommentHandler struct{}

// Call is a single request handler called via client.Call or the generated client code
func (c *CommentHandler) Save(ctx context.Context, req *commentPB.Request, rsp *commentPB.Response) error {
	logger.Info("Not yet implemented")
	return nil
}
