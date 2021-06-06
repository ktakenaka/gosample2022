package sample

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/ktakenaka/gomsx/app/internal/models/v1.0/dao"
	samplesUC "github.com/ktakenaka/gomsx/app/internal/usecases/samples"
)

type sampleHandler struct {
	useCase samplesUC.UseCase
}

func NewSampleHandler(ctx context.Context, txmFact dao.TxManagerFactory) *sampleHandler {
	return &sampleHandler{
		useCase: samplesUC.NewUseCase(ctx, txmFact),
	}
}

func (hdl *sampleHandler) List(ctx echo.Context) error {
	list, err := hdl.useCase.ListSamples(ctx.Request().Context())
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, list)
}
