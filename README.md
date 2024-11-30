# wordcounts

A simple program for counting words (including a per word count) and characters in a text file.

There are two implementations (Python and Go).

The program analyzes a text file and provides various statistics about the content.

## Features

Both versions of the program share the following features:

- **Total Word Count:** Counts the total number of words (including repeated words).
- **Distinct Word Count:** Counts the total number of unique words in the file.
- **Character Count:** Provides:
  - Total number of characters (including spaces).
  - Total number of characters (excluding spaces).
- **Word Frequencies:** Displays the frequency of each word.
- **Top N Words:** Optionally, you can specify an integer `N` to display only the top `N` most frequently occurring words.

---

## Versions

The program is implemented in **Python** and **Go**, with identical functionality.

### Python Version

- **File:** `wordcounts.py`
- **Dependencies:** Python 3.x (no external libraries required).
- **Usage:**
  ```bash
  python wordcounts.py <filename> [-n N]
  ```
  - `<filename>`: Path to the text file to analyze.
  - `-n N`: (Optional) Display the top `N` most frequently occurring words.

- **Example:**
  ```bash
  python wordcounts.py sample.txt -n 3
  ```
  Output:
  ```
  Total word count (including repeated words): 7
  Total number of distinct words: 5
  Total number of characters (including spaces): 30
  Total number of characters (excluding spaces): 25
  Top 3 word frequencies:
      go: 2
      hello: 2
      world: 1
  ```

### Go Version

- **File:** `wordcount.go`
- **Dependencies:** Go 1.18 or later.
- **Usage:**
  ```bash
  go run wordcount.go <filename> [-n N]
  ```
  - `<filename>`: Path to the text file to analyze.
  - `-n N`: (Optional) Display the top `N` most frequently occurring words.

- **Example:**
  ```bash
  go run wordcount.go sample.txt -n 3
  ```
  Output:
  ```
  Total word count (including repeated words): 7
  Total number of distinct words: 5
  Total number of characters (including spaces): 30
  Total number of characters (excluding spaces): 25
  Top 3 word frequencies:
      go: 2
      hello: 2
      world: 1
  ```

---

## How to Run

### Python Version
1. Install Python 3.x if not already installed.
2. Save the `wordcounts.py` script to your local machine.
3. Run the program with the text file and optional `-n` argument.

### Go Version
1. Install Go 1.18 or later if not already installed.
2. Save the `wordcount.go` script to your local machine.
3. Run the program using:
   ```bash
   go run wordcount.go <filename> [-n N]
   ```

---

## Features in Detail

1. **Total Word Count:**
   - Includes all occurrences of words in the file.

2. **Distinct Word Count:**
   - Counts unique words (case-insensitive).

3. **Character Counts:**
   - **Including spaces:** Counts all characters, including spaces, tabs, and newlines.
   - **Excluding spaces:** Counts all characters except spaces, tabs, and newlines.

4. **Word Frequencies:**
   - Outputs each word and its frequency (case-insensitive).
   - Alphabetically sorted when no `-n` argument is provided.

5. **Top N Words:**
   - Use the `-n` option to display the top `N` most frequently occurring words.

---

## Contributing

Feel free to fork the repository and submit pull requests with improvements or bug fixes.

---

## License

This program is open-source and available under the MIT License.
