package main

import (
	"fmt"
	"log"
	"os"

	"github.com/yourusername/trading-app/internal/alpaca"
)

func main() {
	fmt.Println("Go Trading Application")

	// These would typically come from environment variables
	apiKey := os.Getenv("ALPACA_API_KEY")
	apiSecret := os.Getenv("ALPACA_API_SECRET")

	if apiKey == "" || apiSecret == "" {
		log.Fatal("API credentials not set. Please set ALPACA_API_KEY and ALPACA_API_SECRET environment variables.")
	}

	// Create a new Alpaca client (using paper trading)
	client, err := alpaca.NewClient(apiKey, apiSecret, true)
	if err != nil {
		log.Fatalf("Failed to create Alpaca client: %v", err)
	}

	// Get account information
	account, err := client.GetAccount()
	if err != nil {
		log.Fatalf("Failed to get account information: %v", err)
	}

	fmt.Printf("Account ID: %s\n", account.ID)
	fmt.Printf("Cash: $%.2f\n", account.Cash)
	fmt.Printf("Portfolio Value: $%.2f\n", account.PortfolioValue)
}
