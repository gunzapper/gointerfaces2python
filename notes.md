# How to use go in python

### Rigth version of go


To import *.so in python, you need to activate (with gvm) the same version of go used to build.
With other versions `ctypes` fails to import the `sum.so`.

The error raised is:

    OSError: sum.go: cannot open shared object file: No such file or directory

### Explore the dll

By the answers to the question 2891493 of stack over flow.
It is impossible to see the list of functions using `ctypes`,
you can in *nix 

    $ nm -D 'library.so'

If I use `sum.go` I see `Sum`, but with `interfaces.so` I do not find `MyFloat` (a user define `type`) and `Abser` (a `interface`).

#### Other note
With 

    $ go build -buildmode=c-shared -o sum.so sum.go

I produce also `sum.h`, but using the same istruction for `first_example.go`,
I do not have any `*.h`.

