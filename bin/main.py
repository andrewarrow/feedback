#!/usr/bin/env python3

import sys
import os
import fileinput

path = sys.argv[1]
name = sys.argv[2]

def replace_placeholders(directory):
  for dirpath, _, filenames in os.walk(directory):
    for filename in filenames:
      file_path = os.path.join(dirpath, filename)
      if os.path.isfile(file_path):
        replace_in_file(file_path)

def replace_in_file(file_path):
  with fileinput.FileInput(file_path, inplace=True) as f:
    for line in f:
      line = line.replace('{{homeducky}}', name)
        print(line, end='')

def main():
  replace_placeholders("homeducky")

if __name__ == "__main__":
  main()
