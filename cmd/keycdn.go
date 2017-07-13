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
	"github.com/spf13/cobra"
)

var KeyCDNCmd = &cobra.Command{
	Use:   "keycdn",
	Short: "Pass in a hostname or IP and get back the location",
	Run:   getLocation,
}

func getLocation(cmd *cobra.Command, args []string) {
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
