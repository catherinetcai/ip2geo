package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/catherinetcai/ip2geo/geo"
	"github.com/davecgh/go-spew/spew"
	"github.com/nranchev/go-libGeoIP"
	"github.com/spf13/cobra"
)

var KeyCDNCmd = &cobra.Command{
	Use:   "keycdn",
	Short: "Pass in a hostname or IP and get back the location",
	Run:   getKeyCDN,
}

var IPAPICmd = &cobra.Command{
	Use:   "ipapi",
	Short: "Pass in a hostname or IP and get back the location",
	Run:   getIPAPI,
}

var GeoIPCmd = &cobra.Command{
	Use:   "geoip",
	Short: "Pass in a IP and get back the location",
	Run:   getGeoIP,
}

func getKeyCDN(cmd *cobra.Command, args []string) {
	client := &http.Client{}
	if len(args) < 1 {
		log.Fatal(errors.New("Pass in an IP or hostname you doofus"))
	}
	request, err := http.NewRequest("GET", fmt.Sprintf("https://tools.keycdn.com/geo.json?host=%s", args[0]), nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	kresp := &geo.KeyCDNResponse{}
	err = json.Unmarshal(body, kresp)
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump(kresp)
}

func getIPAPI(cmd *cobra.Command, args []string) {
	client := &http.Client{}
	if len(args) < 1 {
		log.Fatal(errors.New("Pass in an IP or hostname you doofus"))
	}
	request, err := http.NewRequest("GET", fmt.Sprintf("http://ip-api.com/json/%s", args[0]), nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	ipresp := &geo.IPAPIResponse{}
	err = json.Unmarshal(body, ipresp)
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump(ipresp)
}

func getGeoIP(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		log.Fatal(errors.New("Could you like pass an IP or a hostname or something? Yeah that would be fantastic."))
	}
	file := "GeoLiteCity.dat"
	gi, err := libgeo.Load(file)
	if err != nil {
		fmt.Printf("I don't know what this is.")
	}
	if gi != nil {
		spew.Dump(gi.GetLocationByIP(args[0]))
	}
}
