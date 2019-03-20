def longest_sub_seq(seq):
    n = len(seq)
    d = [1] * n
    max = 1
    for i in range(0, n):
        for j in range(0, i):
            if seq[j] <= seq[i] and d[j]+1 > d[i]:
                d[i] = d[j] + 1
        if d[i] > max:
            max = d[i]
    return max


if __name__ == '__main__':
    seq = [2,1,5,3,6,4,8,9,7]
    print(longest_sub_seq(seq))
