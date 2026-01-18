### What are Goroutines?

You and a group of friends are renting a massive Airbnb house for the weekend. Y'all are throwing a big party and need to cook up a storm.

Some dishes you need to cook include:
1. Baking a lemon cake
2. Baking strawberry cupcakes
3. Grilling some chicken
4. Cooking some goat stew

Now, the party starts in two hours. If you cook all these meals in one kitchen, it will take 8 hours to finish.

However, this massive house has 4 separate kitchens, so what do you do? You cook each meal separately in different kitchens across the house—that is a Goroutine: the ability to kick off work concurrently.

Another important concept in Go concurrency is **channels**. Imagine that the level of sugar in the lemon cake must be the same for the strawberry cupcakes, and the level is unknown before each group breaks off to their respective kitchens. Whenever the level is decided, the group that decided first can text the other cake group what the level should be. That would be Go channels—sharing memory by communicating (Communicating Sequential Processes).

---

### What is concurrency? What is parallelism?

Let's say you have 2 cores in your computer and you need to run 2 threads. Ideally, each of the CPU's will be running one thread each. This is both concurrency and parellel. 

However, if you have only one CPU, and you can only run one thread at once, even if you can't run them parallely, you can run them concurrently. That's because your computer does a lot of other small and larger tasks that it needs to carry around. This includes accessing files or talking to your network. When one of these things starts, the CPU can go work on something else that it's fast at.

The switching part in fact is much simpler in Go. If one of the goroutines gets stuck in one task, Go automatically hands over the other goroutines that are currently waiting in the OS thread to run.

---

### For this project?

Each ```go handleConn(conn)``` is put into a run queue. A thread of OS pulls these goroutines from the queue and runs them. If a goroutine blocks an I/O (disk, socket read), the scheduler pauses and runs another goroutine. So you can have thousands of idle connections, each with its own goroutine waiting on Read, without killing your machine.

**Go Philosophy:**

> "Don't communicate by sharing memory; share memory by communicating."

So instead of multiple goroutines poking the same map without protection, you usually have one goroutine own the map. Other goroutines send it messages over a channel ("user X said msg Y"). The owner goroutine updates state and maybe broadcasts out.