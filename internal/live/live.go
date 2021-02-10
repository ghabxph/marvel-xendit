package live

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/ghabxph/marvel-xendit/internal/config"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

const CONFIG_PATH_KEY string = "MARVEL_XENDIT_PATH"
const baseurl string = "https://gateway.marvel.com"

type http_iface interface {
	Get(url string) string
}

type memorydb_impl interface {
	CreateCharacter(id int, name string, description string)
	GetCharacter(id int) (string, bool)
}

type live struct {
	db     	memorydb_impl
	config *config.Config
	http	http_iface
}

var instance *live

type response struct {
	Data data `json:data`
}

type data struct {
	Total   int         `json:total`
	Results []character `json:results`
}

type character struct {
	Id          int    `json:id`
	Name        string `json:name`
	Description string `json:description`
}

type http_impl struct{}
func (h *http_impl) Get(url string) string {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

func GetInstance(db ...interface{}) *live {
	if instance == nil {
		instance = &live{db: db[0].(memorydb_impl), config: config.GetInstance(os.Getenv(CONFIG_PATH_KEY)), http:&http_impl{}}
	}
	return instance
}

func (l *live) GetCharacter(id string) string {
	raw := l.call("/v1/public/characters/" + id)
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

// Gets 100 characters. Returns the total number of characters
func (l *live) GetCharacters(offset string) int {
	raw := l.call("/v1/public/characters", "&limit=100&offset=" + offset)
	resp := response{}
	json.Unmarshal([]byte(raw), &resp)
	for _, v := range resp.Data.Results {
		l.db.CreateCharacter(v.Id, v.Name, v.Description)
	}
	return resp.Data.Total
}

// Mocks HttpGet (for test)
func (l *live) MockHttpGet(mock interface{}) {
	l.http = mock.(http_iface)
}

// Call to live endpoint
func (l *live) call(endpoint string, query ...string) string {
	q := ""
	if len(query) > 0 {
		q = query[0]
	}
	return l.http.Get(baseurl + endpoint + l.auth() + q)
}

// Returns auth query string
func (l *live) auth() string {
	ts := strconv.FormatInt(time.Now().Unix(), 10)
	hash := fmt.Sprintf("%x", md5.Sum([]byte(ts+l.config.PrivateKey()+l.config.PublicKey())))
	return "?ts=" + ts + "&apikey=" + l.config.PublicKey() + "&hash=" + hash
}
