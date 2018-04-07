package ubscraper

import (
	"encoding/csv"
	"io"
	"os"

	"github.com/cavaliercoder/grab"
	"github.com/gocarina/gocsv"
)

// Download URLS
const (
	cpuURL = "http://www.userbenchmark.com/resources/download/csv/CPU_UserBenchmarks.csv"
	gpuURL = "http://www.userbenchmark.com/resources/download/csv/GPU_UserBenchmarks.csv"
	ssdURL = "http://www.userbenchmark.com/resources/download/csv/SSD_UserBenchmarks.csv"
	hddURL = "http://www.userbenchmark.com/resources/download/csv/HDD_UserBenchmarks.csv"
	ramURL = "http://www.userbenchmark.com/resources/download/csv/RAM_UserBenchmarks.csv"
	usbURL = "http://www.userbenchmark.com/resources/download/csv/USB_UserBenchmarks.csv"
)

/* Info:
Type							Part Number		Brand		Model		Rank	Benchmark	Samples		URL
enum (CPU|GPU|SSD|HDD|USB|RAM)	string			string		string		int		float		int			string
*/

/*
PartInfo holds the information for a particular part which can be of type CPU, GPU, SSD, HDD, USB, or RAM
*/
type PartInfo struct {
	Type      string  `csv:"Type"`        // Can be CPU, GPU, SSD, HDD, USB, or RAM
	PartNum   string  `csv:"Part Number"` // Part Number or Exact Model
	Brand     string  `csv:"Brand"`       // The brand of the product (Ex. EVGA)
	Model     string  `csv:"Model"`       // Title of item
	Rank      int     `csv:"Rank"`        // Ranking in benchmarks
	Benchmark float64 `csv:"Benchmark"`   // Average benchmark score for the item
	Samples   int     `csv:"Samples"`     // How many benchmarks were taken for the item
	URL       string  `csv:"URL"`         // The URL on UserBenchmark for the part
}

func init() {
	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)

		// Makes it not break due to issues with UserBenchmark
		r.LazyQuotes = true
		r.FieldsPerRecord = -1
		return r
	})
}

// GetCPU downloads and unmarshals all CPU information and returns a slice of the parts of type PartInfo
func GetCPU() ([]PartInfo, error) {
	return genDownload(cpuURL)
}

// GetGPU downloads and unmarshals all GPU information and returns a slice of the parts of type PartInfo
func GetGPU() ([]PartInfo, error) {
	return genDownload(gpuURL)
}

// GetSSD downloads and unmarshals all SSD information and returns a slice of the parts of type PartInfo
func GetSSD() ([]PartInfo, error) {
	return genDownload(ssdURL)
}

// GetHDD downloads and unmarshals all HDD information and returns a slice of the parts of type PartInfo
func GetHDD() ([]PartInfo, error) {
	return genDownload(hddURL)
}

// GetRAM downloads and unmarshals all RAM information and returns a slice of the parts of type PartInfo
func GetRAM() ([]PartInfo, error) {
	return genDownload(ramURL)
}

// GetUSB downloads and unmarshals all USB information and returns a slice of the parts of type PartInfo
func GetUSB() ([]PartInfo, error) {
	return genDownload(usbURL)
}

// genDownload downloads and unmarshals all a UserBenchmark link
func genDownload(url string) ([]PartInfo, error) {
	// Downloads the GPU info
	// create client
	client := grab.NewClient()
	req, _ := grab.NewRequest("./temp.csv", url)
	resp := client.Do(req)
	resp.Wait()
	if resp.Err() != nil {
		return nil, resp.Err()
	}
	// Converts the CSV into a variable of type []PartInfo for easier use
	parts, err := UnmarshalCSV("temp.csv")
	if err != nil {
		return nil, err
	}
	// Deletes the file after it has been converted
	os.Remove("temp.csv")
	// Return a slice of the parts
	return parts, nil
}

// UnmarshalCSV converts a UserBenchmark CSV into a variable of type PartInfo for easier use
func UnmarshalCSV(filename string) ([]PartInfo, error) {
	// Opens the file for reading
	clientsFile, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer clientsFile.Close()

	// Creates placeholder for the gocsv Unmarshal to place the CSV data
	partsPtr := []*PartInfo{}

	// Loads the file information into the variable
	if err := gocsv.UnmarshalFile(clientsFile, &partsPtr); err != nil {
		return nil, err
	}

	// Convert slice of pointers into slice of PartInfo for easier use
	parts := []PartInfo{}
	for i := range partsPtr {
		parts = append(parts, *partsPtr[i])
	}

	return parts, nil

}
