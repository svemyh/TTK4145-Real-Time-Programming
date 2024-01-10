
Exercise 1 

1: If you are not on the lab
----------------------------
✅

2: Version control
------------------
✅

3: Sharing a variable
---------------------

Two threads acting upon a shared variable without proper synchronization will likely produce an 'incorrect' result, contrary to the expectation that the program should print out 0 upon termination. This issue arises because both threads perform read and write operations on the same variable asynchronously, leading to race conditions. Since the increments and decrements are not atomic* and can interleave in unpredictable ways, the final result of the program can appear random and may vary with each execution.

Atomicity - In computing, an operation is considered atomic if it is completed in a single step relative to other threads. Atomic operations are indivisible. Once started, they run to completion without interruption.

4: Sharing a variable, but properly
-----------------------------------
By adding a mutex and locking and unlocking during resource acquisition of the variable i, we enforce proper synchronization of the threads. Only 1 thread can now manipulate i at the time, which prevents the race condition between the threads which we encountered before. 

Semaphores vs Mutex
Mutexes are a type of binary Semaphores, they function as binary locks so that only 1 thread can accuire the resource at a time. They also have a concept of ownership, so 
only the thread that currently has aquired the resource may unlock it (increase the count of the binary semaphore)

Sempaphores function as counting mechanism and keep a overview of how many threads are currently operating on a resource. The count can be larger then 1, so
mult threads can have access to the same resource.

In our case we want to create mutual exclusion (locking the resource for a specific thread), so Mutexes are better



5: Bounded buffer
-----------------

6: Some questions
-----------------

7: Thinking about elevators
---------------------------

8: Thinking about languages
---------------------------

Extra
=====

9: Multithreading in other languages
------------------------------------