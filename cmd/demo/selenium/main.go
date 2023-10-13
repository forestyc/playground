package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/tebeka/selenium"
)

const (
	//指定对应浏览器的driver
	chromeDriverPath = "D:/tools/chromedriver/chromedriver.exe"
	port             = 8080
)

func main() {
	//ServiceOption 配置服务实例
	opts := []selenium.ServiceOption{
		selenium.Output(os.Stderr), // Output debug information to STDERR.
	}
	//SetDebug 设置调试模式
	selenium.SetDebug(true)
	//在后台启动一个ChromeDriver实例
	service, err := selenium.NewChromeDriverService(chromeDriverPath, port, opts...)
	if err != nil {
		panic(err) // panic is used only as an example and is not otherwise recommended.
	}
	defer service.Stop()

	//连接到本地运行的 WebDriver 实例
	//这里的map键值只能为browserName，源码里需要获取这个键的值，来确定连接的是哪个浏览器
	caps := selenium.Capabilities{"browserName": "chrome"}
	//NewRemote 创建新的远程客户端，这也将启动一个新会话。 urlPrefix 是 Selenium 服务器的 URL，必须以协议 (http, https, ...) 为前缀。为urlPrefix提供空字符串会导致使用 DefaultURLPrefix,默认访问4444端口，所以最好自定义，避免端口已经被抢占。后面的路由还是照旧DefaultURLPrefix写
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		panic(err)
	}
	defer wd.Quit()

	// 导航到指定网站页面，浏览器默认get方法
	if err = wd.Get("http://www.czce.com.cn/"); err != nil {
		panic(err)
	}
	// 此步操作为定位元素
	// 在当前页面的 DOM 中准确找到一个元素，使用css选择器，使用id选择器选择文本框
	// 值得注意的是，WebElement接口定义了几乎所有操作Web元素支持的方法，例如清除元素、移动鼠标、截图、鼠标点击、提交按钮等操作
	// 此处对于浏览器的自动化操作可以应用于爬虫登录需要输入密码或者其他应用
	time.Sleep(time.Hour)
	elem, err := wd.FindElement(selenium.ByID, "tab1")
	if err != nil {
		panic(err)
	}
	// 操作元素，删除文本框中已有的样板代码
	if err := elem.Clear(); err != nil {
		panic(err)
	}

	// 在文本框中输入一些新代码
	err = elem.SendKeys(`
        package main
        import "fmt"
        func main() {
            fmt.Println("Hello WebDriver!")
        }
    `)
	if err != nil {
		panic(err)
	}

	// css选择器定位按钮控件
	btn, err := wd.FindElement(selenium.ByCSSSelector, "#run")
	if err != nil {
		panic(err)
	}
	if err := btn.Click(); err != nil {
		panic(err)
	}

	// 等待程序完成运行并获得输出
	outputDiv, err := wd.FindElement(selenium.ByCSSSelector, "#output")
	if err != nil {
		panic(err)
	}

	var output string
	for {
		output, err = outputDiv.Text()
		if err != nil {
			panic(err)
		}
		if output != "Waiting for remote server..." {
			break
		}
		time.Sleep(time.Millisecond * 100)
	}

	fmt.Printf("%s", strings.Replace(output, "\n\n", "\n", -1))

	// Example Output:
	// Hello WebDriver!
	//
	// Program exited.
}
