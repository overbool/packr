package packr

var packData = map[string]map[string]string{}

func PackData(root, path, data string) {
	if _, ok := packData[root]; !ok {
		packData[root] = map[string]string{}
	}

	packData[root][path] = data
}