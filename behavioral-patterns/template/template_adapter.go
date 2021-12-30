package template

// Adapter of MessageRetriever to anonymous func
type TemplateAdapter struct {
	myFunc func() string
}

func (a *TemplateAdapter) Message() string {
	if a.myFunc != nil {
		return a.myFunc()
	}
	return ""
}

func MessageRetrieverAdapter(f func() string) MessageRetriever {
	return &TemplateAdapter{myFunc: f}
}
