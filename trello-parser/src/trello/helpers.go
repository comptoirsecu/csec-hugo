package trello

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

// Search for the card in the Trello board
//
// 	title: string to search for
//	boardId: board ID to search in
//	appId: API app id for this wonderful go tool
//	token: authorized token
func SearchCard(title, boardId, appId, token string) (checklists []string) {

	endpoint := fmt.Sprintf("https://api.trello.com/1/search?query=%%22%s%%22&idBoards=%s&modelTypes=all&board_fields=name%%2CidOrganization&boards_limit=10&card_fields=all&cards_limit=10&cards_page=0&card_attachments=false&organization_fields=name%%2CdisplayName&organizations_limit=10&member_fields=avatarHash%%2CfullName%%2Cinitials%%2Cusername%%2Cconfirmed&members_limit=10&key=%s&token=%s", url.QueryEscape(title), url.QueryEscape(boardId), url.QueryEscape(appId), url.QueryEscape(token))

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode > 200 {
		log.Fatal("Caught HTTP Status ", resp.StatusCode, " / ", resp.Status)
	}

	// Peprare structure to store the result
	var record SearchResult

	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Fatalln(err)
	}

	c := len(record.Cards)

	if c == 0 {
		log.Fatal("No card found for search: ", title)
	}

	if c > 1 {
		// When more than one result is found, propose to pick-up the correct ID
		fmt.Println(c, " results found. Please select:")
		for _, card := range record.Cards {
			fmt.Println(card.ID, "\t", card.Name)
		}

		var cardID string
		fmt.Print("Enter the selected ID: ")
		fmt.Scanln(&cardID)

		// Find the checklists for that ID
		for _, card := range record.Cards {
			if card.ID == cardID {
				checklists = card.IDChecklists
				return
			}
		}
	}

	// Only one result, return the card ID
	checklists = record.Cards[0].IDChecklists

	return
}

// Retrieves all items but those checked from a checklist
//
//	checklist: ID of the checklist
//	appId: API app id for this wonderful go tool
//	token: authorized token
func ParseChecklist(checklist string, appId, token string) (topic Topic) {
	endpoint := fmt.Sprintf("https://api.trello.com/1/checklists/%s?fields=name&cards=all&card_fields=name&key=%s&token=%s", url.QueryEscape(checklist), url.QueryEscape(appId), url.QueryEscape(token))

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode > 200 {
		log.Fatal("Caught HTTP Status ", resp.StatusCode, " / ", resp.Status)
	}

	// Peprare structure to store the result
	var record Checklist

	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Fatalln(err)
	}

	if len(record.CheckItems) == 0 {
		fmt.Println("No item in checklist ", record.Name)
		return
	}

	topic.Title = cleanTopic(record.Name) // Title of this chapter of sechebdo
	log.Println("New topic: ", topic.Title)

	// Iterates over all checklists
	for _, item := range record.CheckItems {
		// Ignore checked items
		if !strings.HasPrefix(item.Name, "http") {
			continue
		}

		link := trlink{}
		link.URL = item.Name
		log.Println("Fetching title for: ", link.URL)
		link.Title = getHtmlTitle(link.URL)

		topic.Links = append(topic.Links, link)
		log.Println("Added link: ", item.Name)
	}

	return
}

// Clear the speaker's name often displaid on the title
// e.g. Morgan : super sujet de la mort qui tue
func cleanTopic(topic string) string {
	r, _ := regexp.Compile("[-:](.*)")
	m := r.FindStringSubmatch(topic)

	if m == nil {
		return topic
	}
	return m[1]
}

// Parse a web page to retrieve its HTML title tag content
func getHtmlTitle(url string) (title string) {
	// Exception cases:
	if strings.HasSuffix(url, ".pdf") {
		title = "PDF Document"
		return
	}

	if !strings.HasPrefix(url, "http") {
		log.Println("Item is not an url: ", url)
	}

	resp, err := http.Get(url)
	if err != nil {
		log.Println("Could nto fetch: ", url)
		title = url
		return
	}
	defer resp.Body.Close()

	// Check content-type for surprises
	if !strings.Contains(resp.Header.Get("Content-Type"), "text/html") {
		log.Println("Not html resource, cannot get title")
		title = url
		return
	}

	body, err := ioutil.ReadAll(resp.Body)

	r, _ := regexp.Compile("<title>([^<]+)</title>")
	m := r.FindStringSubmatch(string(body))
	if len(m) > 0 {
		title = m[1]
	}

	// Revert html entities
	title = html.UnescapeString(m[1])

	// Remove new lines (thank you YouTube for your ugly source code)
	title = strings.Replace(title, "\n", "", -1)

	// Twitter's title includes the whole tweet. Let's strip it
	if strings.Contains(url, "twitter.com") {
		title = strings.Split(title, ":")[0]
	}

	title = strings.Trim(title, " ")

	return
}
