package jobs

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/filedrive-team/filplus-info/models"
	logger "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"
)

type Issue struct {
	Labels []*struct {
		Name string
	}
	Body string
}

type Crawler struct {
}

func NewCrawler() *Crawler {
	return &Crawler{}
}

func (c *Crawler) Run(ctx context.Context) {
	c.startTask()
	ticker := time.NewTicker(1 * time.Hour)
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			c.startTask()
		}
	}
}

func (c *Crawler) startTask() {
	reqUrlFormat := "https://api.github.com/repos/filecoin-project/filecoin-plus-client-onboarding/issues?state=all&page=%d&per_page=100"
	for page := 1; ; page += 1 {
		reqUrl := fmt.Sprintf(reqUrlFormat, page)
		if err := c.query(reqUrl); err != nil {
			break
		}
	}
}

func (c *Crawler) query(url string) error {
	cli := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.19 Safari/537.36")

	resp, err := cli.Do(req)
	if err != nil {
		logger.Errorf("request github api failed, %v", err)
		return err
	}

	body := resp.Body
	defer body.Close()

	payload, err := ioutil.ReadAll(body)
	if err != nil {
		logger.Errorf("request github api failed, %v", err)
		return err
	}

	var data []*Issue
	if err = json.Unmarshal(payload, &data); err != nil {
		logger.WithField("payload", string(payload)).Warning("json unmarshal failed")
		logger.Errorf("request github api failed, %v", err)
		return err
	}

	if len(data) == 0 {
		logger.Info("finished")
		return errors.New("finished")
	}
	for _, issue := range data {
		if len(issue.Body) <= 0 {
			continue
		}
		skip := true
		for _, label := range issue.Labels {
			if label.Name == "state:Granted" {
				skip = false
				break
			}
		}
		if skip {
			continue
		}
		reg := regexp.MustCompile(`(?si)(Name:.*)Notary Requested`)
		params := reg.FindAllStringSubmatch(issue.Body, -1)
		if params == nil {
			reg = regexp.MustCompile(`(?si)(Name:.*Addresses to be Notarized:.*)\r\n`)
			params = reg.FindAllStringSubmatch(issue.Body, -1)
			if params == nil {
				logger.WithField("body", issue.Body).Warn("submatch failed")
				continue
			}
		}
		dst := params[0][1]
		list := strings.Split(dst, "\n")
		c := &models.Client{}
		for _, item := range list {
			const Name = "Name:"
			const Website = "Media:"
			const Region = "Region:"
			const Addresses = "Notarized:"
			if find := strings.Index(item, Name); find >= 0 {
				c.Name = strings.Trim(item[find+len(Name):], " \r\n")
			} else if find := strings.Index(item, Website); find >= 0 {
				c.Website = strings.Trim(item[find+len(Website):], " \r\n")
			} else if find := strings.Index(item, Region); find >= 0 {
				c.Region = strings.Trim(item[find+len(Region):], " \r\n")
			} else if find := strings.Index(item, Addresses); find >= 0 {
				c.Address = strings.Trim(item[find+len(Addresses):], " \r\n")
			} else {
				strs := strings.Split(item, ": ")
				if len(strs) == 2 {
					if strings.Index(strs[0], "Name") >= 0 {
						c.Name = strings.Trim(strs[1], " \r\n")
					} else if strings.Index(strs[0], "Website") >= 0 {
						c.Website = strings.Trim(strs[1], " \r\n")
					} else if strings.Index(strs[0], "Region") >= 0 {
						c.Region = strings.Trim(strs[1], " \r\n")
					} else if strings.Index(strs[0], "Addresses") >= 0 {
						c.Address = strings.Trim(strs[1], " \r\n")
					}
				} else if len(item) > 2 {
					logger.WithField("strs", strs).Info("")
				}
			}
		}

		if c.Name != "" && c.Address != "" && c.Name != "JP Test (Delete)" && strings.ToLower(c.Address[0:1]) == "f" {
			err := models.UpsertClient(c)
			if err != nil {
				logger.Errorf("UpsertClient failed, %v", err)
			}
		}
	}
	return nil
}
