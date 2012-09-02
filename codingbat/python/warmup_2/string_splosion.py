#!/usr/bin/python
# -*- coding: utf-8 -*-
'''
Given a non-empty string like "Code" return a string like "CCoCodCode". 

string_splosion('Code') → 'CCoCodCode'
string_splosion('abc') → 'aababc'
string_splosion('ab') → 'aab'
'''

def string_splosion(str):
    r = [str[:i+1] for i in range(len(str))]
    return ''.join(r)
