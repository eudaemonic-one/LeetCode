class Solution {
    public int[] getStrongest(int[] arr, int k) {
        Arrays.sort(arr);
        int i = 0, j = arr.length - 1, idx = 0;
        int median = arr[(arr.length - 1) / 2];
        int[] res = new int[k];
        while (idx < k)
            if (median - arr[i] > arr[j] - median)
                res[idx++] = arr[i++];  
            else
                res[idx++] = arr[j--];      
        return res;
    }
}