# Network Scanner for Open Ports and Banner Grabbing

This project is a Golang-based network scanner that scans a specified IP range to detect open ports and attempts to grab banners for SMTP and FTP services. The scanner is designed to leverage multi-threading (Goroutines) for enhanced performance on Linux systems.

---

## Table of Contents

- [Project Overview](#project-overview)
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation and Setup](#installation-and-setup)
- [Usage](#usage)
- [Code Structure and Architecture](#code-structure-and-architecture)
- [Performance Testing and Nmap Comparison](#performance-testing-and-nmap-comparison)
- [Future Enhancements](#future-enhancements)
- [License](#license)

---

## Project Overview

This Golang application performs the following tasks:
- Scans a given range of IP addresses to identify open ports.
- Attempts to grab banners from SMTP and FTP servers running on open ports.
- Utilizes multi-threading (Goroutines and Channels) to optimize performance.
- Provides results that can be compared against the popular Nmap tool for accuracy.

---

## Features

- **IP Range Scanning:** Generate and scan a range of IP addresses based on user-specified start and end points.
- **Banner Grabbing:** Retrieve service banners from SMTP (ports 25, 465, 587) and FTP (port 21) services.
- **Multi-threading:** Use of Goroutines and Channels enhances the scanning speed and performance.
- **Performance Optimization:** Designed for efficient scanning even with large IP ranges.
- **Accuracy Comparison:** Results can be compared with Nmap outputs to verify accuracy and completeness.

---

## Prerequisites

- **Operating System:** Linux (the focus is on Linux, though the code can be run on other platforms with minor modifications).
- **Programming Language:** Golang (version 1.13 or later is recommended).
- **Network Access:** Ensure you have network access to scan the target IP range.
- **Basic Networking Knowledge:** Familiarity with TCP/IP, SMTP, and FTP protocols.

---

## Installation and Setup

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/mojtabahadei81/NetSentinel.git
   cd NetSentinel

Build the Application:

2. **Compile the application using the following command:**

   ```bash
   go build -o network_scanner main.go
   Run the Application:

3. **Execute the compiled binary:**

   ```bash
   ./network_scanner

## Usage
Configuring the IP Range:

Modify the startIP and endIP variables in main.go with your desired IP range.

Specifying Ports:

Set the port list in the portsToBeScanned array (e.g., “80”, “25”, “465”, “587”, “21”, “53”).

Viewing Output:

The program will print the list of open ports along with the banners grabbed from SMTP and FTP services in the console output.

The application uses multiple workers running concurrently to rapidly scan and gather data from the target IP addresses.

## Code Structure and Architecture
Project Layout:

main.go: Contains the core logic for scanning IP ranges and grabbing banners.
Functions for IP conversion, port scanning, banner grabbing, and multi-threading are defined within this file.
Key Functions and Their Responsibilities:

generateIPRange(): Generates a list of IP addresses between the start and end IP.
checkPortsConcurrently(): Scans each IP for the specified list of ports concurrently.
getSMTPBanner() and getFTPBanner(): Connect to SMTP/FTP servers and retrieve the service banners.
Worker functions (worker(), bannerWorker()): Manage concurrent tasks via Goroutines and Channels.
Multi-threading Implementation:

Utilizes Go Channels to coordinate between concurrent tasks.
Implements timeouts to prevent blocking when a connection fails.
Performance Testing and Nmap Comparison
Performance Evaluation:

The tool can be tested on various IP ranges (from a few addresses to large blocks) to measure its performance in terms of speed and resource utilization.

## Nmap Comparison:

You can run Nmap on the same IP range and compare the open port results and banner data. Detailed testing reports can be prepared outlining any discrepancies or performance enhancements.

## Future Enhancements
Potential areas for additional features and improvements include:

UDP Port Scanning: Extend scanning to include UDP ports.
Dynamic Configuration: Introduce command-line options or configuration files for easier customization of IP ranges and port lists.
Improved Error Handling: Enhance error messages and logging for better debugging.
User Interface: Develop a GUI to allow non-technical users to interact with the scanner.

Conclusion
This network scanner provides a robust solution for identifying open ports and retrieving banners from SMTP and FTP services. The use of Golang’s concurrency features ensures high performance, making it an excellent tool for network security assessments and learning about network scanning and multi-threaded programming.
