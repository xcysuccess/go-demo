package main

import "fmt"

func main()  {
	/*封装*/
	peo := new(People)
	peo.SetName("derek")
	peo.SetAge(22)
	fmt.Println(peo.GetName(),peo.GetAge())     //derek 22

	/*继承*/
	tea := Teacher{People{"derek-teacher",22},"911"}
	fmt.Println(tea.classroom,tea.name,tea.age)    //911 derek 22

	stu := Student{People{"derek-student",11},"919"}
	fmt.Println(stu.classroom,stu.name,stu.age)    //911 derek 22

	//tea.run()
	//stu.run()

	/*多态*/
	allrun(&tea)
	allrun(&stu)
}

type Live interface {
	run()
}
func allrun(live Live)  {
	live.run()
}

type People struct {
	name string
	age int
}

func (p *People) SetName(name string)  {
	p.name = name
}

func (p *People) SetAge(age int)  {
	p.age = age
}
func (p *People) run()  {
	fmt.Println("%v在跑步",p.name)
}
func (p *People) GetName() string{
	return p.name
}

func (p *People) GetAge() int{
	return p.age
}

type Teacher struct{
	People
	classroom string
}
func (tea *Teacher) run()  {
	fmt.Println(tea.name,"在跑步")
}
type Student struct{
	People
	classroom string
}
func (stu *Student) run()  {
	fmt.Println(stu.name,"在跑步")
}


