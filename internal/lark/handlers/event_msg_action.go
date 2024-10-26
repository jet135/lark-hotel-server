package handlers

type MessageAction struct { /*消息*/
}

func (*MessageAction) Execute(a *ActionInfo) bool {
	replyMsg(*a.ctx, "🤖️：没有什么事的话我先溜了", *a.info.msgId)
	return true
}
