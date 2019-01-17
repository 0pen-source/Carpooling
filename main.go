package main

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/0pen-source/Carpooling/dao"
	"github.com/0pen-source/Carpooling/service/common"
	"github.com/0pen-source/Carpooling/service/drivers"
	"github.com/0pen-source/Carpooling/service/passengers"
	"github.com/0pen-source/Carpooling/service/user"
	"github.com/gin-gonic/gin"
)

func main() {
	dao.InitializeConfiguration()
	dao.InitializeRedis()
	dao.InitializeCache()
	dao.InitOSSClient()
	runtime.GOMAXPROCS(runtime.NumCPU())
	gin.SetMode(dao.GetMODE())
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/aa", func(c *gin.Context) {
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(200, `<p><a href='https://pa-engine.zplayads.com/v1/tracking?ab_test_id=A011A3BB-9316-08D6-06DD-3DADFA6B73C1&ab_test_object_group_id=default&ad_id=476FC7D0-C133-42F5-ED04-20F8C2BF5882&ad_unit_id=19393189-C4EB-3886-60B9-13B39407064E&android_adid=b854d848-b1a9-4860-9fa2-5cebd10233a6&app_id=5C5419C7-A2DE-88BC-A311-C3E7A646F6AF&brand=Nexus+6&bundle_id=com.playableads.demo&channel_id=&channel_share_rate=0&client_id=73AB1CDF-009E-130C-4A15-94901D3470DE&country=%E7%BE%8E%E5%9B%BD&creatives_id=9D7833A2-B939-16F0-A36C-C4F06518B8A0&developer_share_rate=70&device_model=google&device_type=phone&engine_events=&height=0&industry_id=&last_status=0&network_connection_type=wifi&order_id=94799B85-ED77-B6B5-37B2-2EBCC327131E&os=Android&os_version=7.0&pay_by=cpi&pay_event=install&predictive_cpm_cipher=351d6fbfdfc8bb29&predictive_cpm_usd_cipher=213ee1d6dd84207c&predictive_cpms=2c3e8d08ad73a635&promotion_app_id=DF8DD39A-F5B2-A311-BCFE-C6E3CFB75E76&redirect_to=https%3A%2F%2Fapp.appsflyer.com%2Fcom.zplay.beatracer%3Fpid%3Dzplayads%26c%3Dplayable1&region=%E7%BE%8E%E5%9B%BD&request_id=bgvo7rjgr3riid0f9e80&settlement_price_cny=213ee1d6dd84207c&settlement_price_usd=213ee1d6dd84207c&tracking_type=click_from_video_page&unit_price_cny=fdf3d2a3b2e6380d77e13edb&unit_price_usd=6d30843dc8aa7cb0bd86265b&unit_prices=236565a50123ca1c42fe558c9e5dc515314635e89219dcaa8992c4d9bdf71370739e94f4a1ba036aa3d374d2aa9aa16e1b430440d771581759cfb497d92d435c8d1374109cff22a8&width=0'>
<img border="0" src="/i/eg_buttonnext.gif" /></p>`)
	})

	router.Use(
		common.MarkTesting(),
	)
	webServer := router.Group("/v1/webserver")
	{
		webServer.POST("/searchtrip", passengers.SearchTrip)
	}
	noLoginUserGroup := router.Group("/v1/user")
	{
		noLoginUserGroup.POST("/creatpassengerstrip", passengers.CreatTrip)
		noLoginUserGroup.POST("/creatdriverstrip", drivers.CreatTrip)
		noLoginUserGroup.POST("/checkphone", user.Phonetest)
		noLoginUserGroup.POST("/login", user.Login)
		noLoginUserGroup.POST("/getcode", user.GetVerificationCode)
		noLoginUserGroup.POST("/checkcode", user.CheckCode)
		noLoginUserGroup.POST("/register", user.Register)
	}
	LoginUserGroup := router.Group("/v1/user")
	{
		LoginUserGroup.Use(
			common.Auth(),
		)
		LoginUserGroup.POST("/setinformation", user.SetInformation)
		LoginUserGroup.POST("/upload", user.Upload)
	}

	passenger := router.Group("/v1/passengers")
	{
		passenger.Use(
			common.Auth(),
		)
		passenger.POST("/creattrip", passengers.CreatTrip)
		passenger.POST("/upload", user.Upload)
		passenger.POST("/index", passengers.Index)
		passenger.POST("/searchtrip", passengers.SearchTrip)
		passenger.POST("/connected", passengers.Connected)
		passenger.POST("/getconnected", passengers.GetConnecteds)
		passenger.POST("/mytrip", passengers.MyTrip)
		passenger.POST("/getphone", passengers.GetPhone)

	}

	driver := router.Group("/v1/driver")
	{
		driver.Use(
			common.Auth(),
		)
		driver.POST("/creattrip", drivers.CreatTrip)
		driver.POST("/upload", user.Upload)
		driver.POST("/index", drivers.Index)
		driver.POST("/searchtrip", drivers.SearchTrip)
		driver.POST("/connected", drivers.Connected)
		driver.POST("/getconnected", drivers.GetConnecteds)
		driver.POST("/mytrip", drivers.MyTrip)
		driver.POST("/getphone", drivers.GetPhone)
	}

	server := &http.Server{
		Addr:    dao.GetAddress(),
		Handler: router,
	}
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}
