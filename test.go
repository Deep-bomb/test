package main
import (
	"fmt"
	"strconv"
)

func main() {
	//var c1 byte='a';
// 	str :=`
// 	#include "draw.h"
// extern bool gameStart;
// #include <mmsystem.h>
// #pragma comment(lib,"winmm.lib")
// LRESULT CALLBACK MyProc(HWND hwnd, UINT msg, WPARAM wParam, LPARAM lParam)
// {
// 	switch(msg)
// 	{
// 	case WM_PAINT:
// 	{
// 		PAINTSTRUCT ps;
// 		HDC hdc = BeginPaint(hwnd, &ps);
// 		draw(hdc);
// 		EndPaint(hwnd, &ps); 
// 		break;
// 	}
// 	`
	var n int=99
	var isGirl bool=false
	var fMath float64=3.141
	
	//第一种方式，利用fmt包钟的Sprintf函数
	// var st3 string = fmt.Sprintf("%f",fMath)
	// var st1 string = fmt.Sprintf("%v",isGirl)
	// var str2 string = fmt.Sprintf("%v",n)


	//第二种方式，利用strconv包中的 FormatXXX函数
	var st1 string = strconv.FormatInt(int64(n),10)
	var st2 string = strconv.FormatFloat(fMath,'f',10,64)  //参数f 代表一种格式,10代表小数点后面的小数位
	var st3 string = strconv.FormatBool(isGirl)
	fmt.Printf("st1 type is %T str=%q\n",st1,st1)
	fmt.Printf("st2 type is %T str=%q\n",st2,st2)
	fmt.Printf("st3 type is %T str=%q\n",st3,st3)
	fmt.Println("this is the new line")
	//fmt.Println(str);	
}
