#!/bin/sh
./gobuster -m dir -q -k -v -l -lu -w small_wordlist.txt -u https://waldorf.newpace.ca -o results.txt
