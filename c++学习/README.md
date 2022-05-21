# C++学习

## 函数

```c++
return_type function_name( parameter list )
{
   body of the function
}
```

## 字符串

数组实现：

```c++
char site[7] = {'R', 'U', 'N', 'O', 'O', 'B', '\0'};
char site[] = "RUNOOB";
```

## 类和对象

类定义是以关键字 **class** 开头，后跟类的名称。类的主体是包含在一对花括号中。类定义后必须跟着一个分号或一个声明列表。

类的对象的公共数据成员可以使用直接成员访问运算符 **.** 来访问。

```c++
#include <bits/stdc++.h>

using namespace std;

class Person {
public:
    string name;
    int age;

    int getAge() {
        return age;
    }
};

int main(int argc, char *argv[]) {
    Person person;
    person.age = 100;
    printf("%d",person.getAge());
}
```

### 类成员函数

成员函数可以定义在类定义内部，或者单独使用**范围解析运算符 ::** 来定义。在类定义中定义的成员函数把函数声明为**内联**的，即便没有使用 inline 标识符。

### 访问修饰符

- public：**公有**成员在程序中类的外部是可访问的。您可以不使用任何成员函数来设置和获取公有变量的值
- private：**私有**成员变量或函数在类的外部是不可访问的，甚至是不可查看的。只有类和友元函数可以访问私有成员。默认情况下，类的所有成员都是私有的（即不加修饰符的话）。
- protected：受保护成员在派生类（即子类）中是可访问的。

继承中的特点：

- **public 继承：**基类 public 成员，protected 成员，private 成员的访问属性在派生类中分别变成：public, protected, private
- **protected 继承：**基类 public 成员，protected 成员，private 成员的访问属性在派生类中分别变成：protected, protected, private
- **private 继承：**基类 public 成员，protected 成员，private 成员的访问属性在派生类中分别变成：private, private, private

### 构造函数

构造函数的名称与类的名称是完全相同的，并且不会返回任何类型，也不会返回 void。构造函数可用于为某些成员变量设置初始值。

例子：

```c++
#include <bits/stdc++.h>

using namespace std;

class Person {
public:
    string name;
    int age;
    Person(string name, int age){
        this->age = age;
        this->name = name;
    }

    int getAge() {
        return age;
    }
};

int main(int argc, char *argv[]) {
    Person person("feng",18);
    printf("%d",person.getAge());
}
```

### 析构函数

类的**析构函数**是类的一种特殊的成员函数，它会在每次删除所创建的对象时执行。

析构函数的名称与类的名称是完全相同的，只是在前面加了个波浪号（~）作为前缀，它不会返回任何值，也不能带有任何参数。析构函数有助于在跳出程序（比如关闭文件、释放内存等）前释放资源。

```c++
    ~Person(){
        printf("end");
    }
```

### 继承

一个类可以派生自多个类，这意味着，它可以从多个基类继承数据和函数。

```c++
#include <bits/stdc++.h>

using namespace std;

class Person {
public:
    string name;
    int age;
    Person(string name, int age){
        this->age = age;
        this->name = name;
    }
    int getAge() {
        return age;
    }
};

class Student: public Person{
public:
    int grade;
    Student(string name,int age,int grade): Person(name,age){
        this->grade = grade;
    }
};

int main(int argc, char *argv[]) {
    Student student("feng",18,1);
    printf("%d\n",student.getAge());
}
```

访问控制：

| 访问     | public | protected | private |
| :------- | :----- | :-------- | :------ |
| 同一个类 | yes    | yes       | yes     |
| 派生类   | yes    | yes       | no      |
| 外部的类 | yes    | no        | no      |

### 重载运算符和重载函数

在同一个作用域内，可以声明几个功能类似的同名函数，但是这些同名函数的形式参数（指参数的个数、类型或者顺序）必须不同。不能仅通过返回类型的不同来重载函数。



重载的运算符是带有特殊名称的函数，函数名是由关键字 operator 和其后要重载的运算符符号构成的。与其他函数一样，重载运算符有一个返回类型和一个参数列表。

```c++
#include <bits/stdc++.h>

using namespace std;

class Person {
public:
    string name;
    int age;
    Person(string name, int age){
        this->age = age;
        this->name = name;
    }
    int getAge() {
        return age;
    }
};

class Student: public Person{
public:
    int grade;
    Student(string name,int age,int grade): Person(name,age){
        this->grade = grade;
    }
    Student operator+(const Student& s){
        return Student(this->name,this->age+s.age,this->grade);
    }
};

int main(int argc, char *argv[]) {
    Student student1("feng",18,1);
    Student student2("feng",17,1);
    printf("%d\n",(student1+student2).age);
}
```

