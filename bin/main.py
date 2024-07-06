#!/usr/bin/env python3

import sys
import os
import fileinput
import shutil

path = sys.argv[1]
name = sys.argv[2]

def replace_placeholders(destination_directory):
  for dirpath, _, filenames in os.walk(destination_directory):
    for filename in filenames:
      file_path = os.path.join(dirpath, filename)
      if os.path.isfile(file_path):
        try:
          replace_in_file(file_path)
        except UnicodeDecodeError:
          pass

def replace_in_file(file_path):
  with fileinput.FileInput(file_path, inplace=True) as f:
    for line in f:
      line = line.replace('{{homeducky}}', name)
      print(line, end='')

def main():
  source_directory = "homeducky"
  destination_directory = path + "/" + name
  shutil.copytree(source_directory, destination_directory)
  replace_placeholders(destination_directory)

if __name__ == "__main__":
  main()
