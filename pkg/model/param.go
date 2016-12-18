package model

// Parameter of an op
type Param struct {
  Dir       *DirParam `yaml:",omitempty"`
  File      *FileParam `yaml:",omitempty"`
  NetSocket *NetSocketParam `yaml:",omitempty"`
  String    *StringParam `yaml:",omitempty"`
}

// Directory parameter of an op
type DirParam struct {
  Description string `yaml:"description"`
  IsSecret    bool `yaml:"isSecret"`
  Name        string `yaml:"name"`
  Value       string `yaml:"-"`
}

// File parameter of an op
type FileParam struct {
  Description string `yaml:"description"`
  IsSecret    bool `yaml:"isSecret"`
  Name        string `yaml:"name"`
  Value       string `yaml:"-"`
}

// Network socket parameter of an op
type NetSocketParam struct {
  Constraints *NetSocketConstraints `yaml:"constraints"`
  Description string `yaml:"description"`
  IsSecret    bool `yaml:"isSecret"`
  Name        string `yaml:"name"`
  Value       *NetSocketParamValue `yaml:"-"`
}

type NetSocketParamValue struct {
  Host string `json:"host"`
  Port uint `json:"port"`
}

// String parameter of an op
type StringParam struct {
  Constraints *StringConstraints `yaml:"constraints"`
  Default     string `yaml:"default,omitempty"`
  Description string `yaml:"description"`
  IsSecret    bool `yaml:"isSecret"`
  Name        string `yaml:"name"`
  Value       string `yaml:"-"`
}
