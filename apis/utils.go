package apis

import (
	"reflect"
)

func getCleanParams(ginParams map[string][]string) map[string]any {
	// Convert the returned params's VALUE from slice to string
	cleanParams := make(map[string]any)

	for key, value := range ginParams {
		if _, exists := ginParams["page"]; exists {
			continue
		}
		if _, exists := ginParams["page_size"]; exists {
			continue
		}
		cleanParams[key] = value[0]
	}

	return cleanParams
}

func getVaiableType(variable any) reflect.Kind {
	return reflect.ValueOf(variable).Kind()
}
