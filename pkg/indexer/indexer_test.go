package indexer

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

const updirs = "../../"
const sitePath = "dcs/public/dcs/"
const origIndexesPath = "original/"

func TestMain(m *testing.M) {
	err := Index(updirs + sitePath)
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(m.Run())
}

// TestBookIndex compares original/booksindex.html to dcs/public/dcs/booksindex.html
func TestBookIndex(t *testing.T) {
	ob, err := ioutil.ReadFile(updirs + sitePath + "booksindex.html")
	if err != nil {
		t.Error(err)
	}
	obStr := string(ob)
	nb, err := ioutil.ReadFile(updirs + sitePath + "booksindex.html")
	if err != nil {
		t.Error(err)
	}
	nbStr := string(nb)
	if obStr != nbStr {
		t.Errorf("Original booksindex.html and new one not equal: ")
	}
}

// TestServiceIndex compares original/servicesindex.html to dcs/public/dcs/servicesindex.html
func TestServiceIndex(t *testing.T) {
	ob, err := ioutil.ReadFile(updirs + sitePath + "servicesindex.html")
	if err != nil {
		t.Error(err)
	}
	obStr := string(ob)
	nb, err := ioutil.ReadFile(updirs + sitePath + "servicesindex.html")
	if err != nil {
		t.Error(err)
	}
	nbStr := string(nb)
	if obStr != nbStr {
		t.Errorf("Original servicesindex.html and new one not equal: ")
	}
}

// TestDailyIndex compares the daily indexes in the original folder to the ones in dcs/public/dcs/indexes
func TestDailyIndex(t *testing.T) {
	files, err := FileMatcher(updirs+origIndexesPath+"indexes", "html", nil)
	if err != nil {
		t.Error(err)
	}
	for _, file := range files {
		ob, err := ioutil.ReadFile(file)
		if err != nil {
			t.Error(err)
		}
		obStr := string(ob)
		nb, err := ioutil.ReadFile(strings.ReplaceAll(file, origIndexesPath, sitePath))
		if err != nil {
			t.Error(err)
		}
		nbStr := string(nb)
		if obStr != nbStr {
			t.Errorf("Original servicesindex.html and new one not equal: ")
		}
	}
}
