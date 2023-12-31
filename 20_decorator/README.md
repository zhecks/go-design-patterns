# decorator

## 定义

动态地给一个对象添加一些额外的职责。就增加功能来说，装饰模式比生成子类更为灵活

![image-20230818181854099](https://cdn.jsdelivr.net/gh/zhecks/static_resources/images/202308181818130.png)

![image-20230818181907911](https://cdn.jsdelivr.net/gh/zhecks/static_resources/images/202308181819039.png)

## 优缺点

装饰模式有以下优点

* 比继承更灵活

从为对象添加功能的角度来看，装饰模式比继承更灵活。继承是静态的，而且一旦继承所有子类都有一样的功能。而装饰模式采用把功能分离到每个装饰器当中，然后通过对象组合的方式，在运行时动态地组合功能，每个被装饰的对象最终有哪些功能，是由运行期动态组合的功能来决定的

* 更容易复用功能

装饰模式把一些复杂的功能分散到每个装饰器当中，一般一个装饰器只实现一个功能，使实现装饰器变得简单，更重要的是这样有利于装饰器功能的复用，可以给一个对象增加多个同样的装饰器，也可以把一个装饰用来装饰不同的对象，从而实现复用装饰器的功能

* 简化高层定义

装饰模式可以通过组合装饰器的方式，为对象增添任意多的功能。因此在进行高层定义的时候，不用把所有的功能都定义出来，而是定义最基本的就可以了，可以在需要使用的时候，组合相应的装饰器来完成所需的功能

装饰模式的缺点

会产生很多细粒度对象（装饰器）

## 思考

1. 装饰模式的本质

装饰模式的本质：动态组合

2. 何时选用装饰模式

* 如果需要在不影响其他对象的情况下，以动态、透明的方式给对象添加职责，可以使用装饰模式，这几乎就是装饰模式的主要功能
* 如果不适合使用子类进行扩展的时候，可以考虑使用装饰模式