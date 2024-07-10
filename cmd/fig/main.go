package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/KDreynolds/fig/internal/config"
	"github.com/KDreynolds/fig/internal/ssh"
)

func main() {
	applyCmd := flag.NewFlagSet("apply", flag.ExitOnError)
	checkCmd := flag.NewFlagSet("check", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	// Add flags for SSH user, key, and config file
	sshUser := flag.String("user", "", "SSH user")
	sshKey := flag.String("key", "", "Path to SSH private key")
	configFile := flag.String("config", "fig.yaml", "Path to configuration file")

	if len(os.Args) < 2 {
		fmt.Println("expected 'apply', 'check', or 'list' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "apply":
		applyCmd.Parse(os.Args[2:])
		applyConfig(*configFile, *sshUser, *sshKey)
	case "check":
		checkCmd.Parse(os.Args[2:])
		checkConfig(*configFile)
	case "list":
		listCmd.Parse(os.Args[2:])
		listConfigurations(*configFile)
	default:
		fmt.Println("expected 'apply', 'check', or 'list' subcommands")
		os.Exit(1)
	}
}

func applyConfig(configFile, user, keyPath string) {
	cfg, err := config.ParseConfig(configFile)
	if err != nil {
		log.Fatalf("Failed to parse config: %v", err)
	}

	client, err := ssh.NewClient(user, keyPath)
	if err != nil {
		log.Fatalf("Failed to create SSH client: %v", err)
	}

	// TODO: Implement logic to apply configuration using parsed config and SSH client
	fmt.Println("Applying configuration...")
}

func checkConfig(configFile string) {
	cfg, err := config.ParseConfig(configFile)
	if err != nil {
		log.Fatalf("Failed to parse config: %v", err)
	}

	// TODO: Implement more detailed config checking logic
	fmt.Println("Configuration is valid.")
}

func listConfigurations(configFile string) {
	cfg, err := config.ParseConfig(configFile)
	if err != nil {
		log.Fatalf("Failed to parse config: %v", err)
	}

	fmt.Println("Available configurations:")
	for _, conf := range cfg.Configurations {
		fmt.Printf("- %s\n", conf.Name)
	}
}
