## 二叉树中的最大路径和
### Description

给定一个非空二叉树，返回其最大路径和。

本题中，路径被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点。


Input
输入第一行为一个整数 n (1 <= n <= 100000), 表示二叉树的节点个数。

接下来共有 n 行, 每行有 3 个整数 l、r 和 val, 分别表示第 i (i为每一行的行号)个节点的左孩子和右孩子下标以及第 i 个节点的权值, 左/右孩子下标值为 0 表示当前节点没有左孩子/右孩子. 根节点的下标始终为 1.


Output
输出单独的一行, 即当前二叉树的最大路径和

### 解决方案
在 maxPathSumHelper 函数中：
- 我们先检查当前节点是否为 nil，如果是，则返回 0。
- 我们分别递归计算左子树和右子树的最大路径和，并将结果与 0 取较大值。
- 最后计算当前节点为根节点的最大路径和，即当前节点值加上左子树的最大路径和和右子树的最大路径和。我们将这个路径和与当前最大路径和 maxSum 进行比较，并更新 maxSum。