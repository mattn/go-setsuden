package setsuden

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"strings"
)

type Usage struct {
	Datetime string `xml:"attr"`
	Duration int    `xml:"attr"`
	Value    int    `xml:"chardata"`
}

type Peak struct {
	Start string `xml:"attr"`
	End   string `xml:"attr"`
	Value int    `xml:"chardata"`
}

type powerReport struct {
	Actual_usage    []Usage `xml:"actual_usage>usage"`
	Estimated_usage []Usage `xml:"estimated_usage>usage"`
	Instant_usage   []Usage `xml:"instant_usage>usage"`
	Peak_supply     []Peak  `xml:"peak_supply>supply"`
	Peak_demand     []Peak  `xml:"peak_demand>demand"`
}

func getUsage(typ, region, kind, date string) (pr *powerReport, err error) {
	res, err := http.Get("http://api.gosetsuden.jp/" + typ + "/" + region + "/" + kind + "/" + date + "?output=xml")
	if err != nil {
		return
	}
	defer res.Body.Close()
	pr = new(powerReport)
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	err = xml.Unmarshal(b, pr)
	if err != nil {
		return
	}
	return
}

func GetActualUsage(region string, date ...string) (pu []Usage, err error) {
	d := strings.Join(date, "/")
	if len(d) == 0 {
		d = "latest"
	}
	pr, err := getUsage("usage", region, "actual", d)
	if err != nil {
		return
	}
	pu = pr.Actual_usage
	return
}

func GetEstimatedUsage(region string, date ...string) (pu []Usage, err error) {
	d := strings.Join(date, "/")
	if len(d) == 0 {
		d = "latest"
	}
	pr, err := getUsage("usage", region, "estimated", d)
	if err != nil {
		return
	}
	pu = pr.Estimated_usage
	return
}

func GetInstantUsage(region string, date ...string) (pu []Usage, err error) {
	d := strings.Join(date, "/")
	if len(d) == 0 {
		d = "latest"
	}
	pr, err := getUsage("usage", region, "instant", d)
	if err != nil {
		return
	}
	pu = pr.Instant_usage
	return
}

func GetPeakOfSupply(region string, date ...string) (ps []Peak, err error) {
	d := strings.Join(date, "/")
	if len(d) == 0 {
		d = "today"
	}
	pr, err := getUsage("peak", region, "supply", d)
	if err != nil {
		return
	}
	ps = pr.Peak_supply
	return
}

func GetPeakOfDemand(region string, date ...string) (pd []Peak, err error) {
	d := strings.Join(date, "/")
	if len(d) == 0 {
		d = "today"
	}
	pr, err := getUsage("peak", region, "demand", d)
	if err != nil {
		return
	}
	pd = pr.Peak_demand
	return
}
