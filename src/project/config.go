package project

type ConfigProject struct {
	GitlabSecret  string
	TelegramToken string
	TelegramChat  string
}

type ConfigEventProject struct {
	Comment      bool
	Deployment   bool
	FeatureFlag  bool
	Group        bool
	Issue        bool
	Job          bool
	MergeRequest bool
	Pipeline     bool
	Push         bool
	Release      bool
	SubGroup     bool
	Tag          bool
	WikiPage     bool
}
