package example

import (
	"fmt"
	"github.com/spf13/viper"
)

func SayHello() {
	fmt.Print("I am module test")
	viper.SetDefault("ContentDir", "content")
	viper.SetDefault("LayoutDir", "layouts")
	viper.SetDefault("Taxonomies", map[string]string{"tag": "tags", "category": "categories"})
}