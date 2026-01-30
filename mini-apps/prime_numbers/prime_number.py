def is_prime(n):
  prime = True
  for i in range(2, n):
    if n % i == 0:
      prime = False
  return prime


prime_count = 0
for i in range(1, 101):
  if is_prime(i):
    prime_count += 1

print('\nprime count', prime_count)
