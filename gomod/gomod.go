package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"time"
)

type Mod struct {
	Path      string    `json:"Path"`
	Version   string    `json:"Version"`
	Time      time.Time `json:"Time"`
	Indirect  bool      `json:"Indirect"`
	Dir       string    `json:"Dir"`
	GoMod     string    `json:"GoMod"`
	GoVersion string    `json:"GoVersion"`
}

func main() {
	var file string
	var mods []Mod
	flag.StringVar(&file, "file", "", "input json")
	flag.Parse()
	content, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal(content, &mods); err != nil {
		panic(err)
	}
	output := "output"
	f, err := os.Create(output) //创建文件
	if err != nil {
		fmt.Println("create file fail")
	}
	defer f.Close()
	for _, mod := range mods {
		if mod.Version == "" {
			mod.Version = "latest"
		}
		buf := fmt.Sprintf("go get %s@%s\n", mod.Path, mod.Version)
		_, err := f.Write([]byte(buf)) //写入文件(字节数组)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("SEE FILE [output]!")
}
