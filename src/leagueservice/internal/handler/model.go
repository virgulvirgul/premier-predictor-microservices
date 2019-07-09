package handler

type addLeagueRequest struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type leagueRequest struct {
	Id  string `json:"id"`
	Pin int64  `json:"pin"`
}

type renameRequest struct {
	Pin  int64  `json:"pin"`
	Name string `json:"name"`
}
