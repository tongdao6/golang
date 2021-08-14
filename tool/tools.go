package tool
import (
    // "sort"
    "fmt"
)
//统计 号码出现次数
func StaticCnt(s []string) map[string]int{
    ret:=make(map[string]int,33)
    for _,v := range s{
        _,ok := ret[v]
        if ok {
            ret[v]++
        }else{
            ret[v]=1
        }
    }
    return ret
}

//比较两个字符串相同的个数
func compAB(A[]string,B[]string)int{
    s:=0
    for _,a := range A{
        for _,b:= range B{
            if a==b {
                s++
            }
        }
    }
    return s
}

//判断字符串中是否有重复项
func isRepeat(s []string)bool{
    m := make(map[string]int,6)
    for _,v := range s {
        if _,ok:=m[v];ok{
            return true
        }else{
            m[v]=1
        }
    }
    return false
}



//实现排序的三个接口
func (h ArrIssueInf) Len() int{
    return len(h)
}

func (h ArrIssueInf) Less(i,j int) bool{
    return h[i].issue<h[j].issue
}

func (h ArrIssueInf) Swap(i,j int){
    h[i],h[j]=h[j],h[i]
}

//实现排序的三个接口
func (h ArrNumInf) Len() int{
    return len(h)
}

func (h ArrNumInf) Less(i,j int) bool{
    if h[i].Cnt==h[j].Cnt{
        return h[i].Num<h[j].Num
    }
    return h[i].Cnt>h[j].Cnt
}

func (h ArrNumInf) Swap(i,j int){
    h[i],h[j]=h[j],h[i]
}

// //实现排序的三个接口
// func (h StatArrayCodeInf) Len() int{
//     return len(h)
// }

// func (h StatArrayCodeInf) Less(i,j int) bool{
//     return h[i].cnt>h[j].cnt
// }

// func (h StatArrayCodeInf) Swap(i,j int){
//     h[i],h[j]=h[j],h[i]
// }
// //将map转成slice
// func Map2Slice(m map[string]int) arrstrint{
//     ret := make(arrstrint,0,len(m))
//     for key,val :=range m {
//         ret = append(ret,strint{
//             s:key,
//             cnt:val,
//         })
//     }
//     sort.Stable(ret)
//     return ret
// }

//生成红球基础号码
func Reds() [] string{
    ret:=make([]string,0,33)
    for i:=1;i<=33;i++{
        ret=append(ret,fmt.Sprintf("%02d",i))
    }
    return ret
}