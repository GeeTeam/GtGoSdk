package GtGoSdk

import (
	"testing"
)

func TestGeetestLib(t *testing.T) {
	gt:=GeetestLib("","")
	if gt.decodeRandBase("a1d0c6e83f027327d8461063f4ac58a6aa") != 370{
		t.Error("decodeRandBase error")
	}
	if gt.decodeResponse("a1d0c6e83f027327d8461063f4ac58a6aa","fcccccffc3050e")!= 122{
		t.Error("decodeResponse error")
	}
	if gt.validateFailImage(122,4,127) != true{
		t.Error("validateFailImage error")
	}
}