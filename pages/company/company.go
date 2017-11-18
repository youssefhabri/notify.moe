package company

import (
	"net/http"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
	"github.com/animenotifier/notify.moe/utils"
)

// Get company.
func Get(ctx *aero.Context) string {
	user := utils.GetUser(ctx)
	id := ctx.Get("id")
	company, err := arn.GetCompany(id)

	if err != nil {
		return ctx.Error(http.StatusNotFound, "Company not found", err)
	}

	ctx.Data = &arn.OpenGraph{
		Tags: map[string]string{
			"og:title":     company.Name.English,
			"og:url":       "https://" + ctx.App.Config.Domain + company.Link(),
			"og:site_name": "notify.moe",
			"og:image":     company.Image,
		},
	}

	return ctx.HTML(components.CompanyPage(company, user))
}