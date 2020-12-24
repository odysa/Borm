package log

import "testing"

func TestLog(t *testing.T) {
	s:="真不错"
	Error("Nb!")
	Errorf("f is %s",s)
	Info("ABC!")
	Infof("f is %s",s)
}