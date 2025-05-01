# Create a function with two arguments that will return an array of the first n multiples of x.

# Assume both the given number and the number of times to count will be positive numbers greater than 0.

# Return the results as an array or list ( depending on language ).

# Examples
# x = 1, n = 10 --> [1,2,3,4,5,6,7,8,9,10]
# x = 2, n = 5  --> [2,4,6,8,10]

def count_by(x, n):
    """
    Return a sequence of numbers counting by `x` `n` times.
    """
    result = []
    for i in range(n):
        result.append(x*(i+1))
    return result

print(count_by(1, 5), [1, 2, 3, 4, 5])
print(count_by(2, 5), [2, 4, 6, 8, 10])
print(count_by(3, 5), [3, 6, 9, 12, 15])
print(count_by(50, 5), [50, 100, 150, 200, 250])
print(count_by(100, 5), [100, 200, 300, 400, 500])

# https://www.codewars.com/kata/5513795bd3fafb56c200049e/python

# def count_by(x, n):
#    return [i * x for i in range(1, n + 1)]

# def count_by(x, n):
#     """
#     Return a sequence of numbers counting by `x` `n` times.
#     """
#     return list(range(x, n * x + 1, x))

# very sleek solutions - definitely too early days to be able to do this, I wonder if Go has any similar solution