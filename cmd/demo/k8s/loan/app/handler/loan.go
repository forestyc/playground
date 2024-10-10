package handler

import (
	"github.com/forestyc/playground/cmd/demo/k8s/loan/app/context"
	"github.com/forestyc/playground/cmd/demo/k8s/loan/app/model"
	"github.com/forestyc/playground/cmd/demo/k8s/loan/app/service"
	"github.com/forestyc/playground/pkg/core/message"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Loan struct {
	loanLoanService *service.LoanBasicInfo
	loanService     *service.Loan
}

func NewLoan(ctx *context.Context) *Loan {
	return &Loan{
		loanLoanService: service.NewLoanBasicInfo(ctx),
		loanService:     service.NewLoan(ctx),
	}
}

func (l *Loan) Register(engine *gin.Engine) {
	uriLoan := "/loan"
	engine.GET(uriLoan, l.getLoan())
	engine.GET(uriLoan+"/list", l.getLoanList())
	engine.POST(uriLoan, l.createLoan())
	engine.PUT(uriLoan, l.modifyLoan())
	engine.DELETE(uriLoan, l.deleteLoan())
}

func (l *Loan) getLoan() gin.HandlerFunc {
	return func(c *gin.Context) {
		query := c.DefaultQuery("id", "0")
		id, _ := strconv.ParseInt(query, 10, 64)
		if loan, err := l.loanService.GetById(id); err != nil {
			c.JSON(http.StatusOK, message.FailedWithMessage(err.Error()))
		} else {
			c.JSON(http.StatusOK, message.SuccessWithObject(loan))
		}
	}
}

func (l *Loan) createLoan() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.CreateLoanReq
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusOK, message.FailedWithMessage(err.Error()))
			return
		}
		if err := l.loanService.Create(req); err != nil {
			c.JSON(http.StatusOK, message.FailedWithMessage(err.Error()))
			return
		}
		c.JSON(http.StatusOK, message.Success())
	}
}

func (l *Loan) modifyLoan() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.ModifyLoanReq
		if err := c.ShouldBind(&req); err != nil {
			c.JSON(http.StatusOK, message.FailedWithMessage(err.Error()))
			return
		}
		if err := l.loanService.Modify(req); err != nil {
			c.JSON(http.StatusOK, message.FailedWithMessage(err.Error()))
			return
		}
		c.JSON(http.StatusOK, message.Success())
	}
}

func (l *Loan) deleteLoan() gin.HandlerFunc {
	return func(c *gin.Context) {
		query := c.DefaultQuery("id", "0")
		id, _ := strconv.ParseInt(query, 10, 64)
		if err := l.loanService.Delete(id); err != nil {
			c.JSON(http.StatusOK, message.FailedWithMessage(err.Error()))
		} else {
			c.JSON(http.StatusOK, message.Success())
		}
	}
}

func (l *Loan) getLoanList() gin.HandlerFunc {
	return func(c *gin.Context) {
		if loanList, err := l.loanService.GetList(); err != nil {
			c.JSON(http.StatusOK, message.FailedWithMessage(err.Error()))
		} else {
			c.JSON(http.StatusOK, message.SuccessWithObject(loanList))
		}
	}
}
