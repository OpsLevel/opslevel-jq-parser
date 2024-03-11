package opslevel_jq_parser

func RunWithConfig(cfg ServiceRegistrationConfig, json string) ServiceRegistration {
	return ServiceRegistration{
		Aliases:      run[string](NewJQArrayParser(cfg.Aliases), json, nil, defaultStringHandler),
		Description:  NewJQFieldParser(cfg.Description).Run(json),
		Framework:    NewJQFieldParser(cfg.Framework).Run(json),
		Language:     NewJQFieldParser(cfg.Language).Run(json),
		Lifecycle:    NewJQFieldParser(cfg.Lifecycle).Run(json),
		Name:         NewJQFieldParser(cfg.Name).Run(json),
		Owner:        NewJQFieldParser(cfg.Owner).Run(json),
		Product:      NewJQFieldParser(cfg.Product).Run(json),
		Properties:   NewJQDictParser(cfg.Properties).Run(json),
		System:       NewJQFieldParser(cfg.System).Run(json),
		Repositories: RunRepositories(NewJQArrayParser(cfg.Repositories), json),
		TagAssigns:   RunTags(NewJQArrayParser(cfg.Tags.Assign), json),
		TagCreates:   RunTags(NewJQArrayParser(cfg.Tags.Create), json),
		Tier:         NewJQFieldParser(cfg.Tier).Run(json),
		Tools:        RunTools(NewJQArrayParser(cfg.Tools), json),
	}
}
