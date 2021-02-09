package live

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
	"encoding/json"
)

const baseurl string = "https://gateway.marvel.com"

type Memorydb_impl interface {
	CreateCharacter(id int, name string, description string)
	GetCharacter(id int) (string, bool)
}

type live struct {
	db Memorydb_impl
}
var instance *live

type response struct {
	Data data `json:data`
}

type data struct {
	Results []character `json:results`
}

type character struct {
	Id int				`json:id`
	Name string			`json:name`
	Description string	`json:description`
}

func GetInstance(db ...Memorydb_impl) *live {
	if instance == nil {
		instance = &live{db:db[0]}
	}
	return instance
}

func (l *live) GetCharacter(id string) string {
	raw := call("/v1/public/characters/" + id)
	resp := response{}
	json.Unmarshal([]byte(raw), &resp)
	l.db.CreateCharacter(
		resp.Data.Results[0].Id,
		resp.Data.Results[0].Name,
		resp.Data.Results[0].Description,
	)
	ret, _ := l.db.GetCharacter(resp.Data.Results[0].Id)
	return ret
}

func (l *live) GetCharacters(page string) {
	resp := call("/v1/public/characters")
	fmt.Println(resp)
}

// Call to live endpoint
func call(endpoint string) string {
	resp, _ := http.Get(baseurl + endpoint + auth())
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

// Returns auth query string
func auth() string {
	ts := strconv.FormatInt(time.Now().Unix(), 10)
	hash := fmt.Sprintf("%x", md5.Sum([]byte(ts + prkey + pkey)))
	return "?ts=" + ts + "&apikey=" + pkey + "&hash=" + hash
}
