Exercise 1 - Theory questions
-----------------------------

### Concepts

What is the difference between *concurrency* and *parallelism*?
> Concurrency is just that different things can happen at the same time, while parallelism is actually doing multiple things in the CPU.

What is the difference between a *race condition* and a *data race*? 
> Data race is when different threads access the same variable, while a race condition is when the order the variable is accessed affects the outcome of the programme.
 
*Very* roughly - what does a *scheduler* do, and how does it do it?
> The scheduler says who gets to run when, I don't know how. 


### Engineering

Why would we use multiple threads? What kinds of problems do threads solve?
> It enables multiple things to happen at once, often increasing speed and capacity. Also, sometimes we need things to happen at the same time, and they make the code more readable.

Some languages support "fibers" (sometimes called "green threads") or "coroutines"? What are they, and why would we rather use them over threads?
> They are lightweight, and don't always require the CPU, and can run while the CPU is working. Multiple fibers can also run i a thread, and they are inherently synchrised.

Does creating concurrent programs make the programmer's life easier? Harder? Maybe both?
> Concurrent programs are not used to make stuff harder or easier, it is used to solve problems that can not be solved without.

What do you think is best - *shared variables* or *message passing*?
> Message passing, due to the nice queue aspect, rather than getting race conditions and data races.


