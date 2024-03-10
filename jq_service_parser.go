package opslevel_jq_parser

func RunWithConfig(cfg ServiceRegistrationConfig, json string) ServiceRegistration {
	return ServiceRegistration{
		Aliases:      run[string](NewJQArrayParser(cfg.Aliases), json, nil, defaultStringHandler),
		Description:  NewJQFieldParser(cfg.Description).ParseValue(json),
		Framework:    NewJQFieldParser(cfg.Framework).ParseValue(json),
		Language:     NewJQFieldParser(cfg.Language).ParseValue(json),
		Lifecycle:    NewJQFieldParser(cfg.Lifecycle).ParseValue(json),
		Name:         NewJQFieldParser(cfg.Name).ParseValue(json),
		Owner:        NewJQFieldParser(cfg.Owner).ParseValue(json),
		Product:      NewJQFieldParser(cfg.Product).ParseValue(json),
		Properties:   NewJQDictParser(cfg.Properties).RunPropertyAssignments(json),
		System:       NewJQFieldParser(cfg.System).ParseValue(json),
		Repositories: RunRepositories(NewJQArrayParser(cfg.Repositories), json),
		TagAssigns:   RunTags(NewJQArrayParser(cfg.Tags.Assign), json),
		TagCreates:   RunTags(NewJQArrayParser(cfg.Tags.Create), json),
		Tier:         NewJQFieldParser(cfg.Tier).ParseValue(json),
		Tools:        RunTools(NewJQArrayParser(cfg.Tools), json),
	}
}
