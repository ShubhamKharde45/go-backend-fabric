package loadbalancer

type LoadBalancer struct {
	ServerURLs *[]string
}

func NewLoadBalancer(serverURLs []string) *LoadBalancer {
	return &LoadBalancer{
		ServerURLs: &serverURLs,
	}
}

func (lb *LoadBalancer) Start() {

}
