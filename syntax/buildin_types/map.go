package main

func Map() {
	m1 := map[string]string{
		"key1": "value1",
		"key2": "value2",
		"key3": "value3",
	}
	println(m1)

	m2 := make(map[string]string, 4)
	m2["key1"] = "value1"

	val1, ok := m1["key1"]
	println(val1, ok)
}
