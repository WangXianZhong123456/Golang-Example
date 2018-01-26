package library

import (
	"errors"
	"fmt"
)

type Music struct {
	Id     string
	Name   string
	Astist string
	Source string
	Type   string
}
type MusicManager struct {
	musics []Music
}

func NewMusicManager() *MusicManager {
	return &MusicManager{make([]Music, 0)}
}
func (m *MusicManager) Len() int {
	return len(m.musics)
}

func (m *MusicManager) Get(index int) (music *Music, err error) {
	if index < 0 || index >= len(m.musics) {
		return nil, errors.New("Index out of range")
	}
	return &m.musics[index], nil
}
func (m *MusicManager) Find(name string) *Music {
	if len(m.musics) == 0 {
		return nil
	}
	for _, m := range m.musics {
		if m.Name == name {
			return &m
		}
	}
	return nil
}
func (m *MusicManager) Add(music *Music) {
	m.musics = append(m.musics, *music)
}
func (m *MusicManager) List() {
	for _, k := range m.musics {
		fmt.Println(k.Id, k.Name, k.Astist, k.Type)

	}
}

func (m *MusicManager) Remove(id string) *Music {
	var index = -1
	for i := 0; i < len(m.musics); i++ {
		if m.musics[i].Id == id {
			fmt.Println("removeId=:", id)
			index = i
			fmt.Println("index=:", index)
		}
	}
	if index < 0 || index >= len(m.musics) {
		return nil
	}
	removeMusic := &m.musics[index]
	// 如果只有一个元素//可以简化吧应该，如果有简化，请issue我
	if index == 0 && len(m.musics) == 1 {
		m.musics = make([]Music, 0)
	} else if index == 0 && len(m.musics) != 1 {
		m.musics = m.musics[1:]
	} else if index < len(m.musics)-1 {
		m.musics = append(m.musics[:index-1], m.musics[index+1:]...)
	} else {
		m.musics = m.musics[:index-1]
	}
	return removeMusic
}
