import itertools

values = [1, 2, 3, 4, 5, 6]

#for i, v in enumerate(values):
#    print(f'key: {i} ---- value: {v}')

# enumerate mikone az 10 start index az 10 be bad
#for i, v in enumerate(values, start=10):
#    print(f'key: {i} ---- value: {v}')

#def is_even(n: int) -> bool:
#    if n % 2 == 0:
#        return True
#    else:
#        return False


#for n in values:
#    if is_even(n):
#        print(f'number {n} is even')
#        break
#else:
#    print('No even number found!')

#be_tavan_2 = [n**2 for n in values]
#print(be_tavan_2)

#names = ['pouya', 'jadi', 'sara', 'hassan', 'john', 'javad']
#for a, b in itertools.product(values, names):
#    print(f'{a} {b}')

#for subset in itertools.combinations(values, 2):
#    print(subset)

#age = 17
#print('faghat blaye 18 sal' if age < 18 else 'you logged in!!')

lambda_func = lambda n: n * n
print(lambda_func(10))
"""
baraye daryaft name function
"""
def test_func(n): return n * n
print(lambda_func.__name__)
print(test_func.__name__)

















