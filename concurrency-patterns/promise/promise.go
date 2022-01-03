package promise

type CallBackFunc func(string)
type ErrorFunc func(error)
type ExecuteAsyncFunc func() (string, error)

type Promise struct {
	callBackFunc CallBackFunc
	errorFunc    ErrorFunc
}

func (s *Promise) OnSuccess(f CallBackFunc) *Promise {
	s.callBackFunc = f
	return s
}

func (s *Promise) OnError(f ErrorFunc) *Promise {
	s.errorFunc = f
	return s
}

func (s *Promise) Execute(f ExecuteAsyncFunc) {
	go func(s *Promise) {
		results, err := f()
		if err != nil {
			s.errorFunc(err)
		} else {
			s.callBackFunc(results)
		}
	}(s)
}
