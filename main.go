package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"github.com/google/uuid"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DatabaseUsername string `envconfig:"DATABASE_USERNAME"`
	DatabasePassword string `envconfig:"DATABASE_PASSWORD"`
	DatabaseName     string `envconfig:"DATABASE_NAME"`
	DatabaseAddress  string `envconfig:"DATABASE_ADDRESS"`
}

type Widget struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusBadRequest,
		StatusText:     "Bad Request",
		ErrorText:      err.Error(),
	}
}

func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusInternalServerError,
		StatusText:     "Internal Server Error",
		ErrorText:      err.Error(),
	}
}

type Server struct {
	db *pg.DB
}

func NewServer(db *pg.DB) Server {
	return Server{
		db: db,
	}
}

type WidgetRequest struct {
	*Widget
}

func (w *Widget) Bind(r *http.Request) error {
	if w == nil {
		return errors.New("missing required widget fields")
	}

	return nil
}

type WidgetResponse struct {
	*Widget
}

func NewWidgetResponse(widget *Widget) *WidgetResponse {
	return &WidgetResponse{Widget: widget}
}

func (rd *WidgetResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewWidgetListResponse(widgets []*Widget) []render.Renderer {
	list := []render.Renderer{}
	for _, widget := range widgets {
		list = append(list, NewWidgetResponse(widget))
	}
	return list
}

func (s *Server) ListWidgets(w http.ResponseWriter, r *http.Request) {
	var widgets []*Widget
	err := s.db.Model(&widgets).Select()
	if err != nil {
		_ = render.Render(w, r, ErrRender(err))
		return
	}

	if err := render.RenderList(w, r, NewWidgetListResponse(widgets)); err != nil {
		_ = render.Render(w, r, ErrRender(err))
		return
	}
}

func (s *Server) CreateWidget(w http.ResponseWriter, r *http.Request) {
	data := &WidgetRequest{}
	if err := render.Bind(r, data); err != nil {
		_ = render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	widget := data.Widget
	widget.ID = uuid.New().String()
	_, err := s.db.Model(widget).Insert()
	if err != nil {
		_ = render.Render(w, r, ErrRender(err))
		return
	}

	render.Status(r, http.StatusCreated)
	_ = render.Render(w, r, NewWidgetResponse(widget))
}

func main() {
	var c Config
	err := envconfig.Process("api", &c)
	if err != nil {
		fmt.Printf("problem parsing config: %s\n", err.Error())
		os.Exit(1)
	}

	db := pg.Connect(&pg.Options{
		User:     c.DatabaseUsername,
		Password: c.DatabasePassword,
		Database: c.DatabaseName,
		Addr:     c.DatabaseAddress,
	})
	defer db.Close()

	err = createSchema(db)
	if err != nil {
		fmt.Printf("problem creating schema: %s\n", err.Error())
		os.Exit(1)
	}

	srv := NewServer(db)

	r := chi.NewRouter()
	// r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	// r.Use(middleware.Recoverer)
	// r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Route("/widgets", func(r chi.Router) {
		r.Get("/", srv.ListWidgets)
		r.Post("/", srv.CreateWidget)
	})

	if err := http.ListenAndServe(":3000", r); err != nil {
		panic(err)
	}
}

func createSchema(db *pg.DB) error {
	models := []any{
		(*Widget)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
