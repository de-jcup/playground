def standard_arg(arg):
    print(arg)

def pos_only_arg(arg,/):
    print(arg)

def kwd_only(*,arg):
    print(arg)

def combined_example(pos_only, /, standard, *, kwd_only):
    print(pos_only,standard,kwd_only)


standard_arg(1)
standard_arg(arg=1) # both possible

pos_only_arg(2) # only pos based

kwd_only(arg=3) # only keyword

combined_example(1,2,kwd_only=3) # combined ways
combined_example(1,standard=2,kwd_only=3) # combined ways
