#Math-skills#
Download the [file](https://assets.01-edu.org/stats-projects/math-skills) and grant the executable per,issions with this command:
```bash
chmod +x ./math-skills
```
Run the script with `./math-skills`, then run the my program with the created file (`data.txt`) by the previous command.
Run my program with the command:
```golang
go run . data.txt
```
Compare the output.

`Instruction`

Your program must be able to read from a file and print the result of each statistic mentioned above. In other words, your program must be able to read the data present in the path passed as argument. The data in the file will be presented as the following example:

189
113
121
114
145
110
...

This data represents a graph in which the values of the x axis are the number of the lines (0, 1, 2, 3, 4, 5 ...) and the values of the y axis are the actual numbers (189, 113, 121, 114, 145, 110...).

To run your program a command similar to this one will be used if your project is made in Go:

$> go run your-program.go data.txt

After reading the file, your program must execute each of the calculations asked above and print the results in the following manner (the following numbers are only examples):

Average: 35
Median: 4
Variance: 5
Standard Deviation: 65

Please note that the values are rounded integers.# math-skills
