// enumerations (enums)

package main;

import "fmt"

type Color int

const (
    Red Color = iota
    Green
    Blue
)

//Red Color = iota: iota is a special constant in Go that increments by 1 for each constant declaration in a block. So, in this case, Red is assigned the value of iota, which is 0. This means Red is equal to 0.


func main(){

	var palet Color = Blue
    fmt.Println(palet);
	
}

// WHY ENUMS? 
// 1. Readability: By using named constants, your code becomes more readable. Instead of using magic numbers (e.g., 0, 1, 2) in your code, you can use descriptive names like Red, Green, and Blue.
// 2. Type safety: Since Color is a distinct type, the Go compiler will prevent you from assigning a value of the wrong type to a variable of type Color. This helps catch errors at compile-time rather than runtime.
// 3. Enum-like behavior: Enums can be used to define a set of valid values for a variable. In this case, a Color variable can only take on one of the three values: Red, Green, or Blue.

// Enums are used in a variety of situations where you need to define a set of distinct values or categories. Here are some common use cases:

// Defining a set of states or statuses: Enums can be used to define a set of states or statuses, such as Pending, InProgress, Completed, or Failed. This can be useful in workflow management, task tracking, or error handling.
// Representing a set of options or choices: Enums can be used to represent a set of options or choices, such as Yes, No, or Maybe. This can be useful in user interfaces, surveys, or configuration settings.
// Defining a set of error codes or messages: Enums can be used to define a set of error codes or messages, such as Success, Failure, InvalidInput, or Timeout. This can be useful in error handling, logging, or debugging.
// Representing a set of categories or labels: Enums can be used to represent a set of categories or labels, such as Male, Female, or Other. This can be useful in data analysis, reporting, or visualization.
// Defining a set of permissions or access levels: Enums can be used to define a set of permissions or access levels, such as Admin, Moderator, User, or Guest. This can be useful in access control, authentication, or authorization.
// Representing a set of data formats or types: Enums can be used to represent a set of data formats or types, such as JSON, XML, CSV, or PDF. This can be useful in data processing, serialization, or deserialization.
// Defining a set of logging levels or severities: Enums can be used to define a set of logging levels or severities, such as Debug, Info, Warning, Error, or Fatal. This can be useful in logging, monitoring, or debugging.
// Representing a set of currencies or units: Enums can be used to represent a set of currencies or units, such as USD, EUR, GBP, or JPY. This can be useful in financial calculations, conversions, or formatting.
// Defining a set of time zones or regions: Enums can be used to define a set of time zones or regions, such as UTC, EST, PST, or CET. This can be useful in date and time calculations, formatting, or conversions.
// Representing a set of device types or platforms: Enums can be used to represent a set of device types or platforms, such as Desktop, Mobile, Tablet, or TV. This can be useful in device detection, responsive design, or platform-specific development.