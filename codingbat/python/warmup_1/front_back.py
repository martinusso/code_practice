#!/usr/bin/python
# coding: utf-8

'''
Given a string, return a new string where the first and last chars have been exchanged. 

front_back('code') -> 'eodc'
front_back('a') -> 'a'
front_back('ab') -> 'ba'
'''

def front_back(str):
    if not len(str) > 1:
        return str
    first = str[0]
    last = str[-1:]
    mid = str[1:-1]
    return last + mid + first
