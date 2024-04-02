package env

import (
	"errors"
	"fmt"
	"git.ingcreations.com/ingcreations-golang/gokits/utils"
	"os"
	"reflect"
	"strings"
)

const (
	TagEnvParamName         = "env"
	TagEnvParamDefaultValue = "env_default"
	TabEnvParamPrefix       = "env_prefix"
)

var (
	ErrorMustPointer = errors.New("must be struct pointer(*ptr)")
)

func Load(ptr interface{}, prefix string) error {
	refType := reflect.TypeOf(ptr)
	fmt.Println(refType.Elem().Name())
	ek := refType.Elem().Kind()
	if refType.Kind() != reflect.Ptr || ek != reflect.Struct {
		return ErrorMustPointer
	}

	element := reflect.ValueOf(ptr).Elem()
	for i := 0; i < element.NumField(); i++ {
		//获取结构体字段定义
		field := element.Type().Field(i)
		tag := field.Tag
		env := tag.Get(TagEnvParamName)
		envValue := tag.Get(TagEnvParamDefaultValue)

		f := element.Field(i)
		if field.Type.Kind() == reflect.Ptr && field.Type.Elem().Kind() == reflect.Struct {
			envPrefix := tag.Get(TabEnvParamPrefix)
			//子结构体
			child := reflect.New(field.Type.Elem())
			p := child.Interface()
			if err := Load(p, envPrefix); err != nil {
				return err
			}
			f.Set(child)
		}

		if len(env) == 0 {
			//没有定义env注解则直接使用属性名
			env = field.Name
		}

		key := fmt.Sprintf("%s%s", prefix, env)
		ev := os.Getenv(strings.ToUpper(key))
		if len(ev) > 0 {
			envValue = ev
		}
		if field.Type.Kind() == reflect.String {
			element.Field(i).SetString(envValue)
		}

		utils.SetValue(&f, envValue)
	}
	return nil
}
