package runner

import (
	"flag"
	"fmt"
	"github.com/canc3s/gnc/internal/gologger"
	"os"
)

const banner = `
   ▄██████▄       ███▄▄▄▄         ▄████████ 
  ███    ███      ███▀▀▀██▄      ███    ███ 
  ███    █▀       ███   ███      ███    █▀  
 ▄███             ███   ███      ███        
▀▀███ ████▄       ███   ███      ███        
  ███    ███      ███   ███      ███    █▄  
  ███    ███      ███   ███      ███    ███ 
  ████████▀        ▀█   █▀       ████████▀    v`



// Version is the current version of gnc
const Version = `0.0.1`

type Options struct {
	Listen 		bool
	//Udp    		bool
	Security	bool
	Version		bool
	Silent		bool
	Port		int
	Addr    	string
	Exec		string
	Proto		string
	Pass		string
}

func ParseOptions() *Options {
	options := &Options{}

	flag.IntVar(&options.Port, "p", 0, "监听的端口号")
	flag.BoolVar(&options.Listen, "l", false, "纯净模式")
	//flag.BoolVar(&options.Udp, "u", false, "Silent mode")
	flag.BoolVar(&options.Security, "s", false, "开启加密模式")
	flag.BoolVar(&options.Silent, "silent", false, "纯净模式")
	flag.StringVar(&options.Exec, "e", "", "连接后执行的程序")
	flag.StringVar(&options.Pass, "pass", "", "指定加密模式的密码（客服端和服务端需保持一直，通信乱码请检查指定密码是否一致）")

	showBanner()

	flag.Parse()

	if options.Version {
		gologger.Infof("Current Version: %s\n", Version)
		os.Exit(0)
	}

	options.configureOutput()

	options.validateOptions()

	return options
}

func (options *Options) validateOptions() {

	//if options.Udp {
	//	options.Proto = "udp"
	//} else {
	//	options.Proto = "tcp"
	//}
	options.Proto = "tcp"


	if options.Port < 0 || options.Port > 65535  {
		gologger.Fatalf("非法端口，请检查输入")
	}

	options.Addr = fmt.Sprintf("%s:%s", flag.Arg(0), flag.Arg(1))

	//fmt.Println(options.Addr)

	//fmt.Println(flag.Arg(2))
	//if options.CompanyID != "" && len(options.CompanyID) != 10 {
	//	gologger.Fatalf("公司ID %s 不正确!（10位数字）\n", options.CompanyID)
	//}
	//if options.InputFile != "" && !fileutil.FileExists(options.InputFile) {
	//	gologger.Fatalf("文件 %s 不存在!\n", options.InputFile)
	//}
	//if options.CompanyID == "" && options.InputFile == "" {
	//	flag.PrintDefaults()
	//	os.Exit(0)
	//}
}

func showBanner() {
	gologger.Printf("%s %s\n", banner,Version)
	//gologger.Printf("\t\thttps://github.com/canc3s/xx\n\n")

	//gologger.Labelf("请谨慎使用,您应对自己的行为负责\n")
	//gologger.Labelf("开发人员不承担任何责任，也不对任何滥用或损坏负责.\n")
}

// configureOutput configures the output on the screen
func (options *Options) configureOutput() {
	if options.Silent {
		gologger.MaxLevel = gologger.Silent
	}
}