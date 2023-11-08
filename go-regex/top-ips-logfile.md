## Ok, you want to use Go to parse IPv4 and IPv6 addresses from log data?

### Why use regex? 

- By using a regular expression, you can parse out the IPv4 and IPv6 addresses even if the format of the
log file changes over time

### What does this program do? 

- Parse the contents of some log data, assuming it can be processed line by line and extract the IP addresses, 
this is helpful if your log data is in CSV format, or some other delimited text file
- Sort, dedup, and rank IP addresses by traffic, e.g., IP address and count
- Account for any ties, e.g., two IP addresses could have the exact same amount of traffic

