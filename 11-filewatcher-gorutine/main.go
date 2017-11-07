package main

import(
	"fmt"
	"os"
	"io/ioutil"
	"time"
	"strings"
	"encoding/csv"
	"strconv"
)

// consta path to the folder
const watchedPath = "./source"

// main struct
type Invoice struct {
	Number 				string
	Amount              float64
	PurchaseOrderNumber int
	InvoiceDate         time.Time  
}



func main() {
	// we're doing a infinite loop
	for {
		// let open the folder
		d, _ := os.Open(watchedPath)
		// read the file
		files, _ := d.Readdir(-1)

		// loop throw the files are found
		for _, fi := range files {
			// build a file path with the files names
			filePath := watchedPath + "/" + fi.Name()

			// now we need to open the files itself
			f, _ := os.Open(filePath)

			// now we need to read the file content
			data, _ := ioutil.ReadAll(f)

			// close the file
			f.Close()

			// remove the file
			os.Remove(filePath)

			// let run a gorutine to store the information from the files
			go func(data string) {
				// we need to read the information from the file, we use the csv read file
				reader := csv.NewReader(strings.NewReader(data))

				// read all the records
				records, _ := reader.ReadAll()

				// loop throw the records we found
				for _, record := range records {
					// for each record we need to create a new invoice and populate it
					invoice := new(Invoice)
					// we need to add each column on the csv to the struct
					invoice.Number = record[0]
					invoice.Amount, _ = strconv.ParseFloat(record[1], 64)
					invoice.PurchaseOrderNumber, _ = strconv.Atoi(record[2])

					// conver the timestamp from string
					unixTime, _ := strconv.ParseInt(record[3], 10, 64)
					invoice.InvoiceDate = time.Unix(unixTime,0)

					// return the record to the console
					fmt.Printf("Recieve invoice  '%v' for $%.2f  and submitted ", invoice.Number, invoice.Amount)
				}

			}(string(data))
		}
	}
}




