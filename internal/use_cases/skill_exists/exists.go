package skillexists

type existsChecker interface {
	CheckSkillExists(local bool) (bool, error)
}

type SkillExists struct {
	r existsChecker
}

func New(existsChecker existsChecker) SkillExists {
	return SkillExists{
		r: existsChecker,
	}
}

func (s SkillExists) CheckExists(local bool) (bool, error) {
	return s.r.CheckSkillExists(local)
}
