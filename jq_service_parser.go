package opslevel_jq_parser

type JQServiceParser struct {
	aliases      *JQArrayParser
	description  *JQFieldParser
	framework    *JQFieldParser
	language     *JQFieldParser
	lifecycle    *JQFieldParser
	name         *JQFieldParser
	owner        *JQFieldParser
	product      *JQFieldParser
	properties   *JQPropertiesParser
	repositories *JQRepositoryParser
	system       *JQFieldParser
	tags         *JQTagsParser
	tier         *JQFieldParser
	tools        *JQToolsParser
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
		properties:   NewJQPropertiesParser(cfg.Properties),
		repositories: NewJQRepositoryParser(cfg.Repositories),
		system:       NewJQFieldParser(cfg.System),
		tags:         NewJQTagsParser(cfg.Tags),
		tier:         NewJQFieldParser(cfg.Tier),
		tools:        NewJQToolsParser(cfg.Tools),
	}
}

func (p *JQServiceParser) Run(json string) (*ServiceRegistration, error) {
	aliases, err := p.aliases.Run(json)
	if err != nil {
		return nil, err
	}
	description, err := p.description.Run(json)
	if err != nil {
		return nil, err
	}
	framework, err := p.framework.Run(json)
	if err != nil {
		return nil, err
	}
	language, err := p.language.Run(json)
	if err != nil {
		return nil, err
	}
	lifecycle, err := p.lifecycle.Run(json)
	if err != nil {
		return nil, err
	}
	name, err := p.name.Run(json)
	if err != nil {
		return nil, err
	}
	owner, err := p.owner.Run(json)
	if err != nil {
		return nil, err
	}
	product, err := p.product.Run(json)
	if err != nil {
		return nil, err
	}
	properties, err := p.properties.Run(json)
	if err != nil {
		return nil, err
	}
	repositories, err := p.repositories.Run(json)
	if err != nil {
		return nil, err
	}
	system, err := p.system.Run(json)
	if err != nil {
		return nil, err
	}
	tagCreates, tagAssigns, err := p.tags.Run(json)
	if err != nil {
		return nil, err
	}
	tier, err := p.tier.Run(json)
	if err != nil {
		return nil, err
	}
	tools, err := p.tools.Run(json)
	if err != nil {
		return nil, err
	}
	return &ServiceRegistration{
		Aliases:      aliases,
		Description:  description,
		Framework:    framework,
		Language:     language,
		Lifecycle:    lifecycle,
		Name:         name,
		Owner:        owner,
		Product:      product,
		Properties:   properties,
		Repositories: repositories,
		System:       system,
		TagAssigns:   tagAssigns,
		TagCreates:   tagCreates,
		Tier:         tier,
		Tools:        tools,
	}, nil
}

func Deduplicated[T any](objects []T, keyFunc func(object T) string) []T {
	out := make([]T, 0)
	set := make(map[string]struct{})
	for _, obj := range objects {
		key := keyFunc(obj)
		if _, ok := set[key]; ok {
			continue
		}
		set[key] = struct{}{}
		out = append(out, obj)
	}
	return out
}
