// Step 1: Modify the line: netID := "replace with netID"  to contain your netID
// Step 2: Create the module (e.g. go mod init db)
// Step 3: Get the required packages (or run go mod tidy):
//         go get github.com/lib/pq
//         go get github.com/go-echarts/go-echarts/v2/opts
// Step 4: Connect to the NU GlobelProtect vpn
// Step 5: Run the program: go run plot_blocks.go
//         On my desktop it took arund 2.27 seconds to complete.
// Reference on echarts: https://github.com/go-echarts/go-echarts

package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	_ "github.com/lib/pq"
)

func main() {
	// Create DB pool
	netID := "eas141" //  in my case it would be: netID := "eas141"
	DB_DSN := fmt.Sprintf("postgres://%s:@129.105.248.26:5432/chicago_crimes?sslmode=disable", netID)
	// DB_DSN := "postgres://netID:@129.105.248.26:5432/chicago_crimes?sslmode=disable"

	db, err := sql.Open("postgres", DB_DSN)

	if err != nil {
		log.Fatal("Failed to open a DB connection: ", err)
	}
	defer db.Close()
	// TO DO: Write a query to get the slice of district numbers rather than "hard code" the list.
	// districts := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 14, 15, 16, 17, 18, 19, 20, 22, 24, 25, 31}
	blocks_query := "SELECT block FROM crimes GROUP BY block"
	block_function := fmt.Sprintf(blocks_query)

	block_db, err := db.Query(block_function)
	if err != nil {
		log.Fatal("Failed to execute query: ", err)
	}

	defer block_db.Close()

	var blocks []string
	for block_db.Next() {
		var block string
		err = block_db.Scan(&block)
		// fmt.Println(block)

		blocks = append(blocks, block)

	}

	// fmt.Println(blocks)
	start := time.Now()

	for _, block := range blocks {
		print(block)
		qstring := "SELECT primary_type, count(*) FROM crimes WHERE block='%s' AND primary_type in (%s,%s,%s,%s,%s,%s,%s) GROUP BY primary_type"
		userSql := fmt.Sprintf(qstring, block, "'THEFT'", "'ASSAULT'", "'ROBBERY'", "'KIDNAPPING'", "'CRIM SEXUAL ASSAULT'", "'BATTERY'", "'MURDER'")
		fmt.Println(userSql)

		rows, err := db.Query(userSql)
		if err != nil {
			log.Fatal("Failed to execute query: ", err)
		}

		defer rows.Close()

		var ptypes []string
		var counts []int
		for rows.Next() {
			var ptype string
			var count int
			err = rows.Scan(&ptype, &count)

			// Shorten label so it won't cause issues in the plot
			if ptype == "CRIM SEXUAL ASSAULT" {
				ptype = "SEX. ASSAULT"
			}

			ptypes = append(ptypes, ptype)
			counts = append(counts, count)
		}

		plot(string(block), ptypes, counts)

	}
	duration := time.Since(start)
	fmt.Println(duration)
}

func plot(block string, ptypes []string, counts []int) {
	items := make([]opts.BarData, 0)

	for _, v := range counts {
		items = append(items, opts.BarData{Value: v})
	}

	// Put data into instance
	bar := charts.NewBar()
	// set some global options like Title/Legend/ToolTip or anything else
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title: fmt.Sprintf("Violent Crimes By Category (Block %s)", block),
		// Subtitle: "Categories: THEFT, ASSAULT, ROBBERY, KIDNAPPING, CRIM SEXUAL ASSAULT, BATTERY, MURDER",
	}),
		// charts.WithXAxisOpts(opts.XAxis{
		// 	Name: "Violent Crime Categories",
		// }),
		charts.WithYAxisOpts(opts.YAxis{
			Name: "Number of Crimes",
		}),
	)

	bar.SetXAxis(ptypes). // Use the violent crime categories to label the bars
				AddSeries("Number of Crimes", items).
				SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{ // show the number of crimes on top of each bar
				Show:     true,
				Position: "top",
			}),
		)

	// Save the chart with the district number in the name
	f, _ := os.Create(fmt.Sprintf("charts/bar_block_%s.html", block))
	bar.Render(f)
}
