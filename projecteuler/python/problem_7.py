'''
Created on 03/09/2011

@author: Breno Martinusso

By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?

Awnser: 104743
'''

def is_prime_number(number):
    for n in xrange(2, number/2 +1):
        if not number % n:
            return False
    return True

n, i = 1, 0
while i < 10001:
    n += 1
    if is_prime_number(n):
        i += 1
print n