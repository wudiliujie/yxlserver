package msg

import (
	"github.com/golang/protobuf/proto"
	"github.com/wudiliujie/common/network"
	"github.com/wudiliujie/common/network/protobuf"
)

var Processor = protobuf.NewProcessor()

func init() {
	Processor.Register(PCK_S2SServerInfo_ID, func() network.IMessage { return new(S2SServerInfo) })
	Processor.Register(PCK_S2SRequestMsg_ID, func() network.IMessage { return new(S2SRequestMsg) })
	Processor.Register(PCK_S2SResponseMsg_ID, func() network.IMessage { return new(S2SResponseMsg) })
	Processor.Register(PCK_S2SServerPlayerPck_ID, func() network.IMessage { return new(S2SServerPlayerPck) })
	Processor.Register(PCK_S2SServerPlayerHandle_ID, func() network.IMessage { return new(S2SServerPlayerHandle) })
	Processor.Register(PCK_S2KPlayerData_ID, func() network.IMessage { return new(S2KPlayerData) })
	Processor.Register(PCK_Ping_ID, func() network.IMessage { return new(Ping) })
	Processor.Register(PCK_Pong_ID, func() network.IMessage { return new(Pong) })
	Processor.Register(PCK_S2KCrossPing_ID, func() network.IMessage { return new(S2KCrossPing) })
	Processor.Register(PCK_S2CServerTime_ID, func() network.IMessage { return new(S2CServerTime) })
	Processor.Register(PCK_S2CServerAge_ID, func() network.IMessage { return new(S2CServerAge) })
	Processor.Register(PCK_C2SChangeNick_ID, func() network.IMessage { return new(C2SChangeNick) })
	Processor.Register(PCK_S2CChangeNick_ID, func() network.IMessage { return new(S2CChangeNick) })
	Processor.Register(PCK_S2CServerInfo_ID, func() network.IMessage { return new(S2CServerInfo) })
	Processor.Register(PCK_C2SLogin_ID, func() network.IMessage { return new(C2SLogin) })
	Processor.Register(PCK_S2CLogin_ID, func() network.IMessage { return new(S2CLogin) })
	Processor.Register(PCK_C2SReLogin_ID, func() network.IMessage { return new(C2SReLogin) })
	Processor.Register(PCK_S2CReLogin_ID, func() network.IMessage { return new(S2CReLogin) })
	Processor.Register(PCK_C2SLoginEnd_ID, func() network.IMessage { return new(C2SLoginEnd) })
	Processor.Register(PCK_S2CAllData_ID, func() network.IMessage { return new(S2CAllData) })
	Processor.Register(PCK_C2SUserInfo_ID, func() network.IMessage { return new(C2SUserInfo) })
	Processor.Register(PCK_S2CUserInfo_ID, func() network.IMessage { return new(S2CUserInfo) })
	Processor.Register(PCK_S2COfflinePrize_ID, func() network.IMessage { return new(S2COfflinePrize) })
	Processor.Register(PCK_C2SGetOfflinePrize_ID, func() network.IMessage { return new(C2SGetOfflinePrize) })
	Processor.Register(PCK_S2CGetOfflinePrize_ID, func() network.IMessage { return new(S2CGetOfflinePrize) })
	Processor.Register(PCK_LoginFinish_ID, func() network.IMessage { return new(LoginFinish) })
	Processor.Register(PCK_C2SGetGift1_ID, func() network.IMessage { return new(C2SGetGift1) })
	Processor.Register(PCK_S2CGetGift1_ID, func() network.IMessage { return new(S2CGetGift1) })
	Processor.Register(PCK_C2SGetGift2_ID, func() network.IMessage { return new(C2SGetGift2) })
	Processor.Register(PCK_S2CGetGift2_ID, func() network.IMessage { return new(S2CGetGift2) })
	Processor.Register(PCK_C2SFocusPrice_ID, func() network.IMessage { return new(C2SFocusPrice) })
	Processor.Register(PCK_S2CFocusPrice_ID, func() network.IMessage { return new(S2CFocusPrice) })
	Processor.Register(PCK_C2SFirstInvite_ID, func() network.IMessage { return new(C2SFirstInvite) })
	Processor.Register(PCK_S2CFirstInvite_ID, func() network.IMessage { return new(S2CFirstInvite) })
	Processor.Register(PCK_C2SNewStory_ID, func() network.IMessage { return new(C2SNewStory) })
	Processor.Register(PCK_S2CNewStory_ID, func() network.IMessage { return new(S2CNewStory) })
	Processor.Register(PCK_S2CRoleInfo_ID, func() network.IMessage { return new(S2CRoleInfo) })
	Processor.Register(PCK_C2SOtherInfo_ID, func() network.IMessage { return new(C2SOtherInfo) })
	Processor.Register(PCK_S2COtherInfo_ID, func() network.IMessage { return new(S2COtherInfo) })
	Processor.Register(PCK_S2CRoleData_ID, func() network.IMessage { return new(S2CRoleData) })
	Processor.Register(PCK_C2SAchieveShow_ID, func() network.IMessage { return new(C2SAchieveShow) })
	Processor.Register(PCK_S2CAchieveShow_ID, func() network.IMessage { return new(S2CAchieveShow) })
	Processor.Register(PCK_S2CRoleBaseInfo_ID, func() network.IMessage { return new(S2CRoleBaseInfo) })
	Processor.Register(PCK_C2SRoleBaseInfo_ID, func() network.IMessage { return new(C2SRoleBaseInfo) })
	Processor.Register(PCK_S2CKill_ID, func() network.IMessage { return new(S2CKill) })
	Processor.Register(PCK_S2CBattlefieldReport_ID, func() network.IMessage { return new(S2CBattlefieldReport) })
	Processor.Register(PCK_S2KLoginSuccess_ID, func() network.IMessage { return new(S2KLoginSuccess) })
	Processor.Register(PCK_K2SBattlefieldReport_ID, func() network.IMessage { return new(K2SBattlefieldReport) })
	Processor.Register(PCK_K2SEnterMap_ID, func() network.IMessage { return new(K2SEnterMap) })
	Processor.Register(PCK_S2KEnterMap_ID, func() network.IMessage { return new(S2KEnterMap) })
	Processor.Register(PCK_S2KLeaveMap_ID, func() network.IMessage { return new(S2KLeaveMap) })
	Processor.Register(PCK_S2KLeaveInstanceTeam_ID, func() network.IMessage { return new(S2KLeaveInstanceTeam) })
	Processor.Register(PCK_S2KLogout_ID, func() network.IMessage { return new(S2KLogout) })
	Processor.Register(PCK_S2KReviveLife_ID, func() network.IMessage { return new(S2KReviveLife) })
	Processor.Register(PCK_S2KWorldLevel_ID, func() network.IMessage { return new(S2KWorldLevel) })
	Processor.Register(PCK_S2KCreateServerIds_ID, func() network.IMessage { return new(S2KCreateServerIds) })
	Processor.Register(PCK_K2SGetPlayerByLevel_ID, func() network.IMessage { return new(K2SGetPlayerByLevel) })
	Processor.Register(PCK_S2KGetPlayerOtherInfo_ID, func() network.IMessage { return new(S2KGetPlayerOtherInfo) })
	Processor.Register(PCK_K2SAddItem_ID, func() network.IMessage { return new(K2SAddItem) })
	Processor.Register(PCK_K2SAddRedPkg_ID, func() network.IMessage { return new(K2SAddRedPkg) })
	Processor.Register(PCK_K2SChangeAttr_ID, func() network.IMessage { return new(K2SChangeAttr) })
	Processor.Register(PCK_K2SUnlockSkin_ID, func() network.IMessage { return new(K2SUnlockSkin) })
	Processor.Register(PCK_K2SDeleteSkin_ID, func() network.IMessage { return new(K2SDeleteSkin) })
	Processor.Register(PCK_K2SMail_ID, func() network.IMessage { return new(K2SMail) })
	Processor.Register(PCK_K2SActState_ID, func() network.IMessage { return new(K2SActState) })
	Processor.Register(PCK_K2SBroadcastGang_ID, func() network.IMessage { return new(K2SBroadcastGang) })
	Processor.Register(PCK_C2SEndFight_ID, func() network.IMessage { return new(C2SEndFight) })
	Processor.Register(PCK_S2CPrizeReport_ID, func() network.IMessage { return new(S2CPrizeReport) })
	Processor.Register(PCK_C2SStageFight_ID, func() network.IMessage { return new(C2SStageFight) })
	Processor.Register(PCK_S2CStageFight_ID, func() network.IMessage { return new(S2CStageFight) })
	Processor.Register(PCK_C2SAutoStage_ID, func() network.IMessage { return new(C2SAutoStage) })
	Processor.Register(PCK_C2SStageSeek_ID, func() network.IMessage { return new(C2SStageSeek) })
	Processor.Register(PCK_S2CStageSeek_ID, func() network.IMessage { return new(S2CStageSeek) })
	Processor.Register(PCK_C2SStageHelp_ID, func() network.IMessage { return new(C2SStageHelp) })
	Processor.Register(PCK_S2CStageHelp_ID, func() network.IMessage { return new(S2CStageHelp) })
	Processor.Register(PCK_C2SChangeMap_ID, func() network.IMessage { return new(C2SChangeMap) })
	Processor.Register(PCK_S2CChangeMap_ID, func() network.IMessage { return new(S2CChangeMap) })
	Processor.Register(PCK_C2SAutoJump_ID, func() network.IMessage { return new(C2SAutoJump) })
	Processor.Register(PCK_C2SStagePrize_ID, func() network.IMessage { return new(C2SStagePrize) })
	Processor.Register(PCK_S2CStagePrize_ID, func() network.IMessage { return new(S2CStagePrize) })
	Processor.Register(PCK_C2SGetStagePrize_ID, func() network.IMessage { return new(C2SGetStagePrize) })
	Processor.Register(PCK_S2CGetStagePrize_ID, func() network.IMessage { return new(S2CGetStagePrize) })
	Processor.Register(PCK_C2SGetAllStagePrize_ID, func() network.IMessage { return new(C2SGetAllStagePrize) })
	Processor.Register(PCK_S2CGetAllStagePrize_ID, func() network.IMessage { return new(S2CGetAllStagePrize) })
	Processor.Register(PCK_C2SCatchPet_ID, func() network.IMessage { return new(C2SCatchPet) })
	Processor.Register(PCK_S2CCatchPet_ID, func() network.IMessage { return new(S2CCatchPet) })
	Processor.Register(PCK_C2SStartMove_ID, func() network.IMessage { return new(C2SStartMove) })
	Processor.Register(PCK_S2CStartMove_ID, func() network.IMessage { return new(S2CStartMove) })
	Processor.Register(PCK_S2CPlayerMove_ID, func() network.IMessage { return new(S2CPlayerMove) })
	Processor.Register(PCK_C2SStopMove_ID, func() network.IMessage { return new(C2SStopMove) })
	Processor.Register(PCK_S2CStopMove_ID, func() network.IMessage { return new(S2CStopMove) })
	Processor.Register(PCK_S2CPlayerStopMove_ID, func() network.IMessage { return new(S2CPlayerStopMove) })
	Processor.Register(PCK_S2CTransfer_ID, func() network.IMessage { return new(S2CTransfer) })
	Processor.Register(PCK_C2SCheckFight_ID, func() network.IMessage { return new(C2SCheckFight) })
	Processor.Register(PCK_S2CCheckFight_ID, func() network.IMessage { return new(S2CCheckFight) })
	Processor.Register(PCK_S2CPlayerEnterMap_ID, func() network.IMessage { return new(S2CPlayerEnterMap) })
	Processor.Register(PCK_S2CPlayerLeaveMap_ID, func() network.IMessage { return new(S2CPlayerLeaveMap) })
	Processor.Register(PCK_S2CMonsterEnterMap_ID, func() network.IMessage { return new(S2CMonsterEnterMap) })
	Processor.Register(PCK_S2CUpMonsterInfo_ID, func() network.IMessage { return new(S2CUpMonsterInfo) })
	Processor.Register(PCK_S2CMonsterLeaveMap_ID, func() network.IMessage { return new(S2CMonsterLeaveMap) })
	Processor.Register(PCK_C2SStartFight_ID, func() network.IMessage { return new(C2SStartFight) })
	Processor.Register(PCK_S2CStartFight_ID, func() network.IMessage { return new(S2CStartFight) })
	Processor.Register(PCK_C2SStartCollect_ID, func() network.IMessage { return new(C2SStartCollect) })
	Processor.Register(PCK_S2CStartCollect_ID, func() network.IMessage { return new(S2CStartCollect) })
	Processor.Register(PCK_C2SEndCollect_ID, func() network.IMessage { return new(C2SEndCollect) })
	Processor.Register(PCK_S2CEndCollect_ID, func() network.IMessage { return new(S2CEndCollect) })
	Processor.Register(PCK_S2CMonsterTalk_ID, func() network.IMessage { return new(S2CMonsterTalk) })
	Processor.Register(PCK_C2SRoleSkillUp_ID, func() network.IMessage { return new(C2SRoleSkillUp) })
	Processor.Register(PCK_C2SRoleSkillUpAuto_ID, func() network.IMessage { return new(C2SRoleSkillUpAuto) })
	Processor.Register(PCK_S2CRoleSkillUp_ID, func() network.IMessage { return new(S2CRoleSkillUp) })
	Processor.Register(PCK_C2SRoleSkillOrder_ID, func() network.IMessage { return new(C2SRoleSkillOrder) })
	Processor.Register(PCK_S2CRoleSkillOrder_ID, func() network.IMessage { return new(S2CRoleSkillOrder) })
	Processor.Register(PCK_C2SUnlockSkin_ID, func() network.IMessage { return new(C2SUnlockSkin) })
	Processor.Register(PCK_S2CUnlockSkin_ID, func() network.IMessage { return new(S2CUnlockSkin) })
	Processor.Register(PCK_C2SWearSkin_ID, func() network.IMessage { return new(C2SWearSkin) })
	Processor.Register(PCK_S2CWearSkin_ID, func() network.IMessage { return new(S2CWearSkin) })
	Processor.Register(PCK_C2SSkinUp_ID, func() network.IMessage { return new(C2SSkinUp) })
	Processor.Register(PCK_S2CSkinUp_ID, func() network.IMessage { return new(S2CSkinUp) })
	Processor.Register(PCK_C2SGetVipPrize_ID, func() network.IMessage { return new(C2SGetVipPrize) })
	Processor.Register(PCK_S2CGetVipPrize_ID, func() network.IMessage { return new(S2CGetVipPrize) })
	Processor.Register(PCK_C2SBuyVipPrize_ID, func() network.IMessage { return new(C2SBuyVipPrize) })
	Processor.Register(PCK_S2CBuyVipPrize_ID, func() network.IMessage { return new(S2CBuyVipPrize) })
	Processor.Register(PCK_C2SGetNewPrize_ID, func() network.IMessage { return new(C2SGetNewPrize) })
	Processor.Register(PCK_S2CGetNewPrize_ID, func() network.IMessage { return new(S2CGetNewPrize) })
	Processor.Register(PCK_C2SClientSwitch_ID, func() network.IMessage { return new(C2SClientSwitch) })
	Processor.Register(PCK_S2CClientSwitch_ID, func() network.IMessage { return new(S2CClientSwitch) })
	Processor.Register(PCK_C2SShareSuccess_ID, func() network.IMessage { return new(C2SShareSuccess) })
	Processor.Register(PCK_S2CUserGrade_ID, func() network.IMessage { return new(S2CUserGrade) })
	Processor.Register(PCK_C2SRoleLevelUp_ID, func() network.IMessage { return new(C2SRoleLevelUp) })
	Processor.Register(PCK_S2CRoleLevelUp_ID, func() network.IMessage { return new(S2CRoleLevelUp) })
	Processor.Register(PCK_C2SGradeUp_ID, func() network.IMessage { return new(C2SGradeUp) })
	Processor.Register(PCK_S2CGradeUp_ID, func() network.IMessage { return new(S2CGradeUp) })
	Processor.Register(PCK_C2SGrade2_ID, func() network.IMessage { return new(C2SGrade2) })
	Processor.Register(PCK_S2CGrade2_ID, func() network.IMessage { return new(S2CGrade2) })
	Processor.Register(PCK_C2SEquipGradeUp_ID, func() network.IMessage { return new(C2SEquipGradeUp) })
	Processor.Register(PCK_S2CEquipGradeUp_ID, func() network.IMessage { return new(S2CEquipGradeUp) })
	Processor.Register(PCK_C2SRoleGodGradeUp_ID, func() network.IMessage { return new(C2SRoleGodGradeUp) })
	Processor.Register(PCK_S2CRoleGodGradeUp_ID, func() network.IMessage { return new(S2CRoleGodGradeUp) })
	Processor.Register(PCK_C2SGetRoleGodPrize_ID, func() network.IMessage { return new(C2SGetRoleGodPrize) })
	Processor.Register(PCK_S2CGetRoleGodPrize_ID, func() network.IMessage { return new(S2CGetRoleGodPrize) })
	Processor.Register(PCK_S2CRoleGodPrize_ID, func() network.IMessage { return new(S2CRoleGodPrize) })
	Processor.Register(PCK_S2CGetFriends_ID, func() network.IMessage { return new(S2CGetFriends) })
	Processor.Register(PCK_C2SFocus_ID, func() network.IMessage { return new(C2SFocus) })
	Processor.Register(PCK_S2CFocus_ID, func() network.IMessage { return new(S2CFocus) })
	Processor.Register(PCK_C2SCancelFocus_ID, func() network.IMessage { return new(C2SCancelFocus) })
	Processor.Register(PCK_S2CCancelFocus_ID, func() network.IMessage { return new(S2CCancelFocus) })
	Processor.Register(PCK_C2SHate_ID, func() network.IMessage { return new(C2SHate) })
	Processor.Register(PCK_S2CHate_ID, func() network.IMessage { return new(S2CHate) })
	Processor.Register(PCK_C2SCancelHate_ID, func() network.IMessage { return new(C2SCancelHate) })
	Processor.Register(PCK_S2CCancelHate_ID, func() network.IMessage { return new(S2CCancelHate) })
	Processor.Register(PCK_C2SGiveCoin_ID, func() network.IMessage { return new(C2SGiveCoin) })
	Processor.Register(PCK_S2CGiveCoin_ID, func() network.IMessage { return new(S2CGiveCoin) })
	Processor.Register(PCK_C2SGetCoin_ID, func() network.IMessage { return new(C2SGetCoin) })
	Processor.Register(PCK_S2CGetCoin_ID, func() network.IMessage { return new(S2CGetCoin) })
	Processor.Register(PCK_C2SGetSuggest_ID, func() network.IMessage { return new(C2SGetSuggest) })
	Processor.Register(PCK_S2CGetSuggest_ID, func() network.IMessage { return new(S2CGetSuggest) })
	Processor.Register(PCK_C2SOneKeyGiveCoin_ID, func() network.IMessage { return new(C2SOneKeyGiveCoin) })
	Processor.Register(PCK_S2COneKeyGiveCoin_ID, func() network.IMessage { return new(S2COneKeyGiveCoin) })
	Processor.Register(PCK_C2SOneKeyGetCoin_ID, func() network.IMessage { return new(C2SOneKeyGetCoin) })
	Processor.Register(PCK_S2COneKeyGetCoin_ID, func() network.IMessage { return new(S2COneKeyGetCoin) })
	Processor.Register(PCK_C2SOneKeyFocus_ID, func() network.IMessage { return new(C2SOneKeyFocus) })
	Processor.Register(PCK_S2COneKeyFocus_ID, func() network.IMessage { return new(S2COneKeyFocus) })
	Processor.Register(PCK_C2SSendChatMsg_ID, func() network.IMessage { return new(C2SSendChatMsg) })
	Processor.Register(PCK_S2CSendChatMsg_ID, func() network.IMessage { return new(S2CSendChatMsg) })
	Processor.Register(PCK_C2SGetHistoryChat_ID, func() network.IMessage { return new(C2SGetHistoryChat) })
	Processor.Register(PCK_S2CGetHistoryChat_ID, func() network.IMessage { return new(S2CGetHistoryChat) })
	Processor.Register(PCK_S2CNewChatMsg_ID, func() network.IMessage { return new(S2CNewChatMsg) })
	Processor.Register(PCK_C2SWhisper_ID, func() network.IMessage { return new(C2SWhisper) })
	Processor.Register(PCK_S2CWhisper_ID, func() network.IMessage { return new(S2CWhisper) })
	Processor.Register(PCK_C2SGetWhisper_ID, func() network.IMessage { return new(C2SGetWhisper) })
	Processor.Register(PCK_S2CGetWhisper_ID, func() network.IMessage { return new(S2CGetWhisper) })
	Processor.Register(PCK_S2CAllUnreadWhisper_ID, func() network.IMessage { return new(S2CAllUnreadWhisper) })
	Processor.Register(PCK_C2SGetOnlineStatus_ID, func() network.IMessage { return new(C2SGetOnlineStatus) })
	Processor.Register(PCK_S2CGetOnlineStatus_ID, func() network.IMessage { return new(S2CGetOnlineStatus) })
	Processor.Register(PCK_C2SRemoveWhisper_ID, func() network.IMessage { return new(C2SRemoveWhisper) })
	Processor.Register(PCK_S2CRemoveWhisper_ID, func() network.IMessage { return new(S2CRemoveWhisper) })
	Processor.Register(PCK_S2CBagChange_ID, func() network.IMessage { return new(S2CBagChange) })
	Processor.Register(PCK_C2SUserBag_ID, func() network.IMessage { return new(C2SUserBag) })
	Processor.Register(PCK_S2CUserBag_ID, func() network.IMessage { return new(S2CUserBag) })
	Processor.Register(PCK_C2SExtendBag_ID, func() network.IMessage { return new(C2SExtendBag) })
	Processor.Register(PCK_S2CExtendBag_ID, func() network.IMessage { return new(S2CExtendBag) })
	Processor.Register(PCK_C2SExchange_ID, func() network.IMessage { return new(C2SExchange) })
	Processor.Register(PCK_S2CExchange_ID, func() network.IMessage { return new(S2CExchange) })
	Processor.Register(PCK_C2SGoldBag_ID, func() network.IMessage { return new(C2SGoldBag) })
	Processor.Register(PCK_S2CGoldBag_ID, func() network.IMessage { return new(S2CGoldBag) })
	Processor.Register(PCK_C2SGetGoldBagInfo_ID, func() network.IMessage { return new(C2SGetGoldBagInfo) })
	Processor.Register(PCK_S2CGetGoldBagInfo_ID, func() network.IMessage { return new(S2CGetGoldBagInfo) })
	Processor.Register(PCK_C2SWearEquip_ID, func() network.IMessage { return new(C2SWearEquip) })
	Processor.Register(PCK_S2CWearEquip_ID, func() network.IMessage { return new(S2CWearEquip) })
	Processor.Register(PCK_C2SWearOneEquip_ID, func() network.IMessage { return new(C2SWearOneEquip) })
	Processor.Register(PCK_S2CWearOneEquip_ID, func() network.IMessage { return new(S2CWearOneEquip) })
	Processor.Register(PCK_C2SEquipStar_ID, func() network.IMessage { return new(C2SEquipStar) })
	Processor.Register(PCK_S2CEquipStar_ID, func() network.IMessage { return new(S2CEquipStar) })
	Processor.Register(PCK_C2SMeltEquip_ID, func() network.IMessage { return new(C2SMeltEquip) })
	Processor.Register(PCK_S2CMeltEquip_ID, func() network.IMessage { return new(S2CMeltEquip) })
	Processor.Register(PCK_C2SOpenTreasure_ID, func() network.IMessage { return new(C2SOpenTreasure) })
	Processor.Register(PCK_S2COpenTreasure_ID, func() network.IMessage { return new(S2COpenTreasure) })
	Processor.Register(PCK_C2SAutoMelt_ID, func() network.IMessage { return new(C2SAutoMelt) })
	Processor.Register(PCK_S2CAutoMelt_ID, func() network.IMessage { return new(S2CAutoMelt) })
	Processor.Register(PCK_C2SMeltGoldByItem_ID, func() network.IMessage { return new(C2SMeltGoldByItem) })
	Processor.Register(PCK_S2CMeltGoldByItem_ID, func() network.IMessage { return new(S2CMeltGoldByItem) })
	Processor.Register(PCK_C2SGetBattlePrize_ID, func() network.IMessage { return new(C2SGetBattlePrize) })
	Processor.Register(PCK_ItemFly_ID, func() network.IMessage { return new(ItemFly) })
	Processor.Register(PCK_C2SGetShopMallList_ID, func() network.IMessage { return new(C2SGetShopMallList) })
	Processor.Register(PCK_S2CGetShopMallList_ID, func() network.IMessage { return new(S2CGetShopMallList) })
	Processor.Register(PCK_C2SShopBuy_ID, func() network.IMessage { return new(C2SShopBuy) })
	Processor.Register(PCK_S2CShopBuy_ID, func() network.IMessage { return new(S2CShopBuy) })
	Processor.Register(PCK_C2SOrgMallList_ID, func() network.IMessage { return new(C2SOrgMallList) })
	Processor.Register(PCK_S2COrgMallList_ID, func() network.IMessage { return new(S2COrgMallList) })
	Processor.Register(PCK_C2SOrgBuy_ID, func() network.IMessage { return new(C2SOrgBuy) })
	Processor.Register(PCK_S2COrgBuy_ID, func() network.IMessage { return new(S2COrgBuy) })
	Processor.Register(PCK_C2SGetGoodsLimit_ID, func() network.IMessage { return new(C2SGetGoodsLimit) })
	Processor.Register(PCK_S2CGetGoodsLimit_ID, func() network.IMessage { return new(S2CGetGoodsLimit) })
	Processor.Register(PCK_C2SGetSendShopList_ID, func() network.IMessage { return new(C2SGetSendShopList) })
	Processor.Register(PCK_S2KGetSendShopList_ID, func() network.IMessage { return new(S2KGetSendShopList) })
	Processor.Register(PCK_S2CGetSendShopList_ID, func() network.IMessage { return new(S2CGetSendShopList) })
	Processor.Register(PCK_C2SGetSendShopById_ID, func() network.IMessage { return new(C2SGetSendShopById) })
	Processor.Register(PCK_S2CGetSendShopById_ID, func() network.IMessage { return new(S2CGetSendShopById) })
	Processor.Register(PCK_C2SGetMySendShop_ID, func() network.IMessage { return new(C2SGetMySendShop) })
	Processor.Register(PCK_S2KGetMySendShopList_ID, func() network.IMessage { return new(S2KGetMySendShopList) })
	Processor.Register(PCK_S2CGetMySendShop_ID, func() network.IMessage { return new(S2CGetMySendShop) })
	Processor.Register(PCK_C2SSendShopOnSale_ID, func() network.IMessage { return new(C2SSendShopOnSale) })
	Processor.Register(PCK_S2KSendShopOnSale_ID, func() network.IMessage { return new(S2KSendShopOnSale) })
	Processor.Register(PCK_S2CSendShopOnSale_ID, func() network.IMessage { return new(S2CSendShopOnSale) })
	Processor.Register(PCK_C2SSendShopBuy_ID, func() network.IMessage { return new(C2SSendShopBuy) })
	Processor.Register(PCK_S2KSendShopPreBuy_ID, func() network.IMessage { return new(S2KSendShopPreBuy) })
	Processor.Register(PCK_K2SSendShopPreBuy_ID, func() network.IMessage { return new(K2SSendShopPreBuy) })
	Processor.Register(PCK_S2KSendShopBuy_ID, func() network.IMessage { return new(S2KSendShopBuy) })
	Processor.Register(PCK_K2SSendShopBuy_ID, func() network.IMessage { return new(K2SSendShopBuy) })
	Processor.Register(PCK_S2CSendShopBuy_ID, func() network.IMessage { return new(S2CSendShopBuy) })
	Processor.Register(PCK_C2SSendShopClear_ID, func() network.IMessage { return new(C2SSendShopClear) })
	Processor.Register(PCK_S2KSendShopClear_ID, func() network.IMessage { return new(S2KSendShopClear) })
	Processor.Register(PCK_K2SSendShopClear_ID, func() network.IMessage { return new(K2SSendShopClear) })
	Processor.Register(PCK_S2CSendShopClear_ID, func() network.IMessage { return new(S2CSendShopClear) })
	Processor.Register(PCK_C2SSaleLog_ID, func() network.IMessage { return new(C2SSaleLog) })
	Processor.Register(PCK_S2KGetSaleLog_ID, func() network.IMessage { return new(S2KGetSaleLog) })
	Processor.Register(PCK_S2CSaleLog_ID, func() network.IMessage { return new(S2CSaleLog) })
	Processor.Register(PCK_C2SMarkSendShop_ID, func() network.IMessage { return new(C2SMarkSendShop) })
	Processor.Register(PCK_S2CMarkSendShop_ID, func() network.IMessage { return new(S2CMarkSendShop) })
	Processor.Register(PCK_C2SGetSendCfg_ID, func() network.IMessage { return new(C2SGetSendCfg) })
	Processor.Register(PCK_S2CGetSendCfg_ID, func() network.IMessage { return new(S2CGetSendCfg) })
	Processor.Register(PCK_C2SGetSendUserInfo_ID, func() network.IMessage { return new(C2SGetSendUserInfo) })
	Processor.Register(PCK_S2CGetSendUserInfo_ID, func() network.IMessage { return new(S2CGetSendUserInfo) })
	Processor.Register(PCK_C2SSendItem_ID, func() network.IMessage { return new(C2SSendItem) })
	Processor.Register(PCK_S2CSendItem_ID, func() network.IMessage { return new(S2CSendItem) })
	Processor.Register(PCK_C2SGetTaskPrize_ID, func() network.IMessage { return new(C2SGetTaskPrize) })
	Processor.Register(PCK_S2CGetTaskPrize_ID, func() network.IMessage { return new(S2CGetTaskPrize) })
	Processor.Register(PCK_S2CHistoryTaskInfo_ID, func() network.IMessage { return new(S2CHistoryTaskInfo) })
	Processor.Register(PCK_C2SGetHistoryTaskPrize_ID, func() network.IMessage { return new(C2SGetHistoryTaskPrize) })
	Processor.Register(PCK_S2CGetHistoryTaskPrize_ID, func() network.IMessage { return new(S2CGetHistoryTaskPrize) })
	Processor.Register(PCK_S2CLastDayRemain_ID, func() network.IMessage { return new(S2CLastDayRemain) })
	Processor.Register(PCK_C2SLifeFind_ID, func() network.IMessage { return new(C2SLifeFind) })
	Processor.Register(PCK_S2CLifeFind_ID, func() network.IMessage { return new(S2CLifeFind) })
	Processor.Register(PCK_C2SLifeFastFind_ID, func() network.IMessage { return new(C2SLifeFastFind) })
	Processor.Register(PCK_S2CLifeFastFind_ID, func() network.IMessage { return new(S2CLifeFastFind) })
	Processor.Register(PCK_C2SWorldLevel_ID, func() network.IMessage { return new(C2SWorldLevel) })
	Processor.Register(PCK_C2SPeaceFinish_ID, func() network.IMessage { return new(C2SPeaceFinish) })
	Processor.Register(PCK_S2CPeaceFinish_ID, func() network.IMessage { return new(S2CPeaceFinish) })
	Processor.Register(PCK_C2SSignPrize_ID, func() network.IMessage { return new(C2SSignPrize) })
	Processor.Register(PCK_S2CSignPrize_ID, func() network.IMessage { return new(S2CSignPrize) })
	Processor.Register(PCK_S2CSMMonster_ID, func() network.IMessage { return new(S2CSMMonster) })
	Processor.Register(PCK_C2SSMFight_ID, func() network.IMessage { return new(C2SSMFight) })
	Processor.Register(PCK_S2CSMFight_ID, func() network.IMessage { return new(S2CSMFight) })
	Processor.Register(PCK_C2SSMRefreshStar_ID, func() network.IMessage { return new(C2SSMRefreshStar) })
	Processor.Register(PCK_S2CSMRefreshStar_ID, func() network.IMessage { return new(S2CSMRefreshStar) })
	Processor.Register(PCK_C2SSMFastFinish_ID, func() network.IMessage { return new(C2SSMFastFinish) })
	Processor.Register(PCK_S2CSMFastFinish_ID, func() network.IMessage { return new(S2CSMFastFinish) })
	Processor.Register(PCK_C2SMailList_ID, func() network.IMessage { return new(C2SMailList) })
	Processor.Register(PCK_S2CMailList_ID, func() network.IMessage { return new(S2CMailList) })
	Processor.Register(PCK_C2SGetMailAttach_ID, func() network.IMessage { return new(C2SGetMailAttach) })
	Processor.Register(PCK_S2CGetMailAttach_ID, func() network.IMessage { return new(S2CGetMailAttach) })
	Processor.Register(PCK_C2SDelMail_ID, func() network.IMessage { return new(C2SDelMail) })
	Processor.Register(PCK_S2CDelMail_ID, func() network.IMessage { return new(S2CDelMail) })
	Processor.Register(PCK_C2SGM_ID, func() network.IMessage { return new(C2SGM) })
	Processor.Register(PCK_S2CNewMail_ID, func() network.IMessage { return new(S2CNewMail) })
	Processor.Register(PCK_S2CBossPersonalInfo_ID, func() network.IMessage { return new(S2CBossPersonalInfo) })
	Processor.Register(PCK_C2SBossPersonalFight_ID, func() network.IMessage { return new(C2SBossPersonalFight) })
	Processor.Register(PCK_S2CBossPersonalFight_ID, func() network.IMessage { return new(S2CBossPersonalFight) })
	Processor.Register(PCK_C2SBossPersonalSweep_ID, func() network.IMessage { return new(C2SBossPersonalSweep) })
	Processor.Register(PCK_S2CBossPersonalSweep_ID, func() network.IMessage { return new(S2CBossPersonalSweep) })
	Processor.Register(PCK_S2CInstanceMaterialInfo_ID, func() network.IMessage { return new(S2CInstanceMaterialInfo) })
	Processor.Register(PCK_C2SInstanceMaterialFight_ID, func() network.IMessage { return new(C2SInstanceMaterialFight) })
	Processor.Register(PCK_S2CInstanceMaterialFight_ID, func() network.IMessage { return new(S2CInstanceMaterialFight) })
	Processor.Register(PCK_C2SInstanceMaterialSweep_ID, func() network.IMessage { return new(C2SInstanceMaterialSweep) })
	Processor.Register(PCK_S2CInstanceMaterialSweep_ID, func() network.IMessage { return new(S2CInstanceMaterialSweep) })
	Processor.Register(PCK_C2SInstanceMaterialBuy_ID, func() network.IMessage { return new(C2SInstanceMaterialBuy) })
	Processor.Register(PCK_S2CInstanceMaterialBuy_ID, func() network.IMessage { return new(S2CInstanceMaterialBuy) })
	Processor.Register(PCK_C2SInstanceBuBuShengLianInfo_ID, func() network.IMessage { return new(C2SInstanceBuBuShengLianInfo) })
	Processor.Register(PCK_S2CInstanceBuBuShengLianInfo_ID, func() network.IMessage { return new(S2CInstanceBuBuShengLianInfo) })
	Processor.Register(PCK_C2SInstanceBuBuShengLianFight_ID, func() network.IMessage { return new(C2SInstanceBuBuShengLianFight) })
	Processor.Register(PCK_S2CInstanceBuBuShengLianFight_ID, func() network.IMessage { return new(S2CInstanceBuBuShengLianFight) })
	Processor.Register(PCK_C2SInstanceBuBuShengLianJinLianZi_ID, func() network.IMessage { return new(C2SInstanceBuBuShengLianJinLianZi) })
	Processor.Register(PCK_S2CInstanceBuBuShengLianJinLianZi_ID, func() network.IMessage { return new(S2CInstanceBuBuShengLianJinLianZi) })
	Processor.Register(PCK_S2CInstanceTreasureInfo_ID, func() network.IMessage { return new(S2CInstanceTreasureInfo) })
	Processor.Register(PCK_C2SInstanceTreasureFight_ID, func() network.IMessage { return new(C2SInstanceTreasureFight) })
	Processor.Register(PCK_S2CInstanceTreasureFight_ID, func() network.IMessage { return new(S2CInstanceTreasureFight) })
	Processor.Register(PCK_C2SInstanceTreasureSweep_ID, func() network.IMessage { return new(C2SInstanceTreasureSweep) })
	Processor.Register(PCK_S2CInstanceTreasureSweep_ID, func() network.IMessage { return new(S2CInstanceTreasureSweep) })
	Processor.Register(PCK_C2SGetInstanceTreasureBox_ID, func() network.IMessage { return new(C2SGetInstanceTreasureBox) })
	Processor.Register(PCK_S2CGetInstanceTreasureBox_ID, func() network.IMessage { return new(S2CGetInstanceTreasureBox) })
	Processor.Register(PCK_S2CInstanceHeavenlyInfo_ID, func() network.IMessage { return new(S2CInstanceHeavenlyInfo) })
	Processor.Register(PCK_C2SInstanceHeavenlyFight_ID, func() network.IMessage { return new(C2SInstanceHeavenlyFight) })
	Processor.Register(PCK_S2CInstanceHeavenlyFight_ID, func() network.IMessage { return new(S2CInstanceHeavenlyFight) })
	Processor.Register(PCK_C2SInstanceHeavenlySweep_ID, func() network.IMessage { return new(C2SInstanceHeavenlySweep) })
	Processor.Register(PCK_S2CInstanceHeavenlySweep_ID, func() network.IMessage { return new(S2CInstanceHeavenlySweep) })
	Processor.Register(PCK_C2SGetInstanceHeavenlyBox_ID, func() network.IMessage { return new(C2SGetInstanceHeavenlyBox) })
	Processor.Register(PCK_S2CGetInstanceHeavenlyBox_ID, func() network.IMessage { return new(S2CGetInstanceHeavenlyBox) })
	Processor.Register(PCK_S2CInstanceDemonInfo_ID, func() network.IMessage { return new(S2CInstanceDemonInfo) })
	Processor.Register(PCK_C2SInstanceDemonFight_ID, func() network.IMessage { return new(C2SInstanceDemonFight) })
	Processor.Register(PCK_S2CInstanceDemonFight_ID, func() network.IMessage { return new(S2CInstanceDemonFight) })
	Processor.Register(PCK_C2SInstanceDemonSweep_ID, func() network.IMessage { return new(C2SInstanceDemonSweep) })
	Processor.Register(PCK_S2CInstanceDemonSweep_ID, func() network.IMessage { return new(S2CInstanceDemonSweep) })
	Processor.Register(PCK_C2SGetInstanceDemonBox_ID, func() network.IMessage { return new(C2SGetInstanceDemonBox) })
	Processor.Register(PCK_S2CGetInstanceDemonBox_ID, func() network.IMessage { return new(S2CGetInstanceDemonBox) })
	Processor.Register(PCK_S2CInstanceTowerInfo_ID, func() network.IMessage { return new(S2CInstanceTowerInfo) })
	Processor.Register(PCK_C2SInstanceTowerFight_ID, func() network.IMessage { return new(C2SInstanceTowerFight) })
	Processor.Register(PCK_S2CInstanceTowerFight_ID, func() network.IMessage { return new(S2CInstanceTowerFight) })
	Processor.Register(PCK_S2CAllBossInfo_ID, func() network.IMessage { return new(S2CAllBossInfo) })
	Processor.Register(PCK_C2SAllBossStart_ID, func() network.IMessage { return new(C2SAllBossStart) })
	Processor.Register(PCK_S2CAllBossStart_ID, func() network.IMessage { return new(S2CAllBossStart) })
	Processor.Register(PCK_C2SAllBossCfg_ID, func() network.IMessage { return new(C2SAllBossCfg) })
	Processor.Register(PCK_C2SAllBossRelive_ID, func() network.IMessage { return new(C2SAllBossRelive) })
	Processor.Register(PCK_S2CAllBossRelive_ID, func() network.IMessage { return new(S2CAllBossRelive) })
	Processor.Register(PCK_C2SAllBossGetBuyInfo_ID, func() network.IMessage { return new(C2SAllBossGetBuyInfo) })
	Processor.Register(PCK_S2CAllBossGetBuyInfo_ID, func() network.IMessage { return new(S2CAllBossGetBuyInfo) })
	Processor.Register(PCK_C2SAllBossBuyTimes_ID, func() network.IMessage { return new(C2SAllBossBuyTimes) })
	Processor.Register(PCK_S2CAllBossBuyTimes_ID, func() network.IMessage { return new(S2CAllBossBuyTimes) })
	Processor.Register(PCK_C2SAllBossGetDamageLog_ID, func() network.IMessage { return new(C2SAllBossGetDamageLog) })
	Processor.Register(PCK_S2CAllBossGetDamageLog_ID, func() network.IMessage { return new(S2CAllBossGetDamageLog) })
	Processor.Register(PCK_C2SAllBossGetKillLog_ID, func() network.IMessage { return new(C2SAllBossGetKillLog) })
	Processor.Register(PCK_S2CAllBossGetKillLog_ID, func() network.IMessage { return new(S2CAllBossGetKillLog) })
	Processor.Register(PCK_S2CAllBossV2Info_ID, func() network.IMessage { return new(S2CAllBossV2Info) })
	Processor.Register(PCK_C2SAllBossV2PlayerInBoss_ID, func() network.IMessage { return new(C2SAllBossV2PlayerInBoss) })
	Processor.Register(PCK_S2CAllBossV2PlayerInBoss_ID, func() network.IMessage { return new(S2CAllBossV2PlayerInBoss) })
	Processor.Register(PCK_C2SAllBossV2Start_ID, func() network.IMessage { return new(C2SAllBossV2Start) })
	Processor.Register(PCK_S2CAllBossV2Start_ID, func() network.IMessage { return new(S2CAllBossV2Start) })
	Processor.Register(PCK_C2SAllBossV2Cfg_ID, func() network.IMessage { return new(C2SAllBossV2Cfg) })
	Processor.Register(PCK_C2SAllBossV2GetBuyInfo_ID, func() network.IMessage { return new(C2SAllBossV2GetBuyInfo) })
	Processor.Register(PCK_S2CAllBossV2GetBuyInfo_ID, func() network.IMessage { return new(S2CAllBossV2GetBuyInfo) })
	Processor.Register(PCK_C2SAllBossV2BuyTimes_ID, func() network.IMessage { return new(C2SAllBossV2BuyTimes) })
	Processor.Register(PCK_S2CAllBossV2BuyTimes_ID, func() network.IMessage { return new(S2CAllBossV2BuyTimes) })
	Processor.Register(PCK_C2SAllBossV2GetDamageLog_ID, func() network.IMessage { return new(C2SAllBossV2GetDamageLog) })
	Processor.Register(PCK_S2CAllBossV2GetDamageLog_ID, func() network.IMessage { return new(S2CAllBossV2GetDamageLog) })
	Processor.Register(PCK_K2SFightAllBossResult_ID, func() network.IMessage { return new(K2SFightAllBossResult) })
	Processor.Register(PCK_S2CFieldBossV2Info_ID, func() network.IMessage { return new(S2CFieldBossV2Info) })
	Processor.Register(PCK_C2SFieldBossV2PlayerInBoss_ID, func() network.IMessage { return new(C2SFieldBossV2PlayerInBoss) })
	Processor.Register(PCK_S2CFieldBossV2PlayerInBoss_ID, func() network.IMessage { return new(S2CFieldBossV2PlayerInBoss) })
	Processor.Register(PCK_C2SFieldBossV2StartFight_ID, func() network.IMessage { return new(C2SFieldBossV2StartFight) })
	Processor.Register(PCK_S2CFieldBossV2StartFight_ID, func() network.IMessage { return new(S2CFieldBossV2StartFight) })
	Processor.Register(PCK_C2SFieldBossV2Cfg_ID, func() network.IMessage { return new(C2SFieldBossV2Cfg) })
	Processor.Register(PCK_C2SFieldBossV2BuyTimes_ID, func() network.IMessage { return new(C2SFieldBossV2BuyTimes) })
	Processor.Register(PCK_S2CFieldBossV2BuyTimes_ID, func() network.IMessage { return new(S2CFieldBossV2BuyTimes) })
	Processor.Register(PCK_C2SFieldBossV2GetDamageLog_ID, func() network.IMessage { return new(C2SFieldBossV2GetDamageLog) })
	Processor.Register(PCK_S2CFieldBossV2GetDamageLog_ID, func() network.IMessage { return new(S2CFieldBossV2GetDamageLog) })
	Processor.Register(PCK_S2CFieldBossInfo_ID, func() network.IMessage { return new(S2CFieldBossInfo) })
	Processor.Register(PCK_C2SFieldBossStart_ID, func() network.IMessage { return new(C2SFieldBossStart) })
	Processor.Register(PCK_S2CFieldBossStart_ID, func() network.IMessage { return new(S2CFieldBossStart) })
	Processor.Register(PCK_C2SFieldBossItem_ID, func() network.IMessage { return new(C2SFieldBossItem) })
	Processor.Register(PCK_S2CFieldBossItem_ID, func() network.IMessage { return new(S2CFieldBossItem) })
	Processor.Register(PCK_S2CBossVipInfo_ID, func() network.IMessage { return new(S2CBossVipInfo) })
	Processor.Register(PCK_C2SBossVipStart_ID, func() network.IMessage { return new(C2SBossVipStart) })
	Processor.Register(PCK_S2CBossVipStart_ID, func() network.IMessage { return new(S2CBossVipStart) })
	Processor.Register(PCK_C2SGetPartner_ID, func() network.IMessage { return new(C2SGetPartner) })
	Processor.Register(PCK_S2CGetPartner_ID, func() network.IMessage { return new(S2CGetPartner) })
	Processor.Register(PCK_C2SPartnerDeify_ID, func() network.IMessage { return new(C2SPartnerDeify) })
	Processor.Register(PCK_S2CPartnerDeify_ID, func() network.IMessage { return new(S2CPartnerDeify) })
	Processor.Register(PCK_C2SPartnerBuySkillHole_ID, func() network.IMessage { return new(C2SPartnerBuySkillHole) })
	Processor.Register(PCK_S2CPartnerBuySkillHole_ID, func() network.IMessage { return new(S2CPartnerBuySkillHole) })
	Processor.Register(PCK_C2SPartnerLearnSkill_ID, func() network.IMessage { return new(C2SPartnerLearnSkill) })
	Processor.Register(PCK_S2CPartnerLearnSkill_ID, func() network.IMessage { return new(S2CPartnerLearnSkill) })
	Processor.Register(PCK_C2SPartnerChange_ID, func() network.IMessage { return new(C2SPartnerChange) })
	Processor.Register(PCK_S2CPartnerChange_ID, func() network.IMessage { return new(S2CPartnerChange) })
	Processor.Register(PCK_C2SPartnerWearEquip_ID, func() network.IMessage { return new(C2SPartnerWearEquip) })
	Processor.Register(PCK_S2CPartnerWearEquip_ID, func() network.IMessage { return new(S2CPartnerWearEquip) })
	Processor.Register(PCK_C2SPartnerStrongEquip_ID, func() network.IMessage { return new(C2SPartnerStrongEquip) })
	Processor.Register(PCK_S2CPartnerStrongEquip_ID, func() network.IMessage { return new(S2CPartnerStrongEquip) })
	Processor.Register(PCK_C2SPartnerRefineEquip_ID, func() network.IMessage { return new(C2SPartnerRefineEquip) })
	Processor.Register(PCK_S2CPartnerRefineEquip_ID, func() network.IMessage { return new(S2CPartnerRefineEquip) })
	Processor.Register(PCK_C2SPartnerComposeEquip_ID, func() network.IMessage { return new(C2SPartnerComposeEquip) })
	Processor.Register(PCK_S2CPartnerComposeEquip_ID, func() network.IMessage { return new(S2CPartnerComposeEquip) })
	Processor.Register(PCK_C2SPartnerTakeOffEquip_ID, func() network.IMessage { return new(C2SPartnerTakeOffEquip) })
	Processor.Register(PCK_S2CPartnerTakeOffEquip_ID, func() network.IMessage { return new(S2CPartnerTakeOffEquip) })
	Processor.Register(PCK_C2SPartnerLevelUp_ID, func() network.IMessage { return new(C2SPartnerLevelUp) })
	Processor.Register(PCK_S2CPartnerLevelUp_ID, func() network.IMessage { return new(S2CPartnerLevelUp) })
	Processor.Register(PCK_C2SPartnerStarUp_ID, func() network.IMessage { return new(C2SPartnerStarUp) })
	Processor.Register(PCK_S2CPartnerStarUp_ID, func() network.IMessage { return new(S2CPartnerStarUp) })
	Processor.Register(PCK_C2SPartnerGradeUp_ID, func() network.IMessage { return new(C2SPartnerGradeUp) })
	Processor.Register(PCK_S2CPartnerGradeUp_ID, func() network.IMessage { return new(S2CPartnerGradeUp) })
	Processor.Register(PCK_C2SPartnerBattlePos_ID, func() network.IMessage { return new(C2SPartnerBattlePos) })
	Processor.Register(PCK_S2CPartnerBattlePos_ID, func() network.IMessage { return new(S2CPartnerBattlePos) })
	Processor.Register(PCK_C2SPartnerBattlePosRobot_ID, func() network.IMessage { return new(C2SPartnerBattlePosRobot) })
	Processor.Register(PCK_C2SPartnerRefreshSkill_ID, func() network.IMessage { return new(C2SPartnerRefreshSkill) })
	Processor.Register(PCK_S2CPartnerRefreshSkill_ID, func() network.IMessage { return new(S2CPartnerRefreshSkill) })
	Processor.Register(PCK_C2SPartnerDelSkill_ID, func() network.IMessage { return new(C2SPartnerDelSkill) })
	Processor.Register(PCK_S2CPartnerDelSkill_ID, func() network.IMessage { return new(S2CPartnerDelSkill) })
	Processor.Register(PCK_C2SPartnerUpSkill_ID, func() network.IMessage { return new(C2SPartnerUpSkill) })
	Processor.Register(PCK_S2CPartnerUpSkill_ID, func() network.IMessage { return new(S2CPartnerUpSkill) })
	Processor.Register(PCK_C2SPartnerReplaceSkill_ID, func() network.IMessage { return new(C2SPartnerReplaceSkill) })
	Processor.Register(PCK_S2CPartnerReplaceSkill_ID, func() network.IMessage { return new(S2CPartnerReplaceSkill) })
	Processor.Register(PCK_C2SPartnerNick_ID, func() network.IMessage { return new(C2SPartnerNick) })
	Processor.Register(PCK_S2CPartnerNick_ID, func() network.IMessage { return new(S2CPartnerNick) })
	Processor.Register(PCK_C2SPartnerSuit_ID, func() network.IMessage { return new(C2SPartnerSuit) })
	Processor.Register(PCK_S2CPartnerSuit_ID, func() network.IMessage { return new(S2CPartnerSuit) })
	Processor.Register(PCK_C2SPartnerSuitRobot_ID, func() network.IMessage { return new(C2SPartnerSuitRobot) })
	Processor.Register(PCK_C2SUnlockNewPartner_ID, func() network.IMessage { return new(C2SUnlockNewPartner) })
	Processor.Register(PCK_S2CUnlockNewPartner_ID, func() network.IMessage { return new(S2CUnlockNewPartner) })
	Processor.Register(PCK_C2SStickPartner_ID, func() network.IMessage { return new(C2SStickPartner) })
	Processor.Register(PCK_S2CStickPartner_ID, func() network.IMessage { return new(S2CStickPartner) })
	Processor.Register(PCK_S2CPartnerNewSkill_ID, func() network.IMessage { return new(S2CPartnerNewSkill) })
	Processor.Register(PCK_S2CUserPet_ID, func() network.IMessage { return new(S2CUserPet) })
	Processor.Register(PCK_S2CUserPetA_ID, func() network.IMessage { return new(S2CUserPetA) })
	Processor.Register(PCK_S2CUserDevil_ID, func() network.IMessage { return new(S2CUserDevil) })
	Processor.Register(PCK_C2SDevilAwake_ID, func() network.IMessage { return new(C2SDevilAwake) })
	Processor.Register(PCK_S2CDevilAwake_ID, func() network.IMessage { return new(S2CDevilAwake) })
	Processor.Register(PCK_C2SDevilLevelUp_ID, func() network.IMessage { return new(C2SDevilLevelUp) })
	Processor.Register(PCK_S2CDevilLevelUp_ID, func() network.IMessage { return new(S2CDevilLevelUp) })
	Processor.Register(PCK_C2SDevilPos_ID, func() network.IMessage { return new(C2SDevilPos) })
	Processor.Register(PCK_S2CDevilPos_ID, func() network.IMessage { return new(S2CDevilPos) })
	Processor.Register(PCK_C2SDevilSoulPos_ID, func() network.IMessage { return new(C2SDevilSoulPos) })
	Processor.Register(PCK_S2CDevilSoulPos_ID, func() network.IMessage { return new(S2CDevilSoulPos) })
	Processor.Register(PCK_C2SDevilSoulDel_ID, func() network.IMessage { return new(C2SDevilSoulDel) })
	Processor.Register(PCK_S2CDevilSoulDel_ID, func() network.IMessage { return new(S2CDevilSoulDel) })
	Processor.Register(PCK_C2SDevilSoulAuto_ID, func() network.IMessage { return new(C2SDevilSoulAuto) })
	Processor.Register(PCK_S2CDevilSoulAuto_ID, func() network.IMessage { return new(S2CDevilSoulAuto) })
	Processor.Register(PCK_C2SDevilSoulLevelUp_ID, func() network.IMessage { return new(C2SDevilSoulLevelUp) })
	Processor.Register(PCK_S2CDevilSoulLevelUp_ID, func() network.IMessage { return new(S2CDevilSoulLevelUp) })
	Processor.Register(PCK_C2SDevilSoulStarUp_ID, func() network.IMessage { return new(C2SDevilSoulStarUp) })
	Processor.Register(PCK_S2CDevilSoulStarUp_ID, func() network.IMessage { return new(S2CDevilSoulStarUp) })
	Processor.Register(PCK_Beauty_ID, func() network.IMessage { return new(Beauty) })
	Processor.Register(PCK_C2SBeautyNewCloth_ID, func() network.IMessage { return new(C2SBeautyNewCloth) })
	Processor.Register(PCK_S2CBeautyNewCloth_ID, func() network.IMessage { return new(S2CBeautyNewCloth) })
	Processor.Register(PCK_C2SBeautyWearCloth_ID, func() network.IMessage { return new(C2SBeautyWearCloth) })
	Processor.Register(PCK_S2CBeautyWearCloth_ID, func() network.IMessage { return new(S2CBeautyWearCloth) })
	Processor.Register(PCK_C2SBeautyClothUp_ID, func() network.IMessage { return new(C2SBeautyClothUp) })
	Processor.Register(PCK_S2CBeautyClothUp_ID, func() network.IMessage { return new(S2CBeautyClothUp) })
	Processor.Register(PCK_C2SBeautyFlySwitch_ID, func() network.IMessage { return new(C2SBeautyFlySwitch) })
	Processor.Register(PCK_S2CBeautyFlySwitch_ID, func() network.IMessage { return new(S2CBeautyFlySwitch) })
	Processor.Register(PCK_C2SBeautyEquip_ID, func() network.IMessage { return new(C2SBeautyEquip) })
	Processor.Register(PCK_S2CBeautyEquip_ID, func() network.IMessage { return new(S2CBeautyEquip) })
	Processor.Register(PCK_C2SBeautyMagicRefresh_ID, func() network.IMessage { return new(C2SBeautyMagicRefresh) })
	Processor.Register(PCK_S2CBeautyMagicRefresh_ID, func() network.IMessage { return new(S2CBeautyMagicRefresh) })
	Processor.Register(PCK_C2SBeautyMagicChange_ID, func() network.IMessage { return new(C2SBeautyMagicChange) })
	Processor.Register(PCK_S2CBeautyMagicChange_ID, func() network.IMessage { return new(S2CBeautyMagicChange) })
	Processor.Register(PCK_S2CUserKid_ID, func() network.IMessage { return new(S2CUserKid) })
	Processor.Register(PCK_S2CUserToy_ID, func() network.IMessage { return new(S2CUserToy) })
	Processor.Register(PCK_C2SKidTalentUp_ID, func() network.IMessage { return new(C2SKidTalentUp) })
	Processor.Register(PCK_S2CKidTalentUp_ID, func() network.IMessage { return new(S2CKidTalentUp) })
	Processor.Register(PCK_C2SKidUseTalent_ID, func() network.IMessage { return new(C2SKidUseTalent) })
	Processor.Register(PCK_S2CKidUseTalent_ID, func() network.IMessage { return new(S2CKidUseTalent) })
	Processor.Register(PCK_C2SKidGoodUp_ID, func() network.IMessage { return new(C2SKidGoodUp) })
	Processor.Register(PCK_S2CKidGoodUp_ID, func() network.IMessage { return new(S2CKidGoodUp) })
	Processor.Register(PCK_C2SToyLevelUp_ID, func() network.IMessage { return new(C2SToyLevelUp) })
	Processor.Register(PCK_S2CToyLevelUp_ID, func() network.IMessage { return new(S2CToyLevelUp) })
	Processor.Register(PCK_C2SToyOneKeyLevelUp_ID, func() network.IMessage { return new(C2SToyOneKeyLevelUp) })
	Processor.Register(PCK_S2CToyOneKeyLevelUp_ID, func() network.IMessage { return new(S2CToyOneKeyLevelUp) })
	Processor.Register(PCK_S2CUserTherion_ID, func() network.IMessage { return new(S2CUserTherion) })
	Processor.Register(PCK_S2CMasterMessage_ID, func() network.IMessage { return new(S2CMasterMessage) })
	Processor.Register(PCK_C2SSendMasterNotice_ID, func() network.IMessage { return new(C2SSendMasterNotice) })
	Processor.Register(PCK_S2CSendMasterNotice_ID, func() network.IMessage { return new(S2CSendMasterNotice) })
	Processor.Register(PCK_S2CMasterAds_ID, func() network.IMessage { return new(S2CMasterAds) })
	Processor.Register(PCK_C2SMasterInvite_ID, func() network.IMessage { return new(C2SMasterInvite) })
	Processor.Register(PCK_S2CMasterInvite_ID, func() network.IMessage { return new(S2CMasterInvite) })
	Processor.Register(PCK_S2CInvitePupil_ID, func() network.IMessage { return new(S2CInvitePupil) })
	Processor.Register(PCK_C2SHandleMasterInvite_ID, func() network.IMessage { return new(C2SHandleMasterInvite) })
	Processor.Register(PCK_S2CHandleMasterInvite_ID, func() network.IMessage { return new(S2CHandleMasterInvite) })
	Processor.Register(PCK_C2SMasterGiveExp_ID, func() network.IMessage { return new(C2SMasterGiveExp) })
	Processor.Register(PCK_S2CMasterGiveExp_ID, func() network.IMessage { return new(S2CMasterGiveExp) })
	Processor.Register(PCK_C2SGetMasterExp_ID, func() network.IMessage { return new(C2SGetMasterExp) })
	Processor.Register(PCK_S2CGetMasterExp_ID, func() network.IMessage { return new(S2CGetMasterExp) })
	Processor.Register(PCK_C2SDeletePupil_ID, func() network.IMessage { return new(C2SDeletePupil) })
	Processor.Register(PCK_S2CDeletePupil_ID, func() network.IMessage { return new(S2CDeletePupil) })
	Processor.Register(PCK_C2SDeleteMaster_ID, func() network.IMessage { return new(C2SDeleteMaster) })
	Processor.Register(PCK_S2CDeleteMaster_ID, func() network.IMessage { return new(S2CDeleteMaster) })
	Processor.Register(PCK_C2SGetPupilTask_ID, func() network.IMessage { return new(C2SGetPupilTask) })
	Processor.Register(PCK_S2CGetPupilTask_ID, func() network.IMessage { return new(S2CGetPupilTask) })
	Processor.Register(PCK_C2SGetMarryList_ID, func() network.IMessage { return new(C2SGetMarryList) })
	Processor.Register(PCK_S2CGetMarryList_ID, func() network.IMessage { return new(S2CGetMarryList) })
	Processor.Register(PCK_C2SGetMarry_ID, func() network.IMessage { return new(C2SGetMarry) })
	Processor.Register(PCK_S2CGetMarry_ID, func() network.IMessage { return new(S2CGetMarry) })
	Processor.Register(PCK_S2CSendMarry_ID, func() network.IMessage { return new(S2CSendMarry) })
	Processor.Register(PCK_C2SHandelMarry_ID, func() network.IMessage { return new(C2SHandelMarry) })
	Processor.Register(PCK_S2CHandelMarry_ID, func() network.IMessage { return new(S2CHandelMarry) })
	Processor.Register(PCK_C2SSendFlower_ID, func() network.IMessage { return new(C2SSendFlower) })
	Processor.Register(PCK_S2CSendFlower_ID, func() network.IMessage { return new(S2CSendFlower) })
	Processor.Register(PCK_S2CGetFlower_ID, func() network.IMessage { return new(S2CGetFlower) })
	Processor.Register(PCK_S2CMarryStatus_ID, func() network.IMessage { return new(S2CMarryStatus) })
	Processor.Register(PCK_C2SMarryUpdate_ID, func() network.IMessage { return new(C2SMarryUpdate) })
	Processor.Register(PCK_S2CMarryUpdate_ID, func() network.IMessage { return new(S2CMarryUpdate) })
	Processor.Register(PCK_C2SHouseUpdate_ID, func() network.IMessage { return new(C2SHouseUpdate) })
	Processor.Register(PCK_S2CHouseUpdate_ID, func() network.IMessage { return new(S2CHouseUpdate) })
	Processor.Register(PCK_C2SGetUpdate_ID, func() network.IMessage { return new(C2SGetUpdate) })
	Processor.Register(PCK_S2CGetUpdate_ID, func() network.IMessage { return new(S2CGetUpdate) })
	Processor.Register(PCK_S2CNewMarry_ID, func() network.IMessage { return new(S2CNewMarry) })
	Processor.Register(PCK_C2SSendMarryGift_ID, func() network.IMessage { return new(C2SSendMarryGift) })
	Processor.Register(PCK_S2CSendMarryGift_ID, func() network.IMessage { return new(S2CSendMarryGift) })
	Processor.Register(PCK_C2SDeleteWife_ID, func() network.IMessage { return new(C2SDeleteWife) })
	Processor.Register(PCK_S2CDeleteWife_ID, func() network.IMessage { return new(S2CDeleteWife) })
	Processor.Register(PCK_S2CMateOnline_ID, func() network.IMessage { return new(S2CMateOnline) })
	Processor.Register(PCK_S2CQuizAsk_ID, func() network.IMessage { return new(S2CQuizAsk) })
	Processor.Register(PCK_S2CQuizRank_ID, func() network.IMessage { return new(S2CQuizRank) })
	Processor.Register(PCK_C2SAnswerQuiz_ID, func() network.IMessage { return new(C2SAnswerQuiz) })
	Processor.Register(PCK_S2CAnswerQuiz_ID, func() network.IMessage { return new(S2CAnswerQuiz) })
	Processor.Register(PCK_S2CQuizSum_ID, func() network.IMessage { return new(S2CQuizSum) })
	Processor.Register(PCK_S2CQuizFirst_ID, func() network.IMessage { return new(S2CQuizFirst) })
	Processor.Register(PCK_S2CQuizStart_ID, func() network.IMessage { return new(S2CQuizStart) })
	Processor.Register(PCK_S2CWestExp_ID, func() network.IMessage { return new(S2CWestExp) })
	Processor.Register(PCK_C2SGetProtectPlayer_ID, func() network.IMessage { return new(C2SGetProtectPlayer) })
	Processor.Register(PCK_S2CGetProtectPlayer_ID, func() network.IMessage { return new(S2CGetProtectPlayer) })
	Processor.Register(PCK_S2CNewProtectPlayer_ID, func() network.IMessage { return new(S2CNewProtectPlayer) })
	Processor.Register(PCK_S2CEndProtectPlayer_ID, func() network.IMessage { return new(S2CEndProtectPlayer) })
	Processor.Register(PCK_C2SGetWestExp_ID, func() network.IMessage { return new(C2SGetWestExp) })
	Processor.Register(PCK_S2CGetWestExp_ID, func() network.IMessage { return new(S2CGetWestExp) })
	Processor.Register(PCK_C2SQuickFinishWestExp_ID, func() network.IMessage { return new(C2SQuickFinishWestExp) })
	Processor.Register(PCK_S2CQuickFinishWestExp_ID, func() network.IMessage { return new(S2CQuickFinishWestExp) })
	Processor.Register(PCK_C2SStartWestExp_ID, func() network.IMessage { return new(C2SStartWestExp) })
	Processor.Register(PCK_S2CStartWestExp_ID, func() network.IMessage { return new(S2CStartWestExp) })
	Processor.Register(PCK_S2CFinishWestExp_ID, func() network.IMessage { return new(S2CFinishWestExp) })
	Processor.Register(PCK_C2SGetWestPrize_ID, func() network.IMessage { return new(C2SGetWestPrize) })
	Processor.Register(PCK_S2CGetWestPrize_ID, func() network.IMessage { return new(S2CGetWestPrize) })
	Processor.Register(PCK_C2SGetRobbedList_ID, func() network.IMessage { return new(C2SGetRobbedList) })
	Processor.Register(PCK_S2CGetRobbedList_ID, func() network.IMessage { return new(S2CGetRobbedList) })
	Processor.Register(PCK_C2SSendRob_ID, func() network.IMessage { return new(C2SSendRob) })
	Processor.Register(PCK_S2CSendRob_ID, func() network.IMessage { return new(S2CSendRob) })
	Processor.Register(PCK_S2CBeRob_ID, func() network.IMessage { return new(S2CBeRob) })
	Processor.Register(PCK_C2SSendRevenge_ID, func() network.IMessage { return new(C2SSendRevenge) })
	Processor.Register(PCK_S2CSendRevenge_ID, func() network.IMessage { return new(S2CSendRevenge) })
	Processor.Register(PCK_S2CWestExpStart_ID, func() network.IMessage { return new(S2CWestExpStart) })
	Processor.Register(PCK_C2SAllRank_ID, func() network.IMessage { return new(C2SAllRank) })
	Processor.Register(PCK_S2CAllRank_ID, func() network.IMessage { return new(S2CAllRank) })
	Processor.Register(PCK_C2SRespect_ID, func() network.IMessage { return new(C2SRespect) })
	Processor.Register(PCK_S2CRespect_ID, func() network.IMessage { return new(S2CRespect) })
	Processor.Register(PCK_S2CGangRedPot_ID, func() network.IMessage { return new(S2CGangRedPot) })
	Processor.Register(PCK_C2SGangFindExchange_ID, func() network.IMessage { return new(C2SGangFindExchange) })
	Processor.Register(PCK_S2CGangFindExchange_ID, func() network.IMessage { return new(S2CGangFindExchange) })
	Processor.Register(PCK_C2SGangExchange_ID, func() network.IMessage { return new(C2SGangExchange) })
	Processor.Register(PCK_S2CGangExchange_ID, func() network.IMessage { return new(S2CGangExchange) })
	Processor.Register(PCK_C2SGangRefreshExc_ID, func() network.IMessage { return new(C2SGangRefreshExc) })
	Processor.Register(PCK_S2CGangRefreshExc_ID, func() network.IMessage { return new(S2CGangRefreshExc) })
	Processor.Register(PCK_C2SGangSimpleInfo_ID, func() network.IMessage { return new(C2SGangSimpleInfo) })
	Processor.Register(PCK_S2CGangSimpleInfo_ID, func() network.IMessage { return new(S2CGangSimpleInfo) })
	Processor.Register(PCK_C2SGangInfo_ID, func() network.IMessage { return new(C2SGangInfo) })
	Processor.Register(PCK_S2CGangInfo_ID, func() network.IMessage { return new(S2CGangInfo) })
	Processor.Register(PCK_C2SGangList_ID, func() network.IMessage { return new(C2SGangList) })
	Processor.Register(PCK_S2CGangList_ID, func() network.IMessage { return new(S2CGangList) })
	Processor.Register(PCK_C2SCreateGang_ID, func() network.IMessage { return new(C2SCreateGang) })
	Processor.Register(PCK_S2CCreateGang_ID, func() network.IMessage { return new(S2CCreateGang) })
	Processor.Register(PCK_C2SApplyGang_ID, func() network.IMessage { return new(C2SApplyGang) })
	Processor.Register(PCK_S2CApplyGang_ID, func() network.IMessage { return new(S2CApplyGang) })
	Processor.Register(PCK_C2SAgreeApplyGang_ID, func() network.IMessage { return new(C2SAgreeApplyGang) })
	Processor.Register(PCK_S2CAgreeApplyGang_ID, func() network.IMessage { return new(S2CAgreeApplyGang) })
	Processor.Register(PCK_C2SUpNoticeGang_ID, func() network.IMessage { return new(C2SUpNoticeGang) })
	Processor.Register(PCK_S2CUpNoticeGang_ID, func() network.IMessage { return new(S2CUpNoticeGang) })
	Processor.Register(PCK_C2SReNameGang_ID, func() network.IMessage { return new(C2SReNameGang) })
	Processor.Register(PCK_S2CReNameGang_ID, func() network.IMessage { return new(S2CReNameGang) })
	Processor.Register(PCK_C2SGangMember_ID, func() network.IMessage { return new(C2SGangMember) })
	Processor.Register(PCK_S2CGangMember_ID, func() network.IMessage { return new(S2CGangMember) })
	Processor.Register(PCK_C2SGangApplys_ID, func() network.IMessage { return new(C2SGangApplys) })
	Processor.Register(PCK_S2CGangApplys_ID, func() network.IMessage { return new(S2CGangApplys) })
	Processor.Register(PCK_C2SGangKickMasterConsumes_ID, func() network.IMessage { return new(C2SGangKickMasterConsumes) })
	Processor.Register(PCK_S2CGangKickMasterConsumes_ID, func() network.IMessage { return new(S2CGangKickMasterConsumes) })
	Processor.Register(PCK_C2SGangKickMaster_ID, func() network.IMessage { return new(C2SGangKickMaster) })
	Processor.Register(PCK_S2CGangKickMaster_ID, func() network.IMessage { return new(S2CGangKickMaster) })
	Processor.Register(PCK_C2SUpGangFightVal_ID, func() network.IMessage { return new(C2SUpGangFightVal) })
	Processor.Register(PCK_S2CUpGangFightVal_ID, func() network.IMessage { return new(S2CUpGangFightVal) })
	Processor.Register(PCK_C2SUpGangMoney_ID, func() network.IMessage { return new(C2SUpGangMoney) })
	Processor.Register(PCK_S2CUpGangMoney_ID, func() network.IMessage { return new(S2CUpGangMoney) })
	Processor.Register(PCK_C2SGangRecruit_ID, func() network.IMessage { return new(C2SGangRecruit) })
	Processor.Register(PCK_S2CGangRecruit_ID, func() network.IMessage { return new(S2CGangRecruit) })
	Processor.Register(PCK_C2SGangGuardupLevel_ID, func() network.IMessage { return new(C2SGangGuardupLevel) })
	Processor.Register(PCK_S2CGangGuardupLevel_ID, func() network.IMessage { return new(S2CGangGuardupLevel) })
	Processor.Register(PCK_C2SGangUpDownDuty_ID, func() network.IMessage { return new(C2SGangUpDownDuty) })
	Processor.Register(PCK_S2CGangUpDownDuty_ID, func() network.IMessage { return new(S2CGangUpDownDuty) })
	Processor.Register(PCK_C2SGangKick_ID, func() network.IMessage { return new(C2SGangKick) })
	Processor.Register(PCK_S2CGangKick_ID, func() network.IMessage { return new(S2CGangKick) })
	Processor.Register(PCK_C2SGangRecord_ID, func() network.IMessage { return new(C2SGangRecord) })
	Processor.Register(PCK_S2CGangRecord_ID, func() network.IMessage { return new(S2CGangRecord) })
	Processor.Register(PCK_C2SGangExit_ID, func() network.IMessage { return new(C2SGangExit) })
	Processor.Register(PCK_S2CGangExit_ID, func() network.IMessage { return new(S2CGangExit) })
	Processor.Register(PCK_C2SGangEnterMap_ID, func() network.IMessage { return new(C2SGangEnterMap) })
	Processor.Register(PCK_S2CGangEnterMap_ID, func() network.IMessage { return new(S2CGangEnterMap) })
	Processor.Register(PCK_C2SGangBossEnterMap_ID, func() network.IMessage { return new(C2SGangBossEnterMap) })
	Processor.Register(PCK_S2CGangBossEnterMap_ID, func() network.IMessage { return new(S2CGangBossEnterMap) })
	Processor.Register(PCK_C2SGangLeaveMap_ID, func() network.IMessage { return new(C2SGangLeaveMap) })
	Processor.Register(PCK_S2CGangLeaveMap_ID, func() network.IMessage { return new(S2CGangLeaveMap) })
	Processor.Register(PCK_C2SGangGuardReward_ID, func() network.IMessage { return new(C2SGangGuardReward) })
	Processor.Register(PCK_S2CGangGuardReward_ID, func() network.IMessage { return new(S2CGangGuardReward) })
	Processor.Register(PCK_C2SGangFindReward_ID, func() network.IMessage { return new(C2SGangFindReward) })
	Processor.Register(PCK_S2CGangFindReward_ID, func() network.IMessage { return new(S2CGangFindReward) })
	Processor.Register(PCK_C2SGangUpSkill_ID, func() network.IMessage { return new(C2SGangUpSkill) })
	Processor.Register(PCK_S2CGangUpSkill_ID, func() network.IMessage { return new(S2CGangUpSkill) })
	Processor.Register(PCK_C2SGangFindSkill_ID, func() network.IMessage { return new(C2SGangFindSkill) })
	Processor.Register(PCK_S2CGangFindSkill_ID, func() network.IMessage { return new(S2CGangFindSkill) })
	Processor.Register(PCK_C2SGangFindPeach_ID, func() network.IMessage { return new(C2SGangFindPeach) })
	Processor.Register(PCK_S2CGangFindPeach_ID, func() network.IMessage { return new(S2CGangFindPeach) })
	Processor.Register(PCK_C2SGangEatPeach_ID, func() network.IMessage { return new(C2SGangEatPeach) })
	Processor.Register(PCK_S2CGangEatPeach_ID, func() network.IMessage { return new(S2CGangEatPeach) })
	Processor.Register(PCK_C2SGangPeachBox_ID, func() network.IMessage { return new(C2SGangPeachBox) })
	Processor.Register(PCK_S2CGangPeachBox_ID, func() network.IMessage { return new(S2CGangPeachBox) })
	Processor.Register(PCK_C2SGangAutoJoin_ID, func() network.IMessage { return new(C2SGangAutoJoin) })
	Processor.Register(PCK_S2CGangAutoJoin_ID, func() network.IMessage { return new(S2CGangAutoJoin) })
	Processor.Register(PCK_C2SGangMapOnekeyOver_ID, func() network.IMessage { return new(C2SGangMapOnekeyOver) })
	Processor.Register(PCK_S2CGangMapOneKeyOver_ID, func() network.IMessage { return new(S2CGangMapOneKeyOver) })
	Processor.Register(PCK_C2SGangMapResetInfo_ID, func() network.IMessage { return new(C2SGangMapResetInfo) })
	Processor.Register(PCK_S2CGangMapResetInfo_ID, func() network.IMessage { return new(S2CGangMapResetInfo) })
	Processor.Register(PCK_C2SGangMapReset_ID, func() network.IMessage { return new(C2SGangMapReset) })
	Processor.Register(PCK_S2CGangMapReset_ID, func() network.IMessage { return new(S2CGangMapReset) })
	Processor.Register(PCK_C2SRedPkgPool_ID, func() network.IMessage { return new(C2SRedPkgPool) })
	Processor.Register(PCK_S2CRedPkgPool_ID, func() network.IMessage { return new(S2CRedPkgPool) })
	Processor.Register(PCK_C2SRedPkgSend_ID, func() network.IMessage { return new(C2SRedPkgSend) })
	Processor.Register(PCK_S2CRedPkgSend_ID, func() network.IMessage { return new(S2CRedPkgSend) })
	Processor.Register(PCK_S2CRedPkgNotice_ID, func() network.IMessage { return new(S2CRedPkgNotice) })
	Processor.Register(PCK_C2SRedPkgRecieve_ID, func() network.IMessage { return new(C2SRedPkgRecieve) })
	Processor.Register(PCK_S2CRedPkgRecieve_ID, func() network.IMessage { return new(S2CRedPkgRecieve) })
	Processor.Register(PCK_C2SRedPkgSingle_ID, func() network.IMessage { return new(C2SRedPkgSingle) })
	Processor.Register(PCK_S2CRedPkgSingle_ID, func() network.IMessage { return new(S2CRedPkgSingle) })
	Processor.Register(PCK_C2SRedPkgPerRecord_ID, func() network.IMessage { return new(C2SRedPkgPerRecord) })
	Processor.Register(PCK_S2CRedPkgPerRecord_ID, func() network.IMessage { return new(S2CRedPkgPerRecord) })
	Processor.Register(PCK_C2SRedPckShieldLog_ID, func() network.IMessage { return new(C2SRedPckShieldLog) })
	Processor.Register(PCK_S2CRedPckShieldLog_ID, func() network.IMessage { return new(S2CRedPckShieldLog) })
	Processor.Register(PCK_S2KGangSimpleInfo_ID, func() network.IMessage { return new(S2KGangSimpleInfo) })
	Processor.Register(PCK_C2SJJCList_ID, func() network.IMessage { return new(C2SJJCList) })
	Processor.Register(PCK_S2CJJCList_ID, func() network.IMessage { return new(S2CJJCList) })
	Processor.Register(PCK_C2SJJCFight_ID, func() network.IMessage { return new(C2SJJCFight) })
	Processor.Register(PCK_S2CJJCFight_ID, func() network.IMessage { return new(S2CJJCFight) })
	Processor.Register(PCK_C2SJJCGetBuyInfo_ID, func() network.IMessage { return new(C2SJJCGetBuyInfo) })
	Processor.Register(PCK_S2CJJCGetBuyInfo_ID, func() network.IMessage { return new(S2CJJCGetBuyInfo) })
	Processor.Register(PCK_C2SJJCBuyTimes_ID, func() network.IMessage { return new(C2SJJCBuyTimes) })
	Processor.Register(PCK_S2CJJCBuyTimes_ID, func() network.IMessage { return new(S2CJJCBuyTimes) })
	Processor.Register(PCK_C2SGodEquipAwake_ID, func() network.IMessage { return new(C2SGodEquipAwake) })
	Processor.Register(PCK_S2CGodEquipAwake_ID, func() network.IMessage { return new(S2CGodEquipAwake) })
	Processor.Register(PCK_C2SIntoSoulGradeUp_ID, func() network.IMessage { return new(C2SIntoSoulGradeUp) })
	Processor.Register(PCK_S2CIntoSoulGradeUp_ID, func() network.IMessage { return new(S2CIntoSoulGradeUp) })
	Processor.Register(PCK_C2SGodForge_ID, func() network.IMessage { return new(C2SGodForge) })
	Processor.Register(PCK_S2CGodForge_ID, func() network.IMessage { return new(S2CGodForge) })
	Processor.Register(PCK_C2SGodForgeSave_ID, func() network.IMessage { return new(C2SGodForgeSave) })
	Processor.Register(PCK_S2CGodForgeSave_ID, func() network.IMessage { return new(S2CGodForgeSave) })
	Processor.Register(PCK_C2SGodEquipMelt_ID, func() network.IMessage { return new(C2SGodEquipMelt) })
	Processor.Register(PCK_S2CGodEquipMelt_ID, func() network.IMessage { return new(S2CGodEquipMelt) })
	Processor.Register(PCK_C2SGetTeamList_ID, func() network.IMessage { return new(C2SGetTeamList) })
	Processor.Register(PCK_S2CGetTeamList_ID, func() network.IMessage { return new(S2CGetTeamList) })
	Processor.Register(PCK_C2SGetMemberList_ID, func() network.IMessage { return new(C2SGetMemberList) })
	Processor.Register(PCK_S2CGetMemberList_ID, func() network.IMessage { return new(S2CGetMemberList) })
	Processor.Register(PCK_S2KCrossChat_ID, func() network.IMessage { return new(S2KCrossChat) })
	Processor.Register(PCK_K2SCrossChat_ID, func() network.IMessage { return new(K2SCrossChat) })
	Processor.Register(PCK_K2SAddTmpTitle_ID, func() network.IMessage { return new(K2SAddTmpTitle) })
	Processor.Register(PCK_C2SJoinInstance_ID, func() network.IMessage { return new(C2SJoinInstance) })
	Processor.Register(PCK_S2CJoinInstance_ID, func() network.IMessage { return new(S2CJoinInstance) })
	Processor.Register(PCK_C2SLeaveInstance_ID, func() network.IMessage { return new(C2SLeaveInstance) })
	Processor.Register(PCK_S2CLeaveInstance_ID, func() network.IMessage { return new(S2CLeaveInstance) })
	Processor.Register(PCK_C2SLeaveInstanceCopy_ID, func() network.IMessage { return new(C2SLeaveInstanceCopy) })
	Processor.Register(PCK_S2CLeaveInstanceCopy_ID, func() network.IMessage { return new(S2CLeaveInstanceCopy) })
	Processor.Register(PCK_C2SKillInstance_ID, func() network.IMessage { return new(C2SKillInstance) })
	Processor.Register(PCK_S2CKillInstance_ID, func() network.IMessage { return new(S2CKillInstance) })
	Processor.Register(PCK_C2SStartInstance_ID, func() network.IMessage { return new(C2SStartInstance) })
	Processor.Register(PCK_S2CStartInstance_ID, func() network.IMessage { return new(S2CStartInstance) })
	Processor.Register(PCK_C2SGetPreciousPos_ID, func() network.IMessage { return new(C2SGetPreciousPos) })
	Processor.Register(PCK_S2CGetPreciousPos_ID, func() network.IMessage { return new(S2CGetPreciousPos) })
	Processor.Register(PCK_C2SCreatePrecious_ID, func() network.IMessage { return new(C2SCreatePrecious) })
	Processor.Register(PCK_S2CCreatePrecious_ID, func() network.IMessage { return new(S2CCreatePrecious) })
	Processor.Register(PCK_C2SCreatePreciousFast_ID, func() network.IMessage { return new(C2SCreatePreciousFast) })
	Processor.Register(PCK_S2CCreatePreciousFast_ID, func() network.IMessage { return new(S2CCreatePreciousFast) })
	Processor.Register(PCK_C2SMeltPrecious_ID, func() network.IMessage { return new(C2SMeltPrecious) })
	Processor.Register(PCK_S2CMeltPrecious_ID, func() network.IMessage { return new(S2CMeltPrecious) })
	Processor.Register(PCK_C2SWearPrecious_ID, func() network.IMessage { return new(C2SWearPrecious) })
	Processor.Register(PCK_S2CWearPrecious_ID, func() network.IMessage { return new(S2CWearPrecious) })
	Processor.Register(PCK_C2SLockPrecious_ID, func() network.IMessage { return new(C2SLockPrecious) })
	Processor.Register(PCK_S2CLockPrecious_ID, func() network.IMessage { return new(S2CLockPrecious) })
	Processor.Register(PCK_C2SPreciousForge_ID, func() network.IMessage { return new(C2SPreciousForge) })
	Processor.Register(PCK_S2CPreciousForge_ID, func() network.IMessage { return new(S2CPreciousForge) })
	Processor.Register(PCK_C2SPreciousEat_ID, func() network.IMessage { return new(C2SPreciousEat) })
	Processor.Register(PCK_S2CPreciousEat_ID, func() network.IMessage { return new(S2CPreciousEat) })
	Processor.Register(PCK_C2SPreciousSoul_ID, func() network.IMessage { return new(C2SPreciousSoul) })
	Processor.Register(PCK_S2CPreciousSoul_ID, func() network.IMessage { return new(S2CPreciousSoul) })
	Processor.Register(PCK_C2SPreciousSoulUp_ID, func() network.IMessage { return new(C2SPreciousSoulUp) })
	Processor.Register(PCK_S2CPreciousSoulUp_ID, func() network.IMessage { return new(S2CPreciousSoulUp) })
	Processor.Register(PCK_C2SPreciousGive_ID, func() network.IMessage { return new(C2SPreciousGive) })
	Processor.Register(PCK_S2CPreciousGive_ID, func() network.IMessage { return new(S2CPreciousGive) })
	Processor.Register(PCK_C2SPreciousRobot_ID, func() network.IMessage { return new(C2SPreciousRobot) })
	Processor.Register(PCK_S2CPreciousRobot_ID, func() network.IMessage { return new(S2CPreciousRobot) })
	Processor.Register(PCK_S2CActGangWarSettlement_ID, func() network.IMessage { return new(S2CActGangWarSettlement) })
	Processor.Register(PCK_S2CGangWarBeforeCD_ID, func() network.IMessage { return new(S2CGangWarBeforeCD) })
	Processor.Register(PCK_C2SActSountGateFight_ID, func() network.IMessage { return new(C2SActSountGateFight) })
	Processor.Register(PCK_S2CActSountGateFight_ID, func() network.IMessage { return new(S2CActSountGateFight) })
	Processor.Register(PCK_S2CSouthGateStatus_ID, func() network.IMessage { return new(S2CSouthGateStatus) })
	Processor.Register(PCK_K2SActNineDayPreInfo_ID, func() network.IMessage { return new(K2SActNineDayPreInfo) })
	Processor.Register(PCK_C2SActNineDayPreInfo_ID, func() network.IMessage { return new(C2SActNineDayPreInfo) })
	Processor.Register(PCK_S2CActNineDayPreInfo_ID, func() network.IMessage { return new(S2CActNineDayPreInfo) })
	Processor.Register(PCK_C2SActNineDaySort_ID, func() network.IMessage { return new(C2SActNineDaySort) })
	Processor.Register(PCK_S2CActNineDaySort_ID, func() network.IMessage { return new(S2CActNineDaySort) })
	Processor.Register(PCK_S2CNiceDayClose_ID, func() network.IMessage { return new(S2CNiceDayClose) })
	Processor.Register(PCK_K2SActGangWarPreInfo_ID, func() network.IMessage { return new(K2SActGangWarPreInfo) })
	Processor.Register(PCK_C2SActGangWarPreInfo_ID, func() network.IMessage { return new(C2SActGangWarPreInfo) })
	Processor.Register(PCK_S2CActGangWarPreInfo_ID, func() network.IMessage { return new(S2CActGangWarPreInfo) })
	Processor.Register(PCK_S2CActStart_ID, func() network.IMessage { return new(S2CActStart) })
	Processor.Register(PCK_S2CActOver_ID, func() network.IMessage { return new(S2CActOver) })
	Processor.Register(PCK_S2CActIcon_ID, func() network.IMessage { return new(S2CActIcon) })
	Processor.Register(PCK_C2SActGangWarPersonRank_ID, func() network.IMessage { return new(C2SActGangWarPersonRank) })
	Processor.Register(PCK_S2CActGangWarPersonRank_ID, func() network.IMessage { return new(S2CActGangWarPersonRank) })
	Processor.Register(PCK_ActGangWarPersonRank_ID, func() network.IMessage { return new(ActGangWarPersonRank) })
	Processor.Register(PCK_C2SActGangWarGangRank_ID, func() network.IMessage { return new(C2SActGangWarGangRank) })
	Processor.Register(PCK_S2CActGangWarGangRank_ID, func() network.IMessage { return new(S2CActGangWarGangRank) })
	Processor.Register(PCK_ActGangWarGangRank_ID, func() network.IMessage { return new(ActGangWarGangRank) })
	Processor.Register(PCK_C2SActGangWarKillRank_ID, func() network.IMessage { return new(C2SActGangWarKillRank) })
	Processor.Register(PCK_S2CActGangWarKillRank_ID, func() network.IMessage { return new(S2CActGangWarKillRank) })
	Processor.Register(PCK_ActGangWarKillRank_ID, func() network.IMessage { return new(ActGangWarKillRank) })
	Processor.Register(PCK_C2SActGangWarScoreRank_ID, func() network.IMessage { return new(C2SActGangWarScoreRank) })
	Processor.Register(PCK_S2CActGangWarScoreRank_ID, func() network.IMessage { return new(S2CActGangWarScoreRank) })
	Processor.Register(PCK_ActGangWarScoreRank_ID, func() network.IMessage { return new(ActGangWarScoreRank) })
	Processor.Register(PCK_C2SActGangWarScoreExchange_ID, func() network.IMessage { return new(C2SActGangWarScoreExchange) })
	Processor.Register(PCK_S2CActGangWarScoreExchange_ID, func() network.IMessage { return new(S2CActGangWarScoreExchange) })
	Processor.Register(PCK_C2SActGangWarScoreFind_ID, func() network.IMessage { return new(C2SActGangWarScoreFind) })
	Processor.Register(PCK_S2CActGangWarScoreFind_ID, func() network.IMessage { return new(S2CActGangWarScoreFind) })
	Processor.Register(PCK_C2SActGangWarSwitchMap_ID, func() network.IMessage { return new(C2SActGangWarSwitchMap) })
	Processor.Register(PCK_S2CActGangWarSwitchMap_ID, func() network.IMessage { return new(S2CActGangWarSwitchMap) })
	Processor.Register(PCK_S2CActGangWarInDragonGangs_ID, func() network.IMessage { return new(S2CActGangWarInDragonGangs) })
	Processor.Register(PCK_C2SActGiveUpDragon_ID, func() network.IMessage { return new(C2SActGiveUpDragon) })
	Processor.Register(PCK_S2CActGiveUpDragon_ID, func() network.IMessage { return new(S2CActGiveUpDragon) })
	Processor.Register(PCK_C2SActCreateTeam_ID, func() network.IMessage { return new(C2SActCreateTeam) })
	Processor.Register(PCK_S2CActCreateTeam_ID, func() network.IMessage { return new(S2CActCreateTeam) })
	Processor.Register(PCK_C2SActFindTeams_ID, func() network.IMessage { return new(C2SActFindTeams) })
	Processor.Register(PCK_S2CActFindTeams_ID, func() network.IMessage { return new(S2CActFindTeams) })
	Processor.Register(PCK_C2SActJoinTeam_ID, func() network.IMessage { return new(C2SActJoinTeam) })
	Processor.Register(PCK_S2CActJoinTeam_ID, func() network.IMessage { return new(S2CActJoinTeam) })
	Processor.Register(PCK_C2SActFindTeam_ID, func() network.IMessage { return new(C2SActFindTeam) })
	Processor.Register(PCK_S2CActFindTeam_ID, func() network.IMessage { return new(S2CActFindTeam) })
	Processor.Register(PCK_C2SActNeedFightVal_ID, func() network.IMessage { return new(C2SActNeedFightVal) })
	Processor.Register(PCK_S2CActNeedFightVal_ID, func() network.IMessage { return new(S2CActNeedFightVal) })
	Processor.Register(PCK_C2SActExitTeam_ID, func() network.IMessage { return new(C2SActExitTeam) })
	Processor.Register(PCK_S2CActExitTeam_ID, func() network.IMessage { return new(S2CActExitTeam) })
	Processor.Register(PCK_ActTeamInfo_ID, func() network.IMessage { return new(ActTeamInfo) })
	Processor.Register(PCK_ActTeamMemInfo_ID, func() network.IMessage { return new(ActTeamMemInfo) })
	Processor.Register(PCK_C2SActTeamRecruit_ID, func() network.IMessage { return new(C2SActTeamRecruit) })
	Processor.Register(PCK_S2CActTeamRecruit_ID, func() network.IMessage { return new(S2CActTeamRecruit) })
	Processor.Register(PCK_C2SJoinActive_ID, func() network.IMessage { return new(C2SJoinActive) })
	Processor.Register(PCK_S2CJoinActive_ID, func() network.IMessage { return new(S2CJoinActive) })
	Processor.Register(PCK_C2SLeaveActive_ID, func() network.IMessage { return new(C2SLeaveActive) })
	Processor.Register(PCK_S2CLeaveActive_ID, func() network.IMessage { return new(S2CLeaveActive) })
	Processor.Register(PCK_K2SJoinFuncActiveHandleEvent_ID, func() network.IMessage { return new(K2SJoinFuncActiveHandleEvent) })
	Processor.Register(PCK_S2CGangWarDragonPosition_ID, func() network.IMessage { return new(S2CGangWarDragonPosition) })
	Processor.Register(PCK_S2CGangWarGangSort_ID, func() network.IMessage { return new(S2CGangWarGangSort) })
	Processor.Register(PCK_GangHurtInfo_ID, func() network.IMessage { return new(GangHurtInfo) })
	Processor.Register(PCK_S2CGangActWarGangPCount_ID, func() network.IMessage { return new(S2CGangActWarGangPCount) })
	Processor.Register(PCK_S2CGangActWarGangLocal_ID, func() network.IMessage { return new(S2CGangActWarGangLocal) })
	Processor.Register(PCK_S2CActiveGangWarScoreChange_ID, func() network.IMessage { return new(S2CActiveGangWarScoreChange) })
	Processor.Register(PCK_S2CActiveGangWarBeforeInsideScore_ID, func() network.IMessage { return new(S2CActiveGangWarBeforeInsideScore) })
	Processor.Register(PCK_C2SReviveLife_ID, func() network.IMessage { return new(C2SReviveLife) })
	Processor.Register(PCK_S2CReviveLife_ID, func() network.IMessage { return new(S2CReviveLife) })
	Processor.Register(PCK_C2SGangBossOverTime_ID, func() network.IMessage { return new(C2SGangBossOverTime) })
	Processor.Register(PCK_S2CGangBossOverTime_ID, func() network.IMessage { return new(S2CGangBossOverTime) })
	Processor.Register(PCK_K2SAddGangRedPkgMoney_ID, func() network.IMessage { return new(K2SAddGangRedPkgMoney) })
	Processor.Register(PCK_C2SGangBossHurtSort_ID, func() network.IMessage { return new(C2SGangBossHurtSort) })
	Processor.Register(PCK_S2CGangBossHurtSort_ID, func() network.IMessage { return new(S2CGangBossHurtSort) })
	Processor.Register(PCK_S2CGangBossFiveHurt_ID, func() network.IMessage { return new(S2CGangBossFiveHurt) })
	Processor.Register(PCK_C2SGangBossKiller_ID, func() network.IMessage { return new(C2SGangBossKiller) })
	Processor.Register(PCK_S2CGangBossKiller_ID, func() network.IMessage { return new(S2CGangBossKiller) })
	Processor.Register(PCK_S2CGangBossRedPot_ID, func() network.IMessage { return new(S2CGangBossRedPot) })
	Processor.Register(PCK_C2SGangCollect_ID, func() network.IMessage { return new(C2SGangCollect) })
	Processor.Register(PCK_S2CGangCollect_ID, func() network.IMessage { return new(S2CGangCollect) })
	Processor.Register(PCK_C2SActiveNineDayFindExchange_ID, func() network.IMessage { return new(C2SActiveNineDayFindExchange) })
	Processor.Register(PCK_S2CActiveNineDayFindExchange_ID, func() network.IMessage { return new(S2CActiveNineDayFindExchange) })
	Processor.Register(PCK_S2CActiveNineDayScoreChange_ID, func() network.IMessage { return new(S2CActiveNineDayScoreChange) })
	Processor.Register(PCK_C2SActiveNineDayReviceExchange_ID, func() network.IMessage { return new(C2SActiveNineDayReviceExchange) })
	Processor.Register(PCK_S2CActiveNineDayReviceExchange_ID, func() network.IMessage { return new(S2CActiveNineDayReviceExchange) })
	Processor.Register(PCK_C2SGoldTree_ID, func() network.IMessage { return new(C2SGoldTree) })
	Processor.Register(PCK_S2CGoldTree_ID, func() network.IMessage { return new(S2CGoldTree) })
	Processor.Register(PCK_C2SGetGoldTreeInfo_ID, func() network.IMessage { return new(C2SGetGoldTreeInfo) })
	Processor.Register(PCK_S2CGetGoldTreeInfo_ID, func() network.IMessage { return new(S2CGetGoldTreeInfo) })
	Processor.Register(PCK_S2CNotice_ID, func() network.IMessage { return new(S2CNotice) })
	Processor.Register(PCK_C2SGetDragonLog_ID, func() network.IMessage { return new(C2SGetDragonLog) })
	Processor.Register(PCK_S2CGetDragonLog_ID, func() network.IMessage { return new(S2CGetDragonLog) })
	Processor.Register(PCK_C2SStartFightDragon_ID, func() network.IMessage { return new(C2SStartFightDragon) })
	Processor.Register(PCK_S2CStartFightDragon_ID, func() network.IMessage { return new(S2CStartFightDragon) })
	Processor.Register(PCK_S2CContinueRechargeTaskNum_ID, func() network.IMessage { return new(S2CContinueRechargeTaskNum) })
	Processor.Register(PCK_S2CAllActivity_ID, func() network.IMessage { return new(S2CAllActivity) })
	Processor.Register(PCK_S2CActivityInit_ID, func() network.IMessage { return new(S2CActivityInit) })
	Processor.Register(PCK_S2CActivityStart_ID, func() network.IMessage { return new(S2CActivityStart) })
	Processor.Register(PCK_S2CActivityEnd_ID, func() network.IMessage { return new(S2CActivityEnd) })
	Processor.Register(PCK_S2CActivityRedPoint_ID, func() network.IMessage { return new(S2CActivityRedPoint) })
	Processor.Register(PCK_S2CActivityIcon_ID, func() network.IMessage { return new(S2CActivityIcon) })
	Processor.Register(PCK_S2CActivityData_ID, func() network.IMessage { return new(S2CActivityData) })
	Processor.Register(PCK_C2SGetDrawData_ID, func() network.IMessage { return new(C2SGetDrawData) })
	Processor.Register(PCK_S2CGetDrawData_ID, func() network.IMessage { return new(S2CGetDrawData) })
	Processor.Register(PCK_C2SPlayerDrawData_ID, func() network.IMessage { return new(C2SPlayerDrawData) })
	Processor.Register(PCK_S2CPlayerDrawData_ID, func() network.IMessage { return new(S2CPlayerDrawData) })
	Processor.Register(PCK_C2SGetDrawListRate_ID, func() network.IMessage { return new(C2SGetDrawListRate) })
	Processor.Register(PCK_S2CGetDrawListRate_ID, func() network.IMessage { return new(S2CGetDrawListRate) })
	Processor.Register(PCK_C2SBuyDrawItem_ID, func() network.IMessage { return new(C2SBuyDrawItem) })
	Processor.Register(PCK_S2CBuyDrawItem_ID, func() network.IMessage { return new(S2CBuyDrawItem) })
	Processor.Register(PCK_C2SDraw_ID, func() network.IMessage { return new(C2SDraw) })
	Processor.Register(PCK_S2CDraw_ID, func() network.IMessage { return new(S2CDraw) })
	Processor.Register(PCK_C2SGetDrawLog_ID, func() network.IMessage { return new(C2SGetDrawLog) })
	Processor.Register(PCK_S2CGetDrawLog_ID, func() network.IMessage { return new(S2CGetDrawLog) })
	Processor.Register(PCK_C2SGetChargeReturnData_ID, func() network.IMessage { return new(C2SGetChargeReturnData) })
	Processor.Register(PCK_S2CGetChargeReturnData_ID, func() network.IMessage { return new(S2CGetChargeReturnData) })
	Processor.Register(PCK_C2SPlayerChargeReturnData_ID, func() network.IMessage { return new(C2SPlayerChargeReturnData) })
	Processor.Register(PCK_S2CPlayerChargeReturnData_ID, func() network.IMessage { return new(S2CPlayerChargeReturnData) })
	Processor.Register(PCK_C2SChargeReturn_ID, func() network.IMessage { return new(C2SChargeReturn) })
	Processor.Register(PCK_S2CChargeReturn_ID, func() network.IMessage { return new(S2CChargeReturn) })
	Processor.Register(PCK_C2SChargeReturnReceivePrize_ID, func() network.IMessage { return new(C2SChargeReturnReceivePrize) })
	Processor.Register(PCK_S2CChargeReturnReceivePrize_ID, func() network.IMessage { return new(S2CChargeReturnReceivePrize) })
	Processor.Register(PCK_C2SGetChargeReturnLog_ID, func() network.IMessage { return new(C2SGetChargeReturnLog) })
	Processor.Register(PCK_S2CGetChargeReturnLog_ID, func() network.IMessage { return new(S2CGetChargeReturnLog) })
	Processor.Register(PCK_C2SGetActShopList_ID, func() network.IMessage { return new(C2SGetActShopList) })
	Processor.Register(PCK_S2CGetActShopList_ID, func() network.IMessage { return new(S2CGetActShopList) })
	Processor.Register(PCK_C2SActShopBuy_ID, func() network.IMessage { return new(C2SActShopBuy) })
	Processor.Register(PCK_S2CActShopBuy_ID, func() network.IMessage { return new(S2CActShopBuy) })
	Processor.Register(PCK_C2SGetActGiftList_ID, func() network.IMessage { return new(C2SGetActGiftList) })
	Processor.Register(PCK_S2CGetActGiftList_ID, func() network.IMessage { return new(S2CGetActGiftList) })
	Processor.Register(PCK_C2SActGiftBuy_ID, func() network.IMessage { return new(C2SActGiftBuy) })
	Processor.Register(PCK_S2CActGiftBuy_ID, func() network.IMessage { return new(S2CActGiftBuy) })
	Processor.Register(PCK_C2SGetActBossInfo_ID, func() network.IMessage { return new(C2SGetActBossInfo) })
	Processor.Register(PCK_S2CGetActBossInfo_ID, func() network.IMessage { return new(S2CGetActBossInfo) })
	Processor.Register(PCK_C2SPlayerBossInfo_ID, func() network.IMessage { return new(C2SPlayerBossInfo) })
	Processor.Register(PCK_S2CPlayerBossInfo_ID, func() network.IMessage { return new(S2CPlayerBossInfo) })
	Processor.Register(PCK_C2SFightActBoss_ID, func() network.IMessage { return new(C2SFightActBoss) })
	Processor.Register(PCK_S2CFightActBoss_ID, func() network.IMessage { return new(S2CFightActBoss) })
	Processor.Register(PCK_C2SGetActBossPrize_ID, func() network.IMessage { return new(C2SGetActBossPrize) })
	Processor.Register(PCK_S2CGetActBossPrize_ID, func() network.IMessage { return new(S2CGetActBossPrize) })
	Processor.Register(PCK_C2SGetActGiftLimit_ID, func() network.IMessage { return new(C2SGetActGiftLimit) })
	Processor.Register(PCK_S2CGetActGiftLimit_ID, func() network.IMessage { return new(S2CGetActGiftLimit) })
	Processor.Register(PCK_C2SGetActGiftOnline_ID, func() network.IMessage { return new(C2SGetActGiftOnline) })
	Processor.Register(PCK_S2CGetActGiftOnline_ID, func() network.IMessage { return new(S2CGetActGiftOnline) })
	Processor.Register(PCK_C2SActGetPrize_ID, func() network.IMessage { return new(C2SActGetPrize) })
	Processor.Register(PCK_S2CActGetPrize_ID, func() network.IMessage { return new(S2CActGetPrize) })
	Processor.Register(PCK_C2SGetActGiftMiracle_ID, func() network.IMessage { return new(C2SGetActGiftMiracle) })
	Processor.Register(PCK_S2CGetActGiftMiracle_ID, func() network.IMessage { return new(S2CGetActGiftMiracle) })
	Processor.Register(PCK_C2SGetActGiftXRmbInfo_ID, func() network.IMessage { return new(C2SGetActGiftXRmbInfo) })
	Processor.Register(PCK_S2CGetActGiftXRmbInfo_ID, func() network.IMessage { return new(S2CGetActGiftXRmbInfo) })
	Processor.Register(PCK_C2SGetActCollectFontInfo_ID, func() network.IMessage { return new(C2SGetActCollectFontInfo) })
	Processor.Register(PCK_S2CGetActCollectFontInfo_ID, func() network.IMessage { return new(S2CGetActCollectFontInfo) })
	Processor.Register(PCK_C2SGetActLifeVipGiveInfo_ID, func() network.IMessage { return new(C2SGetActLifeVipGiveInfo) })
	Processor.Register(PCK_S2CGetActLifeVipGiveInfo_ID, func() network.IMessage { return new(S2CGetActLifeVipGiveInfo) })
	Processor.Register(PCK_C2SGetActGift1Info_ID, func() network.IMessage { return new(C2SGetActGift1Info) })
	Processor.Register(PCK_S2CGetActGift1Info_ID, func() network.IMessage { return new(S2CGetActGift1Info) })
	Processor.Register(PCK_C2SGetActGift2Info_ID, func() network.IMessage { return new(C2SGetActGift2Info) })
	Processor.Register(PCK_S2CGetActGift2Info_ID, func() network.IMessage { return new(S2CGetActGift2Info) })
	Processor.Register(PCK_C2SGetActCumulativeInfo_ID, func() network.IMessage { return new(C2SGetActCumulativeInfo) })
	Processor.Register(PCK_S2CGetActCumulativeInfo_ID, func() network.IMessage { return new(S2CGetActCumulativeInfo) })
	Processor.Register(PCK_C2SGetActChargeDayInfo_ID, func() network.IMessage { return new(C2SGetActChargeDayInfo) })
	Processor.Register(PCK_S2CGetActChargeDayInfo_ID, func() network.IMessage { return new(S2CGetActChargeDayInfo) })
	Processor.Register(PCK_C2SGetInvestmentInfo_ID, func() network.IMessage { return new(C2SGetInvestmentInfo) })
	Processor.Register(PCK_S2CGetInvestmentInfo_ID, func() network.IMessage { return new(S2CGetInvestmentInfo) })
	Processor.Register(PCK_C2SCumulativeFirstInfo_ID, func() network.IMessage { return new(C2SCumulativeFirstInfo) })
	Processor.Register(PCK_S2CCumulativeFirstInfo_ID, func() network.IMessage { return new(S2CCumulativeFirstInfo) })
	Processor.Register(PCK_C2SGetActDareBossInfo_ID, func() network.IMessage { return new(C2SGetActDareBossInfo) })
	Processor.Register(PCK_S2CGetActDareBossInfo_ID, func() network.IMessage { return new(S2CGetActDareBossInfo) })
	Processor.Register(PCK_C2SPlayerDareBossInfo_ID, func() network.IMessage { return new(C2SPlayerDareBossInfo) })
	Processor.Register(PCK_S2CPlayerDareBossInfo_ID, func() network.IMessage { return new(S2CPlayerDareBossInfo) })
	Processor.Register(PCK_C2SPlayerDareState_ID, func() network.IMessage { return new(C2SPlayerDareState) })
	Processor.Register(PCK_S2CPlayerDareState_ID, func() network.IMessage { return new(S2CPlayerDareState) })
	Processor.Register(PCK_C2SPlayerDareIsOpen_ID, func() network.IMessage { return new(C2SPlayerDareIsOpen) })
	Processor.Register(PCK_S2CPlayerDareIsOpen_ID, func() network.IMessage { return new(S2CPlayerDareIsOpen) })
	Processor.Register(PCK_C2SFightActDareBoss_ID, func() network.IMessage { return new(C2SFightActDareBoss) })
	Processor.Register(PCK_S2CFightActDareBoss_ID, func() network.IMessage { return new(S2CFightActDareBoss) })
	Processor.Register(PCK_C2SGetActDareBossPrize_ID, func() network.IMessage { return new(C2SGetActDareBossPrize) })
	Processor.Register(PCK_S2CGetActDareBossPrize_ID, func() network.IMessage { return new(S2CGetActDareBossPrize) })
	Processor.Register(PCK_C2SGetChargeMallList_ID, func() network.IMessage { return new(C2SGetChargeMallList) })
	Processor.Register(PCK_S2CGetChargeMallList_ID, func() network.IMessage { return new(S2CGetChargeMallList) })
	Processor.Register(PCK_C2SChargeMallBuy_ID, func() network.IMessage { return new(C2SChargeMallBuy) })
	Processor.Register(PCK_S2CChargeMallBuy_ID, func() network.IMessage { return new(S2CChargeMallBuy) })
	Processor.Register(PCK_S2CPayNotify_ID, func() network.IMessage { return new(S2CPayNotify) })
	Processor.Register(PCK_C2SGetActTask_ID, func() network.IMessage { return new(C2SGetActTask) })
	Processor.Register(PCK_S2CGetActTask_ID, func() network.IMessage { return new(S2CGetActTask) })
	Processor.Register(PCK_S2CGetFirstRechargeTask_ID, func() network.IMessage { return new(S2CGetFirstRechargeTask) })
	Processor.Register(PCK_C2SGetGodTowerActTaskDay_ID, func() network.IMessage { return new(C2SGetGodTowerActTaskDay) })
	Processor.Register(PCK_S2CGetGodTowerActTaskDay_ID, func() network.IMessage { return new(S2CGetGodTowerActTaskDay) })
	Processor.Register(PCK_C2SGetInvestInfo_ID, func() network.IMessage { return new(C2SGetInvestInfo) })
	Processor.Register(PCK_S2CGetInvestInfo_ID, func() network.IMessage { return new(S2CGetInvestInfo) })
	Processor.Register(PCK_C2SPlayerInvestData_ID, func() network.IMessage { return new(C2SPlayerInvestData) })
	Processor.Register(PCK_S2CPlayerInvestData_ID, func() network.IMessage { return new(S2CPlayerInvestData) })
	Processor.Register(PCK_C2SBuyInvest_ID, func() network.IMessage { return new(C2SBuyInvest) })
	Processor.Register(PCK_S2CBuyInvest_ID, func() network.IMessage { return new(S2CBuyInvest) })
	Processor.Register(PCK_C2SGetActRankInfo_ID, func() network.IMessage { return new(C2SGetActRankInfo) })
	Processor.Register(PCK_S2CGetActRankInfo_ID, func() network.IMessage { return new(S2CGetActRankInfo) })
	Processor.Register(PCK_C2SGetActRankData_ID, func() network.IMessage { return new(C2SGetActRankData) })
	Processor.Register(PCK_S2CGetActRankData_ID, func() network.IMessage { return new(S2CGetActRankData) })
	Processor.Register(PCK_C2SGetActPicture_ID, func() network.IMessage { return new(C2SGetActPicture) })
	Processor.Register(PCK_S2CGetActPicture_ID, func() network.IMessage { return new(S2CGetActPicture) })
	Processor.Register(PCK_C2SActPictureLight_ID, func() network.IMessage { return new(C2SActPictureLight) })
	Processor.Register(PCK_S2CActPictureLight_ID, func() network.IMessage { return new(S2CActPictureLight) })
	Processor.Register(PCK_C2SActRecoveryData_ID, func() network.IMessage { return new(C2SActRecoveryData) })
	Processor.Register(PCK_S2CActRecoveryData_ID, func() network.IMessage { return new(S2CActRecoveryData) })
	Processor.Register(PCK_C2SActRecoveryEquip_ID, func() network.IMessage { return new(C2SActRecoveryEquip) })
	Processor.Register(PCK_S2CActRecoveryEquip_ID, func() network.IMessage { return new(S2CActRecoveryEquip) })
	Processor.Register(PCK_C2SShowItem_ID, func() network.IMessage { return new(C2SShowItem) })
	Processor.Register(PCK_S2CShowItem_ID, func() network.IMessage { return new(S2CShowItem) })
	Processor.Register(PCK_S2CReportShowItem_ID, func() network.IMessage { return new(S2CReportShowItem) })
	Processor.Register(PCK_C2SGetShowInfo_ID, func() network.IMessage { return new(C2SGetShowInfo) })
	Processor.Register(PCK_S2CGetShowInfo_ID, func() network.IMessage { return new(S2CGetShowInfo) })
	Processor.Register(PCK_S2KShowItem_ID, func() network.IMessage { return new(S2KShowItem) })
	Processor.Register(PCK_C2SRedeemCode_ID, func() network.IMessage { return new(C2SRedeemCode) })
	Processor.Register(PCK_S2CRedeemCode_ID, func() network.IMessage { return new(S2CRedeemCode) })
	Processor.Register(PCK_C2SRealName_ID, func() network.IMessage { return new(C2SRealName) })
	Processor.Register(PCK_S2CRealName_ID, func() network.IMessage { return new(S2CRealName) })
	Processor.Register(PCK_C2SGetRealNamePrize_ID, func() network.IMessage { return new(C2SGetRealNamePrize) })
	Processor.Register(PCK_S2CGetRealNamePrize_ID, func() network.IMessage { return new(S2CGetRealNamePrize) })
	Processor.Register(PCK_S2CReportFcm_ID, func() network.IMessage { return new(S2CReportFcm) })
	Processor.Register(PCK_C2SSetFcm_ID, func() network.IMessage { return new(C2SSetFcm) })
	Processor.Register(PCK_C2SGetLimitCloth_ID, func() network.IMessage { return new(C2SGetLimitCloth) })
	Processor.Register(PCK_S2CGetLimitCloth_ID, func() network.IMessage { return new(S2CGetLimitCloth) })
	Processor.Register(PCK_S2CGetLimitClothEnd_ID, func() network.IMessage { return new(S2CGetLimitClothEnd) })
	Processor.Register(PCK_S2CBossHillData_ID, func() network.IMessage { return new(S2CBossHillData) })
	Processor.Register(PCK_C2SBossHillFight_ID, func() network.IMessage { return new(C2SBossHillFight) })
	Processor.Register(PCK_S2CBossHillFight_ID, func() network.IMessage { return new(S2CBossHillFight) })
	Processor.Register(PCK_C2SBossHillReplace_ID, func() network.IMessage { return new(C2SBossHillReplace) })
	Processor.Register(PCK_S2CBossHillReplace_ID, func() network.IMessage { return new(S2CBossHillReplace) })
	Processor.Register(PCK_C2SBossHillOpen_ID, func() network.IMessage { return new(C2SBossHillOpen) })
	Processor.Register(PCK_S2CBossHillOpen_ID, func() network.IMessage { return new(S2CBossHillOpen) })
	Processor.Register(PCK_S2C81Info_ID, func() network.IMessage { return new(S2C81Info) })
	Processor.Register(PCK_C2SViewInstance81Data_ID, func() network.IMessage { return new(C2SViewInstance81Data) })
	Processor.Register(PCK_S2CViewInstance81Data_ID, func() network.IMessage { return new(S2CViewInstance81Data) })
	Processor.Register(PCK_C2S81Sweep_ID, func() network.IMessage { return new(C2S81Sweep) })
	Processor.Register(PCK_S2C81Sweep_ID, func() network.IMessage { return new(S2C81Sweep) })
	Processor.Register(PCK_C2S81BuyBox_ID, func() network.IMessage { return new(C2S81BuyBox) })
	Processor.Register(PCK_S2C81BuyBox_ID, func() network.IMessage { return new(S2C81BuyBox) })
	Processor.Register(PCK_C2SKingInfo_ID, func() network.IMessage { return new(C2SKingInfo) })
	Processor.Register(PCK_S2CKingInfo_ID, func() network.IMessage { return new(S2CKingInfo) })
	Processor.Register(PCK_C2SKingState_ID, func() network.IMessage { return new(C2SKingState) })
	Processor.Register(PCK_S2CKingState_ID, func() network.IMessage { return new(S2CKingState) })
	Processor.Register(PCK_C2SKingMatch_ID, func() network.IMessage { return new(C2SKingMatch) })
	Processor.Register(PCK_S2CKingMatch_ID, func() network.IMessage { return new(S2CKingMatch) })
	Processor.Register(PCK_C2SKingFight_ID, func() network.IMessage { return new(C2SKingFight) })
	Processor.Register(PCK_S2CKingFight_ID, func() network.IMessage { return new(S2CKingFight) })
	Processor.Register(PCK_C2SKingRank_ID, func() network.IMessage { return new(C2SKingRank) })
	Processor.Register(PCK_S2CKingRank_ID, func() network.IMessage { return new(S2CKingRank) })
	Processor.Register(PCK_C2SKingPlayer_ID, func() network.IMessage { return new(C2SKingPlayer) })
	Processor.Register(PCK_S2CKingPlayer_ID, func() network.IMessage { return new(S2CKingPlayer) })
	Processor.Register(PCK_C2SKingGetBuyInfo_ID, func() network.IMessage { return new(C2SKingGetBuyInfo) })
	Processor.Register(PCK_S2CKingGetBuyInfo_ID, func() network.IMessage { return new(S2CKingGetBuyInfo) })
	Processor.Register(PCK_C2SKingBuyTimes_ID, func() network.IMessage { return new(C2SKingBuyTimes) })
	Processor.Register(PCK_S2CKingBuyTimes_ID, func() network.IMessage { return new(S2CKingBuyTimes) })
	Processor.Register(PCK_C2SKingRespect_ID, func() network.IMessage { return new(C2SKingRespect) })
	Processor.Register(PCK_S2CKingRespect_ID, func() network.IMessage { return new(S2CKingRespect) })
	Processor.Register(PCK_S2CWorldBossRank5_ID, func() network.IMessage { return new(S2CWorldBossRank5) })
	Processor.Register(PCK_S2CActWorldBossSettlement_ID, func() network.IMessage { return new(S2CActWorldBossSettlement) })
	Processor.Register(PCK_K2SWorldBossKillData_ID, func() network.IMessage { return new(K2SWorldBossKillData) })
	Processor.Register(PCK_C2SGetWorldBossKillData_ID, func() network.IMessage { return new(C2SGetWorldBossKillData) })
	Processor.Register(PCK_S2CGetWorldBossKillData_ID, func() network.IMessage { return new(S2CGetWorldBossKillData) })
	Processor.Register(PCK_S2CBossRefreshTime_ID, func() network.IMessage { return new(S2CBossRefreshTime) })
	Processor.Register(PCK_C2SGetWorldBossRank_ID, func() network.IMessage { return new(C2SGetWorldBossRank) })
	Processor.Register(PCK_S2CGetWorldBossRank_ID, func() network.IMessage { return new(S2CGetWorldBossRank) })
	Processor.Register(PCK_C2SExpeditionEnemy_ID, func() network.IMessage { return new(C2SExpeditionEnemy) })
	Processor.Register(PCK_S2CExpeditionEnemy_ID, func() network.IMessage { return new(S2CExpeditionEnemy) })
	Processor.Register(PCK_C2SExpeditionTeam_ID, func() network.IMessage { return new(C2SExpeditionTeam) })
	Processor.Register(PCK_S2CExpeditionTeam_ID, func() network.IMessage { return new(S2CExpeditionTeam) })
	Processor.Register(PCK_S2CExpeditionInfo_ID, func() network.IMessage { return new(S2CExpeditionInfo) })
	Processor.Register(PCK_C2SExpeditionStartFight_ID, func() network.IMessage { return new(C2SExpeditionStartFight) })
	Processor.Register(PCK_S2CExpeditionStartFight_ID, func() network.IMessage { return new(S2CExpeditionStartFight) })
	Processor.Register(PCK_C2SExpeditionGetPrize_ID, func() network.IMessage { return new(C2SExpeditionGetPrize) })
	Processor.Register(PCK_S2CExpeditionGetPrize_ID, func() network.IMessage { return new(S2CExpeditionGetPrize) })
	Processor.Register(PCK_C2SExpeditionRelive_ID, func() network.IMessage { return new(C2SExpeditionRelive) })
	Processor.Register(PCK_S2CExpeditionRelive_ID, func() network.IMessage { return new(S2CExpeditionRelive) })
	Processor.Register(PCK_C2SExpeditionReset_ID, func() network.IMessage { return new(C2SExpeditionReset) })
	Processor.Register(PCK_S2CExpeditionReset_ID, func() network.IMessage { return new(S2CExpeditionReset) })
	Processor.Register(PCK_C2SExpeditionSaveTeam_ID, func() network.IMessage { return new(C2SExpeditionSaveTeam) })
	Processor.Register(PCK_S2CExpeditionSaveTeam_ID, func() network.IMessage { return new(S2CExpeditionSaveTeam) })
	Processor.Register(PCK_C2SExpeditionRefreshBuff_ID, func() network.IMessage { return new(C2SExpeditionRefreshBuff) })
	Processor.Register(PCK_S2CExpeditionRefreshBuff_ID, func() network.IMessage { return new(S2CExpeditionRefreshBuff) })
	Processor.Register(PCK_C2SExpeditionSweep_ID, func() network.IMessage { return new(C2SExpeditionSweep) })
	Processor.Register(PCK_S2CExpeditionSweep_ID, func() network.IMessage { return new(S2CExpeditionSweep) })
	Processor.Register(PCK_C2SWordsWear_ID, func() network.IMessage { return new(C2SWordsWear) })
	Processor.Register(PCK_S2CWordsWear_ID, func() network.IMessage { return new(S2CWordsWear) })
	Processor.Register(PCK_C2SWordsStick_ID, func() network.IMessage { return new(C2SWordsStick) })
	Processor.Register(PCK_S2CWordsStick_ID, func() network.IMessage { return new(S2CWordsStick) })
	Processor.Register(PCK_C2SInstanceSoulHallFight_ID, func() network.IMessage { return new(C2SInstanceSoulHallFight) })
	Processor.Register(PCK_S2CInstanceSoulHallFight_ID, func() network.IMessage { return new(S2CInstanceSoulHallFight) })
	Processor.Register(PCK_C2SHolyBeast_ID, func() network.IMessage { return new(C2SHolyBeast) })
	Processor.Register(PCK_S2CHolyBeast_ID, func() network.IMessage { return new(S2CHolyBeast) })
	Processor.Register(PCK_C2SHolyBeastLevel_ID, func() network.IMessage { return new(C2SHolyBeastLevel) })
	Processor.Register(PCK_S2CHolyBeastLevel_ID, func() network.IMessage { return new(S2CHolyBeastLevel) })
	Processor.Register(PCK_C2SMixHolySoul_ID, func() network.IMessage { return new(C2SMixHolySoul) })
	Processor.Register(PCK_S2CMixHolySoul_ID, func() network.IMessage { return new(S2CMixHolySoul) })
	Processor.Register(PCK_C2SBreakHolySoul_ID, func() network.IMessage { return new(C2SBreakHolySoul) })
	Processor.Register(PCK_S2CBreakHolySoul_ID, func() network.IMessage { return new(S2CBreakHolySoul) })
	Processor.Register(PCK_C2SInjectHolySoul_ID, func() network.IMessage { return new(C2SInjectHolySoul) })
	Processor.Register(PCK_S2CInjectHolySoul_ID, func() network.IMessage { return new(S2CInjectHolySoul) })
	Processor.Register(PCK_C2SOneKeyInjectHolySoul_ID, func() network.IMessage { return new(C2SOneKeyInjectHolySoul) })
	Processor.Register(PCK_S2COneKeyInjectHolySoul_ID, func() network.IMessage { return new(S2COneKeyInjectHolySoul) })
	Processor.Register(PCK_C2SOutHolySoul_ID, func() network.IMessage { return new(C2SOutHolySoul) })
	Processor.Register(PCK_S2COutHolySoul_ID, func() network.IMessage { return new(S2COutHolySoul) })
	Processor.Register(PCK_C2SGodClubSignUp_ID, func() network.IMessage { return new(C2SGodClubSignUp) })
	Processor.Register(PCK_S2CGodClubSignUp_ID, func() network.IMessage { return new(S2CGodClubSignUp) })
	Processor.Register(PCK_C2SGodClubFight_ID, func() network.IMessage { return new(C2SGodClubFight) })
	Processor.Register(PCK_S2CGodClubFight_ID, func() network.IMessage { return new(S2CGodClubFight) })
	Processor.Register(PCK_C2SGodClubFightReport_ID, func() network.IMessage { return new(C2SGodClubFightReport) })
	Processor.Register(PCK_S2CGodClubFightReport_ID, func() network.IMessage { return new(S2CGodClubFightReport) })
	Processor.Register(PCK_C2SGodClub16_ID, func() network.IMessage { return new(C2SGodClub16) })
	Processor.Register(PCK_S2CGodClub16_ID, func() network.IMessage { return new(S2CGodClub16) })
	Processor.Register(PCK_C2SGodClubStakeInfo_ID, func() network.IMessage { return new(C2SGodClubStakeInfo) })
	Processor.Register(PCK_S2CGodClubStakeInfo_ID, func() network.IMessage { return new(S2CGodClubStakeInfo) })
	Processor.Register(PCK_C2SGodClubStake_ID, func() network.IMessage { return new(C2SGodClubStake) })
	Processor.Register(PCK_S2CGodClubStake_ID, func() network.IMessage { return new(S2CGodClubStake) })
	Processor.Register(PCK_C2SGodHerbEnter_ID, func() network.IMessage { return new(C2SGodHerbEnter) })
	Processor.Register(PCK_S2CGodHerbEnter_ID, func() network.IMessage { return new(S2CGodHerbEnter) })
	Processor.Register(PCK_C2SGodHerbRefresh_ID, func() network.IMessage { return new(C2SGodHerbRefresh) })
	Processor.Register(PCK_S2CGodHerbRefresh_ID, func() network.IMessage { return new(S2CGodHerbRefresh) })
	Processor.Register(PCK_C2SGodHerbCollect_ID, func() network.IMessage { return new(C2SGodHerbCollect) })
	Processor.Register(PCK_S2CGodHerbCollect_ID, func() network.IMessage { return new(S2CGodHerbCollect) })
	Processor.Register(PCK_C2SGetGodHerbLog_ID, func() network.IMessage { return new(C2SGetGodHerbLog) })
	Processor.Register(PCK_S2CGetGodHerbLog_ID, func() network.IMessage { return new(S2CGetGodHerbLog) })
	Processor.Register(PCK_S2CRobberyInfo_ID, func() network.IMessage { return new(S2CRobberyInfo) })
	Processor.Register(PCK_C2SGetRobberyData_ID, func() network.IMessage { return new(C2SGetRobberyData) })
	Processor.Register(PCK_S2CGetRobberyData_ID, func() network.IMessage { return new(S2CGetRobberyData) })
	Processor.Register(PCK_C2SRobbery_ID, func() network.IMessage { return new(C2SRobbery) })
	Processor.Register(PCK_S2CRobbery_ID, func() network.IMessage { return new(S2CRobbery) })
	Processor.Register(PCK_C2SRobberyLuck_ID, func() network.IMessage { return new(C2SRobberyLuck) })
	Processor.Register(PCK_S2CRobberyLuck_ID, func() network.IMessage { return new(S2CRobberyLuck) })
	Processor.Register(PCK_C2SGetRobberyLog_ID, func() network.IMessage { return new(C2SGetRobberyLog) })
	Processor.Register(PCK_S2CGetRobberyLog_ID, func() network.IMessage { return new(S2CGetRobberyLog) })
	Processor.Register(PCK_S2CFairyMemoirTrigger_ID, func() network.IMessage { return new(S2CFairyMemoirTrigger) })
	Processor.Register(PCK_S2CFairyAchievementChange_ID, func() network.IMessage { return new(S2CFairyAchievementChange) })
	Processor.Register(PCK_C2SFairyCharacterData_ID, func() network.IMessage { return new(C2SFairyCharacterData) })
	Processor.Register(PCK_S2CFairyCharacterData_ID, func() network.IMessage { return new(S2CFairyCharacterData) })
	Processor.Register(PCK_C2SFairyCharacterFight_ID, func() network.IMessage { return new(C2SFairyCharacterFight) })
	Processor.Register(PCK_S2CFairyCharacterFight_ID, func() network.IMessage { return new(S2CFairyCharacterFight) })
	Processor.Register(PCK_C2SFairyMemoirPrize_ID, func() network.IMessage { return new(C2SFairyMemoirPrize) })
	Processor.Register(PCK_S2CFairyMemoirPrize_ID, func() network.IMessage { return new(S2CFairyMemoirPrize) })
	Processor.Register(PCK_C2SFairyCharacterSweep_ID, func() network.IMessage { return new(C2SFairyCharacterSweep) })
	Processor.Register(PCK_S2CFairyCharacterSweep_ID, func() network.IMessage { return new(S2CFairyCharacterSweep) })
	Processor.Register(PCK_C2SFairyAchieventPrize_ID, func() network.IMessage { return new(C2SFairyAchieventPrize) })
	Processor.Register(PCK_S2CFairyAchieventPrize_ID, func() network.IMessage { return new(S2CFairyAchieventPrize) })
	Processor.Register(PCK_S2CClientDb_ID, func() network.IMessage { return new(S2CClientDb) })
	Processor.Register(PCK_C2SClinetSet_ID, func() network.IMessage { return new(C2SClinetSet) })
	Processor.Register(PCK_S2CClinetSet_ID, func() network.IMessage { return new(S2CClinetSet) })
	Processor.Register(PCK_S2CPassCheckInfo_ID, func() network.IMessage { return new(S2CPassCheckInfo) })
	Processor.Register(PCK_C2SPassCheckGetPrize_ID, func() network.IMessage { return new(C2SPassCheckGetPrize) })
	Processor.Register(PCK_S2CPassCheckGetPrize_ID, func() network.IMessage { return new(S2CPassCheckGetPrize) })
	Processor.Register(PCK_C2SPassCheckFast_ID, func() network.IMessage { return new(C2SPassCheckFast) })
	Processor.Register(PCK_S2CPassCheckFast_ID, func() network.IMessage { return new(S2CPassCheckFast) })
	Processor.Register(PCK_C2SGetZF_ID, func() network.IMessage { return new(C2SGetZF) })
	Processor.Register(PCK_S2CGetZF_ID, func() network.IMessage { return new(S2CGetZF) })
	Processor.Register(PCK_C2SZFPetUp_ID, func() network.IMessage { return new(C2SZFPetUp) })
	Processor.Register(PCK_S2CZFPetUp_ID, func() network.IMessage { return new(S2CZFPetUp) })
	Processor.Register(PCK_C2SZFState_ID, func() network.IMessage { return new(C2SZFState) })
	Processor.Register(PCK_S2CZFState_ID, func() network.IMessage { return new(S2CZFState) })
	Processor.Register(PCK_C2SZFUnlock_ID, func() network.IMessage { return new(C2SZFUnlock) })
	Processor.Register(PCK_S2CZFUnlock_ID, func() network.IMessage { return new(S2CZFUnlock) })
	Processor.Register(PCK_K2SNewEvent_ID, func() network.IMessage { return new(K2SNewEvent) })
	Processor.Register(PCK_C2SFuncOpen_ID, func() network.IMessage { return new(C2SFuncOpen) })
	Processor.Register(PCK_S2CFuncOpen_ID, func() network.IMessage { return new(S2CFuncOpen) })
	Processor.Register(PCK_C2SRobotZF_ID, func() network.IMessage { return new(C2SRobotZF) })
	Processor.Register(PCK_S2CUserPhotoBook_ID, func() network.IMessage { return new(S2CUserPhotoBook) })
	Processor.Register(PCK_C2SPhotoBookLevelUp_ID, func() network.IMessage { return new(C2SPhotoBookLevelUp) })
	Processor.Register(PCK_S2CPhotoBookLevelUp_ID, func() network.IMessage { return new(S2CPhotoBookLevelUp) })
	Processor.Register(PCK_C2SPhotoBookDel_ID, func() network.IMessage { return new(C2SPhotoBookDel) })
	Processor.Register(PCK_S2CPhotoBookDel_ID, func() network.IMessage { return new(S2CPhotoBookDel) })
	Processor.Register(PCK_C2SBuyChargeGift_ID, func() network.IMessage { return new(C2SBuyChargeGift) })
	Processor.Register(PCK_S2CBuyChargeGift_ID, func() network.IMessage { return new(S2CBuyChargeGift) })
	Processor.Register(PCK_S2CAllChargeGift_ID, func() network.IMessage { return new(S2CAllChargeGift) })
	Processor.Register(PCK_S2CVipService_ID, func() network.IMessage { return new(S2CVipService) })
	Processor.Register(PCK_C2SNewPlayerZF_ID, func() network.IMessage { return new(C2SNewPlayerZF) })
	Processor.Register(PCK_C2SGetStateDemons_ID, func() network.IMessage { return new(C2SGetStateDemons) })
	Processor.Register(PCK_S2CGetStateDemons_ID, func() network.IMessage { return new(S2CGetStateDemons) })
	Processor.Register(PCK_C2SStateBreach_ID, func() network.IMessage { return new(C2SStateBreach) })
	Processor.Register(PCK_S2CStateBreach_ID, func() network.IMessage { return new(S2CStateBreach) })
	Processor.Register(PCK_C2SStateUsePill_ID, func() network.IMessage { return new(C2SStateUsePill) })
	Processor.Register(PCK_S2CStateUsePill_ID, func() network.IMessage { return new(S2CStateUsePill) })
	Processor.Register(PCK_C2SBuyStateGift_ID, func() network.IMessage { return new(C2SBuyStateGift) })
	Processor.Register(PCK_S2CBuyStateGift_ID, func() network.IMessage { return new(S2CBuyStateGift) })
	Processor.Register(PCK_C2SEquipStarInherit_ID, func() network.IMessage { return new(C2SEquipStarInherit) })
	Processor.Register(PCK_S2CEquipStarInherit_ID, func() network.IMessage { return new(S2CEquipStarInherit) })
	Processor.Register(PCK_C2SEnterInstance_ID, func() network.IMessage { return new(C2SEnterInstance) })
	Processor.Register(PCK_S2CEnterInstance_ID, func() network.IMessage { return new(S2CEnterInstance) })
	Processor.Register(PCK_S2CInstanceData_ID, func() network.IMessage { return new(S2CInstanceData) })
	Processor.Register(PCK_S2CInstanceSettlement_ID, func() network.IMessage { return new(S2CInstanceSettlement) })
	Processor.Register(PCK_S2CInstanceExpData_ID, func() network.IMessage { return new(S2CInstanceExpData) })
	Processor.Register(PCK_C2SInstanceExpBuyTimes_ID, func() network.IMessage { return new(C2SInstanceExpBuyTimes) })
	Processor.Register(PCK_S2CInstanceExpBuyTimes_ID, func() network.IMessage { return new(S2CInstanceExpBuyTimes) })
	Processor.Register(PCK_C2SGemUpgrade_ID, func() network.IMessage { return new(C2SGemUpgrade) })
	Processor.Register(PCK_S2CGemUpgrade_ID, func() network.IMessage { return new(S2CGemUpgrade) })
	Processor.Register(PCK_C2SGemAutoEquip_ID, func() network.IMessage { return new(C2SGemAutoEquip) })
	Processor.Register(PCK_S2CGemAutoEquip_ID, func() network.IMessage { return new(S2CGemAutoEquip) })
	Processor.Register(PCK_C2SGemAutoUpgradeEquip_ID, func() network.IMessage { return new(C2SGemAutoUpgradeEquip) })
	Processor.Register(PCK_S2CGemAutoUpgradeEquip_ID, func() network.IMessage { return new(S2CGemAutoUpgradeEquip) })
	Processor.Register(PCK_S2CBecomeGodTask_ID, func() network.IMessage { return new(S2CBecomeGodTask) })
	Processor.Register(PCK_C2SGetBecomeGodTaskPrize_ID, func() network.IMessage { return new(C2SGetBecomeGodTaskPrize) })
	Processor.Register(PCK_S2CGetBecomeGodTaskPrize_ID, func() network.IMessage { return new(S2CGetBecomeGodTaskPrize) })
	Processor.Register(PCK_C2SEquipSpecial_ID, func() network.IMessage { return new(C2SEquipSpecial) })
	Processor.Register(PCK_S2CEquipSpecial_ID, func() network.IMessage { return new(S2CEquipSpecial) })
	Processor.Register(PCK_C2SGetEquipSpecialReward_ID, func() network.IMessage { return new(C2SGetEquipSpecialReward) })
	Processor.Register(PCK_S2CGetEquipSpecialReward_ID, func() network.IMessage { return new(S2CGetEquipSpecialReward) })
	Processor.Register(PCK_S2CInstanceTongTianInfo_ID, func() network.IMessage { return new(S2CInstanceTongTianInfo) })
	Processor.Register(PCK_C2SInstanceTongTianFight_ID, func() network.IMessage { return new(C2SInstanceTongTianFight) })
	Processor.Register(PCK_S2CInstanceTongTianFight_ID, func() network.IMessage { return new(S2CInstanceTongTianFight) })
	Processor.Register(PCK_C2SInstanceTongTianSweep_ID, func() network.IMessage { return new(C2SInstanceTongTianSweep) })
	Processor.Register(PCK_S2CInstanceTongTianSweep_ID, func() network.IMessage { return new(S2CInstanceTongTianSweep) })
	Processor.Register(PCK_C2SGetInstanceTongTianBox_ID, func() network.IMessage { return new(C2SGetInstanceTongTianBox) })
	Processor.Register(PCK_S2CGetInstanceTongTianBox_ID, func() network.IMessage { return new(S2CGetInstanceTongTianBox) })
	Processor.Register(PCK_C2SGetInstanceTongTianProcessBox_ID, func() network.IMessage { return new(C2SGetInstanceTongTianProcessBox) })
	Processor.Register(PCK_S2CGetInstanceTongTianProcessBox_ID, func() network.IMessage { return new(S2CGetInstanceTongTianProcessBox) })
	Processor.Register(PCK_C2SInstanceTongTianReset_ID, func() network.IMessage { return new(C2SInstanceTongTianReset) })
	Processor.Register(PCK_S2CInstanceTongTianReset_ID, func() network.IMessage { return new(S2CInstanceTongTianReset) })
	Processor.Register(PCK_C2SHorseEquipStarUp_ID, func() network.IMessage { return new(C2SHorseEquipStarUp) })
	Processor.Register(PCK_S2CHorseEquipStarUp_ID, func() network.IMessage { return new(S2CHorseEquipStarUp) })
	Processor.Register(PCK_C2SStartLockTower_ID, func() network.IMessage { return new(C2SStartLockTower) })
	Processor.Register(PCK_S2CStartLockTower_ID, func() network.IMessage { return new(S2CStartLockTower) })
	Processor.Register(PCK_C2SGirlHillStart_ID, func() network.IMessage { return new(C2SGirlHillStart) })
	Processor.Register(PCK_S2CGirlHillStart_ID, func() network.IMessage { return new(S2CGirlHillStart) })
	Processor.Register(PCK_C2SGirlHillInspire_ID, func() network.IMessage { return new(C2SGirlHillInspire) })
	Processor.Register(PCK_S2CGirlHillInspire_ID, func() network.IMessage { return new(S2CGirlHillInspire) })
	Processor.Register(PCK_C2SGirlHillBuyTimes_ID, func() network.IMessage { return new(C2SGirlHillBuyTimes) })
	Processor.Register(PCK_S2CGirlHillBuyTimes_ID, func() network.IMessage { return new(S2CGirlHillBuyTimes) })
	Processor.Register(PCK_C2SZhenFaActive_ID, func() network.IMessage { return new(C2SZhenFaActive) })
	Processor.Register(PCK_S2CZhenFaActive_ID, func() network.IMessage { return new(S2CZhenFaActive) })
	Processor.Register(PCK_C2SZhenFaUpgrade_ID, func() network.IMessage { return new(C2SZhenFaUpgrade) })
	Processor.Register(PCK_S2CZhenFaUpgrade_ID, func() network.IMessage { return new(S2CZhenFaUpgrade) })
	Processor.Register(PCK_C2SZhenFaEquip_ID, func() network.IMessage { return new(C2SZhenFaEquip) })
	Processor.Register(PCK_S2CZhenFaEquip_ID, func() network.IMessage { return new(S2CZhenFaEquip) })
	Processor.Register(PCK_C2SJiBanActive_ID, func() network.IMessage { return new(C2SJiBanActive) })
	Processor.Register(PCK_S2CJiBanActive_ID, func() network.IMessage { return new(S2CJiBanActive) })
	Processor.Register(PCK_C2SJiBanEquip_ID, func() network.IMessage { return new(C2SJiBanEquip) })
	Processor.Register(PCK_S2CJiBanEquip_ID, func() network.IMessage { return new(S2CJiBanEquip) })
	Processor.Register(PCK_S2CPartnerSportsData_ID, func() network.IMessage { return new(S2CPartnerSportsData) })
	Processor.Register(PCK_C2SPartnerSportsFight_ID, func() network.IMessage { return new(C2SPartnerSportsFight) })
	Processor.Register(PCK_S2CPartnerSportsFight_ID, func() network.IMessage { return new(S2CPartnerSportsFight) })
	Processor.Register(PCK_C2SPartnerSportsRefresh_ID, func() network.IMessage { return new(C2SPartnerSportsRefresh) })
	Processor.Register(PCK_S2CPartnerSportsRefresh_ID, func() network.IMessage { return new(S2CPartnerSportsRefresh) })
	Processor.Register(PCK_C2SPartnerSportsBuyFightTimes_ID, func() network.IMessage { return new(C2SPartnerSportsBuyFightTimes) })
	Processor.Register(PCK_S2CPartnerSportsBuyFightTimes_ID, func() network.IMessage { return new(S2CPartnerSportsBuyFightTimes) })
	Processor.Register(PCK_C2SPartnerSportsSetFightPartner_ID, func() network.IMessage { return new(C2SPartnerSportsSetFightPartner) })
	Processor.Register(PCK_S2CPartnerSportsSetFightPartner_ID, func() network.IMessage { return new(S2CPartnerSportsSetFightPartner) })
	Processor.Register(PCK_C2SPartnerSportsGetWinPrize_ID, func() network.IMessage { return new(C2SPartnerSportsGetWinPrize) })
	Processor.Register(PCK_S2CPartnerSportsGetWinPrize_ID, func() network.IMessage { return new(S2CPartnerSportsGetWinPrize) })
	Processor.Register(PCK_S2KPartnerSportsScore_ID, func() network.IMessage { return new(S2KPartnerSportsScore) })
	Processor.Register(PCK_C2STianShuUnlock_ID, func() network.IMessage { return new(C2STianShuUnlock) })
	Processor.Register(PCK_S2CTianShuUnlock_ID, func() network.IMessage { return new(S2CTianShuUnlock) })
	Processor.Register(PCK_C2STianShuActive_ID, func() network.IMessage { return new(C2STianShuActive) })
	Processor.Register(PCK_S2CTianShuActive_ID, func() network.IMessage { return new(S2CTianShuActive) })
}

const PCK_S2SServerInfo_ID = 9996 //
//
type S2SServerInfo struct {
	//
	ServerId int64 `protobuf:"varint,1,opt,name=ServerId,proto3" json:"ServerId,omitempty"`
	//
	ServerType int64 `protobuf:"varint,2,opt,name=ServerType,proto3" json:"ServerType,omitempty"`
}

func (m *S2SServerInfo) Reset()         { *m = S2SServerInfo{} }
func (m *S2SServerInfo) String() string { return proto.CompactTextString(m) }
func (*S2SServerInfo) ProtoMessage()    {}
func (m *S2SServerInfo) GetId() uint16  { return PCK_S2SServerInfo_ID }

const PCK_S2SRequestMsg_ID = 9999 //
//
type S2SRequestMsg struct {
	//
	RequestID int32 `protobuf:"varint,1,opt,name=RequestID,proto3" json:"RequestID,omitempty"`
	//
	CallType int32 `protobuf:"varint,2,opt,name=CallType,proto3" json:"CallType,omitempty"`
	//
	ServerId int64 `protobuf:"varint,3,opt,name=ServerId,proto3" json:"ServerId,omitempty"`
	//
	UserId int64 `protobuf:"varint,5,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	Pck []byte `protobuf:"bytes,4,opt,name=Pck,proto3" json:"Pck,omitempty"`
}

func (m *S2SRequestMsg) Reset()         { *m = S2SRequestMsg{} }
func (m *S2SRequestMsg) String() string { return proto.CompactTextString(m) }
func (*S2SRequestMsg) ProtoMessage()    {}
func (m *S2SRequestMsg) GetId() uint16  { return PCK_S2SRequestMsg_ID }

const PCK_S2SResponseMsg_ID = 9998 //
//
type S2SResponseMsg struct {
	//
	RequestID int32 `protobuf:"varint,1,opt,name=RequestID,proto3" json:"RequestID,omitempty"`
	//
	Tag int32 `protobuf:"varint,2,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Pck []byte `protobuf:"bytes,3,opt,name=Pck,proto3" json:"Pck,omitempty"`
}

func (m *S2SResponseMsg) Reset()         { *m = S2SResponseMsg{} }
func (m *S2SResponseMsg) String() string { return proto.CompactTextString(m) }
func (*S2SResponseMsg) ProtoMessage()    {}
func (m *S2SResponseMsg) GetId() uint16  { return PCK_S2SResponseMsg_ID }

const PCK_S2SServerPlayerPck_ID = 10000 //
//
type S2SServerPlayerPck struct {
	//
	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	Pck []byte `protobuf:"bytes,2,opt,name=Pck,proto3" json:"Pck,omitempty"`
}

func (m *S2SServerPlayerPck) Reset()         { *m = S2SServerPlayerPck{} }
func (m *S2SServerPlayerPck) String() string { return proto.CompactTextString(m) }
func (*S2SServerPlayerPck) ProtoMessage()    {}
func (m *S2SServerPlayerPck) GetId() uint16  { return PCK_S2SServerPlayerPck_ID }

const PCK_S2SServerPlayerHandle_ID = 9997 //
//
type S2SServerPlayerHandle struct {
	//
	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	Pck []byte `protobuf:"bytes,2,opt,name=Pck,proto3" json:"Pck,omitempty"`
}

func (m *S2SServerPlayerHandle) Reset()         { *m = S2SServerPlayerHandle{} }
func (m *S2SServerPlayerHandle) String() string { return proto.CompactTextString(m) }
func (*S2SServerPlayerHandle) ProtoMessage()    {}
func (m *S2SServerPlayerHandle) GetId() uint16  { return PCK_S2SServerPlayerHandle_ID }

const PCK_S2KPlayerData_ID = 8011 //
//
type S2KPlayerData struct {
	//
	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	Player *RoleDbInfo `protobuf:"bytes,2,opt,name=Player,proto3" json:"Player,omitempty"`
}

func (m *S2KPlayerData) Reset()         { *m = S2KPlayerData{} }
func (m *S2KPlayerData) String() string { return proto.CompactTextString(m) }
func (*S2KPlayerData) ProtoMessage()    {}
func (m *S2KPlayerData) GetId() uint16  { return PCK_S2KPlayerData_ID }

const PCK_Ping_ID = 22 //
//
type Ping struct {
	//
	T int64 `protobuf:"varint,1,opt,name=T,proto3" json:"T,omitempty"`
}

func (m *Ping) Reset()         { *m = Ping{} }
func (m *Ping) String() string { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()    {}
func (m *Ping) GetId() uint16  { return PCK_Ping_ID }

const PCK_Pong_ID = 23 //
//
type Pong struct {
	//
	T int64 `protobuf:"varint,1,opt,name=T,proto3" json:"T,omitempty"`
}

func (m *Pong) Reset()         { *m = Pong{} }
func (m *Pong) String() string { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()    {}
func (m *Pong) GetId() uint16  { return PCK_Pong_ID }

const PCK_S2KCrossPing_ID = 38 //
//
type S2KCrossPing struct {
	//
	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (m *S2KCrossPing) Reset()         { *m = S2KCrossPing{} }
func (m *S2KCrossPing) String() string { return proto.CompactTextString(m) }
func (*S2KCrossPing) ProtoMessage()    {}
func (m *S2KCrossPing) GetId() uint16  { return PCK_S2KCrossPing_ID }

//
type CommonType struct {
	//
	A int32 `protobuf:"varint,1,opt,name=A,proto3" json:"A,omitempty"`
	//
	B int32 `protobuf:"varint,2,opt,name=B,proto3" json:"B,omitempty"`
}

func (m *CommonType) Reset()         { *m = CommonType{} }
func (m *CommonType) String() string { return proto.CompactTextString(m) }
func (*CommonType) ProtoMessage()    {}

const PCK_S2CServerTime_ID = 1001 //
//
type S2CServerTime struct {
	//
	T int64 `protobuf:"varint,1,opt,name=T,proto3" json:"T,omitempty"`
}

func (m *S2CServerTime) Reset()         { *m = S2CServerTime{} }
func (m *S2CServerTime) String() string { return proto.CompactTextString(m) }
func (*S2CServerTime) ProtoMessage()    {}
func (m *S2CServerTime) GetId() uint16  { return PCK_S2CServerTime_ID }

const PCK_S2CServerAge_ID = 1002 //
//
type S2CServerAge struct {
	//
	T int32 `protobuf:"varint,1,opt,name=T,proto3" json:"T,omitempty"`
}

func (m *S2CServerAge) Reset()         { *m = S2CServerAge{} }
func (m *S2CServerAge) String() string { return proto.CompactTextString(m) }
func (*S2CServerAge) ProtoMessage()    {}
func (m *S2CServerAge) GetId() uint16  { return PCK_S2CServerAge_ID }

//
type IntAttr struct {
	//
	K int32 `protobuf:"varint,1,opt,name=K,proto3" json:"K,omitempty"`
	//
	V int64 `protobuf:"varint,2,opt,name=V,proto3" json:"V,omitempty"`
}

func (m *IntAttr) Reset()         { *m = IntAttr{} }
func (m *IntAttr) String() string { return proto.CompactTextString(m) }
func (*IntAttr) ProtoMessage()    {}

//
type StrAttr struct {
	//
	K int32 `protobuf:"varint,1,opt,name=K,proto3" json:"K,omitempty"`
	//
	V string `protobuf:"bytes,2,opt,name=V,proto3" json:"V,omitempty"`
}

func (m *StrAttr) Reset()         { *m = StrAttr{} }
func (m *StrAttr) String() string { return proto.CompactTextString(m) }
func (*StrAttr) ProtoMessage()    {}

//
type ItemInfo struct {
	//
	ItemId int32 `protobuf:"varint,1,opt,name=ItemId,proto3" json:"ItemId,omitempty"`
	//
	ItemNum int64 `protobuf:"varint,2,opt,name=ItemNum,proto3" json:"ItemNum,omitempty"`
	//
	Bind int32 `protobuf:"varint,3,opt,name=Bind,proto3" json:"Bind,omitempty"`
	//
	ItemValid int32 `protobuf:"varint,4,opt,name=ItemValid,proto3" json:"ItemValid,omitempty"`
	//
	DSL int32 `protobuf:"varint,5,opt,name=DSL,proto3" json:"DSL,omitempty"`
	//
	DSS int32 `protobuf:"varint,6,opt,name=DSS,proto3" json:"DSS,omitempty"`
}

func (m *ItemInfo) Reset()         { *m = ItemInfo{} }
func (m *ItemInfo) String() string { return proto.CompactTextString(m) }
func (*ItemInfo) ProtoMessage()    {}

//
type ItemData struct {
	//
	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Pos int32 `protobuf:"varint,2,opt,name=Pos,proto3" json:"Pos,omitempty"`
	//
	IId int32 `protobuf:"varint,3,opt,name=IId,proto3" json:"IId,omitempty"`
	//
	N int64 `protobuf:"varint,4,opt,name=N,proto3" json:"N,omitempty"`
	//
	B int32 `protobuf:"varint,5,opt,name=B,proto3" json:"B,omitempty"`
	//
	A int32 `protobuf:"varint,6,opt,name=A,proto3" json:"A,omitempty"`
	//
	V int64 `protobuf:"varint,7,opt,name=V,proto3" json:"V,omitempty"`
	//
	DSL int32 `protobuf:"varint,8,opt,name=DSL,proto3" json:"DSL,omitempty"`
	//
	DSS int32 `protobuf:"varint,9,opt,name=DSS,proto3" json:"DSS,omitempty"`
	//
	DSPos int32 `protobuf:"varint,10,opt,name=DSPos,proto3" json:"DSPos,omitempty"`
	//
	Attr []*IntAttr `protobuf:"bytes,11,rep,name=Attr,proto3" json:"Attr,omitempty"`
	//
	Star int32 `protobuf:"varint,12,opt,name=Star,proto3" json:"Star,omitempty"`
	//
	PartnerId int32 `protobuf:"varint,13,opt,name=PartnerId,proto3" json:"PartnerId,omitempty"`
}

func (m *ItemData) Reset()         { *m = ItemData{} }
func (m *ItemData) String() string { return proto.CompactTextString(m) }
func (*ItemData) ProtoMessage()    {}

//
type ItemNum struct {
	//
	IId int32 `protobuf:"varint,3,opt,name=IId,proto3" json:"IId,omitempty"`
	//
	N int64 `protobuf:"varint,4,opt,name=N,proto3" json:"N,omitempty"`
}

func (m *ItemNum) Reset()         { *m = ItemNum{} }
func (m *ItemNum) String() string { return proto.CompactTextString(m) }
func (*ItemNum) ProtoMessage()    {}

//
type Pos struct {
	//
	X int64 `protobuf:"varint,1,opt,name=X,proto3" json:"X,omitempty"`
	//
	Y int64 `protobuf:"varint,2,opt,name=Y,proto3" json:"Y,omitempty"`
}

func (m *Pos) Reset()         { *m = Pos{} }
func (m *Pos) String() string { return proto.CompactTextString(m) }
func (*Pos) ProtoMessage()    {}

//
type LimitTimesData struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	P int32 `protobuf:"varint,2,opt,name=P,proto3" json:"P,omitempty"`
	//
	Use int32 `protobuf:"varint,3,opt,name=Use,proto3" json:"Use,omitempty"`
}

func (m *LimitTimesData) Reset()         { *m = LimitTimesData{} }
func (m *LimitTimesData) String() string { return proto.CompactTextString(m) }
func (*LimitTimesData) ProtoMessage()    {}

//
type RoleLimitTimes struct {
	//
	Items []*LimitTimesData `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty"`
}

func (m *RoleLimitTimes) Reset()         { *m = RoleLimitTimes{} }
func (m *RoleLimitTimes) String() string { return proto.CompactTextString(m) }
func (*RoleLimitTimes) ProtoMessage()    {}

//
type RoleDbInfo struct {
	//
	UserId int32 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	Nick string `protobuf:"bytes,2,opt,name=Nick,proto3" json:"Nick,omitempty"`
	//
	RoleId int32 `protobuf:"varint,3,opt,name=RoleId,proto3" json:"RoleId,omitempty"`
	//
	Sex int32 `protobuf:"varint,4,opt,name=Sex,proto3" json:"Sex,omitempty"`
	//
	intAttr []*IntAttr `protobuf:"bytes,5,rep,name=intAttr,proto3" json:"intAttr,omitempty"`
	//
	strAttr []*StrAttr `protobuf:"bytes,6,rep,name=strAttr,proto3" json:"strAttr,omitempty"`
	//
	ActiveSkills []*Skill `protobuf:"bytes,21,rep,name=ActiveSkills,proto3" json:"ActiveSkills,omitempty"`
	//
	Skins []*Skin `protobuf:"bytes,22,rep,name=Skins,proto3" json:"Skins,omitempty"`
	//
	Bag *RoleBag `protobuf:"bytes,7,opt,name=Bag,proto3" json:"Bag,omitempty"`
	//
	Grade *RoleGrade `protobuf:"bytes,9,opt,name=Grade,proto3" json:"Grade,omitempty"`
	//
	LimitTimes *RoleLimitTimes `protobuf:"bytes,10,opt,name=LimitTimes,proto3" json:"LimitTimes,omitempty"`
	//
	Counters []*Counter `protobuf:"bytes,11,rep,name=Counters,proto3" json:"Counters,omitempty"`
	//
	Tasks []*Task `protobuf:"bytes,12,rep,name=Tasks,proto3" json:"Tasks,omitempty"`
	//
	Mails *RoleMail `protobuf:"bytes,13,opt,name=Mails,proto3" json:"Mails,omitempty"`
	//
	Pets []*Pet `protobuf:"bytes,14,rep,name=Pets,proto3" json:"Pets,omitempty"`
	//
	PetAs []*PetA `protobuf:"bytes,15,rep,name=PetAs,proto3" json:"PetAs,omitempty"`
	//
	Kids []*Kid `protobuf:"bytes,17,rep,name=Kids,proto3" json:"Kids,omitempty"`
	//
	SkyGods []*SkyGod `protobuf:"bytes,18,rep,name=SkyGods,proto3" json:"SkyGods,omitempty"`
	//
	Therions []*Therion `protobuf:"bytes,19,rep,name=Therions,proto3" json:"Therions,omitempty"`
	//
	PartnerSuit []int32 `protobuf:"varint,20,rep,name=PartnerSuit,proto3" json:"PartnerSuit,omitempty"`
	//
	VIPPrize []int32 `protobuf:"varint,23,rep,name=VIPPrize,proto3" json:"VIPPrize,omitempty"`
	//
	MarryInfo *DbMarryInfo `protobuf:"bytes,24,opt,name=MarryInfo,proto3" json:"MarryInfo,omitempty"`
	//
	Precious []*Precious `protobuf:"bytes,25,rep,name=Precious,proto3" json:"Precious,omitempty"`
	//
	PlayerWestExp *DbPlayerWestExp `protobuf:"bytes,26,opt,name=PlayerWestExp,proto3" json:"PlayerWestExp,omitempty"`
	//
	GangExc *GangExchange `protobuf:"bytes,27,opt,name=GangExc,proto3" json:"GangExc,omitempty"`
	//
	SendShopItems []int32 `protobuf:"varint,28,rep,name=SendShopItems,proto3" json:"SendShopItems,omitempty"`
	//
	DragonLogs []*DragonLogItem `protobuf:"bytes,29,rep,name=DragonLogs,proto3" json:"DragonLogs,omitempty"`
	//
	StagePrize []*StagePrizeInfo `protobuf:"bytes,30,rep,name=StagePrize,proto3" json:"StagePrize,omitempty"`
	//
	Rprs *RedPkgRecords `protobuf:"bytes,31,opt,name=Rprs,proto3" json:"Rprs,omitempty"`
	//
	SMMonster []*SMMonsterInfo `protobuf:"bytes,32,rep,name=SMMonster,proto3" json:"SMMonster,omitempty"`
	//
	CreateCpsId int32 `protobuf:"varint,33,opt,name=CreateCpsId,proto3" json:"CreateCpsId,omitempty"`
	//
	NewPrize []int32 `protobuf:"varint,34,rep,name=NewPrize,proto3" json:"NewPrize,omitempty"`
	//
	AchieveShow []int32 `protobuf:"varint,35,rep,name=AchieveShow,proto3" json:"AchieveShow,omitempty"`
	//
	FriendInfo *RoleFriend `protobuf:"bytes,36,opt,name=FriendInfo,proto3" json:"FriendInfo,omitempty"`
	//
	MasterInfo *DbMasterData `protobuf:"bytes,37,opt,name=MasterInfo,proto3" json:"MasterInfo,omitempty"`
	//
	DailyTask []*DBDailyTask `protobuf:"bytes,38,rep,name=DailyTask,proto3" json:"DailyTask,omitempty"`
	//
	JJCChange []*JJCChange `protobuf:"bytes,39,rep,name=JJCChange,proto3" json:"JJCChange,omitempty"`
	//
	Devils []*Devil `protobuf:"bytes,40,rep,name=Devils,proto3" json:"Devils,omitempty"`
	//
	Beauty *Beauty `protobuf:"bytes,41,opt,name=Beauty,proto3" json:"Beauty,omitempty"`
	//
	BossHill *BossHillData `protobuf:"bytes,42,opt,name=BossHill,proto3" json:"BossHill,omitempty"`
	//
	Player81 []*Player81Data `protobuf:"bytes,43,rep,name=Player81,proto3" json:"Player81,omitempty"`
	//
	HTI *S2CHistoryTaskInfo `protobuf:"bytes,44,opt,name=HTI,proto3" json:"HTI,omitempty"`
	//
	TreasureData *TreasureData `protobuf:"bytes,45,opt,name=TreasureData,proto3" json:"TreasureData,omitempty"`
	//
	PTasks []*FinishTask `protobuf:"bytes,46,rep,name=PTasks,proto3" json:"PTasks,omitempty"`
	//
	FTasks []*FinishTask `protobuf:"bytes,47,rep,name=FTasks,proto3" json:"FTasks,omitempty"`
	//
	RedPckLog []*RedPckLog `protobuf:"bytes,49,rep,name=RedPckLog,proto3" json:"RedPckLog,omitempty"`
	//
	HolyBeasts []*HolyBeast `protobuf:"bytes,50,rep,name=HolyBeasts,proto3" json:"HolyBeasts,omitempty"`
	//
	ExpeditionData *ExpeditionData `protobuf:"bytes,52,opt,name=ExpeditionData,proto3" json:"ExpeditionData,omitempty"`
	//
	Toys []*Toy `protobuf:"bytes,53,rep,name=Toys,proto3" json:"Toys,omitempty"`
	//
	RobberyItem []*RobberyItem `protobuf:"bytes,54,rep,name=RobberyItem,proto3" json:"RobberyItem,omitempty"`
	//
	FairySheets []*FairySheet `protobuf:"bytes,55,rep,name=FairySheets,proto3" json:"FairySheets,omitempty"`
	//
	FairyAchievements []*FairyAchievement `protobuf:"bytes,56,rep,name=FairyAchievements,proto3" json:"FairyAchievements,omitempty"`
	//
	RobberData []*RobberData `protobuf:"bytes,57,rep,name=RobberData,proto3" json:"RobberData,omitempty"`
	//
	ActPlayerData []*ActPlayerData `protobuf:"bytes,58,rep,name=ActPlayerData,proto3" json:"ActPlayerData,omitempty"`
	//
	ActExit []int32 `protobuf:"varint,59,rep,name=ActExit,proto3" json:"ActExit,omitempty"`
	//
	RobberyLogs []*RobberyLogItem `protobuf:"bytes,60,rep,name=RobberyLogs,proto3" json:"RobberyLogs,omitempty"`
	//
	JJCPrizeTime int32 `protobuf:"varint,61,opt,name=JJCPrizeTime,proto3" json:"JJCPrizeTime,omitempty"`
	//
	VIPGift []int32 `protobuf:"varint,62,rep,name=VIPGift,proto3" json:"VIPGift,omitempty"`
	//
	DevilPos []*DevilPos `protobuf:"bytes,63,rep,name=DevilPos,proto3" json:"DevilPos,omitempty"`
	//
	PlayerDragon *PlayerDragon `protobuf:"bytes,64,opt,name=PlayerDragon,proto3" json:"PlayerDragon,omitempty"`
	//
	ClientDb *S2CClientDb `protobuf:"bytes,65,opt,name=ClientDb,proto3" json:"ClientDb,omitempty"`
	//
	FairyDareData []*S2CPlayerDareBossInfo `protobuf:"bytes,66,rep,name=FairyDareData,proto3" json:"FairyDareData,omitempty"`
	//
	RoleGodPrize []int32 `protobuf:"varint,67,rep,name=RoleGodPrize,proto3" json:"RoleGodPrize,omitempty"`
	//
	CreateServerId int32 `protobuf:"varint,68,opt,name=CreateServerId,proto3" json:"CreateServerId,omitempty"`
	//
	PassCheckData []*PassCheckItem `protobuf:"bytes,69,rep,name=PassCheckData,proto3" json:"PassCheckData,omitempty"`
	//
	ZF []*ZF `protobuf:"bytes,70,rep,name=ZF,proto3" json:"ZF,omitempty"`
	//
	PreciousPos []*PreciousPos `protobuf:"bytes,71,rep,name=PreciousPos,proto3" json:"PreciousPos,omitempty"`
	//
	PhotoBooks []*PhotoBook `protobuf:"bytes,72,rep,name=PhotoBooks,proto3" json:"PhotoBooks,omitempty"`
	//
	ItemAttr []*IntAttr `protobuf:"bytes,73,rep,name=ItemAttr,proto3" json:"ItemAttr,omitempty"`
	//
	ChargeGift []int32 `protobuf:"varint,74,rep,name=ChargeGift,proto3" json:"ChargeGift,omitempty"`
	//
	YesterdayTasks []*Task `protobuf:"bytes,75,rep,name=YesterdayTasks,proto3" json:"YesterdayTasks,omitempty"`
	//
	StarLuck []*IntAttr `protobuf:"bytes,76,rep,name=StarLuck,proto3" json:"StarLuck,omitempty"`
	//
	StageGift []int32 `protobuf:"varint,77,rep,name=StageGift,proto3" json:"StageGift,omitempty"`
	//
	InstanceMaterialData []*InstanceMaterialData `protobuf:"bytes,78,rep,name=InstanceMaterialData,proto3" json:"InstanceMaterialData,omitempty"`
	//
	EquipGems []*EquipGem `protobuf:"bytes,79,rep,name=EquipGems,proto3" json:"EquipGems,omitempty"`
	//
	BecomeGodTask []*BecomeGodTask `protobuf:"bytes,80,rep,name=BecomeGodTask,proto3" json:"BecomeGodTask,omitempty"`
	//
	HorseEquipCollects []*HorseEquipCollect `protobuf:"bytes,81,rep,name=HorseEquipCollects,proto3" json:"HorseEquipCollects,omitempty"`
	//
	TongTianData *TongTianData `protobuf:"bytes,82,opt,name=TongTianData,proto3" json:"TongTianData,omitempty"`
	//
	ZhenFas []*ZhenFa `protobuf:"bytes,83,rep,name=ZhenFas,proto3" json:"ZhenFas,omitempty"`
	//
	JiBans []*JiBan `protobuf:"bytes,84,rep,name=JiBans,proto3" json:"JiBans,omitempty"`
	//
	PartnerSportsData []*PartnerSportsData `protobuf:"bytes,85,rep,name=PartnerSportsData,proto3" json:"PartnerSportsData,omitempty"`
	//
	BuBuShengLianData *InstanceBuBuShengLianData `protobuf:"bytes,86,opt,name=BuBuShengLianData,proto3" json:"BuBuShengLianData,omitempty"`
	//
	TianShus []*TianShu `protobuf:"bytes,87,rep,name=TianShus,proto3" json:"TianShus,omitempty"`
	//
	Vips []*Vip `protobuf:"bytes,88,rep,name=Vips,proto3" json:"Vips,omitempty"`
	//
	NewPrize2 []int32 `protobuf:"varint,89,rep,name=NewPrize2,proto3" json:"NewPrize2,omitempty"`
	//
	ShopNum []*IntAttr `protobuf:"bytes,90,rep,name=ShopNum,proto3" json:"ShopNum,omitempty"`
}

func (m *RoleDbInfo) Reset()         { *m = RoleDbInfo{} }
func (m *RoleDbInfo) String() string { return proto.CompactTextString(m) }
func (*RoleDbInfo) ProtoMessage()    {}

//
type RobberData struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Times int64 `protobuf:"varint,2,opt,name=Times,proto3" json:"Times,omitempty"`
}

func (m *RobberData) Reset()         { *m = RobberData{} }
func (m *RobberData) String() string { return proto.CompactTextString(m) }
func (*RobberData) ProtoMessage()    {}

//
type TreasureData struct {
	//
	Max3Star int64 `protobuf:"varint,1,opt,name=Max3Star,proto3" json:"Max3Star,omitempty"`
	//
	Stage []int64 `protobuf:"varint,2,rep,name=Stage,proto3" json:"Stage,omitempty"`
	//
	MaxPrize int64 `protobuf:"varint,3,opt,name=MaxPrize,proto3" json:"MaxPrize,omitempty"`
	//
	GetPrize []int64 `protobuf:"varint,4,rep,name=GetPrize,proto3" json:"GetPrize,omitempty"`
	//
	MaxBoxId int64 `protobuf:"varint,5,opt,name=MaxBoxId,proto3" json:"MaxBoxId,omitempty"`
	//
	BoxIds []int64 `protobuf:"varint,6,rep,name=BoxIds,proto3" json:"BoxIds,omitempty"`
}

func (m *TreasureData) Reset()         { *m = TreasureData{} }
func (m *TreasureData) String() string { return proto.CompactTextString(m) }
func (*TreasureData) ProtoMessage()    {}

//
type TongTian struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Star int32 `protobuf:"varint,2,opt,name=Star,proto3" json:"Star,omitempty"`
	//
	FightTimes int32 `protobuf:"varint,3,opt,name=FightTimes,proto3" json:"FightTimes,omitempty"`
	//
	IsGetBox int32 `protobuf:"varint,4,opt,name=IsGetBox,proto3" json:"IsGetBox,omitempty"`
	//
	ResetTimes int32 `protobuf:"varint,5,opt,name=ResetTimes,proto3" json:"ResetTimes,omitempty"`
}

func (m *TongTian) Reset()         { *m = TongTian{} }
func (m *TongTian) String() string { return proto.CompactTextString(m) }
func (*TongTian) ProtoMessage()    {}

//
type TongTianData struct {
	//
	tongTina []*TongTian `protobuf:"bytes,1,rep,name=tongTina,proto3" json:"tongTina,omitempty"`
	//
	tongTianBox []int32 `protobuf:"varint,2,rep,name=tongTianBox,proto3" json:"tongTianBox,omitempty"`
}

func (m *TongTianData) Reset()         { *m = TongTianData{} }
func (m *TongTianData) String() string { return proto.CompactTextString(m) }
func (*TongTianData) ProtoMessage()    {}

//
type JJCChange struct {
	//
	T int32 `protobuf:"varint,1,opt,name=T,proto3" json:"T,omitempty"`
	//
	R int32 `protobuf:"varint,2,opt,name=R,proto3" json:"R,omitempty"`
}

func (m *JJCChange) Reset()         { *m = JJCChange{} }
func (m *JJCChange) String() string { return proto.CompactTextString(m) }
func (*JJCChange) ProtoMessage()    {}

//
type BaseDbInfo struct {
	//
	intAttr []*IntAttr `protobuf:"bytes,1,rep,name=intAttr,proto3" json:"intAttr,omitempty"`
	//
	strAttr []*StrAttr `protobuf:"bytes,2,rep,name=strAttr,proto3" json:"strAttr,omitempty"`
}

func (m *BaseDbInfo) Reset()         { *m = BaseDbInfo{} }
func (m *BaseDbInfo) String() string { return proto.CompactTextString(m) }
func (*BaseDbInfo) ProtoMessage()    {}

//
type Skin struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	W int32 `protobuf:"varint,2,opt,name=W,proto3" json:"W,omitempty"`
	//
	V int64 `protobuf:"varint,3,opt,name=V,proto3" json:"V,omitempty"`
	//
	Old int32 `protobuf:"varint,4,opt,name=Old,proto3" json:"Old,omitempty"`
	//
	L int32 `protobuf:"varint,5,opt,name=L,proto3" json:"L,omitempty"`
}

func (m *Skin) Reset()         { *m = Skin{} }
func (m *Skin) String() string { return proto.CompactTextString(m) }
func (*Skin) ProtoMessage()    {}

//
type RoleBaseInfo struct {
	//
	UserId int32 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	Nick string `protobuf:"bytes,2,opt,name=Nick,proto3" json:"Nick,omitempty"`
	//
	LogoutTime int64 `protobuf:"varint,3,opt,name=LogoutTime,proto3" json:"LogoutTime,omitempty"`
	//
	FightValue int64 `protobuf:"varint,4,opt,name=FightValue,proto3" json:"FightValue,omitempty"`
	//
	VipLevel int32 `protobuf:"varint,5,opt,name=VipLevel,proto3" json:"VipLevel,omitempty"`
	//
	RoleId int32 `protobuf:"varint,6,opt,name=RoleId,proto3" json:"RoleId,omitempty"`
	//
	OnLine bool `protobuf:"varint,7,opt,name=OnLine,proto3" json:"OnLine,omitempty"`
}

func (m *RoleBaseInfo) Reset()         { *m = RoleBaseInfo{} }
func (m *RoleBaseInfo) String() string { return proto.CompactTextString(m) }
func (*RoleBaseInfo) ProtoMessage()    {}

const PCK_C2SChangeNick_ID = 420 //
//
type C2SChangeNick struct {
	//
	Nick string `protobuf:"bytes,1,opt,name=Nick,proto3" json:"Nick,omitempty"`
}

func (m *C2SChangeNick) Reset()         { *m = C2SChangeNick{} }
func (m *C2SChangeNick) String() string { return proto.CompactTextString(m) }
func (*C2SChangeNick) ProtoMessage()    {}
func (m *C2SChangeNick) GetId() uint16  { return PCK_C2SChangeNick_ID }

const PCK_S2CChangeNick_ID = 421 //
//
type S2CChangeNick struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CChangeNick) Reset()         { *m = S2CChangeNick{} }
func (m *S2CChangeNick) String() string { return proto.CompactTextString(m) }
func (*S2CChangeNick) ProtoMessage()    {}
func (m *S2CChangeNick) GetId() uint16  { return PCK_S2CChangeNick_ID }

//
type PlayerChargeData struct {
	//
	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	GoodsId int32 `protobuf:"varint,2,opt,name=GoodsId,proto3" json:"GoodsId,omitempty"`
	//
	ChannelOrderId string `protobuf:"bytes,3,opt,name=ChannelOrderId,proto3" json:"ChannelOrderId,omitempty"`
	//
	PfOrderId string `protobuf:"bytes,4,opt,name=PfOrderId,proto3" json:"PfOrderId,omitempty"`
	//
	ChannelId int32 `protobuf:"varint,5,opt,name=ChannelId,proto3" json:"ChannelId,omitempty"`
	//
	RealAmount int32 `protobuf:"varint,6,opt,name=RealAmount,proto3" json:"RealAmount,omitempty"`
	//
	ExtraInfo string `protobuf:"bytes,7,opt,name=ExtraInfo,proto3" json:"ExtraInfo,omitempty"`
}

func (m *PlayerChargeData) Reset()         { *m = PlayerChargeData{} }
func (m *PlayerChargeData) String() string { return proto.CompactTextString(m) }
func (*PlayerChargeData) ProtoMessage()    {}

//
type KingData struct {
	//
	State int32 `protobuf:"varint,1,opt,name=State,proto3" json:"State,omitempty"`
	//
	StartTime int64 `protobuf:"varint,2,opt,name=StartTime,proto3" json:"StartTime,omitempty"`
	//
	PlayerData []*KingPlayerData `protobuf:"bytes,3,rep,name=PlayerData,proto3" json:"PlayerData,omitempty"`
	//
	LastWeekPlayerData []*KingPlayerData `protobuf:"bytes,4,rep,name=LastWeekPlayerData,proto3" json:"LastWeekPlayerData,omitempty"`
}

func (m *KingData) Reset()         { *m = KingData{} }
func (m *KingData) String() string { return proto.CompactTextString(m) }
func (*KingData) ProtoMessage()    {}

//
type KingPlayerData struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	KingRank int32 `protobuf:"varint,2,opt,name=KingRank,proto3" json:"KingRank,omitempty"`
	//
	Win int32 `protobuf:"varint,3,opt,name=Win,proto3" json:"Win,omitempty"`
	//
	ContinueWin int32 `protobuf:"varint,4,opt,name=ContinueWin,proto3" json:"ContinueWin,omitempty"`
	//
	Times int32 `protobuf:"varint,5,opt,name=Times,proto3" json:"Times,omitempty"`
	//
	NextTimes int32 `protobuf:"varint,6,opt,name=NextTimes,proto3" json:"NextTimes,omitempty"`
	//
	IsRobot int32 `protobuf:"varint,7,opt,name=IsRobot,proto3" json:"IsRobot,omitempty"`
}

func (m *KingPlayerData) Reset()         { *m = KingPlayerData{} }
func (m *KingPlayerData) String() string { return proto.CompactTextString(m) }
func (*KingPlayerData) ProtoMessage()    {}

//
type JJCData struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	JJC int32 `protobuf:"varint,2,opt,name=JJC,proto3" json:"JJC,omitempty"`
}

func (m *JJCData) Reset()         { *m = JJCData{} }
func (m *JJCData) String() string { return proto.CompactTextString(m) }
func (*JJCData) ProtoMessage()    {}

//
type ServerData struct {
	//
	DragonLogs []*DragonLogItem `protobuf:"bytes,1,rep,name=DragonLogs,proto3" json:"DragonLogs,omitempty"`
	//
	ActGangWar *ActGangWarPreInfo `protobuf:"bytes,2,opt,name=ActGangWar,proto3" json:"ActGangWar,omitempty"`
	//
	FirstNick string `protobuf:"bytes,4,opt,name=FirstNick,proto3" json:"FirstNick,omitempty"`
	//
	PlayerCharge []*PlayerChargeData `protobuf:"bytes,6,rep,name=PlayerCharge,proto3" json:"PlayerCharge,omitempty"`
	//
	Instance81 []*Instance81Data `protobuf:"bytes,8,rep,name=Instance81,proto3" json:"Instance81,omitempty"`
	//
	WorldBossKillData *WorldBossKillData `protobuf:"bytes,9,opt,name=WorldBossKillData,proto3" json:"WorldBossKillData,omitempty"`
	//
	KingData *KingData `protobuf:"bytes,10,opt,name=KingData,proto3" json:"KingData,omitempty"`
	//
	RobberyLogs []*RobberyLogItem `protobuf:"bytes,11,rep,name=RobberyLogs,proto3" json:"RobberyLogs,omitempty"`
	//
	JJCData []*JJCData `protobuf:"bytes,12,rep,name=JJCData,proto3" json:"JJCData,omitempty"`
	//
	ActGroupData []*ActGroupData `protobuf:"bytes,13,rep,name=ActGroupData,proto3" json:"ActGroupData,omitempty"`
	//
	NowDay int64 `protobuf:"varint,14,opt,name=NowDay,proto3" json:"NowDay,omitempty"`
	//
	GangCreateTime int32 `protobuf:"varint,15,opt,name=GangCreateTime,proto3" json:"GangCreateTime,omitempty"`
	//
	DrawLog []*DrawDbLog `protobuf:"bytes,16,rep,name=DrawLog,proto3" json:"DrawLog,omitempty"`
	//
	ChargeReturnLog []*ChargeReturnDbLog `protobuf:"bytes,17,rep,name=ChargeReturnLog,proto3" json:"ChargeReturnLog,omitempty"`
	//
	ResetChargeTime int64 `protobuf:"varint,18,opt,name=ResetChargeTime,proto3" json:"ResetChargeTime,omitempty"`
}

func (m *ServerData) Reset()         { *m = ServerData{} }
func (m *ServerData) String() string { return proto.CompactTextString(m) }
func (*ServerData) ProtoMessage()    {}

//
type ActGroupData struct {
	//
	GroupId int32 `protobuf:"varint,1,opt,name=GroupId,proto3" json:"GroupId,omitempty"`
	//
	RoundNum int32 `protobuf:"varint,2,opt,name=RoundNum,proto3" json:"RoundNum,omitempty"`
	//
	OrderIdx int32 `protobuf:"varint,3,opt,name=OrderIdx,proto3" json:"OrderIdx,omitempty"`
}

func (m *ActGroupData) Reset()         { *m = ActGroupData{} }
func (m *ActGroupData) String() string { return proto.CompactTextString(m) }
func (*ActGroupData) ProtoMessage()    {}

//
type CrossServerData struct {
	//
	NineLastTimeDay int32 `protobuf:"varint,1,opt,name=NineLastTimeDay,proto3" json:"NineLastTimeDay,omitempty"`
}

func (m *CrossServerData) Reset()         { *m = CrossServerData{} }
func (m *CrossServerData) String() string { return proto.CompactTextString(m) }
func (*CrossServerData) ProtoMessage()    {}

const PCK_S2CServerInfo_ID = 4 //
//
type S2CServerInfo struct {
	//
	IsOpenRealName int32 `protobuf:"varint,1,opt,name=IsOpenRealName,proto3" json:"IsOpenRealName,omitempty"`
	//
	IsOpenFcm int32 `protobuf:"varint,2,opt,name=IsOpenFcm,proto3" json:"IsOpenFcm,omitempty"`
	//
	SendShopMaxNum int64 `protobuf:"varint,3,opt,name=SendShopMaxNum,proto3" json:"SendShopMaxNum,omitempty"`
	//
	OpenDragon int64 `protobuf:"varint,4,opt,name=OpenDragon,proto3" json:"OpenDragon,omitempty"`
	//
	IsOpenSend int64 `protobuf:"varint,5,opt,name=IsOpenSend,proto3" json:"IsOpenSend,omitempty"`
}

func (m *S2CServerInfo) Reset()         { *m = S2CServerInfo{} }
func (m *S2CServerInfo) String() string { return proto.CompactTextString(m) }
func (*S2CServerInfo) ProtoMessage()    {}
func (m *S2CServerInfo) GetId() uint16  { return PCK_S2CServerInfo_ID }

//
type K2SAABB struct {
}

func (m *K2SAABB) Reset()         { *m = K2SAABB{} }
func (m *K2SAABB) String() string { return proto.CompactTextString(m) }
func (*K2SAABB) ProtoMessage()    {}

//
type MAX struct {
}

func (m *MAX) Reset()         { *m = MAX{} }
func (m *MAX) String() string { return proto.CompactTextString(m) }
func (*MAX) ProtoMessage()    {}

const PCK_C2SLogin_ID = 1 //
//
type C2SLogin struct {
	//
	AccountId int64 `protobuf:"varint,1,opt,name=AccountId,proto3" json:"AccountId,omitempty"`
	//
	Token string `protobuf:"bytes,2,opt,name=Token,proto3" json:"Token,omitempty"`
	//
	UserId int64 `protobuf:"varint,3,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	Fcm int32 `protobuf:"varint,4,opt,name=Fcm,proto3" json:"Fcm,omitempty"`
	//
	LoginPf string `protobuf:"bytes,5,opt,name=LoginPf,proto3" json:"LoginPf,omitempty"`
}

func (m *C2SLogin) Reset()         { *m = C2SLogin{} }
func (m *C2SLogin) String() string { return proto.CompactTextString(m) }
func (*C2SLogin) ProtoMessage()    {}
func (m *C2SLogin) GetId() uint16  { return PCK_C2SLogin_ID }

const PCK_S2CLogin_ID = 2 //
//
type S2CLogin struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	UserId int32 `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (m *S2CLogin) Reset()         { *m = S2CLogin{} }
func (m *S2CLogin) String() string { return proto.CompactTextString(m) }
func (*S2CLogin) ProtoMessage()    {}
func (m *S2CLogin) GetId() uint16  { return PCK_S2CLogin_ID }

const PCK_C2SReLogin_ID = 91 //
//
type C2SReLogin struct {
	//
	AccountId int64 `protobuf:"varint,1,opt,name=AccountId,proto3" json:"AccountId,omitempty"`
	//
	Token string `protobuf:"bytes,2,opt,name=Token,proto3" json:"Token,omitempty"`
	//
	UserId int64 `protobuf:"varint,3,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (m *C2SReLogin) Reset()         { *m = C2SReLogin{} }
func (m *C2SReLogin) String() string { return proto.CompactTextString(m) }
func (*C2SReLogin) ProtoMessage()    {}
func (m *C2SReLogin) GetId() uint16  { return PCK_C2SReLogin_ID }

const PCK_S2CReLogin_ID = 92 //
//
type S2CReLogin struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CReLogin) Reset()         { *m = S2CReLogin{} }
func (m *S2CReLogin) String() string { return proto.CompactTextString(m) }
func (*S2CReLogin) ProtoMessage()    {}
func (m *S2CReLogin) GetId() uint16  { return PCK_S2CReLogin_ID }

const PCK_C2SLoginEnd_ID = 29 //
//
type C2SLoginEnd struct {
}

func (m *C2SLoginEnd) Reset()         { *m = C2SLoginEnd{} }
func (m *C2SLoginEnd) String() string { return proto.CompactTextString(m) }
func (*C2SLoginEnd) ProtoMessage()    {}
func (m *C2SLoginEnd) GetId() uint16  { return PCK_C2SLoginEnd_ID }

const PCK_S2CAllData_ID = 39 //
//
type S2CAllData struct {
}

func (m *S2CAllData) Reset()         { *m = S2CAllData{} }
func (m *S2CAllData) String() string { return proto.CompactTextString(m) }
func (*S2CAllData) ProtoMessage()    {}
func (m *S2CAllData) GetId() uint16  { return PCK_S2CAllData_ID }

const PCK_C2SUserInfo_ID = 27 //
//
type C2SUserInfo struct {
	//
	AccountId int32 `protobuf:"varint,1,opt,name=AccountId,proto3" json:"AccountId,omitempty"`
	//
	Token string `protobuf:"bytes,2,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (m *C2SUserInfo) Reset()         { *m = C2SUserInfo{} }
func (m *C2SUserInfo) String() string { return proto.CompactTextString(m) }
func (*C2SUserInfo) ProtoMessage()    {}
func (m *C2SUserInfo) GetId() uint16  { return PCK_C2SUserInfo_ID }

const PCK_S2CUserInfo_ID = 28 //
//
type S2CUserInfo struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	MapId int32 `protobuf:"varint,2,opt,name=MapId,proto3" json:"MapId,omitempty"`
	//
	MapX int32 `protobuf:"varint,3,opt,name=MapX,proto3" json:"MapX,omitempty"`
	//
	MapY int32 `protobuf:"varint,4,opt,name=MapY,proto3" json:"MapY,omitempty"`
	//
	A []*IntAttr `protobuf:"bytes,5,rep,name=A,proto3" json:"A,omitempty"`
}

func (m *S2CUserInfo) Reset()         { *m = S2CUserInfo{} }
func (m *S2CUserInfo) String() string { return proto.CompactTextString(m) }
func (*S2CUserInfo) ProtoMessage()    {}
func (m *S2CUserInfo) GetId() uint16  { return PCK_S2CUserInfo_ID }

const PCK_S2COfflinePrize_ID = 15 //
//
type S2COfflinePrize struct {
	//
	Gold int64 `protobuf:"varint,1,opt,name=Gold,proto3" json:"Gold,omitempty"`
	//
	Exp int64 `protobuf:"varint,2,opt,name=Exp,proto3" json:"Exp,omitempty"`
	//
	Equip int32 `protobuf:"varint,3,opt,name=Equip,proto3" json:"Equip,omitempty"`
	//
	LogoutTime int64 `protobuf:"varint,4,opt,name=LogoutTime,proto3" json:"LogoutTime,omitempty"`
}

func (m *S2COfflinePrize) Reset()         { *m = S2COfflinePrize{} }
func (m *S2COfflinePrize) String() string { return proto.CompactTextString(m) }
func (*S2COfflinePrize) ProtoMessage()    {}
func (m *S2COfflinePrize) GetId() uint16  { return PCK_S2COfflinePrize_ID }

const PCK_C2SGetOfflinePrize_ID = 16 //
//
type C2SGetOfflinePrize struct {
	//
	Multi int32 `protobuf:"varint,1,opt,name=Multi,proto3" json:"Multi,omitempty"`
}

func (m *C2SGetOfflinePrize) Reset()         { *m = C2SGetOfflinePrize{} }
func (m *C2SGetOfflinePrize) String() string { return proto.CompactTextString(m) }
func (*C2SGetOfflinePrize) ProtoMessage()    {}
func (m *C2SGetOfflinePrize) GetId() uint16  { return PCK_C2SGetOfflinePrize_ID }

const PCK_S2CGetOfflinePrize_ID = 17 //
//
type S2CGetOfflinePrize struct {
	//
	Tag int64 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CGetOfflinePrize) Reset()         { *m = S2CGetOfflinePrize{} }
func (m *S2CGetOfflinePrize) String() string { return proto.CompactTextString(m) }
func (*S2CGetOfflinePrize) ProtoMessage()    {}
func (m *S2CGetOfflinePrize) GetId() uint16  { return PCK_S2CGetOfflinePrize_ID }

const PCK_LoginFinish_ID = 24 //
//
type LoginFinish struct {
}

func (m *LoginFinish) Reset()         { *m = LoginFinish{} }
func (m *LoginFinish) String() string { return proto.CompactTextString(m) }
func (*LoginFinish) ProtoMessage()    {}
func (m *LoginFinish) GetId() uint16  { return PCK_LoginFinish_ID }

const PCK_C2SGetGift1_ID = 18 //
//
type C2SGetGift1 struct {
}

func (m *C2SGetGift1) Reset()         { *m = C2SGetGift1{} }
func (m *C2SGetGift1) String() string { return proto.CompactTextString(m) }
func (*C2SGetGift1) ProtoMessage()    {}
func (m *C2SGetGift1) GetId() uint16  { return PCK_C2SGetGift1_ID }

const PCK_S2CGetGift1_ID = 19 //
//
type S2CGetGift1 struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CGetGift1) Reset()         { *m = S2CGetGift1{} }
func (m *S2CGetGift1) String() string { return proto.CompactTextString(m) }
func (*S2CGetGift1) ProtoMessage()    {}
func (m *S2CGetGift1) GetId() uint16  { return PCK_S2CGetGift1_ID }

const PCK_C2SGetGift2_ID = 20 //
//
type C2SGetGift2 struct {
}

func (m *C2SGetGift2) Reset()         { *m = C2SGetGift2{} }
func (m *C2SGetGift2) String() string { return proto.CompactTextString(m) }
func (*C2SGetGift2) ProtoMessage()    {}
func (m *C2SGetGift2) GetId() uint16  { return PCK_C2SGetGift2_ID }

const PCK_S2CGetGift2_ID = 21 //
//
type S2CGetGift2 struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CGetGift2) Reset()         { *m = S2CGetGift2{} }
func (m *S2CGetGift2) String() string { return proto.CompactTextString(m) }
func (*S2CGetGift2) ProtoMessage()    {}
func (m *S2CGetGift2) GetId() uint16  { return PCK_S2CGetGift2_ID }

const PCK_C2SFocusPrice_ID = 32 //
//
type C2SFocusPrice struct {
	//
	Type int32 `protobuf:"varint,1,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (m *C2SFocusPrice) Reset()         { *m = C2SFocusPrice{} }
func (m *C2SFocusPrice) String() string { return proto.CompactTextString(m) }
func (*C2SFocusPrice) ProtoMessage()    {}
func (m *C2SFocusPrice) GetId() uint16  { return PCK_C2SFocusPrice_ID }

const PCK_S2CFocusPrice_ID = 33 //
//
type S2CFocusPrice struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CFocusPrice) Reset()         { *m = S2CFocusPrice{} }
func (m *S2CFocusPrice) String() string { return proto.CompactTextString(m) }
func (*S2CFocusPrice) ProtoMessage()    {}
func (m *S2CFocusPrice) GetId() uint16  { return PCK_S2CFocusPrice_ID }

const PCK_C2SFirstInvite_ID = 34 //
//
type C2SFirstInvite struct {
}

func (m *C2SFirstInvite) Reset()         { *m = C2SFirstInvite{} }
func (m *C2SFirstInvite) String() string { return proto.CompactTextString(m) }
func (*C2SFirstInvite) ProtoMessage()    {}
func (m *C2SFirstInvite) GetId() uint16  { return PCK_C2SFirstInvite_ID }

const PCK_S2CFirstInvite_ID = 35 //
//
type S2CFirstInvite struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CFirstInvite) Reset()         { *m = S2CFirstInvite{} }
func (m *S2CFirstInvite) String() string { return proto.CompactTextString(m) }
func (*S2CFirstInvite) ProtoMessage()    {}
func (m *S2CFirstInvite) GetId() uint16  { return PCK_S2CFirstInvite_ID }

const PCK_C2SNewStory_ID = 36 //
//
type C2SNewStory struct {
}

func (m *C2SNewStory) Reset()         { *m = C2SNewStory{} }
func (m *C2SNewStory) String() string { return proto.CompactTextString(m) }
func (*C2SNewStory) ProtoMessage()    {}
func (m *C2SNewStory) GetId() uint16  { return PCK_C2SNewStory_ID }

const PCK_S2CNewStory_ID = 37 //
//
type S2CNewStory struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CNewStory) Reset()         { *m = S2CNewStory{} }
func (m *S2CNewStory) String() string { return proto.CompactTextString(m) }
func (*S2CNewStory) ProtoMessage()    {}
func (m *S2CNewStory) GetId() uint16  { return PCK_S2CNewStory_ID }

const PCK_S2CRoleInfo_ID = 3 //
//
type S2CRoleInfo struct {
	//
	UserId int32 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	A []*IntAttr `protobuf:"bytes,2,rep,name=A,proto3" json:"A,omitempty"`
	//
	B []*StrAttr `protobuf:"bytes,3,rep,name=B,proto3" json:"B,omitempty"`
}

func (m *S2CRoleInfo) Reset()         { *m = S2CRoleInfo{} }
func (m *S2CRoleInfo) String() string { return proto.CompactTextString(m) }
func (*S2CRoleInfo) ProtoMessage()    {}
func (m *S2CRoleInfo) GetId() uint16  { return PCK_S2CRoleInfo_ID }

const PCK_C2SOtherInfo_ID = 10 //
//
type C2SOtherInfo struct {
	//
	UserId int32 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (m *C2SOtherInfo) Reset()         { *m = C2SOtherInfo{} }
func (m *C2SOtherInfo) String() string { return proto.CompactTextString(m) }
func (*C2SOtherInfo) ProtoMessage()    {}
func (m *C2SOtherInfo) GetId() uint16  { return PCK_C2SOtherInfo_ID }

const PCK_S2COtherInfo_ID = 11 //
//
type S2COtherInfo struct {
	//
	UserId int32 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	A []*IntAttr `protobuf:"bytes,2,rep,name=A,proto3" json:"A,omitempty"`
	//
	B []*StrAttr `protobuf:"bytes,3,rep,name=B,proto3" json:"B,omitempty"`
	//
	MateNick string `protobuf:"bytes,4,opt,name=MateNick,proto3" json:"MateNick,omitempty"`
	//
	Tag int32 `protobuf:"varint,5,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2COtherInfo) Reset()         { *m = S2COtherInfo{} }
func (m *S2COtherInfo) String() string { return proto.CompactTextString(m) }
func (*S2COtherInfo) ProtoMessage()    {}
func (m *S2COtherInfo) GetId() uint16  { return PCK_S2COtherInfo_ID }

const PCK_S2CRoleData_ID = 9 //
//
type S2CRoleData struct {
	//
	ActiveSkills []*Skill `protobuf:"bytes,1,rep,name=ActiveSkills,proto3" json:"ActiveSkills,omitempty"`
	//
	Skins []*Skin `protobuf:"bytes,2,rep,name=Skins,proto3" json:"Skins,omitempty"`
	//
	VIPPrize []int32 `protobuf:"varint,3,rep,name=VIPPrize,proto3" json:"VIPPrize,omitempty"`
	//
	Task []*S2CTask `protobuf:"bytes,4,rep,name=Task,proto3" json:"Task,omitempty"`
	//
	Counters []*S2CCounter `protobuf:"bytes,5,rep,name=Counters,proto3" json:"Counters,omitempty"`
	//
	WorldLevel *WorldLevel `protobuf:"bytes,6,opt,name=WorldLevel,proto3" json:"WorldLevel,omitempty"`
	//
	Precious []*Precious `protobuf:"bytes,7,rep,name=Precious,proto3" json:"Precious,omitempty"`
	//
	NewPrize []int32 `protobuf:"varint,8,rep,name=NewPrize,proto3" json:"NewPrize,omitempty"`
	//
	VIPGift []int32 `protobuf:"varint,9,rep,name=VIPGift,proto3" json:"VIPGift,omitempty"`
	//
	ChargeGift []int32 `protobuf:"varint,10,rep,name=ChargeGift,proto3" json:"ChargeGift,omitempty"`
	//
	StateGift []int32 `protobuf:"varint,11,rep,name=StateGift,proto3" json:"StateGift,omitempty"`
	//
	EquipGems []*EquipGem `protobuf:"bytes,12,rep,name=EquipGems,proto3" json:"EquipGems,omitempty"`
	//
	HorseEquipCollects []*HorseEquipCollect `protobuf:"bytes,13,rep,name=HorseEquipCollects,proto3" json:"HorseEquipCollects,omitempty"`
	//
	ZhenFas []*ZhenFa `protobuf:"bytes,14,rep,name=ZhenFas,proto3" json:"ZhenFas,omitempty"`
	//
	JiBans []*JiBan `protobuf:"bytes,15,rep,name=JiBans,proto3" json:"JiBans,omitempty"`
	//
	TianShus []*TianShu `protobuf:"bytes,16,rep,name=TianShus,proto3" json:"TianShus,omitempty"`
	//
	Vips []*Vip `protobuf:"bytes,17,rep,name=Vips,proto3" json:"Vips,omitempty"`
	//
	NewPrize2 []int32 `protobuf:"varint,18,rep,name=NewPrize2,proto3" json:"NewPrize2,omitempty"`
	//
	ShopData []*IntAttr `protobuf:"bytes,19,rep,name=ShopData,proto3" json:"ShopData,omitempty"`
}

func (m *S2CRoleData) Reset()         { *m = S2CRoleData{} }
func (m *S2CRoleData) String() string { return proto.CompactTextString(m) }
func (*S2CRoleData) ProtoMessage()    {}
func (m *S2CRoleData) GetId() uint16  { return PCK_S2CRoleData_ID }

const PCK_C2SAchieveShow_ID = 25 //
//
type C2SAchieveShow struct {
	//
	AchieveShow int32 `protobuf:"varint,1,opt,name=AchieveShow,proto3" json:"AchieveShow,omitempty"`
}

func (m *C2SAchieveShow) Reset()         { *m = C2SAchieveShow{} }
func (m *C2SAchieveShow) String() string { return proto.CompactTextString(m) }
func (*C2SAchieveShow) ProtoMessage()    {}
func (m *C2SAchieveShow) GetId() uint16  { return PCK_C2SAchieveShow_ID }

const PCK_S2CAchieveShow_ID = 26 //
//
type S2CAchieveShow struct {
	//
	AchieveShow []int32 `protobuf:"varint,1,rep,name=AchieveShow,proto3" json:"AchieveShow,omitempty"`
}

func (m *S2CAchieveShow) Reset()         { *m = S2CAchieveShow{} }
func (m *S2CAchieveShow) String() string { return proto.CompactTextString(m) }
func (*S2CAchieveShow) ProtoMessage()    {}
func (m *S2CAchieveShow) GetId() uint16  { return PCK_S2CAchieveShow_ID }

const PCK_S2CRoleBaseInfo_ID = 8 //
//
type S2CRoleBaseInfo struct {
	//
	UserId int32 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	Nick string `protobuf:"bytes,2,opt,name=Nick,proto3" json:"Nick,omitempty"`
	//
	LogoutTime int64 `protobuf:"varint,3,opt,name=LogoutTime,proto3" json:"LogoutTime,omitempty"`
	//
	FightValue int64 `protobuf:"varint,4,opt,name=FightValue,proto3" json:"FightValue,omitempty"`
	//
	VipLevel int32 `protobuf:"varint,5,opt,name=VipLevel,proto3" json:"VipLevel,omitempty"`
	//
	Level int32 `protobuf:"varint,6,opt,name=Level,proto3" json:"Level,omitempty"`
	//
	RoleId int32 `protobuf:"varint,7,opt,name=RoleId,proto3" json:"RoleId,omitempty"`
	//
	OnLine bool `protobuf:"varint,8,opt,name=OnLine,proto3" json:"OnLine,omitempty"`
}

func (m *S2CRoleBaseInfo) Reset()         { *m = S2CRoleBaseInfo{} }
func (m *S2CRoleBaseInfo) String() string { return proto.CompactTextString(m) }
func (*S2CRoleBaseInfo) ProtoMessage()    {}
func (m *S2CRoleBaseInfo) GetId() uint16  { return PCK_S2CRoleBaseInfo_ID }

const PCK_C2SRoleBaseInfo_ID = 7 //
//
type C2SRoleBaseInfo struct {
	//
	UserId int32 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (m *C2SRoleBaseInfo) Reset()         { *m = C2SRoleBaseInfo{} }
func (m *C2SRoleBaseInfo) String() string { return proto.CompactTextString(m) }
func (*C2SRoleBaseInfo) ProtoMessage()    {}
func (m *C2SRoleBaseInfo) GetId() uint16  { return PCK_C2SRoleBaseInfo_ID }

const PCK_S2CKill_ID = 6 //
//
type S2CKill struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CKill) Reset()         { *m = S2CKill{} }
func (m *S2CKill) String() string { return proto.CompactTextString(m) }
func (*S2CKill) ProtoMessage()    {}
func (m *S2CKill) GetId() uint16  { return PCK_S2CKill_ID }

//
type FightUnit struct {
	//
	P int32 `protobuf:"varint,1,opt,name=P,proto3" json:"P,omitempty"`
	//
	I int64 `protobuf:"varint,2,opt,name=I,proto3" json:"I,omitempty"`
	//
	T int32 `protobuf:"varint,3,opt,name=T,proto3" json:"T,omitempty"`
	//
	A []*IntAttr `protobuf:"bytes,4,rep,name=A,proto3" json:"A,omitempty"`
	//
	B []*StrAttr `protobuf:"bytes,5,rep,name=B,proto3" json:"B,omitempty"`
	//
	F int32 `protobuf:"varint,6,opt,name=F,proto3" json:"F,omitempty"`
	//
	H int64 `protobuf:"varint,7,opt,name=H,proto3" json:"H,omitempty"`
	//
	MH int64 `protobuf:"varint,8,opt,name=MH,proto3" json:"MH,omitempty"`
	//
	H1 int64 `protobuf:"varint,9,opt,name=H1,proto3" json:"H1,omitempty"`
	//
	MH1 int64 `protobuf:"varint,10,opt,name=MH1,proto3" json:"MH1,omitempty"`
	//
	IsMH1 int64 `protobuf:"varint,11,opt,name=IsMH1,proto3" json:"IsMH1,omitempty"`
}

func (m *FightUnit) Reset()         { *m = FightUnit{} }
func (m *FightUnit) String() string { return proto.CompactTextString(m) }
func (*FightUnit) ProtoMessage()    {}

//
type Effect struct {
	//
	K int32 `protobuf:"varint,1,opt,name=K,proto3" json:"K,omitempty"`
	//
	V int64 `protobuf:"varint,2,opt,name=V,proto3" json:"V,omitempty"`
	//
	Protect int64 `protobuf:"varint,3,opt,name=Protect,proto3" json:"Protect,omitempty"`
}

func (m *Effect) Reset()         { *m = Effect{} }
func (m *Effect) String() string { return proto.CompactTextString(m) }
func (*Effect) ProtoMessage()    {}

//
type AtkUnit struct {
	//
	P int32 `protobuf:"varint,1,opt,name=P,proto3" json:"P,omitempty"`
	//
	E []*Effect `protobuf:"bytes,3,rep,name=E,proto3" json:"E,omitempty"`
}

func (m *AtkUnit) Reset()         { *m = AtkUnit{} }
func (m *AtkUnit) String() string { return proto.CompactTextString(m) }
func (*AtkUnit) ProtoMessage()    {}

//
type FightStep struct {
	//
	P int32 `protobuf:"varint,1,opt,name=P,proto3" json:"P,omitempty"`
	//
	S int32 `protobuf:"varint,2,opt,name=S,proto3" json:"S,omitempty"`
	//
	U []*AtkUnit `protobuf:"bytes,3,rep,name=U,proto3" json:"U,omitempty"`
	//
	R int32 `protobuf:"varint,4,opt,name=R,proto3" json:"R,omitempty"`
	//
	E []*Effect `protobuf:"bytes,6,rep,name=E,proto3" json:"E,omitempty"`
}

func (m *FightStep) Reset()         { *m = FightStep{} }
func (m *FightStep) String() string { return proto.CompactTextString(m) }
func (*FightStep) ProtoMessage()    {}

const PCK_S2CBattlefieldReport_ID = 101 //
//
type S2CBattlefieldReport struct {
	//
	U []*FightUnit `protobuf:"bytes,1,rep,name=U,proto3" json:"U,omitempty"`
	//
	S []*FightStep `protobuf:"bytes,2,rep,name=S,proto3" json:"S,omitempty"`
	//
	I []*ItemInfo `protobuf:"bytes,3,rep,name=I,proto3" json:"I,omitempty"`
	//
	T int32 `protobuf:"varint,4,opt,name=T,proto3" json:"T,omitempty"`
	//
	P int64 `protobuf:"varint,5,opt,name=P,proto3" json:"P,omitempty"`
	//
	Win int32 `protobuf:"varint,6,opt,name=Win,proto3" json:"Win,omitempty"`
	//
	Idx int64 `protobuf:"varint,7,opt,name=Idx,proto3" json:"Idx,omitempty"`
	//
	M int32 `protobuf:"varint,8,opt,name=M,proto3" json:"M,omitempty"`
	//
	PVP int32 `protobuf:"varint,9,opt,name=PVP,proto3" json:"PVP,omitempty"`
	//
	Num int32 `protobuf:"varint,10,opt,name=Num,proto3" json:"Num,omitempty"`
	//
	Video int32 `protobuf:"varint,11,opt,name=Video,proto3" json:"Video,omitempty"`
}

func (m *S2CBattlefieldReport) Reset()         { *m = S2CBattlefieldReport{} }
func (m *S2CBattlefieldReport) String() string { return proto.CompactTextString(m) }
func (*S2CBattlefieldReport) ProtoMessage()    {}
func (m *S2CBattlefieldReport) GetId() uint16  { return PCK_S2CBattlefieldReport_ID }

//
type DbBattlefieldReport struct {
	//
	Data [][]byte `protobuf:"bytes,3,rep,name=Data,proto3" json:"Data,omitempty"`
}

func (m *DbBattlefieldReport) Reset()         { *m = DbBattlefieldReport{} }
func (m *DbBattlefieldReport) String() string { return proto.CompactTextString(m) }
func (*DbBattlefieldReport) ProtoMessage()    {}

const PCK_S2KLoginSuccess_ID = 8001 //
//
type S2KLoginSuccess struct {
	//
	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (m *S2KLoginSuccess) Reset()         { *m = S2KLoginSuccess{} }
func (m *S2KLoginSuccess) String() string { return proto.CompactTextString(m) }
func (*S2KLoginSuccess) ProtoMessage()    {}
func (m *S2KLoginSuccess) GetId() uint16  { return PCK_S2KLoginSuccess_ID }

const PCK_K2SBattlefieldReport_ID = 8002 //
//
type K2SBattlefieldReport struct {
	//
	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	IsLeader int64 `protobuf:"varint,2,opt,name=IsLeader,proto3" json:"IsLeader,omitempty"`
	//
	Data []byte `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *K2SBattlefieldReport) Reset()         { *m = K2SBattlefieldReport{} }
func (m *K2SBattlefieldReport) String() string { return proto.CompactTextString(m) }
func (*K2SBattlefieldReport) ProtoMessage()    {}
func (m *K2SBattlefieldReport) GetId() uint16  { return PCK_K2SBattlefieldReport_ID }

const PCK_K2SEnterMap_ID = 8003 //
//
type K2SEnterMap struct {
	//
	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	MapId int64 `protobuf:"varint,2,opt,name=MapId,proto3" json:"MapId,omitempty"`
	//
	InstanceId string `protobuf:"bytes,3,opt,name=InstanceId,proto3" json:"InstanceId,omitempty"`
	//
	X int32 `protobuf:"varint,4,opt,name=X,proto3" json:"X,omitempty"`
	//
	Y int32 `protobuf:"varint,5,opt,name=Y,proto3" json:"Y,omitempty"`
}

func (m *K2SEnterMap) Reset()         { *m = K2SEnterMap{} }
func (m *K2SEnterMap) String() string { return proto.CompactTextString(m) }
func (*K2SEnterMap) ProtoMessage()    {}
func (m *K2SEnterMap) GetId() uint16  { return PCK_K2SEnterMap_ID }

const PCK_S2KEnterMap_ID = 8004 //
//
type S2KEnterMap struct {
	//
	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	MapId int64 `protobuf:"varint,2,opt,name=MapId,proto3" json:"MapId,omitempty"`
	//
	InstanceId string `protobuf:"bytes,3,opt,name=InstanceId,proto3" json:"InstanceId,omitempty"`
	//
	X int32 `protobuf:"varint,4,opt,name=X,proto3" json:"X,omitempty"`
	//
	Y int32 `protobuf:"varint,5,opt,name=Y,proto3" json:"Y,omitempty"`
}

func (m *S2KEnterMap) Reset()         { *m = S2KEnterMap{} }
func (m *S2KEnterMap) String() string { return proto.CompactTextString(m) }
func (*S2KEnterMap) ProtoMessage()    {}
func (m *S2KEnterMap) GetId() uint16  { return PCK_S2KEnterMap_ID }

const PCK_S2KLeaveMap_ID = 8005 //
//
type S2KLeaveMap struct {
	//
	MapId int64 `protobuf:"varint,1,opt,name=MapId,proto3" json:"MapId,omitempty"`
	//
	InstanceId string `protobuf:"bytes,2,opt,name=InstanceId,proto3" json:"InstanceId,omitempty"`
}

func (m *S2KLeaveMap) Reset()         { *m = S2KLeaveMap{} }
func (m *S2KLeaveMap) String() string { return proto.CompactTextString(m) }
func (*S2KLeaveMap) ProtoMessage()    {}
func (m *S2KLeaveMap) GetId() uint16  { return PCK_S2KLeaveMap_ID }

const PCK_S2KLeaveInstanceTeam_ID = 8006 //
//
type S2KLeaveInstanceTeam struct {
	//
	TeamId int64 `protobuf:"varint,1,opt,name=TeamId,proto3" json:"TeamId,omitempty"`
}

func (m *S2KLeaveInstanceTeam) Reset()         { *m = S2KLeaveInstanceTeam{} }
func (m *S2KLeaveInstanceTeam) String() string { return proto.CompactTextString(m) }
func (*S2KLeaveInstanceTeam) ProtoMessage()    {}
func (m *S2KLeaveInstanceTeam) GetId() uint16  { return PCK_S2KLeaveInstanceTeam_ID }

const PCK_S2KLogout_ID = 8009 //
//
type S2KLogout struct {
}

func (m *S2KLogout) Reset()         { *m = S2KLogout{} }
func (m *S2KLogout) String() string { return proto.CompactTextString(m) }
func (*S2KLogout) ProtoMessage()    {}
func (m *S2KLogout) GetId() uint16  { return PCK_S2KLogout_ID }

const PCK_S2KReviveLife_ID = 8017 //
//
type S2KReviveLife struct {
}

func (m *S2KReviveLife) Reset()         { *m = S2KReviveLife{} }
func (m *S2KReviveLife) String() string { return proto.CompactTextString(m) }
func (*S2KReviveLife) ProtoMessage()    {}
func (m *S2KReviveLife) GetId() uint16  { return PCK_S2KReviveLife_ID }

const PCK_S2KWorldLevel_ID = 8018 //
//
type S2KWorldLevel struct {
	//
	Level int64 `protobuf:"varint,1,opt,name=Level,proto3" json:"Level,omitempty"`
}

func (m *S2KWorldLevel) Reset()         { *m = S2KWorldLevel{} }
func (m *S2KWorldLevel) String() string { return proto.CompactTextString(m) }
func (*S2KWorldLevel) ProtoMessage()    {}
func (m *S2KWorldLevel) GetId() uint16  { return PCK_S2KWorldLevel_ID }

const PCK_S2KCreateServerIds_ID = 8120 //
//
type S2KCreateServerIds struct {
	//
	Ids []int64 `protobuf:"varint,1,rep,name=Ids,proto3" json:"Ids,omitempty"`
}

func (m *S2KCreateServerIds) Reset()         { *m = S2KCreateServerIds{} }
func (m *S2KCreateServerIds) String() string { return proto.CompactTextString(m) }
func (*S2KCreateServerIds) ProtoMessage()    {}
func (m *S2KCreateServerIds) GetId() uint16  { return PCK_S2KCreateServerIds_ID }

const PCK_K2SGetPlayerByLevel_ID = 8121 //
//
type K2SGetPlayerByLevel struct {
	//
	MinLevel int64 `protobuf:"varint,1,opt,name=MinLevel,proto3" json:"MinLevel,omitempty"`
	//
	MaxLevel int64 `protobuf:"varint,2,opt,name=MaxLevel,proto3" json:"MaxLevel,omitempty"`
}

func (m *K2SGetPlayerByLevel) Reset()         { *m = K2SGetPlayerByLevel{} }
func (m *K2SGetPlayerByLevel) String() string { return proto.CompactTextString(m) }
func (*K2SGetPlayerByLevel) ProtoMessage()    {}
func (m *K2SGetPlayerByLevel) GetId() uint16  { return PCK_K2SGetPlayerByLevel_ID }

const PCK_S2KGetPlayerOtherInfo_ID = 8122 //
//
type S2KGetPlayerOtherInfo struct {
	//
	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	OtherId int64 `protobuf:"varint,2,opt,name=OtherId,proto3" json:"OtherId,omitempty"`
}

func (m *S2KGetPlayerOtherInfo) Reset()         { *m = S2KGetPlayerOtherInfo{} }
func (m *S2KGetPlayerOtherInfo) String() string { return proto.CompactTextString(m) }
func (*S2KGetPlayerOtherInfo) ProtoMessage()    {}
func (m *S2KGetPlayerOtherInfo) GetId() uint16  { return PCK_S2KGetPlayerOtherInfo_ID }

const PCK_K2SAddItem_ID = 8007 //
//
type K2SAddItem struct {
	//
	Items []*ItemData `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty"`
	//
	Wait bool `protobuf:"varint,2,opt,name=Wait,proto3" json:"Wait,omitempty"`
	//
	ChangType int32 `protobuf:"varint,3,opt,name=ChangType,proto3" json:"ChangType,omitempty"`
	//
	IntParam []int64 `protobuf:"varint,4,rep,name=IntParam,proto3" json:"IntParam,omitempty"`
	//
	StrParam []string `protobuf:"bytes,5,rep,name=StrParam,proto3" json:"StrParam,omitempty"`
	//
	UId int64 `protobuf:"varint,6,opt,name=UId,proto3" json:"UId,omitempty"`
}

func (m *K2SAddItem) Reset()         { *m = K2SAddItem{} }
func (m *K2SAddItem) String() string { return proto.CompactTextString(m) }
func (*K2SAddItem) ProtoMessage()    {}
func (m *K2SAddItem) GetId() uint16  { return PCK_K2SAddItem_ID }

const PCK_K2SAddRedPkg_ID = 8019 //
//
type K2SAddRedPkg struct {
	//
	Type int32 `protobuf:"varint,1,opt,name=Type,proto3" json:"Type,omitempty"`
	//
	Money int64 `protobuf:"varint,2,opt,name=Money,proto3" json:"Money,omitempty"`
}

func (m *K2SAddRedPkg) Reset()         { *m = K2SAddRedPkg{} }
func (m *K2SAddRedPkg) String() string { return proto.CompactTextString(m) }
func (*K2SAddRedPkg) ProtoMessage()    {}
func (m *K2SAddRedPkg) GetId() uint16  { return PCK_K2SAddRedPkg_ID }

const PCK_K2SChangeAttr_ID = 8008 //
//
type K2SChangeAttr struct {
	//
	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	A []*IntAttr `protobuf:"bytes,2,rep,name=A,proto3" json:"A,omitempty"`
	//
	B []*StrAttr `protobuf:"bytes,3,rep,name=B,proto3" json:"B,omitempty"`
}

func (m *K2SChangeAttr) Reset()         { *m = K2SChangeAttr{} }
func (m *K2SChangeAttr) String() string { return proto.CompactTextString(m) }
func (*K2SChangeAttr) ProtoMessage()    {}
func (m *K2SChangeAttr) GetId() uint16  { return PCK_K2SChangeAttr_ID }

const PCK_K2SUnlockSkin_ID = 8012 //
//
type K2SUnlockSkin struct {
	//
	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	SkinId int32 `protobuf:"varint,2,opt,name=SkinId,proto3" json:"SkinId,omitempty"`
}

func (m *K2SUnlockSkin) Reset()         { *m = K2SUnlockSkin{} }
func (m *K2SUnlockSkin) String() string { return proto.CompactTextString(m) }
func (*K2SUnlockSkin) ProtoMessage()    {}
func (m *K2SUnlockSkin) GetId() uint16  { return PCK_K2SUnlockSkin_ID }

const PCK_K2SDeleteSkin_ID = 8013 //
//
type K2SDeleteSkin struct {
	//
	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	SkinId int32 `protobuf:"varint,2,opt,name=SkinId,proto3" json:"SkinId,omitempty"`
}

func (m *K2SDeleteSkin) Reset()         { *m = K2SDeleteSkin{} }
func (m *K2SDeleteSkin) String() string { return proto.CompactTextString(m) }
func (*K2SDeleteSkin) ProtoMessage()    {}
func (m *K2SDeleteSkin) GetId() uint16  { return PCK_K2SDeleteSkin_ID }

const PCK_K2SMail_ID = 8014 //
//
type K2SMail struct {
	//
	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	TmpId int32 `protobuf:"varint,2,opt,name=TmpId,proto3" json:"TmpId,omitempty"`
	//
	Itemdata []*ItemData `protobuf:"bytes,3,rep,name=Itemdata,proto3" json:"Itemdata,omitempty"`
	//
	ChangeType int32 `protobuf:"varint,4,opt,name=ChangeType,proto3" json:"ChangeType,omitempty"`
	//
	Param string `protobuf:"bytes,5,opt,name=Param,proto3" json:"Param,omitempty"`
}

func (m *K2SMail) Reset()         { *m = K2SMail{} }
func (m *K2SMail) String() string { return proto.CompactTextString(m) }
func (*K2SMail) ProtoMessage()    {}
func (m *K2SMail) GetId() uint16  { return PCK_K2SMail_ID }

const PCK_K2SActState_ID = 8015 //
//
type K2SActState struct {
	//
	ActId int32 `protobuf:"varint,1,opt,name=ActId,proto3" json:"ActId,omitempty"`
	//
	State int32 `protobuf:"varint,2,opt,name=State,proto3" json:"State,omitempty"`
	//
	OverTime int64 `protobuf:"varint,3,opt,name=OverTime,proto3" json:"OverTime,omitempty"`
}

func (m *K2SActState) Reset()         { *m = K2SActState{} }
func (m *K2SActState) String() string { return proto.CompactTextString(m) }
func (*K2SActState) ProtoMessage()    {}
func (m *K2SActState) GetId() uint16  { return PCK_K2SActState_ID }

const PCK_K2SBroadcastGang_ID = 8016 //
//
type K2SBroadcastGang struct {
	//
	GangId string `protobuf:"bytes,1,opt,name=GangId,proto3" json:"GangId,omitempty"`
	//
	Pck []byte `protobuf:"bytes,2,opt,name=Pck,proto3" json:"Pck,omitempty"`
}

func (m *K2SBroadcastGang) Reset()         { *m = K2SBroadcastGang{} }
func (m *K2SBroadcastGang) String() string { return proto.CompactTextString(m) }
func (*K2SBroadcastGang) ProtoMessage()    {}
func (m *K2SBroadcastGang) GetId() uint16  { return PCK_K2SBroadcastGang_ID }

const PCK_C2SEndFight_ID = 102 //
//
type C2SEndFight struct {
	//
	Idx int64 `protobuf:"varint,1,opt,name=Idx,proto3" json:"Idx,omitempty"`
}

func (m *C2SEndFight) Reset()         { *m = C2SEndFight{} }
func (m *C2SEndFight) String() string { return proto.CompactTextString(m) }
func (*C2SEndFight) ProtoMessage()    {}
func (m *C2SEndFight) GetId() uint16  { return PCK_C2SEndFight_ID }

const PCK_S2CPrizeReport_ID = 100 //
//
type S2CPrizeReport struct {
	//
	Items []*ItemData `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty"`
	//
	Idx int64 `protobuf:"varint,2,opt,name=Idx,proto3" json:"Idx,omitempty"`
	//
	Type int64 `protobuf:"varint,3,opt,name=Type,proto3" json:"Type,omitempty"`
	//
	Star int64 `protobuf:"varint,4,opt,name=Star,proto3" json:"Star,omitempty"`
	//
	JJC int64 `protobuf:"varint,5,opt,name=JJC,proto3" json:"JJC,omitempty"`
	//
	FBType int32 `protobuf:"varint,6,opt,name=FBType,proto3" json:"FBType,omitempty"`
	//
	FightNick string `protobuf:"bytes,7,opt,name=FightNick,proto3" json:"FightNick,omitempty"`
	//
	Score int32 `protobuf:"varint,8,opt,name=Score,proto3" json:"Score,omitempty"`
	//
	WinTimes int32 `protobuf:"varint,9,opt,name=WinTimes,proto3" json:"WinTimes,omitempty"`
}

func (m *S2CPrizeReport) Reset()         { *m = S2CPrizeReport{} }
func (m *S2CPrizeReport) String() string { return proto.CompactTextString(m) }
func (*S2CPrizeReport) ProtoMessage()    {}
func (m *S2CPrizeReport) GetId() uint16  { return PCK_S2CPrizeReport_ID }

const PCK_C2SStageFight_ID = 103 //
//
type C2SStageFight struct {
}

func (m *C2SStageFight) Reset()         { *m = C2SStageFight{} }
func (m *C2SStageFight) String() string { return proto.CompactTextString(m) }
func (*C2SStageFight) ProtoMessage()    {}
func (m *C2SStageFight) GetId() uint16  { return PCK_C2SStageFight_ID }

const PCK_S2CStageFight_ID = 104 //
//
type S2CStageFight struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CStageFight) Reset()         { *m = S2CStageFight{} }
func (m *S2CStageFight) String() string { return proto.CompactTextString(m) }
func (*S2CStageFight) ProtoMessage()    {}
func (m *S2CStageFight) GetId() uint16  { return PCK_S2CStageFight_ID }

const PCK_C2SAutoStage_ID = 105 //
//
type C2SAutoStage struct {
	//
	State int32 `protobuf:"varint,1,opt,name=State,proto3" json:"State,omitempty"`
}

func (m *C2SAutoStage) Reset()         { *m = C2SAutoStage{} }
func (m *C2SAutoStage) String() string { return proto.CompactTextString(m) }
func (*C2SAutoStage) ProtoMessage()    {}
func (m *C2SAutoStage) GetId() uint16  { return PCK_C2SAutoStage_ID }

const PCK_C2SStageSeek_ID = 106 //
//
type C2SStageSeek struct {
}

func (m *C2SStageSeek) Reset()         { *m = C2SStageSeek{} }
func (m *C2SStageSeek) String() string { return proto.CompactTextString(m) }
func (*C2SStageSeek) ProtoMessage()    {}
func (m *C2SStageSeek) GetId() uint16  { return PCK_C2SStageSeek_ID }

const PCK_S2CStageSeek_ID = 107 //
//
type S2CStageSeek struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	UserId int64 `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	StageId int64 `protobuf:"varint,3,opt,name=StageId,proto3" json:"StageId,omitempty"`
}

func (m *S2CStageSeek) Reset()         { *m = S2CStageSeek{} }
func (m *S2CStageSeek) String() string { return proto.CompactTextString(m) }
func (*S2CStageSeek) ProtoMessage()    {}
func (m *S2CStageSeek) GetId() uint16  { return PCK_S2CStageSeek_ID }

const PCK_C2SStageHelp_ID = 108 //
//
type C2SStageHelp struct {
	//
	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	StageId int64 `protobuf:"varint,2,opt,name=StageId,proto3" json:"StageId,omitempty"`
}

func (m *C2SStageHelp) Reset()         { *m = C2SStageHelp{} }
func (m *C2SStageHelp) String() string { return proto.CompactTextString(m) }
func (*C2SStageHelp) ProtoMessage()    {}
func (m *C2SStageHelp) GetId() uint16  { return PCK_C2SStageHelp_ID }

const PCK_S2CStageHelp_ID = 109 //
//
type S2CStageHelp struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CStageHelp) Reset()         { *m = S2CStageHelp{} }
func (m *S2CStageHelp) String() string { return proto.CompactTextString(m) }
func (*S2CStageHelp) ProtoMessage()    {}
func (m *S2CStageHelp) GetId() uint16  { return PCK_S2CStageHelp_ID }

const PCK_C2SChangeMap_ID = 50 //
//
type C2SChangeMap struct {
	//
	MapId int64 `protobuf:"varint,1,opt,name=MapId,proto3" json:"MapId,omitempty"`
}

func (m *C2SChangeMap) Reset()         { *m = C2SChangeMap{} }
func (m *C2SChangeMap) String() string { return proto.CompactTextString(m) }
func (*C2SChangeMap) ProtoMessage()    {}
func (m *C2SChangeMap) GetId() uint16  { return PCK_C2SChangeMap_ID }

const PCK_S2CChangeMap_ID = 51 //
//
type S2CChangeMap struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	MapId int64 `protobuf:"varint,2,opt,name=MapId,proto3" json:"MapId,omitempty"`
	//
	X int32 `protobuf:"varint,3,opt,name=X,proto3" json:"X,omitempty"`
	//
	Y int32 `protobuf:"varint,4,opt,name=Y,proto3" json:"Y,omitempty"`
}

func (m *S2CChangeMap) Reset()         { *m = S2CChangeMap{} }
func (m *S2CChangeMap) String() string { return proto.CompactTextString(m) }
func (*S2CChangeMap) ProtoMessage()    {}
func (m *S2CChangeMap) GetId() uint16  { return PCK_S2CChangeMap_ID }

const PCK_C2SAutoJump_ID = 41 //
//
type C2SAutoJump struct {
	//
	V int32 `protobuf:"varint,1,opt,name=V,proto3" json:"V,omitempty"`
}

func (m *C2SAutoJump) Reset()         { *m = C2SAutoJump{} }
func (m *C2SAutoJump) String() string { return proto.CompactTextString(m) }
func (*C2SAutoJump) ProtoMessage()    {}
func (m *C2SAutoJump) GetId() uint16  { return PCK_C2SAutoJump_ID }

//
type StagePrizeInfo struct {
	//
	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Num int64 `protobuf:"varint,2,opt,name=Num,proto3" json:"Num,omitempty"`
}

func (m *StagePrizeInfo) Reset()         { *m = StagePrizeInfo{} }
func (m *StagePrizeInfo) String() string { return proto.CompactTextString(m) }
func (*StagePrizeInfo) ProtoMessage()    {}

const PCK_C2SStagePrize_ID = 115 //
//
type C2SStagePrize struct {
}

func (m *C2SStagePrize) Reset()         { *m = C2SStagePrize{} }
func (m *C2SStagePrize) String() string { return proto.CompactTextString(m) }
func (*C2SStagePrize) ProtoMessage()    {}
func (m *C2SStagePrize) GetId() uint16  { return PCK_C2SStagePrize_ID }

const PCK_S2CStagePrize_ID = 110 //
//
type S2CStagePrize struct {
	//
	Items []*StagePrizeInfo `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty"`
}

func (m *S2CStagePrize) Reset()         { *m = S2CStagePrize{} }
func (m *S2CStagePrize) String() string { return proto.CompactTextString(m) }
func (*S2CStagePrize) ProtoMessage()    {}
func (m *S2CStagePrize) GetId() uint16  { return PCK_S2CStagePrize_ID }

const PCK_C2SGetStagePrize_ID = 111 //
//
type C2SGetStagePrize struct {
	//
	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2SGetStagePrize) Reset()         { *m = C2SGetStagePrize{} }
func (m *C2SGetStagePrize) String() string { return proto.CompactTextString(m) }
func (*C2SGetStagePrize) ProtoMessage()    {}
func (m *C2SGetStagePrize) GetId() uint16  { return PCK_C2SGetStagePrize_ID }

const PCK_S2CGetStagePrize_ID = 112 //
//
type S2CGetStagePrize struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id int64 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *S2CGetStagePrize) Reset()         { *m = S2CGetStagePrize{} }
func (m *S2CGetStagePrize) String() string { return proto.CompactTextString(m) }
func (*S2CGetStagePrize) ProtoMessage()    {}
func (m *S2CGetStagePrize) GetId() uint16  { return PCK_S2CGetStagePrize_ID }

const PCK_C2SGetAllStagePrize_ID = 116 //
//
type C2SGetAllStagePrize struct {
}

func (m *C2SGetAllStagePrize) Reset()         { *m = C2SGetAllStagePrize{} }
func (m *C2SGetAllStagePrize) String() string { return proto.CompactTextString(m) }
func (*C2SGetAllStagePrize) ProtoMessage()    {}
func (m *C2SGetAllStagePrize) GetId() uint16  { return PCK_C2SGetAllStagePrize_ID }

const PCK_S2CGetAllStagePrize_ID = 117 //
//
type S2CGetAllStagePrize struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CGetAllStagePrize) Reset()         { *m = S2CGetAllStagePrize{} }
func (m *S2CGetAllStagePrize) String() string { return proto.CompactTextString(m) }
func (*S2CGetAllStagePrize) ProtoMessage()    {}
func (m *S2CGetAllStagePrize) GetId() uint16  { return PCK_S2CGetAllStagePrize_ID }

const PCK_C2SCatchPet_ID = 113 //
//
type C2SCatchPet struct {
}

func (m *C2SCatchPet) Reset()         { *m = C2SCatchPet{} }
func (m *C2SCatchPet) String() string { return proto.CompactTextString(m) }
func (*C2SCatchPet) ProtoMessage()    {}
func (m *C2SCatchPet) GetId() uint16  { return PCK_C2SCatchPet_ID }

const PCK_S2CCatchPet_ID = 114 //
//
type S2CCatchPet struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Win int32 `protobuf:"varint,2,opt,name=Win,proto3" json:"Win,omitempty"`
	//
	ItemId int32 `protobuf:"varint,3,opt,name=ItemId,proto3" json:"ItemId,omitempty"`
	//
	MonsterId int32 `protobuf:"varint,4,opt,name=MonsterId,proto3" json:"MonsterId,omitempty"`
}

func (m *S2CCatchPet) Reset()         { *m = S2CCatchPet{} }
func (m *S2CCatchPet) String() string { return proto.CompactTextString(m) }
func (*S2CCatchPet) ProtoMessage()    {}
func (m *S2CCatchPet) GetId() uint16  { return PCK_S2CCatchPet_ID }

const PCK_C2SStartMove_ID = 52 //
//
type C2SStartMove struct {
	//
	P []int32 `protobuf:"varint,1,rep,name=P,proto3" json:"P,omitempty"`
}

func (m *C2SStartMove) Reset()         { *m = C2SStartMove{} }
func (m *C2SStartMove) String() string { return proto.CompactTextString(m) }
func (*C2SStartMove) ProtoMessage()    {}
func (m *C2SStartMove) GetId() uint16  { return PCK_C2SStartMove_ID }

const PCK_S2CStartMove_ID = 53 //
//
type S2CStartMove struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CStartMove) Reset()         { *m = S2CStartMove{} }
func (m *S2CStartMove) String() string { return proto.CompactTextString(m) }
func (*S2CStartMove) ProtoMessage()    {}
func (m *S2CStartMove) GetId() uint16  { return PCK_S2CStartMove_ID }

const PCK_S2CPlayerMove_ID = 54 //
//
type S2CPlayerMove struct {
	//
	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	P []int32 `protobuf:"varint,2,rep,name=P,proto3" json:"P,omitempty"`
}

func (m *S2CPlayerMove) Reset()         { *m = S2CPlayerMove{} }
func (m *S2CPlayerMove) String() string { return proto.CompactTextString(m) }
func (*S2CPlayerMove) ProtoMessage()    {}
func (m *S2CPlayerMove) GetId() uint16  { return PCK_S2CPlayerMove_ID }

const PCK_C2SStopMove_ID = 55 //
//
type C2SStopMove struct {
	//
	X int32 `protobuf:"varint,1,opt,name=X,proto3" json:"X,omitempty"`
	//
	Y int32 `protobuf:"varint,2,opt,name=Y,proto3" json:"Y,omitempty"`
}

func (m *C2SStopMove) Reset()         { *m = C2SStopMove{} }
func (m *C2SStopMove) String() string { return proto.CompactTextString(m) }
func (*C2SStopMove) ProtoMessage()    {}
func (m *C2SStopMove) GetId() uint16  { return PCK_C2SStopMove_ID }

const PCK_S2CStopMove_ID = 67 //
//
type S2CStopMove struct {
}

func (m *S2CStopMove) Reset()         { *m = S2CStopMove{} }
func (m *S2CStopMove) String() string { return proto.CompactTextString(m) }
func (*S2CStopMove) ProtoMessage()    {}
func (m *S2CStopMove) GetId() uint16  { return PCK_S2CStopMove_ID }

const PCK_S2CPlayerStopMove_ID = 56 //
//
type S2CPlayerStopMove struct {
	//
	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	X int32 `protobuf:"varint,2,opt,name=X,proto3" json:"X,omitempty"`
	//
	Y int32 `protobuf:"varint,3,opt,name=Y,proto3" json:"Y,omitempty"`
}

func (m *S2CPlayerStopMove) Reset()         { *m = S2CPlayerStopMove{} }
func (m *S2CPlayerStopMove) String() string { return proto.CompactTextString(m) }
func (*S2CPlayerStopMove) ProtoMessage()    {}
func (m *S2CPlayerStopMove) GetId() uint16  { return PCK_S2CPlayerStopMove_ID }

const PCK_S2CTransfer_ID = 64 //
//
type S2CTransfer struct {
	//
	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	X int32 `protobuf:"varint,2,opt,name=X,proto3" json:"X,omitempty"`
	//
	Y int32 `protobuf:"varint,3,opt,name=Y,proto3" json:"Y,omitempty"`
}

func (m *S2CTransfer) Reset()         { *m = S2CTransfer{} }
func (m *S2CTransfer) String() string { return proto.CompactTextString(m) }
func (*S2CTransfer) ProtoMessage()    {}
func (m *S2CTransfer) GetId() uint16  { return PCK_S2CTransfer_ID }

const PCK_C2SCheckFight_ID = 65 //
//
type C2SCheckFight struct {
}

func (m *C2SCheckFight) Reset()         { *m = C2SCheckFight{} }
func (m *C2SCheckFight) String() string { return proto.CompactTextString(m) }
func (*C2SCheckFight) ProtoMessage()    {}
func (m *C2SCheckFight) GetId() uint16  { return PCK_C2SCheckFight_ID }

const PCK_S2CCheckFight_ID = 66 //
//
type S2CCheckFight struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	NextTime int64 `protobuf:"varint,2,opt,name=NextTime,proto3" json:"NextTime,omitempty"`
}

func (m *S2CCheckFight) Reset()         { *m = S2CCheckFight{} }
func (m *S2CCheckFight) String() string { return proto.CompactTextString(m) }
func (*S2CCheckFight) ProtoMessage()    {}
func (m *S2CCheckFight) GetId() uint16  { return PCK_S2CCheckFight_ID }

const PCK_S2CPlayerEnterMap_ID = 57 //
//
type S2CPlayerEnterMap struct {
	//
	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	A []*IntAttr `protobuf:"bytes,2,rep,name=A,proto3" json:"A,omitempty"`
	//
	B []*StrAttr `protobuf:"bytes,3,rep,name=B,proto3" json:"B,omitempty"`
	//
	X int32 `protobuf:"varint,4,opt,name=X,proto3" json:"X,omitempty"`
	//
	Y int32 `protobuf:"varint,5,opt,name=Y,proto3" json:"Y,omitempty"`
	//
	P []int32 `protobuf:"varint,6,rep,name=P,proto3" json:"P,omitempty"`
	//
	MT int32 `protobuf:"varint,7,opt,name=MT,proto3" json:"MT,omitempty"`
}

func (m *S2CPlayerEnterMap) Reset()         { *m = S2CPlayerEnterMap{} }
func (m *S2CPlayerEnterMap) String() string { return proto.CompactTextString(m) }
func (*S2CPlayerEnterMap) ProtoMessage()    {}
func (m *S2CPlayerEnterMap) GetId() uint16  { return PCK_S2CPlayerEnterMap_ID }

const PCK_S2CPlayerLeaveMap_ID = 58 //
//
type S2CPlayerLeaveMap struct {
	//
	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (m *S2CPlayerLeaveMap) Reset()         { *m = S2CPlayerLeaveMap{} }
func (m *S2CPlayerLeaveMap) String() string { return proto.CompactTextString(m) }
func (*S2CPlayerLeaveMap) ProtoMessage()    {}
func (m *S2CPlayerLeaveMap) GetId() uint16  { return PCK_S2CPlayerLeaveMap_ID }

const PCK_S2CMonsterEnterMap_ID = 59 //
//
type S2CMonsterEnterMap struct {
	//
	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	IID int64 `protobuf:"varint,2,opt,name=IID,proto3" json:"IID,omitempty"`
	//
	X int64 `protobuf:"varint,3,opt,name=X,proto3" json:"X,omitempty"`
	//
	Y int64 `protobuf:"varint,4,opt,name=Y,proto3" json:"Y,omitempty"`
	//
	MHp int64 `protobuf:"varint,5,opt,name=MHp,proto3" json:"MHp,omitempty"`
	//
	Hp int64 `protobuf:"varint,6,opt,name=Hp,proto3" json:"Hp,omitempty"`
	//
	MHp1 int64 `protobuf:"varint,7,opt,name=MHp1,proto3" json:"MHp1,omitempty"`
	//
	Hp1 int64 `protobuf:"varint,8,opt,name=Hp1,proto3" json:"Hp1,omitempty"`
	//
	Fx int32 `protobuf:"varint,9,opt,name=Fx,proto3" json:"Fx,omitempty"`
}

func (m *S2CMonsterEnterMap) Reset()         { *m = S2CMonsterEnterMap{} }
func (m *S2CMonsterEnterMap) String() string { return proto.CompactTextString(m) }
func (*S2CMonsterEnterMap) ProtoMessage()    {}
func (m *S2CMonsterEnterMap) GetId() uint16  { return PCK_S2CMonsterEnterMap_ID }

const PCK_S2CUpMonsterInfo_ID = 63 //
//
type S2CUpMonsterInfo struct {
	//
	MHp int64 `protobuf:"varint,1,opt,name=MHp,proto3" json:"MHp,omitempty"`
	//
	Hp int64 `protobuf:"varint,2,opt,name=Hp,proto3" json:"Hp,omitempty"`
	//
	MHp1 int64 `protobuf:"varint,3,opt,name=MHp1,proto3" json:"MHp1,omitempty"`
	//
	Hp1 int64 `protobuf:"varint,4,opt,name=Hp1,proto3" json:"Hp1,omitempty"`
	//
	Id int64 `protobuf:"varint,5,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	LifeState int32 `protobuf:"varint,6,opt,name=LifeState,proto3" json:"LifeState,omitempty"`
	//
	IsCollect int32 `protobuf:"varint,7,opt,name=IsCollect,proto3" json:"IsCollect,omitempty"`
}

func (m *S2CUpMonsterInfo) Reset()         { *m = S2CUpMonsterInfo{} }
func (m *S2CUpMonsterInfo) String() string { return proto.CompactTextString(m) }
func (*S2CUpMonsterInfo) ProtoMessage()    {}
func (m *S2CUpMonsterInfo) GetId() uint16  { return PCK_S2CUpMonsterInfo_ID }

const PCK_S2CMonsterLeaveMap_ID = 60 //
//
type S2CMonsterLeaveMap struct {
	//
	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *S2CMonsterLeaveMap) Reset()         { *m = S2CMonsterLeaveMap{} }
func (m *S2CMonsterLeaveMap) String() string { return proto.CompactTextString(m) }
func (*S2CMonsterLeaveMap) ProtoMessage()    {}
func (m *S2CMonsterLeaveMap) GetId() uint16  { return PCK_S2CMonsterLeaveMap_ID }

const PCK_C2SStartFight_ID = 61 //
//
type C2SStartFight struct {
	//
	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Type int64 `protobuf:"varint,2,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (m *C2SStartFight) Reset()         { *m = C2SStartFight{} }
func (m *C2SStartFight) String() string { return proto.CompactTextString(m) }
func (*C2SStartFight) ProtoMessage()    {}
func (m *C2SStartFight) GetId() uint16  { return PCK_C2SStartFight_ID }

const PCK_S2CStartFight_ID = 62 //
//
type S2CStartFight struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CStartFight) Reset()         { *m = S2CStartFight{} }
func (m *S2CStartFight) String() string { return proto.CompactTextString(m) }
func (*S2CStartFight) ProtoMessage()    {}
func (m *S2CStartFight) GetId() uint16  { return PCK_S2CStartFight_ID }

const PCK_C2SStartCollect_ID = 93 //
//
type C2SStartCollect struct {
	//
	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2SStartCollect) Reset()         { *m = C2SStartCollect{} }
func (m *C2SStartCollect) String() string { return proto.CompactTextString(m) }
func (*C2SStartCollect) ProtoMessage()    {}
func (m *C2SStartCollect) GetId() uint16  { return PCK_C2SStartCollect_ID }

const PCK_S2CStartCollect_ID = 94 //
//
type S2CStartCollect struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CStartCollect) Reset()         { *m = S2CStartCollect{} }
func (m *S2CStartCollect) String() string { return proto.CompactTextString(m) }
func (*S2CStartCollect) ProtoMessage()    {}
func (m *S2CStartCollect) GetId() uint16  { return PCK_S2CStartCollect_ID }

const PCK_C2SEndCollect_ID = 95 //
//
type C2SEndCollect struct {
	//
	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Ret int32 `protobuf:"varint,2,opt,name=Ret,proto3" json:"Ret,omitempty"`
}

func (m *C2SEndCollect) Reset()         { *m = C2SEndCollect{} }
func (m *C2SEndCollect) String() string { return proto.CompactTextString(m) }
func (*C2SEndCollect) ProtoMessage()    {}
func (m *C2SEndCollect) GetId() uint16  { return PCK_C2SEndCollect_ID }

const PCK_S2CEndCollect_ID = 96 //
//
type S2CEndCollect struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Ret int32 `protobuf:"varint,2,opt,name=Ret,proto3" json:"Ret,omitempty"`
}

func (m *S2CEndCollect) Reset()         { *m = S2CEndCollect{} }
func (m *S2CEndCollect) String() string { return proto.CompactTextString(m) }
func (*S2CEndCollect) ProtoMessage()    {}
func (m *S2CEndCollect) GetId() uint16  { return PCK_S2CEndCollect_ID }

const PCK_S2CMonsterTalk_ID = 98 //
//
type S2CMonsterTalk struct {
	//
	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	TalkId int32 `protobuf:"varint,2,opt,name=TalkId,proto3" json:"TalkId,omitempty"`
}

func (m *S2CMonsterTalk) Reset()         { *m = S2CMonsterTalk{} }
func (m *S2CMonsterTalk) String() string { return proto.CompactTextString(m) }
func (*S2CMonsterTalk) ProtoMessage()    {}
func (m *S2CMonsterTalk) GetId() uint16  { return PCK_S2CMonsterTalk_ID }

//
type Skill struct {
	//
	I int32 `protobuf:"varint,1,opt,name=I,proto3" json:"I,omitempty"`
	//
	L int32 `protobuf:"varint,2,opt,name=L,proto3" json:"L,omitempty"`
}

func (m *Skill) Reset()         { *m = Skill{} }
func (m *Skill) String() string { return proto.CompactTextString(m) }
func (*Skill) ProtoMessage()    {}

const PCK_C2SRoleSkillUp_ID = 71 //
//
type C2SRoleSkillUp struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2SRoleSkillUp) Reset()         { *m = C2SRoleSkillUp{} }
func (m *C2SRoleSkillUp) String() string { return proto.CompactTextString(m) }
func (*C2SRoleSkillUp) ProtoMessage()    {}
func (m *C2SRoleSkillUp) GetId() uint16  { return PCK_C2SRoleSkillUp_ID }

const PCK_C2SRoleSkillUpAuto_ID = 72 //
//
type C2SRoleSkillUpAuto struct {
}

func (m *C2SRoleSkillUpAuto) Reset()         { *m = C2SRoleSkillUpAuto{} }
func (m *C2SRoleSkillUpAuto) String() string { return proto.CompactTextString(m) }
func (*C2SRoleSkillUpAuto) ProtoMessage()    {}
func (m *C2SRoleSkillUpAuto) GetId() uint16  { return PCK_C2SRoleSkillUpAuto_ID }

const PCK_S2CRoleSkillUp_ID = 73 //
//
type S2CRoleSkillUp struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Skill []*Skill `protobuf:"bytes,2,rep,name=Skill,proto3" json:"Skill,omitempty"`
}

func (m *S2CRoleSkillUp) Reset()         { *m = S2CRoleSkillUp{} }
func (m *S2CRoleSkillUp) String() string { return proto.CompactTextString(m) }
func (*S2CRoleSkillUp) ProtoMessage()    {}
func (m *S2CRoleSkillUp) GetId() uint16  { return PCK_S2CRoleSkillUp_ID }

const PCK_C2SRoleSkillOrder_ID = 74 //
//
type C2SRoleSkillOrder struct {
	//
	Skill []*Skill `protobuf:"bytes,1,rep,name=Skill,proto3" json:"Skill,omitempty"`
}

func (m *C2SRoleSkillOrder) Reset()         { *m = C2SRoleSkillOrder{} }
func (m *C2SRoleSkillOrder) String() string { return proto.CompactTextString(m) }
func (*C2SRoleSkillOrder) ProtoMessage()    {}
func (m *C2SRoleSkillOrder) GetId() uint16  { return PCK_C2SRoleSkillOrder_ID }

const PCK_S2CRoleSkillOrder_ID = 75 //
//
type S2CRoleSkillOrder struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CRoleSkillOrder) Reset()         { *m = S2CRoleSkillOrder{} }
func (m *S2CRoleSkillOrder) String() string { return proto.CompactTextString(m) }
func (*S2CRoleSkillOrder) ProtoMessage()    {}
func (m *S2CRoleSkillOrder) GetId() uint16  { return PCK_S2CRoleSkillOrder_ID }

const PCK_C2SUnlockSkin_ID = 81 //
//
type C2SUnlockSkin struct {
	//
	SkinId int32 `protobuf:"varint,1,opt,name=SkinId,proto3" json:"SkinId,omitempty"`
}

func (m *C2SUnlockSkin) Reset()         { *m = C2SUnlockSkin{} }
func (m *C2SUnlockSkin) String() string { return proto.CompactTextString(m) }
func (*C2SUnlockSkin) ProtoMessage()    {}
func (m *C2SUnlockSkin) GetId() uint16  { return PCK_C2SUnlockSkin_ID }

const PCK_S2CUnlockSkin_ID = 82 //
//
type S2CUnlockSkin struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Skins []*Skin `protobuf:"bytes,2,rep,name=Skins,proto3" json:"Skins,omitempty"`
}

func (m *S2CUnlockSkin) Reset()         { *m = S2CUnlockSkin{} }
func (m *S2CUnlockSkin) String() string { return proto.CompactTextString(m) }
func (*S2CUnlockSkin) ProtoMessage()    {}
func (m *S2CUnlockSkin) GetId() uint16  { return PCK_S2CUnlockSkin_ID }

const PCK_C2SWearSkin_ID = 83 //
//
type C2SWearSkin struct {
	//
	SkinId int32 `protobuf:"varint,1,opt,name=SkinId,proto3" json:"SkinId,omitempty"`
}

func (m *C2SWearSkin) Reset()         { *m = C2SWearSkin{} }
func (m *C2SWearSkin) String() string { return proto.CompactTextString(m) }
func (*C2SWearSkin) ProtoMessage()    {}
func (m *C2SWearSkin) GetId() uint16  { return PCK_C2SWearSkin_ID }

const PCK_S2CWearSkin_ID = 84 //
//
type S2CWearSkin struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	SkinId int32 `protobuf:"varint,2,opt,name=SkinId,proto3" json:"SkinId,omitempty"`
}

func (m *S2CWearSkin) Reset()         { *m = S2CWearSkin{} }
func (m *S2CWearSkin) String() string { return proto.CompactTextString(m) }
func (*S2CWearSkin) ProtoMessage()    {}
func (m *S2CWearSkin) GetId() uint16  { return PCK_S2CWearSkin_ID }

const PCK_C2SSkinUp_ID = 127 //
//
type C2SSkinUp struct {
	//
	SkinId int32 `protobuf:"varint,1,opt,name=SkinId,proto3" json:"SkinId,omitempty"`
}

func (m *C2SSkinUp) Reset()         { *m = C2SSkinUp{} }
func (m *C2SSkinUp) String() string { return proto.CompactTextString(m) }
func (*C2SSkinUp) ProtoMessage()    {}
func (m *C2SSkinUp) GetId() uint16  { return PCK_C2SSkinUp_ID }

const PCK_S2CSkinUp_ID = 128 //
//
type S2CSkinUp struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	SkinId int32 `protobuf:"varint,2,opt,name=SkinId,proto3" json:"SkinId,omitempty"`
	//
	SkinLevel int32 `protobuf:"varint,3,opt,name=SkinLevel,proto3" json:"SkinLevel,omitempty"`
}

func (m *S2CSkinUp) Reset()         { *m = S2CSkinUp{} }
func (m *S2CSkinUp) String() string { return proto.CompactTextString(m) }
func (*S2CSkinUp) ProtoMessage()    {}
func (m *S2CSkinUp) GetId() uint16  { return PCK_S2CSkinUp_ID }

const PCK_C2SGetVipPrize_ID = 85 //
//
type C2SGetVipPrize struct {
	//
	VIPLevel int32 `protobuf:"varint,1,opt,name=VIPLevel,proto3" json:"VIPLevel,omitempty"`
}

func (m *C2SGetVipPrize) Reset()         { *m = C2SGetVipPrize{} }
func (m *C2SGetVipPrize) String() string { return proto.CompactTextString(m) }
func (*C2SGetVipPrize) ProtoMessage()    {}
func (m *C2SGetVipPrize) GetId() uint16  { return PCK_C2SGetVipPrize_ID }

const PCK_S2CGetVipPrize_ID = 86 //
//
type S2CGetVipPrize struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	VIPLevel int32 `protobuf:"varint,2,opt,name=VIPLevel,proto3" json:"VIPLevel,omitempty"`
}

func (m *S2CGetVipPrize) Reset()         { *m = S2CGetVipPrize{} }
func (m *S2CGetVipPrize) String() string { return proto.CompactTextString(m) }
func (*S2CGetVipPrize) ProtoMessage()    {}
func (m *S2CGetVipPrize) GetId() uint16  { return PCK_S2CGetVipPrize_ID }

const PCK_C2SBuyVipPrize_ID = 130 //
//
type C2SBuyVipPrize struct {
	//
	VIPLevel int32 `protobuf:"varint,1,opt,name=VIPLevel,proto3" json:"VIPLevel,omitempty"`
}

func (m *C2SBuyVipPrize) Reset()         { *m = C2SBuyVipPrize{} }
func (m *C2SBuyVipPrize) String() string { return proto.CompactTextString(m) }
func (*C2SBuyVipPrize) ProtoMessage()    {}
func (m *C2SBuyVipPrize) GetId() uint16  { return PCK_C2SBuyVipPrize_ID }

const PCK_S2CBuyVipPrize_ID = 131 //
//
type S2CBuyVipPrize struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	VIPLevel int32 `protobuf:"varint,2,opt,name=VIPLevel,proto3" json:"VIPLevel,omitempty"`
}

func (m *S2CBuyVipPrize) Reset()         { *m = S2CBuyVipPrize{} }
func (m *S2CBuyVipPrize) String() string { return proto.CompactTextString(m) }
func (*S2CBuyVipPrize) ProtoMessage()    {}
func (m *S2CBuyVipPrize) GetId() uint16  { return PCK_S2CBuyVipPrize_ID }

const PCK_C2SGetNewPrize_ID = 87 //
//
type C2SGetNewPrize struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Type int32 `protobuf:"varint,2,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (m *C2SGetNewPrize) Reset()         { *m = C2SGetNewPrize{} }
func (m *C2SGetNewPrize) String() string { return proto.CompactTextString(m) }
func (*C2SGetNewPrize) ProtoMessage()    {}
func (m *C2SGetNewPrize) GetId() uint16  { return PCK_C2SGetNewPrize_ID }

const PCK_S2CGetNewPrize_ID = 88 //
//
type S2CGetNewPrize struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Type int32 `protobuf:"varint,3,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (m *S2CGetNewPrize) Reset()         { *m = S2CGetNewPrize{} }
func (m *S2CGetNewPrize) String() string { return proto.CompactTextString(m) }
func (*S2CGetNewPrize) ProtoMessage()    {}
func (m *S2CGetNewPrize) GetId() uint16  { return PCK_S2CGetNewPrize_ID }

const PCK_C2SClientSwitch_ID = 89 //
//
type C2SClientSwitch struct {
	//
	Value int32 `protobuf:"varint,1,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (m *C2SClientSwitch) Reset()         { *m = C2SClientSwitch{} }
func (m *C2SClientSwitch) String() string { return proto.CompactTextString(m) }
func (*C2SClientSwitch) ProtoMessage()    {}
func (m *C2SClientSwitch) GetId() uint16  { return PCK_C2SClientSwitch_ID }

const PCK_S2CClientSwitch_ID = 90 //
//
type S2CClientSwitch struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CClientSwitch) Reset()         { *m = S2CClientSwitch{} }
func (m *S2CClientSwitch) String() string { return proto.CompactTextString(m) }
func (*S2CClientSwitch) ProtoMessage()    {}
func (m *S2CClientSwitch) GetId() uint16  { return PCK_S2CClientSwitch_ID }

const PCK_C2SShareSuccess_ID = 97 //
//
type C2SShareSuccess struct {
}

func (m *C2SShareSuccess) Reset()         { *m = C2SShareSuccess{} }
func (m *C2SShareSuccess) String() string { return proto.CompactTextString(m) }
func (*C2SShareSuccess) ProtoMessage()    {}
func (m *C2SShareSuccess) GetId() uint16  { return PCK_C2SShareSuccess_ID }

const PCK_S2CUserGrade_ID = 210 //
//
type S2CUserGrade struct {
	//
	Grade []*Grade `protobuf:"bytes,1,rep,name=Grade,proto3" json:"Grade,omitempty"`
}

func (m *S2CUserGrade) Reset()         { *m = S2CUserGrade{} }
func (m *S2CUserGrade) String() string { return proto.CompactTextString(m) }
func (*S2CUserGrade) ProtoMessage()    {}
func (m *S2CUserGrade) GetId() uint16  { return PCK_S2CUserGrade_ID }

const PCK_C2SRoleLevelUp_ID = 203 //
//
type C2SRoleLevelUp struct {
}

func (m *C2SRoleLevelUp) Reset()         { *m = C2SRoleLevelUp{} }
func (m *C2SRoleLevelUp) String() string { return proto.CompactTextString(m) }
func (*C2SRoleLevelUp) ProtoMessage()    {}
func (m *C2SRoleLevelUp) GetId() uint16  { return PCK_C2SRoleLevelUp_ID }

const PCK_S2CRoleLevelUp_ID = 204 //
//
type S2CRoleLevelUp struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CRoleLevelUp) Reset()         { *m = S2CRoleLevelUp{} }
func (m *S2CRoleLevelUp) String() string { return proto.CompactTextString(m) }
func (*S2CRoleLevelUp) ProtoMessage()    {}
func (m *S2CRoleLevelUp) GetId() uint16  { return PCK_S2CRoleLevelUp_ID }

//
type RoleGrade struct {
	//
	Grade []*Grade `protobuf:"bytes,1,rep,name=Grade,proto3" json:"Grade,omitempty"`
}

func (m *RoleGrade) Reset()         { *m = RoleGrade{} }
func (m *RoleGrade) String() string { return proto.CompactTextString(m) }
func (*RoleGrade) ProtoMessage()    {}

//
type Grade struct {
	//
	DT int32 `protobuf:"varint,1,opt,name=DT,proto3" json:"DT,omitempty"`
	//
	PT int32 `protobuf:"varint,2,opt,name=PT,proto3" json:"PT,omitempty"`
	//
	L int32 `protobuf:"varint,3,opt,name=L,proto3" json:"L,omitempty"`
	//
	P int32 `protobuf:"varint,4,opt,name=P,proto3" json:"P,omitempty"`
}

func (m *Grade) Reset()         { *m = Grade{} }
func (m *Grade) String() string { return proto.CompactTextString(m) }
func (*Grade) ProtoMessage()    {}

const PCK_C2SGradeUp_ID = 201 //
//
type C2SGradeUp struct {
	//
	GradeType int32 `protobuf:"varint,1,opt,name=GradeType,proto3" json:"GradeType,omitempty"`
	//
	PartType int32 `protobuf:"varint,2,opt,name=PartType,proto3" json:"PartType,omitempty"`
	//
	Times int32 `protobuf:"varint,3,opt,name=Times,proto3" json:"Times,omitempty"`
	//
	AutoBuy int32 `protobuf:"varint,4,opt,name=AutoBuy,proto3" json:"AutoBuy,omitempty"`
}

func (m *C2SGradeUp) Reset()         { *m = C2SGradeUp{} }
func (m *C2SGradeUp) String() string { return proto.CompactTextString(m) }
func (*C2SGradeUp) ProtoMessage()    {}
func (m *C2SGradeUp) GetId() uint16  { return PCK_C2SGradeUp_ID }

const PCK_S2CGradeUp_ID = 202 //
//
type S2CGradeUp struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Grade *Grade `protobuf:"bytes,2,opt,name=Grade,proto3" json:"Grade,omitempty"`
	//
	Multi int32 `protobuf:"varint,3,opt,name=Multi,proto3" json:"Multi,omitempty"`
	//
	New int32 `protobuf:"varint,4,opt,name=New,proto3" json:"New,omitempty"`
}

func (m *S2CGradeUp) Reset()         { *m = S2CGradeUp{} }
func (m *S2CGradeUp) String() string { return proto.CompactTextString(m) }
func (*S2CGradeUp) ProtoMessage()    {}
func (m *S2CGradeUp) GetId() uint16  { return PCK_S2CGradeUp_ID }

const PCK_C2SGrade2_ID = 207 //
//
type C2SGrade2 struct {
	//
	T int32 `protobuf:"varint,1,opt,name=T,proto3" json:"T,omitempty"`
}

func (m *C2SGrade2) Reset()         { *m = C2SGrade2{} }
func (m *C2SGrade2) String() string { return proto.CompactTextString(m) }
func (*C2SGrade2) ProtoMessage()    {}
func (m *C2SGrade2) GetId() uint16  { return PCK_C2SGrade2_ID }

const PCK_S2CGrade2_ID = 208 //
//
type S2CGrade2 struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	T int32 `protobuf:"varint,2,opt,name=T,proto3" json:"T,omitempty"`
}

func (m *S2CGrade2) Reset()         { *m = S2CGrade2{} }
func (m *S2CGrade2) String() string { return proto.CompactTextString(m) }
func (*S2CGrade2) ProtoMessage()    {}
func (m *S2CGrade2) GetId() uint16  { return PCK_S2CGrade2_ID }

const PCK_C2SEquipGradeUp_ID = 205 //
//
type C2SEquipGradeUp struct {
	//
	GradeType int32 `protobuf:"varint,1,opt,name=GradeType,proto3" json:"GradeType,omitempty"`
}

func (m *C2SEquipGradeUp) Reset()         { *m = C2SEquipGradeUp{} }
func (m *C2SEquipGradeUp) String() string { return proto.CompactTextString(m) }
func (*C2SEquipGradeUp) ProtoMessage()    {}
func (m *C2SEquipGradeUp) GetId() uint16  { return PCK_C2SEquipGradeUp_ID }

const PCK_S2CEquipGradeUp_ID = 206 //
//
type S2CEquipGradeUp struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Grade []*Grade `protobuf:"bytes,2,rep,name=Grade,proto3" json:"Grade,omitempty"`
}

func (m *S2CEquipGradeUp) Reset()         { *m = S2CEquipGradeUp{} }
func (m *S2CEquipGradeUp) String() string { return proto.CompactTextString(m) }
func (*S2CEquipGradeUp) ProtoMessage()    {}
func (m *S2CEquipGradeUp) GetId() uint16  { return PCK_S2CEquipGradeUp_ID }

const PCK_C2SRoleGodGradeUp_ID = 211 //
//
type C2SRoleGodGradeUp struct {
}

func (m *C2SRoleGodGradeUp) Reset()         { *m = C2SRoleGodGradeUp{} }
func (m *C2SRoleGodGradeUp) String() string { return proto.CompactTextString(m) }
func (*C2SRoleGodGradeUp) ProtoMessage()    {}
func (m *C2SRoleGodGradeUp) GetId() uint16  { return PCK_C2SRoleGodGradeUp_ID }

const PCK_S2CRoleGodGradeUp_ID = 212 //
//
type S2CRoleGodGradeUp struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Level int32 `protobuf:"varint,2,opt,name=Level,proto3" json:"Level,omitempty"`
}

func (m *S2CRoleGodGradeUp) Reset()         { *m = S2CRoleGodGradeUp{} }
func (m *S2CRoleGodGradeUp) String() string { return proto.CompactTextString(m) }
func (*S2CRoleGodGradeUp) ProtoMessage()    {}
func (m *S2CRoleGodGradeUp) GetId() uint16  { return PCK_S2CRoleGodGradeUp_ID }

const PCK_C2SGetRoleGodPrize_ID = 213 //
//
type C2SGetRoleGodPrize struct {
	//
	Level int32 `protobuf:"varint,1,opt,name=Level,proto3" json:"Level,omitempty"`
}

func (m *C2SGetRoleGodPrize) Reset()         { *m = C2SGetRoleGodPrize{} }
func (m *C2SGetRoleGodPrize) String() string { return proto.CompactTextString(m) }
func (*C2SGetRoleGodPrize) ProtoMessage()    {}
func (m *C2SGetRoleGodPrize) GetId() uint16  { return PCK_C2SGetRoleGodPrize_ID }

const PCK_S2CGetRoleGodPrize_ID = 214 //
//
type S2CGetRoleGodPrize struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Level int32 `protobuf:"varint,2,opt,name=Level,proto3" json:"Level,omitempty"`
}

func (m *S2CGetRoleGodPrize) Reset()         { *m = S2CGetRoleGodPrize{} }
func (m *S2CGetRoleGodPrize) String() string { return proto.CompactTextString(m) }
func (*S2CGetRoleGodPrize) ProtoMessage()    {}
func (m *S2CGetRoleGodPrize) GetId() uint16  { return PCK_S2CGetRoleGodPrize_ID }

const PCK_S2CRoleGodPrize_ID = 215 //
//
type S2CRoleGodPrize struct {
	//
	Prize []int32 `protobuf:"varint,1,rep,name=Prize,proto3" json:"Prize,omitempty"`
}

func (m *S2CRoleGodPrize) Reset()         { *m = S2CRoleGodPrize{} }
func (m *S2CRoleGodPrize) String() string { return proto.CompactTextString(m) }
func (*S2CRoleGodPrize) ProtoMessage()    {}
func (m *S2CRoleGodPrize) GetId() uint16  { return PCK_S2CRoleGodPrize_ID }

//
type RoleFriend struct {
	//
	FriendList []*DBFriend `protobuf:"bytes,1,rep,name=FriendList,proto3" json:"FriendList,omitempty"`
}

func (m *RoleFriend) Reset()         { *m = RoleFriend{} }
func (m *RoleFriend) String() string { return proto.CompactTextString(m) }
func (*RoleFriend) ProtoMessage()    {}

//
type DBFriend struct {
	//
	FId int32 `protobuf:"varint,1,opt,name=FId,proto3" json:"FId,omitempty"`
	//
	T int32 `protobuf:"varint,2,opt,name=T,proto3" json:"T,omitempty"`
	//
	Cs int32 `protobuf:"varint,3,opt,name=Cs,proto3" json:"Cs,omitempty"`
	//
	E int32 `protobuf:"varint,4,opt,name=E,proto3" json:"E,omitempty"`
}

func (m *DBFriend) Reset()         { *m = DBFriend{} }
func (m *DBFriend) String() string { return proto.CompactTextString(m) }
func (*DBFriend) ProtoMessage()    {}

//
type Friend struct {
	//
	F *DBFriend `protobuf:"bytes,1,opt,name=F,proto3" json:"F,omitempty"`
	//
	A int32 `protobuf:"varint,2,opt,name=A,proto3" json:"A,omitempty"`
	//
	N string `protobuf:"bytes,3,opt,name=N,proto3" json:"N,omitempty"`
	//
	L int32 `protobuf:"varint,4,opt,name=L,proto3" json:"L,omitempty"`
	//
	S int64 `protobuf:"varint,5,opt,name=S,proto3" json:"S,omitempty"`
	//
	V int32 `protobuf:"varint,6,opt,name=V,proto3" json:"V,omitempty"`
	//
	O int64 `protobuf:"varint,7,opt,name=O,proto3" json:"O,omitempty"`
	//
	Gl int64 `protobuf:"varint,8,opt,name=Gl,proto3" json:"Gl,omitempty"`
	//
	Gn string `protobuf:"bytes,9,opt,name=Gn,proto3" json:"Gn,omitempty"`
}

func (m *Friend) Reset()         { *m = Friend{} }
func (m *Friend) String() string { return proto.CompactTextString(m) }
func (*Friend) ProtoMessage()    {}

const PCK_S2CGetFriends_ID = 302 //
//
type S2CGetFriends struct {
	//
	F int32 `protobuf:"varint,1,opt,name=F,proto3" json:"F,omitempty"`
	//
	Tf int32 `protobuf:"varint,2,opt,name=Tf,proto3" json:"Tf,omitempty"`
	//
	A int32 `protobuf:"varint,3,opt,name=A,proto3" json:"A,omitempty"`
	//
	Ta int32 `protobuf:"varint,4,opt,name=Ta,proto3" json:"Ta,omitempty"`
	//
	B int32 `protobuf:"varint,5,opt,name=B,proto3" json:"B,omitempty"`
	//
	Tb int32 `protobuf:"varint,6,opt,name=Tb,proto3" json:"Tb,omitempty"`
	//
	S int32 `protobuf:"varint,7,opt,name=S,proto3" json:"S,omitempty"`
	//
	Ts int32 `protobuf:"varint,8,opt,name=Ts,proto3" json:"Ts,omitempty"`
	//
	R int32 `protobuf:"varint,9,opt,name=R,proto3" json:"R,omitempty"`
	//
	Tr int32 `protobuf:"varint,10,opt,name=Tr,proto3" json:"Tr,omitempty"`
	//
	List []*Friend `protobuf:"bytes,11,rep,name=List,proto3" json:"List,omitempty"`
}

func (m *S2CGetFriends) Reset()         { *m = S2CGetFriends{} }
func (m *S2CGetFriends) String() string { return proto.CompactTextString(m) }
func (*S2CGetFriends) ProtoMessage()    {}
func (m *S2CGetFriends) GetId() uint16  { return PCK_S2CGetFriends_ID }

//
type FriendSendMessage struct {
	//
	List []*Friend `protobuf:"bytes,1,rep,name=List,proto3" json:"List,omitempty"`
	//
	FriendId int32 `protobuf:"varint,2,opt,name=FriendId,proto3" json:"FriendId,omitempty"`
}

func (m *FriendSendMessage) Reset()         { *m = FriendSendMessage{} }
func (m *FriendSendMessage) String() string { return proto.CompactTextString(m) }
func (*FriendSendMessage) ProtoMessage()    {}

const PCK_C2SFocus_ID = 303 //
//
type C2SFocus struct {
	//
	FriendId int32 `protobuf:"varint,1,opt,name=FriendId,proto3" json:"FriendId,omitempty"`
}

func (m *C2SFocus) Reset()         { *m = C2SFocus{} }
func (m *C2SFocus) String() string { return proto.CompactTextString(m) }
func (*C2SFocus) ProtoMessage()    {}
func (m *C2SFocus) GetId() uint16  { return PCK_C2SFocus_ID }

const PCK_S2CFocus_ID = 304 //
//
type S2CFocus struct {
	//
	Tag int32 `protobuf:"varint,2,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	FriendId int32 `protobuf:"varint,1,opt,name=FriendId,proto3" json:"FriendId,omitempty"`
}

func (m *S2CFocus) Reset()         { *m = S2CFocus{} }
func (m *S2CFocus) String() string { return proto.CompactTextString(m) }
func (*S2CFocus) ProtoMessage()    {}
func (m *S2CFocus) GetId() uint16  { return PCK_S2CFocus_ID }

const PCK_C2SCancelFocus_ID = 305 //
//
type C2SCancelFocus struct {
	//
	FriendId int32 `protobuf:"varint,1,opt,name=FriendId,proto3" json:"FriendId,omitempty"`
}

func (m *C2SCancelFocus) Reset()         { *m = C2SCancelFocus{} }
func (m *C2SCancelFocus) String() string { return proto.CompactTextString(m) }
func (*C2SCancelFocus) ProtoMessage()    {}
func (m *C2SCancelFocus) GetId() uint16  { return PCK_C2SCancelFocus_ID }

const PCK_S2CCancelFocus_ID = 306 //
//
type S2CCancelFocus struct {
	//
	Tag int32 `protobuf:"varint,2,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	FriendId int32 `protobuf:"varint,1,opt,name=FriendId,proto3" json:"FriendId,omitempty"`
}

func (m *S2CCancelFocus) Reset()         { *m = S2CCancelFocus{} }
func (m *S2CCancelFocus) String() string { return proto.CompactTextString(m) }
func (*S2CCancelFocus) ProtoMessage()    {}
func (m *S2CCancelFocus) GetId() uint16  { return PCK_S2CCancelFocus_ID }

const PCK_C2SHate_ID = 307 //
//
type C2SHate struct {
	//
	FriendId int32 `protobuf:"varint,1,opt,name=FriendId,proto3" json:"FriendId,omitempty"`
}

func (m *C2SHate) Reset()         { *m = C2SHate{} }
func (m *C2SHate) String() string { return proto.CompactTextString(m) }
func (*C2SHate) ProtoMessage()    {}
func (m *C2SHate) GetId() uint16  { return PCK_C2SHate_ID }

const PCK_S2CHate_ID = 308 //
//
type S2CHate struct {
	//
	Tag int32 `protobuf:"varint,2,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	FriendId int32 `protobuf:"varint,1,opt,name=FriendId,proto3" json:"FriendId,omitempty"`
}

func (m *S2CHate) Reset()         { *m = S2CHate{} }
func (m *S2CHate) String() string { return proto.CompactTextString(m) }
func (*S2CHate) ProtoMessage()    {}
func (m *S2CHate) GetId() uint16  { return PCK_S2CHate_ID }

const PCK_C2SCancelHate_ID = 309 //
//
type C2SCancelHate struct {
	//
	FriendId int32 `protobuf:"varint,1,opt,name=FriendId,proto3" json:"FriendId,omitempty"`
}

func (m *C2SCancelHate) Reset()         { *m = C2SCancelHate{} }
func (m *C2SCancelHate) String() string { return proto.CompactTextString(m) }
func (*C2SCancelHate) ProtoMessage()    {}
func (m *C2SCancelHate) GetId() uint16  { return PCK_C2SCancelHate_ID }

const PCK_S2CCancelHate_ID = 310 //
//
type S2CCancelHate struct {
	//
	Tag int32 `protobuf:"varint,2,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	FriendId int32 `protobuf:"varint,1,opt,name=FriendId,proto3" json:"FriendId,omitempty"`
}

func (m *S2CCancelHate) Reset()         { *m = S2CCancelHate{} }
func (m *S2CCancelHate) String() string { return proto.CompactTextString(m) }
func (*S2CCancelHate) ProtoMessage()    {}
func (m *S2CCancelHate) GetId() uint16  { return PCK_S2CCancelHate_ID }

const PCK_C2SGiveCoin_ID = 311 //
//
type C2SGiveCoin struct {
	//
	FriendId int32 `protobuf:"varint,1,opt,name=FriendId,proto3" json:"FriendId,omitempty"`
}

func (m *C2SGiveCoin) Reset()         { *m = C2SGiveCoin{} }
func (m *C2SGiveCoin) String() string { return proto.CompactTextString(m) }
func (*C2SGiveCoin) ProtoMessage()    {}
func (m *C2SGiveCoin) GetId() uint16  { return PCK_C2SGiveCoin_ID }

const PCK_S2CGiveCoin_ID = 312 //
//
type S2CGiveCoin struct {
	//
	Tag int32 `protobuf:"varint,2,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	FriendId int32 `protobuf:"varint,1,opt,name=FriendId,proto3" json:"FriendId,omitempty"`
}

func (m *S2CGiveCoin) Reset()         { *m = S2CGiveCoin{} }
func (m *S2CGiveCoin) String() string { return proto.CompactTextString(m) }
func (*S2CGiveCoin) ProtoMessage()    {}
func (m *S2CGiveCoin) GetId() uint16  { return PCK_S2CGiveCoin_ID }

const PCK_C2SGetCoin_ID = 313 //
//
type C2SGetCoin struct {
	//
	FriendId int32 `protobuf:"varint,1,opt,name=FriendId,proto3" json:"FriendId,omitempty"`
}

func (m *C2SGetCoin) Reset()         { *m = C2SGetCoin{} }
func (m *C2SGetCoin) String() string { return proto.CompactTextString(m) }
func (*C2SGetCoin) ProtoMessage()    {}
func (m *C2SGetCoin) GetId() uint16  { return PCK_C2SGetCoin_ID }

const PCK_S2CGetCoin_ID = 314 //
//
type S2CGetCoin struct {
	//
	Tag int32 `protobuf:"varint,2,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	FriendId int32 `protobuf:"varint,1,opt,name=FriendId,proto3" json:"FriendId,omitempty"`
}

func (m *S2CGetCoin) Reset()         { *m = S2CGetCoin{} }
func (m *S2CGetCoin) String() string { return proto.CompactTextString(m) }
func (*S2CGetCoin) ProtoMessage()    {}
func (m *S2CGetCoin) GetId() uint16  { return PCK_S2CGetCoin_ID }

const PCK_C2SGetSuggest_ID = 315 //
//
type C2SGetSuggest struct {
}

func (m *C2SGetSuggest) Reset()         { *m = C2SGetSuggest{} }
func (m *C2SGetSuggest) String() string { return proto.CompactTextString(m) }
func (*C2SGetSuggest) ProtoMessage()    {}
func (m *C2SGetSuggest) GetId() uint16  { return PCK_C2SGetSuggest_ID }

const PCK_S2CGetSuggest_ID = 316 //
//
type S2CGetSuggest struct {
	//
	List []*Friend `protobuf:"bytes,1,rep,name=List,proto3" json:"List,omitempty"`
}

func (m *S2CGetSuggest) Reset()         { *m = S2CGetSuggest{} }
func (m *S2CGetSuggest) String() string { return proto.CompactTextString(m) }
func (*S2CGetSuggest) ProtoMessage()    {}
func (m *S2CGetSuggest) GetId() uint16  { return PCK_S2CGetSuggest_ID }

const PCK_C2SOneKeyGiveCoin_ID = 317 //
//
type C2SOneKeyGiveCoin struct {
}

func (m *C2SOneKeyGiveCoin) Reset()         { *m = C2SOneKeyGiveCoin{} }
func (m *C2SOneKeyGiveCoin) String() string { return proto.CompactTextString(m) }
func (*C2SOneKeyGiveCoin) ProtoMessage()    {}
func (m *C2SOneKeyGiveCoin) GetId() uint16  { return PCK_C2SOneKeyGiveCoin_ID }

const PCK_S2COneKeyGiveCoin_ID = 318 //
//
type S2COneKeyGiveCoin struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Num int32 `protobuf:"varint,2,opt,name=Num,proto3" json:"Num,omitempty"`
}

func (m *S2COneKeyGiveCoin) Reset()         { *m = S2COneKeyGiveCoin{} }
func (m *S2COneKeyGiveCoin) String() string { return proto.CompactTextString(m) }
func (*S2COneKeyGiveCoin) ProtoMessage()    {}
func (m *S2COneKeyGiveCoin) GetId() uint16  { return PCK_S2COneKeyGiveCoin_ID }

const PCK_C2SOneKeyGetCoin_ID = 319 //
//
type C2SOneKeyGetCoin struct {
}

func (m *C2SOneKeyGetCoin) Reset()         { *m = C2SOneKeyGetCoin{} }
func (m *C2SOneKeyGetCoin) String() string { return proto.CompactTextString(m) }
func (*C2SOneKeyGetCoin) ProtoMessage()    {}
func (m *C2SOneKeyGetCoin) GetId() uint16  { return PCK_C2SOneKeyGetCoin_ID }

const PCK_S2COneKeyGetCoin_ID = 320 //
//
type S2COneKeyGetCoin struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Num int32 `protobuf:"varint,2,opt,name=Num,proto3" json:"Num,omitempty"`
}

func (m *S2COneKeyGetCoin) Reset()         { *m = S2COneKeyGetCoin{} }
func (m *S2COneKeyGetCoin) String() string { return proto.CompactTextString(m) }
func (*S2COneKeyGetCoin) ProtoMessage()    {}
func (m *S2COneKeyGetCoin) GetId() uint16  { return PCK_S2COneKeyGetCoin_ID }

const PCK_C2SOneKeyFocus_ID = 321 //
//
type C2SOneKeyFocus struct {
}

func (m *C2SOneKeyFocus) Reset()         { *m = C2SOneKeyFocus{} }
func (m *C2SOneKeyFocus) String() string { return proto.CompactTextString(m) }
func (*C2SOneKeyFocus) ProtoMessage()    {}
func (m *C2SOneKeyFocus) GetId() uint16  { return PCK_C2SOneKeyFocus_ID }

const PCK_S2COneKeyFocus_ID = 322 //
//
type S2COneKeyFocus struct {
	//
	Tags []int32 `protobuf:"varint,1,rep,name=Tags,proto3" json:"Tags,omitempty"`
	//
	Num int32 `protobuf:"varint,2,opt,name=Num,proto3" json:"Num,omitempty"`
}

func (m *S2COneKeyFocus) Reset()         { *m = S2COneKeyFocus{} }
func (m *S2COneKeyFocus) String() string { return proto.CompactTextString(m) }
func (*S2COneKeyFocus) ProtoMessage()    {}
func (m *S2COneKeyFocus) GetId() uint16  { return PCK_S2COneKeyFocus_ID }

//
type ChatMsg struct {
	//
	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	//
	ChannelId int32 `protobuf:"varint,2,opt,name=ChannelId,proto3" json:"ChannelId,omitempty"`
	//
	Type int32 `protobuf:"varint,3,opt,name=Type,proto3" json:"Type,omitempty"`
	//
	Content string `protobuf:"bytes,4,opt,name=Content,proto3" json:"Content,omitempty"`
	//
	SenderId int32 `protobuf:"varint,5,opt,name=SenderId,proto3" json:"SenderId,omitempty"`
	//
	SenderNick string `protobuf:"bytes,6,opt,name=SenderNick,proto3" json:"SenderNick,omitempty"`
	//
	SenderRoleId int32 `protobuf:"varint,7,opt,name=SenderRoleId,proto3" json:"SenderRoleId,omitempty"`
	//
	SenderGodLevel int64 `protobuf:"varint,8,opt,name=SenderGodLevel,proto3" json:"SenderGodLevel,omitempty"`
	//
	AreaId int64 `protobuf:"varint,9,opt,name=AreaId,proto3" json:"AreaId,omitempty"`
}

func (m *ChatMsg) Reset()         { *m = ChatMsg{} }
func (m *ChatMsg) String() string { return proto.CompactTextString(m) }
func (*ChatMsg) ProtoMessage()    {}

const PCK_C2SSendChatMsg_ID = 404 //
//
type C2SSendChatMsg struct {
	//
	ChannelId int32 `protobuf:"varint,1,opt,name=ChannelId,proto3" json:"ChannelId,omitempty"`
	//
	Content string `protobuf:"bytes,2,opt,name=Content,proto3" json:"Content,omitempty"`
}

func (m *C2SSendChatMsg) Reset()         { *m = C2SSendChatMsg{} }
func (m *C2SSendChatMsg) String() string { return proto.CompactTextString(m) }
func (*C2SSendChatMsg) ProtoMessage()    {}
func (m *C2SSendChatMsg) GetId() uint16  { return PCK_C2SSendChatMsg_ID }

const PCK_S2CSendChatMsg_ID = 405 //
//
type S2CSendChatMsg struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CSendChatMsg) Reset()         { *m = S2CSendChatMsg{} }
func (m *S2CSendChatMsg) String() string { return proto.CompactTextString(m) }
func (*S2CSendChatMsg) ProtoMessage()    {}
func (m *S2CSendChatMsg) GetId() uint16  { return PCK_S2CSendChatMsg_ID }

const PCK_C2SGetHistoryChat_ID = 401 //
//
type C2SGetHistoryChat struct {
}

func (m *C2SGetHistoryChat) Reset()         { *m = C2SGetHistoryChat{} }
func (m *C2SGetHistoryChat) String() string { return proto.CompactTextString(m) }
func (*C2SGetHistoryChat) ProtoMessage()    {}
func (m *C2SGetHistoryChat) GetId() uint16  { return PCK_C2SGetHistoryChat_ID }

const PCK_S2CGetHistoryChat_ID = 402 //
//
type S2CGetHistoryChat struct {
	//
	Chatmessage []*ChatMsg `protobuf:"bytes,1,rep,name=Chatmessage,proto3" json:"Chatmessage,omitempty"`
}

func (m *S2CGetHistoryChat) Reset()         { *m = S2CGetHistoryChat{} }
func (m *S2CGetHistoryChat) String() string { return proto.CompactTextString(m) }
func (*S2CGetHistoryChat) ProtoMessage()    {}
func (m *S2CGetHistoryChat) GetId() uint16  { return PCK_S2CGetHistoryChat_ID }

const PCK_S2CNewChatMsg_ID = 403 //
//
type S2CNewChatMsg struct {
	//
	Chatmessage *ChatMsg `protobuf:"bytes,1,opt,name=Chatmessage,proto3" json:"Chatmessage,omitempty"`
}

func (m *S2CNewChatMsg) Reset()         { *m = S2CNewChatMsg{} }
func (m *S2CNewChatMsg) String() string { return proto.CompactTextString(m) }
func (*S2CNewChatMsg) ProtoMessage()    {}
func (m *S2CNewChatMsg) GetId() uint16  { return PCK_S2CNewChatMsg_ID }

//
type WhisperMsg struct {
	//
	SenderId int32 `protobuf:"varint,1,opt,name=SenderId,proto3" json:"SenderId,omitempty"`
	//
	SenderNick string `protobuf:"bytes,2,opt,name=SenderNick,proto3" json:"SenderNick,omitempty"`
	//
	SenderRoleId int32 `protobuf:"varint,3,opt,name=SenderRoleId,proto3" json:"SenderRoleId,omitempty"`
	//
	SenderGodLevel int64 `protobuf:"varint,4,opt,name=SenderGodLevel,proto3" json:"SenderGodLevel,omitempty"`
	//
	ReceiverId int32 `protobuf:"varint,5,opt,name=ReceiverId,proto3" json:"ReceiverId,omitempty"`
	//
	ReceiverNick string `protobuf:"bytes,6,opt,name=ReceiverNick,proto3" json:"ReceiverNick,omitempty"`
	//
	ReceiverRoleId int32 `protobuf:"varint,7,opt,name=ReceiverRoleId,proto3" json:"ReceiverRoleId,omitempty"`
	//
	ReceiverGodLevel int64 `protobuf:"varint,8,opt,name=ReceiverGodLevel,proto3" json:"ReceiverGodLevel,omitempty"`
	//
	Content string `protobuf:"bytes,9,opt,name=Content,proto3" json:"Content,omitempty"`
	//
	Time int64 `protobuf:"varint,10,opt,name=Time,proto3" json:"Time,omitempty"`
}

func (m *WhisperMsg) Reset()         { *m = WhisperMsg{} }
func (m *WhisperMsg) String() string { return proto.CompactTextString(m) }
func (*WhisperMsg) ProtoMessage()    {}

const PCK_C2SWhisper_ID = 411 //
//
type C2SWhisper struct {
	//
	ReceiverId int32 `protobuf:"varint,1,opt,name=ReceiverId,proto3" json:"ReceiverId,omitempty"`
	//
	Content string `protobuf:"bytes,2,opt,name=Content,proto3" json:"Content,omitempty"`
}

func (m *C2SWhisper) Reset()         { *m = C2SWhisper{} }
func (m *C2SWhisper) String() string { return proto.CompactTextString(m) }
func (*C2SWhisper) ProtoMessage()    {}
func (m *C2SWhisper) GetId() uint16  { return PCK_C2SWhisper_ID }

const PCK_S2CWhisper_ID = 412 //
//
type S2CWhisper struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	M *WhisperMsg `protobuf:"bytes,2,opt,name=M,proto3" json:"M,omitempty"`
}

func (m *S2CWhisper) Reset()         { *m = S2CWhisper{} }
func (m *S2CWhisper) String() string { return proto.CompactTextString(m) }
func (*S2CWhisper) ProtoMessage()    {}
func (m *S2CWhisper) GetId() uint16  { return PCK_S2CWhisper_ID }

const PCK_C2SGetWhisper_ID = 413 //
//
type C2SGetWhisper struct {
	//
	PlayerId int32 `protobuf:"varint,1,opt,name=PlayerId,proto3" json:"PlayerId,omitempty"`
}

func (m *C2SGetWhisper) Reset()         { *m = C2SGetWhisper{} }
func (m *C2SGetWhisper) String() string { return proto.CompactTextString(m) }
func (*C2SGetWhisper) ProtoMessage()    {}
func (m *C2SGetWhisper) GetId() uint16  { return PCK_C2SGetWhisper_ID }

const PCK_S2CGetWhisper_ID = 414 //
//
type S2CGetWhisper struct {
	//
	m []*WhisperMsg `protobuf:"bytes,1,rep,name=m,proto3" json:"m,omitempty"`
	//
	TargetId int32 `protobuf:"varint,2,opt,name=TargetId,proto3" json:"TargetId,omitempty"`
}

func (m *S2CGetWhisper) Reset()         { *m = S2CGetWhisper{} }
func (m *S2CGetWhisper) String() string { return proto.CompactTextString(m) }
func (*S2CGetWhisper) ProtoMessage()    {}
func (m *S2CGetWhisper) GetId() uint16  { return PCK_S2CGetWhisper_ID }

const PCK_S2CAllUnreadWhisper_ID = 415 //
//
type S2CAllUnreadWhisper struct {
	//
	m []*WhisperMsg `protobuf:"bytes,1,rep,name=m,proto3" json:"m,omitempty"`
}

func (m *S2CAllUnreadWhisper) Reset()         { *m = S2CAllUnreadWhisper{} }
func (m *S2CAllUnreadWhisper) String() string { return proto.CompactTextString(m) }
func (*S2CAllUnreadWhisper) ProtoMessage()    {}
func (m *S2CAllUnreadWhisper) GetId() uint16  { return PCK_S2CAllUnreadWhisper_ID }

const PCK_C2SGetOnlineStatus_ID = 416 //
//
type C2SGetOnlineStatus struct {
	//
	UidList []int32 `protobuf:"varint,1,rep,name=UidList,proto3" json:"UidList,omitempty"`
}

func (m *C2SGetOnlineStatus) Reset()         { *m = C2SGetOnlineStatus{} }
func (m *C2SGetOnlineStatus) String() string { return proto.CompactTextString(m) }
func (*C2SGetOnlineStatus) ProtoMessage()    {}
func (m *C2SGetOnlineStatus) GetId() uint16  { return PCK_C2SGetOnlineStatus_ID }

//
type PlayerOnlineStatus struct {
	//
	Uid int32 `protobuf:"varint,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
	//
	Status int32 `protobuf:"varint,2,opt,name=Status,proto3" json:"Status,omitempty"`
}

func (m *PlayerOnlineStatus) Reset()         { *m = PlayerOnlineStatus{} }
func (m *PlayerOnlineStatus) String() string { return proto.CompactTextString(m) }
func (*PlayerOnlineStatus) ProtoMessage()    {}

const PCK_S2CGetOnlineStatus_ID = 417 //
//
type S2CGetOnlineStatus struct {
	//
	List []*PlayerOnlineStatus `protobuf:"bytes,1,rep,name=List,proto3" json:"List,omitempty"`
}

func (m *S2CGetOnlineStatus) Reset()         { *m = S2CGetOnlineStatus{} }
func (m *S2CGetOnlineStatus) String() string { return proto.CompactTextString(m) }
func (*S2CGetOnlineStatus) ProtoMessage()    {}
func (m *S2CGetOnlineStatus) GetId() uint16  { return PCK_S2CGetOnlineStatus_ID }

const PCK_C2SRemoveWhisper_ID = 418 //
//
type C2SRemoveWhisper struct {
	//
	PlayerId int32 `protobuf:"varint,1,opt,name=PlayerId,proto3" json:"PlayerId,omitempty"`
}

func (m *C2SRemoveWhisper) Reset()         { *m = C2SRemoveWhisper{} }
func (m *C2SRemoveWhisper) String() string { return proto.CompactTextString(m) }
func (*C2SRemoveWhisper) ProtoMessage()    {}
func (m *C2SRemoveWhisper) GetId() uint16  { return PCK_C2SRemoveWhisper_ID }

const PCK_S2CRemoveWhisper_ID = 419 //
//
type S2CRemoveWhisper struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CRemoveWhisper) Reset()         { *m = S2CRemoveWhisper{} }
func (m *S2CRemoveWhisper) String() string { return proto.CompactTextString(m) }
func (*S2CRemoveWhisper) ProtoMessage()    {}
func (m *S2CRemoveWhisper) GetId() uint16  { return PCK_S2CRemoveWhisper_ID }

//
type RoleBag struct {
	//
	Items []*ItemData `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty"`
	//
	Size int32 `protobuf:"varint,2,opt,name=Size,proto3" json:"Size,omitempty"`
	//
	Auto int32 `protobuf:"varint,3,opt,name=Auto,proto3" json:"Auto,omitempty"`
	//
	ExtendBagSize int32 `protobuf:"varint,4,opt,name=ExtendBagSize,proto3" json:"ExtendBagSize,omitempty"`
}

func (m *RoleBag) Reset()         { *m = RoleBag{} }
func (m *RoleBag) String() string { return proto.CompactTextString(m) }
func (*RoleBag) ProtoMessage()    {}

//
type BagChange struct {
	//
	T int32 `protobuf:"varint,1,opt,name=T,proto3" json:"T,omitempty"`
	//
	Item *ItemData `protobuf:"bytes,2,opt,name=Item,proto3" json:"Item,omitempty"`
}

func (m *BagChange) Reset()         { *m = BagChange{} }
func (m *BagChange) String() string { return proto.CompactTextString(m) }
func (*BagChange) ProtoMessage()    {}

const PCK_S2CBagChange_ID = 520 //
//
type S2CBagChange struct {
	//
	Change []*BagChange `protobuf:"bytes,1,rep,name=Change,proto3" json:"Change,omitempty"`
}

func (m *S2CBagChange) Reset()         { *m = S2CBagChange{} }
func (m *S2CBagChange) String() string { return proto.CompactTextString(m) }
func (*S2CBagChange) ProtoMessage()    {}
func (m *S2CBagChange) GetId() uint16  { return PCK_S2CBagChange_ID }

//
type S2CNewItem struct {
	//
	Change []*BagChange `protobuf:"bytes,1,rep,name=Change,proto3" json:"Change,omitempty"`
}

func (m *S2CNewItem) Reset()         { *m = S2CNewItem{} }
func (m *S2CNewItem) String() string { return proto.CompactTextString(m) }
func (*S2CNewItem) ProtoMessage()    {}

const PCK_C2SUserBag_ID = 500 //
//
type C2SUserBag struct {
}

func (m *C2SUserBag) Reset()         { *m = C2SUserBag{} }
func (m *C2SUserBag) String() string { return proto.CompactTextString(m) }
func (*C2SUserBag) ProtoMessage()    {}
func (m *C2SUserBag) GetId() uint16  { return PCK_C2SUserBag_ID }

const PCK_S2CUserBag_ID = 501 //
//
type S2CUserBag struct {
	//
	Bag *RoleBag `protobuf:"bytes,1,opt,name=Bag,proto3" json:"Bag,omitempty"`
}

func (m *S2CUserBag) Reset()         { *m = S2CUserBag{} }
func (m *S2CUserBag) String() string { return proto.CompactTextString(m) }
func (*S2CUserBag) ProtoMessage()    {}
func (m *S2CUserBag) GetId() uint16  { return PCK_S2CUserBag_ID }

const PCK_C2SExtendBag_ID = 504 //
//
type C2SExtendBag struct {
	//
	Count int32 `protobuf:"varint,1,opt,name=Count,proto3" json:"Count,omitempty"`
}

func (m *C2SExtendBag) Reset()         { *m = C2SExtendBag{} }
func (m *C2SExtendBag) String() string { return proto.CompactTextString(m) }
func (*C2SExtendBag) ProtoMessage()    {}
func (m *C2SExtendBag) GetId() uint16  { return PCK_C2SExtendBag_ID }

const PCK_S2CExtendBag_ID = 505 //
//
type S2CExtendBag struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Size int32 `protobuf:"varint,2,opt,name=Size,proto3" json:"Size,omitempty"`
}

func (m *S2CExtendBag) Reset()         { *m = S2CExtendBag{} }
func (m *S2CExtendBag) String() string { return proto.CompactTextString(m) }
func (*S2CExtendBag) ProtoMessage()    {}
func (m *S2CExtendBag) GetId() uint16  { return PCK_S2CExtendBag_ID }

const PCK_C2SExchange_ID = 502 //
//
type C2SExchange struct {
	//
	ItemID string `protobuf:"bytes,2,opt,name=ItemID,proto3" json:"ItemID,omitempty"`
	//
	Count int64 `protobuf:"varint,1,opt,name=Count,proto3" json:"Count,omitempty"`
	//
	Param1 int32 `protobuf:"varint,3,opt,name=Param1,proto3" json:"Param1,omitempty"`
}

func (m *C2SExchange) Reset()         { *m = C2SExchange{} }
func (m *C2SExchange) String() string { return proto.CompactTextString(m) }
func (*C2SExchange) ProtoMessage()    {}
func (m *C2SExchange) GetId() uint16  { return PCK_C2SExchange_ID }

const PCK_S2CExchange_ID = 503 //
//
type S2CExchange struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CExchange) Reset()         { *m = S2CExchange{} }
func (m *S2CExchange) String() string { return proto.CompactTextString(m) }
func (*S2CExchange) ProtoMessage()    {}
func (m *S2CExchange) GetId() uint16  { return PCK_S2CExchange_ID }

const PCK_C2SGoldBag_ID = 530 //
//
type C2SGoldBag struct {
	//
	Buy int32 `protobuf:"varint,1,opt,name=Buy,proto3" json:"Buy,omitempty"`
}

func (m *C2SGoldBag) Reset()         { *m = C2SGoldBag{} }
func (m *C2SGoldBag) String() string { return proto.CompactTextString(m) }
func (*C2SGoldBag) ProtoMessage()    {}
func (m *C2SGoldBag) GetId() uint16  { return PCK_C2SGoldBag_ID }

const PCK_S2CGoldBag_ID = 531 //
//
type S2CGoldBag struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CGoldBag) Reset()         { *m = S2CGoldBag{} }
func (m *S2CGoldBag) String() string { return proto.CompactTextString(m) }
func (*S2CGoldBag) ProtoMessage()    {}
func (m *S2CGoldBag) GetId() uint16  { return PCK_S2CGoldBag_ID }

const PCK_C2SGetGoldBagInfo_ID = 532 //
//
type C2SGetGoldBagInfo struct {
}

func (m *C2SGetGoldBagInfo) Reset()         { *m = C2SGetGoldBagInfo{} }
func (m *C2SGetGoldBagInfo) String() string { return proto.CompactTextString(m) }
func (*C2SGetGoldBagInfo) ProtoMessage()    {}
func (m *C2SGetGoldBagInfo) GetId() uint16  { return PCK_C2SGetGoldBagInfo_ID }

const PCK_S2CGetGoldBagInfo_ID = 533 //
//
type S2CGetGoldBagInfo struct {
	//
	TodayRamin int32 `protobuf:"varint,1,opt,name=TodayRamin,proto3" json:"TodayRamin,omitempty"`
	//
	NextVipTimes int32 `protobuf:"varint,2,opt,name=NextVipTimes,proto3" json:"NextVipTimes,omitempty"`
}

func (m *S2CGetGoldBagInfo) Reset()         { *m = S2CGetGoldBagInfo{} }
func (m *S2CGetGoldBagInfo) String() string { return proto.CompactTextString(m) }
func (*S2CGetGoldBagInfo) ProtoMessage()    {}
func (m *S2CGetGoldBagInfo) GetId() uint16  { return PCK_S2CGetGoldBagInfo_ID }

const PCK_C2SWearEquip_ID = 506 //
//
type C2SWearEquip struct {
	//
	Type int32 `protobuf:"varint,1,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (m *C2SWearEquip) Reset()         { *m = C2SWearEquip{} }
func (m *C2SWearEquip) String() string { return proto.CompactTextString(m) }
func (*C2SWearEquip) ProtoMessage()    {}
func (m *C2SWearEquip) GetId() uint16  { return PCK_C2SWearEquip_ID }

const PCK_S2CWearEquip_ID = 507 //
//
type S2CWearEquip struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CWearEquip) Reset()         { *m = S2CWearEquip{} }
func (m *S2CWearEquip) String() string { return proto.CompactTextString(m) }
func (*S2CWearEquip) ProtoMessage()    {}
func (m *S2CWearEquip) GetId() uint16  { return PCK_S2CWearEquip_ID }

const PCK_C2SWearOneEquip_ID = 525 //
//
type C2SWearOneEquip struct {
	//
	ItemId string `protobuf:"bytes,1,opt,name=ItemId,proto3" json:"ItemId,omitempty"`
}

func (m *C2SWearOneEquip) Reset()         { *m = C2SWearOneEquip{} }
func (m *C2SWearOneEquip) String() string { return proto.CompactTextString(m) }
func (*C2SWearOneEquip) ProtoMessage()    {}
func (m *C2SWearOneEquip) GetId() uint16  { return PCK_C2SWearOneEquip_ID }

const PCK_S2CWearOneEquip_ID = 526 //
//
type S2CWearOneEquip struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CWearOneEquip) Reset()         { *m = S2CWearOneEquip{} }
func (m *S2CWearOneEquip) String() string { return proto.CompactTextString(m) }
func (*S2CWearOneEquip) ProtoMessage()    {}
func (m *S2CWearOneEquip) GetId() uint16  { return PCK_S2CWearOneEquip_ID }

const PCK_C2SEquipStar_ID = 527 //
//
type C2SEquipStar struct {
	//
	ItemId string `protobuf:"bytes,1,opt,name=ItemId,proto3" json:"ItemId,omitempty"`
	//
	Auto int32 `protobuf:"varint,2,opt,name=Auto,proto3" json:"Auto,omitempty"`
}

func (m *C2SEquipStar) Reset()         { *m = C2SEquipStar{} }
func (m *C2SEquipStar) String() string { return proto.CompactTextString(m) }
func (*C2SEquipStar) ProtoMessage()    {}
func (m *C2SEquipStar) GetId() uint16  { return PCK_C2SEquipStar_ID }

const PCK_S2CEquipStar_ID = 528 //
//
type S2CEquipStar struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Ret int32 `protobuf:"varint,2,opt,name=Ret,proto3" json:"Ret,omitempty"`
}

func (m *S2CEquipStar) Reset()         { *m = S2CEquipStar{} }
func (m *S2CEquipStar) String() string { return proto.CompactTextString(m) }
func (*S2CEquipStar) ProtoMessage()    {}
func (m *S2CEquipStar) GetId() uint16  { return PCK_S2CEquipStar_ID }

const PCK_C2SMeltEquip_ID = 508 //
//
type C2SMeltEquip struct {
	//
	Items []string `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty"`
}

func (m *C2SMeltEquip) Reset()         { *m = C2SMeltEquip{} }
func (m *C2SMeltEquip) String() string { return proto.CompactTextString(m) }
func (*C2SMeltEquip) ProtoMessage()    {}
func (m *C2SMeltEquip) GetId() uint16  { return PCK_C2SMeltEquip_ID }

const PCK_S2CMeltEquip_ID = 509 //
//
type S2CMeltEquip struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CMeltEquip) Reset()         { *m = S2CMeltEquip{} }
func (m *S2CMeltEquip) String() string { return proto.CompactTextString(m) }
func (*S2CMeltEquip) ProtoMessage()    {}
func (m *S2CMeltEquip) GetId() uint16  { return PCK_S2CMeltEquip_ID }

const PCK_C2SOpenTreasure_ID = 510 //
//
type C2SOpenTreasure struct {
	//
	ItemID string `protobuf:"bytes,2,opt,name=ItemID,proto3" json:"ItemID,omitempty"`
	//
	Count int64 `protobuf:"varint,1,opt,name=Count,proto3" json:"Count,omitempty"`
}

func (m *C2SOpenTreasure) Reset()         { *m = C2SOpenTreasure{} }
func (m *C2SOpenTreasure) String() string { return proto.CompactTextString(m) }
func (*C2SOpenTreasure) ProtoMessage()    {}
func (m *C2SOpenTreasure) GetId() uint16  { return PCK_C2SOpenTreasure_ID }

const PCK_S2COpenTreasure_ID = 511 //
//
type S2COpenTreasure struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2COpenTreasure) Reset()         { *m = S2COpenTreasure{} }
func (m *S2COpenTreasure) String() string { return proto.CompactTextString(m) }
func (*S2COpenTreasure) ProtoMessage()    {}
func (m *S2COpenTreasure) GetId() uint16  { return PCK_S2COpenTreasure_ID }

const PCK_C2SAutoMelt_ID = 512 //
//
type C2SAutoMelt struct {
}

func (m *C2SAutoMelt) Reset()         { *m = C2SAutoMelt{} }
func (m *C2SAutoMelt) String() string { return proto.CompactTextString(m) }
func (*C2SAutoMelt) ProtoMessage()    {}
func (m *C2SAutoMelt) GetId() uint16  { return PCK_C2SAutoMelt_ID }

const PCK_S2CAutoMelt_ID = 513 //
//
type S2CAutoMelt struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Auto int32 `protobuf:"varint,2,opt,name=Auto,proto3" json:"Auto,omitempty"`
	//
	Melt int32 `protobuf:"varint,3,opt,name=Melt,proto3" json:"Melt,omitempty"`
}

func (m *S2CAutoMelt) Reset()         { *m = S2CAutoMelt{} }
func (m *S2CAutoMelt) String() string { return proto.CompactTextString(m) }
func (*S2CAutoMelt) ProtoMessage()    {}
func (m *S2CAutoMelt) GetId() uint16  { return PCK_S2CAutoMelt_ID }

const PCK_C2SMeltGoldByItem_ID = 521 //
//
type C2SMeltGoldByItem struct {
	//
	PartId int32 `protobuf:"varint,1,opt,name=PartId,proto3" json:"PartId,omitempty"`
}

func (m *C2SMeltGoldByItem) Reset()         { *m = C2SMeltGoldByItem{} }
func (m *C2SMeltGoldByItem) String() string { return proto.CompactTextString(m) }
func (*C2SMeltGoldByItem) ProtoMessage()    {}
func (m *C2SMeltGoldByItem) GetId() uint16  { return PCK_C2SMeltGoldByItem_ID }

const PCK_S2CMeltGoldByItem_ID = 522 //
//
type S2CMeltGoldByItem struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CMeltGoldByItem) Reset()         { *m = S2CMeltGoldByItem{} }
func (m *S2CMeltGoldByItem) String() string { return proto.CompactTextString(m) }
func (*S2CMeltGoldByItem) ProtoMessage()    {}
func (m *S2CMeltGoldByItem) GetId() uint16  { return PCK_S2CMeltGoldByItem_ID }

const PCK_C2SGetBattlePrize_ID = 523 //
//
type C2SGetBattlePrize struct {
}

func (m *C2SGetBattlePrize) Reset()         { *m = C2SGetBattlePrize{} }
func (m *C2SGetBattlePrize) String() string { return proto.CompactTextString(m) }
func (*C2SGetBattlePrize) ProtoMessage()    {}
func (m *C2SGetBattlePrize) GetId() uint16  { return PCK_C2SGetBattlePrize_ID }

const PCK_ItemFly_ID = 524 //
//
type ItemFly struct {
	//
	Item []*ItemNum `protobuf:"bytes,1,rep,name=Item,proto3" json:"Item,omitempty"`
}

func (m *ItemFly) Reset()         { *m = ItemFly{} }
func (m *ItemFly) String() string { return proto.CompactTextString(m) }
func (*ItemFly) ProtoMessage()    {}
func (m *ItemFly) GetId() uint16  { return PCK_ItemFly_ID }

const PCK_C2SGetShopMallList_ID = 430 //
//
type C2SGetShopMallList struct {
	//
	MallType int32 `protobuf:"varint,1,opt,name=MallType,proto3" json:"MallType,omitempty"`
}

func (m *C2SGetShopMallList) Reset()         { *m = C2SGetShopMallList{} }
func (m *C2SGetShopMallList) String() string { return proto.CompactTextString(m) }
func (*C2SGetShopMallList) ProtoMessage()    {}
func (m *C2SGetShopMallList) GetId() uint16  { return PCK_C2SGetShopMallList_ID }

//
type ShopGoodsInfo struct {
	//
	GoodsId int32 `protobuf:"varint,1,opt,name=GoodsId,proto3" json:"GoodsId,omitempty"`
	//
	LimitBuy int32 `protobuf:"varint,2,opt,name=LimitBuy,proto3" json:"LimitBuy,omitempty"`
	//
	LimitUseTimes int32 `protobuf:"varint,3,opt,name=LimitUseTimes,proto3" json:"LimitUseTimes,omitempty"`
	//
	LimitTotalTimes int32 `protobuf:"varint,4,opt,name=LimitTotalTimes,proto3" json:"LimitTotalTimes,omitempty"`
	//
	MallType int32 `protobuf:"varint,5,opt,name=MallType,proto3" json:"MallType,omitempty"`
	//
	Sort int32 `protobuf:"varint,6,opt,name=Sort,proto3" json:"Sort,omitempty"`
	//
	AddFight int64 `protobuf:"varint,7,opt,name=AddFight,proto3" json:"AddFight,omitempty"`
}

func (m *ShopGoodsInfo) Reset()         { *m = ShopGoodsInfo{} }
func (m *ShopGoodsInfo) String() string { return proto.CompactTextString(m) }
func (*ShopGoodsInfo) ProtoMessage()    {}

const PCK_S2CGetShopMallList_ID = 431 //
//
type S2CGetShopMallList struct {
	//
	GoodsData []*ShopGoodsInfo `protobuf:"bytes,1,rep,name=GoodsData,proto3" json:"GoodsData,omitempty"`
	//
	LimitValue int32 `protobuf:"varint,2,opt,name=LimitValue,proto3" json:"LimitValue,omitempty"`
	//
	MallType int32 `protobuf:"varint,3,opt,name=MallType,proto3" json:"MallType,omitempty"`
}

func (m *S2CGetShopMallList) Reset()         { *m = S2CGetShopMallList{} }
func (m *S2CGetShopMallList) String() string { return proto.CompactTextString(m) }
func (*S2CGetShopMallList) ProtoMessage()    {}
func (m *S2CGetShopMallList) GetId() uint16  { return PCK_S2CGetShopMallList_ID }

const PCK_C2SShopBuy_ID = 432 //
//
type C2SShopBuy struct {
	//
	GoodsId int32 `protobuf:"varint,1,opt,name=GoodsId,proto3" json:"GoodsId,omitempty"`
	//
	Num int32 `protobuf:"varint,2,opt,name=Num,proto3" json:"Num,omitempty"`
}

func (m *C2SShopBuy) Reset()         { *m = C2SShopBuy{} }
func (m *C2SShopBuy) String() string { return proto.CompactTextString(m) }
func (*C2SShopBuy) ProtoMessage()    {}
func (m *C2SShopBuy) GetId() uint16  { return PCK_C2SShopBuy_ID }

const PCK_S2CShopBuy_ID = 433 //
//
type S2CShopBuy struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	GoodsId int32 `protobuf:"varint,2,opt,name=GoodsId,proto3" json:"GoodsId,omitempty"`
	//
	LimitBuy int32 `protobuf:"varint,3,opt,name=LimitBuy,proto3" json:"LimitBuy,omitempty"`
	//
	LimitUseTimes int32 `protobuf:"varint,4,opt,name=LimitUseTimes,proto3" json:"LimitUseTimes,omitempty"`
	//
	LimitTotalTimes int32 `protobuf:"varint,5,opt,name=LimitTotalTimes,proto3" json:"LimitTotalTimes,omitempty"`
	//
	MallType int32 `protobuf:"varint,6,opt,name=MallType,proto3" json:"MallType,omitempty"`
	//
	Num int32 `protobuf:"varint,8,opt,name=Num,proto3" json:"Num,omitempty"`
}

func (m *S2CShopBuy) Reset()         { *m = S2CShopBuy{} }
func (m *S2CShopBuy) String() string { return proto.CompactTextString(m) }
func (*S2CShopBuy) ProtoMessage()    {}
func (m *S2CShopBuy) GetId() uint16  { return PCK_S2CShopBuy_ID }

const PCK_C2SOrgMallList_ID = 434 //
//
type C2SOrgMallList struct {
	//
	MallType int32 `protobuf:"varint,1,opt,name=MallType,proto3" json:"MallType,omitempty"`
}

func (m *C2SOrgMallList) Reset()         { *m = C2SOrgMallList{} }
func (m *C2SOrgMallList) String() string { return proto.CompactTextString(m) }
func (*C2SOrgMallList) ProtoMessage()    {}
func (m *C2SOrgMallList) GetId() uint16  { return PCK_C2SOrgMallList_ID }

const PCK_S2COrgMallList_ID = 435 //
//
type S2COrgMallList struct {
	//
	GoodsData []*ShopGoodsInfo `protobuf:"bytes,1,rep,name=GoodsData,proto3" json:"GoodsData,omitempty"`
}

func (m *S2COrgMallList) Reset()         { *m = S2COrgMallList{} }
func (m *S2COrgMallList) String() string { return proto.CompactTextString(m) }
func (*S2COrgMallList) ProtoMessage()    {}
func (m *S2COrgMallList) GetId() uint16  { return PCK_S2COrgMallList_ID }

const PCK_C2SOrgBuy_ID = 436 //
//
type C2SOrgBuy struct {
	//
	GoodsId int32 `protobuf:"varint,1,opt,name=GoodsId,proto3" json:"GoodsId,omitempty"`
	//
	Num int32 `protobuf:"varint,2,opt,name=Num,proto3" json:"Num,omitempty"`
}

func (m *C2SOrgBuy) Reset()         { *m = C2SOrgBuy{} }
func (m *C2SOrgBuy) String() string { return proto.CompactTextString(m) }
func (*C2SOrgBuy) ProtoMessage()    {}
func (m *C2SOrgBuy) GetId() uint16  { return PCK_C2SOrgBuy_ID }

const PCK_S2COrgBuy_ID = 437 //
//
type S2COrgBuy struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	MallType int32 `protobuf:"varint,2,opt,name=MallType,proto3" json:"MallType,omitempty"`
}

func (m *S2COrgBuy) Reset()         { *m = S2COrgBuy{} }
func (m *S2COrgBuy) String() string { return proto.CompactTextString(m) }
func (*S2COrgBuy) ProtoMessage()    {}
func (m *S2COrgBuy) GetId() uint16  { return PCK_S2COrgBuy_ID }

const PCK_C2SGetGoodsLimit_ID = 438 //
//
type C2SGetGoodsLimit struct {
	//
	GoodsId int32 `protobuf:"varint,1,opt,name=GoodsId,proto3" json:"GoodsId,omitempty"`
}

func (m *C2SGetGoodsLimit) Reset()         { *m = C2SGetGoodsLimit{} }
func (m *C2SGetGoodsLimit) String() string { return proto.CompactTextString(m) }
func (*C2SGetGoodsLimit) ProtoMessage()    {}
func (m *C2SGetGoodsLimit) GetId() uint16  { return PCK_C2SGetGoodsLimit_ID }

const PCK_S2CGetGoodsLimit_ID = 439 //
//
type S2CGetGoodsLimit struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	LimitUseTimes int32 `protobuf:"varint,2,opt,name=LimitUseTimes,proto3" json:"LimitUseTimes,omitempty"`
	//
	LimitTotalTimes int32 `protobuf:"varint,3,opt,name=LimitTotalTimes,proto3" json:"LimitTotalTimes,omitempty"`
	//
	GoodsId int32 `protobuf:"varint,4,opt,name=GoodsId,proto3" json:"GoodsId,omitempty"`
}

func (m *S2CGetGoodsLimit) Reset()         { *m = S2CGetGoodsLimit{} }
func (m *S2CGetGoodsLimit) String() string { return proto.CompactTextString(m) }
func (*S2CGetGoodsLimit) ProtoMessage()    {}
func (m *S2CGetGoodsLimit) GetId() uint16  { return PCK_S2CGetGoodsLimit_ID }

//
type DbSendShop struct {
	//
	DbId string `protobuf:"bytes,1,opt,name=DbId,proto3" json:"DbId,omitempty"`
	//
	ItemData *ItemData `protobuf:"bytes,2,opt,name=ItemData,proto3" json:"ItemData,omitempty"`
	//
	TotalNum int32 `protobuf:"varint,3,opt,name=TotalNum,proto3" json:"TotalNum,omitempty"`
	//
	LeftNum int32 `protobuf:"varint,4,opt,name=LeftNum,proto3" json:"LeftNum,omitempty"`
	//
	EndTime int64 `protobuf:"varint,5,opt,name=EndTime,proto3" json:"EndTime,omitempty"`
	//
	Uid int32 `protobuf:"varint,6,opt,name=Uid,proto3" json:"Uid,omitempty"`
	//
	Sid int32 `protobuf:"varint,7,opt,name=Sid,proto3" json:"Sid,omitempty"`
	//
	Price int32 `protobuf:"varint,8,opt,name=Price,proto3" json:"Price,omitempty"`
	//
	Status int32 `protobuf:"varint,10,opt,name=Status,proto3" json:"Status,omitempty"`
	//
	LeftCoin3 int64 `protobuf:"varint,11,opt,name=LeftCoin3,proto3" json:"LeftCoin3,omitempty"`
	//
	LockTime int64 `protobuf:"varint,12,opt,name=LockTime,proto3" json:"LockTime,omitempty"`
}

func (m *DbSendShop) Reset()         { *m = DbSendShop{} }
func (m *DbSendShop) String() string { return proto.CompactTextString(m) }
func (*DbSendShop) ProtoMessage()    {}

//
type DbSaleLog struct {
	//
	BuyerId int64 `protobuf:"varint,1,opt,name=BuyerId,proto3" json:"BuyerId,omitempty"`
	//
	ItemId int32 `protobuf:"varint,2,opt,name=ItemId,proto3" json:"ItemId,omitempty"`
	//
	Num int64 `protobuf:"varint,3,opt,name=Num,proto3" json:"Num,omitempty"`
	//
	BuyTime int64 `protobuf:"varint,4,opt,name=BuyTime,proto3" json:"BuyTime,omitempty"`
	//
	Price int64 `protobuf:"varint,5,opt,name=Price,proto3" json:"Price,omitempty"`
}

func (m *DbSaleLog) Reset()         { *m = DbSaleLog{} }
func (m *DbSaleLog) String() string { return proto.CompactTextString(m) }
func (*DbSaleLog) ProtoMessage()    {}

const PCK_C2SGetSendShopList_ID = 1601 //
//
type C2SGetSendShopList struct {
	//
	Td int32 `protobuf:"varint,1,opt,name=Td,proto3" json:"Td,omitempty"`
	//
	Page int64 `protobuf:"varint,2,opt,name=Page,proto3" json:"Page,omitempty"`
	//
	Size int64 `protobuf:"varint,3,opt,name=Size,proto3" json:"Size,omitempty"`
}

func (m *C2SGetSendShopList) Reset()         { *m = C2SGetSendShopList{} }
func (m *C2SGetSendShopList) String() string { return proto.CompactTextString(m) }
func (*C2SGetSendShopList) ProtoMessage()    {}
func (m *C2SGetSendShopList) GetId() uint16  { return PCK_C2SGetSendShopList_ID }

const PCK_S2KGetSendShopList_ID = 1617 //
//
type S2KGetSendShopList struct {
	//
	Items []int32 `protobuf:"varint,1,rep,name=Items,proto3" json:"Items,omitempty"`
	//
	Td int32 `protobuf:"varint,2,opt,name=Td,proto3" json:"Td,omitempty"`
	//
	Page int64 `protobuf:"varint,3,opt,name=Page,proto3" json:"Page,omitempty"`
	//
	Size int64 `protobuf:"varint,4,opt,name=Size,proto3" json:"Size,omitempty"`
}

func (m *S2KGetSendShopList) Reset()         { *m = S2KGetSendShopList{} }
func (m *S2KGetSendShopList) String() string { return proto.CompactTextString(m) }
func (*S2KGetSendShopList) ProtoMessage()    {}
func (m *S2KGetSendShopList) GetId() uint16  { return PCK_S2KGetSendShopList_ID }

//
type SendShopGroup struct {
	//
	Iid int32 `protobuf:"varint,1,opt,name=Iid,proto3" json:"Iid,omitempty"`
	//
	Sn int32 `protobuf:"varint,2,opt,name=Sn,proto3" json:"Sn,omitempty"`
	//
	M int32 `protobuf:"varint,3,opt,name=M,proto3" json:"M,omitempty"`
}

func (m *SendShopGroup) Reset()         { *m = SendShopGroup{} }
func (m *SendShopGroup) String() string { return proto.CompactTextString(m) }
func (*SendShopGroup) ProtoMessage()    {}

const PCK_S2CGetSendShopList_ID = 1602 //
//
type S2CGetSendShopList struct {
	//
	List []*SendShopGroup `protobuf:"bytes,1,rep,name=List,proto3" json:"List,omitempty"`
	//
	Td int32 `protobuf:"varint,2,opt,name=Td,proto3" json:"Td,omitempty"`
	//
	Tag int32 `protobuf:"varint,3,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Page int64 `protobuf:"varint,4,opt,name=Page,proto3" json:"Page,omitempty"`
	//
	Total int64 `protobuf:"varint,5,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (m *S2CGetSendShopList) Reset()         { *m = S2CGetSendShopList{} }
func (m *S2CGetSendShopList) String() string { return proto.CompactTextString(m) }
func (*S2CGetSendShopList) ProtoMessage()    {}
func (m *S2CGetSendShopList) GetId() uint16  { return PCK_S2CGetSendShopList_ID }

const PCK_C2SGetSendShopById_ID = 1603 //
//
type C2SGetSendShopById struct {
	//
	Iid int32 `protobuf:"varint,1,opt,name=Iid,proto3" json:"Iid,omitempty"`
	//
	Page int64 `protobuf:"varint,2,opt,name=Page,proto3" json:"Page,omitempty"`
	//
	Size int64 `protobuf:"varint,3,opt,name=Size,proto3" json:"Size,omitempty"`
}

func (m *C2SGetSendShopById) Reset()         { *m = C2SGetSendShopById{} }
func (m *C2SGetSendShopById) String() string { return proto.CompactTextString(m) }
func (*C2SGetSendShopById) ProtoMessage()    {}
func (m *C2SGetSendShopById) GetId() uint16  { return PCK_C2SGetSendShopById_ID }

//
type SendShopSingle struct {
	//
	ItemData *ItemData `protobuf:"bytes,1,opt,name=ItemData,proto3" json:"ItemData,omitempty"`
	//
	P int32 `protobuf:"varint,2,opt,name=P,proto3" json:"P,omitempty"`
	//
	Ln int32 `protobuf:"varint,3,opt,name=Ln,proto3" json:"Ln,omitempty"`
	//
	N int32 `protobuf:"varint,4,opt,name=N,proto3" json:"N,omitempty"`
	//
	Et int64 `protobuf:"varint,5,opt,name=Et,proto3" json:"Et,omitempty"`
	//
	Lid string `protobuf:"bytes,6,opt,name=Lid,proto3" json:"Lid,omitempty"`
}

func (m *SendShopSingle) Reset()         { *m = SendShopSingle{} }
func (m *SendShopSingle) String() string { return proto.CompactTextString(m) }
func (*SendShopSingle) ProtoMessage()    {}

const PCK_S2CGetSendShopById_ID = 1604 //
//
type S2CGetSendShopById struct {
	//
	Iid int32 `protobuf:"varint,1,opt,name=Iid,proto3" json:"Iid,omitempty"`
	//
	List []*SendShopSingle `protobuf:"bytes,2,rep,name=List,proto3" json:"List,omitempty"`
	//
	Page int64 `protobuf:"varint,4,opt,name=Page,proto3" json:"Page,omitempty"`
	//
	Total int64 `protobuf:"varint,5,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (m *S2CGetSendShopById) Reset()         { *m = S2CGetSendShopById{} }
func (m *S2CGetSendShopById) String() string { return proto.CompactTextString(m) }
func (*S2CGetSendShopById) ProtoMessage()    {}
func (m *S2CGetSendShopById) GetId() uint16  { return PCK_S2CGetSendShopById_ID }

//
type SellItem struct {
	//
	ItemData *ItemData `protobuf:"bytes,1,opt,name=ItemData,proto3" json:"ItemData,omitempty"`
	//
	L int32 `protobuf:"varint,3,opt,name=L,proto3" json:"L,omitempty"`
	//
	U int32 `protobuf:"varint,4,opt,name=U,proto3" json:"U,omitempty"`
	//
	P int32 `protobuf:"varint,5,opt,name=P,proto3" json:"P,omitempty"`
	//
	OtherList []*SendShopSingle `protobuf:"bytes,6,rep,name=OtherList,proto3" json:"OtherList,omitempty"`
}

func (m *SellItem) Reset()         { *m = SellItem{} }
func (m *SellItem) String() string { return proto.CompactTextString(m) }
func (*SellItem) ProtoMessage()    {}

const PCK_C2SGetMySendShop_ID = 1611 //
//
type C2SGetMySendShop struct {
}

func (m *C2SGetMySendShop) Reset()         { *m = C2SGetMySendShop{} }
func (m *C2SGetMySendShop) String() string { return proto.CompactTextString(m) }
func (*C2SGetMySendShop) ProtoMessage()    {}
func (m *C2SGetMySendShop) GetId() uint16  { return PCK_C2SGetMySendShop_ID }

const PCK_S2KGetMySendShopList_ID = 1618 //
//
type S2KGetMySendShopList struct {
	//
	PlayerId int32 `protobuf:"varint,1,opt,name=PlayerId,proto3" json:"PlayerId,omitempty"`
	//
	Items []*ItemData `protobuf:"bytes,2,rep,name=Items,proto3" json:"Items,omitempty"`
	//
	Lt int32 `protobuf:"varint,3,opt,name=Lt,proto3" json:"Lt,omitempty"`
	//
	Page int64 `protobuf:"varint,4,opt,name=Page,proto3" json:"Page,omitempty"`
	//
	Total int64 `protobuf:"varint,5,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (m *S2KGetMySendShopList) Reset()         { *m = S2KGetMySendShopList{} }
func (m *S2KGetMySendShopList) String() string { return proto.CompactTextString(m) }
func (*S2KGetMySendShopList) ProtoMessage()    {}
func (m *S2KGetMySendShopList) GetId() uint16  { return PCK_S2KGetMySendShopList_ID }

const PCK_S2CGetMySendShop_ID = 1612 //
//
type S2CGetMySendShop struct {
	//
	List []*SendShopSingle `protobuf:"bytes,1,rep,name=List,proto3" json:"List,omitempty"`
	//
	ItemList []*SellItem `protobuf:"bytes,2,rep,name=ItemList,proto3" json:"ItemList,omitempty"`
	//
	Lt int32 `protobuf:"varint,3,opt,name=Lt,proto3" json:"Lt,omitempty"`
}

func (m *S2CGetMySendShop) Reset()         { *m = S2CGetMySendShop{} }
func (m *S2CGetMySendShop) String() string { return proto.CompactTextString(m) }
func (*S2CGetMySendShop) ProtoMessage()    {}
func (m *S2CGetMySendShop) GetId() uint16  { return PCK_S2CGetMySendShop_ID }

const PCK_C2SSendShopOnSale_ID = 1609 //
//
type C2SSendShopOnSale struct {
	//
	ItemData *ItemData `protobuf:"bytes,1,opt,name=ItemData,proto3" json:"ItemData,omitempty"`
	//
	P int32 `protobuf:"varint,3,opt,name=P,proto3" json:"P,omitempty"`
}

func (m *C2SSendShopOnSale) Reset()         { *m = C2SSendShopOnSale{} }
func (m *C2SSendShopOnSale) String() string { return proto.CompactTextString(m) }
func (*C2SSendShopOnSale) ProtoMessage()    {}
func (m *C2SSendShopOnSale) GetId() uint16  { return PCK_C2SSendShopOnSale_ID }

const PCK_S2KSendShopOnSale_ID = 1625 //
//
type S2KSendShopOnSale struct {
	//
	PlayerId int32 `protobuf:"varint,1,opt,name=PlayerId,proto3" json:"PlayerId,omitempty"`
	//
	ServerId int32 `protobuf:"varint,2,opt,name=ServerId,proto3" json:"ServerId,omitempty"`
	//
	ItemData *ItemData `protobuf:"bytes,3,opt,name=ItemData,proto3" json:"ItemData,omitempty"`
	//
	Price int32 `protobuf:"varint,5,opt,name=Price,proto3" json:"Price,omitempty"`
	//
	OrderId string `protobuf:"bytes,6,opt,name=OrderId,proto3" json:"OrderId,omitempty"`
}

func (m *S2KSendShopOnSale) Reset()         { *m = S2KSendShopOnSale{} }
func (m *S2KSendShopOnSale) String() string { return proto.CompactTextString(m) }
func (*S2KSendShopOnSale) ProtoMessage()    {}
func (m *S2KSendShopOnSale) GetId() uint16  { return PCK_S2KSendShopOnSale_ID }

const PCK_S2CSendShopOnSale_ID = 1610 //
//
type S2CSendShopOnSale struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	SaleCount int64 `protobuf:"varint,2,opt,name=SaleCount,proto3" json:"SaleCount,omitempty"`
}

func (m *S2CSendShopOnSale) Reset()         { *m = S2CSendShopOnSale{} }
func (m *S2CSendShopOnSale) String() string { return proto.CompactTextString(m) }
func (*S2CSendShopOnSale) ProtoMessage()    {}
func (m *S2CSendShopOnSale) GetId() uint16  { return PCK_S2CSendShopOnSale_ID }

const PCK_C2SSendShopBuy_ID = 1605 //
//
type C2SSendShopBuy struct {
	//
	Lid string `protobuf:"bytes,1,opt,name=Lid,proto3" json:"Lid,omitempty"`
	//
	N int32 `protobuf:"varint,2,opt,name=N,proto3" json:"N,omitempty"`
}

func (m *C2SSendShopBuy) Reset()         { *m = C2SSendShopBuy{} }
func (m *C2SSendShopBuy) String() string { return proto.CompactTextString(m) }
func (*C2SSendShopBuy) ProtoMessage()    {}
func (m *C2SSendShopBuy) GetId() uint16  { return PCK_C2SSendShopBuy_ID }

const PCK_S2KSendShopPreBuy_ID = 1619 //
//
type S2KSendShopPreBuy struct {
	//
	Lid string `protobuf:"bytes,1,opt,name=Lid,proto3" json:"Lid,omitempty"`
	//
	Num int32 `protobuf:"varint,2,opt,name=Num,proto3" json:"Num,omitempty"`
}

func (m *S2KSendShopPreBuy) Reset()         { *m = S2KSendShopPreBuy{} }
func (m *S2KSendShopPreBuy) String() string { return proto.CompactTextString(m) }
func (*S2KSendShopPreBuy) ProtoMessage()    {}
func (m *S2KSendShopPreBuy) GetId() uint16  { return PCK_S2KSendShopPreBuy_ID }

const PCK_K2SSendShopPreBuy_ID = 1620 //
//
type K2SSendShopPreBuy struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Price int32 `protobuf:"varint,2,opt,name=Price,proto3" json:"Price,omitempty"`
	//
	ItemId int32 `protobuf:"varint,3,opt,name=ItemId,proto3" json:"ItemId,omitempty"`
	//
	Num int32 `protobuf:"varint,4,opt,name=Num,proto3" json:"Num,omitempty"`
	//
	Lid string `protobuf:"bytes,5,opt,name=Lid,proto3" json:"Lid,omitempty"`
}

func (m *K2SSendShopPreBuy) Reset()         { *m = K2SSendShopPreBuy{} }
func (m *K2SSendShopPreBuy) String() string { return proto.CompactTextString(m) }
func (*K2SSendShopPreBuy) ProtoMessage()    {}
func (m *K2SSendShopPreBuy) GetId() uint16  { return PCK_K2SSendShopPreBuy_ID }

const PCK_S2KSendShopBuy_ID = 1621 //
//
type S2KSendShopBuy struct {
	//
	Lid string `protobuf:"bytes,1,opt,name=Lid,proto3" json:"Lid,omitempty"`
	//
	Num int32 `protobuf:"varint,2,opt,name=Num,proto3" json:"Num,omitempty"`
	//
	PlayerId int32 `protobuf:"varint,3,opt,name=PlayerId,proto3" json:"PlayerId,omitempty"`
	//
	OrderId string `protobuf:"bytes,4,opt,name=OrderId,proto3" json:"OrderId,omitempty"`
}

func (m *S2KSendShopBuy) Reset()         { *m = S2KSendShopBuy{} }
func (m *S2KSendShopBuy) String() string { return proto.CompactTextString(m) }
func (*S2KSendShopBuy) ProtoMessage()    {}
func (m *S2KSendShopBuy) GetId() uint16  { return PCK_S2KSendShopBuy_ID }

const PCK_K2SSendShopBuy_ID = 1633 //
//
type K2SSendShopBuy struct {
	//
	Lid string `protobuf:"bytes,1,opt,name=Lid,proto3" json:"Lid,omitempty"`
	//
	ItemData *ItemData `protobuf:"bytes,2,opt,name=ItemData,proto3" json:"ItemData,omitempty"`
	//
	Coin3 int32 `protobuf:"varint,4,opt,name=Coin3,proto3" json:"Coin3,omitempty"`
	//
	PlayerId int32 `protobuf:"varint,5,opt,name=PlayerId,proto3" json:"PlayerId,omitempty"`
	//
	Tag int32 `protobuf:"varint,6,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	OrderId string `protobuf:"bytes,7,opt,name=OrderId,proto3" json:"OrderId,omitempty"`
}

func (m *K2SSendShopBuy) Reset()         { *m = K2SSendShopBuy{} }
func (m *K2SSendShopBuy) String() string { return proto.CompactTextString(m) }
func (*K2SSendShopBuy) ProtoMessage()    {}
func (m *K2SSendShopBuy) GetId() uint16  { return PCK_K2SSendShopBuy_ID }

const PCK_S2CSendShopBuy_ID = 1606 //
//
type S2CSendShopBuy struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CSendShopBuy) Reset()         { *m = S2CSendShopBuy{} }
func (m *S2CSendShopBuy) String() string { return proto.CompactTextString(m) }
func (*S2CSendShopBuy) ProtoMessage()    {}
func (m *S2CSendShopBuy) GetId() uint16  { return PCK_S2CSendShopBuy_ID }

const PCK_C2SSendShopClear_ID = 1607 //
//
type C2SSendShopClear struct {
	//
	Lid string `protobuf:"bytes,1,opt,name=Lid,proto3" json:"Lid,omitempty"`
	//
	IsLower int64 `protobuf:"varint,2,opt,name=IsLower,proto3" json:"IsLower,omitempty"`
}

func (m *C2SSendShopClear) Reset()         { *m = C2SSendShopClear{} }
func (m *C2SSendShopClear) String() string { return proto.CompactTextString(m) }
func (*C2SSendShopClear) ProtoMessage()    {}
func (m *C2SSendShopClear) GetId() uint16  { return PCK_C2SSendShopClear_ID }

const PCK_S2KSendShopClear_ID = 1623 //
//
type S2KSendShopClear struct {
	//
	PlayerId int32 `protobuf:"varint,1,opt,name=PlayerId,proto3" json:"PlayerId,omitempty"`
	//
	Lid string `protobuf:"bytes,2,opt,name=Lid,proto3" json:"Lid,omitempty"`
	//
	IsLower int64 `protobuf:"varint,3,opt,name=IsLower,proto3" json:"IsLower,omitempty"`
}

func (m *S2KSendShopClear) Reset()         { *m = S2KSendShopClear{} }
func (m *S2KSendShopClear) String() string { return proto.CompactTextString(m) }
func (*S2KSendShopClear) ProtoMessage()    {}
func (m *S2KSendShopClear) GetId() uint16  { return PCK_S2KSendShopClear_ID }

//
type ClearInfo struct {
	//
	ItemData *ItemData `protobuf:"bytes,1,opt,name=ItemData,proto3" json:"ItemData,omitempty"`
	//
	Status int32 `protobuf:"varint,3,opt,name=Status,proto3" json:"Status,omitempty"`
	//
	Money int64 `protobuf:"varint,4,opt,name=Money,proto3" json:"Money,omitempty"`
	//
	LeftNum int64 `protobuf:"varint,5,opt,name=LeftNum,proto3" json:"LeftNum,omitempty"`
	//
	EndTime int64 `protobuf:"varint,6,opt,name=EndTime,proto3" json:"EndTime,omitempty"`
	//
	TotalNum int32 `protobuf:"varint,7,opt,name=TotalNum,proto3" json:"TotalNum,omitempty"`
	//
	Lid string `protobuf:"bytes,8,opt,name=Lid,proto3" json:"Lid,omitempty"`
}

func (m *ClearInfo) Reset()         { *m = ClearInfo{} }
func (m *ClearInfo) String() string { return proto.CompactTextString(m) }
func (*ClearInfo) ProtoMessage()    {}

const PCK_K2SSendShopClear_ID = 1624 //
//
type K2SSendShopClear struct {
	//
	Items []*ClearInfo `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty"`
	//
	SaleCount int64 `protobuf:"varint,2,opt,name=SaleCount,proto3" json:"SaleCount,omitempty"`
	//
	OrderId string `protobuf:"bytes,3,opt,name=OrderId,proto3" json:"OrderId,omitempty"`
	//
	IsLower int64 `protobuf:"varint,4,opt,name=IsLower,proto3" json:"IsLower,omitempty"`
}

func (m *K2SSendShopClear) Reset()         { *m = K2SSendShopClear{} }
func (m *K2SSendShopClear) String() string { return proto.CompactTextString(m) }
func (*K2SSendShopClear) ProtoMessage()    {}
func (m *K2SSendShopClear) GetId() uint16  { return PCK_K2SSendShopClear_ID }

const PCK_S2CSendShopClear_ID = 1608 //
//
type S2CSendShopClear struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	IsLower int64 `protobuf:"varint,2,opt,name=IsLower,proto3" json:"IsLower,omitempty"`
}

func (m *S2CSendShopClear) Reset()         { *m = S2CSendShopClear{} }
func (m *S2CSendShopClear) String() string { return proto.CompactTextString(m) }
func (*S2CSendShopClear) ProtoMessage()    {}
func (m *S2CSendShopClear) GetId() uint16  { return PCK_S2CSendShopClear_ID }

//
type SaleItemLog struct {
	//
	Iid int32 `protobuf:"varint,1,opt,name=Iid,proto3" json:"Iid,omitempty"`
	//
	N int32 `protobuf:"varint,2,opt,name=N,proto3" json:"N,omitempty"`
	//
	Gn int32 `protobuf:"varint,3,opt,name=Gn,proto3" json:"Gn,omitempty"`
	//
	St int64 `protobuf:"varint,4,opt,name=St,proto3" json:"St,omitempty"`
}

func (m *SaleItemLog) Reset()         { *m = SaleItemLog{} }
func (m *SaleItemLog) String() string { return proto.CompactTextString(m) }
func (*SaleItemLog) ProtoMessage()    {}

const PCK_C2SSaleLog_ID = 1613 //
//
type C2SSaleLog struct {
}

func (m *C2SSaleLog) Reset()         { *m = C2SSaleLog{} }
func (m *C2SSaleLog) String() string { return proto.CompactTextString(m) }
func (*C2SSaleLog) ProtoMessage()    {}
func (m *C2SSaleLog) GetId() uint16  { return PCK_C2SSaleLog_ID }

const PCK_S2KGetSaleLog_ID = 1626 //
//
type S2KGetSaleLog struct {
	//
	PlayerId int32 `protobuf:"varint,1,opt,name=PlayerId,proto3" json:"PlayerId,omitempty"`
}

func (m *S2KGetSaleLog) Reset()         { *m = S2KGetSaleLog{} }
func (m *S2KGetSaleLog) String() string { return proto.CompactTextString(m) }
func (*S2KGetSaleLog) ProtoMessage()    {}
func (m *S2KGetSaleLog) GetId() uint16  { return PCK_S2KGetSaleLog_ID }

const PCK_S2CSaleLog_ID = 1614 //
//
type S2CSaleLog struct {
	//
	List []*SaleItemLog `protobuf:"bytes,1,rep,name=List,proto3" json:"List,omitempty"`
}

func (m *S2CSaleLog) Reset()         { *m = S2CSaleLog{} }
func (m *S2CSaleLog) String() string { return proto.CompactTextString(m) }
func (*S2CSaleLog) ProtoMessage()    {}
func (m *S2CSaleLog) GetId() uint16  { return PCK_S2CSaleLog_ID }

const PCK_C2SMarkSendShop_ID = 1615 //
//
type C2SMarkSendShop struct {
	//
	Iid int32 `protobuf:"varint,1,opt,name=Iid,proto3" json:"Iid,omitempty"`
	//
	A int32 `protobuf:"varint,2,opt,name=A,proto3" json:"A,omitempty"`
}

func (m *C2SMarkSendShop) Reset()         { *m = C2SMarkSendShop{} }
func (m *C2SMarkSendShop) String() string { return proto.CompactTextString(m) }
func (*C2SMarkSendShop) ProtoMessage()    {}
func (m *C2SMarkSendShop) GetId() uint16  { return PCK_C2SMarkSendShop_ID }

const PCK_S2CMarkSendShop_ID = 1616 //
//
type S2CMarkSendShop struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Iid int32 `protobuf:"varint,2,opt,name=Iid,proto3" json:"Iid,omitempty"`
	//
	A int32 `protobuf:"varint,3,opt,name=A,proto3" json:"A,omitempty"`
}

func (m *S2CMarkSendShop) Reset()         { *m = S2CMarkSendShop{} }
func (m *S2CMarkSendShop) String() string { return proto.CompactTextString(m) }
func (*S2CMarkSendShop) ProtoMessage()    {}
func (m *S2CMarkSendShop) GetId() uint16  { return PCK_S2CMarkSendShop_ID }

const PCK_C2SGetSendCfg_ID = 1627 //
//
type C2SGetSendCfg struct {
	//
	ItemId int32 `protobuf:"varint,1,opt,name=ItemId,proto3" json:"ItemId,omitempty"`
}

func (m *C2SGetSendCfg) Reset()         { *m = C2SGetSendCfg{} }
func (m *C2SGetSendCfg) String() string { return proto.CompactTextString(m) }
func (*C2SGetSendCfg) ProtoMessage()    {}
func (m *C2SGetSendCfg) GetId() uint16  { return PCK_C2SGetSendCfg_ID }

const PCK_S2CGetSendCfg_ID = 1628 //
//
type S2CGetSendCfg struct {
	//
	T int32 `protobuf:"varint,1,opt,name=T,proto3" json:"T,omitempty"`
}

func (m *S2CGetSendCfg) Reset()         { *m = S2CGetSendCfg{} }
func (m *S2CGetSendCfg) String() string { return proto.CompactTextString(m) }
func (*S2CGetSendCfg) ProtoMessage()    {}
func (m *S2CGetSendCfg) GetId() uint16  { return PCK_S2CGetSendCfg_ID }

//
type SendUserInfo struct {
	//
	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	Nick string `protobuf:"bytes,2,opt,name=Nick,proto3" json:"Nick,omitempty"`
	//
	FightValue int64 `protobuf:"varint,3,opt,name=FightValue,proto3" json:"FightValue,omitempty"`
	//
	Level int64 `protobuf:"varint,4,opt,name=Level,proto3" json:"Level,omitempty"`
	//
	VipLevel int64 `protobuf:"varint,5,opt,name=VipLevel,proto3" json:"VipLevel,omitempty"`
	//
	RoleId int64 `protobuf:"varint,6,opt,name=RoleId,proto3" json:"RoleId,omitempty"`
	//
	GodLevel int64 `protobuf:"varint,7,opt,name=GodLevel,proto3" json:"GodLevel,omitempty"`
	//
	AreaId int64 `protobuf:"varint,8,opt,name=AreaId,proto3" json:"AreaId,omitempty"`
	//
	ServerId int64 `protobuf:"varint,9,opt,name=ServerId,proto3" json:"ServerId,omitempty"`
}

func (m *SendUserInfo) Reset()         { *m = SendUserInfo{} }
func (m *SendUserInfo) String() string { return proto.CompactTextString(m) }
func (*SendUserInfo) ProtoMessage()    {}

const PCK_C2SGetSendUserInfo_ID = 1631 //
//
type C2SGetSendUserInfo struct {
	//
	UseId int64 `protobuf:"varint,1,opt,name=UseId,proto3" json:"UseId,omitempty"`
}

func (m *C2SGetSendUserInfo) Reset()         { *m = C2SGetSendUserInfo{} }
func (m *C2SGetSendUserInfo) String() string { return proto.CompactTextString(m) }
func (*C2SGetSendUserInfo) ProtoMessage()    {}
func (m *C2SGetSendUserInfo) GetId() uint16  { return PCK_C2SGetSendUserInfo_ID }

const PCK_S2CGetSendUserInfo_ID = 1632 //
//
type S2CGetSendUserInfo struct {
	//
	Tag int64 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Info *SendUserInfo `protobuf:"bytes,2,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (m *S2CGetSendUserInfo) Reset()         { *m = S2CGetSendUserInfo{} }
func (m *S2CGetSendUserInfo) String() string { return proto.CompactTextString(m) }
func (*S2CGetSendUserInfo) ProtoMessage()    {}
func (m *S2CGetSendUserInfo) GetId() uint16  { return PCK_S2CGetSendUserInfo_ID }

const PCK_C2SSendItem_ID = 1629 //
//
type C2SSendItem struct {
	//
	Uid int32 `protobuf:"varint,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
	//
	Iid int32 `protobuf:"varint,2,opt,name=Iid,proto3" json:"Iid,omitempty"`
	//
	In int32 `protobuf:"varint,3,opt,name=In,proto3" json:"In,omitempty"`
}

func (m *C2SSendItem) Reset()         { *m = C2SSendItem{} }
func (m *C2SSendItem) String() string { return proto.CompactTextString(m) }
func (*C2SSendItem) ProtoMessage()    {}
func (m *C2SSendItem) GetId() uint16  { return PCK_C2SSendItem_ID }

const PCK_S2CSendItem_ID = 1630 //
//
type S2CSendItem struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CSendItem) Reset()         { *m = S2CSendItem{} }
func (m *S2CSendItem) String() string { return proto.CompactTextString(m) }
func (*S2CSendItem) ProtoMessage()    {}
func (m *S2CSendItem) GetId() uint16  { return PCK_S2CSendItem_ID }

//
type Counter struct {
	//
	T int32 `protobuf:"varint,1,opt,name=T,proto3" json:"T,omitempty"`
	//
	C int64 `protobuf:"varint,2,opt,name=C,proto3" json:"C,omitempty"`
	//
	UT string `protobuf:"bytes,3,opt,name=UT,proto3" json:"UT,omitempty"`
	//
	P int32 `protobuf:"varint,4,opt,name=P,proto3" json:"P,omitempty"`
}

func (m *Counter) Reset()         { *m = Counter{} }
func (m *Counter) String() string { return proto.CompactTextString(m) }
func (*Counter) ProtoMessage()    {}

//
type S2CCounter struct {
	//
	T int32 `protobuf:"varint,1,opt,name=T,proto3" json:"T,omitempty"`
	//
	C int64 `protobuf:"varint,2,opt,name=C,proto3" json:"C,omitempty"`
	//
	P int32 `protobuf:"varint,4,opt,name=P,proto3" json:"P,omitempty"`
}

func (m *S2CCounter) Reset()         { *m = S2CCounter{} }
func (m *S2CCounter) String() string { return proto.CompactTextString(m) }
func (*S2CCounter) ProtoMessage()    {}

//
type Task struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	T int32 `protobuf:"varint,8,opt,name=T,proto3" json:"T,omitempty"`
	//
	IC int64 `protobuf:"varint,4,opt,name=IC,proto3" json:"IC,omitempty"`
	//
	C int32 `protobuf:"varint,5,opt,name=C,proto3" json:"C,omitempty"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}

//
type FinishTask struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	T int32 `protobuf:"varint,2,opt,name=T,proto3" json:"T,omitempty"`
	//
	IC int64 `protobuf:"varint,3,opt,name=IC,proto3" json:"IC,omitempty"`
	//
	C int32 `protobuf:"varint,4,opt,name=C,proto3" json:"C,omitempty"`
}

func (m *FinishTask) Reset()         { *m = FinishTask{} }
func (m *FinishTask) String() string { return proto.CompactTextString(m) }
func (*FinishTask) ProtoMessage()    {}

//
type S2CTask struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	T int32 `protobuf:"varint,4,opt,name=T,proto3" json:"T,omitempty"`
	//
	S int32 `protobuf:"varint,2,opt,name=S,proto3" json:"S,omitempty"`
	//
	IC int64 `protobuf:"varint,3,opt,name=IC,proto3" json:"IC,omitempty"`
}

func (m *S2CTask) Reset()         { *m = S2CTask{} }
func (m *S2CTask) String() string { return proto.CompactTextString(m) }
func (*S2CTask) ProtoMessage()    {}

//
type DBDailyTask struct {
	//
	T []*DailyTask `protobuf:"bytes,1,rep,name=T,proto3" json:"T,omitempty"`
	//
	Key string `protobuf:"bytes,2,opt,name=Key,proto3" json:"Key,omitempty"`
}

func (m *DBDailyTask) Reset()         { *m = DBDailyTask{} }
func (m *DBDailyTask) String() string { return proto.CompactTextString(m) }
func (*DBDailyTask) ProtoMessage()    {}

//
type DailyTask struct {
	//
	Uid int32 `protobuf:"varint,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
	//
	ID int32 `protobuf:"varint,2,opt,name=ID,proto3" json:"ID,omitempty"`
	//
	T int32 `protobuf:"varint,3,opt,name=T,proto3" json:"T,omitempty"`
	//
	S int32 `protobuf:"varint,4,opt,name=S,proto3" json:"S,omitempty"`
	//
	I int64 `protobuf:"varint,5,opt,name=I,proto3" json:"I,omitempty"`
	//
	C int32 `protobuf:"varint,6,opt,name=C,proto3" json:"C,omitempty"`
}

func (m *DailyTask) Reset()         { *m = DailyTask{} }
func (m *DailyTask) String() string { return proto.CompactTextString(m) }
func (*DailyTask) ProtoMessage()    {}

const PCK_C2SGetTaskPrize_ID = 703 //
//
type C2SGetTaskPrize struct {
	//
	TaskId int32 `protobuf:"varint,1,opt,name=TaskId,proto3" json:"TaskId,omitempty"`
	//
	TaskType int32 `protobuf:"varint,3,opt,name=TaskType,proto3" json:"TaskType,omitempty"`
	//
	Multi int32 `protobuf:"varint,2,opt,name=Multi,proto3" json:"Multi,omitempty"`
}

func (m *C2SGetTaskPrize) Reset()         { *m = C2SGetTaskPrize{} }
func (m *C2SGetTaskPrize) String() string { return proto.CompactTextString(m) }
func (*C2SGetTaskPrize) ProtoMessage()    {}
func (m *C2SGetTaskPrize) GetId() uint16  { return PCK_C2SGetTaskPrize_ID }

const PCK_S2CGetTaskPrize_ID = 704 //
//
type S2CGetTaskPrize struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Prize []*ItemInfo `protobuf:"bytes,2,rep,name=Prize,proto3" json:"Prize,omitempty"`
}

func (m *S2CGetTaskPrize) Reset()         { *m = S2CGetTaskPrize{} }
func (m *S2CGetTaskPrize) String() string { return proto.CompactTextString(m) }
func (*S2CGetTaskPrize) ProtoMessage()    {}
func (m *S2CGetTaskPrize) GetId() uint16  { return PCK_S2CGetTaskPrize_ID }

const PCK_S2CHistoryTaskInfo_ID = 741 //
//
type S2CHistoryTaskInfo struct {
	//
	MainTaskId int32 `protobuf:"varint,1,opt,name=MainTaskId,proto3" json:"MainTaskId,omitempty"`
	//
	NowStage int32 `protobuf:"varint,2,opt,name=NowStage,proto3" json:"NowStage,omitempty"`
	//
	NewMapId int32 `protobuf:"varint,3,opt,name=NewMapId,proto3" json:"NewMapId,omitempty"`
	//
	StateLevel int32 `protobuf:"varint,4,opt,name=StateLevel,proto3" json:"StateLevel,omitempty"`
}

func (m *S2CHistoryTaskInfo) Reset()         { *m = S2CHistoryTaskInfo{} }
func (m *S2CHistoryTaskInfo) String() string { return proto.CompactTextString(m) }
func (*S2CHistoryTaskInfo) ProtoMessage()    {}
func (m *S2CHistoryTaskInfo) GetId() uint16  { return PCK_S2CHistoryTaskInfo_ID }

const PCK_C2SGetHistoryTaskPrize_ID = 713 //
//
type C2SGetHistoryTaskPrize struct {
	//
	TaskId int32 `protobuf:"varint,1,opt,name=TaskId,proto3" json:"TaskId,omitempty"`
}

func (m *C2SGetHistoryTaskPrize) Reset()         { *m = C2SGetHistoryTaskPrize{} }
func (m *C2SGetHistoryTaskPrize) String() string { return proto.CompactTextString(m) }
func (*C2SGetHistoryTaskPrize) ProtoMessage()    {}
func (m *C2SGetHistoryTaskPrize) GetId() uint16  { return PCK_C2SGetHistoryTaskPrize_ID }

const PCK_S2CGetHistoryTaskPrize_ID = 714 //
//
type S2CGetHistoryTaskPrize struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Prize []*ItemInfo `protobuf:"bytes,2,rep,name=Prize,proto3" json:"Prize,omitempty"`
	//
	NextTask *S2CTask `protobuf:"bytes,3,opt,name=NextTask,proto3" json:"NextTask,omitempty"`
}

func (m *S2CGetHistoryTaskPrize) Reset()         { *m = S2CGetHistoryTaskPrize{} }
func (m *S2CGetHistoryTaskPrize) String() string { return proto.CompactTextString(m) }
func (*S2CGetHistoryTaskPrize) ProtoMessage()    {}
func (m *S2CGetHistoryTaskPrize) GetId() uint16  { return PCK_S2CGetHistoryTaskPrize_ID }

//
type LifeRemain struct {
	//
	TaskId int32 `protobuf:"varint,1,opt,name=TaskId,proto3" json:"TaskId,omitempty"`
	//
	Type int32 `protobuf:"varint,2,opt,name=Type,proto3" json:"Type,omitempty"`
	//
	Times int32 `protobuf:"varint,3,opt,name=Times,proto3" json:"Times,omitempty"`
}

func (m *LifeRemain) Reset()         { *m = LifeRemain{} }
func (m *LifeRemain) String() string { return proto.CompactTextString(m) }
func (*LifeRemain) ProtoMessage()    {}

const PCK_S2CLastDayRemain_ID = 710 //
//
type S2CLastDayRemain struct {
	//
	List []*LifeRemain `protobuf:"bytes,1,rep,name=List,proto3" json:"List,omitempty"`
}

func (m *S2CLastDayRemain) Reset()         { *m = S2CLastDayRemain{} }
func (m *S2CLastDayRemain) String() string { return proto.CompactTextString(m) }
func (*S2CLastDayRemain) ProtoMessage()    {}
func (m *S2CLastDayRemain) GetId() uint16  { return PCK_S2CLastDayRemain_ID }

const PCK_C2SLifeFind_ID = 711 //
//
type C2SLifeFind struct {
	//
	TaskId int32 `protobuf:"varint,1,opt,name=TaskId,proto3" json:"TaskId,omitempty"`
	//
	Type int32 `protobuf:"varint,2,opt,name=Type,proto3" json:"Type,omitempty"`
	//
	Count int32 `protobuf:"varint,3,opt,name=Count,proto3" json:"Count,omitempty"`
}

func (m *C2SLifeFind) Reset()         { *m = C2SLifeFind{} }
func (m *C2SLifeFind) String() string { return proto.CompactTextString(m) }
func (*C2SLifeFind) ProtoMessage()    {}
func (m *C2SLifeFind) GetId() uint16  { return PCK_C2SLifeFind_ID }

const PCK_S2CLifeFind_ID = 712 //
//
type S2CLifeFind struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Remain *LifeRemain `protobuf:"bytes,2,opt,name=Remain,proto3" json:"Remain,omitempty"`
}

func (m *S2CLifeFind) Reset()         { *m = S2CLifeFind{} }
func (m *S2CLifeFind) String() string { return proto.CompactTextString(m) }
func (*S2CLifeFind) ProtoMessage()    {}
func (m *S2CLifeFind) GetId() uint16  { return PCK_S2CLifeFind_ID }

const PCK_C2SLifeFastFind_ID = 717 //
//
type C2SLifeFastFind struct {
}

func (m *C2SLifeFastFind) Reset()         { *m = C2SLifeFastFind{} }
func (m *C2SLifeFastFind) String() string { return proto.CompactTextString(m) }
func (*C2SLifeFastFind) ProtoMessage()    {}
func (m *C2SLifeFastFind) GetId() uint16  { return PCK_C2SLifeFastFind_ID }

const PCK_S2CLifeFastFind_ID = 718 //
//
type S2CLifeFastFind struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CLifeFastFind) Reset()         { *m = S2CLifeFastFind{} }
func (m *S2CLifeFastFind) String() string { return proto.CompactTextString(m) }
func (*S2CLifeFastFind) ProtoMessage()    {}
func (m *S2CLifeFastFind) GetId() uint16  { return PCK_S2CLifeFastFind_ID }

const PCK_C2SWorldLevel_ID = 719 //
//
type C2SWorldLevel struct {
}

func (m *C2SWorldLevel) Reset()         { *m = C2SWorldLevel{} }
func (m *C2SWorldLevel) String() string { return proto.CompactTextString(m) }
func (*C2SWorldLevel) ProtoMessage()    {}
func (m *C2SWorldLevel) GetId() uint16  { return PCK_C2SWorldLevel_ID }

//
type WorldLevel struct {
	//
	Level int32 `protobuf:"varint,1,opt,name=Level,proto3" json:"Level,omitempty"`
	//
	Players []*Rank `protobuf:"bytes,2,rep,name=Players,proto3" json:"Players,omitempty"`
}

func (m *WorldLevel) Reset()         { *m = WorldLevel{} }
func (m *WorldLevel) String() string { return proto.CompactTextString(m) }
func (*WorldLevel) ProtoMessage()    {}

const PCK_C2SPeaceFinish_ID = 715 //
//
type C2SPeaceFinish struct {
}

func (m *C2SPeaceFinish) Reset()         { *m = C2SPeaceFinish{} }
func (m *C2SPeaceFinish) String() string { return proto.CompactTextString(m) }
func (*C2SPeaceFinish) ProtoMessage()    {}
func (m *C2SPeaceFinish) GetId() uint16  { return PCK_C2SPeaceFinish_ID }

const PCK_S2CPeaceFinish_ID = 716 //
//
type S2CPeaceFinish struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CPeaceFinish) Reset()         { *m = S2CPeaceFinish{} }
func (m *S2CPeaceFinish) String() string { return proto.CompactTextString(m) }
func (*S2CPeaceFinish) ProtoMessage()    {}
func (m *S2CPeaceFinish) GetId() uint16  { return PCK_S2CPeaceFinish_ID }

const PCK_C2SSignPrize_ID = 730 //
//
type C2SSignPrize struct {
}

func (m *C2SSignPrize) Reset()         { *m = C2SSignPrize{} }
func (m *C2SSignPrize) String() string { return proto.CompactTextString(m) }
func (*C2SSignPrize) ProtoMessage()    {}
func (m *C2SSignPrize) GetId() uint16  { return PCK_C2SSignPrize_ID }

const PCK_S2CSignPrize_ID = 731 //
//
type S2CSignPrize struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CSignPrize) Reset()         { *m = S2CSignPrize{} }
func (m *S2CSignPrize) String() string { return proto.CompactTextString(m) }
func (*S2CSignPrize) ProtoMessage()    {}
func (m *S2CSignPrize) GetId() uint16  { return PCK_S2CSignPrize_ID }

//
type SMMonsterInfo struct {
	//
	RefreshId int32 `protobuf:"varint,1,opt,name=RefreshId,proto3" json:"RefreshId,omitempty"`
	//
	Star int32 `protobuf:"varint,2,opt,name=Star,proto3" json:"Star,omitempty"`
}

func (m *SMMonsterInfo) Reset()         { *m = SMMonsterInfo{} }
func (m *SMMonsterInfo) String() string { return proto.CompactTextString(m) }
func (*SMMonsterInfo) ProtoMessage()    {}

//
type SMMonster struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Star int32 `protobuf:"varint,2,opt,name=Star,proto3" json:"Star,omitempty"`
}

func (m *SMMonster) Reset()         { *m = SMMonster{} }
func (m *SMMonster) String() string { return proto.CompactTextString(m) }
func (*SMMonster) ProtoMessage()    {}

const PCK_S2CSMMonster_ID = 720 //
//
type S2CSMMonster struct {
	//
	Monster []*SMMonster `protobuf:"bytes,1,rep,name=Monster,proto3" json:"Monster,omitempty"`
}

func (m *S2CSMMonster) Reset()         { *m = S2CSMMonster{} }
func (m *S2CSMMonster) String() string { return proto.CompactTextString(m) }
func (*S2CSMMonster) ProtoMessage()    {}
func (m *S2CSMMonster) GetId() uint16  { return PCK_S2CSMMonster_ID }

const PCK_C2SSMFight_ID = 721 //
//
type C2SSMFight struct {
	//
	Monster *SMMonster `protobuf:"bytes,1,opt,name=Monster,proto3" json:"Monster,omitempty"`
}

func (m *C2SSMFight) Reset()         { *m = C2SSMFight{} }
func (m *C2SSMFight) String() string { return proto.CompactTextString(m) }
func (*C2SSMFight) ProtoMessage()    {}
func (m *C2SSMFight) GetId() uint16  { return PCK_C2SSMFight_ID }

const PCK_S2CSMFight_ID = 722 //
//
type S2CSMFight struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Prize []*ItemInfo `protobuf:"bytes,2,rep,name=Prize,proto3" json:"Prize,omitempty"`
}

func (m *S2CSMFight) Reset()         { *m = S2CSMFight{} }
func (m *S2CSMFight) String() string { return proto.CompactTextString(m) }
func (*S2CSMFight) ProtoMessage()    {}
func (m *S2CSMFight) GetId() uint16  { return PCK_S2CSMFight_ID }

const PCK_C2SSMRefreshStar_ID = 723 //
//
type C2SSMRefreshStar struct {
	//
	Monster *SMMonster `protobuf:"bytes,1,opt,name=Monster,proto3" json:"Monster,omitempty"`
}

func (m *C2SSMRefreshStar) Reset()         { *m = C2SSMRefreshStar{} }
func (m *C2SSMRefreshStar) String() string { return proto.CompactTextString(m) }
func (*C2SSMRefreshStar) ProtoMessage()    {}
func (m *C2SSMRefreshStar) GetId() uint16  { return PCK_C2SSMRefreshStar_ID }

const PCK_S2CSMRefreshStar_ID = 724 //
//
type S2CSMRefreshStar struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	OldMonster *SMMonster `protobuf:"bytes,2,opt,name=OldMonster,proto3" json:"OldMonster,omitempty"`
	//
	NewMonster *SMMonster `protobuf:"bytes,3,opt,name=NewMonster,proto3" json:"NewMonster,omitempty"`
}

func (m *S2CSMRefreshStar) Reset()         { *m = S2CSMRefreshStar{} }
func (m *S2CSMRefreshStar) String() string { return proto.CompactTextString(m) }
func (*S2CSMRefreshStar) ProtoMessage()    {}
func (m *S2CSMRefreshStar) GetId() uint16  { return PCK_S2CSMRefreshStar_ID }

const PCK_C2SSMFastFinish_ID = 725 //
//
type C2SSMFastFinish struct {
}

func (m *C2SSMFastFinish) Reset()         { *m = C2SSMFastFinish{} }
func (m *C2SSMFastFinish) String() string { return proto.CompactTextString(m) }
func (*C2SSMFastFinish) ProtoMessage()    {}
func (m *C2SSMFastFinish) GetId() uint16  { return PCK_C2SSMFastFinish_ID }

const PCK_S2CSMFastFinish_ID = 726 //
//
type S2CSMFastFinish struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Prize []*ItemInfo `protobuf:"bytes,2,rep,name=Prize,proto3" json:"Prize,omitempty"`
}

func (m *S2CSMFastFinish) Reset()         { *m = S2CSMFastFinish{} }
func (m *S2CSMFastFinish) String() string { return proto.CompactTextString(m) }
func (*S2CSMFastFinish) ProtoMessage()    {}
func (m *S2CSMFastFinish) GetId() uint16  { return PCK_S2CSMFastFinish_ID }

//
type MailInfo struct {
	//
	MailId int32 `protobuf:"varint,1,opt,name=MailId,proto3" json:"MailId,omitempty"`
	//
	MailTplId int32 `protobuf:"varint,2,opt,name=MailTplId,proto3" json:"MailTplId,omitempty"`
	//
	MailTplParam string `protobuf:"bytes,4,opt,name=MailTplParam,proto3" json:"MailTplParam,omitempty"`
	//
	Title string `protobuf:"bytes,5,opt,name=Title,proto3" json:"Title,omitempty"`
	//
	MailType int32 `protobuf:"varint,6,opt,name=MailType,proto3" json:"MailType,omitempty"`
	//
	IsReceive int32 `protobuf:"varint,7,opt,name=IsReceive,proto3" json:"IsReceive,omitempty"`
	//
	ReceiveTime string `protobuf:"bytes,8,opt,name=ReceiveTime,proto3" json:"ReceiveTime,omitempty"`
	//
	Content string `protobuf:"bytes,9,opt,name=Content,proto3" json:"Content,omitempty"`
	//
	AttachInfo []*ItemInfo `protobuf:"bytes,10,rep,name=AttachInfo,proto3" json:"AttachInfo,omitempty"`
	//
	AttachData []*ItemData `protobuf:"bytes,11,rep,name=AttachData,proto3" json:"AttachData,omitempty"`
}

func (m *MailInfo) Reset()         { *m = MailInfo{} }
func (m *MailInfo) String() string { return proto.CompactTextString(m) }
func (*MailInfo) ProtoMessage()    {}

//
type DbMailInfo struct {
	//
	MailId int32 `protobuf:"varint,1,opt,name=MailId,proto3" json:"MailId,omitempty"`
	//
	MailTplId int32 `protobuf:"varint,2,opt,name=MailTplId,proto3" json:"MailTplId,omitempty"`
	//
	MailTplParam string `protobuf:"bytes,3,opt,name=MailTplParam,proto3" json:"MailTplParam,omitempty"`
	//
	MailType int32 `protobuf:"varint,4,opt,name=MailType,proto3" json:"MailType,omitempty"`
	//
	IsReceive int32 `protobuf:"varint,5,opt,name=IsReceive,proto3" json:"IsReceive,omitempty"`
	//
	Content string `protobuf:"bytes,6,opt,name=Content,proto3" json:"Content,omitempty"`
	//
	ReceiveTime int64 `protobuf:"varint,7,opt,name=ReceiveTime,proto3" json:"ReceiveTime,omitempty"`
	//
	ChangeType int32 `protobuf:"varint,8,opt,name=ChangeType,proto3" json:"ChangeType,omitempty"`
	//
	SendUserId int32 `protobuf:"varint,9,opt,name=SendUserId,proto3" json:"SendUserId,omitempty"`
	//
	SendUserNick string `protobuf:"bytes,10,opt,name=SendUserNick,proto3" json:"SendUserNick,omitempty"`
	//
	Title string `protobuf:"bytes,11,opt,name=Title,proto3" json:"Title,omitempty"`
	//
	AttachInfo []*ItemInfo `protobuf:"bytes,12,rep,name=AttachInfo,proto3" json:"AttachInfo,omitempty"`
	//
	AttachData []*ItemData `protobuf:"bytes,13,rep,name=AttachData,proto3" json:"AttachData,omitempty"`
}

func (m *DbMailInfo) Reset()         { *m = DbMailInfo{} }
func (m *DbMailInfo) String() string { return proto.CompactTextString(m) }
func (*DbMailInfo) ProtoMessage()    {}

//
type RoleMail struct {
	//
	SystemMailId int32 `protobuf:"varint,1,opt,name=SystemMailId,proto3" json:"SystemMailId,omitempty"`
	//
	UserMailId int32 `protobuf:"varint,2,opt,name=UserMailId,proto3" json:"UserMailId,omitempty"`
	//
	DbMailId int32 `protobuf:"varint,3,opt,name=DbMailId,proto3" json:"DbMailId,omitempty"`
	//
	DbMailInfo []*DbMailInfo `protobuf:"bytes,4,rep,name=DbMailInfo,proto3" json:"DbMailInfo,omitempty"`
}

func (m *RoleMail) Reset()         { *m = RoleMail{} }
func (m *RoleMail) String() string { return proto.CompactTextString(m) }
func (*RoleMail) ProtoMessage()    {}

const PCK_C2SMailList_ID = 440 //
//
type C2SMailList struct {
}

func (m *C2SMailList) Reset()         { *m = C2SMailList{} }
func (m *C2SMailList) String() string { return proto.CompactTextString(m) }
func (*C2SMailList) ProtoMessage()    {}
func (m *C2SMailList) GetId() uint16  { return PCK_C2SMailList_ID }

const PCK_S2CMailList_ID = 441 //
//
type S2CMailList struct {
	//
	MailList []*MailInfo `protobuf:"bytes,1,rep,name=MailList,proto3" json:"MailList,omitempty"`
}

func (m *S2CMailList) Reset()         { *m = S2CMailList{} }
func (m *S2CMailList) String() string { return proto.CompactTextString(m) }
func (*S2CMailList) ProtoMessage()    {}
func (m *S2CMailList) GetId() uint16  { return PCK_S2CMailList_ID }

const PCK_C2SGetMailAttach_ID = 444 //
//
type C2SGetMailAttach struct {
	//
	MailId int32 `protobuf:"varint,1,opt,name=MailId,proto3" json:"MailId,omitempty"`
}

func (m *C2SGetMailAttach) Reset()         { *m = C2SGetMailAttach{} }
func (m *C2SGetMailAttach) String() string { return proto.CompactTextString(m) }
func (*C2SGetMailAttach) ProtoMessage()    {}
func (m *C2SGetMailAttach) GetId() uint16  { return PCK_C2SGetMailAttach_ID }

const PCK_S2CGetMailAttach_ID = 445 //
//
type S2CGetMailAttach struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CGetMailAttach) Reset()         { *m = S2CGetMailAttach{} }
func (m *S2CGetMailAttach) String() string { return proto.CompactTextString(m) }
func (*S2CGetMailAttach) ProtoMessage()    {}
func (m *S2CGetMailAttach) GetId() uint16  { return PCK_S2CGetMailAttach_ID }

const PCK_C2SDelMail_ID = 457 //
//
type C2SDelMail struct {
	//
	MailId int32 `protobuf:"varint,1,opt,name=MailId,proto3" json:"MailId,omitempty"`
}

func (m *C2SDelMail) Reset()         { *m = C2SDelMail{} }
func (m *C2SDelMail) String() string { return proto.CompactTextString(m) }
func (*C2SDelMail) ProtoMessage()    {}
func (m *C2SDelMail) GetId() uint16  { return PCK_C2SDelMail_ID }

const PCK_S2CDelMail_ID = 458 //
//
type S2CDelMail struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CDelMail) Reset()         { *m = S2CDelMail{} }
func (m *S2CDelMail) String() string { return proto.CompactTextString(m) }
func (*S2CDelMail) ProtoMessage()    {}
func (m *S2CDelMail) GetId() uint16  { return PCK_S2CDelMail_ID }

const PCK_C2SGM_ID = 5 //
//
type C2SGM struct {
	//
	Data string `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *C2SGM) Reset()         { *m = C2SGM{} }
func (m *C2SGM) String() string { return proto.CompactTextString(m) }
func (*C2SGM) ProtoMessage()    {}
func (m *C2SGM) GetId() uint16  { return PCK_C2SGM_ID }

const PCK_S2CNewMail_ID = 446 //
//
type S2CNewMail struct {
}

func (m *S2CNewMail) Reset()         { *m = S2CNewMail{} }
func (m *S2CNewMail) String() string { return proto.CompactTextString(m) }
func (*S2CNewMail) ProtoMessage()    {}
func (m *S2CNewMail) GetId() uint16  { return PCK_S2CNewMail_ID }

//
type BossPersonalInfo struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	K int32 `protobuf:"varint,2,opt,name=K,proto3" json:"K,omitempty"`
}

func (m *BossPersonalInfo) Reset()         { *m = BossPersonalInfo{} }
func (m *BossPersonalInfo) String() string { return proto.CompactTextString(m) }
func (*BossPersonalInfo) ProtoMessage()    {}

const PCK_S2CBossPersonalInfo_ID = 601 //
//
type S2CBossPersonalInfo struct {
	//
	Items []*BossPersonalInfo `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty"`
}

func (m *S2CBossPersonalInfo) Reset()         { *m = S2CBossPersonalInfo{} }
func (m *S2CBossPersonalInfo) String() string { return proto.CompactTextString(m) }
func (*S2CBossPersonalInfo) ProtoMessage()    {}
func (m *S2CBossPersonalInfo) GetId() uint16  { return PCK_S2CBossPersonalInfo_ID }

const PCK_C2SBossPersonalFight_ID = 602 //
//
type C2SBossPersonalFight struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2SBossPersonalFight) Reset()         { *m = C2SBossPersonalFight{} }
func (m *C2SBossPersonalFight) String() string { return proto.CompactTextString(m) }
func (*C2SBossPersonalFight) ProtoMessage()    {}
func (m *C2SBossPersonalFight) GetId() uint16  { return PCK_C2SBossPersonalFight_ID }

const PCK_S2CBossPersonalFight_ID = 603 //
//
type S2CBossPersonalFight struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Win int64 `protobuf:"varint,3,opt,name=Win,proto3" json:"Win,omitempty"`
}

func (m *S2CBossPersonalFight) Reset()         { *m = S2CBossPersonalFight{} }
func (m *S2CBossPersonalFight) String() string { return proto.CompactTextString(m) }
func (*S2CBossPersonalFight) ProtoMessage()    {}
func (m *S2CBossPersonalFight) GetId() uint16  { return PCK_S2CBossPersonalFight_ID }

const PCK_C2SBossPersonalSweep_ID = 604 //
//
type C2SBossPersonalSweep struct {
}

func (m *C2SBossPersonalSweep) Reset()         { *m = C2SBossPersonalSweep{} }
func (m *C2SBossPersonalSweep) String() string { return proto.CompactTextString(m) }
func (*C2SBossPersonalSweep) ProtoMessage()    {}
func (m *C2SBossPersonalSweep) GetId() uint16  { return PCK_C2SBossPersonalSweep_ID }

const PCK_S2CBossPersonalSweep_ID = 605 //
//
type S2CBossPersonalSweep struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Ids []int32 `protobuf:"varint,2,rep,name=Ids,proto3" json:"Ids,omitempty"`
}

func (m *S2CBossPersonalSweep) Reset()         { *m = S2CBossPersonalSweep{} }
func (m *S2CBossPersonalSweep) String() string { return proto.CompactTextString(m) }
func (*S2CBossPersonalSweep) ProtoMessage()    {}
func (m *S2CBossPersonalSweep) GetId() uint16  { return PCK_S2CBossPersonalSweep_ID }

//
type InstanceMaterialData struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Kill int32 `protobuf:"varint,2,opt,name=Kill,proto3" json:"Kill,omitempty"`
	//
	FightTimes int32 `protobuf:"varint,3,opt,name=FightTimes,proto3" json:"FightTimes,omitempty"`
	//
	BuyTimes int32 `protobuf:"varint,4,opt,name=BuyTimes,proto3" json:"BuyTimes,omitempty"`
	//
	NextTime int64 `protobuf:"varint,5,opt,name=NextTime,proto3" json:"NextTime,omitempty"`
}

func (m *InstanceMaterialData) Reset()         { *m = InstanceMaterialData{} }
func (m *InstanceMaterialData) String() string { return proto.CompactTextString(m) }
func (*InstanceMaterialData) ProtoMessage()    {}

//
type InstanceMaterialInfo struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Kill int32 `protobuf:"varint,2,opt,name=Kill,proto3" json:"Kill,omitempty"`
	//
	LeftTimes int32 `protobuf:"varint,3,opt,name=LeftTimes,proto3" json:"LeftTimes,omitempty"`
	//
	BuyTimes int32 `protobuf:"varint,4,opt,name=BuyTimes,proto3" json:"BuyTimes,omitempty"`
	//
	VipBuyMax int32 `protobuf:"varint,5,opt,name=VipBuyMax,proto3" json:"VipBuyMax,omitempty"`
	//
	Coin int64 `protobuf:"varint,6,opt,name=Coin,proto3" json:"Coin,omitempty"`
	//
	NextTime int64 `protobuf:"varint,7,opt,name=NextTime,proto3" json:"NextTime,omitempty"`
}

func (m *InstanceMaterialInfo) Reset()         { *m = InstanceMaterialInfo{} }
func (m *InstanceMaterialInfo) String() string { return proto.CompactTextString(m) }
func (*InstanceMaterialInfo) ProtoMessage()    {}

const PCK_S2CInstanceMaterialInfo_ID = 606 //
//
type S2CInstanceMaterialInfo struct {
	//
	Items []*InstanceMaterialInfo `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty"`
}

func (m *S2CInstanceMaterialInfo) Reset()         { *m = S2CInstanceMaterialInfo{} }
func (m *S2CInstanceMaterialInfo) String() string { return proto.CompactTextString(m) }
func (*S2CInstanceMaterialInfo) ProtoMessage()    {}
func (m *S2CInstanceMaterialInfo) GetId() uint16  { return PCK_S2CInstanceMaterialInfo_ID }

const PCK_C2SInstanceMaterialFight_ID = 607 //
//
type C2SInstanceMaterialFight struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Param int32 `protobuf:"varint,2,opt,name=Param,proto3" json:"Param,omitempty"`
}

func (m *C2SInstanceMaterialFight) Reset()         { *m = C2SInstanceMaterialFight{} }
func (m *C2SInstanceMaterialFight) String() string { return proto.CompactTextString(m) }
func (*C2SInstanceMaterialFight) ProtoMessage()    {}
func (m *C2SInstanceMaterialFight) GetId() uint16  { return PCK_C2SInstanceMaterialFight_ID }

const PCK_S2CInstanceMaterialFight_ID = 608 //
//
type S2CInstanceMaterialFight struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	info *InstanceMaterialInfo `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
}

func (m *S2CInstanceMaterialFight) Reset()         { *m = S2CInstanceMaterialFight{} }
func (m *S2CInstanceMaterialFight) String() string { return proto.CompactTextString(m) }
func (*S2CInstanceMaterialFight) ProtoMessage()    {}
func (m *S2CInstanceMaterialFight) GetId() uint16  { return PCK_S2CInstanceMaterialFight_ID }

const PCK_C2SInstanceMaterialSweep_ID = 609 //
//
type C2SInstanceMaterialSweep struct {
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2SInstanceMaterialSweep) Reset()         { *m = C2SInstanceMaterialSweep{} }
func (m *C2SInstanceMaterialSweep) String() string { return proto.CompactTextString(m) }
func (*C2SInstanceMaterialSweep) ProtoMessage()    {}
func (m *C2SInstanceMaterialSweep) GetId() uint16  { return PCK_C2SInstanceMaterialSweep_ID }

const PCK_S2CInstanceMaterialSweep_ID = 610 //
//
type S2CInstanceMaterialSweep struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	info []*InstanceMaterialInfo `protobuf:"bytes,2,rep,name=info,proto3" json:"info,omitempty"`
}

func (m *S2CInstanceMaterialSweep) Reset()         { *m = S2CInstanceMaterialSweep{} }
func (m *S2CInstanceMaterialSweep) String() string { return proto.CompactTextString(m) }
func (*S2CInstanceMaterialSweep) ProtoMessage()    {}
func (m *S2CInstanceMaterialSweep) GetId() uint16  { return PCK_S2CInstanceMaterialSweep_ID }

const PCK_C2SInstanceMaterialBuy_ID = 628 //
//
type C2SInstanceMaterialBuy struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Num int32 `protobuf:"varint,2,opt,name=Num,proto3" json:"Num,omitempty"`
}

func (m *C2SInstanceMaterialBuy) Reset()         { *m = C2SInstanceMaterialBuy{} }
func (m *C2SInstanceMaterialBuy) String() string { return proto.CompactTextString(m) }
func (*C2SInstanceMaterialBuy) ProtoMessage()    {}
func (m *C2SInstanceMaterialBuy) GetId() uint16  { return PCK_C2SInstanceMaterialBuy_ID }

const PCK_S2CInstanceMaterialBuy_ID = 629 //
//
type S2CInstanceMaterialBuy struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	info []*InstanceMaterialInfo `protobuf:"bytes,2,rep,name=info,proto3" json:"info,omitempty"`
}

func (m *S2CInstanceMaterialBuy) Reset()         { *m = S2CInstanceMaterialBuy{} }
func (m *S2CInstanceMaterialBuy) String() string { return proto.CompactTextString(m) }
func (*S2CInstanceMaterialBuy) ProtoMessage()    {}
func (m *S2CInstanceMaterialBuy) GetId() uint16  { return PCK_S2CInstanceMaterialBuy_ID }

//
type InstanceBuBuShengLianData struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	ChooseNextId int32 `protobuf:"varint,2,opt,name=ChooseNextId,proto3" json:"ChooseNextId,omitempty"`
	//
	TeamData []*TeamMemberData `protobuf:"bytes,3,rep,name=TeamData,proto3" json:"TeamData,omitempty"`
	//
	FvUp int32 `protobuf:"varint,4,opt,name=FvUp,proto3" json:"FvUp,omitempty"`
	//
	passedId []int32 `protobuf:"varint,5,rep,name=passedId,proto3" json:"passedId,omitempty"`
}

func (m *InstanceBuBuShengLianData) Reset()         { *m = InstanceBuBuShengLianData{} }
func (m *InstanceBuBuShengLianData) String() string { return proto.CompactTextString(m) }
func (*InstanceBuBuShengLianData) ProtoMessage()    {}

const PCK_C2SInstanceBuBuShengLianInfo_ID = 930 //
//
type C2SInstanceBuBuShengLianInfo struct {
}

func (m *C2SInstanceBuBuShengLianInfo) Reset()         { *m = C2SInstanceBuBuShengLianInfo{} }
func (m *C2SInstanceBuBuShengLianInfo) String() string { return proto.CompactTextString(m) }
func (*C2SInstanceBuBuShengLianInfo) ProtoMessage()    {}
func (m *C2SInstanceBuBuShengLianInfo) GetId() uint16  { return PCK_C2SInstanceBuBuShengLianInfo_ID }

const PCK_S2CInstanceBuBuShengLianInfo_ID = 931 //
//
type S2CInstanceBuBuShengLianInfo struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	ChooseNextId int32 `protobuf:"varint,2,opt,name=ChooseNextId,proto3" json:"ChooseNextId,omitempty"`
	//
	NextTime int64 `protobuf:"varint,3,opt,name=NextTime,proto3" json:"NextTime,omitempty"`
	//
	TeamData []*TeamMemberData `protobuf:"bytes,4,rep,name=TeamData,proto3" json:"TeamData,omitempty"`
	//
	passedId []int32 `protobuf:"varint,5,rep,name=passedId,proto3" json:"passedId,omitempty"`
}

func (m *S2CInstanceBuBuShengLianInfo) Reset()         { *m = S2CInstanceBuBuShengLianInfo{} }
func (m *S2CInstanceBuBuShengLianInfo) String() string { return proto.CompactTextString(m) }
func (*S2CInstanceBuBuShengLianInfo) ProtoMessage()    {}
func (m *S2CInstanceBuBuShengLianInfo) GetId() uint16  { return PCK_S2CInstanceBuBuShengLianInfo_ID }

const PCK_C2SInstanceBuBuShengLianFight_ID = 932 //
//
type C2SInstanceBuBuShengLianFight struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2SInstanceBuBuShengLianFight) Reset()         { *m = C2SInstanceBuBuShengLianFight{} }
func (m *C2SInstanceBuBuShengLianFight) String() string { return proto.CompactTextString(m) }
func (*C2SInstanceBuBuShengLianFight) ProtoMessage()    {}
func (m *C2SInstanceBuBuShengLianFight) GetId() uint16  { return PCK_C2SInstanceBuBuShengLianFight_ID }

const PCK_S2CInstanceBuBuShengLianFight_ID = 933 //
//
type S2CInstanceBuBuShengLianFight struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	info *S2CInstanceBuBuShengLianInfo `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
}

func (m *S2CInstanceBuBuShengLianFight) Reset()         { *m = S2CInstanceBuBuShengLianFight{} }
func (m *S2CInstanceBuBuShengLianFight) String() string { return proto.CompactTextString(m) }
func (*S2CInstanceBuBuShengLianFight) ProtoMessage()    {}
func (m *S2CInstanceBuBuShengLianFight) GetId() uint16  { return PCK_S2CInstanceBuBuShengLianFight_ID }

const PCK_C2SInstanceBuBuShengLianJinLianZi_ID = 934 //
//
type C2SInstanceBuBuShengLianJinLianZi struct {
	//
	AutoBuy int32 `protobuf:"varint,1,opt,name=AutoBuy,proto3" json:"AutoBuy,omitempty"`
}

func (m *C2SInstanceBuBuShengLianJinLianZi) Reset()         { *m = C2SInstanceBuBuShengLianJinLianZi{} }
func (m *C2SInstanceBuBuShengLianJinLianZi) String() string { return proto.CompactTextString(m) }
func (*C2SInstanceBuBuShengLianJinLianZi) ProtoMessage()    {}
func (m *C2SInstanceBuBuShengLianJinLianZi) GetId() uint16 {
	return PCK_C2SInstanceBuBuShengLianJinLianZi_ID
}

const PCK_S2CInstanceBuBuShengLianJinLianZi_ID = 935 //
//
type S2CInstanceBuBuShengLianJinLianZi struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CInstanceBuBuShengLianJinLianZi) Reset()         { *m = S2CInstanceBuBuShengLianJinLianZi{} }
func (m *S2CInstanceBuBuShengLianJinLianZi) String() string { return proto.CompactTextString(m) }
func (*S2CInstanceBuBuShengLianJinLianZi) ProtoMessage()    {}
func (m *S2CInstanceBuBuShengLianJinLianZi) GetId() uint16 {
	return PCK_S2CInstanceBuBuShengLianJinLianZi_ID
}

//
type TeamMemberData struct {
	//
	GradeType int32 `protobuf:"varint,1,opt,name=GradeType,proto3" json:"GradeType,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Hp int64 `protobuf:"varint,3,opt,name=Hp,proto3" json:"Hp,omitempty"`
	//
	MaxHp int64 `protobuf:"varint,4,opt,name=MaxHp,proto3" json:"MaxHp,omitempty"`
	//
	Pos int32 `protobuf:"varint,5,opt,name=Pos,proto3" json:"Pos,omitempty"`
}

func (m *TeamMemberData) Reset()         { *m = TeamMemberData{} }
func (m *TeamMemberData) String() string { return proto.CompactTextString(m) }
func (*TeamMemberData) ProtoMessage()    {}

//
type InstanceTreasureInfo struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	S int64 `protobuf:"varint,2,opt,name=S,proto3" json:"S,omitempty"`
	//
	L int32 `protobuf:"varint,3,opt,name=L,proto3" json:"L,omitempty"`
}

func (m *InstanceTreasureInfo) Reset()         { *m = InstanceTreasureInfo{} }
func (m *InstanceTreasureInfo) String() string { return proto.CompactTextString(m) }
func (*InstanceTreasureInfo) ProtoMessage()    {}

const PCK_S2CInstanceTreasureInfo_ID = 611 //
//
type S2CInstanceTreasureInfo struct {
	//
	Items []*InstanceTreasureInfo `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty"`
	//
	B []int32 `protobuf:"varint,2,rep,name=B,proto3" json:"B,omitempty"`
}

func (m *S2CInstanceTreasureInfo) Reset()         { *m = S2CInstanceTreasureInfo{} }
func (m *S2CInstanceTreasureInfo) String() string { return proto.CompactTextString(m) }
func (*S2CInstanceTreasureInfo) ProtoMessage()    {}
func (m *S2CInstanceTreasureInfo) GetId() uint16  { return PCK_S2CInstanceTreasureInfo_ID }

const PCK_C2SInstanceTreasureFight_ID = 612 //
//
type C2SInstanceTreasureFight struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2SInstanceTreasureFight) Reset()         { *m = C2SInstanceTreasureFight{} }
func (m *C2SInstanceTreasureFight) String() string { return proto.CompactTextString(m) }
func (*C2SInstanceTreasureFight) ProtoMessage()    {}
func (m *C2SInstanceTreasureFight) GetId() uint16  { return PCK_C2SInstanceTreasureFight_ID }

const PCK_S2CInstanceTreasureFight_ID = 613 //
//
type S2CInstanceTreasureFight struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Star int32 `protobuf:"varint,2,opt,name=Star,proto3" json:"Star,omitempty"`
	//
	Id int32 `protobuf:"varint,3,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *S2CInstanceTreasureFight) Reset()         { *m = S2CInstanceTreasureFight{} }
func (m *S2CInstanceTreasureFight) String() string { return proto.CompactTextString(m) }
func (*S2CInstanceTreasureFight) ProtoMessage()    {}
func (m *S2CInstanceTreasureFight) GetId() uint16  { return PCK_S2CInstanceTreasureFight_ID }

const PCK_C2SInstanceTreasureSweep_ID = 614 //
//
type C2SInstanceTreasureSweep struct {
}

func (m *C2SInstanceTreasureSweep) Reset()         { *m = C2SInstanceTreasureSweep{} }
func (m *C2SInstanceTreasureSweep) String() string { return proto.CompactTextString(m) }
func (*C2SInstanceTreasureSweep) ProtoMessage()    {}
func (m *C2SInstanceTreasureSweep) GetId() uint16  { return PCK_C2SInstanceTreasureSweep_ID }

const PCK_S2CInstanceTreasureSweep_ID = 615 //
//
type S2CInstanceTreasureSweep struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Ids []int32 `protobuf:"varint,2,rep,name=Ids,proto3" json:"Ids,omitempty"`
}

func (m *S2CInstanceTreasureSweep) Reset()         { *m = S2CInstanceTreasureSweep{} }
func (m *S2CInstanceTreasureSweep) String() string { return proto.CompactTextString(m) }
func (*S2CInstanceTreasureSweep) ProtoMessage()    {}
func (m *S2CInstanceTreasureSweep) GetId() uint16  { return PCK_S2CInstanceTreasureSweep_ID }

const PCK_C2SGetInstanceTreasureBox_ID = 616 //
//
type C2SGetInstanceTreasureBox struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2SGetInstanceTreasureBox) Reset()         { *m = C2SGetInstanceTreasureBox{} }
func (m *C2SGetInstanceTreasureBox) String() string { return proto.CompactTextString(m) }
func (*C2SGetInstanceTreasureBox) ProtoMessage()    {}
func (m *C2SGetInstanceTreasureBox) GetId() uint16  { return PCK_C2SGetInstanceTreasureBox_ID }

const PCK_S2CGetInstanceTreasureBox_ID = 617 //
//
type S2CGetInstanceTreasureBox struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Ids []int32 `protobuf:"varint,3,rep,name=Ids,proto3" json:"Ids,omitempty"`
}

func (m *S2CGetInstanceTreasureBox) Reset()         { *m = S2CGetInstanceTreasureBox{} }
func (m *S2CGetInstanceTreasureBox) String() string { return proto.CompactTextString(m) }
func (*S2CGetInstanceTreasureBox) ProtoMessage()    {}
func (m *S2CGetInstanceTreasureBox) GetId() uint16  { return PCK_S2CGetInstanceTreasureBox_ID }

const PCK_S2CInstanceHeavenlyInfo_ID = 618 //
//
type S2CInstanceHeavenlyInfo struct {
	//
	M int32 `protobuf:"varint,1,opt,name=M,proto3" json:"M,omitempty"`
	//
	D int32 `protobuf:"varint,2,opt,name=D,proto3" json:"D,omitempty"`
	//
	B []int32 `protobuf:"varint,3,rep,name=B,proto3" json:"B,omitempty"`
}

func (m *S2CInstanceHeavenlyInfo) Reset()         { *m = S2CInstanceHeavenlyInfo{} }
func (m *S2CInstanceHeavenlyInfo) String() string { return proto.CompactTextString(m) }
func (*S2CInstanceHeavenlyInfo) ProtoMessage()    {}
func (m *S2CInstanceHeavenlyInfo) GetId() uint16  { return PCK_S2CInstanceHeavenlyInfo_ID }

const PCK_C2SInstanceHeavenlyFight_ID = 619 //
//
type C2SInstanceHeavenlyFight struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2SInstanceHeavenlyFight) Reset()         { *m = C2SInstanceHeavenlyFight{} }
func (m *C2SInstanceHeavenlyFight) String() string { return proto.CompactTextString(m) }
func (*C2SInstanceHeavenlyFight) ProtoMessage()    {}
func (m *C2SInstanceHeavenlyFight) GetId() uint16  { return PCK_C2SInstanceHeavenlyFight_ID }

const PCK_S2CInstanceHeavenlyFight_ID = 620 //
//
type S2CInstanceHeavenlyFight struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Win int32 `protobuf:"varint,3,opt,name=Win,proto3" json:"Win,omitempty"`
}

func (m *S2CInstanceHeavenlyFight) Reset()         { *m = S2CInstanceHeavenlyFight{} }
func (m *S2CInstanceHeavenlyFight) String() string { return proto.CompactTextString(m) }
func (*S2CInstanceHeavenlyFight) ProtoMessage()    {}
func (m *S2CInstanceHeavenlyFight) GetId() uint16  { return PCK_S2CInstanceHeavenlyFight_ID }

const PCK_C2SInstanceHeavenlySweep_ID = 621 //
//
type C2SInstanceHeavenlySweep struct {
}

func (m *C2SInstanceHeavenlySweep) Reset()         { *m = C2SInstanceHeavenlySweep{} }
func (m *C2SInstanceHeavenlySweep) String() string { return proto.CompactTextString(m) }
func (*C2SInstanceHeavenlySweep) ProtoMessage()    {}
func (m *C2SInstanceHeavenlySweep) GetId() uint16  { return PCK_C2SInstanceHeavenlySweep_ID }

const PCK_S2CInstanceHeavenlySweep_ID = 622 //
//
type S2CInstanceHeavenlySweep struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	D int32 `protobuf:"varint,2,opt,name=D,proto3" json:"D,omitempty"`
}

func (m *S2CInstanceHeavenlySweep) Reset()         { *m = S2CInstanceHeavenlySweep{} }
func (m *S2CInstanceHeavenlySweep) String() string { return proto.CompactTextString(m) }
func (*S2CInstanceHeavenlySweep) ProtoMessage()    {}
func (m *S2CInstanceHeavenlySweep) GetId() uint16  { return PCK_S2CInstanceHeavenlySweep_ID }

const PCK_C2SGetInstanceHeavenlyBox_ID = 623 //
//
type C2SGetInstanceHeavenlyBox struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2SGetInstanceHeavenlyBox) Reset()         { *m = C2SGetInstanceHeavenlyBox{} }
func (m *C2SGetInstanceHeavenlyBox) String() string { return proto.CompactTextString(m) }
func (*C2SGetInstanceHeavenlyBox) ProtoMessage()    {}
func (m *C2SGetInstanceHeavenlyBox) GetId() uint16  { return PCK_C2SGetInstanceHeavenlyBox_ID }

const PCK_S2CGetInstanceHeavenlyBox_ID = 624 //
//
type S2CGetInstanceHeavenlyBox struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id int32 `protobuf:"varint,3,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *S2CGetInstanceHeavenlyBox) Reset()         { *m = S2CGetInstanceHeavenlyBox{} }
func (m *S2CGetInstanceHeavenlyBox) String() string { return proto.CompactTextString(m) }
func (*S2CGetInstanceHeavenlyBox) ProtoMessage()    {}
func (m *S2CGetInstanceHeavenlyBox) GetId() uint16  { return PCK_S2CGetInstanceHeavenlyBox_ID }

const PCK_S2CInstanceDemonInfo_ID = 684 //
//
type S2CInstanceDemonInfo struct {
	//
	M int32 `protobuf:"varint,1,opt,name=M,proto3" json:"M,omitempty"`
	//
	D int32 `protobuf:"varint,2,opt,name=D,proto3" json:"D,omitempty"`
	//
	B []int32 `protobuf:"varint,3,rep,name=B,proto3" json:"B,omitempty"`
}

func (m *S2CInstanceDemonInfo) Reset()         { *m = S2CInstanceDemonInfo{} }
func (m *S2CInstanceDemonInfo) String() string { return proto.CompactTextString(m) }
func (*S2CInstanceDemonInfo) ProtoMessage()    {}
func (m *S2CInstanceDemonInfo) GetId() uint16  { return PCK_S2CInstanceDemonInfo_ID }

const PCK_C2SInstanceDemonFight_ID = 685 //
//
type C2SInstanceDemonFight struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2SInstanceDemonFight) Reset()         { *m = C2SInstanceDemonFight{} }
func (m *C2SInstanceDemonFight) String() string { return proto.CompactTextString(m) }
func (*C2SInstanceDemonFight) ProtoMessage()    {}
func (m *C2SInstanceDemonFight) GetId() uint16  { return PCK_C2SInstanceDemonFight_ID }

const PCK_S2CInstanceDemonFight_ID = 686 //
//
type S2CInstanceDemonFight struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Win int32 `protobuf:"varint,3,opt,name=Win,proto3" json:"Win,omitempty"`
}

func (m *S2CInstanceDemonFight) Reset()         { *m = S2CInstanceDemonFight{} }
func (m *S2CInstanceDemonFight) String() string { return proto.CompactTextString(m) }
func (*S2CInstanceDemonFight) ProtoMessage()    {}
func (m *S2CInstanceDemonFight) GetId() uint16  { return PCK_S2CInstanceDemonFight_ID }

const PCK_C2SInstanceDemonSweep_ID = 687 //
//
type C2SInstanceDemonSweep struct {
}

func (m *C2SInstanceDemonSweep) Reset()         { *m = C2SInstanceDemonSweep{} }
func (m *C2SInstanceDemonSweep) String() string { return proto.CompactTextString(m) }
func (*C2SInstanceDemonSweep) ProtoMessage()    {}
func (m *C2SInstanceDemonSweep) GetId() uint16  { return PCK_C2SInstanceDemonSweep_ID }

const PCK_S2CInstanceDemonSweep_ID = 688 //
//
type S2CInstanceDemonSweep struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	D int32 `protobuf:"varint,2,opt,name=D,proto3" json:"D,omitempty"`
}

func (m *S2CInstanceDemonSweep) Reset()         { *m = S2CInstanceDemonSweep{} }
func (m *S2CInstanceDemonSweep) String() string { return proto.CompactTextString(m) }
func (*S2CInstanceDemonSweep) ProtoMessage()    {}
func (m *S2CInstanceDemonSweep) GetId() uint16  { return PCK_S2CInstanceDemonSweep_ID }

const PCK_C2SGetInstanceDemonBox_ID = 689 //
//
type C2SGetInstanceDemonBox struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2SGetInstanceDemonBox) Reset()         { *m = C2SGetInstanceDemonBox{} }
func (m *C2SGetInstanceDemonBox) String() string { return proto.CompactTextString(m) }
func (*C2SGetInstanceDemonBox) ProtoMessage()    {}
func (m *C2SGetInstanceDemonBox) GetId() uint16  { return PCK_C2SGetInstanceDemonBox_ID }

const PCK_S2CGetInstanceDemonBox_ID = 690 //
//
type S2CGetInstanceDemonBox struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id int32 `protobuf:"varint,3,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *S2CGetInstanceDemonBox) Reset()         { *m = S2CGetInstanceDemonBox{} }
func (m *S2CGetInstanceDemonBox) String() string { return proto.CompactTextString(m) }
func (*S2CGetInstanceDemonBox) ProtoMessage()    {}
func (m *S2CGetInstanceDemonBox) GetId() uint16  { return PCK_S2CGetInstanceDemonBox_ID }

const PCK_S2CInstanceTowerInfo_ID = 625 //
//
type S2CInstanceTowerInfo struct {
	//
	M int32 `protobuf:"varint,1,opt,name=M,proto3" json:"M,omitempty"`
}

func (m *S2CInstanceTowerInfo) Reset()         { *m = S2CInstanceTowerInfo{} }
func (m *S2CInstanceTowerInfo) String() string { return proto.CompactTextString(m) }
func (*S2CInstanceTowerInfo) ProtoMessage()    {}
func (m *S2CInstanceTowerInfo) GetId() uint16  { return PCK_S2CInstanceTowerInfo_ID }

const PCK_C2SInstanceTowerFight_ID = 626 //
//
type C2SInstanceTowerFight struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2SInstanceTowerFight) Reset()         { *m = C2SInstanceTowerFight{} }
func (m *C2SInstanceTowerFight) String() string { return proto.CompactTextString(m) }
func (*C2SInstanceTowerFight) ProtoMessage()    {}
func (m *C2SInstanceTowerFight) GetId() uint16  { return PCK_C2SInstanceTowerFight_ID }

const PCK_S2CInstanceTowerFight_ID = 627 //
//
type S2CInstanceTowerFight struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Win int32 `protobuf:"varint,3,opt,name=Win,proto3" json:"Win,omitempty"`
}

func (m *S2CInstanceTowerFight) Reset()         { *m = S2CInstanceTowerFight{} }
func (m *S2CInstanceTowerFight) String() string { return proto.CompactTextString(m) }
func (*S2CInstanceTowerFight) ProtoMessage()    {}
func (m *S2CInstanceTowerFight) GetId() uint16  { return PCK_S2CInstanceTowerFight_ID }

//
type AllBossInfo struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	CurrHp int64 `protobuf:"varint,2,opt,name=CurrHp,proto3" json:"CurrHp,omitempty"`
	//
	MaxHP int64 `protobuf:"varint,3,opt,name=MaxHP,proto3" json:"MaxHP,omitempty"`
	//
	Num int32 `protobuf:"varint,4,opt,name=Num,proto3" json:"Num,omitempty"`
	//
	Relive int64 `protobuf:"varint,5,opt,name=Relive,proto3" json:"Relive,omitempty"`
}

func (m *AllBossInfo) Reset()         { *m = AllBossInfo{} }
func (m *AllBossInfo) String() string { return proto.CompactTextString(m) }
func (*AllBossInfo) ProtoMessage()    {}

const PCK_S2CAllBossInfo_ID = 631 //
//
type S2CAllBossInfo struct {
	//
	Items []*AllBossInfo `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty"`
}

func (m *S2CAllBossInfo) Reset()         { *m = S2CAllBossInfo{} }
func (m *S2CAllBossInfo) String() string { return proto.CompactTextString(m) }
func (*S2CAllBossInfo) ProtoMessage()    {}
func (m *S2CAllBossInfo) GetId() uint16  { return PCK_S2CAllBossInfo_ID }

const PCK_C2SAllBossStart_ID = 632 //
//
type C2SAllBossStart struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2SAllBossStart) Reset()         { *m = C2SAllBossStart{} }
func (m *C2SAllBossStart) String() string { return proto.CompactTextString(m) }
func (*C2SAllBossStart) ProtoMessage()    {}
func (m *C2SAllBossStart) GetId() uint16  { return PCK_C2SAllBossStart_ID }

const PCK_S2CAllBossStart_ID = 633 //
//
type S2CAllBossStart struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CAllBossStart) Reset()         { *m = S2CAllBossStart{} }
func (m *S2CAllBossStart) String() string { return proto.CompactTextString(m) }
func (*S2CAllBossStart) ProtoMessage()    {}
func (m *S2CAllBossStart) GetId() uint16  { return PCK_S2CAllBossStart_ID }

const PCK_C2SAllBossCfg_ID = 634 //
//
type C2SAllBossCfg struct {
	//
	Cfg string `protobuf:"bytes,1,opt,name=Cfg,proto3" json:"Cfg,omitempty"`
}

func (m *C2SAllBossCfg) Reset()         { *m = C2SAllBossCfg{} }
func (m *C2SAllBossCfg) String() string { return proto.CompactTextString(m) }
func (*C2SAllBossCfg) ProtoMessage()    {}
func (m *C2SAllBossCfg) GetId() uint16  { return PCK_C2SAllBossCfg_ID }

const PCK_C2SAllBossRelive_ID = 635 //
//
type C2SAllBossRelive struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2SAllBossRelive) Reset()         { *m = C2SAllBossRelive{} }
func (m *C2SAllBossRelive) String() string { return proto.CompactTextString(m) }
func (*C2SAllBossRelive) ProtoMessage()    {}
func (m *C2SAllBossRelive) GetId() uint16  { return PCK_C2SAllBossRelive_ID }

const PCK_S2CAllBossRelive_ID = 636 //
//
type S2CAllBossRelive struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CAllBossRelive) Reset()         { *m = S2CAllBossRelive{} }
func (m *S2CAllBossRelive) String() string { return proto.CompactTextString(m) }
func (*S2CAllBossRelive) ProtoMessage()    {}
func (m *S2CAllBossRelive) GetId() uint16  { return PCK_S2CAllBossRelive_ID }

const PCK_C2SAllBossGetBuyInfo_ID = 637 //
//
type C2SAllBossGetBuyInfo struct {
}

func (m *C2SAllBossGetBuyInfo) Reset()         { *m = C2SAllBossGetBuyInfo{} }
func (m *C2SAllBossGetBuyInfo) String() string { return proto.CompactTextString(m) }
func (*C2SAllBossGetBuyInfo) ProtoMessage()    {}
func (m *C2SAllBossGetBuyInfo) GetId() uint16  { return PCK_C2SAllBossGetBuyInfo_ID }

const PCK_S2CAllBossGetBuyInfo_ID = 638 //
//
type S2CAllBossGetBuyInfo struct {
	//
	Coin3 int32 `protobuf:"varint,1,opt,name=Coin3,proto3" json:"Coin3,omitempty"`
	//
	Times int32 `protobuf:"varint,2,opt,name=Times,proto3" json:"Times,omitempty"`
	//
	LeftTimes int32 `protobuf:"varint,3,opt,name=LeftTimes,proto3" json:"LeftTimes,omitempty"`
	//
	NoticeTimes int32 `protobuf:"varint,4,opt,name=NoticeTimes,proto3" json:"NoticeTimes,omitempty"`
	//
	NoticeVip int32 `protobuf:"varint,5,opt,name=NoticeVip,proto3" json:"NoticeVip,omitempty"`
}

func (m *S2CAllBossGetBuyInfo) Reset()         { *m = S2CAllBossGetBuyInfo{} }
func (m *S2CAllBossGetBuyInfo) String() string { return proto.CompactTextString(m) }
func (*S2CAllBossGetBuyInfo) ProtoMessage()    {}
func (m *S2CAllBossGetBuyInfo) GetId() uint16  { return PCK_S2CAllBossGetBuyInfo_ID }

const PCK_C2SAllBossBuyTimes_ID = 639 //
//
type C2SAllBossBuyTimes struct {
}

func (m *C2SAllBossBuyTimes) Reset()         { *m = C2SAllBossBuyTimes{} }
func (m *C2SAllBossBuyTimes) String() string { return proto.CompactTextString(m) }
func (*C2SAllBossBuyTimes) ProtoMessage()    {}
func (m *C2SAllBossBuyTimes) GetId() uint16  { return PCK_C2SAllBossBuyTimes_ID }

const PCK_S2CAllBossBuyTimes_ID = 640 //
//
type S2CAllBossBuyTimes struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CAllBossBuyTimes) Reset()         { *m = S2CAllBossBuyTimes{} }
func (m *S2CAllBossBuyTimes) String() string { return proto.CompactTextString(m) }
func (*S2CAllBossBuyTimes) ProtoMessage()    {}
func (m *S2CAllBossBuyTimes) GetId() uint16  { return PCK_S2CAllBossBuyTimes_ID }

//
type DamageLog struct {
	//
	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	Nick string `protobuf:"bytes,2,opt,name=Nick,proto3" json:"Nick,omitempty"`
	//
	Damage int64 `protobuf:"varint,3,opt,name=Damage,proto3" json:"Damage,omitempty"`
}

func (m *DamageLog) Reset()         { *m = DamageLog{} }
func (m *DamageLog) String() string { return proto.CompactTextString(m) }
func (*DamageLog) ProtoMessage()    {}

const PCK_C2SAllBossGetDamageLog_ID = 641 //
//
type C2SAllBossGetDamageLog struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2SAllBossGetDamageLog) Reset()         { *m = C2SAllBossGetDamageLog{} }
func (m *C2SAllBossGetDamageLog) String() string { return proto.CompactTextString(m) }
func (*C2SAllBossGetDamageLog) ProtoMessage()    {}
func (m *C2SAllBossGetDamageLog) GetId() uint16  { return PCK_C2SAllBossGetDamageLog_ID }

const PCK_S2CAllBossGetDamageLog_ID = 642 //
//
type S2CAllBossGetDamageLog struct {
	//
	Items []*DamageLog `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty"`
}

func (m *S2CAllBossGetDamageLog) Reset()         { *m = S2CAllBossGetDamageLog{} }
func (m *S2CAllBossGetDamageLog) String() string { return proto.CompactTextString(m) }
func (*S2CAllBossGetDamageLog) ProtoMessage()    {}
func (m *S2CAllBossGetDamageLog) GetId() uint16  { return PCK_S2CAllBossGetDamageLog_ID }

//
type KillLog struct {
	//
	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	Nick string `protobuf:"bytes,2,opt,name=Nick,proto3" json:"Nick,omitempty"`
	//
	FightValue int64 `protobuf:"varint,3,opt,name=FightValue,proto3" json:"FightValue,omitempty"`
	//
	KillTime int64 `protobuf:"varint,4,opt,name=KillTime,proto3" json:"KillTime,omitempty"`
}

func (m *KillLog) Reset()         { *m = KillLog{} }
func (m *KillLog) String() string { return proto.CompactTextString(m) }
func (*KillLog) ProtoMessage()    {}

const PCK_C2SAllBossGetKillLog_ID = 643 //
//
type C2SAllBossGetKillLog struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2SAllBossGetKillLog) Reset()         { *m = C2SAllBossGetKillLog{} }
func (m *C2SAllBossGetKillLog) String() string { return proto.CompactTextString(m) }
func (*C2SAllBossGetKillLog) ProtoMessage()    {}
func (m *C2SAllBossGetKillLog) GetId() uint16  { return PCK_C2SAllBossGetKillLog_ID }

const PCK_S2CAllBossGetKillLog_ID = 644 //
//
type S2CAllBossGetKillLog struct {
	//
	Items []*KillLog `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty"`
}

func (m *S2CAllBossGetKillLog) Reset()         { *m = S2CAllBossGetKillLog{} }
func (m *S2CAllBossGetKillLog) String() string { return proto.CompactTextString(m) }
func (*S2CAllBossGetKillLog) ProtoMessage()    {}
func (m *S2CAllBossGetKillLog) GetId() uint16  { return PCK_S2CAllBossGetKillLog_ID }

//
type AllBossV2Info struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	CurrHp int64 `protobuf:"varint,2,opt,name=CurrHp,proto3" json:"CurrHp,omitempty"`
	//
	MaxHP int64 `protobuf:"varint,3,opt,name=MaxHP,proto3" json:"MaxHP,omitempty"`
	//
	Num int32 `protobuf:"varint,4,opt,name=Num,proto3" json:"Num,omitempty"`
	//
	Relive int64 `protobuf:"varint,5,opt,name=Relive,proto3" json:"Relive,omitempty"`
	//
	RunAway int64 `protobuf:"varint,6,opt,name=RunAway,proto3" json:"RunAway,omitempty"`
	//
	State int32 `protobuf:"varint,7,opt,name=State,proto3" json:"State,omitempty"`
}

func (m *AllBossV2Info) Reset()         { *m = AllBossV2Info{} }
func (m *AllBossV2Info) String() string { return proto.CompactTextString(m) }
func (*AllBossV2Info) ProtoMessage()    {}

const PCK_S2CAllBossV2Info_ID = 1111 //
//
type S2CAllBossV2Info struct {
	//
	Items []*AllBossV2Info `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty"`
}

func (m *S2CAllBossV2Info) Reset()         { *m = S2CAllBossV2Info{} }
func (m *S2CAllBossV2Info) String() string { return proto.CompactTextString(m) }
func (*S2CAllBossV2Info) ProtoMessage()    {}
func (m *S2CAllBossV2Info) GetId() uint16  { return PCK_S2CAllBossV2Info_ID }

const PCK_C2SAllBossV2PlayerInBoss_ID = 1125 //
//
type C2SAllBossV2PlayerInBoss struct {
}

func (m *C2SAllBossV2PlayerInBoss) Reset()         { *m = C2SAllBossV2PlayerInBoss{} }
func (m *C2SAllBossV2PlayerInBoss) String() string { return proto.CompactTextString(m) }
func (*C2SAllBossV2PlayerInBoss) ProtoMessage()    {}
func (m *C2SAllBossV2PlayerInBoss) GetId() uint16  { return PCK_C2SAllBossV2PlayerInBoss_ID }

const PCK_S2CAllBossV2PlayerInBoss_ID = 1126 //
//
type S2CAllBossV2PlayerInBoss struct {
	//
	BossId int32 `protobuf:"varint,1,opt,name=BossId,proto3" json:"BossId,omitempty"`
	//
	DamageOrder int32 `protobuf:"varint,2,opt,name=DamageOrder,proto3" json:"DamageOrder,omitempty"`
	//
	Damage int64 `protobuf:"varint,3,opt,name=Damage,proto3" json:"Damage,omitempty"`
}

func (m *S2CAllBossV2PlayerInBoss) Reset()         { *m = S2CAllBossV2PlayerInBoss{} }
func (m *S2CAllBossV2PlayerInBoss) String() string { return proto.CompactTextString(m) }
func (*S2CAllBossV2PlayerInBoss) ProtoMessage()    {}
func (m *S2CAllBossV2PlayerInBoss) GetId() uint16  { return PCK_S2CAllBossV2PlayerInBoss_ID }

const PCK_C2SAllBossV2Start_ID = 1112 //
//
type C2SAllBossV2Start struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2SAllBossV2Start) Reset()         { *m = C2SAllBossV2Start{} }
func (m *C2SAllBossV2Start) String() string { return proto.CompactTextString(m) }
func (*C2SAllBossV2Start) ProtoMessage()    {}
func (m *C2SAllBossV2Start) GetId() uint16  { return PCK_C2SAllBossV2Start_ID }

const PCK_S2CAllBossV2Start_ID = 1113 //
//
type S2CAllBossV2Start struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CAllBossV2Start) Reset()         { *m = S2CAllBossV2Start{} }
func (m *S2CAllBossV2Start) String() string { return proto.CompactTextString(m) }
func (*S2CAllBossV2Start) ProtoMessage()    {}
func (m *S2CAllBossV2Start) GetId() uint16  { return PCK_S2CAllBossV2Start_ID }

const PCK_C2SAllBossV2Cfg_ID = 1114 //
//
type C2SAllBossV2Cfg struct {
	//
	Cfg string `protobuf:"bytes,1,opt,name=Cfg,proto3" json:"Cfg,omitempty"`
}

func (m *C2SAllBossV2Cfg) Reset()         { *m = C2SAllBossV2Cfg{} }
func (m *C2SAllBossV2Cfg) String() string { return proto.CompactTextString(m) }
func (*C2SAllBossV2Cfg) ProtoMessage()    {}
func (m *C2SAllBossV2Cfg) GetId() uint16  { return PCK_C2SAllBossV2Cfg_ID }

const PCK_C2SAllBossV2GetBuyInfo_ID = 1117 //
//
type C2SAllBossV2GetBuyInfo struct {
}

func (m *C2SAllBossV2GetBuyInfo) Reset()         { *m = C2SAllBossV2GetBuyInfo{} }
func (m *C2SAllBossV2GetBuyInfo) String() string { return proto.CompactTextString(m) }
func (*C2SAllBossV2GetBuyInfo) ProtoMessage()    {}
func (m *C2SAllBossV2GetBuyInfo) GetId() uint16  { return PCK_C2SAllBossV2GetBuyInfo_ID }

const PCK_S2CAllBossV2GetBuyInfo_ID = 1118 //
//
type S2CAllBossV2GetBuyInfo struct {
	//
	Coin3 int32 `protobuf:"varint,1,opt,name=Coin3,proto3" json:"Coin3,omitempty"`
	//
	Times int32 `protobuf:"varint,2,opt,name=Times,proto3" json:"Times,omitempty"`
	//
	LeftTimes int32 `protobuf:"varint,3,opt,name=LeftTimes,proto3" json:"LeftTimes,omitempty"`
	//
	NoticeTimes int32 `protobuf:"varint,4,opt,name=NoticeTimes,proto3" json:"NoticeTimes,omitempty"`
	//
	NoticeVip int32 `protobuf:"varint,5,opt,name=NoticeVip,proto3" json:"NoticeVip,omitempty"`
}

func (m *S2CAllBossV2GetBuyInfo) Reset()         { *m = S2CAllBossV2GetBuyInfo{} }
func (m *S2CAllBossV2GetBuyInfo) String() string { return proto.CompactTextString(m) }
func (*S2CAllBossV2GetBuyInfo) ProtoMessage()    {}
func (m *S2CAllBossV2GetBuyInfo) GetId() uint16  { return PCK_S2CAllBossV2GetBuyInfo_ID }

const PCK_C2SAllBossV2BuyTimes_ID = 1119 //
//
type C2SAllBossV2BuyTimes struct {
}

func (m *C2SAllBossV2BuyTimes) Reset()         { *m = C2SAllBossV2BuyTimes{} }
func (m *C2SAllBossV2BuyTimes) String() string { return proto.CompactTextString(m) }
func (*C2SAllBossV2BuyTimes) ProtoMessage()    {}
func (m *C2SAllBossV2BuyTimes) GetId() uint16  { return PCK_C2SAllBossV2BuyTimes_ID }

const PCK_S2CAllBossV2BuyTimes_ID = 1120 //
//
type S2CAllBossV2BuyTimes struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CAllBossV2BuyTimes) Reset()         { *m = S2CAllBossV2BuyTimes{} }
func (m *S2CAllBossV2BuyTimes) String() string { return proto.CompactTextString(m) }
func (*S2CAllBossV2BuyTimes) ProtoMessage()    {}
func (m *S2CAllBossV2BuyTimes) GetId() uint16  { return PCK_S2CAllBossV2BuyTimes_ID }

//
type DamageRank struct {
	//
	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	DamageTime int64 `protobuf:"varint,2,opt,name=DamageTime,proto3" json:"DamageTime,omitempty"`
	//
	Nick string `protobuf:"bytes,3,opt,name=Nick,proto3" json:"Nick,omitempty"`
	//
	Damage int64 `protobuf:"varint,4,opt,name=Damage,proto3" json:"Damage,omitempty"`
	//
	AreaId int64 `protobuf:"varint,5,opt,name=AreaId,proto3" json:"AreaId,omitempty"`
}

func (m *DamageRank) Reset()         { *m = DamageRank{} }
func (m *DamageRank) String() string { return proto.CompactTextString(m) }
func (*DamageRank) ProtoMessage()    {}

const PCK_C2SAllBossV2GetDamageLog_ID = 1121 //
//
type C2SAllBossV2GetDamageLog struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2SAllBossV2GetDamageLog) Reset()         { *m = C2SAllBossV2GetDamageLog{} }
func (m *C2SAllBossV2GetDamageLog) String() string { return proto.CompactTextString(m) }
func (*C2SAllBossV2GetDamageLog) ProtoMessage()    {}
func (m *C2SAllBossV2GetDamageLog) GetId() uint16  { return PCK_C2SAllBossV2GetDamageLog_ID }

const PCK_S2CAllBossV2GetDamageLog_ID = 1122 //
//
type S2CAllBossV2GetDamageLog struct {
	//
	BossId int32 `protobuf:"varint,1,opt,name=BossId,proto3" json:"BossId,omitempty"`
	//
	BossState int32 `protobuf:"varint,2,opt,name=BossState,proto3" json:"BossState,omitempty"`
	//
	Items []*DamageRank `protobuf:"bytes,3,rep,name=Items,proto3" json:"Items,omitempty"`
}

func (m *S2CAllBossV2GetDamageLog) Reset()         { *m = S2CAllBossV2GetDamageLog{} }
func (m *S2CAllBossV2GetDamageLog) String() string { return proto.CompactTextString(m) }
func (*S2CAllBossV2GetDamageLog) ProtoMessage()    {}
func (m *S2CAllBossV2GetDamageLog) GetId() uint16  { return PCK_S2CAllBossV2GetDamageLog_ID }

const PCK_K2SFightAllBossResult_ID = 1127 //
//
type K2SFightAllBossResult struct {
	//
	BossId int32 `protobuf:"varint,1,opt,name=BossId,proto3" json:"BossId,omitempty"`
	//
	IsFirst int32 `protobuf:"varint,2,opt,name=IsFirst,proto3" json:"IsFirst,omitempty"`
	//
	RunAwayTime int64 `protobuf:"varint,3,opt,name=RunAwayTime,proto3" json:"RunAwayTime,omitempty"`
}

func (m *K2SFightAllBossResult) Reset()         { *m = K2SFightAllBossResult{} }
func (m *K2SFightAllBossResult) String() string { return proto.CompactTextString(m) }
func (*K2SFightAllBossResult) ProtoMessage()    {}
func (m *K2SFightAllBossResult) GetId() uint16  { return PCK_K2SFightAllBossResult_ID }

//
type FieldBossV2Info struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	CurrHp int64 `protobuf:"varint,2,opt,name=CurrHp,proto3" json:"CurrHp,omitempty"`
	//
	MaxHP int64 `protobuf:"varint,3,opt,name=MaxHP,proto3" json:"MaxHP,omitempty"`
	//
	Num int32 `protobuf:"varint,4,opt,name=Num,proto3" json:"Num,omitempty"`
	//
	Relive int64 `protobuf:"varint,5,opt,name=Relive,proto3" json:"Relive,omitempty"`
	//
	RunAway int64 `protobuf:"varint,6,opt,name=RunAway,proto3" json:"RunAway,omitempty"`
	//
	State int32 `protobuf:"varint,7,opt,name=State,proto3" json:"State,omitempty"`
}

func (m *FieldBossV2Info) Reset()         { *m = FieldBossV2Info{} }
func (m *FieldBossV2Info) String() string { return proto.CompactTextString(m) }
func (*FieldBossV2Info) ProtoMessage()    {}

const PCK_S2CFieldBossV2Info_ID = 1131 //
//
type S2CFieldBossV2Info struct {
	//
	Items []*FieldBossV2Info `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty"`
}

func (m *S2CFieldBossV2Info) Reset()         { *m = S2CFieldBossV2Info{} }
func (m *S2CFieldBossV2Info) String() string { return proto.CompactTextString(m) }
func (*S2CFieldBossV2Info) ProtoMessage()    {}
func (m *S2CFieldBossV2Info) GetId() uint16  { return PCK_S2CFieldBossV2Info_ID }

const PCK_C2SFieldBossV2PlayerInBoss_ID = 1141 //
//
type C2SFieldBossV2PlayerInBoss struct {
}

func (m *C2SFieldBossV2PlayerInBoss) Reset()         { *m = C2SFieldBossV2PlayerInBoss{} }
func (m *C2SFieldBossV2PlayerInBoss) String() string { return proto.CompactTextString(m) }
func (*C2SFieldBossV2PlayerInBoss) ProtoMessage()    {}
func (m *C2SFieldBossV2PlayerInBoss) GetId() uint16  { return PCK_C2SFieldBossV2PlayerInBoss_ID }

const PCK_S2CFieldBossV2PlayerInBoss_ID = 1142 //
//
type S2CFieldBossV2PlayerInBoss struct {
	//
	BossId int32 `protobuf:"varint,1,opt,name=BossId,proto3" json:"BossId,omitempty"`
	//
	DamageOrder int32 `protobuf:"varint,2,opt,name=DamageOrder,proto3" json:"DamageOrder,omitempty"`
	//
	Damage int64 `protobuf:"varint,3,opt,name=Damage,proto3" json:"Damage,omitempty"`
}

func (m *S2CFieldBossV2PlayerInBoss) Reset()         { *m = S2CFieldBossV2PlayerInBoss{} }
func (m *S2CFieldBossV2PlayerInBoss) String() string { return proto.CompactTextString(m) }
func (*S2CFieldBossV2PlayerInBoss) ProtoMessage()    {}
func (m *S2CFieldBossV2PlayerInBoss) GetId() uint16  { return PCK_S2CFieldBossV2PlayerInBoss_ID }

const PCK_C2SFieldBossV2StartFight_ID = 1132 //
//
type C2SFieldBossV2StartFight struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2SFieldBossV2StartFight) Reset()         { *m = C2SFieldBossV2StartFight{} }
func (m *C2SFieldBossV2StartFight) String() string { return proto.CompactTextString(m) }
func (*C2SFieldBossV2StartFight) ProtoMessage()    {}
func (m *C2SFieldBossV2StartFight) GetId() uint16  { return PCK_C2SFieldBossV2StartFight_ID }

const PCK_S2CFieldBossV2StartFight_ID = 1133 //
//
type S2CFieldBossV2StartFight struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CFieldBossV2StartFight) Reset()         { *m = S2CFieldBossV2StartFight{} }
func (m *S2CFieldBossV2StartFight) String() string { return proto.CompactTextString(m) }
func (*S2CFieldBossV2StartFight) ProtoMessage()    {}
func (m *S2CFieldBossV2StartFight) GetId() uint16  { return PCK_S2CFieldBossV2StartFight_ID }

const PCK_C2SFieldBossV2Cfg_ID = 1134 //
//
type C2SFieldBossV2Cfg struct {
	//
	Cfg string `protobuf:"bytes,1,opt,name=Cfg,proto3" json:"Cfg,omitempty"`
}

func (m *C2SFieldBossV2Cfg) Reset()         { *m = C2SFieldBossV2Cfg{} }
func (m *C2SFieldBossV2Cfg) String() string { return proto.CompactTextString(m) }
func (*C2SFieldBossV2Cfg) ProtoMessage()    {}
func (m *C2SFieldBossV2Cfg) GetId() uint16  { return PCK_C2SFieldBossV2Cfg_ID }

const PCK_C2SFieldBossV2BuyTimes_ID = 1137 //
//
type C2SFieldBossV2BuyTimes struct {
}

func (m *C2SFieldBossV2BuyTimes) Reset()         { *m = C2SFieldBossV2BuyTimes{} }
func (m *C2SFieldBossV2BuyTimes) String() string { return proto.CompactTextString(m) }
func (*C2SFieldBossV2BuyTimes) ProtoMessage()    {}
func (m *C2SFieldBossV2BuyTimes) GetId() uint16  { return PCK_C2SFieldBossV2BuyTimes_ID }

const PCK_S2CFieldBossV2BuyTimes_ID = 1138 //
//
type S2CFieldBossV2BuyTimes struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CFieldBossV2BuyTimes) Reset()         { *m = S2CFieldBossV2BuyTimes{} }
func (m *S2CFieldBossV2BuyTimes) String() string { return proto.CompactTextString(m) }
func (*S2CFieldBossV2BuyTimes) ProtoMessage()    {}
func (m *S2CFieldBossV2BuyTimes) GetId() uint16  { return PCK_S2CFieldBossV2BuyTimes_ID }

//
type FieldBossDamageRank struct {
	//
	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	DamageTime int64 `protobuf:"varint,2,opt,name=DamageTime,proto3" json:"DamageTime,omitempty"`
	//
	Nick string `protobuf:"bytes,3,opt,name=Nick,proto3" json:"Nick,omitempty"`
	//
	Damage int64 `protobuf:"varint,4,opt,name=Damage,proto3" json:"Damage,omitempty"`
}

func (m *FieldBossDamageRank) Reset()         { *m = FieldBossDamageRank{} }
func (m *FieldBossDamageRank) String() string { return proto.CompactTextString(m) }
func (*FieldBossDamageRank) ProtoMessage()    {}

const PCK_C2SFieldBossV2GetDamageLog_ID = 1139 //
//
type C2SFieldBossV2GetDamageLog struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2SFieldBossV2GetDamageLog) Reset()         { *m = C2SFieldBossV2GetDamageLog{} }
func (m *C2SFieldBossV2GetDamageLog) String() string { return proto.CompactTextString(m) }
func (*C2SFieldBossV2GetDamageLog) ProtoMessage()    {}
func (m *C2SFieldBossV2GetDamageLog) GetId() uint16  { return PCK_C2SFieldBossV2GetDamageLog_ID }

const PCK_S2CFieldBossV2GetDamageLog_ID = 1140 //
//
type S2CFieldBossV2GetDamageLog struct {
	//
	BossId int32 `protobuf:"varint,1,opt,name=BossId,proto3" json:"BossId,omitempty"`
	//
	BossState int32 `protobuf:"varint,2,opt,name=BossState,proto3" json:"BossState,omitempty"`
	//
	Items []*FieldBossDamageRank `protobuf:"bytes,3,rep,name=Items,proto3" json:"Items,omitempty"`
}

func (m *S2CFieldBossV2GetDamageLog) Reset()         { *m = S2CFieldBossV2GetDamageLog{} }
func (m *S2CFieldBossV2GetDamageLog) String() string { return proto.CompactTextString(m) }
func (*S2CFieldBossV2GetDamageLog) ProtoMessage()    {}
func (m *S2CFieldBossV2GetDamageLog) GetId() uint16  { return PCK_S2CFieldBossV2GetDamageLog_ID }

//
type FieldBossInfo struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	S int32 `protobuf:"varint,2,opt,name=S,proto3" json:"S,omitempty"`
	//
	R int64 `protobuf:"varint,3,opt,name=R,proto3" json:"R,omitempty"`
	//
	F int64 `protobuf:"varint,4,opt,name=F,proto3" json:"F,omitempty"`
	//
	U int64 `protobuf:"varint,5,opt,name=U,proto3" json:"U,omitempty"`
	//
	N string `protobuf:"bytes,6,opt,name=N,proto3" json:"N,omitempty"`
	//
	O int64 `protobuf:"varint,7,opt,name=O,proto3" json:"O,omitempty"`
	//
	RoleId int64 `protobuf:"varint,8,opt,name=RoleId,proto3" json:"RoleId,omitempty"`
	//
	Hp int64 `protobuf:"varint,9,opt,name=Hp,proto3" json:"Hp,omitempty"`
	//
	MaxHp int64 `protobuf:"varint,10,opt,name=MaxHp,proto3" json:"MaxHp,omitempty"`
}

func (m *FieldBossInfo) Reset()         { *m = FieldBossInfo{} }
func (m *FieldBossInfo) String() string { return proto.CompactTextString(m) }
func (*FieldBossInfo) ProtoMessage()    {}

const PCK_S2CFieldBossInfo_ID = 651 //
//
type S2CFieldBossInfo struct {
	//
	Items []*FieldBossInfo `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty"`
}

func (m *S2CFieldBossInfo) Reset()         { *m = S2CFieldBossInfo{} }
func (m *S2CFieldBossInfo) String() string { return proto.CompactTextString(m) }
func (*S2CFieldBossInfo) ProtoMessage()    {}
func (m *S2CFieldBossInfo) GetId() uint16  { return PCK_S2CFieldBossInfo_ID }

const PCK_C2SFieldBossStart_ID = 652 //
//
type C2SFieldBossStart struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2SFieldBossStart) Reset()         { *m = C2SFieldBossStart{} }
func (m *C2SFieldBossStart) String() string { return proto.CompactTextString(m) }
func (*C2SFieldBossStart) ProtoMessage()    {}
func (m *C2SFieldBossStart) GetId() uint16  { return PCK_C2SFieldBossStart_ID }

const PCK_S2CFieldBossStart_ID = 653 //
//
type S2CFieldBossStart struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CFieldBossStart) Reset()         { *m = S2CFieldBossStart{} }
func (m *S2CFieldBossStart) String() string { return proto.CompactTextString(m) }
func (*S2CFieldBossStart) ProtoMessage()    {}
func (m *S2CFieldBossStart) GetId() uint16  { return PCK_S2CFieldBossStart_ID }

const PCK_C2SFieldBossItem_ID = 682 //
//
type C2SFieldBossItem struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2SFieldBossItem) Reset()         { *m = C2SFieldBossItem{} }
func (m *C2SFieldBossItem) String() string { return proto.CompactTextString(m) }
func (*C2SFieldBossItem) ProtoMessage()    {}
func (m *C2SFieldBossItem) GetId() uint16  { return PCK_C2SFieldBossItem_ID }

const PCK_S2CFieldBossItem_ID = 683 //
//
type S2CFieldBossItem struct {
	//
	Ret int32 `protobuf:"varint,1,opt,name=Ret,proto3" json:"Ret,omitempty"`
}

func (m *S2CFieldBossItem) Reset()         { *m = S2CFieldBossItem{} }
func (m *S2CFieldBossItem) String() string { return proto.CompactTextString(m) }
func (*S2CFieldBossItem) ProtoMessage()    {}
func (m *S2CFieldBossItem) GetId() uint16  { return PCK_S2CFieldBossItem_ID }

//
type BossVipInfo struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	K int64 `protobuf:"varint,2,opt,name=K,proto3" json:"K,omitempty"`
	//
	L int32 `protobuf:"varint,3,opt,name=L,proto3" json:"L,omitempty"`
}

func (m *BossVipInfo) Reset()         { *m = BossVipInfo{} }
func (m *BossVipInfo) String() string { return proto.CompactTextString(m) }
func (*BossVipInfo) ProtoMessage()    {}

const PCK_S2CBossVipInfo_ID = 661 //
//
type S2CBossVipInfo struct {
	//
	Items []*BossVipInfo `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty"`
}

func (m *S2CBossVipInfo) Reset()         { *m = S2CBossVipInfo{} }
func (m *S2CBossVipInfo) String() string { return proto.CompactTextString(m) }
func (*S2CBossVipInfo) ProtoMessage()    {}
func (m *S2CBossVipInfo) GetId() uint16  { return PCK_S2CBossVipInfo_ID }

const PCK_C2SBossVipStart_ID = 662 //
//
type C2SBossVipStart struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2SBossVipStart) Reset()         { *m = C2SBossVipStart{} }
func (m *C2SBossVipStart) String() string { return proto.CompactTextString(m) }
func (*C2SBossVipStart) ProtoMessage()    {}
func (m *C2SBossVipStart) GetId() uint16  { return PCK_C2SBossVipStart_ID }

const PCK_S2CBossVipStart_ID = 663 //
//
type S2CBossVipStart struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Win int32 `protobuf:"varint,3,opt,name=Win,proto3" json:"Win,omitempty"`
	//
	Sweep int32 `protobuf:"varint,4,opt,name=Sweep,proto3" json:"Sweep,omitempty"`
}

func (m *S2CBossVipStart) Reset()         { *m = S2CBossVipStart{} }
func (m *S2CBossVipStart) String() string { return proto.CompactTextString(m) }
func (*S2CBossVipStart) ProtoMessage()    {}
func (m *S2CBossVipStart) GetId() uint16  { return PCK_S2CBossVipStart_ID }

//
type Partner struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	T int32 `protobuf:"varint,2,opt,name=T,proto3" json:"T,omitempty"`
	//
	L int32 `protobuf:"varint,3,opt,name=L,proto3" json:"L,omitempty"`
	//
	E int32 `protobuf:"varint,4,opt,name=E,proto3" json:"E,omitempty"`
	//
	AS []*Skill `protobuf:"bytes,5,rep,name=AS,proto3" json:"AS,omitempty"`
	//
	PS []*Skill `protobuf:"bytes,6,rep,name=PS,proto3" json:"PS,omitempty"`
	//
	PS2 []*Skill `protobuf:"bytes,7,rep,name=PS2,proto3" json:"PS2,omitempty"`
	//
	R int32 `protobuf:"varint,8,opt,name=R,proto3" json:"R,omitempty"`
	//
	B int32 `protobuf:"varint,9,opt,name=B,proto3" json:"B,omitempty"`
	//
	S int32 `protobuf:"varint,10,opt,name=S,proto3" json:"S,omitempty"`
	//
	SE int32 `protobuf:"varint,11,opt,name=SE,proto3" json:"SE,omitempty"`
	//
	G int32 `protobuf:"varint,12,opt,name=G,proto3" json:"G,omitempty"`
	//
	GE int32 `protobuf:"varint,13,opt,name=GE,proto3" json:"GE,omitempty"`
	//
	N string `protobuf:"bytes,14,opt,name=N,proto3" json:"N,omitempty"`
	//
	State int32 `protobuf:"varint,15,opt,name=State,proto3" json:"State,omitempty"`
	//
	BuyHoleNum int32 `protobuf:"varint,16,opt,name=BuyHoleNum,proto3" json:"BuyHoleNum,omitempty"`
	//
	BookSkill []*Skill `protobuf:"bytes,17,rep,name=BookSkill,proto3" json:"BookSkill,omitempty"`
	//
	ChangeLevel int32 `protobuf:"varint,18,opt,name=ChangeLevel,proto3" json:"ChangeLevel,omitempty"`
	//
	Equip []*PartnerEquip `protobuf:"bytes,19,rep,name=Equip,proto3" json:"Equip,omitempty"`
	//
	Forbid int32 `protobuf:"varint,20,opt,name=Forbid,proto3" json:"Forbid,omitempty"`
	//
	FightValue int32 `protobuf:"varint,40,opt,name=FightValue,proto3" json:"FightValue,omitempty"`
}

func (m *Partner) Reset()         { *m = Partner{} }
func (m *Partner) String() string { return proto.CompactTextString(m) }
func (*Partner) ProtoMessage()    {}

//
type PartnerEquip struct {
	//
	EquipPart int32 `protobuf:"varint,1,opt,name=EquipPart,proto3" json:"EquipPart,omitempty"`
	//
	StrongLevel int32 `protobuf:"varint,2,opt,name=StrongLevel,proto3" json:"StrongLevel,omitempty"`
	//
	RefineLevel int32 `protobuf:"varint,3,opt,name=RefineLevel,proto3" json:"RefineLevel,omitempty"`
	//
	RefineExp int32 `protobuf:"varint,4,opt,name=RefineExp,proto3" json:"RefineExp,omitempty"`
}

func (m *PartnerEquip) Reset()         { *m = PartnerEquip{} }
func (m *PartnerEquip) String() string { return proto.CompactTextString(m) }
func (*PartnerEquip) ProtoMessage()    {}

const PCK_C2SGetPartner_ID = 896 //
//
type C2SGetPartner struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Type int32 `protobuf:"varint,2,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (m *C2SGetPartner) Reset()         { *m = C2SGetPartner{} }
func (m *C2SGetPartner) String() string { return proto.CompactTextString(m) }
func (*C2SGetPartner) ProtoMessage()    {}
func (m *C2SGetPartner) GetId() uint16  { return PCK_C2SGetPartner_ID }

const PCK_S2CGetPartner_ID = 897 //
//
type S2CGetPartner struct {
	//
	P *Partner `protobuf:"bytes,1,opt,name=P,proto3" json:"P,omitempty"`
}

func (m *S2CGetPartner) Reset()         { *m = S2CGetPartner{} }
func (m *S2CGetPartner) String() string { return proto.CompactTextString(m) }
func (*S2CGetPartner) ProtoMessage()    {}
func (m *S2CGetPartner) GetId() uint16  { return PCK_S2CGetPartner_ID }

const PCK_C2SPartnerDeify_ID = 898 //
//
type C2SPartnerDeify struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Type int32 `protobuf:"varint,2,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (m *C2SPartnerDeify) Reset()         { *m = C2SPartnerDeify{} }
func (m *C2SPartnerDeify) String() string { return proto.CompactTextString(m) }
func (*C2SPartnerDeify) ProtoMessage()    {}
func (m *C2SPartnerDeify) GetId() uint16  { return PCK_C2SPartnerDeify_ID }

const PCK_S2CPartnerDeify_ID = 899 //
//
type S2CPartnerDeify struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Type int32 `protobuf:"varint,3,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (m *S2CPartnerDeify) Reset()         { *m = S2CPartnerDeify{} }
func (m *S2CPartnerDeify) String() string { return proto.CompactTextString(m) }
func (*S2CPartnerDeify) ProtoMessage()    {}
func (m *S2CPartnerDeify) GetId() uint16  { return PCK_S2CPartnerDeify_ID }

const PCK_C2SPartnerBuySkillHole_ID = 902 //
//
type C2SPartnerBuySkillHole struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Type int32 `protobuf:"varint,2,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (m *C2SPartnerBuySkillHole) Reset()         { *m = C2SPartnerBuySkillHole{} }
func (m *C2SPartnerBuySkillHole) String() string { return proto.CompactTextString(m) }
func (*C2SPartnerBuySkillHole) ProtoMessage()    {}
func (m *C2SPartnerBuySkillHole) GetId() uint16  { return PCK_C2SPartnerBuySkillHole_ID }

const PCK_S2CPartnerBuySkillHole_ID = 903 //
//
type S2CPartnerBuySkillHole struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Type int32 `protobuf:"varint,3,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (m *S2CPartnerBuySkillHole) Reset()         { *m = S2CPartnerBuySkillHole{} }
func (m *S2CPartnerBuySkillHole) String() string { return proto.CompactTextString(m) }
func (*S2CPartnerBuySkillHole) ProtoMessage()    {}
func (m *S2CPartnerBuySkillHole) GetId() uint16  { return PCK_S2CPartnerBuySkillHole_ID }

const PCK_C2SPartnerLearnSkill_ID = 904 //
//
type C2SPartnerLearnSkill struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Type int32 `protobuf:"varint,2,opt,name=Type,proto3" json:"Type,omitempty"`
	//
	BookId int32 `protobuf:"varint,3,opt,name=BookId,proto3" json:"BookId,omitempty"`
}

func (m *C2SPartnerLearnSkill) Reset()         { *m = C2SPartnerLearnSkill{} }
func (m *C2SPartnerLearnSkill) String() string { return proto.CompactTextString(m) }
func (*C2SPartnerLearnSkill) ProtoMessage()    {}
func (m *C2SPartnerLearnSkill) GetId() uint16  { return PCK_C2SPartnerLearnSkill_ID }

const PCK_S2CPartnerLearnSkill_ID = 905 //
//
type S2CPartnerLearnSkill struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Type int32 `protobuf:"varint,3,opt,name=Type,proto3" json:"Type,omitempty"`
	//
	BookSkill []*Skill `protobuf:"bytes,4,rep,name=BookSkill,proto3" json:"BookSkill,omitempty"`
}

func (m *S2CPartnerLearnSkill) Reset()         { *m = S2CPartnerLearnSkill{} }
func (m *S2CPartnerLearnSkill) String() string { return proto.CompactTextString(m) }
func (*S2CPartnerLearnSkill) ProtoMessage()    {}
func (m *S2CPartnerLearnSkill) GetId() uint16  { return PCK_S2CPartnerLearnSkill_ID }

const PCK_C2SPartnerChange_ID = 906 //
//
type C2SPartnerChange struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Type int32 `protobuf:"varint,2,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (m *C2SPartnerChange) Reset()         { *m = C2SPartnerChange{} }
func (m *C2SPartnerChange) String() string { return proto.CompactTextString(m) }
func (*C2SPartnerChange) ProtoMessage()    {}
func (m *C2SPartnerChange) GetId() uint16  { return PCK_C2SPartnerChange_ID }

const PCK_S2CPartnerChange_ID = 907 //
//
type S2CPartnerChange struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Type int32 `protobuf:"varint,3,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (m *S2CPartnerChange) Reset()         { *m = S2CPartnerChange{} }
func (m *S2CPartnerChange) String() string { return proto.CompactTextString(m) }
func (*S2CPartnerChange) ProtoMessage()    {}
func (m *S2CPartnerChange) GetId() uint16  { return PCK_S2CPartnerChange_ID }

const PCK_C2SPartnerWearEquip_ID = 908 //
//
type C2SPartnerWearEquip struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Type int32 `protobuf:"varint,2,opt,name=Type,proto3" json:"Type,omitempty"`
	//
	OnlyId string `protobuf:"bytes,3,opt,name=OnlyId,proto3" json:"OnlyId,omitempty"`
}

func (m *C2SPartnerWearEquip) Reset()         { *m = C2SPartnerWearEquip{} }
func (m *C2SPartnerWearEquip) String() string { return proto.CompactTextString(m) }
func (*C2SPartnerWearEquip) ProtoMessage()    {}
func (m *C2SPartnerWearEquip) GetId() uint16  { return PCK_C2SPartnerWearEquip_ID }

const PCK_S2CPartnerWearEquip_ID = 909 //
//
type S2CPartnerWearEquip struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CPartnerWearEquip) Reset()         { *m = S2CPartnerWearEquip{} }
func (m *S2CPartnerWearEquip) String() string { return proto.CompactTextString(m) }
func (*S2CPartnerWearEquip) ProtoMessage()    {}
func (m *S2CPartnerWearEquip) GetId() uint16  { return PCK_S2CPartnerWearEquip_ID }

const PCK_C2SPartnerStrongEquip_ID = 910 //
//
type C2SPartnerStrongEquip struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Type int32 `protobuf:"varint,2,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (m *C2SPartnerStrongEquip) Reset()         { *m = C2SPartnerStrongEquip{} }
func (m *C2SPartnerStrongEquip) String() string { return proto.CompactTextString(m) }
func (*C2SPartnerStrongEquip) ProtoMessage()    {}
func (m *C2SPartnerStrongEquip) GetId() uint16  { return PCK_C2SPartnerStrongEquip_ID }

const PCK_S2CPartnerStrongEquip_ID = 911 //
//
type S2CPartnerStrongEquip struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Type int32 `protobuf:"varint,3,opt,name=Type,proto3" json:"Type,omitempty"`
	//
	Equips []*PartnerEquip `protobuf:"bytes,4,rep,name=Equips,proto3" json:"Equips,omitempty"`
}

func (m *S2CPartnerStrongEquip) Reset()         { *m = S2CPartnerStrongEquip{} }
func (m *S2CPartnerStrongEquip) String() string { return proto.CompactTextString(m) }
func (*S2CPartnerStrongEquip) ProtoMessage()    {}
func (m *S2CPartnerStrongEquip) GetId() uint16  { return PCK_S2CPartnerStrongEquip_ID }

const PCK_C2SPartnerRefineEquip_ID = 912 //
//
type C2SPartnerRefineEquip struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Type int32 `protobuf:"varint,2,opt,name=Type,proto3" json:"Type,omitempty"`
	//
	EquipPart int32 `protobuf:"varint,3,opt,name=EquipPart,proto3" json:"EquipPart,omitempty"`
}

func (m *C2SPartnerRefineEquip) Reset()         { *m = C2SPartnerRefineEquip{} }
func (m *C2SPartnerRefineEquip) String() string { return proto.CompactTextString(m) }
func (*C2SPartnerRefineEquip) ProtoMessage()    {}
func (m *C2SPartnerRefineEquip) GetId() uint16  { return PCK_C2SPartnerRefineEquip_ID }

const PCK_S2CPartnerRefineEquip_ID = 913 //
//
type S2CPartnerRefineEquip struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Type int32 `protobuf:"varint,3,opt,name=Type,proto3" json:"Type,omitempty"`
	//
	Equips []*PartnerEquip `protobuf:"bytes,4,rep,name=Equips,proto3" json:"Equips,omitempty"`
}

func (m *S2CPartnerRefineEquip) Reset()         { *m = S2CPartnerRefineEquip{} }
func (m *S2CPartnerRefineEquip) String() string { return proto.CompactTextString(m) }
func (*S2CPartnerRefineEquip) ProtoMessage()    {}
func (m *S2CPartnerRefineEquip) GetId() uint16  { return PCK_S2CPartnerRefineEquip_ID }

const PCK_C2SPartnerComposeEquip_ID = 914 //
//
type C2SPartnerComposeEquip struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Type int32 `protobuf:"varint,2,opt,name=Type,proto3" json:"Type,omitempty"`
	//
	EquipPart int32 `protobuf:"varint,3,opt,name=EquipPart,proto3" json:"EquipPart,omitempty"`
}

func (m *C2SPartnerComposeEquip) Reset()         { *m = C2SPartnerComposeEquip{} }
func (m *C2SPartnerComposeEquip) String() string { return proto.CompactTextString(m) }
func (*C2SPartnerComposeEquip) ProtoMessage()    {}
func (m *C2SPartnerComposeEquip) GetId() uint16  { return PCK_C2SPartnerComposeEquip_ID }

const PCK_S2CPartnerComposeEquip_ID = 915 //
//
type S2CPartnerComposeEquip struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Type int32 `protobuf:"varint,3,opt,name=Type,proto3" json:"Type,omitempty"`
	//
	EquipPart int32 `protobuf:"varint,4,opt,name=EquipPart,proto3" json:"EquipPart,omitempty"`
}

func (m *S2CPartnerComposeEquip) Reset()         { *m = S2CPartnerComposeEquip{} }
func (m *S2CPartnerComposeEquip) String() string { return proto.CompactTextString(m) }
func (*S2CPartnerComposeEquip) ProtoMessage()    {}
func (m *S2CPartnerComposeEquip) GetId() uint16  { return PCK_S2CPartnerComposeEquip_ID }

const PCK_C2SPartnerTakeOffEquip_ID = 916 //
//
type C2SPartnerTakeOffEquip struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Type int32 `protobuf:"varint,2,opt,name=Type,proto3" json:"Type,omitempty"`
	//
	OnlyId string `protobuf:"bytes,3,opt,name=OnlyId,proto3" json:"OnlyId,omitempty"`
}

func (m *C2SPartnerTakeOffEquip) Reset()         { *m = C2SPartnerTakeOffEquip{} }
func (m *C2SPartnerTakeOffEquip) String() string { return proto.CompactTextString(m) }
func (*C2SPartnerTakeOffEquip) ProtoMessage()    {}
func (m *C2SPartnerTakeOffEquip) GetId() uint16  { return PCK_C2SPartnerTakeOffEquip_ID }

const PCK_S2CPartnerTakeOffEquip_ID = 917 //
//
type S2CPartnerTakeOffEquip struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Type int32 `protobuf:"varint,3,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (m *S2CPartnerTakeOffEquip) Reset()         { *m = S2CPartnerTakeOffEquip{} }
func (m *S2CPartnerTakeOffEquip) String() string { return proto.CompactTextString(m) }
func (*S2CPartnerTakeOffEquip) ProtoMessage()    {}
func (m *S2CPartnerTakeOffEquip) GetId() uint16  { return PCK_S2CPartnerTakeOffEquip_ID }

const PCK_C2SPartnerLevelUp_ID = 801 //
//
type C2SPartnerLevelUp struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Type int32 `protobuf:"varint,2,opt,name=Type,proto3" json:"Type,omitempty"`
	//
	Auto int32 `protobuf:"varint,3,opt,name=Auto,proto3" json:"Auto,omitempty"`
}

func (m *C2SPartnerLevelUp) Reset()         { *m = C2SPartnerLevelUp{} }
func (m *C2SPartnerLevelUp) String() string { return proto.CompactTextString(m) }
func (*C2SPartnerLevelUp) ProtoMessage()    {}
func (m *C2SPartnerLevelUp) GetId() uint16  { return PCK_C2SPartnerLevelUp_ID }

const PCK_S2CPartnerLevelUp_ID = 802 //
//
type S2CPartnerLevelUp struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Type int32 `protobuf:"varint,3,opt,name=Type,proto3" json:"Type,omitempty"`
	//
	L int32 `protobuf:"varint,4,opt,name=L,proto3" json:"L,omitempty"`
	//
	E int32 `protobuf:"varint,5,opt,name=E,proto3" json:"E,omitempty"`
	//
	Multi int32 `protobuf:"varint,6,opt,name=Multi,proto3" json:"Multi,omitempty"`
}

func (m *S2CPartnerLevelUp) Reset()         { *m = S2CPartnerLevelUp{} }
func (m *S2CPartnerLevelUp) String() string { return proto.CompactTextString(m) }
func (*S2CPartnerLevelUp) ProtoMessage()    {}
func (m *S2CPartnerLevelUp) GetId() uint16  { return PCK_S2CPartnerLevelUp_ID }

const PCK_C2SPartnerStarUp_ID = 803 //
//
type C2SPartnerStarUp struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Type int32 `protobuf:"varint,2,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (m *C2SPartnerStarUp) Reset()         { *m = C2SPartnerStarUp{} }
func (m *C2SPartnerStarUp) String() string { return proto.CompactTextString(m) }
func (*C2SPartnerStarUp) ProtoMessage()    {}
func (m *C2SPartnerStarUp) GetId() uint16  { return PCK_C2SPartnerStarUp_ID }

const PCK_S2CPartnerStarUp_ID = 804 //
//
type S2CPartnerStarUp struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Type int32 `protobuf:"varint,3,opt,name=Type,proto3" json:"Type,omitempty"`
	//
	S int32 `protobuf:"varint,4,opt,name=S,proto3" json:"S,omitempty"`
	//
	SE int32 `protobuf:"varint,5,opt,name=SE,proto3" json:"SE,omitempty"`
	//
	ActiveSkill *Skill `protobuf:"bytes,6,opt,name=ActiveSkill,proto3" json:"ActiveSkill,omitempty"`
}

func (m *S2CPartnerStarUp) Reset()         { *m = S2CPartnerStarUp{} }
func (m *S2CPartnerStarUp) String() string { return proto.CompactTextString(m) }
func (*S2CPartnerStarUp) ProtoMessage()    {}
func (m *S2CPartnerStarUp) GetId() uint16  { return PCK_S2CPartnerStarUp_ID }

const PCK_C2SPartnerGradeUp_ID = 817 //
//
type C2SPartnerGradeUp struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Type int32 `protobuf:"varint,2,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (m *C2SPartnerGradeUp) Reset()         { *m = C2SPartnerGradeUp{} }
func (m *C2SPartnerGradeUp) String() string { return proto.CompactTextString(m) }
func (*C2SPartnerGradeUp) ProtoMessage()    {}
func (m *C2SPartnerGradeUp) GetId() uint16  { return PCK_C2SPartnerGradeUp_ID }

const PCK_S2CPartnerGradeUp_ID = 818 //
//
type S2CPartnerGradeUp struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Type int32 `protobuf:"varint,3,opt,name=Type,proto3" json:"Type,omitempty"`
	//
	G int32 `protobuf:"varint,4,opt,name=G,proto3" json:"G,omitempty"`
	//
	GE int32 `protobuf:"varint,5,opt,name=GE,proto3" json:"GE,omitempty"`
}

func (m *S2CPartnerGradeUp) Reset()         { *m = S2CPartnerGradeUp{} }
func (m *S2CPartnerGradeUp) String() string { return proto.CompactTextString(m) }
func (*S2CPartnerGradeUp) ProtoMessage()    {}
func (m *S2CPartnerGradeUp) GetId() uint16  { return PCK_S2CPartnerGradeUp_ID }

const PCK_C2SPartnerBattlePos_ID = 805 //
//
type C2SPartnerBattlePos struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Type int32 `protobuf:"varint,2,opt,name=Type,proto3" json:"Type,omitempty"`
	//
	Pos int32 `protobuf:"varint,3,opt,name=Pos,proto3" json:"Pos,omitempty"`
}

func (m *C2SPartnerBattlePos) Reset()         { *m = C2SPartnerBattlePos{} }
func (m *C2SPartnerBattlePos) String() string { return proto.CompactTextString(m) }
func (*C2SPartnerBattlePos) ProtoMessage()    {}
func (m *C2SPartnerBattlePos) GetId() uint16  { return PCK_C2SPartnerBattlePos_ID }

const PCK_S2CPartnerBattlePos_ID = 806 //
//
type S2CPartnerBattlePos struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Type int32 `protobuf:"varint,3,opt,name=Type,proto3" json:"Type,omitempty"`
	//
	RestId int32 `protobuf:"varint,5,opt,name=RestId,proto3" json:"RestId,omitempty"`
	//
	Pos int32 `protobuf:"varint,4,opt,name=Pos,proto3" json:"Pos,omitempty"`
}

func (m *S2CPartnerBattlePos) Reset()         { *m = S2CPartnerBattlePos{} }
func (m *S2CPartnerBattlePos) String() string { return proto.CompactTextString(m) }
func (*S2CPartnerBattlePos) ProtoMessage()    {}
func (m *S2CPartnerBattlePos) GetId() uint16  { return PCK_S2CPartnerBattlePos_ID }

const PCK_C2SPartnerBattlePosRobot_ID = 895 //
//
type C2SPartnerBattlePosRobot struct {
	//
	Type int32 `protobuf:"varint,1,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (m *C2SPartnerBattlePosRobot) Reset()         { *m = C2SPartnerBattlePosRobot{} }
func (m *C2SPartnerBattlePosRobot) String() string { return proto.CompactTextString(m) }
func (*C2SPartnerBattlePosRobot) ProtoMessage()    {}
func (m *C2SPartnerBattlePosRobot) GetId() uint16  { return PCK_C2SPartnerBattlePosRobot_ID }

const PCK_C2SPartnerRefreshSkill_ID = 807 //
//
type C2SPartnerRefreshSkill struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Type int32 `protobuf:"varint,2,opt,name=Type,proto3" json:"Type,omitempty"`
	//
	RefreshType int32 `protobuf:"varint,3,opt,name=RefreshType,proto3" json:"RefreshType,omitempty"`
	//
	LockSkills []*Skill `protobuf:"bytes,4,rep,name=LockSkills,proto3" json:"LockSkills,omitempty"`
	//
	Auto int32 `protobuf:"varint,5,opt,name=Auto,proto3" json:"Auto,omitempty"`
}

func (m *C2SPartnerRefreshSkill) Reset()         { *m = C2SPartnerRefreshSkill{} }
func (m *C2SPartnerRefreshSkill) String() string { return proto.CompactTextString(m) }
func (*C2SPartnerRefreshSkill) ProtoMessage()    {}
func (m *C2SPartnerRefreshSkill) GetId() uint16  { return PCK_C2SPartnerRefreshSkill_ID }

const PCK_S2CPartnerRefreshSkill_ID = 808 //
//
type S2CPartnerRefreshSkill struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Skills []*Skill `protobuf:"bytes,2,rep,name=Skills,proto3" json:"Skills,omitempty"`
	//
	Id int32 `protobuf:"varint,3,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Type int32 `protobuf:"varint,4,opt,name=Type,proto3" json:"Type,omitempty"`
	//
	RefreshType int32 `protobuf:"varint,5,opt,name=RefreshType,proto3" json:"RefreshType,omitempty"`
	//
	LockSkills []*Skill `protobuf:"bytes,6,rep,name=LockSkills,proto3" json:"LockSkills,omitempty"`
	//
	Times int32 `protobuf:"varint,7,opt,name=Times,proto3" json:"Times,omitempty"`
}

func (m *S2CPartnerRefreshSkill) Reset()         { *m = S2CPartnerRefreshSkill{} }
func (m *S2CPartnerRefreshSkill) String() string { return proto.CompactTextString(m) }
func (*S2CPartnerRefreshSkill) ProtoMessage()    {}
func (m *S2CPartnerRefreshSkill) GetId() uint16  { return PCK_S2CPartnerRefreshSkill_ID }

const PCK_C2SPartnerDelSkill_ID = 860 //
//
type C2SPartnerDelSkill struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	SkillId int32 `protobuf:"varint,2,opt,name=SkillId,proto3" json:"SkillId,omitempty"`
}

func (m *C2SPartnerDelSkill) Reset()         { *m = C2SPartnerDelSkill{} }
func (m *C2SPartnerDelSkill) String() string { return proto.CompactTextString(m) }
func (*C2SPartnerDelSkill) ProtoMessage()    {}
func (m *C2SPartnerDelSkill) GetId() uint16  { return PCK_C2SPartnerDelSkill_ID }

const PCK_S2CPartnerDelSkill_ID = 861 //
//
type S2CPartnerDelSkill struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Skills []*Skill `protobuf:"bytes,3,rep,name=Skills,proto3" json:"Skills,omitempty"`
}

func (m *S2CPartnerDelSkill) Reset()         { *m = S2CPartnerDelSkill{} }
func (m *S2CPartnerDelSkill) String() string { return proto.CompactTextString(m) }
func (*S2CPartnerDelSkill) ProtoMessage()    {}
func (m *S2CPartnerDelSkill) GetId() uint16  { return PCK_S2CPartnerDelSkill_ID }

const PCK_C2SPartnerUpSkill_ID = 862 //
//
type C2SPartnerUpSkill struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	SkillId int32 `protobuf:"varint,2,opt,name=SkillId,proto3" json:"SkillId,omitempty"`
}

func (m *C2SPartnerUpSkill) Reset()         { *m = C2SPartnerUpSkill{} }
func (m *C2SPartnerUpSkill) String() string { return proto.CompactTextString(m) }
func (*C2SPartnerUpSkill) ProtoMessage()    {}
func (m *C2SPartnerUpSkill) GetId() uint16  { return PCK_C2SPartnerUpSkill_ID }

const PCK_S2CPartnerUpSkill_ID = 863 //
//
type S2CPartnerUpSkill struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Skills []*Skill `protobuf:"bytes,3,rep,name=Skills,proto3" json:"Skills,omitempty"`
}

func (m *S2CPartnerUpSkill) Reset()         { *m = S2CPartnerUpSkill{} }
func (m *S2CPartnerUpSkill) String() string { return proto.CompactTextString(m) }
func (*S2CPartnerUpSkill) ProtoMessage()    {}
func (m *S2CPartnerUpSkill) GetId() uint16  { return PCK_S2CPartnerUpSkill_ID }

const PCK_C2SPartnerReplaceSkill_ID = 809 //
//
type C2SPartnerReplaceSkill struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Type int32 `protobuf:"varint,2,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (m *C2SPartnerReplaceSkill) Reset()         { *m = C2SPartnerReplaceSkill{} }
func (m *C2SPartnerReplaceSkill) String() string { return proto.CompactTextString(m) }
func (*C2SPartnerReplaceSkill) ProtoMessage()    {}
func (m *C2SPartnerReplaceSkill) GetId() uint16  { return PCK_C2SPartnerReplaceSkill_ID }

const PCK_S2CPartnerReplaceSkill_ID = 810 //
//
type S2CPartnerReplaceSkill struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Type int32 `protobuf:"varint,3,opt,name=Type,proto3" json:"Type,omitempty"`
	//
	PS []*Skill `protobuf:"bytes,4,rep,name=PS,proto3" json:"PS,omitempty"`
}

func (m *S2CPartnerReplaceSkill) Reset()         { *m = S2CPartnerReplaceSkill{} }
func (m *S2CPartnerReplaceSkill) String() string { return proto.CompactTextString(m) }
func (*S2CPartnerReplaceSkill) ProtoMessage()    {}
func (m *S2CPartnerReplaceSkill) GetId() uint16  { return PCK_S2CPartnerReplaceSkill_ID }

const PCK_C2SPartnerNick_ID = 811 //
//
type C2SPartnerNick struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Type int32 `protobuf:"varint,2,opt,name=Type,proto3" json:"Type,omitempty"`
	//
	N string `protobuf:"bytes,3,opt,name=N,proto3" json:"N,omitempty"`
}

func (m *C2SPartnerNick) Reset()         { *m = C2SPartnerNick{} }
func (m *C2SPartnerNick) String() string { return proto.CompactTextString(m) }
func (*C2SPartnerNick) ProtoMessage()    {}
func (m *C2SPartnerNick) GetId() uint16  { return PCK_C2SPartnerNick_ID }

const PCK_S2CPartnerNick_ID = 812 //
//
type S2CPartnerNick struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id int32 `protobuf:"varint,4,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Type int32 `protobuf:"varint,2,opt,name=Type,proto3" json:"Type,omitempty"`
	//
	N string `protobuf:"bytes,3,opt,name=N,proto3" json:"N,omitempty"`
}

func (m *S2CPartnerNick) Reset()         { *m = S2CPartnerNick{} }
func (m *S2CPartnerNick) String() string { return proto.CompactTextString(m) }
func (*S2CPartnerNick) ProtoMessage()    {}
func (m *S2CPartnerNick) GetId() uint16  { return PCK_S2CPartnerNick_ID }

const PCK_C2SPartnerSuit_ID = 813 //
//
type C2SPartnerSuit struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2SPartnerSuit) Reset()         { *m = C2SPartnerSuit{} }
func (m *C2SPartnerSuit) String() string { return proto.CompactTextString(m) }
func (*C2SPartnerSuit) ProtoMessage()    {}
func (m *C2SPartnerSuit) GetId() uint16  { return PCK_C2SPartnerSuit_ID }

const PCK_S2CPartnerSuit_ID = 814 //
//
type S2CPartnerSuit struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *S2CPartnerSuit) Reset()         { *m = S2CPartnerSuit{} }
func (m *S2CPartnerSuit) String() string { return proto.CompactTextString(m) }
func (*S2CPartnerSuit) ProtoMessage()    {}
func (m *S2CPartnerSuit) GetId() uint16  { return PCK_S2CPartnerSuit_ID }

const PCK_C2SPartnerSuitRobot_ID = 822 //
//
type C2SPartnerSuitRobot struct {
}

func (m *C2SPartnerSuitRobot) Reset()         { *m = C2SPartnerSuitRobot{} }
func (m *C2SPartnerSuitRobot) String() string { return proto.CompactTextString(m) }
func (*C2SPartnerSuitRobot) ProtoMessage()    {}
func (m *C2SPartnerSuitRobot) GetId() uint16  { return PCK_C2SPartnerSuitRobot_ID }

const PCK_C2SUnlockNewPartner_ID = 815 //
//
type C2SUnlockNewPartner struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Type int32 `protobuf:"varint,2,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (m *C2SUnlockNewPartner) Reset()         { *m = C2SUnlockNewPartner{} }
func (m *C2SUnlockNewPartner) String() string { return proto.CompactTextString(m) }
func (*C2SUnlockNewPartner) ProtoMessage()    {}
func (m *C2SUnlockNewPartner) GetId() uint16  { return PCK_C2SUnlockNewPartner_ID }

const PCK_S2CUnlockNewPartner_ID = 816 //
//
type S2CUnlockNewPartner struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Type int32 `protobuf:"varint,3,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (m *S2CUnlockNewPartner) Reset()         { *m = S2CUnlockNewPartner{} }
func (m *S2CUnlockNewPartner) String() string { return proto.CompactTextString(m) }
func (*S2CUnlockNewPartner) ProtoMessage()    {}
func (m *S2CUnlockNewPartner) GetId() uint16  { return PCK_S2CUnlockNewPartner_ID }

const PCK_C2SStickPartner_ID = 819 //
//
type C2SStickPartner struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Type int32 `protobuf:"varint,2,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (m *C2SStickPartner) Reset()         { *m = C2SStickPartner{} }
func (m *C2SStickPartner) String() string { return proto.CompactTextString(m) }
func (*C2SStickPartner) ProtoMessage()    {}
func (m *C2SStickPartner) GetId() uint16  { return PCK_C2SStickPartner_ID }

const PCK_S2CStickPartner_ID = 820 //
//
type S2CStickPartner struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Type int32 `protobuf:"varint,3,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (m *S2CStickPartner) Reset()         { *m = S2CStickPartner{} }
func (m *S2CStickPartner) String() string { return proto.CompactTextString(m) }
func (*S2CStickPartner) ProtoMessage()    {}
func (m *S2CStickPartner) GetId() uint16  { return PCK_S2CStickPartner_ID }

const PCK_S2CPartnerNewSkill_ID = 864 //
//
type S2CPartnerNewSkill struct {
	//
	Type int32 `protobuf:"varint,1,opt,name=Type,proto3" json:"Type,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	PS []*Skill `protobuf:"bytes,4,rep,name=PS,proto3" json:"PS,omitempty"`
}

func (m *S2CPartnerNewSkill) Reset()         { *m = S2CPartnerNewSkill{} }
func (m *S2CPartnerNewSkill) String() string { return proto.CompactTextString(m) }
func (*S2CPartnerNewSkill) ProtoMessage()    {}
func (m *S2CPartnerNewSkill) GetId() uint16  { return PCK_S2CPartnerNewSkill_ID }

//
type Pet struct {
	//
	P *Partner `protobuf:"bytes,1,opt,name=P,proto3" json:"P,omitempty"`
}

func (m *Pet) Reset()         { *m = Pet{} }
func (m *Pet) String() string { return proto.CompactTextString(m) }
func (*Pet) ProtoMessage()    {}

const PCK_S2CUserPet_ID = 821 //
//
type S2CUserPet struct {
	//
	P []*Pet `protobuf:"bytes,1,rep,name=P,proto3" json:"P,omitempty"`
	//
	Suit []int32 `protobuf:"varint,2,rep,name=Suit,proto3" json:"Suit,omitempty"`
}

func (m *S2CUserPet) Reset()         { *m = S2CUserPet{} }
func (m *S2CUserPet) String() string { return proto.CompactTextString(m) }
func (*S2CUserPet) ProtoMessage()    {}
func (m *S2CUserPet) GetId() uint16  { return PCK_S2CUserPet_ID }

//
type PetA struct {
	//
	P *Partner `protobuf:"bytes,1,opt,name=P,proto3" json:"P,omitempty"`
}

func (m *PetA) Reset()         { *m = PetA{} }
func (m *PetA) String() string { return proto.CompactTextString(m) }
func (*PetA) ProtoMessage()    {}

const PCK_S2CUserPetA_ID = 831 //
//
type S2CUserPetA struct {
	//
	P []*PetA `protobuf:"bytes,1,rep,name=P,proto3" json:"P,omitempty"`
}

func (m *S2CUserPetA) Reset()         { *m = S2CUserPetA{} }
func (m *S2CUserPetA) String() string { return proto.CompactTextString(m) }
func (*S2CUserPetA) ProtoMessage()    {}
func (m *S2CUserPetA) GetId() uint16  { return PCK_S2CUserPetA_ID }

//
type Devil struct {
	//
	P *Partner `protobuf:"bytes,1,opt,name=P,proto3" json:"P,omitempty"`
}

func (m *Devil) Reset()         { *m = Devil{} }
func (m *Devil) String() string { return proto.CompactTextString(m) }
func (*Devil) ProtoMessage()    {}

const PCK_S2CUserDevil_ID = 841 //
//
type S2CUserDevil struct {
	//
	P []*Devil `protobuf:"bytes,1,rep,name=P,proto3" json:"P,omitempty"`
	//
	DevilPos []*DevilPos `protobuf:"bytes,2,rep,name=DevilPos,proto3" json:"DevilPos,omitempty"`
}

func (m *S2CUserDevil) Reset()         { *m = S2CUserDevil{} }
func (m *S2CUserDevil) String() string { return proto.CompactTextString(m) }
func (*S2CUserDevil) ProtoMessage()    {}
func (m *S2CUserDevil) GetId() uint16  { return PCK_S2CUserDevil_ID }

const PCK_C2SDevilAwake_ID = 842 //
//
type C2SDevilAwake struct {
}

func (m *C2SDevilAwake) Reset()         { *m = C2SDevilAwake{} }
func (m *C2SDevilAwake) String() string { return proto.CompactTextString(m) }
func (*C2SDevilAwake) ProtoMessage()    {}
func (m *C2SDevilAwake) GetId() uint16  { return PCK_C2SDevilAwake_ID }

const PCK_S2CDevilAwake_ID = 843 //
//
type S2CDevilAwake struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CDevilAwake) Reset()         { *m = S2CDevilAwake{} }
func (m *S2CDevilAwake) String() string { return proto.CompactTextString(m) }
func (*S2CDevilAwake) ProtoMessage()    {}
func (m *S2CDevilAwake) GetId() uint16  { return PCK_S2CDevilAwake_ID }

const PCK_C2SDevilLevelUp_ID = 844 //
//
type C2SDevilLevelUp struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Item []*ItemInfo `protobuf:"bytes,2,rep,name=Item,proto3" json:"Item,omitempty"`
}

func (m *C2SDevilLevelUp) Reset()         { *m = C2SDevilLevelUp{} }
func (m *C2SDevilLevelUp) String() string { return proto.CompactTextString(m) }
func (*C2SDevilLevelUp) ProtoMessage()    {}
func (m *C2SDevilLevelUp) GetId() uint16  { return PCK_C2SDevilLevelUp_ID }

const PCK_S2CDevilLevelUp_ID = 845 //
//
type S2CDevilLevelUp struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	L int32 `protobuf:"varint,3,opt,name=L,proto3" json:"L,omitempty"`
	//
	E int32 `protobuf:"varint,4,opt,name=E,proto3" json:"E,omitempty"`
}

func (m *S2CDevilLevelUp) Reset()         { *m = S2CDevilLevelUp{} }
func (m *S2CDevilLevelUp) String() string { return proto.CompactTextString(m) }
func (*S2CDevilLevelUp) ProtoMessage()    {}
func (m *S2CDevilLevelUp) GetId() uint16  { return PCK_S2CDevilLevelUp_ID }

//
type DevilPos struct {
	//
	T int32 `protobuf:"varint,1,opt,name=T,proto3" json:"T,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *DevilPos) Reset()         { *m = DevilPos{} }
func (m *DevilPos) String() string { return proto.CompactTextString(m) }
func (*DevilPos) ProtoMessage()    {}

const PCK_C2SDevilPos_ID = 846 //
//
type C2SDevilPos struct {
	//
	T int32 `protobuf:"varint,1,opt,name=T,proto3" json:"T,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2SDevilPos) Reset()         { *m = C2SDevilPos{} }
func (m *C2SDevilPos) String() string { return proto.CompactTextString(m) }
func (*C2SDevilPos) ProtoMessage()    {}
func (m *C2SDevilPos) GetId() uint16  { return PCK_C2SDevilPos_ID }

const PCK_S2CDevilPos_ID = 847 //
//
type S2CDevilPos struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	DevilPos []*DevilPos `protobuf:"bytes,2,rep,name=DevilPos,proto3" json:"DevilPos,omitempty"`
}

func (m *S2CDevilPos) Reset()         { *m = S2CDevilPos{} }
func (m *S2CDevilPos) String() string { return proto.CompactTextString(m) }
func (*S2CDevilPos) ProtoMessage()    {}
func (m *S2CDevilPos) GetId() uint16  { return PCK_S2CDevilPos_ID }

const PCK_C2SDevilSoulPos_ID = 17202 //
//
type C2SDevilSoulPos struct {
	//
	Pos int32 `protobuf:"varint,1,opt,name=Pos,proto3" json:"Pos,omitempty"`
	//
	Id string `protobuf:"bytes,2,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2SDevilSoulPos) Reset()         { *m = C2SDevilSoulPos{} }
func (m *C2SDevilSoulPos) String() string { return proto.CompactTextString(m) }
func (*C2SDevilSoulPos) ProtoMessage()    {}
func (m *C2SDevilSoulPos) GetId() uint16  { return PCK_C2SDevilSoulPos_ID }

const PCK_S2CDevilSoulPos_ID = 17203 //
//
type S2CDevilSoulPos struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CDevilSoulPos) Reset()         { *m = S2CDevilSoulPos{} }
func (m *S2CDevilSoulPos) String() string { return proto.CompactTextString(m) }
func (*S2CDevilSoulPos) ProtoMessage()    {}
func (m *S2CDevilSoulPos) GetId() uint16  { return PCK_S2CDevilSoulPos_ID }

const PCK_C2SDevilSoulDel_ID = 17208 //
//
type C2SDevilSoulDel struct {
	//
	Id []string `protobuf:"bytes,1,rep,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2SDevilSoulDel) Reset()         { *m = C2SDevilSoulDel{} }
func (m *C2SDevilSoulDel) String() string { return proto.CompactTextString(m) }
func (*C2SDevilSoulDel) ProtoMessage()    {}
func (m *C2SDevilSoulDel) GetId() uint16  { return PCK_C2SDevilSoulDel_ID }

const PCK_S2CDevilSoulDel_ID = 17209 //
//
type S2CDevilSoulDel struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CDevilSoulDel) Reset()         { *m = S2CDevilSoulDel{} }
func (m *S2CDevilSoulDel) String() string { return proto.CompactTextString(m) }
func (*S2CDevilSoulDel) ProtoMessage()    {}
func (m *S2CDevilSoulDel) GetId() uint16  { return PCK_S2CDevilSoulDel_ID }

const PCK_C2SDevilSoulAuto_ID = 17210 //
//
type C2SDevilSoulAuto struct {
	//
	Auto int32 `protobuf:"varint,1,opt,name=Auto,proto3" json:"Auto,omitempty"`
}

func (m *C2SDevilSoulAuto) Reset()         { *m = C2SDevilSoulAuto{} }
func (m *C2SDevilSoulAuto) String() string { return proto.CompactTextString(m) }
func (*C2SDevilSoulAuto) ProtoMessage()    {}
func (m *C2SDevilSoulAuto) GetId() uint16  { return PCK_C2SDevilSoulAuto_ID }

const PCK_S2CDevilSoulAuto_ID = 17211 //
//
type S2CDevilSoulAuto struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CDevilSoulAuto) Reset()         { *m = S2CDevilSoulAuto{} }
func (m *S2CDevilSoulAuto) String() string { return proto.CompactTextString(m) }
func (*S2CDevilSoulAuto) ProtoMessage()    {}
func (m *S2CDevilSoulAuto) GetId() uint16  { return PCK_S2CDevilSoulAuto_ID }

const PCK_C2SDevilSoulLevelUp_ID = 17204 //
//
type C2SDevilSoulLevelUp struct {
	//
	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2SDevilSoulLevelUp) Reset()         { *m = C2SDevilSoulLevelUp{} }
func (m *C2SDevilSoulLevelUp) String() string { return proto.CompactTextString(m) }
func (*C2SDevilSoulLevelUp) ProtoMessage()    {}
func (m *C2SDevilSoulLevelUp) GetId() uint16  { return PCK_C2SDevilSoulLevelUp_ID }

const PCK_S2CDevilSoulLevelUp_ID = 17205 //
//
type S2CDevilSoulLevelUp struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id string `protobuf:"bytes,2,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	DevilSoul *ItemData `protobuf:"bytes,3,opt,name=DevilSoul,proto3" json:"DevilSoul,omitempty"`
}

func (m *S2CDevilSoulLevelUp) Reset()         { *m = S2CDevilSoulLevelUp{} }
func (m *S2CDevilSoulLevelUp) String() string { return proto.CompactTextString(m) }
func (*S2CDevilSoulLevelUp) ProtoMessage()    {}
func (m *S2CDevilSoulLevelUp) GetId() uint16  { return PCK_S2CDevilSoulLevelUp_ID }

const PCK_C2SDevilSoulStarUp_ID = 17206 //
//
type C2SDevilSoulStarUp struct {
	//
	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	SameId []string `protobuf:"bytes,2,rep,name=SameId,proto3" json:"SameId,omitempty"`
	//
	AnyId []string `protobuf:"bytes,3,rep,name=AnyId,proto3" json:"AnyId,omitempty"`
}

func (m *C2SDevilSoulStarUp) Reset()         { *m = C2SDevilSoulStarUp{} }
func (m *C2SDevilSoulStarUp) String() string { return proto.CompactTextString(m) }
func (*C2SDevilSoulStarUp) ProtoMessage()    {}
func (m *C2SDevilSoulStarUp) GetId() uint16  { return PCK_C2SDevilSoulStarUp_ID }

const PCK_S2CDevilSoulStarUp_ID = 17207 //
//
type S2CDevilSoulStarUp struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id string `protobuf:"bytes,2,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	DevilSoul *ItemData `protobuf:"bytes,3,opt,name=DevilSoul,proto3" json:"DevilSoul,omitempty"`
}

func (m *S2CDevilSoulStarUp) Reset()         { *m = S2CDevilSoulStarUp{} }
func (m *S2CDevilSoulStarUp) String() string { return proto.CompactTextString(m) }
func (*S2CDevilSoulStarUp) ProtoMessage()    {}
func (m *S2CDevilSoulStarUp) GetId() uint16  { return PCK_S2CDevilSoulStarUp_ID }

const PCK_Beauty_ID = 851 //
//
type Beauty struct {
	//
	P *Partner `protobuf:"bytes,1,opt,name=P,proto3" json:"P,omitempty"`
	//
	Magics []*BeautyMagic `protobuf:"bytes,2,rep,name=Magics,proto3" json:"Magics,omitempty"`
	//
	Cloth []*BeautyCloth `protobuf:"bytes,3,rep,name=Cloth,proto3" json:"Cloth,omitempty"`
}

func (m *Beauty) Reset()         { *m = Beauty{} }
func (m *Beauty) String() string { return proto.CompactTextString(m) }
func (*Beauty) ProtoMessage()    {}
func (m *Beauty) GetId() uint16  { return PCK_Beauty_ID }

//
type BeautyMagic struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Items []int32 `protobuf:"varint,2,rep,name=Items,proto3" json:"Items,omitempty"`
	//
	TempItems []int32 `protobuf:"varint,3,rep,name=TempItems,proto3" json:"TempItems,omitempty"`
	//
	Times int32 `protobuf:"varint,4,opt,name=Times,proto3" json:"Times,omitempty"`
}

func (m *BeautyMagic) Reset()         { *m = BeautyMagic{} }
func (m *BeautyMagic) String() string { return proto.CompactTextString(m) }
func (*BeautyMagic) ProtoMessage()    {}

//
type BeautyCloth struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	On int32 `protobuf:"varint,2,opt,name=On,proto3" json:"On,omitempty"`
	//
	L int32 `protobuf:"varint,3,opt,name=L,proto3" json:"L,omitempty"`
}

func (m *BeautyCloth) Reset()         { *m = BeautyCloth{} }
func (m *BeautyCloth) String() string { return proto.CompactTextString(m) }
func (*BeautyCloth) ProtoMessage()    {}

const PCK_C2SBeautyNewCloth_ID = 871 //
//
type C2SBeautyNewCloth struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2SBeautyNewCloth) Reset()         { *m = C2SBeautyNewCloth{} }
func (m *C2SBeautyNewCloth) String() string { return proto.CompactTextString(m) }
func (*C2SBeautyNewCloth) ProtoMessage()    {}
func (m *C2SBeautyNewCloth) GetId() uint16  { return PCK_C2SBeautyNewCloth_ID }

const PCK_S2CBeautyNewCloth_ID = 872 //
//
type S2CBeautyNewCloth struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Cloth []*BeautyCloth `protobuf:"bytes,2,rep,name=Cloth,proto3" json:"Cloth,omitempty"`
}

func (m *S2CBeautyNewCloth) Reset()         { *m = S2CBeautyNewCloth{} }
func (m *S2CBeautyNewCloth) String() string { return proto.CompactTextString(m) }
func (*S2CBeautyNewCloth) ProtoMessage()    {}
func (m *S2CBeautyNewCloth) GetId() uint16  { return PCK_S2CBeautyNewCloth_ID }

const PCK_C2SBeautyWearCloth_ID = 873 //
//
type C2SBeautyWearCloth struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2SBeautyWearCloth) Reset()         { *m = C2SBeautyWearCloth{} }
func (m *C2SBeautyWearCloth) String() string { return proto.CompactTextString(m) }
func (*C2SBeautyWearCloth) ProtoMessage()    {}
func (m *C2SBeautyWearCloth) GetId() uint16  { return PCK_C2SBeautyWearCloth_ID }

const PCK_S2CBeautyWearCloth_ID = 874 //
//
type S2CBeautyWearCloth struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Cloth []*BeautyCloth `protobuf:"bytes,2,rep,name=Cloth,proto3" json:"Cloth,omitempty"`
}

func (m *S2CBeautyWearCloth) Reset()         { *m = S2CBeautyWearCloth{} }
func (m *S2CBeautyWearCloth) String() string { return proto.CompactTextString(m) }
func (*S2CBeautyWearCloth) ProtoMessage()    {}
func (m *S2CBeautyWearCloth) GetId() uint16  { return PCK_S2CBeautyWearCloth_ID }

const PCK_C2SBeautyClothUp_ID = 875 //
//
type C2SBeautyClothUp struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2SBeautyClothUp) Reset()         { *m = C2SBeautyClothUp{} }
func (m *C2SBeautyClothUp) String() string { return proto.CompactTextString(m) }
func (*C2SBeautyClothUp) ProtoMessage()    {}
func (m *C2SBeautyClothUp) GetId() uint16  { return PCK_C2SBeautyClothUp_ID }

const PCK_S2CBeautyClothUp_ID = 876 //
//
type S2CBeautyClothUp struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Cloth []*BeautyCloth `protobuf:"bytes,2,rep,name=Cloth,proto3" json:"Cloth,omitempty"`
}

func (m *S2CBeautyClothUp) Reset()         { *m = S2CBeautyClothUp{} }
func (m *S2CBeautyClothUp) String() string { return proto.CompactTextString(m) }
func (*S2CBeautyClothUp) ProtoMessage()    {}
func (m *S2CBeautyClothUp) GetId() uint16  { return PCK_S2CBeautyClothUp_ID }

const PCK_C2SBeautyFlySwitch_ID = 852 //
//
type C2SBeautyFlySwitch struct {
}

func (m *C2SBeautyFlySwitch) Reset()         { *m = C2SBeautyFlySwitch{} }
func (m *C2SBeautyFlySwitch) String() string { return proto.CompactTextString(m) }
func (*C2SBeautyFlySwitch) ProtoMessage()    {}
func (m *C2SBeautyFlySwitch) GetId() uint16  { return PCK_C2SBeautyFlySwitch_ID }

const PCK_S2CBeautyFlySwitch_ID = 853 //
//
type S2CBeautyFlySwitch struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CBeautyFlySwitch) Reset()         { *m = S2CBeautyFlySwitch{} }
func (m *S2CBeautyFlySwitch) String() string { return proto.CompactTextString(m) }
func (*S2CBeautyFlySwitch) ProtoMessage()    {}
func (m *S2CBeautyFlySwitch) GetId() uint16  { return PCK_S2CBeautyFlySwitch_ID }

const PCK_C2SBeautyEquip_ID = 854 //
//
type C2SBeautyEquip struct {
	//
	Level int32 `protobuf:"varint,1,opt,name=Level,proto3" json:"Level,omitempty"`
}

func (m *C2SBeautyEquip) Reset()         { *m = C2SBeautyEquip{} }
func (m *C2SBeautyEquip) String() string { return proto.CompactTextString(m) }
func (*C2SBeautyEquip) ProtoMessage()    {}
func (m *C2SBeautyEquip) GetId() uint16  { return PCK_C2SBeautyEquip_ID }

const PCK_S2CBeautyEquip_ID = 855 //
//
type S2CBeautyEquip struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CBeautyEquip) Reset()         { *m = S2CBeautyEquip{} }
func (m *S2CBeautyEquip) String() string { return proto.CompactTextString(m) }
func (*S2CBeautyEquip) ProtoMessage()    {}
func (m *S2CBeautyEquip) GetId() uint16  { return PCK_S2CBeautyEquip_ID }

const PCK_C2SBeautyMagicRefresh_ID = 856 //
//
type C2SBeautyMagicRefresh struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	LockedItem []int32 `protobuf:"varint,2,rep,name=LockedItem,proto3" json:"LockedItem,omitempty"`
	//
	Type int32 `protobuf:"varint,3,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (m *C2SBeautyMagicRefresh) Reset()         { *m = C2SBeautyMagicRefresh{} }
func (m *C2SBeautyMagicRefresh) String() string { return proto.CompactTextString(m) }
func (*C2SBeautyMagicRefresh) ProtoMessage()    {}
func (m *C2SBeautyMagicRefresh) GetId() uint16  { return PCK_C2SBeautyMagicRefresh_ID }

const PCK_S2CBeautyMagicRefresh_ID = 857 //
//
type S2CBeautyMagicRefresh struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	TempItems []int32 `protobuf:"varint,3,rep,name=TempItems,proto3" json:"TempItems,omitempty"`
	//
	Times int32 `protobuf:"varint,4,opt,name=Times,proto3" json:"Times,omitempty"`
}

func (m *S2CBeautyMagicRefresh) Reset()         { *m = S2CBeautyMagicRefresh{} }
func (m *S2CBeautyMagicRefresh) String() string { return proto.CompactTextString(m) }
func (*S2CBeautyMagicRefresh) ProtoMessage()    {}
func (m *S2CBeautyMagicRefresh) GetId() uint16  { return PCK_S2CBeautyMagicRefresh_ID }

const PCK_C2SBeautyMagicChange_ID = 858 //
//
type C2SBeautyMagicChange struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2SBeautyMagicChange) Reset()         { *m = C2SBeautyMagicChange{} }
func (m *C2SBeautyMagicChange) String() string { return proto.CompactTextString(m) }
func (*C2SBeautyMagicChange) ProtoMessage()    {}
func (m *C2SBeautyMagicChange) GetId() uint16  { return PCK_C2SBeautyMagicChange_ID }

const PCK_S2CBeautyMagicChange_ID = 859 //
//
type S2CBeautyMagicChange struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Items []int32 `protobuf:"varint,3,rep,name=Items,proto3" json:"Items,omitempty"`
	//
	TempItems []int32 `protobuf:"varint,4,rep,name=TempItems,proto3" json:"TempItems,omitempty"`
}

func (m *S2CBeautyMagicChange) Reset()         { *m = S2CBeautyMagicChange{} }
func (m *S2CBeautyMagicChange) String() string { return proto.CompactTextString(m) }
func (*S2CBeautyMagicChange) ProtoMessage()    {}
func (m *S2CBeautyMagicChange) GetId() uint16  { return PCK_S2CBeautyMagicChange_ID }

//
type Kid struct {
	//
	P *Partner `protobuf:"bytes,1,opt,name=P,proto3" json:"P,omitempty"`
	//
	Good int32 `protobuf:"varint,2,opt,name=Good,proto3" json:"Good,omitempty"`
	//
	T1 []int32 `protobuf:"varint,3,rep,name=T1,proto3" json:"T1,omitempty"`
	//
	T2 []int32 `protobuf:"varint,4,rep,name=T2,proto3" json:"T2,omitempty"`
	//
	T3 []int32 `protobuf:"varint,5,rep,name=T3,proto3" json:"T3,omitempty"`
	//
	T4 []int32 `protobuf:"varint,6,rep,name=T4,proto3" json:"T4,omitempty"`
	//
	TS int32 `protobuf:"varint,7,opt,name=TS,proto3" json:"TS,omitempty"`
}

func (m *Kid) Reset()         { *m = Kid{} }
func (m *Kid) String() string { return proto.CompactTextString(m) }
func (*Kid) ProtoMessage()    {}

//
type Toy struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	L int32 `protobuf:"varint,2,opt,name=L,proto3" json:"L,omitempty"`
}

func (m *Toy) Reset()         { *m = Toy{} }
func (m *Toy) String() string { return proto.CompactTextString(m) }
func (*Toy) ProtoMessage()    {}

const PCK_S2CUserKid_ID = 880 //
//
type S2CUserKid struct {
	//
	P []*Kid `protobuf:"bytes,1,rep,name=P,proto3" json:"P,omitempty"`
}

func (m *S2CUserKid) Reset()         { *m = S2CUserKid{} }
func (m *S2CUserKid) String() string { return proto.CompactTextString(m) }
func (*S2CUserKid) ProtoMessage()    {}
func (m *S2CUserKid) GetId() uint16  { return PCK_S2CUserKid_ID }

const PCK_S2CUserToy_ID = 891 //
//
type S2CUserToy struct {
	//
	T []*Toy `protobuf:"bytes,1,rep,name=T,proto3" json:"T,omitempty"`
}

func (m *S2CUserToy) Reset()         { *m = S2CUserToy{} }
func (m *S2CUserToy) String() string { return proto.CompactTextString(m) }
func (*S2CUserToy) ProtoMessage()    {}
func (m *S2CUserToy) GetId() uint16  { return PCK_S2CUserToy_ID }

const PCK_C2SKidTalentUp_ID = 882 //
//
type C2SKidTalentUp struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Type int32 `protobuf:"varint,2,opt,name=Type,proto3" json:"Type,omitempty"`
	//
	Part int32 `protobuf:"varint,3,opt,name=Part,proto3" json:"Part,omitempty"`
}

func (m *C2SKidTalentUp) Reset()         { *m = C2SKidTalentUp{} }
func (m *C2SKidTalentUp) String() string { return proto.CompactTextString(m) }
func (*C2SKidTalentUp) ProtoMessage()    {}
func (m *C2SKidTalentUp) GetId() uint16  { return PCK_C2SKidTalentUp_ID }

const PCK_S2CKidTalentUp_ID = 881 //
//
type S2CKidTalentUp struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Type int32 `protobuf:"varint,3,opt,name=Type,proto3" json:"Type,omitempty"`
	//
	Part int32 `protobuf:"varint,4,opt,name=Part,proto3" json:"Part,omitempty"`
	//
	T []int32 `protobuf:"varint,5,rep,name=T,proto3" json:"T,omitempty"`
}

func (m *S2CKidTalentUp) Reset()         { *m = S2CKidTalentUp{} }
func (m *S2CKidTalentUp) String() string { return proto.CompactTextString(m) }
func (*S2CKidTalentUp) ProtoMessage()    {}
func (m *S2CKidTalentUp) GetId() uint16  { return PCK_S2CKidTalentUp_ID }

const PCK_C2SKidUseTalent_ID = 883 //
//
type C2SKidUseTalent struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Type int32 `protobuf:"varint,2,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (m *C2SKidUseTalent) Reset()         { *m = C2SKidUseTalent{} }
func (m *C2SKidUseTalent) String() string { return proto.CompactTextString(m) }
func (*C2SKidUseTalent) ProtoMessage()    {}
func (m *C2SKidUseTalent) GetId() uint16  { return PCK_C2SKidUseTalent_ID }

const PCK_S2CKidUseTalent_ID = 884 //
//
type S2CKidUseTalent struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Type int32 `protobuf:"varint,3,opt,name=Type,proto3" json:"Type,omitempty"`
	//
	TS int32 `protobuf:"varint,4,opt,name=TS,proto3" json:"TS,omitempty"`
	//
	OldTS int32 `protobuf:"varint,5,opt,name=OldTS,proto3" json:"OldTS,omitempty"`
}

func (m *S2CKidUseTalent) Reset()         { *m = S2CKidUseTalent{} }
func (m *S2CKidUseTalent) String() string { return proto.CompactTextString(m) }
func (*S2CKidUseTalent) ProtoMessage()    {}
func (m *S2CKidUseTalent) GetId() uint16  { return PCK_S2CKidUseTalent_ID }

const PCK_C2SKidGoodUp_ID = 885 //
//
type C2SKidGoodUp struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2SKidGoodUp) Reset()         { *m = C2SKidGoodUp{} }
func (m *C2SKidGoodUp) String() string { return proto.CompactTextString(m) }
func (*C2SKidGoodUp) ProtoMessage()    {}
func (m *C2SKidGoodUp) GetId() uint16  { return PCK_C2SKidGoodUp_ID }

const PCK_S2CKidGoodUp_ID = 886 //
//
type S2CKidGoodUp struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Good int32 `protobuf:"varint,3,opt,name=Good,proto3" json:"Good,omitempty"`
}

func (m *S2CKidGoodUp) Reset()         { *m = S2CKidGoodUp{} }
func (m *S2CKidGoodUp) String() string { return proto.CompactTextString(m) }
func (*S2CKidGoodUp) ProtoMessage()    {}
func (m *S2CKidGoodUp) GetId() uint16  { return PCK_S2CKidGoodUp_ID }

const PCK_C2SToyLevelUp_ID = 887 //
//
type C2SToyLevelUp struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2SToyLevelUp) Reset()         { *m = C2SToyLevelUp{} }
func (m *C2SToyLevelUp) String() string { return proto.CompactTextString(m) }
func (*C2SToyLevelUp) ProtoMessage()    {}
func (m *C2SToyLevelUp) GetId() uint16  { return PCK_C2SToyLevelUp_ID }

const PCK_S2CToyLevelUp_ID = 888 //
//
type S2CToyLevelUp struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	L int32 `protobuf:"varint,3,opt,name=L,proto3" json:"L,omitempty"`
}

func (m *S2CToyLevelUp) Reset()         { *m = S2CToyLevelUp{} }
func (m *S2CToyLevelUp) String() string { return proto.CompactTextString(m) }
func (*S2CToyLevelUp) ProtoMessage()    {}
func (m *S2CToyLevelUp) GetId() uint16  { return PCK_S2CToyLevelUp_ID }

const PCK_C2SToyOneKeyLevelUp_ID = 889 //
//
type C2SToyOneKeyLevelUp struct {
	//
	Quality int32 `protobuf:"varint,1,opt,name=Quality,proto3" json:"Quality,omitempty"`
}

func (m *C2SToyOneKeyLevelUp) Reset()         { *m = C2SToyOneKeyLevelUp{} }
func (m *C2SToyOneKeyLevelUp) String() string { return proto.CompactTextString(m) }
func (*C2SToyOneKeyLevelUp) ProtoMessage()    {}
func (m *C2SToyOneKeyLevelUp) GetId() uint16  { return PCK_C2SToyOneKeyLevelUp_ID }

const PCK_S2CToyOneKeyLevelUp_ID = 890 //
//
type S2CToyOneKeyLevelUp struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Quality int32 `protobuf:"varint,2,opt,name=Quality,proto3" json:"Quality,omitempty"`
	//
	Toys []*Toy `protobuf:"bytes,3,rep,name=Toys,proto3" json:"Toys,omitempty"`
}

func (m *S2CToyOneKeyLevelUp) Reset()         { *m = S2CToyOneKeyLevelUp{} }
func (m *S2CToyOneKeyLevelUp) String() string { return proto.CompactTextString(m) }
func (*S2CToyOneKeyLevelUp) ProtoMessage()    {}
func (m *S2CToyOneKeyLevelUp) GetId() uint16  { return PCK_S2CToyOneKeyLevelUp_ID }

//
type SkyGod struct {
	//
	P *Partner `protobuf:"bytes,1,opt,name=P,proto3" json:"P,omitempty"`
}

func (m *SkyGod) Reset()         { *m = SkyGod{} }
func (m *SkyGod) String() string { return proto.CompactTextString(m) }
func (*SkyGod) ProtoMessage()    {}

//
type S2CUserSkyGod struct {
	//
	P []*SkyGod `protobuf:"bytes,1,rep,name=P,proto3" json:"P,omitempty"`
}

func (m *S2CUserSkyGod) Reset()         { *m = S2CUserSkyGod{} }
func (m *S2CUserSkyGod) String() string { return proto.CompactTextString(m) }
func (*S2CUserSkyGod) ProtoMessage()    {}

//
type Therion struct {
	//
	P *Partner `protobuf:"bytes,1,opt,name=P,proto3" json:"P,omitempty"`
}

func (m *Therion) Reset()         { *m = Therion{} }
func (m *Therion) String() string { return proto.CompactTextString(m) }
func (*Therion) ProtoMessage()    {}

const PCK_S2CUserTherion_ID = 918 //
//
type S2CUserTherion struct {
	//
	P []*Therion `protobuf:"bytes,1,rep,name=P,proto3" json:"P,omitempty"`
}

func (m *S2CUserTherion) Reset()         { *m = S2CUserTherion{} }
func (m *S2CUserTherion) String() string { return proto.CompactTextString(m) }
func (*S2CUserTherion) ProtoMessage()    {}
func (m *S2CUserTherion) GetId() uint16  { return PCK_S2CUserTherion_ID }

//
type MasterPupilList struct {
	//
	U int32 `protobuf:"varint,1,opt,name=U,proto3" json:"U,omitempty"`
	//
	N string `protobuf:"bytes,2,opt,name=N,proto3" json:"N,omitempty"`
	//
	L int32 `protobuf:"varint,3,opt,name=L,proto3" json:"L,omitempty"`
	//
	OnlineState int32 `protobuf:"varint,4,opt,name=OnlineState,proto3" json:"OnlineState,omitempty"`
	//
	IsMaster int32 `protobuf:"varint,5,opt,name=IsMaster,proto3" json:"IsMaster,omitempty"`
	//
	IsGiveExp int32 `protobuf:"varint,6,opt,name=IsGiveExp,proto3" json:"IsGiveExp,omitempty"`
	//
	RoleId int32 `protobuf:"varint,7,opt,name=RoleId,proto3" json:"RoleId,omitempty"`
}

func (m *MasterPupilList) Reset()         { *m = MasterPupilList{} }
func (m *MasterPupilList) String() string { return proto.CompactTextString(m) }
func (*MasterPupilList) ProtoMessage()    {}

//
type DbMasterInfo struct {
	//
	U int32 `protobuf:"varint,1,opt,name=U,proto3" json:"U,omitempty"`
	//
	IsMaster int32 `protobuf:"varint,2,opt,name=IsMaster,proto3" json:"IsMaster,omitempty"`
	//
	IsGiveExp int32 `protobuf:"varint,3,opt,name=IsGiveExp,proto3" json:"IsGiveExp,omitempty"`
}

func (m *DbMasterInfo) Reset()         { *m = DbMasterInfo{} }
func (m *DbMasterInfo) String() string { return proto.CompactTextString(m) }
func (*DbMasterInfo) ProtoMessage()    {}

//
type DbMasterData struct {
	//
	JoinTime int64 `protobuf:"varint,1,opt,name=JoinTime,proto3" json:"JoinTime,omitempty"`
	//
	MasterList []*DbMasterInfo `protobuf:"bytes,2,rep,name=MasterList,proto3" json:"MasterList,omitempty"`
	//
	InviteList []*MasterArray `protobuf:"bytes,3,rep,name=InviteList,proto3" json:"InviteList,omitempty"`
}

func (m *DbMasterData) Reset()         { *m = DbMasterData{} }
func (m *DbMasterData) String() string { return proto.CompactTextString(m) }
func (*DbMasterData) ProtoMessage()    {}

const PCK_S2CMasterMessage_ID = 350 //
//
type S2CMasterMessage struct {
	//
	OutTime int32 `protobuf:"varint,1,opt,name=OutTime,proto3" json:"OutTime,omitempty"`
	//
	List []*MasterPupilList `protobuf:"bytes,2,rep,name=List,proto3" json:"List,omitempty"`
	//
	InviteList []*MasterArray `protobuf:"bytes,3,rep,name=InviteList,proto3" json:"InviteList,omitempty"`
}

func (m *S2CMasterMessage) Reset()         { *m = S2CMasterMessage{} }
func (m *S2CMasterMessage) String() string { return proto.CompactTextString(m) }
func (*S2CMasterMessage) ProtoMessage()    {}
func (m *S2CMasterMessage) GetId() uint16  { return PCK_S2CMasterMessage_ID }

const PCK_C2SSendMasterNotice_ID = 351 //
//
type C2SSendMasterNotice struct {
}

func (m *C2SSendMasterNotice) Reset()         { *m = C2SSendMasterNotice{} }
func (m *C2SSendMasterNotice) String() string { return proto.CompactTextString(m) }
func (*C2SSendMasterNotice) ProtoMessage()    {}
func (m *C2SSendMasterNotice) GetId() uint16  { return PCK_C2SSendMasterNotice_ID }

const PCK_S2CSendMasterNotice_ID = 352 //
//
type S2CSendMasterNotice struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CSendMasterNotice) Reset()         { *m = S2CSendMasterNotice{} }
func (m *S2CSendMasterNotice) String() string { return proto.CompactTextString(m) }
func (*S2CSendMasterNotice) ProtoMessage()    {}
func (m *S2CSendMasterNotice) GetId() uint16  { return PCK_S2CSendMasterNotice_ID }

//
type MasterArray struct {
	//
	U int32 `protobuf:"varint,1,opt,name=U,proto3" json:"U,omitempty"`
	//
	N string `protobuf:"bytes,2,opt,name=N,proto3" json:"N,omitempty"`
	//
	L int32 `protobuf:"varint,3,opt,name=L,proto3" json:"L,omitempty"`
	//
	D int32 `protobuf:"varint,4,opt,name=D,proto3" json:"D,omitempty"`
	//
	S int32 `protobuf:"varint,5,opt,name=S,proto3" json:"S,omitempty"`
	//
	T int64 `protobuf:"varint,6,opt,name=T,proto3" json:"T,omitempty"`
	//
	DeleteTime int64 `protobuf:"varint,7,opt,name=DeleteTime,proto3" json:"DeleteTime,omitempty"`
}

func (m *MasterArray) Reset()         { *m = MasterArray{} }
func (m *MasterArray) String() string { return proto.CompactTextString(m) }
func (*MasterArray) ProtoMessage()    {}

const PCK_S2CMasterAds_ID = 353 //
//
type S2CMasterAds struct {
	//
	PupilList []*MasterArray `protobuf:"bytes,1,rep,name=PupilList,proto3" json:"PupilList,omitempty"`
}

func (m *S2CMasterAds) Reset()         { *m = S2CMasterAds{} }
func (m *S2CMasterAds) String() string { return proto.CompactTextString(m) }
func (*S2CMasterAds) ProtoMessage()    {}
func (m *S2CMasterAds) GetId() uint16  { return PCK_S2CMasterAds_ID }

const PCK_C2SMasterInvite_ID = 354 //
//
type C2SMasterInvite struct {
	//
	U int32 `protobuf:"varint,1,opt,name=U,proto3" json:"U,omitempty"`
}

func (m *C2SMasterInvite) Reset()         { *m = C2SMasterInvite{} }
func (m *C2SMasterInvite) String() string { return proto.CompactTextString(m) }
func (*C2SMasterInvite) ProtoMessage()    {}
func (m *C2SMasterInvite) GetId() uint16  { return PCK_C2SMasterInvite_ID }

const PCK_S2CMasterInvite_ID = 355 //
//
type S2CMasterInvite struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CMasterInvite) Reset()         { *m = S2CMasterInvite{} }
func (m *S2CMasterInvite) String() string { return proto.CompactTextString(m) }
func (*S2CMasterInvite) ProtoMessage()    {}
func (m *S2CMasterInvite) GetId() uint16  { return PCK_S2CMasterInvite_ID }

const PCK_S2CInvitePupil_ID = 356 //
//
type S2CInvitePupil struct {
	//
	U int32 `protobuf:"varint,1,opt,name=U,proto3" json:"U,omitempty"`
	//
	N string `protobuf:"bytes,2,opt,name=N,proto3" json:"N,omitempty"`
}

func (m *S2CInvitePupil) Reset()         { *m = S2CInvitePupil{} }
func (m *S2CInvitePupil) String() string { return proto.CompactTextString(m) }
func (*S2CInvitePupil) ProtoMessage()    {}
func (m *S2CInvitePupil) GetId() uint16  { return PCK_S2CInvitePupil_ID }

const PCK_C2SHandleMasterInvite_ID = 357 //
//
type C2SHandleMasterInvite struct {
	//
	U int32 `protobuf:"varint,1,opt,name=U,proto3" json:"U,omitempty"`
	//
	Y int32 `protobuf:"varint,2,opt,name=Y,proto3" json:"Y,omitempty"`
}

func (m *C2SHandleMasterInvite) Reset()         { *m = C2SHandleMasterInvite{} }
func (m *C2SHandleMasterInvite) String() string { return proto.CompactTextString(m) }
func (*C2SHandleMasterInvite) ProtoMessage()    {}
func (m *C2SHandleMasterInvite) GetId() uint16  { return PCK_C2SHandleMasterInvite_ID }

const PCK_S2CHandleMasterInvite_ID = 358 //
//
type S2CHandleMasterInvite struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CHandleMasterInvite) Reset()         { *m = S2CHandleMasterInvite{} }
func (m *S2CHandleMasterInvite) String() string { return proto.CompactTextString(m) }
func (*S2CHandleMasterInvite) ProtoMessage()    {}
func (m *S2CHandleMasterInvite) GetId() uint16  { return PCK_S2CHandleMasterInvite_ID }

const PCK_C2SMasterGiveExp_ID = 359 //
//
type C2SMasterGiveExp struct {
	//
	U int32 `protobuf:"varint,1,opt,name=U,proto3" json:"U,omitempty"`
}

func (m *C2SMasterGiveExp) Reset()         { *m = C2SMasterGiveExp{} }
func (m *C2SMasterGiveExp) String() string { return proto.CompactTextString(m) }
func (*C2SMasterGiveExp) ProtoMessage()    {}
func (m *C2SMasterGiveExp) GetId() uint16  { return PCK_C2SMasterGiveExp_ID }

const PCK_S2CMasterGiveExp_ID = 360 //
//
type S2CMasterGiveExp struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CMasterGiveExp) Reset()         { *m = S2CMasterGiveExp{} }
func (m *S2CMasterGiveExp) String() string { return proto.CompactTextString(m) }
func (*S2CMasterGiveExp) ProtoMessage()    {}
func (m *S2CMasterGiveExp) GetId() uint16  { return PCK_S2CMasterGiveExp_ID }

const PCK_C2SGetMasterExp_ID = 361 //
//
type C2SGetMasterExp struct {
}

func (m *C2SGetMasterExp) Reset()         { *m = C2SGetMasterExp{} }
func (m *C2SGetMasterExp) String() string { return proto.CompactTextString(m) }
func (*C2SGetMasterExp) ProtoMessage()    {}
func (m *C2SGetMasterExp) GetId() uint16  { return PCK_C2SGetMasterExp_ID }

const PCK_S2CGetMasterExp_ID = 362 //
//
type S2CGetMasterExp struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CGetMasterExp) Reset()         { *m = S2CGetMasterExp{} }
func (m *S2CGetMasterExp) String() string { return proto.CompactTextString(m) }
func (*S2CGetMasterExp) ProtoMessage()    {}
func (m *S2CGetMasterExp) GetId() uint16  { return PCK_S2CGetMasterExp_ID }

const PCK_C2SDeletePupil_ID = 363 //
//
type C2SDeletePupil struct {
	//
	U int32 `protobuf:"varint,1,opt,name=U,proto3" json:"U,omitempty"`
}

func (m *C2SDeletePupil) Reset()         { *m = C2SDeletePupil{} }
func (m *C2SDeletePupil) String() string { return proto.CompactTextString(m) }
func (*C2SDeletePupil) ProtoMessage()    {}
func (m *C2SDeletePupil) GetId() uint16  { return PCK_C2SDeletePupil_ID }

const PCK_S2CDeletePupil_ID = 364 //
//
type S2CDeletePupil struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CDeletePupil) Reset()         { *m = S2CDeletePupil{} }
func (m *S2CDeletePupil) String() string { return proto.CompactTextString(m) }
func (*S2CDeletePupil) ProtoMessage()    {}
func (m *S2CDeletePupil) GetId() uint16  { return PCK_S2CDeletePupil_ID }

const PCK_C2SDeleteMaster_ID = 365 //
//
type C2SDeleteMaster struct {
	//
	T int32 `protobuf:"varint,1,opt,name=T,proto3" json:"T,omitempty"`
}

func (m *C2SDeleteMaster) Reset()         { *m = C2SDeleteMaster{} }
func (m *C2SDeleteMaster) String() string { return proto.CompactTextString(m) }
func (*C2SDeleteMaster) ProtoMessage()    {}
func (m *C2SDeleteMaster) GetId() uint16  { return PCK_C2SDeleteMaster_ID }

const PCK_S2CDeleteMaster_ID = 366 //
//
type S2CDeleteMaster struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CDeleteMaster) Reset()         { *m = S2CDeleteMaster{} }
func (m *S2CDeleteMaster) String() string { return proto.CompactTextString(m) }
func (*S2CDeleteMaster) ProtoMessage()    {}
func (m *S2CDeleteMaster) GetId() uint16  { return PCK_S2CDeleteMaster_ID }

//
type MasterTask struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	State int32 `protobuf:"varint,2,opt,name=State,proto3" json:"State,omitempty"`
	//
	Count int32 `protobuf:"varint,3,opt,name=Count,proto3" json:"Count,omitempty"`
}

func (m *MasterTask) Reset()         { *m = MasterTask{} }
func (m *MasterTask) String() string { return proto.CompactTextString(m) }
func (*MasterTask) ProtoMessage()    {}

const PCK_C2SGetPupilTask_ID = 367 //
//
type C2SGetPupilTask struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2SGetPupilTask) Reset()         { *m = C2SGetPupilTask{} }
func (m *C2SGetPupilTask) String() string { return proto.CompactTextString(m) }
func (*C2SGetPupilTask) ProtoMessage()    {}
func (m *C2SGetPupilTask) GetId() uint16  { return PCK_C2SGetPupilTask_ID }

const PCK_S2CGetPupilTask_ID = 368 //
//
type S2CGetPupilTask struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Task []*MasterTask `protobuf:"bytes,2,rep,name=Task,proto3" json:"Task,omitempty"`
}

func (m *S2CGetPupilTask) Reset()         { *m = S2CGetPupilTask{} }
func (m *S2CGetPupilTask) String() string { return proto.CompactTextString(m) }
func (*S2CGetPupilTask) ProtoMessage()    {}
func (m *S2CGetPupilTask) GetId() uint16  { return PCK_S2CGetPupilTask_ID }

//
type RoleMarryInfo struct {
	//
	MarryList []*DbMarryInfo `protobuf:"bytes,1,rep,name=MarryList,proto3" json:"MarryList,omitempty"`
}

func (m *RoleMarryInfo) Reset()         { *m = RoleMarryInfo{} }
func (m *RoleMarryInfo) String() string { return proto.CompactTextString(m) }
func (*RoleMarryInfo) ProtoMessage()    {}

//
type DbMarryInfo struct {
	//
	U int32 `protobuf:"varint,1,opt,name=U,proto3" json:"U,omitempty"`
	//
	T int64 `protobuf:"varint,2,opt,name=T,proto3" json:"T,omitempty"`
	//
	H int32 `protobuf:"varint,3,opt,name=H,proto3" json:"H,omitempty"`
	//
	HL int32 `protobuf:"varint,4,opt,name=HL,proto3" json:"HL,omitempty"`
	//
	MV int32 `protobuf:"varint,5,opt,name=MV,proto3" json:"MV,omitempty"`
	//
	HM int32 `protobuf:"varint,6,opt,name=HM,proto3" json:"HM,omitempty"`
	//
	HouseUpList []*HouseUpList `protobuf:"bytes,7,rep,name=HouseUpList,proto3" json:"HouseUpList,omitempty"`
	//
	Pos int32 `protobuf:"varint,8,opt,name=Pos,proto3" json:"Pos,omitempty"`
	//
	InviteList []*MarryInvite `protobuf:"bytes,9,rep,name=InviteList,proto3" json:"InviteList,omitempty"`
	//
	CanAddHouseValue int64 `protobuf:"varint,10,opt,name=CanAddHouseValue,proto3" json:"CanAddHouseValue,omitempty"`
}

func (m *DbMarryInfo) Reset()         { *m = DbMarryInfo{} }
func (m *DbMarryInfo) String() string { return proto.CompactTextString(m) }
func (*DbMarryInfo) ProtoMessage()    {}

//
type MarryInvite struct {
	//
	Pos int32 `protobuf:"varint,1,opt,name=Pos,proto3" json:"Pos,omitempty"`
	//
	HouseId int32 `protobuf:"varint,2,opt,name=HouseId,proto3" json:"HouseId,omitempty"`
	//
	InviteUserId int32 `protobuf:"varint,3,opt,name=InviteUserId,proto3" json:"InviteUserId,omitempty"`
}

func (m *MarryInvite) Reset()         { *m = MarryInvite{} }
func (m *MarryInvite) String() string { return proto.CompactTextString(m) }
func (*MarryInvite) ProtoMessage()    {}

const PCK_C2SGetMarryList_ID = 371 //
//
type C2SGetMarryList struct {
}

func (m *C2SGetMarryList) Reset()         { *m = C2SGetMarryList{} }
func (m *C2SGetMarryList) String() string { return proto.CompactTextString(m) }
func (*C2SGetMarryList) ProtoMessage()    {}
func (m *C2SGetMarryList) GetId() uint16  { return PCK_C2SGetMarryList_ID }

//
type MarryList struct {
	//
	U int32 `protobuf:"varint,1,opt,name=U,proto3" json:"U,omitempty"`
	//
	N string `protobuf:"bytes,2,opt,name=N,proto3" json:"N,omitempty"`
	//
	L int32 `protobuf:"varint,3,opt,name=L,proto3" json:"L,omitempty"`
	//
	F int64 `protobuf:"varint,4,opt,name=F,proto3" json:"F,omitempty"`
	//
	R int32 `protobuf:"varint,5,opt,name=R,proto3" json:"R,omitempty"`
}

func (m *MarryList) Reset()         { *m = MarryList{} }
func (m *MarryList) String() string { return proto.CompactTextString(m) }
func (*MarryList) ProtoMessage()    {}

const PCK_S2CGetMarryList_ID = 372 //
//
type S2CGetMarryList struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	List []*MarryList `protobuf:"bytes,2,rep,name=List,proto3" json:"List,omitempty"`
}

func (m *S2CGetMarryList) Reset()         { *m = S2CGetMarryList{} }
func (m *S2CGetMarryList) String() string { return proto.CompactTextString(m) }
func (*S2CGetMarryList) ProtoMessage()    {}
func (m *S2CGetMarryList) GetId() uint16  { return PCK_S2CGetMarryList_ID }

const PCK_C2SGetMarry_ID = 373 //
//
type C2SGetMarry struct {
	//
	I int32 `protobuf:"varint,1,opt,name=I,proto3" json:"I,omitempty"`
	//
	P int32 `protobuf:"varint,2,opt,name=P,proto3" json:"P,omitempty"`
	//
	U int32 `protobuf:"varint,3,opt,name=U,proto3" json:"U,omitempty"`
}

func (m *C2SGetMarry) Reset()         { *m = C2SGetMarry{} }
func (m *C2SGetMarry) String() string { return proto.CompactTextString(m) }
func (*C2SGetMarry) ProtoMessage()    {}
func (m *C2SGetMarry) GetId() uint16  { return PCK_C2SGetMarry_ID }

const PCK_S2CGetMarry_ID = 374 //
//
type S2CGetMarry struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CGetMarry) Reset()         { *m = S2CGetMarry{} }
func (m *S2CGetMarry) String() string { return proto.CompactTextString(m) }
func (*S2CGetMarry) ProtoMessage()    {}
func (m *S2CGetMarry) GetId() uint16  { return PCK_S2CGetMarry_ID }

const PCK_S2CSendMarry_ID = 375 //
//
type S2CSendMarry struct {
	//
	I int32 `protobuf:"varint,1,opt,name=I,proto3" json:"I,omitempty"`
	//
	N string `protobuf:"bytes,2,opt,name=N,proto3" json:"N,omitempty"`
	//
	L int64 `protobuf:"varint,3,opt,name=L,proto3" json:"L,omitempty"`
	//
	F int64 `protobuf:"varint,4,opt,name=F,proto3" json:"F,omitempty"`
	//
	R int32 `protobuf:"varint,5,opt,name=R,proto3" json:"R,omitempty"`
	//
	M int32 `protobuf:"varint,6,opt,name=M,proto3" json:"M,omitempty"`
	//
	U int32 `protobuf:"varint,7,opt,name=U,proto3" json:"U,omitempty"`
}

func (m *S2CSendMarry) Reset()         { *m = S2CSendMarry{} }
func (m *S2CSendMarry) String() string { return proto.CompactTextString(m) }
func (*S2CSendMarry) ProtoMessage()    {}
func (m *S2CSendMarry) GetId() uint16  { return PCK_S2CSendMarry_ID }

const PCK_C2SHandelMarry_ID = 391 //
//
type C2SHandelMarry struct {
	//
	U int32 `protobuf:"varint,1,opt,name=U,proto3" json:"U,omitempty"`
	//
	Y int32 `protobuf:"varint,2,opt,name=Y,proto3" json:"Y,omitempty"`
}

func (m *C2SHandelMarry) Reset()         { *m = C2SHandelMarry{} }
func (m *C2SHandelMarry) String() string { return proto.CompactTextString(m) }
func (*C2SHandelMarry) ProtoMessage()    {}
func (m *C2SHandelMarry) GetId() uint16  { return PCK_C2SHandelMarry_ID }

const PCK_S2CHandelMarry_ID = 392 //
//
type S2CHandelMarry struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CHandelMarry) Reset()         { *m = S2CHandelMarry{} }
func (m *S2CHandelMarry) String() string { return proto.CompactTextString(m) }
func (*S2CHandelMarry) ProtoMessage()    {}
func (m *S2CHandelMarry) GetId() uint16  { return PCK_S2CHandelMarry_ID }

const PCK_C2SSendFlower_ID = 377 //
//
type C2SSendFlower struct {
	//
	F int32 `protobuf:"varint,1,opt,name=F,proto3" json:"F,omitempty"`
	//
	N int32 `protobuf:"varint,2,opt,name=N,proto3" json:"N,omitempty"`
	//
	Bn int32 `protobuf:"varint,3,opt,name=Bn,proto3" json:"Bn,omitempty"`
}

func (m *C2SSendFlower) Reset()         { *m = C2SSendFlower{} }
func (m *C2SSendFlower) String() string { return proto.CompactTextString(m) }
func (*C2SSendFlower) ProtoMessage()    {}
func (m *C2SSendFlower) GetId() uint16  { return PCK_C2SSendFlower_ID }

const PCK_S2CSendFlower_ID = 378 //
//
type S2CSendFlower struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Mv int32 `protobuf:"varint,2,opt,name=Mv,proto3" json:"Mv,omitempty"`
}

func (m *S2CSendFlower) Reset()         { *m = S2CSendFlower{} }
func (m *S2CSendFlower) String() string { return proto.CompactTextString(m) }
func (*S2CSendFlower) ProtoMessage()    {}
func (m *S2CSendFlower) GetId() uint16  { return PCK_S2CSendFlower_ID }

const PCK_S2CGetFlower_ID = 379 //
//
type S2CGetFlower struct {
	//
	N string `protobuf:"bytes,1,opt,name=N,proto3" json:"N,omitempty"`
	//
	F int32 `protobuf:"varint,2,opt,name=F,proto3" json:"F,omitempty"`
	//
	NM int32 `protobuf:"varint,3,opt,name=NM,proto3" json:"NM,omitempty"`
	//
	MV int32 `protobuf:"varint,4,opt,name=MV,proto3" json:"MV,omitempty"`
}

func (m *S2CGetFlower) Reset()         { *m = S2CGetFlower{} }
func (m *S2CGetFlower) String() string { return proto.CompactTextString(m) }
func (*S2CGetFlower) ProtoMessage()    {}
func (m *S2CGetFlower) GetId() uint16  { return PCK_S2CGetFlower_ID }

const PCK_S2CMarryStatus_ID = 370 //
//
type S2CMarryStatus struct {
	//
	U int32 `protobuf:"varint,1,opt,name=U,proto3" json:"U,omitempty"`
	//
	N string `protobuf:"bytes,2,opt,name=N,proto3" json:"N,omitempty"`
	//
	R int32 `protobuf:"varint,3,opt,name=R,proto3" json:"R,omitempty"`
	//
	D int64 `protobuf:"varint,4,opt,name=D,proto3" json:"D,omitempty"`
	//
	M int32 `protobuf:"varint,5,opt,name=M,proto3" json:"M,omitempty"`
	//
	HL int32 `protobuf:"varint,6,opt,name=HL,proto3" json:"HL,omitempty"`
	//
	HS int32 `protobuf:"varint,7,opt,name=HS,proto3" json:"HS,omitempty"`
	//
	MV int32 `protobuf:"varint,8,opt,name=MV,proto3" json:"MV,omitempty"`
	//
	Pos int32 `protobuf:"varint,9,opt,name=Pos,proto3" json:"Pos,omitempty"`
	//
	LT int32 `protobuf:"varint,10,opt,name=LT,proto3" json:"LT,omitempty"`
	//
	List []*HouseUpList `protobuf:"bytes,11,rep,name=List,proto3" json:"List,omitempty"`
}

func (m *S2CMarryStatus) Reset()         { *m = S2CMarryStatus{} }
func (m *S2CMarryStatus) String() string { return proto.CompactTextString(m) }
func (*S2CMarryStatus) ProtoMessage()    {}
func (m *S2CMarryStatus) GetId() uint16  { return PCK_S2CMarryStatus_ID }

const PCK_C2SMarryUpdate_ID = 381 //
//
type C2SMarryUpdate struct {
}

func (m *C2SMarryUpdate) Reset()         { *m = C2SMarryUpdate{} }
func (m *C2SMarryUpdate) String() string { return proto.CompactTextString(m) }
func (*C2SMarryUpdate) ProtoMessage()    {}
func (m *C2SMarryUpdate) GetId() uint16  { return PCK_C2SMarryUpdate_ID }

const PCK_S2CMarryUpdate_ID = 382 //
//
type S2CMarryUpdate struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CMarryUpdate) Reset()         { *m = S2CMarryUpdate{} }
func (m *S2CMarryUpdate) String() string { return proto.CompactTextString(m) }
func (*S2CMarryUpdate) ProtoMessage()    {}
func (m *S2CMarryUpdate) GetId() uint16  { return PCK_S2CMarryUpdate_ID }

const PCK_C2SHouseUpdate_ID = 383 //
//
type C2SHouseUpdate struct {
	//
	M int32 `protobuf:"varint,1,opt,name=M,proto3" json:"M,omitempty"`
}

func (m *C2SHouseUpdate) Reset()         { *m = C2SHouseUpdate{} }
func (m *C2SHouseUpdate) String() string { return proto.CompactTextString(m) }
func (*C2SHouseUpdate) ProtoMessage()    {}
func (m *C2SHouseUpdate) GetId() uint16  { return PCK_C2SHouseUpdate_ID }

const PCK_S2CHouseUpdate_ID = 384 //
//
type S2CHouseUpdate struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CHouseUpdate) Reset()         { *m = S2CHouseUpdate{} }
func (m *S2CHouseUpdate) String() string { return proto.CompactTextString(m) }
func (*S2CHouseUpdate) ProtoMessage()    {}
func (m *S2CHouseUpdate) GetId() uint16  { return PCK_S2CHouseUpdate_ID }

//
type HouseUpList struct {
	//
	T int64 `protobuf:"varint,1,opt,name=T,proto3" json:"T,omitempty"`
	//
	V int32 `protobuf:"varint,2,opt,name=V,proto3" json:"V,omitempty"`
}

func (m *HouseUpList) Reset()         { *m = HouseUpList{} }
func (m *HouseUpList) String() string { return proto.CompactTextString(m) }
func (*HouseUpList) ProtoMessage()    {}

const PCK_C2SGetUpdate_ID = 386 //
//
type C2SGetUpdate struct {
}

func (m *C2SGetUpdate) Reset()         { *m = C2SGetUpdate{} }
func (m *C2SGetUpdate) String() string { return proto.CompactTextString(m) }
func (*C2SGetUpdate) ProtoMessage()    {}
func (m *C2SGetUpdate) GetId() uint16  { return PCK_C2SGetUpdate_ID }

const PCK_S2CGetUpdate_ID = 387 //
//
type S2CGetUpdate struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CGetUpdate) Reset()         { *m = S2CGetUpdate{} }
func (m *S2CGetUpdate) String() string { return proto.CompactTextString(m) }
func (*S2CGetUpdate) ProtoMessage()    {}
func (m *S2CGetUpdate) GetId() uint16  { return PCK_S2CGetUpdate_ID }

const PCK_S2CNewMarry_ID = 388 //
//
type S2CNewMarry struct {
	//
	U int32 `protobuf:"varint,1,opt,name=U,proto3" json:"U,omitempty"`
	//
	RM int32 `protobuf:"varint,2,opt,name=RM,proto3" json:"RM,omitempty"`
	//
	RF int32 `protobuf:"varint,3,opt,name=RF,proto3" json:"RF,omitempty"`
	//
	HN string `protobuf:"bytes,4,opt,name=HN,proto3" json:"HN,omitempty"`
	//
	WN string `protobuf:"bytes,5,opt,name=WN,proto3" json:"WN,omitempty"`
}

func (m *S2CNewMarry) Reset()         { *m = S2CNewMarry{} }
func (m *S2CNewMarry) String() string { return proto.CompactTextString(m) }
func (*S2CNewMarry) ProtoMessage()    {}
func (m *S2CNewMarry) GetId() uint16  { return PCK_S2CNewMarry_ID }

const PCK_C2SSendMarryGift_ID = 389 //
//
type C2SSendMarryGift struct {
	//
	I int32 `protobuf:"varint,1,opt,name=I,proto3" json:"I,omitempty"`
	//
	U int32 `protobuf:"varint,2,opt,name=U,proto3" json:"U,omitempty"`
}

func (m *C2SSendMarryGift) Reset()         { *m = C2SSendMarryGift{} }
func (m *C2SSendMarryGift) String() string { return proto.CompactTextString(m) }
func (*C2SSendMarryGift) ProtoMessage()    {}
func (m *C2SSendMarryGift) GetId() uint16  { return PCK_C2SSendMarryGift_ID }

const PCK_S2CSendMarryGift_ID = 390 //
//
type S2CSendMarryGift struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CSendMarryGift) Reset()         { *m = S2CSendMarryGift{} }
func (m *S2CSendMarryGift) String() string { return proto.CompactTextString(m) }
func (*S2CSendMarryGift) ProtoMessage()    {}
func (m *S2CSendMarryGift) GetId() uint16  { return PCK_S2CSendMarryGift_ID }

const PCK_C2SDeleteWife_ID = 393 //
//
type C2SDeleteWife struct {
}

func (m *C2SDeleteWife) Reset()         { *m = C2SDeleteWife{} }
func (m *C2SDeleteWife) String() string { return proto.CompactTextString(m) }
func (*C2SDeleteWife) ProtoMessage()    {}
func (m *C2SDeleteWife) GetId() uint16  { return PCK_C2SDeleteWife_ID }

const PCK_S2CDeleteWife_ID = 394 //
//
type S2CDeleteWife struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CDeleteWife) Reset()         { *m = S2CDeleteWife{} }
func (m *S2CDeleteWife) String() string { return proto.CompactTextString(m) }
func (*S2CDeleteWife) ProtoMessage()    {}
func (m *S2CDeleteWife) GetId() uint16  { return PCK_S2CDeleteWife_ID }

const PCK_S2CMateOnline_ID = 395 //
//
type S2CMateOnline struct {
	//
	R int32 `protobuf:"varint,1,opt,name=R,proto3" json:"R,omitempty"`
}

func (m *S2CMateOnline) Reset()         { *m = S2CMateOnline{} }
func (m *S2CMateOnline) String() string { return proto.CompactTextString(m) }
func (*S2CMateOnline) ProtoMessage()    {}
func (m *S2CMateOnline) GetId() uint16  { return PCK_S2CMateOnline_ID }

const PCK_S2CQuizAsk_ID = 450 //
//
type S2CQuizAsk struct {
	//
	Qid int32 `protobuf:"varint,1,opt,name=Qid,proto3" json:"Qid,omitempty"`
	//
	Num int32 `protobuf:"varint,2,opt,name=Num,proto3" json:"Num,omitempty"`
	//
	Aid []int32 `protobuf:"varint,3,rep,name=Aid,proto3" json:"Aid,omitempty"`
	//
	AserTime int32 `protobuf:"varint,4,opt,name=AserTime,proto3" json:"AserTime,omitempty"`
	//
	NextTime int32 `protobuf:"varint,5,opt,name=NextTime,proto3" json:"NextTime,omitempty"`
}

func (m *S2CQuizAsk) Reset()         { *m = S2CQuizAsk{} }
func (m *S2CQuizAsk) String() string { return proto.CompactTextString(m) }
func (*S2CQuizAsk) ProtoMessage()    {}
func (m *S2CQuizAsk) GetId() uint16  { return PCK_S2CQuizAsk_ID }

const PCK_S2CQuizRank_ID = 454 //
//
type S2CQuizRank struct {
	//
	RankList []*QuizRank `protobuf:"bytes,1,rep,name=RankList,proto3" json:"RankList,omitempty"`
	//
	Mn int32 `protobuf:"varint,2,opt,name=Mn,proto3" json:"Mn,omitempty"`
	//
	Ms int32 `protobuf:"varint,3,opt,name=Ms,proto3" json:"Ms,omitempty"`
}

func (m *S2CQuizRank) Reset()         { *m = S2CQuizRank{} }
func (m *S2CQuizRank) String() string { return proto.CompactTextString(m) }
func (*S2CQuizRank) ProtoMessage()    {}
func (m *S2CQuizRank) GetId() uint16  { return PCK_S2CQuizRank_ID }

//
type QuizRank struct {
	//
	Nick string `protobuf:"bytes,1,opt,name=Nick,proto3" json:"Nick,omitempty"`
	//
	Score int32 `protobuf:"varint,2,opt,name=Score,proto3" json:"Score,omitempty"`
	//
	Num int32 `protobuf:"varint,3,opt,name=Num,proto3" json:"Num,omitempty"`
	//
	Id int32 `protobuf:"varint,4,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *QuizRank) Reset()         { *m = QuizRank{} }
func (m *QuizRank) String() string { return proto.CompactTextString(m) }
func (*QuizRank) ProtoMessage()    {}

const PCK_C2SAnswerQuiz_ID = 451 //
//
type C2SAnswerQuiz struct {
	//
	Aid int32 `protobuf:"varint,1,opt,name=Aid,proto3" json:"Aid,omitempty"`
}

func (m *C2SAnswerQuiz) Reset()         { *m = C2SAnswerQuiz{} }
func (m *C2SAnswerQuiz) String() string { return proto.CompactTextString(m) }
func (*C2SAnswerQuiz) ProtoMessage()    {}
func (m *C2SAnswerQuiz) GetId() uint16  { return PCK_C2SAnswerQuiz_ID }

const PCK_S2CAnswerQuiz_ID = 452 //
//
type S2CAnswerQuiz struct {
	//
	Y int32 `protobuf:"varint,1,opt,name=Y,proto3" json:"Y,omitempty"`
}

func (m *S2CAnswerQuiz) Reset()         { *m = S2CAnswerQuiz{} }
func (m *S2CAnswerQuiz) String() string { return proto.CompactTextString(m) }
func (*S2CAnswerQuiz) ProtoMessage()    {}
func (m *S2CAnswerQuiz) GetId() uint16  { return PCK_S2CAnswerQuiz_ID }

const PCK_S2CQuizSum_ID = 453 //
//
type S2CQuizSum struct {
	//
	Mn int32 `protobuf:"varint,1,opt,name=Mn,proto3" json:"Mn,omitempty"`
	//
	Ms int32 `protobuf:"varint,2,opt,name=Ms,proto3" json:"Ms,omitempty"`
	//
	I string `protobuf:"bytes,3,opt,name=I,proto3" json:"I,omitempty"`
}

func (m *S2CQuizSum) Reset()         { *m = S2CQuizSum{} }
func (m *S2CQuizSum) String() string { return proto.CompactTextString(m) }
func (*S2CQuizSum) ProtoMessage()    {}
func (m *S2CQuizSum) GetId() uint16  { return PCK_S2CQuizSum_ID }

const PCK_S2CQuizFirst_ID = 455 //
//
type S2CQuizFirst struct {
	//
	Nick string `protobuf:"bytes,1,opt,name=Nick,proto3" json:"Nick,omitempty"`
}

func (m *S2CQuizFirst) Reset()         { *m = S2CQuizFirst{} }
func (m *S2CQuizFirst) String() string { return proto.CompactTextString(m) }
func (*S2CQuizFirst) ProtoMessage()    {}
func (m *S2CQuizFirst) GetId() uint16  { return PCK_S2CQuizFirst_ID }

const PCK_S2CQuizStart_ID = 456 //
//
type S2CQuizStart struct {
}

func (m *S2CQuizStart) Reset()         { *m = S2CQuizStart{} }
func (m *S2CQuizStart) String() string { return proto.CompactTextString(m) }
func (*S2CQuizStart) ProtoMessage()    {}
func (m *S2CQuizStart) GetId() uint16  { return PCK_S2CQuizStart_ID }

//
type DbPlayerWestExp struct {
	//
	List []*DbWestExp `protobuf:"bytes,1,rep,name=List,proto3" json:"List,omitempty"`
}

func (m *DbPlayerWestExp) Reset()         { *m = DbPlayerWestExp{} }
func (m *DbPlayerWestExp) String() string { return proto.CompactTextString(m) }
func (*DbPlayerWestExp) ProtoMessage()    {}

//
type DbWestExp struct {
	//
	ProId int32 `protobuf:"varint,1,opt,name=ProId,proto3" json:"ProId,omitempty"`
	//
	StartTime int64 `protobuf:"varint,2,opt,name=StartTime,proto3" json:"StartTime,omitempty"`
	//
	RobberList []*RobbedPlayer `protobuf:"bytes,3,rep,name=RobberList,proto3" json:"RobberList,omitempty"`
	//
	IsHavePrize int32 `protobuf:"varint,4,opt,name=IsHavePrize,proto3" json:"IsHavePrize,omitempty"`
	//
	Flag int32 `protobuf:"varint,5,opt,name=Flag,proto3" json:"Flag,omitempty"`
	//
	Prize []*ItemInfo `protobuf:"bytes,6,rep,name=Prize,proto3" json:"Prize,omitempty"`
	//
	IsDouble int32 `protobuf:"varint,7,opt,name=IsDouble,proto3" json:"IsDouble,omitempty"`
	//
	IsHaveSend int32 `protobuf:"varint,8,opt,name=IsHaveSend,proto3" json:"IsHaveSend,omitempty"`
	//
	RobbedTimes int64 `protobuf:"varint,9,opt,name=RobbedTimes,proto3" json:"RobbedTimes,omitempty"`
	//
	LogIncId int64 `protobuf:"varint,10,opt,name=LogIncId,proto3" json:"LogIncId,omitempty"`
	//
	Idx int64 `protobuf:"varint,11,opt,name=Idx,proto3" json:"Idx,omitempty"`
	//
	FinishTime int64 `protobuf:"varint,12,opt,name=FinishTime,proto3" json:"FinishTime,omitempty"`
	//
	TotalRobberList []*RobbedPlayer `protobuf:"bytes,13,rep,name=TotalRobberList,proto3" json:"TotalRobberList,omitempty"`
}

func (m *DbWestExp) Reset()         { *m = DbWestExp{} }
func (m *DbWestExp) String() string { return proto.CompactTextString(m) }
func (*DbWestExp) ProtoMessage()    {}

//
type RobbedPlayer struct {
	//
	EnemyId int64 `protobuf:"varint,1,opt,name=EnemyId,proto3" json:"EnemyId,omitempty"`
	//
	ProId int32 `protobuf:"varint,2,opt,name=ProId,proto3" json:"ProId,omitempty"`
	//
	Y int32 `protobuf:"varint,4,opt,name=Y,proto3" json:"Y,omitempty"`
	//
	R int32 `protobuf:"varint,5,opt,name=R,proto3" json:"R,omitempty"`
	//
	LogId int64 `protobuf:"varint,6,opt,name=LogId,proto3" json:"LogId,omitempty"`
	//
	N string `protobuf:"bytes,7,opt,name=N,proto3" json:"N,omitempty"`
	//
	Idx int64 `protobuf:"varint,8,opt,name=Idx,proto3" json:"Idx,omitempty"`
}

func (m *RobbedPlayer) Reset()         { *m = RobbedPlayer{} }
func (m *RobbedPlayer) String() string { return proto.CompactTextString(m) }
func (*RobbedPlayer) ProtoMessage()    {}

const PCK_S2CWestExp_ID = 460 //
//
type S2CWestExp struct {
	//
	Lt int32 `protobuf:"varint,1,opt,name=Lt,proto3" json:"Lt,omitempty"`
	//
	Tt int32 `protobuf:"varint,2,opt,name=Tt,proto3" json:"Tt,omitempty"`
	//
	Rl int32 `protobuf:"varint,3,opt,name=Rl,proto3" json:"Rl,omitempty"`
	//
	Rt int32 `protobuf:"varint,4,opt,name=Rt,proto3" json:"Rt,omitempty"`
	//
	IsHavePrize int32 `protobuf:"varint,5,opt,name=IsHavePrize,proto3" json:"IsHavePrize,omitempty"`
	//
	St int64 `protobuf:"varint,6,opt,name=St,proto3" json:"St,omitempty"`
	//
	RobberList []*RobbedPlayer `protobuf:"bytes,7,rep,name=RobberList,proto3" json:"RobberList,omitempty"`
}

func (m *S2CWestExp) Reset()         { *m = S2CWestExp{} }
func (m *S2CWestExp) String() string { return proto.CompactTextString(m) }
func (*S2CWestExp) ProtoMessage()    {}
func (m *S2CWestExp) GetId() uint16  { return PCK_S2CWestExp_ID }

//
type ProtectingPlayer struct {
	//
	Uid int32 `protobuf:"varint,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
	//
	Rid int32 `protobuf:"varint,2,opt,name=Rid,proto3" json:"Rid,omitempty"`
	//
	N string `protobuf:"bytes,3,opt,name=N,proto3" json:"N,omitempty"`
	//
	Gn string `protobuf:"bytes,4,opt,name=Gn,proto3" json:"Gn,omitempty"`
	//
	Fv int64 `protobuf:"varint,5,opt,name=Fv,proto3" json:"Fv,omitempty"`
	//
	I int32 `protobuf:"varint,6,opt,name=I,proto3" json:"I,omitempty"`
	//
	Pt int64 `protobuf:"varint,7,opt,name=Pt,proto3" json:"Pt,omitempty"`
	//
	C int32 `protobuf:"varint,8,opt,name=C,proto3" json:"C,omitempty"`
}

func (m *ProtectingPlayer) Reset()         { *m = ProtectingPlayer{} }
func (m *ProtectingPlayer) String() string { return proto.CompactTextString(m) }
func (*ProtectingPlayer) ProtoMessage()    {}

const PCK_C2SGetProtectPlayer_ID = 461 //
//
type C2SGetProtectPlayer struct {
}

func (m *C2SGetProtectPlayer) Reset()         { *m = C2SGetProtectPlayer{} }
func (m *C2SGetProtectPlayer) String() string { return proto.CompactTextString(m) }
func (*C2SGetProtectPlayer) ProtoMessage()    {}
func (m *C2SGetProtectPlayer) GetId() uint16  { return PCK_C2SGetProtectPlayer_ID }

const PCK_S2CGetProtectPlayer_ID = 462 //
//
type S2CGetProtectPlayer struct {
	//
	List []*ProtectingPlayer `protobuf:"bytes,1,rep,name=List,proto3" json:"List,omitempty"`
}

func (m *S2CGetProtectPlayer) Reset()         { *m = S2CGetProtectPlayer{} }
func (m *S2CGetProtectPlayer) String() string { return proto.CompactTextString(m) }
func (*S2CGetProtectPlayer) ProtoMessage()    {}
func (m *S2CGetProtectPlayer) GetId() uint16  { return PCK_S2CGetProtectPlayer_ID }

const PCK_S2CNewProtectPlayer_ID = 481 //
//
type S2CNewProtectPlayer struct {
	//
	List *ProtectingPlayer `protobuf:"bytes,1,opt,name=List,proto3" json:"List,omitempty"`
}

func (m *S2CNewProtectPlayer) Reset()         { *m = S2CNewProtectPlayer{} }
func (m *S2CNewProtectPlayer) String() string { return proto.CompactTextString(m) }
func (*S2CNewProtectPlayer) ProtoMessage()    {}
func (m *S2CNewProtectPlayer) GetId() uint16  { return PCK_S2CNewProtectPlayer_ID }

const PCK_S2CEndProtectPlayer_ID = 482 //
//
type S2CEndProtectPlayer struct {
	//
	Uid int32 `protobuf:"varint,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
}

func (m *S2CEndProtectPlayer) Reset()         { *m = S2CEndProtectPlayer{} }
func (m *S2CEndProtectPlayer) String() string { return proto.CompactTextString(m) }
func (*S2CEndProtectPlayer) ProtoMessage()    {}
func (m *S2CEndProtectPlayer) GetId() uint16  { return PCK_S2CEndProtectPlayer_ID }

const PCK_C2SGetWestExp_ID = 463 //
//
type C2SGetWestExp struct {
	//
	GetType int32 `protobuf:"varint,1,opt,name=GetType,proto3" json:"GetType,omitempty"`
}

func (m *C2SGetWestExp) Reset()         { *m = C2SGetWestExp{} }
func (m *C2SGetWestExp) String() string { return proto.CompactTextString(m) }
func (*C2SGetWestExp) ProtoMessage()    {}
func (m *C2SGetWestExp) GetId() uint16  { return PCK_C2SGetWestExp_ID }

const PCK_S2CGetWestExp_ID = 464 //
//
type S2CGetWestExp struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	I int32 `protobuf:"varint,2,opt,name=I,proto3" json:"I,omitempty"`
}

func (m *S2CGetWestExp) Reset()         { *m = S2CGetWestExp{} }
func (m *S2CGetWestExp) String() string { return proto.CompactTextString(m) }
func (*S2CGetWestExp) ProtoMessage()    {}
func (m *S2CGetWestExp) GetId() uint16  { return PCK_S2CGetWestExp_ID }

const PCK_C2SQuickFinishWestExp_ID = 465 //
//
type C2SQuickFinishWestExp struct {
}

func (m *C2SQuickFinishWestExp) Reset()         { *m = C2SQuickFinishWestExp{} }
func (m *C2SQuickFinishWestExp) String() string { return proto.CompactTextString(m) }
func (*C2SQuickFinishWestExp) ProtoMessage()    {}
func (m *C2SQuickFinishWestExp) GetId() uint16  { return PCK_C2SQuickFinishWestExp_ID }

const PCK_S2CQuickFinishWestExp_ID = 466 //
//
type S2CQuickFinishWestExp struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CQuickFinishWestExp) Reset()         { *m = S2CQuickFinishWestExp{} }
func (m *S2CQuickFinishWestExp) String() string { return proto.CompactTextString(m) }
func (*S2CQuickFinishWestExp) ProtoMessage()    {}
func (m *S2CQuickFinishWestExp) GetId() uint16  { return PCK_S2CQuickFinishWestExp_ID }

const PCK_C2SStartWestExp_ID = 474 //
//
type C2SStartWestExp struct {
}

func (m *C2SStartWestExp) Reset()         { *m = C2SStartWestExp{} }
func (m *C2SStartWestExp) String() string { return proto.CompactTextString(m) }
func (*C2SStartWestExp) ProtoMessage()    {}
func (m *C2SStartWestExp) GetId() uint16  { return PCK_C2SStartWestExp_ID }

const PCK_S2CStartWestExp_ID = 475 //
//
type S2CStartWestExp struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CStartWestExp) Reset()         { *m = S2CStartWestExp{} }
func (m *S2CStartWestExp) String() string { return proto.CompactTextString(m) }
func (*S2CStartWestExp) ProtoMessage()    {}
func (m *S2CStartWestExp) GetId() uint16  { return PCK_S2CStartWestExp_ID }

const PCK_S2CFinishWestExp_ID = 467 //
//
type S2CFinishWestExp struct {
	//
	I int32 `protobuf:"varint,1,opt,name=I,proto3" json:"I,omitempty"`
	//
	P string `protobuf:"bytes,3,opt,name=P,proto3" json:"P,omitempty"`
	//
	List []*RobbedPlayer `protobuf:"bytes,2,rep,name=List,proto3" json:"List,omitempty"`
	//
	D int32 `protobuf:"varint,4,opt,name=D,proto3" json:"D,omitempty"`
}

func (m *S2CFinishWestExp) Reset()         { *m = S2CFinishWestExp{} }
func (m *S2CFinishWestExp) String() string { return proto.CompactTextString(m) }
func (*S2CFinishWestExp) ProtoMessage()    {}
func (m *S2CFinishWestExp) GetId() uint16  { return PCK_S2CFinishWestExp_ID }

const PCK_C2SGetWestPrize_ID = 468 //
//
type C2SGetWestPrize struct {
}

func (m *C2SGetWestPrize) Reset()         { *m = C2SGetWestPrize{} }
func (m *C2SGetWestPrize) String() string { return proto.CompactTextString(m) }
func (*C2SGetWestPrize) ProtoMessage()    {}
func (m *C2SGetWestPrize) GetId() uint16  { return PCK_C2SGetWestPrize_ID }

const PCK_S2CGetWestPrize_ID = 469 //
//
type S2CGetWestPrize struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CGetWestPrize) Reset()         { *m = S2CGetWestPrize{} }
func (m *S2CGetWestPrize) String() string { return proto.CompactTextString(m) }
func (*S2CGetWestPrize) ProtoMessage()    {}
func (m *S2CGetWestPrize) GetId() uint16  { return PCK_S2CGetWestPrize_ID }

const PCK_C2SGetRobbedList_ID = 470 //
//
type C2SGetRobbedList struct {
}

func (m *C2SGetRobbedList) Reset()         { *m = C2SGetRobbedList{} }
func (m *C2SGetRobbedList) String() string { return proto.CompactTextString(m) }
func (*C2SGetRobbedList) ProtoMessage()    {}
func (m *C2SGetRobbedList) GetId() uint16  { return PCK_C2SGetRobbedList_ID }

//
type Robber struct {
	//
	U int32 `protobuf:"varint,1,opt,name=U,proto3" json:"U,omitempty"`
	//
	Fv int64 `protobuf:"varint,2,opt,name=Fv,proto3" json:"Fv,omitempty"`
	//
	N string `protobuf:"bytes,3,opt,name=N,proto3" json:"N,omitempty"`
	//
	I int32 `protobuf:"varint,4,opt,name=I,proto3" json:"I,omitempty"`
	//
	Y int32 `protobuf:"varint,5,opt,name=Y,proto3" json:"Y,omitempty"`
	//
	Lid string `protobuf:"bytes,6,opt,name=Lid,proto3" json:"Lid,omitempty"`
	//
	G string `protobuf:"bytes,7,opt,name=G,proto3" json:"G,omitempty"`
	//
	Uid int32 `protobuf:"varint,8,opt,name=Uid,proto3" json:"Uid,omitempty"`
}

func (m *Robber) Reset()         { *m = Robber{} }
func (m *Robber) String() string { return proto.CompactTextString(m) }
func (*Robber) ProtoMessage()    {}

const PCK_S2CGetRobbedList_ID = 471 //
//
type S2CGetRobbedList struct {
	//
	List []*Robber `protobuf:"bytes,1,rep,name=List,proto3" json:"List,omitempty"`
}

func (m *S2CGetRobbedList) Reset()         { *m = S2CGetRobbedList{} }
func (m *S2CGetRobbedList) String() string { return proto.CompactTextString(m) }
func (*S2CGetRobbedList) ProtoMessage()    {}
func (m *S2CGetRobbedList) GetId() uint16  { return PCK_S2CGetRobbedList_ID }

const PCK_C2SSendRob_ID = 476 //
//
type C2SSendRob struct {
	//
	U int32 `protobuf:"varint,1,opt,name=U,proto3" json:"U,omitempty"`
}

func (m *C2SSendRob) Reset()         { *m = C2SSendRob{} }
func (m *C2SSendRob) String() string { return proto.CompactTextString(m) }
func (*C2SSendRob) ProtoMessage()    {}
func (m *C2SSendRob) GetId() uint16  { return PCK_C2SSendRob_ID }

const PCK_S2CSendRob_ID = 477 //
//
type S2CSendRob struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CSendRob) Reset()         { *m = S2CSendRob{} }
func (m *S2CSendRob) String() string { return proto.CompactTextString(m) }
func (*S2CSendRob) ProtoMessage()    {}
func (m *S2CSendRob) GetId() uint16  { return PCK_S2CSendRob_ID }

const PCK_S2CBeRob_ID = 478 //
//
type S2CBeRob struct {
}

func (m *S2CBeRob) Reset()         { *m = S2CBeRob{} }
func (m *S2CBeRob) String() string { return proto.CompactTextString(m) }
func (*S2CBeRob) ProtoMessage()    {}
func (m *S2CBeRob) GetId() uint16  { return PCK_S2CBeRob_ID }

const PCK_C2SSendRevenge_ID = 472 //
//
type C2SSendRevenge struct {
	//
	Lid string `protobuf:"bytes,1,opt,name=Lid,proto3" json:"Lid,omitempty"`
}

func (m *C2SSendRevenge) Reset()         { *m = C2SSendRevenge{} }
func (m *C2SSendRevenge) String() string { return proto.CompactTextString(m) }
func (*C2SSendRevenge) ProtoMessage()    {}
func (m *C2SSendRevenge) GetId() uint16  { return PCK_C2SSendRevenge_ID }

const PCK_S2CSendRevenge_ID = 473 //
//
type S2CSendRevenge struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	P string `protobuf:"bytes,2,opt,name=P,proto3" json:"P,omitempty"`
	//
	Y int32 `protobuf:"varint,3,opt,name=Y,proto3" json:"Y,omitempty"`
	//
	Lid string `protobuf:"bytes,4,opt,name=Lid,proto3" json:"Lid,omitempty"`
}

func (m *S2CSendRevenge) Reset()         { *m = S2CSendRevenge{} }
func (m *S2CSendRevenge) String() string { return proto.CompactTextString(m) }
func (*S2CSendRevenge) ProtoMessage()    {}
func (m *S2CSendRevenge) GetId() uint16  { return PCK_S2CSendRevenge_ID }

const PCK_S2CWestExpStart_ID = 480 //
//
type S2CWestExpStart struct {
}

func (m *S2CWestExpStart) Reset()         { *m = S2CWestExpStart{} }
func (m *S2CWestExpStart) String() string { return proto.CompactTextString(m) }
func (*S2CWestExpStart) ProtoMessage()    {}
func (m *S2CWestExpStart) GetId() uint16  { return PCK_S2CWestExpStart_ID }

//
type Rank struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	R int32 `protobuf:"varint,2,opt,name=R,proto3" json:"R,omitempty"`
	//
	A []*IntAttr `protobuf:"bytes,3,rep,name=A,proto3" json:"A,omitempty"`
	//
	B []*StrAttr `protobuf:"bytes,4,rep,name=B,proto3" json:"B,omitempty"`
	//
	SortValue int64 `protobuf:"varint,5,opt,name=SortValue,proto3" json:"SortValue,omitempty"`
	//
	AdditionalOk int32 `protobuf:"varint,6,opt,name=AdditionalOk,proto3" json:"AdditionalOk,omitempty"`
}

func (m *Rank) Reset()         { *m = Rank{} }
func (m *Rank) String() string { return proto.CompactTextString(m) }
func (*Rank) ProtoMessage()    {}

//
type SimpleRank struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	R int32 `protobuf:"varint,2,opt,name=R,proto3" json:"R,omitempty"`
	//
	Name string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	//
	Level int32 `protobuf:"varint,4,opt,name=Level,proto3" json:"Level,omitempty"`
	//
	Fv int64 `protobuf:"varint,5,opt,name=Fv,proto3" json:"Fv,omitempty"`
	//
	SortValue int64 `protobuf:"varint,6,opt,name=SortValue,proto3" json:"SortValue,omitempty"`
	//
	Vip int64 `protobuf:"varint,7,opt,name=Vip,proto3" json:"Vip,omitempty"`
	//
	AdditionalOk int32 `protobuf:"varint,8,opt,name=AdditionalOk,proto3" json:"AdditionalOk,omitempty"`
}

func (m *SimpleRank) Reset()         { *m = SimpleRank{} }
func (m *SimpleRank) String() string { return proto.CompactTextString(m) }
func (*SimpleRank) ProtoMessage()    {}

const PCK_C2SAllRank_ID = 900 //
//
type C2SAllRank struct {
	//
	Type int32 `protobuf:"varint,1,opt,name=Type,proto3" json:"Type,omitempty"`
	//
	Param int32 `protobuf:"varint,2,opt,name=Param,proto3" json:"Param,omitempty"`
	//
	FullDataRank []int32 `protobuf:"varint,3,rep,name=FullDataRank,proto3" json:"FullDataRank,omitempty"`
	//
	SimpleDataRank []int32 `protobuf:"varint,4,rep,name=SimpleDataRank,proto3" json:"SimpleDataRank,omitempty"`
}

func (m *C2SAllRank) Reset()         { *m = C2SAllRank{} }
func (m *C2SAllRank) String() string { return proto.CompactTextString(m) }
func (*C2SAllRank) ProtoMessage()    {}
func (m *C2SAllRank) GetId() uint16  { return PCK_C2SAllRank_ID }

const PCK_S2CAllRank_ID = 901 //
//
type S2CAllRank struct {
	//
	Type int32 `protobuf:"varint,1,opt,name=Type,proto3" json:"Type,omitempty"`
	//
	Param int32 `protobuf:"varint,2,opt,name=Param,proto3" json:"Param,omitempty"`
	//
	FullDataRank []*Rank `protobuf:"bytes,3,rep,name=FullDataRank,proto3" json:"FullDataRank,omitempty"`
	//
	SimpleDataRank []*SimpleRank `protobuf:"bytes,4,rep,name=SimpleDataRank,proto3" json:"SimpleDataRank,omitempty"`
	//
	MyData *SimpleRank `protobuf:"bytes,5,opt,name=MyData,proto3" json:"MyData,omitempty"`
}

func (m *S2CAllRank) Reset()         { *m = S2CAllRank{} }
func (m *S2CAllRank) String() string { return proto.CompactTextString(m) }
func (*S2CAllRank) ProtoMessage()    {}
func (m *S2CAllRank) GetId() uint16  { return PCK_S2CAllRank_ID }

const PCK_C2SRespect_ID = 13 //
//
type C2SRespect struct {
}

func (m *C2SRespect) Reset()         { *m = C2SRespect{} }
func (m *C2SRespect) String() string { return proto.CompactTextString(m) }
func (*C2SRespect) ProtoMessage()    {}
func (m *C2SRespect) GetId() uint16  { return PCK_C2SRespect_ID }

const PCK_S2CRespect_ID = 14 //
//
type S2CRespect struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Prize []*ItemInfo `protobuf:"bytes,2,rep,name=Prize,proto3" json:"Prize,omitempty"`
}

func (m *S2CRespect) Reset()         { *m = S2CRespect{} }
func (m *S2CRespect) String() string { return proto.CompactTextString(m) }
func (*S2CRespect) ProtoMessage()    {}
func (m *S2CRespect) GetId() uint16  { return PCK_S2CRespect_ID }

//
type RankSortValue struct {
	//
	PlayerId int32 `protobuf:"varint,1,opt,name=PlayerId,proto3" json:"PlayerId,omitempty"`
	//
	SortValue int64 `protobuf:"varint,2,opt,name=SortValue,proto3" json:"SortValue,omitempty"`
}

func (m *RankSortValue) Reset()         { *m = RankSortValue{} }
func (m *RankSortValue) String() string { return proto.CompactTextString(m) }
func (*RankSortValue) ProtoMessage()    {}

//
type NewRankDbInfo struct {
	//
	players []*Rank `protobuf:"bytes,1,rep,name=players,proto3" json:"players,omitempty"`
	//
	allplayersData []*RankSortValue `protobuf:"bytes,2,rep,name=allplayersData,proto3" json:"allplayersData,omitempty"`
}

func (m *NewRankDbInfo) Reset()         { *m = NewRankDbInfo{} }
func (m *NewRankDbInfo) String() string { return proto.CompactTextString(m) }
func (*NewRankDbInfo) ProtoMessage()    {}

//
type GangDbMember struct {
	//
	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	JoinTime int64 `protobuf:"varint,2,opt,name=JoinTime,proto3" json:"JoinTime,omitempty"`
	//
	DutyId int32 `protobuf:"varint,3,opt,name=DutyId,proto3" json:"DutyId,omitempty"`
}

func (m *GangDbMember) Reset()         { *m = GangDbMember{} }
func (m *GangDbMember) String() string { return proto.CompactTextString(m) }
func (*GangDbMember) ProtoMessage()    {}

//
type GangDbApply struct {
	//
	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	ApplyTime int64 `protobuf:"varint,2,opt,name=ApplyTime,proto3" json:"ApplyTime,omitempty"`
}

func (m *GangDbApply) Reset()         { *m = GangDbApply{} }
func (m *GangDbApply) String() string { return proto.CompactTextString(m) }
func (*GangDbApply) ProtoMessage()    {}

//
type GangDbAction struct {
	//
	Type int32 `protobuf:"varint,1,opt,name=Type,proto3" json:"Type,omitempty"`
	//
	UserId int64 `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	ActionTime int64 `protobuf:"varint,3,opt,name=ActionTime,proto3" json:"ActionTime,omitempty"`
	//
	DutyId int32 `protobuf:"varint,4,opt,name=DutyId,proto3" json:"DutyId,omitempty"`
	//
	Param string `protobuf:"bytes,5,opt,name=Param,proto3" json:"Param,omitempty"`
}

func (m *GangDbAction) Reset()         { *m = GangDbAction{} }
func (m *GangDbAction) String() string { return proto.CompactTextString(m) }
func (*GangDbAction) ProtoMessage()    {}

//
type GangDbPeach struct {
	//
	EatTime int64 `protobuf:"varint,1,opt,name=EatTime,proto3" json:"EatTime,omitempty"`
	//
	UserId int64 `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	PeachId int32 `protobuf:"varint,3,opt,name=PeachId,proto3" json:"PeachId,omitempty"`
}

func (m *GangDbPeach) Reset()         { *m = GangDbPeach{} }
func (m *GangDbPeach) String() string { return proto.CompactTextString(m) }
func (*GangDbPeach) ProtoMessage()    {}

//
type GangDbInfo struct {
	//
	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Members []*GangDbMember `protobuf:"bytes,2,rep,name=Members,proto3" json:"Members,omitempty"`
	//
	Applys []*GangDbApply `protobuf:"bytes,3,rep,name=Applys,proto3" json:"Applys,omitempty"`
	//
	Name string `protobuf:"bytes,4,opt,name=Name,proto3" json:"Name,omitempty"`
	//
	Money int64 `protobuf:"varint,5,opt,name=Money,proto3" json:"Money,omitempty"`
	//
	Level int32 `protobuf:"varint,6,opt,name=Level,proto3" json:"Level,omitempty"`
	//
	NoticeId int32 `protobuf:"varint,7,opt,name=NoticeId,proto3" json:"NoticeId,omitempty"`
	//
	NeedFightValue int64 `protobuf:"varint,8,opt,name=NeedFightValue,proto3" json:"NeedFightValue,omitempty"`
	//
	Actions []*GangDbAction `protobuf:"bytes,9,rep,name=Actions,proto3" json:"Actions,omitempty"`
	//
	PeachVal int64 `protobuf:"varint,10,opt,name=PeachVal,proto3" json:"PeachVal,omitempty"`
	//
	RecordPeach []*GangDbPeach `protobuf:"bytes,11,rep,name=RecordPeach,proto3" json:"RecordPeach,omitempty"`
	//
	AutoJoin bool `protobuf:"varint,12,opt,name=AutoJoin,proto3" json:"AutoJoin,omitempty"`
	//
	CreateServerId int64 `protobuf:"varint,13,opt,name=CreateServerId,proto3" json:"CreateServerId,omitempty"`
	//
	CreateAreaId int64 `protobuf:"varint,14,opt,name=CreateAreaId,proto3" json:"CreateAreaId,omitempty"`
	//
	RedPkgPool int64 `protobuf:"varint,15,opt,name=RedPkgPool,proto3" json:"RedPkgPool,omitempty"`
	//
	GangFightVal int64 `protobuf:"varint,16,opt,name=GangFightVal,proto3" json:"GangFightVal,omitempty"`
}

func (m *GangDbInfo) Reset()         { *m = GangDbInfo{} }
func (m *GangDbInfo) String() string { return proto.CompactTextString(m) }
func (*GangDbInfo) ProtoMessage()    {}

const PCK_S2CGangRedPot_ID = 1288 //
//
type S2CGangRedPot struct {
	//
	A bool `protobuf:"varint,1,opt,name=A,proto3" json:"A,omitempty"`
	//
	E bool `protobuf:"varint,2,opt,name=E,proto3" json:"E,omitempty"`
	//
	P bool `protobuf:"varint,3,opt,name=P,proto3" json:"P,omitempty"`
	//
	G bool `protobuf:"varint,4,opt,name=G,proto3" json:"G,omitempty"`
	//
	kM bool `protobuf:"varint,5,opt,name=kM,proto3" json:"kM,omitempty"`
}

func (m *S2CGangRedPot) Reset()         { *m = S2CGangRedPot{} }
func (m *S2CGangRedPot) String() string { return proto.CompactTextString(m) }
func (*S2CGangRedPot) ProtoMessage()    {}
func (m *S2CGangRedPot) GetId() uint16  { return PCK_S2CGangRedPot_ID }

//
type GangExchange struct {
	//
	NextReTime int64 `protobuf:"varint,1,opt,name=NextReTime,proto3" json:"NextReTime,omitempty"`
	//
	ExchangeIds []int32 `protobuf:"varint,2,rep,name=ExchangeIds,proto3" json:"ExchangeIds,omitempty"`
	//
	ReExchangeIds []int32 `protobuf:"varint,3,rep,name=ReExchangeIds,proto3" json:"ReExchangeIds,omitempty"`
	//
	RefreshCount int32 `protobuf:"varint,4,opt,name=RefreshCount,proto3" json:"RefreshCount,omitempty"`
}

func (m *GangExchange) Reset()         { *m = GangExchange{} }
func (m *GangExchange) String() string { return proto.CompactTextString(m) }
func (*GangExchange) ProtoMessage()    {}

const PCK_C2SGangFindExchange_ID = 1252 //
//
type C2SGangFindExchange struct {
	//
	GangId string `protobuf:"bytes,1,opt,name=GangId,proto3" json:"GangId,omitempty"`
}

func (m *C2SGangFindExchange) Reset()         { *m = C2SGangFindExchange{} }
func (m *C2SGangFindExchange) String() string { return proto.CompactTextString(m) }
func (*C2SGangFindExchange) ProtoMessage()    {}
func (m *C2SGangFindExchange) GetId() uint16  { return PCK_C2SGangFindExchange_ID }

const PCK_S2CGangFindExchange_ID = 1253 //
//
type S2CGangFindExchange struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Exchange *GangExchange `protobuf:"bytes,2,opt,name=Exchange,proto3" json:"Exchange,omitempty"`
}

func (m *S2CGangFindExchange) Reset()         { *m = S2CGangFindExchange{} }
func (m *S2CGangFindExchange) String() string { return proto.CompactTextString(m) }
func (*S2CGangFindExchange) ProtoMessage()    {}
func (m *S2CGangFindExchange) GetId() uint16  { return PCK_S2CGangFindExchange_ID }

const PCK_C2SGangExchange_ID = 1254 //
//
type C2SGangExchange struct {
	//
	GangId string `protobuf:"bytes,1,opt,name=GangId,proto3" json:"GangId,omitempty"`
	//
	ExcId int32 `protobuf:"varint,2,opt,name=ExcId,proto3" json:"ExcId,omitempty"`
}

func (m *C2SGangExchange) Reset()         { *m = C2SGangExchange{} }
func (m *C2SGangExchange) String() string { return proto.CompactTextString(m) }
func (*C2SGangExchange) ProtoMessage()    {}
func (m *C2SGangExchange) GetId() uint16  { return PCK_C2SGangExchange_ID }

const PCK_S2CGangExchange_ID = 1255 //
//
type S2CGangExchange struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Exchange *GangExchange `protobuf:"bytes,2,opt,name=Exchange,proto3" json:"Exchange,omitempty"`
}

func (m *S2CGangExchange) Reset()         { *m = S2CGangExchange{} }
func (m *S2CGangExchange) String() string { return proto.CompactTextString(m) }
func (*S2CGangExchange) ProtoMessage()    {}
func (m *S2CGangExchange) GetId() uint16  { return PCK_S2CGangExchange_ID }

const PCK_C2SGangRefreshExc_ID = 1256 //
//
type C2SGangRefreshExc struct {
	//
	GangId string `protobuf:"bytes,1,opt,name=GangId,proto3" json:"GangId,omitempty"`
}

func (m *C2SGangRefreshExc) Reset()         { *m = C2SGangRefreshExc{} }
func (m *C2SGangRefreshExc) String() string { return proto.CompactTextString(m) }
func (*C2SGangRefreshExc) ProtoMessage()    {}
func (m *C2SGangRefreshExc) GetId() uint16  { return PCK_C2SGangRefreshExc_ID }

const PCK_S2CGangRefreshExc_ID = 1257 //
//
type S2CGangRefreshExc struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Exchange *GangExchange `protobuf:"bytes,2,opt,name=Exchange,proto3" json:"Exchange,omitempty"`
}

func (m *S2CGangRefreshExc) Reset()         { *m = S2CGangRefreshExc{} }
func (m *S2CGangRefreshExc) String() string { return proto.CompactTextString(m) }
func (*S2CGangRefreshExc) ProtoMessage()    {}
func (m *S2CGangRefreshExc) GetId() uint16  { return PCK_S2CGangRefreshExc_ID }

const PCK_C2SGangSimpleInfo_ID = 1286 //
//
type C2SGangSimpleInfo struct {
	//
	GId string `protobuf:"bytes,1,opt,name=GId,proto3" json:"GId,omitempty"`
}

func (m *C2SGangSimpleInfo) Reset()         { *m = C2SGangSimpleInfo{} }
func (m *C2SGangSimpleInfo) String() string { return proto.CompactTextString(m) }
func (*C2SGangSimpleInfo) ProtoMessage()    {}
func (m *C2SGangSimpleInfo) GetId() uint16  { return PCK_C2SGangSimpleInfo_ID }

const PCK_S2CGangSimpleInfo_ID = 1287 //
//
type S2CGangSimpleInfo struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	GId string `protobuf:"bytes,2,opt,name=GId,proto3" json:"GId,omitempty"`
	//
	GN string `protobuf:"bytes,3,opt,name=GN,proto3" json:"GN,omitempty"`
	//
	GL int32 `protobuf:"varint,4,opt,name=GL,proto3" json:"GL,omitempty"`
}

func (m *S2CGangSimpleInfo) Reset()         { *m = S2CGangSimpleInfo{} }
func (m *S2CGangSimpleInfo) String() string { return proto.CompactTextString(m) }
func (*S2CGangSimpleInfo) ProtoMessage()    {}
func (m *S2CGangSimpleInfo) GetId() uint16  { return PCK_S2CGangSimpleInfo_ID }

const PCK_C2SGangInfo_ID = 1201 //
//
type C2SGangInfo struct {
	//
	GangId string `protobuf:"bytes,1,opt,name=GangId,proto3" json:"GangId,omitempty"`
}

func (m *C2SGangInfo) Reset()         { *m = C2SGangInfo{} }
func (m *C2SGangInfo) String() string { return proto.CompactTextString(m) }
func (*C2SGangInfo) ProtoMessage()    {}
func (m *C2SGangInfo) GetId() uint16  { return PCK_C2SGangInfo_ID }

const PCK_S2CGangInfo_ID = 1202 //
//
type S2CGangInfo struct {
	//
	Gang *GangInfo `protobuf:"bytes,1,opt,name=Gang,proto3" json:"Gang,omitempty"`
	//
	Tag int32 `protobuf:"varint,2,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CGangInfo) Reset()         { *m = S2CGangInfo{} }
func (m *S2CGangInfo) String() string { return proto.CompactTextString(m) }
func (*S2CGangInfo) ProtoMessage()    {}
func (m *S2CGangInfo) GetId() uint16  { return PCK_S2CGangInfo_ID }

//
type GangInfo struct {
	//
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	//
	MasterName string `protobuf:"bytes,2,opt,name=MasterName,proto3" json:"MasterName,omitempty"`
	//
	Nums int32 `protobuf:"varint,3,opt,name=Nums,proto3" json:"Nums,omitempty"`
	//
	MaxNums int32 `protobuf:"varint,4,opt,name=MaxNums,proto3" json:"MaxNums,omitempty"`
	//
	Money int32 `protobuf:"varint,5,opt,name=Money,proto3" json:"Money,omitempty"`
	//
	UpMoney int32 `protobuf:"varint,6,opt,name=UpMoney,proto3" json:"UpMoney,omitempty"`
	//
	Level int32 `protobuf:"varint,7,opt,name=Level,proto3" json:"Level,omitempty"`
	//
	NoticeId int32 `protobuf:"varint,8,opt,name=NoticeId,proto3" json:"NoticeId,omitempty"`
	//
	Id string `protobuf:"bytes,9,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	NeedFightValue int64 `protobuf:"varint,10,opt,name=NeedFightValue,proto3" json:"NeedFightValue,omitempty"`
	//
	IsJoin int32 `protobuf:"varint,11,opt,name=IsJoin,proto3" json:"IsJoin,omitempty"`
	//
	MasterId int64 `protobuf:"varint,12,opt,name=MasterId,proto3" json:"MasterId,omitempty"`
	//
	MasterRoleId int64 `protobuf:"varint,13,opt,name=MasterRoleId,proto3" json:"MasterRoleId,omitempty"`
	//
	IsAutoJoin bool `protobuf:"varint,14,opt,name=IsAutoJoin,proto3" json:"IsAutoJoin,omitempty"`
	//
	GangFightVal int64 `protobuf:"varint,15,opt,name=GangFightVal,proto3" json:"GangFightVal,omitempty"`
	//
	KM int32 `protobuf:"varint,16,opt,name=KM,proto3" json:"KM,omitempty"`
}

func (m *GangInfo) Reset()         { *m = GangInfo{} }
func (m *GangInfo) String() string { return proto.CompactTextString(m) }
func (*GangInfo) ProtoMessage()    {}

const PCK_C2SGangList_ID = 1203 //
//
type C2SGangList struct {
	//
	Page int32 `protobuf:"varint,1,opt,name=Page,proto3" json:"Page,omitempty"`
}

func (m *C2SGangList) Reset()         { *m = C2SGangList{} }
func (m *C2SGangList) String() string { return proto.CompactTextString(m) }
func (*C2SGangList) ProtoMessage()    {}
func (m *C2SGangList) GetId() uint16  { return PCK_C2SGangList_ID }

const PCK_S2CGangList_ID = 1204 //
//
type S2CGangList struct {
	//
	Gangs []*GangInfo `protobuf:"bytes,1,rep,name=Gangs,proto3" json:"Gangs,omitempty"`
	//
	Tag int32 `protobuf:"varint,2,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Page int32 `protobuf:"varint,3,opt,name=Page,proto3" json:"Page,omitempty"`
	//
	Row int32 `protobuf:"varint,4,opt,name=Row,proto3" json:"Row,omitempty"`
	//
	Total int32 `protobuf:"varint,5,opt,name=Total,proto3" json:"Total,omitempty"`
	//
	Prize int32 `protobuf:"varint,6,opt,name=Prize,proto3" json:"Prize,omitempty"`
}

func (m *S2CGangList) Reset()         { *m = S2CGangList{} }
func (m *S2CGangList) String() string { return proto.CompactTextString(m) }
func (*S2CGangList) ProtoMessage()    {}
func (m *S2CGangList) GetId() uint16  { return PCK_S2CGangList_ID }

const PCK_C2SCreateGang_ID = 1205 //
//
type C2SCreateGang struct {
	//
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	//
	Level int32 `protobuf:"varint,2,opt,name=Level,proto3" json:"Level,omitempty"`
}

func (m *C2SCreateGang) Reset()         { *m = C2SCreateGang{} }
func (m *C2SCreateGang) String() string { return proto.CompactTextString(m) }
func (*C2SCreateGang) ProtoMessage()    {}
func (m *C2SCreateGang) GetId() uint16  { return PCK_C2SCreateGang_ID }

const PCK_S2CCreateGang_ID = 1233 //
//
type S2CCreateGang struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Gang *GangInfo `protobuf:"bytes,2,opt,name=Gang,proto3" json:"Gang,omitempty"`
	//
	DropId int32 `protobuf:"varint,3,opt,name=DropId,proto3" json:"DropId,omitempty"`
}

func (m *S2CCreateGang) Reset()         { *m = S2CCreateGang{} }
func (m *S2CCreateGang) String() string { return proto.CompactTextString(m) }
func (*S2CCreateGang) ProtoMessage()    {}
func (m *S2CCreateGang) GetId() uint16  { return PCK_S2CCreateGang_ID }

const PCK_C2SApplyGang_ID = 1206 //
//
type C2SApplyGang struct {
	//
	GangId string `protobuf:"bytes,1,opt,name=GangId,proto3" json:"GangId,omitempty"`
}

func (m *C2SApplyGang) Reset()         { *m = C2SApplyGang{} }
func (m *C2SApplyGang) String() string { return proto.CompactTextString(m) }
func (*C2SApplyGang) ProtoMessage()    {}
func (m *C2SApplyGang) GetId() uint16  { return PCK_C2SApplyGang_ID }

const PCK_S2CApplyGang_ID = 1207 //
//
type S2CApplyGang struct {
	//
	GangId string `protobuf:"bytes,1,opt,name=GangId,proto3" json:"GangId,omitempty"`
	//
	Tag int32 `protobuf:"varint,2,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	IsJoin int32 `protobuf:"varint,3,opt,name=IsJoin,proto3" json:"IsJoin,omitempty"`
}

func (m *S2CApplyGang) Reset()         { *m = S2CApplyGang{} }
func (m *S2CApplyGang) String() string { return proto.CompactTextString(m) }
func (*S2CApplyGang) ProtoMessage()    {}
func (m *S2CApplyGang) GetId() uint16  { return PCK_S2CApplyGang_ID }

const PCK_C2SAgreeApplyGang_ID = 1208 //
//
type C2SAgreeApplyGang struct {
	//
	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	GangId string `protobuf:"bytes,2,opt,name=GangId,proto3" json:"GangId,omitempty"`
	//
	ActionType int32 `protobuf:"varint,3,opt,name=ActionType,proto3" json:"ActionType,omitempty"`
}

func (m *C2SAgreeApplyGang) Reset()         { *m = C2SAgreeApplyGang{} }
func (m *C2SAgreeApplyGang) String() string { return proto.CompactTextString(m) }
func (*C2SAgreeApplyGang) ProtoMessage()    {}
func (m *C2SAgreeApplyGang) GetId() uint16  { return PCK_C2SAgreeApplyGang_ID }

const PCK_S2CAgreeApplyGang_ID = 1209 //
//
type S2CAgreeApplyGang struct {
	//
	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	Tag int32 `protobuf:"varint,2,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Nums int32 `protobuf:"varint,3,opt,name=Nums,proto3" json:"Nums,omitempty"`
	//
	GangId string `protobuf:"bytes,4,opt,name=GangId,proto3" json:"GangId,omitempty"`
	//
	ActionType int32 `protobuf:"varint,5,opt,name=ActionType,proto3" json:"ActionType,omitempty"`
	//
	RoleId int64 `protobuf:"varint,6,opt,name=RoleId,proto3" json:"RoleId,omitempty"`
	//
	GN string `protobuf:"bytes,7,opt,name=GN,proto3" json:"GN,omitempty"`
}

func (m *S2CAgreeApplyGang) Reset()         { *m = S2CAgreeApplyGang{} }
func (m *S2CAgreeApplyGang) String() string { return proto.CompactTextString(m) }
func (*S2CAgreeApplyGang) ProtoMessage()    {}
func (m *S2CAgreeApplyGang) GetId() uint16  { return PCK_S2CAgreeApplyGang_ID }

const PCK_C2SUpNoticeGang_ID = 1210 //
//
type C2SUpNoticeGang struct {
	//
	GangId string `protobuf:"bytes,1,opt,name=GangId,proto3" json:"GangId,omitempty"`
	//
	NoticeId int32 `protobuf:"varint,2,opt,name=NoticeId,proto3" json:"NoticeId,omitempty"`
}

func (m *C2SUpNoticeGang) Reset()         { *m = C2SUpNoticeGang{} }
func (m *C2SUpNoticeGang) String() string { return proto.CompactTextString(m) }
func (*C2SUpNoticeGang) ProtoMessage()    {}
func (m *C2SUpNoticeGang) GetId() uint16  { return PCK_C2SUpNoticeGang_ID }

const PCK_S2CUpNoticeGang_ID = 1211 //
//
type S2CUpNoticeGang struct {
	//
	GangId string `protobuf:"bytes,1,opt,name=GangId,proto3" json:"GangId,omitempty"`
	//
	NoticeId int32 `protobuf:"varint,2,opt,name=NoticeId,proto3" json:"NoticeId,omitempty"`
	//
	Tag int32 `protobuf:"varint,3,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CUpNoticeGang) Reset()         { *m = S2CUpNoticeGang{} }
func (m *S2CUpNoticeGang) String() string { return proto.CompactTextString(m) }
func (*S2CUpNoticeGang) ProtoMessage()    {}
func (m *S2CUpNoticeGang) GetId() uint16  { return PCK_S2CUpNoticeGang_ID }

const PCK_C2SReNameGang_ID = 1212 //
//
type C2SReNameGang struct {
	//
	GangId string `protobuf:"bytes,1,opt,name=GangId,proto3" json:"GangId,omitempty"`
	//
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (m *C2SReNameGang) Reset()         { *m = C2SReNameGang{} }
func (m *C2SReNameGang) String() string { return proto.CompactTextString(m) }
func (*C2SReNameGang) ProtoMessage()    {}
func (m *C2SReNameGang) GetId() uint16  { return PCK_C2SReNameGang_ID }

const PCK_S2CReNameGang_ID = 1213 //
//
type S2CReNameGang struct {
	//
	GangId string `protobuf:"bytes,1,opt,name=GangId,proto3" json:"GangId,omitempty"`
	//
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	//
	Tag int32 `protobuf:"varint,3,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CReNameGang) Reset()         { *m = S2CReNameGang{} }
func (m *S2CReNameGang) String() string { return proto.CompactTextString(m) }
func (*S2CReNameGang) ProtoMessage()    {}
func (m *S2CReNameGang) GetId() uint16  { return PCK_S2CReNameGang_ID }

//
type GangMember struct {
	//
	UseId int64 `protobuf:"varint,1,opt,name=UseId,proto3" json:"UseId,omitempty"`
	//
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	//
	JoinTime int64 `protobuf:"varint,3,opt,name=JoinTime,proto3" json:"JoinTime,omitempty"`
	//
	FightValue int64 `protobuf:"varint,4,opt,name=FightValue,proto3" json:"FightValue,omitempty"`
	//
	Contribution int64 `protobuf:"varint,5,opt,name=Contribution,proto3" json:"Contribution,omitempty"`
	//
	Status int32 `protobuf:"varint,6,opt,name=Status,proto3" json:"Status,omitempty"`
	//
	DutyId int32 `protobuf:"varint,7,opt,name=DutyId,proto3" json:"DutyId,omitempty"`
	//
	PreOutTime int64 `protobuf:"varint,8,opt,name=PreOutTime,proto3" json:"PreOutTime,omitempty"`
	//
	RoleId int64 `protobuf:"varint,9,opt,name=RoleId,proto3" json:"RoleId,omitempty"`
	//
	Level int64 `protobuf:"varint,10,opt,name=Level,proto3" json:"Level,omitempty"`
	//
	MonthCard int64 `protobuf:"varint,11,opt,name=MonthCard,proto3" json:"MonthCard,omitempty"`
	//
	YearCard int64 `protobuf:"varint,12,opt,name=YearCard,proto3" json:"YearCard,omitempty"`
	//
	HisD int64 `protobuf:"varint,13,opt,name=HisD,proto3" json:"HisD,omitempty"`
}

func (m *GangMember) Reset()         { *m = GangMember{} }
func (m *GangMember) String() string { return proto.CompactTextString(m) }
func (*GangMember) ProtoMessage()    {}

const PCK_C2SGangMember_ID = 1214 //
//
type C2SGangMember struct {
	//
	GangId string `protobuf:"bytes,1,opt,name=GangId,proto3" json:"GangId,omitempty"`
}

func (m *C2SGangMember) Reset()         { *m = C2SGangMember{} }
func (m *C2SGangMember) String() string { return proto.CompactTextString(m) }
func (*C2SGangMember) ProtoMessage()    {}
func (m *C2SGangMember) GetId() uint16  { return PCK_C2SGangMember_ID }

const PCK_S2CGangMember_ID = 1215 //
//
type S2CGangMember struct {
	//
	Members []*GangMember `protobuf:"bytes,1,rep,name=Members,proto3" json:"Members,omitempty"`
	//
	Tag int32 `protobuf:"varint,2,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CGangMember) Reset()         { *m = S2CGangMember{} }
func (m *S2CGangMember) String() string { return proto.CompactTextString(m) }
func (*S2CGangMember) ProtoMessage()    {}
func (m *S2CGangMember) GetId() uint16  { return PCK_S2CGangMember_ID }

const PCK_C2SGangApplys_ID = 1216 //
//
type C2SGangApplys struct {
	//
	GangId string `protobuf:"bytes,1,opt,name=GangId,proto3" json:"GangId,omitempty"`
}

func (m *C2SGangApplys) Reset()         { *m = C2SGangApplys{} }
func (m *C2SGangApplys) String() string { return proto.CompactTextString(m) }
func (*C2SGangApplys) ProtoMessage()    {}
func (m *C2SGangApplys) GetId() uint16  { return PCK_C2SGangApplys_ID }

const PCK_S2CGangApplys_ID = 1217 //
//
type S2CGangApplys struct {
	//
	Applys []*GangMember `protobuf:"bytes,1,rep,name=Applys,proto3" json:"Applys,omitempty"`
	//
	Tag int32 `protobuf:"varint,2,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CGangApplys) Reset()         { *m = S2CGangApplys{} }
func (m *S2CGangApplys) String() string { return proto.CompactTextString(m) }
func (*S2CGangApplys) ProtoMessage()    {}
func (m *S2CGangApplys) GetId() uint16  { return PCK_S2CGangApplys_ID }

const PCK_C2SGangKickMasterConsumes_ID = 1220 //
//
type C2SGangKickMasterConsumes struct {
}

func (m *C2SGangKickMasterConsumes) Reset()         { *m = C2SGangKickMasterConsumes{} }
func (m *C2SGangKickMasterConsumes) String() string { return proto.CompactTextString(m) }
func (*C2SGangKickMasterConsumes) ProtoMessage()    {}
func (m *C2SGangKickMasterConsumes) GetId() uint16  { return PCK_C2SGangKickMasterConsumes_ID }

const PCK_S2CGangKickMasterConsumes_ID = 1221 //
//
type S2CGangKickMasterConsumes struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Consumes string `protobuf:"bytes,2,opt,name=Consumes,proto3" json:"Consumes,omitempty"`
}

func (m *S2CGangKickMasterConsumes) Reset()         { *m = S2CGangKickMasterConsumes{} }
func (m *S2CGangKickMasterConsumes) String() string { return proto.CompactTextString(m) }
func (*S2CGangKickMasterConsumes) ProtoMessage()    {}
func (m *S2CGangKickMasterConsumes) GetId() uint16  { return PCK_S2CGangKickMasterConsumes_ID }

const PCK_C2SGangKickMaster_ID = 1231 //
//
type C2SGangKickMaster struct {
	//
	GangId string `protobuf:"bytes,1,opt,name=GangId,proto3" json:"GangId,omitempty"`
	//
	UserId int64 `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (m *C2SGangKickMaster) Reset()         { *m = C2SGangKickMaster{} }
func (m *C2SGangKickMaster) String() string { return proto.CompactTextString(m) }
func (*C2SGangKickMaster) ProtoMessage()    {}
func (m *C2SGangKickMaster) GetId() uint16  { return PCK_C2SGangKickMaster_ID }

const PCK_S2CGangKickMaster_ID = 1232 //
//
type S2CGangKickMaster struct {
	//
	GangId string `protobuf:"bytes,1,opt,name=GangId,proto3" json:"GangId,omitempty"`
	//
	GangMasterId int64 `protobuf:"varint,2,opt,name=GangMasterId,proto3" json:"GangMasterId,omitempty"`
	//
	Tag int32 `protobuf:"varint,3,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Consumes string `protobuf:"bytes,4,opt,name=Consumes,proto3" json:"Consumes,omitempty"`
	//
	Nums int32 `protobuf:"varint,5,opt,name=Nums,proto3" json:"Nums,omitempty"`
	//
	MasterName string `protobuf:"bytes,6,opt,name=MasterName,proto3" json:"MasterName,omitempty"`
}

func (m *S2CGangKickMaster) Reset()         { *m = S2CGangKickMaster{} }
func (m *S2CGangKickMaster) String() string { return proto.CompactTextString(m) }
func (*S2CGangKickMaster) ProtoMessage()    {}
func (m *S2CGangKickMaster) GetId() uint16  { return PCK_S2CGangKickMaster_ID }

const PCK_C2SUpGangFightVal_ID = 1218 //
//
type C2SUpGangFightVal struct {
	//
	GangId string `protobuf:"bytes,1,opt,name=GangId,proto3" json:"GangId,omitempty"`
	//
	FightValue int64 `protobuf:"varint,2,opt,name=FightValue,proto3" json:"FightValue,omitempty"`
}

func (m *C2SUpGangFightVal) Reset()         { *m = C2SUpGangFightVal{} }
func (m *C2SUpGangFightVal) String() string { return proto.CompactTextString(m) }
func (*C2SUpGangFightVal) ProtoMessage()    {}
func (m *C2SUpGangFightVal) GetId() uint16  { return PCK_C2SUpGangFightVal_ID }

const PCK_S2CUpGangFightVal_ID = 1219 //
//
type S2CUpGangFightVal struct {
	//
	GangId string `protobuf:"bytes,1,opt,name=GangId,proto3" json:"GangId,omitempty"`
	//
	FightValue int64 `protobuf:"varint,2,opt,name=FightValue,proto3" json:"FightValue,omitempty"`
	//
	Tag int32 `protobuf:"varint,3,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CUpGangFightVal) Reset()         { *m = S2CUpGangFightVal{} }
func (m *S2CUpGangFightVal) String() string { return proto.CompactTextString(m) }
func (*S2CUpGangFightVal) ProtoMessage()    {}
func (m *S2CUpGangFightVal) GetId() uint16  { return PCK_S2CUpGangFightVal_ID }

const PCK_C2SUpGangMoney_ID = 1222 //
//
type C2SUpGangMoney struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	GangId string `protobuf:"bytes,2,opt,name=GangId,proto3" json:"GangId,omitempty"`
}

func (m *C2SUpGangMoney) Reset()         { *m = C2SUpGangMoney{} }
func (m *C2SUpGangMoney) String() string { return proto.CompactTextString(m) }
func (*C2SUpGangMoney) ProtoMessage()    {}
func (m *C2SUpGangMoney) GetId() uint16  { return PCK_C2SUpGangMoney_ID }

const PCK_S2CUpGangMoney_ID = 1223 //
//
type S2CUpGangMoney struct {
	//
	Money int32 `protobuf:"varint,1,opt,name=Money,proto3" json:"Money,omitempty"`
	//
	GangId string `protobuf:"bytes,2,opt,name=GangId,proto3" json:"GangId,omitempty"`
	//
	Tag int32 `protobuf:"varint,3,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Level int32 `protobuf:"varint,4,opt,name=Level,proto3" json:"Level,omitempty"`
}

func (m *S2CUpGangMoney) Reset()         { *m = S2CUpGangMoney{} }
func (m *S2CUpGangMoney) String() string { return proto.CompactTextString(m) }
func (*S2CUpGangMoney) ProtoMessage()    {}
func (m *S2CUpGangMoney) GetId() uint16  { return PCK_S2CUpGangMoney_ID }

const PCK_C2SGangRecruit_ID = 1225 //
//
type C2SGangRecruit struct {
	//
	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (m *C2SGangRecruit) Reset()         { *m = C2SGangRecruit{} }
func (m *C2SGangRecruit) String() string { return proto.CompactTextString(m) }
func (*C2SGangRecruit) ProtoMessage()    {}
func (m *C2SGangRecruit) GetId() uint16  { return PCK_C2SGangRecruit_ID }

const PCK_S2CGangRecruit_ID = 1226 //
//
type S2CGangRecruit struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CGangRecruit) Reset()         { *m = S2CGangRecruit{} }
func (m *S2CGangRecruit) String() string { return proto.CompactTextString(m) }
func (*S2CGangRecruit) ProtoMessage()    {}
func (m *S2CGangRecruit) GetId() uint16  { return PCK_S2CGangRecruit_ID }

const PCK_C2SGangGuardupLevel_ID = 1234 //
//
type C2SGangGuardupLevel struct {
	//
	GangId string `protobuf:"bytes,1,opt,name=GangId,proto3" json:"GangId,omitempty"`
}

func (m *C2SGangGuardupLevel) Reset()         { *m = C2SGangGuardupLevel{} }
func (m *C2SGangGuardupLevel) String() string { return proto.CompactTextString(m) }
func (*C2SGangGuardupLevel) ProtoMessage()    {}
func (m *C2SGangGuardupLevel) GetId() uint16  { return PCK_C2SGangGuardupLevel_ID }

const PCK_S2CGangGuardupLevel_ID = 1235 //
//
type S2CGangGuardupLevel struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Level int32 `protobuf:"varint,2,opt,name=Level,proto3" json:"Level,omitempty"`
	//
	Experience int32 `protobuf:"varint,3,opt,name=Experience,proto3" json:"Experience,omitempty"`
	//
	Reward []*ItemInfo `protobuf:"bytes,4,rep,name=Reward,proto3" json:"Reward,omitempty"`
}

func (m *S2CGangGuardupLevel) Reset()         { *m = S2CGangGuardupLevel{} }
func (m *S2CGangGuardupLevel) String() string { return proto.CompactTextString(m) }
func (*S2CGangGuardupLevel) ProtoMessage()    {}
func (m *S2CGangGuardupLevel) GetId() uint16  { return PCK_S2CGangGuardupLevel_ID }

const PCK_C2SGangUpDownDuty_ID = 1227 //
//
type C2SGangUpDownDuty struct {
	//
	DutyId int32 `protobuf:"varint,1,opt,name=DutyId,proto3" json:"DutyId,omitempty"`
	//
	UserId int64 `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	GangId string `protobuf:"bytes,3,opt,name=GangId,proto3" json:"GangId,omitempty"`
}

func (m *C2SGangUpDownDuty) Reset()         { *m = C2SGangUpDownDuty{} }
func (m *C2SGangUpDownDuty) String() string { return proto.CompactTextString(m) }
func (*C2SGangUpDownDuty) ProtoMessage()    {}
func (m *C2SGangUpDownDuty) GetId() uint16  { return PCK_C2SGangUpDownDuty_ID }

const PCK_S2CGangUpDownDuty_ID = 1228 //
//
type S2CGangUpDownDuty struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	DutyId int32 `protobuf:"varint,2,opt,name=DutyId,proto3" json:"DutyId,omitempty"`
	//
	UserId int64 `protobuf:"varint,3,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (m *S2CGangUpDownDuty) Reset()         { *m = S2CGangUpDownDuty{} }
func (m *S2CGangUpDownDuty) String() string { return proto.CompactTextString(m) }
func (*S2CGangUpDownDuty) ProtoMessage()    {}
func (m *S2CGangUpDownDuty) GetId() uint16  { return PCK_S2CGangUpDownDuty_ID }

const PCK_C2SGangKick_ID = 1229 //
//
type C2SGangKick struct {
	//
	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	GangId string `protobuf:"bytes,2,opt,name=GangId,proto3" json:"GangId,omitempty"`
}

func (m *C2SGangKick) Reset()         { *m = C2SGangKick{} }
func (m *C2SGangKick) String() string { return proto.CompactTextString(m) }
func (*C2SGangKick) ProtoMessage()    {}
func (m *C2SGangKick) GetId() uint16  { return PCK_C2SGangKick_ID }

const PCK_S2CGangKick_ID = 1230 //
//
type S2CGangKick struct {
	//
	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	Tag int32 `protobuf:"varint,2,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Nums int32 `protobuf:"varint,3,opt,name=Nums,proto3" json:"Nums,omitempty"`
}

func (m *S2CGangKick) Reset()         { *m = S2CGangKick{} }
func (m *S2CGangKick) String() string { return proto.CompactTextString(m) }
func (*S2CGangKick) ProtoMessage()    {}
func (m *S2CGangKick) GetId() uint16  { return PCK_S2CGangKick_ID }

//
type GangAction struct {
	//
	Type int32 `protobuf:"varint,1,opt,name=Type,proto3" json:"Type,omitempty"`
	//
	UserId int64 `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	Name string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	//
	ActionTime int64 `protobuf:"varint,4,opt,name=ActionTime,proto3" json:"ActionTime,omitempty"`
	//
	DutyId int32 `protobuf:"varint,5,opt,name=DutyId,proto3" json:"DutyId,omitempty"`
	//
	Param string `protobuf:"bytes,6,opt,name=Param,proto3" json:"Param,omitempty"`
}

func (m *GangAction) Reset()         { *m = GangAction{} }
func (m *GangAction) String() string { return proto.CompactTextString(m) }
func (*GangAction) ProtoMessage()    {}

const PCK_C2SGangRecord_ID = 1242 //
//
type C2SGangRecord struct {
	//
	GangId string `protobuf:"bytes,1,opt,name=GangId,proto3" json:"GangId,omitempty"`
}

func (m *C2SGangRecord) Reset()         { *m = C2SGangRecord{} }
func (m *C2SGangRecord) String() string { return proto.CompactTextString(m) }
func (*C2SGangRecord) ProtoMessage()    {}
func (m *C2SGangRecord) GetId() uint16  { return PCK_C2SGangRecord_ID }

const PCK_S2CGangRecord_ID = 1243 //
//
type S2CGangRecord struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Actions []*GangAction `protobuf:"bytes,2,rep,name=Actions,proto3" json:"Actions,omitempty"`
}

func (m *S2CGangRecord) Reset()         { *m = S2CGangRecord{} }
func (m *S2CGangRecord) String() string { return proto.CompactTextString(m) }
func (*S2CGangRecord) ProtoMessage()    {}
func (m *S2CGangRecord) GetId() uint16  { return PCK_S2CGangRecord_ID }

const PCK_C2SGangExit_ID = 1236 //
//
type C2SGangExit struct {
	//
	GangId string `protobuf:"bytes,1,opt,name=GangId,proto3" json:"GangId,omitempty"`
}

func (m *C2SGangExit) Reset()         { *m = C2SGangExit{} }
func (m *C2SGangExit) String() string { return proto.CompactTextString(m) }
func (*C2SGangExit) ProtoMessage()    {}
func (m *C2SGangExit) GetId() uint16  { return PCK_C2SGangExit_ID }

const PCK_S2CGangExit_ID = 1237 //
//
type S2CGangExit struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Nums int32 `protobuf:"varint,2,opt,name=Nums,proto3" json:"Nums,omitempty"`
}

func (m *S2CGangExit) Reset()         { *m = S2CGangExit{} }
func (m *S2CGangExit) String() string { return proto.CompactTextString(m) }
func (*S2CGangExit) ProtoMessage()    {}
func (m *S2CGangExit) GetId() uint16  { return PCK_S2CGangExit_ID }

const PCK_C2SGangEnterMap_ID = 1238 //
//
type C2SGangEnterMap struct {
}

func (m *C2SGangEnterMap) Reset()         { *m = C2SGangEnterMap{} }
func (m *C2SGangEnterMap) String() string { return proto.CompactTextString(m) }
func (*C2SGangEnterMap) ProtoMessage()    {}
func (m *C2SGangEnterMap) GetId() uint16  { return PCK_C2SGangEnterMap_ID }

const PCK_S2CGangEnterMap_ID = 1239 //
//
type S2CGangEnterMap struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CGangEnterMap) Reset()         { *m = S2CGangEnterMap{} }
func (m *S2CGangEnterMap) String() string { return proto.CompactTextString(m) }
func (*S2CGangEnterMap) ProtoMessage()    {}
func (m *S2CGangEnterMap) GetId() uint16  { return PCK_S2CGangEnterMap_ID }

const PCK_C2SGangBossEnterMap_ID = 1281 //
//
type C2SGangBossEnterMap struct {
}

func (m *C2SGangBossEnterMap) Reset()         { *m = C2SGangBossEnterMap{} }
func (m *C2SGangBossEnterMap) String() string { return proto.CompactTextString(m) }
func (*C2SGangBossEnterMap) ProtoMessage()    {}
func (m *C2SGangBossEnterMap) GetId() uint16  { return PCK_C2SGangBossEnterMap_ID }

const PCK_S2CGangBossEnterMap_ID = 1282 //
//
type S2CGangBossEnterMap struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CGangBossEnterMap) Reset()         { *m = S2CGangBossEnterMap{} }
func (m *S2CGangBossEnterMap) String() string { return proto.CompactTextString(m) }
func (*S2CGangBossEnterMap) ProtoMessage()    {}
func (m *S2CGangBossEnterMap) GetId() uint16  { return PCK_S2CGangBossEnterMap_ID }

const PCK_C2SGangLeaveMap_ID = 1240 //
//
type C2SGangLeaveMap struct {
}

func (m *C2SGangLeaveMap) Reset()         { *m = C2SGangLeaveMap{} }
func (m *C2SGangLeaveMap) String() string { return proto.CompactTextString(m) }
func (*C2SGangLeaveMap) ProtoMessage()    {}
func (m *C2SGangLeaveMap) GetId() uint16  { return PCK_C2SGangLeaveMap_ID }

const PCK_S2CGangLeaveMap_ID = 1241 //
//
type S2CGangLeaveMap struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CGangLeaveMap) Reset()         { *m = S2CGangLeaveMap{} }
func (m *S2CGangLeaveMap) String() string { return proto.CompactTextString(m) }
func (*S2CGangLeaveMap) ProtoMessage()    {}
func (m *S2CGangLeaveMap) GetId() uint16  { return PCK_S2CGangLeaveMap_ID }

const PCK_C2SGangGuardReward_ID = 1244 //
//
type C2SGangGuardReward struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	GangId string `protobuf:"bytes,2,opt,name=GangId,proto3" json:"GangId,omitempty"`
}

func (m *C2SGangGuardReward) Reset()         { *m = C2SGangGuardReward{} }
func (m *C2SGangGuardReward) String() string { return proto.CompactTextString(m) }
func (*C2SGangGuardReward) ProtoMessage()    {}
func (m *C2SGangGuardReward) GetId() uint16  { return PCK_C2SGangGuardReward_ID }

const PCK_S2CGangGuardReward_ID = 1245 //
//
type S2CGangGuardReward struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	ReceiverId []int32 `protobuf:"varint,2,rep,name=ReceiverId,proto3" json:"ReceiverId,omitempty"`
	//
	Id int32 `protobuf:"varint,3,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *S2CGangGuardReward) Reset()         { *m = S2CGangGuardReward{} }
func (m *S2CGangGuardReward) String() string { return proto.CompactTextString(m) }
func (*S2CGangGuardReward) ProtoMessage()    {}
func (m *S2CGangGuardReward) GetId() uint16  { return PCK_S2CGangGuardReward_ID }

const PCK_C2SGangFindReward_ID = 1246 //
//
type C2SGangFindReward struct {
	//
	GangId string `protobuf:"bytes,1,opt,name=GangId,proto3" json:"GangId,omitempty"`
}

func (m *C2SGangFindReward) Reset()         { *m = C2SGangFindReward{} }
func (m *C2SGangFindReward) String() string { return proto.CompactTextString(m) }
func (*C2SGangFindReward) ProtoMessage()    {}
func (m *C2SGangFindReward) GetId() uint16  { return PCK_C2SGangFindReward_ID }

const PCK_S2CGangFindReward_ID = 1247 //
//
type S2CGangFindReward struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	ReceiverId []int32 `protobuf:"varint,2,rep,name=ReceiverId,proto3" json:"ReceiverId,omitempty"`
}

func (m *S2CGangFindReward) Reset()         { *m = S2CGangFindReward{} }
func (m *S2CGangFindReward) String() string { return proto.CompactTextString(m) }
func (*S2CGangFindReward) ProtoMessage()    {}
func (m *S2CGangFindReward) GetId() uint16  { return PCK_S2CGangFindReward_ID }

const PCK_C2SGangUpSkill_ID = 1248 //
//
type C2SGangUpSkill struct {
	//
	GangId string `protobuf:"bytes,1,opt,name=GangId,proto3" json:"GangId,omitempty"`
	//
	SkillId int32 `protobuf:"varint,2,opt,name=SkillId,proto3" json:"SkillId,omitempty"`
}

func (m *C2SGangUpSkill) Reset()         { *m = C2SGangUpSkill{} }
func (m *C2SGangUpSkill) String() string { return proto.CompactTextString(m) }
func (*C2SGangUpSkill) ProtoMessage()    {}
func (m *C2SGangUpSkill) GetId() uint16  { return PCK_C2SGangUpSkill_ID }

const PCK_S2CGangUpSkill_ID = 1249 //
//
type S2CGangUpSkill struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	SkillId int32 `protobuf:"varint,2,opt,name=SkillId,proto3" json:"SkillId,omitempty"`
}

func (m *S2CGangUpSkill) Reset()         { *m = S2CGangUpSkill{} }
func (m *S2CGangUpSkill) String() string { return proto.CompactTextString(m) }
func (*S2CGangUpSkill) ProtoMessage()    {}
func (m *S2CGangUpSkill) GetId() uint16  { return PCK_S2CGangUpSkill_ID }

const PCK_C2SGangFindSkill_ID = 1250 //
//
type C2SGangFindSkill struct {
	//
	GangId string `protobuf:"bytes,1,opt,name=GangId,proto3" json:"GangId,omitempty"`
}

func (m *C2SGangFindSkill) Reset()         { *m = C2SGangFindSkill{} }
func (m *C2SGangFindSkill) String() string { return proto.CompactTextString(m) }
func (*C2SGangFindSkill) ProtoMessage()    {}
func (m *C2SGangFindSkill) GetId() uint16  { return PCK_C2SGangFindSkill_ID }

const PCK_S2CGangFindSkill_ID = 1251 //
//
type S2CGangFindSkill struct {
	//
	CurSkill int32 `protobuf:"varint,1,opt,name=CurSkill,proto3" json:"CurSkill,omitempty"`
	//
	CurHeiSkill int32 `protobuf:"varint,2,opt,name=CurHeiSkill,proto3" json:"CurHeiSkill,omitempty"`
	//
	Skills string `protobuf:"bytes,3,opt,name=Skills,proto3" json:"Skills,omitempty"`
	//
	HeiSkills string `protobuf:"bytes,4,opt,name=HeiSkills,proto3" json:"HeiSkills,omitempty"`
}

func (m *S2CGangFindSkill) Reset()         { *m = S2CGangFindSkill{} }
func (m *S2CGangFindSkill) String() string { return proto.CompactTextString(m) }
func (*S2CGangFindSkill) ProtoMessage()    {}
func (m *S2CGangFindSkill) GetId() uint16  { return PCK_S2CGangFindSkill_ID }

//
type GangPeachRecord struct {
	//
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	//
	PeachId int32 `protobuf:"varint,2,opt,name=PeachId,proto3" json:"PeachId,omitempty"`
	//
	EatTime int64 `protobuf:"varint,3,opt,name=EatTime,proto3" json:"EatTime,omitempty"`
}

func (m *GangPeachRecord) Reset()         { *m = GangPeachRecord{} }
func (m *GangPeachRecord) String() string { return proto.CompactTextString(m) }
func (*GangPeachRecord) ProtoMessage()    {}

const PCK_C2SGangFindPeach_ID = 1258 //
//
type C2SGangFindPeach struct {
	//
	GangId string `protobuf:"bytes,1,opt,name=GangId,proto3" json:"GangId,omitempty"`
}

func (m *C2SGangFindPeach) Reset()         { *m = C2SGangFindPeach{} }
func (m *C2SGangFindPeach) String() string { return proto.CompactTextString(m) }
func (*C2SGangFindPeach) ProtoMessage()    {}
func (m *C2SGangFindPeach) GetId() uint16  { return PCK_C2SGangFindPeach_ID }

const PCK_S2CGangFindPeach_ID = 1259 //
//
type S2CGangFindPeach struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Records []*GangPeachRecord `protobuf:"bytes,2,rep,name=Records,proto3" json:"Records,omitempty"`
	//
	Boxs string `protobuf:"bytes,3,opt,name=Boxs,proto3" json:"Boxs,omitempty"`
	//
	Eats string `protobuf:"bytes,4,opt,name=Eats,proto3" json:"Eats,omitempty"`
	//
	GangPeachVal int64 `protobuf:"varint,5,opt,name=GangPeachVal,proto3" json:"GangPeachVal,omitempty"`
}

func (m *S2CGangFindPeach) Reset()         { *m = S2CGangFindPeach{} }
func (m *S2CGangFindPeach) String() string { return proto.CompactTextString(m) }
func (*S2CGangFindPeach) ProtoMessage()    {}
func (m *S2CGangFindPeach) GetId() uint16  { return PCK_S2CGangFindPeach_ID }

const PCK_C2SGangEatPeach_ID = 1260 //
//
type C2SGangEatPeach struct {
	//
	GangId string `protobuf:"bytes,1,opt,name=GangId,proto3" json:"GangId,omitempty"`
	//
	PeachId int32 `protobuf:"varint,2,opt,name=PeachId,proto3" json:"PeachId,omitempty"`
}

func (m *C2SGangEatPeach) Reset()         { *m = C2SGangEatPeach{} }
func (m *C2SGangEatPeach) String() string { return proto.CompactTextString(m) }
func (*C2SGangEatPeach) ProtoMessage()    {}
func (m *C2SGangEatPeach) GetId() uint16  { return PCK_C2SGangEatPeach_ID }

const PCK_S2CGangEatPeach_ID = 1261 //
//
type S2CGangEatPeach struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Records []*GangPeachRecord `protobuf:"bytes,2,rep,name=Records,proto3" json:"Records,omitempty"`
	//
	PeachId int32 `protobuf:"varint,3,opt,name=PeachId,proto3" json:"PeachId,omitempty"`
	//
	Eats string `protobuf:"bytes,4,opt,name=Eats,proto3" json:"Eats,omitempty"`
	//
	GangPeachVal int64 `protobuf:"varint,5,opt,name=GangPeachVal,proto3" json:"GangPeachVal,omitempty"`
	//
	GM int32 `protobuf:"varint,6,opt,name=GM,proto3" json:"GM,omitempty"`
}

func (m *S2CGangEatPeach) Reset()         { *m = S2CGangEatPeach{} }
func (m *S2CGangEatPeach) String() string { return proto.CompactTextString(m) }
func (*S2CGangEatPeach) ProtoMessage()    {}
func (m *S2CGangEatPeach) GetId() uint16  { return PCK_S2CGangEatPeach_ID }

const PCK_C2SGangPeachBox_ID = 1262 //
//
type C2SGangPeachBox struct {
	//
	GangId string `protobuf:"bytes,1,opt,name=GangId,proto3" json:"GangId,omitempty"`
	//
	BoxId int32 `protobuf:"varint,2,opt,name=BoxId,proto3" json:"BoxId,omitempty"`
}

func (m *C2SGangPeachBox) Reset()         { *m = C2SGangPeachBox{} }
func (m *C2SGangPeachBox) String() string { return proto.CompactTextString(m) }
func (*C2SGangPeachBox) ProtoMessage()    {}
func (m *C2SGangPeachBox) GetId() uint16  { return PCK_C2SGangPeachBox_ID }

const PCK_S2CGangPeachBox_ID = 1263 //
//
type S2CGangPeachBox struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	BoxId int32 `protobuf:"varint,2,opt,name=BoxId,proto3" json:"BoxId,omitempty"`
	//
	Boxs string `protobuf:"bytes,3,opt,name=Boxs,proto3" json:"Boxs,omitempty"`
}

func (m *S2CGangPeachBox) Reset()         { *m = S2CGangPeachBox{} }
func (m *S2CGangPeachBox) String() string { return proto.CompactTextString(m) }
func (*S2CGangPeachBox) ProtoMessage()    {}
func (m *S2CGangPeachBox) GetId() uint16  { return PCK_S2CGangPeachBox_ID }

const PCK_C2SGangAutoJoin_ID = 1264 //
//
type C2SGangAutoJoin struct {
	//
	GangId string `protobuf:"bytes,1,opt,name=GangId,proto3" json:"GangId,omitempty"`
	//
	ActionType int32 `protobuf:"varint,2,opt,name=ActionType,proto3" json:"ActionType,omitempty"`
}

func (m *C2SGangAutoJoin) Reset()         { *m = C2SGangAutoJoin{} }
func (m *C2SGangAutoJoin) String() string { return proto.CompactTextString(m) }
func (*C2SGangAutoJoin) ProtoMessage()    {}
func (m *C2SGangAutoJoin) GetId() uint16  { return PCK_C2SGangAutoJoin_ID }

const PCK_S2CGangAutoJoin_ID = 1265 //
//
type S2CGangAutoJoin struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	GangId string `protobuf:"bytes,2,opt,name=GangId,proto3" json:"GangId,omitempty"`
	//
	ActionType int32 `protobuf:"varint,3,opt,name=ActionType,proto3" json:"ActionType,omitempty"`
}

func (m *S2CGangAutoJoin) Reset()         { *m = S2CGangAutoJoin{} }
func (m *S2CGangAutoJoin) String() string { return proto.CompactTextString(m) }
func (*S2CGangAutoJoin) ProtoMessage()    {}
func (m *S2CGangAutoJoin) GetId() uint16  { return PCK_S2CGangAutoJoin_ID }

const PCK_C2SGangMapOnekeyOver_ID = 1266 //
//
type C2SGangMapOnekeyOver struct {
	//
	GangId string `protobuf:"bytes,1,opt,name=GangId,proto3" json:"GangId,omitempty"`
	//
	TaskId int32 `protobuf:"varint,2,opt,name=TaskId,proto3" json:"TaskId,omitempty"`
}

func (m *C2SGangMapOnekeyOver) Reset()         { *m = C2SGangMapOnekeyOver{} }
func (m *C2SGangMapOnekeyOver) String() string { return proto.CompactTextString(m) }
func (*C2SGangMapOnekeyOver) ProtoMessage()    {}
func (m *C2SGangMapOnekeyOver) GetId() uint16  { return PCK_C2SGangMapOnekeyOver_ID }

const PCK_S2CGangMapOneKeyOver_ID = 1267 //
//
type S2CGangMapOneKeyOver struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	TaskId int32 `protobuf:"varint,2,opt,name=TaskId,proto3" json:"TaskId,omitempty"`
}

func (m *S2CGangMapOneKeyOver) Reset()         { *m = S2CGangMapOneKeyOver{} }
func (m *S2CGangMapOneKeyOver) String() string { return proto.CompactTextString(m) }
func (*S2CGangMapOneKeyOver) ProtoMessage()    {}
func (m *S2CGangMapOneKeyOver) GetId() uint16  { return PCK_S2CGangMapOneKeyOver_ID }

const PCK_C2SGangMapResetInfo_ID = 1294 //
//
type C2SGangMapResetInfo struct {
	//
	GangId string `protobuf:"bytes,1,opt,name=GangId,proto3" json:"GangId,omitempty"`
	//
	TaskId int32 `protobuf:"varint,2,opt,name=TaskId,proto3" json:"TaskId,omitempty"`
}

func (m *C2SGangMapResetInfo) Reset()         { *m = C2SGangMapResetInfo{} }
func (m *C2SGangMapResetInfo) String() string { return proto.CompactTextString(m) }
func (*C2SGangMapResetInfo) ProtoMessage()    {}
func (m *C2SGangMapResetInfo) GetId() uint16  { return PCK_C2SGangMapResetInfo_ID }

const PCK_S2CGangMapResetInfo_ID = 1295 //
//
type S2CGangMapResetInfo struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	LeftTimes int32 `protobuf:"varint,2,opt,name=LeftTimes,proto3" json:"LeftTimes,omitempty"`
	//
	NextVipTimes int32 `protobuf:"varint,3,opt,name=NextVipTimes,proto3" json:"NextVipTimes,omitempty"`
	//
	NextVip int32 `protobuf:"varint,4,opt,name=NextVip,proto3" json:"NextVip,omitempty"`
	//
	TaskId int32 `protobuf:"varint,5,opt,name=TaskId,proto3" json:"TaskId,omitempty"`
}

func (m *S2CGangMapResetInfo) Reset()         { *m = S2CGangMapResetInfo{} }
func (m *S2CGangMapResetInfo) String() string { return proto.CompactTextString(m) }
func (*S2CGangMapResetInfo) ProtoMessage()    {}
func (m *S2CGangMapResetInfo) GetId() uint16  { return PCK_S2CGangMapResetInfo_ID }

const PCK_C2SGangMapReset_ID = 1268 //
//
type C2SGangMapReset struct {
	//
	GangId string `protobuf:"bytes,1,opt,name=GangId,proto3" json:"GangId,omitempty"`
	//
	TaskId int32 `protobuf:"varint,2,opt,name=TaskId,proto3" json:"TaskId,omitempty"`
}

func (m *C2SGangMapReset) Reset()         { *m = C2SGangMapReset{} }
func (m *C2SGangMapReset) String() string { return proto.CompactTextString(m) }
func (*C2SGangMapReset) ProtoMessage()    {}
func (m *C2SGangMapReset) GetId() uint16  { return PCK_C2SGangMapReset_ID }

const PCK_S2CGangMapReset_ID = 1269 //
//
type S2CGangMapReset struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	TaskId int32 `protobuf:"varint,2,opt,name=TaskId,proto3" json:"TaskId,omitempty"`
}

func (m *S2CGangMapReset) Reset()         { *m = S2CGangMapReset{} }
func (m *S2CGangMapReset) String() string { return proto.CompactTextString(m) }
func (*S2CGangMapReset) ProtoMessage()    {}
func (m *S2CGangMapReset) GetId() uint16  { return PCK_S2CGangMapReset_ID }

//
type RedPkgRecords struct {
	//
	Rds []*RedPkgRecord `protobuf:"bytes,1,rep,name=Rds,proto3" json:"Rds,omitempty"`
}

func (m *RedPkgRecords) Reset()         { *m = RedPkgRecords{} }
func (m *RedPkgRecords) String() string { return proto.CompactTextString(m) }
func (*RedPkgRecords) ProtoMessage()    {}

//
type RedPkgRecord struct {
	//
	SId int64 `protobuf:"varint,1,opt,name=SId,proto3" json:"SId,omitempty"`
	//
	M int64 `protobuf:"varint,2,opt,name=M,proto3" json:"M,omitempty"`
	//
	T int32 `protobuf:"varint,3,opt,name=T,proto3" json:"T,omitempty"`
	//
	Id string `protobuf:"bytes,4,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	RId int64 `protobuf:"varint,5,opt,name=RId,proto3" json:"RId,omitempty"`
	//
	RT int32 `protobuf:"varint,6,opt,name=RT,proto3" json:"RT,omitempty"`
	//
	UT int64 `protobuf:"varint,7,opt,name=UT,proto3" json:"UT,omitempty"`
}

func (m *RedPkgRecord) Reset()         { *m = RedPkgRecord{} }
func (m *RedPkgRecord) String() string { return proto.CompactTextString(m) }
func (*RedPkgRecord) ProtoMessage()    {}

const PCK_C2SRedPkgPool_ID = 1270 //
//
type C2SRedPkgPool struct {
}

func (m *C2SRedPkgPool) Reset()         { *m = C2SRedPkgPool{} }
func (m *C2SRedPkgPool) String() string { return proto.CompactTextString(m) }
func (*C2SRedPkgPool) ProtoMessage()    {}
func (m *C2SRedPkgPool) GetId() uint16  { return PCK_C2SRedPkgPool_ID }

const PCK_S2CRedPkgPool_ID = 1271 //
//
type S2CRedPkgPool struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	LC int32 `protobuf:"varint,2,opt,name=LC,proto3" json:"LC,omitempty"`
	//
	M int64 `protobuf:"varint,3,opt,name=M,proto3" json:"M,omitempty"`
}

func (m *S2CRedPkgPool) Reset()         { *m = S2CRedPkgPool{} }
func (m *S2CRedPkgPool) String() string { return proto.CompactTextString(m) }
func (*S2CRedPkgPool) ProtoMessage()    {}
func (m *S2CRedPkgPool) GetId() uint16  { return PCK_S2CRedPkgPool_ID }

const PCK_C2SRedPkgSend_ID = 1272 //
//
type C2SRedPkgSend struct {
	//
	T int32 `protobuf:"varint,1,opt,name=T,proto3" json:"T,omitempty"`
	//
	M int64 `protobuf:"varint,2,opt,name=M,proto3" json:"M,omitempty"`
	//
	C int32 `protobuf:"varint,3,opt,name=C,proto3" json:"C,omitempty"`
	//
	Des string `protobuf:"bytes,4,opt,name=Des,proto3" json:"Des,omitempty"`
}

func (m *C2SRedPkgSend) Reset()         { *m = C2SRedPkgSend{} }
func (m *C2SRedPkgSend) String() string { return proto.CompactTextString(m) }
func (*C2SRedPkgSend) ProtoMessage()    {}
func (m *C2SRedPkgSend) GetId() uint16  { return PCK_C2SRedPkgSend_ID }

const PCK_S2CRedPkgSend_ID = 1273 //
//
type S2CRedPkgSend struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CRedPkgSend) Reset()         { *m = S2CRedPkgSend{} }
func (m *S2CRedPkgSend) String() string { return proto.CompactTextString(m) }
func (*S2CRedPkgSend) ProtoMessage()    {}
func (m *S2CRedPkgSend) GetId() uint16  { return PCK_S2CRedPkgSend_ID }

const PCK_S2CRedPkgNotice_ID = 1274 //
//
type S2CRedPkgNotice struct {
	//
	UN string `protobuf:"bytes,1,opt,name=UN,proto3" json:"UN,omitempty"`
	//
	ET int64 `protobuf:"varint,2,opt,name=ET,proto3" json:"ET,omitempty"`
	//
	RId string `protobuf:"bytes,3,opt,name=RId,proto3" json:"RId,omitempty"`
	//
	URId int32 `protobuf:"varint,4,opt,name=URId,proto3" json:"URId,omitempty"`
	//
	RT int32 `protobuf:"varint,5,opt,name=RT,proto3" json:"RT,omitempty"`
	//
	Money int32 `protobuf:"varint,6,opt,name=Money,proto3" json:"Money,omitempty"`
}

func (m *S2CRedPkgNotice) Reset()         { *m = S2CRedPkgNotice{} }
func (m *S2CRedPkgNotice) String() string { return proto.CompactTextString(m) }
func (*S2CRedPkgNotice) ProtoMessage()    {}
func (m *S2CRedPkgNotice) GetId() uint16  { return PCK_S2CRedPkgNotice_ID }

const PCK_C2SRedPkgRecieve_ID = 1275 //
//
type C2SRedPkgRecieve struct {
	//
	RId string `protobuf:"bytes,1,opt,name=RId,proto3" json:"RId,omitempty"`
	//
	RedType int32 `protobuf:"varint,2,opt,name=RedType,proto3" json:"RedType,omitempty"`
}

func (m *C2SRedPkgRecieve) Reset()         { *m = C2SRedPkgRecieve{} }
func (m *C2SRedPkgRecieve) String() string { return proto.CompactTextString(m) }
func (*C2SRedPkgRecieve) ProtoMessage()    {}
func (m *C2SRedPkgRecieve) GetId() uint16  { return PCK_C2SRedPkgRecieve_ID }

const PCK_S2CRedPkgRecieve_ID = 1276 //
//
type S2CRedPkgRecieve struct {
	//
	RId string `protobuf:"bytes,1,opt,name=RId,proto3" json:"RId,omitempty"`
	//
	M int64 `protobuf:"varint,4,opt,name=M,proto3" json:"M,omitempty"`
	//
	Tag int32 `protobuf:"varint,5,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Desc string `protobuf:"bytes,6,opt,name=Desc,proto3" json:"Desc,omitempty"`
	//
	T int32 `protobuf:"varint,7,opt,name=T,proto3" json:"T,omitempty"`
}

func (m *S2CRedPkgRecieve) Reset()         { *m = S2CRedPkgRecieve{} }
func (m *S2CRedPkgRecieve) String() string { return proto.CompactTextString(m) }
func (*S2CRedPkgRecieve) ProtoMessage()    {}
func (m *S2CRedPkgRecieve) GetId() uint16  { return PCK_S2CRedPkgRecieve_ID }

const PCK_C2SRedPkgSingle_ID = 1277 //
//
type C2SRedPkgSingle struct {
	//
	RId string `protobuf:"bytes,1,opt,name=RId,proto3" json:"RId,omitempty"`
}

func (m *C2SRedPkgSingle) Reset()         { *m = C2SRedPkgSingle{} }
func (m *C2SRedPkgSingle) String() string { return proto.CompactTextString(m) }
func (*C2SRedPkgSingle) ProtoMessage()    {}
func (m *C2SRedPkgSingle) GetId() uint16  { return PCK_C2SRedPkgSingle_ID }

const PCK_S2CRedPkgSingle_ID = 1278 //
//
type S2CRedPkgSingle struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	UN string `protobuf:"bytes,2,opt,name=UN,proto3" json:"UN,omitempty"`
	//
	URId int32 `protobuf:"varint,3,opt,name=URId,proto3" json:"URId,omitempty"`
	//
	Rds []*RedPkgShowRecord `protobuf:"bytes,4,rep,name=Rds,proto3" json:"Rds,omitempty"`
	//
	AM int64 `protobuf:"varint,5,opt,name=AM,proto3" json:"AM,omitempty"`
	//
	Desc string `protobuf:"bytes,6,opt,name=Desc,proto3" json:"Desc,omitempty"`
}

func (m *S2CRedPkgSingle) Reset()         { *m = S2CRedPkgSingle{} }
func (m *S2CRedPkgSingle) String() string { return proto.CompactTextString(m) }
func (*S2CRedPkgSingle) ProtoMessage()    {}
func (m *S2CRedPkgSingle) GetId() uint16  { return PCK_S2CRedPkgSingle_ID }

//
type RedPkgShowRecord struct {
	//
	SN string `protobuf:"bytes,1,opt,name=SN,proto3" json:"SN,omitempty"`
	//
	M int64 `protobuf:"varint,2,opt,name=M,proto3" json:"M,omitempty"`
	//
	T int32 `protobuf:"varint,3,opt,name=T,proto3" json:"T,omitempty"`
	//
	RT int32 `protobuf:"varint,4,opt,name=RT,proto3" json:"RT,omitempty"`
	//
	RN string `protobuf:"bytes,5,opt,name=RN,proto3" json:"RN,omitempty"`
	//
	RId int32 `protobuf:"varint,6,opt,name=RId,proto3" json:"RId,omitempty"`
	//
	SRId int32 `protobuf:"varint,7,opt,name=SRId,proto3" json:"SRId,omitempty"`
	//
	RedId string `protobuf:"bytes,8,opt,name=RedId,proto3" json:"RedId,omitempty"`
	//
	RUId int64 `protobuf:"varint,9,opt,name=RUId,proto3" json:"RUId,omitempty"`
	//
	UT int64 `protobuf:"varint,10,opt,name=UT,proto3" json:"UT,omitempty"`
}

func (m *RedPkgShowRecord) Reset()         { *m = RedPkgShowRecord{} }
func (m *RedPkgShowRecord) String() string { return proto.CompactTextString(m) }
func (*RedPkgShowRecord) ProtoMessage()    {}

const PCK_C2SRedPkgPerRecord_ID = 1279 //
//
type C2SRedPkgPerRecord struct {
}

func (m *C2SRedPkgPerRecord) Reset()         { *m = C2SRedPkgPerRecord{} }
func (m *C2SRedPkgPerRecord) String() string { return proto.CompactTextString(m) }
func (*C2SRedPkgPerRecord) ProtoMessage()    {}
func (m *C2SRedPkgPerRecord) GetId() uint16  { return PCK_C2SRedPkgPerRecord_ID }

const PCK_S2CRedPkgPerRecord_ID = 1280 //
//
type S2CRedPkgPerRecord struct {
	//
	Rds []*RedPkgShowRecord `protobuf:"bytes,1,rep,name=Rds,proto3" json:"Rds,omitempty"`
}

func (m *S2CRedPkgPerRecord) Reset()         { *m = S2CRedPkgPerRecord{} }
func (m *S2CRedPkgPerRecord) String() string { return proto.CompactTextString(m) }
func (*S2CRedPkgPerRecord) ProtoMessage()    {}
func (m *S2CRedPkgPerRecord) GetId() uint16  { return PCK_S2CRedPkgPerRecord_ID }

//
type RedPckLog struct {
	//
	Type int32 `protobuf:"varint,1,opt,name=Type,proto3" json:"Type,omitempty"`
	//
	Money int64 `protobuf:"varint,2,opt,name=Money,proto3" json:"Money,omitempty"`
	//
	Time int64 `protobuf:"varint,3,opt,name=Time,proto3" json:"Time,omitempty"`
}

func (m *RedPckLog) Reset()         { *m = RedPckLog{} }
func (m *RedPckLog) String() string { return proto.CompactTextString(m) }
func (*RedPckLog) ProtoMessage()    {}

const PCK_C2SRedPckShieldLog_ID = 1297 //
//
type C2SRedPckShieldLog struct {
}

func (m *C2SRedPckShieldLog) Reset()         { *m = C2SRedPckShieldLog{} }
func (m *C2SRedPckShieldLog) String() string { return proto.CompactTextString(m) }
func (*C2SRedPckShieldLog) ProtoMessage()    {}
func (m *C2SRedPckShieldLog) GetId() uint16  { return PCK_C2SRedPckShieldLog_ID }

const PCK_S2CRedPckShieldLog_ID = 1298 //
//
type S2CRedPckShieldLog struct {
	//
	Logs []*RedPckLog `protobuf:"bytes,1,rep,name=Logs,proto3" json:"Logs,omitempty"`
}

func (m *S2CRedPckShieldLog) Reset()         { *m = S2CRedPckShieldLog{} }
func (m *S2CRedPckShieldLog) String() string { return proto.CompactTextString(m) }
func (*S2CRedPckShieldLog) ProtoMessage()    {}
func (m *S2CRedPckShieldLog) GetId() uint16  { return PCK_S2CRedPckShieldLog_ID }

//
type GangSimpleInfo struct {
	//
	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	//
	Level int64 `protobuf:"varint,3,opt,name=Level,proto3" json:"Level,omitempty"`
	//
	MasterId int64 `protobuf:"varint,4,opt,name=MasterId,proto3" json:"MasterId,omitempty"`
	//
	MasterName string `protobuf:"bytes,5,opt,name=MasterName,proto3" json:"MasterName,omitempty"`
	//
	CreateServerId int64 `protobuf:"varint,6,opt,name=CreateServerId,proto3" json:"CreateServerId,omitempty"`
	//
	CreateAreaId int64 `protobuf:"varint,7,opt,name=CreateAreaId,proto3" json:"CreateAreaId,omitempty"`
	//
	CurrServerId int64 `protobuf:"varint,8,opt,name=CurrServerId,proto3" json:"CurrServerId,omitempty"`
}

func (m *GangSimpleInfo) Reset()         { *m = GangSimpleInfo{} }
func (m *GangSimpleInfo) String() string { return proto.CompactTextString(m) }
func (*GangSimpleInfo) ProtoMessage()    {}

const PCK_S2KGangSimpleInfo_ID = 1296 //
//
type S2KGangSimpleInfo struct {
	//
	Info *GangSimpleInfo `protobuf:"bytes,1,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (m *S2KGangSimpleInfo) Reset()         { *m = S2KGangSimpleInfo{} }
func (m *S2KGangSimpleInfo) String() string { return proto.CompactTextString(m) }
func (*S2KGangSimpleInfo) ProtoMessage()    {}
func (m *S2KGangSimpleInfo) GetId() uint16  { return PCK_S2KGangSimpleInfo_ID }

//
type JJCRole struct {
	//
	UserId int32 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	A []*IntAttr `protobuf:"bytes,2,rep,name=A,proto3" json:"A,omitempty"`
	//
	B []*StrAttr `protobuf:"bytes,3,rep,name=B,proto3" json:"B,omitempty"`
	//
	Kill int32 `protobuf:"varint,4,opt,name=Kill,proto3" json:"Kill,omitempty"`
	//
	Robot int32 `protobuf:"varint,5,opt,name=Robot,proto3" json:"Robot,omitempty"`
}

func (m *JJCRole) Reset()         { *m = JJCRole{} }
func (m *JJCRole) String() string { return proto.CompactTextString(m) }
func (*JJCRole) ProtoMessage()    {}

const PCK_C2SJJCList_ID = 1101 //
//
type C2SJJCList struct {
}

func (m *C2SJJCList) Reset()         { *m = C2SJJCList{} }
func (m *C2SJJCList) String() string { return proto.CompactTextString(m) }
func (*C2SJJCList) ProtoMessage()    {}
func (m *C2SJJCList) GetId() uint16  { return PCK_C2SJJCList_ID }

const PCK_S2CJJCList_ID = 1102 //
//
type S2CJJCList struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Roles []*JJCRole `protobuf:"bytes,2,rep,name=Roles,proto3" json:"Roles,omitempty"`
}

func (m *S2CJJCList) Reset()         { *m = S2CJJCList{} }
func (m *S2CJJCList) String() string { return proto.CompactTextString(m) }
func (*S2CJJCList) ProtoMessage()    {}
func (m *S2CJJCList) GetId() uint16  { return PCK_S2CJJCList_ID }

const PCK_C2SJJCFight_ID = 1103 //
//
type C2SJJCFight struct {
	//
	TargetId int32 `protobuf:"varint,1,opt,name=TargetId,proto3" json:"TargetId,omitempty"`
	//
	TargetRank int32 `protobuf:"varint,2,opt,name=TargetRank,proto3" json:"TargetRank,omitempty"`
	//
	Kill int32 `protobuf:"varint,3,opt,name=Kill,proto3" json:"Kill,omitempty"`
}

func (m *C2SJJCFight) Reset()         { *m = C2SJJCFight{} }
func (m *C2SJJCFight) String() string { return proto.CompactTextString(m) }
func (*C2SJJCFight) ProtoMessage()    {}
func (m *C2SJJCFight) GetId() uint16  { return PCK_C2SJJCFight_ID }

const PCK_S2CJJCFight_ID = 1104 //
//
type S2CJJCFight struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	IsWin int32 `protobuf:"varint,2,opt,name=IsWin,proto3" json:"IsWin,omitempty"`
	//
	HonorPrize int32 `protobuf:"varint,3,opt,name=HonorPrize,proto3" json:"HonorPrize,omitempty"`
	//
	Coin1Prize int32 `protobuf:"varint,4,opt,name=Coin1Prize,proto3" json:"Coin1Prize,omitempty"`
	//
	Coin4Prize int32 `protobuf:"varint,5,opt,name=Coin4Prize,proto3" json:"Coin4Prize,omitempty"`
	//
	HistoryRank int32 `protobuf:"varint,6,opt,name=HistoryRank,proto3" json:"HistoryRank,omitempty"`
	//
	NowRank int32 `protobuf:"varint,7,opt,name=NowRank,proto3" json:"NowRank,omitempty"`
	//
	Kill int32 `protobuf:"varint,8,opt,name=Kill,proto3" json:"Kill,omitempty"`
}

func (m *S2CJJCFight) Reset()         { *m = S2CJJCFight{} }
func (m *S2CJJCFight) String() string { return proto.CompactTextString(m) }
func (*S2CJJCFight) ProtoMessage()    {}
func (m *S2CJJCFight) GetId() uint16  { return PCK_S2CJJCFight_ID }

const PCK_C2SJJCGetBuyInfo_ID = 1107 //
//
type C2SJJCGetBuyInfo struct {
}

func (m *C2SJJCGetBuyInfo) Reset()         { *m = C2SJJCGetBuyInfo{} }
func (m *C2SJJCGetBuyInfo) String() string { return proto.CompactTextString(m) }
func (*C2SJJCGetBuyInfo) ProtoMessage()    {}
func (m *C2SJJCGetBuyInfo) GetId() uint16  { return PCK_C2SJJCGetBuyInfo_ID }

const PCK_S2CJJCGetBuyInfo_ID = 1108 //
//
type S2CJJCGetBuyInfo struct {
	//
	Coin3 int32 `protobuf:"varint,1,opt,name=Coin3,proto3" json:"Coin3,omitempty"`
	//
	Times int32 `protobuf:"varint,2,opt,name=Times,proto3" json:"Times,omitempty"`
	//
	LeftTimes int32 `protobuf:"varint,3,opt,name=LeftTimes,proto3" json:"LeftTimes,omitempty"`
	//
	NoticeTimes int32 `protobuf:"varint,4,opt,name=NoticeTimes,proto3" json:"NoticeTimes,omitempty"`
	//
	NoticeVip int32 `protobuf:"varint,5,opt,name=NoticeVip,proto3" json:"NoticeVip,omitempty"`
}

func (m *S2CJJCGetBuyInfo) Reset()         { *m = S2CJJCGetBuyInfo{} }
func (m *S2CJJCGetBuyInfo) String() string { return proto.CompactTextString(m) }
func (*S2CJJCGetBuyInfo) ProtoMessage()    {}
func (m *S2CJJCGetBuyInfo) GetId() uint16  { return PCK_S2CJJCGetBuyInfo_ID }

const PCK_C2SJJCBuyTimes_ID = 1105 //
//
type C2SJJCBuyTimes struct {
}

func (m *C2SJJCBuyTimes) Reset()         { *m = C2SJJCBuyTimes{} }
func (m *C2SJJCBuyTimes) String() string { return proto.CompactTextString(m) }
func (*C2SJJCBuyTimes) ProtoMessage()    {}
func (m *C2SJJCBuyTimes) GetId() uint16  { return PCK_C2SJJCBuyTimes_ID }

const PCK_S2CJJCBuyTimes_ID = 1106 //
//
type S2CJJCBuyTimes struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CJJCBuyTimes) Reset()         { *m = S2CJJCBuyTimes{} }
func (m *S2CJJCBuyTimes) String() string { return proto.CompactTextString(m) }
func (*S2CJJCBuyTimes) ProtoMessage()    {}
func (m *S2CJJCBuyTimes) GetId() uint16  { return PCK_S2CJJCBuyTimes_ID }

const PCK_C2SGodEquipAwake_ID = 1353 //
//
type C2SGodEquipAwake struct {
	//
	PartType int32 `protobuf:"varint,1,opt,name=PartType,proto3" json:"PartType,omitempty"`
}

func (m *C2SGodEquipAwake) Reset()         { *m = C2SGodEquipAwake{} }
func (m *C2SGodEquipAwake) String() string { return proto.CompactTextString(m) }
func (*C2SGodEquipAwake) ProtoMessage()    {}
func (m *C2SGodEquipAwake) GetId() uint16  { return PCK_C2SGodEquipAwake_ID }

const PCK_S2CGodEquipAwake_ID = 1354 //
//
type S2CGodEquipAwake struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CGodEquipAwake) Reset()         { *m = S2CGodEquipAwake{} }
func (m *S2CGodEquipAwake) String() string { return proto.CompactTextString(m) }
func (*S2CGodEquipAwake) ProtoMessage()    {}
func (m *S2CGodEquipAwake) GetId() uint16  { return PCK_S2CGodEquipAwake_ID }

const PCK_C2SIntoSoulGradeUp_ID = 1351 //
//
type C2SIntoSoulGradeUp struct {
	//
	PartType int32 `protobuf:"varint,1,opt,name=PartType,proto3" json:"PartType,omitempty"`
}

func (m *C2SIntoSoulGradeUp) Reset()         { *m = C2SIntoSoulGradeUp{} }
func (m *C2SIntoSoulGradeUp) String() string { return proto.CompactTextString(m) }
func (*C2SIntoSoulGradeUp) ProtoMessage()    {}
func (m *C2SIntoSoulGradeUp) GetId() uint16  { return PCK_C2SIntoSoulGradeUp_ID }

const PCK_S2CIntoSoulGradeUp_ID = 1352 //
//
type S2CIntoSoulGradeUp struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Directly int32 `protobuf:"varint,2,opt,name=Directly,proto3" json:"Directly,omitempty"`
	//
	Grade *Grade `protobuf:"bytes,3,opt,name=Grade,proto3" json:"Grade,omitempty"`
}

func (m *S2CIntoSoulGradeUp) Reset()         { *m = S2CIntoSoulGradeUp{} }
func (m *S2CIntoSoulGradeUp) String() string { return proto.CompactTextString(m) }
func (*S2CIntoSoulGradeUp) ProtoMessage()    {}
func (m *S2CIntoSoulGradeUp) GetId() uint16  { return PCK_S2CIntoSoulGradeUp_ID }

const PCK_C2SGodForge_ID = 1355 //
//
type C2SGodForge struct {
	//
	PartType int32 `protobuf:"varint,1,opt,name=PartType,proto3" json:"PartType,omitempty"`
	//
	Type int32 `protobuf:"varint,2,opt,name=Type,proto3" json:"Type,omitempty"`
	//
	Times int32 `protobuf:"varint,3,opt,name=Times,proto3" json:"Times,omitempty"`
}

func (m *C2SGodForge) Reset()         { *m = C2SGodForge{} }
func (m *C2SGodForge) String() string { return proto.CompactTextString(m) }
func (*C2SGodForge) ProtoMessage()    {}
func (m *C2SGodForge) GetId() uint16  { return PCK_C2SGodForge_ID }

const PCK_S2CGodForge_ID = 1356 //
//
type S2CGodForge struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Grade1 *Grade `protobuf:"bytes,2,opt,name=Grade1,proto3" json:"Grade1,omitempty"`
	//
	Grade2 *Grade `protobuf:"bytes,3,opt,name=Grade2,proto3" json:"Grade2,omitempty"`
}

func (m *S2CGodForge) Reset()         { *m = S2CGodForge{} }
func (m *S2CGodForge) String() string { return proto.CompactTextString(m) }
func (*S2CGodForge) ProtoMessage()    {}
func (m *S2CGodForge) GetId() uint16  { return PCK_S2CGodForge_ID }

const PCK_C2SGodForgeSave_ID = 1357 //
//
type C2SGodForgeSave struct {
	//
	PartType int32 `protobuf:"varint,1,opt,name=PartType,proto3" json:"PartType,omitempty"`
}

func (m *C2SGodForgeSave) Reset()         { *m = C2SGodForgeSave{} }
func (m *C2SGodForgeSave) String() string { return proto.CompactTextString(m) }
func (*C2SGodForgeSave) ProtoMessage()    {}
func (m *C2SGodForgeSave) GetId() uint16  { return PCK_C2SGodForgeSave_ID }

const PCK_S2CGodForgeSave_ID = 1358 //
//
type S2CGodForgeSave struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Grade1 *Grade `protobuf:"bytes,2,opt,name=Grade1,proto3" json:"Grade1,omitempty"`
	//
	Grade2 *Grade `protobuf:"bytes,3,opt,name=Grade2,proto3" json:"Grade2,omitempty"`
}

func (m *S2CGodForgeSave) Reset()         { *m = S2CGodForgeSave{} }
func (m *S2CGodForgeSave) String() string { return proto.CompactTextString(m) }
func (*S2CGodForgeSave) ProtoMessage()    {}
func (m *S2CGodForgeSave) GetId() uint16  { return PCK_S2CGodForgeSave_ID }

const PCK_C2SGodEquipMelt_ID = 1359 //
//
type C2SGodEquipMelt struct {
	//
	ItemId string `protobuf:"bytes,1,opt,name=ItemId,proto3" json:"ItemId,omitempty"`
}

func (m *C2SGodEquipMelt) Reset()         { *m = C2SGodEquipMelt{} }
func (m *C2SGodEquipMelt) String() string { return proto.CompactTextString(m) }
func (*C2SGodEquipMelt) ProtoMessage()    {}
func (m *C2SGodEquipMelt) GetId() uint16  { return PCK_C2SGodEquipMelt_ID }

const PCK_S2CGodEquipMelt_ID = 1360 //
//
type S2CGodEquipMelt struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CGodEquipMelt) Reset()         { *m = S2CGodEquipMelt{} }
func (m *S2CGodEquipMelt) String() string { return proto.CompactTextString(m) }
func (*S2CGodEquipMelt) ProtoMessage()    {}
func (m *S2CGodEquipMelt) GetId() uint16  { return PCK_S2CGodEquipMelt_ID }

//
type ExpeditionTeamUnit struct {
	//
	UnitType int32 `protobuf:"varint,1,opt,name=UnitType,proto3" json:"UnitType,omitempty"`
	//
	Hp int64 `protobuf:"varint,2,opt,name=Hp,proto3" json:"Hp,omitempty"`
	//
	MaxHp int64 `protobuf:"varint,3,opt,name=MaxHp,proto3" json:"MaxHp,omitempty"`
	//
	Pos int32 `protobuf:"varint,4,opt,name=Pos,proto3" json:"Pos,omitempty"`
	//
	Id int64 `protobuf:"varint,5,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	State int64 `protobuf:"varint,6,opt,name=State,proto3" json:"State,omitempty"`
}

func (m *ExpeditionTeamUnit) Reset()         { *m = ExpeditionTeamUnit{} }
func (m *ExpeditionTeamUnit) String() string { return proto.CompactTextString(m) }
func (*ExpeditionTeamUnit) ProtoMessage()    {}

//
type ExpeditionData struct {
	//
	Team []*ExpeditionTeamUnit `protobuf:"bytes,1,rep,name=Team,proto3" json:"Team,omitempty"`
	//
	CurrId int32 `protobuf:"varint,2,opt,name=CurrId,proto3" json:"CurrId,omitempty"`
	//
	CurrTeam []*ExpeditionTeamUnit `protobuf:"bytes,3,rep,name=CurrTeam,proto3" json:"CurrTeam,omitempty"`
	//
	CurrFightUserId int64 `protobuf:"varint,4,opt,name=CurrFightUserId,proto3" json:"CurrFightUserId,omitempty"`
	//
	CheckPointRank []int32 `protobuf:"varint,5,rep,name=CheckPointRank,proto3" json:"CheckPointRank,omitempty"`
	//
	Times int32 `protobuf:"varint,6,opt,name=Times,proto3" json:"Times,omitempty"`
	//
	IsOpen bool `protobuf:"varint,7,opt,name=IsOpen,proto3" json:"IsOpen,omitempty"`
	//
	BoxState int32 `protobuf:"varint,8,opt,name=BoxState,proto3" json:"BoxState,omitempty"`
	//
	DayResetTimes int32 `protobuf:"varint,9,opt,name=DayResetTimes,proto3" json:"DayResetTimes,omitempty"`
	//
	Buff []*ExpeditionBuff `protobuf:"bytes,10,rep,name=Buff,proto3" json:"Buff,omitempty"`
	//
	GoodsId int32 `protobuf:"varint,11,opt,name=GoodsId,proto3" json:"GoodsId,omitempty"`
	//
	DayMaxStage int32 `protobuf:"varint,12,opt,name=DayMaxStage,proto3" json:"DayMaxStage,omitempty"`
}

func (m *ExpeditionData) Reset()         { *m = ExpeditionData{} }
func (m *ExpeditionData) String() string { return proto.CompactTextString(m) }
func (*ExpeditionData) ProtoMessage()    {}

//
type TeamInfo struct {
	//
	TeamId int64 `protobuf:"varint,1,opt,name=TeamId,proto3" json:"TeamId,omitempty"`
	//
	Nick string `protobuf:"bytes,2,opt,name=Nick,proto3" json:"Nick,omitempty"`
	//
	Num int32 `protobuf:"varint,3,opt,name=Num,proto3" json:"Num,omitempty"`
}

func (m *TeamInfo) Reset()         { *m = TeamInfo{} }
func (m *TeamInfo) String() string { return proto.CompactTextString(m) }
func (*TeamInfo) ProtoMessage()    {}

//
type MemberInfo struct {
	//
	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	Nick string `protobuf:"bytes,2,opt,name=Nick,proto3" json:"Nick,omitempty"`
	//
	Level int64 `protobuf:"varint,3,opt,name=Level,proto3" json:"Level,omitempty"`
	//
	RoleId int64 `protobuf:"varint,4,opt,name=RoleId,proto3" json:"RoleId,omitempty"`
	//
	Value int64 `protobuf:"varint,5,opt,name=Value,proto3" json:"Value,omitempty"`
	//
	Leader int64 `protobuf:"varint,7,opt,name=Leader,proto3" json:"Leader,omitempty"`
}

func (m *MemberInfo) Reset()         { *m = MemberInfo{} }
func (m *MemberInfo) String() string { return proto.CompactTextString(m) }
func (*MemberInfo) ProtoMessage()    {}

const PCK_C2SGetTeamList_ID = 8101 //
//
type C2SGetTeamList struct {
	//
	InstanceType int32 `protobuf:"varint,1,opt,name=InstanceType,proto3" json:"InstanceType,omitempty"`
	//
	InstanceId int64 `protobuf:"varint,2,opt,name=InstanceId,proto3" json:"InstanceId,omitempty"`
}

func (m *C2SGetTeamList) Reset()         { *m = C2SGetTeamList{} }
func (m *C2SGetTeamList) String() string { return proto.CompactTextString(m) }
func (*C2SGetTeamList) ProtoMessage()    {}
func (m *C2SGetTeamList) GetId() uint16  { return PCK_C2SGetTeamList_ID }

const PCK_S2CGetTeamList_ID = 8102 //
//
type S2CGetTeamList struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	InstanceType int32 `protobuf:"varint,2,opt,name=InstanceType,proto3" json:"InstanceType,omitempty"`
	//
	InstanceId int64 `protobuf:"varint,3,opt,name=InstanceId,proto3" json:"InstanceId,omitempty"`
	//
	Teams []*TeamInfo `protobuf:"bytes,4,rep,name=Teams,proto3" json:"Teams,omitempty"`
}

func (m *S2CGetTeamList) Reset()         { *m = S2CGetTeamList{} }
func (m *S2CGetTeamList) String() string { return proto.CompactTextString(m) }
func (*S2CGetTeamList) ProtoMessage()    {}
func (m *S2CGetTeamList) GetId() uint16  { return PCK_S2CGetTeamList_ID }

const PCK_C2SGetMemberList_ID = 8103 //
//
type C2SGetMemberList struct {
	//
	InstanceType int32 `protobuf:"varint,1,opt,name=InstanceType,proto3" json:"InstanceType,omitempty"`
	//
	TeamId int64 `protobuf:"varint,2,opt,name=TeamId,proto3" json:"TeamId,omitempty"`
}

func (m *C2SGetMemberList) Reset()         { *m = C2SGetMemberList{} }
func (m *C2SGetMemberList) String() string { return proto.CompactTextString(m) }
func (*C2SGetMemberList) ProtoMessage()    {}
func (m *C2SGetMemberList) GetId() uint16  { return PCK_C2SGetMemberList_ID }

const PCK_S2CGetMemberList_ID = 8104 //
//
type S2CGetMemberList struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	TeamId int64 `protobuf:"varint,2,opt,name=TeamId,proto3" json:"TeamId,omitempty"`
	//
	Member []*MemberInfo `protobuf:"bytes,3,rep,name=Member,proto3" json:"Member,omitempty"`
}

func (m *S2CGetMemberList) Reset()         { *m = S2CGetMemberList{} }
func (m *S2CGetMemberList) String() string { return proto.CompactTextString(m) }
func (*S2CGetMemberList) ProtoMessage()    {}
func (m *S2CGetMemberList) GetId() uint16  { return PCK_S2CGetMemberList_ID }

const PCK_S2KCrossChat_ID = 8105 //
//
type S2KCrossChat struct {
	//
	Msg *ChatMsg `protobuf:"bytes,1,opt,name=Msg,proto3" json:"Msg,omitempty"`
}

func (m *S2KCrossChat) Reset()         { *m = S2KCrossChat{} }
func (m *S2KCrossChat) String() string { return proto.CompactTextString(m) }
func (*S2KCrossChat) ProtoMessage()    {}
func (m *S2KCrossChat) GetId() uint16  { return PCK_S2KCrossChat_ID }

const PCK_K2SCrossChat_ID = 8106 //
//
type K2SCrossChat struct {
	//
	Msg *ChatMsg `protobuf:"bytes,1,opt,name=Msg,proto3" json:"Msg,omitempty"`
}

func (m *K2SCrossChat) Reset()         { *m = K2SCrossChat{} }
func (m *K2SCrossChat) String() string { return proto.CompactTextString(m) }
func (*K2SCrossChat) ProtoMessage()    {}
func (m *K2SCrossChat) GetId() uint16  { return PCK_K2SCrossChat_ID }

const PCK_K2SAddTmpTitle_ID = 8107 //
//
type K2SAddTmpTitle struct {
	//
	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	SkinId int32 `protobuf:"varint,2,opt,name=SkinId,proto3" json:"SkinId,omitempty"`
	//
	Opp int32 `protobuf:"varint,3,opt,name=Opp,proto3" json:"Opp,omitempty"`
}

func (m *K2SAddTmpTitle) Reset()         { *m = K2SAddTmpTitle{} }
func (m *K2SAddTmpTitle) String() string { return proto.CompactTextString(m) }
func (*K2SAddTmpTitle) ProtoMessage()    {}
func (m *K2SAddTmpTitle) GetId() uint16  { return PCK_K2SAddTmpTitle_ID }

const PCK_C2SJoinInstance_ID = 8108 //
//
type C2SJoinInstance struct {
	//
	InstanceType int32 `protobuf:"varint,1,opt,name=InstanceType,proto3" json:"InstanceType,omitempty"`
	//
	InstanceId int64 `protobuf:"varint,2,opt,name=InstanceId,proto3" json:"InstanceId,omitempty"`
	//
	TeamId int64 `protobuf:"varint,3,opt,name=TeamId,proto3" json:"TeamId,omitempty"`
	//
	Type int32 `protobuf:"varint,5,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (m *C2SJoinInstance) Reset()         { *m = C2SJoinInstance{} }
func (m *C2SJoinInstance) String() string { return proto.CompactTextString(m) }
func (*C2SJoinInstance) ProtoMessage()    {}
func (m *C2SJoinInstance) GetId() uint16  { return PCK_C2SJoinInstance_ID }

const PCK_S2CJoinInstance_ID = 8109 //
//
type S2CJoinInstance struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	TeamId int64 `protobuf:"varint,2,opt,name=TeamId,proto3" json:"TeamId,omitempty"`
	//
	LeaderId int64 `protobuf:"varint,3,opt,name=LeaderId,proto3" json:"LeaderId,omitempty"`
	//
	InstanceType int32 `protobuf:"varint,4,opt,name=InstanceType,proto3" json:"InstanceType,omitempty"`
}

func (m *S2CJoinInstance) Reset()         { *m = S2CJoinInstance{} }
func (m *S2CJoinInstance) String() string { return proto.CompactTextString(m) }
func (*S2CJoinInstance) ProtoMessage()    {}
func (m *S2CJoinInstance) GetId() uint16  { return PCK_S2CJoinInstance_ID }

const PCK_C2SLeaveInstance_ID = 8110 //
//
type C2SLeaveInstance struct {
	//
	InstanceType int32 `protobuf:"varint,1,opt,name=InstanceType,proto3" json:"InstanceType,omitempty"`
	//
	TeamId int64 `protobuf:"varint,2,opt,name=TeamId,proto3" json:"TeamId,omitempty"`
}

func (m *C2SLeaveInstance) Reset()         { *m = C2SLeaveInstance{} }
func (m *C2SLeaveInstance) String() string { return proto.CompactTextString(m) }
func (*C2SLeaveInstance) ProtoMessage()    {}
func (m *C2SLeaveInstance) GetId() uint16  { return PCK_C2SLeaveInstance_ID }

const PCK_S2CLeaveInstance_ID = 8111 //
//
type S2CLeaveInstance struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CLeaveInstance) Reset()         { *m = S2CLeaveInstance{} }
func (m *S2CLeaveInstance) String() string { return proto.CompactTextString(m) }
func (*S2CLeaveInstance) ProtoMessage()    {}
func (m *S2CLeaveInstance) GetId() uint16  { return PCK_S2CLeaveInstance_ID }

const PCK_C2SLeaveInstanceCopy_ID = 8118 //
//
type C2SLeaveInstanceCopy struct {
}

func (m *C2SLeaveInstanceCopy) Reset()         { *m = C2SLeaveInstanceCopy{} }
func (m *C2SLeaveInstanceCopy) String() string { return proto.CompactTextString(m) }
func (*C2SLeaveInstanceCopy) ProtoMessage()    {}
func (m *C2SLeaveInstanceCopy) GetId() uint16  { return PCK_C2SLeaveInstanceCopy_ID }

const PCK_S2CLeaveInstanceCopy_ID = 8119 //
//
type S2CLeaveInstanceCopy struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CLeaveInstanceCopy) Reset()         { *m = S2CLeaveInstanceCopy{} }
func (m *S2CLeaveInstanceCopy) String() string { return proto.CompactTextString(m) }
func (*S2CLeaveInstanceCopy) ProtoMessage()    {}
func (m *S2CLeaveInstanceCopy) GetId() uint16  { return PCK_S2CLeaveInstanceCopy_ID }

const PCK_C2SKillInstance_ID = 8112 //
//
type C2SKillInstance struct {
	//
	InstanceType int32 `protobuf:"varint,1,opt,name=InstanceType,proto3" json:"InstanceType,omitempty"`
	//
	TeamId int64 `protobuf:"varint,2,opt,name=TeamId,proto3" json:"TeamId,omitempty"`
	//
	KillId int64 `protobuf:"varint,3,opt,name=KillId,proto3" json:"KillId,omitempty"`
}

func (m *C2SKillInstance) Reset()         { *m = C2SKillInstance{} }
func (m *C2SKillInstance) String() string { return proto.CompactTextString(m) }
func (*C2SKillInstance) ProtoMessage()    {}
func (m *C2SKillInstance) GetId() uint16  { return PCK_C2SKillInstance_ID }

const PCK_S2CKillInstance_ID = 8113 //
//
type S2CKillInstance struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CKillInstance) Reset()         { *m = S2CKillInstance{} }
func (m *S2CKillInstance) String() string { return proto.CompactTextString(m) }
func (*S2CKillInstance) ProtoMessage()    {}
func (m *S2CKillInstance) GetId() uint16  { return PCK_S2CKillInstance_ID }

const PCK_C2SStartInstance_ID = 8114 //
//
type C2SStartInstance struct {
	//
	InstanceType int32 `protobuf:"varint,1,opt,name=InstanceType,proto3" json:"InstanceType,omitempty"`
	//
	TeamId int64 `protobuf:"varint,2,opt,name=TeamId,proto3" json:"TeamId,omitempty"`
}

func (m *C2SStartInstance) Reset()         { *m = C2SStartInstance{} }
func (m *C2SStartInstance) String() string { return proto.CompactTextString(m) }
func (*C2SStartInstance) ProtoMessage()    {}
func (m *C2SStartInstance) GetId() uint16  { return PCK_C2SStartInstance_ID }

const PCK_S2CStartInstance_ID = 8115 //
//
type S2CStartInstance struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CStartInstance) Reset()         { *m = S2CStartInstance{} }
func (m *S2CStartInstance) String() string { return proto.CompactTextString(m) }
func (*S2CStartInstance) ProtoMessage()    {}
func (m *S2CStartInstance) GetId() uint16  { return PCK_S2CStartInstance_ID }

//
type Precious struct {
	//
	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	IId int32 `protobuf:"varint,2,opt,name=IId,proto3" json:"IId,omitempty"`
	//
	Quality int32 `protobuf:"varint,3,opt,name=Quality,proto3" json:"Quality,omitempty"`
	//
	Pos int32 `protobuf:"varint,4,opt,name=Pos,proto3" json:"Pos,omitempty"`
	//
	Skills []*Skill `protobuf:"bytes,5,rep,name=Skills,proto3" json:"Skills,omitempty"`
	//
	Lock int32 `protobuf:"varint,6,opt,name=Lock,proto3" json:"Lock,omitempty"`
	//
	Creator string `protobuf:"bytes,7,opt,name=Creator,proto3" json:"Creator,omitempty"`
	//
	EatLevel int32 `protobuf:"varint,8,opt,name=EatLevel,proto3" json:"EatLevel,omitempty"`
	//
	EatExp int32 `protobuf:"varint,9,opt,name=EatExp,proto3" json:"EatExp,omitempty"`
	//
	SoulId int32 `protobuf:"varint,10,opt,name=SoulId,proto3" json:"SoulId,omitempty"`
	//
	SoulLevel int32 `protobuf:"varint,11,opt,name=SoulLevel,proto3" json:"SoulLevel,omitempty"`
	//
	SoulExp int32 `protobuf:"varint,12,opt,name=SoulExp,proto3" json:"SoulExp,omitempty"`
}

func (m *Precious) Reset()         { *m = Precious{} }
func (m *Precious) String() string { return proto.CompactTextString(m) }
func (*Precious) ProtoMessage()    {}

//
type PreciousPos struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	ForgeTime int32 `protobuf:"varint,2,opt,name=ForgeTime,proto3" json:"ForgeTime,omitempty"`
	//
	Attr []*IntAttr `protobuf:"bytes,3,rep,name=Attr,proto3" json:"Attr,omitempty"`
}

func (m *PreciousPos) Reset()         { *m = PreciousPos{} }
func (m *PreciousPos) String() string { return proto.CompactTextString(m) }
func (*PreciousPos) ProtoMessage()    {}

const PCK_C2SGetPreciousPos_ID = 1409 //
//
type C2SGetPreciousPos struct {
}

func (m *C2SGetPreciousPos) Reset()         { *m = C2SGetPreciousPos{} }
func (m *C2SGetPreciousPos) String() string { return proto.CompactTextString(m) }
func (*C2SGetPreciousPos) ProtoMessage()    {}
func (m *C2SGetPreciousPos) GetId() uint16  { return PCK_C2SGetPreciousPos_ID }

const PCK_S2CGetPreciousPos_ID = 1410 //
//
type S2CGetPreciousPos struct {
	//
	Pos []*PreciousPos `protobuf:"bytes,1,rep,name=Pos,proto3" json:"Pos,omitempty"`
}

func (m *S2CGetPreciousPos) Reset()         { *m = S2CGetPreciousPos{} }
func (m *S2CGetPreciousPos) String() string { return proto.CompactTextString(m) }
func (*S2CGetPreciousPos) ProtoMessage()    {}
func (m *S2CGetPreciousPos) GetId() uint16  { return PCK_S2CGetPreciousPos_ID }

const PCK_C2SCreatePrecious_ID = 1401 //
//
type C2SCreatePrecious struct {
	//
	Type int32 `protobuf:"varint,1,opt,name=Type,proto3" json:"Type,omitempty"`
	//
	Auto int32 `protobuf:"varint,2,opt,name=Auto,proto3" json:"Auto,omitempty"`
}

func (m *C2SCreatePrecious) Reset()         { *m = C2SCreatePrecious{} }
func (m *C2SCreatePrecious) String() string { return proto.CompactTextString(m) }
func (*C2SCreatePrecious) ProtoMessage()    {}
func (m *C2SCreatePrecious) GetId() uint16  { return PCK_C2SCreatePrecious_ID }

const PCK_S2CCreatePrecious_ID = 1402 //
//
type S2CCreatePrecious struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	NewPrecious *Precious `protobuf:"bytes,2,opt,name=NewPrecious,proto3" json:"NewPrecious,omitempty"`
}

func (m *S2CCreatePrecious) Reset()         { *m = S2CCreatePrecious{} }
func (m *S2CCreatePrecious) String() string { return proto.CompactTextString(m) }
func (*S2CCreatePrecious) ProtoMessage()    {}
func (m *S2CCreatePrecious) GetId() uint16  { return PCK_S2CCreatePrecious_ID }

const PCK_C2SCreatePreciousFast_ID = 1421 //
//
type C2SCreatePreciousFast struct {
	//
	Type int32 `protobuf:"varint,1,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (m *C2SCreatePreciousFast) Reset()         { *m = C2SCreatePreciousFast{} }
func (m *C2SCreatePreciousFast) String() string { return proto.CompactTextString(m) }
func (*C2SCreatePreciousFast) ProtoMessage()    {}
func (m *C2SCreatePreciousFast) GetId() uint16  { return PCK_C2SCreatePreciousFast_ID }

const PCK_S2CCreatePreciousFast_ID = 1422 //
//
type S2CCreatePreciousFast struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	NewPrecious []*Precious `protobuf:"bytes,2,rep,name=NewPrecious,proto3" json:"NewPrecious,omitempty"`
	//
	Precious []*Precious `protobuf:"bytes,3,rep,name=Precious,proto3" json:"Precious,omitempty"`
}

func (m *S2CCreatePreciousFast) Reset()         { *m = S2CCreatePreciousFast{} }
func (m *S2CCreatePreciousFast) String() string { return proto.CompactTextString(m) }
func (*S2CCreatePreciousFast) ProtoMessage()    {}
func (m *S2CCreatePreciousFast) GetId() uint16  { return PCK_S2CCreatePreciousFast_ID }

const PCK_C2SMeltPrecious_ID = 1403 //
//
type C2SMeltPrecious struct {
	//
	Precious []string `protobuf:"bytes,1,rep,name=Precious,proto3" json:"Precious,omitempty"`
}

func (m *C2SMeltPrecious) Reset()         { *m = C2SMeltPrecious{} }
func (m *C2SMeltPrecious) String() string { return proto.CompactTextString(m) }
func (*C2SMeltPrecious) ProtoMessage()    {}
func (m *C2SMeltPrecious) GetId() uint16  { return PCK_C2SMeltPrecious_ID }

const PCK_S2CMeltPrecious_ID = 1404 //
//
type S2CMeltPrecious struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Precious []string `protobuf:"bytes,2,rep,name=Precious,proto3" json:"Precious,omitempty"`
}

func (m *S2CMeltPrecious) Reset()         { *m = S2CMeltPrecious{} }
func (m *S2CMeltPrecious) String() string { return proto.CompactTextString(m) }
func (*S2CMeltPrecious) ProtoMessage()    {}
func (m *S2CMeltPrecious) GetId() uint16  { return PCK_S2CMeltPrecious_ID }

const PCK_C2SWearPrecious_ID = 1405 //
//
type C2SWearPrecious struct {
	//
	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Pos int32 `protobuf:"varint,2,opt,name=Pos,proto3" json:"Pos,omitempty"`
}

func (m *C2SWearPrecious) Reset()         { *m = C2SWearPrecious{} }
func (m *C2SWearPrecious) String() string { return proto.CompactTextString(m) }
func (*C2SWearPrecious) ProtoMessage()    {}
func (m *C2SWearPrecious) GetId() uint16  { return PCK_C2SWearPrecious_ID }

const PCK_S2CWearPrecious_ID = 1406 //
//
type S2CWearPrecious struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id string `protobuf:"bytes,2,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Pos int32 `protobuf:"varint,3,opt,name=Pos,proto3" json:"Pos,omitempty"`
}

func (m *S2CWearPrecious) Reset()         { *m = S2CWearPrecious{} }
func (m *S2CWearPrecious) String() string { return proto.CompactTextString(m) }
func (*S2CWearPrecious) ProtoMessage()    {}
func (m *S2CWearPrecious) GetId() uint16  { return PCK_S2CWearPrecious_ID }

const PCK_C2SLockPrecious_ID = 1407 //
//
type C2SLockPrecious struct {
	//
	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2SLockPrecious) Reset()         { *m = C2SLockPrecious{} }
func (m *C2SLockPrecious) String() string { return proto.CompactTextString(m) }
func (*C2SLockPrecious) ProtoMessage()    {}
func (m *C2SLockPrecious) GetId() uint16  { return PCK_C2SLockPrecious_ID }

const PCK_S2CLockPrecious_ID = 1408 //
//
type S2CLockPrecious struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id string `protobuf:"bytes,2,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Lock int32 `protobuf:"varint,3,opt,name=Lock,proto3" json:"Lock,omitempty"`
}

func (m *S2CLockPrecious) Reset()         { *m = S2CLockPrecious{} }
func (m *S2CLockPrecious) String() string { return proto.CompactTextString(m) }
func (*S2CLockPrecious) ProtoMessage()    {}
func (m *S2CLockPrecious) GetId() uint16  { return PCK_S2CLockPrecious_ID }

const PCK_C2SPreciousForge_ID = 1411 //
//
type C2SPreciousForge struct {
	//
	Pos int32 `protobuf:"varint,1,opt,name=Pos,proto3" json:"Pos,omitempty"`
	//
	T int32 `protobuf:"varint,2,opt,name=T,proto3" json:"T,omitempty"`
	//
	LockAttr []int32 `protobuf:"varint,3,rep,name=LockAttr,proto3" json:"LockAttr,omitempty"`
}

func (m *C2SPreciousForge) Reset()         { *m = C2SPreciousForge{} }
func (m *C2SPreciousForge) String() string { return proto.CompactTextString(m) }
func (*C2SPreciousForge) ProtoMessage()    {}
func (m *C2SPreciousForge) GetId() uint16  { return PCK_C2SPreciousForge_ID }

const PCK_S2CPreciousForge_ID = 1412 //
//
type S2CPreciousForge struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Pos int32 `protobuf:"varint,2,opt,name=Pos,proto3" json:"Pos,omitempty"`
	//
	T int32 `protobuf:"varint,3,opt,name=T,proto3" json:"T,omitempty"`
	//
	LockAttr []int32 `protobuf:"varint,4,rep,name=LockAttr,proto3" json:"LockAttr,omitempty"`
	//
	Attr []*IntAttr `protobuf:"bytes,5,rep,name=Attr,proto3" json:"Attr,omitempty"`
	//
	Times int32 `protobuf:"varint,6,opt,name=Times,proto3" json:"Times,omitempty"`
}

func (m *S2CPreciousForge) Reset()         { *m = S2CPreciousForge{} }
func (m *S2CPreciousForge) String() string { return proto.CompactTextString(m) }
func (*S2CPreciousForge) ProtoMessage()    {}
func (m *S2CPreciousForge) GetId() uint16  { return PCK_S2CPreciousForge_ID }

const PCK_C2SPreciousEat_ID = 1413 //
//
type C2SPreciousEat struct {
	//
	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	AutoBuy int32 `protobuf:"varint,2,opt,name=AutoBuy,proto3" json:"AutoBuy,omitempty"`
}

func (m *C2SPreciousEat) Reset()         { *m = C2SPreciousEat{} }
func (m *C2SPreciousEat) String() string { return proto.CompactTextString(m) }
func (*C2SPreciousEat) ProtoMessage()    {}
func (m *C2SPreciousEat) GetId() uint16  { return PCK_C2SPreciousEat_ID }

const PCK_S2CPreciousEat_ID = 1414 //
//
type S2CPreciousEat struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id string `protobuf:"bytes,2,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	EatLevel int32 `protobuf:"varint,3,opt,name=EatLevel,proto3" json:"EatLevel,omitempty"`
	//
	EatExp int32 `protobuf:"varint,4,opt,name=EatExp,proto3" json:"EatExp,omitempty"`
}

func (m *S2CPreciousEat) Reset()         { *m = S2CPreciousEat{} }
func (m *S2CPreciousEat) String() string { return proto.CompactTextString(m) }
func (*S2CPreciousEat) ProtoMessage()    {}
func (m *S2CPreciousEat) GetId() uint16  { return PCK_S2CPreciousEat_ID }

const PCK_C2SPreciousSoul_ID = 1415 //
//
type C2SPreciousSoul struct {
	//
	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	SoulId int32 `protobuf:"varint,2,opt,name=SoulId,proto3" json:"SoulId,omitempty"`
}

func (m *C2SPreciousSoul) Reset()         { *m = C2SPreciousSoul{} }
func (m *C2SPreciousSoul) String() string { return proto.CompactTextString(m) }
func (*C2SPreciousSoul) ProtoMessage()    {}
func (m *C2SPreciousSoul) GetId() uint16  { return PCK_C2SPreciousSoul_ID }

const PCK_S2CPreciousSoul_ID = 1416 //
//
type S2CPreciousSoul struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id string `protobuf:"bytes,2,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	SoulId int32 `protobuf:"varint,3,opt,name=SoulId,proto3" json:"SoulId,omitempty"`
}

func (m *S2CPreciousSoul) Reset()         { *m = S2CPreciousSoul{} }
func (m *S2CPreciousSoul) String() string { return proto.CompactTextString(m) }
func (*S2CPreciousSoul) ProtoMessage()    {}
func (m *S2CPreciousSoul) GetId() uint16  { return PCK_S2CPreciousSoul_ID }

const PCK_C2SPreciousSoulUp_ID = 1417 //
//
type C2SPreciousSoulUp struct {
	//
	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	SoulId []*ItemInfo `protobuf:"bytes,2,rep,name=SoulId,proto3" json:"SoulId,omitempty"`
}

func (m *C2SPreciousSoulUp) Reset()         { *m = C2SPreciousSoulUp{} }
func (m *C2SPreciousSoulUp) String() string { return proto.CompactTextString(m) }
func (*C2SPreciousSoulUp) ProtoMessage()    {}
func (m *C2SPreciousSoulUp) GetId() uint16  { return PCK_C2SPreciousSoulUp_ID }

const PCK_S2CPreciousSoulUp_ID = 1418 //
//
type S2CPreciousSoulUp struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id string `protobuf:"bytes,2,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	SoulLevel int32 `protobuf:"varint,3,opt,name=SoulLevel,proto3" json:"SoulLevel,omitempty"`
	//
	SoulExp int32 `protobuf:"varint,4,opt,name=SoulExp,proto3" json:"SoulExp,omitempty"`
}

func (m *S2CPreciousSoulUp) Reset()         { *m = S2CPreciousSoulUp{} }
func (m *S2CPreciousSoulUp) String() string { return proto.CompactTextString(m) }
func (*S2CPreciousSoulUp) ProtoMessage()    {}
func (m *S2CPreciousSoulUp) GetId() uint16  { return PCK_S2CPreciousSoulUp_ID }

const PCK_C2SPreciousGive_ID = 1419 //
//
type C2SPreciousGive struct {
	//
	Id string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Id2 string `protobuf:"bytes,2,opt,name=Id2,proto3" json:"Id2,omitempty"`
	//
	T int32 `protobuf:"varint,3,opt,name=T,proto3" json:"T,omitempty"`
}

func (m *C2SPreciousGive) Reset()         { *m = C2SPreciousGive{} }
func (m *C2SPreciousGive) String() string { return proto.CompactTextString(m) }
func (*C2SPreciousGive) ProtoMessage()    {}
func (m *C2SPreciousGive) GetId() uint16  { return PCK_C2SPreciousGive_ID }

const PCK_S2CPreciousGive_ID = 1420 //
//
type S2CPreciousGive struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id string `protobuf:"bytes,2,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Id2 string `protobuf:"bytes,3,opt,name=Id2,proto3" json:"Id2,omitempty"`
	//
	T int32 `protobuf:"varint,4,opt,name=T,proto3" json:"T,omitempty"`
}

func (m *S2CPreciousGive) Reset()         { *m = S2CPreciousGive{} }
func (m *S2CPreciousGive) String() string { return proto.CompactTextString(m) }
func (*S2CPreciousGive) ProtoMessage()    {}
func (m *S2CPreciousGive) GetId() uint16  { return PCK_S2CPreciousGive_ID }

const PCK_C2SPreciousRobot_ID = 1423 //
//
type C2SPreciousRobot struct {
}

func (m *C2SPreciousRobot) Reset()         { *m = C2SPreciousRobot{} }
func (m *C2SPreciousRobot) String() string { return proto.CompactTextString(m) }
func (*C2SPreciousRobot) ProtoMessage()    {}
func (m *C2SPreciousRobot) GetId() uint16  { return PCK_C2SPreciousRobot_ID }

const PCK_S2CPreciousRobot_ID = 1424 //
//
type S2CPreciousRobot struct {
	//
	Precious []*Precious `protobuf:"bytes,1,rep,name=Precious,proto3" json:"Precious,omitempty"`
}

func (m *S2CPreciousRobot) Reset()         { *m = S2CPreciousRobot{} }
func (m *S2CPreciousRobot) String() string { return proto.CompactTextString(m) }
func (*S2CPreciousRobot) ProtoMessage()    {}
func (m *S2CPreciousRobot) GetId() uint16  { return PCK_S2CPreciousRobot_ID }

//
type GangWarIncGetDragonIds struct {
}

func (m *GangWarIncGetDragonIds) Reset()         { *m = GangWarIncGetDragonIds{} }
func (m *GangWarIncGetDragonIds) String() string { return proto.CompactTextString(m) }
func (*GangWarIncGetDragonIds) ProtoMessage()    {}

//
type S2CActPlayerDeath struct {
	//
	UId int64 `protobuf:"varint,1,opt,name=UId,proto3" json:"UId,omitempty"`
}

func (m *S2CActPlayerDeath) Reset()         { *m = S2CActPlayerDeath{} }
func (m *S2CActPlayerDeath) String() string { return proto.CompactTextString(m) }
func (*S2CActPlayerDeath) ProtoMessage()    {}

//
type ReqActGangWarCDTime struct {
	//
	UId int64 `protobuf:"varint,1,opt,name=UId,proto3" json:"UId,omitempty"`
	//
	MId int64 `protobuf:"varint,2,opt,name=MId,proto3" json:"MId,omitempty"`
}

func (m *ReqActGangWarCDTime) Reset()         { *m = ReqActGangWarCDTime{} }
func (m *ReqActGangWarCDTime) String() string { return proto.CompactTextString(m) }
func (*ReqActGangWarCDTime) ProtoMessage()    {}

const PCK_S2CActGangWarSettlement_ID = 1584 //
//
type S2CActGangWarSettlement struct {
	//
	Gid string `protobuf:"bytes,1,opt,name=Gid,proto3" json:"Gid,omitempty"`
	//
	GN string `protobuf:"bytes,2,opt,name=GN,proto3" json:"GN,omitempty"`
	//
	GR int64 `protobuf:"varint,3,opt,name=GR,proto3" json:"GR,omitempty"`
	//
	MR int64 `protobuf:"varint,4,opt,name=MR,proto3" json:"MR,omitempty"`
	//
	SId int64 `protobuf:"varint,5,opt,name=SId,proto3" json:"SId,omitempty"`
}

func (m *S2CActGangWarSettlement) Reset()         { *m = S2CActGangWarSettlement{} }
func (m *S2CActGangWarSettlement) String() string { return proto.CompactTextString(m) }
func (*S2CActGangWarSettlement) ProtoMessage()    {}
func (m *S2CActGangWarSettlement) GetId() uint16  { return PCK_S2CActGangWarSettlement_ID }

const PCK_S2CGangWarBeforeCD_ID = 1583 //
//
type S2CGangWarBeforeCD struct {
	//
	CDs []*CDInfo `protobuf:"bytes,1,rep,name=CDs,proto3" json:"CDs,omitempty"`
}

func (m *S2CGangWarBeforeCD) Reset()         { *m = S2CGangWarBeforeCD{} }
func (m *S2CGangWarBeforeCD) String() string { return proto.CompactTextString(m) }
func (*S2CGangWarBeforeCD) ProtoMessage()    {}
func (m *S2CGangWarBeforeCD) GetId() uint16  { return PCK_S2CGangWarBeforeCD_ID }

//
type CDInfo struct {
	//
	FId int64 `protobuf:"varint,1,opt,name=FId,proto3" json:"FId,omitempty"`
	//
	SId int64 `protobuf:"varint,2,opt,name=SId,proto3" json:"SId,omitempty"`
	//
	T int64 `protobuf:"varint,3,opt,name=T,proto3" json:"T,omitempty"`
}

func (m *CDInfo) Reset()         { *m = CDInfo{} }
func (m *CDInfo) String() string { return proto.CompactTextString(m) }
func (*CDInfo) ProtoMessage()    {}

//
type GangWarFightInfo struct {
	//
	Win bool `protobuf:"varint,1,opt,name=Win,proto3" json:"Win,omitempty"`
	//
	Pvp bool `protobuf:"varint,2,opt,name=Pvp,proto3" json:"Pvp,omitempty"`
	//
	AHV int64 `protobuf:"varint,3,opt,name=AHV,proto3" json:"AHV,omitempty"`
	//
	PS []*GangWarPlayerInfo `protobuf:"bytes,4,rep,name=PS,proto3" json:"PS,omitempty"`
	//
	IncType int32 `protobuf:"varint,5,opt,name=IncType,proto3" json:"IncType,omitempty"`
	//
	MId int64 `protobuf:"varint,6,opt,name=MId,proto3" json:"MId,omitempty"`
	//
	EPS []*GangWarPlayerInfo `protobuf:"bytes,7,rep,name=EPS,proto3" json:"EPS,omitempty"`
}

func (m *GangWarFightInfo) Reset()         { *m = GangWarFightInfo{} }
func (m *GangWarFightInfo) String() string { return proto.CompactTextString(m) }
func (*GangWarFightInfo) ProtoMessage()    {}

//
type GangWarPlayerInfo struct {
	//
	UId int64 `protobuf:"varint,1,opt,name=UId,proto3" json:"UId,omitempty"`
	//
	HV int64 `protobuf:"varint,2,opt,name=HV,proto3" json:"HV,omitempty"`
}

func (m *GangWarPlayerInfo) Reset()         { *m = GangWarPlayerInfo{} }
func (m *GangWarPlayerInfo) String() string { return proto.CompactTextString(m) }
func (*GangWarPlayerInfo) ProtoMessage()    {}

const PCK_C2SActSountGateFight_ID = 1574 //
//
type C2SActSountGateFight struct {
}

func (m *C2SActSountGateFight) Reset()         { *m = C2SActSountGateFight{} }
func (m *C2SActSountGateFight) String() string { return proto.CompactTextString(m) }
func (*C2SActSountGateFight) ProtoMessage()    {}
func (m *C2SActSountGateFight) GetId() uint16  { return PCK_C2SActSountGateFight_ID }

const PCK_S2CActSountGateFight_ID = 1575 //
//
type S2CActSountGateFight struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CActSountGateFight) Reset()         { *m = S2CActSountGateFight{} }
func (m *S2CActSountGateFight) String() string { return proto.CompactTextString(m) }
func (*S2CActSountGateFight) ProtoMessage()    {}
func (m *S2CActSountGateFight) GetId() uint16  { return PCK_S2CActSountGateFight_ID }

const PCK_S2CSouthGateStatus_ID = 1573 //
//
type S2CSouthGateStatus struct {
	//
	HP int64 `protobuf:"varint,1,opt,name=HP,proto3" json:"HP,omitempty"`
	//
	MHP int64 `protobuf:"varint,2,opt,name=MHP,proto3" json:"MHP,omitempty"`
}

func (m *S2CSouthGateStatus) Reset()         { *m = S2CSouthGateStatus{} }
func (m *S2CSouthGateStatus) String() string { return proto.CompactTextString(m) }
func (*S2CSouthGateStatus) ProtoMessage()    {}
func (m *S2CSouthGateStatus) GetId() uint16  { return PCK_S2CSouthGateStatus_ID }

const PCK_K2SActNineDayPreInfo_ID = 1572 //
//
type K2SActNineDayPreInfo struct {
	//
	Rk []*ActNineDayInfo `protobuf:"bytes,1,rep,name=Rk,proto3" json:"Rk,omitempty"`
	//
	DataType int32 `protobuf:"varint,2,opt,name=DataType,proto3" json:"DataType,omitempty"`
}

func (m *K2SActNineDayPreInfo) Reset()         { *m = K2SActNineDayPreInfo{} }
func (m *K2SActNineDayPreInfo) String() string { return proto.CompactTextString(m) }
func (*K2SActNineDayPreInfo) ProtoMessage()    {}
func (m *K2SActNineDayPreInfo) GetId() uint16  { return PCK_K2SActNineDayPreInfo_ID }

const PCK_C2SActNineDayPreInfo_ID = 1569 //
//
type C2SActNineDayPreInfo struct {
}

func (m *C2SActNineDayPreInfo) Reset()         { *m = C2SActNineDayPreInfo{} }
func (m *C2SActNineDayPreInfo) String() string { return proto.CompactTextString(m) }
func (*C2SActNineDayPreInfo) ProtoMessage()    {}
func (m *C2SActNineDayPreInfo) GetId() uint16  { return PCK_C2SActNineDayPreInfo_ID }

const PCK_S2CActNineDayPreInfo_ID = 1570 //
//
type S2CActNineDayPreInfo struct {
	//
	PreInfo *ActNineDayPreInfo `protobuf:"bytes,1,opt,name=PreInfo,proto3" json:"PreInfo,omitempty"`
}

func (m *S2CActNineDayPreInfo) Reset()         { *m = S2CActNineDayPreInfo{} }
func (m *S2CActNineDayPreInfo) String() string { return proto.CompactTextString(m) }
func (*S2CActNineDayPreInfo) ProtoMessage()    {}
func (m *S2CActNineDayPreInfo) GetId() uint16  { return PCK_S2CActNineDayPreInfo_ID }

//
type ActNineDayPreInfo struct {
	//
	Rk []*ActNineDayInfo `protobuf:"bytes,1,rep,name=Rk,proto3" json:"Rk,omitempty"`
	//
	WRk []*ActNineDayInfo `protobuf:"bytes,2,rep,name=WRk,proto3" json:"WRk,omitempty"`
	//
	Day int32 `protobuf:"varint,3,opt,name=Day,proto3" json:"Day,omitempty"`
}

func (m *ActNineDayPreInfo) Reset()         { *m = ActNineDayPreInfo{} }
func (m *ActNineDayPreInfo) String() string { return proto.CompactTextString(m) }
func (*ActNineDayPreInfo) ProtoMessage()    {}

//
type ActNineDayInfo struct {
	//
	SI int64 `protobuf:"varint,1,opt,name=SI,proto3" json:"SI,omitempty"`
	//
	LD int32 `protobuf:"varint,2,opt,name=LD,proto3" json:"LD,omitempty"`
	//
	SCV int64 `protobuf:"varint,3,opt,name=SCV,proto3" json:"SCV,omitempty"`
	//
	UId int64 `protobuf:"varint,4,opt,name=UId,proto3" json:"UId,omitempty"`
	//
	UN string `protobuf:"bytes,5,opt,name=UN,proto3" json:"UN,omitempty"`
	//
	A []*IntAttr `protobuf:"bytes,6,rep,name=A,proto3" json:"A,omitempty"`
	//
	B []*StrAttr `protobuf:"bytes,7,rep,name=B,proto3" json:"B,omitempty"`
	//
	CreateServerId int64 `protobuf:"varint,8,opt,name=CreateServerId,proto3" json:"CreateServerId,omitempty"`
}

func (m *ActNineDayInfo) Reset()         { *m = ActNineDayInfo{} }
func (m *ActNineDayInfo) String() string { return proto.CompactTextString(m) }
func (*ActNineDayInfo) ProtoMessage()    {}

const PCK_C2SActNineDaySort_ID = 1581 //
//
type C2SActNineDaySort struct {
}

func (m *C2SActNineDaySort) Reset()         { *m = C2SActNineDaySort{} }
func (m *C2SActNineDaySort) String() string { return proto.CompactTextString(m) }
func (*C2SActNineDaySort) ProtoMessage()    {}
func (m *C2SActNineDaySort) GetId() uint16  { return PCK_C2SActNineDaySort_ID }

const PCK_S2CActNineDaySort_ID = 1582 //
//
type S2CActNineDaySort struct {
	//
	Rk []*ActNineDayInfo `protobuf:"bytes,1,rep,name=Rk,proto3" json:"Rk,omitempty"`
	//
	Tag int32 `protobuf:"varint,2,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CActNineDaySort) Reset()         { *m = S2CActNineDaySort{} }
func (m *S2CActNineDaySort) String() string { return proto.CompactTextString(m) }
func (*S2CActNineDaySort) ProtoMessage()    {}
func (m *S2CActNineDaySort) GetId() uint16  { return PCK_S2CActNineDaySort_ID }

const PCK_S2CNiceDayClose_ID = 1579 //
//
type S2CNiceDayClose struct {
	//
	R int64 `protobuf:"varint,1,opt,name=R,proto3" json:"R,omitempty"`
	//
	F *ActNineDayInfo `protobuf:"bytes,2,opt,name=F,proto3" json:"F,omitempty"`
}

func (m *S2CNiceDayClose) Reset()         { *m = S2CNiceDayClose{} }
func (m *S2CNiceDayClose) String() string { return proto.CompactTextString(m) }
func (*S2CNiceDayClose) ProtoMessage()    {}
func (m *S2CNiceDayClose) GetId() uint16  { return PCK_S2CNiceDayClose_ID }

const PCK_K2SActGangWarPreInfo_ID = 1571 //
//
type K2SActGangWarPreInfo struct {
	//
	PreInfo *ActGangWarPreInfo `protobuf:"bytes,1,opt,name=PreInfo,proto3" json:"PreInfo,omitempty"`
}

func (m *K2SActGangWarPreInfo) Reset()         { *m = K2SActGangWarPreInfo{} }
func (m *K2SActGangWarPreInfo) String() string { return proto.CompactTextString(m) }
func (*K2SActGangWarPreInfo) ProtoMessage()    {}
func (m *K2SActGangWarPreInfo) GetId() uint16  { return PCK_K2SActGangWarPreInfo_ID }

const PCK_C2SActGangWarPreInfo_ID = 1567 //
//
type C2SActGangWarPreInfo struct {
}

func (m *C2SActGangWarPreInfo) Reset()         { *m = C2SActGangWarPreInfo{} }
func (m *C2SActGangWarPreInfo) String() string { return proto.CompactTextString(m) }
func (*C2SActGangWarPreInfo) ProtoMessage()    {}
func (m *C2SActGangWarPreInfo) GetId() uint16  { return PCK_C2SActGangWarPreInfo_ID }

const PCK_S2CActGangWarPreInfo_ID = 1568 //
//
type S2CActGangWarPreInfo struct {
	//
	PreInfo *ActGangWarPreInfo `protobuf:"bytes,1,opt,name=PreInfo,proto3" json:"PreInfo,omitempty"`
}

func (m *S2CActGangWarPreInfo) Reset()         { *m = S2CActGangWarPreInfo{} }
func (m *S2CActGangWarPreInfo) String() string { return proto.CompactTextString(m) }
func (*S2CActGangWarPreInfo) ProtoMessage()    {}
func (m *S2CActGangWarPreInfo) GetId() uint16  { return PCK_S2CActGangWarPreInfo_ID }

//
type ActGangWarPreInfo struct {
	//
	SI int64 `protobuf:"varint,1,opt,name=SI,proto3" json:"SI,omitempty"`
	//
	UId int64 `protobuf:"varint,2,opt,name=UId,proto3" json:"UId,omitempty"`
	//
	UN string `protobuf:"bytes,3,opt,name=UN,proto3" json:"UN,omitempty"`
	//
	GId string `protobuf:"bytes,4,opt,name=GId,proto3" json:"GId,omitempty"`
	//
	GN string `protobuf:"bytes,5,opt,name=GN,proto3" json:"GN,omitempty"`
	//
	SCV int64 `protobuf:"varint,6,opt,name=SCV,proto3" json:"SCV,omitempty"`
}

func (m *ActGangWarPreInfo) Reset()         { *m = ActGangWarPreInfo{} }
func (m *ActGangWarPreInfo) String() string { return proto.CompactTextString(m) }
func (*ActGangWarPreInfo) ProtoMessage()    {}

const PCK_S2CActStart_ID = 1561 //
//
type S2CActStart struct {
	//
	AId int32 `protobuf:"varint,1,opt,name=AId,proto3" json:"AId,omitempty"`
	//
	OT int64 `protobuf:"varint,2,opt,name=OT,proto3" json:"OT,omitempty"`
}

func (m *S2CActStart) Reset()         { *m = S2CActStart{} }
func (m *S2CActStart) String() string { return proto.CompactTextString(m) }
func (*S2CActStart) ProtoMessage()    {}
func (m *S2CActStart) GetId() uint16  { return PCK_S2CActStart_ID }

const PCK_S2CActOver_ID = 1565 //
//
type S2CActOver struct {
	//
	AId int32 `protobuf:"varint,1,opt,name=AId,proto3" json:"AId,omitempty"`
}

func (m *S2CActOver) Reset()         { *m = S2CActOver{} }
func (m *S2CActOver) String() string { return proto.CompactTextString(m) }
func (*S2CActOver) ProtoMessage()    {}
func (m *S2CActOver) GetId() uint16  { return PCK_S2CActOver_ID }

const PCK_S2CActIcon_ID = 1566 //
//
type S2CActIcon struct {
	//
	AIds []*ActIcon `protobuf:"bytes,1,rep,name=AIds,proto3" json:"AIds,omitempty"`
}

func (m *S2CActIcon) Reset()         { *m = S2CActIcon{} }
func (m *S2CActIcon) String() string { return proto.CompactTextString(m) }
func (*S2CActIcon) ProtoMessage()    {}
func (m *S2CActIcon) GetId() uint16  { return PCK_S2CActIcon_ID }

//
type ActIcon struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	B int64 `protobuf:"varint,2,opt,name=B,proto3" json:"B,omitempty"`
	//
	E int64 `protobuf:"varint,3,opt,name=E,proto3" json:"E,omitempty"`
}

func (m *ActIcon) Reset()         { *m = ActIcon{} }
func (m *ActIcon) String() string { return proto.CompactTextString(m) }
func (*ActIcon) ProtoMessage()    {}

const PCK_C2SActGangWarPersonRank_ID = 1548 //
//
type C2SActGangWarPersonRank struct {
}

func (m *C2SActGangWarPersonRank) Reset()         { *m = C2SActGangWarPersonRank{} }
func (m *C2SActGangWarPersonRank) String() string { return proto.CompactTextString(m) }
func (*C2SActGangWarPersonRank) ProtoMessage()    {}
func (m *C2SActGangWarPersonRank) GetId() uint16  { return PCK_C2SActGangWarPersonRank_ID }

const PCK_S2CActGangWarPersonRank_ID = 1549 //
//
type S2CActGangWarPersonRank struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	RS []*ActGangWarPersonRank `protobuf:"bytes,2,rep,name=RS,proto3" json:"RS,omitempty"`
}

func (m *S2CActGangWarPersonRank) Reset()         { *m = S2CActGangWarPersonRank{} }
func (m *S2CActGangWarPersonRank) String() string { return proto.CompactTextString(m) }
func (*S2CActGangWarPersonRank) ProtoMessage()    {}
func (m *S2CActGangWarPersonRank) GetId() uint16  { return PCK_S2CActGangWarPersonRank_ID }

const PCK_ActGangWarPersonRank_ID = 1550 //
//
type ActGangWarPersonRank struct {
	//
	GN string `protobuf:"bytes,1,opt,name=GN,proto3" json:"GN,omitempty"`
	//
	UN string `protobuf:"bytes,2,opt,name=UN,proto3" json:"UN,omitempty"`
	//
	H int64 `protobuf:"varint,3,opt,name=H,proto3" json:"H,omitempty"`
	//
	UId int64 `protobuf:"varint,4,opt,name=UId,proto3" json:"UId,omitempty"`
	//
	SId int64 `protobuf:"varint,5,opt,name=SId,proto3" json:"SId,omitempty"`
}

func (m *ActGangWarPersonRank) Reset()         { *m = ActGangWarPersonRank{} }
func (m *ActGangWarPersonRank) String() string { return proto.CompactTextString(m) }
func (*ActGangWarPersonRank) ProtoMessage()    {}
func (m *ActGangWarPersonRank) GetId() uint16  { return PCK_ActGangWarPersonRank_ID }

const PCK_C2SActGangWarGangRank_ID = 1551 //
//
type C2SActGangWarGangRank struct {
}

func (m *C2SActGangWarGangRank) Reset()         { *m = C2SActGangWarGangRank{} }
func (m *C2SActGangWarGangRank) String() string { return proto.CompactTextString(m) }
func (*C2SActGangWarGangRank) ProtoMessage()    {}
func (m *C2SActGangWarGangRank) GetId() uint16  { return PCK_C2SActGangWarGangRank_ID }

const PCK_S2CActGangWarGangRank_ID = 1552 //
//
type S2CActGangWarGangRank struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	GS []*ActGangWarGangRank `protobuf:"bytes,2,rep,name=GS,proto3" json:"GS,omitempty"`
}

func (m *S2CActGangWarGangRank) Reset()         { *m = S2CActGangWarGangRank{} }
func (m *S2CActGangWarGangRank) String() string { return proto.CompactTextString(m) }
func (*S2CActGangWarGangRank) ProtoMessage()    {}
func (m *S2CActGangWarGangRank) GetId() uint16  { return PCK_S2CActGangWarGangRank_ID }

const PCK_ActGangWarGangRank_ID = 1553 //
//
type ActGangWarGangRank struct {
	//
	GN string `protobuf:"bytes,1,opt,name=GN,proto3" json:"GN,omitempty"`
	//
	H int64 `protobuf:"varint,2,opt,name=H,proto3" json:"H,omitempty"`
	//
	GId string `protobuf:"bytes,3,opt,name=GId,proto3" json:"GId,omitempty"`
	//
	SId int64 `protobuf:"varint,4,opt,name=SId,proto3" json:"SId,omitempty"`
}

func (m *ActGangWarGangRank) Reset()         { *m = ActGangWarGangRank{} }
func (m *ActGangWarGangRank) String() string { return proto.CompactTextString(m) }
func (*ActGangWarGangRank) ProtoMessage()    {}
func (m *ActGangWarGangRank) GetId() uint16  { return PCK_ActGangWarGangRank_ID }

const PCK_C2SActGangWarKillRank_ID = 1554 //
//
type C2SActGangWarKillRank struct {
}

func (m *C2SActGangWarKillRank) Reset()         { *m = C2SActGangWarKillRank{} }
func (m *C2SActGangWarKillRank) String() string { return proto.CompactTextString(m) }
func (*C2SActGangWarKillRank) ProtoMessage()    {}
func (m *C2SActGangWarKillRank) GetId() uint16  { return PCK_C2SActGangWarKillRank_ID }

const PCK_S2CActGangWarKillRank_ID = 1555 //
//
type S2CActGangWarKillRank struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	KS []*ActGangWarKillRank `protobuf:"bytes,2,rep,name=KS,proto3" json:"KS,omitempty"`
}

func (m *S2CActGangWarKillRank) Reset()         { *m = S2CActGangWarKillRank{} }
func (m *S2CActGangWarKillRank) String() string { return proto.CompactTextString(m) }
func (*S2CActGangWarKillRank) ProtoMessage()    {}
func (m *S2CActGangWarKillRank) GetId() uint16  { return PCK_S2CActGangWarKillRank_ID }

const PCK_ActGangWarKillRank_ID = 1556 //
//
type ActGangWarKillRank struct {
	//
	GN string `protobuf:"bytes,1,opt,name=GN,proto3" json:"GN,omitempty"`
	//
	UN string `protobuf:"bytes,2,opt,name=UN,proto3" json:"UN,omitempty"`
	//
	K int64 `protobuf:"varint,3,opt,name=K,proto3" json:"K,omitempty"`
	//
	UId int64 `protobuf:"varint,4,opt,name=UId,proto3" json:"UId,omitempty"`
	//
	SId int64 `protobuf:"varint,5,opt,name=SId,proto3" json:"SId,omitempty"`
}

func (m *ActGangWarKillRank) Reset()         { *m = ActGangWarKillRank{} }
func (m *ActGangWarKillRank) String() string { return proto.CompactTextString(m) }
func (*ActGangWarKillRank) ProtoMessage()    {}
func (m *ActGangWarKillRank) GetId() uint16  { return PCK_ActGangWarKillRank_ID }

const PCK_C2SActGangWarScoreRank_ID = 1557 //
//
type C2SActGangWarScoreRank struct {
}

func (m *C2SActGangWarScoreRank) Reset()         { *m = C2SActGangWarScoreRank{} }
func (m *C2SActGangWarScoreRank) String() string { return proto.CompactTextString(m) }
func (*C2SActGangWarScoreRank) ProtoMessage()    {}
func (m *C2SActGangWarScoreRank) GetId() uint16  { return PCK_C2SActGangWarScoreRank_ID }

const PCK_S2CActGangWarScoreRank_ID = 1558 //
//
type S2CActGangWarScoreRank struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	SS []*ActGangWarScoreRank `protobuf:"bytes,2,rep,name=SS,proto3" json:"SS,omitempty"`
}

func (m *S2CActGangWarScoreRank) Reset()         { *m = S2CActGangWarScoreRank{} }
func (m *S2CActGangWarScoreRank) String() string { return proto.CompactTextString(m) }
func (*S2CActGangWarScoreRank) ProtoMessage()    {}
func (m *S2CActGangWarScoreRank) GetId() uint16  { return PCK_S2CActGangWarScoreRank_ID }

const PCK_ActGangWarScoreRank_ID = 1559 //
//
type ActGangWarScoreRank struct {
	//
	GN string `protobuf:"bytes,1,opt,name=GN,proto3" json:"GN,omitempty"`
	//
	UN string `protobuf:"bytes,2,opt,name=UN,proto3" json:"UN,omitempty"`
	//
	S int64 `protobuf:"varint,3,opt,name=S,proto3" json:"S,omitempty"`
	//
	UId int64 `protobuf:"varint,4,opt,name=UId,proto3" json:"UId,omitempty"`
	//
	SId int64 `protobuf:"varint,5,opt,name=SId,proto3" json:"SId,omitempty"`
}

func (m *ActGangWarScoreRank) Reset()         { *m = ActGangWarScoreRank{} }
func (m *ActGangWarScoreRank) String() string { return proto.CompactTextString(m) }
func (*ActGangWarScoreRank) ProtoMessage()    {}
func (m *ActGangWarScoreRank) GetId() uint16  { return PCK_ActGangWarScoreRank_ID }

const PCK_C2SActGangWarScoreExchange_ID = 1545 //
//
type C2SActGangWarScoreExchange struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2SActGangWarScoreExchange) Reset()         { *m = C2SActGangWarScoreExchange{} }
func (m *C2SActGangWarScoreExchange) String() string { return proto.CompactTextString(m) }
func (*C2SActGangWarScoreExchange) ProtoMessage()    {}
func (m *C2SActGangWarScoreExchange) GetId() uint16  { return PCK_C2SActGangWarScoreExchange_ID }

const PCK_S2CActGangWarScoreExchange_ID = 1546 //
//
type S2CActGangWarScoreExchange struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Ids []int32 `protobuf:"varint,2,rep,name=Ids,proto3" json:"Ids,omitempty"`
}

func (m *S2CActGangWarScoreExchange) Reset()         { *m = S2CActGangWarScoreExchange{} }
func (m *S2CActGangWarScoreExchange) String() string { return proto.CompactTextString(m) }
func (*S2CActGangWarScoreExchange) ProtoMessage()    {}
func (m *S2CActGangWarScoreExchange) GetId() uint16  { return PCK_S2CActGangWarScoreExchange_ID }

const PCK_C2SActGangWarScoreFind_ID = 1513 //
//
type C2SActGangWarScoreFind struct {
}

func (m *C2SActGangWarScoreFind) Reset()         { *m = C2SActGangWarScoreFind{} }
func (m *C2SActGangWarScoreFind) String() string { return proto.CompactTextString(m) }
func (*C2SActGangWarScoreFind) ProtoMessage()    {}
func (m *C2SActGangWarScoreFind) GetId() uint16  { return PCK_C2SActGangWarScoreFind_ID }

const PCK_S2CActGangWarScoreFind_ID = 1514 //
//
type S2CActGangWarScoreFind struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Ids []int32 `protobuf:"varint,2,rep,name=Ids,proto3" json:"Ids,omitempty"`
}

func (m *S2CActGangWarScoreFind) Reset()         { *m = S2CActGangWarScoreFind{} }
func (m *S2CActGangWarScoreFind) String() string { return proto.CompactTextString(m) }
func (*S2CActGangWarScoreFind) ProtoMessage()    {}
func (m *S2CActGangWarScoreFind) GetId() uint16  { return PCK_S2CActGangWarScoreFind_ID }

const PCK_C2SActGangWarSwitchMap_ID = 1543 //
//
type C2SActGangWarSwitchMap struct {
	//
	AT int64 `protobuf:"varint,1,opt,name=AT,proto3" json:"AT,omitempty"`
}

func (m *C2SActGangWarSwitchMap) Reset()         { *m = C2SActGangWarSwitchMap{} }
func (m *C2SActGangWarSwitchMap) String() string { return proto.CompactTextString(m) }
func (*C2SActGangWarSwitchMap) ProtoMessage()    {}
func (m *C2SActGangWarSwitchMap) GetId() uint16  { return PCK_C2SActGangWarSwitchMap_ID }

const PCK_S2CActGangWarSwitchMap_ID = 1544 //
//
type S2CActGangWarSwitchMap struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CActGangWarSwitchMap) Reset()         { *m = S2CActGangWarSwitchMap{} }
func (m *S2CActGangWarSwitchMap) String() string { return proto.CompactTextString(m) }
func (*S2CActGangWarSwitchMap) ProtoMessage()    {}
func (m *S2CActGangWarSwitchMap) GetId() uint16  { return PCK_S2CActGangWarSwitchMap_ID }

const PCK_S2CActGangWarInDragonGangs_ID = 1541 //
//
type S2CActGangWarInDragonGangs struct {
	//
	GS []*ActGangWarInDragonGang `protobuf:"bytes,1,rep,name=GS,proto3" json:"GS,omitempty"`
}

func (m *S2CActGangWarInDragonGangs) Reset()         { *m = S2CActGangWarInDragonGangs{} }
func (m *S2CActGangWarInDragonGangs) String() string { return proto.CompactTextString(m) }
func (*S2CActGangWarInDragonGangs) ProtoMessage()    {}
func (m *S2CActGangWarInDragonGangs) GetId() uint16  { return PCK_S2CActGangWarInDragonGangs_ID }

//
type ActGangWarInDragonGang struct {
	//
	Gid string `protobuf:"bytes,1,opt,name=Gid,proto3" json:"Gid,omitempty"`
	//
	GN string `protobuf:"bytes,2,opt,name=GN,proto3" json:"GN,omitempty"`
	//
	SId int64 `protobuf:"varint,3,opt,name=SId,proto3" json:"SId,omitempty"`
}

func (m *ActGangWarInDragonGang) Reset()         { *m = ActGangWarInDragonGang{} }
func (m *ActGangWarInDragonGang) String() string { return proto.CompactTextString(m) }
func (*ActGangWarInDragonGang) ProtoMessage()    {}

const PCK_C2SActGiveUpDragon_ID = 1540 //
//
type C2SActGiveUpDragon struct {
}

func (m *C2SActGiveUpDragon) Reset()         { *m = C2SActGiveUpDragon{} }
func (m *C2SActGiveUpDragon) String() string { return proto.CompactTextString(m) }
func (*C2SActGiveUpDragon) ProtoMessage()    {}
func (m *C2SActGiveUpDragon) GetId() uint16  { return PCK_C2SActGiveUpDragon_ID }

const PCK_S2CActGiveUpDragon_ID = 1539 //
//
type S2CActGiveUpDragon struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CActGiveUpDragon) Reset()         { *m = S2CActGiveUpDragon{} }
func (m *S2CActGiveUpDragon) String() string { return proto.CompactTextString(m) }
func (*S2CActGiveUpDragon) ProtoMessage()    {}
func (m *S2CActGiveUpDragon) GetId() uint16  { return PCK_S2CActGiveUpDragon_ID }

const PCK_C2SActCreateTeam_ID = 1527 //
//
type C2SActCreateTeam struct {
	//
	TeamType int32 `protobuf:"varint,1,opt,name=TeamType,proto3" json:"TeamType,omitempty"`
}

func (m *C2SActCreateTeam) Reset()         { *m = C2SActCreateTeam{} }
func (m *C2SActCreateTeam) String() string { return proto.CompactTextString(m) }
func (*C2SActCreateTeam) ProtoMessage()    {}
func (m *C2SActCreateTeam) GetId() uint16  { return PCK_C2SActCreateTeam_ID }

const PCK_S2CActCreateTeam_ID = 1528 //
//
type S2CActCreateTeam struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	TI *ActTeamInfo `protobuf:"bytes,2,opt,name=TI,proto3" json:"TI,omitempty"`
}

func (m *S2CActCreateTeam) Reset()         { *m = S2CActCreateTeam{} }
func (m *S2CActCreateTeam) String() string { return proto.CompactTextString(m) }
func (*S2CActCreateTeam) ProtoMessage()    {}
func (m *S2CActCreateTeam) GetId() uint16  { return PCK_S2CActCreateTeam_ID }

const PCK_C2SActFindTeams_ID = 1529 //
//
type C2SActFindTeams struct {
	//
	TeamType int32 `protobuf:"varint,1,opt,name=TeamType,proto3" json:"TeamType,omitempty"`
}

func (m *C2SActFindTeams) Reset()         { *m = C2SActFindTeams{} }
func (m *C2SActFindTeams) String() string { return proto.CompactTextString(m) }
func (*C2SActFindTeams) ProtoMessage()    {}
func (m *C2SActFindTeams) GetId() uint16  { return PCK_C2SActFindTeams_ID }

const PCK_S2CActFindTeams_ID = 1530 //
//
type S2CActFindTeams struct {
	//
	TIS []*ActTeamInfo `protobuf:"bytes,1,rep,name=TIS,proto3" json:"TIS,omitempty"`
	//
	Tag int32 `protobuf:"varint,2,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CActFindTeams) Reset()         { *m = S2CActFindTeams{} }
func (m *S2CActFindTeams) String() string { return proto.CompactTextString(m) }
func (*S2CActFindTeams) ProtoMessage()    {}
func (m *S2CActFindTeams) GetId() uint16  { return PCK_S2CActFindTeams_ID }

const PCK_C2SActJoinTeam_ID = 1531 //
//
type C2SActJoinTeam struct {
	//
	TId int64 `protobuf:"varint,1,opt,name=TId,proto3" json:"TId,omitempty"`
}

func (m *C2SActJoinTeam) Reset()         { *m = C2SActJoinTeam{} }
func (m *C2SActJoinTeam) String() string { return proto.CompactTextString(m) }
func (*C2SActJoinTeam) ProtoMessage()    {}
func (m *C2SActJoinTeam) GetId() uint16  { return PCK_C2SActJoinTeam_ID }

const PCK_S2CActJoinTeam_ID = 1532 //
//
type S2CActJoinTeam struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	TI *ActTeamInfo `protobuf:"bytes,2,opt,name=TI,proto3" json:"TI,omitempty"`
}

func (m *S2CActJoinTeam) Reset()         { *m = S2CActJoinTeam{} }
func (m *S2CActJoinTeam) String() string { return proto.CompactTextString(m) }
func (*S2CActJoinTeam) ProtoMessage()    {}
func (m *S2CActJoinTeam) GetId() uint16  { return PCK_S2CActJoinTeam_ID }

const PCK_C2SActFindTeam_ID = 1533 //
//
type C2SActFindTeam struct {
}

func (m *C2SActFindTeam) Reset()         { *m = C2SActFindTeam{} }
func (m *C2SActFindTeam) String() string { return proto.CompactTextString(m) }
func (*C2SActFindTeam) ProtoMessage()    {}
func (m *C2SActFindTeam) GetId() uint16  { return PCK_C2SActFindTeam_ID }

const PCK_S2CActFindTeam_ID = 1534 //
//
type S2CActFindTeam struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	TI *ActTeamInfo `protobuf:"bytes,2,opt,name=TI,proto3" json:"TI,omitempty"`
}

func (m *S2CActFindTeam) Reset()         { *m = S2CActFindTeam{} }
func (m *S2CActFindTeam) String() string { return proto.CompactTextString(m) }
func (*S2CActFindTeam) ProtoMessage()    {}
func (m *S2CActFindTeam) GetId() uint16  { return PCK_S2CActFindTeam_ID }

const PCK_C2SActNeedFightVal_ID = 1537 //
//
type C2SActNeedFightVal struct {
	//
	FV int64 `protobuf:"varint,1,opt,name=FV,proto3" json:"FV,omitempty"`
}

func (m *C2SActNeedFightVal) Reset()         { *m = C2SActNeedFightVal{} }
func (m *C2SActNeedFightVal) String() string { return proto.CompactTextString(m) }
func (*C2SActNeedFightVal) ProtoMessage()    {}
func (m *C2SActNeedFightVal) GetId() uint16  { return PCK_C2SActNeedFightVal_ID }

const PCK_S2CActNeedFightVal_ID = 1538 //
//
type S2CActNeedFightVal struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	TI *ActTeamInfo `protobuf:"bytes,2,opt,name=TI,proto3" json:"TI,omitempty"`
}

func (m *S2CActNeedFightVal) Reset()         { *m = S2CActNeedFightVal{} }
func (m *S2CActNeedFightVal) String() string { return proto.CompactTextString(m) }
func (*S2CActNeedFightVal) ProtoMessage()    {}
func (m *S2CActNeedFightVal) GetId() uint16  { return PCK_S2CActNeedFightVal_ID }

const PCK_C2SActExitTeam_ID = 1576 //
//
type C2SActExitTeam struct {
	//
	Uid int64 `protobuf:"varint,1,opt,name=Uid,proto3" json:"Uid,omitempty"`
}

func (m *C2SActExitTeam) Reset()         { *m = C2SActExitTeam{} }
func (m *C2SActExitTeam) String() string { return proto.CompactTextString(m) }
func (*C2SActExitTeam) ProtoMessage()    {}
func (m *C2SActExitTeam) GetId() uint16  { return PCK_C2SActExitTeam_ID }

const PCK_S2CActExitTeam_ID = 1577 //
//
type S2CActExitTeam struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Uid int64 `protobuf:"varint,2,opt,name=Uid,proto3" json:"Uid,omitempty"`
}

func (m *S2CActExitTeam) Reset()         { *m = S2CActExitTeam{} }
func (m *S2CActExitTeam) String() string { return proto.CompactTextString(m) }
func (*S2CActExitTeam) ProtoMessage()    {}
func (m *S2CActExitTeam) GetId() uint16  { return PCK_S2CActExitTeam_ID }

const PCK_ActTeamInfo_ID = 1535 //
//
type ActTeamInfo struct {
	//
	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	NFV int64 `protobuf:"varint,2,opt,name=NFV,proto3" json:"NFV,omitempty"`
	//
	MS []*ActTeamMemInfo `protobuf:"bytes,4,rep,name=MS,proto3" json:"MS,omitempty"`
}

func (m *ActTeamInfo) Reset()         { *m = ActTeamInfo{} }
func (m *ActTeamInfo) String() string { return proto.CompactTextString(m) }
func (*ActTeamInfo) ProtoMessage()    {}
func (m *ActTeamInfo) GetId() uint16  { return PCK_ActTeamInfo_ID }

const PCK_ActTeamMemInfo_ID = 1536 //
//
type ActTeamMemInfo struct {
	//
	UId int64 `protobuf:"varint,1,opt,name=UId,proto3" json:"UId,omitempty"`
	//
	UN string `protobuf:"bytes,2,opt,name=UN,proto3" json:"UN,omitempty"`
	//
	L int32 `protobuf:"varint,3,opt,name=L,proto3" json:"L,omitempty"`
	//
	RId int32 `protobuf:"varint,4,opt,name=RId,proto3" json:"RId,omitempty"`
	//
	LV int64 `protobuf:"varint,5,opt,name=LV,proto3" json:"LV,omitempty"`
	//
	FV int64 `protobuf:"varint,6,opt,name=FV,proto3" json:"FV,omitempty"`
	//
	MHP int64 `protobuf:"varint,7,opt,name=MHP,proto3" json:"MHP,omitempty"`
	//
	HP int64 `protobuf:"varint,8,opt,name=HP,proto3" json:"HP,omitempty"`
}

func (m *ActTeamMemInfo) Reset()         { *m = ActTeamMemInfo{} }
func (m *ActTeamMemInfo) String() string { return proto.CompactTextString(m) }
func (*ActTeamMemInfo) ProtoMessage()    {}
func (m *ActTeamMemInfo) GetId() uint16  { return PCK_ActTeamMemInfo_ID }

const PCK_C2SActTeamRecruit_ID = 1515 //
//
type C2SActTeamRecruit struct {
}

func (m *C2SActTeamRecruit) Reset()         { *m = C2SActTeamRecruit{} }
func (m *C2SActTeamRecruit) String() string { return proto.CompactTextString(m) }
func (*C2SActTeamRecruit) ProtoMessage()    {}
func (m *C2SActTeamRecruit) GetId() uint16  { return PCK_C2SActTeamRecruit_ID }

const PCK_S2CActTeamRecruit_ID = 1516 //
//
type S2CActTeamRecruit struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CActTeamRecruit) Reset()         { *m = S2CActTeamRecruit{} }
func (m *S2CActTeamRecruit) String() string { return proto.CompactTextString(m) }
func (*S2CActTeamRecruit) ProtoMessage()    {}
func (m *S2CActTeamRecruit) GetId() uint16  { return PCK_S2CActTeamRecruit_ID }

const PCK_C2SJoinActive_ID = 1507 //
//
type C2SJoinActive struct {
	//
	AId int32 `protobuf:"varint,1,opt,name=AId,proto3" json:"AId,omitempty"`
}

func (m *C2SJoinActive) Reset()         { *m = C2SJoinActive{} }
func (m *C2SJoinActive) String() string { return proto.CompactTextString(m) }
func (*C2SJoinActive) ProtoMessage()    {}
func (m *C2SJoinActive) GetId() uint16  { return PCK_C2SJoinActive_ID }

const PCK_S2CJoinActive_ID = 1508 //
//
type S2CJoinActive struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	AId int32 `protobuf:"varint,2,opt,name=AId,proto3" json:"AId,omitempty"`
}

func (m *S2CJoinActive) Reset()         { *m = S2CJoinActive{} }
func (m *S2CJoinActive) String() string { return proto.CompactTextString(m) }
func (*S2CJoinActive) ProtoMessage()    {}
func (m *S2CJoinActive) GetId() uint16  { return PCK_S2CJoinActive_ID }

const PCK_C2SLeaveActive_ID = 1509 //
//
type C2SLeaveActive struct {
	//
	AId int32 `protobuf:"varint,1,opt,name=AId,proto3" json:"AId,omitempty"`
}

func (m *C2SLeaveActive) Reset()         { *m = C2SLeaveActive{} }
func (m *C2SLeaveActive) String() string { return proto.CompactTextString(m) }
func (*C2SLeaveActive) ProtoMessage()    {}
func (m *C2SLeaveActive) GetId() uint16  { return PCK_C2SLeaveActive_ID }

const PCK_S2CLeaveActive_ID = 1510 //
//
type S2CLeaveActive struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CLeaveActive) Reset()         { *m = S2CLeaveActive{} }
func (m *S2CLeaveActive) String() string { return proto.CompactTextString(m) }
func (*S2CLeaveActive) ProtoMessage()    {}
func (m *S2CLeaveActive) GetId() uint16  { return PCK_S2CLeaveActive_ID }

const PCK_K2SJoinFuncActiveHandleEvent_ID = 1517 //
//
type K2SJoinFuncActiveHandleEvent struct {
	//
	ActType int32 `protobuf:"varint,1,opt,name=ActType,proto3" json:"ActType,omitempty"`
}

func (m *K2SJoinFuncActiveHandleEvent) Reset()         { *m = K2SJoinFuncActiveHandleEvent{} }
func (m *K2SJoinFuncActiveHandleEvent) String() string { return proto.CompactTextString(m) }
func (*K2SJoinFuncActiveHandleEvent) ProtoMessage()    {}
func (m *K2SJoinFuncActiveHandleEvent) GetId() uint16  { return PCK_K2SJoinFuncActiveHandleEvent_ID }

const PCK_S2CGangWarDragonPosition_ID = 1511 //
//
type S2CGangWarDragonPosition struct {
	//
	LUId int64 `protobuf:"varint,1,opt,name=LUId,proto3" json:"LUId,omitempty"`
	//
	Us []*DragonPosition `protobuf:"bytes,2,rep,name=Us,proto3" json:"Us,omitempty"`
	//
	Mhp int64 `protobuf:"varint,3,opt,name=Mhp,proto3" json:"Mhp,omitempty"`
	//
	Hp int64 `protobuf:"varint,4,opt,name=Hp,proto3" json:"Hp,omitempty"`
}

func (m *S2CGangWarDragonPosition) Reset()         { *m = S2CGangWarDragonPosition{} }
func (m *S2CGangWarDragonPosition) String() string { return proto.CompactTextString(m) }
func (*S2CGangWarDragonPosition) ProtoMessage()    {}
func (m *S2CGangWarDragonPosition) GetId() uint16  { return PCK_S2CGangWarDragonPosition_ID }

//
type DragonPosition struct {
	//
	UId int64 `protobuf:"varint,1,opt,name=UId,proto3" json:"UId,omitempty"`
	//
	UN string `protobuf:"bytes,2,opt,name=UN,proto3" json:"UN,omitempty"`
	//
	SId int64 `protobuf:"varint,3,opt,name=SId,proto3" json:"SId,omitempty"`
	//
	RId int64 `protobuf:"varint,4,opt,name=RId,proto3" json:"RId,omitempty"`
	//
	Mhp int64 `protobuf:"varint,5,opt,name=Mhp,proto3" json:"Mhp,omitempty"`
	//
	Hp int64 `protobuf:"varint,6,opt,name=Hp,proto3" json:"Hp,omitempty"`
}

func (m *DragonPosition) Reset()         { *m = DragonPosition{} }
func (m *DragonPosition) String() string { return proto.CompactTextString(m) }
func (*DragonPosition) ProtoMessage()    {}

const PCK_S2CGangWarGangSort_ID = 1526 //
//
type S2CGangWarGangSort struct {
	//
	Infos []*GangHurtInfo `protobuf:"bytes,1,rep,name=Infos,proto3" json:"Infos,omitempty"`
	//
	RId int32 `protobuf:"varint,2,opt,name=RId,proto3" json:"RId,omitempty"`
	//
	V int64 `protobuf:"varint,3,opt,name=V,proto3" json:"V,omitempty"`
	//
	T int32 `protobuf:"varint,4,opt,name=T,proto3" json:"T,omitempty"`
}

func (m *S2CGangWarGangSort) Reset()         { *m = S2CGangWarGangSort{} }
func (m *S2CGangWarGangSort) String() string { return proto.CompactTextString(m) }
func (*S2CGangWarGangSort) ProtoMessage()    {}
func (m *S2CGangWarGangSort) GetId() uint16  { return PCK_S2CGangWarGangSort_ID }

const PCK_GangHurtInfo_ID = 1512 //
//
type GangHurtInfo struct {
	//
	GId string `protobuf:"bytes,1,opt,name=GId,proto3" json:"GId,omitempty"`
	//
	SId int64 `protobuf:"varint,3,opt,name=SId,proto3" json:"SId,omitempty"`
	//
	GN string `protobuf:"bytes,4,opt,name=GN,proto3" json:"GN,omitempty"`
	//
	H int64 `protobuf:"varint,5,opt,name=H,proto3" json:"H,omitempty"`
	//
	PC int64 `protobuf:"varint,6,opt,name=PC,proto3" json:"PC,omitempty"`
	//
	APC int64 `protobuf:"varint,7,opt,name=APC,proto3" json:"APC,omitempty"`
	//
	SC int64 `protobuf:"varint,8,opt,name=SC,proto3" json:"SC,omitempty"`
}

func (m *GangHurtInfo) Reset()         { *m = GangHurtInfo{} }
func (m *GangHurtInfo) String() string { return proto.CompactTextString(m) }
func (*GangHurtInfo) ProtoMessage()    {}
func (m *GangHurtInfo) GetId() uint16  { return PCK_GangHurtInfo_ID }

const PCK_S2CGangActWarGangPCount_ID = 1503 //
//
type S2CGangActWarGangPCount struct {
	//
	GId string `protobuf:"bytes,1,opt,name=GId,proto3" json:"GId,omitempty"`
	//
	PC int64 `protobuf:"varint,2,opt,name=PC,proto3" json:"PC,omitempty"`
}

func (m *S2CGangActWarGangPCount) Reset()         { *m = S2CGangActWarGangPCount{} }
func (m *S2CGangActWarGangPCount) String() string { return proto.CompactTextString(m) }
func (*S2CGangActWarGangPCount) ProtoMessage()    {}
func (m *S2CGangActWarGangPCount) GetId() uint16  { return PCK_S2CGangActWarGangPCount_ID }

const PCK_S2CGangActWarGangLocal_ID = 1504 //
//
type S2CGangActWarGangLocal struct {
	//
	L int32 `protobuf:"varint,1,opt,name=L,proto3" json:"L,omitempty"`
}

func (m *S2CGangActWarGangLocal) Reset()         { *m = S2CGangActWarGangLocal{} }
func (m *S2CGangActWarGangLocal) String() string { return proto.CompactTextString(m) }
func (*S2CGangActWarGangLocal) ProtoMessage()    {}
func (m *S2CGangActWarGangLocal) GetId() uint16  { return PCK_S2CGangActWarGangLocal_ID }

const PCK_S2CActiveGangWarScoreChange_ID = 1580 //
//
type S2CActiveGangWarScoreChange struct {
	//
	V int64 `protobuf:"varint,1,opt,name=V,proto3" json:"V,omitempty"`
}

func (m *S2CActiveGangWarScoreChange) Reset()         { *m = S2CActiveGangWarScoreChange{} }
func (m *S2CActiveGangWarScoreChange) String() string { return proto.CompactTextString(m) }
func (*S2CActiveGangWarScoreChange) ProtoMessage()    {}
func (m *S2CActiveGangWarScoreChange) GetId() uint16  { return PCK_S2CActiveGangWarScoreChange_ID }

const PCK_S2CActiveGangWarBeforeInsideScore_ID = 1585 //
//
type S2CActiveGangWarBeforeInsideScore struct {
	//
	V int64 `protobuf:"varint,1,opt,name=V,proto3" json:"V,omitempty"`
}

func (m *S2CActiveGangWarBeforeInsideScore) Reset()         { *m = S2CActiveGangWarBeforeInsideScore{} }
func (m *S2CActiveGangWarBeforeInsideScore) String() string { return proto.CompactTextString(m) }
func (*S2CActiveGangWarBeforeInsideScore) ProtoMessage()    {}
func (m *S2CActiveGangWarBeforeInsideScore) GetId() uint16 {
	return PCK_S2CActiveGangWarBeforeInsideScore_ID
}

const PCK_C2SReviveLife_ID = 1501 //
//
type C2SReviveLife struct {
	//
	ReviveType int32 `protobuf:"varint,1,opt,name=ReviveType,proto3" json:"ReviveType,omitempty"`
}

func (m *C2SReviveLife) Reset()         { *m = C2SReviveLife{} }
func (m *C2SReviveLife) String() string { return proto.CompactTextString(m) }
func (*C2SReviveLife) ProtoMessage()    {}
func (m *C2SReviveLife) GetId() uint16  { return PCK_C2SReviveLife_ID }

const PCK_S2CReviveLife_ID = 1502 //
//
type S2CReviveLife struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	ReviveType int32 `protobuf:"varint,2,opt,name=ReviveType,proto3" json:"ReviveType,omitempty"`
}

func (m *S2CReviveLife) Reset()         { *m = S2CReviveLife{} }
func (m *S2CReviveLife) String() string { return proto.CompactTextString(m) }
func (*S2CReviveLife) ProtoMessage()    {}
func (m *S2CReviveLife) GetId() uint16  { return PCK_S2CReviveLife_ID }

const PCK_C2SGangBossOverTime_ID = 1283 //
//
type C2SGangBossOverTime struct {
}

func (m *C2SGangBossOverTime) Reset()         { *m = C2SGangBossOverTime{} }
func (m *C2SGangBossOverTime) String() string { return proto.CompactTextString(m) }
func (*C2SGangBossOverTime) ProtoMessage()    {}
func (m *C2SGangBossOverTime) GetId() uint16  { return PCK_C2SGangBossOverTime_ID }

const PCK_S2CGangBossOverTime_ID = 1284 //
//
type S2CGangBossOverTime struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	OT int64 `protobuf:"varint,2,opt,name=OT,proto3" json:"OT,omitempty"`
	//
	RT int64 `protobuf:"varint,3,opt,name=RT,proto3" json:"RT,omitempty"`
	//
	S int64 `protobuf:"varint,4,opt,name=S,proto3" json:"S,omitempty"`
	//
	AS int32 `protobuf:"varint,5,opt,name=AS,proto3" json:"AS,omitempty"`
}

func (m *S2CGangBossOverTime) Reset()         { *m = S2CGangBossOverTime{} }
func (m *S2CGangBossOverTime) String() string { return proto.CompactTextString(m) }
func (*S2CGangBossOverTime) ProtoMessage()    {}
func (m *S2CGangBossOverTime) GetId() uint16  { return PCK_S2CGangBossOverTime_ID }

const PCK_K2SAddGangRedPkgMoney_ID = 1291 //
//
type K2SAddGangRedPkgMoney struct {
	//
	M int64 `protobuf:"varint,1,opt,name=M,proto3" json:"M,omitempty"`
	//
	Gid string `protobuf:"bytes,2,opt,name=Gid,proto3" json:"Gid,omitempty"`
}

func (m *K2SAddGangRedPkgMoney) Reset()         { *m = K2SAddGangRedPkgMoney{} }
func (m *K2SAddGangRedPkgMoney) String() string { return proto.CompactTextString(m) }
func (*K2SAddGangRedPkgMoney) ProtoMessage()    {}
func (m *K2SAddGangRedPkgMoney) GetId() uint16  { return PCK_K2SAddGangRedPkgMoney_ID }

const PCK_C2SGangBossHurtSort_ID = 1290 //
//
type C2SGangBossHurtSort struct {
}

func (m *C2SGangBossHurtSort) Reset()         { *m = C2SGangBossHurtSort{} }
func (m *C2SGangBossHurtSort) String() string { return proto.CompactTextString(m) }
func (*C2SGangBossHurtSort) ProtoMessage()    {}
func (m *C2SGangBossHurtSort) GetId() uint16  { return PCK_C2SGangBossHurtSort_ID }

const PCK_S2CGangBossHurtSort_ID = 1285 //
//
type S2CGangBossHurtSort struct {
	//
	HTs []*GangBossHurt `protobuf:"bytes,1,rep,name=HTs,proto3" json:"HTs,omitempty"`
	//
	Tag int32 `protobuf:"varint,2,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CGangBossHurtSort) Reset()         { *m = S2CGangBossHurtSort{} }
func (m *S2CGangBossHurtSort) String() string { return proto.CompactTextString(m) }
func (*S2CGangBossHurtSort) ProtoMessage()    {}
func (m *S2CGangBossHurtSort) GetId() uint16  { return PCK_S2CGangBossHurtSort_ID }

const PCK_S2CGangBossFiveHurt_ID = 1299 //
//
type S2CGangBossFiveHurt struct {
	//
	HTs []*GangBossHurt `protobuf:"bytes,1,rep,name=HTs,proto3" json:"HTs,omitempty"`
}

func (m *S2CGangBossFiveHurt) Reset()         { *m = S2CGangBossFiveHurt{} }
func (m *S2CGangBossFiveHurt) String() string { return proto.CompactTextString(m) }
func (*S2CGangBossFiveHurt) ProtoMessage()    {}
func (m *S2CGangBossFiveHurt) GetId() uint16  { return PCK_S2CGangBossFiveHurt_ID }

//
type GangBossHurt struct {
	//
	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	N string `protobuf:"bytes,2,opt,name=N,proto3" json:"N,omitempty"`
	//
	RId int32 `protobuf:"varint,3,opt,name=RId,proto3" json:"RId,omitempty"`
	//
	Hurt int64 `protobuf:"varint,4,opt,name=Hurt,proto3" json:"Hurt,omitempty"`
}

func (m *GangBossHurt) Reset()         { *m = GangBossHurt{} }
func (m *GangBossHurt) String() string { return proto.CompactTextString(m) }
func (*GangBossHurt) ProtoMessage()    {}

const PCK_C2SGangBossKiller_ID = 1292 //
//
type C2SGangBossKiller struct {
}

func (m *C2SGangBossKiller) Reset()         { *m = C2SGangBossKiller{} }
func (m *C2SGangBossKiller) String() string { return proto.CompactTextString(m) }
func (*C2SGangBossKiller) ProtoMessage()    {}
func (m *C2SGangBossKiller) GetId() uint16  { return PCK_C2SGangBossKiller_ID }

const PCK_S2CGangBossKiller_ID = 1293 //
//
type S2CGangBossKiller struct {
	//
	N string `protobuf:"bytes,1,opt,name=N,proto3" json:"N,omitempty"`
	//
	GangName string `protobuf:"bytes,2,opt,name=GangName,proto3" json:"GangName,omitempty"`
}

func (m *S2CGangBossKiller) Reset()         { *m = S2CGangBossKiller{} }
func (m *S2CGangBossKiller) String() string { return proto.CompactTextString(m) }
func (*S2CGangBossKiller) ProtoMessage()    {}
func (m *S2CGangBossKiller) GetId() uint16  { return PCK_S2CGangBossKiller_ID }

const PCK_S2CGangBossRedPot_ID = 1289 //
//
type S2CGangBossRedPot struct {
	//
	RP bool `protobuf:"varint,1,opt,name=RP,proto3" json:"RP,omitempty"`
}

func (m *S2CGangBossRedPot) Reset()         { *m = S2CGangBossRedPot{} }
func (m *S2CGangBossRedPot) String() string { return proto.CompactTextString(m) }
func (*S2CGangBossRedPot) ProtoMessage()    {}
func (m *S2CGangBossRedPot) GetId() uint16  { return PCK_S2CGangBossRedPot_ID }

const PCK_C2SGangCollect_ID = 1300 //
//
type C2SGangCollect struct {
}

func (m *C2SGangCollect) Reset()         { *m = C2SGangCollect{} }
func (m *C2SGangCollect) String() string { return proto.CompactTextString(m) }
func (*C2SGangCollect) ProtoMessage()    {}
func (m *C2SGangCollect) GetId() uint16  { return PCK_C2SGangCollect_ID }

const PCK_S2CGangCollect_ID = 1301 //
//
type S2CGangCollect struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CGangCollect) Reset()         { *m = S2CGangCollect{} }
func (m *S2CGangCollect) String() string { return proto.CompactTextString(m) }
func (*S2CGangCollect) ProtoMessage()    {}
func (m *S2CGangCollect) GetId() uint16  { return PCK_S2CGangCollect_ID }

//
type ActiveNineDayChangeInstance struct {
}

func (m *ActiveNineDayChangeInstance) Reset()         { *m = ActiveNineDayChangeInstance{} }
func (m *ActiveNineDayChangeInstance) String() string { return proto.CompactTextString(m) }
func (*ActiveNineDayChangeInstance) ProtoMessage()    {}

const PCK_C2SActiveNineDayFindExchange_ID = 1562 //
//
type C2SActiveNineDayFindExchange struct {
}

func (m *C2SActiveNineDayFindExchange) Reset()         { *m = C2SActiveNineDayFindExchange{} }
func (m *C2SActiveNineDayFindExchange) String() string { return proto.CompactTextString(m) }
func (*C2SActiveNineDayFindExchange) ProtoMessage()    {}
func (m *C2SActiveNineDayFindExchange) GetId() uint16  { return PCK_C2SActiveNineDayFindExchange_ID }

const PCK_S2CActiveNineDayFindExchange_ID = 1563 //
//
type S2CActiveNineDayFindExchange struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *S2CActiveNineDayFindExchange) Reset()         { *m = S2CActiveNineDayFindExchange{} }
func (m *S2CActiveNineDayFindExchange) String() string { return proto.CompactTextString(m) }
func (*S2CActiveNineDayFindExchange) ProtoMessage()    {}
func (m *S2CActiveNineDayFindExchange) GetId() uint16  { return PCK_S2CActiveNineDayFindExchange_ID }

const PCK_S2CActiveNineDayScoreChange_ID = 1564 //
//
type S2CActiveNineDayScoreChange struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	V int32 `protobuf:"varint,2,opt,name=V,proto3" json:"V,omitempty"`
}

func (m *S2CActiveNineDayScoreChange) Reset()         { *m = S2CActiveNineDayScoreChange{} }
func (m *S2CActiveNineDayScoreChange) String() string { return proto.CompactTextString(m) }
func (*S2CActiveNineDayScoreChange) ProtoMessage()    {}
func (m *S2CActiveNineDayScoreChange) GetId() uint16  { return PCK_S2CActiveNineDayScoreChange_ID }

const PCK_C2SActiveNineDayReviceExchange_ID = 1505 //
//
type C2SActiveNineDayReviceExchange struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2SActiveNineDayReviceExchange) Reset()         { *m = C2SActiveNineDayReviceExchange{} }
func (m *C2SActiveNineDayReviceExchange) String() string { return proto.CompactTextString(m) }
func (*C2SActiveNineDayReviceExchange) ProtoMessage()    {}
func (m *C2SActiveNineDayReviceExchange) GetId() uint16  { return PCK_C2SActiveNineDayReviceExchange_ID }

const PCK_S2CActiveNineDayReviceExchange_ID = 1506 //
//
type S2CActiveNineDayReviceExchange struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *S2CActiveNineDayReviceExchange) Reset()         { *m = S2CActiveNineDayReviceExchange{} }
func (m *S2CActiveNineDayReviceExchange) String() string { return proto.CompactTextString(m) }
func (*S2CActiveNineDayReviceExchange) ProtoMessage()    {}
func (m *S2CActiveNineDayReviceExchange) GetId() uint16  { return PCK_S2CActiveNineDayReviceExchange_ID }

const PCK_C2SGoldTree_ID = 8201 //
//
type C2SGoldTree struct {
}

func (m *C2SGoldTree) Reset()         { *m = C2SGoldTree{} }
func (m *C2SGoldTree) String() string { return proto.CompactTextString(m) }
func (*C2SGoldTree) ProtoMessage()    {}
func (m *C2SGoldTree) GetId() uint16  { return PCK_C2SGoldTree_ID }

const PCK_S2CGoldTree_ID = 8202 //
//
type S2CGoldTree struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Gold int32 `protobuf:"varint,3,opt,name=Gold,proto3" json:"Gold,omitempty"`
	//
	NextMulti int32 `protobuf:"varint,2,opt,name=NextMulti,proto3" json:"NextMulti,omitempty"`
}

func (m *S2CGoldTree) Reset()         { *m = S2CGoldTree{} }
func (m *S2CGoldTree) String() string { return proto.CompactTextString(m) }
func (*S2CGoldTree) ProtoMessage()    {}
func (m *S2CGoldTree) GetId() uint16  { return PCK_S2CGoldTree_ID }

const PCK_C2SGetGoldTreeInfo_ID = 8203 //
//
type C2SGetGoldTreeInfo struct {
}

func (m *C2SGetGoldTreeInfo) Reset()         { *m = C2SGetGoldTreeInfo{} }
func (m *C2SGetGoldTreeInfo) String() string { return proto.CompactTextString(m) }
func (*C2SGetGoldTreeInfo) ProtoMessage()    {}
func (m *C2SGetGoldTreeInfo) GetId() uint16  { return PCK_C2SGetGoldTreeInfo_ID }

const PCK_S2CGetGoldTreeInfo_ID = 8204 //
//
type S2CGetGoldTreeInfo struct {
	//
	TodayUse int32 `protobuf:"varint,1,opt,name=TodayUse,proto3" json:"TodayUse,omitempty"`
	//
	TodayTotal int32 `protobuf:"varint,2,opt,name=TodayTotal,proto3" json:"TodayTotal,omitempty"`
	//
	NextVipTimes int32 `protobuf:"varint,3,opt,name=NextVipTimes,proto3" json:"NextVipTimes,omitempty"`
	//
	ActId int32 `protobuf:"varint,4,opt,name=ActId,proto3" json:"ActId,omitempty"`
}

func (m *S2CGetGoldTreeInfo) Reset()         { *m = S2CGetGoldTreeInfo{} }
func (m *S2CGetGoldTreeInfo) String() string { return proto.CompactTextString(m) }
func (*S2CGetGoldTreeInfo) ProtoMessage()    {}
func (m *S2CGetGoldTreeInfo) GetId() uint16  { return PCK_S2CGetGoldTreeInfo_ID }

//
type NoticeUser struct {
	//
	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Nick string `protobuf:"bytes,2,opt,name=Nick,proto3" json:"Nick,omitempty"`
	//
	AreaId int32 `protobuf:"varint,3,opt,name=AreaId,proto3" json:"AreaId,omitempty"`
}

func (m *NoticeUser) Reset()         { *m = NoticeUser{} }
func (m *NoticeUser) String() string { return proto.CompactTextString(m) }
func (*NoticeUser) ProtoMessage()    {}

//
type NoticeItem struct {
	//
	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Num int64 `protobuf:"varint,2,opt,name=Num,proto3" json:"Num,omitempty"`
	//
	OnlyId string `protobuf:"bytes,3,opt,name=OnlyId,proto3" json:"OnlyId,omitempty"`
}

func (m *NoticeItem) Reset()         { *m = NoticeItem{} }
func (m *NoticeItem) String() string { return proto.CompactTextString(m) }
func (*NoticeItem) ProtoMessage()    {}

//
type NoticeFaBao struct {
	//
	IId int32 `protobuf:"varint,1,opt,name=IId,proto3" json:"IId,omitempty"`
	//
	Id string `protobuf:"bytes,2,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Quality int32 `protobuf:"varint,3,opt,name=Quality,proto3" json:"Quality,omitempty"`
}

func (m *NoticeFaBao) Reset()         { *m = NoticeFaBao{} }
func (m *NoticeFaBao) String() string { return proto.CompactTextString(m) }
func (*NoticeFaBao) ProtoMessage()    {}

const PCK_S2CNotice_ID = 12 //
//
type S2CNotice struct {
	//
	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	U []*NoticeUser `protobuf:"bytes,2,rep,name=U,proto3" json:"U,omitempty"`
	//
	I []*NoticeItem `protobuf:"bytes,3,rep,name=I,proto3" json:"I,omitempty"`
	//
	Stage []int64 `protobuf:"varint,4,rep,name=Stage,proto3" json:"Stage,omitempty"`
	//
	P []string `protobuf:"bytes,5,rep,name=P,proto3" json:"P,omitempty"`
	//
	FaBao *NoticeFaBao `protobuf:"bytes,6,opt,name=FaBao,proto3" json:"FaBao,omitempty"`
	//
	Skill []int32 `protobuf:"varint,7,rep,name=Skill,proto3" json:"Skill,omitempty"`
	//
	ItemData []*ItemData `protobuf:"bytes,8,rep,name=ItemData,proto3" json:"ItemData,omitempty"`
}

func (m *S2CNotice) Reset()         { *m = S2CNotice{} }
func (m *S2CNotice) String() string { return proto.CompactTextString(m) }
func (*S2CNotice) ProtoMessage()    {}
func (m *S2CNotice) GetId() uint16  { return PCK_S2CNotice_ID }

//
type DragonLogItem struct {
	//
	Time int64 `protobuf:"varint,1,opt,name=Time,proto3" json:"Time,omitempty"`
	//
	UserId int64 `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	InstanceId int64 `protobuf:"varint,3,opt,name=InstanceId,proto3" json:"InstanceId,omitempty"`
	//
	Item []*ItemData `protobuf:"bytes,4,rep,name=Item,proto3" json:"Item,omitempty"`
	//
	DragonType int32 `protobuf:"varint,5,opt,name=DragonType,proto3" json:"DragonType,omitempty"`
}

func (m *DragonLogItem) Reset()         { *m = DragonLogItem{} }
func (m *DragonLogItem) String() string { return proto.CompactTextString(m) }
func (*DragonLogItem) ProtoMessage()    {}

//
type DragonLogItemClient struct {
	//
	Time int64 `protobuf:"varint,1,opt,name=Time,proto3" json:"Time,omitempty"`
	//
	User *NoticeUser `protobuf:"bytes,2,opt,name=User,proto3" json:"User,omitempty"`
	//
	InstanceId int64 `protobuf:"varint,3,opt,name=InstanceId,proto3" json:"InstanceId,omitempty"`
	//
	Item []*ItemData `protobuf:"bytes,4,rep,name=Item,proto3" json:"Item,omitempty"`
}

func (m *DragonLogItemClient) Reset()         { *m = DragonLogItemClient{} }
func (m *DragonLogItemClient) String() string { return proto.CompactTextString(m) }
func (*DragonLogItemClient) ProtoMessage()    {}

const PCK_C2SGetDragonLog_ID = 1701 //
//
type C2SGetDragonLog struct {
	//
	Type int32 `protobuf:"varint,1,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (m *C2SGetDragonLog) Reset()         { *m = C2SGetDragonLog{} }
func (m *C2SGetDragonLog) String() string { return proto.CompactTextString(m) }
func (*C2SGetDragonLog) ProtoMessage()    {}
func (m *C2SGetDragonLog) GetId() uint16  { return PCK_C2SGetDragonLog_ID }

const PCK_S2CGetDragonLog_ID = 1702 //
//
type S2CGetDragonLog struct {
	//
	Type int32 `protobuf:"varint,1,opt,name=Type,proto3" json:"Type,omitempty"`
	//
	Logs []*DragonLogItemClient `protobuf:"bytes,2,rep,name=Logs,proto3" json:"Logs,omitempty"`
}

func (m *S2CGetDragonLog) Reset()         { *m = S2CGetDragonLog{} }
func (m *S2CGetDragonLog) String() string { return proto.CompactTextString(m) }
func (*S2CGetDragonLog) ProtoMessage()    {}
func (m *S2CGetDragonLog) GetId() uint16  { return PCK_S2CGetDragonLog_ID }

const PCK_C2SStartFightDragon_ID = 1703 //
//
type C2SStartFightDragon struct {
	//
	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Auto int64 `protobuf:"varint,2,opt,name=Auto,proto3" json:"Auto,omitempty"`
	//
	ItemId int64 `protobuf:"varint,3,opt,name=ItemId,proto3" json:"ItemId,omitempty"`
	//
	Num int64 `protobuf:"varint,4,opt,name=Num,proto3" json:"Num,omitempty"`
	//
	Jump int64 `protobuf:"varint,5,opt,name=Jump,proto3" json:"Jump,omitempty"`
	//
	Guide int64 `protobuf:"varint,6,opt,name=Guide,proto3" json:"Guide,omitempty"`
}

func (m *C2SStartFightDragon) Reset()         { *m = C2SStartFightDragon{} }
func (m *C2SStartFightDragon) String() string { return proto.CompactTextString(m) }
func (*C2SStartFightDragon) ProtoMessage()    {}
func (m *C2SStartFightDragon) GetId() uint16  { return PCK_C2SStartFightDragon_ID }

const PCK_S2CStartFightDragon_ID = 1704 //
//
type S2CStartFightDragon struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id int64 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	ItemId int64 `protobuf:"varint,3,opt,name=ItemId,proto3" json:"ItemId,omitempty"`
}

func (m *S2CStartFightDragon) Reset()         { *m = S2CStartFightDragon{} }
func (m *S2CStartFightDragon) String() string { return proto.CompactTextString(m) }
func (*S2CStartFightDragon) ProtoMessage()    {}
func (m *S2CStartFightDragon) GetId() uint16  { return PCK_S2CStartFightDragon_ID }

//
type MiracleItem struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Multi int32 `protobuf:"varint,2,opt,name=Multi,proto3" json:"Multi,omitempty"`
	//
	State int32 `protobuf:"varint,3,opt,name=State,proto3" json:"State,omitempty"`
}

func (m *MiracleItem) Reset()         { *m = MiracleItem{} }
func (m *MiracleItem) String() string { return proto.CompactTextString(m) }
func (*MiracleItem) ProtoMessage()    {}

//
type ActPlayerData struct {
	//
	GoodsInfo []*BuyGoods `protobuf:"bytes,6,rep,name=GoodsInfo,proto3" json:"GoodsInfo,omitempty"`
	//
	Log []int32 `protobuf:"varint,9,rep,name=Log,proto3" json:"Log,omitempty"`
	//
	BossData []*PlayerBossData `protobuf:"bytes,13,rep,name=BossData,proto3" json:"BossData,omitempty"`
	//
	LuckDraw1 []*LuckDraw `protobuf:"bytes,14,rep,name=LuckDraw1,proto3" json:"LuckDraw1,omitempty"`
	//
	PicturePart []int32 `protobuf:"varint,15,rep,name=PicturePart,proto3" json:"PicturePart,omitempty"`
	//
	DareBossData []*PlayerDareBossData `protobuf:"bytes,17,rep,name=DareBossData,proto3" json:"DareBossData,omitempty"`
	//
	ActId int32 `protobuf:"varint,20,opt,name=ActId,proto3" json:"ActId,omitempty"`
	//
	ReturnDetail []*ReturnDetail `protobuf:"bytes,23,rep,name=ReturnDetail,proto3" json:"ReturnDetail,omitempty"`
	//
	ReturnLog []*ChargeReturnLog `protobuf:"bytes,24,rep,name=ReturnLog,proto3" json:"ReturnLog,omitempty"`
	//
	ChargeReturnPrizeCache *ChargeReturnPrizeCache `protobuf:"bytes,30,opt,name=ChargeReturnPrizeCache,proto3" json:"ChargeReturnPrizeCache,omitempty"`
	//
	IntAttr []*IntAttr `protobuf:"bytes,40,rep,name=IntAttr,proto3" json:"IntAttr,omitempty"`
	//
	MiracleItem []*MiracleItem `protobuf:"bytes,41,rep,name=MiracleItem,proto3" json:"MiracleItem,omitempty"`
}

func (m *ActPlayerData) Reset()         { *m = ActPlayerData{} }
func (m *ActPlayerData) String() string { return proto.CompactTextString(m) }
func (*ActPlayerData) ProtoMessage()    {}

const PCK_S2CContinueRechargeTaskNum_ID = 8307 //
//
type S2CContinueRechargeTaskNum struct {
	//
	Number int32 `protobuf:"varint,1,opt,name=Number,proto3" json:"Number,omitempty"`
}

func (m *S2CContinueRechargeTaskNum) Reset()         { *m = S2CContinueRechargeTaskNum{} }
func (m *S2CContinueRechargeTaskNum) String() string { return proto.CompactTextString(m) }
func (*S2CContinueRechargeTaskNum) ProtoMessage()    {}
func (m *S2CContinueRechargeTaskNum) GetId() uint16  { return PCK_S2CContinueRechargeTaskNum_ID }

//
type BuyGoods struct {
	//
	Gid int32 `protobuf:"varint,1,opt,name=Gid,proto3" json:"Gid,omitempty"`
	//
	BuyTime int64 `protobuf:"varint,2,opt,name=BuyTime,proto3" json:"BuyTime,omitempty"`
	//
	Status int32 `protobuf:"varint,3,opt,name=Status,proto3" json:"Status,omitempty"`
	//
	BuyNumTimes int32 `protobuf:"varint,4,opt,name=BuyNumTimes,proto3" json:"BuyNumTimes,omitempty"`
}

func (m *BuyGoods) Reset()         { *m = BuyGoods{} }
func (m *BuyGoods) String() string { return proto.CompactTextString(m) }
func (*BuyGoods) ProtoMessage()    {}

//
type Activity struct {
	//
	ActId int32 `protobuf:"varint,1,opt,name=ActId,proto3" json:"ActId,omitempty"`
	//
	Type int32 `protobuf:"varint,2,opt,name=Type,proto3" json:"Type,omitempty"`
	//
	InitTime int32 `protobuf:"varint,3,opt,name=InitTime,proto3" json:"InitTime,omitempty"`
	//
	StartTime int64 `protobuf:"varint,4,opt,name=StartTime,proto3" json:"StartTime,omitempty"`
	//
	EndTime int64 `protobuf:"varint,5,opt,name=EndTime,proto3" json:"EndTime,omitempty"`
	//
	Pos int32 `protobuf:"varint,6,opt,name=Pos,proto3" json:"Pos,omitempty"`
	//
	Container int32 `protobuf:"varint,7,opt,name=Container,proto3" json:"Container,omitempty"`
	//
	RedPoint int32 `protobuf:"varint,8,opt,name=RedPoint,proto3" json:"RedPoint,omitempty"`
	//
	Name string `protobuf:"bytes,10,opt,name=Name,proto3" json:"Name,omitempty"`
	//
	Order int32 `protobuf:"varint,11,opt,name=Order,proto3" json:"Order,omitempty"`
	//
	Icon []*ActivityIcon `protobuf:"bytes,12,rep,name=Icon,proto3" json:"Icon,omitempty"`
	//
	RMB int32 `protobuf:"varint,13,opt,name=RMB,proto3" json:"RMB,omitempty"`
	//
	Param int32 `protobuf:"varint,14,opt,name=Param,proto3" json:"Param,omitempty"`
	//
	Banner int32 `protobuf:"varint,15,opt,name=Banner,proto3" json:"Banner,omitempty"`
}

func (m *Activity) Reset()         { *m = Activity{} }
func (m *Activity) String() string { return proto.CompactTextString(m) }
func (*Activity) ProtoMessage()    {}

//
type ActivityIcon struct {
	//
	Pos int32 `protobuf:"varint,1,opt,name=Pos,proto3" json:"Pos,omitempty"`
	//
	Container int32 `protobuf:"varint,2,opt,name=Container,proto3" json:"Container,omitempty"`
	//
	Order int32 `protobuf:"varint,3,opt,name=Order,proto3" json:"Order,omitempty"`
	//
	Icon int32 `protobuf:"varint,4,opt,name=Icon,proto3" json:"Icon,omitempty"`
	//
	Circle int32 `protobuf:"varint,5,opt,name=Circle,proto3" json:"Circle,omitempty"`
	//
	NoHide int32 `protobuf:"varint,6,opt,name=NoHide,proto3" json:"NoHide,omitempty"`
}

func (m *ActivityIcon) Reset()         { *m = ActivityIcon{} }
func (m *ActivityIcon) String() string { return proto.CompactTextString(m) }
func (*ActivityIcon) ProtoMessage()    {}

//
type NoticePoint struct {
	//
	TimePoint int32 `protobuf:"varint,1,opt,name=TimePoint,proto3" json:"TimePoint,omitempty"`
	//
	LoopNum int32 `protobuf:"varint,2,opt,name=LoopNum,proto3" json:"LoopNum,omitempty"`
	//
	Sended bool `protobuf:"varint,3,opt,name=Sended,proto3" json:"Sended,omitempty"`
}

func (m *NoticePoint) Reset()         { *m = NoticePoint{} }
func (m *NoticePoint) String() string { return proto.CompactTextString(m) }
func (*NoticePoint) ProtoMessage()    {}

const PCK_S2CAllActivity_ID = 8300 //
//
type S2CAllActivity struct {
	//
	Activity []*Activity `protobuf:"bytes,1,rep,name=Activity,proto3" json:"Activity,omitempty"`
}

func (m *S2CAllActivity) Reset()         { *m = S2CAllActivity{} }
func (m *S2CAllActivity) String() string { return proto.CompactTextString(m) }
func (*S2CAllActivity) ProtoMessage()    {}
func (m *S2CAllActivity) GetId() uint16  { return PCK_S2CAllActivity_ID }

const PCK_S2CActivityInit_ID = 8303 //
//
type S2CActivityInit struct {
	//
	Activity []*Activity `protobuf:"bytes,1,rep,name=Activity,proto3" json:"Activity,omitempty"`
}

func (m *S2CActivityInit) Reset()         { *m = S2CActivityInit{} }
func (m *S2CActivityInit) String() string { return proto.CompactTextString(m) }
func (*S2CActivityInit) ProtoMessage()    {}
func (m *S2CActivityInit) GetId() uint16  { return PCK_S2CActivityInit_ID }

const PCK_S2CActivityStart_ID = 8301 //
//
type S2CActivityStart struct {
	//
	Activity []*Activity `protobuf:"bytes,1,rep,name=Activity,proto3" json:"Activity,omitempty"`
}

func (m *S2CActivityStart) Reset()         { *m = S2CActivityStart{} }
func (m *S2CActivityStart) String() string { return proto.CompactTextString(m) }
func (*S2CActivityStart) ProtoMessage()    {}
func (m *S2CActivityStart) GetId() uint16  { return PCK_S2CActivityStart_ID }

const PCK_S2CActivityEnd_ID = 8302 //
//
type S2CActivityEnd struct {
	//
	ActId []int32 `protobuf:"varint,1,rep,name=ActId,proto3" json:"ActId,omitempty"`
}

func (m *S2CActivityEnd) Reset()         { *m = S2CActivityEnd{} }
func (m *S2CActivityEnd) String() string { return proto.CompactTextString(m) }
func (*S2CActivityEnd) ProtoMessage()    {}
func (m *S2CActivityEnd) GetId() uint16  { return PCK_S2CActivityEnd_ID }

const PCK_S2CActivityRedPoint_ID = 8304 //
//
type S2CActivityRedPoint struct {
	//
	ActId int32 `protobuf:"varint,1,opt,name=ActId,proto3" json:"ActId,omitempty"`
	//
	RedPoint int32 `protobuf:"varint,2,opt,name=RedPoint,proto3" json:"RedPoint,omitempty"`
}

func (m *S2CActivityRedPoint) Reset()         { *m = S2CActivityRedPoint{} }
func (m *S2CActivityRedPoint) String() string { return proto.CompactTextString(m) }
func (*S2CActivityRedPoint) ProtoMessage()    {}
func (m *S2CActivityRedPoint) GetId() uint16  { return PCK_S2CActivityRedPoint_ID }

const PCK_S2CActivityIcon_ID = 8306 //
//
type S2CActivityIcon struct {
	//
	ActId int32 `protobuf:"varint,1,opt,name=ActId,proto3" json:"ActId,omitempty"`
	//
	Pos int32 `protobuf:"varint,2,opt,name=Pos,proto3" json:"Pos,omitempty"`
	//
	Icon int32 `protobuf:"varint,3,opt,name=Icon,proto3" json:"Icon,omitempty"`
}

func (m *S2CActivityIcon) Reset()         { *m = S2CActivityIcon{} }
func (m *S2CActivityIcon) String() string { return proto.CompactTextString(m) }
func (*S2CActivityIcon) ProtoMessage()    {}
func (m *S2CActivityIcon) GetId() uint16  { return PCK_S2CActivityIcon_ID }

const PCK_S2CActivityData_ID = 8305 //
//
type S2CActivityData struct {
	//
	S2CGetDrawData []*S2CGetDrawData `protobuf:"bytes,1,rep,name=S2CGetDrawData,proto3" json:"S2CGetDrawData,omitempty"`
	//
	S2CPlayerDrawData []*S2CPlayerDrawData `protobuf:"bytes,2,rep,name=S2CPlayerDrawData,proto3" json:"S2CPlayerDrawData,omitempty"`
	//
	S2CGetActShopList []*S2CGetActShopList `protobuf:"bytes,4,rep,name=S2CGetActShopList,proto3" json:"S2CGetActShopList,omitempty"`
	//
	S2CGetActGiftList []*S2CGetActGiftList `protobuf:"bytes,5,rep,name=S2CGetActGiftList,proto3" json:"S2CGetActGiftList,omitempty"`
	//
	S2CGetActBossInfo []*S2CGetActBossInfo `protobuf:"bytes,6,rep,name=S2CGetActBossInfo,proto3" json:"S2CGetActBossInfo,omitempty"`
	//
	S2CPlayerBossInfo []*S2CPlayerBossInfo `protobuf:"bytes,7,rep,name=S2CPlayerBossInfo,proto3" json:"S2CPlayerBossInfo,omitempty"`
	//
	S2CGetActTask []*S2CGetActTask `protobuf:"bytes,8,rep,name=S2CGetActTask,proto3" json:"S2CGetActTask,omitempty"`
	//
	S2CGetInvestInfo []*S2CGetInvestInfo `protobuf:"bytes,9,rep,name=S2CGetInvestInfo,proto3" json:"S2CGetInvestInfo,omitempty"`
	//
	S2CPlayerInvestData []*S2CPlayerInvestData `protobuf:"bytes,10,rep,name=S2CPlayerInvestData,proto3" json:"S2CPlayerInvestData,omitempty"`
	//
	S2CGetActRankInfo []*S2CGetActRankInfo `protobuf:"bytes,11,rep,name=S2CGetActRankInfo,proto3" json:"S2CGetActRankInfo,omitempty"`
	//
	S2CGetGoldTreeInfo []*S2CGetGoldTreeInfo `protobuf:"bytes,12,rep,name=S2CGetGoldTreeInfo,proto3" json:"S2CGetGoldTreeInfo,omitempty"`
	//
	S2CGetActPicture []*S2CGetActPicture `protobuf:"bytes,13,rep,name=S2CGetActPicture,proto3" json:"S2CGetActPicture,omitempty"`
	//
	S2CGetChargeReturnData []*S2CGetChargeReturnData `protobuf:"bytes,14,rep,name=S2CGetChargeReturnData,proto3" json:"S2CGetChargeReturnData,omitempty"`
	//
	S2CActRecoveryData []*S2CActRecoveryData `protobuf:"bytes,15,rep,name=S2CActRecoveryData,proto3" json:"S2CActRecoveryData,omitempty"`
	//
	S2CGetActGiftLimit []*S2CGetActGiftLimit `protobuf:"bytes,16,rep,name=S2CGetActGiftLimit,proto3" json:"S2CGetActGiftLimit,omitempty"`
	//
	S2CGetActGiftOnline []*S2CGetActGiftOnline `protobuf:"bytes,17,rep,name=S2CGetActGiftOnline,proto3" json:"S2CGetActGiftOnline,omitempty"`
	//
	S2CGetActGiftMiracle []*S2CGetActGiftMiracle `protobuf:"bytes,18,rep,name=S2CGetActGiftMiracle,proto3" json:"S2CGetActGiftMiracle,omitempty"`
	//
	S2CGetActGiftXRmbInfo []*S2CGetActGiftXRmbInfo `protobuf:"bytes,19,rep,name=S2CGetActGiftXRmbInfo,proto3" json:"S2CGetActGiftXRmbInfo,omitempty"`
	//
	S2CGetActCollectFontInfo []*S2CGetActCollectFontInfo `protobuf:"bytes,20,rep,name=S2CGetActCollectFontInfo,proto3" json:"S2CGetActCollectFontInfo,omitempty"`
	//
	S2CGetActLifeVipGiveInfo []*S2CGetActLifeVipGiveInfo `protobuf:"bytes,21,rep,name=S2CGetActLifeVipGiveInfo,proto3" json:"S2CGetActLifeVipGiveInfo,omitempty"`
	//
	S2CGetActGift1Info []*S2CGetActGift1Info `protobuf:"bytes,22,rep,name=S2CGetActGift1Info,proto3" json:"S2CGetActGift1Info,omitempty"`
	//
	S2CGetActGift2Info []*S2CGetActGift2Info `protobuf:"bytes,23,rep,name=S2CGetActGift2Info,proto3" json:"S2CGetActGift2Info,omitempty"`
	//
	S2CGetActCumulativeInfo []*S2CGetActCumulativeInfo `protobuf:"bytes,24,rep,name=S2CGetActCumulativeInfo,proto3" json:"S2CGetActCumulativeInfo,omitempty"`
	//
	S2CGetActChargeDayInfo []*S2CGetActChargeDayInfo `protobuf:"bytes,25,rep,name=S2CGetActChargeDayInfo,proto3" json:"S2CGetActChargeDayInfo,omitempty"`
	//
	S2CCumulativeFirstInfo []*S2CCumulativeFirstInfo `protobuf:"bytes,26,rep,name=S2CCumulativeFirstInfo,proto3" json:"S2CCumulativeFirstInfo,omitempty"`
	//
	S2CGetInvestmentInfo []*S2CGetInvestmentInfo `protobuf:"bytes,27,rep,name=S2CGetInvestmentInfo,proto3" json:"S2CGetInvestmentInfo,omitempty"`
}

func (m *S2CActivityData) Reset()         { *m = S2CActivityData{} }
func (m *S2CActivityData) String() string { return proto.CompactTextString(m) }
func (*S2CActivityData) ProtoMessage()    {}
func (m *S2CActivityData) GetId() uint16  { return PCK_S2CActivityData_ID }

const PCK_C2SGetDrawData_ID = 11007 //
//
type C2SGetDrawData struct {
	//
	ActId int32 `protobuf:"varint,1,opt,name=ActId,proto3" json:"ActId,omitempty"`
}

func (m *C2SGetDrawData) Reset()         { *m = C2SGetDrawData{} }
func (m *C2SGetDrawData) String() string { return proto.CompactTextString(m) }
func (*C2SGetDrawData) ProtoMessage()    {}
func (m *C2SGetDrawData) GetId() uint16  { return PCK_C2SGetDrawData_ID }

const PCK_S2CGetDrawData_ID = 11008 //
//
type S2CGetDrawData struct {
	//
	ActId int32 `protobuf:"varint,1,opt,name=ActId,proto3" json:"ActId,omitempty"`
	//
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	//
	Luck int32 `protobuf:"varint,3,opt,name=Luck,proto3" json:"Luck,omitempty"`
	//
	LowDraw string `protobuf:"bytes,4,opt,name=LowDraw,proto3" json:"LowDraw,omitempty"`
	//
	HighDrawOne string `protobuf:"bytes,5,opt,name=HighDrawOne,proto3" json:"HighDrawOne,omitempty"`
	//
	HighDrawTen string `protobuf:"bytes,6,opt,name=HighDrawTen,proto3" json:"HighDrawTen,omitempty"`
	//
	GaiID int32 `protobuf:"varint,7,opt,name=GaiID,proto3" json:"GaiID,omitempty"`
	//
	Detail []*DrawDetail `protobuf:"bytes,8,rep,name=Detail,proto3" json:"Detail,omitempty"`
	//
	DoubleRate int32 `protobuf:"varint,9,opt,name=DoubleRate,proto3" json:"DoubleRate,omitempty"`
}

func (m *S2CGetDrawData) Reset()         { *m = S2CGetDrawData{} }
func (m *S2CGetDrawData) String() string { return proto.CompactTextString(m) }
func (*S2CGetDrawData) ProtoMessage()    {}
func (m *S2CGetDrawData) GetId() uint16  { return PCK_S2CGetDrawData_ID }

//
type DrawDetail struct {
	//
	ID int32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	//
	ItemId int32 `protobuf:"varint,2,opt,name=ItemId,proto3" json:"ItemId,omitempty"`
	//
	ItemCount int32 `protobuf:"varint,3,opt,name=ItemCount,proto3" json:"ItemCount,omitempty"`
	//
	TitleType int32 `protobuf:"varint,4,opt,name=TitleType,proto3" json:"TitleType,omitempty"`
	//
	SortId int32 `protobuf:"varint,5,opt,name=SortId,proto3" json:"SortId,omitempty"`
	//
	ShowRate string `protobuf:"bytes,6,opt,name=ShowRate,proto3" json:"ShowRate,omitempty"`
	//
	Name string `protobuf:"bytes,7,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (m *DrawDetail) Reset()         { *m = DrawDetail{} }
func (m *DrawDetail) String() string { return proto.CompactTextString(m) }
func (*DrawDetail) ProtoMessage()    {}

const PCK_C2SPlayerDrawData_ID = 11009 //
//
type C2SPlayerDrawData struct {
	//
	ActId int32 `protobuf:"varint,1,opt,name=ActId,proto3" json:"ActId,omitempty"`
}

func (m *C2SPlayerDrawData) Reset()         { *m = C2SPlayerDrawData{} }
func (m *C2SPlayerDrawData) String() string { return proto.CompactTextString(m) }
func (*C2SPlayerDrawData) ProtoMessage()    {}
func (m *C2SPlayerDrawData) GetId() uint16  { return PCK_C2SPlayerDrawData_ID }

const PCK_S2CPlayerDrawData_ID = 11010 //
//
type S2CPlayerDrawData struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	ActId int32 `protobuf:"varint,2,opt,name=ActId,proto3" json:"ActId,omitempty"`
	//
	Luck int32 `protobuf:"varint,3,opt,name=Luck,proto3" json:"Luck,omitempty"`
	//
	Score int32 `protobuf:"varint,4,opt,name=Score,proto3" json:"Score,omitempty"`
	//
	Free int32 `protobuf:"varint,5,opt,name=Free,proto3" json:"Free,omitempty"`
	//
	MustDrawId int32 `protobuf:"varint,6,opt,name=MustDrawId,proto3" json:"MustDrawId,omitempty"`
	//
	ActType int32 `protobuf:"varint,7,opt,name=ActType,proto3" json:"ActType,omitempty"`
	//
	LuckDraw []*LuckDraw `protobuf:"bytes,8,rep,name=LuckDraw,proto3" json:"LuckDraw,omitempty"`
	//
	LowDrawTimes int32 `protobuf:"varint,9,opt,name=LowDrawTimes,proto3" json:"LowDrawTimes,omitempty"`
	//
	DrawTimes int32 `protobuf:"varint,10,opt,name=DrawTimes,proto3" json:"DrawTimes,omitempty"`
	//
	Discount int32 `protobuf:"varint,11,opt,name=Discount,proto3" json:"Discount,omitempty"`
}

func (m *S2CPlayerDrawData) Reset()         { *m = S2CPlayerDrawData{} }
func (m *S2CPlayerDrawData) String() string { return proto.CompactTextString(m) }
func (*S2CPlayerDrawData) ProtoMessage()    {}
func (m *S2CPlayerDrawData) GetId() uint16  { return PCK_S2CPlayerDrawData_ID }

//
type LuckDraw struct {
	//
	Value int32 `protobuf:"varint,1,opt,name=Value,proto3" json:"Value,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Have int32 `protobuf:"varint,3,opt,name=Have,proto3" json:"Have,omitempty"`
}

func (m *LuckDraw) Reset()         { *m = LuckDraw{} }
func (m *LuckDraw) String() string { return proto.CompactTextString(m) }
func (*LuckDraw) ProtoMessage()    {}

const PCK_C2SGetDrawListRate_ID = 11011 //
//
type C2SGetDrawListRate struct {
	//
	ActId int32 `protobuf:"varint,1,opt,name=ActId,proto3" json:"ActId,omitempty"`
}

func (m *C2SGetDrawListRate) Reset()         { *m = C2SGetDrawListRate{} }
func (m *C2SGetDrawListRate) String() string { return proto.CompactTextString(m) }
func (*C2SGetDrawListRate) ProtoMessage()    {}
func (m *C2SGetDrawListRate) GetId() uint16  { return PCK_C2SGetDrawListRate_ID }

const PCK_S2CGetDrawListRate_ID = 11012 //
//
type S2CGetDrawListRate struct {
	//
	ActId int32 `protobuf:"varint,1,opt,name=ActId,proto3" json:"ActId,omitempty"`
	//
	Item []string `protobuf:"bytes,2,rep,name=Item,proto3" json:"Item,omitempty"`
	//
	Rate []string `protobuf:"bytes,3,rep,name=Rate,proto3" json:"Rate,omitempty"`
}

func (m *S2CGetDrawListRate) Reset()         { *m = S2CGetDrawListRate{} }
func (m *S2CGetDrawListRate) String() string { return proto.CompactTextString(m) }
func (*S2CGetDrawListRate) ProtoMessage()    {}
func (m *S2CGetDrawListRate) GetId() uint16  { return PCK_S2CGetDrawListRate_ID }

const PCK_C2SBuyDrawItem_ID = 11001 //
//
type C2SBuyDrawItem struct {
	//
	ActId int32 `protobuf:"varint,1,opt,name=ActId,proto3" json:"ActId,omitempty"`
	//
	Type int32 `protobuf:"varint,2,opt,name=Type,proto3" json:"Type,omitempty"`
	//
	Count int32 `protobuf:"varint,3,opt,name=Count,proto3" json:"Count,omitempty"`
}

func (m *C2SBuyDrawItem) Reset()         { *m = C2SBuyDrawItem{} }
func (m *C2SBuyDrawItem) String() string { return proto.CompactTextString(m) }
func (*C2SBuyDrawItem) ProtoMessage()    {}
func (m *C2SBuyDrawItem) GetId() uint16  { return PCK_C2SBuyDrawItem_ID }

const PCK_S2CBuyDrawItem_ID = 11002 //
//
type S2CBuyDrawItem struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	ActId int32 `protobuf:"varint,2,opt,name=ActId,proto3" json:"ActId,omitempty"`
}

func (m *S2CBuyDrawItem) Reset()         { *m = S2CBuyDrawItem{} }
func (m *S2CBuyDrawItem) String() string { return proto.CompactTextString(m) }
func (*S2CBuyDrawItem) ProtoMessage()    {}
func (m *S2CBuyDrawItem) GetId() uint16  { return PCK_S2CBuyDrawItem_ID }

const PCK_C2SDraw_ID = 11003 //
//
type C2SDraw struct {
	//
	ActId int32 `protobuf:"varint,1,opt,name=ActId,proto3" json:"ActId,omitempty"`
	//
	Type int32 `protobuf:"varint,2,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (m *C2SDraw) Reset()         { *m = C2SDraw{} }
func (m *C2SDraw) String() string { return proto.CompactTextString(m) }
func (*C2SDraw) ProtoMessage()    {}
func (m *C2SDraw) GetId() uint16  { return PCK_C2SDraw_ID }

const PCK_S2CDraw_ID = 11004 //
//
type S2CDraw struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	ActId int32 `protobuf:"varint,2,opt,name=ActId,proto3" json:"ActId,omitempty"`
	//
	Prize []*ItemInfo `protobuf:"bytes,3,rep,name=Prize,proto3" json:"Prize,omitempty"`
	//
	CritLuck int32 `protobuf:"varint,4,opt,name=CritLuck,proto3" json:"CritLuck,omitempty"`
	//
	SpecialShow []int32 `protobuf:"varint,5,rep,name=SpecialShow,proto3" json:"SpecialShow,omitempty"`
}

func (m *S2CDraw) Reset()         { *m = S2CDraw{} }
func (m *S2CDraw) String() string { return proto.CompactTextString(m) }
func (*S2CDraw) ProtoMessage()    {}
func (m *S2CDraw) GetId() uint16  { return PCK_S2CDraw_ID }

//
type DrawLog struct {
	//
	Nick string `protobuf:"bytes,1,opt,name=Nick,proto3" json:"Nick,omitempty"`
	//
	DataId int32 `protobuf:"varint,2,opt,name=DataId,proto3" json:"DataId,omitempty"`
}

func (m *DrawLog) Reset()         { *m = DrawLog{} }
func (m *DrawLog) String() string { return proto.CompactTextString(m) }
func (*DrawLog) ProtoMessage()    {}

//
type DrawDbLog struct {
	//
	ActId int32 `protobuf:"varint,1,opt,name=ActId,proto3" json:"ActId,omitempty"`
	//
	Nick string `protobuf:"bytes,2,opt,name=Nick,proto3" json:"Nick,omitempty"`
	//
	DataId int32 `protobuf:"varint,3,opt,name=DataId,proto3" json:"DataId,omitempty"`
}

func (m *DrawDbLog) Reset()         { *m = DrawDbLog{} }
func (m *DrawDbLog) String() string { return proto.CompactTextString(m) }
func (*DrawDbLog) ProtoMessage()    {}

const PCK_C2SGetDrawLog_ID = 11005 //
//
type C2SGetDrawLog struct {
	//
	ActId int32 `protobuf:"varint,1,opt,name=ActId,proto3" json:"ActId,omitempty"`
}

func (m *C2SGetDrawLog) Reset()         { *m = C2SGetDrawLog{} }
func (m *C2SGetDrawLog) String() string { return proto.CompactTextString(m) }
func (*C2SGetDrawLog) ProtoMessage()    {}
func (m *C2SGetDrawLog) GetId() uint16  { return PCK_C2SGetDrawLog_ID }

const PCK_S2CGetDrawLog_ID = 11006 //
//
type S2CGetDrawLog struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Log []*DrawLog `protobuf:"bytes,2,rep,name=Log,proto3" json:"Log,omitempty"`
	//
	ActId int32 `protobuf:"varint,3,opt,name=ActId,proto3" json:"ActId,omitempty"`
}

func (m *S2CGetDrawLog) Reset()         { *m = S2CGetDrawLog{} }
func (m *S2CGetDrawLog) String() string { return proto.CompactTextString(m) }
func (*S2CGetDrawLog) ProtoMessage()    {}
func (m *S2CGetDrawLog) GetId() uint16  { return PCK_S2CGetDrawLog_ID }

const PCK_C2SGetChargeReturnData_ID = 11150 //
//
type C2SGetChargeReturnData struct {
	//
	ActId int32 `protobuf:"varint,1,opt,name=ActId,proto3" json:"ActId,omitempty"`
}

func (m *C2SGetChargeReturnData) Reset()         { *m = C2SGetChargeReturnData{} }
func (m *C2SGetChargeReturnData) String() string { return proto.CompactTextString(m) }
func (*C2SGetChargeReturnData) ProtoMessage()    {}
func (m *C2SGetChargeReturnData) GetId() uint16  { return PCK_C2SGetChargeReturnData_ID }

const PCK_S2CGetChargeReturnData_ID = 11151 //
//
type S2CGetChargeReturnData struct {
	//
	ActId int32 `protobuf:"varint,1,opt,name=ActId,proto3" json:"ActId,omitempty"`
	//
	IsReset int32 `protobuf:"varint,2,opt,name=IsReset,proto3" json:"IsReset,omitempty"`
	//
	MaxTime int32 `protobuf:"varint,3,opt,name=MaxTime,proto3" json:"MaxTime,omitempty"`
	//
	Mutil []int32 `protobuf:"varint,4,rep,name=Mutil,proto3" json:"Mutil,omitempty"`
}

func (m *S2CGetChargeReturnData) Reset()         { *m = S2CGetChargeReturnData{} }
func (m *S2CGetChargeReturnData) String() string { return proto.CompactTextString(m) }
func (*S2CGetChargeReturnData) ProtoMessage()    {}
func (m *S2CGetChargeReturnData) GetId() uint16  { return PCK_S2CGetChargeReturnData_ID }

const PCK_C2SPlayerChargeReturnData_ID = 11152 //
//
type C2SPlayerChargeReturnData struct {
	//
	ActId int32 `protobuf:"varint,1,opt,name=ActId,proto3" json:"ActId,omitempty"`
}

func (m *C2SPlayerChargeReturnData) Reset()         { *m = C2SPlayerChargeReturnData{} }
func (m *C2SPlayerChargeReturnData) String() string { return proto.CompactTextString(m) }
func (*C2SPlayerChargeReturnData) ProtoMessage()    {}
func (m *C2SPlayerChargeReturnData) GetId() uint16  { return PCK_C2SPlayerChargeReturnData_ID }

const PCK_S2CPlayerChargeReturnData_ID = 11153 //
//
type S2CPlayerChargeReturnData struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	ActId int32 `protobuf:"varint,2,opt,name=ActId,proto3" json:"ActId,omitempty"`
	//
	Times int32 `protobuf:"varint,3,opt,name=Times,proto3" json:"Times,omitempty"`
	//
	TotalCharge int32 `protobuf:"varint,4,opt,name=TotalCharge,proto3" json:"TotalCharge,omitempty"`
	//
	Need int32 `protobuf:"varint,5,opt,name=Need,proto3" json:"Need,omitempty"`
	//
	Base int32 `protobuf:"varint,6,opt,name=Base,proto3" json:"Base,omitempty"`
	//
	TodayUseTimes int32 `protobuf:"varint,7,opt,name=TodayUseTimes,proto3" json:"TodayUseTimes,omitempty"`
}

func (m *S2CPlayerChargeReturnData) Reset()         { *m = S2CPlayerChargeReturnData{} }
func (m *S2CPlayerChargeReturnData) String() string { return proto.CompactTextString(m) }
func (*S2CPlayerChargeReturnData) ProtoMessage()    {}
func (m *S2CPlayerChargeReturnData) GetId() uint16  { return PCK_S2CPlayerChargeReturnData_ID }

//
type ReturnDetail struct {
	//
	Item int32 `protobuf:"varint,1,opt,name=Item,proto3" json:"Item,omitempty"`
	//
	T int32 `protobuf:"varint,2,opt,name=T,proto3" json:"T,omitempty"`
	//
	Have int32 `protobuf:"varint,3,opt,name=Have,proto3" json:"Have,omitempty"`
}

func (m *ReturnDetail) Reset()         { *m = ReturnDetail{} }
func (m *ReturnDetail) String() string { return proto.CompactTextString(m) }
func (*ReturnDetail) ProtoMessage()    {}

const PCK_C2SChargeReturn_ID = 11154 //
//
type C2SChargeReturn struct {
	//
	ActId int32 `protobuf:"varint,1,opt,name=ActId,proto3" json:"ActId,omitempty"`
}

func (m *C2SChargeReturn) Reset()         { *m = C2SChargeReturn{} }
func (m *C2SChargeReturn) String() string { return proto.CompactTextString(m) }
func (*C2SChargeReturn) ProtoMessage()    {}
func (m *C2SChargeReturn) GetId() uint16  { return PCK_C2SChargeReturn_ID }

const PCK_S2CChargeReturn_ID = 11155 //
//
type S2CChargeReturn struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	ActId int32 `protobuf:"varint,2,opt,name=ActId,proto3" json:"ActId,omitempty"`
	//
	Multi int32 `protobuf:"varint,3,opt,name=Multi,proto3" json:"Multi,omitempty"`
}

func (m *S2CChargeReturn) Reset()         { *m = S2CChargeReturn{} }
func (m *S2CChargeReturn) String() string { return proto.CompactTextString(m) }
func (*S2CChargeReturn) ProtoMessage()    {}
func (m *S2CChargeReturn) GetId() uint16  { return PCK_S2CChargeReturn_ID }

//
type ChargeReturnPrizeCache struct {
	//
	Cached int32 `protobuf:"varint,1,opt,name=Cached,proto3" json:"Cached,omitempty"`
	//
	Multi int32 `protobuf:"varint,2,opt,name=Multi,proto3" json:"Multi,omitempty"`
	//
	Base int32 `protobuf:"varint,3,opt,name=Base,proto3" json:"Base,omitempty"`
}

func (m *ChargeReturnPrizeCache) Reset()         { *m = ChargeReturnPrizeCache{} }
func (m *ChargeReturnPrizeCache) String() string { return proto.CompactTextString(m) }
func (*ChargeReturnPrizeCache) ProtoMessage()    {}

const PCK_C2SChargeReturnReceivePrize_ID = 11158 //
//
type C2SChargeReturnReceivePrize struct {
	//
	ActId int32 `protobuf:"varint,1,opt,name=ActId,proto3" json:"ActId,omitempty"`
}

func (m *C2SChargeReturnReceivePrize) Reset()         { *m = C2SChargeReturnReceivePrize{} }
func (m *C2SChargeReturnReceivePrize) String() string { return proto.CompactTextString(m) }
func (*C2SChargeReturnReceivePrize) ProtoMessage()    {}
func (m *C2SChargeReturnReceivePrize) GetId() uint16  { return PCK_C2SChargeReturnReceivePrize_ID }

const PCK_S2CChargeReturnReceivePrize_ID = 11159 //
//
type S2CChargeReturnReceivePrize struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	ActId int32 `protobuf:"varint,2,opt,name=ActId,proto3" json:"ActId,omitempty"`
	//
	Multi int32 `protobuf:"varint,3,opt,name=Multi,proto3" json:"Multi,omitempty"`
}

func (m *S2CChargeReturnReceivePrize) Reset()         { *m = S2CChargeReturnReceivePrize{} }
func (m *S2CChargeReturnReceivePrize) String() string { return proto.CompactTextString(m) }
func (*S2CChargeReturnReceivePrize) ProtoMessage()    {}
func (m *S2CChargeReturnReceivePrize) GetId() uint16  { return PCK_S2CChargeReturnReceivePrize_ID }

//
type ChargeReturnLog struct {
	//
	Nick string `protobuf:"bytes,1,opt,name=Nick,proto3" json:"Nick,omitempty"`
	//
	Base int32 `protobuf:"varint,2,opt,name=Base,proto3" json:"Base,omitempty"`
	//
	Multi int32 `protobuf:"varint,3,opt,name=Multi,proto3" json:"Multi,omitempty"`
}

func (m *ChargeReturnLog) Reset()         { *m = ChargeReturnLog{} }
func (m *ChargeReturnLog) String() string { return proto.CompactTextString(m) }
func (*ChargeReturnLog) ProtoMessage()    {}

//
type ChargeReturnDbLog struct {
	//
	ActId int32 `protobuf:"varint,1,opt,name=ActId,proto3" json:"ActId,omitempty"`
	//
	Nick string `protobuf:"bytes,2,opt,name=Nick,proto3" json:"Nick,omitempty"`
	//
	Base int32 `protobuf:"varint,3,opt,name=Base,proto3" json:"Base,omitempty"`
	//
	Multi int32 `protobuf:"varint,4,opt,name=Multi,proto3" json:"Multi,omitempty"`
}

func (m *ChargeReturnDbLog) Reset()         { *m = ChargeReturnDbLog{} }
func (m *ChargeReturnDbLog) String() string { return proto.CompactTextString(m) }
func (*ChargeReturnDbLog) ProtoMessage()    {}

const PCK_C2SGetChargeReturnLog_ID = 11156 //
//
type C2SGetChargeReturnLog struct {
	//
	ActId int32 `protobuf:"varint,1,opt,name=ActId,proto3" json:"ActId,omitempty"`
}

func (m *C2SGetChargeReturnLog) Reset()         { *m = C2SGetChargeReturnLog{} }
func (m *C2SGetChargeReturnLog) String() string { return proto.CompactTextString(m) }
func (*C2SGetChargeReturnLog) ProtoMessage()    {}
func (m *C2SGetChargeReturnLog) GetId() uint16  { return PCK_C2SGetChargeReturnLog_ID }

const PCK_S2CGetChargeReturnLog_ID = 11157 //
//
type S2CGetChargeReturnLog struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Log []*ChargeReturnLog `protobuf:"bytes,2,rep,name=Log,proto3" json:"Log,omitempty"`
	//
	ActId int32 `protobuf:"varint,3,opt,name=ActId,proto3" json:"ActId,omitempty"`
}

func (m *S2CGetChargeReturnLog) Reset()         { *m = S2CGetChargeReturnLog{} }
func (m *S2CGetChargeReturnLog) String() string { return proto.CompactTextString(m) }
func (*S2CGetChargeReturnLog) ProtoMessage()    {}
func (m *S2CGetChargeReturnLog) GetId() uint16  { return PCK_S2CGetChargeReturnLog_ID }

const PCK_C2SGetActShopList_ID = 12001 //
//
type C2SGetActShopList struct {
	//
	Aid int32 `protobuf:"varint,1,opt,name=Aid,proto3" json:"Aid,omitempty"`
}

func (m *C2SGetActShopList) Reset()         { *m = C2SGetActShopList{} }
func (m *C2SGetActShopList) String() string { return proto.CompactTextString(m) }
func (*C2SGetActShopList) ProtoMessage()    {}
func (m *C2SGetActShopList) GetId() uint16  { return PCK_C2SGetActShopList_ID }

//
type ActShopGoods struct {
	//
	Gid int32 `protobuf:"varint,1,opt,name=Gid,proto3" json:"Gid,omitempty"`
	//
	F int32 `protobuf:"varint,2,opt,name=F,proto3" json:"F,omitempty"`
	//
	Bt int32 `protobuf:"varint,3,opt,name=Bt,proto3" json:"Bt,omitempty"`
	//
	T int32 `protobuf:"varint,4,opt,name=T,proto3" json:"T,omitempty"`
	//
	Sort int32 `protobuf:"varint,5,opt,name=Sort,proto3" json:"Sort,omitempty"`
}

func (m *ActShopGoods) Reset()         { *m = ActShopGoods{} }
func (m *ActShopGoods) String() string { return proto.CompactTextString(m) }
func (*ActShopGoods) ProtoMessage()    {}

const PCK_S2CGetActShopList_ID = 12002 //
//
type S2CGetActShopList struct {
	//
	Aid int32 `protobuf:"varint,1,opt,name=Aid,proto3" json:"Aid,omitempty"`
	//
	List []*ActShopGoods `protobuf:"bytes,2,rep,name=List,proto3" json:"List,omitempty"`
}

func (m *S2CGetActShopList) Reset()         { *m = S2CGetActShopList{} }
func (m *S2CGetActShopList) String() string { return proto.CompactTextString(m) }
func (*S2CGetActShopList) ProtoMessage()    {}
func (m *S2CGetActShopList) GetId() uint16  { return PCK_S2CGetActShopList_ID }

const PCK_C2SActShopBuy_ID = 12003 //
//
type C2SActShopBuy struct {
	//
	Gid int32 `protobuf:"varint,1,opt,name=Gid,proto3" json:"Gid,omitempty"`
}

func (m *C2SActShopBuy) Reset()         { *m = C2SActShopBuy{} }
func (m *C2SActShopBuy) String() string { return proto.CompactTextString(m) }
func (*C2SActShopBuy) ProtoMessage()    {}
func (m *C2SActShopBuy) GetId() uint16  { return PCK_C2SActShopBuy_ID }

const PCK_S2CActShopBuy_ID = 12004 //
//
type S2CActShopBuy struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Gid int32 `protobuf:"varint,2,opt,name=Gid,proto3" json:"Gid,omitempty"`
	//
	MsgParam int32 `protobuf:"varint,3,opt,name=MsgParam,proto3" json:"MsgParam,omitempty"`
	//
	Goods *ActShopGoods `protobuf:"bytes,4,opt,name=Goods,proto3" json:"Goods,omitempty"`
}

func (m *S2CActShopBuy) Reset()         { *m = S2CActShopBuy{} }
func (m *S2CActShopBuy) String() string { return proto.CompactTextString(m) }
func (*S2CActShopBuy) ProtoMessage()    {}
func (m *S2CActShopBuy) GetId() uint16  { return PCK_S2CActShopBuy_ID }

const PCK_C2SGetActGiftList_ID = 12005 //
//
type C2SGetActGiftList struct {
	//
	Aid int32 `protobuf:"varint,1,opt,name=Aid,proto3" json:"Aid,omitempty"`
}

func (m *C2SGetActGiftList) Reset()         { *m = C2SGetActGiftList{} }
func (m *C2SGetActGiftList) String() string { return proto.CompactTextString(m) }
func (*C2SGetActGiftList) ProtoMessage()    {}
func (m *C2SGetActGiftList) GetId() uint16  { return PCK_C2SGetActGiftList_ID }

//
type ActGift struct {
	//
	Gid int32 `protobuf:"varint,1,opt,name=Gid,proto3" json:"Gid,omitempty"`
	//
	Items []*ItemInfo `protobuf:"bytes,2,rep,name=Items,proto3" json:"Items,omitempty"`
	//
	CostType int32 `protobuf:"varint,3,opt,name=CostType,proto3" json:"CostType,omitempty"`
	//
	CostNum int32 `protobuf:"varint,4,opt,name=CostNum,proto3" json:"CostNum,omitempty"`
	//
	Price int32 `protobuf:"varint,5,opt,name=Price,proto3" json:"Price,omitempty"`
	//
	Price2 int32 `protobuf:"varint,6,opt,name=Price2,proto3" json:"Price2,omitempty"`
	//
	Sort int32 `protobuf:"varint,7,opt,name=Sort,proto3" json:"Sort,omitempty"`
	//
	CanBuy int32 `protobuf:"varint,8,opt,name=CanBuy,proto3" json:"CanBuy,omitempty"`
	//
	Content string `protobuf:"bytes,9,opt,name=Content,proto3" json:"Content,omitempty"`
	//
	TabName string `protobuf:"bytes,10,opt,name=TabName,proto3" json:"TabName,omitempty"`
	//
	IsRec int32 `protobuf:"varint,11,opt,name=IsRec,proto3" json:"IsRec,omitempty"`
	//
	OpenDay int32 `protobuf:"varint,12,opt,name=OpenDay,proto3" json:"OpenDay,omitempty"`
	//
	OpenDayBuy int32 `protobuf:"varint,13,opt,name=OpenDayBuy,proto3" json:"OpenDayBuy,omitempty"`
	//
	LoginDay int32 `protobuf:"varint,14,opt,name=LoginDay,proto3" json:"LoginDay,omitempty"`
	//
	LoginDayBuy int32 `protobuf:"varint,15,opt,name=LoginDayBuy,proto3" json:"LoginDayBuy,omitempty"`
	//
	UseYuan int32 `protobuf:"varint,16,opt,name=UseYuan,proto3" json:"UseYuan,omitempty"`
	//
	DevilSoulgift string `protobuf:"bytes,17,opt,name=DevilSoulgift,proto3" json:"DevilSoulgift,omitempty"`
	//
	BuyTime int32 `protobuf:"varint,18,opt,name=BuyTime,proto3" json:"BuyTime,omitempty"`
	//
	BuyTimeLimit int32 `protobuf:"varint,19,opt,name=BuyTimeLimit,proto3" json:"BuyTimeLimit,omitempty"`
}

func (m *ActGift) Reset()         { *m = ActGift{} }
func (m *ActGift) String() string { return proto.CompactTextString(m) }
func (*ActGift) ProtoMessage()    {}

const PCK_S2CGetActGiftList_ID = 12006 //
//
type S2CGetActGiftList struct {
	//
	Aid int32 `protobuf:"varint,1,opt,name=Aid,proto3" json:"Aid,omitempty"`
	//
	ActType int32 `protobuf:"varint,3,opt,name=ActType,proto3" json:"ActType,omitempty"`
	//
	List []*ActGift `protobuf:"bytes,2,rep,name=List,proto3" json:"List,omitempty"`
}

func (m *S2CGetActGiftList) Reset()         { *m = S2CGetActGiftList{} }
func (m *S2CGetActGiftList) String() string { return proto.CompactTextString(m) }
func (*S2CGetActGiftList) ProtoMessage()    {}
func (m *S2CGetActGiftList) GetId() uint16  { return PCK_S2CGetActGiftList_ID }

const PCK_C2SActGiftBuy_ID = 12007 //
//
type C2SActGiftBuy struct {
	//
	Gid int32 `protobuf:"varint,1,opt,name=Gid,proto3" json:"Gid,omitempty"`
	//
	Aid int32 `protobuf:"varint,2,opt,name=Aid,proto3" json:"Aid,omitempty"`
}

func (m *C2SActGiftBuy) Reset()         { *m = C2SActGiftBuy{} }
func (m *C2SActGiftBuy) String() string { return proto.CompactTextString(m) }
func (*C2SActGiftBuy) ProtoMessage()    {}
func (m *C2SActGiftBuy) GetId() uint16  { return PCK_C2SActGiftBuy_ID }

const PCK_S2CActGiftBuy_ID = 12008 //
//
type S2CActGiftBuy struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	P string `protobuf:"bytes,2,opt,name=P,proto3" json:"P,omitempty"`
	//
	Gid int32 `protobuf:"varint,3,opt,name=Gid,proto3" json:"Gid,omitempty"`
	//
	Close int32 `protobuf:"varint,4,opt,name=Close,proto3" json:"Close,omitempty"`
	//
	Aid int32 `protobuf:"varint,5,opt,name=Aid,proto3" json:"Aid,omitempty"`
}

func (m *S2CActGiftBuy) Reset()         { *m = S2CActGiftBuy{} }
func (m *S2CActGiftBuy) String() string { return proto.CompactTextString(m) }
func (*S2CActGiftBuy) ProtoMessage()    {}
func (m *S2CActGiftBuy) GetId() uint16  { return PCK_S2CActGiftBuy_ID }

//
type ActBossInfo struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	//
	Order int32 `protobuf:"varint,3,opt,name=Order,proto3" json:"Order,omitempty"`
	//
	BossId int32 `protobuf:"varint,4,opt,name=BossId,proto3" json:"BossId,omitempty"`
	//
	PetId int32 `protobuf:"varint,5,opt,name=PetId,proto3" json:"PetId,omitempty"`
	//
	Prize string `protobuf:"bytes,6,opt,name=Prize,proto3" json:"Prize,omitempty"`
}

func (m *ActBossInfo) Reset()         { *m = ActBossInfo{} }
func (m *ActBossInfo) String() string { return proto.CompactTextString(m) }
func (*ActBossInfo) ProtoMessage()    {}

const PCK_C2SGetActBossInfo_ID = 11101 //
//
type C2SGetActBossInfo struct {
	//
	ActId int32 `protobuf:"varint,1,opt,name=ActId,proto3" json:"ActId,omitempty"`
}

func (m *C2SGetActBossInfo) Reset()         { *m = C2SGetActBossInfo{} }
func (m *C2SGetActBossInfo) String() string { return proto.CompactTextString(m) }
func (*C2SGetActBossInfo) ProtoMessage()    {}
func (m *C2SGetActBossInfo) GetId() uint16  { return PCK_C2SGetActBossInfo_ID }

const PCK_S2CGetActBossInfo_ID = 11102 //
//
type S2CGetActBossInfo struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	ActId int32 `protobuf:"varint,2,opt,name=ActId,proto3" json:"ActId,omitempty"`
	//
	Info []*ActBossInfo `protobuf:"bytes,3,rep,name=Info,proto3" json:"Info,omitempty"`
}

func (m *S2CGetActBossInfo) Reset()         { *m = S2CGetActBossInfo{} }
func (m *S2CGetActBossInfo) String() string { return proto.CompactTextString(m) }
func (*S2CGetActBossInfo) ProtoMessage()    {}
func (m *S2CGetActBossInfo) GetId() uint16  { return PCK_S2CGetActBossInfo_ID }

//
type PlayerBossData struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	State int32 `protobuf:"varint,2,opt,name=State,proto3" json:"State,omitempty"`
	//
	Prize int32 `protobuf:"varint,3,opt,name=Prize,proto3" json:"Prize,omitempty"`
}

func (m *PlayerBossData) Reset()         { *m = PlayerBossData{} }
func (m *PlayerBossData) String() string { return proto.CompactTextString(m) }
func (*PlayerBossData) ProtoMessage()    {}

const PCK_C2SPlayerBossInfo_ID = 11107 //
//
type C2SPlayerBossInfo struct {
	//
	ActId int32 `protobuf:"varint,1,opt,name=ActId,proto3" json:"ActId,omitempty"`
}

func (m *C2SPlayerBossInfo) Reset()         { *m = C2SPlayerBossInfo{} }
func (m *C2SPlayerBossInfo) String() string { return proto.CompactTextString(m) }
func (*C2SPlayerBossInfo) ProtoMessage()    {}
func (m *C2SPlayerBossInfo) GetId() uint16  { return PCK_C2SPlayerBossInfo_ID }

const PCK_S2CPlayerBossInfo_ID = 11108 //
//
type S2CPlayerBossInfo struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	ActId int32 `protobuf:"varint,2,opt,name=ActId,proto3" json:"ActId,omitempty"`
	//
	Info []*PlayerBossData `protobuf:"bytes,3,rep,name=Info,proto3" json:"Info,omitempty"`
}

func (m *S2CPlayerBossInfo) Reset()         { *m = S2CPlayerBossInfo{} }
func (m *S2CPlayerBossInfo) String() string { return proto.CompactTextString(m) }
func (*S2CPlayerBossInfo) ProtoMessage()    {}
func (m *S2CPlayerBossInfo) GetId() uint16  { return PCK_S2CPlayerBossInfo_ID }

const PCK_C2SFightActBoss_ID = 11103 //
//
type C2SFightActBoss struct {
	//
	ActId int32 `protobuf:"varint,1,opt,name=ActId,proto3" json:"ActId,omitempty"`
	//
	BossId int32 `protobuf:"varint,2,opt,name=BossId,proto3" json:"BossId,omitempty"`
}

func (m *C2SFightActBoss) Reset()         { *m = C2SFightActBoss{} }
func (m *C2SFightActBoss) String() string { return proto.CompactTextString(m) }
func (*C2SFightActBoss) ProtoMessage()    {}
func (m *C2SFightActBoss) GetId() uint16  { return PCK_C2SFightActBoss_ID }

const PCK_S2CFightActBoss_ID = 11104 //
//
type S2CFightActBoss struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	ActId int32 `protobuf:"varint,2,opt,name=ActId,proto3" json:"ActId,omitempty"`
	//
	BossId int32 `protobuf:"varint,3,opt,name=BossId,proto3" json:"BossId,omitempty"`
	//
	Prize []*ItemInfo `protobuf:"bytes,4,rep,name=Prize,proto3" json:"Prize,omitempty"`
}

func (m *S2CFightActBoss) Reset()         { *m = S2CFightActBoss{} }
func (m *S2CFightActBoss) String() string { return proto.CompactTextString(m) }
func (*S2CFightActBoss) ProtoMessage()    {}
func (m *S2CFightActBoss) GetId() uint16  { return PCK_S2CFightActBoss_ID }

const PCK_C2SGetActBossPrize_ID = 11105 //
//
type C2SGetActBossPrize struct {
	//
	ActId int32 `protobuf:"varint,1,opt,name=ActId,proto3" json:"ActId,omitempty"`
	//
	BossId int32 `protobuf:"varint,2,opt,name=BossId,proto3" json:"BossId,omitempty"`
}

func (m *C2SGetActBossPrize) Reset()         { *m = C2SGetActBossPrize{} }
func (m *C2SGetActBossPrize) String() string { return proto.CompactTextString(m) }
func (*C2SGetActBossPrize) ProtoMessage()    {}
func (m *C2SGetActBossPrize) GetId() uint16  { return PCK_C2SGetActBossPrize_ID }

const PCK_S2CGetActBossPrize_ID = 11106 //
//
type S2CGetActBossPrize struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	ActId int32 `protobuf:"varint,2,opt,name=ActId,proto3" json:"ActId,omitempty"`
	//
	BossId int32 `protobuf:"varint,3,opt,name=BossId,proto3" json:"BossId,omitempty"`
	//
	Prize []*ItemInfo `protobuf:"bytes,4,rep,name=Prize,proto3" json:"Prize,omitempty"`
}

func (m *S2CGetActBossPrize) Reset()         { *m = S2CGetActBossPrize{} }
func (m *S2CGetActBossPrize) String() string { return proto.CompactTextString(m) }
func (*S2CGetActBossPrize) ProtoMessage()    {}
func (m *S2CGetActBossPrize) GetId() uint16  { return PCK_S2CGetActBossPrize_ID }

//
type GiftLimitInfo struct {
	//
	GoodId int32 `protobuf:"varint,1,opt,name=GoodId,proto3" json:"GoodId,omitempty"`
	//
	Items []*ItemInfo `protobuf:"bytes,2,rep,name=Items,proto3" json:"Items,omitempty"`
	//
	Price int32 `protobuf:"varint,3,opt,name=Price,proto3" json:"Price,omitempty"`
}

func (m *GiftLimitInfo) Reset()         { *m = GiftLimitInfo{} }
func (m *GiftLimitInfo) String() string { return proto.CompactTextString(m) }
func (*GiftLimitInfo) ProtoMessage()    {}

const PCK_C2SGetActGiftLimit_ID = 12801 //
//
type C2SGetActGiftLimit struct {
	//
	AId int32 `protobuf:"varint,1,opt,name=AId,proto3" json:"AId,omitempty"`
}

func (m *C2SGetActGiftLimit) Reset()         { *m = C2SGetActGiftLimit{} }
func (m *C2SGetActGiftLimit) String() string { return proto.CompactTextString(m) }
func (*C2SGetActGiftLimit) ProtoMessage()    {}
func (m *C2SGetActGiftLimit) GetId() uint16  { return PCK_C2SGetActGiftLimit_ID }

const PCK_S2CGetActGiftLimit_ID = 12802 //
//
type S2CGetActGiftLimit struct {
	//
	AId int32 `protobuf:"varint,1,opt,name=AId,proto3" json:"AId,omitempty"`
	//
	List []*GiftLimitInfo `protobuf:"bytes,2,rep,name=List,proto3" json:"List,omitempty"`
	//
	ActType int32 `protobuf:"varint,3,opt,name=ActType,proto3" json:"ActType,omitempty"`
	//
	ResId int32 `protobuf:"varint,5,opt,name=ResId,proto3" json:"ResId,omitempty"`
}

func (m *S2CGetActGiftLimit) Reset()         { *m = S2CGetActGiftLimit{} }
func (m *S2CGetActGiftLimit) String() string { return proto.CompactTextString(m) }
func (*S2CGetActGiftLimit) ProtoMessage()    {}
func (m *S2CGetActGiftLimit) GetId() uint16  { return PCK_S2CGetActGiftLimit_ID }

//
type GiftOnlineInfo struct {
	//
	GoodId int32 `protobuf:"varint,1,opt,name=GoodId,proto3" json:"GoodId,omitempty"`
	//
	Items []*ItemInfo `protobuf:"bytes,2,rep,name=Items,proto3" json:"Items,omitempty"`
	//
	State int32 `protobuf:"varint,3,opt,name=State,proto3" json:"State,omitempty"`
	//
	GetTime int64 `protobuf:"varint,4,opt,name=GetTime,proto3" json:"GetTime,omitempty"`
	//
	OnlineTime int32 `protobuf:"varint,5,opt,name=OnlineTime,proto3" json:"OnlineTime,omitempty"`
}

func (m *GiftOnlineInfo) Reset()         { *m = GiftOnlineInfo{} }
func (m *GiftOnlineInfo) String() string { return proto.CompactTextString(m) }
func (*GiftOnlineInfo) ProtoMessage()    {}

const PCK_C2SGetActGiftOnline_ID = 12601 //
//
type C2SGetActGiftOnline struct {
	//
	AId int32 `protobuf:"varint,1,opt,name=AId,proto3" json:"AId,omitempty"`
}

func (m *C2SGetActGiftOnline) Reset()         { *m = C2SGetActGiftOnline{} }
func (m *C2SGetActGiftOnline) String() string { return proto.CompactTextString(m) }
func (*C2SGetActGiftOnline) ProtoMessage()    {}
func (m *C2SGetActGiftOnline) GetId() uint16  { return PCK_C2SGetActGiftOnline_ID }

const PCK_S2CGetActGiftOnline_ID = 12602 //
//
type S2CGetActGiftOnline struct {
	//
	AId int32 `protobuf:"varint,1,opt,name=AId,proto3" json:"AId,omitempty"`
	//
	List []*GiftOnlineInfo `protobuf:"bytes,2,rep,name=List,proto3" json:"List,omitempty"`
	//
	Online int64 `protobuf:"varint,3,opt,name=Online,proto3" json:"Online,omitempty"`
}

func (m *S2CGetActGiftOnline) Reset()         { *m = S2CGetActGiftOnline{} }
func (m *S2CGetActGiftOnline) String() string { return proto.CompactTextString(m) }
func (*S2CGetActGiftOnline) ProtoMessage()    {}
func (m *S2CGetActGiftOnline) GetId() uint16  { return PCK_S2CGetActGiftOnline_ID }

const PCK_C2SActGetPrize_ID = 12009 //
//
type C2SActGetPrize struct {
	//
	AId int32 `protobuf:"varint,1,opt,name=AId,proto3" json:"AId,omitempty"`
	//
	Param int32 `protobuf:"varint,2,opt,name=Param,proto3" json:"Param,omitempty"`
}

func (m *C2SActGetPrize) Reset()         { *m = C2SActGetPrize{} }
func (m *C2SActGetPrize) String() string { return proto.CompactTextString(m) }
func (*C2SActGetPrize) ProtoMessage()    {}
func (m *C2SActGetPrize) GetId() uint16  { return PCK_C2SActGetPrize_ID }

const PCK_S2CActGetPrize_ID = 12010 //
//
type S2CActGetPrize struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	AId int32 `protobuf:"varint,2,opt,name=AId,proto3" json:"AId,omitempty"`
	//
	Param int32 `protobuf:"varint,3,opt,name=Param,proto3" json:"Param,omitempty"`
	//
	Param2 int64 `protobuf:"varint,4,opt,name=Param2,proto3" json:"Param2,omitempty"`
}

func (m *S2CActGetPrize) Reset()         { *m = S2CActGetPrize{} }
func (m *S2CActGetPrize) String() string { return proto.CompactTextString(m) }
func (*S2CActGetPrize) ProtoMessage()    {}
func (m *S2CActGetPrize) GetId() uint16  { return PCK_S2CActGetPrize_ID }

//
type GiftMiracleInfo struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Items []*ItemInfo `protobuf:"bytes,2,rep,name=Items,proto3" json:"Items,omitempty"`
	//
	BuyMul int32 `protobuf:"varint,4,opt,name=BuyMul,proto3" json:"BuyMul,omitempty"`
	//
	State int32 `protobuf:"varint,5,opt,name=State,proto3" json:"State,omitempty"`
	//
	CostItemId int32 `protobuf:"varint,6,opt,name=CostItemId,proto3" json:"CostItemId,omitempty"`
	//
	CostItemNum int64 `protobuf:"varint,7,opt,name=CostItemNum,proto3" json:"CostItemNum,omitempty"`
}

func (m *GiftMiracleInfo) Reset()         { *m = GiftMiracleInfo{} }
func (m *GiftMiracleInfo) String() string { return proto.CompactTextString(m) }
func (*GiftMiracleInfo) ProtoMessage()    {}

const PCK_C2SGetActGiftMiracle_ID = 12701 //
//
type C2SGetActGiftMiracle struct {
	//
	AId int32 `protobuf:"varint,1,opt,name=AId,proto3" json:"AId,omitempty"`
}

func (m *C2SGetActGiftMiracle) Reset()         { *m = C2SGetActGiftMiracle{} }
func (m *C2SGetActGiftMiracle) String() string { return proto.CompactTextString(m) }
func (*C2SGetActGiftMiracle) ProtoMessage()    {}
func (m *C2SGetActGiftMiracle) GetId() uint16  { return PCK_C2SGetActGiftMiracle_ID }

const PCK_S2CGetActGiftMiracle_ID = 12702 //
//
type S2CGetActGiftMiracle struct {
	//
	AId int32 `protobuf:"varint,1,opt,name=AId,proto3" json:"AId,omitempty"`
	//
	List []*GiftMiracleInfo `protobuf:"bytes,2,rep,name=List,proto3" json:"List,omitempty"`
	//
	T int32 `protobuf:"varint,3,opt,name=T,proto3" json:"T,omitempty"`
}

func (m *S2CGetActGiftMiracle) Reset()         { *m = S2CGetActGiftMiracle{} }
func (m *S2CGetActGiftMiracle) String() string { return proto.CompactTextString(m) }
func (*S2CGetActGiftMiracle) ProtoMessage()    {}
func (m *S2CGetActGiftMiracle) GetId() uint16  { return PCK_S2CGetActGiftMiracle_ID }

//
type GiftXRmbInfo struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Items []*ItemInfo `protobuf:"bytes,2,rep,name=Items,proto3" json:"Items,omitempty"`
	//
	TableName string `protobuf:"bytes,4,opt,name=TableName,proto3" json:"TableName,omitempty"`
	//
	Icon int32 `protobuf:"varint,5,opt,name=Icon,proto3" json:"Icon,omitempty"`
	//
	Bg int32 `protobuf:"varint,6,opt,name=Bg,proto3" json:"Bg,omitempty"`
	//
	AdId int32 `protobuf:"varint,7,opt,name=AdId,proto3" json:"AdId,omitempty"`
	//
	AnimId int32 `protobuf:"varint,8,opt,name=AnimId,proto3" json:"AnimId,omitempty"`
	//
	State int32 `protobuf:"varint,9,opt,name=State,proto3" json:"State,omitempty"`
	//
	AnimType string `protobuf:"bytes,10,opt,name=AnimType,proto3" json:"AnimType,omitempty"`
	//
	CostItem []*ItemInfo `protobuf:"bytes,11,rep,name=CostItem,proto3" json:"CostItem,omitempty"`
}

func (m *GiftXRmbInfo) Reset()         { *m = GiftXRmbInfo{} }
func (m *GiftXRmbInfo) String() string { return proto.CompactTextString(m) }
func (*GiftXRmbInfo) ProtoMessage()    {}

const PCK_C2SGetActGiftXRmbInfo_ID = 12901 //
//
type C2SGetActGiftXRmbInfo struct {
	//
	AId int32 `protobuf:"varint,1,opt,name=AId,proto3" json:"AId,omitempty"`
}

func (m *C2SGetActGiftXRmbInfo) Reset()         { *m = C2SGetActGiftXRmbInfo{} }
func (m *C2SGetActGiftXRmbInfo) String() string { return proto.CompactTextString(m) }
func (*C2SGetActGiftXRmbInfo) ProtoMessage()    {}
func (m *C2SGetActGiftXRmbInfo) GetId() uint16  { return PCK_C2SGetActGiftXRmbInfo_ID }

const PCK_S2CGetActGiftXRmbInfo_ID = 12902 //
//
type S2CGetActGiftXRmbInfo struct {
	//
	AId int32 `protobuf:"varint,1,opt,name=AId,proto3" json:"AId,omitempty"`
	//
	List []*GiftXRmbInfo `protobuf:"bytes,2,rep,name=List,proto3" json:"List,omitempty"`
}

func (m *S2CGetActGiftXRmbInfo) Reset()         { *m = S2CGetActGiftXRmbInfo{} }
func (m *S2CGetActGiftXRmbInfo) String() string { return proto.CompactTextString(m) }
func (*S2CGetActGiftXRmbInfo) ProtoMessage()    {}
func (m *S2CGetActGiftXRmbInfo) GetId() uint16  { return PCK_S2CGetActGiftXRmbInfo_ID }

//
type CollectFontInfo struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	SourceItems []*ItemInfo `protobuf:"bytes,2,rep,name=SourceItems,proto3" json:"SourceItems,omitempty"`
	//
	TargetItems []*ItemInfo `protobuf:"bytes,3,rep,name=TargetItems,proto3" json:"TargetItems,omitempty"`
	//
	CurrTimes int32 `protobuf:"varint,4,opt,name=CurrTimes,proto3" json:"CurrTimes,omitempty"`
	//
	MaxTimes int32 `protobuf:"varint,5,opt,name=MaxTimes,proto3" json:"MaxTimes,omitempty"`
}

func (m *CollectFontInfo) Reset()         { *m = CollectFontInfo{} }
func (m *CollectFontInfo) String() string { return proto.CompactTextString(m) }
func (*CollectFontInfo) ProtoMessage()    {}

const PCK_C2SGetActCollectFontInfo_ID = 12911 //
//
type C2SGetActCollectFontInfo struct {
	//
	AId int32 `protobuf:"varint,1,opt,name=AId,proto3" json:"AId,omitempty"`
}

func (m *C2SGetActCollectFontInfo) Reset()         { *m = C2SGetActCollectFontInfo{} }
func (m *C2SGetActCollectFontInfo) String() string { return proto.CompactTextString(m) }
func (*C2SGetActCollectFontInfo) ProtoMessage()    {}
func (m *C2SGetActCollectFontInfo) GetId() uint16  { return PCK_C2SGetActCollectFontInfo_ID }

const PCK_S2CGetActCollectFontInfo_ID = 12912 //
//
type S2CGetActCollectFontInfo struct {
	//
	AId int32 `protobuf:"varint,1,opt,name=AId,proto3" json:"AId,omitempty"`
	//
	List []*CollectFontInfo `protobuf:"bytes,2,rep,name=List,proto3" json:"List,omitempty"`
}

func (m *S2CGetActCollectFontInfo) Reset()         { *m = S2CGetActCollectFontInfo{} }
func (m *S2CGetActCollectFontInfo) String() string { return proto.CompactTextString(m) }
func (*S2CGetActCollectFontInfo) ProtoMessage()    {}
func (m *S2CGetActCollectFontInfo) GetId() uint16  { return PCK_S2CGetActCollectFontInfo_ID }

const PCK_C2SGetActLifeVipGiveInfo_ID = 12921 //
//
type C2SGetActLifeVipGiveInfo struct {
	//
	AId int32 `protobuf:"varint,1,opt,name=AId,proto3" json:"AId,omitempty"`
}

func (m *C2SGetActLifeVipGiveInfo) Reset()         { *m = C2SGetActLifeVipGiveInfo{} }
func (m *C2SGetActLifeVipGiveInfo) String() string { return proto.CompactTextString(m) }
func (*C2SGetActLifeVipGiveInfo) ProtoMessage()    {}
func (m *C2SGetActLifeVipGiveInfo) GetId() uint16  { return PCK_C2SGetActLifeVipGiveInfo_ID }

const PCK_S2CGetActLifeVipGiveInfo_ID = 12922 //
//
type S2CGetActLifeVipGiveInfo struct {
	//
	AId int32 `protobuf:"varint,1,opt,name=AId,proto3" json:"AId,omitempty"`
	//
	State int32 `protobuf:"varint,2,opt,name=State,proto3" json:"State,omitempty"`
	//
	Buy int32 `protobuf:"varint,3,opt,name=Buy,proto3" json:"Buy,omitempty"`
	//
	Give int32 `protobuf:"varint,4,opt,name=Give,proto3" json:"Give,omitempty"`
	//
	Price int32 `protobuf:"varint,5,opt,name=Price,proto3" json:"Price,omitempty"`
}

func (m *S2CGetActLifeVipGiveInfo) Reset()         { *m = S2CGetActLifeVipGiveInfo{} }
func (m *S2CGetActLifeVipGiveInfo) String() string { return proto.CompactTextString(m) }
func (*S2CGetActLifeVipGiveInfo) ProtoMessage()    {}
func (m *S2CGetActLifeVipGiveInfo) GetId() uint16  { return PCK_S2CGetActLifeVipGiveInfo_ID }

//
type Gift1Info struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Items []*ItemInfo `protobuf:"bytes,2,rep,name=Items,proto3" json:"Items,omitempty"`
	//
	TableName string `protobuf:"bytes,3,opt,name=TableName,proto3" json:"TableName,omitempty"`
	//
	Bg int32 `protobuf:"varint,4,opt,name=Bg,proto3" json:"Bg,omitempty"`
	//
	AdId int32 `protobuf:"varint,5,opt,name=AdId,proto3" json:"AdId,omitempty"`
	//
	ImgId int32 `protobuf:"varint,6,opt,name=ImgId,proto3" json:"ImgId,omitempty"`
	//
	DiscountImg int32 `protobuf:"varint,7,opt,name=DiscountImg,proto3" json:"DiscountImg,omitempty"`
	//
	State int32 `protobuf:"varint,8,opt,name=State,proto3" json:"State,omitempty"`
}

func (m *Gift1Info) Reset()         { *m = Gift1Info{} }
func (m *Gift1Info) String() string { return proto.CompactTextString(m) }
func (*Gift1Info) ProtoMessage()    {}

const PCK_C2SGetActGift1Info_ID = 12923 //
//
type C2SGetActGift1Info struct {
	//
	AId int32 `protobuf:"varint,1,opt,name=AId,proto3" json:"AId,omitempty"`
}

func (m *C2SGetActGift1Info) Reset()         { *m = C2SGetActGift1Info{} }
func (m *C2SGetActGift1Info) String() string { return proto.CompactTextString(m) }
func (*C2SGetActGift1Info) ProtoMessage()    {}
func (m *C2SGetActGift1Info) GetId() uint16  { return PCK_C2SGetActGift1Info_ID }

const PCK_S2CGetActGift1Info_ID = 12924 //
//
type S2CGetActGift1Info struct {
	//
	AId int32 `protobuf:"varint,1,opt,name=AId,proto3" json:"AId,omitempty"`
	//
	List *Gift1Info `protobuf:"bytes,2,opt,name=List,proto3" json:"List,omitempty"`
}

func (m *S2CGetActGift1Info) Reset()         { *m = S2CGetActGift1Info{} }
func (m *S2CGetActGift1Info) String() string { return proto.CompactTextString(m) }
func (*S2CGetActGift1Info) ProtoMessage()    {}
func (m *S2CGetActGift1Info) GetId() uint16  { return PCK_S2CGetActGift1Info_ID }

//
type Gift2Info struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Items []*ItemInfo `protobuf:"bytes,2,rep,name=Items,proto3" json:"Items,omitempty"`
	//
	TableName string `protobuf:"bytes,3,opt,name=TableName,proto3" json:"TableName,omitempty"`
	//
	Bg int32 `protobuf:"varint,4,opt,name=Bg,proto3" json:"Bg,omitempty"`
	//
	AdId int32 `protobuf:"varint,5,opt,name=AdId,proto3" json:"AdId,omitempty"`
	//
	ImgId int32 `protobuf:"varint,6,opt,name=ImgId,proto3" json:"ImgId,omitempty"`
	//
	State int32 `protobuf:"varint,7,opt,name=State,proto3" json:"State,omitempty"`
}

func (m *Gift2Info) Reset()         { *m = Gift2Info{} }
func (m *Gift2Info) String() string { return proto.CompactTextString(m) }
func (*Gift2Info) ProtoMessage()    {}

const PCK_C2SGetActGift2Info_ID = 12925 //
//
type C2SGetActGift2Info struct {
	//
	AId int32 `protobuf:"varint,1,opt,name=AId,proto3" json:"AId,omitempty"`
}

func (m *C2SGetActGift2Info) Reset()         { *m = C2SGetActGift2Info{} }
func (m *C2SGetActGift2Info) String() string { return proto.CompactTextString(m) }
func (*C2SGetActGift2Info) ProtoMessage()    {}
func (m *C2SGetActGift2Info) GetId() uint16  { return PCK_C2SGetActGift2Info_ID }

const PCK_S2CGetActGift2Info_ID = 12926 //
//
type S2CGetActGift2Info struct {
	//
	AId int32 `protobuf:"varint,1,opt,name=AId,proto3" json:"AId,omitempty"`
	//
	List []*Gift2Info `protobuf:"bytes,2,rep,name=List,proto3" json:"List,omitempty"`
}

func (m *S2CGetActGift2Info) Reset()         { *m = S2CGetActGift2Info{} }
func (m *S2CGetActGift2Info) String() string { return proto.CompactTextString(m) }
func (*S2CGetActGift2Info) ProtoMessage()    {}
func (m *S2CGetActGift2Info) GetId() uint16  { return PCK_S2CGetActGift2Info_ID }

//
type CumulativeInfo struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	TabName string `protobuf:"bytes,2,opt,name=TabName,proto3" json:"TabName,omitempty"`
	//
	ItemName string `protobuf:"bytes,3,opt,name=ItemName,proto3" json:"ItemName,omitempty"`
	//
	AdId int32 `protobuf:"varint,4,opt,name=AdId,proto3" json:"AdId,omitempty"`
	//
	Charge int32 `protobuf:"varint,6,opt,name=Charge,proto3" json:"Charge,omitempty"`
	//
	FightValue int64 `protobuf:"varint,7,opt,name=FightValue,proto3" json:"FightValue,omitempty"`
	//
	AnimId int32 `protobuf:"varint,8,opt,name=AnimId,proto3" json:"AnimId,omitempty"`
	//
	AnimType string `protobuf:"bytes,9,opt,name=AnimType,proto3" json:"AnimType,omitempty"`
	//
	Items []*ItemInfo `protobuf:"bytes,10,rep,name=Items,proto3" json:"Items,omitempty"`
	//
	State int32 `protobuf:"varint,12,opt,name=State,proto3" json:"State,omitempty"`
}

func (m *CumulativeInfo) Reset()         { *m = CumulativeInfo{} }
func (m *CumulativeInfo) String() string { return proto.CompactTextString(m) }
func (*CumulativeInfo) ProtoMessage()    {}

const PCK_C2SGetActCumulativeInfo_ID = 12927 //
//
type C2SGetActCumulativeInfo struct {
	//
	AId int32 `protobuf:"varint,1,opt,name=AId,proto3" json:"AId,omitempty"`
}

func (m *C2SGetActCumulativeInfo) Reset()         { *m = C2SGetActCumulativeInfo{} }
func (m *C2SGetActCumulativeInfo) String() string { return proto.CompactTextString(m) }
func (*C2SGetActCumulativeInfo) ProtoMessage()    {}
func (m *C2SGetActCumulativeInfo) GetId() uint16  { return PCK_C2SGetActCumulativeInfo_ID }

const PCK_S2CGetActCumulativeInfo_ID = 12928 //
//
type S2CGetActCumulativeInfo struct {
	//
	AId int32 `protobuf:"varint,1,opt,name=AId,proto3" json:"AId,omitempty"`
	//
	List []*CumulativeInfo `protobuf:"bytes,2,rep,name=List,proto3" json:"List,omitempty"`
	//
	FinalReward []*ItemInfo `protobuf:"bytes,3,rep,name=FinalReward,proto3" json:"FinalReward,omitempty"`
	//
	BannerId int32 `protobuf:"varint,4,opt,name=BannerId,proto3" json:"BannerId,omitempty"`
	//
	ActRmb int32 `protobuf:"varint,5,opt,name=ActRmb,proto3" json:"ActRmb,omitempty"`
}

func (m *S2CGetActCumulativeInfo) Reset()         { *m = S2CGetActCumulativeInfo{} }
func (m *S2CGetActCumulativeInfo) String() string { return proto.CompactTextString(m) }
func (*S2CGetActCumulativeInfo) ProtoMessage()    {}
func (m *S2CGetActCumulativeInfo) GetId() uint16  { return PCK_S2CGetActCumulativeInfo_ID }

//
type ChargeDayInfo struct {
	//
	Day int32 `protobuf:"varint,1,opt,name=Day,proto3" json:"Day,omitempty"`
	//
	Items []*ItemInfo `protobuf:"bytes,2,rep,name=Items,proto3" json:"Items,omitempty"`
	//
	State int32 `protobuf:"varint,3,opt,name=State,proto3" json:"State,omitempty"`
	//
	Desc string `protobuf:"bytes,4,opt,name=Desc,proto3" json:"Desc,omitempty"`
}

func (m *ChargeDayInfo) Reset()         { *m = ChargeDayInfo{} }
func (m *ChargeDayInfo) String() string { return proto.CompactTextString(m) }
func (*ChargeDayInfo) ProtoMessage()    {}

const PCK_C2SGetActChargeDayInfo_ID = 12929 //
//
type C2SGetActChargeDayInfo struct {
	//
	AId int32 `protobuf:"varint,1,opt,name=AId,proto3" json:"AId,omitempty"`
}

func (m *C2SGetActChargeDayInfo) Reset()         { *m = C2SGetActChargeDayInfo{} }
func (m *C2SGetActChargeDayInfo) String() string { return proto.CompactTextString(m) }
func (*C2SGetActChargeDayInfo) ProtoMessage()    {}
func (m *C2SGetActChargeDayInfo) GetId() uint16  { return PCK_C2SGetActChargeDayInfo_ID }

const PCK_S2CGetActChargeDayInfo_ID = 12930 //
//
type S2CGetActChargeDayInfo struct {
	//
	AId int32 `protobuf:"varint,1,opt,name=AId,proto3" json:"AId,omitempty"`
	//
	List []*ChargeDayInfo `protobuf:"bytes,2,rep,name=List,proto3" json:"List,omitempty"`
	//
	ActChargeDay int32 `protobuf:"varint,3,opt,name=ActChargeDay,proto3" json:"ActChargeDay,omitempty"`
}

func (m *S2CGetActChargeDayInfo) Reset()         { *m = S2CGetActChargeDayInfo{} }
func (m *S2CGetActChargeDayInfo) String() string { return proto.CompactTextString(m) }
func (*S2CGetActChargeDayInfo) ProtoMessage()    {}
func (m *S2CGetActChargeDayInfo) GetId() uint16  { return PCK_S2CGetActChargeDayInfo_ID }

//
type InvestmentInfo struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Items []*ItemInfo `protobuf:"bytes,2,rep,name=Items,proto3" json:"Items,omitempty"`
	//
	ItemDesc string `protobuf:"bytes,3,opt,name=ItemDesc,proto3" json:"ItemDesc,omitempty"`
	//
	AttrType int32 `protobuf:"varint,4,opt,name=AttrType,proto3" json:"AttrType,omitempty"`
	//
	AttrValue int64 `protobuf:"varint,5,opt,name=AttrValue,proto3" json:"AttrValue,omitempty"`
	//
	State int32 `protobuf:"varint,6,opt,name=State,proto3" json:"State,omitempty"`
}

func (m *InvestmentInfo) Reset()         { *m = InvestmentInfo{} }
func (m *InvestmentInfo) String() string { return proto.CompactTextString(m) }
func (*InvestmentInfo) ProtoMessage()    {}

const PCK_C2SGetInvestmentInfo_ID = 12933 //
//
type C2SGetInvestmentInfo struct {
	//
	AId int32 `protobuf:"varint,1,opt,name=AId,proto3" json:"AId,omitempty"`
}

func (m *C2SGetInvestmentInfo) Reset()         { *m = C2SGetInvestmentInfo{} }
func (m *C2SGetInvestmentInfo) String() string { return proto.CompactTextString(m) }
func (*C2SGetInvestmentInfo) ProtoMessage()    {}
func (m *C2SGetInvestmentInfo) GetId() uint16  { return PCK_C2SGetInvestmentInfo_ID }

const PCK_S2CGetInvestmentInfo_ID = 12934 //
//
type S2CGetInvestmentInfo struct {
	//
	AId int32 `protobuf:"varint,1,opt,name=AId,proto3" json:"AId,omitempty"`
	//
	List []*InvestmentInfo `protobuf:"bytes,2,rep,name=List,proto3" json:"List,omitempty"`
	//
	IsBuy int32 `protobuf:"varint,3,opt,name=IsBuy,proto3" json:"IsBuy,omitempty"`
	//
	GoodId int32 `protobuf:"varint,4,opt,name=GoodId,proto3" json:"GoodId,omitempty"`
}

func (m *S2CGetInvestmentInfo) Reset()         { *m = S2CGetInvestmentInfo{} }
func (m *S2CGetInvestmentInfo) String() string { return proto.CompactTextString(m) }
func (*S2CGetInvestmentInfo) ProtoMessage()    {}
func (m *S2CGetInvestmentInfo) GetId() uint16  { return PCK_S2CGetInvestmentInfo_ID }

const PCK_C2SCumulativeFirstInfo_ID = 12931 //
//
type C2SCumulativeFirstInfo struct {
	//
	AId int32 `protobuf:"varint,1,opt,name=AId,proto3" json:"AId,omitempty"`
}

func (m *C2SCumulativeFirstInfo) Reset()         { *m = C2SCumulativeFirstInfo{} }
func (m *C2SCumulativeFirstInfo) String() string { return proto.CompactTextString(m) }
func (*C2SCumulativeFirstInfo) ProtoMessage()    {}
func (m *C2SCumulativeFirstInfo) GetId() uint16  { return PCK_C2SCumulativeFirstInfo_ID }

const PCK_S2CCumulativeFirstInfo_ID = 12932 //
//
type S2CCumulativeFirstInfo struct {
	//
	AId int32 `protobuf:"varint,1,opt,name=AId,proto3" json:"AId,omitempty"`
	//
	Items []*ItemInfo `protobuf:"bytes,2,rep,name=Items,proto3" json:"Items,omitempty"`
	//
	Rmb int32 `protobuf:"varint,3,opt,name=Rmb,proto3" json:"Rmb,omitempty"`
	//
	AdId int32 `protobuf:"varint,4,opt,name=AdId,proto3" json:"AdId,omitempty"`
	//
	FightValue int64 `protobuf:"varint,5,opt,name=FightValue,proto3" json:"FightValue,omitempty"`
	//
	AnimId int32 `protobuf:"varint,6,opt,name=AnimId,proto3" json:"AnimId,omitempty"`
	//
	AnimType string `protobuf:"bytes,7,opt,name=AnimType,proto3" json:"AnimType,omitempty"`
}

func (m *S2CCumulativeFirstInfo) Reset()         { *m = S2CCumulativeFirstInfo{} }
func (m *S2CCumulativeFirstInfo) String() string { return proto.CompactTextString(m) }
func (*S2CCumulativeFirstInfo) ProtoMessage()    {}
func (m *S2CCumulativeFirstInfo) GetId() uint16  { return PCK_S2CCumulativeFirstInfo_ID }

//
type ActDareBossInfo struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	//
	Order int32 `protobuf:"varint,3,opt,name=Order,proto3" json:"Order,omitempty"`
	//
	BossId int32 `protobuf:"varint,4,opt,name=BossId,proto3" json:"BossId,omitempty"`
	//
	ObjType int32 `protobuf:"varint,5,opt,name=ObjType,proto3" json:"ObjType,omitempty"`
	//
	TypeId int32 `protobuf:"varint,6,opt,name=TypeId,proto3" json:"TypeId,omitempty"`
	//
	Prize string `protobuf:"bytes,7,opt,name=Prize,proto3" json:"Prize,omitempty"`
}

func (m *ActDareBossInfo) Reset()         { *m = ActDareBossInfo{} }
func (m *ActDareBossInfo) String() string { return proto.CompactTextString(m) }
func (*ActDareBossInfo) ProtoMessage()    {}

const PCK_C2SGetActDareBossInfo_ID = 11201 //
//
type C2SGetActDareBossInfo struct {
	//
	ActId int32 `protobuf:"varint,1,opt,name=ActId,proto3" json:"ActId,omitempty"`
}

func (m *C2SGetActDareBossInfo) Reset()         { *m = C2SGetActDareBossInfo{} }
func (m *C2SGetActDareBossInfo) String() string { return proto.CompactTextString(m) }
func (*C2SGetActDareBossInfo) ProtoMessage()    {}
func (m *C2SGetActDareBossInfo) GetId() uint16  { return PCK_C2SGetActDareBossInfo_ID }

const PCK_S2CGetActDareBossInfo_ID = 11202 //
//
type S2CGetActDareBossInfo struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	ActId int32 `protobuf:"varint,2,opt,name=ActId,proto3" json:"ActId,omitempty"`
	//
	Info []*ActDareBossInfo `protobuf:"bytes,3,rep,name=Info,proto3" json:"Info,omitempty"`
}

func (m *S2CGetActDareBossInfo) Reset()         { *m = S2CGetActDareBossInfo{} }
func (m *S2CGetActDareBossInfo) String() string { return proto.CompactTextString(m) }
func (*S2CGetActDareBossInfo) ProtoMessage()    {}
func (m *S2CGetActDareBossInfo) GetId() uint16  { return PCK_S2CGetActDareBossInfo_ID }

//
type PlayerDareBossData struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	State int32 `protobuf:"varint,2,opt,name=State,proto3" json:"State,omitempty"`
	//
	Prize int32 `protobuf:"varint,3,opt,name=Prize,proto3" json:"Prize,omitempty"`
}

func (m *PlayerDareBossData) Reset()         { *m = PlayerDareBossData{} }
func (m *PlayerDareBossData) String() string { return proto.CompactTextString(m) }
func (*PlayerDareBossData) ProtoMessage()    {}

const PCK_C2SPlayerDareBossInfo_ID = 11207 //
//
type C2SPlayerDareBossInfo struct {
	//
	ActId int32 `protobuf:"varint,1,opt,name=ActId,proto3" json:"ActId,omitempty"`
}

func (m *C2SPlayerDareBossInfo) Reset()         { *m = C2SPlayerDareBossInfo{} }
func (m *C2SPlayerDareBossInfo) String() string { return proto.CompactTextString(m) }
func (*C2SPlayerDareBossInfo) ProtoMessage()    {}
func (m *C2SPlayerDareBossInfo) GetId() uint16  { return PCK_C2SPlayerDareBossInfo_ID }

const PCK_S2CPlayerDareBossInfo_ID = 11208 //
//
type S2CPlayerDareBossInfo struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	ActId int32 `protobuf:"varint,2,opt,name=ActId,proto3" json:"ActId,omitempty"`
	//
	State int32 `protobuf:"varint,3,opt,name=State,proto3" json:"State,omitempty"`
	//
	DeadLine int32 `protobuf:"varint,4,opt,name=DeadLine,proto3" json:"DeadLine,omitempty"`
	//
	Info []*PlayerDareBossData `protobuf:"bytes,5,rep,name=Info,proto3" json:"Info,omitempty"`
	//
	EndTimestamp int64 `protobuf:"varint,6,opt,name=EndTimestamp,proto3" json:"EndTimestamp,omitempty"`
	//
	ActIsOpen int32 `protobuf:"varint,7,opt,name=ActIsOpen,proto3" json:"ActIsOpen,omitempty"`
}

func (m *S2CPlayerDareBossInfo) Reset()         { *m = S2CPlayerDareBossInfo{} }
func (m *S2CPlayerDareBossInfo) String() string { return proto.CompactTextString(m) }
func (*S2CPlayerDareBossInfo) ProtoMessage()    {}
func (m *S2CPlayerDareBossInfo) GetId() uint16  { return PCK_S2CPlayerDareBossInfo_ID }

//
type DareState struct {
	//
	ActId int32 `protobuf:"varint,1,opt,name=ActId,proto3" json:"ActId,omitempty"`
	//
	State int32 `protobuf:"varint,2,opt,name=State,proto3" json:"State,omitempty"`
}

func (m *DareState) Reset()         { *m = DareState{} }
func (m *DareState) String() string { return proto.CompactTextString(m) }
func (*DareState) ProtoMessage()    {}

const PCK_C2SPlayerDareState_ID = 11209 //
//
type C2SPlayerDareState struct {
}

func (m *C2SPlayerDareState) Reset()         { *m = C2SPlayerDareState{} }
func (m *C2SPlayerDareState) String() string { return proto.CompactTextString(m) }
func (*C2SPlayerDareState) ProtoMessage()    {}
func (m *C2SPlayerDareState) GetId() uint16  { return PCK_C2SPlayerDareState_ID }

const PCK_S2CPlayerDareState_ID = 11210 //
//
type S2CPlayerDareState struct {
	//
	States []*DareState `protobuf:"bytes,1,rep,name=States,proto3" json:"States,omitempty"`
}

func (m *S2CPlayerDareState) Reset()         { *m = S2CPlayerDareState{} }
func (m *S2CPlayerDareState) String() string { return proto.CompactTextString(m) }
func (*S2CPlayerDareState) ProtoMessage()    {}
func (m *S2CPlayerDareState) GetId() uint16  { return PCK_S2CPlayerDareState_ID }

const PCK_C2SPlayerDareIsOpen_ID = 11211 //
//
type C2SPlayerDareIsOpen struct {
}

func (m *C2SPlayerDareIsOpen) Reset()         { *m = C2SPlayerDareIsOpen{} }
func (m *C2SPlayerDareIsOpen) String() string { return proto.CompactTextString(m) }
func (*C2SPlayerDareIsOpen) ProtoMessage()    {}
func (m *C2SPlayerDareIsOpen) GetId() uint16  { return PCK_C2SPlayerDareIsOpen_ID }

const PCK_S2CPlayerDareIsOpen_ID = 11212 //
//
type S2CPlayerDareIsOpen struct {
	//
	Open int32 `protobuf:"varint,1,opt,name=Open,proto3" json:"Open,omitempty"`
	//
	EndTimestamp []*DareEndTimestamp `protobuf:"bytes,2,rep,name=EndTimestamp,proto3" json:"EndTimestamp,omitempty"`
}

func (m *S2CPlayerDareIsOpen) Reset()         { *m = S2CPlayerDareIsOpen{} }
func (m *S2CPlayerDareIsOpen) String() string { return proto.CompactTextString(m) }
func (*S2CPlayerDareIsOpen) ProtoMessage()    {}
func (m *S2CPlayerDareIsOpen) GetId() uint16  { return PCK_S2CPlayerDareIsOpen_ID }

//
type DareEndTimestamp struct {
	//
	ActId int32 `protobuf:"varint,1,opt,name=ActId,proto3" json:"ActId,omitempty"`
	//
	EndTimestamp int64 `protobuf:"varint,2,opt,name=EndTimestamp,proto3" json:"EndTimestamp,omitempty"`
	//
	State int32 `protobuf:"varint,3,opt,name=State,proto3" json:"State,omitempty"`
}

func (m *DareEndTimestamp) Reset()         { *m = DareEndTimestamp{} }
func (m *DareEndTimestamp) String() string { return proto.CompactTextString(m) }
func (*DareEndTimestamp) ProtoMessage()    {}

const PCK_C2SFightActDareBoss_ID = 11203 //
//
type C2SFightActDareBoss struct {
	//
	ActId int32 `protobuf:"varint,1,opt,name=ActId,proto3" json:"ActId,omitempty"`
	//
	BossId int32 `protobuf:"varint,2,opt,name=BossId,proto3" json:"BossId,omitempty"`
}

func (m *C2SFightActDareBoss) Reset()         { *m = C2SFightActDareBoss{} }
func (m *C2SFightActDareBoss) String() string { return proto.CompactTextString(m) }
func (*C2SFightActDareBoss) ProtoMessage()    {}
func (m *C2SFightActDareBoss) GetId() uint16  { return PCK_C2SFightActDareBoss_ID }

const PCK_S2CFightActDareBoss_ID = 11204 //
//
type S2CFightActDareBoss struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	ActId int32 `protobuf:"varint,2,opt,name=ActId,proto3" json:"ActId,omitempty"`
	//
	BossId int32 `protobuf:"varint,3,opt,name=BossId,proto3" json:"BossId,omitempty"`
	//
	Prize []*ItemInfo `protobuf:"bytes,4,rep,name=Prize,proto3" json:"Prize,omitempty"`
	//
	Win int32 `protobuf:"varint,5,opt,name=Win,proto3" json:"Win,omitempty"`
}

func (m *S2CFightActDareBoss) Reset()         { *m = S2CFightActDareBoss{} }
func (m *S2CFightActDareBoss) String() string { return proto.CompactTextString(m) }
func (*S2CFightActDareBoss) ProtoMessage()    {}
func (m *S2CFightActDareBoss) GetId() uint16  { return PCK_S2CFightActDareBoss_ID }

const PCK_C2SGetActDareBossPrize_ID = 11205 //
//
type C2SGetActDareBossPrize struct {
	//
	ActId int32 `protobuf:"varint,1,opt,name=ActId,proto3" json:"ActId,omitempty"`
	//
	BossId int32 `protobuf:"varint,2,opt,name=BossId,proto3" json:"BossId,omitempty"`
}

func (m *C2SGetActDareBossPrize) Reset()         { *m = C2SGetActDareBossPrize{} }
func (m *C2SGetActDareBossPrize) String() string { return proto.CompactTextString(m) }
func (*C2SGetActDareBossPrize) ProtoMessage()    {}
func (m *C2SGetActDareBossPrize) GetId() uint16  { return PCK_C2SGetActDareBossPrize_ID }

const PCK_S2CGetActDareBossPrize_ID = 11206 //
//
type S2CGetActDareBossPrize struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	ActId int32 `protobuf:"varint,2,opt,name=ActId,proto3" json:"ActId,omitempty"`
	//
	BossId int32 `protobuf:"varint,3,opt,name=BossId,proto3" json:"BossId,omitempty"`
	//
	Prize []*ItemInfo `protobuf:"bytes,4,rep,name=Prize,proto3" json:"Prize,omitempty"`
}

func (m *S2CGetActDareBossPrize) Reset()         { *m = S2CGetActDareBossPrize{} }
func (m *S2CGetActDareBossPrize) String() string { return proto.CompactTextString(m) }
func (*S2CGetActDareBossPrize) ProtoMessage()    {}
func (m *S2CGetActDareBossPrize) GetId() uint16  { return PCK_S2CGetActDareBossPrize_ID }

const PCK_C2SGetChargeMallList_ID = 12101 //
//
type C2SGetChargeMallList struct {
	//
	T int32 `protobuf:"varint,1,opt,name=T,proto3" json:"T,omitempty"`
}

func (m *C2SGetChargeMallList) Reset()         { *m = C2SGetChargeMallList{} }
func (m *C2SGetChargeMallList) String() string { return proto.CompactTextString(m) }
func (*C2SGetChargeMallList) ProtoMessage()    {}
func (m *C2SGetChargeMallList) GetId() uint16  { return PCK_C2SGetChargeMallList_ID }

//
type GoodsList struct {
	//
	Gid int32 `protobuf:"varint,1,opt,name=Gid,proto3" json:"Gid,omitempty"`
	//
	Items []*ItemInfo `protobuf:"bytes,2,rep,name=Items,proto3" json:"Items,omitempty"`
	//
	Mt int32 `protobuf:"varint,3,opt,name=Mt,proto3" json:"Mt,omitempty"`
	//
	S int32 `protobuf:"varint,4,opt,name=S,proto3" json:"S,omitempty"`
	//
	P int32 `protobuf:"varint,5,opt,name=P,proto3" json:"P,omitempty"`
	//
	Pid int32 `protobuf:"varint,6,opt,name=Pid,proto3" json:"Pid,omitempty"`
	//
	Gt int32 `protobuf:"varint,7,opt,name=Gt,proto3" json:"Gt,omitempty"`
	//
	I int32 `protobuf:"varint,8,opt,name=I,proto3" json:"I,omitempty"`
}

func (m *GoodsList) Reset()         { *m = GoodsList{} }
func (m *GoodsList) String() string { return proto.CompactTextString(m) }
func (*GoodsList) ProtoMessage()    {}

const PCK_S2CGetChargeMallList_ID = 12102 //
//
type S2CGetChargeMallList struct {
	//
	List []*GoodsList `protobuf:"bytes,1,rep,name=List,proto3" json:"List,omitempty"`
}

func (m *S2CGetChargeMallList) Reset()         { *m = S2CGetChargeMallList{} }
func (m *S2CGetChargeMallList) String() string { return proto.CompactTextString(m) }
func (*S2CGetChargeMallList) ProtoMessage()    {}
func (m *S2CGetChargeMallList) GetId() uint16  { return PCK_S2CGetChargeMallList_ID }

const PCK_C2SChargeMallBuy_ID = 12103 //
//
type C2SChargeMallBuy struct {
	//
	Gid int32 `protobuf:"varint,1,opt,name=Gid,proto3" json:"Gid,omitempty"`
}

func (m *C2SChargeMallBuy) Reset()         { *m = C2SChargeMallBuy{} }
func (m *C2SChargeMallBuy) String() string { return proto.CompactTextString(m) }
func (*C2SChargeMallBuy) ProtoMessage()    {}
func (m *C2SChargeMallBuy) GetId() uint16  { return PCK_C2SChargeMallBuy_ID }

const PCK_S2CChargeMallBuy_ID = 12104 //
//
type S2CChargeMallBuy struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	data string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *S2CChargeMallBuy) Reset()         { *m = S2CChargeMallBuy{} }
func (m *S2CChargeMallBuy) String() string { return proto.CompactTextString(m) }
func (*S2CChargeMallBuy) ProtoMessage()    {}
func (m *S2CChargeMallBuy) GetId() uint16  { return PCK_S2CChargeMallBuy_ID }

const PCK_S2CPayNotify_ID = 12105 //
//
type S2CPayNotify struct {
	//
	Type int32 `protobuf:"varint,1,opt,name=Type,proto3" json:"Type,omitempty"`
	//
	GoodsID int32 `protobuf:"varint,2,opt,name=GoodsID,proto3" json:"GoodsID,omitempty"`
}

func (m *S2CPayNotify) Reset()         { *m = S2CPayNotify{} }
func (m *S2CPayNotify) String() string { return proto.CompactTextString(m) }
func (*S2CPayNotify) ProtoMessage()    {}
func (m *S2CPayNotify) GetId() uint16  { return PCK_S2CPayNotify_ID }

//
type ActTask struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	//
	Type int32 `protobuf:"varint,3,opt,name=Type,proto3" json:"Type,omitempty"`
	//
	CounterType int32 `protobuf:"varint,4,opt,name=CounterType,proto3" json:"CounterType,omitempty"`
	//
	Param int32 `protobuf:"varint,5,opt,name=Param,proto3" json:"Param,omitempty"`
	//
	Target int32 `protobuf:"varint,6,opt,name=Target,proto3" json:"Target,omitempty"`
	//
	Prize string `protobuf:"bytes,7,opt,name=Prize,proto3" json:"Prize,omitempty"`
	//
	Order int32 `protobuf:"varint,8,opt,name=Order,proto3" json:"Order,omitempty"`
	//
	AnimId string `protobuf:"bytes,9,opt,name=AnimId,proto3" json:"AnimId,omitempty"`
	//
	ClientFuncId int32 `protobuf:"varint,10,opt,name=ClientFuncId,proto3" json:"ClientFuncId,omitempty"`
	//
	RevelValue int32 `protobuf:"varint,11,opt,name=RevelValue,proto3" json:"RevelValue,omitempty"`
}

func (m *ActTask) Reset()         { *m = ActTask{} }
func (m *ActTask) String() string { return proto.CompactTextString(m) }
func (*ActTask) ProtoMessage()    {}

const PCK_C2SGetActTask_ID = 12151 //
//
type C2SGetActTask struct {
	//
	ActId int32 `protobuf:"varint,1,opt,name=ActId,proto3" json:"ActId,omitempty"`
}

func (m *C2SGetActTask) Reset()         { *m = C2SGetActTask{} }
func (m *C2SGetActTask) String() string { return proto.CompactTextString(m) }
func (*C2SGetActTask) ProtoMessage()    {}
func (m *C2SGetActTask) GetId() uint16  { return PCK_C2SGetActTask_ID }

const PCK_S2CGetActTask_ID = 12152 //
//
type S2CGetActTask struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	ActId int32 `protobuf:"varint,2,opt,name=ActId,proto3" json:"ActId,omitempty"`
	//
	Task []*ActTask `protobuf:"bytes,3,rep,name=Task,proto3" json:"Task,omitempty"`
}

func (m *S2CGetActTask) Reset()         { *m = S2CGetActTask{} }
func (m *S2CGetActTask) String() string { return proto.CompactTextString(m) }
func (*S2CGetActTask) ProtoMessage()    {}
func (m *S2CGetActTask) GetId() uint16  { return PCK_S2CGetActTask_ID }

const PCK_S2CGetFirstRechargeTask_ID = 12153 //
//
type S2CGetFirstRechargeTask struct {
	//
	Task []*ActTask `protobuf:"bytes,1,rep,name=Task,proto3" json:"Task,omitempty"`
}

func (m *S2CGetFirstRechargeTask) Reset()         { *m = S2CGetFirstRechargeTask{} }
func (m *S2CGetFirstRechargeTask) String() string { return proto.CompactTextString(m) }
func (*S2CGetFirstRechargeTask) ProtoMessage()    {}
func (m *S2CGetFirstRechargeTask) GetId() uint16  { return PCK_S2CGetFirstRechargeTask_ID }

//
type GodTowerDayInfo struct {
	//
	Day int32 `protobuf:"varint,1,opt,name=Day,proto3" json:"Day,omitempty"`
	//
	stage int32 `protobuf:"varint,2,opt,name=stage,proto3" json:"stage,omitempty"`
}

func (m *GodTowerDayInfo) Reset()         { *m = GodTowerDayInfo{} }
func (m *GodTowerDayInfo) String() string { return proto.CompactTextString(m) }
func (*GodTowerDayInfo) ProtoMessage()    {}

const PCK_C2SGetGodTowerActTaskDay_ID = 12161 //
//
type C2SGetGodTowerActTaskDay struct {
}

func (m *C2SGetGodTowerActTaskDay) Reset()         { *m = C2SGetGodTowerActTaskDay{} }
func (m *C2SGetGodTowerActTaskDay) String() string { return proto.CompactTextString(m) }
func (*C2SGetGodTowerActTaskDay) ProtoMessage()    {}
func (m *C2SGetGodTowerActTaskDay) GetId() uint16  { return PCK_C2SGetGodTowerActTaskDay_ID }

const PCK_S2CGetGodTowerActTaskDay_ID = 12162 //
//
type S2CGetGodTowerActTaskDay struct {
	//
	Task []*GodTowerDayInfo `protobuf:"bytes,3,rep,name=Task,proto3" json:"Task,omitempty"`
}

func (m *S2CGetGodTowerActTaskDay) Reset()         { *m = S2CGetGodTowerActTaskDay{} }
func (m *S2CGetGodTowerActTaskDay) String() string { return proto.CompactTextString(m) }
func (*S2CGetGodTowerActTaskDay) ProtoMessage()    {}
func (m *S2CGetGodTowerActTaskDay) GetId() uint16  { return PCK_S2CGetGodTowerActTaskDay_ID }

const PCK_C2SGetInvestInfo_ID = 12201 //
//
type C2SGetInvestInfo struct {
	//
	ActId int32 `protobuf:"varint,1,opt,name=ActId,proto3" json:"ActId,omitempty"`
}

func (m *C2SGetInvestInfo) Reset()         { *m = C2SGetInvestInfo{} }
func (m *C2SGetInvestInfo) String() string { return proto.CompactTextString(m) }
func (*C2SGetInvestInfo) ProtoMessage()    {}
func (m *C2SGetInvestInfo) GetId() uint16  { return PCK_C2SGetInvestInfo_ID }

const PCK_S2CGetInvestInfo_ID = 12202 //
//
type S2CGetInvestInfo struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	ActId int32 `protobuf:"varint,2,opt,name=ActId,proto3" json:"ActId,omitempty"`
	//
	AnimId int32 `protobuf:"varint,3,opt,name=AnimId,proto3" json:"AnimId,omitempty"`
	//
	Task []*ActTask `protobuf:"bytes,4,rep,name=Task,proto3" json:"Task,omitempty"`
	//
	Name string `protobuf:"bytes,5,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (m *S2CGetInvestInfo) Reset()         { *m = S2CGetInvestInfo{} }
func (m *S2CGetInvestInfo) String() string { return proto.CompactTextString(m) }
func (*S2CGetInvestInfo) ProtoMessage()    {}
func (m *S2CGetInvestInfo) GetId() uint16  { return PCK_S2CGetInvestInfo_ID }

const PCK_C2SPlayerInvestData_ID = 12205 //
//
type C2SPlayerInvestData struct {
	//
	ActId int32 `protobuf:"varint,1,opt,name=ActId,proto3" json:"ActId,omitempty"`
}

func (m *C2SPlayerInvestData) Reset()         { *m = C2SPlayerInvestData{} }
func (m *C2SPlayerInvestData) String() string { return proto.CompactTextString(m) }
func (*C2SPlayerInvestData) ProtoMessage()    {}
func (m *C2SPlayerInvestData) GetId() uint16  { return PCK_C2SPlayerInvestData_ID }

const PCK_S2CPlayerInvestData_ID = 12206 //
//
type S2CPlayerInvestData struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	ActId int32 `protobuf:"varint,2,opt,name=ActId,proto3" json:"ActId,omitempty"`
	//
	Buy int32 `protobuf:"varint,3,opt,name=Buy,proto3" json:"Buy,omitempty"`
}

func (m *S2CPlayerInvestData) Reset()         { *m = S2CPlayerInvestData{} }
func (m *S2CPlayerInvestData) String() string { return proto.CompactTextString(m) }
func (*S2CPlayerInvestData) ProtoMessage()    {}
func (m *S2CPlayerInvestData) GetId() uint16  { return PCK_S2CPlayerInvestData_ID }

const PCK_C2SBuyInvest_ID = 12203 //
//
type C2SBuyInvest struct {
	//
	ActId int32 `protobuf:"varint,1,opt,name=ActId,proto3" json:"ActId,omitempty"`
}

func (m *C2SBuyInvest) Reset()         { *m = C2SBuyInvest{} }
func (m *C2SBuyInvest) String() string { return proto.CompactTextString(m) }
func (*C2SBuyInvest) ProtoMessage()    {}
func (m *C2SBuyInvest) GetId() uint16  { return PCK_C2SBuyInvest_ID }

const PCK_S2CBuyInvest_ID = 12204 //
//
type S2CBuyInvest struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	ActId int32 `protobuf:"varint,2,opt,name=ActId,proto3" json:"ActId,omitempty"`
}

func (m *S2CBuyInvest) Reset()         { *m = S2CBuyInvest{} }
func (m *S2CBuyInvest) String() string { return proto.CompactTextString(m) }
func (*S2CBuyInvest) ProtoMessage()    {}
func (m *S2CBuyInvest) GetId() uint16  { return PCK_S2CBuyInvest_ID }

//
type RankData struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	RankMin int32 `protobuf:"varint,2,opt,name=RankMin,proto3" json:"RankMin,omitempty"`
	//
	RankMax int32 `protobuf:"varint,3,opt,name=RankMax,proto3" json:"RankMax,omitempty"`
	//
	BasePrize string `protobuf:"bytes,4,opt,name=BasePrize,proto3" json:"BasePrize,omitempty"`
	//
	AddPrize string `protobuf:"bytes,5,opt,name=AddPrize,proto3" json:"AddPrize,omitempty"`
	//
	PrizeTitle string `protobuf:"bytes,6,opt,name=PrizeTitle,proto3" json:"PrizeTitle,omitempty"`
	//
	PrizeContent string `protobuf:"bytes,7,opt,name=PrizeContent,proto3" json:"PrizeContent,omitempty"`
	//
	CompareParam int32 `protobuf:"varint,8,opt,name=CompareParam,proto3" json:"CompareParam,omitempty"`
}

func (m *RankData) Reset()         { *m = RankData{} }
func (m *RankData) String() string { return proto.CompactTextString(m) }
func (*RankData) ProtoMessage()    {}

const PCK_C2SGetActRankInfo_ID = 12301 //
//
type C2SGetActRankInfo struct {
	//
	ActId int32 `protobuf:"varint,1,opt,name=ActId,proto3" json:"ActId,omitempty"`
}

func (m *C2SGetActRankInfo) Reset()         { *m = C2SGetActRankInfo{} }
func (m *C2SGetActRankInfo) String() string { return proto.CompactTextString(m) }
func (*C2SGetActRankInfo) ProtoMessage()    {}
func (m *C2SGetActRankInfo) GetId() uint16  { return PCK_C2SGetActRankInfo_ID }

const PCK_S2CGetActRankInfo_ID = 12302 //
//
type S2CGetActRankInfo struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	ActId int32 `protobuf:"varint,2,opt,name=ActId,proto3" json:"ActId,omitempty"`
	//
	Name string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	//
	RankType int32 `protobuf:"varint,4,opt,name=RankType,proto3" json:"RankType,omitempty"`
	//
	Param int32 `protobuf:"varint,5,opt,name=Param,proto3" json:"Param,omitempty"`
	//
	Data []*RankData `protobuf:"bytes,6,rep,name=Data,proto3" json:"Data,omitempty"`
}

func (m *S2CGetActRankInfo) Reset()         { *m = S2CGetActRankInfo{} }
func (m *S2CGetActRankInfo) String() string { return proto.CompactTextString(m) }
func (*S2CGetActRankInfo) ProtoMessage()    {}
func (m *S2CGetActRankInfo) GetId() uint16  { return PCK_S2CGetActRankInfo_ID }

const PCK_C2SGetActRankData_ID = 12303 //
//
type C2SGetActRankData struct {
	//
	ActId int32 `protobuf:"varint,1,opt,name=ActId,proto3" json:"ActId,omitempty"`
	//
	FullDataRank []int32 `protobuf:"varint,2,rep,name=FullDataRank,proto3" json:"FullDataRank,omitempty"`
	//
	SimpleDataRank []int32 `protobuf:"varint,3,rep,name=SimpleDataRank,proto3" json:"SimpleDataRank,omitempty"`
}

func (m *C2SGetActRankData) Reset()         { *m = C2SGetActRankData{} }
func (m *C2SGetActRankData) String() string { return proto.CompactTextString(m) }
func (*C2SGetActRankData) ProtoMessage()    {}
func (m *C2SGetActRankData) GetId() uint16  { return PCK_C2SGetActRankData_ID }

const PCK_S2CGetActRankData_ID = 12304 //
//
type S2CGetActRankData struct {
	//
	ActId int32 `protobuf:"varint,1,opt,name=ActId,proto3" json:"ActId,omitempty"`
	//
	FullDataRank []*Rank `protobuf:"bytes,2,rep,name=FullDataRank,proto3" json:"FullDataRank,omitempty"`
	//
	SimpleDataRank []*SimpleRank `protobuf:"bytes,3,rep,name=SimpleDataRank,proto3" json:"SimpleDataRank,omitempty"`
	//
	MyData *SimpleRank `protobuf:"bytes,4,opt,name=MyData,proto3" json:"MyData,omitempty"`
}

func (m *S2CGetActRankData) Reset()         { *m = S2CGetActRankData{} }
func (m *S2CGetActRankData) String() string { return proto.CompactTextString(m) }
func (*S2CGetActRankData) ProtoMessage()    {}
func (m *S2CGetActRankData) GetId() uint16  { return PCK_S2CGetActRankData_ID }

const PCK_C2SGetActPicture_ID = 12350 //
//
type C2SGetActPicture struct {
	//
	ActId int32 `protobuf:"varint,1,opt,name=ActId,proto3" json:"ActId,omitempty"`
}

func (m *C2SGetActPicture) Reset()         { *m = C2SGetActPicture{} }
func (m *C2SGetActPicture) String() string { return proto.CompactTextString(m) }
func (*C2SGetActPicture) ProtoMessage()    {}
func (m *C2SGetActPicture) GetId() uint16  { return PCK_C2SGetActPicture_ID }

const PCK_S2CGetActPicture_ID = 12351 //
//
type S2CGetActPicture struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	ActId int32 `protobuf:"varint,2,opt,name=ActId,proto3" json:"ActId,omitempty"`
	//
	Task []*ActTask `protobuf:"bytes,3,rep,name=Task,proto3" json:"Task,omitempty"`
	//
	Pictures []int32 `protobuf:"varint,4,rep,name=Pictures,proto3" json:"Pictures,omitempty"`
	//
	PicAnimId int32 `protobuf:"varint,5,opt,name=PicAnimId,proto3" json:"PicAnimId,omitempty"`
	//
	PartCount int32 `protobuf:"varint,6,opt,name=PartCount,proto3" json:"PartCount,omitempty"`
	//
	PartAnimId []int32 `protobuf:"varint,7,rep,name=PartAnimId,proto3" json:"PartAnimId,omitempty"`
	//
	StickCount int32 `protobuf:"varint,8,opt,name=StickCount,proto3" json:"StickCount,omitempty"`
	//
	Prize int32 `protobuf:"varint,9,opt,name=Prize,proto3" json:"Prize,omitempty"`
	//
	Intro int32 `protobuf:"varint,10,opt,name=Intro,proto3" json:"Intro,omitempty"`
	//
	PartItem []int32 `protobuf:"varint,11,rep,name=PartItem,proto3" json:"PartItem,omitempty"`
}

func (m *S2CGetActPicture) Reset()         { *m = S2CGetActPicture{} }
func (m *S2CGetActPicture) String() string { return proto.CompactTextString(m) }
func (*S2CGetActPicture) ProtoMessage()    {}
func (m *S2CGetActPicture) GetId() uint16  { return PCK_S2CGetActPicture_ID }

const PCK_C2SActPictureLight_ID = 12352 //
//
type C2SActPictureLight struct {
	//
	ActId int32 `protobuf:"varint,1,opt,name=ActId,proto3" json:"ActId,omitempty"`
	//
	Part int32 `protobuf:"varint,2,opt,name=Part,proto3" json:"Part,omitempty"`
}

func (m *C2SActPictureLight) Reset()         { *m = C2SActPictureLight{} }
func (m *C2SActPictureLight) String() string { return proto.CompactTextString(m) }
func (*C2SActPictureLight) ProtoMessage()    {}
func (m *C2SActPictureLight) GetId() uint16  { return PCK_C2SActPictureLight_ID }

const PCK_S2CActPictureLight_ID = 12353 //
//
type S2CActPictureLight struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	ActId int32 `protobuf:"varint,2,opt,name=ActId,proto3" json:"ActId,omitempty"`
	//
	Pictures []int32 `protobuf:"varint,3,rep,name=Pictures,proto3" json:"Pictures,omitempty"`
}

func (m *S2CActPictureLight) Reset()         { *m = S2CActPictureLight{} }
func (m *S2CActPictureLight) String() string { return proto.CompactTextString(m) }
func (*S2CActPictureLight) ProtoMessage()    {}
func (m *S2CActPictureLight) GetId() uint16  { return PCK_S2CActPictureLight_ID }

const PCK_C2SActRecoveryData_ID = 12501 //
//
type C2SActRecoveryData struct {
	//
	ActId int32 `protobuf:"varint,1,opt,name=ActId,proto3" json:"ActId,omitempty"`
}

func (m *C2SActRecoveryData) Reset()         { *m = C2SActRecoveryData{} }
func (m *C2SActRecoveryData) String() string { return proto.CompactTextString(m) }
func (*C2SActRecoveryData) ProtoMessage()    {}
func (m *C2SActRecoveryData) GetId() uint16  { return PCK_C2SActRecoveryData_ID }

//
type QualityPrice struct {
	//
	Quality int32 `protobuf:"varint,1,opt,name=Quality,proto3" json:"Quality,omitempty"`
	//
	ItemInfo []*ItemInfo `protobuf:"bytes,2,rep,name=ItemInfo,proto3" json:"ItemInfo,omitempty"`
}

func (m *QualityPrice) Reset()         { *m = QualityPrice{} }
func (m *QualityPrice) String() string { return proto.CompactTextString(m) }
func (*QualityPrice) ProtoMessage()    {}

//
type RecoverData struct {
	//
	LevelMin int64 `protobuf:"varint,1,opt,name=LevelMin,proto3" json:"LevelMin,omitempty"`
	//
	LevelMax int64 `protobuf:"varint,2,opt,name=LevelMax,proto3" json:"LevelMax,omitempty"`
	//
	Name string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	//
	QualityPrice []*QualityPrice `protobuf:"bytes,4,rep,name=QualityPrice,proto3" json:"QualityPrice,omitempty"`
	//
	Id int32 `protobuf:"varint,5,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	StateLevel int64 `protobuf:"varint,6,opt,name=StateLevel,proto3" json:"StateLevel,omitempty"`
}

func (m *RecoverData) Reset()         { *m = RecoverData{} }
func (m *RecoverData) String() string { return proto.CompactTextString(m) }
func (*RecoverData) ProtoMessage()    {}

const PCK_S2CActRecoveryData_ID = 12502 //
//
type S2CActRecoveryData struct {
	//
	items []*RecoverData `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	//
	FreeTimes int32 `protobuf:"varint,2,opt,name=FreeTimes,proto3" json:"FreeTimes,omitempty"`
	//
	MonthTimes int32 `protobuf:"varint,3,opt,name=MonthTimes,proto3" json:"MonthTimes,omitempty"`
	//
	ActId int32 `protobuf:"varint,5,opt,name=ActId,proto3" json:"ActId,omitempty"`
}

func (m *S2CActRecoveryData) Reset()         { *m = S2CActRecoveryData{} }
func (m *S2CActRecoveryData) String() string { return proto.CompactTextString(m) }
func (*S2CActRecoveryData) ProtoMessage()    {}
func (m *S2CActRecoveryData) GetId() uint16  { return PCK_S2CActRecoveryData_ID }

const PCK_C2SActRecoveryEquip_ID = 12503 //
//
type C2SActRecoveryEquip struct {
	//
	ActId int32 `protobuf:"varint,1,opt,name=ActId,proto3" json:"ActId,omitempty"`
	//
	Id string `protobuf:"bytes,2,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	RecoveryId int32 `protobuf:"varint,3,opt,name=RecoveryId,proto3" json:"RecoveryId,omitempty"`
}

func (m *C2SActRecoveryEquip) Reset()         { *m = C2SActRecoveryEquip{} }
func (m *C2SActRecoveryEquip) String() string { return proto.CompactTextString(m) }
func (*C2SActRecoveryEquip) ProtoMessage()    {}
func (m *C2SActRecoveryEquip) GetId() uint16  { return PCK_C2SActRecoveryEquip_ID }

const PCK_S2CActRecoveryEquip_ID = 12504 //
//
type S2CActRecoveryEquip struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Mul []int32 `protobuf:"varint,2,rep,name=Mul,proto3" json:"Mul,omitempty"`
}

func (m *S2CActRecoveryEquip) Reset()         { *m = S2CActRecoveryEquip{} }
func (m *S2CActRecoveryEquip) String() string { return proto.CompactTextString(m) }
func (*S2CActRecoveryEquip) ProtoMessage()    {}
func (m *S2CActRecoveryEquip) GetId() uint16  { return PCK_S2CActRecoveryEquip_ID }

const PCK_C2SShowItem_ID = 1801 //
//
type C2SShowItem struct {
	//
	Type int32 `protobuf:"varint,1,opt,name=Type,proto3" json:"Type,omitempty"`
	//
	Param1 int64 `protobuf:"varint,2,opt,name=Param1,proto3" json:"Param1,omitempty"`
	//
	Param2 string `protobuf:"bytes,3,opt,name=Param2,proto3" json:"Param2,omitempty"`
	//
	ChannelType int32 `protobuf:"varint,4,opt,name=ChannelType,proto3" json:"ChannelType,omitempty"`
}

func (m *C2SShowItem) Reset()         { *m = C2SShowItem{} }
func (m *C2SShowItem) String() string { return proto.CompactTextString(m) }
func (*C2SShowItem) ProtoMessage()    {}
func (m *C2SShowItem) GetId() uint16  { return PCK_C2SShowItem_ID }

const PCK_S2CShowItem_ID = 1802 //
//
type S2CShowItem struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CShowItem) Reset()         { *m = S2CShowItem{} }
func (m *S2CShowItem) String() string { return proto.CompactTextString(m) }
func (*S2CShowItem) ProtoMessage()    {}
func (m *S2CShowItem) GetId() uint16  { return PCK_S2CShowItem_ID }

const PCK_S2CReportShowItem_ID = 1803 //
//
type S2CReportShowItem struct {
	//
	UserId int32 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	Nick string `protobuf:"bytes,2,opt,name=Nick,proto3" json:"Nick,omitempty"`
	//
	RoleId int64 `protobuf:"varint,3,opt,name=RoleId,proto3" json:"RoleId,omitempty"`
	//
	GodLevel int64 `protobuf:"varint,4,opt,name=GodLevel,proto3" json:"GodLevel,omitempty"`
	//
	Type int32 `protobuf:"varint,5,opt,name=Type,proto3" json:"Type,omitempty"`
	//
	Param1 int64 `protobuf:"varint,6,opt,name=Param1,proto3" json:"Param1,omitempty"`
	//
	Param2 string `protobuf:"bytes,7,opt,name=Param2,proto3" json:"Param2,omitempty"`
	//
	ChannelType int32 `protobuf:"varint,8,opt,name=ChannelType,proto3" json:"ChannelType,omitempty"`
	//
	ShowId int32 `protobuf:"varint,9,opt,name=ShowId,proto3" json:"ShowId,omitempty"`
}

func (m *S2CReportShowItem) Reset()         { *m = S2CReportShowItem{} }
func (m *S2CReportShowItem) String() string { return proto.CompactTextString(m) }
func (*S2CReportShowItem) ProtoMessage()    {}
func (m *S2CReportShowItem) GetId() uint16  { return PCK_S2CReportShowItem_ID }

const PCK_C2SGetShowInfo_ID = 1804 //
//
type C2SGetShowInfo struct {
	//
	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	Type int32 `protobuf:"varint,2,opt,name=Type,proto3" json:"Type,omitempty"`
	//
	Param1 int64 `protobuf:"varint,3,opt,name=Param1,proto3" json:"Param1,omitempty"`
	//
	Param2 string `protobuf:"bytes,4,opt,name=Param2,proto3" json:"Param2,omitempty"`
	//
	ShowId int32 `protobuf:"varint,6,opt,name=ShowId,proto3" json:"ShowId,omitempty"`
}

func (m *C2SGetShowInfo) Reset()         { *m = C2SGetShowInfo{} }
func (m *C2SGetShowInfo) String() string { return proto.CompactTextString(m) }
func (*C2SGetShowInfo) ProtoMessage()    {}
func (m *C2SGetShowInfo) GetId() uint16  { return PCK_C2SGetShowInfo_ID }

const PCK_S2CGetShowInfo_ID = 1805 //
//
type S2CGetShowInfo struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Type int32 `protobuf:"varint,2,opt,name=Type,proto3" json:"Type,omitempty"`
	//
	ItemInfo *ItemData `protobuf:"bytes,3,opt,name=ItemInfo,proto3" json:"ItemInfo,omitempty"`
	//
	PartnerInfo *Partner `protobuf:"bytes,4,opt,name=PartnerInfo,proto3" json:"PartnerInfo,omitempty"`
	//
	PreciousInfo *Precious `protobuf:"bytes,5,opt,name=PreciousInfo,proto3" json:"PreciousInfo,omitempty"`
	//
	GodEquip []*Grade `protobuf:"bytes,6,rep,name=GodEquip,proto3" json:"GodEquip,omitempty"`
	//
	PreciousPosInfo *PreciousPos `protobuf:"bytes,7,opt,name=PreciousPosInfo,proto3" json:"PreciousPosInfo,omitempty"`
	//
	ShowId int32 `protobuf:"varint,8,opt,name=ShowId,proto3" json:"ShowId,omitempty"`
}

func (m *S2CGetShowInfo) Reset()         { *m = S2CGetShowInfo{} }
func (m *S2CGetShowInfo) String() string { return proto.CompactTextString(m) }
func (*S2CGetShowInfo) ProtoMessage()    {}
func (m *S2CGetShowInfo) GetId() uint16  { return PCK_S2CGetShowInfo_ID }

const PCK_S2KShowItem_ID = 1806 //
//
type S2KShowItem struct {
	//
	ReportMsg *S2CReportShowItem `protobuf:"bytes,1,opt,name=ReportMsg,proto3" json:"ReportMsg,omitempty"`
	//
	ShowMsg *S2CGetShowInfo `protobuf:"bytes,2,opt,name=ShowMsg,proto3" json:"ShowMsg,omitempty"`
}

func (m *S2KShowItem) Reset()         { *m = S2KShowItem{} }
func (m *S2KShowItem) String() string { return proto.CompactTextString(m) }
func (*S2KShowItem) ProtoMessage()    {}
func (m *S2KShowItem) GetId() uint16  { return PCK_S2KShowItem_ID }

//
type OnlineUser struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Nick string `protobuf:"bytes,2,opt,name=Nick,proto3" json:"Nick,omitempty"`
	//
	Level int64 `protobuf:"varint,3,opt,name=Level,proto3" json:"Level,omitempty"`
}

func (m *OnlineUser) Reset()         { *m = OnlineUser{} }
func (m *OnlineUser) String() string { return proto.CompactTextString(m) }
func (*OnlineUser) ProtoMessage()    {}

//
type ApiResult struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Msg string `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
}

func (m *ApiResult) Reset()         { *m = ApiResult{} }
func (m *ApiResult) String() string { return proto.CompactTextString(m) }
func (*ApiResult) ProtoMessage()    {}

const PCK_C2SRedeemCode_ID = 12401 //
//
type C2SRedeemCode struct {
	//
	Code string `protobuf:"bytes,1,opt,name=Code,proto3" json:"Code,omitempty"`
	//
	ActivityId int32 `protobuf:"varint,2,opt,name=ActivityId,proto3" json:"ActivityId,omitempty"`
}

func (m *C2SRedeemCode) Reset()         { *m = C2SRedeemCode{} }
func (m *C2SRedeemCode) String() string { return proto.CompactTextString(m) }
func (*C2SRedeemCode) ProtoMessage()    {}
func (m *C2SRedeemCode) GetId() uint16  { return PCK_C2SRedeemCode_ID }

const PCK_S2CRedeemCode_ID = 12402 //
//
type S2CRedeemCode struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	RewardList string `protobuf:"bytes,2,opt,name=RewardList,proto3" json:"RewardList,omitempty"`
	//
	ActivityId int32 `protobuf:"varint,3,opt,name=ActivityId,proto3" json:"ActivityId,omitempty"`
}

func (m *S2CRedeemCode) Reset()         { *m = S2CRedeemCode{} }
func (m *S2CRedeemCode) String() string { return proto.CompactTextString(m) }
func (*S2CRedeemCode) ProtoMessage()    {}
func (m *S2CRedeemCode) GetId() uint16  { return PCK_S2CRedeemCode_ID }

const PCK_C2SRealName_ID = 120 //
//
type C2SRealName struct {
	//
	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	//
	Card string `protobuf:"bytes,2,opt,name=Card,proto3" json:"Card,omitempty"`
}

func (m *C2SRealName) Reset()         { *m = C2SRealName{} }
func (m *C2SRealName) String() string { return proto.CompactTextString(m) }
func (*C2SRealName) ProtoMessage()    {}
func (m *C2SRealName) GetId() uint16  { return PCK_C2SRealName_ID }

const PCK_S2CRealName_ID = 121 //
//
type S2CRealName struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CRealName) Reset()         { *m = S2CRealName{} }
func (m *S2CRealName) String() string { return proto.CompactTextString(m) }
func (*S2CRealName) ProtoMessage()    {}
func (m *S2CRealName) GetId() uint16  { return PCK_S2CRealName_ID }

const PCK_C2SGetRealNamePrize_ID = 122 //
//
type C2SGetRealNamePrize struct {
	//
	RealName int32 `protobuf:"varint,1,opt,name=RealName,proto3" json:"RealName,omitempty"`
}

func (m *C2SGetRealNamePrize) Reset()         { *m = C2SGetRealNamePrize{} }
func (m *C2SGetRealNamePrize) String() string { return proto.CompactTextString(m) }
func (*C2SGetRealNamePrize) ProtoMessage()    {}
func (m *C2SGetRealNamePrize) GetId() uint16  { return PCK_C2SGetRealNamePrize_ID }

const PCK_S2CGetRealNamePrize_ID = 123 //
//
type S2CGetRealNamePrize struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CGetRealNamePrize) Reset()         { *m = S2CGetRealNamePrize{} }
func (m *S2CGetRealNamePrize) String() string { return proto.CompactTextString(m) }
func (*S2CGetRealNamePrize) ProtoMessage()    {}
func (m *S2CGetRealNamePrize) GetId() uint16  { return PCK_S2CGetRealNamePrize_ID }

const PCK_S2CReportFcm_ID = 30 //
//
type S2CReportFcm struct {
	//
	FcmMinute int32 `protobuf:"varint,1,opt,name=FcmMinute,proto3" json:"FcmMinute,omitempty"`
	//
	FcmStatus int32 `protobuf:"varint,2,opt,name=FcmStatus,proto3" json:"FcmStatus,omitempty"`
}

func (m *S2CReportFcm) Reset()         { *m = S2CReportFcm{} }
func (m *S2CReportFcm) String() string { return proto.CompactTextString(m) }
func (*S2CReportFcm) ProtoMessage()    {}
func (m *S2CReportFcm) GetId() uint16  { return PCK_S2CReportFcm_ID }

const PCK_C2SSetFcm_ID = 31 //
//
type C2SSetFcm struct {
	//
	Fcm int32 `protobuf:"varint,1,opt,name=Fcm,proto3" json:"Fcm,omitempty"`
}

func (m *C2SSetFcm) Reset()         { *m = C2SSetFcm{} }
func (m *C2SSetFcm) String() string { return proto.CompactTextString(m) }
func (*C2SSetFcm) ProtoMessage()    {}
func (m *C2SSetFcm) GetId() uint16  { return PCK_C2SSetFcm_ID }

const PCK_C2SGetLimitCloth_ID = 124 //
//
type C2SGetLimitCloth struct {
}

func (m *C2SGetLimitCloth) Reset()         { *m = C2SGetLimitCloth{} }
func (m *C2SGetLimitCloth) String() string { return proto.CompactTextString(m) }
func (*C2SGetLimitCloth) ProtoMessage()    {}
func (m *C2SGetLimitCloth) GetId() uint16  { return PCK_C2SGetLimitCloth_ID }

const PCK_S2CGetLimitCloth_ID = 125 //
//
type S2CGetLimitCloth struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CGetLimitCloth) Reset()         { *m = S2CGetLimitCloth{} }
func (m *S2CGetLimitCloth) String() string { return proto.CompactTextString(m) }
func (*S2CGetLimitCloth) ProtoMessage()    {}
func (m *S2CGetLimitCloth) GetId() uint16  { return PCK_S2CGetLimitCloth_ID }

const PCK_S2CGetLimitClothEnd_ID = 126 //
//
type S2CGetLimitClothEnd struct {
}

func (m *S2CGetLimitClothEnd) Reset()         { *m = S2CGetLimitClothEnd{} }
func (m *S2CGetLimitClothEnd) String() string { return proto.CompactTextString(m) }
func (*S2CGetLimitClothEnd) ProtoMessage()    {}
func (m *S2CGetLimitClothEnd) GetId() uint16  { return PCK_S2CGetLimitClothEnd_ID }

//
type BossHillBox struct {
	//
	BoxId int64 `protobuf:"varint,1,opt,name=BoxId,proto3" json:"BoxId,omitempty"`
	//
	Pos int32 `protobuf:"varint,2,opt,name=Pos,proto3" json:"Pos,omitempty"`
	//
	OpenTime int64 `protobuf:"varint,3,opt,name=OpenTime,proto3" json:"OpenTime,omitempty"`
}

func (m *BossHillBox) Reset()         { *m = BossHillBox{} }
func (m *BossHillBox) String() string { return proto.CompactTextString(m) }
func (*BossHillBox) ProtoMessage()    {}

//
type BossHillData struct {
	//
	Box []*BossHillBox `protobuf:"bytes,1,rep,name=Box,proto3" json:"Box,omitempty"`
	//
	TmpBoxId int64 `protobuf:"varint,2,opt,name=TmpBoxId,proto3" json:"TmpBoxId,omitempty"`
	//
	Instances []*IntAttr `protobuf:"bytes,3,rep,name=Instances,proto3" json:"Instances,omitempty"`
}

func (m *BossHillData) Reset()         { *m = BossHillData{} }
func (m *BossHillData) String() string { return proto.CompactTextString(m) }
func (*BossHillData) ProtoMessage()    {}

const PCK_S2CBossHillData_ID = 671 //
//
type S2CBossHillData struct {
	//
	Box []*BossHillBox `protobuf:"bytes,1,rep,name=Box,proto3" json:"Box,omitempty"`
	//
	Instances []*IntAttr `protobuf:"bytes,2,rep,name=Instances,proto3" json:"Instances,omitempty"`
}

func (m *S2CBossHillData) Reset()         { *m = S2CBossHillData{} }
func (m *S2CBossHillData) String() string { return proto.CompactTextString(m) }
func (*S2CBossHillData) ProtoMessage()    {}
func (m *S2CBossHillData) GetId() uint16  { return PCK_S2CBossHillData_ID }

const PCK_C2SBossHillFight_ID = 672 //
//
type C2SBossHillFight struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2SBossHillFight) Reset()         { *m = C2SBossHillFight{} }
func (m *C2SBossHillFight) String() string { return proto.CompactTextString(m) }
func (*C2SBossHillFight) ProtoMessage()    {}
func (m *C2SBossHillFight) GetId() uint16  { return PCK_C2SBossHillFight_ID }

const PCK_S2CBossHillFight_ID = 673 //
//
type S2CBossHillFight struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	BoxId int64 `protobuf:"varint,2,opt,name=BoxId,proto3" json:"BoxId,omitempty"`
	//
	Id int32 `protobuf:"varint,3,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Win bool `protobuf:"varint,4,opt,name=Win,proto3" json:"Win,omitempty"`
}

func (m *S2CBossHillFight) Reset()         { *m = S2CBossHillFight{} }
func (m *S2CBossHillFight) String() string { return proto.CompactTextString(m) }
func (*S2CBossHillFight) ProtoMessage()    {}
func (m *S2CBossHillFight) GetId() uint16  { return PCK_S2CBossHillFight_ID }

const PCK_C2SBossHillReplace_ID = 674 //
//
type C2SBossHillReplace struct {
	//
	Pos int32 `protobuf:"varint,1,opt,name=Pos,proto3" json:"Pos,omitempty"`
}

func (m *C2SBossHillReplace) Reset()         { *m = C2SBossHillReplace{} }
func (m *C2SBossHillReplace) String() string { return proto.CompactTextString(m) }
func (*C2SBossHillReplace) ProtoMessage()    {}
func (m *C2SBossHillReplace) GetId() uint16  { return PCK_C2SBossHillReplace_ID }

const PCK_S2CBossHillReplace_ID = 675 //
//
type S2CBossHillReplace struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	BoxId int64 `protobuf:"varint,2,opt,name=BoxId,proto3" json:"BoxId,omitempty"`
	//
	OpenTime int64 `protobuf:"varint,3,opt,name=OpenTime,proto3" json:"OpenTime,omitempty"`
	//
	Pos int32 `protobuf:"varint,4,opt,name=Pos,proto3" json:"Pos,omitempty"`
}

func (m *S2CBossHillReplace) Reset()         { *m = S2CBossHillReplace{} }
func (m *S2CBossHillReplace) String() string { return proto.CompactTextString(m) }
func (*S2CBossHillReplace) ProtoMessage()    {}
func (m *S2CBossHillReplace) GetId() uint16  { return PCK_S2CBossHillReplace_ID }

const PCK_C2SBossHillOpen_ID = 676 //
//
type C2SBossHillOpen struct {
	//
	IsCoin int32 `protobuf:"varint,1,opt,name=IsCoin,proto3" json:"IsCoin,omitempty"`
	//
	Pos int32 `protobuf:"varint,2,opt,name=Pos,proto3" json:"Pos,omitempty"`
}

func (m *C2SBossHillOpen) Reset()         { *m = C2SBossHillOpen{} }
func (m *C2SBossHillOpen) String() string { return proto.CompactTextString(m) }
func (*C2SBossHillOpen) ProtoMessage()    {}
func (m *C2SBossHillOpen) GetId() uint16  { return PCK_C2SBossHillOpen_ID }

const PCK_S2CBossHillOpen_ID = 677 //
//
type S2CBossHillOpen struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Pos int32 `protobuf:"varint,2,opt,name=Pos,proto3" json:"Pos,omitempty"`
	//
	Items []*ItemData `protobuf:"bytes,3,rep,name=Items,proto3" json:"Items,omitempty"`
}

func (m *S2CBossHillOpen) Reset()         { *m = S2CBossHillOpen{} }
func (m *S2CBossHillOpen) String() string { return proto.CompactTextString(m) }
func (*S2CBossHillOpen) ProtoMessage()    {}
func (m *S2CBossHillOpen) GetId() uint16  { return PCK_S2CBossHillOpen_ID }

//
type Instance81Data struct {
	//
	InstanceId int64 `protobuf:"varint,1,opt,name=InstanceId,proto3" json:"InstanceId,omitempty"`
	//
	FirstTeam []string `protobuf:"bytes,2,rep,name=FirstTeam,proto3" json:"FirstTeam,omitempty"`
	//
	FirstTime int64 `protobuf:"varint,3,opt,name=FirstTime,proto3" json:"FirstTime,omitempty"`
	//
	FirstRound int64 `protobuf:"varint,4,opt,name=FirstRound,proto3" json:"FirstRound,omitempty"`
	//
	QuickTeam []string `protobuf:"bytes,5,rep,name=QuickTeam,proto3" json:"QuickTeam,omitempty"`
	//
	QuickTime int64 `protobuf:"varint,6,opt,name=QuickTime,proto3" json:"QuickTime,omitempty"`
	//
	QuickRound int64 `protobuf:"varint,7,opt,name=QuickRound,proto3" json:"QuickRound,omitempty"`
}

func (m *Instance81Data) Reset()         { *m = Instance81Data{} }
func (m *Instance81Data) String() string { return proto.CompactTextString(m) }
func (*Instance81Data) ProtoMessage()    {}

//
type Player81Data struct {
	//
	InstanceId int64 `protobuf:"varint,1,opt,name=InstanceId,proto3" json:"InstanceId,omitempty"`
	//
	State int64 `protobuf:"varint,2,opt,name=State,proto3" json:"State,omitempty"`
	//
	FightTimes int64 `protobuf:"varint,3,opt,name=FightTimes,proto3" json:"FightTimes,omitempty"`
	//
	IsBoxOpen int64 `protobuf:"varint,4,opt,name=IsBoxOpen,proto3" json:"IsBoxOpen,omitempty"`
}

func (m *Player81Data) Reset()         { *m = Player81Data{} }
func (m *Player81Data) String() string { return proto.CompactTextString(m) }
func (*Player81Data) ProtoMessage()    {}

const PCK_S2C81Info_ID = 1901 //
//
type S2C81Info struct {
	//
	Data []*Player81Data `protobuf:"bytes,1,rep,name=Data,proto3" json:"Data,omitempty"`
	//
	HelpTimes int64 `protobuf:"varint,2,opt,name=HelpTimes,proto3" json:"HelpTimes,omitempty"`
}

func (m *S2C81Info) Reset()         { *m = S2C81Info{} }
func (m *S2C81Info) String() string { return proto.CompactTextString(m) }
func (*S2C81Info) ProtoMessage()    {}
func (m *S2C81Info) GetId() uint16  { return PCK_S2C81Info_ID }

const PCK_C2SViewInstance81Data_ID = 1906 //
//
type C2SViewInstance81Data struct {
	//
	InstanceId int64 `protobuf:"varint,1,opt,name=InstanceId,proto3" json:"InstanceId,omitempty"`
}

func (m *C2SViewInstance81Data) Reset()         { *m = C2SViewInstance81Data{} }
func (m *C2SViewInstance81Data) String() string { return proto.CompactTextString(m) }
func (*C2SViewInstance81Data) ProtoMessage()    {}
func (m *C2SViewInstance81Data) GetId() uint16  { return PCK_C2SViewInstance81Data_ID }

const PCK_S2CViewInstance81Data_ID = 1907 //
//
type S2CViewInstance81Data struct {
	//
	InstanceId int64 `protobuf:"varint,1,opt,name=InstanceId,proto3" json:"InstanceId,omitempty"`
	//
	Instance *Instance81Data `protobuf:"bytes,2,opt,name=Instance,proto3" json:"Instance,omitempty"`
}

func (m *S2CViewInstance81Data) Reset()         { *m = S2CViewInstance81Data{} }
func (m *S2CViewInstance81Data) String() string { return proto.CompactTextString(m) }
func (*S2CViewInstance81Data) ProtoMessage()    {}
func (m *S2CViewInstance81Data) GetId() uint16  { return PCK_S2CViewInstance81Data_ID }

const PCK_C2S81Sweep_ID = 1902 //
//
type C2S81Sweep struct {
	//
	InstanceId int64 `protobuf:"varint,1,opt,name=InstanceId,proto3" json:"InstanceId,omitempty"`
	//
	isSweepAll int32 `protobuf:"varint,2,opt,name=isSweepAll,proto3" json:"isSweepAll,omitempty"`
}

func (m *C2S81Sweep) Reset()         { *m = C2S81Sweep{} }
func (m *C2S81Sweep) String() string { return proto.CompactTextString(m) }
func (*C2S81Sweep) ProtoMessage()    {}
func (m *C2S81Sweep) GetId() uint16  { return PCK_C2S81Sweep_ID }

const PCK_S2C81Sweep_ID = 1903 //
//
type S2C81Sweep struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	InstanceId []int64 `protobuf:"varint,2,rep,name=InstanceId,proto3" json:"InstanceId,omitempty"`
	//
	Items []*ItemData `protobuf:"bytes,3,rep,name=Items,proto3" json:"Items,omitempty"`
}

func (m *S2C81Sweep) Reset()         { *m = S2C81Sweep{} }
func (m *S2C81Sweep) String() string { return proto.CompactTextString(m) }
func (*S2C81Sweep) ProtoMessage()    {}
func (m *S2C81Sweep) GetId() uint16  { return PCK_S2C81Sweep_ID }

const PCK_C2S81BuyBox_ID = 1904 //
//
type C2S81BuyBox struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Chapter int32 `protobuf:"varint,2,opt,name=Chapter,proto3" json:"Chapter,omitempty"`
}

func (m *C2S81BuyBox) Reset()         { *m = C2S81BuyBox{} }
func (m *C2S81BuyBox) String() string { return proto.CompactTextString(m) }
func (*C2S81BuyBox) ProtoMessage()    {}
func (m *C2S81BuyBox) GetId() uint16  { return PCK_C2S81BuyBox_ID }

const PCK_S2C81BuyBox_ID = 1905 //
//
type S2C81BuyBox struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	id int32 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	//
	Chapter int32 `protobuf:"varint,3,opt,name=Chapter,proto3" json:"Chapter,omitempty"`
	//
	Items []*ItemData `protobuf:"bytes,4,rep,name=Items,proto3" json:"Items,omitempty"`
}

func (m *S2C81BuyBox) Reset()         { *m = S2C81BuyBox{} }
func (m *S2C81BuyBox) String() string { return proto.CompactTextString(m) }
func (*S2C81BuyBox) ProtoMessage()    {}
func (m *S2C81BuyBox) GetId() uint16  { return PCK_S2C81BuyBox_ID }

const PCK_C2SKingInfo_ID = 14014 //
//
type C2SKingInfo struct {
}

func (m *C2SKingInfo) Reset()         { *m = C2SKingInfo{} }
func (m *C2SKingInfo) String() string { return proto.CompactTextString(m) }
func (*C2SKingInfo) ProtoMessage()    {}
func (m *C2SKingInfo) GetId() uint16  { return PCK_C2SKingInfo_ID }

const PCK_S2CKingInfo_ID = 14015 //
//
type S2CKingInfo struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	KingRank int32 `protobuf:"varint,2,opt,name=KingRank,proto3" json:"KingRank,omitempty"`
	//
	Win int32 `protobuf:"varint,3,opt,name=Win,proto3" json:"Win,omitempty"`
	//
	ContinueWin int32 `protobuf:"varint,4,opt,name=ContinueWin,proto3" json:"ContinueWin,omitempty"`
	//
	Respect int32 `protobuf:"varint,5,opt,name=Respect,proto3" json:"Respect,omitempty"`
	//
	Times int32 `protobuf:"varint,6,opt,name=Times,proto3" json:"Times,omitempty"`
	//
	NextTime int32 `protobuf:"varint,7,opt,name=NextTime,proto3" json:"NextTime,omitempty"`
}

func (m *S2CKingInfo) Reset()         { *m = S2CKingInfo{} }
func (m *S2CKingInfo) String() string { return proto.CompactTextString(m) }
func (*S2CKingInfo) ProtoMessage()    {}
func (m *S2CKingInfo) GetId() uint16  { return PCK_S2CKingInfo_ID }

const PCK_C2SKingState_ID = 14000 //
//
type C2SKingState struct {
}

func (m *C2SKingState) Reset()         { *m = C2SKingState{} }
func (m *C2SKingState) String() string { return proto.CompactTextString(m) }
func (*C2SKingState) ProtoMessage()    {}
func (m *C2SKingState) GetId() uint16  { return PCK_C2SKingState_ID }

const PCK_S2CKingState_ID = 14001 //
//
type S2CKingState struct {
	//
	State int32 `protobuf:"varint,1,opt,name=State,proto3" json:"State,omitempty"`
}

func (m *S2CKingState) Reset()         { *m = S2CKingState{} }
func (m *S2CKingState) String() string { return proto.CompactTextString(m) }
func (*S2CKingState) ProtoMessage()    {}
func (m *S2CKingState) GetId() uint16  { return PCK_S2CKingState_ID }

//
type KingMatch struct {
	//
	UserId int32 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	ServerId int32 `protobuf:"varint,2,opt,name=ServerId,proto3" json:"ServerId,omitempty"`
	//
	RoleId int32 `protobuf:"varint,3,opt,name=RoleId,proto3" json:"RoleId,omitempty"`
	//
	Nick string `protobuf:"bytes,4,opt,name=Nick,proto3" json:"Nick,omitempty"`
	//
	Rank int32 `protobuf:"varint,5,opt,name=Rank,proto3" json:"Rank,omitempty"`
}

func (m *KingMatch) Reset()         { *m = KingMatch{} }
func (m *KingMatch) String() string { return proto.CompactTextString(m) }
func (*KingMatch) ProtoMessage()    {}

const PCK_C2SKingMatch_ID = 14002 //
//
type C2SKingMatch struct {
}

func (m *C2SKingMatch) Reset()         { *m = C2SKingMatch{} }
func (m *C2SKingMatch) String() string { return proto.CompactTextString(m) }
func (*C2SKingMatch) ProtoMessage()    {}
func (m *C2SKingMatch) GetId() uint16  { return PCK_C2SKingMatch_ID }

const PCK_S2CKingMatch_ID = 14003 //
//
type S2CKingMatch struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	MatchList []*KingMatch `protobuf:"bytes,2,rep,name=MatchList,proto3" json:"MatchList,omitempty"`
	//
	Target *KingMatch `protobuf:"bytes,3,opt,name=Target,proto3" json:"Target,omitempty"`
}

func (m *S2CKingMatch) Reset()         { *m = S2CKingMatch{} }
func (m *S2CKingMatch) String() string { return proto.CompactTextString(m) }
func (*S2CKingMatch) ProtoMessage()    {}
func (m *S2CKingMatch) GetId() uint16  { return PCK_S2CKingMatch_ID }

const PCK_C2SKingFight_ID = 14004 //
//
type C2SKingFight struct {
}

func (m *C2SKingFight) Reset()         { *m = C2SKingFight{} }
func (m *C2SKingFight) String() string { return proto.CompactTextString(m) }
func (*C2SKingFight) ProtoMessage()    {}
func (m *C2SKingFight) GetId() uint16  { return PCK_C2SKingFight_ID }

const PCK_S2CKingFight_ID = 14005 //
//
type S2CKingFight struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CKingFight) Reset()         { *m = S2CKingFight{} }
func (m *S2CKingFight) String() string { return proto.CompactTextString(m) }
func (*S2CKingFight) ProtoMessage()    {}
func (m *S2CKingFight) GetId() uint16  { return PCK_S2CKingFight_ID }

//
type KingRank struct {
	//
	KingRank int32 `protobuf:"varint,1,opt,name=KingRank,proto3" json:"KingRank,omitempty"`
	//
	ServerId int32 `protobuf:"varint,2,opt,name=ServerId,proto3" json:"ServerId,omitempty"`
	//
	Nick string `protobuf:"bytes,3,opt,name=Nick,proto3" json:"Nick,omitempty"`
	//
	Win int32 `protobuf:"varint,4,opt,name=Win,proto3" json:"Win,omitempty"`
	//
	Sort int32 `protobuf:"varint,5,opt,name=Sort,proto3" json:"Sort,omitempty"`
}

func (m *KingRank) Reset()         { *m = KingRank{} }
func (m *KingRank) String() string { return proto.CompactTextString(m) }
func (*KingRank) ProtoMessage()    {}

const PCK_C2SKingRank_ID = 14006 //
//
type C2SKingRank struct {
	//
	Type int32 `protobuf:"varint,1,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (m *C2SKingRank) Reset()         { *m = C2SKingRank{} }
func (m *C2SKingRank) String() string { return proto.CompactTextString(m) }
func (*C2SKingRank) ProtoMessage()    {}
func (m *C2SKingRank) GetId() uint16  { return PCK_C2SKingRank_ID }

const PCK_S2CKingRank_ID = 14007 //
//
type S2CKingRank struct {
	//
	Type int32 `protobuf:"varint,1,opt,name=Type,proto3" json:"Type,omitempty"`
	//
	Rank []*KingRank `protobuf:"bytes,2,rep,name=Rank,proto3" json:"Rank,omitempty"`
	//
	MyRank *KingRank `protobuf:"bytes,3,opt,name=MyRank,proto3" json:"MyRank,omitempty"`
}

func (m *S2CKingRank) Reset()         { *m = S2CKingRank{} }
func (m *S2CKingRank) String() string { return proto.CompactTextString(m) }
func (*S2CKingRank) ProtoMessage()    {}
func (m *S2CKingRank) GetId() uint16  { return PCK_S2CKingRank_ID }

const PCK_C2SKingPlayer_ID = 14010 //
//
type C2SKingPlayer struct {
}

func (m *C2SKingPlayer) Reset()         { *m = C2SKingPlayer{} }
func (m *C2SKingPlayer) String() string { return proto.CompactTextString(m) }
func (*C2SKingPlayer) ProtoMessage()    {}
func (m *C2SKingPlayer) GetId() uint16  { return PCK_C2SKingPlayer_ID }

const PCK_S2CKingPlayer_ID = 14011 //
//
type S2CKingPlayer struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	ServerId int32 `protobuf:"varint,2,opt,name=ServerId,proto3" json:"ServerId,omitempty"`
	//
	A []*IntAttr `protobuf:"bytes,3,rep,name=A,proto3" json:"A,omitempty"`
	//
	B []*StrAttr `protobuf:"bytes,4,rep,name=B,proto3" json:"B,omitempty"`
}

func (m *S2CKingPlayer) Reset()         { *m = S2CKingPlayer{} }
func (m *S2CKingPlayer) String() string { return proto.CompactTextString(m) }
func (*S2CKingPlayer) ProtoMessage()    {}
func (m *S2CKingPlayer) GetId() uint16  { return PCK_S2CKingPlayer_ID }

const PCK_C2SKingGetBuyInfo_ID = 14018 //
//
type C2SKingGetBuyInfo struct {
}

func (m *C2SKingGetBuyInfo) Reset()         { *m = C2SKingGetBuyInfo{} }
func (m *C2SKingGetBuyInfo) String() string { return proto.CompactTextString(m) }
func (*C2SKingGetBuyInfo) ProtoMessage()    {}
func (m *C2SKingGetBuyInfo) GetId() uint16  { return PCK_C2SKingGetBuyInfo_ID }

const PCK_S2CKingGetBuyInfo_ID = 14019 //
//
type S2CKingGetBuyInfo struct {
	//
	Coin3 int32 `protobuf:"varint,1,opt,name=Coin3,proto3" json:"Coin3,omitempty"`
	//
	Times int32 `protobuf:"varint,2,opt,name=Times,proto3" json:"Times,omitempty"`
	//
	TotleTimes int32 `protobuf:"varint,3,opt,name=TotleTimes,proto3" json:"TotleTimes,omitempty"`
}

func (m *S2CKingGetBuyInfo) Reset()         { *m = S2CKingGetBuyInfo{} }
func (m *S2CKingGetBuyInfo) String() string { return proto.CompactTextString(m) }
func (*S2CKingGetBuyInfo) ProtoMessage()    {}
func (m *S2CKingGetBuyInfo) GetId() uint16  { return PCK_S2CKingGetBuyInfo_ID }

const PCK_C2SKingBuyTimes_ID = 14012 //
//
type C2SKingBuyTimes struct {
}

func (m *C2SKingBuyTimes) Reset()         { *m = C2SKingBuyTimes{} }
func (m *C2SKingBuyTimes) String() string { return proto.CompactTextString(m) }
func (*C2SKingBuyTimes) ProtoMessage()    {}
func (m *C2SKingBuyTimes) GetId() uint16  { return PCK_C2SKingBuyTimes_ID }

const PCK_S2CKingBuyTimes_ID = 14013 //
//
type S2CKingBuyTimes struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CKingBuyTimes) Reset()         { *m = S2CKingBuyTimes{} }
func (m *S2CKingBuyTimes) String() string { return proto.CompactTextString(m) }
func (*S2CKingBuyTimes) ProtoMessage()    {}
func (m *S2CKingBuyTimes) GetId() uint16  { return PCK_S2CKingBuyTimes_ID }

const PCK_C2SKingRespect_ID = 14016 //
//
type C2SKingRespect struct {
}

func (m *C2SKingRespect) Reset()         { *m = C2SKingRespect{} }
func (m *C2SKingRespect) String() string { return proto.CompactTextString(m) }
func (*C2SKingRespect) ProtoMessage()    {}
func (m *C2SKingRespect) GetId() uint16  { return PCK_C2SKingRespect_ID }

const PCK_S2CKingRespect_ID = 14017 //
//
type S2CKingRespect struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CKingRespect) Reset()         { *m = S2CKingRespect{} }
func (m *S2CKingRespect) String() string { return proto.CompactTextString(m) }
func (*S2CKingRespect) ProtoMessage()    {}
func (m *S2CKingRespect) GetId() uint16  { return PCK_S2CKingRespect_ID }

//
type WorldBossRankData struct {
	//
	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	Nick string `protobuf:"bytes,2,opt,name=Nick,proto3" json:"Nick,omitempty"`
	//
	RoleId int32 `protobuf:"varint,3,opt,name=RoleId,proto3" json:"RoleId,omitempty"`
	//
	Damage int64 `protobuf:"varint,4,opt,name=Damage,proto3" json:"Damage,omitempty"`
	//
	AreaId int64 `protobuf:"varint,5,opt,name=AreaId,proto3" json:"AreaId,omitempty"`
}

func (m *WorldBossRankData) Reset()         { *m = WorldBossRankData{} }
func (m *WorldBossRankData) String() string { return proto.CompactTextString(m) }
func (*WorldBossRankData) ProtoMessage()    {}

const PCK_S2CWorldBossRank5_ID = 15001 //
//
type S2CWorldBossRank5 struct {
	//
	Rank []*WorldBossRankData `protobuf:"bytes,1,rep,name=Rank,proto3" json:"Rank,omitempty"`
}

func (m *S2CWorldBossRank5) Reset()         { *m = S2CWorldBossRank5{} }
func (m *S2CWorldBossRank5) String() string { return proto.CompactTextString(m) }
func (*S2CWorldBossRank5) ProtoMessage()    {}
func (m *S2CWorldBossRank5) GetId() uint16  { return PCK_S2CWorldBossRank5_ID }

const PCK_S2CActWorldBossSettlement_ID = 15002 //
//
type S2CActWorldBossSettlement struct {
	//
	KillNick string `protobuf:"bytes,1,opt,name=KillNick,proto3" json:"KillNick,omitempty"`
	//
	KillGang string `protobuf:"bytes,2,opt,name=KillGang,proto3" json:"KillGang,omitempty"`
	//
	KillAreaId int32 `protobuf:"varint,3,opt,name=KillAreaId,proto3" json:"KillAreaId,omitempty"`
	//
	FirstNick string `protobuf:"bytes,4,opt,name=FirstNick,proto3" json:"FirstNick,omitempty"`
	//
	FirstGang string `protobuf:"bytes,5,opt,name=FirstGang,proto3" json:"FirstGang,omitempty"`
	//
	FirstAreaId int32 `protobuf:"varint,6,opt,name=FirstAreaId,proto3" json:"FirstAreaId,omitempty"`
	//
	Rank int32 `protobuf:"varint,7,opt,name=Rank,proto3" json:"Rank,omitempty"`
	//
	IsKill int32 `protobuf:"varint,8,opt,name=IsKill,proto3" json:"IsKill,omitempty"`
	//
	KillRoleId int32 `protobuf:"varint,9,opt,name=KillRoleId,proto3" json:"KillRoleId,omitempty"`
	//
	FirstRoleId int32 `protobuf:"varint,10,opt,name=FirstRoleId,proto3" json:"FirstRoleId,omitempty"`
}

func (m *S2CActWorldBossSettlement) Reset()         { *m = S2CActWorldBossSettlement{} }
func (m *S2CActWorldBossSettlement) String() string { return proto.CompactTextString(m) }
func (*S2CActWorldBossSettlement) ProtoMessage()    {}
func (m *S2CActWorldBossSettlement) GetId() uint16  { return PCK_S2CActWorldBossSettlement_ID }

//
type WorldBossKillData struct {
	//
	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	Nick string `protobuf:"bytes,2,opt,name=Nick,proto3" json:"Nick,omitempty"`
	//
	GangName string `protobuf:"bytes,3,opt,name=GangName,proto3" json:"GangName,omitempty"`
	//
	AreaId int64 `protobuf:"varint,4,opt,name=AreaId,proto3" json:"AreaId,omitempty"`
}

func (m *WorldBossKillData) Reset()         { *m = WorldBossKillData{} }
func (m *WorldBossKillData) String() string { return proto.CompactTextString(m) }
func (*WorldBossKillData) ProtoMessage()    {}

const PCK_K2SWorldBossKillData_ID = 15003 //
//
type K2SWorldBossKillData struct {
	//
	Data *WorldBossKillData `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *K2SWorldBossKillData) Reset()         { *m = K2SWorldBossKillData{} }
func (m *K2SWorldBossKillData) String() string { return proto.CompactTextString(m) }
func (*K2SWorldBossKillData) ProtoMessage()    {}
func (m *K2SWorldBossKillData) GetId() uint16  { return PCK_K2SWorldBossKillData_ID }

const PCK_C2SGetWorldBossKillData_ID = 15004 //
//
type C2SGetWorldBossKillData struct {
}

func (m *C2SGetWorldBossKillData) Reset()         { *m = C2SGetWorldBossKillData{} }
func (m *C2SGetWorldBossKillData) String() string { return proto.CompactTextString(m) }
func (*C2SGetWorldBossKillData) ProtoMessage()    {}
func (m *C2SGetWorldBossKillData) GetId() uint16  { return PCK_C2SGetWorldBossKillData_ID }

const PCK_S2CGetWorldBossKillData_ID = 15006 //
//
type S2CGetWorldBossKillData struct {
	//
	Data *WorldBossKillData `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *S2CGetWorldBossKillData) Reset()         { *m = S2CGetWorldBossKillData{} }
func (m *S2CGetWorldBossKillData) String() string { return proto.CompactTextString(m) }
func (*S2CGetWorldBossKillData) ProtoMessage()    {}
func (m *S2CGetWorldBossKillData) GetId() uint16  { return PCK_S2CGetWorldBossKillData_ID }

const PCK_S2CBossRefreshTime_ID = 15007 //
//
type S2CBossRefreshTime struct {
	//
	RefreshTime int64 `protobuf:"varint,1,opt,name=RefreshTime,proto3" json:"RefreshTime,omitempty"`
}

func (m *S2CBossRefreshTime) Reset()         { *m = S2CBossRefreshTime{} }
func (m *S2CBossRefreshTime) String() string { return proto.CompactTextString(m) }
func (*S2CBossRefreshTime) ProtoMessage()    {}
func (m *S2CBossRefreshTime) GetId() uint16  { return PCK_S2CBossRefreshTime_ID }

const PCK_C2SGetWorldBossRank_ID = 15008 //
//
type C2SGetWorldBossRank struct {
	//
	Num int32 `protobuf:"varint,1,opt,name=Num,proto3" json:"Num,omitempty"`
}

func (m *C2SGetWorldBossRank) Reset()         { *m = C2SGetWorldBossRank{} }
func (m *C2SGetWorldBossRank) String() string { return proto.CompactTextString(m) }
func (*C2SGetWorldBossRank) ProtoMessage()    {}
func (m *C2SGetWorldBossRank) GetId() uint16  { return PCK_C2SGetWorldBossRank_ID }

const PCK_S2CGetWorldBossRank_ID = 15009 //
//
type S2CGetWorldBossRank struct {
	//
	Rank []*WorldBossRankData `protobuf:"bytes,1,rep,name=Rank,proto3" json:"Rank,omitempty"`
	//
	SelfRank int32 `protobuf:"varint,2,opt,name=SelfRank,proto3" json:"SelfRank,omitempty"`
	//
	SelfDamage int64 `protobuf:"varint,3,opt,name=SelfDamage,proto3" json:"SelfDamage,omitempty"`
}

func (m *S2CGetWorldBossRank) Reset()         { *m = S2CGetWorldBossRank{} }
func (m *S2CGetWorldBossRank) String() string { return proto.CompactTextString(m) }
func (*S2CGetWorldBossRank) ProtoMessage()    {}
func (m *S2CGetWorldBossRank) GetId() uint16  { return PCK_S2CGetWorldBossRank_ID }

//
type ExpeditionTeamMember struct {
	//
	UnitType int32 `protobuf:"varint,1,opt,name=UnitType,proto3" json:"UnitType,omitempty"`
	//
	ObjId int64 `protobuf:"varint,2,opt,name=ObjId,proto3" json:"ObjId,omitempty"`
	//
	Nick string `protobuf:"bytes,3,opt,name=Nick,proto3" json:"Nick,omitempty"`
	//
	CurrHp int64 `protobuf:"varint,4,opt,name=CurrHp,proto3" json:"CurrHp,omitempty"`
	//
	MaxHp int64 `protobuf:"varint,5,opt,name=MaxHp,proto3" json:"MaxHp,omitempty"`
	//
	RoleId int32 `protobuf:"varint,6,opt,name=RoleId,proto3" json:"RoleId,omitempty"`
	//
	Pos int32 `protobuf:"varint,7,opt,name=Pos,proto3" json:"Pos,omitempty"`
	//
	Level int32 `protobuf:"varint,8,opt,name=Level,proto3" json:"Level,omitempty"`
	//
	FightValue int64 `protobuf:"varint,9,opt,name=FightValue,proto3" json:"FightValue,omitempty"`
	//
	State int32 `protobuf:"varint,10,opt,name=State,proto3" json:"State,omitempty"`
}

func (m *ExpeditionTeamMember) Reset()         { *m = ExpeditionTeamMember{} }
func (m *ExpeditionTeamMember) String() string { return proto.CompactTextString(m) }
func (*ExpeditionTeamMember) ProtoMessage()    {}

const PCK_C2SExpeditionEnemy_ID = 2001 //
//
type C2SExpeditionEnemy struct {
}

func (m *C2SExpeditionEnemy) Reset()         { *m = C2SExpeditionEnemy{} }
func (m *C2SExpeditionEnemy) String() string { return proto.CompactTextString(m) }
func (*C2SExpeditionEnemy) ProtoMessage()    {}
func (m *C2SExpeditionEnemy) GetId() uint16  { return PCK_C2SExpeditionEnemy_ID }

const PCK_S2CExpeditionEnemy_ID = 2002 //
//
type S2CExpeditionEnemy struct {
	//
	Items []*ExpeditionTeamMember `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty"`
}

func (m *S2CExpeditionEnemy) Reset()         { *m = S2CExpeditionEnemy{} }
func (m *S2CExpeditionEnemy) String() string { return proto.CompactTextString(m) }
func (*S2CExpeditionEnemy) ProtoMessage()    {}
func (m *S2CExpeditionEnemy) GetId() uint16  { return PCK_S2CExpeditionEnemy_ID }

const PCK_C2SExpeditionTeam_ID = 2003 //
//
type C2SExpeditionTeam struct {
}

func (m *C2SExpeditionTeam) Reset()         { *m = C2SExpeditionTeam{} }
func (m *C2SExpeditionTeam) String() string { return proto.CompactTextString(m) }
func (*C2SExpeditionTeam) ProtoMessage()    {}
func (m *C2SExpeditionTeam) GetId() uint16  { return PCK_C2SExpeditionTeam_ID }

const PCK_S2CExpeditionTeam_ID = 2004 //
//
type S2CExpeditionTeam struct {
	//
	Items []*ExpeditionTeamMember `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty"`
}

func (m *S2CExpeditionTeam) Reset()         { *m = S2CExpeditionTeam{} }
func (m *S2CExpeditionTeam) String() string { return proto.CompactTextString(m) }
func (*S2CExpeditionTeam) ProtoMessage()    {}
func (m *S2CExpeditionTeam) GetId() uint16  { return PCK_S2CExpeditionTeam_ID }

const PCK_S2CExpeditionInfo_ID = 2005 //
//
type S2CExpeditionInfo struct {
	//
	CurrId int32 `protobuf:"varint,1,opt,name=CurrId,proto3" json:"CurrId,omitempty"`
	//
	BoxId []int32 `protobuf:"varint,2,rep,name=BoxId,proto3" json:"BoxId,omitempty"`
	//
	ResetTimes int32 `protobuf:"varint,3,opt,name=ResetTimes,proto3" json:"ResetTimes,omitempty"`
	//
	ResetMaxTimes int32 `protobuf:"varint,4,opt,name=ResetMaxTimes,proto3" json:"ResetMaxTimes,omitempty"`
	//
	Times int32 `protobuf:"varint,5,opt,name=Times,proto3" json:"Times,omitempty"`
	//
	Buff1 *ExpeditionBuff `protobuf:"bytes,6,opt,name=Buff1,proto3" json:"Buff1,omitempty"`
	//
	Buff2 *ExpeditionBuff `protobuf:"bytes,7,opt,name=Buff2,proto3" json:"Buff2,omitempty"`
	//
	GoodsId int32 `protobuf:"varint,8,opt,name=GoodsId,proto3" json:"GoodsId,omitempty"`
	//
	CanSweep int32 `protobuf:"varint,9,opt,name=CanSweep,proto3" json:"CanSweep,omitempty"`
}

func (m *S2CExpeditionInfo) Reset()         { *m = S2CExpeditionInfo{} }
func (m *S2CExpeditionInfo) String() string { return proto.CompactTextString(m) }
func (*S2CExpeditionInfo) ProtoMessage()    {}
func (m *S2CExpeditionInfo) GetId() uint16  { return PCK_S2CExpeditionInfo_ID }

const PCK_C2SExpeditionStartFight_ID = 2006 //
//
type C2SExpeditionStartFight struct {
}

func (m *C2SExpeditionStartFight) Reset()         { *m = C2SExpeditionStartFight{} }
func (m *C2SExpeditionStartFight) String() string { return proto.CompactTextString(m) }
func (*C2SExpeditionStartFight) ProtoMessage()    {}
func (m *C2SExpeditionStartFight) GetId() uint16  { return PCK_C2SExpeditionStartFight_ID }

const PCK_S2CExpeditionStartFight_ID = 2007 //
//
type S2CExpeditionStartFight struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Win int32 `protobuf:"varint,2,opt,name=Win,proto3" json:"Win,omitempty"`
}

func (m *S2CExpeditionStartFight) Reset()         { *m = S2CExpeditionStartFight{} }
func (m *S2CExpeditionStartFight) String() string { return proto.CompactTextString(m) }
func (*S2CExpeditionStartFight) ProtoMessage()    {}
func (m *S2CExpeditionStartFight) GetId() uint16  { return PCK_S2CExpeditionStartFight_ID }

const PCK_C2SExpeditionGetPrize_ID = 2008 //
//
type C2SExpeditionGetPrize struct {
	//
	BoxId int32 `protobuf:"varint,1,opt,name=BoxId,proto3" json:"BoxId,omitempty"`
	//
	Param int32 `protobuf:"varint,2,opt,name=Param,proto3" json:"Param,omitempty"`
}

func (m *C2SExpeditionGetPrize) Reset()         { *m = C2SExpeditionGetPrize{} }
func (m *C2SExpeditionGetPrize) String() string { return proto.CompactTextString(m) }
func (*C2SExpeditionGetPrize) ProtoMessage()    {}
func (m *C2SExpeditionGetPrize) GetId() uint16  { return PCK_C2SExpeditionGetPrize_ID }

const PCK_S2CExpeditionGetPrize_ID = 2009 //
//
type S2CExpeditionGetPrize struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	BoxId int32 `protobuf:"varint,2,opt,name=BoxId,proto3" json:"BoxId,omitempty"`
	//
	Param int32 `protobuf:"varint,3,opt,name=Param,proto3" json:"Param,omitempty"`
}

func (m *S2CExpeditionGetPrize) Reset()         { *m = S2CExpeditionGetPrize{} }
func (m *S2CExpeditionGetPrize) String() string { return proto.CompactTextString(m) }
func (*S2CExpeditionGetPrize) ProtoMessage()    {}
func (m *S2CExpeditionGetPrize) GetId() uint16  { return PCK_S2CExpeditionGetPrize_ID }

const PCK_C2SExpeditionRelive_ID = 2010 //
//
type C2SExpeditionRelive struct {
}

func (m *C2SExpeditionRelive) Reset()         { *m = C2SExpeditionRelive{} }
func (m *C2SExpeditionRelive) String() string { return proto.CompactTextString(m) }
func (*C2SExpeditionRelive) ProtoMessage()    {}
func (m *C2SExpeditionRelive) GetId() uint16  { return PCK_C2SExpeditionRelive_ID }

const PCK_S2CExpeditionRelive_ID = 2011 //
//
type S2CExpeditionRelive struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CExpeditionRelive) Reset()         { *m = S2CExpeditionRelive{} }
func (m *S2CExpeditionRelive) String() string { return proto.CompactTextString(m) }
func (*S2CExpeditionRelive) ProtoMessage()    {}
func (m *S2CExpeditionRelive) GetId() uint16  { return PCK_S2CExpeditionRelive_ID }

const PCK_C2SExpeditionReset_ID = 2014 //
//
type C2SExpeditionReset struct {
}

func (m *C2SExpeditionReset) Reset()         { *m = C2SExpeditionReset{} }
func (m *C2SExpeditionReset) String() string { return proto.CompactTextString(m) }
func (*C2SExpeditionReset) ProtoMessage()    {}
func (m *C2SExpeditionReset) GetId() uint16  { return PCK_C2SExpeditionReset_ID }

const PCK_S2CExpeditionReset_ID = 2015 //
//
type S2CExpeditionReset struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CExpeditionReset) Reset()         { *m = S2CExpeditionReset{} }
func (m *S2CExpeditionReset) String() string { return proto.CompactTextString(m) }
func (*S2CExpeditionReset) ProtoMessage()    {}
func (m *S2CExpeditionReset) GetId() uint16  { return PCK_S2CExpeditionReset_ID }

const PCK_C2SExpeditionSaveTeam_ID = 2012 //
//
type C2SExpeditionSaveTeam struct {
	//
	UnitType int32 `protobuf:"varint,1,opt,name=UnitType,proto3" json:"UnitType,omitempty"`
	//
	Ids []int64 `protobuf:"varint,2,rep,name=Ids,proto3" json:"Ids,omitempty"`
}

func (m *C2SExpeditionSaveTeam) Reset()         { *m = C2SExpeditionSaveTeam{} }
func (m *C2SExpeditionSaveTeam) String() string { return proto.CompactTextString(m) }
func (*C2SExpeditionSaveTeam) ProtoMessage()    {}
func (m *C2SExpeditionSaveTeam) GetId() uint16  { return PCK_C2SExpeditionSaveTeam_ID }

const PCK_S2CExpeditionSaveTeam_ID = 2013 //
//
type S2CExpeditionSaveTeam struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Items []*ExpeditionTeamMember `protobuf:"bytes,2,rep,name=Items,proto3" json:"Items,omitempty"`
}

func (m *S2CExpeditionSaveTeam) Reset()         { *m = S2CExpeditionSaveTeam{} }
func (m *S2CExpeditionSaveTeam) String() string { return proto.CompactTextString(m) }
func (*S2CExpeditionSaveTeam) ProtoMessage()    {}
func (m *S2CExpeditionSaveTeam) GetId() uint16  { return PCK_S2CExpeditionSaveTeam_ID }

//
type ExpeditionBuff struct {
	//
	Attr []int32 `protobuf:"varint,1,rep,name=Attr,proto3" json:"Attr,omitempty"`
	//
	Use int32 `protobuf:"varint,2,opt,name=Use,proto3" json:"Use,omitempty"`
}

func (m *ExpeditionBuff) Reset()         { *m = ExpeditionBuff{} }
func (m *ExpeditionBuff) String() string { return proto.CompactTextString(m) }
func (*ExpeditionBuff) ProtoMessage()    {}

const PCK_C2SExpeditionRefreshBuff_ID = 2016 //
//
type C2SExpeditionRefreshBuff struct {
	//
	BoxId int32 `protobuf:"varint,1,opt,name=BoxId,proto3" json:"BoxId,omitempty"`
}

func (m *C2SExpeditionRefreshBuff) Reset()         { *m = C2SExpeditionRefreshBuff{} }
func (m *C2SExpeditionRefreshBuff) String() string { return proto.CompactTextString(m) }
func (*C2SExpeditionRefreshBuff) ProtoMessage()    {}
func (m *C2SExpeditionRefreshBuff) GetId() uint16  { return PCK_C2SExpeditionRefreshBuff_ID }

const PCK_S2CExpeditionRefreshBuff_ID = 2017 //
//
type S2CExpeditionRefreshBuff struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	BoxId int32 `protobuf:"varint,2,opt,name=BoxId,proto3" json:"BoxId,omitempty"`
	//
	Buff *ExpeditionBuff `protobuf:"bytes,3,opt,name=Buff,proto3" json:"Buff,omitempty"`
}

func (m *S2CExpeditionRefreshBuff) Reset()         { *m = S2CExpeditionRefreshBuff{} }
func (m *S2CExpeditionRefreshBuff) String() string { return proto.CompactTextString(m) }
func (*S2CExpeditionRefreshBuff) ProtoMessage()    {}
func (m *S2CExpeditionRefreshBuff) GetId() uint16  { return PCK_S2CExpeditionRefreshBuff_ID }

const PCK_C2SExpeditionSweep_ID = 2018 //
//
type C2SExpeditionSweep struct {
}

func (m *C2SExpeditionSweep) Reset()         { *m = C2SExpeditionSweep{} }
func (m *C2SExpeditionSweep) String() string { return proto.CompactTextString(m) }
func (*C2SExpeditionSweep) ProtoMessage()    {}
func (m *C2SExpeditionSweep) GetId() uint16  { return PCK_C2SExpeditionSweep_ID }

const PCK_S2CExpeditionSweep_ID = 2019 //
//
type S2CExpeditionSweep struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CExpeditionSweep) Reset()         { *m = S2CExpeditionSweep{} }
func (m *S2CExpeditionSweep) String() string { return proto.CompactTextString(m) }
func (*S2CExpeditionSweep) ProtoMessage()    {}
func (m *S2CExpeditionSweep) GetId() uint16  { return PCK_S2CExpeditionSweep_ID }

const PCK_C2SWordsWear_ID = 16001 //
//
type C2SWordsWear struct {
	//
	Part int32 `protobuf:"varint,1,opt,name=Part,proto3" json:"Part,omitempty"`
}

func (m *C2SWordsWear) Reset()         { *m = C2SWordsWear{} }
func (m *C2SWordsWear) String() string { return proto.CompactTextString(m) }
func (*C2SWordsWear) ProtoMessage()    {}
func (m *C2SWordsWear) GetId() uint16  { return PCK_C2SWordsWear_ID }

const PCK_S2CWordsWear_ID = 16002 //
//
type S2CWordsWear struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Parts []int32 `protobuf:"varint,2,rep,name=Parts,proto3" json:"Parts,omitempty"`
}

func (m *S2CWordsWear) Reset()         { *m = S2CWordsWear{} }
func (m *S2CWordsWear) String() string { return proto.CompactTextString(m) }
func (*S2CWordsWear) ProtoMessage()    {}
func (m *S2CWordsWear) GetId() uint16  { return PCK_S2CWordsWear_ID }

const PCK_C2SWordsStick_ID = 16003 //
//
type C2SWordsStick struct {
	//
	Part int32 `protobuf:"varint,1,opt,name=Part,proto3" json:"Part,omitempty"`
	//
	Times int32 `protobuf:"varint,2,opt,name=Times,proto3" json:"Times,omitempty"`
	//
	Level int32 `protobuf:"varint,3,opt,name=Level,proto3" json:"Level,omitempty"`
}

func (m *C2SWordsStick) Reset()         { *m = C2SWordsStick{} }
func (m *C2SWordsStick) String() string { return proto.CompactTextString(m) }
func (*C2SWordsStick) ProtoMessage()    {}
func (m *C2SWordsStick) GetId() uint16  { return PCK_C2SWordsStick_ID }

const PCK_S2CWordsStick_ID = 16004 //
//
type S2CWordsStick struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CWordsStick) Reset()         { *m = S2CWordsStick{} }
func (m *S2CWordsStick) String() string { return proto.CompactTextString(m) }
func (*S2CWordsStick) ProtoMessage()    {}
func (m *S2CWordsStick) GetId() uint16  { return PCK_S2CWordsStick_ID }

const PCK_C2SInstanceSoulHallFight_ID = 691 //
//
type C2SInstanceSoulHallFight struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2SInstanceSoulHallFight) Reset()         { *m = C2SInstanceSoulHallFight{} }
func (m *C2SInstanceSoulHallFight) String() string { return proto.CompactTextString(m) }
func (*C2SInstanceSoulHallFight) ProtoMessage()    {}
func (m *C2SInstanceSoulHallFight) GetId() uint16  { return PCK_C2SInstanceSoulHallFight_ID }

const PCK_S2CInstanceSoulHallFight_ID = 692 //
//
type S2CInstanceSoulHallFight struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Win int32 `protobuf:"varint,3,opt,name=Win,proto3" json:"Win,omitempty"`
}

func (m *S2CInstanceSoulHallFight) Reset()         { *m = S2CInstanceSoulHallFight{} }
func (m *S2CInstanceSoulHallFight) String() string { return proto.CompactTextString(m) }
func (*S2CInstanceSoulHallFight) ProtoMessage()    {}
func (m *S2CInstanceSoulHallFight) GetId() uint16  { return PCK_S2CInstanceSoulHallFight_ID }

//
type HolyBeast struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	ItemId int32 `protobuf:"varint,2,opt,name=ItemId,proto3" json:"ItemId,omitempty"`
	//
	Level int32 `protobuf:"varint,3,opt,name=Level,proto3" json:"Level,omitempty"`
}

func (m *HolyBeast) Reset()         { *m = HolyBeast{} }
func (m *HolyBeast) String() string { return proto.CompactTextString(m) }
func (*HolyBeast) ProtoMessage()    {}

const PCK_C2SHolyBeast_ID = 701 //
//
type C2SHolyBeast struct {
	//
	Param int32 `protobuf:"varint,1,opt,name=Param,proto3" json:"Param,omitempty"`
}

func (m *C2SHolyBeast) Reset()         { *m = C2SHolyBeast{} }
func (m *C2SHolyBeast) String() string { return proto.CompactTextString(m) }
func (*C2SHolyBeast) ProtoMessage()    {}
func (m *C2SHolyBeast) GetId() uint16  { return PCK_C2SHolyBeast_ID }

const PCK_S2CHolyBeast_ID = 702 //
//
type S2CHolyBeast struct {
	//
	HolyBeasts []*HolyBeast `protobuf:"bytes,1,rep,name=HolyBeasts,proto3" json:"HolyBeasts,omitempty"`
	//
	SuitLevel int32 `protobuf:"varint,2,opt,name=SuitLevel,proto3" json:"SuitLevel,omitempty"`
	//
	Param int32 `protobuf:"varint,3,opt,name=Param,proto3" json:"Param,omitempty"`
}

func (m *S2CHolyBeast) Reset()         { *m = S2CHolyBeast{} }
func (m *S2CHolyBeast) String() string { return proto.CompactTextString(m) }
func (*S2CHolyBeast) ProtoMessage()    {}
func (m *S2CHolyBeast) GetId() uint16  { return PCK_S2CHolyBeast_ID }

const PCK_C2SHolyBeastLevel_ID = 705 //
//
type C2SHolyBeastLevel struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2SHolyBeastLevel) Reset()         { *m = C2SHolyBeastLevel{} }
func (m *C2SHolyBeastLevel) String() string { return proto.CompactTextString(m) }
func (*C2SHolyBeastLevel) ProtoMessage()    {}
func (m *C2SHolyBeastLevel) GetId() uint16  { return PCK_C2SHolyBeastLevel_ID }

const PCK_S2CHolyBeastLevel_ID = 706 //
//
type S2CHolyBeastLevel struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *S2CHolyBeastLevel) Reset()         { *m = S2CHolyBeastLevel{} }
func (m *S2CHolyBeastLevel) String() string { return proto.CompactTextString(m) }
func (*S2CHolyBeastLevel) ProtoMessage()    {}
func (m *S2CHolyBeastLevel) GetId() uint16  { return PCK_S2CHolyBeastLevel_ID }

const PCK_C2SMixHolySoul_ID = 693 //
//
type C2SMixHolySoul struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Type int32 `protobuf:"varint,2,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (m *C2SMixHolySoul) Reset()         { *m = C2SMixHolySoul{} }
func (m *C2SMixHolySoul) String() string { return proto.CompactTextString(m) }
func (*C2SMixHolySoul) ProtoMessage()    {}
func (m *C2SMixHolySoul) GetId() uint16  { return PCK_C2SMixHolySoul_ID }

const PCK_S2CMixHolySoul_ID = 694 //
//
type S2CMixHolySoul struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	MixOk int32 `protobuf:"varint,2,opt,name=MixOk,proto3" json:"MixOk,omitempty"`
	//
	NewHolySoul *ItemData `protobuf:"bytes,3,opt,name=NewHolySoul,proto3" json:"NewHolySoul,omitempty"`
}

func (m *S2CMixHolySoul) Reset()         { *m = S2CMixHolySoul{} }
func (m *S2CMixHolySoul) String() string { return proto.CompactTextString(m) }
func (*S2CMixHolySoul) ProtoMessage()    {}
func (m *S2CMixHolySoul) GetId() uint16  { return PCK_S2CMixHolySoul_ID }

//
type BreakHolySoulInfo struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Num int64 `protobuf:"varint,2,opt,name=Num,proto3" json:"Num,omitempty"`
}

func (m *BreakHolySoulInfo) Reset()         { *m = BreakHolySoulInfo{} }
func (m *BreakHolySoulInfo) String() string { return proto.CompactTextString(m) }
func (*BreakHolySoulInfo) ProtoMessage()    {}

const PCK_C2SBreakHolySoul_ID = 695 //
//
type C2SBreakHolySoul struct {
	//
	breaks []*BreakHolySoulInfo `protobuf:"bytes,1,rep,name=breaks,proto3" json:"breaks,omitempty"`
}

func (m *C2SBreakHolySoul) Reset()         { *m = C2SBreakHolySoul{} }
func (m *C2SBreakHolySoul) String() string { return proto.CompactTextString(m) }
func (*C2SBreakHolySoul) ProtoMessage()    {}
func (m *C2SBreakHolySoul) GetId() uint16  { return PCK_C2SBreakHolySoul_ID }

const PCK_S2CBreakHolySoul_ID = 696 //
//
type S2CBreakHolySoul struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	breaks []*BreakHolySoulInfo `protobuf:"bytes,2,rep,name=breaks,proto3" json:"breaks,omitempty"`
}

func (m *S2CBreakHolySoul) Reset()         { *m = S2CBreakHolySoul{} }
func (m *S2CBreakHolySoul) String() string { return proto.CompactTextString(m) }
func (*S2CBreakHolySoul) ProtoMessage()    {}
func (m *S2CBreakHolySoul) GetId() uint16  { return PCK_S2CBreakHolySoul_ID }

const PCK_C2SInjectHolySoul_ID = 697 //
//
type C2SInjectHolySoul struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Pos int32 `protobuf:"varint,2,opt,name=Pos,proto3" json:"Pos,omitempty"`
}

func (m *C2SInjectHolySoul) Reset()         { *m = C2SInjectHolySoul{} }
func (m *C2SInjectHolySoul) String() string { return proto.CompactTextString(m) }
func (*C2SInjectHolySoul) ProtoMessage()    {}
func (m *C2SInjectHolySoul) GetId() uint16  { return PCK_C2SInjectHolySoul_ID }

const PCK_S2CInjectHolySoul_ID = 698 //
//
type S2CInjectHolySoul struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Pos int32 `protobuf:"varint,3,opt,name=Pos,proto3" json:"Pos,omitempty"`
}

func (m *S2CInjectHolySoul) Reset()         { *m = S2CInjectHolySoul{} }
func (m *S2CInjectHolySoul) String() string { return proto.CompactTextString(m) }
func (*S2CInjectHolySoul) ProtoMessage()    {}
func (m *S2CInjectHolySoul) GetId() uint16  { return PCK_S2CInjectHolySoul_ID }

//
type InjectHolySoul struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Pos int32 `protobuf:"varint,2,opt,name=Pos,proto3" json:"Pos,omitempty"`
}

func (m *InjectHolySoul) Reset()         { *m = InjectHolySoul{} }
func (m *InjectHolySoul) String() string { return proto.CompactTextString(m) }
func (*InjectHolySoul) ProtoMessage()    {}

const PCK_C2SOneKeyInjectHolySoul_ID = 732 //
//
type C2SOneKeyInjectHolySoul struct {
	//
	souls []*InjectHolySoul `protobuf:"bytes,1,rep,name=souls,proto3" json:"souls,omitempty"`
}

func (m *C2SOneKeyInjectHolySoul) Reset()         { *m = C2SOneKeyInjectHolySoul{} }
func (m *C2SOneKeyInjectHolySoul) String() string { return proto.CompactTextString(m) }
func (*C2SOneKeyInjectHolySoul) ProtoMessage()    {}
func (m *C2SOneKeyInjectHolySoul) GetId() uint16  { return PCK_C2SOneKeyInjectHolySoul_ID }

const PCK_S2COneKeyInjectHolySoul_ID = 733 //
//
type S2COneKeyInjectHolySoul struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	souls []*InjectHolySoul `protobuf:"bytes,2,rep,name=souls,proto3" json:"souls,omitempty"`
}

func (m *S2COneKeyInjectHolySoul) Reset()         { *m = S2COneKeyInjectHolySoul{} }
func (m *S2COneKeyInjectHolySoul) String() string { return proto.CompactTextString(m) }
func (*S2COneKeyInjectHolySoul) ProtoMessage()    {}
func (m *S2COneKeyInjectHolySoul) GetId() uint16  { return PCK_S2COneKeyInjectHolySoul_ID }

const PCK_C2SOutHolySoul_ID = 699 //
//
type C2SOutHolySoul struct {
	//
	Pos int32 `protobuf:"varint,1,opt,name=Pos,proto3" json:"Pos,omitempty"`
}

func (m *C2SOutHolySoul) Reset()         { *m = C2SOutHolySoul{} }
func (m *C2SOutHolySoul) String() string { return proto.CompactTextString(m) }
func (*C2SOutHolySoul) ProtoMessage()    {}
func (m *C2SOutHolySoul) GetId() uint16  { return PCK_C2SOutHolySoul_ID }

const PCK_S2COutHolySoul_ID = 700 //
//
type S2COutHolySoul struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Pos int32 `protobuf:"varint,3,opt,name=Pos,proto3" json:"Pos,omitempty"`
}

func (m *S2COutHolySoul) Reset()         { *m = S2COutHolySoul{} }
func (m *S2COutHolySoul) String() string { return proto.CompactTextString(m) }
func (*S2COutHolySoul) ProtoMessage()    {}
func (m *S2COutHolySoul) GetId() uint16  { return PCK_S2COutHolySoul_ID }

//
type StakeData struct {
	//
	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	Pos int32 `protobuf:"varint,2,opt,name=Pos,proto3" json:"Pos,omitempty"`
	//
	StakeId int32 `protobuf:"varint,3,opt,name=StakeId,proto3" json:"StakeId,omitempty"`
	//
	Win int32 `protobuf:"varint,4,opt,name=Win,proto3" json:"Win,omitempty"`
	//
	FightId int32 `protobuf:"varint,5,opt,name=FightId,proto3" json:"FightId,omitempty"`
}

func (m *StakeData) Reset()         { *m = StakeData{} }
func (m *StakeData) String() string { return proto.CompactTextString(m) }
func (*StakeData) ProtoMessage()    {}

//
type GodClubDbData struct {
	//
	State int32 `protobuf:"varint,1,opt,name=State,proto3" json:"State,omitempty"`
	//
	HandleState []int32 `protobuf:"varint,2,rep,name=HandleState,proto3" json:"HandleState,omitempty"`
	//
	SignUpUser []int64 `protobuf:"varint,3,rep,name=SignUpUser,proto3" json:"SignUpUser,omitempty"`
	//
	Session int32 `protobuf:"varint,4,opt,name=Session,proto3" json:"Session,omitempty"`
	//
	Stakes []*StakeData `protobuf:"bytes,5,rep,name=Stakes,proto3" json:"Stakes,omitempty"`
}

func (m *GodClubDbData) Reset()         { *m = GodClubDbData{} }
func (m *GodClubDbData) String() string { return proto.CompactTextString(m) }
func (*GodClubDbData) ProtoMessage()    {}

//
type GodClubUserInfo struct {
	//
	Session int32 `protobuf:"varint,1,opt,name=Session,proto3" json:"Session,omitempty"`
	//
	BattlefieldId int32 `protobuf:"varint,2,opt,name=BattlefieldId,proto3" json:"BattlefieldId,omitempty"`
	//
	UserId int64 `protobuf:"varint,3,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	Win []int32 `protobuf:"varint,4,rep,name=Win,proto3" json:"Win,omitempty"`
	//
	Rank int32 `protobuf:"varint,5,opt,name=Rank,proto3" json:"Rank,omitempty"`
	//
	Pos int32 `protobuf:"varint,6,opt,name=Pos,proto3" json:"Pos,omitempty"`
	//
	Data *BaseDbInfo `protobuf:"bytes,7,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *GodClubUserInfo) Reset()         { *m = GodClubUserInfo{} }
func (m *GodClubUserInfo) String() string { return proto.CompactTextString(m) }
func (*GodClubUserInfo) ProtoMessage()    {}

//
type GodClubFightReportItem struct {
	//
	Win int32 `protobuf:"varint,3,opt,name=Win,proto3" json:"Win,omitempty"`
	//
	GameIdx int32 `protobuf:"varint,4,opt,name=GameIdx,proto3" json:"GameIdx,omitempty"`
	//
	Report *S2CBattlefieldReport `protobuf:"bytes,5,opt,name=Report,proto3" json:"Report,omitempty"`
}

func (m *GodClubFightReportItem) Reset()         { *m = GodClubFightReportItem{} }
func (m *GodClubFightReportItem) String() string { return proto.CompactTextString(m) }
func (*GodClubFightReportItem) ProtoMessage()    {}

//
type GodClubFightReport struct {
	//
	Session int32 `protobuf:"varint,1,opt,name=Session,proto3" json:"Session,omitempty"`
	//
	BattlefieldId int32 `protobuf:"varint,2,opt,name=BattlefieldId,proto3" json:"BattlefieldId,omitempty"`
	//
	FightIdx int32 `protobuf:"varint,3,opt,name=FightIdx,proto3" json:"FightIdx,omitempty"`
	//
	UserId1 int64 `protobuf:"varint,4,opt,name=UserId1,proto3" json:"UserId1,omitempty"`
	//
	UserId2 int64 `protobuf:"varint,5,opt,name=UserId2,proto3" json:"UserId2,omitempty"`
	//
	Items []*GodClubFightReportItem `protobuf:"bytes,6,rep,name=Items,proto3" json:"Items,omitempty"`
}

func (m *GodClubFightReport) Reset()         { *m = GodClubFightReport{} }
func (m *GodClubFightReport) String() string { return proto.CompactTextString(m) }
func (*GodClubFightReport) ProtoMessage()    {}

const PCK_C2SGodClubSignUp_ID = 2101 //
//
type C2SGodClubSignUp struct {
}

func (m *C2SGodClubSignUp) Reset()         { *m = C2SGodClubSignUp{} }
func (m *C2SGodClubSignUp) String() string { return proto.CompactTextString(m) }
func (*C2SGodClubSignUp) ProtoMessage()    {}
func (m *C2SGodClubSignUp) GetId() uint16  { return PCK_C2SGodClubSignUp_ID }

const PCK_S2CGodClubSignUp_ID = 2102 //
//
type S2CGodClubSignUp struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CGodClubSignUp) Reset()         { *m = S2CGodClubSignUp{} }
func (m *S2CGodClubSignUp) String() string { return proto.CompactTextString(m) }
func (*S2CGodClubSignUp) ProtoMessage()    {}
func (m *S2CGodClubSignUp) GetId() uint16  { return PCK_S2CGodClubSignUp_ID }

//
type GodClubFightUser struct {
	//
	UserId int64 `protobuf:"varint,3,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	A []*IntAttr `protobuf:"bytes,1,rep,name=A,proto3" json:"A,omitempty"`
	//
	B []*StrAttr `protobuf:"bytes,2,rep,name=B,proto3" json:"B,omitempty"`
}

func (m *GodClubFightUser) Reset()         { *m = GodClubFightUser{} }
func (m *GodClubFightUser) String() string { return proto.CompactTextString(m) }
func (*GodClubFightUser) ProtoMessage()    {}

const PCK_C2SGodClubFight_ID = 2103 //
//
type C2SGodClubFight struct {
	//
	FightIdx int32 `protobuf:"varint,1,opt,name=FightIdx,proto3" json:"FightIdx,omitempty"`
}

func (m *C2SGodClubFight) Reset()         { *m = C2SGodClubFight{} }
func (m *C2SGodClubFight) String() string { return proto.CompactTextString(m) }
func (*C2SGodClubFight) ProtoMessage()    {}
func (m *C2SGodClubFight) GetId() uint16  { return PCK_C2SGodClubFight_ID }

const PCK_S2CGodClubFight_ID = 2104 //
//
type S2CGodClubFight struct {
	//
	A *GodClubFightUser `protobuf:"bytes,1,opt,name=A,proto3" json:"A,omitempty"`
	//
	B *GodClubFightUser `protobuf:"bytes,2,opt,name=B,proto3" json:"B,omitempty"`
	//
	Win []int32 `protobuf:"varint,3,rep,name=Win,proto3" json:"Win,omitempty"`
	//
	FightIdx int32 `protobuf:"varint,4,opt,name=FightIdx,proto3" json:"FightIdx,omitempty"`
}

func (m *S2CGodClubFight) Reset()         { *m = S2CGodClubFight{} }
func (m *S2CGodClubFight) String() string { return proto.CompactTextString(m) }
func (*S2CGodClubFight) ProtoMessage()    {}
func (m *S2CGodClubFight) GetId() uint16  { return PCK_S2CGodClubFight_ID }

const PCK_C2SGodClubFightReport_ID = 2105 //
//
type C2SGodClubFightReport struct {
	//
	FightIdx int32 `protobuf:"varint,1,opt,name=FightIdx,proto3" json:"FightIdx,omitempty"`
	//
	GameIdx int32 `protobuf:"varint,2,opt,name=GameIdx,proto3" json:"GameIdx,omitempty"`
}

func (m *C2SGodClubFightReport) Reset()         { *m = C2SGodClubFightReport{} }
func (m *C2SGodClubFightReport) String() string { return proto.CompactTextString(m) }
func (*C2SGodClubFightReport) ProtoMessage()    {}
func (m *C2SGodClubFightReport) GetId() uint16  { return PCK_C2SGodClubFightReport_ID }

const PCK_S2CGodClubFightReport_ID = 2106 //
//
type S2CGodClubFightReport struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CGodClubFightReport) Reset()         { *m = S2CGodClubFightReport{} }
func (m *S2CGodClubFightReport) String() string { return proto.CompactTextString(m) }
func (*S2CGodClubFightReport) ProtoMessage()    {}
func (m *S2CGodClubFightReport) GetId() uint16  { return PCK_S2CGodClubFightReport_ID }

//
type GodClub16User struct {
	//
	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	RoleId int32 `protobuf:"varint,2,opt,name=RoleId,proto3" json:"RoleId,omitempty"`
	//
	Nick string `protobuf:"bytes,3,opt,name=Nick,proto3" json:"Nick,omitempty"`
	//
	AreaId int32 `protobuf:"varint,4,opt,name=AreaId,proto3" json:"AreaId,omitempty"`
	//
	Pos int32 `protobuf:"varint,5,opt,name=Pos,proto3" json:"Pos,omitempty"`
	//
	Win []int32 `protobuf:"varint,6,rep,name=Win,proto3" json:"Win,omitempty"`
}

func (m *GodClub16User) Reset()         { *m = GodClub16User{} }
func (m *GodClub16User) String() string { return proto.CompactTextString(m) }
func (*GodClub16User) ProtoMessage()    {}

const PCK_C2SGodClub16_ID = 2107 //
//
type C2SGodClub16 struct {
}

func (m *C2SGodClub16) Reset()         { *m = C2SGodClub16{} }
func (m *C2SGodClub16) String() string { return proto.CompactTextString(m) }
func (*C2SGodClub16) ProtoMessage()    {}
func (m *C2SGodClub16) GetId() uint16  { return PCK_C2SGodClub16_ID }

const PCK_S2CGodClub16_ID = 2108 //
//
type S2CGodClub16 struct {
	//
	Users []*GodClub16User `protobuf:"bytes,1,rep,name=Users,proto3" json:"Users,omitempty"`
	//
	Stakes []int32 `protobuf:"varint,2,rep,name=Stakes,proto3" json:"Stakes,omitempty"`
}

func (m *S2CGodClub16) Reset()         { *m = S2CGodClub16{} }
func (m *S2CGodClub16) String() string { return proto.CompactTextString(m) }
func (*S2CGodClub16) ProtoMessage()    {}
func (m *S2CGodClub16) GetId() uint16  { return PCK_S2CGodClub16_ID }

const PCK_C2SGodClubStakeInfo_ID = 2109 //
//
type C2SGodClubStakeInfo struct {
	//
	FightId int32 `protobuf:"varint,1,opt,name=FightId,proto3" json:"FightId,omitempty"`
}

func (m *C2SGodClubStakeInfo) Reset()         { *m = C2SGodClubStakeInfo{} }
func (m *C2SGodClubStakeInfo) String() string { return proto.CompactTextString(m) }
func (*C2SGodClubStakeInfo) ProtoMessage()    {}
func (m *C2SGodClubStakeInfo) GetId() uint16  { return PCK_C2SGodClubStakeInfo_ID }

const PCK_S2CGodClubStakeInfo_ID = 2110 //
//
type S2CGodClubStakeInfo struct {
	//
	FightId int32 `protobuf:"varint,1,opt,name=FightId,proto3" json:"FightId,omitempty"`
	//
	Pos int32 `protobuf:"varint,2,opt,name=Pos,proto3" json:"Pos,omitempty"`
	//
	StakeId int32 `protobuf:"varint,3,opt,name=StakeId,proto3" json:"StakeId,omitempty"`
	//
	A *GodClubFightUser `protobuf:"bytes,4,opt,name=A,proto3" json:"A,omitempty"`
	//
	B *GodClubFightUser `protobuf:"bytes,5,opt,name=B,proto3" json:"B,omitempty"`
}

func (m *S2CGodClubStakeInfo) Reset()         { *m = S2CGodClubStakeInfo{} }
func (m *S2CGodClubStakeInfo) String() string { return proto.CompactTextString(m) }
func (*S2CGodClubStakeInfo) ProtoMessage()    {}
func (m *S2CGodClubStakeInfo) GetId() uint16  { return PCK_S2CGodClubStakeInfo_ID }

const PCK_C2SGodClubStake_ID = 2111 //
//
type C2SGodClubStake struct {
	//
	FightId int32 `protobuf:"varint,1,opt,name=FightId,proto3" json:"FightId,omitempty"`
	//
	Pos int32 `protobuf:"varint,2,opt,name=Pos,proto3" json:"Pos,omitempty"`
	//
	StackId int32 `protobuf:"varint,3,opt,name=StackId,proto3" json:"StackId,omitempty"`
}

func (m *C2SGodClubStake) Reset()         { *m = C2SGodClubStake{} }
func (m *C2SGodClubStake) String() string { return proto.CompactTextString(m) }
func (*C2SGodClubStake) ProtoMessage()    {}
func (m *C2SGodClubStake) GetId() uint16  { return PCK_C2SGodClubStake_ID }

const PCK_S2CGodClubStake_ID = 2112 //
//
type S2CGodClubStake struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CGodClubStake) Reset()         { *m = S2CGodClubStake{} }
func (m *S2CGodClubStake) String() string { return proto.CompactTextString(m) }
func (*S2CGodClubStake) ProtoMessage()    {}
func (m *S2CGodClubStake) GetId() uint16  { return PCK_S2CGodClubStake_ID }

const PCK_C2SGodHerbEnter_ID = 17001 //
//
type C2SGodHerbEnter struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2SGodHerbEnter) Reset()         { *m = C2SGodHerbEnter{} }
func (m *C2SGodHerbEnter) String() string { return proto.CompactTextString(m) }
func (*C2SGodHerbEnter) ProtoMessage()    {}
func (m *C2SGodHerbEnter) GetId() uint16  { return PCK_C2SGodHerbEnter_ID }

const PCK_S2CGodHerbEnter_ID = 17002 //
//
type S2CGodHerbEnter struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *S2CGodHerbEnter) Reset()         { *m = S2CGodHerbEnter{} }
func (m *S2CGodHerbEnter) String() string { return proto.CompactTextString(m) }
func (*S2CGodHerbEnter) ProtoMessage()    {}
func (m *S2CGodHerbEnter) GetId() uint16  { return PCK_S2CGodHerbEnter_ID }

const PCK_C2SGodHerbRefresh_ID = 17003 //
//
type C2SGodHerbRefresh struct {
}

func (m *C2SGodHerbRefresh) Reset()         { *m = C2SGodHerbRefresh{} }
func (m *C2SGodHerbRefresh) String() string { return proto.CompactTextString(m) }
func (*C2SGodHerbRefresh) ProtoMessage()    {}
func (m *C2SGodHerbRefresh) GetId() uint16  { return PCK_C2SGodHerbRefresh_ID }

const PCK_S2CGodHerbRefresh_ID = 17004 //
//
type S2CGodHerbRefresh struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CGodHerbRefresh) Reset()         { *m = S2CGodHerbRefresh{} }
func (m *S2CGodHerbRefresh) String() string { return proto.CompactTextString(m) }
func (*S2CGodHerbRefresh) ProtoMessage()    {}
func (m *S2CGodHerbRefresh) GetId() uint16  { return PCK_S2CGodHerbRefresh_ID }

const PCK_C2SGodHerbCollect_ID = 17005 //
//
type C2SGodHerbCollect struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Num int32 `protobuf:"varint,2,opt,name=Num,proto3" json:"Num,omitempty"`
	//
	AutoBuy int32 `protobuf:"varint,3,opt,name=AutoBuy,proto3" json:"AutoBuy,omitempty"`
}

func (m *C2SGodHerbCollect) Reset()         { *m = C2SGodHerbCollect{} }
func (m *C2SGodHerbCollect) String() string { return proto.CompactTextString(m) }
func (*C2SGodHerbCollect) ProtoMessage()    {}
func (m *C2SGodHerbCollect) GetId() uint16  { return PCK_C2SGodHerbCollect_ID }

const PCK_S2CGodHerbCollect_ID = 17006 //
//
type S2CGodHerbCollect struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Num int32 `protobuf:"varint,2,opt,name=Num,proto3" json:"Num,omitempty"`
	//
	Data []*ItemData `protobuf:"bytes,3,rep,name=Data,proto3" json:"Data,omitempty"`
}

func (m *S2CGodHerbCollect) Reset()         { *m = S2CGodHerbCollect{} }
func (m *S2CGodHerbCollect) String() string { return proto.CompactTextString(m) }
func (*S2CGodHerbCollect) ProtoMessage()    {}
func (m *S2CGodHerbCollect) GetId() uint16  { return PCK_S2CGodHerbCollect_ID }

const PCK_C2SGetGodHerbLog_ID = 17007 //
//
type C2SGetGodHerbLog struct {
	//
	Type int32 `protobuf:"varint,1,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (m *C2SGetGodHerbLog) Reset()         { *m = C2SGetGodHerbLog{} }
func (m *C2SGetGodHerbLog) String() string { return proto.CompactTextString(m) }
func (*C2SGetGodHerbLog) ProtoMessage()    {}
func (m *C2SGetGodHerbLog) GetId() uint16  { return PCK_C2SGetGodHerbLog_ID }

//
type GodHerbLogItemClient struct {
	//
	Time int64 `protobuf:"varint,1,opt,name=Time,proto3" json:"Time,omitempty"`
	//
	User *NoticeUser `protobuf:"bytes,2,opt,name=User,proto3" json:"User,omitempty"`
	//
	GodHerbId int64 `protobuf:"varint,3,opt,name=GodHerbId,proto3" json:"GodHerbId,omitempty"`
	//
	Item []*ItemData `protobuf:"bytes,4,rep,name=Item,proto3" json:"Item,omitempty"`
}

func (m *GodHerbLogItemClient) Reset()         { *m = GodHerbLogItemClient{} }
func (m *GodHerbLogItemClient) String() string { return proto.CompactTextString(m) }
func (*GodHerbLogItemClient) ProtoMessage()    {}

const PCK_S2CGetGodHerbLog_ID = 17008 //
//
type S2CGetGodHerbLog struct {
	//
	Type int32 `protobuf:"varint,1,opt,name=Type,proto3" json:"Type,omitempty"`
	//
	Logs []*GodHerbLogItemClient `protobuf:"bytes,2,rep,name=Logs,proto3" json:"Logs,omitempty"`
}

func (m *S2CGetGodHerbLog) Reset()         { *m = S2CGetGodHerbLog{} }
func (m *S2CGetGodHerbLog) String() string { return proto.CompactTextString(m) }
func (*S2CGetGodHerbLog) ProtoMessage()    {}
func (m *S2CGetGodHerbLog) GetId() uint16  { return PCK_S2CGetGodHerbLog_ID }

//
type RobberyItem struct {
	//
	ItemId int32 `protobuf:"varint,1,opt,name=ItemId,proto3" json:"ItemId,omitempty"`
	//
	RobberyId int32 `protobuf:"varint,2,opt,name=RobberyId,proto3" json:"RobberyId,omitempty"`
	//
	Coin3 int64 `protobuf:"varint,3,opt,name=Coin3,proto3" json:"Coin3,omitempty"`
}

func (m *RobberyItem) Reset()         { *m = RobberyItem{} }
func (m *RobberyItem) String() string { return proto.CompactTextString(m) }
func (*RobberyItem) ProtoMessage()    {}

//
type RobberyLogItem struct {
	//
	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	Nick string `protobuf:"bytes,2,opt,name=Nick,proto3" json:"Nick,omitempty"`
	//
	AreaId int32 `protobuf:"varint,3,opt,name=AreaId,proto3" json:"AreaId,omitempty"`
	//
	Time int64 `protobuf:"varint,4,opt,name=Time,proto3" json:"Time,omitempty"`
	//
	Pos int32 `protobuf:"varint,5,opt,name=Pos,proto3" json:"Pos,omitempty"`
	//
	Mul int32 `protobuf:"varint,6,opt,name=Mul,proto3" json:"Mul,omitempty"`
	//
	Coin3 int64 `protobuf:"varint,7,opt,name=Coin3,proto3" json:"Coin3,omitempty"`
}

func (m *RobberyLogItem) Reset()         { *m = RobberyLogItem{} }
func (m *RobberyLogItem) String() string { return proto.CompactTextString(m) }
func (*RobberyLogItem) ProtoMessage()    {}

const PCK_S2CRobberyInfo_ID = 17101 //
//
type S2CRobberyInfo struct {
	//
	Items []*RobberData `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty"`
}

func (m *S2CRobberyInfo) Reset()         { *m = S2CRobberyInfo{} }
func (m *S2CRobberyInfo) String() string { return proto.CompactTextString(m) }
func (*S2CRobberyInfo) ProtoMessage()    {}
func (m *S2CRobberyInfo) GetId() uint16  { return PCK_S2CRobberyInfo_ID }

const PCK_C2SGetRobberyData_ID = 17102 //
//
type C2SGetRobberyData struct {
	//
	ItemId int32 `protobuf:"varint,1,opt,name=ItemId,proto3" json:"ItemId,omitempty"`
	//
	Num int32 `protobuf:"varint,2,opt,name=Num,proto3" json:"Num,omitempty"`
}

func (m *C2SGetRobberyData) Reset()         { *m = C2SGetRobberyData{} }
func (m *C2SGetRobberyData) String() string { return proto.CompactTextString(m) }
func (*C2SGetRobberyData) ProtoMessage()    {}
func (m *C2SGetRobberyData) GetId() uint16  { return PCK_C2SGetRobberyData_ID }

const PCK_S2CGetRobberyData_ID = 17103 //
//
type S2CGetRobberyData struct {
	//
	Coin3 int64 `protobuf:"varint,1,opt,name=Coin3,proto3" json:"Coin3,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Mul int32 `protobuf:"varint,3,opt,name=Mul,proto3" json:"Mul,omitempty"`
}

func (m *S2CGetRobberyData) Reset()         { *m = S2CGetRobberyData{} }
func (m *S2CGetRobberyData) String() string { return proto.CompactTextString(m) }
func (*S2CGetRobberyData) ProtoMessage()    {}
func (m *S2CGetRobberyData) GetId() uint16  { return PCK_S2CGetRobberyData_ID }

const PCK_C2SRobbery_ID = 17104 //
//
type C2SRobbery struct {
	//
	ItemId int32 `protobuf:"varint,1,opt,name=ItemId,proto3" json:"ItemId,omitempty"`
	//
	Num int32 `protobuf:"varint,2,opt,name=Num,proto3" json:"Num,omitempty"`
	//
	Pos int32 `protobuf:"varint,3,opt,name=Pos,proto3" json:"Pos,omitempty"`
}

func (m *C2SRobbery) Reset()         { *m = C2SRobbery{} }
func (m *C2SRobbery) String() string { return proto.CompactTextString(m) }
func (*C2SRobbery) ProtoMessage()    {}
func (m *C2SRobbery) GetId() uint16  { return PCK_C2SRobbery_ID }

const PCK_S2CRobbery_ID = 17105 //
//
type S2CRobbery struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Items []*RobberData `protobuf:"bytes,3,rep,name=Items,proto3" json:"Items,omitempty"`
}

func (m *S2CRobbery) Reset()         { *m = S2CRobbery{} }
func (m *S2CRobbery) String() string { return proto.CompactTextString(m) }
func (*S2CRobbery) ProtoMessage()    {}
func (m *S2CRobbery) GetId() uint16  { return PCK_S2CRobbery_ID }

const PCK_C2SRobberyLuck_ID = 17106 //
//
type C2SRobberyLuck struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2SRobberyLuck) Reset()         { *m = C2SRobberyLuck{} }
func (m *C2SRobberyLuck) String() string { return proto.CompactTextString(m) }
func (*C2SRobberyLuck) ProtoMessage()    {}
func (m *C2SRobberyLuck) GetId() uint16  { return PCK_C2SRobberyLuck_ID }

const PCK_S2CRobberyLuck_ID = 17107 //
//
type S2CRobberyLuck struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *S2CRobberyLuck) Reset()         { *m = S2CRobberyLuck{} }
func (m *S2CRobberyLuck) String() string { return proto.CompactTextString(m) }
func (*S2CRobberyLuck) ProtoMessage()    {}
func (m *S2CRobberyLuck) GetId() uint16  { return PCK_S2CRobberyLuck_ID }

const PCK_C2SGetRobberyLog_ID = 17108 //
//
type C2SGetRobberyLog struct {
	//
	Type int32 `protobuf:"varint,1,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (m *C2SGetRobberyLog) Reset()         { *m = C2SGetRobberyLog{} }
func (m *C2SGetRobberyLog) String() string { return proto.CompactTextString(m) }
func (*C2SGetRobberyLog) ProtoMessage()    {}
func (m *C2SGetRobberyLog) GetId() uint16  { return PCK_C2SGetRobberyLog_ID }

const PCK_S2CGetRobberyLog_ID = 17109 //
//
type S2CGetRobberyLog struct {
	//
	Items []*RobberyLogItem `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty"`
	//
	Type int32 `protobuf:"varint,2,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (m *S2CGetRobberyLog) Reset()         { *m = S2CGetRobberyLog{} }
func (m *S2CGetRobberyLog) String() string { return proto.CompactTextString(m) }
func (*S2CGetRobberyLog) ProtoMessage()    {}
func (m *S2CGetRobberyLog) GetId() uint16  { return PCK_S2CGetRobberyLog_ID }

//
type FairyMemoir struct {
	//
	MemoirId int32 `protobuf:"varint,1,opt,name=MemoirId,proto3" json:"MemoirId,omitempty"`
	//
	State int32 `protobuf:"varint,2,opt,name=State,proto3" json:"State,omitempty"`
	//
	IsReceivedPrize int32 `protobuf:"varint,3,opt,name=IsReceivedPrize,proto3" json:"IsReceivedPrize,omitempty"`
	//
	NeedExploredSection int32 `protobuf:"varint,4,opt,name=NeedExploredSection,proto3" json:"NeedExploredSection,omitempty"`
}

func (m *FairyMemoir) Reset()         { *m = FairyMemoir{} }
func (m *FairyMemoir) String() string { return proto.CompactTextString(m) }
func (*FairyMemoir) ProtoMessage()    {}

//
type FairySheet struct {
	//
	SheetId int32 `protobuf:"varint,1,opt,name=SheetId,proto3" json:"SheetId,omitempty"`
	//
	Memoirs []*FairyMemoir `protobuf:"bytes,2,rep,name=Memoirs,proto3" json:"Memoirs,omitempty"`
}

func (m *FairySheet) Reset()         { *m = FairySheet{} }
func (m *FairySheet) String() string { return proto.CompactTextString(m) }
func (*FairySheet) ProtoMessage()    {}

//
type FairyAchievement struct {
	//
	AchievementId int32 `protobuf:"varint,1,opt,name=AchievementId,proto3" json:"AchievementId,omitempty"`
	//
	State int32 `protobuf:"varint,2,opt,name=State,proto3" json:"State,omitempty"`
	//
	Count int32 `protobuf:"varint,3,opt,name=Count,proto3" json:"Count,omitempty"`
	//
	TimeStamp int64 `protobuf:"varint,4,opt,name=TimeStamp,proto3" json:"TimeStamp,omitempty"`
}

func (m *FairyAchievement) Reset()         { *m = FairyAchievement{} }
func (m *FairyAchievement) String() string { return proto.CompactTextString(m) }
func (*FairyAchievement) ProtoMessage()    {}

const PCK_S2CFairyMemoirTrigger_ID = 781 //
//
type S2CFairyMemoirTrigger struct {
	//
	SheetId int32 `protobuf:"varint,1,opt,name=SheetId,proto3" json:"SheetId,omitempty"`
	//
	MemoirId int32 `protobuf:"varint,2,opt,name=MemoirId,proto3" json:"MemoirId,omitempty"`
	//
	State int32 `protobuf:"varint,3,opt,name=State,proto3" json:"State,omitempty"`
}

func (m *S2CFairyMemoirTrigger) Reset()         { *m = S2CFairyMemoirTrigger{} }
func (m *S2CFairyMemoirTrigger) String() string { return proto.CompactTextString(m) }
func (*S2CFairyMemoirTrigger) ProtoMessage()    {}
func (m *S2CFairyMemoirTrigger) GetId() uint16  { return PCK_S2CFairyMemoirTrigger_ID }

const PCK_S2CFairyAchievementChange_ID = 780 //
//
type S2CFairyAchievementChange struct {
	//
	achievement []*FairyAchievement `protobuf:"bytes,1,rep,name=achievement,proto3" json:"achievement,omitempty"`
}

func (m *S2CFairyAchievementChange) Reset()         { *m = S2CFairyAchievementChange{} }
func (m *S2CFairyAchievementChange) String() string { return proto.CompactTextString(m) }
func (*S2CFairyAchievementChange) ProtoMessage()    {}
func (m *S2CFairyAchievementChange) GetId() uint16  { return PCK_S2CFairyAchievementChange_ID }

const PCK_C2SFairyCharacterData_ID = 750 //
//
type C2SFairyCharacterData struct {
}

func (m *C2SFairyCharacterData) Reset()         { *m = C2SFairyCharacterData{} }
func (m *C2SFairyCharacterData) String() string { return proto.CompactTextString(m) }
func (*C2SFairyCharacterData) ProtoMessage()    {}
func (m *C2SFairyCharacterData) GetId() uint16  { return PCK_C2SFairyCharacterData_ID }

const PCK_S2CFairyCharacterData_ID = 751 //
//
type S2CFairyCharacterData struct {
	//
	FightTimes int32 `protobuf:"varint,1,opt,name=FightTimes,proto3" json:"FightTimes,omitempty"`
	//
	Sheets []*FairySheet `protobuf:"bytes,2,rep,name=Sheets,proto3" json:"Sheets,omitempty"`
	//
	achievement []*FairyAchievement `protobuf:"bytes,3,rep,name=achievement,proto3" json:"achievement,omitempty"`
}

func (m *S2CFairyCharacterData) Reset()         { *m = S2CFairyCharacterData{} }
func (m *S2CFairyCharacterData) String() string { return proto.CompactTextString(m) }
func (*S2CFairyCharacterData) ProtoMessage()    {}
func (m *S2CFairyCharacterData) GetId() uint16  { return PCK_S2CFairyCharacterData_ID }

const PCK_C2SFairyCharacterFight_ID = 752 //
//
type C2SFairyCharacterFight struct {
	//
	SheetId int32 `protobuf:"varint,1,opt,name=SheetId,proto3" json:"SheetId,omitempty"`
	//
	MemoirId int32 `protobuf:"varint,2,opt,name=MemoirId,proto3" json:"MemoirId,omitempty"`
	//
	SectionId int32 `protobuf:"varint,3,opt,name=SectionId,proto3" json:"SectionId,omitempty"`
}

func (m *C2SFairyCharacterFight) Reset()         { *m = C2SFairyCharacterFight{} }
func (m *C2SFairyCharacterFight) String() string { return proto.CompactTextString(m) }
func (*C2SFairyCharacterFight) ProtoMessage()    {}
func (m *C2SFairyCharacterFight) GetId() uint16  { return PCK_C2SFairyCharacterFight_ID }

const PCK_S2CFairyCharacterFight_ID = 753 //
//
type S2CFairyCharacterFight struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	SheetId int32 `protobuf:"varint,2,opt,name=SheetId,proto3" json:"SheetId,omitempty"`
	//
	MemoirId int32 `protobuf:"varint,3,opt,name=MemoirId,proto3" json:"MemoirId,omitempty"`
	//
	SectionId int32 `protobuf:"varint,4,opt,name=SectionId,proto3" json:"SectionId,omitempty"`
	//
	Win bool `protobuf:"varint,5,opt,name=Win,proto3" json:"Win,omitempty"`
}

func (m *S2CFairyCharacterFight) Reset()         { *m = S2CFairyCharacterFight{} }
func (m *S2CFairyCharacterFight) String() string { return proto.CompactTextString(m) }
func (*S2CFairyCharacterFight) ProtoMessage()    {}
func (m *S2CFairyCharacterFight) GetId() uint16  { return PCK_S2CFairyCharacterFight_ID }

const PCK_C2SFairyMemoirPrize_ID = 754 //
//
type C2SFairyMemoirPrize struct {
	//
	SheetId int32 `protobuf:"varint,1,opt,name=SheetId,proto3" json:"SheetId,omitempty"`
	//
	MemoirId int32 `protobuf:"varint,2,opt,name=MemoirId,proto3" json:"MemoirId,omitempty"`
}

func (m *C2SFairyMemoirPrize) Reset()         { *m = C2SFairyMemoirPrize{} }
func (m *C2SFairyMemoirPrize) String() string { return proto.CompactTextString(m) }
func (*C2SFairyMemoirPrize) ProtoMessage()    {}
func (m *C2SFairyMemoirPrize) GetId() uint16  { return PCK_C2SFairyMemoirPrize_ID }

const PCK_S2CFairyMemoirPrize_ID = 755 //
//
type S2CFairyMemoirPrize struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	SheetId int32 `protobuf:"varint,2,opt,name=SheetId,proto3" json:"SheetId,omitempty"`
	//
	MemoirId int32 `protobuf:"varint,3,opt,name=MemoirId,proto3" json:"MemoirId,omitempty"`
	//
	Items []*ItemData `protobuf:"bytes,4,rep,name=Items,proto3" json:"Items,omitempty"`
}

func (m *S2CFairyMemoirPrize) Reset()         { *m = S2CFairyMemoirPrize{} }
func (m *S2CFairyMemoirPrize) String() string { return proto.CompactTextString(m) }
func (*S2CFairyMemoirPrize) ProtoMessage()    {}
func (m *S2CFairyMemoirPrize) GetId() uint16  { return PCK_S2CFairyMemoirPrize_ID }

const PCK_C2SFairyCharacterSweep_ID = 756 //
//
type C2SFairyCharacterSweep struct {
	//
	SheetId int32 `protobuf:"varint,1,opt,name=SheetId,proto3" json:"SheetId,omitempty"`
	//
	MemoirId int32 `protobuf:"varint,2,opt,name=MemoirId,proto3" json:"MemoirId,omitempty"`
}

func (m *C2SFairyCharacterSweep) Reset()         { *m = C2SFairyCharacterSweep{} }
func (m *C2SFairyCharacterSweep) String() string { return proto.CompactTextString(m) }
func (*C2SFairyCharacterSweep) ProtoMessage()    {}
func (m *C2SFairyCharacterSweep) GetId() uint16  { return PCK_C2SFairyCharacterSweep_ID }

const PCK_S2CFairyCharacterSweep_ID = 757 //
//
type S2CFairyCharacterSweep struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	SheetId int32 `protobuf:"varint,2,opt,name=SheetId,proto3" json:"SheetId,omitempty"`
	//
	MemoirId int32 `protobuf:"varint,3,opt,name=MemoirId,proto3" json:"MemoirId,omitempty"`
	//
	Items []*ItemData `protobuf:"bytes,4,rep,name=Items,proto3" json:"Items,omitempty"`
}

func (m *S2CFairyCharacterSweep) Reset()         { *m = S2CFairyCharacterSweep{} }
func (m *S2CFairyCharacterSweep) String() string { return proto.CompactTextString(m) }
func (*S2CFairyCharacterSweep) ProtoMessage()    {}
func (m *S2CFairyCharacterSweep) GetId() uint16  { return PCK_S2CFairyCharacterSweep_ID }

const PCK_C2SFairyAchieventPrize_ID = 758 //
//
type C2SFairyAchieventPrize struct {
	//
	AchieventId int32 `protobuf:"varint,1,opt,name=AchieventId,proto3" json:"AchieventId,omitempty"`
}

func (m *C2SFairyAchieventPrize) Reset()         { *m = C2SFairyAchieventPrize{} }
func (m *C2SFairyAchieventPrize) String() string { return proto.CompactTextString(m) }
func (*C2SFairyAchieventPrize) ProtoMessage()    {}
func (m *C2SFairyAchieventPrize) GetId() uint16  { return PCK_C2SFairyAchieventPrize_ID }

const PCK_S2CFairyAchieventPrize_ID = 759 //
//
type S2CFairyAchieventPrize struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	AchieventId int32 `protobuf:"varint,2,opt,name=AchieventId,proto3" json:"AchieventId,omitempty"`
	//
	Items []*ItemData `protobuf:"bytes,3,rep,name=Items,proto3" json:"Items,omitempty"`
}

func (m *S2CFairyAchieventPrize) Reset()         { *m = S2CFairyAchieventPrize{} }
func (m *S2CFairyAchieventPrize) String() string { return proto.CompactTextString(m) }
func (*S2CFairyAchieventPrize) ProtoMessage()    {}
func (m *S2CFairyAchieventPrize) GetId() uint16  { return PCK_S2CFairyAchieventPrize_ID }

//
type PlayerDragonTimes struct {
	//
	InstanceType int32 `protobuf:"varint,1,opt,name=InstanceType,proto3" json:"InstanceType,omitempty"`
	//
	InstanceId int32 `protobuf:"varint,2,opt,name=InstanceId,proto3" json:"InstanceId,omitempty"`
	//
	Times int64 `protobuf:"varint,3,opt,name=Times,proto3" json:"Times,omitempty"`
}

func (m *PlayerDragonTimes) Reset()         { *m = PlayerDragonTimes{} }
func (m *PlayerDragonTimes) String() string { return proto.CompactTextString(m) }
func (*PlayerDragonTimes) ProtoMessage()    {}

//
type PlayerDragon struct {
	//
	CloseStock int32 `protobuf:"varint,1,opt,name=CloseStock,proto3" json:"CloseStock,omitempty"`
	//
	DragonOdd int64 `protobuf:"varint,2,opt,name=DragonOdd,proto3" json:"DragonOdd,omitempty"`
	//
	TotalProfit int64 `protobuf:"varint,3,opt,name=TotalProfit,proto3" json:"TotalProfit,omitempty"`
	//
	TotalProfits []int64 `protobuf:"varint,4,rep,name=TotalProfits,proto3" json:"TotalProfits,omitempty"`
	//
	TotalFlowingWater int64 `protobuf:"varint,5,opt,name=TotalFlowingWater,proto3" json:"TotalFlowingWater,omitempty"`
	//
	TotalFlowingWaters []int64 `protobuf:"varint,6,rep,name=TotalFlowingWaters,proto3" json:"TotalFlowingWaters,omitempty"`
	//
	DayProfit int64 `protobuf:"varint,7,opt,name=DayProfit,proto3" json:"DayProfit,omitempty"`
	//
	DayProfits []int64 `protobuf:"varint,8,rep,name=DayProfits,proto3" json:"DayProfits,omitempty"`
	//
	DayFlowingWater int64 `protobuf:"varint,9,opt,name=DayFlowingWater,proto3" json:"DayFlowingWater,omitempty"`
	//
	DayFlowingWaters []int64 `protobuf:"varint,10,rep,name=DayFlowingWaters,proto3" json:"DayFlowingWaters,omitempty"`
	//
	TotalGuard int64 `protobuf:"varint,11,opt,name=TotalGuard,proto3" json:"TotalGuard,omitempty"`
	//
	TotalRebate int64 `protobuf:"varint,12,opt,name=TotalRebate,proto3" json:"TotalRebate,omitempty"`
	//
	DayGuard int64 `protobuf:"varint,13,opt,name=DayGuard,proto3" json:"DayGuard,omitempty"`
	//
	DayRebate int64 `protobuf:"varint,14,opt,name=DayRebate,proto3" json:"DayRebate,omitempty"`
	//
	CurrGuard int64 `protobuf:"varint,15,opt,name=CurrGuard,proto3" json:"CurrGuard,omitempty"`
	//
	CurrRebate int64 `protobuf:"varint,16,opt,name=CurrRebate,proto3" json:"CurrRebate,omitempty"`
	//
	MaxFlowingWater int64 `protobuf:"varint,17,opt,name=MaxFlowingWater,proto3" json:"MaxFlowingWater,omitempty"`
	//
	DayMaxFlowingWater int64 `protobuf:"varint,18,opt,name=DayMaxFlowingWater,proto3" json:"DayMaxFlowingWater,omitempty"`
	//
	Day int32 `protobuf:"varint,19,opt,name=Day,proto3" json:"Day,omitempty"`
	//
	UserId int64 `protobuf:"varint,20,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	DragonTimes []*PlayerDragonTimes `protobuf:"bytes,21,rep,name=DragonTimes,proto3" json:"DragonTimes,omitempty"`
	//
	CurrNewGuard int64 `protobuf:"varint,22,opt,name=CurrNewGuard,proto3" json:"CurrNewGuard,omitempty"`
}

func (m *PlayerDragon) Reset()         { *m = PlayerDragon{} }
func (m *PlayerDragon) String() string { return proto.CompactTextString(m) }
func (*PlayerDragon) ProtoMessage()    {}

//
type ClientDb struct {
	//
	Key string `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	//
	Value string `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (m *ClientDb) Reset()         { *m = ClientDb{} }
func (m *ClientDb) String() string { return proto.CompactTextString(m) }
func (*ClientDb) ProtoMessage()    {}

const PCK_S2CClientDb_ID = 540 //
//
type S2CClientDb struct {
	//
	Items []*ClientDb `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty"`
}

func (m *S2CClientDb) Reset()         { *m = S2CClientDb{} }
func (m *S2CClientDb) String() string { return proto.CompactTextString(m) }
func (*S2CClientDb) ProtoMessage()    {}
func (m *S2CClientDb) GetId() uint16  { return PCK_S2CClientDb_ID }

const PCK_C2SClinetSet_ID = 541 //
//
type C2SClinetSet struct {
	//
	Key string `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	//
	Value string `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (m *C2SClinetSet) Reset()         { *m = C2SClinetSet{} }
func (m *C2SClinetSet) String() string { return proto.CompactTextString(m) }
func (*C2SClinetSet) ProtoMessage()    {}
func (m *C2SClinetSet) GetId() uint16  { return PCK_C2SClinetSet_ID }

const PCK_S2CClinetSet_ID = 542 //
//
type S2CClinetSet struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CClinetSet) Reset()         { *m = S2CClinetSet{} }
func (m *S2CClinetSet) String() string { return proto.CompactTextString(m) }
func (*S2CClinetSet) ProtoMessage()    {}
func (m *S2CClinetSet) GetId() uint16  { return PCK_S2CClinetSet_ID }

//
type PassCheckItem struct {
	//
	Level int64 `protobuf:"varint,1,opt,name=Level,proto3" json:"Level,omitempty"`
	//
	Prize1 int32 `protobuf:"varint,2,opt,name=Prize1,proto3" json:"Prize1,omitempty"`
	//
	Prize2 int32 `protobuf:"varint,3,opt,name=Prize2,proto3" json:"Prize2,omitempty"`
}

func (m *PassCheckItem) Reset()         { *m = PassCheckItem{} }
func (m *PassCheckItem) String() string { return proto.CompactTextString(m) }
func (*PassCheckItem) ProtoMessage()    {}

const PCK_S2CPassCheckInfo_ID = 2201 //
//
type S2CPassCheckInfo struct {
	//
	Data []*PassCheckItem `protobuf:"bytes,1,rep,name=Data,proto3" json:"Data,omitempty"`
}

func (m *S2CPassCheckInfo) Reset()         { *m = S2CPassCheckInfo{} }
func (m *S2CPassCheckInfo) String() string { return proto.CompactTextString(m) }
func (*S2CPassCheckInfo) ProtoMessage()    {}
func (m *S2CPassCheckInfo) GetId() uint16  { return PCK_S2CPassCheckInfo_ID }

const PCK_C2SPassCheckGetPrize_ID = 2202 //
//
type C2SPassCheckGetPrize struct {
	//
	Level int32 `protobuf:"varint,1,opt,name=Level,proto3" json:"Level,omitempty"`
	//
	PrizeType int32 `protobuf:"varint,2,opt,name=PrizeType,proto3" json:"PrizeType,omitempty"`
}

func (m *C2SPassCheckGetPrize) Reset()         { *m = C2SPassCheckGetPrize{} }
func (m *C2SPassCheckGetPrize) String() string { return proto.CompactTextString(m) }
func (*C2SPassCheckGetPrize) ProtoMessage()    {}
func (m *C2SPassCheckGetPrize) GetId() uint16  { return PCK_C2SPassCheckGetPrize_ID }

const PCK_S2CPassCheckGetPrize_ID = 2203 //
//
type S2CPassCheckGetPrize struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Level int32 `protobuf:"varint,2,opt,name=Level,proto3" json:"Level,omitempty"`
	//
	PrizeType int32 `protobuf:"varint,3,opt,name=PrizeType,proto3" json:"PrizeType,omitempty"`
}

func (m *S2CPassCheckGetPrize) Reset()         { *m = S2CPassCheckGetPrize{} }
func (m *S2CPassCheckGetPrize) String() string { return proto.CompactTextString(m) }
func (*S2CPassCheckGetPrize) ProtoMessage()    {}
func (m *S2CPassCheckGetPrize) GetId() uint16  { return PCK_S2CPassCheckGetPrize_ID }

const PCK_C2SPassCheckFast_ID = 2204 //
//
type C2SPassCheckFast struct {
	//
	Exp int32 `protobuf:"varint,1,opt,name=Exp,proto3" json:"Exp,omitempty"`
}

func (m *C2SPassCheckFast) Reset()         { *m = C2SPassCheckFast{} }
func (m *C2SPassCheckFast) String() string { return proto.CompactTextString(m) }
func (*C2SPassCheckFast) ProtoMessage()    {}
func (m *C2SPassCheckFast) GetId() uint16  { return PCK_C2SPassCheckFast_ID }

const PCK_S2CPassCheckFast_ID = 2205 //
//
type S2CPassCheckFast struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CPassCheckFast) Reset()         { *m = S2CPassCheckFast{} }
func (m *S2CPassCheckFast) String() string { return proto.CompactTextString(m) }
func (*S2CPassCheckFast) ProtoMessage()    {}
func (m *S2CPassCheckFast) GetId() uint16  { return PCK_S2CPassCheckFast_ID }

//
type ZF struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	State int32 `protobuf:"varint,2,opt,name=State,proto3" json:"State,omitempty"`
	//
	PetId []int32 `protobuf:"varint,3,rep,name=PetId,proto3" json:"PetId,omitempty"`
}

func (m *ZF) Reset()         { *m = ZF{} }
func (m *ZF) String() string { return proto.CompactTextString(m) }
func (*ZF) ProtoMessage()    {}

const PCK_C2SGetZF_ID = 3001 //
//
type C2SGetZF struct {
}

func (m *C2SGetZF) Reset()         { *m = C2SGetZF{} }
func (m *C2SGetZF) String() string { return proto.CompactTextString(m) }
func (*C2SGetZF) ProtoMessage()    {}
func (m *C2SGetZF) GetId() uint16  { return PCK_C2SGetZF_ID }

const PCK_S2CGetZF_ID = 3002 //
//
type S2CGetZF struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	ZF []*ZF `protobuf:"bytes,2,rep,name=ZF,proto3" json:"ZF,omitempty"`
}

func (m *S2CGetZF) Reset()         { *m = S2CGetZF{} }
func (m *S2CGetZF) String() string { return proto.CompactTextString(m) }
func (*S2CGetZF) ProtoMessage()    {}
func (m *S2CGetZF) GetId() uint16  { return PCK_S2CGetZF_ID }

const PCK_C2SZFPetUp_ID = 3003 //
//
type C2SZFPetUp struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Pos int32 `protobuf:"varint,2,opt,name=Pos,proto3" json:"Pos,omitempty"`
	//
	PetId int32 `protobuf:"varint,3,opt,name=PetId,proto3" json:"PetId,omitempty"`
}

func (m *C2SZFPetUp) Reset()         { *m = C2SZFPetUp{} }
func (m *C2SZFPetUp) String() string { return proto.CompactTextString(m) }
func (*C2SZFPetUp) ProtoMessage()    {}
func (m *C2SZFPetUp) GetId() uint16  { return PCK_C2SZFPetUp_ID }

const PCK_S2CZFPetUp_ID = 3004 //
//
type S2CZFPetUp struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Pos int32 `protobuf:"varint,3,opt,name=Pos,proto3" json:"Pos,omitempty"`
	//
	PetId int32 `protobuf:"varint,4,opt,name=PetId,proto3" json:"PetId,omitempty"`
}

func (m *S2CZFPetUp) Reset()         { *m = S2CZFPetUp{} }
func (m *S2CZFPetUp) String() string { return proto.CompactTextString(m) }
func (*S2CZFPetUp) ProtoMessage()    {}
func (m *S2CZFPetUp) GetId() uint16  { return PCK_S2CZFPetUp_ID }

const PCK_C2SZFState_ID = 3005 //
//
type C2SZFState struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	State int32 `protobuf:"varint,2,opt,name=State,proto3" json:"State,omitempty"`
}

func (m *C2SZFState) Reset()         { *m = C2SZFState{} }
func (m *C2SZFState) String() string { return proto.CompactTextString(m) }
func (*C2SZFState) ProtoMessage()    {}
func (m *C2SZFState) GetId() uint16  { return PCK_C2SZFState_ID }

const PCK_S2CZFState_ID = 3006 //
//
type S2CZFState struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	State int32 `protobuf:"varint,3,opt,name=State,proto3" json:"State,omitempty"`
}

func (m *S2CZFState) Reset()         { *m = S2CZFState{} }
func (m *S2CZFState) String() string { return proto.CompactTextString(m) }
func (*S2CZFState) ProtoMessage()    {}
func (m *S2CZFState) GetId() uint16  { return PCK_S2CZFState_ID }

const PCK_C2SZFUnlock_ID = 3104 //
//
type C2SZFUnlock struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2SZFUnlock) Reset()         { *m = C2SZFUnlock{} }
func (m *C2SZFUnlock) String() string { return proto.CompactTextString(m) }
func (*C2SZFUnlock) ProtoMessage()    {}
func (m *C2SZFUnlock) GetId() uint16  { return PCK_C2SZFUnlock_ID }

const PCK_S2CZFUnlock_ID = 3105 //
//
type S2CZFUnlock struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *S2CZFUnlock) Reset()         { *m = S2CZFUnlock{} }
func (m *S2CZFUnlock) String() string { return proto.CompactTextString(m) }
func (*S2CZFUnlock) ProtoMessage()    {}
func (m *S2CZFUnlock) GetId() uint16  { return PCK_S2CZFUnlock_ID }

//
type RankDataItem struct {
	//
	RankKey int32 `protobuf:"varint,1,opt,name=RankKey,proto3" json:"RankKey,omitempty"`
	//
	UserId int64 `protobuf:"varint,2,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	Value1 int64 `protobuf:"varint,3,opt,name=Value1,proto3" json:"Value1,omitempty"`
	//
	Value2 int64 `protobuf:"varint,4,opt,name=Value2,proto3" json:"Value2,omitempty"`
	//
	Ext int64 `protobuf:"varint,5,opt,name=Ext,proto3" json:"Ext,omitempty"`
}

func (m *RankDataItem) Reset()         { *m = RankDataItem{} }
func (m *RankDataItem) String() string { return proto.CompactTextString(m) }
func (*RankDataItem) ProtoMessage()    {}

//
type RankDbData struct {
	//
	RankData []*RankDataItem `protobuf:"bytes,1,rep,name=RankData,proto3" json:"RankData,omitempty"`
}

func (m *RankDbData) Reset()         { *m = RankDbData{} }
func (m *RankDbData) String() string { return proto.CompactTextString(m) }
func (*RankDbData) ProtoMessage()    {}

const PCK_K2SNewEvent_ID = 40 //
//
type K2SNewEvent struct {
	//
	ServiceType int32 `protobuf:"varint,1,opt,name=ServiceType,proto3" json:"ServiceType,omitempty"`
	//
	Count int64 `protobuf:"varint,2,opt,name=Count,proto3" json:"Count,omitempty"`
	//
	Param2 int32 `protobuf:"varint,3,opt,name=Param2,proto3" json:"Param2,omitempty"`
	//
	Param3 int32 `protobuf:"varint,4,opt,name=Param3,proto3" json:"Param3,omitempty"`
}

func (m *K2SNewEvent) Reset()         { *m = K2SNewEvent{} }
func (m *K2SNewEvent) String() string { return proto.CompactTextString(m) }
func (*K2SNewEvent) ProtoMessage()    {}
func (m *K2SNewEvent) GetId() uint16  { return PCK_K2SNewEvent_ID }

const PCK_C2SFuncOpen_ID = 3101 //
//
type C2SFuncOpen struct {
	//
	FuncId int32 `protobuf:"varint,1,opt,name=FuncId,proto3" json:"FuncId,omitempty"`
}

func (m *C2SFuncOpen) Reset()         { *m = C2SFuncOpen{} }
func (m *C2SFuncOpen) String() string { return proto.CompactTextString(m) }
func (*C2SFuncOpen) ProtoMessage()    {}
func (m *C2SFuncOpen) GetId() uint16  { return PCK_C2SFuncOpen_ID }

const PCK_S2CFuncOpen_ID = 3102 //
//
type S2CFuncOpen struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CFuncOpen) Reset()         { *m = S2CFuncOpen{} }
func (m *S2CFuncOpen) String() string { return proto.CompactTextString(m) }
func (*S2CFuncOpen) ProtoMessage()    {}
func (m *S2CFuncOpen) GetId() uint16  { return PCK_S2CFuncOpen_ID }

const PCK_C2SRobotZF_ID = 3103 //
//
type C2SRobotZF struct {
}

func (m *C2SRobotZF) Reset()         { *m = C2SRobotZF{} }
func (m *C2SRobotZF) String() string { return proto.CompactTextString(m) }
func (*C2SRobotZF) ProtoMessage()    {}
func (m *C2SRobotZF) GetId() uint16  { return PCK_C2SRobotZF_ID }

//
type PhotoBook struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	L int32 `protobuf:"varint,2,opt,name=L,proto3" json:"L,omitempty"`
}

func (m *PhotoBook) Reset()         { *m = PhotoBook{} }
func (m *PhotoBook) String() string { return proto.CompactTextString(m) }
func (*PhotoBook) ProtoMessage()    {}

const PCK_S2CUserPhotoBook_ID = 4101 //
//
type S2CUserPhotoBook struct {
	//
	T []*PhotoBook `protobuf:"bytes,1,rep,name=T,proto3" json:"T,omitempty"`
}

func (m *S2CUserPhotoBook) Reset()         { *m = S2CUserPhotoBook{} }
func (m *S2CUserPhotoBook) String() string { return proto.CompactTextString(m) }
func (*S2CUserPhotoBook) ProtoMessage()    {}
func (m *S2CUserPhotoBook) GetId() uint16  { return PCK_S2CUserPhotoBook_ID }

const PCK_C2SPhotoBookLevelUp_ID = 4102 //
//
type C2SPhotoBookLevelUp struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2SPhotoBookLevelUp) Reset()         { *m = C2SPhotoBookLevelUp{} }
func (m *C2SPhotoBookLevelUp) String() string { return proto.CompactTextString(m) }
func (*C2SPhotoBookLevelUp) ProtoMessage()    {}
func (m *C2SPhotoBookLevelUp) GetId() uint16  { return PCK_C2SPhotoBookLevelUp_ID }

const PCK_S2CPhotoBookLevelUp_ID = 4103 //
//
type S2CPhotoBookLevelUp struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	L int32 `protobuf:"varint,3,opt,name=L,proto3" json:"L,omitempty"`
}

func (m *S2CPhotoBookLevelUp) Reset()         { *m = S2CPhotoBookLevelUp{} }
func (m *S2CPhotoBookLevelUp) String() string { return proto.CompactTextString(m) }
func (*S2CPhotoBookLevelUp) ProtoMessage()    {}
func (m *S2CPhotoBookLevelUp) GetId() uint16  { return PCK_S2CPhotoBookLevelUp_ID }

const PCK_C2SPhotoBookDel_ID = 4104 //
//
type C2SPhotoBookDel struct {
}

func (m *C2SPhotoBookDel) Reset()         { *m = C2SPhotoBookDel{} }
func (m *C2SPhotoBookDel) String() string { return proto.CompactTextString(m) }
func (*C2SPhotoBookDel) ProtoMessage()    {}
func (m *C2SPhotoBookDel) GetId() uint16  { return PCK_C2SPhotoBookDel_ID }

const PCK_S2CPhotoBookDel_ID = 4105 //
//
type S2CPhotoBookDel struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CPhotoBookDel) Reset()         { *m = S2CPhotoBookDel{} }
func (m *S2CPhotoBookDel) String() string { return proto.CompactTextString(m) }
func (*S2CPhotoBookDel) ProtoMessage()    {}
func (m *S2CPhotoBookDel) GetId() uint16  { return PCK_S2CPhotoBookDel_ID }

const PCK_C2SBuyChargeGift_ID = 17300 //
//
type C2SBuyChargeGift struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2SBuyChargeGift) Reset()         { *m = C2SBuyChargeGift{} }
func (m *C2SBuyChargeGift) String() string { return proto.CompactTextString(m) }
func (*C2SBuyChargeGift) ProtoMessage()    {}
func (m *C2SBuyChargeGift) GetId() uint16  { return PCK_C2SBuyChargeGift_ID }

const PCK_S2CBuyChargeGift_ID = 17301 //
//
type S2CBuyChargeGift struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *S2CBuyChargeGift) Reset()         { *m = S2CBuyChargeGift{} }
func (m *S2CBuyChargeGift) String() string { return proto.CompactTextString(m) }
func (*S2CBuyChargeGift) ProtoMessage()    {}
func (m *S2CBuyChargeGift) GetId() uint16  { return PCK_S2CBuyChargeGift_ID }

const PCK_S2CAllChargeGift_ID = 17302 //
//
type S2CAllChargeGift struct {
	//
	List []int32 `protobuf:"varint,1,rep,name=List,proto3" json:"List,omitempty"`
}

func (m *S2CAllChargeGift) Reset()         { *m = S2CAllChargeGift{} }
func (m *S2CAllChargeGift) String() string { return proto.CompactTextString(m) }
func (*S2CAllChargeGift) ProtoMessage()    {}
func (m *S2CAllChargeGift) GetId() uint16  { return PCK_S2CAllChargeGift_ID }

const PCK_S2CVipService_ID = 12106 //
//
type S2CVipService struct {
	//
	QQ string `protobuf:"bytes,1,opt,name=QQ,proto3" json:"QQ,omitempty"`
}

func (m *S2CVipService) Reset()         { *m = S2CVipService{} }
func (m *S2CVipService) String() string { return proto.CompactTextString(m) }
func (*S2CVipService) ProtoMessage()    {}
func (m *S2CVipService) GetId() uint16  { return PCK_S2CVipService_ID }

const PCK_C2SNewPlayerZF_ID = 99 //
//
type C2SNewPlayerZF struct {
}

func (m *C2SNewPlayerZF) Reset()         { *m = C2SNewPlayerZF{} }
func (m *C2SNewPlayerZF) String() string { return proto.CompactTextString(m) }
func (*C2SNewPlayerZF) ProtoMessage()    {}
func (m *C2SNewPlayerZF) GetId() uint16  { return PCK_C2SNewPlayerZF_ID }

const PCK_C2SGetStateDemons_ID = 17403 //
//
type C2SGetStateDemons struct {
}

func (m *C2SGetStateDemons) Reset()         { *m = C2SGetStateDemons{} }
func (m *C2SGetStateDemons) String() string { return proto.CompactTextString(m) }
func (*C2SGetStateDemons) ProtoMessage()    {}
func (m *C2SGetStateDemons) GetId() uint16  { return PCK_C2SGetStateDemons_ID }

const PCK_S2CGetStateDemons_ID = 17404 //
//
type S2CGetStateDemons struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CGetStateDemons) Reset()         { *m = S2CGetStateDemons{} }
func (m *S2CGetStateDemons) String() string { return proto.CompactTextString(m) }
func (*S2CGetStateDemons) ProtoMessage()    {}
func (m *S2CGetStateDemons) GetId() uint16  { return PCK_S2CGetStateDemons_ID }

const PCK_C2SStateBreach_ID = 17405 //
//
type C2SStateBreach struct {
}

func (m *C2SStateBreach) Reset()         { *m = C2SStateBreach{} }
func (m *C2SStateBreach) String() string { return proto.CompactTextString(m) }
func (*C2SStateBreach) ProtoMessage()    {}
func (m *C2SStateBreach) GetId() uint16  { return PCK_C2SStateBreach_ID }

const PCK_S2CStateBreach_ID = 17406 //
//
type S2CStateBreach struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CStateBreach) Reset()         { *m = S2CStateBreach{} }
func (m *S2CStateBreach) String() string { return proto.CompactTextString(m) }
func (*S2CStateBreach) ProtoMessage()    {}
func (m *S2CStateBreach) GetId() uint16  { return PCK_S2CStateBreach_ID }

const PCK_C2SStateUsePill_ID = 17407 //
//
type C2SStateUsePill struct {
}

func (m *C2SStateUsePill) Reset()         { *m = C2SStateUsePill{} }
func (m *C2SStateUsePill) String() string { return proto.CompactTextString(m) }
func (*C2SStateUsePill) ProtoMessage()    {}
func (m *C2SStateUsePill) GetId() uint16  { return PCK_C2SStateUsePill_ID }

const PCK_S2CStateUsePill_ID = 17408 //
//
type S2CStateUsePill struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CStateUsePill) Reset()         { *m = S2CStateUsePill{} }
func (m *S2CStateUsePill) String() string { return proto.CompactTextString(m) }
func (*S2CStateUsePill) ProtoMessage()    {}
func (m *S2CStateUsePill) GetId() uint16  { return PCK_S2CStateUsePill_ID }

const PCK_C2SBuyStateGift_ID = 17409 //
//
type C2SBuyStateGift struct {
	//
	Gift int32 `protobuf:"varint,1,opt,name=Gift,proto3" json:"Gift,omitempty"`
}

func (m *C2SBuyStateGift) Reset()         { *m = C2SBuyStateGift{} }
func (m *C2SBuyStateGift) String() string { return proto.CompactTextString(m) }
func (*C2SBuyStateGift) ProtoMessage()    {}
func (m *C2SBuyStateGift) GetId() uint16  { return PCK_C2SBuyStateGift_ID }

const PCK_S2CBuyStateGift_ID = 17410 //
//
type S2CBuyStateGift struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Gift int32 `protobuf:"varint,2,opt,name=Gift,proto3" json:"Gift,omitempty"`
}

func (m *S2CBuyStateGift) Reset()         { *m = S2CBuyStateGift{} }
func (m *S2CBuyStateGift) String() string { return proto.CompactTextString(m) }
func (*S2CBuyStateGift) ProtoMessage()    {}
func (m *S2CBuyStateGift) GetId() uint16  { return PCK_S2CBuyStateGift_ID }

const PCK_C2SEquipStarInherit_ID = 534 //
//
type C2SEquipStarInherit struct {
	//
	FromItemId string `protobuf:"bytes,1,opt,name=FromItemId,proto3" json:"FromItemId,omitempty"`
	//
	ToItemId string `protobuf:"bytes,2,opt,name=ToItemId,proto3" json:"ToItemId,omitempty"`
}

func (m *C2SEquipStarInherit) Reset()         { *m = C2SEquipStarInherit{} }
func (m *C2SEquipStarInherit) String() string { return proto.CompactTextString(m) }
func (*C2SEquipStarInherit) ProtoMessage()    {}
func (m *C2SEquipStarInherit) GetId() uint16  { return PCK_C2SEquipStarInherit_ID }

const PCK_S2CEquipStarInherit_ID = 535 //
//
type S2CEquipStarInherit struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CEquipStarInherit) Reset()         { *m = S2CEquipStarInherit{} }
func (m *S2CEquipStarInherit) String() string { return proto.CompactTextString(m) }
func (*S2CEquipStarInherit) ProtoMessage()    {}
func (m *S2CEquipStarInherit) GetId() uint16  { return PCK_S2CEquipStarInherit_ID }

const PCK_C2SEnterInstance_ID = 18001 //
//
type C2SEnterInstance struct {
	//
	InstanceType int32 `protobuf:"varint,1,opt,name=InstanceType,proto3" json:"InstanceType,omitempty"`
	//
	Param int32 `protobuf:"varint,2,opt,name=Param,proto3" json:"Param,omitempty"`
}

func (m *C2SEnterInstance) Reset()         { *m = C2SEnterInstance{} }
func (m *C2SEnterInstance) String() string { return proto.CompactTextString(m) }
func (*C2SEnterInstance) ProtoMessage()    {}
func (m *C2SEnterInstance) GetId() uint16  { return PCK_C2SEnterInstance_ID }

const PCK_S2CEnterInstance_ID = 18002 //
//
type S2CEnterInstance struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	InstanceType int32 `protobuf:"varint,2,opt,name=InstanceType,proto3" json:"InstanceType,omitempty"`
	//
	Param int32 `protobuf:"varint,3,opt,name=Param,proto3" json:"Param,omitempty"`
}

func (m *S2CEnterInstance) Reset()         { *m = S2CEnterInstance{} }
func (m *S2CEnterInstance) String() string { return proto.CompactTextString(m) }
func (*S2CEnterInstance) ProtoMessage()    {}
func (m *S2CEnterInstance) GetId() uint16  { return PCK_S2CEnterInstance_ID }

const PCK_S2CInstanceData_ID = 18003 //
//
type S2CInstanceData struct {
	//
	MapId int64 `protobuf:"varint,1,opt,name=MapId,proto3" json:"MapId,omitempty"`
	//
	Attr []*IntAttr `protobuf:"bytes,2,rep,name=Attr,proto3" json:"Attr,omitempty"`
}

func (m *S2CInstanceData) Reset()         { *m = S2CInstanceData{} }
func (m *S2CInstanceData) String() string { return proto.CompactTextString(m) }
func (*S2CInstanceData) ProtoMessage()    {}
func (m *S2CInstanceData) GetId() uint16  { return PCK_S2CInstanceData_ID }

const PCK_S2CInstanceSettlement_ID = 18004 //
//
type S2CInstanceSettlement struct {
	//
	MapId int64 `protobuf:"varint,1,opt,name=MapId,proto3" json:"MapId,omitempty"`
	//
	Attr []*IntAttr `protobuf:"bytes,2,rep,name=Attr,proto3" json:"Attr,omitempty"`
	//
	Items []*ItemData `protobuf:"bytes,3,rep,name=Items,proto3" json:"Items,omitempty"`
}

func (m *S2CInstanceSettlement) Reset()         { *m = S2CInstanceSettlement{} }
func (m *S2CInstanceSettlement) String() string { return proto.CompactTextString(m) }
func (*S2CInstanceSettlement) ProtoMessage()    {}
func (m *S2CInstanceSettlement) GetId() uint16  { return PCK_S2CInstanceSettlement_ID }

const PCK_S2CInstanceExpData_ID = 19001 //
//
type S2CInstanceExpData struct {
	//
	LeftTime int32 `protobuf:"varint,1,opt,name=LeftTime,proto3" json:"LeftTime,omitempty"`
	//
	LeftBuyTime int32 `protobuf:"varint,2,opt,name=LeftBuyTime,proto3" json:"LeftBuyTime,omitempty"`
	//
	NextVipLevel int32 `protobuf:"varint,3,opt,name=NextVipLevel,proto3" json:"NextVipLevel,omitempty"`
	//
	TotalBuyTimes int32 `protobuf:"varint,4,opt,name=TotalBuyTimes,proto3" json:"TotalBuyTimes,omitempty"`
	//
	DiffBuyTimes int32 `protobuf:"varint,5,opt,name=DiffBuyTimes,proto3" json:"DiffBuyTimes,omitempty"`
	//
	InstanceEndTime int64 `protobuf:"varint,6,opt,name=InstanceEndTime,proto3" json:"InstanceEndTime,omitempty"`
	//
	BuyCoin4 int64 `protobuf:"varint,7,opt,name=BuyCoin4,proto3" json:"BuyCoin4,omitempty"`
}

func (m *S2CInstanceExpData) Reset()         { *m = S2CInstanceExpData{} }
func (m *S2CInstanceExpData) String() string { return proto.CompactTextString(m) }
func (*S2CInstanceExpData) ProtoMessage()    {}
func (m *S2CInstanceExpData) GetId() uint16  { return PCK_S2CInstanceExpData_ID }

const PCK_C2SInstanceExpBuyTimes_ID = 19002 //
//
type C2SInstanceExpBuyTimes struct {
}

func (m *C2SInstanceExpBuyTimes) Reset()         { *m = C2SInstanceExpBuyTimes{} }
func (m *C2SInstanceExpBuyTimes) String() string { return proto.CompactTextString(m) }
func (*C2SInstanceExpBuyTimes) ProtoMessage()    {}
func (m *C2SInstanceExpBuyTimes) GetId() uint16  { return PCK_C2SInstanceExpBuyTimes_ID }

const PCK_S2CInstanceExpBuyTimes_ID = 19003 //
//
type S2CInstanceExpBuyTimes struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CInstanceExpBuyTimes) Reset()         { *m = S2CInstanceExpBuyTimes{} }
func (m *S2CInstanceExpBuyTimes) String() string { return proto.CompactTextString(m) }
func (*S2CInstanceExpBuyTimes) ProtoMessage()    {}
func (m *S2CInstanceExpBuyTimes) GetId() uint16  { return PCK_S2CInstanceExpBuyTimes_ID }

//
type HorseEquipCollect struct {
	//
	ID int32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	//
	CollectId []int32 `protobuf:"varint,2,rep,name=CollectId,proto3" json:"CollectId,omitempty"`
}

func (m *HorseEquipCollect) Reset()         { *m = HorseEquipCollect{} }
func (m *HorseEquipCollect) String() string { return proto.CompactTextString(m) }
func (*HorseEquipCollect) ProtoMessage()    {}

//
type ZhenFa struct {
	//
	ZhenFaId int32 `protobuf:"varint,1,opt,name=ZhenFaId,proto3" json:"ZhenFaId,omitempty"`
	//
	Level int32 `protobuf:"varint,2,opt,name=Level,proto3" json:"Level,omitempty"`
	//
	IsUse int32 `protobuf:"varint,3,opt,name=IsUse,proto3" json:"IsUse,omitempty"`
	//
	ZhenYanSkillId int32 `protobuf:"varint,4,opt,name=ZhenYanSkillId,proto3" json:"ZhenYanSkillId,omitempty"`
	//
	JiBanId int32 `protobuf:"varint,5,opt,name=JiBanId,proto3" json:"JiBanId,omitempty"`
}

func (m *ZhenFa) Reset()         { *m = ZhenFa{} }
func (m *ZhenFa) String() string { return proto.CompactTextString(m) }
func (*ZhenFa) ProtoMessage()    {}

//
type Vip struct {
	//
	VipType int32 `protobuf:"varint,1,opt,name=VipType,proto3" json:"VipType,omitempty"`
	//
	IsForever int32 `protobuf:"varint,2,opt,name=IsForever,proto3" json:"IsForever,omitempty"`
	//
	LimitTime int64 `protobuf:"varint,3,opt,name=LimitTime,proto3" json:"LimitTime,omitempty"`
	//
	MaxVipLevel int32 `protobuf:"varint,4,opt,name=MaxVipLevel,proto3" json:"MaxVipLevel,omitempty"`
	//
	PrizeGet []int32 `protobuf:"varint,5,rep,name=PrizeGet,proto3" json:"PrizeGet,omitempty"`
	//
	GiftBuy []int32 `protobuf:"varint,6,rep,name=GiftBuy,proto3" json:"GiftBuy,omitempty"`
}

func (m *Vip) Reset()         { *m = Vip{} }
func (m *Vip) String() string { return proto.CompactTextString(m) }
func (*Vip) ProtoMessage()    {}

//
type TianShu struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Active int32 `protobuf:"varint,2,opt,name=Active,proto3" json:"Active,omitempty"`
}

func (m *TianShu) Reset()         { *m = TianShu{} }
func (m *TianShu) String() string { return proto.CompactTextString(m) }
func (*TianShu) ProtoMessage()    {}

//
type JiBan struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Actived int32 `protobuf:"varint,2,opt,name=Actived,proto3" json:"Actived,omitempty"`
	//
	Equiped int32 `protobuf:"varint,3,opt,name=Equiped,proto3" json:"Equiped,omitempty"`
	//
	Partners []*CommonType `protobuf:"bytes,4,rep,name=Partners,proto3" json:"Partners,omitempty"`
}

func (m *JiBan) Reset()         { *m = JiBan{} }
func (m *JiBan) String() string { return proto.CompactTextString(m) }
func (*JiBan) ProtoMessage()    {}

//
type EquipGem struct {
	//
	EquipPart int32 `protobuf:"varint,1,opt,name=EquipPart,proto3" json:"EquipPart,omitempty"`
	//
	GemCell int32 `protobuf:"varint,2,opt,name=GemCell,proto3" json:"GemCell,omitempty"`
	//
	GemID int32 `protobuf:"varint,3,opt,name=GemID,proto3" json:"GemID,omitempty"`
}

func (m *EquipGem) Reset()         { *m = EquipGem{} }
func (m *EquipGem) String() string { return proto.CompactTextString(m) }
func (*EquipGem) ProtoMessage()    {}

const PCK_C2SGemUpgrade_ID = 550 //
//
type C2SGemUpgrade struct {
	//
	TargetGemID int32 `protobuf:"varint,1,opt,name=TargetGemID,proto3" json:"TargetGemID,omitempty"`
	//
	Num int32 `protobuf:"varint,2,opt,name=Num,proto3" json:"Num,omitempty"`
}

func (m *C2SGemUpgrade) Reset()         { *m = C2SGemUpgrade{} }
func (m *C2SGemUpgrade) String() string { return proto.CompactTextString(m) }
func (*C2SGemUpgrade) ProtoMessage()    {}
func (m *C2SGemUpgrade) GetId() uint16  { return PCK_C2SGemUpgrade_ID }

const PCK_S2CGemUpgrade_ID = 551 //
//
type S2CGemUpgrade struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CGemUpgrade) Reset()         { *m = S2CGemUpgrade{} }
func (m *S2CGemUpgrade) String() string { return proto.CompactTextString(m) }
func (*S2CGemUpgrade) ProtoMessage()    {}
func (m *S2CGemUpgrade) GetId() uint16  { return PCK_S2CGemUpgrade_ID }

const PCK_C2SGemAutoEquip_ID = 552 //
//
type C2SGemAutoEquip struct {
}

func (m *C2SGemAutoEquip) Reset()         { *m = C2SGemAutoEquip{} }
func (m *C2SGemAutoEquip) String() string { return proto.CompactTextString(m) }
func (*C2SGemAutoEquip) ProtoMessage()    {}
func (m *C2SGemAutoEquip) GetId() uint16  { return PCK_C2SGemAutoEquip_ID }

const PCK_S2CGemAutoEquip_ID = 553 //
//
type S2CGemAutoEquip struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CGemAutoEquip) Reset()         { *m = S2CGemAutoEquip{} }
func (m *S2CGemAutoEquip) String() string { return proto.CompactTextString(m) }
func (*S2CGemAutoEquip) ProtoMessage()    {}
func (m *S2CGemAutoEquip) GetId() uint16  { return PCK_S2CGemAutoEquip_ID }

const PCK_C2SGemAutoUpgradeEquip_ID = 554 //
//
type C2SGemAutoUpgradeEquip struct {
}

func (m *C2SGemAutoUpgradeEquip) Reset()         { *m = C2SGemAutoUpgradeEquip{} }
func (m *C2SGemAutoUpgradeEquip) String() string { return proto.CompactTextString(m) }
func (*C2SGemAutoUpgradeEquip) ProtoMessage()    {}
func (m *C2SGemAutoUpgradeEquip) GetId() uint16  { return PCK_C2SGemAutoUpgradeEquip_ID }

const PCK_S2CGemAutoUpgradeEquip_ID = 555 //
//
type S2CGemAutoUpgradeEquip struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CGemAutoUpgradeEquip) Reset()         { *m = S2CGemAutoUpgradeEquip{} }
func (m *S2CGemAutoUpgradeEquip) String() string { return proto.CompactTextString(m) }
func (*S2CGemAutoUpgradeEquip) ProtoMessage()    {}
func (m *S2CGemAutoUpgradeEquip) GetId() uint16  { return PCK_S2CGemAutoUpgradeEquip_ID }

//
type BecomeGodTask struct {
	//
	GroupId int32 `protobuf:"varint,1,opt,name=GroupId,proto3" json:"GroupId,omitempty"`
	//
	TaskId int32 `protobuf:"varint,2,opt,name=TaskId,proto3" json:"TaskId,omitempty"`
	//
	Value int64 `protobuf:"varint,3,opt,name=Value,proto3" json:"Value,omitempty"`
	//
	State int32 `protobuf:"varint,4,opt,name=State,proto3" json:"State,omitempty"`
}

func (m *BecomeGodTask) Reset()         { *m = BecomeGodTask{} }
func (m *BecomeGodTask) String() string { return proto.CompactTextString(m) }
func (*BecomeGodTask) ProtoMessage()    {}

const PCK_S2CBecomeGodTask_ID = 19101 //
//
type S2CBecomeGodTask struct {
	//
	Tasks []*BecomeGodTask `protobuf:"bytes,1,rep,name=Tasks,proto3" json:"Tasks,omitempty"`
}

func (m *S2CBecomeGodTask) Reset()         { *m = S2CBecomeGodTask{} }
func (m *S2CBecomeGodTask) String() string { return proto.CompactTextString(m) }
func (*S2CBecomeGodTask) ProtoMessage()    {}
func (m *S2CBecomeGodTask) GetId() uint16  { return PCK_S2CBecomeGodTask_ID }

const PCK_C2SGetBecomeGodTaskPrize_ID = 19102 //
//
type C2SGetBecomeGodTaskPrize struct {
	//
	GroupId int32 `protobuf:"varint,1,opt,name=GroupId,proto3" json:"GroupId,omitempty"`
}

func (m *C2SGetBecomeGodTaskPrize) Reset()         { *m = C2SGetBecomeGodTaskPrize{} }
func (m *C2SGetBecomeGodTaskPrize) String() string { return proto.CompactTextString(m) }
func (*C2SGetBecomeGodTaskPrize) ProtoMessage()    {}
func (m *C2SGetBecomeGodTaskPrize) GetId() uint16  { return PCK_C2SGetBecomeGodTaskPrize_ID }

const PCK_S2CGetBecomeGodTaskPrize_ID = 19103 //
//
type S2CGetBecomeGodTaskPrize struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CGetBecomeGodTaskPrize) Reset()         { *m = S2CGetBecomeGodTaskPrize{} }
func (m *S2CGetBecomeGodTaskPrize) String() string { return proto.CompactTextString(m) }
func (*S2CGetBecomeGodTaskPrize) ProtoMessage()    {}
func (m *S2CGetBecomeGodTaskPrize) GetId() uint16  { return PCK_S2CGetBecomeGodTaskPrize_ID }

const PCK_C2SEquipSpecial_ID = 4151 //
//
type C2SEquipSpecial struct {
	//
	ItemID int32 `protobuf:"varint,1,opt,name=ItemID,proto3" json:"ItemID,omitempty"`
	//
	ActionType int32 `protobuf:"varint,2,opt,name=ActionType,proto3" json:"ActionType,omitempty"`
}

func (m *C2SEquipSpecial) Reset()         { *m = C2SEquipSpecial{} }
func (m *C2SEquipSpecial) String() string { return proto.CompactTextString(m) }
func (*C2SEquipSpecial) ProtoMessage()    {}
func (m *C2SEquipSpecial) GetId() uint16  { return PCK_C2SEquipSpecial_ID }

const PCK_S2CEquipSpecial_ID = 4152 //
//
type S2CEquipSpecial struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	ItemID int32 `protobuf:"varint,2,opt,name=ItemID,proto3" json:"ItemID,omitempty"`
	//
	ActionType int32 `protobuf:"varint,3,opt,name=ActionType,proto3" json:"ActionType,omitempty"`
}

func (m *S2CEquipSpecial) Reset()         { *m = S2CEquipSpecial{} }
func (m *S2CEquipSpecial) String() string { return proto.CompactTextString(m) }
func (*S2CEquipSpecial) ProtoMessage()    {}
func (m *S2CEquipSpecial) GetId() uint16  { return PCK_S2CEquipSpecial_ID }

const PCK_C2SGetEquipSpecialReward_ID = 4253 //
//
type C2SGetEquipSpecialReward struct {
	//
	ItemID int32 `protobuf:"varint,1,opt,name=ItemID,proto3" json:"ItemID,omitempty"`
}

func (m *C2SGetEquipSpecialReward) Reset()         { *m = C2SGetEquipSpecialReward{} }
func (m *C2SGetEquipSpecialReward) String() string { return proto.CompactTextString(m) }
func (*C2SGetEquipSpecialReward) ProtoMessage()    {}
func (m *C2SGetEquipSpecialReward) GetId() uint16  { return PCK_C2SGetEquipSpecialReward_ID }

const PCK_S2CGetEquipSpecialReward_ID = 4254 //
//
type S2CGetEquipSpecialReward struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	ItemID int32 `protobuf:"varint,2,opt,name=ItemID,proto3" json:"ItemID,omitempty"`
}

func (m *S2CGetEquipSpecialReward) Reset()         { *m = S2CGetEquipSpecialReward{} }
func (m *S2CGetEquipSpecialReward) String() string { return proto.CompactTextString(m) }
func (*S2CGetEquipSpecialReward) ProtoMessage()    {}
func (m *S2CGetEquipSpecialReward) GetId() uint16  { return PCK_S2CGetEquipSpecialReward_ID }

const PCK_S2CInstanceTongTianInfo_ID = 783 //
//
type S2CInstanceTongTianInfo struct {
	//
	Items []*TongTian `protobuf:"bytes,1,rep,name=Items,proto3" json:"Items,omitempty"`
	//
	TongTianBox []int32 `protobuf:"varint,2,rep,name=TongTianBox,proto3" json:"TongTianBox,omitempty"`
}

func (m *S2CInstanceTongTianInfo) Reset()         { *m = S2CInstanceTongTianInfo{} }
func (m *S2CInstanceTongTianInfo) String() string { return proto.CompactTextString(m) }
func (*S2CInstanceTongTianInfo) ProtoMessage()    {}
func (m *S2CInstanceTongTianInfo) GetId() uint16  { return PCK_S2CInstanceTongTianInfo_ID }

const PCK_C2SInstanceTongTianFight_ID = 784 //
//
type C2SInstanceTongTianFight struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2SInstanceTongTianFight) Reset()         { *m = C2SInstanceTongTianFight{} }
func (m *C2SInstanceTongTianFight) String() string { return proto.CompactTextString(m) }
func (*C2SInstanceTongTianFight) ProtoMessage()    {}
func (m *C2SInstanceTongTianFight) GetId() uint16  { return PCK_C2SInstanceTongTianFight_ID }

const PCK_S2CInstanceTongTianFight_ID = 785 //
//
type S2CInstanceTongTianFight struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Star int32 `protobuf:"varint,2,opt,name=Star,proto3" json:"Star,omitempty"`
	//
	Data *TongTian `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *S2CInstanceTongTianFight) Reset()         { *m = S2CInstanceTongTianFight{} }
func (m *S2CInstanceTongTianFight) String() string { return proto.CompactTextString(m) }
func (*S2CInstanceTongTianFight) ProtoMessage()    {}
func (m *S2CInstanceTongTianFight) GetId() uint16  { return PCK_S2CInstanceTongTianFight_ID }

const PCK_C2SInstanceTongTianSweep_ID = 786 //
//
type C2SInstanceTongTianSweep struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2SInstanceTongTianSweep) Reset()         { *m = C2SInstanceTongTianSweep{} }
func (m *C2SInstanceTongTianSweep) String() string { return proto.CompactTextString(m) }
func (*C2SInstanceTongTianSweep) ProtoMessage()    {}
func (m *C2SInstanceTongTianSweep) GetId() uint16  { return PCK_C2SInstanceTongTianSweep_ID }

const PCK_S2CInstanceTongTianSweep_ID = 787 //
//
type S2CInstanceTongTianSweep struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Data *TongTian `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *S2CInstanceTongTianSweep) Reset()         { *m = S2CInstanceTongTianSweep{} }
func (m *S2CInstanceTongTianSweep) String() string { return proto.CompactTextString(m) }
func (*S2CInstanceTongTianSweep) ProtoMessage()    {}
func (m *S2CInstanceTongTianSweep) GetId() uint16  { return PCK_S2CInstanceTongTianSweep_ID }

const PCK_C2SGetInstanceTongTianBox_ID = 788 //
//
type C2SGetInstanceTongTianBox struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	MapId int32 `protobuf:"varint,2,opt,name=MapId,proto3" json:"MapId,omitempty"`
}

func (m *C2SGetInstanceTongTianBox) Reset()         { *m = C2SGetInstanceTongTianBox{} }
func (m *C2SGetInstanceTongTianBox) String() string { return proto.CompactTextString(m) }
func (*C2SGetInstanceTongTianBox) ProtoMessage()    {}
func (m *C2SGetInstanceTongTianBox) GetId() uint16  { return PCK_C2SGetInstanceTongTianBox_ID }

const PCK_S2CGetInstanceTongTianBox_ID = 789 //
//
type S2CGetInstanceTongTianBox struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Ids []int32 `protobuf:"varint,2,rep,name=Ids,proto3" json:"Ids,omitempty"`
}

func (m *S2CGetInstanceTongTianBox) Reset()         { *m = S2CGetInstanceTongTianBox{} }
func (m *S2CGetInstanceTongTianBox) String() string { return proto.CompactTextString(m) }
func (*S2CGetInstanceTongTianBox) ProtoMessage()    {}
func (m *S2CGetInstanceTongTianBox) GetId() uint16  { return PCK_S2CGetInstanceTongTianBox_ID }

const PCK_C2SGetInstanceTongTianProcessBox_ID = 790 //
//
type C2SGetInstanceTongTianProcessBox struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2SGetInstanceTongTianProcessBox) Reset()         { *m = C2SGetInstanceTongTianProcessBox{} }
func (m *C2SGetInstanceTongTianProcessBox) String() string { return proto.CompactTextString(m) }
func (*C2SGetInstanceTongTianProcessBox) ProtoMessage()    {}
func (m *C2SGetInstanceTongTianProcessBox) GetId() uint16 {
	return PCK_C2SGetInstanceTongTianProcessBox_ID
}

const PCK_S2CGetInstanceTongTianProcessBox_ID = 791 //
//
type S2CGetInstanceTongTianProcessBox struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Data *TongTian `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *S2CGetInstanceTongTianProcessBox) Reset()         { *m = S2CGetInstanceTongTianProcessBox{} }
func (m *S2CGetInstanceTongTianProcessBox) String() string { return proto.CompactTextString(m) }
func (*S2CGetInstanceTongTianProcessBox) ProtoMessage()    {}
func (m *S2CGetInstanceTongTianProcessBox) GetId() uint16 {
	return PCK_S2CGetInstanceTongTianProcessBox_ID
}

const PCK_C2SInstanceTongTianReset_ID = 792 //
//
type C2SInstanceTongTianReset struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2SInstanceTongTianReset) Reset()         { *m = C2SInstanceTongTianReset{} }
func (m *C2SInstanceTongTianReset) String() string { return proto.CompactTextString(m) }
func (*C2SInstanceTongTianReset) ProtoMessage()    {}
func (m *C2SInstanceTongTianReset) GetId() uint16  { return PCK_C2SInstanceTongTianReset_ID }

const PCK_S2CInstanceTongTianReset_ID = 793 //
//
type S2CInstanceTongTianReset struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Data *TongTian `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *S2CInstanceTongTianReset) Reset()         { *m = S2CInstanceTongTianReset{} }
func (m *S2CInstanceTongTianReset) String() string { return proto.CompactTextString(m) }
func (*S2CInstanceTongTianReset) ProtoMessage()    {}
func (m *S2CInstanceTongTianReset) GetId() uint16  { return PCK_S2CInstanceTongTianReset_ID }

const PCK_C2SHorseEquipStarUp_ID = 3151 //
//
type C2SHorseEquipStarUp struct {
	//
	SelectItemID string `protobuf:"bytes,1,opt,name=SelectItemID,proto3" json:"SelectItemID,omitempty"`
	//
	CostItemIds []string `protobuf:"bytes,2,rep,name=CostItemIds,proto3" json:"CostItemIds,omitempty"`
}

func (m *C2SHorseEquipStarUp) Reset()         { *m = C2SHorseEquipStarUp{} }
func (m *C2SHorseEquipStarUp) String() string { return proto.CompactTextString(m) }
func (*C2SHorseEquipStarUp) ProtoMessage()    {}
func (m *C2SHorseEquipStarUp) GetId() uint16  { return PCK_C2SHorseEquipStarUp_ID }

const PCK_S2CHorseEquipStarUp_ID = 3152 //
//
type S2CHorseEquipStarUp struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	SelectItemID string `protobuf:"bytes,2,opt,name=SelectItemID,proto3" json:"SelectItemID,omitempty"`
	//
	GetItemID string `protobuf:"bytes,3,opt,name=GetItemID,proto3" json:"GetItemID,omitempty"`
}

func (m *S2CHorseEquipStarUp) Reset()         { *m = S2CHorseEquipStarUp{} }
func (m *S2CHorseEquipStarUp) String() string { return proto.CompactTextString(m) }
func (*S2CHorseEquipStarUp) ProtoMessage()    {}
func (m *S2CHorseEquipStarUp) GetId() uint16  { return PCK_S2CHorseEquipStarUp_ID }

const PCK_C2SStartLockTower_ID = 19201 //
//
type C2SStartLockTower struct {
	//
	Sweep int32 `protobuf:"varint,1,opt,name=Sweep,proto3" json:"Sweep,omitempty"`
}

func (m *C2SStartLockTower) Reset()         { *m = C2SStartLockTower{} }
func (m *C2SStartLockTower) String() string { return proto.CompactTextString(m) }
func (*C2SStartLockTower) ProtoMessage()    {}
func (m *C2SStartLockTower) GetId() uint16  { return PCK_C2SStartLockTower_ID }

const PCK_S2CStartLockTower_ID = 19202 //
//
type S2CStartLockTower struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CStartLockTower) Reset()         { *m = S2CStartLockTower{} }
func (m *S2CStartLockTower) String() string { return proto.CompactTextString(m) }
func (*S2CStartLockTower) ProtoMessage()    {}
func (m *S2CStartLockTower) GetId() uint16  { return PCK_S2CStartLockTower_ID }

const PCK_C2SGirlHillStart_ID = 19301 //
//
type C2SGirlHillStart struct {
}

func (m *C2SGirlHillStart) Reset()         { *m = C2SGirlHillStart{} }
func (m *C2SGirlHillStart) String() string { return proto.CompactTextString(m) }
func (*C2SGirlHillStart) ProtoMessage()    {}
func (m *C2SGirlHillStart) GetId() uint16  { return PCK_C2SGirlHillStart_ID }

const PCK_S2CGirlHillStart_ID = 19302 //
//
type S2CGirlHillStart struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CGirlHillStart) Reset()         { *m = S2CGirlHillStart{} }
func (m *S2CGirlHillStart) String() string { return proto.CompactTextString(m) }
func (*S2CGirlHillStart) ProtoMessage()    {}
func (m *S2CGirlHillStart) GetId() uint16  { return PCK_S2CGirlHillStart_ID }

const PCK_C2SGirlHillInspire_ID = 19303 //
//
type C2SGirlHillInspire struct {
}

func (m *C2SGirlHillInspire) Reset()         { *m = C2SGirlHillInspire{} }
func (m *C2SGirlHillInspire) String() string { return proto.CompactTextString(m) }
func (*C2SGirlHillInspire) ProtoMessage()    {}
func (m *C2SGirlHillInspire) GetId() uint16  { return PCK_C2SGirlHillInspire_ID }

const PCK_S2CGirlHillInspire_ID = 19304 //
//
type S2CGirlHillInspire struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CGirlHillInspire) Reset()         { *m = S2CGirlHillInspire{} }
func (m *S2CGirlHillInspire) String() string { return proto.CompactTextString(m) }
func (*S2CGirlHillInspire) ProtoMessage()    {}
func (m *S2CGirlHillInspire) GetId() uint16  { return PCK_S2CGirlHillInspire_ID }

const PCK_C2SGirlHillBuyTimes_ID = 19305 //
//
type C2SGirlHillBuyTimes struct {
}

func (m *C2SGirlHillBuyTimes) Reset()         { *m = C2SGirlHillBuyTimes{} }
func (m *C2SGirlHillBuyTimes) String() string { return proto.CompactTextString(m) }
func (*C2SGirlHillBuyTimes) ProtoMessage()    {}
func (m *C2SGirlHillBuyTimes) GetId() uint16  { return PCK_C2SGirlHillBuyTimes_ID }

const PCK_S2CGirlHillBuyTimes_ID = 19306 //
//
type S2CGirlHillBuyTimes struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CGirlHillBuyTimes) Reset()         { *m = S2CGirlHillBuyTimes{} }
func (m *S2CGirlHillBuyTimes) String() string { return proto.CompactTextString(m) }
func (*S2CGirlHillBuyTimes) ProtoMessage()    {}
func (m *S2CGirlHillBuyTimes) GetId() uint16  { return PCK_S2CGirlHillBuyTimes_ID }

const PCK_C2SZhenFaActive_ID = 3171 //
//
type C2SZhenFaActive struct {
	//
	ZhenFaId int32 `protobuf:"varint,1,opt,name=ZhenFaId,proto3" json:"ZhenFaId,omitempty"`
}

func (m *C2SZhenFaActive) Reset()         { *m = C2SZhenFaActive{} }
func (m *C2SZhenFaActive) String() string { return proto.CompactTextString(m) }
func (*C2SZhenFaActive) ProtoMessage()    {}
func (m *C2SZhenFaActive) GetId() uint16  { return PCK_C2SZhenFaActive_ID }

const PCK_S2CZhenFaActive_ID = 3172 //
//
type S2CZhenFaActive struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	ZhenFaId int32 `protobuf:"varint,2,opt,name=ZhenFaId,proto3" json:"ZhenFaId,omitempty"`
}

func (m *S2CZhenFaActive) Reset()         { *m = S2CZhenFaActive{} }
func (m *S2CZhenFaActive) String() string { return proto.CompactTextString(m) }
func (*S2CZhenFaActive) ProtoMessage()    {}
func (m *S2CZhenFaActive) GetId() uint16  { return PCK_S2CZhenFaActive_ID }

const PCK_C2SZhenFaUpgrade_ID = 3173 //
//
type C2SZhenFaUpgrade struct {
	//
	ZhenFaId int32 `protobuf:"varint,1,opt,name=ZhenFaId,proto3" json:"ZhenFaId,omitempty"`
}

func (m *C2SZhenFaUpgrade) Reset()         { *m = C2SZhenFaUpgrade{} }
func (m *C2SZhenFaUpgrade) String() string { return proto.CompactTextString(m) }
func (*C2SZhenFaUpgrade) ProtoMessage()    {}
func (m *C2SZhenFaUpgrade) GetId() uint16  { return PCK_C2SZhenFaUpgrade_ID }

const PCK_S2CZhenFaUpgrade_ID = 3174 //
//
type S2CZhenFaUpgrade struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	ZhenFaId int32 `protobuf:"varint,2,opt,name=ZhenFaId,proto3" json:"ZhenFaId,omitempty"`
}

func (m *S2CZhenFaUpgrade) Reset()         { *m = S2CZhenFaUpgrade{} }
func (m *S2CZhenFaUpgrade) String() string { return proto.CompactTextString(m) }
func (*S2CZhenFaUpgrade) ProtoMessage()    {}
func (m *S2CZhenFaUpgrade) GetId() uint16  { return PCK_S2CZhenFaUpgrade_ID }

const PCK_C2SZhenFaEquip_ID = 3175 //
//
type C2SZhenFaEquip struct {
	//
	ZhenFaId int32 `protobuf:"varint,1,opt,name=ZhenFaId,proto3" json:"ZhenFaId,omitempty"`
}

func (m *C2SZhenFaEquip) Reset()         { *m = C2SZhenFaEquip{} }
func (m *C2SZhenFaEquip) String() string { return proto.CompactTextString(m) }
func (*C2SZhenFaEquip) ProtoMessage()    {}
func (m *C2SZhenFaEquip) GetId() uint16  { return PCK_C2SZhenFaEquip_ID }

const PCK_S2CZhenFaEquip_ID = 3176 //
//
type S2CZhenFaEquip struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	ZhenFaId int32 `protobuf:"varint,2,opt,name=ZhenFaId,proto3" json:"ZhenFaId,omitempty"`
}

func (m *S2CZhenFaEquip) Reset()         { *m = S2CZhenFaEquip{} }
func (m *S2CZhenFaEquip) String() string { return proto.CompactTextString(m) }
func (*S2CZhenFaEquip) ProtoMessage()    {}
func (m *S2CZhenFaEquip) GetId() uint16  { return PCK_S2CZhenFaEquip_ID }

const PCK_C2SJiBanActive_ID = 3201 //
//
type C2SJiBanActive struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2SJiBanActive) Reset()         { *m = C2SJiBanActive{} }
func (m *C2SJiBanActive) String() string { return proto.CompactTextString(m) }
func (*C2SJiBanActive) ProtoMessage()    {}
func (m *C2SJiBanActive) GetId() uint16  { return PCK_C2SJiBanActive_ID }

const PCK_S2CJiBanActive_ID = 3202 //
//
type S2CJiBanActive struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *S2CJiBanActive) Reset()         { *m = S2CJiBanActive{} }
func (m *S2CJiBanActive) String() string { return proto.CompactTextString(m) }
func (*S2CJiBanActive) ProtoMessage()    {}
func (m *S2CJiBanActive) GetId() uint16  { return PCK_S2CJiBanActive_ID }

const PCK_C2SJiBanEquip_ID = 3203 //
//
type C2SJiBanEquip struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2SJiBanEquip) Reset()         { *m = C2SJiBanEquip{} }
func (m *C2SJiBanEquip) String() string { return proto.CompactTextString(m) }
func (*C2SJiBanEquip) ProtoMessage()    {}
func (m *C2SJiBanEquip) GetId() uint16  { return PCK_C2SJiBanEquip_ID }

const PCK_S2CJiBanEquip_ID = 3204 //
//
type S2CJiBanEquip struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *S2CJiBanEquip) Reset()         { *m = S2CJiBanEquip{} }
func (m *S2CJiBanEquip) String() string { return proto.CompactTextString(m) }
func (*S2CJiBanEquip) ProtoMessage()    {}
func (m *S2CJiBanEquip) GetId() uint16  { return PCK_S2CJiBanEquip_ID }

//
type FightUnitData struct {
	//
	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	UnitType int32 `protobuf:"varint,2,opt,name=UnitType,proto3" json:"UnitType,omitempty"`
	//
	Id int64 `protobuf:"varint,3,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Skill []*Skill `protobuf:"bytes,4,rep,name=Skill,proto3" json:"Skill,omitempty"`
	//
	ItemAttr []*IntAttr `protobuf:"bytes,5,rep,name=ItemAttr,proto3" json:"ItemAttr,omitempty"`
	//
	A []*IntAttr `protobuf:"bytes,6,rep,name=A,proto3" json:"A,omitempty"`
	//
	B []*StrAttr `protobuf:"bytes,7,rep,name=B,proto3" json:"B,omitempty"`
}

func (m *FightUnitData) Reset()         { *m = FightUnitData{} }
func (m *FightUnitData) String() string { return proto.CompactTextString(m) }
func (*FightUnitData) ProtoMessage()    {}

//
type PartnerSportsEnemy struct {
	//
	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	Nick string `protobuf:"bytes,2,opt,name=Nick,proto3" json:"Nick,omitempty"`
	//
	AreaId int64 `protobuf:"varint,3,opt,name=AreaId,proto3" json:"AreaId,omitempty"`
	//
	GradeType int32 `protobuf:"varint,4,opt,name=GradeType,proto3" json:"GradeType,omitempty"`
	//
	FightValue int64 `protobuf:"varint,5,opt,name=FightValue,proto3" json:"FightValue,omitempty"`
	//
	PartnerData *FightUnitData `protobuf:"bytes,6,opt,name=PartnerData,proto3" json:"PartnerData,omitempty"`
	//
	FightResult int32 `protobuf:"varint,7,opt,name=FightResult,proto3" json:"FightResult,omitempty"`
	//
	AddScore int32 `protobuf:"varint,8,opt,name=AddScore,proto3" json:"AddScore,omitempty"`
	//
	PartnerId int32 `protobuf:"varint,9,opt,name=PartnerId,proto3" json:"PartnerId,omitempty"`
	//
	LoseScore int32 `protobuf:"varint,10,opt,name=LoseScore,proto3" json:"LoseScore,omitempty"`
}

func (m *PartnerSportsEnemy) Reset()         { *m = PartnerSportsEnemy{} }
func (m *PartnerSportsEnemy) String() string { return proto.CompactTextString(m) }
func (*PartnerSportsEnemy) ProtoMessage()    {}

//
type PartnerSportsData struct {
	//
	GradeType int32 `protobuf:"varint,1,opt,name=GradeType,proto3" json:"GradeType,omitempty"`
	//
	Score int32 `protobuf:"varint,2,opt,name=Score,proto3" json:"Score,omitempty"`
	//
	DayWin int32 `protobuf:"varint,3,opt,name=DayWin,proto3" json:"DayWin,omitempty"`
	//
	WinPrizeState int32 `protobuf:"varint,4,opt,name=WinPrizeState,proto3" json:"WinPrizeState,omitempty"`
	//
	PartnerId int32 `protobuf:"varint,5,opt,name=PartnerId,proto3" json:"PartnerId,omitempty"`
	//
	FightTimes int32 `protobuf:"varint,6,opt,name=FightTimes,proto3" json:"FightTimes,omitempty"`
	//
	RefreshTimes int32 `protobuf:"varint,7,opt,name=RefreshTimes,proto3" json:"RefreshTimes,omitempty"`
	//
	BuyTimes int32 `protobuf:"varint,8,opt,name=BuyTimes,proto3" json:"BuyTimes,omitempty"`
	//
	Enemy []*PartnerSportsEnemy `protobuf:"bytes,9,rep,name=Enemy,proto3" json:"Enemy,omitempty"`
	//
	Stage int32 `protobuf:"varint,10,opt,name=Stage,proto3" json:"Stage,omitempty"`
	//
	EndTime int64 `protobuf:"varint,11,opt,name=EndTime,proto3" json:"EndTime,omitempty"`
}

func (m *PartnerSportsData) Reset()         { *m = PartnerSportsData{} }
func (m *PartnerSportsData) String() string { return proto.CompactTextString(m) }
func (*PartnerSportsData) ProtoMessage()    {}

const PCK_S2CPartnerSportsData_ID = 20002 //
//
type S2CPartnerSportsData struct {
	//
	Data []*PartnerSportsData `protobuf:"bytes,1,rep,name=Data,proto3" json:"Data,omitempty"`
}

func (m *S2CPartnerSportsData) Reset()         { *m = S2CPartnerSportsData{} }
func (m *S2CPartnerSportsData) String() string { return proto.CompactTextString(m) }
func (*S2CPartnerSportsData) ProtoMessage()    {}
func (m *S2CPartnerSportsData) GetId() uint16  { return PCK_S2CPartnerSportsData_ID }

const PCK_C2SPartnerSportsFight_ID = 20003 //
//
type C2SPartnerSportsFight struct {
	//
	GradeType int32 `protobuf:"varint,1,opt,name=GradeType,proto3" json:"GradeType,omitempty"`
	//
	Idx int32 `protobuf:"varint,2,opt,name=Idx,proto3" json:"Idx,omitempty"`
}

func (m *C2SPartnerSportsFight) Reset()         { *m = C2SPartnerSportsFight{} }
func (m *C2SPartnerSportsFight) String() string { return proto.CompactTextString(m) }
func (*C2SPartnerSportsFight) ProtoMessage()    {}
func (m *C2SPartnerSportsFight) GetId() uint16  { return PCK_C2SPartnerSportsFight_ID }

const PCK_S2CPartnerSportsFight_ID = 20004 //
//
type S2CPartnerSportsFight struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CPartnerSportsFight) Reset()         { *m = S2CPartnerSportsFight{} }
func (m *S2CPartnerSportsFight) String() string { return proto.CompactTextString(m) }
func (*S2CPartnerSportsFight) ProtoMessage()    {}
func (m *S2CPartnerSportsFight) GetId() uint16  { return PCK_S2CPartnerSportsFight_ID }

const PCK_C2SPartnerSportsRefresh_ID = 20005 //
//
type C2SPartnerSportsRefresh struct {
	//
	GradeType int32 `protobuf:"varint,1,opt,name=GradeType,proto3" json:"GradeType,omitempty"`
}

func (m *C2SPartnerSportsRefresh) Reset()         { *m = C2SPartnerSportsRefresh{} }
func (m *C2SPartnerSportsRefresh) String() string { return proto.CompactTextString(m) }
func (*C2SPartnerSportsRefresh) ProtoMessage()    {}
func (m *C2SPartnerSportsRefresh) GetId() uint16  { return PCK_C2SPartnerSportsRefresh_ID }

const PCK_S2CPartnerSportsRefresh_ID = 20006 //
//
type S2CPartnerSportsRefresh struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CPartnerSportsRefresh) Reset()         { *m = S2CPartnerSportsRefresh{} }
func (m *S2CPartnerSportsRefresh) String() string { return proto.CompactTextString(m) }
func (*S2CPartnerSportsRefresh) ProtoMessage()    {}
func (m *S2CPartnerSportsRefresh) GetId() uint16  { return PCK_S2CPartnerSportsRefresh_ID }

const PCK_C2SPartnerSportsBuyFightTimes_ID = 20007 //
//
type C2SPartnerSportsBuyFightTimes struct {
	//
	GradeType int32 `protobuf:"varint,1,opt,name=GradeType,proto3" json:"GradeType,omitempty"`
}

func (m *C2SPartnerSportsBuyFightTimes) Reset()         { *m = C2SPartnerSportsBuyFightTimes{} }
func (m *C2SPartnerSportsBuyFightTimes) String() string { return proto.CompactTextString(m) }
func (*C2SPartnerSportsBuyFightTimes) ProtoMessage()    {}
func (m *C2SPartnerSportsBuyFightTimes) GetId() uint16  { return PCK_C2SPartnerSportsBuyFightTimes_ID }

const PCK_S2CPartnerSportsBuyFightTimes_ID = 20008 //
//
type S2CPartnerSportsBuyFightTimes struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CPartnerSportsBuyFightTimes) Reset()         { *m = S2CPartnerSportsBuyFightTimes{} }
func (m *S2CPartnerSportsBuyFightTimes) String() string { return proto.CompactTextString(m) }
func (*S2CPartnerSportsBuyFightTimes) ProtoMessage()    {}
func (m *S2CPartnerSportsBuyFightTimes) GetId() uint16  { return PCK_S2CPartnerSportsBuyFightTimes_ID }

const PCK_C2SPartnerSportsSetFightPartner_ID = 20009 //
//
type C2SPartnerSportsSetFightPartner struct {
	//
	GradeType int32 `protobuf:"varint,1,opt,name=GradeType,proto3" json:"GradeType,omitempty"`
	//
	PartnerId int32 `protobuf:"varint,2,opt,name=PartnerId,proto3" json:"PartnerId,omitempty"`
}

func (m *C2SPartnerSportsSetFightPartner) Reset()         { *m = C2SPartnerSportsSetFightPartner{} }
func (m *C2SPartnerSportsSetFightPartner) String() string { return proto.CompactTextString(m) }
func (*C2SPartnerSportsSetFightPartner) ProtoMessage()    {}
func (m *C2SPartnerSportsSetFightPartner) GetId() uint16 {
	return PCK_C2SPartnerSportsSetFightPartner_ID
}

const PCK_S2CPartnerSportsSetFightPartner_ID = 20010 //
//
type S2CPartnerSportsSetFightPartner struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CPartnerSportsSetFightPartner) Reset()         { *m = S2CPartnerSportsSetFightPartner{} }
func (m *S2CPartnerSportsSetFightPartner) String() string { return proto.CompactTextString(m) }
func (*S2CPartnerSportsSetFightPartner) ProtoMessage()    {}
func (m *S2CPartnerSportsSetFightPartner) GetId() uint16 {
	return PCK_S2CPartnerSportsSetFightPartner_ID
}

const PCK_C2SPartnerSportsGetWinPrize_ID = 20011 //
//
type C2SPartnerSportsGetWinPrize struct {
	//
	GradeType int32 `protobuf:"varint,1,opt,name=GradeType,proto3" json:"GradeType,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2SPartnerSportsGetWinPrize) Reset()         { *m = C2SPartnerSportsGetWinPrize{} }
func (m *C2SPartnerSportsGetWinPrize) String() string { return proto.CompactTextString(m) }
func (*C2SPartnerSportsGetWinPrize) ProtoMessage()    {}
func (m *C2SPartnerSportsGetWinPrize) GetId() uint16  { return PCK_C2SPartnerSportsGetWinPrize_ID }

const PCK_S2CPartnerSportsGetWinPrize_ID = 20012 //
//
type S2CPartnerSportsGetWinPrize struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
}

func (m *S2CPartnerSportsGetWinPrize) Reset()         { *m = S2CPartnerSportsGetWinPrize{} }
func (m *S2CPartnerSportsGetWinPrize) String() string { return proto.CompactTextString(m) }
func (*S2CPartnerSportsGetWinPrize) ProtoMessage()    {}
func (m *S2CPartnerSportsGetWinPrize) GetId() uint16  { return PCK_S2CPartnerSportsGetWinPrize_ID }

//
type UserPartnerSportsScore struct {
	//
	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	Score int32 `protobuf:"varint,2,opt,name=Score,proto3" json:"Score,omitempty"`
	//
	FightValue int64 `protobuf:"varint,3,opt,name=FightValue,proto3" json:"FightValue,omitempty"`
}

func (m *UserPartnerSportsScore) Reset()         { *m = UserPartnerSportsScore{} }
func (m *UserPartnerSportsScore) String() string { return proto.CompactTextString(m) }
func (*UserPartnerSportsScore) ProtoMessage()    {}

//
type ServerPartnerSportsScore struct {
	//
	Users []*UserPartnerSportsScore `protobuf:"bytes,1,rep,name=Users,proto3" json:"Users,omitempty"`
	//
	Stage int32 `protobuf:"varint,2,opt,name=Stage,proto3" json:"Stage,omitempty"`
	//
	GradeType int32 `protobuf:"varint,3,opt,name=GradeType,proto3" json:"GradeType,omitempty"`
}

func (m *ServerPartnerSportsScore) Reset()         { *m = ServerPartnerSportsScore{} }
func (m *ServerPartnerSportsScore) String() string { return proto.CompactTextString(m) }
func (*ServerPartnerSportsScore) ProtoMessage()    {}

//
type ServerPartnerSportsData struct {
	//
	Scores []*ServerPartnerSportsScore `protobuf:"bytes,1,rep,name=Scores,proto3" json:"Scores,omitempty"`
}

func (m *ServerPartnerSportsData) Reset()         { *m = ServerPartnerSportsData{} }
func (m *ServerPartnerSportsData) String() string { return proto.CompactTextString(m) }
func (*ServerPartnerSportsData) ProtoMessage()    {}

//
type PartnerSportsRankData struct {
	//
	UserId int64 `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	//
	GradeType int32 `protobuf:"varint,2,opt,name=GradeType,proto3" json:"GradeType,omitempty"`
	//
	Score int32 `protobuf:"varint,3,opt,name=Score,proto3" json:"Score,omitempty"`
	//
	Name string `protobuf:"bytes,4,opt,name=Name,proto3" json:"Name,omitempty"`
	//
	VipLevel int32 `protobuf:"varint,5,opt,name=VipLevel,proto3" json:"VipLevel,omitempty"`
	//
	AreaId int64 `protobuf:"varint,6,opt,name=AreaId,proto3" json:"AreaId,omitempty"`
	//
	FightValue int64 `protobuf:"varint,7,opt,name=FightValue,proto3" json:"FightValue,omitempty"`
}

func (m *PartnerSportsRankData) Reset()         { *m = PartnerSportsRankData{} }
func (m *PartnerSportsRankData) String() string { return proto.CompactTextString(m) }
func (*PartnerSportsRankData) ProtoMessage()    {}

const PCK_S2KPartnerSportsScore_ID = 20001 //
//
type S2KPartnerSportsScore struct {
	//
	Scores []*PartnerSportsRankData `protobuf:"bytes,1,rep,name=Scores,proto3" json:"Scores,omitempty"`
}

func (m *S2KPartnerSportsScore) Reset()         { *m = S2KPartnerSportsScore{} }
func (m *S2KPartnerSportsScore) String() string { return proto.CompactTextString(m) }
func (*S2KPartnerSportsScore) ProtoMessage()    {}
func (m *S2KPartnerSportsScore) GetId() uint16  { return PCK_S2KPartnerSportsScore_ID }

const PCK_C2STianShuUnlock_ID = 3221 //
//
type C2STianShuUnlock struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2STianShuUnlock) Reset()         { *m = C2STianShuUnlock{} }
func (m *C2STianShuUnlock) String() string { return proto.CompactTextString(m) }
func (*C2STianShuUnlock) ProtoMessage()    {}
func (m *C2STianShuUnlock) GetId() uint16  { return PCK_C2STianShuUnlock_ID }

const PCK_S2CTianShuUnlock_ID = 3222 //
//
type S2CTianShuUnlock struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *S2CTianShuUnlock) Reset()         { *m = S2CTianShuUnlock{} }
func (m *S2CTianShuUnlock) String() string { return proto.CompactTextString(m) }
func (*S2CTianShuUnlock) ProtoMessage()    {}
func (m *S2CTianShuUnlock) GetId() uint16  { return PCK_S2CTianShuUnlock_ID }

const PCK_C2STianShuActive_ID = 3223 //
//
type C2STianShuActive struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *C2STianShuActive) Reset()         { *m = C2STianShuActive{} }
func (m *C2STianShuActive) String() string { return proto.CompactTextString(m) }
func (*C2STianShuActive) ProtoMessage()    {}
func (m *C2STianShuActive) GetId() uint16  { return PCK_C2STianShuActive_ID }

const PCK_S2CTianShuActive_ID = 3224 //
//
type S2CTianShuActive struct {
	//
	Tag int32 `protobuf:"varint,1,opt,name=Tag,proto3" json:"Tag,omitempty"`
	//
	Id int32 `protobuf:"varint,2,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (m *S2CTianShuActive) Reset()         { *m = S2CTianShuActive{} }
func (m *S2CTianShuActive) String() string { return proto.CompactTextString(m) }
func (*S2CTianShuActive) ProtoMessage()    {}
func (m *S2CTianShuActive) GetId() uint16  { return PCK_S2CTianShuActive_ID }

//
type C2SCommon struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Param []int64 `protobuf:"varint,2,rep,name=Param,proto3" json:"Param,omitempty"`
}

func (m *C2SCommon) Reset()         { *m = C2SCommon{} }
func (m *C2SCommon) String() string { return proto.CompactTextString(m) }
func (*C2SCommon) ProtoMessage()    {}

//
type S2CCommon struct {
	//
	Id int32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	//
	Param []int64 `protobuf:"varint,2,rep,name=Param,proto3" json:"Param,omitempty"`
}

func (m *S2CCommon) Reset()         { *m = S2CCommon{} }
func (m *S2CCommon) String() string { return proto.CompactTextString(m) }
func (*S2CCommon) ProtoMessage()    {}
