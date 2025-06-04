package main

import (
"fmt"
"os"
"time"

"github.com/rollout/rox-go/v5/client"
"github.com/rollout/rox-go/v5/core"
"github.com/rollout/rox-go/v5/rox"
)

// Define a struct to hold feature flags
type AppConfig struct {
DashboardTheme *rox.String
}

var config AppConfig

func main() {
// Replace with your actual environment key
envKey := os.Getenv("ROX_ENV_KEY")
if envKey == "" {
fmt.Println("ERROR: Please set the environment variable ROX_ENV_KEY with your CloudBees SDK key.")
os.Exit(1)
}

// Define the feature flags with default values and allowed options
config = AppConfig{
	DashboardTheme: rox.NewString("light", []string{"light", "dark", "synthwave"}),
}

// Register your config with the SDK
rox.Register("", config)

// Initialize the SDK with optional settings
err := rox.Setup(envKey, &client.RoxOptions{
	FetchInterval:         60 * time.Second,
	RoxOptionsConfiguration: client.RoxOptionsConfiguration{
		DevModeKey: "",
	},
})
if err != nil {
	fmt.Println("Failed to initialize CloudBees Feature Management SDK:", err)
	os.Exit(1)
}

// Create context with simulated user properties
context := core.NewContext()
context.SetCustomStringProperty("user_id", "123")
context.SetCustomStringProperty("user_group", "beta-users")

// Instead of evaluating the WelcomeMessage flag, directly use the value
welcome := "Hii!"

theme := config.DashboardTheme.GetValue(context)

fmt.Println("Welcome Message:", welcome)
fmt.Println("Dashboard Theme:", theme)

// Keep running so updates can be pulled dynamically
select {}
}