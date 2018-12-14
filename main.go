package main

import (
	"io/ioutil"

	"gitlab.com/avtnl/golang-tts/golang_tts"
)

func main() {
	text := "Cessna 9130 Delta, Oakland Center, radio check"
	polly := NewPolly()
	// polly.TTS(text, "en-gb-brian-radiocheck.mp3", golang_tts.Brian)
	// polly.TTS(text, "en-gb-amy-radiocheck.mp3", golang_tts.Amy)
	// polly.TTS(text, "en-gb-emma-radiocheck.mp3", golang_tts.Emma)
	polly.TTS(text, "en-au-russell-radiocheck.mp3", golang_tts.Russell)
	polly.TTS(text, "en-au-nicole-radiocheck.mp3", golang_tts.Nicole)
	polly.TTS(text, "en-in-aditi-radiocheck.mp3", golang_tts.Aditi)
	polly.TTS(text, "en-in-raveena-radiocheck.mp3", golang_tts.Raveena)

	polly.TTS(text, "en-us-joey-radiocheck.mp3", golang_tts.Joey)
	polly.TTS(text, "en-us-justin-radiocheck.mp3", golang_tts.Justin)
	polly.TTS(text, "en-us-matthew-radiocheck.mp3", golang_tts.Matthew)
	polly.TTS(text, "en-us-ivy-radiocheck.mp3", golang_tts.Ivy)
	polly.TTS(text, "en-us-joanna-radiocheck.mp3", golang_tts.Joanna)
	polly.TTS(text, "en-us-kendra-radiocheck.mp3", golang_tts.Kendra)
	polly.TTS(text, "en-us-kimberly-radiocheck.mp3", golang_tts.Kimberly)
	polly.TTS(text, "en-us-salli-radiocheck.mp3", golang_tts.Salli)
	polly.TTS(text, "en-gb-wls-geraint-radiocheck.mp3", golang_tts.Geraint)
	polly.TTS(text, "es-us-miguel-radiocheck.mp3", golang_tts.Miguel)
	polly.TTS(text, "es-us-penelope-radiocheck.mp3", golang_tts.Penelope)
	polly.TTS(text, "fr-fr-mathieu-radiocheck.mp3", golang_tts.Mathieu)
}

type PollyTTS struct {
	Polly *golang_tts.TTS
}

func NewPolly() *PollyTTS {
	polly := golang_tts.New("AKIAIxxxxxxxxxxxxxxx", "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	polly.Format(golang_tts.MP3)
	polly.Voice(golang_tts.Brian)
	return &PollyTTS{Polly: polly}
}

func (polly *PollyTTS) TTS(text, outputFile, voice string) {
	polly.Polly.Voice(voice)
	bytes, err := polly.Polly.Speech(text)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile(outputFile, bytes, 0644)
}
