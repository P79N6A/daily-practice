package me.moye;

import me.moye.utils.Tuple;

import java.util.ArrayList;
import java.util.List;
import java.util.Random;
import java.util.concurrent.locks.Condition;
import java.util.concurrent.locks.ReentrantLock;

public class Philosopher extends Thread {
    private boolean eating;
    private Philosopher left;
    private Philosopher right;
    private ReentrantLock table;
    private Condition condition;
    private Random random;

    public Philosopher(ReentrantLock table) {
        eating = false;
        this.table = table;
        condition = table.newCondition();
        random = new Random();
    }

    public void setLeft(Philosopher left) { this.left = left; }
    public void setRight(Philosopher right) { this.right = right; }

    public void run() {
        try {
            while(true) {
                think();
                eat();
            }
        } catch (InterruptedException e) { }
    }

    private void think() throws InterruptedException {
        table.lock();
        try {
            eating = false;
            left.condition.signal();
            right.condition.signal();
        } finally { table.unlock(); }
        Thread.sleep(1000);
    }

    private void eat() throws InterruptedException {
        table.lock();
        try {
            while (left.eating || right.eating)
                condition.await();
            eating = true;
        } finally { table.unlock(); }
        Thread.sleep(1000);
    }

    private static Tuple<List<Philosopher>, ReentrantLock> init(int n) {
        var table = new ReentrantLock();
        var list = new ArrayList<Philosopher>();
        var result = new Tuple(list, table);

        for (var i = 0; i < n; i++) {
            var p = new Philosopher(table);
            list.add(p);
        }

        var total = list.size();
        for (var i = 0; i < total; i++) {
            var left = i - 1;
            var right = i + 1;
            if (left < 0)
                left = total - 1;
            if (right == total)
                right = 0;
            list.get(i).setLeft(list.get(left));
            list.get(i).setRight(list.get(right));
        }

        return result;
    }

    public static void exec(int n) throws InterruptedException {
        var tuple = Philosopher.init(n);
        for (Philosopher p : tuple.x) {
            p.run();
        }
    }
}
