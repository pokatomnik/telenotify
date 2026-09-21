package prompts

import _ "embed"

//go:generate cp ../../.agents/skills/telenotify/SKILL.md ./SKILL.md

//go:embed SKILL.md
var Skill string
