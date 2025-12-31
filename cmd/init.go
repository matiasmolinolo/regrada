package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var initCmd = &cobra.Command{
	Use:   "init [path]",
	Short: "Initialize a new regrada project",
	Long: `Initialize a new regrada project in the specified directory.
If no path is provided, the current directory will be used.

Example:
  regrada init
  regrada init ./my-project`,
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := "."
		if len(args) > 0 {
			path = args[0]
		}

		force, _ := cmd.Flags().GetBool("force")
		config, _ := cmd.Flags().GetString("config")

		runInit(path, force, config)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Add flags specific to init command
	initCmd.Flags().BoolP("force", "f", false, "Force initialization even if project exists")
	initCmd.Flags().StringP("config", "c", "", "Specify a custom config file")
}

type RegradaConfig struct {
	Project  string `yaml:"project"`
	Env      string `yaml:"env"`
	Provider struct {
		Type    string `yaml:"type"`
		BaseURL string `yaml:"base_url"`
	} `yaml:"provider"`
	Capture struct {
		Requests  bool `yaml:"requests"`
		Responses bool `yaml:"responses"`
		Traces    bool `yaml:"traces"`
	} `yaml:"capture"`
}

func runInit(path string, force bool, config string) {
	fmt.Printf("🛠  Initializing regrada project in: %s\n", path)

	// Ensure the path exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Printf("📁 Creating directory: %s\n", path)
		if err := os.MkdirAll(path, 0755); err != nil {
			fmt.Printf("❌ Error creating directory: %v\n", err)
			os.Exit(1)
		}
	}

	// Determine config file name
	configFile := ".regrada.yaml"
	if config != "" {
		configFile = config
	}
	configPath := filepath.Join(path, configFile)

	// Prevent overwrite unless --force
	if _, err := os.Stat(configPath); err == nil && !force {
		fmt.Printf("⚠️  %s already exists. Use --force to overwrite.\n", configFile)
		return
	}

	// Build default config object
	defaultCfg := RegradaConfig{
		Project: filepath.Base(path),
		Env:     "local",
	}
	defaultCfg.Provider.Type = "openai"
	defaultCfg.Provider.BaseURL = "https://api.openai.com/v1"

	defaultCfg.Capture.Requests = true
	defaultCfg.Capture.Responses = true
	defaultCfg.Capture.Traces = true

	// Marshal to YAML
	data, err := yaml.Marshal(&defaultCfg)
	if err != nil {
		fmt.Printf("❌ Failed to generate config: %v\n", err)
		os.Exit(1)
	}

	// Write file
	if err := os.WriteFile(configPath, data, 0644); err != nil {
		fmt.Printf("❌ Failed to write config: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("📄 Created %s\n", configFile)
	fmt.Println("🎯 Project initialized successfully!")
}
