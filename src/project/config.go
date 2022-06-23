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

type ConfigTitlesProject struct {
	Comment      string
	Deployment   string
	FeatureFlag  string
	Group        string
	Issue        string
	Job          string
	MergeRequest string
	Pipeline     string
	Push         string
	Release      string
	SubGroup     string
	Tag          string
	WikiPage     string
}
