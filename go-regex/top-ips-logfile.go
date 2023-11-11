/* Read log file entries for IPv4 and IPv6 addresses using regex, and rank by traffic */

package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

var elapsed time.Time
var proc_counter int

// Parse the IP addresses from the log file entries
func parse(log string) []string {
	elapsed = time.Now()

	ipv4_regex := regexp.MustCompile(`^(?:[0-9]{1,3}\.){3}[0-9]{1,3}$`)
	ipv6_regex := regexp.MustCompile(`^(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))$`)

	var ip_list []string

	for _, line := range strings.Split(log, "\n") {
		proc_counter++

		contents := strings.Split(line, ",")

		for i := len(contents) - 1; i == len(contents)-1; i++ {
			ipv4_match := ipv4_regex.FindStringSubmatch(contents[i])
			if ipv4_match != nil {
				ip_list = append(ip_list, ipv4_match[0])
			}

			ipv6_match := ipv6_regex.FindStringSubmatch(contents[i])
			if ipv6_match != nil {
				ip_list = append(ip_list, ipv6_match[0])
			}
		}
	}
	return ip_list
}

// Dedup the data and obtain the distinct counts for each IPv4/IPv6 address
func rank(ip_list []string) {
	var ip_addresses []string
	ranking := make(map[int]string)

	for _, ip_address := range ip_list {
		ip_addresses = append(ip_addresses, ip_address)
	}

	sort.Strings(ip_addresses)

	var copy_arr []string

	for i, v := range ip_addresses {
		if i == len(ip_addresses)-1 {
			if ip_addresses[i] != "" {
				copy_arr = append(copy_arr, v)
			}
		}
		if i+1 < len(ip_addresses) {
			if v != ip_addresses[i+1] {
				if ip_addresses[i] != "" {
					copy_arr = append(copy_arr, v)
				}
			}
		}
	}

	ip_addresses = copy_arr

	fmt.Println("Unique IPs: ", ip_addresses, "\n")

	counter := 0

	var tie []string

	for _, ip := range ip_addresses {
		for i := len(ip_list) - 1; i >= 0; i-- {
			if ip == ip_list[i] {
				ip_addresses = ip_addresses[:len(ip_addresses)]
				counter++
			}
		}
		if _, is_tied := ranking[counter]; is_tied {
			entry := strconv.Itoa(counter) + "\t" + " " + ip
			tie = append(tie, entry)
			counter = 0
		} else {
			ranking[counter] = ip
			counter = 0
		}
	}

	var top_n []int

	for k, _ := range ranking {
		top_n = append(top_n, k)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(top_n)))

	fmt.Println("Count | IPAddress")

	for _, n := range top_n {
		fmt.Println(n, "\t", ranking[n])
	}

	fmt.Println("These IP addresses have the exact same count value:")
	for _, v := range tie {
		fmt.Println(v)
	}

	fmt.Println("Processed ", proc_counter, " log entries in ", time.Now().Sub(elapsed).String())
}

func main() {
	// Dummy data, e.g., what you might expect in a web server log file
	data, err := ioutil.ReadFile("logfile.txt")
	if err != nil {
		panic(err)
	}

	ip_addresses := parse(string(data))

	rank(ip_addresses)
}
