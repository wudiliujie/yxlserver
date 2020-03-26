package consts

type AttrType int32

const (
	AttrType_Hp          AttrType = 1   //生命
	AttrType_Atk         AttrType = 2   //攻击
	AttrType_Def         AttrType = 3   //防御
	AttrType_Hit         AttrType = 4   //命中(万分比)
	AttrType_Avd         AttrType = 5   //闪避(万分比)
	AttrType_Cri         AttrType = 6   //暴击(万分比)
	AttrType_ACri        AttrType = 7   //防暴击(万分比)
	AttrType_Speed       AttrType = 8   //速度
	AttrType_Atk1        AttrType = 9   //无视防御
	AttrType_Def1        AttrType = 10  //减免无视
	AttrType_Atk2        AttrType = 11  //伤害加深
	AttrType_Def2        AttrType = 12  //伤害减少
	AttrType_Atk3        AttrType = 13  //伤害增加%
	AttrType_Def3        AttrType = 14  //伤害减少%
	AttrType_Atk4        AttrType = 15  //暴击伤害增加
	AttrType_Def4        AttrType = 16  //暴击伤害减少
	AttrType_Atk5        AttrType = 17  //魂概率增加
	AttrType_Def5        AttrType = 18  //魂概率减少
	AttrType_Atk6        AttrType = 19  //PVP伤害增加(万分比)
	AttrType_Def6        AttrType = 20  //PVP伤害减少(万分比)
	AttrType_Atk7        AttrType = 21  //PVE伤害增加(万分比)
	AttrType_Def7        AttrType = 22  //PVE伤害减少(万分比)
	AttrType_ExtraHurt   AttrType = 23  //额外伤害
	AttrType_Atk8        AttrType = 24  //PVE首领伤害增加万分比
	AttrType_Hp_P        AttrType = 51  //生命
	AttrType_Atk_P       AttrType = 52  //攻击
	AttrType_Def_P       AttrType = 53  //防御
	AttrType_Hit_P       AttrType = 54  //命中(万分比)
	AttrType_Avd_P       AttrType = 55  //闪避(万分比)
	AttrType_Cri_P       AttrType = 56  //暴击(万分比)
	AttrType_ACri_P      AttrType = 57  //防暴击(万分比)
	AttrType_Speed_P     AttrType = 58  //速度
	AttrType_Atk1_P      AttrType = 59  //无视防御
	AttrType_Def1_P      AttrType = 60  //减免无视
	AttrType_Atk2_P      AttrType = 61  //伤害加深
	AttrType_Def2_P      AttrType = 62  //伤害减少
	AttrType_Atk3_P      AttrType = 63  //伤害增加%
	AttrType_Def3_P      AttrType = 64  //伤害减少%
	AttrType_Atk4_P      AttrType = 65  //暴击伤害增加
	AttrType_Def4_P      AttrType = 66  //暴击伤害减少
	AttrType_Atk5_P      AttrType = 67  //魂概率增加
	AttrType_Def5_P      AttrType = 68  //魂概率减少
	AttrType_Atk6_P      AttrType = 69  //PVP伤害增加(万分比)
	AttrType_Def6_P      AttrType = 70  //PVP伤害减少(万分比)
	AttrType_Atk7_P      AttrType = 71  //PVE伤害增加(万分比)
	AttrType_Def7_P      AttrType = 72  //PVE伤害减少(万分比)
	AttrType_UserId      AttrType = 100 //用户编号,纯客户端用户
	AttrType_Coin1       AttrType = 101 //金币(银元宝)
	AttrType_Coin2       AttrType = 102 //绑定金币（没有用）
	AttrType_Coin3       AttrType = 103 //仙玉
	AttrType_Level       AttrType = 104 //等级
	AttrType_Exp         AttrType = 105 //经验
	AttrType_Nick        AttrType = 106 //昵称
	AttrType_Sex         AttrType = 107 //性别
	AttrType_LoginAreaId AttrType = 108 //登录区编号
	AttrType_LoginSign   AttrType = 109 //登录签名,断线重连用
)

