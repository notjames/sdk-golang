package model

type StartOpRunReq struct {
  Args  map[string]string `json:"args"`
  OpUrl string `json:"opUrl"`
}