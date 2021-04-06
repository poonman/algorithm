package main

/*标题：基于有序单链接的 Set 部分函数实现
描述信息
题目详情
基于有序单链接的 Set 部分函数实现
class Node {
    int value;
    Node *next;
}
class Set {
    Node *head;
    public Set():head(NULL) {}
    ...
    bool remove(int v);
    Set& operator=(const Set &other);
}
考察对单链表的删除操作和对 C++operator=的理解
参考答案
remove 重点考察对链表和指针的操作
operator=考察对 C++的赋值操作符的理解，要考虑对象本身赋值的情况，还 fsdv 考虑释放左边对象的空间然后再 deepcopy*/
