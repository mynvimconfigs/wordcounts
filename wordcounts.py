#!/usr/bin/env python

import sys
import re
from os.path import exists
from collections import Counter
import argparse

def wordcounts(filename):
    """
    Counts the total number of words, characters, and their frequencies in the given file.

    Args:
        filename (str): Path to the file.

    Returns:
        tuple: Total word count, total distinct word count, total characters with spaces,
               total characters without spaces, dictionary of word frequencies.
    """
    try:
        word_counter = Counter()
        total_words = 0
        total_chars_with_spaces = 0
        total_chars_without_spaces = 0

        with open(filename, 'r') as file:
            for line in file:
                total_chars_with_spaces += len(line)
                total_chars_without_spaces += len(line.replace(" ", "").replace("\t", "").replace("\n", ""))
                words = re.findall(r'\b\w+\b', line.lower())
                word_counter.update(words)
                total_words += len(words)
        distinct_words = len(word_counter)
        return total_words, distinct_words, total_chars_with_spaces, total_chars_without_spaces, dict(word_counter)
    except FileNotFoundError:
        print(f"Error: File '{filename}' not found.")
        return None, None, None, None, None
    except Exception as e:
        print(f"Error reading file: {e}")
        return None, None, None, None, None

def main():
    """
    Main function to handle the command-line interface and logic.
    """
    parser = argparse.ArgumentParser(description="Counts total words, characters, and word frequencies in a file.")
    parser.add_argument('filename', type=str, help='Path to the file to analyze.')
    parser.add_argument('-n', '--number', type=int, default=5, 
                        help='Number of top occurring words to display.')
    args = parser.parse_args()
    filename = args.filename
    top_n = args.number
    if not exists(filename):
        print(f"Error: File '{filename}' not found.")
        return
    total_words, distinct_words, chars_with_spaces, chars_without_spaces, word_counts = wordcounts(filename)
    if total_words is not None:
        print(f"Total word count (including repeated words): {total_words}")
        print(f"Total number of distinct words: {distinct_words}")
        print(f"Total number of characters (including spaces): {chars_with_spaces}")
        print(f"Total number of characters (excluding spaces): {chars_without_spaces}")
        if top_n:
            print(f"Top {top_n} word frequencies:")
            top_words = Counter(word_counts).most_common(top_n)
            for word, count in top_words:
                print(f"    {word}: {count}")
        else:
            print("Word frequencies (sorted alphabetically):")
            for word, count in sorted(word_counts.items()):
                print(f"    {word}: {count}")

if __name__ == '__main__':
    main()
