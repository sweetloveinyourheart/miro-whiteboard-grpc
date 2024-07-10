package handlers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	pb "github.com/sweetloveinyourheart/miro-whiteboard/common/api"
	"github.com/sweetloveinyourheart/miro-whiteboard/gateway/internal/requests"
	"github.com/sweetloveinyourheart/miro-whiteboard/gateway/internal/responses"
	"github.com/sweetloveinyourheart/miro-whiteboard/gateway/internal/utils"
)

type BoardHandler struct {
	c pb.BoardServiceClient
}

type IBoardHandler interface {
	CreateBoard(ctx *fiber.Ctx) error
}

func NewBoardHandler(client *pb.BoardServiceClient) IBoardHandler {
	return &BoardHandler{
		c: *client,
	}
}

func (h *BoardHandler) CreateBoard(ctx *fiber.Ctx) error {
	user := ctx.Get("user")

	var newBoard requests.CreateBoardRequest
	unMashallError := json.Unmarshal(ctx.Body(), &newBoard)
	if unMashallError != nil {
		return ctx.Status(400).JSON(responses.AppResponse{
			Success: false,
			Message: "Bad request",
		})
	}

	// Validate request body
	if errs := utils.Validate(newBoard); len(errs) > 0 && errs[0].Error {
		validationMessage := utils.CreateValidationMessage(errs)
		return ctx.Status(400).JSON(responses.AppResponse{
			Success: false,
			Message: validationMessage,
		})
	}

	grpcData := pb.CreateBoardRequest{
		Title:       newBoard.Title,
		Description: newBoard.Description,
	}
	grpcContext := utils.CreateAuthContext(user)
	result, grpcErr := h.c.CreateBoard(grpcContext, &grpcData)
	if grpcErr != nil {
		return fiber.ErrInternalServerError
	}

	response := responses.AppResponse{
		Success: result.Success,
		Message: result.Message,
	}
	return ctx.Status(201).JSON(response)
}
