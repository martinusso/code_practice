'''
Created on 03/09/2011

@author: Breno Martinusso

Project Euler - Problem 6

The sum of the squares of the first ten natural numbers is,

1**2 + 2**2 + ... + 10**2 = 385
The square of the sum of the first ten natural numbers is,

(1 + 2 + ... + 10)**2 = 55**2 = 3025
Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 - 385 = 2640.

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

Awnser: 25164150
'''

sum_squares = sum(n**2 for n in xrange(1, 101))
square_sum = sum(xrange(1, 101))**2

print square_sum - sum_squares
    

