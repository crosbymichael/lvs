package lvs

import (
	"fmt"
	"os/exec"
)

type Proto string

const (
	TCP Proto = "-t"
	UDP Proto = "-u"
)

type Service struct {
	Proto Proto
	IP    string
	Port  int
}

// CreateService creates a new lvs server on the specific ip and port
func CreateService(ip string, proto Proto, port int) (*Service, error) {
	s := &Service{
		Proto: proto,
		IP:    ip,
		Port:  port,
	}

	if err := run("-A", string(proto), s.combine()); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Service) Delete() error {
	return run("-D", string(s.Proto), s.combine())
}

func (s *Service) AddServer(ip string, port int) error {
	return run("-a", string(s.Proto), s.combine(), "-r", fmt.Sprintf("%s:%d", ip, port), "-m")
}

func (s *Service) RemoveServer(ip string, port int) error {
	return run("-d", string(s.Proto), s.combine(), "-r", fmt.Sprintf("%s:%d", ip, port), "-m")
}

func (s *Service) combine() string {
	return fmt.Sprintf("%s:%d", s.IP, s.Port)
}

func run(args ...string) error {
	return exec.Command("ipvsadm", args...).Run()
}
