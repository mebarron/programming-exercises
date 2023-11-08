## Ok, you want to use Go to parse IPv4 and IPv6 addresses from log data?

### Why use regex? 

- By using a regular expression, you can parse out the IPv4 and IPv6 addresses even if the format of the
log file changes over time

### What does this program do? 

- Parse the contents of some log data, assuming it can be processed line by line and extract the IP addresses, 
this is helpful if your log data is in CSV format, or some other delimited text file
- Sort, dedup, and rank IP addresses by traffic, e.g., IP address and count
- Account for any ties, e.g., two IP addresses could have the exact same amount of traffic

### Example output:

Unique IPs:  [10.1.1.1 10.1.1.2 10.1.1.3 172.16.21.40 192.168.0.2] 

Count | IPAddress
7 	 10.1.1.3
5 	 10.1.1.1
2 	 192.168.0.2
1 	 10.1.1.2
These IP addresses have the exact same count value:
1	 172.16.21.40


