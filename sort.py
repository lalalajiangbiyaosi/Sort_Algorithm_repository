import sys
def Insert_sort(A) :
    
    for j, val in enumerate(A) :
        
        i = j - 1
        while i >= 0 and A[i] > A[i+1] :
            A[i],A[i+1] = A[i+1],A[i]
            i = i - 1
        
    # print(A)
    return A
def Select_sort(A) :
    for i in range(len(A)):
        minVal = A[i]
        minValIndex = i
        j = i + 1
        while j < len(A) :
            if A[j] < minVal :
                minVal = A[j]
                minValIndex = j
            j = j + 1
        A[i],A[minValIndex] = A[minValIndex],A[i]
    print(A)
def Merge_sort(A) :
    if len(A) < 2 :
        return A
    divNum = int(len(A)/2)
    X = A[0:divNum]
    Y = A[divNum:]
    if len(X) == 1 :
        gredientA = X
    else:
        gredientA = Merge_sort(X)
    if len(Y) == 1 :
        gredientB = Y
    else:
        gredientB = Merge_sort(Y)
    result_slice = []
    length_of_all = len(gredientA) + len(gredientB)
    gredientA.append(sys.maxsize)
    gredientB.append(sys.maxsize)
    pointA,pointB = 0,0
    for i in range(length_of_all):
        if gredientA[pointA] < gredientB[pointB] :
            result_slice.append(gredientA[pointA])
            pointA = pointA + 1
        else :
            result_slice.append(gredientB[pointB])
            pointB = pointB + 1
    return result_slice
if __name__ == '__main__' :
    a = [31, 41, 59, 26, 41, 58, 90, 100, 29, 11, 25]
    print(a)
    Insert_sort(a)
    Select_sort(a)
    print(Merge_sort(a))