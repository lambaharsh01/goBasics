package main

import (
    "encoding/json"
    "fmt"
)

type Person struct {
    Name    string `json:"name"`
    Age     int    `json:"age"`
    Email   string `json:"email"`
}

func main() {

	// convertingToJson()
	// convertingFromJson()
	// handleUnknownFields()
   
}


func convertingToJson(){ //Marshaling JSON

	person := Person{Name: "Alice", Age: 30, Email: "alice@example.com"}
    
    jsonData, err := json.Marshal(person)
    if err != nil {
        fmt.Println("Error marshaling JSON:", err)
        return
    }
    
    fmt.Println(string(jsonData));

}

func convertingFromJson(){ //Unmarshaling
	jsonData := `{"name":"Bob","age":25,"email":"bob@example.com", "extraField":"extraValue"}`;
	//will not take extraField into consideration as the Person Struct dosent have it

    var person Person
    err := json.Unmarshal([]byte(jsonData), &person)
    if err != nil {
        fmt.Println("Error unmarshaling JSON:", err)
        return
    }
    
    fmt.Println(person) // Output: {Bob 25 bob@example.com}
}


func handleUnknownFields(){
	jsonData := `{"name":"Charlie","age":40,"extraField":"extraValue"}`
    
    var result map[string]interface{}
    err := json.Unmarshal([]byte(jsonData), &result)
    if err != nil {
        fmt.Println("Error unmarshaling JSON:", err)
        return
    }
    
    fmt.Println(result) // Output: map[age:40 extraField:extraValue name:Charlie]



	// In Go, when you unmarshal JSON into a map[string]interface{}, the values are stored as interface{}, which means you need to type assert or convert them to the appropriate type to access them. In the case of the age field, it's a number, so you should first assert it as a float64, since Go's JSON unmarshaler treats numbers as float64 by default
	
	// Accessing the 'age' field

	// if age, ok := result["age"].(float64); ok {
    //     fmt.Println("Age:", age) // Output: Age: 40
    // } else {
    //     fmt.Println("Age not found or not a float64")
    // }

	// age:= result["age"].(float64);
	// fmt.Println(age)
	
	// extraField:= result["extraField"].(string);
	// fmt.Println(extraField)
	
	// result["age"]


	// for key, value := range result {
    //     fmt.Printf("Key: %s, Value: %v\n", key, value)
    // }

	for key, value := range result {
        if strVal, ok := value.(string); ok {
            fmt.Printf("%s is a string: %s\n", key, strVal)
        } else if floatVal, ok := value.(float64); ok {
            fmt.Printf("%s is a float64: %f\n", key, floatVal)
        } else if boolVal, ok := value.(bool); ok {
            fmt.Printf("%s is a bool: %t\n", key, boolVal)
        } else if value == nil {
            fmt.Printf("%s is null\n", key)
        } else {
            fmt.Printf("%s is of an unknown type\n", key)
        }
    }

}
