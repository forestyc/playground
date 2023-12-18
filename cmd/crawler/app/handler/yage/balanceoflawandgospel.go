package yage

import (
	"encoding/json"
	"github.com/forestyc/playground/cmd/crawler/app/context"
	"github.com/forestyc/playground/pkg/crawler"
	"github.com/forestyc/playground/pkg/log/zap"
	"github.com/gocolly/colly/v2"
	rawZap "go.uber.org/zap"
	"os"
	"os/exec"
	"strings"
)

var urls = []string{
	//"http://api.yageapp.com/api/web/share/postor.php?aid=19143&sid=702424&bundleid=&base_uid=1532937",
	//"http://api.yageapp.com/api/web/share/postor.php?aid=4045&sid=123771&bundleid=&base_uid=-1",
	//"http://api.yageapp.com/api/web/share/postor.php?aid=2521&sid=71364&bundleid=&base_uid=-1",
	//"http://api.yageapp.com/api/web/share/postor.php?aid=4117&sid=126811&bundleid=&base_uid=-1",
	"http://api.yageapp.com/api/web/share/postor.php?aid=4119&sid=127035&bundleid=&base_uid=-1",
	"http://api.yageapp.com/api/web/share/postor.php?aid=4158&sid=129217&bundleid=&base_uid=-1",
	"http://api.yageapp.com/api/web/share/postor.php?aid=4985&sid=149346&bundleid=&base_uid=-1",
	"http://api.yageapp.com/api/web/share/postor.php?aid=4986&sid=149804&bundleid=&base_uid=-1",
	"http://api.yageapp.com/api/web/share/postor.php?aid=4987&sid=149834&bundleid=&base_uid=-1",
	"http://api.yageapp.com/api/web/share/postor.php?aid=7123&sid=219488&bundleid=&base_uid=-1",
	"http://api.yageapp.com/api/web/share/postor.php?aid=7539&sid=233624&bundleid=&base_uid=-1",
	"http://api.yageapp.com/api/web/share/postor.php?aid=8187&sid=260838&bundleid=&base_uid=-1",
	"http://api.yageapp.com/api/web/share/postor.php?aid=10271&sid=338085&bundleid=&base_uid=-1",
	"http://api.yageapp.com/api/web/share/postor.php?aid=19143&sid=702424&bundleid=&base_uid=-1",
	"http://api.yageapp.com/api/web/share/postor.php?aid=22953&sid=837655&bundleid=&base_uid=-1",
	"http://api.yageapp.com/api/web/share/postor.php?aid=23103&sid=841286&bundleid=&base_uid=-1",
	"http://api.yageapp.com/api/web/share/postor.php?aid=28272&sid=1088165&bundleid=&base_uid=-1",
	"http://api.yageapp.com/api/web/share/postor.php?aid=31175&sid=1241633&bundleid=&base_uid=-1",
	"http://api.yageapp.com/api/web/share/postor.php?aid=33517&sid=1360687&bundleid=&base_uid=-1",
	"http://api.yageapp.com/api/web/share/postor.php?aid=34652&sid=1430759&bundleid=&base_uid=-1",
	"http://api.yageapp.com/api/web/share/postor.php?aid=22660&sid=823832&bundleid=&base_uid=-1",
	"http://api.yageapp.com/api/web/share/postor.php?aid=27625&sid=1065396&bundleid=&base_uid=-1",
	"http://api.yageapp.com/api/web/share/postor.php?aid=28277&sid=1088372&bundleid=&base_uid=-1",
	"http://api.yageapp.com/api/web/share/postor.php?aid=37675&sid=1616932&bundleid=&base_uid=-1",
}

type Element struct {
	Playurl string `json:"playurl"`
	Media   string `json:"-"`
	Cover   string `json:"-"`
}

type BalanceOfLawAndGospel struct {
	task      string
	ctx       context.Context
	audios    map[string]Element
	targets   map[string]string
	directory string
	crawlers  []*crawler.Colly
}

func (blg *BalanceOfLawAndGospel) Init(ctx context.Context, config zap.Config, task string) {
	blg.ctx = ctx
	blg.ctx.C.Log = config
	// 初始化日志
	blg.ctx.Logger = zap.NewZap(blg.ctx.C.Log)
	blg.task = task
	blg.audios = make(map[string]Element)
	blg.targets = make(map[string]string)
	for _, url := range urls {
		c := crawler.NewColly(
			blg.task,
			url,
		)
		blg.crawlers = append(blg.crawlers, c)
	}
}

