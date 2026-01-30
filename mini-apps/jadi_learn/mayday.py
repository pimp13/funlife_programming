#!/bin/python3

INPUT = 'Whiskey Hotel Four Tango Dash Alpha Romeo Three Dash Yankee Oscar Uniform Dash Sierra One November Kilo India November Golf Dash Four Bravo Zero Uniform Seven'

numbers_char = {
  'zero': 0,
  'one': 1,
  'tow': 2,
  'three': 3,
  'four': 4,
  'five': 5,
  'six': 6,
  'seven': 7,
  'eight': 8,
  'nine': 9,
  'dash': '-',
}

words = INPUT.split(' ')
# for w in words:
#   if w.lower() in numbers_char.keys():
#     print(numbers_char[w.lower()], end='')
#   else:
#     print(w[0], end='')
    
for w in words:
  print(numbers_char.get(w.lower(), w[0]), end='')

print('\n')
    
  # print (numbers_char.get(c.lower()))
  # print(c.lower())