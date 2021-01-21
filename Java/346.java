// Approach #2: Double-ended Queue
class MovingAverage {
    private int size, windowSum = 0, count = 0;
    private Deque<Integer> queue;
    
    /** Initialize your data structure here. */
    public MovingAverage(int size) {
        this.size = size;
        this.queue = new ArrayDeque<>();
    }
    
    public double next(int val) {
        ++count;
        queue.add(val);
        int tail = count > size ? queue.poll() : 0;
        windowSum = windowSum - tail + val;
        return windowSum * 1.0 / Math.min(size, count);
    }
}

// Approach #1: Array or List
// class MovingAverage {
//     private int size;
//     private List<Integer> queue;
    
//     /** Initialize your data structure here. */
//     public MovingAverage(int size) {
//         this.size = size;
//         this.queue = new ArrayList<>();
//     }
    
//     public double next(int val) {
//         queue.add(val);
//         int sum = 0;
//         for (int i = Math.max(0, queue.size()-size); i < queue.size(); i++) {
//             sum += queue.get(i);
//         }
//         return sum * 1.0 / Math.min(queue.size(), size);
//     }
// }

/**
 * Your MovingAverage object will be instantiated and called as such:
 * MovingAverage obj = new MovingAverage(size);
 * double param_1 = obj.next(val);
 */