Exercise 1 - Theory questions
-----------------------------

### Concepts

What is the difference between *concurrency* and *parallelism*?
> Concurrency: Exploiting the down-time/dead-time of the system. So while the system is waiting for a process to finish or start, it is able to complete some other task.
Parallelism: Doing several things at once. Tasks are executed simultaneously - in parallell.


What is the difference between a *race condition* and a *data race*? 
> Race condition: Order of task-completion affects the final result.
Data race: A type of race condition where more than one function/process/routine acts upon the same object. If two different processes both write and read from the same object in the same time-frame unexpected errors can occur.  
 
*Very* roughly - what does a *scheduler* do, and how does it do it?
> A scheduler decides which threads to run, which is determined by factors like priority and resource availability.

### Engineering

Why would we use multiple threads? What kinds of problems do threads solve?
> Faster completion of tasks. Utilize more of the CPU by distributing threads over several cpu cores. Threads make the most out of the down-time of a program. It can make the code more readable.

Some languages support "fibers" (sometimes called "green threads") or "coroutines"? What are they, and why would we rather use them over threads?
> *Your answer here*

Useful in environments with high levels of I/O-bound tasks or where minimizing the overhead of context switching is important.

Does creating concurrent programs make the programmer's life easier? Harder? Maybe both?
> Both. It can be harder due increased code-complexity (in some cases) but it can make programs more efficient

What do you think is best - *shared variables* or *message passing*?
> Shared variables: Faster to write, more possibilities for incorrect implementations. Vulnerable to synchronization errors.
Message passing: Provides better structure and encapsulation, reducing the likelihood of data races.


