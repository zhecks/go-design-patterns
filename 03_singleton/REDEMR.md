# singleton

## 定义

保证一个类仅有一个实例，并提供一个访问它的全局访问点（减少内存消耗，共享变量）

![](https://cdn.jsdelivr.net/gh/zhecks/static_resources/images/202308161519049.png)

## 命名规范

一般建议单例模式的方法命名为getInstance()，这个方法的返回类型肯定是单例类的类型

## 懒汉式和饿汉式

懒汉式直到需要使用再创建实例，时间换空间

饿汉式不管有没有用到在类初始化的过程中就创建一份实例，空间还时间

## 思考

1. 单例模式的本质

单例模式的本质：控制实例数目（减少内存消耗，共享变量）

2. 何时选用单例模式

当需要控制一个类的实例职能有一个，而且客户职能从一个全局访问点访问它时，可以选用单例模式，这些功能恰好是单例模式要解决的问题