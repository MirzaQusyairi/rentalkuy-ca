package items

import (
	"net/http"
	"rentalkuy-ca/app/middlewares"
	"rentalkuy-ca/business/items"
	controller "rentalkuy-ca/controllers"
	"rentalkuy-ca/controllers/items/request"
	"rentalkuy-ca/controllers/items/response"

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
