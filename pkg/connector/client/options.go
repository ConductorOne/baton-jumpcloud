package client

type Options struct {
	Page    *int32
	Limit   *int32
	Targets []string
}

func (o *Options) WithSkip(skip int32) *Options {
	o.Page = &skip
	return o
}

func (o *Options) WithLimit(limit int32) *Options {
	o.Limit = &limit
	return o
}

func (o *Options) WithTargets(targets []string) *Options {
	o.Targets = targets
	return o
}

// getPage returns the page value or 0 if nil
func (o *Options) getPage() int32 {
	if o.Page == nil {
		return 0
	}
	return *o.Page
}

// getLimit returns the limit value or 100 if nil
func (o *Options) getLimit() int32 {
	if o.Limit == nil {
		return 100
	}
	return *o.Limit
}

// getTargets returns the targets value or empty slice if nil
func (o *Options) getTargets() []string {
	if o.Targets == nil {
		return []string{}
	}
	return o.Targets
}
