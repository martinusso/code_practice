#!/usr/bin/python
# -*- coding: utf-8 -*-

'''
Given two strings, a and b, return the result of putting them together in the order abba, e.g. "Hi" and "Bye" returns "HiByeByeHi". 

make_abba('Hi', 'Bye') → 'HiByeByeHi'
make_abba('Yo', 'Alice') → 'YoAliceAliceYo'
make_abba('x', 'y') → 'xyyx'
'''

def make_abba(a, b):
    return a + b + b + a
