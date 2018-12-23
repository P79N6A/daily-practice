package me.moye;

import java.util.concurrent.locks.ReentrantLock;

public class UnInterruptible {
    private static void runThread(ReentrantLock own, ReentrantLock others, String name) {
        try {
            own.lockInterruptibly();
            Thread.sleep(1000);
            others.lockInterruptibly();
        } catch (InterruptedException e) {
            System.out.println(name + " interrupted");
        }
    }

    public static void exec() throws InterruptedException{
        final var o1 = new ReentrantLock();
        final var o2 = new ReentrantLock();

        var t1 = new Thread(() -> UnInterruptible.runThread(o1, o2, "t1"));
        var t2 = new Thread(() -> UnInterruptible.runThread(o2, o1, "t2"));

        t1.start();
        t2.start();
        Thread.sleep(2000);
        t1.interrupt();
        t2.interrupt();
        t1.join();
        t2.join();
    }
}
