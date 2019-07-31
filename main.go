package main
import (
"fmt"
)
//把要默写的单词放入切片中
var slice1= [][]string{
{"clone","英 [kləʊn]  美 [klon] ","n. 克隆"},
{"origin","英 ['ɒrɪdʒɪn]  美 ['ɔrɪdʒɪn] ","n. 起源；原点；出身；开端"},
{"master","英 ['mɑːstə]  美 ['mæstɚ] ","原件"},
{"branch","英 [brɑːn(t)ʃ]  美 [bræntʃ] ","n. 分支；分公司；"},
{"develop","英 [dɪ'veləp]  美 [dɪ'vɛləp] ","vt. 开发；进步；使成长；使显影"},
{"ignore","英 [ɪg'nɔː]  美 [ɪɡ'nɔr] ","vt. 驳回诉讼；忽视；不理睬"},
{"name","英 [kləʊn]  美 [klon] ","n. 名字"},
}
var errs= [][]string{}
func wordOut(a int ){
//最后一个单词的索引
var len1 =len(slice1)-1
//最后一个单词默写后，提示信息
if a == len1{
//输出出错的单词
fmt.Println("不正确的单词如下:共",len(errs),"个")
for _,v:= range errs{
	fmt.Println(v[0],v[2])
}
fmt.Println("单词默写结束，程序结束")
}
}
func main(){
	var name string
	var word string
	fmt.Println("请输入你的名字")
	fmt.Scanln(&name)
	fmt.Println(name,":请根据提示的中文输入英文单词")
	//循环输出单词的中文与英文
	for i, v:= range slice1{
	fmt.Println(v[2])
	fmt.Scanln(&word)
		//第一次判断输入是否正解
		if word !=v[0]{
			fmt.Println("错误，请再次输入")
			fmt.Scanln(&word)
			//第二次判断输入是否正解
			if word !=v[0]{
				fmt.Println("错误,正确答案",v[0],"请根据提示的中文输入接下来的英文单词")
				//将第二次出错的单词放入另一个切片保存起来
				errs=append(errs,v)
				// fmt.Println(errs)
				//最后一个单词默写后，提示信息
				wordOut(i)
				}else{
				fmt.Println("正确")
				//最后一个单词默写后，提示信息
				wordOut(i)
			}
			}else{
			fmt.Println("正确")
			//最后一个单词默写后，提示信息
			wordOut(i)
		}
	}
}