package ui

import "webmalc/no-more-excuses/internal/dto"

type uiInterface interface {
	ShowConfig()
}

type appsGetter interface {
	GetApps() map[string]dto.App
}
