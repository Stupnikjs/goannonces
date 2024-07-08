package api 


import (
  "testing"
  )



func TestParseJsonReq(t *testing.T){
    // create mockrequest 
    req := httptest.NewRequest() 
    /* type POST
    
             {
                "action": "update",
                "object": {
                    "type" : "track",
                    "id": trackid,
                    "field" : "tag",
                    "body": input.value
                }
            } 

     */

    // pass json to request 

    jsonReq, err := ParseJsonReq(req)

  if err != nil {
  t.Errorf("expected no error but got %s", err.Error())
    }

  if jsonReq.Action != "update" {
    t.Errorf("expected update action but go %s", jsonReq.Action) 
    }

             
    }
  
