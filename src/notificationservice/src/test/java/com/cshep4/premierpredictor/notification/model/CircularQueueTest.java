package com.cshep4.premierpredictor.notification.model;

import lombok.val;
import org.junit.Test;

import static org.hamcrest.MatcherAssert.assertThat;
import static org.hamcrest.CoreMatchers.is;

public class CircularQueueTest {
    private static final int ONE = 1;
    private static final int TWO = 2;
    private static final int THREE = 3;
    private static final int FOUR = 4;

    @Test
    public void offerWillAddRecordAtTheEndOfTheQueueAndRemoveTheFirstElementIfTheLimitIsReached() {
        val queue = new CircularQueue<Integer>(2);

        queue.offer(ONE);
        queue.offer(TWO);
        queue.offer(THREE);
        queue.offer(FOUR);

        assertThat(queue.get(0), is(THREE));
        assertThat(queue.get(1), is(FOUR));
    }
}
