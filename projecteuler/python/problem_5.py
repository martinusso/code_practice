'''
Created on 03/09/2011

@author: Breno

Project Euler - Problem 5

2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

Answer: 232792560
'''

def is_divisible_by_1_to_20(number):
    for n in xrange(1, 21):
        if (number % n):
            return False
    return True

positive_number = 1
while not (is_divisible_by_1_to_20(positive_number)):
    positive_number+=1
print positive_number