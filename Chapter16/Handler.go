package main

import (
	"book-code/Chapter13/13-4/conf"
	"book-code/Chapter13/13-4/config"
	"book-code/Chapter13/13-4/handler"
	"book-code/Chapter13/13-4/model"
	"book-code/Chapter13/13-4/repository"
	"book-code/Chapter13/13-4/service"
	"book-code/Chapter13/13-4/service/dp_service"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"log"
)

var (
	DB                   *gorm.DB
	AccountHandler       handler.AccountHandler
	DiscountHandler      handler.DiscountHandler
	RestaurantNavHandler handler.RestaurantNavHandler
	HotelDetailHandler   handler.HotelDetailHandler
	TeamDetailHandler    handler.TeamDetailHandler
	OrderSeatHandler     handler.OrderSeatHandler
	TakeOutHandler       handler.TakeOutHandler
	MeHandler            handler.MeHandler
	SuggestFoodHandler   handler.SuggestFoodHandler
	GuessHandler         handler.GuessHandler
	NavHandler           handler.NavHandler
	PostTeamOrderHandler handler.PostTeamOrderHandler
)

func initViper() {
	if err := config.Init(""); err != nil {
		panic(err)
	}
}

func initDB() {
	fmt.Println("数据库 init")
	var err error
	conf := &conf.DBConf{
		Host:     viper.GetString("database.host"),
		User:     viper.GetString("database.username"),
		Password: viper.GetString("database.password"),
		DbName:   viper.GetString("database.name"),
	}

	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&charset=utf8&parseTime=%t&loc=%s",
		conf.User,
		conf.Password,
		conf.Host,
		conf.DbName,
		true,
		"Local")

	DB, err = gorm.Open("mysql", config)
	if err != nil {
		log.Fatalf("connect error: %v\n", err)
	}
	DB.SingularTable(true)
	fmt.Println("数据库 init 结束...")
}

func initHandler() {
	AccountHandler = handler.AccountHandler{
		Srv: &service.AccountService{
			Repo: &repository.AccountModelRepo{
				DB: model.DataBase{MyDB: DB},
			},
		}}
	DiscountHandler = handler.DiscountHandler{
		Srv: &dp_service.DiscountService{
			Repo: &repository.Discount{
				DB: model.DataBase{
					MyDB: DB,
				},
			},
		},
	}
	RestaurantNavHandler = handler.RestaurantNavHandler{
		Srv: &dp_service.RestaurantService{
			Repo: &repository.RestaurantNavRepo{
				DB: model.DataBase{
					MyDB: DB,
				},
			},
			ItemRepo: &repository.RestaurantTabItemRepo{
				DB: model.DataBase{
					MyDB: DB,
				},
			},
		},
	}
	HotelDetailHandler = handler.HotelDetailHandler{
		Srv: &dp_service.HotelService{
			Repo: &repository.HotelRepo{
				DB: model.DataBase{
					MyDB: DB,
				},
			},
			TeamRepo: &repository.TeamRepo{
				DB: model.DataBase{
					MyDB: DB,
				},
			},
			SuggestFoodRepo: &repository.SuggestFoodRepo{
				DB: model.DataBase{
					MyDB: DB,
				},
			},
			CommentTagRepo: &repository.CommentTagRepo{
				DB: model.DataBase{
					MyDB: DB,
				},
			},
			CommentRepo: &repository.CommentRepo{
				DB: model.DataBase{
					MyDB: DB,
				},
			},
			MarketRepo: &repository.MarketRepo{
				DB: model.DataBase{
					MyDB: DB,
				},
			},
		},
	}
	TeamDetailHandler = handler.TeamDetailHandler{
		Srv: &dp_service.TeamService{
			Repo: &repository.TeamRepo{
				DB: model.DataBase{
					MyDB: DB,
				},
			},
		},
	}
	OrderSeatHandler = handler.OrderSeatHandler{
		Srv: &dp_service.OrderSeatService{
			Repo: &repository.OrderSeatRepo{
				DB: model.DataBase{
					MyDB: DB,
				},
			},
		},
	}
	TakeOutHandler = handler.TakeOutHandler{
		Srv: &dp_service.TakeOutService{
			Repo: &repository.TakeOutItemRepo{
				DB: model.DataBase{
					MyDB: DB,
				},
			},
		},
	}
	MeHandler = handler.MeHandler{
		Srv: &dp_service.MeService{
			Repo: &repository.MeItemRepo{
				DB: model.DataBase{
					MyDB: DB,
				},
			},
		},
	}
	SuggestFoodHandler = handler.SuggestFoodHandler{
		Srv: &dp_service.SuggestFoodService{
			Repo: &repository.SuggestFoodRepo{
				DB: model.DataBase{
					MyDB: DB,
				},
			},
		},
	}
	GuessHandler = handler.GuessHandler{
		Srv: &dp_service.GuessService{
			Repo: &repository.GuessRepo{
				DB: model.DataBase{
					MyDB: DB,
				},
			},
		},
	}
	NavHandler = handler.NavHandler{
		Srv: &dp_service.ListNavItemService{
			Repo: &repository.ListNavItemRepo{
				DB: model.DataBase{
					MyDB: DB,
				},
			},
		},
	}
	PostTeamOrderHandler = handler.PostTeamOrderHandler{
		Srv: &dp_service.PostTeamOrderService{
			Repo: &repository.TeamPostOrderRepo{
				DB: model.DataBase{
					MyDB: DB,
				},
			},
		},
	}
}

func init() {
	initViper()
	initDB()
	initHandler()
}
