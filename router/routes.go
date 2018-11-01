package router

import (
	"github.com/gin-gonic/gin"

	"github.com/sovrinbloc/kairos/config"
	"github.com/sovrinbloc/kairos/controller/status"
	"github.com/sovrinbloc/kairos/controller"
	"github.com/sovrinbloc/kairos/model"
	"fmt"
)

func Route(router *gin.Engine) {
	api := router.Group(config.APIDefaultRoute)
	{
		api.GET("/is-up", status.IsUp)
		api.POST("", controller.GetIntent)
	}

	test := router.Group(config.APITestRoute)
	{
		bools := test.Group("/success")
		{
			bools.POST("/true", func(c *gin.Context) {
				c.JSON(200, "true")
			})
			bools.POST("/false", func(c *gin.Context) {
				c.JSON(400, "false")
			})
		}
		stripe := test.Group("/stripe")
		{
			stripe.POST("/createcharge", func(c *gin.Context) {

				response := model.IntentResponse{
					Functions: map[int]string{
						1: "StripeCreateCharge",
					},
					Data: map[string]interface{}{
						"StripeCreateCharge": map[string]string{
							"Email":       "josephalai@gmail.com",
							"Amount":      "1999",
							"Title":       "Joseph Alai's Sovrin Mind",
							"Name":        "Joseph Alai",
							"Currency":    "usd",
							"Description": "Something",
							"PageData":    "sk_test_kqcgtNglf8UIlDtPs7ziTKzO",
						}, "Slack": map[string]string{
							"API":          "slack_ck4chs7",
							"Notification": "Hello world!",
						},
					},
				}

				c.JSON(200, response)
			})

			stripe.POST("/addkey", func(c *gin.Context) {
				response := model.IntentResponse{
					Functions: map[int]string{
						1: "StripeSaveAPIKey",
					},
					Data: map[string]interface{}{
						"StripeSaveAPIKey": map[string]string{},
						"Endpoint": fmt.Sprintf("http://localhost:%s/test/success/true", config.ServerPort),

					},
				}

				c.JSON(200, response)
			})

			test.POST("/get-intent", func(c *gin.Context) {
				type ReqData struct {
					KairosKey string `json:"kairosKey" binding:"required,min=1"`
					IntentKey string `json:"intentKey" binding:"required,min=1"`
				}
				reqData := ReqData{}
				if err := c.ShouldBindJSON(&reqData); err != nil {
					return
				}

				if reqData.IntentKey == "tok_stripe_save_key" {
					response := model.IntentResponse{
						Functions: map[int]string{
							1: "StripeSaveAPIKey",
						},
						Data: map[string]interface{}{
							"StripeSaveAPIKey": map[string]string{
								"Endpoint": fmt.Sprintf("http://localhost:%s/test/success/true", config.ServerPort),

							},
						},
					}

					c.JSON(200, response)
					return
				}
				if reqData.IntentKey == "tok_stripe_subscription" {
					response := model.IntentResponse{
						Functions: map[int]string{
							1: "StripeCreateSubscription",
						},
						Data: map[string]interface{}{
							"StripeCreateSubscription": map[string]string{
								"plan":     "plan_123",
								"apiKey":   "sk_test_kqcgtNglf8UIlDtPs7ziTKzO",
								"Endpoint": fmt.Sprintf("http://localhost:%s/test/success/true", config.ServerPort),

							},
						},
					}

					c.JSON(200, response)
					return
				}

				if reqData.IntentKey == "tok_stripe_charge" {
					response := model.IntentResponse{
						Functions: map[int]string{
							1: "StripeCreateCharge",
						},
						Data: map[string]interface{}{
							"StripeCreateCharge": map[string]string{
								"Email":       "josephalai@gmail.com",
								"Amount":      "1999",
								"Title":       "Joseph Alai's Sovrin Mind",
								"Name":        "Joseph Alai",
								"Currency":    "usd",
								"Description": "Something",
								"Endpoint": fmt.Sprintf("http://localhost:%s/test/success/true", config.ServerPort),

								"apiKey":      "sk_test_kqcgtNglf8UIlDtPs7ziTKzO",
							}, "Slack": map[string]string{
								"API":          "slack_ck4chs7",
								"Notification": "Hello world!",
								"Endpoint": fmt.Sprintf("http://localhost:%s/test/success/true", config.ServerPort),

							},
						},
					}

					c.JSON(200, response)
				}
				if reqData.IntentKey == "tok_slack_webhook" {
					response := model.IntentResponse{
						Functions: map[int]string{
							1: "SlackSendNotification",
						},
						Data: map[string]interface{}{
							"SlackSendNotification": map[string]string{
								"text":     "new purchase",
								"channel":  "general",
								"apiKey":   "xoxp-388044397831-386300920016-466668367249-aeef50807e1f333200f0e340949542e2",
								"Endpoint": fmt.Sprintf("http://localhost:%s/test/success/true", config.ServerPort),

							},
						},
					}

					c.JSON(200, response)
				}
				if reqData.IntentKey == "tok_combo" {
					response := model.IntentResponse{
						Functions: map[int]string{
							1: "SlackSendNotification",
							2: "StripeCreateCharge",
							3: "StripeSaveAPIKey",
							4: "StripeCreateCharge",
						},
						Data: map[string]interface{}{
							"SlackSendNotification": map[string]string{
								"text":     "new purchase",
								"channel":  "general",
								"apiKey":   "xoxp-388044397831-386300920016-466668367249-aeef50807e1f333200f0e340949542e2",
								"Endpoint": fmt.Sprintf("http://localhost:%s/test/success/true", config.ServerPort),

							},
							"StripeCreateCharge": map[string]string{
								"Email":       "josephalai@gmail.com",
								"Amount":      "1999",
								"Title":       "Joseph Alai's Sovrin Mind",
								"Name":        "Joseph Alai",
								"Currency":    "usd",
								"Description": "Something",
								"apiKey":      "sk_test_kqcgtNglf8UIlDtPs7ziTKzO",
								"Endpoint": fmt.Sprintf("http://localhost:%s/test/success/true", config.ServerPort),

							},
							"StripeSaveAPIKey": map[string]string{
								"Endpoint": fmt.Sprintf("http://localhost:%s/test/success/true", config.ServerPort),

							},
						},
					}

					c.JSON(200, response)
				}
				if reqData.IntentKey == "tok_cms" {
					response := model.IntentResponse{
						Functions: map[int]string{
							1: "GetPageData",
						},
						Data: map[string]interface{}{
							"GetPageData": map[string]interface{}{
								"Endpoint": fmt.Sprintf("http://localhost:%s/test/success/true", config.ServerPort),
								"pageData": map[string]interface{}{
									"h1": "Hello",
									"h2": "World",
									"p": map[string]interface{}{
										"class":   "joseph",
										"id":      123,
										"v-model": "func(joe){}",
									},
								},
							},
						},
					}

					c.JSON(200, response)
				}
			})
		}

	}
}
