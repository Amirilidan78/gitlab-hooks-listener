package project

import (
	"gitlab-hooks-listener/config"
)

type Project interface {
	initial()
	GetConfig() ConfigProject
	GetEvents() ConfigEventProject
	GetTitles() ConfigTitlesProject
}

type project struct {
	Name   string
	Config ConfigProject
	Events ConfigEventProject
	Titles ConfigTitlesProject
}

func (p *project) initial() {

	c := config.NewConfig()

	if len(c.GetMap("projects."+p.Name)) == 0 {
		panic("project not found")
	}

	p.Config = ConfigProject{
		GitlabSecret:  c.GetString("projects." + p.Name + ".gitlab-secret"),
		TelegramToken: c.GetString("projects." + p.Name + ".telegram-token"),
		TelegramChat:  c.GetString("projects." + p.Name + ".telegram-chat"),
	}

	p.Events = ConfigEventProject{
		Comment:      c.GetBool("projects." + p.Name + ".events.comment"),
		Deployment:   c.GetBool("projects." + p.Name + ".events.deployment"),
		FeatureFlag:  c.GetBool("projects." + p.Name + ".events.feature-flag"),
		Group:        c.GetBool("projects." + p.Name + ".events.group"),
		Issue:        c.GetBool("projects." + p.Name + ".events.issue"),
		Job:          c.GetBool("projects." + p.Name + ".events.job"),
		MergeRequest: c.GetBool("projects." + p.Name + ".events.merge-request"),
		Pipeline:     c.GetBool("projects." + p.Name + ".events.pipeline"),
		Push:         c.GetBool("projects." + p.Name + ".events.push"),
		Release:      c.GetBool("projects." + p.Name + ".events.release"),
		SubGroup:     c.GetBool("projects." + p.Name + ".events.sub-group"),
		Tag:          c.GetBool("projects." + p.Name + ".events.tag"),
		WikiPage:     c.GetBool("projects." + p.Name + ".events.wiki-page"),
	}

	p.Titles = ConfigTitlesProject{
		Comment:      c.GetString("titles.comment"),
		Deployment:   c.GetString("titles.deployment"),
		FeatureFlag:  c.GetString("titles.feature-flag"),
		Group:        c.GetString("titles.group"),
		Issue:        c.GetString("titles.issue"),
		Job:          c.GetString("titles.job"),
		MergeRequest: c.GetString("titles.merge-request"),
		Pipeline:     c.GetString("titles.pipeline"),
		Push:         c.GetString("titles.push"),
		Release:      c.GetString("titles.release"),
		SubGroup:     c.GetString("titles.sub-group"),
		Tag:          c.GetString("titles.tag"),
		WikiPage:     c.GetString("titles.wiki-page"),
	}
}

func (p *project) GetConfig() ConfigProject {

	return p.Config
}

func (p *project) GetEvents() ConfigEventProject {

	return p.Events
}

func (p *project) GetTitles() ConfigTitlesProject {

	return p.Titles
}

func GetProject(name string) Project {
	p := project{Name: name}
	p.initial()
	return &p
}
