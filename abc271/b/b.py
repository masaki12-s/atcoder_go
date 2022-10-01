n,q = map(int,input().split(" "))
a = [[]for i in range(n)]
for i in range(n):
    la = list(map(int,input().split(" ")))
    a[i] = la[1:]

for i in range(q):
    s,t = map(int,input().split(" "))
    print(a[s-1][t-1])