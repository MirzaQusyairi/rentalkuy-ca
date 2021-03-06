package photos

import (
	"net/http"
	"rentalkuy-ca/business/photos"
	controller "rentalkuy-ca/controllers"
	"rentalkuy-ca/controllers/photos/request"
	"rentalkuy-ca/controllers/photos/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PhotoController struct {
	PhotoService photos.Service
}

func NewPhotoController(service photos.Service) *PhotoController {
	return &PhotoController{
		PhotoService: service,
	}
}

func (ctrl *PhotoController) Create(c echo.Context) error {
	req := request.Photos{}

	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	data, err := ctrl.PhotoService.Create(req.ToDomain())

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomainCreatePhoto(data))

}

func (ctrl *PhotoController) Delete(c echo.Context) error {
	deletedId, _ := strconv.Atoi(c.Param("id"))

	result, err := ctrl.PhotoService.Delete(deletedId)

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controller.NewSuccessResponse(c, result)
}

func (ctrl *PhotoController) GetByID(c echo.Context) error {
	photoID, _ := strconv.Atoi(c.Param("id"))

	result, err := ctrl.PhotoService.GetByID(photoID)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controller.NewSuccessResponse(c, response.FromDomainPhoto(result))
}

func (ctrl *PhotoController) GetAllByID(c echo.Context) error {
	itemID, _ := strconv.Atoi(c.Param("id"))

	data, err := ctrl.PhotoService.GetAllByID(int(itemID))
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusNotFound, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomainListItem(data))
}
