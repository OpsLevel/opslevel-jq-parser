package opslevel_jq_parser

import "github.com/opslevel/opslevel-go/v2024"

func RunWithConfig(cfg ServiceRegistrationConfig, json string) ServiceRegistration {
	return ServiceRegistration{
		//Aliases: NewJQArrayParser(cfg.Aliases).ParseValues(json),
		Aliases: Run[string](NewJQArrayParser(cfg.Aliases), json, nil, defaultStringHandler),
		//Description: NewJQFieldParser(cfg.Description).ParseValue(json),
		//Framework:   NewJQFieldParser(cfg.Framework).ParseValue(json),
		//Language:    NewJQFieldParser(cfg.Language).ParseValue(json),
		//Lifecycle:   NewJQFieldParser(cfg.Lifecycle).ParseValue(json),
		//Name:        NewJQFieldParser(cfg.Name).ParseValue(json),
		//Owner:       NewJQFieldParser(cfg.Owner).ParseValue(json),
		//Product:     NewJQFieldParser(cfg.Product).ParseValue(json),
		//Properties:  NewJQDictParser(cfg.Properties).ParsePropertyAssignments(json),
		//System:      NewJQFieldParser(cfg.System).ParseValue(json),
		//Repositories: NewJQArrayParser(cfg.Repositories).ParseRepositories(json),
		//TagAssigns:   NewJQArrayParser(cfg.Tags.Assign).ParseTags(json),
		//TagCreates:   NewJQArrayParser(cfg.Tags.Create).ParseTags(json),
		//Tier:         NewJQFieldParser(cfg.Tier).ParseValue(json),
		Tools: Run[opslevel.ToolCreateInput](NewJQArrayParser(cfg.Tools), json, defaultObjectHandler[opslevel.ToolCreateInput], nil),
	}
}
