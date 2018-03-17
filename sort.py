def Insert_sort(A) :
    print(A)
    for j, val in enumerate(A) :
        
        i = j - 1
        while i >= 0 and A[i] > A[i+1] :
            A[i],A[i+1] = A[i+1],A[i]
            i = i - 1
        
    print(A)
    return A

if __name__ == '__main__' :
    a = [31, 41, 59, 26, 41, 58, 90, 100, 29, 11, 25]
    Insert_sort(a)
        