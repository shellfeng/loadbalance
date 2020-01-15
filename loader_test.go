package loadbalance

import (
	"fmt"
	"testing"
)

type Node struct {
	Name string
	Weight int
}

func (node *Node) GetWeight() int {
	return node.Weight
}

func TestWeightRound_Get(t *testing.T) {
	loader := WeightRoundLoader()

	loader.Add(&Node{
		Name:   "test01",
		Weight: 10,
	}, &Node{
		Name:   "test02",
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
