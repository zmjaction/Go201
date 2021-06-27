/*
* @Author: ZhaoMingJun
* @Date:   2021/6/25 3:43 下午
 */
package main

import (
	"fmt"
	"homework/global"
	"homework/internal/model"
	"homework/setting"
)

func init() {
	var err error
	err = setupSetting()
	if err != nil {
		fmt.Println(err)
	}

	err = setupDBEngine()
	if err != nil {
		fmt.Println(err)
	}

}


func main() {
	fmt.Println(global.DatabaseSetting.Host)

}


func setupSetting() error{

	newSetting, err := setting.NewSetting()
	if err != nil {
		fmt.Println(err)
	}
	// 解析普通的yaml
	err = newSetting.ReadSection("Database", &global.DatabaseSetting) // 解析到结构体
	if err != nil {
		return err
	}
	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}

	return nil
}