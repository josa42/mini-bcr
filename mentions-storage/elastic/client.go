package elastic

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch"
	"github.com/elastic/go-elasticsearch/esapi"
)

type Storable interface {
	ToJSON() string
}

type Client struct {
	client *elasticsearch.Client
}

func NewClient(host string) *Client {
	es, _ := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{fmt.Sprintf("http://%s", host)},
	})

	return &Client{
		client: es,
	}
}

func (c *Client) Info() *esapi.Response {
	r, _ := c.client.Info()
	return r
}

func (c *Client) Store(m Storable) {

	req := esapi.IndexRequest{
		Index:   "mentions",
		Body:    strings.NewReader(m.ToJSON()),
		Refresh: "true",
	}

	res, err := req.Do(context.Background(), c.client)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Printf("[%s] Error indexing document", res.Status())
	} else {
		// Deserialize the response into a map.
		var r map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			log.Printf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and indexed document version.
			log.Printf("[%s] %s; version=%d", res.Status(), r["result"], int(r["_version"].(float64)))
		}
	}
}
