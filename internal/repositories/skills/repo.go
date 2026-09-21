package skills

import (
	"errors"
	"os"
	"path"

	apperrors "github.com/pokatomnik/telenotify/internal/entities/errors"
)

const (
	allSkillsPath             = ".agents/skills"
	telenotifySkillPath       = "telenotify"
	telenotifySkillFileName   = "SKILL.md"
	defaultDirPermissionMode  = 0o755
	defaultFilePermissionMode = 0o644
)

type (
	Repo           struct{}
	skillRepoPaths struct {
		allSkillsPath string
		skillDirPath  string
		skillFilePath string
	}
)

func New() Repo {
	return Repo{}
}

func getPath(local bool) (string, error) {
	if local {
		return os.Getwd()
	}
	return os.UserHomeDir()
}

func (Repo) paths(local bool) skillRepoPaths {
	basePath, err := getPath(local)
	if err != nil {
		return skillRepoPaths{}
	}

	// Глобальный путь до каталога скиллов:
	// ~/.agents/skills
	// или локальный:
	// ./.agents/skills
	allSkillsPath := path.Join(basePath, allSkillsPath)

	// Глобальный путь до каталога скилла telenotify:
	// ~/.agents/skills/telenotify
	// или локальный:
	// ./.agents/skills/telenotify
	skillDirPath := path.Join(allSkillsPath, telenotifySkillPath)

	// Глобальный путь до файла скилла:
	//  ~/.agents/skills/telenotify/SKILL.md
	// или локальный:
	// ./.agents/skills/telenotify/SKILL.md
	skillFilePath := path.Join(skillDirPath, telenotifySkillFileName)

	return skillRepoPaths{
		allSkillsPath: allSkillsPath,
		skillDirPath:  skillDirPath,
		skillFilePath: skillFilePath,
	}

}

func (r Repo) CheckSkillExists(local bool) (bool, error) {
	paths := r.paths(local)
	stat, err := os.Stat(paths.skillFilePath)

	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	isFile := !stat.IsDir()
	if !isFile {
		return false, apperrors.ErrorSkillPathIsNotAFile
	}

	return true, nil
}

func (r Repo) WriteSkill(local bool, contents string) error {
	paths := r.paths(local)
	err := os.MkdirAll(paths.skillDirPath, defaultDirPermissionMode)
	if err != nil {
		return err
	}

	return os.WriteFile(
		paths.skillFilePath,
		[]byte(contents),
		defaultFilePermissionMode,
	)
}

func (r Repo) RemoveSkill(local bool) error {
	paths := r.paths(local)
	return os.RemoveAll(paths.skillDirPath)
}
