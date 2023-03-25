def toBinary(n, len):
    binary = ''
    i = 1 << len - 1
    while i > 0:
        binary += '1' if (n & i) else '0'
        i = i // 2
    return binary
 
if __name__ == '__main__':
 
    n = 121
    len = 16
    print(f'The binary representation of {n} is', toBinary(n, len))
    print(f'The binary representation of 623 is', toBinary(623, len))
    print(f'The binary representation of 198 is', toBinary(198, len))
    print(f'The binary representation of 5321 is', toBinary(5321, len))
    print(f'The binary representation of 52125 is', toBinary(52125, 32))
