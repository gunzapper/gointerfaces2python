# How to use go in python

To import *.so in python, you need to activate (with gvm) the same version of go used to build.
With other versions `ctypes` fails to import the `sum.so`.

The error raised is:

    OSError: sum.go: cannot open shared object file: No such file or directory
