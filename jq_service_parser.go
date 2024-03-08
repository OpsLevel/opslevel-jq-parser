package opslevel_jq_parser

func RunWithConfig(cfg ServiceRegistrationConfig, json string) ServiceRegistration {
	return ServiceRegistration{
		Aliases:      NewJQArrayParser(cfg.Aliases).Run(json),
		Description:  NewJQFieldParser(cfg.Description).Run(json),
		Framework:    NewJQFieldParser(cfg.Framework).Run(json),
		Language:     NewJQFieldParser(cfg.Language).Run(json),
		Lifecycle:    NewJQFieldParser(cfg.Lifecycle).Run(json),
		Name:         NewJQFieldParser(cfg.Name).Run(json),
		Owner:        NewJQFieldParser(cfg.Owner).Run(json),
		Product:      NewJQFieldParser(cfg.Product).Run(json),
		Properties:   NewJQDictParser(cfg.Properties).Run(json),
		System:       NewJQFieldParser(cfg.System).Run(json),
		Repositories: NewJQArrayParser(cfg.Repositories).Run(json),
		TagAssigns:   NewJQArrayParser(cfg.Tags["assign"]).Run(json),
		TagCreates:   NewJQArrayParser(cfg.Tags["create"]).Run(json),
		Tier:         NewJQFieldParser(cfg.Tier).Run(json),
		Tools:        NewJQArrayParser(cfg.Tools).Run(json),
	}
}
