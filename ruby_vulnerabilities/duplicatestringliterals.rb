#!/usr/bin/ruby
def foo()
  @ip="192.168.1.1"
  prepare('action random1')    #Noncompliant - "action random1" is duplicated 3 times
  execute('action random1')
  release('action random1')
end
