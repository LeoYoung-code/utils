package main

import "fmt"

// OldPlayer 是旧播放器接口，只能播放 .mp3 文件
type OldPlayer interface {
	PlayMP3(filename string)
}

// OldMusicPlayer 是旧播放器的实现
type OldMusicPlayer struct{}

func (p *OldMusicPlayer) PlayMP3(filename string) {
	fmt.Printf("Playing MP3 file: %s\n", filename)
}

// NewPlayer 是新播放器接口，可以播放 .mp4 文件
type NewPlayer interface {
	PlayMP4(filename string)
}

// NewMusicPlayer 是新播放器的实现
type NewMusicPlayer struct{}

func (p *NewMusicPlayer) PlayMP4(filename string) {
	fmt.Printf("Playing MP4 file: %s\n", filename)
}

// PlayerAdapter 是适配器，实现了 OldPlayer 接口
type PlayerAdapter struct {
	newPlayer NewPlayer
}

func (adapter *PlayerAdapter) PlayMP3(filename string) {
	fmt.Println("Converting MP3 to MP4 format...")
	// 使用新播放器播放 .mp4 文件
	adapter.newPlayer.PlayMP4(filename)
}

func main() {
	// 使用旧播放器
	oldPlayer := &OldMusicPlayer{}
	oldPlayer.PlayMP3("song.mp3")

	// 使用适配器来播放 .mp3 文件
	newPlayer := &NewMusicPlayer{}
	adapter := &PlayerAdapter{newPlayer: newPlayer}
	adapter.PlayMP3("song.mp3")
}
