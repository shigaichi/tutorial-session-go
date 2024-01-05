package orm

import (
	"time"

	"github.com/cockroachdb/errors"
	"github.com/shigaichi/tutorial-session-go/internal/domain/model"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&model.Account{}, &model.Category{}, &model.Goods{}, &model.Order{}, &model.OrderLine{})
	if err != nil {
		return errors.Wrap(err, "failed to migrate tables")
	}

	account := model.Account{
		ID:                 "488d6755-9ad9-4475-b652-18651b9a15a9",
		Email:              "a@example.com",
		EncodedPassword:    "$2a$12$/4/49055N0aVM5Innm7aSOhREYqxPcpWbt5mZJFV.iuYOJOxujgBq",
		Name:               "xxx",
		Birthday:           time.Date(2015, 8, 1, 0, 0, 0, 0, time.UTC),
		Zip:                "1111111",
		Address:            "Tokyo",
		CardNumber:         "1111111111111111",
		CardExpirationDate: time.Date(2015, 8, 1, 0, 0, 0, 0, time.UTC),
		CardSecurityCode:   "111",
	}

	result := db.Save(&account)
	if result.Error != nil {
		return errors.Wrap(result.Error, "failed to seed accounts")
	}

	categories := []*model.Category{
		{CategoryID: 1, CategoryName: "book"},
		{CategoryID: 2, CategoryName: "music"},
		{CategoryID: 3, CategoryName: "appliance"},
		{CategoryID: 4, CategoryName: "PC"},
	}
	result = db.Save(&categories)
	if result.Error != nil {
		return errors.Wrap(result.Error, "failed to seed categories")
	}

	goods := []*model.Goods{
		// 本
		{GoodsID: "366cf3a4-68c5-4dae-a557-673769f76840", GoodsName: "Kokoro", Description: "Souseki Natsume wrote this book", CategoryID: 1, Price: 900},
		{GoodsID: "366cf3a4-68c5-4dae-a557-673769f76841", GoodsName: "〔Ame ni mo Makezu〕", Description: "Kenji Miyazawa worte this book", CategoryID: 1, Price: 800},
		{GoodsID: "366cf3a4-68c5-4dae-a557-673769f76842", GoodsName: "Run, Melos!", Description: "Osamu Dazai wrote this book", CategoryID: 1, Price: 880},
		{GoodsID: "366cf3a4-68c5-4dae-a557-673769f76843", GoodsName: "I am a cat", Description: "Souseki Natsume wrote this book", CategoryID: 1, Price: 900},
		{GoodsID: "366cf3a4-68c5-4dae-a557-673769f76844", GoodsName: "No Longer Human", Description: "Osamu Dazai wrote this book", CategoryID: 1, Price: 880},
		// 音楽
		{GoodsID: "366cf3a4-68c5-4dae-a557-673769f76850", GoodsName: "Symphony No. 5 in C minor (Fate)", Description: "Beethoven composed this music", CategoryID: 2, Price: 1200},
		{GoodsID: "366cf3a4-68c5-4dae-a557-673769f76851", GoodsName: "Eine kleine Nachtmusik", Description: "Mozart composed this music", CategoryID: 2, Price: 1000},
		{GoodsID: "366cf3a4-68c5-4dae-a557-673769f76852", GoodsName: "Swan Lake", Description: "Tchaikovsky composed this music", CategoryID: 2, Price: 900},
		{GoodsID: "366cf3a4-68c5-4dae-a557-673769f76853", GoodsName: "Nocturnes", Description: "Chopin composed this music", CategoryID: 2, Price: 1000},
		{GoodsID: "366cf3a4-68c5-4dae-a557-673769f76854", GoodsName: "Air on the G String", Description: "Bach composed this music", CategoryID: 2, Price: 800},
		// 家電
		{GoodsID: "366cf3a4-68c5-4dae-a557-673769f76860", GoodsName: "Refrigerator", Description: "This refrigerator has a small power consumption", CategoryID: 3, Price: 108000},
		{GoodsID: "366cf3a4-68c5-4dae-a557-673769f76861", GoodsName: "Washing machine", Description: "This washing machine remove any stains", CategoryID: 3, Price: 216000},
		{GoodsID: "366cf3a4-68c5-4dae-a557-673769f76862", GoodsName: "Microwave", Description: "This microwave has over 10 options for cooking", CategoryID: 3, Price: 10800},
		{GoodsID: "366cf3a4-68c5-4dae-a557-673769f76863", GoodsName: "TV", Description: "This TV has a large screen and high image quality", CategoryID: 3, Price: 10800},
		{GoodsID: "366cf3a4-68c5-4dae-a557-673769f76864", GoodsName: "Digital Camera", Description: "This camera has shake correction device", CategoryID: 3, Price: 54000},
		// パソコン
		{GoodsID: "366cf3a4-68c5-4dae-a557-673769f76870", GoodsName: "Desktop PC", Description: "This Desktop PC has the latest CPU", CategoryID: 4, Price: 216000},
		{GoodsID: "366cf3a4-68c5-4dae-a557-673769f76871", GoodsName: "Laptop PC", Description: "This Laptop PC is convenient to carry", CategoryID: 4, Price: 162000},
		{GoodsID: "366cf3a4-68c5-4dae-a557-673769f76872", GoodsName: "Tablet", Description: "This tablet has an option of burglary insurance", CategoryID: 4, Price: 108000},
		{GoodsID: "366cf3a4-68c5-4dae-a557-673769f76873", GoodsName: "Integrated PC", Description: "This integreated PC has large screen and high performance", CategoryID: 4, Price: 183600},

		// ページネーション動作確認のための追加商品
		{GoodsID: "466cf3a4-68c5-4dae-a557-673769f76840", GoodsName: "Kokoro", Description: "Souseki Natsume wrote this book", CategoryID: 1, Price: 900},
		{GoodsID: "466cf3a4-68c5-4dae-a557-673769f76841", GoodsName: "〔Ame ni mo Makezu〕", Description: "Kenji Miyazawa worte this book", CategoryID: 1, Price: 800},
		{GoodsID: "466cf3a4-68c5-4dae-a557-673769f76842", GoodsName: "Run, Melos!", Description: "Osamu Dazai wrote this book", CategoryID: 1, Price: 880},
		{GoodsID: "466cf3a4-68c5-4dae-a557-673769f76843", GoodsName: "I am a cat", Description: "Souseki Natsume wrote this book", CategoryID: 1, Price: 900},
		{GoodsID: "466cf3a4-68c5-4dae-a557-673769f76844", GoodsName: "No Longer Human", Description: "Osamu Dazai wrote this book", CategoryID: 1, Price: 880},
		{GoodsID: "566cf3a4-68c5-4dae-a557-673769f76840", GoodsName: "Kokoro", Description: "Souseki Natsume wrote this book", CategoryID: 1, Price: 900},
		{GoodsID: "566cf3a4-68c5-4dae-a557-673769f76841", GoodsName: "〔Ame ni mo Makezu〕", Description: "Kenji Miyazawa worte this book", CategoryID: 1, Price: 800},
		{GoodsID: "566cf3a4-68c5-4dae-a557-673769f76842", GoodsName: "Run, Melos!", Description: "Osamu Dazai wrote this book", CategoryID: 1, Price: 880},
		{GoodsID: "566cf3a4-68c5-4dae-a557-673769f76843", GoodsName: "I am a cat", Description: "Souseki Natsume wrote this book", CategoryID: 1, Price: 900},
		{GoodsID: "566cf3a4-68c5-4dae-a557-673769f76844", GoodsName: "No Longer Human", Description: "Osamu Dazai wrote this book", CategoryID: 1, Price: 880},
		{GoodsID: "666cf3a4-68c5-4dae-a557-673769f76840", GoodsName: "Kokoro", Description: "Souseki Natsume wrote this book", CategoryID: 1, Price: 900},
		{GoodsID: "666cf3a4-68c5-4dae-a557-673769f76841", GoodsName: "〔Ame ni mo Makezu〕", Description: "Kenji Miyazawa worte this book", CategoryID: 1, Price: 800},
		{GoodsID: "666cf3a4-68c5-4dae-a557-673769f76842", GoodsName: "Run, Melos!", Description: "Osamu Dazai wrote this book", CategoryID: 1, Price: 880},
		{GoodsID: "666cf3a4-68c5-4dae-a557-673769f76843", GoodsName: "I am a cat", Description: "Souseki Natsume wrote this book", CategoryID: 1, Price: 900},
		{GoodsID: "666cf3a4-68c5-4dae-a557-673769f76844", GoodsName: "No Longer Human", Description: "Osamu Dazai wrote this book", CategoryID: 1, Price: 880},
		{GoodsID: "766cf3a4-68c5-4dae-a557-673769f76840", GoodsName: "Kokoro", Description: "Souseki Natsume wrote this book", CategoryID: 1, Price: 900},
		{GoodsID: "766cf3a4-68c5-4dae-a557-673769f76841", GoodsName: "〔Ame ni mo Makezu〕", Description: "Kenji Miyazawa worte this book", CategoryID: 1, Price: 800},
		{GoodsID: "766cf3a4-68c5-4dae-a557-673769f76842", GoodsName: "Run, Melos!", Description: "Osamu Dazai wrote this book", CategoryID: 1, Price: 880},
		{GoodsID: "766cf3a4-68c5-4dae-a557-673769f76843", GoodsName: "I am a cat", Description: "Souseki Natsume wrote this book", CategoryID: 1, Price: 900},
		{GoodsID: "766cf3a4-68c5-4dae-a557-673769f76844", GoodsName: "No Longer Human", Description: "Osamu Dazai wrote this book", CategoryID: 1, Price: 880},
		{GoodsID: "866cf3a4-68c5-4dae-a557-673769f76840", GoodsName: "Kokoro", Description: "Souseki Natsume wrote this book", CategoryID: 1, Price: 900},
		{GoodsID: "866cf3a4-68c5-4dae-a557-673769f76841", GoodsName: "〔Ame ni mo Makezu〕", Description: "Kenji Miyazawa worte this book", CategoryID: 1, Price: 800},
		{GoodsID: "866cf3a4-68c5-4dae-a557-673769f76842", GoodsName: "Run, Melos!", Description: "Osamu Dazai wrote this book", CategoryID: 1, Price: 880},
		{GoodsID: "866cf3a4-68c5-4dae-a557-673769f76843", GoodsName: "I am a cat", Description: "Souseki Natsume wrote this book", CategoryID: 1, Price: 900},
		{GoodsID: "866cf3a4-68c5-4dae-a557-673769f76844", GoodsName: "No Longer Human", Description: "Osamu Dazai wrote this book", CategoryID: 1, Price: 880},
		{GoodsID: "966cf3a4-68c5-4dae-a557-673769f76840", GoodsName: "Kokoro", Description: "Souseki Natsume wrote this book", CategoryID: 1, Price: 900},
		{GoodsID: "966cf3a4-68c5-4dae-a557-673769f76841", GoodsName: "〔Ame ni mo Makezu〕", Description: "Kenji Miyazawa worte this book", CategoryID: 1, Price: 800},
		{GoodsID: "966cf3a4-68c5-4dae-a557-673769f76842", GoodsName: "Run, Melos!", Description: "Osamu Dazai wrote this book", CategoryID: 1, Price: 880},
		{GoodsID: "966cf3a4-68c5-4dae-a557-673769f76843", GoodsName: "I am a cat", Description: "Souseki Natsume wrote this book", CategoryID: 1, Price: 900},
		{GoodsID: "966cf3a4-68c5-4dae-a557-673769f76844", GoodsName: "No Longer Human", Description: "Osamu Dazai wrote this book", CategoryID: 1, Price: 880},
	}

	result = db.Save(&goods)
	if result.Error != nil {
		return errors.Wrap(result.Error, "failed to seed goods")
	}

	return nil
}
