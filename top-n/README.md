## Parse and rank top n entries in a file by index 

### What does this program do?

This builds on the previous example (go-regex) to attempt a significant speed up in the performance, except we assume to *know exactly* where to look for the value we want to count in the file, (e.g., a log file, a file containing scores for a game, or some other stream of data we could read line by line and split into an array.)

### Example output: 

10   10.1.1.4   2014234
9   192.168.0.3   593270
8   192.168.0.1   68009
7   192.168.0.2   10129
6   10.1.1.1   7235
5   10.1.1.2   2894
4   10.1.1.5   2080
3   10.1.1.3   1447
2   10.1.1.44   3
1   10.1.1.45   2
Processed  2699304  lines in  349.754583ms

### Performance improvement: 

The function call to parse() passes three arguments, 1) the data as a string, 2) the delimeter (e.g., comma for CSV), and 3) the index to look for in the resulting array for each line.

The counting is condensed into the same block of code that parses out the values, which is why this approach is so much faster.

The ranking function accepts the resulting map as input, sorts, then prints the results to the console. This is ok if you're confident the data will always be formatted the same, and you know exactly what/where to look in the data/stream.