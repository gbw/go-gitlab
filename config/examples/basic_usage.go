package main

import (
	"fmt"
	"log"

	"gitlab.com/gitlab-org/api/client-go"
	"gitlab.com/gitlab-org/api/client-go/config"
)

func main() {
	// Example 0: empty config
	_ = config.Empty()

	// Example 1: Using the default config location
	basicConfigExample()

	// Example 2: auto CI support
	autoCISupportExample()
}

func basicConfigExample() {
	fmt.Println("=== Basic Config Example ===")

	// Create a config with default location (~/.config/gitlab/config.yaml)
	cfg := config.New(
		config.WithOAuth2Settings(config.OAuth2Settings{
			AuthorizationFlowEnabled: true,
			CallbackServerListenAddr: ":7171",
			Browser: func(url string) error {
				fmt.Printf("Open: %s\n", url)
				return nil
			},
			ClientID:    "41d48f9422ebd655dd9cf2947d6979681dfaddc6d0c56f7628f6ada59559af1e",
			RedirectURL: "http://localhost:7171/auth/redirect",
			Scopes:      []string{"openid", "profile", "read_user", "write_repository", "api"},
		}),
	)

	// Load the configuration
	if err := cfg.Load(); err != nil {
		log.Printf("Failed to load config: %v", err)
		return
	}

	client, err := cfg.NewClient(gitlab.WithUserAgent("my-app"))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Use the client
	user, _, err := client.Users.CurrentUser()
	if err != nil {
		log.Fatalf("Failed to get current user: %v", err)
	}

	fmt.Printf("Authenticated as: %s (%s)\n", user.Name, user.Username)
}

func autoCISupportExample() {
	fmt.Println("=== Auto CI Support Example ===")

	// Create a config with precedence to CI support and falls back to the default paths for configs otherwise
	// cfg := config.New(config.WithAutoCISupport())
	cfg := config.Empty(config.WithAutoCISupport())

	// Load the configuration
	if err := cfg.Load(); err != nil {
		log.Printf("Failed to load config: %v", err)
		return
	}

	client, err := cfg.NewClient(gitlab.WithUserAgent("my-app"))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Use the client
	user, _, err := client.Users.CurrentUser()
	if err != nil {
		log.Fatalf("Failed to get current user: %v", err)
	}

	fmt.Printf("Authenticated as: %s (%s)\n", user.Name, user.Username)
}
