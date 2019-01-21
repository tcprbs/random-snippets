def decompress(s):
    result = ""
    temp = []
    n_times = ""
    i = 0

    while i < len(s):
        char = s[i]
        if ord(char) >= 48 and ord(char) <= 57:
            n_times += char
        elif char == '[':
            while s[i] != ']':
                i += 1
                if s[i] != ']':
                    temp.append(s[i])
            if temp:
               repeated_str = "".join(temp)
               result += (repeated_str * int(n_times))
               n_times = ""
               temp = []
        else:
            result += char
        i += 1
    return result
