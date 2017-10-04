/* crud for the bookings using sqlite 3 */
package main

import (
	"net/http"
	"io/ioutil"
	"net/url"
	"encoding/xml"
	"fmt"
)

// defingin the struct or interfaces
type ClassifyBookResponse struct {
  BookData struct {
    Title string `xml:"title,attr"`
    Author string `xml:"author,attr"`
    ID string `xml:"owi,attr"`
  } `xml:"work"`
  Classification struct {
    MostPopular string `xml:"sfa,attr"`
  } `xml:"recommendations>ddc>mostPopular"`
}

/*
	Method: ClassifyAPI
	Gets: url as string
	Return: bytes slice and error	
	public
*/
func ClassifyApi(url string) ([]byte,  error) {
	// get url information
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("ERROR URL: ", err)
		return []byte{}, err
	}

	// defer
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

// find one api
func Find(id string) (ClassifyBookResponse, error) {

	var c ClassifyBookResponse

	fmt.Println("GETTER: ", id)

    body, err := ClassifyApi("http://classify.oclc.org/classify2/Classify?summary=true&owi=" + url.QueryEscape(id))

    if err != nil {
    	fmt.Println("ERROR ON FIND: ", err)
    	return ClassifyBookResponse{}, err
    }

 	err = xml.Unmarshal(body, &c)
    return c, err
}



