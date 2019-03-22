from typing import List


def common_length(s):
    l = len(s)
    prefix = s[0:l-1]
    suffix = s[1:l]
    while len(prefix) > 0:
        l = len(prefix)
        if prefix == suffix:
            return l
        prefix = prefix[0:l-1]
        suffix = suffix[1:l]
    return 0


def prefix_table(s: str):
    result = []
    for i in range(1, len(s)+1):
        sub = s[0:i]
        result.append(common_length(sub))
    result.insert(0, -1)
    result.pop()
    return result


def kmp(s: str, q: str):
    t = prefix_table(s)
    ls = len(s)
    lq = len(q)
    i = 0
    j = 0
    while i < ls:
        if s[i] != q[j]:
            p = t[j]
            if p == -1:
                j = 0
                i += 1
            else:
                j = p
        else:
            j += 1
            i += 1            
        if j == lq:
            return i-j
    return -1


if __name__ == "__main__":
    s = 'abaacababcac'
    q = 'ababc'
    print(kmp(s, q))
