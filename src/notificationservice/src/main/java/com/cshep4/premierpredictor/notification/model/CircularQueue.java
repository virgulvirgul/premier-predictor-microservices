package com.cshep4.premierpredictor.notification.model;

import lombok.val;

import java.util.LinkedList;

public class CircularQueue<T> extends LinkedList<T> {
    private int limit;

    public CircularQueue(int limit) {
        this.limit = limit;
    }

    @Override
    public boolean offer(T element) {
        val result = super.offer(element);

        if (!result) {
            return false;
        }

        if (size() > limit) {
            poll();
        }

        return true;
    }

//    @Override
//    public T getFirst() {
//        return super.getLast();
//    }
//
//    @Override
//    public T getLast() {
//        return super.getFirst();
//    }
//
//    @Override
//    public T removeFirst() {
//        return super.removeLast();
//    }
//
//    @Override
//    public T removeLast() {
//        return super.removeFirst();
//    }
//
//    @Override
//    public void addFirst(T t) {
//        super.addLast(t);
//    }
//
//    @Override
//    public void addLast(T t) {
//        super.addFirst(t);
//    }
//
//    @Override
//    public boolean offerFirst(T t) {
//        return super.offerLast(t);
//    }
//
//    @Override
//    public boolean offerLast(T t) {
//        return super.offerFirst(t);
//    }
//
    @Override
    public T peek() {
        return super.peekLast();
    }
//
//    @Override
//    public T peekFirst() {
//        return super.peekLast();
//    }
//
//    @Override
//    public T peekLast() {
//        return super.peekFirst();
//    }
//
//    @Override
//    public T poll() {
//        return super.pollLast();
//    }
//
//    @Override
//    public T pollFirst() {
//        return super.pollLast();
//    }
//
//    @Override
//    public T pollLast() {
//        return super.pollFirst();
//    }
//
//    @Override
//    public boolean removeFirstOccurrence(Object o) {
//        return super.removeLastOccurrence(o);
//    }
//
//    @Override
//    public boolean removeLastOccurrence(Object o) {
//        return super.removeFirstOccurrence(o);
//    }
}
