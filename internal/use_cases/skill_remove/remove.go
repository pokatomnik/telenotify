package skillremove

type skillRemover interface {
	RemoveSkill(local bool) error
}

type RemoveSkill struct {
	r skillRemover
}

func New(skillRemover skillRemover) RemoveSkill {
	return RemoveSkill{
		r: skillRemover,
	}
}

func (r RemoveSkill) RemoveSkill(local bool) error {
	return r.r.RemoveSkill(local)
}
