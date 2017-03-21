package cache



var mapCache map[string]string = make(map[string]string)

func Get(key string) string {
	return mapCache[key]
}

func Set(key,value string) bool {
	mapCache[key] = value
	return true;
}