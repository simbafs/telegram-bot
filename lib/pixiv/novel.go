package pixiv

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var (
	ErrIDEmpty = errors.New("ID is empty")
)

type Novel struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Body    struct {
		BookmarkCount int       `json:"bookmarkCount"`
		CommentCount  int       `json:"commentCount"`
		MarkerCount   int       `json:"markerCount"`
		CreateDate    time.Time `json:"createDate"`
		UploadDate    time.Time `json:"uploadDate"`
		Description   string    `json:"description"`
		ID            string    `json:"id"`
		Title         string    `json:"title"`
		LikeCount     int       `json:"likeCount"`
		PageCount     int       `json:"pageCount"`
		UserID        string    `json:"userId"`
		UserName      string    `json:"userName"`
		ViewCount     int       `json:"viewCount"`
		IsOriginal    bool      `json:"isOriginal"`
		IsBungei      bool      `json:"isBungei"`
		XRestrict     int       `json:"xRestrict"`
		Restrict      int       `json:"restrict"`
		Content       string    `json:"content"`
		CoverURL      string    `json:"coverUrl"`
		LikeData      bool      `json:"likeData"`
		Tags          struct {
			AuthorID string `json:"authorId"`
			IsLocked bool   `json:"isLocked"`
			Tags     []struct {
				Tag       string `json:"tag"`
				Locked    bool   `json:"locked"`
				Deletable bool   `json:"deletable"`
				UserID    string `json:"userId"`
				UserName  string `json:"userName"`
			} `json:"tags"`
			Writable bool `json:"writable"`
		} `json:"tags"`
		IsUnlisted     bool   `json:"isUnlisted"`
		Language       string `json:"language"`
		CommentOff     int    `json:"commentOff"`
		CharacterCount int    `json:"characterCount"`
		WordCount      int    `json:"wordCount"`
		UseWordCount   bool   `json:"useWordCount"`
		ReadingTime    int    `json:"readingTime"`
		AiType         int    `json:"aiType"`
	} `json:"body"`
}

func GetNovelByID(ID string) (*Novel, error) {
	if strings.Trim(ID, " ") == "" {
		return nil, ErrIDEmpty
	}

	data, err := http.Get(fmt.Sprintf("https://www.pixiv.net/ajax/novel/%s", ID))
	if err != nil {
		return nil, err
	}
	defer data.Body.Close()

	body, err := ioutil.ReadAll(data.Body)

    var novel Novel 
    if err := json.Unmarshal(body, &novel); err != nil {
        return nil, err 
    }

	return &novel, nil
}
