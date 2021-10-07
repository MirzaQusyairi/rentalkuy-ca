package items

import (
	"net/http"
	"rentalkuy-ca/app/middlewares"
	"rentalkuy-ca/business/items"
	controller "rentalkuy-ca/controllers"
	"rentalkuy-ca/controllers/items/request"
	"rentalkuy-ca/controllers/items/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ItemController struct {
	ItemService items.Service
}

func NewItemController(service items.Service) *ItemController {
	return &ItemController{
		ItemService: service,
	}
}

func (ctrl *ItemController) Create(c echo.Context) error {
	req := request.Items{}

	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	user := middlewares.GetUser(c) //get id from jwt
	data, err := ctrl.ItemService.Create(user.ID, req.ToDomain())

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomainCreateItem(data))

}

func (ctrl *ItemController) Update(c echo.Context) error {
	req := request.Items{}

	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	id, _ := strconv.Atoi(c.Param("id"))
	user := middlewares.GetUser(c)

	result, err := ctrl.ItemService.Update(user.ID, id, req.ToDomain())

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomainUpdateItem(result))
}

func (ctrl *ItemController) Delete(c echo.Context) error {
	userID := middlewares.GetUser(c)
	deletedId, _ := strconv.Atoi(c.Param("id"))

	result, err := ctrl.ItemService.Delete(userID.ID, deletedId)

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controller.NewSuccessResponse(c, result)
}

func (ctrl *ItemController) GetByID(c echo.Context) error {
	itemID, _ := strconv.Atoi(c.Param("id"))

	result, err := ctrl.ItemService.GetByID(itemID)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controller.NewSuccessResponse(c, response.FromDomainItem(result))
}

func (ctrl *ItemController) GetAllByUserID(c echo.Context) error {
	userID := middlewares.GetUser(c).ID

	data, err := ctrl.ItemService.GetAllByUserID(int(userID))
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusNotFound, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomainListItem(data))
}
