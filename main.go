package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

var banners [][4]string

var openPorts [][2]string

func main() {
	startIP := "74.125.133.100"
	endIP := "74.125.133.150"

	portsToBeScanned := []string{"80", "25", "465", "587", "21", "53"}

	ipList, _ := generateIPRange(startIP, endIP)

	checkPortsConcurrently(ipList, portsToBeScanned)

	for _, port := range openPorts {
		fmt.Println(port)
	}

	time.Sleep(2 * time.Second)

	find_SMTP_And_FTP_Banners(openPorts)

	for _, banner := range banners {
		fmt.Println(banner)
	}
}

func getSMTPBanner(ip string, port string, results chan<- [4]string) {
	var banner [4]string
	banner[0] = ip
	banner[1] = port

	// conn, err := net.DialTimeout("tcp", "smtp.gmail.com"+":"+"25", 2*time.Second)
	conn, err := net.DialTimeout("tcp", ip+":"+port, 5*time.Second)
	if err != nil {
		fmt.Println("connection failed:", err)
		return
	}
	defer conn.Close()

	buffer := make([]byte, 1024)
	_, err = conn.Read(buffer)
	if err != nil {
		fmt.Println("error in getting banner:", err)
		return
	}
	banner[2] = string(buffer)

	_, err = conn.Write([]byte("EHLO " + ip + "\r\n"))
	if err != nil {
		fmt.Println("error in sending HELO request:", err)
		return
	}

	_, err = conn.Read(buffer)
	if err != nil {
		fmt.Println("error in getting EHLO response:", err)
		return
	}
	banner[3] = string(buffer)
	results <- banner
}

func getFTPBanner(ip string, port string, results chan<- [4]string) {
	var banner [4]string
	banner[0] = ip
	banner[1] = port

	conn, err := net.DialTimeout("tcp", ip+":"+port, 5*time.Second)
	if err != nil {
		fmt.Println("connection failed:", err)
		return
	}
	defer conn.Close()

	buffer := make([]byte, 1024)
	_, err = conn.Read(buffer)
	if err != nil {
		fmt.Println("error in getting banner:", err)
		return
	}
	banner[2] = string(buffer)

	_, err = conn.Write([]byte("USER anonymous\r\n"))
	if err != nil {
		fmt.Println("error in sending USER command:", err)
		return
	}

	_, err = conn.Read(buffer)
	if err != nil {
		fmt.Println("error in getting USER response:", err)
		return
	}
	banner[3] = string(buffer)
	results <- banner
}

func writeBanners(results chan [4]string) {
	for i := range results {
		banners = append(banners, i)
	}
}

func bannerWorker(ip_ports <-chan [2]string, results chan<- [4]string) {
	for ip_port := range ip_ports {
		ip := ip_port[0]
		port := ip_port[1]
		if port == "25" || port == "465" || port == "587" {
			//if port == "80" {
			getSMTPBanner(ip, port, results)
		} else if port == "21" {
			getFTPBanner(ip, port, results)
		}
	}

}

func find_SMTP_And_FTP_Banners(ip_ports [][2]string) {
	jobsNo := len(ip_ports)
	jobs := make(chan [2]string, jobsNo)
	results := make(chan [4]string, jobsNo)

	go writeBanners(results)

	for w := 0; w < 100; w++ {
		go bannerWorker(jobs, results)
	}

	for j := 0; j < jobsNo; j++ {
		jobs <- ip_ports[j]
	}

	time.Sleep(20 * time.Second)
	close(jobs)
	close(results)
}

func checkPortsConcurrently(ipRange []string, portsToBeScanned []string) {
	numIPs := len(ipRange) * len(portsToBeScanned)
	jobs := make(chan string, numIPs)
	results := make(chan [2]string, numIPs)
	var ip_ports []string

	for i := 0; i < len(ipRange); i++ {
		for j := 0; j < len(portsToBeScanned); j++ {
			ip_ports = append(ip_ports, ipRange[i]+":"+portsToBeScanned[j])
		}
	}

	go writeOpenPorts(results)

	for w := 0; w < 300; w++ {
		go worker(w, jobs, results)
	}

	for j := 0; j < numIPs; j++ {
		jobs <- ip_ports[j]
	}

	time.Sleep(6 * time.Second)
	close(jobs)
	close(results)
}

func worker(index int, jobs <-chan string, results chan<- [2]string) {
	for ip_port := range jobs {
		ip_port_splited := strings.Split(ip_port, ":")
		ip := ip_port_splited[0]
		port := ip_port_splited[1]
		checkPortIsOpen(ip, port, results, index)
	}
}

func checkPortIsOpen(ip, port string, results chan<- [2]string, index int) {
	var ip_port [2]string

	ip_port[0] = ip
	ip_port[1] = port

	conn, err := net.DialTimeout("tcp", ip+":"+port, 5*time.Second)
	if err != nil {
		fmt.Println(index, ": connection failed:", err)
		return
	} else {
		fmt.Println(index, ": connection to "+ip+":"+port+" successful!")
		results <- ip_port
	}
	defer conn.Close()
}

func writeOpenPorts(ip_ports chan [2]string) {
	for i := range ip_ports {
		openPorts = append(openPorts, i)
	}
}

func convertIpToDecimal(ipStr string) (ip uint32) {
	// var ip uint32

	result := strings.Split(ipStr, ".")

	for i := 3; i >= 0; i-- {
		number, err := strconv.Atoi(result[3-i])

		if err != nil {
			panic(err)
		}
		ip += uint32(number) << (i * 8)
	}
	return
}

func generateIPRange(startIP string, endIP string) ([]string, error) {
	startInt := convertIpToDecimal(startIP)
	endInt := convertIpToDecimal(endIP)

	var ipList []string
	for i := startInt; i <= endInt; i++ {
		ip := net.IPv4(byte(i>>24), byte(i>>16&0xFF), byte(i>>8&0xFF), byte(i&0xFF))
		ipList = append(ipList, ip.String())
	}

	return ipList, nil
}
