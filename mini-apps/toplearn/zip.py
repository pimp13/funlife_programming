#!/bin/python3


names = ['pouya', 'ali', 'sara', 'hassan']
num1 = [910, 911, 912, 903]
num2 = [23, 34, 19, 32]

final_user = {t[0]: max(t[1], t[2]) for t in zip(names, num1, num2)}
print(final_user)