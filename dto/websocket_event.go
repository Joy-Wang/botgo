package dto

func init() {
	eventIntentMap = transposeIntentEventMap(intentEventMap)
}

// 事件类型
const (
	EventGuildCreate       EventType = "GUILD_CREATE"
	EventGuildUpdate       EventType = "GUILD_UPDATE"
	EventGuildDelete       EventType = "GUILD_DELETE"
	EventChannelCreate     EventType = "CHANNEL_CREATE"
	EventChannelUpdate     EventType = "CHANNEL_UPDATE"
	EventChannelDelete     EventType = "CHANNEL_DELETE"
	EventGuildMemberAdd    EventType = "GUILD_MEMBER_ADD"
	EventGuildMemberUpdate EventType = "GUILD_MEMBER_UPDATE"
	EventGuildMemberRemove EventType = "GUILD_MEMBER_REMOVE"
	EventMessageCreate     EventType = "MESSAGE_CREATE"
	EventAtMessageCreate   EventType = "AT_MESSAGE_CREATE"
	EventAudioStart        EventType = "AUDIO_START"
	EventAudioFinish       EventType = "AUDIO_FINISH"
	EventAudioOnMic        EventType = "AUDIO_ON_MIC"
	EventAudioOffMic       EventType = "AUDIO_OFF_MIC"
)

// intentEventMap 不同 intent 对应的事件定义
var intentEventMap = map[Intent][]EventType{
	IntentGuilds: {EventGuildCreate, EventGuildUpdate, EventGuildDelete,
		EventChannelCreate, EventChannelUpdate, EventChannelDelete},
	IntentGuildMembers:   {EventGuildMemberAdd, EventGuildMemberUpdate, EventGuildMemberRemove},
	IntentGuildMessages:  {EventMessageCreate},
	IntentGuildAtMessage: {EventAtMessageCreate},
	IntentAudio:          {EventAudioStart, EventAudioFinish, EventAudioOnMic, EventAudioOffMic},
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
