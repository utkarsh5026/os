### Threads in Java 🧵

Threads in Java are essential for building concurrent applications. Java provides a robust and easy-to-use model for multithreading, allowing developers to create applications that can perform multiple tasks simultaneously. Let's explore everything you need to know about threads in Java.

---

![1720637232298](image/threads/1720637232298.png)

#### Definition

**Thread**: A thread is a lightweight subprocess, the smallest unit of processing. It is a separate path of execution within a program. Java's `java.lang.Thread` class is used to create and manage threads.

---

#### Why Use Threads?

1. **Concurrency** 🌐

   - Allows multiple tasks to be performed simultaneously, improving application performance and responsiveness.
2. **Resource Sharing** 🛠️

   - Threads within the same process share resources, such as memory, which can be more efficient than using multiple processes.
3. **Responsiveness** 🏃‍♂️

   - Keeping the user interface responsive while performing background tasks.
4. **Utilizing Multi-core Processors** 💻

   - Modern processors have multiple cores, and multithreading allows applications to utilize these cores efficiently.

---

#### Creating Threads in Java

There are two main ways to create threads in Java:

1. **Extending the `Thread` Class**

   - **Mechanism**: Create a new class that extends `Thread` and override the `run()` method.
   - **Example**:
     ```java
     class MyThread extends Thread {
         public void run() {
             System.out.println("Thread is running");
         }
     }

     public class TestThread {
         public static void main(String[] args) {
             MyThread t1 = new MyThread();
             t1.start();  // Start the thread
         }
     }
     ```
2. **Implementing the `Runnable` Interface**

   - **Mechanism**: Create a class that implements the `Runnable` interface and implement the `run()` method. Pass an instance of the class to a `Thread` object and call `start()`.
   - **Example**:
     ```java
     class MyRunnable implements Runnable {
         public void run() {
             System.out.println("Thread is running");
         }
     }

     public class TestRunnable {
         public static void main(String[] args) {
             MyRunnable myRunnable = new MyRunnable();
             Thread t1 = new Thread(myRunnable);
             t1.start();  // Start the thread
         }
     }
     ```

---

#### Thread Lifecycle

Threads in Java go through several states during their lifecycle:

1. **New** 🌱

   - A thread that is created but not yet started.
   - **Code**: `Thread t = new Thread();`
2. **Runnable** 🎬

   - A thread that is ready to run and waiting for CPU allocation.
   - **Code**: `t.start();`
3. **Running** 🏃‍♂️

   - A thread that is currently executing.
   - **Code**: The `run()` method is being executed.
4. **Blocked/Waiting** ⏳

   - A thread that is waiting for a resource or another thread to perform an action.
   - **Code**: `t.wait();` or `t.join();`
5. **Timed Waiting** ⏲️

   - A thread that is waiting for a specified amount of time.
   - **Code**: `t.sleep(1000);`
6. **Terminated** 🛑

   - A thread that has completed its execution.
   - **Code**: The `run()` method has finished.

---

#### Thread Methods

Java provides several methods to control the behavior of threads:

1. **`start()`** 🎬

   - Starts the thread.
   - **Code**: `t.start();`
2. **`run()`** 🏃‍♂️

   - Contains the code to be executed by the thread.
   - **Code**: `public void run() { ... }`
3. **`sleep(long millis)`** 😴

   - Makes the thread sleep for the specified number of milliseconds.
   - **Code**: `Thread.sleep(1000);`
4. **`join()`** ⏳

   - Waits for the thread to die.
   - **Code**: `t.join();`
5. **`yield()`** 🔄

   - Pauses the currently executing thread to allow other threads to execute.
   - **Code**: `Thread.yield();`
6. **`interrupt()`** 🚨

   - Interrupts the thread.
   - **Code**: `t.interrupt();`
7. **`isAlive()`** 🔍

   - Checks if the thread is still running.
   - **Code**: `t.isAlive();`

---

#### Synchronization

Synchronization is crucial when multiple threads access shared resources to prevent data inconsistency and ensure thread safety.

1. **Synchronized Methods** 🔒

   - Methods can be synchronized to allow only one thread to access them at a time.
   - **Code**:
     ```java
     public synchronized void synchronizedMethod() {
         // Code
     }
     ```
2. **Synchronized Blocks** 🔒

   - Blocks within methods can be synchronized to lock on a particular object.
   - **Code**:
     ```java
     public void method() {
         synchronized(this) {
             // Code
         }
     }
     ```
3. **Static Synchronization** 🔒

   - Static methods can also be synchronized.
   - **Code**:
     ```java
     public static synchronized void staticSynchronizedMethod() {
         // Code
     }
     ```

