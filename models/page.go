package models

type PageInfo struct {
	HasNext   bool
	DocsCount int
}

type Page struct {
	PageInfo  PageInfo
	AnimeList []Media
}
