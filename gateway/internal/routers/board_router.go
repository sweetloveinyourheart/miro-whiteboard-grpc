package routers

import (
	"github.com/gofiber/fiber/v2"
	pb "github.com/sweetloveinyourheart/miro-whiteboard/common/api"
	"github.com/sweetloveinyourheart/miro-whiteboard/gateway/internal/handlers"
	"github.com/sweetloveinyourheart/miro-whiteboard/gateway/internal/middlewares"
)

func CreateBoardRouters(r fiber.Router, c *pb.BoardServiceClient) {
	routes := r.Group("/boards")
	boardHandler := handlers.NewBoardHandler(c)

	routes.Get("/get-board/:id", middlewares.AuthGuard, boardHandler.GetBoardById)
	routes.Post("/new", middlewares.AuthGuard, boardHandler.CreateBoard)
	routes.Delete("/terminate/:id", middlewares.AuthGuard, boardHandler.DeleteBoard)
}
