package config

import "fmt"

func Parse(format Format, data string) {
	// el switch pero con map
	parseByFunc := parsersMap[format]

	// parsear ya sea yaml o json obtendré un map[string]interface{} para validar después con mi struct
	parseData, err := parseByFunc(data)
	if err != nil {
		fmt.Println("Error parsing data:", err)
		return
	}

	fmt.Println("Parsed data:", parseData)
}