### 多态

虚函数是指一个类中你希望重载的成员函数，当你用一个基类指针或引用指向一个继承类对象的时候，你调用一个虚函数，实际调用的是继承类的版本。

虚函数最关键的特点是“动态联编”，它可以在运行时判断指针指向的对象，并自动调用相应的函数。

**虚函数** 是在基类中使用关键字 **virtual** 声明的函数。在派生类中重新定义基类中定义的虚函数时，会告诉编译器不要静态链接到该函数。

我们想要的是在程序中任意点可以根据所调用的对象类型来选择调用的函数，这种操作被称为**动态链接**，或**后期绑定**。



```c++
virtual int area() = 0;
```

`=0`告诉编译器，函数没有主体，因此这个函数就是纯虚函数。

### 接口

如果类中至少有一个函数被声明为纯虚函数，则这个类就是抽象类。纯虚函数是通过在声明中使用 "= 0" 来指定的。

抽象类不能被用于实例化对象，它只能作为**接口**使用。如果试图实例化一个抽象类的对象，会导致编译错误。

## STL

### vector

数组遵循静态方法，这意味着在运行时不能更改其大小，而vector实现动态数组意味着在添加元素时会自动调整其大小。



| 函数                                                         | 描述                                                    |
| :----------------------------------------------------------- | :------------------------------------------------------ |
| [at(idx)](https://www.nhooo.com/cpp/cpp-vector-at-function.html) | 传回索引idx所指的数据，如果idx越界，抛出out_of_range。  |
| [back()](https://www.nhooo.com/cpp/cpp-vector-back-function.html) | 返回最后一个原始，不检查这个数据是否存在。              |
| [front()](https://www.nhooo.com/cpp/cpp-vector-front-function.html) | 返回第一个元素。                                        |
| [swap()](https://www.nhooo.com/cpp/cpp-vector-swap-function.html) | 交换两个Vector。                                        |
| [push_back()](https://www.nhooo.com/cpp/cpp-vector-push-back-function.html) | 在Vector最后添加一个元素。                              |
| [pop_back()](https://www.nhooo.com/cpp/cpp-vector-pop-back-function.html) | 它从向量中删除最后一个元素。                            |
| [empty()](https://www.nhooo.com/cpp/cpp-vector-empty-function.html) | 判断Vector是否为空（返回true时为空）                    |
| [insert()](https://www.nhooo.com/cpp/cpp-vector-insert-function.html) | 它将在指定位置插入新元素。                              |
| [erase()](https://www.nhooo.com/cpp/cpp-vector-erase-function.html) | 删除指定的元素。                                        |
| [resize()](https://www.nhooo.com/cpp/cpp-vector-resize-function.html) | 它修改向量的大小。                                      |
| [clear()](https://www.nhooo.com/cpp/cpp-vector-clear-function.html) | 它从向量中删除所有元素。                                |
| [size()](https://www.nhooo.com/cpp/cpp-vector-size-function.html) | 返回Vector元素数量的大小。                              |
| [capacity()](https://www.nhooo.com/cpp/cpp-vector-capacity-function.html) | 返回vector所能容纳的元素数量(在不重新分配内存的情况下） |
| [assign()](https://www.nhooo.com/cpp/cpp-vector-assign-function.html) | 它将新值分配给向量。                                    |
| [operator=()](https://www.nhooo.com/cpp/cpp-vector-operator-equal-function.html) | 它将新值分配给向量容器。                                |
| [operator[]()](https://www.nhooo.com/cpp/cpp-vector-operator-select-function.html) | 它访问指定的元素。                                      |
| [end()](https://www.nhooo.com/cpp/cpp-vector-end-function.html) | 返回最末元素的迭代器(实指向最末元素的下一个位置)        |
| [emplace()](https://www.nhooo.com/cpp/cpp-vector-emplace-function.html) | 它将在位置pos之前插入一个新元素。                       |
| [emplace_back()](https://www.nhooo.com/cpp/cpp-vector-emplace-back-function.html) | 它在末尾插入一个新元素。                                |
| [rend()](https://www.nhooo.com/cpp/cpp-vector-rend-function.html) | 它指向向量的第一个元素之前的元素。                      |
| [rbegin()](https://www.nhooo.com/cpp/cpp-vector-rbegin-function.html) | 它指向向量的最后一个元素。                              |
| [begin()](https://www.nhooo.com/cpp/cpp-vector-begin-function.html) | 返回第一个元素的迭代器。                                |
| [max_size()](https://www.nhooo.com/cpp/cpp-vector-max-size-function.html) | 返回Vector所能容纳元素的最大数量（上限）。              |
| [cend()](https://www.nhooo.com/cpp/cpp-vector-cend-function.html) | 它指向量中的last-last-element。                         |
| [cbegin()](https://www.nhooo.com/cpp/cpp-vector-cbegin-function.html) | 它指向量的第一个元素。                                  |
| [crbegin()](https://www.nhooo.com/cpp/cpp-vector-crbegin-function.html) | 它指的是向量的最后一个字符。                            |
| [crend()](https://www.nhooo.com/cpp/cpp-vector-crend-function.html) | 它指的是向量的第一个元素之前的元素。                    |
| [data()](https://www.nhooo.com/cpp/cpp-vector-data-function.html) | 它将向量的数据写入数组。                                |
| [shrink_to_fit()](https://www.nhooo.com/cpp/cpp-vector-shrink-to-fit-function.html) | 它减少了容量并使它等于向量的大小。                      |



==初始化方式==：

1. `vector<int> ilist1;`
2. `vector<int> ilist2(ilist);`
3. `vector<int> ilist = {1,2,3.0,4,5,6,7};`
4. `vector<int> ilist3(ilist.begin()+2,ilist.end()-1);`
5. `vector<int> ilist4(7);`，7个值为0的int
6. `vector<int> ilist5(7,3);`，7个值为3的int。



==遍历==：

1. 下标
2. 迭代器
3. auto关键字



```c++
    vector<int> v1;
    v1.push_back(1);
    v1.push_back(2);
    for (int i = 0; i < v1.size(); ++i) {
        printf("%d ",v1[i]);
    }
    printf("\n");
    for(vector<int>::iterator iter = v1.begin();iter!=v1.end();iter++){
        printf("%d ",*iter);
    }
    printf("\n");
    for(auto iter = v1.begin();iter!=v1.end();iter++){
        printf("%d ",*iter);
    }
    printf("\n");
    for(auto i : v1){
        printf("%d ",i);
    }
```

### deque

Deque是双端队列。

==初始化==：

```c++
std::deque<int> dq;	//创建一个empty的int型队列
std::deque<int> dq(8);  //创建一个有8个元素的int型队列，默认初始化值(value)为0
std::deque<int> dq(8, 50);  //创建一个有8个元素的int型队列，默认初始化值(value)都设为50
std::deque<int> dq(dq.begin(), dq.end()); //通过迭代器创建队列
std::deque<int> dq1(dq);	//通过拷贝构造创建队列
```



==方法==：

```c++
push_back();  	//在队列末尾增加一个元素， 参数为拷贝或移动的元素
push_front();  	//在队列头部增加一个元素，参数可以是拷贝或移动的元素
emplace();		//在队列指定的元素位置前插入新的元素
emplace_back();	//在队列尾部增加新的元素
emplace_front();// 在队列头部增加新的元素
insert();		//在队列莫一元素前增加新的元素

pop_front()	//从队列头部移除第一个元素
pop_back()  //从队列尾部移除最末尾的元素
erase(); //从队列指定的元素位置删除元素，可以指定一个范围删除。
clear();	//清空队列所有元素，size将为0

at()		// 在队列中返回指定索引元素的引用
front()		// 在队列中返回头部第一个元素的引用
back()		//在队列中返回尾部最后一个元素的引用
size()		//返回队列中元素的个数
max_size()	//返回队列最大容量
resize()	// 重新扩展容器大小， 如果重新扩展的容量小于现在的容量，多出的将会丢失，如果大于现在的容量，将会在现在容量的基础进行扩充并以0填充默认值
```

其中`emplac_back/front`和`push_back/front`功能是一样的都是用来增加新元素到队列，前者在效率上要好一些。

### queue

先进先出的数据结构。

- front()：返回 queue 中第一个元素的引用。如果 queue 是常量，就返回一个常引用；如果 queue 为空，返回值是未定义的。
- back()：返回 queue 中最后一个元素的引用。如果 queue 是常量，就返回一个常引用；如果 queue 为空，返回值是未定义的。
- push(const T& obj)：在 queue 的尾部添加一个元素的副本。这是通过调用底层容器的成员函数 push_back() 来完成的。
- push(T&& obj)：以移动的方式在 queue 的尾部添加元素。这是通过调用底层容器的具有右值引用参数的成员函数 push_back() 来完成的。
- pop()：删除 queue 中的第一个元素。
- size()：返回 queue 中元素的个数。
- empty()：如果 queue 中没有元素的话，返回 true。
- emplace()：用传给 emplace() 的参数调用 T 的构造函数，在 queue 的尾部生成对象。
- swap(queue<T> &other_q)：将当前 queue 中的元素和参数 queue 中的元素交换。它们需要包含相同类型的元素。也可以调用全局函数模板 swap() 来完成同样的操作。