package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	l := len(os.Args)
	if l != 2 {
		//fmt.Println(l)
		fmt.Println("\nUsage:\n\tlog-analysis.exe file.log\n\nauthor:kevin\nkevinfwp@126.com\n")
		return
	}
	file := os.Args[1]
	//file := "log.txt"
	ip_rgx := `(?i:ip)=(\d+).(\d+).(\d+).(\d+)`
	port_rgx := `(?i:port)=(\d+)`
	r_ip := regexp.MustCompile(ip_rgx)
	r_port := regexp.MustCompile(port_rgx)
	//var nums int64
	// rgx := os.Args[2]
	// sep := os.Args[3]
	// num, _ := strconv.Atoi(os.Args[4])
	res := make(map[string]int64)
	f, e := os.Open(file)
	defer f.Close()
	if e != nil {
		fmt.Println("Error: ", e)
		return
	}
	buf := bufio.NewReader(f)

	for {
	next:
		b, e := buf.ReadBytes('\n')
		s := string(b)
		//fmt.Println(s)

		//		r := regexp.MustCompile("f([a-z]+)g")
		if r_ip.MatchString(s) {
			ips := r_ip.FindAllString(s, -1)
			if len(ips) != 2 {
				goto next
			}
			ports := r_port.FindAllString(s, -1)
			if len(ports) != 2 {
				goto next
			}
			ip1 := ips[0][3:]
			ip2 := ips[1][3:]
			//port1 := ports[0][5:]
			port2 := ports[1][5:]
			str := ip1 + "\t" + ip2 + "\t" + port2
			//fmt.Println(str)
			//r.fi

			//fmt.Print(s)
			//fmt.Println("a")
			// 		w := strings.Split(s, sep)
			// 		wn := w[num]
			if res[str] >= 1 {
				res[str] += 1
			} else {
				res[str] = 1
			}
		}
		if e != nil {
			if e == io.EOF {
				break
			}
			panic(e)
		}
		// 	//time.Sleep(time.Second * 1)
		// }
	}

	for k, v := range res {
		fmt.Println(k, "\t=>\t", v)
	}
}
