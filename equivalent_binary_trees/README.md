# [Exercise: Equivalent Binary Trees](https://go.dev/tour/concurrency/8)

1. Implement the Walk function.

2. Test the Walk function.

    The function tree.New(k) constructs a randomly-structured (but always sorted) binary tree holding the values k, 2k, 3k, ..., 10k.

    Create a new channel ch and kick off the walker:

    go Walk(tree.New(1), ch)
    Then read and print 10 values from the channel. It should be the numbers 1, 2, 3, ..., 10.

3. Implement the Same function using Walk to determine whether t1 and t2 store the same values.

4. Test the Same function.

Same(tree.New(1), tree.New(1)) should return true, and Same(tree.New(1), tree.New(2)) should return false.

The documentation for Tree can be found here.
