package main

import (
	"bufio"
	"os"
)

func main() {

	/*

			var n int
			fmt.Scan(&n)
			fmt.Println(n)



		in := bufio.NewReader(os.Stdin)
		out := bufio.NewWriter(os.Stdin)

		line, _ := in.ReadString('\n')
		line = strings.TrimSpace(line)
		a, _ := strconv.Atoi(line)

		fmt.Fprintln(out, a)

		out.Flush()


		file, _ := os.Create("output.txt")
		fmt.Fprintln(file, "This is how to ")
		file.Close()

	*/

	file2, _ := os.Open("output.txt")
	reader := bufio.NewReader(file2)
	// line, _ := reader.ReadString("\n")

}
