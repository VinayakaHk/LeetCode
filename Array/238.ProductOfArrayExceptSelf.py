def productExceptSelf(nums):
    # Get the length of input array
    n = len(nums)
    # Initialize result array with 1s
    res = [1] * (n)
    print("res", res)

    # Calculate prefix products (products of all elements to the left)
    prefix = 1
    for i in range(n):
        print("i", i)
        # Store the product of all elements to the left of current position
        res[i] = prefix
        print("res[i]", res[i])
        # Update prefix by multiplying with current number
        prefix = prefix * nums[i]
        print("prefix", prefix)

    print('res',res)
    print("-------")

    # Calculate postfix products (products of all elements to the right)
    postfix = 1
    # Iterate from right to left
    for i in range(n -1, -1, -1):
        # Multiply current position with postfix to get final result
        res[i]  = res[i] *  postfix
        print("res[i]", res[i])
        # Update postfix by multiplying with current number
        postfix = postfix * nums[i]
        print("postfix", postfix)

    print('res',res)
    print("-------")
    return res

# Test the function with example array [1,2,3,4]
print(productExceptSelf([1,2,3,4]))