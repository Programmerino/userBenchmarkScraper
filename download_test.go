package ubscraper

import (
	"os"
	"testing"
)

//(CPU|GPU|SSD|HDD|USB|RAM)
func TestCPUDownload(t *testing.T) {
	parts, err := GetCPU()
	if err != nil {
		t.Error("Error downloading CPU info! ", err)
	}
	if parts == nil || len(parts) == 0 {
		t.Error("No parts in parts slice!")
	}
}

func TestGPUDownload(t *testing.T) {
	parts, err := GetGPU()
	if err != nil {
		t.Error("Error downloading GPU info! ", err)
	}
	if parts == nil || len(parts) == 0 {
		t.Error("No parts in parts slice!")
	}
}

func TestSSDDownload(t *testing.T) {
	parts, err := GetSSD()
	if err != nil {
		t.Error("Error downloading SSD info! ", err)
	}
	if parts == nil || len(parts) == 0 {
		t.Error("No parts in parts slice!")
	}
}

func TestHDDDownload(t *testing.T) {
	parts, err := GetHDD()
	if err != nil {
		t.Error("Error downloading HDD info! ", err)
	}
	if parts == nil || len(parts) == 0 {
		t.Error("No parts in parts slice!")
	}
}
func TestUSBDownload(t *testing.T) {
	parts, err := GetUSB()
	if err != nil {
		t.Error("Error downloading USB info! ", err)
	}
	if parts == nil || len(parts) == 0 {
		t.Error("No parts in parts slice!")
	}
}

func TestRAMDownload(t *testing.T) {
	parts, err := GetRAM()
	if err != nil {
		t.Error("Error downloading RAM info! ", err)
	}
	if parts == nil || len(parts) == 0 {
		t.Error("No parts in parts slice!")
	}
}

func TestUnmarshalCSV(t *testing.T) {
	_, err := UnmarshalCSV("")
	if err == nil {
		t.Error("Didn't throw error when couldn't find file!")
	}
}

func TestGenDownload(t *testing.T) {
	_, err := genDownload("")
	if err == nil {
		t.Error("Didn't throw error when no URL was provided")
	}
	_, err = genDownload("asdfasdf")
	if err == nil {
		t.Error("Didn't throw error when invalid URL was provided")
	}
	_, err = genDownload("https://raw.githubusercontent.com/SkyrisBactera/randomFiles/master/empty.csv")
	if err == nil {
		t.Error("Didn't throw error for empty csv")
	}
	_, err = genDownload("http: //a b.com/")
	if err == nil {
		t.Error("Didn't throw error for bad URL")
	}
	os.Remove("temp.csv")
}
