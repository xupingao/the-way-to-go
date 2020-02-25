package main
import (
	"fmt"
	"reflect"
	"time"
)
type OutType struct { // tags
	field1 bool   "This is field1's tag"
	Field2 string "This is field2's tag"
	Field3 int    "This is field3's tag"
	field4 *InType "This is field4's tag"
}
type InType struct {
	a int "this is tag of a"
	b int "this is tag of b"
}
func (o OutType)Function1(){
	fmt.Println("this is function1,it's unexported.")
}
func (o OutType)Function2(arg1 string,arg2 *int)(res1 int,res *string){
	fmt.Println("this is function2,it's exported.")
	return 0,nil
}
func (o *OutType)Function3(a,b int)(InType){
	fmt.Println("this is function2,it's exported.")
	return InType{a,b}
}
func main() {
	var obj interface{}=OutType{true,"hello",1,&InType{1,2}}
	var objPtr interface{}=&OutType{true,"hello",1,&InType{1,2}}

	/////////////////////////////////////////////////////////////////////////////////////////////////
	//1、Type类型：包含了一个结构的类型信息（各数据成员和方法信息），不包含实际值，接口中包含了大量类型信息读取方法
	fmt.Println(reflect.TypeOf(obj))   //main.OutType
	fmt.Println(reflect.TypeOf(objPtr)) //*main.OutType

	objType:=reflect.TypeOf(obj)
	fmt.Println(objType.Size())   //40
	fmt.Println(objType.Kind())   //struct    kind()总显示底层类型
	fmt.Println(objType.Name())   //OutType
	//用指针初始化Type实例
	objPtrType:=reflect.TypeOf(objPtr)
	fmt.Println(objPtrType.Size())  //8
	fmt.Println(objPtrType.Kind())  //ptr
	fmt.Println(objPtrType.Name())  //空值
	/////////////////////////////////////////////////////////////////////////////////////////////////
	//2、Value类型：包含了Type类型和对应的实际值，是reflect包中代表一个实例的类型
	//   ValueOf returns a new Value initialized to the concrete value stored in the interface
	//   源码注释：返回一个根据输入的接口中存储的值初始化的新Value类型，其包含的Type与输入接口相同
	//   可以理解为将输入的接口初始化为reflect包中的实例
	objPtrValue:=reflect.ValueOf(objPtr)
	objPtr.(*OutType).field1=false   //对指针指向属性的修改，会同时修改之前生成的Value对象
	fmt.Println(objPtr)       //&{false hello 1 0xc00000a0e0}
	fmt.Println(objPtrValue)   //&{false hello 1 0xc00000a0e0}
	objValue:=reflect.ValueOf(obj)    //用值对象初始化Value实例，只复制了一份值，相当于按值传递
	obj=OutType{false,"hello",1,&InType{1,2}}
	fmt.Println(obj)   //{false hello 1 0xc00000a140}
	fmt.Println(objValue)  //{true hello 1 0xc00000a0d0}
	//同样可以从Value实例中获得对应的Type信息
	typeFromobjValue:=objValue.Type()
	fmt.Println(typeFromobjValue.Name())  //OutType
	/////////////////////////////////////////////////////////////////////////////////////////////////
	//3、对于使用包含指针型值的接口初始化的Value和Type，无法直接获得指向类型的信息
	//   可以通过一下方式，提取valuel类型指针指向的值，到一个新的value类型中，此时就获得了非指针类型
	//   在输入不确定是否为指针型的情况下，可以使用这种方法确保操作的是非指针型实例
	fmt.Println(objPtrType.Name())   //空
	elem1:=reflect.ValueOf(objPtr).Elem()
    elem2:=reflect.Indirect(reflect.ValueOf(objPtr))
    fmt.Println(elem1.Type().Name())   //OutType
	fmt.Println(elem2.Type().Name())   //OutType
	fmt.Println(reflect.ValueOf(objPtr).Elem().Type().Name())   //OutType
	fmt.Println(reflect.Indirect(reflect.ValueOf(objPtr)).Type().Name())   //OutType
	/////////////////////////////////////////////////////////////////////////////////////////////////
	//4、同样的，reflect包 支持将非指针类型转化为指针类型进行处理
	ptrValue:=reflect.PtrTo(objType)
	fmt.Println(ptrValue.Kind())   //ptr
	fmt.Println(ptrValue)          //*main.OutType
	/////////////////////////////////////////////////////////////////////////////////////////////////
	//5、使用type类型：获得类型内各字段信息Field实例，其结构如下：
	//type StructField struct {
	//	Name string   字段的名称
	//	PkgPath string  标志着一个非导出（首字母小写）字段的包路径，如果是导出量，该值为空（因此可以用于判断是否为导出量）
	//	Type      Type      字段类型，与外部的的实例类型没什么不同，因此可以做域中数据的进一步判断
	//	Tag       StructTag  字段后标签的字符串
	//	Offset    uintptr   // offset within struct, in bytes
	//	Index     []int     // index sequence for Type.FieldByIndex
	//	Anonymous bool     是否是一个嵌入的（匿名的）域
	//}
	for i:=1;i<=objType.NumField();i++{
		field:=objType.Field(i-1)
		fType:=field.Type
		fName:=field.Name
		fPkgPath:=field.PkgPath
		fTag:=field.Tag
		fmt.Printf("Field%v  type:%v  name:%v  PkgPath:%v  tag:%v \n",i,fType,fName,fPkgPath,fTag)
		//Field1  type:bool  name:field1  PkgPath:main  tag:This is field1's tag
		//Field2  type:string  name:Field2  PkgPath:  tag:This is field2's tag
		//Field3  type:int  name:Field3  PkgPath:  tag:This is field3's tag
		//Field4  type:*main.InType  name:field4  PkgPath:main  tag:This is field4's tag
	}
	//得到Field的Type字段，就可以查看类型内类型的具体信息
	intype:=objType.Field(3).Type
	fmt.Println(intype)   //*main.InType
	inObj:=intype.Elem()
	fmt.Println(inObj.Field(0).Name)  //a
	//可以通过函数FieldByName(name string) (StructField, bool),用字段名称得到字段类型
	field,_:=objType.FieldByName("field1");
	fmt.Println(field.Tag)   //This is field1's tag
	/////////////////////////////////////////////////////////////////////////////////////////////////
	//6、使用Type类型：获得实例下的方法信息
	//type Method struct {
	//	Name    string  方法名
	//	PkgPath string   包路径名，同样可以被用来判断是否是被导出函数
	//	Type  Type  方法类型
	//	Func  Value 以接收者为第一参数的函数体
	//	Index int   方法下标
	//}
	fmt.Println(objType.NumMethod())   //2
	for i:=1;i<=objType.NumMethod();i++{
		mtd:=objType.Method(i-1)
		mName:=mtd.Name
		mType:=mtd.Type
		mPkgPath:=mtd.PkgPath
		fmt.Printf("Method%v  %v %v %v \n",i,mName,mType,mPkgPath)
		//Method1  Function1 func(main.OutType)
		//Method2  Function2 func(main.OutType, string, *int) (int, *string)
	}
	//与正常方法调用要求相同，以指针作为接收体的方法必需以指针实例调用
	//因此，只有指针化的结构才能反射到第三个函数
	fmt.Println(reflect.PtrTo(objType).Method(2).Name)    //Function3
	/////////////////////////////////////////////////////////////////////////////////////////////////
	//7、对方法类型的处理
	//   方法对应的类型同样实现了Type接口，可以获得其中的字段信息
	mtdType:=objType.Method(1).Type
	fmt.Println(mtdType)   //func(main.OutType, string, *int) (int, *string)
	fmt.Println(mtdType.Kind()) //func
	//同样提供了按方法名获得方法类型的函数
	fun,_:=objType.MethodByName("Function1")
	fmt.Println(fun.Type)//func(main.OutType)
	//在方法中最关键的当然是参数和返回值返回值信息，使用Out和In函数，同样可以得到对应的Type类型
	//值得注意的是，方法的接收者类型被作为第一个参数进行处理
	for i:=0;i<mtdType.NumIn();i++{
		argType:=mtdType.In(i)
		if argType.Kind()==reflect.Ptr{   //对于有可能是指针的量，可以与relect.Ptr进行对比做出判断
			fmt.Println(argType.Elem().Name())
		}else{
			fmt.Println(argType.Name())
		}
	}
	for i:=0;i<mtdType.NumOut();i++{
		resType:=mtdType.Out(i)
		if resType.Kind()==reflect.Ptr{
			fmt.Println(resType.Elem().Name())
		}else{
			fmt.Println(resType.Name())
		}
	}
	/////////////////////////////////////////////////////////////////////////////////////////////////
	//8、除解析一个接口的结构字段和方法外，还可以对注册在结构上的方法进行调用
	//   调用过程中需要注意接收者是否为指针型，被调用函数应当是被导出类型
	//   当然，这个调用过程只能在包含了结构实例值的Value类型上使用，Type类型无法嗲用
	//    func (v Value) Call(in []Value) []Value
	//   参数和返回值都是reflect包中Value型的切片，需要经过转换
	objValue.Method(0).Call(nil)   //函数打印：this is function1,it's unexported.
	res:=objPtrValue.Method(2).Call(
		[]reflect.Value{reflect.ValueOf(1),reflect.ValueOf(2)})//this is function2,it's exported.
	fmt.Println(res[0])  //{1 2}
	//对于Type类型包含的方法，因为没有实际值作为接收者，需要吧传入的第一个参数作为接收者
	function:=objType.Method(0).Func
	function.Call([]reflect.Value{reflect.ValueOf(obj)})//this is function1,it's unexported.
	/////////////////////////////////////////////////////////////////////////////////////////////////
	//9、reflect包提供了方法对Value实例中的字段值做出修改
	//   To modify a reflection object, the value must be settable.
	//   值得注意的是，根据官方描述，为了改变一个反射对象，其值必需是可修改的
	//   这里的可修改值得注意，我们知道，通过输入的接口初始化的Value实例
	//   返回一个根据输入的接口中存储的值初始化的新Value类型（跟按值传递参数一致）
	//   因此，对Value中的值做出修改并没有改变外部接口值，因此并不支持这样做
	//   objValue.Field(1).SetString("hello,world！")
	//   panic: reflect: reflect.flag.mustBeAssignable using unaddressable value
	//   可行的用法是使用Elem()函数将用指针量初始化的Value实例中实际存储的值（相当于引用传递，这里的Elem()相当于一个解引用）
	fmt.Println(objPtrValue.Elem().Field(1))   //hello
	objPtrValue.Elem().Field(1).SetString("hello,world！")
	fmt.Println(objPtrValue.Elem().Field(1))    //hello,world！
	//可以使用CanSet()函数判断是否能够修改值
	fmt.Println(objValue.Field(1).CanSet())   //false
	fmt.Println(objPtrValue.Elem().Field(1).CanSet())  //true
 	///////////////////////////////////////////////////////////////////////////////////////////////////
 	//11、对切片进行的一些操作
 	slice:=[]int{}
 	slice=append(slice,1,2,3,4,5,6,7,8,9)
 	sliceValue:=reflect.ValueOf(&slice).Elem()
	fmt.Println(sliceValue.Len())       //9
	fmt.Println(sliceValue.Cap())       //10
	sliceValue.SetLen(10)
	fmt.Println(sliceValue.Len())       //10
 	fmt.Println(sliceValue.Index(3)) //4
	///////////////////////////////////////////////////////////////////////////////////////////////////
	//11、补充
	//对于kind()返回的底层类型的判断举例
	fmt.Println(objPtrType.Kind()==reflect.Ptr)   //true
	fmt.Println(objType.Kind()==reflect.Struct)   //true
	var m=make(map[int]int)
	fmt.Println(reflect.TypeOf(m).Kind()==reflect.Map) //true
	//reflect包提供func New(typ Type) Value函数，根据输入类型默认初始化一个Value
	// 该Value量包含的是一个指针型实例，也就是New函数和包外的New一样，申请好了内存空间，其中值可修改
	newType:=reflect.New(objType)
	fmt.Println(newType)   //&{false  0 <nil>}
	fmt.Println(newType.Elem().Field(0).CanSet())  //true
	//可以使用Interface()函数将Value值转化回对应的interface
	fmt.Println(objValue.Interface().(OutType).Field2)//hello
	//对于管道值的操作,reflect提供了send函数
	ch:=make(chan int)
	go reflect.ValueOf(ch).Send(reflect.ValueOf(1))
	fmt.Println(<-ch)  //1
	time.Sleep(1e9)
}

