package main 


import (
"testing" 
"os"
)

func TestLoadToBucket(t testing.T)  {
    
}

func TestCreateBucket(t testing.T)  {

}


func TestGetBucketObject(t testing.T){

  
   mockFile , err := os.Create("test.txt") 
   mockFile.Write([]byte("this is test files content"))
   defer mockFile.Close()
   defer os.Remove(mockFile)
   
    // Push mock object to bucket 

   
   
   // Call get bucket method 
   if expectedObjectName != actualObjectName {
  t.Errorf("expected %s as object name but got %s", expectedObjectName, actualObjectName)
}

} 
