package resources

func (src *Source) FetchEntity(key string) (interface{}, error) {
	return src.get("entity", "entities", key)
}

func (src *Source) FetchEntities(pagesize string) (map[string]interface{}, error) {
	res, err := src.get("entity", "entities", pagesize)

	if err != nil {
		return nil, err
	}

	return res.(map[string]interface{}), nil
}
