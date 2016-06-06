# learning

This is a very simple go program while I was playing with channels and goroutines, trying to learn how this concepts is implemented in Go. Simple and exciting feature. 

Be careful with the details of this feature:
-unbuffered channel do not need sync but need a sender and a receiver to be available for a value/msg to be passed through the pipe, or the program will be blocked until both available.
-buffered channel do not block before filling the "buffer". If the buffer is filled faster than values are received/read from the pipe/channel, then the programm will still be blocked.

Be careful with deadlocks also.
