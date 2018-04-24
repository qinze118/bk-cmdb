package example

import (
	"configcenter/src/framework/api"
	"configcenter/src/framework/core/types"
	"fmt"
)

func init() {

	_, sender, _ := api.CreateCustomOutputer("example_output", func(data types.MapStr) error {
		fmt.Println("outputer:", data)
		return nil
	})

	api.RegisterInputer(target, sender, nil)
	api.RegisterTimingInputer()
}

var target = &myInputer{}

type myInputer struct {
}

// Description the Inputer description.
// This information will be printed when the Inputer is abnormal, which is convenient for debugging.
func (cli *myInputer) Name() string {
	return "name_myinputer"
}

// Input the input should not be blocked
func (cli *myInputer) Input() interface{} {

	fmt.Println("my_inputer")

	// 1. 返回 MapStr对象，此方法用于有Inputer绑定了自定义Outputer的时候使用，内置Outputer不采用此方法传递数据。
	/**
	return types.MapStr{
		"test": "outputer",
		"hoid": "",
	}
	*/

	// 此方法仅用于内置Outputer 的数据返回
	// 1. 构建模型分类
	// 2. 通过模型分类构建model
	// 3. 通过model 构建模型属性
	// 4. 利用包装器对要返回的数据做处理。

	//api.CreateBusiness()

	// TODO: inputer 内部显示调用 save
	// 多inputer 运行
	// inputer 定时跑  ，长期跑
	// 事务有cc 底层api 实现。（一时半会搞不了，不在框架里做）
	cls := api.CreateClassification() // 必填参数

	//cls.FindModelsByCondition()

	model := cls.CreateModel()

	grp := model.CreateGroup()

	grp.CreateAttribute()

	attr := model.CreateAttribute()
	attr.SetName("test")

	return nil

}