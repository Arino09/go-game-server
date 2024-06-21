package player

func (p *Player) HandlerRegister() {
	p.handlers["add_friend"] = p.AddFriend
	p.handlers["remove_friend"] = p.DelFriend
	p.handlers["resolve_chat_msg"] = p.ResolveChatMsg
}
