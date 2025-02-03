# lwc (Line, Word, Count)

`lwc` is a simple command-line tool to count the number of lines, words, bytes, and characters in a text file. It's a Go implementation of `wc` command available on Unix-based systems .

**This project was created as part of the challenge on [Coding Challenges](https://codingchallenges.fyi/challenges/challenge-wc).**

## Features

- **Line Count**: Counts the number of lines in the file.
- **Word Count**: Counts the number of words in the file.
- **Byte Count**: Counts the total number of bytes in the file.
- **Character Count**: Counts the number of characters in the file.

## Usage

You can use `lwc` to process a file by running the following command:

```bash
go run cmd/lwc/main.go -file <your-file-name>
```

## Example

```bash
go run cmd/lwc/main.go -file test.txt
```

## Output:

```
Total ByteCount: <count>
Total Characters: <count>
Total Words: <count>
Total Lines: <count>
```

## Installation

If you'd like to run the program locally, clone the repository and run the following commands:

```bash
git clone https://github.com/suveshmoza/lwc.git

cd lwc

go run cmd/lwc/main.go -file <your-file-name>
```

## Testing

To run unit tests:

```bash
go test ./tests
```
