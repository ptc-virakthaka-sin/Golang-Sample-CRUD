package gracefulshutdown

import (
	"fmt"
	"learn-fiber/config"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
)

var defaultSignals = []os.Signal{
	os.Interrupt,
	syscall.SIGINT,
	syscall.SIGTERM,
}

/*
Credit
https://github.com/gkampitakis/fiber-modules/tree/v1.1.10/gracefulshutdown
*/
func Listen(app *fiber.App, cleanupFns []func() error) {
	// _signals := defaultSignals
	// if signals != nil {
	// 	_signals = signals
	// }

	c := make(chan os.Signal, 1)
	signal.Notify(c, defaultSignals...)

	go func() {
		_ = <-c // This blocks the main thread until an interrupt is received
		fmt.Println("Gracefully shutting down...")
		_ = app.Shutdown()
	}()

	addr := getServerAddr()
	if err := app.Listen(addr); err != nil {
		log.Panic(err)
	}

	fmt.Println("Running cleanup tasks...")
	for _, fn := range cleanupFns {
		executeFn(fn)
	}
	fmt.Println("Server was successful shutdown.")
}

func getServerAddr() string {
	host := config.Cfg.Host
	port := config.Cfg.Port

	if host != "" && port != "" {
		return fmt.Sprintf("%s:%s", host, port)
	} else if port != "" {
		return fmt.Sprintf(":%s", port)
	}
	// default port
	return ":3000"
}

func executeFn(fn func() error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	if err := fn(); err != nil {
		log.Println(err)
	}
}
