package db

import (
	"fmt"
	"log"
	"time"

	"opslaundry/pkg/models"
	"opslaundry/pkg/models/views"
	"opslaundry/pkg/utils"

	"github.com/google/uuid"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, string) {
	fmt.Println(" __________________________________________________________________________________")
	fmt.Println("|                                                                                  |")
	fmt.Println("|   #       #  # # # # #  #           # # #     # # # #   #       #  # # # # #     |")
	fmt.Println("|   #       #  #          #         #          #       #  # #   # #  #             |")
	fmt.Println("|   #   #   #  #          #         #          #       #  #  # #  #  #             |")
	fmt.Println("|   #  # #  #  # # # #    #         #          #       #  #   #   #  # # # #       |")
	fmt.Println("|   # #   # #  #          #         #          #       #  #       #  #             |")
	fmt.Println("|   #       #  # # # # #  # # # #     # # #     # # # #   #       #  # # # # #     |")
	fmt.Println("|                                                                                  |")
	fmt.Println("|                                       to                                         |")
	fmt.Println("|                                   OpsLaundry                                     |")
	fmt.Println("|                         created by: putraops@gmail.com                           |")
	fmt.Println("|                                   @putraops                                      |")
	fmt.Println("|__________________________________________________________________________________|")
	fmt.Println("|                                                                                  |")
	fmt.Println("|                                Open Connection                                   |")
	fmt.Println("|                                  Connecting...                                   |")

	config, err := LoadDBConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
		panic("")
	}

	DB_DIAL := config.DB_DIAL
	DB_DRIVER := config.DB_DRIVER
	_ = config.DB_USERNAME
	DB_PASSWORD := config.DB_PASSWORD
	DB_HOST := config.DB_HOST
	DB_NAME := config.DB_NAME
	DB_PORT := config.DB_PORT

	dsn := fmt.Sprintf("%v://%v:%v@%v:%v/%v",
		DB_DIAL,
		DB_DRIVER,
		DB_PASSWORD,
		DB_HOST,
		DB_PORT,
		DB_NAME)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
		fmt.Println("|__________________________________________________________________________________|")
		fmt.Println("|                                                                                  |")
		fmt.Println("|                          Failed to connect to database!                          |")
		fmt.Println("|__________________________________________________________________________________|")
		panic("")
	}

	fmt.Println("|                                    Connected...                                  |")
	fmt.Println("|__________________________________________________________________________________|")
	fmt.Println("")
	fmt.Println("")

	dbMigration(db)
	initApplication(db)

	return db, ":3000"
}

