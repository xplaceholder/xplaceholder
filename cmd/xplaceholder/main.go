package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/xplaceholder/xplaceholder/config"
	"github.com/xplaceholder/xplaceholder/storage"

	"github.com/xplaceholder/xplaceholder/application"
)

func main() {

	log.SetFlags(0)

	logger := application.NewLogger(os.Stdout, os.Stdin)
	stderrLogger := application.NewLogger(os.Stderr, os.Stdin)
	stateBootstrap := storage.NewStateBootstrap(stderrLogger, Version)

	globals, remainingArgs, err := config.ParseArgs(os.Args)
	if err != nil {
		log.Fatalf("\n\n%s\n", err)
	}

	if globals.NoConfirm {
		logger.NoConfirm()
	}
	stateJSON, _ := json.Marshal(globals)
	stderrLogger.Println(string(stateJSON))

	// // File IO
	// fs := afero.NewOsFs()
	// afs := &afero.Afero{Fs: fs}
	newConfig := config.NewConfig(stateBootstrap, stderrLogger, afs)

	appConfig, err := newConfig.Bootstrap(globals, remainingArgs, len(os.Args))
	if err != nil {
		log.Fatalf("\n\n%s\n", err)
	}

	// // Utilities
	// envIDGenerator := helpers.NewEnvIDGenerator(rand.Reader)

}
