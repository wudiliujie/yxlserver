package units

import (
	"leaf/util"
	"leaf/log"
)

type House struct {
	pos  	int32
	x    	int32
	y    	int32
	owner	int32
}
const MAX_ACTION_NUM =6
type Unit struct{
	id 			int32
	gold    	int32
	currPos		int32		//当前位置
	currCircle		int32		//当前圈
	actionNum 	int32		//行动次数
	actionNumTimer util.CTimer //剩余次数计时器
	state   int32  // 0正常,1破产
}
func CreateUnit(id int32) *Unit {
	m:=&Unit{}
	m.id= id
	m.currCircle=1
	m.gold=1000
	return  m
}

func (m *Unit)Start(){
	m.actionNumTimer.BeginTimer(2000,0) //10秒
}
//是否投掷
func (m *Unit)CanActions()bool{
	if m.state==1{
		return  false
	}
	if m.actionNum>0{
		return  true
	}else {
		return  false
	}
}
//diff 毫秒
func (m *Unit)Update(diff int32){
	if m.state==1{ //破产
		return
	}

	//增加投掷次数
	if m.actionNumTimer.CountingTimer(diff){
		m.AddActionNum(1)
	}
}
func (m *Unit) AddActionNum(value int32){
	m.actionNum+=value
	if m.actionNum>=MAX_ACTION_NUM{
		m.actionNumTimer.CleanUp()
	}else{
		if !m.actionNumTimer.IsRun(){
			m.actionNumTimer.BeginTimer(10000,0)
		}

	}
}
func (m*Unit)AddPos(value int32,max_house_num int32){
	m.currPos+=value
	for{
		if m.currPos> max_house_num{
			m.currPos= m.currPos-max_house_num
			m.currCircle++
		}
		if m.currPos<=max_house_num{
			break
		}
	}

}

func (m*Unit) GetPos()int32{
	return  m.currPos
}
func (m*Unit) ResetPos(){
	 m.currPos=0
}
func (m*Unit)GetGold()int32{
	return  m.gold
}
func (m*Unit)GetCircle()int32{
	return  m.currCircle
}
func (m*Unit)GetState()int32{
	return  m.state
}
func (m *Unit)AddGold(value int32){
	m.gold+=value
}
//破产
func (m *Unit)Bankruptcy(){
	log.Debug("破产:%v",*m)
	m.gold=0
	m.state=1
}
func (m*Unit) GetId()int32{
	return  m.id
}
