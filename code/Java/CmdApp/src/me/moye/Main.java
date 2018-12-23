package me.moye;

public class Main {

    public static void main(String[] args) throws InterruptedException {
	    Puzzle.t1.start();
	    Puzzle.t2.start();
	    Puzzle.t1.join();
	    Puzzle.t2.join();
    }
}
