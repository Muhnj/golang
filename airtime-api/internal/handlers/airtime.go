package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PurchaseRequest struct {
	MSISDN      string `json:"msisdn"`
	Network     string `json:"network"`
	Value       string `json:"value"`
	Transaction string `json:"transaction"`
}

func (h *Handler) Purchase(c *gin.Context) {

	var req PurchaseRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp, err := h.Service.Purchase(
		req.MSISDN,
		req.Network,
		req.Value,
		req.Transaction,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}