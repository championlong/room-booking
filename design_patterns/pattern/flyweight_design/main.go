package main

/*
享元模式
“享元”，顾名思义就是被共享的单元。享元模式的意图是复用对象，节省内存，前提是享元对象是不 可变对象。
当一个系统中存在大量重复对象的时候，如果这些重复的对象是不可变对象，我们就可以利用享 元模式将对象设计成享元，在内存中只保留一份实例，供多处代码引用。

享元模式跟缓存的区别：缓存，主要是为了提高访问效率，而非复用
享元模式跟单例的区别：享元模式中，一个类可以创建多个对象，每个对象被多处代 码引用共享
 */
func main() {

}