class LRUCache {
    class DoublyLinkedNode {
        int key;
        int value;
        DoublyLinkedNode prev;
        DoublyLinkedNode next;
    }
    
    private void addNode(DoublyLinkedNode node) {
        node.prev = head;
        node.next = head.next;
        head.next.prev = node;
        head.next = node;
    }
    
    private void removeNode(DoublyLinkedNode node) {
        DoublyLinkedNode prev = node.prev;
        DoublyLinkedNode next = node.next;
        prev.next = next;
        next.prev = prev;
    }
    
    private void moveToHead(DoublyLinkedNode node) {
        removeNode(node);
        addNode(node);
    }
    
    private DoublyLinkedNode popTail() {
        DoublyLinkedNode node = tail.prev;
        removeNode(node);
        return node;
    }
    
    private Map<Integer, DoublyLinkedNode> cache;
    private int size;
    private int capacity;
    private DoublyLinkedNode head, tail;

    public LRUCache(int capacity) {
        this.cache = new HashMap<>();
        this.size = 0;
        this.capacity = capacity;
        head = new DoublyLinkedNode();
        tail = new DoublyLinkedNode();
        head.next = tail;
        tail.prev = head;
    }
    
    public int get(int key) {
        DoublyLinkedNode node = cache.get(key);
        if (node == null) {
            return -1;
        }
        moveToHead(node);
        return node.value;
    }
    
    public void put(int key, int value) {
        DoublyLinkedNode node = cache.get(key);
        if (node == null) {
            if (size >= capacity) {
                DoublyLinkedNode tail = popTail();
                if (tail != null) {
                    cache.remove(tail.key);
                    size -= 1;
                }
            }
            DoublyLinkedNode newNode = new DoublyLinkedNode();
            newNode.key = key;
            newNode.value = value;
            cache.put(key, newNode);
            addNode(newNode);
            size += 1;
        } else {
            node.value = value;
            moveToHead(node);
        }
    }
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * LRUCache obj = new LRUCache(capacity);
 * int param_1 = obj.get(key);
 * obj.put(key,value);
 */