package internal

import (
	"gameserver/units"
	"leaf/util"
	"leaf/log"
	"consts"
)


func (r *RoomData)Update(diff int32){
	if r.RoomState==RoomState_Wait{
		r.CheckStart()
	}else if  r.RoomState==RoomState_Gameing{
		count:=0
		for _,v:=range r.Units{
			v.Update(diff)
			if v.GetState()==0{
				count++
			}
			if v.CanActions(){ //需要投掷
				r.Throw(v,util.RandInterval(1,10))
			}
		}
		if r.debugTimer.CountingTimer(diff){
			log.Debug("剩余:%v",count)
		}
	}
}
func (r *RoomData) GetPlayerNum()int32{
	return  int32(len(r.Players))
}
//获取指定用户的所有房子的价值
func (r *RoomData)GetUnitAllHouseValue(unitId int32)int32{
	count:=int32(0)
	for i:=int32(1);i<=consts.MAX_HOUSE_NUM;i++{
		if r.Houses[i].unit!=nil && r.Houses[i].unit.GetId()==unitId{
			count+=r.Houses[i].GetPrice()
		}
	}
	return  count
}
//出售房子到指定的金币数
func (r* RoomData)SellHouseToGold(unit *units.Unit,gold int32){
	for i:=int32(1);i<=consts.MAX_HOUSE_NUM;i++{
		if r.Houses[i]!=nil && r.Houses[i].unit==unit{
			unit.AddGold(r.Houses[i].GetPrice())
			r.Houses[i].unit=nil
			r.Houses[i].issell=1
			if unit.GetGold()>=gold{
				break
			}
		}
	}
}
//出售价格小宇多少的金币
func (r* RoomData)SellHouseMinGold(unit *units.Unit,gold int32){
	for i:=int32(1);i<=consts.MAX_HOUSE_NUM;i++{
		if r.Houses[i]!=nil && r.Houses[i].unit==unit{
			if r.Houses[i].gold<=gold{
				unit.AddGold(r.Houses[i].GetPrice())
				r.Houses[i].unit=nil
				r.Houses[i].issell=1
			}
		}
	}
}
//检测是否开始
func (r *RoomData)CheckStart(){
	if r.GetPlayerNum()==1{
		//人数够4个人，开始假如机器人
		var i int32 =1
		for ;i<=99;i++{
			m:= units.CreateUnit(i)
			r.Units[i]=m
			m.Start()
		}
		//创建房间
		i=1
		for ;i<=consts.MAX_HOUSE_NUM;i++{
			r.Houses[i]=CreateHouse()
		}
		r.RoomState=RoomState_Gameing
		log.Debug("开始")
	}
}

//投掷：power: 力度,1-10
func (r *RoomData)Throw(u *units.Unit, power int32){
	u.AddActionNum(-1)
	rand:=util.RandInterval(1,power*10)
	//log.Debug("投掷:%v>>%v",power,rand)
	//出售的位置不能走
/*	for i:=int32(0);i<rand;i++{
		for {
			u.AddPos(1,r.currRoomNum)
			if r.Houses[u.GetPos()].issell==0{
				break
			}
		}
	}*/
	//出售价格小宇10

	//r.SellHouseMinGold(u,(u.GetCircle()-10)*100)
	u.AddPos(rand,consts.MAX_HOUSE_NUM)
	//判断当前位置有人占领没有
	h:=r.Houses[u.GetPos()]
	//log.Debug("u.getpos:%v",u.GetPos())
	if h.issell==1 {
		return
	}
	if h.unit==nil{ //没有占领
		//判断金币是否可以购买
		buyGold :=h.GetBuyGold(u.GetCircle())
		if  u.GetGold()>= buyGold{
			log.Debug("购买%v>>%v",*u,buyGold)
			u.AddGold(-buyGold)
			h.Buy(u.GetCircle())
			h.unit= u
		}
	}else if h.unit.GetId() ==u.GetId(){ //能升级不能
		levelGold :=h.GeteLevelGold(u.GetCircle())
		if u.GetGold()>=levelGold{
			log.Debug("升级%v>>%v",*u,levelGold)
			u.AddGold(-levelGold)
			h.LevelUp()
		}
	}else { //扣钱
		taxesGold:=h.GetTaxes()

		if u.GetGold()<taxesGold{
			//出售房子
			r.SellHouseToGold(u,taxesGold)
		}
		if u.GetGold()>=taxesGold{
			u.AddGold(-taxesGold)
			h.unit.AddGold(taxesGold)
			log.Debug("扣税:%v->%v>>%v>>circle:%v >>%v",u.GetId(),h.unit.GetId(),taxesGold,u.GetCircle(),h.unit.GetGold()+r.GetUnitAllHouseValue(h.unit.GetId()))
		}else{
			h.unit.AddGold(u.GetGold())
			u.Bankruptcy()
		}
	}
}