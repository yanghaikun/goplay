package example2

import (
	"fmt"
	"github.com/spf13/viper"
)

func SayHello2() {
	fmt.Print("I am module test example 2")
	viper.SetDefault("ContentDir", "content")
	viper.SetDefault("LayoutDir", "layouts")
	viper.SetDefault("Taxonomies", map[string]string{"tag": "tags", "category": "categories"})
}