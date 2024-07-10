package config

type ServerGroup struct {
	Name    string   `yaml:"name"`
	Hosts   []string `yaml:"hosts"`
	SSHUser string   `yaml:"ssh_user"`
	SSHKey  string   `yaml:"ssh_key"`
}

type Task struct {
	Name    string            `yaml:"name"`
	Command string            `yaml:"command"`
	Vars    map[string]string `yaml:"vars,omitempty"`
}

type Configuration struct {
	Name    string   `yaml:"name"`
	Servers []string `yaml:"servers"`
	Tasks   []string `yaml:"tasks"`
}

type Config struct {
	ServerGroups   []ServerGroup   `yaml:"server_groups"`
	Tasks          []Task          `yaml:"tasks"`
	Configurations []Configuration `yaml:"configurations"`
}