func (blg *BalanceOfLawAndGospel) Run() {
	for _, c := range blg.crawlers {
		if err := c.Run(crawler.WithCrawlCallback(blg.getList(c)), crawler.WithPipeline(blg.download())); err != nil {
			blg.ctx.Logger.Error("Run fail", rawZap.Error(err), rawZap.String("task", blg.task))
		}
		blg.audios = make(map[string]Element)
		blg.targets = make(map[string]string)
	}
	blg.ctx.Wg.Done()
	blg.ctx.Logger.Info("Run success", rawZap.String("task", blg.task))
}

func (blg *BalanceOfLawAndGospel) getList(c *crawler.Colly) crawler.Callback {
	return func() {
		// 获取专辑名
		c.Crawler.OnHTML(`#js_scroller > div.count_box > div`, func(element *colly.HTMLElement) {
			directory := strings.ReplaceAll(strings.TrimSpace(element.Text), `\`, "")
			directory = strings.ReplaceAll(directory, `/`, "")
			directory = strings.ReplaceAll(directory, `:`, "")
			directory = strings.ReplaceAll(directory, `*`, "")
			directory = strings.ReplaceAll(directory, `?`, "")
			directory = strings.ReplaceAll(directory, `<`, "")
			directory = strings.ReplaceAll(directory, `>`, "")
			directory = strings.ReplaceAll(directory, `|`, "")
			blg.directory = directory
		})
		// 获取音视频文件地址
		c.Crawler.OnHTML(`script:last-of-type`, func(element *colly.HTMLElement) {
			script := element.Text
			s1 := strings.Split(script, `var patharray = `)
			if len(s1) == 2 {
				s2 := strings.Split(s1[1], `;`)
				if len(s2) > 0 {
					targetJson := strings.ReplaceAll(s2[0], `\/`, "/")
					blg.audios = make(map[string]Element)
					if err := json.Unmarshal([]byte(targetJson), &blg.audios); err != nil {
						blg.ctx.Logger.Error("json unmarshal fail", rawZap.Error(err), rawZap.String("task", blg.task))
						return
					}
				}
			}
		})
		// 获取音频名称
		c.Crawler.OnHTML(`#js_scroller > ul > li`, func(li *colly.HTMLElement) {
			li.ForEach(`div.qui_list__bd > div.qui_list__box > p > span`, func(i int, p *colly.HTMLElement) {
				name := strings.TrimSpace(p.Text)
				name = strings.ReplaceAll(name, `\`, "")
				name = strings.ReplaceAll(name, `/`, "")
				name = strings.ReplaceAll(name, `:`, "")
				name = strings.ReplaceAll(name, `*`, "")
				name = strings.ReplaceAll(name, `?`, "")
				name = strings.ReplaceAll(name, `<`, "")
				name = strings.ReplaceAll(name, `>`, "")
				name = strings.ReplaceAll(name, `|`, "")
				blg.targets[li.Attr(`data-songid`)] = name
			})
		})
	}
}

func (blg *BalanceOfLawAndGospel) download() crawler.Pipeline {
	return func() error {
		// 创建文件夹
		if err := os.Mkdir(blg.directory, os.ModeDir); err != nil {
			blg.ctx.Logger.Error("mkdir fail", rawZap.Error(err), rawZap.String("task", blg.task))
		}
		// 下载
		for k, v := range blg.targets {
			filename := v
			idx := strings.LastIndex(blg.audios[k].Playurl, ".")
			suffix := blg.audios[k].Playurl[idx:]
			filename = strings.Join([]string{blg.directory, "/", filename, suffix}, "")
			cmd := exec.Command("curl", "-o", filename, blg.audios[k].Playurl)
			if err := cmd.Run(); err != nil {
				blg.ctx.Logger.Error("download fail", rawZap.Error(err), rawZap.String("task", blg.task), rawZap.String("cmd", cmd.String()))
			}
			blg.ctx.Logger.Info("download", rawZap.String("task", blg.task), rawZap.String("cmd", cmd.String()))
		}
		return nil
	}
}
