package opslevel_jq_parser

type JQServiceParser struct {
	aliases      JQArrayParser
	description  JQFieldParser
	framework    JQFieldParser
	language     JQFieldParser
	lifecycle    JQFieldParser
	name         JQFieldParser
	owner        JQFieldParser
	product      JQFieldParser
	properties   JQDictParser
	repositories JQRepositoryParser
	system       JQFieldParser
	tags         JQTagsParser
	tier         JQFieldParser
	tools        JQToolsParser
}

func NewJQServiceParser(cfg ServiceRegistrationConfig) *JQServiceParser {
	return &JQServiceParser{
		aliases:      NewJQArrayParser(cfg.Aliases),
		description:  NewJQFieldParser(cfg.Description),
		framework:    NewJQFieldParser(cfg.Framework),
		language:     NewJQFieldParser(cfg.Language),
		lifecycle:    NewJQFieldParser(cfg.Lifecycle),
		name:         NewJQFieldParser(cfg.Name),
		owner:        NewJQFieldParser(cfg.Owner),
		product:      NewJQFieldParser(cfg.Product),
		properties:   NewJQDictParser(cfg.Properties),
		repositories: NewJQRepositoryParser(cfg.Repositories),
		system:       NewJQFieldParser(cfg.System),
		tags:         NewJQTagsParser(cfg.Tags),
		tier:         NewJQFieldParser(cfg.Tier),
		tools:        NewJQToolsParser(cfg.Tools),
	}
}

func (p *JQServiceParser) Run(json string) ServiceRegistration {
	tagCreates, tagAssigns := p.tags.Run(json)
	return ServiceRegistration{
		Aliases:      p.aliases.Run(json),
		Description:  p.description.Run(json),
		Framework:    p.framework.Run(json),
		Language:     p.language.Run(json),
		Lifecycle:    p.lifecycle.Run(json),
		Name:         p.name.Run(json),
		Owner:        p.owner.Run(json),
		Product:      p.product.Run(json),
		Properties:   p.properties.Run(json),
		Repositories: p.repositories.Run(json),
		System:       p.system.Run(json),
		TagAssigns:   tagAssigns,
		TagCreates:   tagCreates,
		Tier:         p.tier.Run(json),
		Tools:        p.tools.Run(json),
	}
}
