package http

import (
  "encoding/json"
  "fmt"
  nethttp "net/http"

  "github.com/golang/glog"
)

func GetRequestParameters(r *nethttp.Request) (map[string]interface{}, error) {
  contentType := r.Header.Get("content-type")
  glog.V(1).Infof("content-type: %v\n", contentType)
  if contentType != "application/json" {
    return nil, fmt.Errorf("POST requires content-type=application/json")
  }
  decoder := json.NewDecoder(r.Body)
  jsonBody := make(map[string]interface{}, 0)
  if err := decoder.Decode(&jsonBody); err != nil {
    return nil, fmt.Errorf("Error decoding JSON body: %v", err)
  }
  return jsonBody, nil
}

func GetJsonStringParameter(jsonBody map[string]interface{}, name string) string {
  val, ok := jsonBody[name]
  if !ok {
    return ""
  }
  s, ok := val.(string)
  if !ok {
    return ""
  }
  return s
}
