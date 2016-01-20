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

    $ nm -D 'sum.so'

And filter the results

    $ nm -D 'sum.so' | grep "Sum"

### first_example.go 

It is the [first example](https://tour.golang.org/methods/4) of the tour.
  

## Further readings

  - (Go for python programmers) [https://golang-for-python-programmers.readthedocs.org/en/latest/]



