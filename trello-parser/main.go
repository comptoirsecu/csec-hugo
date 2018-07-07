package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"
)

type SearchResult struct {
	Options struct {
		Terms []struct {
			Text    string `json:"text"`
			Phrase  bool   `json:"phrase"`
			Partial bool   `json:"partial"`
		} `json:"terms"`
		Modifiers  []interface{} `json:"modifiers"`
		ModelTypes []string      `json:"modelTypes"`
		Partial    bool          `json:"partial"`
	} `json:"options"`
	Boards []interface{} `json:"boards"`
	Cards  []struct {
		ID     string `json:"id"`
		Badges struct {
			Votes             int `json:"votes"`
			AttachmentsByType struct {
				Trello struct {
					Board int `json:"board"`
					Card  int `json:"card"`
				} `json:"trello"`
			} `json:"attachmentsByType"`
			ViewingMemberVoted bool        `json:"viewingMemberVoted"`
			Subscribed         bool        `json:"subscribed"`
			Fogbugz            string      `json:"fogbugz"`
			CheckItems         int         `json:"checkItems"`
			CheckItemsChecked  int         `json:"checkItemsChecked"`
			Comments           int         `json:"comments"`
			Attachments        int         `json:"attachments"`
			Description        bool        `json:"description"`
			Due                interface{} `json:"due"`
			DueComplete        bool        `json:"dueComplete"`
		} `json:"badges"`
		CheckItemStates []struct {
			IDCheckItem string `json:"idCheckItem"`
			State       string `json:"state"`
		} `json:"checkItemStates"`
		Closed            bool          `json:"closed"`
		DueComplete       bool          `json:"dueComplete"`
		DateLastActivity  time.Time     `json:"dateLastActivity"`
		Desc              string        `json:"desc"`
		DescData          interface{}   `json:"descData"`
		Due               interface{}   `json:"due"`
		Email             interface{}   `json:"email"`
		IDBoard           string        `json:"idBoard"`
		IDChecklists      []string      `json:"idChecklists"`
		IDList            string        `json:"idList"`
		IDMembers         []interface{} `json:"idMembers"`
		IDMembersVoted    []interface{} `json:"idMembersVoted"`
		IDShort           int           `json:"idShort"`
		IDAttachmentCover interface{}   `json:"idAttachmentCover"`
		Labels            []interface{} `json:"labels"`
		Limits            struct {
			Attachments struct {
				PerCard struct {
					Status    string `json:"status"`
					DisableAt int    `json:"disableAt"`
					WarnAt    int    `json:"warnAt"`
				} `json:"perCard"`
			} `json:"attachments"`
			Checklists struct {
				PerCard struct {
					Status    string `json:"status"`
					DisableAt int    `json:"disableAt"`
					WarnAt    int    `json:"warnAt"`
				} `json:"perCard"`
			} `json:"checklists"`
			Stickers struct {
				PerCard struct {
					Status    string `json:"status"`
					DisableAt int    `json:"disableAt"`
					WarnAt    int    `json:"warnAt"`
				} `json:"perCard"`
			} `json:"stickers"`
		} `json:"limits"`
		IDLabels              []interface{} `json:"idLabels"`
		ManualCoverAttachment bool          `json:"manualCoverAttachment"`
		Name                  string        `json:"name"`
		Pos                   float64       `json:"pos"`
		ShortLink             string        `json:"shortLink"`
		ShortURL              string        `json:"shortUrl"`
		Subscribed            bool          `json:"subscribed"`
		URL                   string        `json:"url"`
	} `json:"cards"`
	Organizations []interface{} `json:"organizations"`
	Members       []interface{} `json:"members"`
}

type Checklist struct {
	CheckItems []struct {
		ID          string      `json:"id"`
		IDChecklist string      `json:"idChecklist"`
		Name        string      `json:"name"`
		NameData    interface{} `json:"nameData"`
		Pos         int         `json:"pos"`
		State       string      `json:"state"`
	} `json:"checkItems"`
	ID      string `json:"id"`
	IDBoard string `json:"idBoard"`
	IDCard  string `json:"idCard"`
	Name    string `json:"name"`
	Pos     int    `json:"pos"`
}

// Topic structure stores the links section of the sechebdo md
type trlink struct {
	URL   string
	Title string
}

type Topic struct {
	Title string
	Links []trlink
}

func main() {
	//log.SetOutput(ioutil.Discard)

	appId := flag.String("appid", "", "Application ID from Trello")
	token := flag.String("token", "", "Token from Trello")
	cardSearch := flag.String("card", "", "Exact title of the card so the search returns one row")
	boardId := flag.String("board", "5200f6bf8dd457fe09002ff3", "Board ID")

	flag.Parse()

	if len(*appId) == 0 {
		log.Fatal("APPId required!")
	}
	if len(*token) == 0 {
		log.Fatal("Token required! Get one from developers.trello.com")
	}
	if len(*cardSearch) == 0 {
		log.Fatal("Enter the card's name, you fool! ;)")
	}

	checklists := searchCard(*cardSearch, *boardId, *appId, *token)

	var topics []Topic
	for _, checklist := range checklists {
		topics = append(topics, parseChecklist(checklist, *appId, *token))
	}

	// Output md:
	output, err := os.Create("output.md")
	if err != nil {
		log.Fatal("output.md file is not writable")
	}
	defer output.Close()

	for _, topic := range topics {
		fmt.Fprintf(output, "* %s\n", topic.Title)

		for _, link := range topic.Links {
			fmt.Fprintf(output, "	* [%s](%s)\n", link.Title, link.URL)
		}
	}

	log.Println("Scan completed. Check output.md")

}

// Search for the card in the Trello board
//
// 	title: string to search for
//	boardId: board ID to search in
//	appId: API app id for this wonderful go tool
//	token: authorized token
func searchCard(title, boardId, appId, token string) (checklists []string) {

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
func parseChecklist(checklist string, appId, token string) (topic Topic) {
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
		if item.State == "complete" {
			continue
		}

		link := trlink{}
		link.URL = item.Name
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
	body, err := ioutil.ReadAll(resp.Body)

	r, _ := regexp.Compile("<title>([^<]+)</title>")
	m := r.FindStringSubmatch(string(body))
	title = m[1]

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
