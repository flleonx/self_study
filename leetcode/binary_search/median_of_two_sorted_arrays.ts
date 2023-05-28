function findMedianSortedArrays(nums1: number[], nums2: number[]): number {
  let A = nums1;
  let B = nums2;
  const half = Math.floor((A.length + B.length) / 2);

  if (A.length < B.length) {
    A = nums2;
    B = nums1;
  }

  let l = 0;
  let r = A.length - 1;

  while (true) {
    const cut1 = l + Math.floor((r - l) / 2);
    const cut2 = half - cut1 - 2;

    const l1 = cut1 < 0 ? -Infinity : A[cut1];
    const l2 = cut2 < 0 ? -Infinity : B[cut2];
    const r1 = cut1 + 1 < A.length ? A[cut1 + 1] : Infinity;
    const r2 = cut2 + 1 < B.length ? B[cut2 + 1] : Infinity;

    // l1 < r2
    // l2 < r1

    if (l1 <= r2 && l2 <= r1) {
      return (A.length + B.length) % 2
        ? Math.min(r1, r2)
        : (Math.max(l1, l2) + Math.min(r1, r2)) / 2;
    } else if (l1 > r2) {
      r = cut1 - 1;
    } else if (l2 > r1) {
      l = cut1 + 1;
    }
  }
}

// console.log("findMedianSortedArrays", findMedianSortedArrays([1, 3], [2]));
// console.log("findMedianSortedArrays", findMedianSortedArrays([1, 2], [3, 4]));
// console.log("findMedianSortedArrays", findMedianSortedArrays([0, 0], [0, 0]));
console.log("findMedianSortedArrays", findMedianSortedArrays([], [1]));
