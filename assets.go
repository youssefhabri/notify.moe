package main

import (
	"io/ioutil"
	"strings"

	"github.com/aerogo/aero"
	"github.com/animenotifier/notify.moe/components/js"
)

// configureAssets adds all the routes used for media assets.
func configureAssets(app *aero.Application) {
	// Script bundle
	scriptBundle := js.Bundle()

	// Service worker
	serviceWorkerBytes, err := ioutil.ReadFile("sw/service-worker.js")
	serviceWorker := string(serviceWorkerBytes)

	if err != nil {
		panic("Couldn't load service worker")
	}

	app.Get("/scripts", func(ctx *aero.Context) string {
		return ctx.JavaScript(scriptBundle)
	})

	app.Get("/scripts.js", func(ctx *aero.Context) string {
		return ctx.JavaScript(scriptBundle)
	})

	app.Get("/service-worker", func(ctx *aero.Context) string {
		return ctx.JavaScript(serviceWorker)
	})

	// Web manifest
	app.Get("/manifest.json", func(ctx *aero.Context) string {
		return ctx.JSON(app.Config.Manifest)
	})

	// Favicon
	app.Get("/favicon.ico", func(ctx *aero.Context) string {
		return ctx.TryWebP("images/brand/64", ".png")
	})

	// Brand icons
	app.Get("/images/brand/:file", func(ctx *aero.Context) string {
		return ctx.File("images/brand/" + ctx.Get("file"))
	})

	// Cover image
	app.Get("/images/cover/:file", func(ctx *aero.Context) string {
		file := strings.TrimSuffix(ctx.Get("file"), ".webp")
		return ctx.TryWebP("images/cover/"+file, ".jpg")
	})

	// Login buttons
	app.Get("/images/login/:file", func(ctx *aero.Context) string {
		return ctx.File("images/login/" + ctx.Get("file") + ".png")
	})

	// Avatars
	app.Get("/images/avatars/large/:file", func(ctx *aero.Context) string {
		return ctx.File("images/avatars/large/" + ctx.Get("file"))
	})

	// Avatars
	app.Get("/images/avatars/small/:file", func(ctx *aero.Context) string {
		return ctx.File("images/avatars/small/" + ctx.Get("file"))
	})

	// Elements
	app.Get("/images/elements/:file", func(ctx *aero.Context) string {
		return ctx.File("images/elements/" + ctx.Get("file"))
	})

	// For benchmarks
	app.Get("/hello", func(ctx *aero.Context) string {
		return ctx.Text("Hello World")
	})
}
