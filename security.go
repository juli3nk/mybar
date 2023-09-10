package main

import (
	"fmt"

	"barista.run/bar"
	"barista.run/base/click"
	"barista.run/colors"
	"barista.run/outputs"
	"barista.run/pango"

	"github.com/juli3nk/barista-module-security"
)

func outputSecurity(i security.Info) bar.Output {
	if i.Status == battery.Disconnected || i.Status == battery.Unknown {
		return nil
	}

	iconName := "battery"
	if i.Status == battery.Charging {
		iconName += "-charging"
	}
	tenth := i.RemainingPct() / 10
	switch {
	case tenth == 0:
		iconName += "-outline"
	case tenth < 10:
		iconName += fmt.Sprintf("-%d0", tenth)
	}

	mainModalController.SetOutput("battery", makeIconOutput("mdi-"+iconName))
	rem := i.RemainingTime()
	out := outputs.Group()

	// First segment will be used in summary mode.
	out.Append(outputs.Pango(
		pango.Icon("mdi-"+iconName).Alpha(0.6),
		pango.Textf("%d:%02d", int(rem.Hours()), int(rem.Minutes())%60),
	).OnClick(click.Left(func() {
		mainModalController.Toggle("battery")
	})))

	// Others in detail mode.
	out.Append(outputs.Pango(
		pango.Icon("mdi-"+iconName).Alpha(0.6),
		pango.Textf("%d%%", i.RemainingPct()),
		spacer,
		pango.Textf("(%d:%02d)", int(rem.Hours()), int(rem.Minutes())%60),
	).OnClick(click.Left(func() {
		mainModalController.Toggle("battery")
	})))
	out.Append(outputs.Pango(
		pango.Textf("%4.1f/%4.1f", i.EnergyNow, i.EnergyFull),
		pango.Text("Wh").Smaller(),
	))
	out.Append(outputs.Pango(
		pango.Textf("% +6.2f", i.SignedPower()),
		pango.Text("W").Smaller(),
	))

	switch {
	case i.RemainingPct() <= 5:
		out.Urgent(true)
	case i.RemainingPct() <= 15:
		out.Color(colors.Scheme("bad"))
	case i.RemainingPct() <= 25:
		out.Color(colors.Scheme("degraded"))
	}

	return out
}
