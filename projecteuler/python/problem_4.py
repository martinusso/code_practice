'''
Created on 03/09/2011

@author: Breno Martinusso

PRoject Euler - Problem 4

A palindromic number reads the same both ways. 
The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 x 99.

Find the largest palindrome made from the product of two 3-digit numbers.

Answer: 906609
'''

def is_palindrome(number):
    s = str(number)
    return s == s[::-1]

largest = 0
for n1 in xrange(999, 100, -1):
    for n2 in xrange(n1, 100, -1):
        calc = n1 * n2
        if is_palindrome(calc) and (calc > largest):
            largest = calc
print largest
