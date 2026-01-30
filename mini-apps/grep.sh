#!/bin/zsh

cat /usr/share/dict/words | tr '[:upper:]' '[:lower:]' | egrep -v '[anie]' | egrep '^.{5}$' | egrep '^s...t$' | egrep -v 'poh' | egrep '[r]'
