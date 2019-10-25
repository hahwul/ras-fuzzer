package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net"
	"os/exec"
	"strings"
	"time"
)

func banner() {
	fmt.Println("____ ___  ___       ____ _    ___  ___  ____ ____")
	fmt.Println("| . \\|  \\ | _\\  ___ |  _\\|| \\ |_ \\ |_ \\ | __\\| . \\")
	fmt.Println("|  <_| . \\[__ \\|___\\| _\\ ||_|\\| __]| __]|  ]_|  <_")
	fmt.Println("|/\\_/|/\\_/|___/     |/   |___/|___/|___/|___/|/\\_/      by hahwul")
}

func dQuery(sd string, target string) bool {
	ips, err := net.LookupIP(sd + "." + target)
	_ = ips
	if err != nil {
		return false
	} else {
		//fmt.Println(sd+"."+target)
		return true
	}
}

func main() {
	//var dict []string

	target := flag.String("target", "", "Target domain (e.g hahwul.com)")
	length := flag.Int("length", 8, "Max Length of domain")
	callback := flag.String("callback", "", "runnning command if me find\n- Pattern: **PARAM**\n- e.g: **PARAM**")
	verbose := flag.Int("verbose", 1, "(Not Supported) Show/Unshow Log(1=show log, 0=only result)")

	flag.Parse()
	banner()
	if flag.NFlag() == 0 {
		flag.Usage()
		return
	}

	dict := [36]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p",
		"q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	fmt.Println("* Fuzzing Information")
	fmt.Printf("* Your Target: *.%s\n", *target)
	fmt.Printf("")
	fmt.Printf("* Max Length : %d\n", *length)
	fmt.Printf("* Dictionary  : %s\n", dict)
	fmt.Printf("* Combinatorial: (36+%d-1)!/((36-1)!*%d!) = ༼ つ ◕_◕ ༽つ <= !@#.. many case..\n", *length, *length)
	fmt.Println("-----------------------------------------------------------------------------------------")
	duplicated := []string{"09785645789"}

	for {
		//fmt.Println(dict[r1.Intn(len(dict))])
		sd := ""
		for i := 0; i < r1.Intn(*length)+1; i++ {
			sd += dict[r1.Intn(len(dict))]
		}
		if dQuery(sd, *target) {
			isDu := false
			for k, v := range duplicated {
				_ = k
				if sd == v {
					isDu = true
					break
				} else {

				}
			}
			if isDu {
				//fmt.Println(duplicated)
			} else {
				duplicated = append(duplicated, sd)
				fmt.Printf("\r[+] %s                                                         \n", sd+"."+*target)
				if len(*callback) > 0 {
					cmd := strings.Replace(*callback, "**PARAM**", sd+"."+*target, 1)
					out, err := exec.Command("sh", "-c", cmd).Output()
					_ = out
					_ = err
				}
			}
		}
		fmt.Printf("\r@ Testing to %s", sd)
	}
	_ = *verbose
}
