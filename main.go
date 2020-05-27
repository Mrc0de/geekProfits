package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"time"
)

// Iteration 5,203,002 of BitProphet, Fourth language
// MrC0de@geekprojex.com

// Web Server w/ crypto services
// influx market data storage
// websocket client data
// Web Client
// Coinbase (Primary Integration)
// Account Information Display
// Automatic Trading

var (
	// Globals
	Config     Configuration
	logger     *log.Logger
	WebService *httpService

	// Channels
	DebugChannel  chan string
	WWWLogChannel chan string

	// Cmdline Flags
	ConfigFile string
	Debug      bool
	Verbose    bool
)

func main() {
	flag.StringVar(&ConfigFile, "c", "geekProfits.yml", "Alternate Config (Default: geekProfits.yml)")
	flag.BoolVar(&Debug, "debug", false, "Most Verbose Output")
	flag.BoolVar(&Verbose, "v", false, "Verbose Output")
	flag.Parse()
	logger = log.New(os.Stdout, "", log.LstdFlags)
	err := Config.load(ConfigFile)
	if err != nil {
		logger.Printf("Error, Cannot Load Config: %s", err)
		os.Exit(1)
	}
	defer func() {
		if r := recover(); r != nil {
			logger.Printf("[geekProfits] [UNHANDLED_ERROR]: %s", r)
			os.Exit(1)
		}
	}()

	// Channels
	DebugChannel := make(chan string)
	WWWLogChannel := make(chan string)
	quitKey := make(chan os.Signal, 1)

	// Start
	WebService = &httpService{}
	WebService.Init()
	signal.Notify(quitKey, os.Interrupt)
	mainQuit := false
	go func() {
		for x := 0; x < 10; x++ {
			time.Sleep(500 * time.Millisecond)
			logger.Printf("Fired [%d]", x)
			DebugChannel <- "Fired!"
		}
		WWWLogChannel <- "Web isnt even alive yet."
	}()

	// Loop
	for {
		select {
		case d := <-DebugChannel:
			{
				if Debug {
					logger.Printf("[DEBUG] %s", d)
				}
			}
		case logData := <-WWWLogChannel:
			{
				logger.Printf("[WWW] %s", logData)
			}
		case <-quitKey:
			{
				// Start shutting down
				mainQuit = true
			}
		}
		if mainQuit {
			logger.Println("[geekProfits] Shutdown Finished.")
			break
		}
	}

}
