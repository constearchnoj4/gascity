// gascity - A fork of gastownhall/gascity
// Gas price monitoring and analysis tool for EVM-compatible networks
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gascity/gascity/internal/config"
	"github.com/gascity/gascity/internal/monitor"
)

const (
	appName    = "gascity"
	appVersion = "0.1.0"
)

func main() {
	// Parse command-line flags
	cfgPath := flag.String("config", "config.yaml", "Path to configuration file")
	showVersion := flag.Bool("version", false, "Print version and exit")
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()

	if *showVersion {
		fmt.Printf("%s v%s\n", appName, appVersion)
		os.Exit(0)
	}

	// Set up logger
	logger := log.New(os.Stdout, "[gascity] ", log.LstdFlags)
	if *verbose {
		logger.SetFlags(log.LstdFlags | log.Lshortfile)
	}

	logger.Printf("Starting %s v%s", appName, appVersion)

	// Load configuration
	cfg, err := config.Load(*cfgPath)
	if err != nil {
		logger.Fatalf("Failed to load configuration: %v", err)
	}

	if *verbose {
		logger.Printf("Loaded configuration from %s", *cfgPath)
	}

	// Set up context with cancellation for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle OS signals for graceful shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-sigCh
		logger.Printf("Received signal %s, shutting down...", sig)
		cancel()
	}()

	// Initialize and start the gas price monitor
	m, err := monitor.New(cfg, logger)
	if err != nil {
		logger.Fatalf("Failed to initialize monitor: %v", err)
	}

	if err := m.Run(ctx); err != nil && err != context.Canceled {
		logger.Fatalf("Monitor exited with error: %v", err)
	}

	logger.Println("Shutdown complete")
}
