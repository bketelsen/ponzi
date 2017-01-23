package ponzi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/patrickmn/go-cache"
)

// Cached provides an in-memory read- and write-through cache for ponzu
type Cache struct {
	cache   *cache.Cache
	baseURL string
	client  *http.Client
}

// New returns a new Cache with `ttl` duration that purges at `purge` intervals.
//    New("https://mycms.com", 5*time.Minute, 30*time.Second)
func New(baseURL string, ttl time.Duration, purge time.Duration) *Cache {
	c := &http.Client{}
	c.Timeout = 1 * time.Second
	return &Cache{
		baseURL: baseURL,
		cache:   cache.New(ttl, purge),
		client:  c,
	}
}

// GetBySlug returns a single item of content
func (c *Cache) GetBySlug(slug string, contentType string, result interface{}) error {
	cached, ok := c.cache.Get(slug)
	if ok {
		err := json.Unmarshal(cached.([]byte), result)
		return err
	}
	var body []byte
	url := fmt.Sprintf("%s/api/content?slug=%s", c.baseURL, slug)
	response, err := c.client.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	body, err = ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return err
	}
	c.cache.SetDefault(slug, body)
	return err
}

// Get returns a single item of content
func (c *Cache) Get(id int, contentType string, result interface{}) error {
	cached, ok := c.cache.Get(strconv.Itoa(id))
	if ok {
		err := json.Unmarshal(cached.([]byte), result)
		return err
	}
	var body []byte
	url := fmt.Sprintf("%s/api/content?type=%s&id=%d", c.baseURL, contentType, id)
	response, err := c.client.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	body, err = ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return err
	}
	c.cache.SetDefault(strconv.Itoa(id), body)
	return err
}

// GetAll returns all items
func (c *Cache) GetAll(contentType string, result interface{}) error {
	cached, ok := c.cache.Get(contentType)
	if ok {
		err := json.Unmarshal(cached.([]byte), result)
		return err
	}
	var body []byte
	url := fmt.Sprintf("%s/api/contents?type=%s&count=-1", c.baseURL, contentType)
	response, err := c.client.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	body, err = ioutil.ReadAll(response.Body)

	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return err
	}
	c.cache.SetDefault(contentType, body)
	return err
}
