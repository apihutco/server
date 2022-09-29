package config

type Open struct {
	*Tencent `mapstructure:"tencent"`
	*Gaode   `mapstructure:"gaode"`
}

type Tencent struct {
	Key string
}

type Gaode struct {
	Key string
}
