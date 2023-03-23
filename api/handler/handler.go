package handler

import (
	"github.com/labstack/echo"
	"github.com/lukamindo/pet-reminder/app/domain"
	"github.com/lukamindo/pet-reminder/helper/conn"
	"github.com/lukamindo/pet-reminder/helper/server"
)

func New(e *echo.Echo) {
	authenticationGroup(e.Group("/auth"))
	adminGroup(e.Group("/admin"))
}

func adminGroup(g *echo.Group) {
	g.GET("/test", test())
}

func authenticationGroup(g *echo.Group) {
	as := domain.NewAuthService(conn.New())
	g.POST("/signin", signInUser(as))
	// g.POST("/signup", test())
	// g.GET("/userDetails", test())
}

func test() echo.HandlerFunc {
	return func(c echo.Context) error {
		ret := "hi"
		return server.Success(c, ret)
	}
}

func signInUser(s domain.AuthService) echo.HandlerFunc {
	return func(c echo.Context) error {
		token, err := s.Login(c.Request().Context())
		if token == nil {
			return err
		}
		if err != nil {
			return err
		}
		return server.Success(c, "shemovida")
	}
}

// // SignInUser Used for Signing In the Users
// func SignInUser(response http.ResponseWriter, request *http.Request) {
// 	var loginRequest LoginParams
// 	var result UserDetails
// 	var errorResponse = ErrorResponse{
// 		Code: http.StatusInternalServerError, Message: "It's not you it's me.",
// 	}

// 	decoder := json.NewDecoder(request.Body)
// 	decoderErr := decoder.Decode(&loginRequest)
// 	defer request.Body.Close()

// 	if decoderErr != nil {
// 		returnErrorResponse(response, request, errorResponse)
// 	} else {
// 		errorResponse.Code = http.StatusBadRequest
// 		if loginRequest.Email == "" {
// 			errorResponse.Message = "Last Name can't be empty"
// 			returnErrorResponse(response, request, errorResponse)
// 		} else if loginRequest.Password == "" {
// 			errorResponse.Message = "Password can't be empty"
// 			returnErrorResponse(response, request, errorResponse)
// 		} else {

// 			collection := Client.Database("test").Collection("users")

// 			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 			var err = collection.FindOne(ctx, bson.M{
// 				"email":    loginRequest.Email,
// 				"password": loginRequest.Password,
// 			}).Decode(&result)

// 			defer cancel()

// 			if err != nil {
// 				returnErrorResponse(response, request, errorResponse)
// 			} else {
// 				tokenString, _ := CreateJWT(loginRequest.Email)

// 				if tokenString == "" {
// 					returnErrorResponse(response, request, errorResponse)
// 				}

// 				var successResponse = SuccessResponse{
// 					Code:    http.StatusOK,
// 					Message: "You are registered, login again",
// 					Response: SuccessfulLoginResponse{
// 						AuthToken: tokenString,
// 						Email:     loginRequest.Email,
// 					},
// 				}

// 				successJSONResponse, jsonError := json.Marshal(successResponse)

// 				if jsonError != nil {
// 					returnErrorResponse(response, request, errorResponse)
// 				}
// 				response.Header().Set("Content-Type", "application/json")
// 				response.Write(successJSONResponse)
// 			}
// 		}
// 	}
// }
