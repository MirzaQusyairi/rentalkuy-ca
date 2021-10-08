package packets

import (
	"net/http"
	"rentalkuy-ca/app/middlewares"
	"rentalkuy-ca/business/packets"
	controller "rentalkuy-ca/controllers"
	"rentalkuy-ca/controllers/packets/request"
	"rentalkuy-ca/controllers/packets/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PacketController struct {
	PacketService packets.Service
}

func NewPacketController(service packets.Service) *PacketController {
	return &PacketController{
		PacketService: service,
	}
}

func (ctrl *PacketController) Create(c echo.Context) error {
	req := request.Packets{}

	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	data, err := ctrl.PacketService.Create(req.ToDomain())

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomainCreatePacket(data))

}

func (ctrl *PacketController) Update(c echo.Context) error {
	req := request.Packets{}

	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	id, _ := strconv.Atoi(c.Param("id"))
	user := middlewares.GetUser(c)

	result, err := ctrl.PacketService.Update(user.ID, id, req.ToDomain())

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomainUpdatePacket(result))
}

func (ctrl *PacketController) Delete(c echo.Context) error {
	deletedId, _ := strconv.Atoi(c.Param("id"))

	result, err := ctrl.PacketService.Delete(deletedId)

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controller.NewSuccessResponse(c, result)
}

func (ctrl *PacketController) GetByID(c echo.Context) error {
	packetID, _ := strconv.Atoi(c.Param("id"))

	result, err := ctrl.PacketService.GetByID(packetID)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controller.NewSuccessResponse(c, response.FromDomainPacket(result))
}

func (ctrl *PacketController) GetAllByID(c echo.Context) error {
	itemID, _ := strconv.Atoi(c.Param("id"))

	data, err := ctrl.PacketService.GetAllByID(int(itemID))
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusNotFound, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomainListPacket(data))
}
