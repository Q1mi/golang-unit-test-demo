package gomock_demo

import (
	"bou.ke/monkey"
	"github.com/golang/mock/gomock"
	"gomock_demo/mocks"
	"reflect"
	"testing"
)

// db_test.go

func TestGetFromDB(t *testing.T) {
	// 创建gomock控制器
	ctrl := gomock.NewController(t)
	// 断言期望的方法都被执行
	// Go1.14+,mockgen1.5.0+的单测中不再需要手动调用该方法
	defer ctrl.Finish()
	// 调用mockgen生成代码中的NewMockDB方法
	// 这里mocks是我们生成代码时指定的package名称
	m := mocks.NewMockDB(ctrl)
	// 打桩（stub）
	//// 当传入Get函数的参数为liwenzhou.com时返回1和nil
	//m.
	//	EXPECT().
	//	Get(gomock.Eq("liwenzhou.com")). // 参数
	//	Return(1, nil).                  // 返回值
	//	Times(1)                         // 调用次数
	//
	//// 调用GetFromDB函数时传入上面的mock对象m
	//if v := GetFromDB(m, "liwenzhou.com"); v != 1 {
	//	t.Fatal()
	//}
	//// 再次调用上方mock的Get方法时不满足调用次数为1的期望
	//if v := GetFromDB(m, "liwenzhou.com"); v != 1 {
	//	t.Fatal()
	//}

	//m.EXPECT().Get(gomock.Not("q1mi")).Return(10, nil)
	//m.EXPECT().Get(gomock.Any()).Return(20, nil)
	//m.EXPECT().Get(gomock.Any()).Do(func(key string) {
	//	t.Logf("input key is %v\n", key)
	//})
	//m.EXPECT().Get(gomock.Any()).DoAndReturn(func(key string)(int, error) {
	//	t.Logf("input key is %v\n", key)
	//	return 10, nil
	//})
	//m.EXPECT().Get(gomock.Nil()).Return(-1, nil)
	//m.EXPECT().
	//	Get(gomock.Eq("liwenzhou")).
	//	SetArg(0, "liwenzhou.com").
	//	Return(1, nil)

	// 指定顺序
	gomock.InOrder(
		m.EXPECT().Get("1"),
		m.EXPECT().Get("2"),
		m.EXPECT().Get("3"),
	)

	// 按顺序调用
	GetFromDB(m, "1")
	GetFromDB(m, "2")
	GetFromDB(m, "3")
}




func TestGetFromDB2(t *testing.T) {
	var db DB
	db = &MySQL{}
	monkey.PatchInstanceMethod(reflect.TypeOf(db), "Get", func(db DB, key string) (int, error) {
		return 100, nil
	})

	if v := GetFromDB(db, "liwenzhou.com"); v != 100 {
		t.Fatalf("expect 100, got %v\n", v)
	}
}
