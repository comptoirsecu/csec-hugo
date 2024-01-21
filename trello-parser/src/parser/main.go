package main

import (
	"flag"
	"log"
	"os"
	"text/template"
	"trello"
)

func main() {
	//log.SetOutput(ioutil.Discard)

	// Get Trello content
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

	checklists := trello.SearchCard(*cardSearch, *boardId, *appId, *token)

	var topics []trello.Topic
	for _, checklist := range checklists {
		topics = append(topics, trello.ParseChecklist(checklist, *appId, *token))
	}

	// Load template
	t, err := template.ParseFiles("template-sechebdo.md")

	if err != nil {
		panic(err)
	}

	// Fill template
	err = t.Execute(os.Stdout, topics)

	if err != nil {
		panic(err)
	}

	// Output md:
	output, err := os.Create("output.md")
	if err != nil {
		log.Fatal("output.md file is not writable")
	}
	defer output.Close()

	/*for _, topic := range topics {
		fmt.Fprintf(output, "* %s\n", topic.Title)

		for _, link := range topic.Links {
			fmt.Fprintf(output, "	* [%s](%s)\n", link.Title, link.URL)
		}
	}*/

	log.Println("Scan completed. Check output.md")
}
