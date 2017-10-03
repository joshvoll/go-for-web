package main

import (
	"net/url"
	"encoding/xml"
)

// new struct for the results
type SearchResult struct {
	Title string `xml:"title,attr"`
	Author string `xml:"author,attr"`
	Year string `xml:"hyr,attr"`
	ID string `xml:"owi,attr"`
}


// classify response body for response code
type ClassifySearchResponse struct {
	Results []SearchResult `xml:"works>work"`
}


// new function search for bringing back books
func Search(query string) ([]SearchResult, error) {
	// local properties
    var c ClassifySearchResponse
	var err error

	// doing the search rstuls
	resp, err := ClassifyApi("http://classify.oclc.org/classify2/Classify?summary=true&title=" + url.QueryEscape(query))

	if err != nil {
		return []SearchResult{} , err
	}
 	
	// pasring the xml error code
	err = xml.Unmarshal(resp, &c)

	return c.Results, err
}




