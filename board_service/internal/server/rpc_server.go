package server

import (
	"context"
	"fmt"

	"github.com/sweetloveinyourheart/miro-whiteboard/board_service/internal/services"
	pb "github.com/sweetloveinyourheart/miro-whiteboard/common/api"
	"github.com/sweetloveinyourheart/miro-whiteboard/common/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	pb.UnimplementedBoardServiceServer
	svc services.IBoardService
}

func CreateUserServer(client *mongo.Client) *Server {
	return &Server{
		svc: services.NewBoardService(client),
	}
}

func (s *Server) CreateBoard(ctx context.Context, in *pb.CreateBoardRequest) (*pb.CreateBoardResponse, error) {
	user, err := utils.GetAuthorizedUser(ctx)
	if err != nil {
		return nil, fmt.Errorf("unauthorized")
	}

	board := services.BoardInfo{
		Title:       in.Title,
		Description: in.Description,
		CreatedBy:   user,
	}

	success, err := s.svc.CreateBoard(board)
	if err != nil {
		return nil, err
	}

	response := pb.CreateBoardResponse{
		Success: success,
		Message: "New board created",
	}

	return &response, nil
}

func (s *Server) GetBoardById(ctx context.Context, in *pb.GetBoardByIdRequest) (*pb.BoardResponse, error) {
	user, err := utils.GetAuthorizedUser(ctx)
	if err != nil {
		return nil, fmt.Errorf("unauthorized")
	}

	board, err := s.svc.GetBoardById(user, in.Id)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	response := pb.BoardResponse{
		Id:          board.ID.String(),
		Title:       board.Title,
		Description: board.Description,
		CreatedBy:   board.CreatedBy,
		CreatedAt:   board.CreatedAt,
		UpdatedAt:   board.UpdatedAt,
	}

	return &response, nil
}

func (s *Server) DeleteBoard(ctx context.Context, in *pb.DeleteBoardRequest) (*pb.DeleteBoardResponse, error) {
	user, err := utils.GetAuthorizedUser(ctx)
	if err != nil {
		return nil, fmt.Errorf("unauthorized")
	}

	success, err := s.svc.DeleteBoard(user, in.Id)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	response := pb.DeleteBoardResponse{
		Success: success,
		Message: "board deleted",
	}

	return &response, nil
}
