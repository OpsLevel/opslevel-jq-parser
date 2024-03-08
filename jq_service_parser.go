package opslevel_jq_parser

type JQServiceParser struct {
	aliases     JQArrayParser
	description JQFieldParser
	framework   JQFieldParser
	language    JQFieldParser
	lifecycle   JQFieldParser
	name        JQFieldParser
	owner       JQFieldParser
	product     JQFieldParser
	properties  JQDictParser
	system      JQFieldParser
	tier        JQFieldParser
}

func NewJQServiceParser(cfg ServiceRegistrationConfig) JQServiceParser {
	return JQServiceParser{
		aliases:     NewJQArrayParser(cfg.Aliases),
		description: NewJQFieldParser(cfg.Description),
		framework:   NewJQFieldParser(cfg.Framework),
		language:    NewJQFieldParser(cfg.Language),
		lifecycle:   NewJQFieldParser(cfg.Lifecycle),
		name:        NewJQFieldParser(cfg.Name),
		owner:       NewJQFieldParser(cfg.Owner),
		product:     NewJQFieldParser(cfg.Product),
		properties:  NewJQDictParser(cfg.Properties),
		system:      NewJQFieldParser(cfg.System),
		tier:        NewJQFieldParser(cfg.Tier),
	}
}

func (p JQServiceParser) Run(json string) (ServiceRegistration, error) {
	aliases, err := p.aliases.Run(json)
	if err != nil {
		return ServiceRegistration{}, err
	}
	description, err := p.description.Run(json)
	if err != nil {
		return ServiceRegistration{}, err
	}
	framework, err := p.framework.Run(json)
	if err != nil {
		return ServiceRegistration{}, err
	}
	language, err := p.language.Run(json)
	if err != nil {
		return ServiceRegistration{}, err
	}
	lifecycle, err := p.lifecycle.Run(json)
	if err != nil {
		return ServiceRegistration{}, err
	}
	name, err := p.name.Run(json)
	if err != nil {
		return ServiceRegistration{}, err
	}
	owner, err := p.owner.Run(json)
	if err != nil {
		return ServiceRegistration{}, err
	}
	product, err := p.product.Run(json)
	if err != nil {
		return ServiceRegistration{}, err
	}
	properties, err := p.properties.Run(json)
	if err != nil {
		return ServiceRegistration{}, err
	}
	system, err := p.system.Run(json)
	if err != nil {
		return ServiceRegistration{}, err
	}
	tier, err := p.tier.Run(json)
	if err != nil {
		return ServiceRegistration{}, err
	}
	return ServiceRegistration{
		Aliases:     aliases,
		Description: description,
		Framework:   framework,
		Language:    language,
		Lifecycle:   lifecycle,
		Name:        name,
		Owner:       owner,
		Product:     product,
		Properties:  properties,
		System:      system,
		Tier:        tier,
	}, nil
}
