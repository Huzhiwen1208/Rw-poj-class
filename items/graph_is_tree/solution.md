## 判断无向图G是否是一棵树
### Description

给定一个无向图 G, 判断 G 是否为一棵树, 如果是, 输出 "YES", 否则输出 "NO".


Input
输入的第一行为一个正整数 n, 表示图中边的数量.

接下来 n 行输入, 每行为两个正整数 a, b, 表示节点 a 到节点 b 存在一条边.


Output
输出单独的一行, 如果给定无向图是一棵树, 输出 "YES", 否则输出 "NO".

### 解决方案
我们只需判断以下几个情况即可：
- 边数是否为N-1，不是则不是树
- 是否连通图？不是则不是树（只需要DFS解决）
- 以上均不返回False，就说明该无向图是树。