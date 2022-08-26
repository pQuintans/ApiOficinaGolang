package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pquintans/OficinaGolang/model"
	"github.com/pquintans/OficinaGolang/repository"
)

func PostAluno(c *gin.Context) error {
	var usuario model.Usuario
	if err := c.BindJSON(&usuario); err != nil {
		return err
	}

	if err := repository.InsertUsuario(usuario); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	c.String(http.StatusOK, "Deu tudo certo")
	return nil
}
