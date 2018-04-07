# userBenchmarkScraper [![GoDoc](https://godoc.org/github.com/SkyrisBactera/userBenchmarkScraper?status.svg)](https://godoc.org/github.com/SkyrisBactera/userBenchmarkScraper)
A Golang library that scrapes devices and all available info for those devices that has been recorded on [UserBenchmark](http://userbenchmark.com).

## Usage
```go
/*
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
*/
// This finds the most tested CPU. Just change it from GetCPU() to getGPU(), etc. to get other device types
parts, _ := ubscraper.GetCPU() // Returns a slice of parts with the information above
sort.Slice(parts, func(i, j int) bool { return parts[i].Samples > parts[j].Samples })
fmt.Println(parts[0])
```