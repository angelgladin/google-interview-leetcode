import (
    "sort"
)

func intersect(nums1 []int, nums2 []int) []int {
    sort.Ints(nums1)
    sort.Ints(nums2)

    var intersection []int
    i, j := 0, 0
    n, m := len(nums1), len(nums2)
    for i < n && j < m {
        if (nums1[i] < nums2[j]) {
            i++
        } else if (nums1[i] > nums2[j]) {
            j++
        } else {
            intersection = append(intersection, nums1[i])
            i++
            j++
        }
    }
    return intersection
}
