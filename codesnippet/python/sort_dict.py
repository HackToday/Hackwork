#! /usr/bin/env python

my_list = [{'name': 'Tom', 'age': 23}, {'name': 'Kate', 'age': 30}, {'name': 'Poke', 'age': 36}]
sorted_list = sorted(my_list, key=lambda k: k['name'])

print "Before sort, the list is:"
print my_list
print "After sort, the list is:"
print sorted_list
