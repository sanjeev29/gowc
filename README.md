# gowc
A command-line tool written in Go that replicates the functionality of the Unix `wc` (word count) utility. It provides an efficient and fast way to count lines, words, and characters in text files, leveraging the power of Go.

## Usage
1. **Clone the Repository**:
   ```bash
   git clone https://github.com/sanjeev29/gowc.git
   cd gowc
   ```

2. **Build the Executable**:
   Use the Go build tool to create the executable.
   ```bash
   go build -o gowc main.go
   ```

3. **Run the Tool**:
   Run the generated executable file with a text file as input.
   ```bash
   ./gowc test.txt
   ```
   This will output the number of lines, words, and characters in `test.txt`.
