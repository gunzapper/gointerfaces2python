Import using ctypes
===================

For the moment I use ctypes to import .so file, compiled from go

  >>> from ctypes import CDLL
  >>> foo = CDLL("sum.so")
  >>> foo.Sum(40, 2)
  42
  >>> foo.Sum(2, 40)
  42
  >>> foo.Sum(88, 2)
  90
