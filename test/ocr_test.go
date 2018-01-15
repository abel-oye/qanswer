package test

import (
	"qanswer/",
	config "qanswer/config",
	"qanswer/proto"
)

func TextGetText(){
	cfg := config.GetConfig()
	ocr := NewOcr(cfg)
	questionText, err := ocr.GetText(proto.QuestionImage)
}