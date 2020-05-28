package middleware

import "github.com/gin-gonic/gin"

type InstallInterceptorImpl struct {

}

func NewInstallInterceptor() InstallInterceptor {
	return &InstallInterceptorImpl{}
}

func (m *InstallInterceptorImpl) Handle(c *gin.Context) {
	panic("implement me")
}

func (m *InstallInterceptorImpl) IsNeededInstall() (isNeeded bool, err error) {
	panic("implement me")
}
