## Parse IPv4 and IPv6 addresses from log data using regular expressions

### Why use regexp? 

- By using a regular expression, you can parse out the IPv4 and IPv6 addresses even if the format of the
log file changes over time

### What does this program do? 

- Parse the contents of log data line by line and extract the IP addresses, this is helpful if your log data 
is in CSV format, or some other delimited text file format
- Dedup, sort, and rank IP addresses by traffic, e.g., IP address and traffic count
- Account for ties, e.g., two IP addresses could have the exact same amount of traffic

### Example output:

Unique IPs:  [10.1.1.1 10.1.1.2 10.1.1.3 172.16.21.40 192.168.0.2] 

Count | IPAddress

7 	 10.1.1.3

5 	 10.1.1.1

2 	 192.168.0.2

1 	 10.1.1.2

These IP addresses have the exact same count value:

1	 172.16.21.40

### Performance Improvement: 

The first loop splits the log file by line, and the inner loop splits each line by the delimeter. If you know
that your IP address is always going to be in a certain position (but not the exact position), e.g., it will always be 
the last entry, you can improve the performance. 

This test was run using a 100MB log file:

Evaluating each text element (brute force):

Processed  2699299  log entries in  8.266018625s

Evaluating each text element starting at the last element:

Processed  2699299  log entries in  3.136688s

