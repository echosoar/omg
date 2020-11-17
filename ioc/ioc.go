// Package ioc
package ioc

import (
	"reflect"
	"strings"
)

type IoC struct {
	Name string
	container map[string]interface{}
}

var globalIocContainer = make(map[string]*IoC);

// Provide
func (i *IoC) Provide(anyThing interface{}) error {
	provideName, provideType := GetNameAndType(anyThing);

	if provideType != "struct" {
		return &IoCError{"Not support" + provideType};
	}

	i.ProvideByName(provideName, anyThing);
	return nil;
}

func (i *IoC) ProvideByName(provideName string, anyThing interface{}) error {
	i.container[provideName] = anyThing;
	return nil;
}


// Inject
func (i *IoC) Inject(anyThing interface{}) (interface {}, error) {
	strt := reflect.TypeOf(anyThing);
	// inject by id
	if strt.String() == "string" {
		data := i.container[anyThing.(string)];
		if data == nil {
			return nil, &IoCError{"Not Exists"};
		}
		return data, nil;
	}
	for index := 0; index < strt.NumField(); index++ {
		value := strt.Field(index);
		injectTag := value.Tag.Get("inject");
		if injectTag != "" {
			iocValue, getIoCErr := i.GetValue(injectTag);
			if getIoCErr != nil {
				return nil, getIoCErr;
			}
			_, iocValueType := GetNameAndTypeByRfType(iocValue.Type());
			// https://stackoverflow.com/questions/63421976/panic-reflect-call-of-reflect-value-fieldbyname-on-interface-value
			v := reflect.ValueOf(&anyThing).Elem()
			tmp := reflect.New(v.Elem().Type()).Elem()
			tmp.Set(v.Elem())
			tmpValue := tmp.FieldByName(value.Name);
			_, tmpValueType := GetNameAndTypeByRfType(tmpValue.Type());
			
			if iocValueType != tmpValueType {
				return nil, &IoCError{"Inject type need equal"}
			}
			
			switch tmpValueType {
			case "string":
				tmpValue.SetString(iocValue.String());
			case "bool":
				tmpValue.SetBool(iocValue.Bool());
			case "int":
				tmpValue.SetInt(iocValue.Int());
			case "float":
				tmpValue.SetFloat(iocValue.Float());
			}
			v.Set(tmp)
		}
	}
	return anyThing, nil;
}

// ClearIoC
func (i *IoC) Clear() {
	i.container = make(map[string]interface{});
}


// GetIoCValue
func (i *IoC) GetValue(iocId string) (reflect.Value, error) {
	idList := strings.Split(iocId, ".");
	idLen := len(idList);
	var ioc *IoC;
	if idLen == 3 {
		ioc = Get(idList[0])
		idList = idList[1:];
	} else {
		ioc = i;
	}

	iocValue := ioc.container[idList[0]]
	injectedValue, injectErr := ioc.Inject(iocValue);
	if injectErr != nil {
		return reflect.Value{}, injectErr;
	}
	elem := reflect.ValueOf(&injectedValue).Elem();
	if len(idList) == 1 {
		return elem, nil;
	}
	tmp := reflect.New(elem.Elem().Type()).Elem();
	tmp.Set(elem.Elem())
	value := tmp.FieldByName(idList[1])
	return value, nil;
}

func New(name string) *IoC {
	newIoc := &IoC{
		name,
		make(map[string]interface{}),
	};
	globalIocContainer[name] = newIoc;
	return newIoc;
}

func Get(name string) *IoC {
	ioc := globalIocContainer[name];
	if ioc == nil {
		return New(name);
	}
	return ioc;
}