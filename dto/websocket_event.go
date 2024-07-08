package dto

func init() {
	eventIntentMap = transposeIntentEventMap(intentEventMap)
}

// 事件类型
const (
	EventGuildCreate           EventType = "GUILD_CREATE"
	EventGuildUpdate           EventType = "GUILD_UPDATE"
	EventGuildDelete           EventType = "GUILD_DELETE"
	EventChannelCreate         EventType = "CHANNEL_CREATE"
	EventChannelUpdate         EventType = "CHANNEL_UPDATE"
	EventChannelDelete         EventType = "CHANNEL_DELETE"
	EventGuildMemberAdd        EventType = "GUILD_MEMBER_ADD"
	EventGuildMemberUpdate     EventType = "GUILD_MEMBER_UPDATE"
	EventGuildMemberRemove     EventType = "GUILD_MEMBER_REMOVE"
	EventMessageCreate         EventType = "MESSAGE_CREATE"
	EventMessageReactionAdd    EventType = "MESSAGE_REACTION_ADD"
	EventMessageReactionRemove EventType = "MESSAGE_REACTION_REMOVE"
	EventAtMessageCreate       EventType = "AT_MESSAGE_CREATE"
	EventPublicMessageDelete   EventType = "PUBLIC_MESSAGE_DELETE"
	EventDirectMessageCreate   EventType = "DIRECT_MESSAGE_CREATE"
	EventDirectMessageDelete   EventType = "DIRECT_MESSAGE_DELETE"
	EventAudioStart            EventType = "AUDIO_START"
	EventAudioFinish           EventType = "AUDIO_FINISH"
	EventAudioOnMic            EventType = "AUDIO_ON_MIC"
	EventAudioOffMic           EventType = "AUDIO_OFF_MIC"
	EventMessageAuditPass      EventType = "MESSAGE_AUDIT_PASS"
	EventMessageAuditReject    EventType = "MESSAGE_AUDIT_REJECT"
	EventMessageDelete         EventType = "MESSAGE_DELETE"
	EventForumThreadCreate     EventType = "FORUM_THREAD_CREATE"
	EventForumThreadUpdate     EventType = "FORUM_THREAD_UPDATE"
	EventForumThreadDelete     EventType = "FORUM_THREAD_DELETE"
	EventForumPostCreate       EventType = "FORUM_POST_CREATE"
	EventForumPostDelete       EventType = "FORUM_POST_DELETE"
	EventForumReplyCreate      EventType = "FORUM_REPLY_CREATE"
	EventForumReplyDelete      EventType = "FORUM_REPLY_DELETE"
	EventForumAuditResult      EventType = "FORUM_PUBLISH_AUDIT_RESULT"
	EventInteractionCreate     EventType = "INTERACTION_CREATE"
	EventC2CMessageCreate      EventType = "C2C_MESSAGE_CREATE"      // 用户单聊发消息给机器人时候
	EventFriendAdd             EventType = "FRIEND_ADD"              // 用户添加使用机器人
	EventFriendDel             EventType = "FRIEND_DEL"              // 用户删除机器人
	EventC2CMsgReject          EventType = "C2C_MSG_REJECT"          // 用户在机器人资料卡手动关闭"主动消息"推送
	EventC2CMsgReceive         EventType = "C2C_MSG_RECEIVE"         // 用户在机器人资料卡手动开启"主动消息"推送开关
	EventGroupAtMessageCreate  EventType = "GROUP_AT_MESSAGE_CREATE" // 用户在群里@机器人时收到的消息
	EventGroupAddRobot         EventType = "GROUP_ADD_ROBOT"         // 机器人被添加到群聊
	EventGroupDelRobot         EventType = "GROUP_DEL_ROBOT"         // 机器人被移出群聊
	EventGroupMsgReject        EventType = "GROUP_MSG_REJECT"        // 群管理员主动在机器人资料页操作关闭通知
	EventGroupMsgReceive       EventType = "GROUP_MSG_RECEIVE"       // 群管理员主动在机器人资料页操作开启通知
)

// intentEventMap 不同 intent 对应的事件定义
var intentEventMap = map[Intent][]EventType{
	IntentGuilds: {
		EventGuildCreate, EventGuildUpdate, EventGuildDelete,
		EventChannelCreate, EventChannelUpdate, EventChannelDelete,
	},
	IntentGroupAndC2CEvent: {
		EventC2CMessageCreate, EventFriendAdd,
		EventFriendDel, EventC2CMsgReject,
		EventC2CMsgReceive, EventGroupAtMessageCreate,
		EventGroupAddRobot, EventGroupDelRobot,
		EventGroupMsgReject, EventGroupMsgReceive,
	},
	IntentGuildMembers:          {EventGuildMemberAdd, EventGuildMemberUpdate, EventGuildMemberRemove},
	IntentGuildMessages:         {EventMessageCreate, EventMessageDelete},
	IntentGuildMessageReactions: {EventMessageReactionAdd, EventMessageReactionRemove},
	IntentGuildAtMessage:        {EventAtMessageCreate, EventPublicMessageDelete},
	IntentDirectMessages:        {EventDirectMessageCreate, EventDirectMessageDelete},
	IntentAudio:                 {EventAudioStart, EventAudioFinish, EventAudioOnMic, EventAudioOffMic},
	IntentAudit:                 {EventMessageAuditPass, EventMessageAuditReject},
	IntentForum: {
		EventForumThreadCreate, EventForumThreadUpdate, EventForumThreadDelete, EventForumPostCreate,
		EventForumPostDelete, EventForumReplyCreate, EventForumReplyDelete, EventForumAuditResult,
	},
	IntentInteraction: {EventInteractionCreate},
}

var eventIntentMap = transposeIntentEventMap(intentEventMap)

// transposeIntentEventMap 转置 intent 与 event 的关系，用于根据 event 找到 intent
func transposeIntentEventMap(input map[Intent][]EventType) map[EventType]Intent {
	result := make(map[EventType]Intent)
	for i, eventTypes := range input {
		for _, s := range eventTypes {
			result[s] = i
		}
	}
	return result
}

// EventToIntent 事件转换对应的Intent
func EventToIntent(events ...EventType) Intent {
	var i Intent
	for _, event := range events {
		i = i | eventIntentMap[event]
	}
	return i
}
