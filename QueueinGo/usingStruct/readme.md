# Implementing a Queue using a Struct

to resolve some of the problems encountered with using a slice, we can create a structure that will hold the underlaying slice which here we called ###Elements and its ###Size

- Using this implementation we don't need to pass the slice as an argument in all the function, instead we'll use ###methods and all changes will be directly made to our Queue using a pointer under the hood.
- Another great advantage to this implementation, is that even though we are creating a new slice when Enqueueing and Dequeueing we are not changing the address of the Queue itself.
