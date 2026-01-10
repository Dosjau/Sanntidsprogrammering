Exercise 1 - Theory questions
-----------------------------

### Concepts

What is the difference between *concurrency* and *parallelism*?
> Concurrency is managing several taks with different threads and interleaving to create an illusion of parallell processing. Parallelism is to actually handle these tasks in parallell, but this requires multiple CPU-cores which can allow the tasks to run truly parallell.

What is the difference between a *race condition* and a *data race*? 
> A race condition is a flaw with a program that is the outcome of unpredictable timing or interleaving of events. A data race is more specific where it is the issue of two or more threads that access the same memory location at the same time. At least one access is a write, and it happens without proper synchronization. 
 
*Very* roughly - what does a *scheduler* do, and how does it do it?
> A schedulers job is to manage all the different tasks of the CPU. It descides when each process or thread should run for the computer to run optimally. It does this by switching between processes, making sure that all are handled efficiently.


### Engineering

Why would we use multiple threads? What kinds of problems do threads solve?
> Multiple threads are used to allow different tasks in the program to be handled concurrently. This solves a problem where large or repeating tasks in the program occupy all the runtime, making it so other tasks can't run. By creating different threads they can be scheduled such that all tasks can be handled optimally.

Some languages support "fibers" (sometimes called "green threads") or "coroutines"? What are they, and why would we rather use them over threads?
> Fibers are a different type of execution that is handled in user space, rather than the operating system. An example is Go's goroutines, and they are scheduled by the language runtime. This makes them cheaper to create and switch between. Fibers are preferable when there are a lot of concurrent tasks that need to be handled, as they reduce overhead and avoid many of the synchronization and context-switching costs that are associated with threads.

Does creating concurrent programs make the programmer's life easier? Harder? Maybe both?
> Creating a concurrent program can make the programmer's life both easier and harder, depending on the program. If the program requires several tasks to be handled at the same time then it can be impossible or at least very hard to create without using concurrency. However, for programs that does not require simultanious processing then it can just be extra work and overhead. If all tasks can just as well be done in chronological order, then it is best to solve it without concurrency.

What do you think is best - *shared variables* or *message passing*?
> I don't have a preference for either shared variables or message passing. They are two different types of ways to manage the resources in a program, that both have advantages and disadvantages. Shared memory is often faster, but introduces the risk of race conditions that are not present for message passing.


