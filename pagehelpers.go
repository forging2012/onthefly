package main

import (
	"github.com/hoisie/web"
)

// Create a blank HTML5 page
func NewHTML5Page(titleText string) *Page {
	page := NewPage(titleText, "<!DOCTYPE html>")
	html := page.root.AddNewTag("html")
	head := html.AddNewTag("head")
	title := head.AddNewTag("title")
	title.AddContent(titleText)
	html.AddNewTag("body")
	return page
}

// Get a function that returns a string that is the html for this page
func HTML(page *Page) func(*web.Context) string {
	return func(ctx *web.Context) string {
		return page.GetHTML()
	}
}

// Get a function that returns a string that is the css for this page
func CSS(page *Page) func(*web.Context) string {
	return func(ctx *web.Context) string {
		ctx.ContentType("css")
		return page.GetCSS()
	}
}

// Set the margins of the body
func (page *Page) SetMargin(em int) (*Tag, error) {
	tag, err := page.root.GetTag("body")
	if err == nil {
		tag.SetMargin(em)
	}
	return tag, err
}

// Set one of the css styles of the body
func (page *Page) bodyAttr(key, value string) (*Tag, error) {
	tag, err := page.root.GetTag("body")
	if err == nil {
		tag.AddStyle(key, value)
	}
	return tag, err
}

// Set the foreground and background color of the body
func (page *Page) SetColor(fgColor string, bgColor string) (*Tag, error) {
	tag, err := page.root.GetTag("body")
	if err == nil {
		tag.AddStyle("color", fgColor)
		tag.AddStyle("background-color", bgColor)
	}
	return tag, err
}

// Set the font family
func (page *Page) SetFont(fontFamily string) (*Tag, error) {
	return page.bodyAttr("font-family", fontFamily)
}

// Add a box, for testing
func (page *Page) AddBox(id string, rounded bool) (*Tag, error) {
	tag, err := page.root.GetTag("body")
	if err == nil {
		return tag.AddBox(id, rounded, "0.9em", "Speaks browser so you don't have to", "white", "black", "3em"), nil
	}
	return tag, err
}

// Used for debugging, should be a template instead
func message(title, msg string) string {
	return "<!DOCTYPE html><html><head><title>" + title + "</title></head><body style=\"margin:4em; font-family:courier; color:gray; background-color:light gray;\"><h2>" + title + "</h2><hr style=\"margin-top:-1em; margin-bottom: 2em; margin-right: 20%; text-align: left; border: 1px dotted #b0b0b0; height:1px;\"><div style=\"margin-left: 2em;\">" + msg + "</div></body></html>"
}
