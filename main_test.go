package main

import (
	"haggle/models"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func Test_bet(t *testing.T) {

}

func Test_stoiximan(t *testing.T) {
	if _, err := scrapeSite(models.ParseSiteConfig("stoiximan")); err != nil {
		t.Error(err.Error())
	}
}

func Test_pokerstars(t *testing.T) {

}
