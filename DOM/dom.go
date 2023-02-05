package DOM

import (
	"fmt"
	"regexp"
	"strings"
	"sync"
)

func DomSinker(str string, Url string) {
	var b bool
	regexList := []string{`.*\.innerHTML.*`, `eval\(.*\)`, `setInterval\(.*\)`, `setImmediate\(.*\)`, `execScript\(.*\)`, `Range\.createContextualFragment\(.*\)`, `window\.location.*?(\ )=.*`, `document\.domain.*?(\ )=.*`, `window\.location\.hash.*?(\ )=.*`, `window\.open\(.*\)`, `setTimeout\(.*\)`, `.*\.outerHTML.*?(\ )=.*?.*?(\+).*`, `.*\.insertAdjacentHTML.*?(\ )=.*?.*?(\+).*`, `.*\.onEventName.*?(\ )=.*?.*?(\+).*`, `crypto\.generateCRMFRequest\(.*\)`, `setImmediate\(.*\)`, `setInterval\(.*\)`}
	b = false
	fmt.Println(Url)
	var wg sync.WaitGroup
	for _, regex := range regexList {
		wg.Add(1)
		go func(regex string) {
			defer wg.Done()
			re := regexp.MustCompile(regex)
			matches := re.FindAllString(str, -1)

			if len(matches) > 0 {
				b = true
				fmt.Printf("Matches found for regex %q:\n", regex)
				fmt.Println(strings.Join(matches, "\n"))
			}
		}(regex)
	}
	wg.Wait()

	if b != true {
		fmt.Println("No result!")
	}

}
