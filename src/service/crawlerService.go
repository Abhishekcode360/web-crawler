package service

import (
	"Crawler/utils"
	"context"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/chromedp/chromedp"
	"github.com/gin-gonic/gin"
)

type Result struct {
	Domain string
	Urls   []string
}

func CrawlDomains(c *gin.Context, alldomain []string) utils.ProductURLs {
	productURLs := make(utils.ProductURLs)

	var mu sync.Mutex
	domainCh := make(chan string, 100)
	resultCh := make(chan Result, 100)
	var wg sync.WaitGroup
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for domain := range domainCh {
				urls := crawlDomain(domain)
				resultCh <- Result{Domain: domain, Urls: urls}
			}
		}()
	}

	go func() {
		for _, domain := range alldomain {
			domainCh <- domain
		}
		close(domainCh)
	}()

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	for result := range resultCh {
		mu.Lock()
		productURLs[result.Domain] = result.Urls
		mu.Unlock()
	}
	return productURLs
}

func crawlDomain(domain string) []string {
	var htmlContent string
	var ExtractedproductUrls []string
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	err := chromedp.Run(ctx,
		chromedp.Navigate(fmt.Sprintf("https://www.%s", domain)),
		chromedp.WaitVisible("body", chromedp.ByQuery),
		chromedp.OuterHTML("html", &htmlContent),
	)
	if err != nil {
		log.Printf("Failed to fetch domain %s: %v", domain, err)
	} else {
		urls := ExtractURLs(strings.NewReader(htmlContent), domain)
		ExtractedproductUrls = FilterProductURLs(urls)
	}
	return Unique(ExtractedproductUrls)
}
