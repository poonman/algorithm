package main
/*
标题：用两个栈实现一个队列
描述信息
用两个堆栈模拟队列的功能，实现 push，pop，count 三个方法
参考答案
简单的做法：栈 s1 和 s2，始终维护 s1 作为存储空间，以 s2 作为临时缓冲区，push 直接进 s1，pop 时 s1 导入 s2，栈顶出栈，导回 s1 优化做法：入队时，将元素压入 s1，出队时，判断 s2 是否为空，如不为空，则直接弹出顶元素；如为空，则将 s1 的元素逐个“倒入”s2，把最后一个元素弹出并出队

评分标准
2.5 分及以下：没有思路
3.0 分：正确的写出基本的算法，且代码风格优秀
3.5 分：写出优化的算法，且逻辑严谨，边界检查健全
4.0 分：
 */


