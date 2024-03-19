package main

import (
    "fmt"
    "log"
    "time"
    "github.com/ryanwclark1/cmd/config" // Ensure this import path matches your project's structure
)

func main() {
    // Parse CLI arguments to get initial configuration
    cliConfig := config.ParseCLIArgs()

    // Load configuration from the file and apply CLI overrides
    finalConfig := config.Load(cliConfig)

    // Print out the final configuration to verify it's loaded correctly
    fmt.Printf("Final Configuration:\n %+v\n", finalConfig)

    // Example usage of the configuration
    if finalConfig.Debug {
        log.Println("Debugging is enabled.")
    }
    log.Printf("Listening on %s:%d\n", finalConfig.HTTP.Listen, finalConfig.HTTP.Port)

    // Simulate application logic
    log.Println("Application started. Press Ctrl+C to exit.")
    for {
        time.Sleep(10 * time.Second) // Dummy loop to simulate application work
    }
}