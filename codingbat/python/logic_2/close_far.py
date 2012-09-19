#!/usr/bin/python
# coding: utf-8

'''
Given three ints, a b c, return True if one of b or c is "close" (differing from a by at most 1), while the other is "far", differing from both other values by 2 or more. Note: abs(num) computes the absolute value of a number. 

close_far(1, 2, 10) → True
close_far(1, 2, 3) → False
close_far(4, 1, 3) → True
'''

def close_far(a, b, c):
    b_a_diff = abs(b - a)
    c_a_diff = abs(c - a)
    b_c_diff = abs(b - c)
    is_close = (b_a_diff <= 1) or (c_a_diff <= 1)
    is_far = ((b_a_diff >= 2) or (c_a_diff >= 2)) and (b_c_diff >= 2)
    return is_close and is_far
