package loadbalance

import (
	"errors"
	"github.com/ebar-go/loadbalance/utils"
)

// WeightRound 权重轮训算法
type WeightRound struct {
	count         int            // 节点数
	maxWeight     int            // 最大权值
	lastNodeIndex int            // 上一次的节点
	currentWeight int            // 当前权值
	gcd           int            // 公约数
	items []Item
}



// Init
func (loader *WeightRound) Init() {
	loader.lastNodeIndex = -1
	loader.currentWeight = 0

	maxWeight := 0

	var numbers []int
	for _, item := range loader.items {
		weight := item.GetWeight()
		numbers = append(numbers, weight)
		if weight > maxWeight {
			maxWeight = weight
		}
	}

	if len(numbers) > 1 {
		// 计算最大公约数
		loader.gcd = utils.NGcd(numbers, len(numbers))
	}else {
		loader.gcd = 1
	}

	loader.maxWeight = maxWeight
}

// Get
func (loader *WeightRound) Get() Item {
	if len(loader.items) == 0 {
		return nil
	}

	// 使用权重轮询调度算法计算当前节点
	for {
		// 计算权值
		loader.lastNodeIndex = (loader.lastNodeIndex + 1) % loader.count
		if loader.lastNodeIndex == 0 {
			loader.currentWeight = loader.currentWeight - loader.gcd
			if loader.currentWeight <= 0 {
				loader.currentWeight = loader.maxWeight
				if loader.currentWeight == 0 {
					return nil
				}
			}
		}

		// 判断权值，如果当前节点的权重大于当前权值，则返回该节点
		if loader.items[loader.lastNodeIndex].GetWeight() >= loader.currentWeight {
			return loader.items[loader.lastNodeIndex]
		}
	}
}

// Add
func (loader *WeightRound) Add(items ...Item) {
	if len(items) == 0 {
		panic(errors.New("empty arguments"))
	}

	for _, item := range items {
		loader.items = append(loader.items, items...)
		loader.maxWeight = utils.Max(loader.maxWeight, item.GetWeight())

		loader.count++
	}

	loader.Init()
}

// delete
func (loader *WeightRound) Delete(index int) {
	if index < 0 || index >= loader.count {
		panic(errors.New("not found"))
	}

	// 删除节点
	loader.items = append(loader.items[:index], loader.items[index+1:]...)
	loader.count--
	// 重新初始化权重参数
	loader.Init()

}