**Example**:

```java
class Counter {
    private int count = 0;

    public synchronized void increment() {
        count++;
    }

    public int getCount() {
        return count;
    }
}

public class TestSync {
    public static void main(String[] args) throws InterruptedException {
        Counter counter = new Counter();

        Thread t1 = new Thread(() -> {
            for (int i = 0; i < 1000; i++) {
                counter.increment();
            }
        });

        Thread t2 = new Thread(() -> {
            for (int i = 0; i < 1000; i++) {
                counter.increment();
            }
        });

        t1.start();
        t2.start();

        t1.join();
        t2.join();

        System.out.println("Count: " + counter.getCount());
    }
}
```

---

#### Inter-Thread Communication

Java provides a way for threads to communicate with each other using the `wait()`, `notify()`, and `notifyAll()` methods.

1. **`wait()`** ⏳

   - Causes the current thread to wait until another thread invokes `notify()` or `notifyAll()` on the same object.
   - **Code**: `synchronized(obj) { obj.wait(); }`
2. **`notify()`** 🔔

   - Wakes up a single thread that is waiting on the object's monitor.
   - **Code**: `synchronized(obj) { obj.notify(); }`
3. **`notifyAll()`** 🔔🔔

   - Wakes up all threads that are waiting on the object's monitor.
   - **Code**: `synchronized(obj) { obj.notifyAll(); }`

**Example**:

```java
class SharedResource {
    private boolean flag = false;

    public synchronized void produce() throws InterruptedException {
        while (flag) {
            wait();
        }
        System.out.println("Produced");
        flag = true;
        notify();
    }

    public synchronized void consume() throws InterruptedException {
        while (!flag) {
            wait();
        }
        System.out.println("Consumed");
        flag = false;
        notify();
    }
}

public class TestInterThread {
    public static void main(String[] args) {
        SharedResource resource = new SharedResource();

        Thread producer = new Thread(() -> {
            try {
                resource.produce();
            } catch (InterruptedException e) {
                Thread.currentThread().interrupt();
            }
        });

        Thread consumer = new Thread(() -> {
            try {
                resource.consume();
            } catch (InterruptedException e) {
                Thread.currentThread().interrupt();
            }
        });

        producer.start();
        consumer.start();
    }
}
```

---

#### Thread Pools

Java provides the `Executor` framework to manage a pool of worker threads.

1. **Fixed Thread Pool** 🔢

   - A thread pool with a fixed number of threads.
   - **Code**:
     ```java
     ExecutorService executor = Executors.newFixedThreadPool(5);
     for (int i = 0; i < 10; i++) {
         executor.submit(() -> {
             System.out.println("Task executed by " + Thread.currentThread().getName());
         });
     }
     executor.shutdown();
     ```
2. **Cached Thread Pool** 📈

   - A thread pool that creates new threads as needed but reuses previously constructed threads when they are available.
   - **Code**:
     ```java
     ExecutorService executor = Executors.newCachedThreadPool();
     for (int i = 0; i < 10; i++) {
         executor.submit(() -> {
             System.out.println("Task executed by " + Thread.currentThread().getName());
         });
     }
     executor.shutdown();
     ```
3. **Scheduled Thread Pool** 📅

   - A thread pool that can schedule commands to run after a given delay or to execute periodically.
   - **Code**:
     ```java
     ScheduledExecutorService executor = Executors.newScheduledThreadPool(5);
     executor.schedule(() -> {
         System.out.println("Scheduled task executed by " + Thread.currentThread().getName());
     }, 5,

     ```

 TimeUnit.SECONDS);
     executor.shutdown();
     ```

---

#### Best Practices

1. **Use Thread Pools** 🏊‍♂️

   - Use the `Executor` framework instead of creating threads manually to manage thread life cycle efficiently.
2. **Minimize Synchronization** 🔒

   - Synchronize only the necessary parts of the code to reduce contention and improve performance.
3. **Handle InterruptedException** ⚠️

   - Properly handle `InterruptedException` to maintain thread responsiveness and manage cleanup.
4. **Avoid Deadlocks** 🚧

   - Design thread synchronization carefully to prevent deadlocks.
5. **Use High-Level Concurrency Utilities** 🛠️

   - Utilize `java.util.concurrent` package utilities such as `CountDownLatch`, `CyclicBarrier`, and `Semaphore` for advanced synchronization.

---

Understanding threads in Java is crucial for developing efficient, responsive, and high-performance applications. If you have any questions or need further explanations, feel free to ask!
