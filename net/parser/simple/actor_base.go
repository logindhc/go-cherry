package simple

import (
	cfacade "cherry/facade"
	clog "cherry/logger"
	cactor "cherry/net/actor"
	cproto "cherry/net/proto"
)

type ActorBase struct {
	cactor.Base
}

func (p *ActorBase) Response(session *cproto.Session, mid uint32, v interface{}) {
	Response(p, session, mid, v)
}

func Response(iActor cfacade.IActor, session *cproto.Session, mid uint32, v interface{}) {
	data, err := iActor.App().Serializer().Marshal(v)
	if err != nil {
		clog.Warnf("[Response] Marshal error. v = %+v", v)
		return
	}

	rsp := &cproto.PomeloResponse{
		Sid:  session.Sid,
		Mid:  mid,
		Data: data,
	}

	iActor.Call(session.AgentPath, ResponseFuncName, rsp)
}
