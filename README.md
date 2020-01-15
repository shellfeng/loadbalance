## 说明
实现负载均衡算法,提供多种方案

### 加权轮训调度算法(WeightRoundLoader)
具体原理请百度,实现方式看源码，以下是示例:
```go
package main

// define node struct
type Node struct {
	Name string
	Weight int
}

func (node *Node) GetWeight() int {
	return node.Weight
}

func main () {
    // get WeightRoundLoader
    loader := WeightRoundLoader()
    
    loader.Add(&Node{
    		Name:   "test01",
    		Weight: 10,
    })
    loader.Add(&Node{
    		Name:   "test002",
    		Weight: 20,
    })
    
    var count1, count2 int
    for i:=0; i<1000;i++  {
    	selectNode := loader.Get().(*Node)
    	if selectNode.Name == "test01" {
    		count1++
    	}else {
    		count2++
    	}
    }
    
    fmt.Println(count1, count2)
    
}
```