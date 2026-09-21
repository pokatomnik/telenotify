package skillwrite

import "github.com/pokatomnik/telenotify/internal/prompts"

type skillWriter interface {
	WriteSkill(local bool, contents string) error
}

type WriteSkill struct {
	r skillWriter
}

func New(skillWriter skillWriter) WriteSkill {
	return WriteSkill{
		r: skillWriter,
	}
}

func (w WriteSkill) WriteSkill(local bool) error {
	return w.r.WriteSkill(local, prompts.Skill)
}
