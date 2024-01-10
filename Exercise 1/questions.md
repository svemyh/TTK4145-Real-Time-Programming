Exercise 1 - Theory questions
-----------------------------

### Concepts

What is the difference between *concurrency* and *parallelism*?
> *Your answer here*
    Concurrency: Refers to the ability of different parts or tasks of a system to be executed out-of-order or in partial order without affecting the final outcome. 
    It's about dealing with multiple tasks being in progress simultaneously, but not necessarily executing at the exact same time.

    Parallelism: Involves the simultaneous execution of multiple tasks or processes to increase the overall speed and efficiency of a system. 
    It implies the simultaneous execution of multiple instructions at the same time, often with the goal of improving performance by utilizing multiple processors or cores.

What is the difference between a *race condition* and a *data race*? 
> *Your answer here* 
    Race Condition: A race condition occurs when the behavior of a program depends on the relative timing of events, such as the order of execution of concurrent threads. 
    It can lead to unpredictable and undesirable outcomes if the timing of events is not properly controlled.

    Data Race: A data race specifically refers to a situation in concurrent programming where two or more threads access shared data concurrently, 
    and at least one of them modifies the data. 
    If the access is not properly synchronized, it can lead to unexpected behavior and data corruption.

*Very* roughly - what does a *scheduler* do, and how does it do it?
> *Your answer here* 
    A scheduler is a component of an operating system responsible for determining the execution order of processes or threads. It decides which tasks are executed and when. 
    The scheduler uses scheduling algorithms to allocate resources and manage the execution of tasks efficiently.

### Engineering

Why would we use multiple threads? What kinds of problems do threads solve?
> *Your answer here*
    Threads are used to solve problems that can be decomposed into smaller, independent tasks that can be executed concurrently.
    Improved Responsiveness: In graphical user interfaces, threads can be used to keep the user interface responsive while performing background tasks

Some languages support "fibers" (sometimes called "green threads") or "coroutines"? What are they, and why would we rather use them over threads?
> *Your answer here*
    These are lightweight, cooperative concurrency units. 
    They are managed by the application rather than the operating system, allowing for more control over execution. They are suitable for tasks with many blocking operations.

Does creating concurrent programs make the programmer's life easier? Harder? Maybe both?
> *Your answer here*
    Easier for tasks that naturally lend themselves to parallel execution. 
    Harder due to the complexity of managing synchronization, avoiding race conditions, and ensuring correct behavior in a concurrent environment. 
    Bugs that may not be discovered in the code, only use

What do you think is best - *shared variables* or *message passing*?
> *Your answer here*
    Shared Variables: Involves threads communicating by sharing variables. It requires synchronization mechanisms to avoid race conditions.
    Message Passing: Involves communication between threads by passing messages. It can simplify synchronization but may introduce overhead.


