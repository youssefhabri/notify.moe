package statistics

import (
	"fmt"
	"strings"

	"github.com/aerogo/aero"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/components"
)

type stats map[string]float64

// Get ...
func Get(ctx *aero.Context) string {
	pieCharts := getUserStats()
	return ctx.HTML(components.Statistics(pieCharts))
}

func getUserStats() []*arn.PieChart {
	screenSize := stats{}
	pixelRatio := stats{}
	browser := stats{}
	country := stats{}
	gender := stats{}
	os := stats{}
	notifications := stats{}
	avatar := stats{}
	ip := stats{}
	pro := stats{}

	for info := range arn.StreamAnalytics() {
		user, err := arn.GetUser(info.UserID)
		arn.PanicOnError(err)

		if !user.IsActive() {
			continue
		}

		pixelRatio[fmt.Sprintf("%.0f", info.Screen.PixelRatio)]++

		size := arn.ToString(info.Screen.Width) + " x " + arn.ToString(info.Screen.Height)
		screenSize[size]++
	}

	for user := range arn.StreamUsers() {
		if !user.IsActive() {
			continue
		}

		if user.Gender != "" && user.Gender != "other" {
			gender[user.Gender]++
		}

		if user.Browser.Name != "" {
			browser[user.Browser.Name]++
		}

		if user.Location.CountryName != "" {
			country[user.Location.CountryName]++
		}

		if user.OS.Name != "" {
			if strings.HasPrefix(user.OS.Name, "CrOS") {
				user.OS.Name = "Chrome OS"
			}

			os[user.OS.Name]++
		}

		if len(user.PushSubscriptions().Items) > 0 {
			notifications["Enabled"]++
		} else {
			notifications["Disabled"]++
		}

		if user.Avatar.Source == "" {
			avatar["none"]++
		} else {
			avatar[user.Avatar.Source]++
		}

		if arn.IsIPv6(user.IP) {
			ip["IPv6"]++
		} else {
			ip["IPv4"]++
		}

		if user.IsPro() {
			pro["PRO accounts"]++
		} else {
			pro["Free accounts"]++
		}
	}

	return []*arn.PieChart{
		arn.NewPieChart("OS", os),
		arn.NewPieChart("Screen size", screenSize),
		arn.NewPieChart("Browser", browser),
		arn.NewPieChart("Country", country),
		arn.NewPieChart("Avatar", avatar),
		arn.NewPieChart("Notifications", notifications),
		arn.NewPieChart("Gender", gender),
		arn.NewPieChart("Pixel ratio", pixelRatio),
		arn.NewPieChart("IP version", ip),
		arn.NewPieChart("PRO accounts", pro),
	}
}
