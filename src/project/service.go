package project

import (
	"gitlab-hooks-listener/config"
)

type Project interface {
	initial()
	GetConfig() ConfigProject
}

type project struct {
	Name   string
	Config ConfigProject
	Events ConfigEventProject
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
		Comment:      c.GetBool("projects." + p.Name + "events.comment"),
		Deployment:   c.GetBool("projects." + p.Name + "events.deployment"),
		FeatureFlag:  c.GetBool("projects." + p.Name + "events.feature-flag"),
		Group:        c.GetBool("projects." + p.Name + "events.group"),
		Issue:        c.GetBool("projects." + p.Name + "events.issue"),
		Job:          c.GetBool("projects." + p.Name + "events.job"),
		MergeRequest: c.GetBool("projects." + p.Name + "events.merge-request"),
		Pipeline:     c.GetBool("projects." + p.Name + "events.pipeline"),
		Push:         c.GetBool("projects." + p.Name + "events.push"),
		Release:      c.GetBool("projects." + p.Name + "events.release"),
		SubGroup:     c.GetBool("projects." + p.Name + "events.sub-group"),
		Tag:          c.GetBool("projects." + p.Name + "events.tag"),
		WikiPage:     c.GetBool("projects." + p.Name + "events.wiki-page"),
	}
}

func (p *project) GetConfig() ConfigProject {

	return p.Config
}

func GetProject(name string) Project {
	p := project{Name: name}
	p.initial()
	return &p
}