var SendPlayerInt = []AttrType{
	AttrType_Hp,
	AttrType_Atk,
	AttrType_Def,
	AttrType_Hit,
	AttrType_Avd,
	AttrType_Cri,
	AttrType_ACri,
	AttrType_Speed,
	AttrType_Atk1,
	AttrType_Def1,
	AttrType_Atk2,
	AttrType_Def2,
	AttrType_Atk3,
	AttrType_Def3,
	AttrType_Atk4,
	AttrType_Def4,
	AttrType_Atk5,
	AttrType_Def5,
	AttrType_Atk6,
	AttrType_Def6,
	AttrType_Atk7,
	AttrType_Def7,
	AttrType_ExtraHurt,
	AttrType_Atk8,
	AttrType_Hp_P,
	AttrType_Atk_P,
	AttrType_Def_P,
	AttrType_Hit_P,
	AttrType_Avd_P,
	AttrType_Cri_P,
	AttrType_ACri_P,
	AttrType_Speed_P,
	AttrType_Atk1_P,
	AttrType_Def1_P,
	AttrType_Atk2_P,
	AttrType_Def2_P,
	AttrType_Atk3_P,
	AttrType_Def3_P,
	AttrType_Atk4_P,
	AttrType_Def4_P,
	AttrType_Atk5_P,
	AttrType_Def5_P,
	AttrType_Atk6_P,
	AttrType_Def6_P,
	AttrType_Atk7_P,
	AttrType_Def7_P,
	AttrType_UserId,
	AttrType_Coin1,
	AttrType_Coin2,
	AttrType_Coin3,
	AttrType_Level,
	AttrType_Exp,
	AttrType_Sex,
	AttrType_LoginAreaId,
}
var SendPlayerStr = []AttrType{
	AttrType_Nick,
}
var SendOtherInt = []AttrType{
	AttrType_Level,
	AttrType_Sex,
}
var SendOtherStr = []AttrType{
	AttrType_Nick,
}
var SaveInt = []AttrType{
	AttrType_Hp,
	AttrType_Atk,
	AttrType_Def,
	AttrType_Hit,
	AttrType_Avd,
	AttrType_Cri,
	AttrType_ACri,
	AttrType_Speed,
	AttrType_Atk1,
	AttrType_Def1,
	AttrType_Atk2,
	AttrType_Def2,
	AttrType_Atk3,
	AttrType_Def3,
	AttrType_Atk4,
	AttrType_Def4,
	AttrType_Atk5,
	AttrType_Def5,
	AttrType_Atk6,
	AttrType_Def6,
	AttrType_Atk7,
	AttrType_Def7,
	AttrType_ExtraHurt,
	AttrType_Atk8,
	AttrType_Hp_P,
	AttrType_Atk_P,
	AttrType_Def_P,
	AttrType_Hit_P,
	AttrType_Avd_P,
	AttrType_Cri_P,
	AttrType_ACri_P,
	AttrType_Speed_P,
	AttrType_Atk1_P,
	AttrType_Def1_P,
	AttrType_Atk2_P,
	AttrType_Def2_P,
	AttrType_Atk3_P,
	AttrType_Def3_P,
	AttrType_Atk4_P,
	AttrType_Def4_P,
	AttrType_Atk5_P,
	AttrType_Def5_P,
	AttrType_Atk6_P,
	AttrType_Def6_P,
	AttrType_Atk7_P,
	AttrType_Def7_P,
	AttrType_UserId,
	AttrType_Coin1,
	AttrType_Coin2,
	AttrType_Coin3,
	AttrType_Level,
	AttrType_Exp,
	AttrType_Sex,
	AttrType_LoginAreaId,
}
var SaveStr = []AttrType{
	AttrType_Nick,
}
var FightAttr = []AttrType{
	AttrType_Hp,
	AttrType_Atk,
	AttrType_Def,
	AttrType_Hit,
	AttrType_Avd,
	AttrType_Cri,
	AttrType_ACri,
	AttrType_Speed,
	AttrType_Atk1,
	AttrType_Def1,
	AttrType_Atk2,
	AttrType_Def2,
	AttrType_Atk3,
	AttrType_Def3,
	AttrType_Atk4,
	AttrType_Def4,
	AttrType_Atk5,
	AttrType_Def5,
	AttrType_Atk6,
	AttrType_Def6,
	AttrType_Atk7,
	AttrType_Def7,
	AttrType_ExtraHurt,
	AttrType_Atk8,
	AttrType_Hp_P,
	AttrType_Atk_P,
	AttrType_Def_P,
	AttrType_Hit_P,
	AttrType_Avd_P,
	AttrType_Cri_P,
	AttrType_ACri_P,
	AttrType_Speed_P,
	AttrType_Atk1_P,
	AttrType_Def1_P,
	AttrType_Atk2_P,
	AttrType_Def2_P,
	AttrType_Atk3_P,
	AttrType_Def3_P,
	AttrType_Atk4_P,
	AttrType_Def4_P,
	AttrType_Atk5_P,
	AttrType_Def5_P,
	AttrType_Atk6_P,
	AttrType_Def6_P,
	AttrType_Atk7_P,
	AttrType_Def7_P,
}
var ResetAttr = []AttrType{}
var StrResetAttr = []AttrType{}
var IntBaseAttr = []AttrType{
	AttrType_UserId,
	AttrType_Level,
	AttrType_Sex,
	AttrType_LoginAreaId,
}
var StrBaseAttr = []AttrType{
	AttrType_Nick,
}
var IntCorssAttr = []AttrType{}
var StrCrossAttr = []AttrType{
	AttrType_Nick,
}
var BaseFightAttr = []AttrType{
	AttrType_Hp,
	AttrType_Atk,
	AttrType_Def,
}
var IntRankAttr = []AttrType{
	AttrType_Level,
}
var CounterAttr = []AttrType{}
var AttrName = map[AttrType]string{
	AttrType_Hp:          "生命",
	AttrType_Atk:         "攻击",
	AttrType_Def:         "防御",
	AttrType_Hit:         "命中(万分比)",
	AttrType_Avd:         "闪避(万分比)",
	AttrType_Cri:         "暴击(万分比)",
	AttrType_ACri:        "防暴击(万分比)",
	AttrType_Speed:       "速度",
	AttrType_Atk1:        "无视防御",
	AttrType_Def1:        "减免无视",
	AttrType_Atk2:        "伤害加深",
	AttrType_Def2:        "伤害减少",
	AttrType_Atk3:        "伤害增加%",
	AttrType_Def3:        "伤害减少%",
	AttrType_Atk4:        "暴击伤害增加",
	AttrType_Def4:        "暴击伤害减少",
	AttrType_Atk5:        "魂概率增加",
	AttrType_Def5:        "魂概率减少",
	AttrType_Atk6:        "PVP伤害增加(万分比)",
	AttrType_Def6:        "PVP伤害减少(万分比)",
	AttrType_Atk7:        "PVE伤害增加(万分比)",
	AttrType_Def7:        "PVE伤害减少(万分比)",
	AttrType_ExtraHurt:   "额外伤害",
	AttrType_Atk8:        "PVE首领伤害增加万分比",
	AttrType_Hp_P:        "生命",
	AttrType_Atk_P:       "攻击",
	AttrType_Def_P:       "防御",
	AttrType_Hit_P:       "命中(万分比)",
	AttrType_Avd_P:       "闪避(万分比)",
	AttrType_Cri_P:       "暴击(万分比)",
	AttrType_ACri_P:      "防暴击(万分比)",
	AttrType_Speed_P:     "速度",
	AttrType_Atk1_P:      "无视防御",
	AttrType_Def1_P:      "减免无视",
	AttrType_Atk2_P:      "伤害加深",
	AttrType_Def2_P:      "伤害减少",
	AttrType_Atk3_P:      "伤害增加%",
	AttrType_Def3_P:      "伤害减少%",
	AttrType_Atk4_P:      "暴击伤害增加",
	AttrType_Def4_P:      "暴击伤害减少",
	AttrType_Atk5_P:      "魂概率增加",
	AttrType_Def5_P:      "魂概率减少",
	AttrType_Atk6_P:      "PVP伤害增加(万分比)",
	AttrType_Def6_P:      "PVP伤害减少(万分比)",
	AttrType_Atk7_P:      "PVE伤害增加(万分比)",
	AttrType_Def7_P:      "PVE伤害减少(万分比)",
	AttrType_UserId:      "用户编号,纯客户端用户",
	AttrType_Coin1:       "金币(银元宝)",
	AttrType_Coin2:       "绑定金币（没有用）",
	AttrType_Coin3:       "仙玉",
	AttrType_Level:       "等级",
	AttrType_Exp:         "经验",
	AttrType_Nick:        "昵称",
	AttrType_Sex:         "性别",
	AttrType_LoginAreaId: "登录区编号",
	AttrType_LoginSign:   "登录签名,断线重连用",
}
