#!/bin/python3

import math

def is_prime(n):
    if n <= 1:
        return False
    if n == 2:
        return True
    if n % 2 == 0:
        return False

    limit = int(math.sqrt(n))
    for i in range(3, limit + 1, 2):
        if n % i == 0:
            return  False

    return True


#number = int(input("check is prime number, please enter number: "))

for i in range(100, 10000):
    if is_prime(i):
        print(f"number {i} is prime")
    else:
        print("not prime")
