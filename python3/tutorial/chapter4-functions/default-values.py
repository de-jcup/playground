# set 42 as default for i
i = 42

def myDefaultFunction(arg=i): # uses here 42 as default
    print(arg)

i = 4711 #switching to 4711
myDefaultFunction() # will still print 42


# default value evaluation only once! and reused/singleton
def f(a, L=[]):
    L.append(a)
    return L

print(f(1))
print(f(2))
print(f(3)) # prints 1,2,3...


def f2(a, L=None):
    if L is None:
        L = []
    L.append(a)
    return L


print(f2(1))
print(f2(2))
print(f2(3)) # prints 3