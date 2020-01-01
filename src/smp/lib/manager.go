package lib

import "errors"

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
	if index < 0 || index > m.Len() {
		return nil, errors.New("Index out of range.")
	}
	return &m.musics[index], nil
}

func (m *MusicManager) Find(name string) *Music {
	if m.Len() == 0 {
		return nil
	}

	for _, music := range m.musics {
		if music.Name == name {
			return &music
		}
	}
	return nil
}

func (m *MusicManager) Add(music *Music) {
	m.musics = append(m.musics, *music)
}

func (m *MusicManager) Remove(index int) *Music {
	if index < 0 || index > m.Len() {
		return nil
	}
	music := &m.musics[index]
	if m.Len() == 1 {
		m.musics = make([]Music, 0)
		return music
	}
	if index > 0 && index < m.Len()-1 {
		m.musics = append(m.musics[:index-1], m.musics[index+1:]...)
	} else if index == m.Len()-1 {
		m.musics = m.musics[:index-1]
	} else {
		m.musics = m.musics[index+1:]
	}
	return music
}

func (m *MusicManager) RemoveByName(name string) *Music {
	if m.Len() == 0 {
		return nil
	}

	for i, music := range m.musics {
		if music.Name == name {
			return m.Remove(i)
		}
	}

	return nil
}
