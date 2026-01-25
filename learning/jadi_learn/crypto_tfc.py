#!/bin/python3

crypto = "T4 l16 _36 510 _27 s26 _11 320 414 {6 }39 C2 T0 m28 y35 d31 F1 m22 g19 d38 z34 423 l15 329 c12 ;37 19 h13 _30 F5 t7 C3 325 z33 _21 h8 n18 132 k24"
parts = crypto.split(' ')

answer = []
for i in range(100):
  answer.append('*')
  
for p in parts:
  char = p[0]
  i = int(p[1:])
  # answer.insert(i, char)
  answer[i] = char
  # print(''.join(answer))
  
  
print(''.join(answer))