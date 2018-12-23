package me.moye;

import java.util.concurrent.atomic.AtomicInteger;

public class Counting {
    public static void exec() throws InterruptedException {
        final AtomicInteger counter = new AtomicInteger();

        class CountingThread extends Thread {
            public void run() {
                for (var x = 0; x < 10000; ++x) {
                    counter.incrementAndGet();
                }
            }
        }

        var t1 = new CountingThread();
        var t2 = new CountingThread();

        t1.start();
        t2.start();
        t1.join();
        t2.join();

        System.out.println(counter.get());
    }
}
