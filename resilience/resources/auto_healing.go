package resources

func (res *InfraResources) StartAutohealing() {
	
}

func (res *InfraResources) RecoverCache() {
	res.mu.Lock()
}

func (res *InfraResources) RecoverDatabase() {
	res.mu.Lock()
}
