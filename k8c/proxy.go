type ProxyValue string

func NewProxyValue(value string) *ProxyValue {
	val := ProxyValue(value)
	return &val
}

func (p *ProxyValue) Empty() bool {
	return p == nil || *p == ""
}

func (p *ProxyValue) String() string {
	if p.Empty() {
		return ""
	}

	return string(*p)
}