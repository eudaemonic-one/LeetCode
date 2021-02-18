class Solution {
    public int[] sortByBits(int[] arr) {
        Integer[] arr1 = new Integer[arr.length];
        for (int i = 0; i < arr.length; i++) {
            arr1[i] = arr[i];
        }
        Arrays.sort(arr1, new Comparator<>(){
            @Override
            public int compare(Integer a, Integer b) {
                int onesA = Integer.bitCount(a);
                int onesB = Integer.bitCount(b);
                if (onesA == onesB) {
                    return a - b;
                }
                return onesA - onesB;
            }
        });
        for (int i = 0; i < arr.length; i++) {
            arr[i] = arr1[i];
        }
        return arr;
    }
}