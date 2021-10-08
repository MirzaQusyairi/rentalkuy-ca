package rents

import (
	"net/http"
	"rentalkuy-ca/app/middlewares"
	"rentalkuy-ca/business/rents"
	controller "rentalkuy-ca/controllers"
	"rentalkuy-ca/controllers/rents/request"
	"rentalkuy-ca/controllers/rents/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

type RentController struct {
	RentService rents.Service
}

func NewRentController(service rents.Service) *RentController {
	return &RentController{
		RentService: service,
	}
}

func (ctrl *RentController) Create(c echo.Context) error {
	req := request.Rents{}

	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	user := middlewares.GetUser(c) //get id from jwt
	data, err := ctrl.RentService.Create(user.ID, req.ToDomain())

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomainCreateRent(data))

}

// func (ctrl *RentController) Update(c echo.Context) error {
// 	req := request.Rents{}

// 	if err := c.Bind(&req); err != nil {
// 		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
// 	}

// 	id, _ := strconv.Atoi(c.Param("id"))
// 	user := middlewares.GetUser(c)

// 	result, err := ctrl.RentService.Update(user.ID, id, req.ToDomain())

// 	if err != nil {
// 		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
// 	}

// 	return controller.NewSuccessResponse(c, response.FromDomainUpdateRent(result))
// }

func (ctrl *RentController) Delete(c echo.Context) error {
	userID := middlewares.GetUser(c)
	deletedId, _ := strconv.Atoi(c.Param("id"))

	result, err := ctrl.RentService.Delete(userID.ID, deletedId)

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controller.NewSuccessResponse(c, result)
}

func (ctrl *RentController) GetByID(c echo.Context) error {
	itemID, _ := strconv.Atoi(c.Param("id"))

	result, err := ctrl.RentService.GetByID(itemID)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controller.NewSuccessResponse(c, response.FromDomainItem(result))
}

func (ctrl *RentController) GetAllByUserID(c echo.Context) error {
	userID := middlewares.GetUser(c).ID

	data, err := ctrl.RentService.GetAllByUserID(int(userID))
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusNotFound, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomainListItem(data))
}
