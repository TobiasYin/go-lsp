package jsonrpc

import (
	"context"
	"sync"
)

type MethodInfo struct {
	Name       string
	NewRequest func() interface{}
	Handler    func(ctx context.Context, req interface{}) (interface{}, error)
}

type Server struct {
	session     map[int]*Session
	nowId       int
	methods     map[string]MethodInfo
	sessionLock sync.Mutex
}

func NewServer() *Server {
	s := &Server{}
	s.session = make(map[int]*Session)
	s.methods = make(map[string]MethodInfo)

	// Register Builtin
	s.RegisterMethod(CancelRequest())

	return s
}

func (s *Server) RegisterMethod(m MethodInfo) {
	s.methods[m.Name] = m
}

func (s *Server) ConnComeIn(conn ReaderWriter) {
	session := s.newSession(conn)
	session.Start()
}

func (s *Server) removeSession(id int) {
	s.sessionLock.Lock()
	defer s.sessionLock.Unlock()
	delete(s.session, id)
}

func (s *Server) newSession(conn ReaderWriter) *Session {
	s.sessionLock.Lock()
	defer s.sessionLock.Unlock()
	id := s.nowId
	s.nowId += 1
	session := newSession(id, s, conn)
	s.session[id] = session
	return session
}

func (s *Server) GetAllSession() []*Session {
	s.sessionLock.Lock()
	defer s.sessionLock.Unlock()
	sessions := make([]*Session, 0, len(s.session))
	for _, session := range s.session {
		sessions = append(sessions, session)
	}
	return sessions
}
