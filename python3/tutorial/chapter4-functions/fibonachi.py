# https://docs.python.org/3/tutorial/controlflow.html
def fib(n):
    """
    Print a Fibonacci series up to n.
    """
    a, b = 0, 1
    while a < n:
         print(a, end=' ')
         a, b = b, a+b
    print()

def fib2(n):
    """
    Collect and return a Fibonacci series up to n as an array
    """
    result = []
    a, b = 0, 1
    while a < n:
         result.append(a)
         a, b = b, a+b
    return result

print("----")
print("fib - method 1: directly print")
fib(21000)
print("----")
print("fib - method 2: move to array")
print(fib2(21000))
print("----")