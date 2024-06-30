#!/usr/bin/env python

#
#   Print the 'n'th number in the fibonacci sequence
#       Took me an embarrassingly long time to figure 
#       that out and why it just wasn't printing the
#       sequence.
#
import functools

@functools.cache
def fib(n):
    if n <= 1:
        return n
    else:
        sum = fib(n-1) + fib(n-2)
        return sum
    
print(fib(40))