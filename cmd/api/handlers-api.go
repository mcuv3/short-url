package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mcuv3/demo/internal/link"
	"github.com/mcuv3/demo/internal/models"
	"github.com/mcuv3/demo/internal/validation"
)

// GetWidgetByID gets one widget by id and returns as JSON
func (app *application) AliasRedirect(w http.ResponseWriter, r *http.Request) {

	alias := chi.URLParam(r, "alias")

	l, err := app.repo.LinkService.GetLinkByShort(alias)
	if err != nil {
		app.badRequest(w, r, errors.New("unable to get link"))
		return
	}

	http.Redirect(w, r, l.FullURL, http.StatusMovedPermanently)
}

type createAliasRequest struct {
	FullURL string `json:"full_url"`
}

type createAliasResponse struct {
	Alias string `json:"alias"`
	Link  models.Link
}

func (app *application) CreateAlias(w http.ResponseWriter, r *http.Request) {
	req := createAliasRequest{}
	if err := app.readJSON(w, r, &req); err != nil {
		app.badRequest(w, r, err)
		return
	}

	if ok := validation.ValidateURL(req.FullURL); !ok {
		app.badRequest(w, r, errors.New("invalid url"))
		return
	}

	urlShort := link.FromURL(req.FullURL)
	alias := link.Encode(urlShort)

	link, err := app.repo.LinkService.CreateLink(models.CreateLinkParams{
		FullURL: req.FullURL,
		Short:   alias,
		Title:   "to_implement",
	})

	if err != nil {
		app.errorLog.Printf("error creating link: %v", err)
		app.badRequest(w, r, errors.New("unable to create link"))
		return
	}

	app.writeJSON(w, http.StatusCreated, createAliasResponse{
		Alias: fmt.Sprintf("http://localhost:5000/%s", alias),
		Link:  *link,
	})

}
