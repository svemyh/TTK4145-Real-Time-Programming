
Exercise 1 

1: If you are not on the lab ✅
----------------------------

2: Version control ✅
------------------

3: Sharing a variable ✅
---------------------

Two threads acting upon a shared variable without proper synchronization will likely produce an 'incorrect' result, contrary to the expectation that the program should print out 0 upon termination. This issue arises because both threads perform read and write operations on the same variable asynchronously, leading to race conditions. Since the increments and decrements are not atomic* and can interleave in unpredictable ways, the final result of the program can appear random and may vary with each execution.

Atomicity - In computing, an operation is considered atomic if it is completed in a single step relative to other threads. Atomic operations are indivisible. Once started, they run to completion without interruption.

4: Sharing a variable, but properly
-----------------------------------



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