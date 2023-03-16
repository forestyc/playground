// Command click is a chromedp example demonstrating how to use a selector to
// click on an element.
package main

import (
	"archive/zip"
	"context"
	"encoding/csv"
	"fmt"
	"github.com/chromedp/cdproto/browser"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/target"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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
		chromedp.WithErrorf(log.Printf),
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
		chromedp.WithErrorf(log.Printf),
	)
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()

	// progress
	done := DownloadListener(ctx)
	// get working directory
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	err = chromedp.Run(ctx,
		chromedp.Navigate("https://www.sgx.com/research-education/derivatives"),
		chromedp.WaitVisible(`body`, chromedp.ByQuery),
		chromedp.Sleep(30*time.Second), // 防止下拉列表值未初始化
		// 展开type of data下拉框
		chromedp.Click(`#page-container > template-base > div > div > section.col-xxs-12.col-md-9.template-widgets-section > div > sgx-widgets-wrapper > widget-research-and-reports-download:nth-child(6) > widget-reports-derivatives-settlement > div > sgx-input-select:nth-child(1) > label > span.sgx-input-select-filter-wrapper > input`,
			chromedp.ByQuery, chromedp.NodeReady),
		// 选择futures
		chromedp.Click(`#sgx-select-dialog > div.sgx-dialog-box.tether-element.tether-enabled.tether-element-attached-top.tether-element-attached-left.tether-target-attached-bottom.tether-target-attached-left > sgx-select-picker > sgx-list > div > div > sgx-select-picker-option:nth-child(1)`,
			chromedp.ByQuery),
		//chromedp.Click(`#sgx-select-dialog > div.sgx-dialog-box.tether-element.tether-enabled.tether-element-attached-top.tether-element-attached-left.tether-target-attached-bottom.tether-target-attached-left > sgx-select-picker > sgx-list > div > div > sgx-select-picker-option:nth-child(2)`,
		//	chromedp.ByQuery),	// options
		chromedp.ActionFunc(func(ctx context.Context) error {
			c := chromedp.FromContext(ctx)
			return browser.SetDownloadBehavior(browser.SetDownloadBehaviorBehaviorAllowAndName).
				WithDownloadPath(wd).
				WithEventsEnabled(true).Do(cdp.WithExecutor(ctx, c.Browser))
		}),
		chromedp.Click(`#page-container > template-base > div > div > section.col-xxs-12.col-md-9.template-widgets-section > div > sgx-widgets-wrapper > widget-research-and-reports-download:nth-child(6) > widget-reports-derivatives-settlement > div > button`,
			chromedp.ByQuery, chromedp.NodeVisible),
	)
	if err != nil {
		log.Printf("wrote %s")
	}
	guid := <-done
	if err = HandleJob(filepath.Join(wd, guid)); err != nil {
		log.Printf("wrote %s")
	}
}

func DownloadListener(ctx context.Context) <-chan string {
	done := make(chan string, 1)
	chromedp.ListenBrowser(ctx, func(v interface{}) {
		select {
		case <-ctx.Done():
			done <- ""
			close(done)
		default:
		}
		if ev, ok := v.(*browser.EventDownloadProgress); ok {
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
	return done
}

func Unzip(path string) (string, error) {
	archive, err := zip.OpenReader(path)
	if err != nil {
		panic(err)
	}
	dstFileName := ""
	for _, f := range archive.File {
		dstFileName = f.Name
		dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return dstFileName, err
		}

		fileInArchive, err := f.Open()
		if err != nil {
			return dstFileName, err
		}

		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			return dstFileName, err
		}

		dstFile.Close()
		fileInArchive.Close()
	}
	archive.Close()
	return dstFileName, nil
}

type DayQuotsSgx struct {
	SpeciesId      string
	Exchange       string
	QuotSource     string
	ContractId     string
	SettlePrice    string
	PreSettlePrice string
	HighPrice      string
	LowPrice       string
	Volumns        string
	Interest       string
	Date           string
}

type SgxFuture map[string]DayQuotsSgx

var (
	nameMap = map[string]string{
		"FEF":  "SGX铁矿石62%(汇总)",
		"M65F": "SGX铁矿石65%(汇总)",
		"LPF":  "SGX铁矿石块矿(汇总)",
	}
)

func ReadCsv(path string) (SgxFuture, error) {
	result := make(SgxFuture)

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1

	for {
		csvdata, err := reader.Read() // 按行读取数据,可控制读取部分
		if err == io.EOF {
			break
		}

		com := strings.TrimSpace(csvdata[1])
		if com == "FEF" || com == "M65F" || com == "LPF" {
			year := csvdata[3][2:]
			month, _ := strconv.Atoi(csvdata[2])
			quot := DayQuotsSgx{
				SpeciesId:      nameMap[com],
				Exchange:       "SGX",
				QuotSource:     "SGX",
				ContractId:     year + fmt.Sprintf("%02d", month),
				SettlePrice:    csvdata[8],
				PreSettlePrice: "",
				HighPrice:      csvdata[5],
				LowPrice:       csvdata[6],
				Volumns:        csvdata[9],
				Interest:       csvdata[10],
				//Date:           "",
			}
			key := quot.SpeciesId + quot.ContractId
			result[key] = quot
		}
	}
	return result, nil
}

func HandleJob(path string) error {
	defer os.Remove(path)
	archive, err := Unzip(path)
	if err != nil {
		log.Println(err)
		return err
	}
	defer os.Remove(archive)
	result, err := ReadCsv(archive)
	if err != nil {
		return err
	}
	// todo
	for k, v := range result {
		log.Println("key", k, "value", v)
	}
	return nil
}
