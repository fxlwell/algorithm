# design mode
```
	单例模式 sync.Once
	工厂模式 New一个变量
	策略模式	interface{}
	模板模式 interface{}
	代理模式 参数是func的方法


选项模式：
Option 模式
Option 模式的专业术语为：Functional Options Pattern（函数式选项模式）函数的参数传递是一个 Options。
为开发者提供了将一个函数的参数设置为可选的功能，每次新增选项时，可以不改变接口保持兼容，并且参数无顺序要求。


package main

import "log"

type Hero struct {
	HeroId  int
	Name    string // 名字
	SkillId int    // 技能Id
	Speed   int    // 移动速度
	Extra   map[string]string
}

type Option struct {
	Change func(hero *Hero)
}

func NewHeroTemplate() *Hero {
	return &Hero{
		HeroId:  1001,
		Name:    "Default",
		SkillId: 1,
		Speed:   1,
		Extra:   make(map[string]string),
	}
}

// 改变英雄名字
func WithName(name string) *Option {
	return &Option{Change: func(hero *Hero) {
		hero.Name = name
	}}
}

// 改变技能Id
func WithSkillId(skillId int) *Option {
	return &Option{Change: func(hero *Hero) {
		hero.SkillId = skillId
	}}
}

// 改变速度
func WithSpeed(speed int) *Option {
	return &Option{Change: func(hero *Hero) {
		hero.Speed = speed
	}}
}

// 增加扩展信息
func WithExtra(key string, value string) *Option {
	return &Option{Change: func(hero *Hero) {
		hero.Extra[key] = value
	}}
}

func GetHero(options ...*Option) *Hero {
	hero := NewHeroTemplate()
	for _, opt := range options {
		opt.Change(hero)
	}
	return hero
}

func main() {
	hero:=GetHero(
		WithName("libai"),
		WithSkillId(2),
		WithSkillId(3),
		WithSpeed(2),
		WithSpeed(10),
		WithExtra("name", "libai"))
	log.Println("英雄创造后属性为:",hero)
}


```
