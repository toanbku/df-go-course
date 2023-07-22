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

func reorderNames(firstName string, lastName string, countryCode string, middleName string) string {
  var result []string;
  switch (countryCode) {
    case "VN": {
      result = []string{lastName};
      if middleName != "" {
        result = append(result, middleName);
      }
      result = append(result, firstName)
    }
    case "US": {
      result = []string{firstName};
      if middleName != "" {
        result = append(result, middleName);
      }
      result = append(result, lastName)
    }
  }
  
  return strings.Join(result, " ");
} 

func main() {
  props := os.Args[1:]
  
  lenProps := len(props)
  if lenProps < 3 {
    fmt.Println("Not enough required params")
    return;
  }
  
  countryCode := props[lenProps-1]
  if (!isSupportedCountry(countryCode)){
    fmt.Println("Not supported country")
    return;
  }

  firstName := props[0]
  lastName := props[1]
  middleName := strings.Join(props[2:lenProps-1], " ")

  fmt.Println(reorderNames(firstName, lastName, countryCode, middleName))
}
