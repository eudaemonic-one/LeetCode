class H2O {
    private Semaphore semOxygen;
    private Semaphore semHydrogen;
    private Phaser phaser;

    public H2O() {
        semOxygen = new Semaphore(1);
        semHydrogen = new Semaphore(2);
        phaser = new Phaser(3);
    }

    public void hydrogen(Runnable releaseHydrogen) throws InterruptedException {
		semHydrogen.acquire();
        releaseHydrogen.run();
        phaser.arriveAndAwaitAdvance();
        phaser.arriveAndAwaitAdvance();
        semHydrogen.release();
    }

    public void oxygen(Runnable releaseOxygen) throws InterruptedException {
		semOxygen.acquire();
        phaser.arriveAndAwaitAdvance();
		releaseOxygen.run();
        phaser.arriveAndAwaitAdvance();
        semOxygen.release();
    }
}