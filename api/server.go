package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	db "github.com/VinCPR/backend/db/sqlc"
	"github.com/VinCPR/backend/token"
	"github.com/VinCPR/backend/util"
)

type Server struct {
	config     util.Config
	tokenMaker token.ITokenMaker
	store      *db.Store
	router     *gin.Engine
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, POST, PUT, DELETE, OPTIONS, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func NewServer(config util.Config, store *db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot initialize token maker")
	}

	server := &Server{config: config, tokenMaker: tokenMaker, store: store}
	router := gin.Default()
	router.Use(CORS())
	routerV1 := router.Group(config.BasePath)
	{
		routerV1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	// Authentication
	{
		routerV1.POST("/users", server.createUser)
		routerV1.POST("/users/login", server.loginUser)
	}
	{
		routerV1.POST("/academic_year", server.createAcademicYear)
		routerV1.GET("/academic_year/list", server.listAcademicYears)
		routerV1.GET("/academic_year/calendar", server.getAcademicCalendar)
	}
	{
		routerV1.GET("/group/list", server.listGroupsByAcademicYear)
	}
	{
		routerV1.GET("/period/list", server.listPeriodsByAcademicYear)
	}
	{
		routerV1.GET("/block/list", server.listBlocksByAcademicYear)
	}
	{
		routerV1.GET("/hospital/list", server.listHospitalsByName)
	}
	{
		routerV1.GET("/specialty/list", server.listSpecialtiesByName)
	}
	{
		routerV1.GET("/service/list/specialty", server.listServicesbySpecialtyID)
		routerV1.GET("/service/list/hospital", server.listServicesbyHospitalID)
		routerV1.GET("/service/list/specialty_and_hospital", server.listServicesBySpecialtyIDAndHospitalID)
	}
	{
		routerV1.GET("/student/list/name", server.listStudentsByName)
		routerV1.GET("/student/list/studentID", server.listStudentsByStudentID)
	}
	{
		routerV1.GET("/attending/list", server.listAttendingsByName)
	}
	// authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	server.router = router
	return server, nil
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
