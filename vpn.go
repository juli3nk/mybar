package main

import (
	"barista.run/bar"
	"barista.run/outputs"
	"barista.run/pango"

	"github.com/juli3nk/barista-module-vpn"
)

func outputVpn(s vpn.State) bar.Output {
	iconName := "lock"
	if s.Disconnected() {
		iconName += "-off"
	}
	return outputs.Pango(
		pango.Icon("mdi-"+iconName).Alpha(0.8),
	)
}
