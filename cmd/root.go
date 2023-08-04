package cmd

import (
	"chrome-tabs/internal/config"
	"flag"
)

func Execute() {
	cfg := &config.Config{}
	flag.StringVar(&cfg.RodBin, "rod-bin", "", "rod bin path")
	flag.IntVar(&cfg.Port, "port", 8787, "rod bin port")
	flag.Parse()
	app := newApp(cfg)
	app.Start()
}
