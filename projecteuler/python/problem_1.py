'''
Created on 22/08/2011

@author: Breno Martinusso <breno@martinusso.com>

Problem 1
05 October 2001

If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
'''

s = 0
for i in xrange(1, 1000):# In Python 3 use range()
    s+= i if not i % 3 or not i % 5 else 0
print s
