import bubble_sort from "@code/BubbleSort";

test("bubble-sort", function () {
    const arr = [9, 3, 7, 4, 69, 420, 42];

    debugger;
    bubble_sort(arr);
    expect(arr).toEqual([3, 4, 7, 9, 42, 69, 420]);
});

test("bubble-sort", function () {
    const arr = [3, 7, 1, 8, 2];

    debugger;
    bubble_sort(arr);
    expect(arr).toEqual([1, 2, 3, 7, 8]);
});


