package cmd

import (
	"chrome-tabs/internal/config"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
)

//var exitCh = make(chan int)

func Execute() {
	cfg := &config.Config{}
	flag.StringVar(&cfg.RodBin, "rod-bin", "", "rod bin path")
	flag.IntVar(&cfg.Port, "port", 8787, "rod bin port")
	flag.Parse()
	app := newApp(cfg)
	go app.Start()

	gracefulExit(app)

}

func gracefulExit(app *application) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, os.Kill)

	sig := <-signalChan
	log.Printf("catch signal, %+v", sig)
	app.browser.Close()
}
