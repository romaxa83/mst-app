package v1

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"github.com/romaxa83/mst-app/gateway/config"
	"github.com/romaxa83/mst-app/gateway/internal/author/commands"
	"github.com/romaxa83/mst-app/gateway/internal/author/service"
	"github.com/romaxa83/mst-app/gateway/internal/dto"
	"github.com/romaxa83/mst-app/gateway/internal/metrics"
	"github.com/romaxa83/mst-app/gateway/internal/middlewares"
	httpErrors "github.com/romaxa83/mst-app/pkg/http_errors"
	"github.com/romaxa83/mst-app/pkg/logger"
	"github.com/romaxa83/mst-app/pkg/tracing"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

type authorsHandlers struct {
	group   *echo.Group
	log     logger.Logger
	mw      middlewares.MiddlewareManager
	cfg     *config.Config
	as      *service.AuthorService
	v       *validator.Validate
	metrics *metrics.ApiGatewayMetrics
}

func NewAuthorsHandlers(
	group *echo.Group,
	log logger.Logger,
	mw middlewares.MiddlewareManager,
	cfg *config.Config,
	as *service.AuthorService,
	v *validator.Validate,
	metrics *metrics.ApiGatewayMetrics,
) *authorsHandlers {
	return &authorsHandlers{
		group:   group,
		log:     log,
		mw:      mw,
		cfg:     cfg,
		as:      as,
		v:       v,
		metrics: metrics,
	}
}

// CreateAuthor
// @Tags Products
// @Summary Create author
// @Description Create new author
// @Accept json
// @Produce json
// @Success 201 {object} dto.CreateAuthorResponseDto
// @Router /authors [post]
func (h *authorsHandlers) CreateAuthor() echo.HandlerFunc {
	return func(c echo.Context) error {
		h.metrics.CreateAuthorHttpRequests.Inc()

		ctx, span := tracing.StartHttpServerTracerSpan(c, "authorsHandlers.CreateAuthor")
		defer span.Finish()

		dtoAuthor := &dto.CreateAuthorDto{}
		if err := c.Bind(dtoAuthor); err != nil {
			h.log.WarnMsg("Bind", err)
			h.traceErr(span, err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		dtoAuthor.ID = uuid.NewV4()
		if err := h.v.StructCtx(ctx, dtoAuthor); err != nil {
			h.log.WarnMsg("validate", err)
			h.traceErr(span, err)
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		//h.log.Warnf("CREATE_AUTHOR_DTO %+v", dtoAuthor)
		if err := h.as.Commands.CreateAuthor.Handle(ctx, commands.NewCreateAuthorCmd(dtoAuthor)); err != nil {
			h.log.WarnMsg("CreateAuthor", err)
			h.metrics.ErrorHttpRequests.Inc()
			return httpErrors.ErrorCtxResponse(c, err, h.cfg.Http.DebugErrorsResponse)
		}

		h.metrics.SuccessHttpRequests.Inc()
		return c.JSON(http.StatusCreated, dto.CreateAuthorResponseDto{ID: dtoAuthor.ID})

		//return nil
	}
}

func (h *authorsHandlers) traceErr(span opentracing.Span, err error) {
	span.SetTag("error", true)
	span.LogKV("error_code", err.Error())
	h.metrics.ErrorHttpRequests.Inc()
}
