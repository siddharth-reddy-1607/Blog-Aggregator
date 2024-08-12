package utils

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/siddharth-reddy-1607/Blog-Aggregator/internal/database"
)

type CustomTime struct{
    time.Time
}

func (ct *CustomTime) UnmarshalXML(d *xml.Decoder,start xml.StartElement) error{
    var s string
    if err := d.DecodeElement(&s,&start); err != nil{
        return err
    }
    t,err := time.Parse(time.RFC1123Z,s)
    if err != nil{
        return err
    }
    ct.Time = t
    return nil
}

type Posts struct{
    Title string `xml:"title"`
    Link string `xml:"link"`
    Description string `xml:"description"`
    PublishedDate CustomTime `xml:"pubDate"`
}
type RSSFeed struct{
    XMLName xml.Name `xml:"rss"`
    Name string `xml:"channel>title"`
    Link string `xml:"channel>link"`
    Description string `xml:"channel>description"`
    Posts []Posts `xml:"channel>item"`
}

func ProcessFeed(dbFeedID uuid.UUID,url string,wg *sync.WaitGroup,dbQueries *database.Queries){
    response,err := http.Get(url)
    defer wg.Done()
    if err != nil{
        log.Printf("Error while getting the RSS feed %v: %v\n",url,err)
        return
    }
    feed := RSSFeed{}
    data,err := io.ReadAll(response.Body)
    xml.Unmarshal(data,&feed)
    if err != nil{
        log.Printf("Error while reading the XML %v: %v\n",url,err)
        return 
    }
    if err := xml.Unmarshal(data,&feed); err != nil{
        log.Printf("Error while decoding XML %v: %v\n",url,err)
        return
    }
    for _,post := range(feed.Posts){
        curTime := time.Now()
        _,err := dbQueries.CreatePost(context.Background(),database.CreatePostParams{
                    ID: uuid.New(),
                    CreatedAt: curTime,
                    UpdatedAt: curTime,
                    Title: post.Title,
                    Url: post.Link,
                    Description: post.Title,
                    PublishedAt : post.PublishedDate.Time,
                    FeedID: dbFeedID,}) 
        if err != nil{
            if found := strings.Contains(strings.ToLower(err.Error()),"unique constraint"); !found{
                log.Printf("Error while posting feed to DB: %v\n",err)
            }
        }
    }
}

func ProcessFeeds(dbQueries *database.Queries,limit int32){
    for {
        fmt.Println("---Fetching latest posts---")
        dbFeeds,err := dbQueries.GetNextNFeedsToFetch(context.Background(),limit)
        if err != nil{
            log.Printf("Error while getting the latest feeds to fetch: %v\n",err)
            return
        } 
        log.Println("Starting processing all feeds...")
        var wg sync.WaitGroup
        for _,dbFeed := range dbFeeds{
            wg.Add(1)
            go ProcessFeed(dbFeed.ID,dbFeed.Url,&wg,dbQueries)
            dbFeed,err = dbQueries.MarkFeedFetched(context.Background(),dbFeed.ID)
            if err != nil{
                log.Printf("Error while getting the marking the feed as fetched: %v\n",err)
                return
            }
        }
        wg.Wait()
        log.Println("--------------Done processing all feeds--------------")
        time.Sleep(60*time.Second)
    }
}
