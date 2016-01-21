How protocols work
==================

This little example is also from Fluent Python, the Ramalho's Book.
It is from an example of Alex Martelli's excerpt.
He conied the expression **duck typing** - far far away. 
Now describe a new way to compare objects in Python. The **goose typing** - based on the comparing thorgth a ABC.
Let's see the example:

   >>> class Struggle:
   ...     def __len__(self): return 23
   ...
   >>> from collections import abc
   >>> isinstance(Struggle(), abc.Sized)
   True
   >>> issubclass(Struggle, abc.Sized)
   True

`Struggle` is a virtual subclass of `abc.Sized`

This resemble how interfaces work in go! If I defined `__len__` the instance of the class are also a `abc.Sized`! 

How it works
------------
The class `abc.Sized` has a `@classmethod` called `__subclasshook__`, that `return True` 

    if any("__len__" in B.__dict__ for B in C.__mro__)

`__mro__` is Method Resolution Order, i.e. the list of the class and its super classes.

No more `singledispatch`. Welcome to informal protocols.
