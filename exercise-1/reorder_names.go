package main

import (
  "fmt"
  "os"
  "strings"
)

func isSupportedCountry(countryCode string) bool{
  switch (countryCode) {
    case "VN", "US":
      return true;
    default:
      return false
  }
}

func buildName(firstName string, middleName string, lastName string) string {
  var result []string;
  result = []string{firstName};
  if middleName != "" {
    result = append(result, middleName);
  }
  result = append(result, lastName)

  return strings.Join(result, " ");
}

func reorderNames(firstName string, lastName string, countryCode string, middleName string) string {
  switch (countryCode) {
    case "VN": {
      return buildName(lastName, middleName, firstName)
    }
    case "US": {
      return buildName(firstName, middleName, lastName)
    }
  }

  return ""
} 

func main() {
  props := os.Args[1:]
  
  lenProps := len(props)
  if lenProps < 3 {
    fmt.Println("Not enough required params")
    return;
  }
  
  countryCode := strings.ToUpper(props[lenProps-1])
  if (!isSupportedCountry(countryCode)){
    fmt.Println("Not supported country")
    return;
  }

  firstName := props[0]
  lastName := props[1]
  middleName := strings.Join(props[2:lenProps-1], " ")

  fmt.Println(reorderNames(firstName, lastName, countryCode, middleName))
}
