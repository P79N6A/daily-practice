def find_insertion(sub_seq, start, end, value):
    print(f'end:{sub_seq[end]} current:{value}, sub_seq: {sub_seq[0:end]}')
    if sub_seq[end] < value:
        return end+1
    while start < end:
        middle = int(start + (end - start)/2)
        if sub_seq[middle] <= value:
            start = middle + 1
        else:
            end = middle
    return start

def longest_sub_seq(seq):
    n = len(seq)
    sub_seq = [0] * n
    sub_seq[0] = seq[0]
    end = 0
    for i in range(0, n):
        pos = find_insertion(sub_seq, 0, end, seq[i])
        sub_seq[pos] = seq[i]
        if end < pos:
            end = pos

    print(sub_seq[0:end+1])
    return end+1


if __name__ == '__main__':
    seq = [2,1,5,3,6,4,8,9,7]
    print(longest_sub_seq(seq))
