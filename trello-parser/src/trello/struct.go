package trello

import "time"

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