func LoadDBConfig() (config Config, err error) {
	viper.AddConfigPath("./pkg/config/envs")
	viper.SetConfigName("db")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

func LoadInitConfig() (config InitConfig, err error) {
	viper.AddConfigPath("./pkg/config/envs")
	viper.SetConfigName("init")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

func initApplication(db *gorm.DB) {
	initConfig, err := LoadInitConfig()
	if err != nil {
		log.Fatal("Failed to applicationConfig config:", err)
		panic("")
	}

	DEFAULT_ORGANIZATION := initConfig.DEFAULT_ORGANIZATION
	DEFAULT_ADMINISTRATOR_USERNAME := initConfig.DEFAULT_ADMINISTRATOR_USERNAME
	DEFAULT_ADMINISTRATOR_EMAIL := initConfig.DEFAULT_ADMINISTRATOR_EMAIL
	DEFAULT_ADMINISTRATOR_PASSWORD := initConfig.DEFAULT_ADMINISTRATOR_PASSWORD

	var defaultOrganization = models.Organization{}
	if err := db.Session(&gorm.Session{SkipHooks: true}).Where("name = ?", DEFAULT_ORGANIZATION).First(&defaultOrganization).Error; err != nil {
		organizationId := uuid.New().String()
		tenantId := uuid.New().String()
		teamId := uuid.New().String()
		adminId := uuid.New().String()
		isActive := true
		isDefault := true
		isLocked := false
		dateNow := time.Now().UTC()

		fmt.Println("================================================================================")
		fmt.Println("============================ Setup Organization... =============================")
		fmt.Println("================================================================================")
		fmt.Println("Creating Default Organization...")
		fmt.Println("================================================================================")
		defaultOrganization = models.Organization{
			Id:          organizationId,
			Name:        DEFAULT_ORGANIZATION,
			Description: "",
			CreatedBy:   adminId,
			CreatedAt:   &dateNow,
			EntityId:    uuid.New().String(),
			IsActive:    &isActive,
			IsDefault:   &isDefault,
		}
		db.Session(&gorm.Session{SkipHooks: true}).Create(&defaultOrganization)
		fmt.Println("default organization has been created!")
		fmt.Println("================================================================================")

		fmt.Println("Creating Default Tenant...")
		var defaultTenant = models.Tenant{
			Id:             tenantId,
			OrganizationId: organizationId,
			Name:           DEFAULT_ORGANIZATION,
			Description:    "",
			CreatedBy:      adminId,
			CreatedAt:      &dateNow,
			EntityId:       uuid.New().String(),
			IsActive:       &isActive,
			IsDefault:      &isDefault,
		}
		db.Session(&gorm.Session{SkipHooks: true}).Create(&defaultTenant)
		fmt.Println("default tenant has been created!")
		fmt.Println("================================================================================")

		fmt.Println("Creating default Team...")
		var defaultTeam = models.Team{
			Id:             teamId,
			OrganizationId: organizationId,
			Name:           DEFAULT_ORGANIZATION,
			Description:    "",
			CreatedBy:      adminId,
			CreatedAt:      &dateNow,
			IsActive:       &isActive,
			IsDefault:      &isDefault,
		}
		db.Session(&gorm.Session{SkipHooks: true}).Create(&defaultTeam)
		fmt.Println("default team has been created!")
		fmt.Println("================================================================================")

		fmt.Println("Checking Administrator Account...")
		var superAdmin = models.ApplicationUser{}
		if err := db.Session(&gorm.Session{SkipHooks: true}).Where("Username = ?", DEFAULT_ADMINISTRATOR_USERNAME).First(&superAdmin).Error; err != nil {
			fmt.Println("Creating Administrator Account...")
			superAdmin = models.ApplicationUser{
				Id:             adminId,
				OwnerId:        adminId,
				IsActive:       &isActive,
				IsDefault:      &isDefault,
				IsLocked:       &isLocked,
				OrganizationId: organizationId,
				Password:       utils.HashAndSalt([]byte(DEFAULT_ADMINISTRATOR_PASSWORD)),
				FirstName:      "System",
				LastName:       "Administrator",
				Username:       DEFAULT_ADMINISTRATOR_PASSWORD,
				Address:        "",
				Phone:          "",
				IsSystemAdmin:  true,
				IsAdmin:        true,
				UserType:       1,
				CreatedBy:      adminId,
				CreatedAt:      &dateNow,
				Email:          DEFAULT_ADMINISTRATOR_EMAIL,
			}

			db.Session(&gorm.Session{SkipHooks: true}).Create(&superAdmin)
			fmt.Println("administrator account has ben created!")
			fmt.Println("Finished and Enjoy.")
			fmt.Println("================================== Completed... ================================")
			fmt.Println("================================================================================")
		}
	}
}

func dbMigration(db *gorm.DB) {
	db.AutoMigrate(
		&models.Organization{},
		&models.ApplicationUser{},
		&models.Team{},
		&models.Tenant{},
		&models.ServiceType{},
		&models.Uom{},
		&models.ProductCategory{},
		&models.Product{},
		&models.ProductDetail{},
	)

	viewList := make(map[string]map[string]string)
	viewList[views.Organization{}.TableName()] = views.Organization{}.Migration()
	viewList[views.ApplicationUser{}.TableName()] = views.ApplicationUser{}.Migration()
	viewList[views.Team{}.TableName()] = views.Team{}.Migration()
	viewList[views.Tenant{}.TableName()] = views.Tenant{}.Migration()
	viewList[views.Uom{}.TableName()] = views.Uom{}.Migration()
	viewList[views.ServiceType{}.TableName()] = views.ServiceType{}.Migration()
	viewList[views.ProductCategory{}.TableName()] = views.ProductCategory{}.Migration()
	viewList[views.Product{}.TableName()] = views.Product{}.Migration()
	viewList[views.ProductDetail{}.TableName()] = views.ProductDetail{}.Migration()

	if len(viewList) > 0 {
		for _, detail := range viewList {
			db.Exec(fmt.Sprintf("DROP VIEW IF EXISTS public.%v; CREATE OR REPLACE VIEW %v AS %s", detail["view_name"], detail["view_name"], detail["query"]))
		}
	}
}

type Config struct {
	DB_DRIVER     string `mapstructure:"DB_DRIVER"`
	DB_USERNAME   string `mapstructure:"DB_USERNAME"`
	DB_PASSWORD   string `mapstructure:"DB_PASSWORD"`
	DB_HOST       string `mapstructure:"DB_HOST"`
	DB_NAME       string `mapstructure:"DB_NAME"`
	DB_DIAL       string `mapstructure:"DB_DIAL"`
	DB_PORT       string `mapstructure:"DB_PORT"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

type InitConfig struct {
	DEFAULT_ORGANIZATION           string `mapstructure:"DEFAULT_ORGANIZATION"`
	DEFAULT_ADMINISTRATOR_USERNAME string `mapstructure:"DEFAULT_ADMINISTRATOR_USERNAME"`
	DEFAULT_ADMINISTRATOR_PASSWORD string `mapstructure:"DEFAULT_ADMINISTRATOR_PASSWORD"`
	DEFAULT_ADMINISTRATOR_EMAIL    string `mapstructure:"DEFAULT_ADMINISTRATOR_EMAIL"`
}
