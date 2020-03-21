// Copyright (c) 2019-2020 The iexcloud developers. All rights reserved.
// Project site: https://github.com/goinvest/iexcloud
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE file for the project.

package cmd

import (
	"context"
	"fmt"
	"log"
	"sort"
	"strings"

	iex "github.com/goinvest/iexcloud/v2"
	"github.com/spf13/cobra"
)

// QuoteByMktCap is a wrapper type used for sorting arrays of Quotes
type QuoteByMktCap []iex.Quote

func (q QuoteByMktCap) Len() int           { return len(q) }
func (q QuoteByMktCap) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }
func (q QuoteByMktCap) Less(i, j int) bool { return q[i].MarketCap < q[j].MarketCap }

func init() {
	rootCmd.AddCommand(collectionSectorCmd)
	rootCmd.AddCommand(collectionTagCmd)
}

func printNumber(n int) string {
	s := "N/A"
	if n > 1000000000 {
		s = fmt.Sprintf("%.2fB", float32(n)/1000000000)
	} else if n > 1000000 {
		s = fmt.Sprintf("%.2fM", float32(n)/1000000)
	} else if n > 0 {
		s = fmt.Sprintf("%d", n)
	}
	return s
}

func printQuote(q iex.Quote) {
	fmt.Printf("%s : %s <Mkt Cap %s - %.2f %d>\n",
		q.Symbol, q.CompanyName, printNumber(q.MarketCap), q.LatestPrice, q.LatestVolume)
}

var collectionSectorCmd = &cobra.Command{
	Use:   "collection-sector [sector]",
	Short: "Retrieve quotes for each company in a specified sector",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		sector := strings.ToLower(args[0])
		client := GetClient()
		sectors, err := client.Sectors(context.Background())
		if err != nil {
			log.Fatalf("Error getting sectors: %s", err)
		}

		foundSectors := []iex.Sector{}
		for _, s := range sectors {
			if strings.ToLower(s.Name) == sector {
				foundSectors = []iex.Sector{s}
				break
			}
			if strings.Contains(strings.ToLower(s.Name), sector) {
				foundSectors = append(foundSectors, s)
			}
		}

		switch len(foundSectors) {
		case 0:
			log.Fatalf("Could not find specified sector")
		case 1:
			sector := foundSectors[0]
			quotes, err := client.CollectionBySector(context.Background(), sector)
			if err != nil {
				log.Fatalf("Error getting quotes for sector %s: %s", sector.Name, err)
			}
			sort.Sort(QuoteByMktCap(quotes))
			for _, q := range quotes {
				printQuote(q)
			}
			fmt.Printf("%d quotes for sector <%s>\n", len(quotes), sector.Name)
		default:
			fmt.Println("Found multiple matching sectors:")
			for _, s := range foundSectors {
				fmt.Println(s.Name)
			}
		}
	},
}

var collectionTagCmd = &cobra.Command{
	Use:   "collection-tag [tag]",
	Short: "Retrieve quotes for each company in a specified tag",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tag := strings.ToLower(args[0])
		client := GetClient()
		tags, err := client.Tags(context.Background())
		if err != nil {
			log.Fatalf("Error getting tags: %s", err)
		}

		foundTags := []iex.Tag{}
		for _, t := range tags {
			if strings.ToLower(t.Name) == tag {
				foundTags = []iex.Tag{t}
				break
			}
			if strings.Contains(strings.ToLower(t.Name), tag) {
				foundTags = append(foundTags, t)
			}
		}

		switch len(foundTags) {
		case 0:
			log.Fatalf("Could not find specified tag")
		case 1:
			tag := foundTags[0]
			quotes, err := client.CollectionByTag(context.Background(), tag)
			if err != nil {
				log.Fatalf("Error getting quotes for tag <%s>: %s", tag.Name, err)
			}
			sort.Sort(QuoteByMktCap(quotes))
			for _, q := range quotes {
				printQuote(q)
			}
			fmt.Printf("%d quotes for tag <%s>\n", len(quotes), tag.Name)
		default:
			fmt.Println("Found multiple matching tags:")
			for _, t := range foundTags {
				fmt.Println(t.Name)
			}
		}
	},
}
