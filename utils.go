package main

func IsInSlice[T any](str T, arr []T) bool {
for _,s := range arr {
  if str == s {
  return true 
}
  
}
return false 
}

