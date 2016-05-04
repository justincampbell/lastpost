package main

import (
	"fmt"
	"io/ioutil"
	"testing"

	rss "github.com/jteeuwen/go-pkg-rss"
	"github.com/stretchr/testify/require"
)

var items []*rss.Item
var fixtures = []string{
	"aws-ec2-us-east-1",
	"crossfitwc",
	"github-status",
	"heroku-status",
	"weather.gov-KPHL",
}

func Test_formatContent(t *testing.T) {
	for _, fixture := range fixtures {
		feed := rss.New(1, true, nil, testItemHandler)
		items = []*rss.Item{}

		xml, err := ioutil.ReadFile(fmt.Sprintf("fixtures/%s.xml", fixture))
		require.NoError(t, err)

		b, err := ioutil.ReadFile(fmt.Sprintf("fixtures/%s.out", fixture))
		require.NoError(t, err)
		expected := string(b)

		err = feed.FetchBytes("http://example.com", xml, charsetReader)
		require.NoError(t, err)
		require.NotEmpty(t, items)

		item := items[0]
		c, err := extractContent(item)
		require.NoError(t, err)
		require.Equal(t, expected, formatContent(c))
	}
}

func testItemHandler(feed *rss.Feed, ch *rss.Channel, newitems []*rss.Item) {
	items = newitems
}
