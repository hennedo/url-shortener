package main

import "github.com/go-bongo/bongo"

type Link struct {
	bongo.DocumentBase `bson:",inline"`
	Name string
	Url string
	Scam bool
	Password string
	ClicksFacebook int `bson:"clicksFacebook"`
	ClicksInstagram int `bson:"clicksInstagram"`
	ClicksOther int `bson:"clicksOther"`
	ClicksNone int `bson:"clicksNone"`
	Clicks int `bson:"clicks"`
}

