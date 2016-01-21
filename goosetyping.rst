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

This resemble how interfaces work in go! If I defined `__len__` the instance of the class are also a `abc.Sized`! 
No more `singledispatch`. Welcome to informal protocols.
