// Command click is a chromedp example demonstrating how to use a selector to
// click on an element.
package main

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/browser"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/target"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	// get list
	//GetList()
	// click
	// Click()
	// Download
	Download()
}

func GetList() {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		//chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	var nodes []*cdp.Node
	err := chromedp.Run(ctx,
		//chromedp.Navigate(`https://www.cnblogs.com/`),
		//chromedp.WaitVisible(`body`),
		//chromedp.Nodes(`.//a[@class="post-item-title"]`, &nodes, chromedp.ByQuery),
		chromedp.Navigate("https://www.cnblogs.com/"),
		chromedp.WaitVisible(`body`, chromedp.ByQuery),
		chromedp.Nodes(`.//a[@class="post-item-title"]`, &nodes),
	)
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range nodes {
		log.Println(e.Children[0].NodeValue, ":", e.AttributeValue("href"))
	}
}

func Click() {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		//chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	ch := addNewTabListener(ctx)
	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://www.cnblogs.com/"),
		chromedp.WaitVisible(`body`, chromedp.ByQuery),
		chromedp.Click(`//*[@id="post_list"]/article[1]/section/div/a`),
	)
	if err != nil {
		log.Fatal(err)
	}

	newCtx, cancel2 := chromedp.NewContext(ctx, chromedp.WithTargetID(<-ch))
	defer cancel2()

	var res string
	err = chromedp.Run(newCtx,
		chromedp.OuterHTML(`html`, &res, chromedp.ByQuery),
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res)
}

/**
 * 注册新tab标签的监听服务
 */
func addNewTabListener(ctx context.Context) <-chan target.ID {
	mux := http.NewServeMux()
	ts := httptest.NewServer(mux)
	defer ts.Close()

	return chromedp.WaitNewTarget(ctx, func(info *target.Info) bool {
		return true
	})
}

func Download() {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	//	ch := addNewTabListener(ctx)
	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()

	// set up a channel, so we can block later while we monitor the download
	// progress
	done := make(chan string, 1)
	// set up a listener to watch the download events and close the channel
	// when complete this could be expanded to handle multiple downloads
	// through creating a guid map, monitor download urls via
	// EventDownloadWillBegin, etc
	chromedp.ListenTarget(ctx, func(v interface{}) {
		select {
		case <-ctx.Done():
			done <- ""
			close(done)
		}
		if ev, ok := v.(*browser.EventDownloadWillBegin); ok {
			ev.SuggestedFilename = "asdfa"
		} else if ev, ok := v.(*browser.EventDownloadProgress); ok {
			completed := "(unknown)"
			if ev.TotalBytes != 0 {
				completed = fmt.Sprintf("%0.2f%%", ev.ReceivedBytes/ev.TotalBytes*100.0)
			}
			log.Printf("state: %s, completed: %s\n", ev.State.String(), completed)
			if ev.State == browser.DownloadProgressStateCompleted {
				done <- ev.GUID
				close(done)
			}
		}
	})

	// get working directory
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	err = chromedp.Run(ctx,
		chromedp.Navigate("https://www.sgx.com/research-education/derivatives"),
		chromedp.WaitVisible(`#page-container > template-base > div > div > section.col-xxs-12.col-md-9.template-widgets-section > div > sgx-widgets-wrapper > widget-research-and-reports-download:nth-child(6) > widget-reports-derivatives-settlement > div > button`,
			chromedp.ByQuery),
		chromedp.ActionFunc(func(ctx context.Context) error {
			c := chromedp.FromContext(ctx)
			return browser.SetDownloadBehavior(browser.SetDownloadBehaviorBehaviorAllowAndName).
				WithDownloadPath(wd).
				WithEventsEnabled(true).Do(cdp.WithExecutor(ctx, c.Browser))
		}),
		chromedp.Click(`#page-container > template-base > div > div > section.col-xxs-12.col-md-9.template-widgets-section > div > sgx-widgets-wrapper > widget-research-and-reports-download:nth-child(6) > widget-reports-derivatives-settlement > div > button`,
			chromedp.ByQuery),
	)
	if err != nil {
		log.Fatal(err)
	}
	// This will block until the chromedp listener closes the channel
	guid := <-done

	// We can predict the exact file location and name here because of how we
	// configured SetDownloadBehavior and WithDownloadPath
	log.Printf("wrote %s", filepath.Join(wd, guid))
}
