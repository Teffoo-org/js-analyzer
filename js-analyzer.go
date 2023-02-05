package main

import (
	"fmt"
	"github.com/Teffoo-org/js-analyzer/DOM"
	"github.com/Teffoo-org/js-analyzer/function"
	"github.com/Teffoo-org/js-analyzer/regex"
	"github.com/projectdiscovery/goflags"
	"github.com/projectdiscovery/gologger"
	"log"
)

type options struct {
	url    string
	file   string
	silent bool
	sink   bool
}

func main() {
	opt := &options{}
	flagSet := goflags.NewFlagSet()
	flagSet.SetDescription(`
     ██╗███████╗       █████╗ ███╗   ██╗ █████╗ ██╗  ██╗   ██╗███████╗███████╗██████╗ 
     ██║██╔════╝      ██╔══██╗████╗  ██║██╔══██╗██║  ╚██╗ ██╔╝╚══███╔╝██╔════╝██╔══██╗
     ██║███████╗█████╗███████║██╔██╗ ██║███████║██║   ╚████╔╝   ███╔╝ █████╗  ██████╔╝
██   ██║╚════██║╚════╝██╔══██║██║╚██╗██║██╔══██║██║    ╚██╔╝   ███╔╝  ██╔══╝  ██╔══██╗
╚█████╔╝███████║      ██║  ██║██║ ╚████║██║  ██║███████╗██║   ███████╗███████╗██║  ██║
 ╚════╝ ╚══════╝      ╚═╝  ╚═╝╚═╝  ╚═══╝╚═╝  ╚═╝╚══════╝╚═╝   ╚══════╝╚══════╝╚═╝  ╚═╝
                        Created By Zero To Hero :)`)
	flagSet.StringVarP(&opt.url, "url", "u", "", "url for analyze")
	flagSet.StringVarP(&opt.file, "file", "f", "", "js file to analyze")
	flagSet.BoolVarP(&opt.silent, "silent", "s", false, "show silent output")
	flagSet.BoolVarP(&opt.sink, "sink", "S", false, "crawl dom sinks")
	if err := flagSet.Parse(); err != nil {
		log.Fatalf("Could not parse flags: %s\n", err)
	}

	if opt.silent == false {
		showBanner()
	}
	if opt.url != "" && opt.sink != false {
		resp, err := function.Get(opt.url)
		if err != nil {
			fmt.Println(err)
		}
		DOM.DomSinker(resp, opt.url)
		return
	}
	if opt.url != "" {
		resp, err := function.Get(opt.url)
		if err != nil {
			fmt.Println(err)
		}
		regex.Regex(resp)
		return
	}

	if opt.file != "" {
		file, err := function.File(opt.file)
		if err != nil {
			fmt.Println(err)
		}
		regex.Regex(file)
		return
	}
	return
}

func showBanner() {
	gologger.Print().Msgf(`
     ██╗███████╗       █████╗ ███╗   ██╗ █████╗ ██╗  ██╗   ██╗███████╗███████╗██████╗ 
     ██║██╔════╝      ██╔══██╗████╗  ██║██╔══██╗██║  ╚██╗ ██╔╝╚══███╔╝██╔════╝██╔══██╗
     ██║███████╗█████╗███████║██╔██╗ ██║███████║██║   ╚████╔╝   ███╔╝ █████╗  ██████╔╝
██   ██║╚════██║╚════╝██╔══██║██║╚██╗██║██╔══██║██║    ╚██╔╝   ███╔╝  ██╔══╝  ██╔══██╗
╚█████╔╝███████║      ██║  ██║██║ ╚████║██║  ██║███████╗██║   ███████╗███████╗██║  ██║
 ╚════╝ ╚══════╝      ╚═╝  ╚═╝╚═╝  ╚═══╝╚═╝  ╚═╝╚══════╝╚═╝   ╚══════╝╚══════╝╚═╝  ╚═╝
`)
	gologger.Print().Msgf("   Created By Zero To Hero :)\n\n")
}
