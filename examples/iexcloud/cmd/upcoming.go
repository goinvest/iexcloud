package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	iex "github.com/goinvest/iexcloud/v2"
	"github.com/goinvest/iexcloud/v2/examples/iexcloud/domain"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

func init() {
	rootCmd.AddCommand(upcomingEventsCmd)
	rootCmd.AddCommand(upcomingEarningsCmd)
	rootCmd.AddCommand(upcomingDividendsCmd)
	rootCmd.AddCommand(upcomingSplitsCmd)
	rootCmd.AddCommand(upcomingIPOsCmd)
}

var upcomingEventsCmd = &cobra.Command{
	Use:   "upcoming-events [stock] [fullUpcomingEarnings]",
	Short: "Retrieve the upcoming events for stock symbol",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		stock := args[0]
		fullUpcomingEarnings, err := strconv.ParseBool(args[1])
		if err != nil {
			log.Fatalf("Error parsing boolean: %v", args[1])
		}
		cfg, err := domain.ReadConfig(configFileFlag)
		if err != nil {
			log.Fatalf("Error reading config file: %s", err)
		}
		client := iex.NewClient(cfg.Token, iex.WithBaseURL(cfg.BaseURL))
		events, err := client.UpcomingEvents(context.Background(), stock, fullUpcomingEarnings)
		if err != nil {
			log.Fatalf("Error getting upcoming events: %s", err)
		}
		b, err := json.MarshalIndent(events, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling into JSON: %s", err)
		}
		fmt.Println("## Upcoming Events ##")
		fmt.Println(string(b))
	},
}

var upcomingEarningsCmd = &cobra.Command{
	Use:   "upcoming-earnings [stock] [fullUpcomingEarnings]",
	Short: "Retrieve the upcoming earnings for stock symbol",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		stock := args[0]
		fullUpcomingEarnings, err := strconv.ParseBool(args[1])
		if err != nil {
			log.Fatalf("Error parsing boolean: %v", args[1])
		}
		cfg, err := domain.ReadConfig(configFileFlag)
		if err != nil {
			log.Fatalf("Error reading config file: %s", err)
		}
		client := iex.NewClient(cfg.Token, iex.WithBaseURL(cfg.BaseURL))
		events, err := client.UpcomingEarnings(context.Background(), stock, fullUpcomingEarnings)
		if err != nil {
			log.Fatalf("Error getting upcoming earnings: %s", err)
		}
		b, err := json.MarshalIndent(events, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling into JSON: %s", err)
		}
		fmt.Println("## Upcoming Earnings ##")
		fmt.Println(string(b))
	},
}

var upcomingDividendsCmd = &cobra.Command{
	Use:   "upcoming-dividends [stock]",
	Short: "Retrieve the upcoming dividends for stock symbol",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		stock := args[0]
		cfg, err := domain.ReadConfig(configFileFlag)
		if err != nil {
			log.Fatalf("Error reading config file: %s", err)
		}
		client := iex.NewClient(cfg.Token, iex.WithBaseURL(cfg.BaseURL))
		events, err := client.UpcomingDividends(context.Background(), stock)
		if err != nil {
			log.Fatalf("Error getting upcoming dividends: %s", err)
		}
		b, err := json.MarshalIndent(events, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling into JSON: %s", err)
		}
		fmt.Println("## Upcoming Dividends ##")
		fmt.Println(string(b))
	},
}

var upcomingSplitsCmd = &cobra.Command{
	Use:   "upcoming-splits [stock]",
	Short: "Retrieve the upcoming splits for stock symbol",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		stock := args[0]
		cfg, err := domain.ReadConfig(configFileFlag)
		if err != nil {
			log.Fatalf("Error reading config file: %s", err)
		}
		client := iex.NewClient(cfg.Token, iex.WithBaseURL(cfg.BaseURL))
		events, err := client.UpcomingSplits(context.Background(), stock)
		if err != nil {
			log.Fatalf("Error getting upcoming splits: %s", err)
		}
		b, err := json.MarshalIndent(events, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling into JSON: %s", err)
		}
		fmt.Println("## Upcoming Splits ##")
		fmt.Println(string(b))
	},
}

var upcomingIPOsCmd = &cobra.Command{
	Use:   "upcoming-ipos",
	Short: "Retrieve the upcoming dividends for stock symbol",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := domain.ReadConfig(configFileFlag)
		if err != nil {
			log.Fatalf("Error reading config file: %s", err)
		}
		client := iex.NewClient(cfg.Token, iex.WithBaseURL(cfg.BaseURL))
		events, err := client.UpcomingIPOs(context.Background())
		if err != nil {
			log.Fatalf("Error getting upcoming IPOs: %s", err)
		}
		b, err := json.MarshalIndent(events, "", "  ")
		if err != nil {
			log.Fatalf("Error marshaling into JSON: %s", err)
		}
		fmt.Println("## Upcoming IPOs ##")
		fmt.Println(string(b))
	},
}
