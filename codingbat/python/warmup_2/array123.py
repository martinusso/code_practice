#!/usr/bin/python
# -*- coding: utf-8 -*-

'''
Given an array of ints, return True if .. 1, 2, 3, .. appears in the array somewhere. 

array123([1, 1, 2, 3, 1]) → True
array123([1, 1, 2, 4, 1]) → False
array123([1, 1, 2, 1, 2, 3]) → True
'''

def array123(nums):
    exists_1 = nums.count(1) > 0
    exists_2 = nums.count(2) > 0
    exists_3 = nums.count(3) > 0
    return exists_1 and exists_2 and exists_3
