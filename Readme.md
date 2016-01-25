# Go interfaces in python

This repository contains code used for my presentation at Milano python meetup.
I would try to realize something similar to go interfaces in python.
The first thing is to start for an go example. I download and modified the code from go tour.

To manage the version I have to install [gvm](http://www.hostingadvice.com/how-to/install-golang-on-ubuntu/).

###sum.go 


It is the first file from a [tutorial](https://blog.filippo.io/building-python-modules-with-go-1-5/) that learn how to build python modules with go.
Basically, it uses `buildmode` option to have a static library

      $ go build -buildmode=c-shared -o sum.so sum.go

It produces also a `*.h` file, with also the functions defined in the go file.
You can inspect the contain of a `*.so` file with

    $ nm -D sum.so

And filter the results

    $ nm -D sum.so | grep "Sum"

The `main.c` is only used to test `sum.so`

    $ gcc -Wall -o main main.c ./sum.so

As described in the tutorial mentioned.
Now, we can build a *C* python module with go...

Or with `ctypes` we can directly use the functions from a `*.so` file. 
The `import.rst` tests with `sum.so` exactly this opportunity.

To launch the test

    $ python -m doctest -v import.rst


### abser.go 

It is the [interface example](https://tour.golang.org/methods/4) of the go-tour.

If I build the `abser.so` I don't obtain the `abser.h`, and inspecting `abser.so` there are not `Abs` or `Abser`.
The `so` contains only functions not relative to an interface!


### singledispatch.py

This file is from the github of [Ramalho's book](https://github.com/fluentpython/example-code/blob/a3bf32b9fa8b80f714a899ee832145e1091cc635/07-closure-deco/generic.py).

I want to mimic in go this file...


### Wrong file

It works, but Is it what I want? Maybe no...but...


### The right file

But is it what we want?

### Goose typing

The last file describe how duck typing and the new coined **goose typing**work. 
It show that informal protocols in python are exactly how the interfaces in go. Use `singlesdispatch` in some case help to write code with
less `if\less` blocks.

#### N.B.:

Because python is a language **EAFP** - *Easier to ask forgiveness than permission*.
On the contrary of C's **LBYL** - *Look before you leap*.


## Further readings

  - (Go for python programmers) [https://golang-for-python-programmers.readthedocs.org/en/latest/]



