#!/usr/bin/python
# -*- coding: utf-8 -*-

'''
Return the sum of the numbers in the array, returning 0 for an empty array. Except the number 13 is very unlucky, so it does not count and numbers that come immediately after a 13 also do not count. 

sum13([1, 2, 2, 1]) → 6
sum13([1, 1]) → 2
sum13([1, 2, 2, 1, 13]) → 6
'''

def sum13(nums):
    def is_lucky():
        return nums[i] != 13 and ((nums[i-1] != 13 and i > 0) or i == 0)
    return sum([nums[i] for i in range(len(nums)) if is_lucky()])
