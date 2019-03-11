class Stack(object):
    def __init__(self, array):
        self.array = array

    def push(self, a):
        self.array.append(a)

    def pop(self):
        return self.array.pop()

    def is_empty(self) -> bool:
        return len(self.array) == 0

    def __repr__(self) -> str:
        return self.array.__repr__()


def reverse_impl(s: Stack, data) -> Stack:
    if s.is_empty():
        return Stack(data)
    
    data.append(s.pop())
    return reverse_impl(s, data)


def reverse(s: Stack) -> Stack: 
    return reverse_impl(s, [])


if __name__ == '__main__':
    s = Stack([1, 2, 3, 5])
    print(reverse(s))