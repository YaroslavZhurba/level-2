package main

import (
	"fmt"
)

type Cms int64

const (
	WORDPRESS Cms = iota
	ALFREDO
)

type Website struct {
	name  string
	cms   Cms
	price int
}

type WebsiteBuilder interface {
	CreateWebsite()
	BuildName()
	BuildCms()
	BuildPrice()
	GetWebsite() *Website
}

// implementation of WebsiteBuilder
type VisitCardWebsiteBuilder struct {
	website *Website
}

func (ws *VisitCardWebsiteBuilder) CreateWebsite() {
	ws.website = &Website{}
}

func (ws *VisitCardWebsiteBuilder) BuildName() {
	ws.website.name = "Visit card"
}

func (ws *VisitCardWebsiteBuilder) BuildCms() {
	ws.website.cms = WORDPRESS
}

func (ws *VisitCardWebsiteBuilder) BuildPrice() {
	ws.website.price = 500
}

func (ws *VisitCardWebsiteBuilder) GetWebsite() *Website {
	return ws.website
}

// implementation of WebsiteBuilder
type EnterpriseWebsiteBuilder struct {
	website *Website
}

func (ws *EnterpriseWebsiteBuilder) CreateWebsite() {
	ws.website = &Website{}
}

func (ws *EnterpriseWebsiteBuilder) BuildName() {
	ws.website.name = "Enterprise web site"
}

func (ws *EnterpriseWebsiteBuilder) BuildCms() {
	ws.website.cms = ALFREDO
}

func (ws *EnterpriseWebsiteBuilder) BuildPrice() {
	ws.website.price = 7500
}

func (ws *EnterpriseWebsiteBuilder) GetWebsite() *Website {
	return ws.website
}

// Director
type Director struct {
	websiteBuilder WebsiteBuilder
}

func (director *Director) SetBuilder(builder WebsiteBuilder) {
	director.websiteBuilder = builder
}

func (director *Director) BuildWebsite() *Website {
	if director.websiteBuilder == nil {
		return nil
	}
	director.websiteBuilder.CreateWebsite()
	director.websiteBuilder.BuildName()
	director.websiteBuilder.BuildCms()
	director.websiteBuilder.BuildPrice()
	return director.websiteBuilder.GetWebsite()
}

func main() {
	director := Director{}
	director.SetBuilder(&VisitCardWebsiteBuilder{})
	// director.SetBuilder(&EnterpriseWebsiteBuilder{})
	site := director.BuildWebsite()
	fmt.Println(site)
}
