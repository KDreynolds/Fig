package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/KDreynolds/fig/internal/config"
	"github.com/KDreynolds/fig/internal/ssh"
	"github.com/KDreynolds/fig/internal/template"
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

	engine := template.New()

	// Example: Apply the first configuration
	if len(cfg.Configurations) > 0 {
		conf := cfg.Configurations[0]
		fmt.Printf("Applying configuration: %s\n", conf.Name)

		for _, serverGroupName := range conf.Servers {
			serverGroup, err := cfg.GetServerGroup(serverGroupName)
			if err != nil {
				log.Fatalf("Failed to get server group: %v", err)
			}

			for _, host := range serverGroup.Hosts {
				for _, taskName := range conf.Tasks {
					task, err := cfg.GetTask(taskName)
					if err != nil {
						log.Fatalf("Failed to get task: %v", err)
					}

					renderedCommand, err := engine.RenderTask(task.Command, task.Vars, cfg.GlobalVars)
					if err != nil {
						log.Fatalf("Failed to render task command: %v", err)
					}

					output, err := client.RunCommand(host, renderedCommand)
					if err != nil {
						log.Fatalf("Failed to run command on %s: %v", host, err)
					}

					fmt.Printf("Output from %s:\n%s\n", host, output)
				}
			}
		}
	}
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
