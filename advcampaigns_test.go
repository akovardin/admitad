package admitad

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/suite"
)

type AdvCampaignsSuite struct {
	suite.Suite
	server *httptest.Server
}

func (s *AdvCampaignsSuite) SetupSuite() {
	s.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.RequestURI {
		case "/advcampaigns/?":
			fmt.Fprint(w, advCampaignsTestData)
		case "/advcampaigns/website/1/?":
			fmt.Fprint(w, advCampaignsByWebsiteTestData)
		case "/advcampaigns/1/?":
			fmt.Fprint(w, advCampaignTestData)
		}
	}))
}

func (s *AdvCampaignsSuite) TearDownSuite() {
	if s.server != nil {
		s.server.Close()
	}
}

func (s *AdvCampaignsSuite) TestAdvCampaigns() {
	client := NewClient(
		s.server.URL,
		"xxx",
		"168d7bd617162c4ff00ae33e17eb5e",
		[]string{"advcampaigns", "banners", "websites"},
	)
	client.Init(&Token{})

	type Campaigns struct {
		Results []map[string]interface{} `json:"results"`
	}

	campaigns := Campaigns{}

	err := client.Call("advcampaigns", "GET", url.Values{}, &campaigns)
	s.Assertions.NoError(err)
	s.Assertions.Len(campaigns.Results, 2)
}

func (s *AdvCampaignsSuite) TestAdvCampaignsByWebsite() {
	client := NewClient(
		s.server.URL,
		"xxx",
		"168d7bd617162c4ff00ae33e17eb5e",
		[]string{"advcampaigns", "banners", "websites"},
	)
	client.Init(&Token{})

	type Campaigns struct {
		Results []map[string]interface{} `json:"results"`
	}

	campaigns := Campaigns{}

	err := client.Call("advcampaigns/website/1", "GET", url.Values{}, &campaigns)
	s.Assertions.NoError(err)
	s.Assertions.Len(campaigns.Results, 2)
}

func (s *AdvCampaignsSuite) TestAdvCampaign() {
	client := NewClient(
		s.server.URL,
		"xxx",
		"168d7bd617162c4ff00ae33e17eb5e",
		[]string{"advcampaigns", "banners", "websites"},
	)
	client.Init(&Token{})

	type Campaign struct {
		Id int `json:"id"`
	}

	campaign := Campaign{}

	err := client.Call("advcampaigns/1", "GET", url.Values{}, &campaign)
	s.Assertions.NoError(err)
	s.Assertions.EqualValues(3063, campaign.Id)
}

func TestAdvCampaignsSuite(t *testing.T) {
	suite.Run(t, new(AdvCampaignsSuite))
}

const (
	advCampaignsTestData = `{
    "results": [
        {
            "goto_cookie_lifetime": 60,
            "rating": "2.6",
            "rate_of_approve": "98",
            "more_rules": "",
            "exclusive": false,
            "image": "http://cdn.admitad.com/campaign/images/2011/03/05/8c855a8197e3f5b0067734c72efffb86.jpg",
            "actions": [
                {
                    "hold_time": 0,
                    "payment_size": "80RUB",
                    "type": "sale",
                    "name": "Эффективная регистрация",
                    "id": 1617
                },
                {
                    "hold_time": 0,
                    "payment_size": "0RUB",
                    "type": "lead",
                    "name": "Регистрация",
                    "id": 524
                }
            ],
            "avg_money_transfer_time": 14,
            "currency": "RUB",
            "activation_date": "2015-11-09T14:56:16",
            "retag": false,
            "cr": 2.51,
            "ecpc": 0.73,
            "id": 92,
            "description": "ЗАПРЕЩЕНА&nbsp;КОНТЕКСТНАЯ РЕКЛАМА НА БРЕНД!\r\n\r\nПодключайтесь к новой партнерской программе -&nbsp;&quot;Lineage 2 RU&quot;&nbsp;\r\nLineage 2&nbsp;&mdash; это не&nbsp;просто онлайн-игра с&nbsp;реалистичной графикой и&nbsp;продуманной системой сражений. Это целый мир, в&nbsp;котором развитие персонажа ничем не&nbsp;ограничено. Тебя ждут зрелищные, масштабные PvP-битвы и&nbsp;самые посещаемые серверы, каждый с&nbsp;онлайном до&nbsp;5500 игроков одновременно.\r\n\r\nЧто привлекает пользователей в игре&nbsp;&quot;Lineage 2 RU&quot;\r\n\r\nLineage 2 &mdash; легендарная онлайн-игра, которая безостановочно развивается вот уже на протяжении 10 лет. Это населенный фэнтезийными расами онлайн-мир, в котором есть место дружбе и предательству, масштабным сражениям кланов и дуэльным схваткам игроков. Настоящая арена для миллионов людей.\r\n\r\nСоздавая своего персонажа, пользователь может выбрать из 37 уникальных классов.\r\n\r\nКоличество зарегистрированных пользователей Lineage 2 превышает 6 миллиона человек. Возраст игроков от 13 до 70 лет.\r\n\r\n&nbsp;Список минус - слов по контекстной рекламе:\r\n\r\n\r\nС Уважением к Вам, admitad afilliate team!",
            "modified_date": "2018-02-13T18:39:41",
            "cr_trend": "0.0000",
            "site_url": "https://ru.4game.com/lineage2/install",
            "regions": [
                {
                    "region": "AM"
                },
                {
                    "region": "AZ"
                },
                {
                    "region": "BY"
                },
                {
                    "region": "GE"
                },
                {
                    "region": "KG"
                },
                {
                    "region": "KZ"
                },
                {
                    "region": "MD"
                },
                {
                    "region": "RU"
                },
                {
                    "region": "TJ"
                },
                {
                    "region": "TM"
                },
                {
                    "region": "UA"
                },
                {
                    "region": "UZ"
                }
            ],
            "epc_trend": "0.0000",
            "geotargeting": false,
            "status": "active",
            "coupon_iframe_denied": false,
            "traffics": [
                {
                    "enabled": false,
                    "name": "Cashback",
                    "id": 1
                },
                {
                    "enabled": false,
                    "name": "PopUp / ClickUnder",
                    "id": 2
                },
                {
                    "enabled": true,
                    "name": "Контекстная реклама",
                    "id": 3
                },
                {
                    "enabled": true,
                    "name": "Дорвей - трафик",
                    "id": 4
                },
                {
                    "enabled": true,
                    "name": "Email - рассылка",
                    "id": 5
                },
                {
                    "enabled": false,
                    "name": "Контекстная реклама на Бренд",
                    "id": 6
                },
                {
                    "enabled": true,
                    "name": "Реклама в социальных сетях",
                    "id": 7
                },
                {
                    "enabled": false,
                    "name": "Мотивированный трафик",
                    "id": 8
                },
                {
                    "enabled": true,
                    "name": "Toolbar",
                    "id": 9
                },
                {
                    "enabled": false,
                    "name": "Adult - трафик",
                    "id": 14
                },
                {
                    "enabled": true,
                    "name": "Тизерные сети",
                    "id": 18
                },
                {
                    "enabled": true,
                    "name": "Youtube Канал",
                    "id": 19
                }
            ],
            "individual_terms": false,
            "avg_hold_time": 14,
            "raw_description": "<p><span style=\"color:#FF0000\"><strong>ЗАПРЕЩЕНА&nbsp;КОНТЕКСТНАЯ РЕКЛАМА НА БРЕНД!</strong></span></p>\r\n\r\n<p><span style=\"color:#FF0000\"><strong><span style=\"font-size:16px\">Подключайтесь к новой партнерской программе -&nbsp;&quot;Lineage 2 RU&quot;&nbsp;</span></strong></span><br />\r\nLineage 2&nbsp;&mdash; это не&nbsp;просто онлайн-игра с&nbsp;реалистичной графикой и&nbsp;продуманной системой сражений. Это целый мир, в&nbsp;котором развитие персонажа ничем не&nbsp;ограничено. Тебя ждут зрелищные, масштабные PvP-битвы и&nbsp;самые посещаемые серверы, каждый с&nbsp;онлайном до&nbsp;5500 игроков одновременно.</p>\r\n\r\n<p><span style=\"color:#FF0000\"><span style=\"font-size:16px\"><strong>Что привлекает пользователей в игре&nbsp;</strong></span><strong><span style=\"font-size:16px\">&quot;</span></strong><strong><span style=\"font-size:16px\">Lineage 2 RU</span></strong><strong><span style=\"font-size:16px\">&quot;</span></strong></span></p>\r\n\r\n<p>Lineage 2 &mdash; легендарная онлайн-игра, которая безостановочно развивается вот уже на протяжении 10 лет. Это населенный фэнтезийными расами онлайн-мир, в котором есть место дружбе и предательству, масштабным сражениям кланов и дуэльным схваткам игроков. Настоящая арена для миллионов людей.</p>\r\n\r\n<p>Создавая своего персонажа, пользователь может выбрать из 37 уникальных классов.</p>\r\n\r\n<p>Количество зарегистрированных пользователей Lineage 2 превышает 6 миллиона человек. Возраст игроков от 13 до 70 лет.</p>\r\n\r\n<p><a href=\"https://cdn.admitad-connect.com/public/storage/2017/12/20/Lineage_2_RU_Spisok_minus_slov.docx\">&nbsp;Список минус - слов по контекстной рекламе:</a></p>\r\n\r\n<p><br />\r\n<strong>С Уважением к Вам, admitad afilliate team!</strong></p>\r\n",
            "denynewwms": false,
            "connected": false,
            "max_hold_time": null,
            "categories": [
                {
                    "language": "ru",
                    "id": 6,
                    "parent": null,
                    "name": "Онлайн Игры"
                },
                {
                    "language": "ru",
                    "id": 15,
                    "parent": {
                        "language": "ru",
                        "id": 6,
                        "parent": null,
                        "name": "Онлайн Игры"
                    },
                    "name": "Клиентские"
                }
            ],
            "name_aliases": "lineage2, innova, 4game, 4гейм, Линяга",
            "name": "Lineage 2 [CPP] RU +12 countries",
            "landing_code": null,
            "ecpc_trend": "0.0000",
            "landing_title": null,
            "action_type": "sale",
            "epc": 73.0,
            "allow_deeplink": false,
            "show_products_links": false
        },
        {
            "goto_cookie_lifetime": 30,
            "rating": "2.7",
            "rate_of_approve": "56",
            "more_rules": "<p>1) Запрещено использовать<span style=\"background-color:rgba(255, 255, 255, 0.917969); font-size:13px\">&nbsp;в рекламе бренд компании Yves Rocher.</span></p>\r\n\r\n<p><strong>2) Напоминаем, запрещены следующие виды трафика</strong><span style=\"color:rgb(0, 0, 0); font-size:13px\">:&nbsp;</span><strong><span style=\"color:rgb(255, 0, 0)\">Контекстная реклама на Бренд, Контекстная реклама, PopUp реклама, ClickUnder реклама, приложения/игры в соц.сетях, Push - реклама, мотифированный трафик, Aduit - трафик, любые виды ретаргетинга!</span></strong></p>\r\n\r\n<p>Будьте внимательны, заказа от запрещенного вида трафика - будут отклоняться!</p>\r\n\r\n<p><span style=\"color:#ff0000\"><strong>3) ВНИМАНИЕ, разрешается!</strong></span><br />\r\n<span style=\"color:rgb(80, 0, 80); font-size:13.333333969116211px\">Рассылка (по согласованию)&nbsp;cash-</span><span style=\"color:rgb(80, 0, 80); font-size:13.333333969116211px\">back&nbsp;(только крупных, остальных по согласованию), мотивированный трафик (по согласованию)!</span></p>\r\n\r\n<p><u><strong><span style=\"color:rgb(255, 0, 0)\"><span style=\"font-size:12.8px\">4) Запрещен автоматический редирект с сайта вэб-мастера на сайт рекламодателя!</span></span></strong></u></p>\r\n\r\n<p><span style=\"color:#FF0000\"><strong><u>5) Отключение промокодов на некоторые бренды:</u></strong></span><br />\r\n<br />\r\nНиже&nbsp;black-лист. Технически промокода не будут действовать на следующие бренды:</p>\r\n\r\n<p>oodji, Deri&amp;Mod, Madeleine, Pandora, GILLETTE VENUS, SIMPLY VENUS, PAMPERS, ALWAYS, BLEND A MED, DISCREET, FAIRY, FUSION, GILLETTE, GILLETTE VENUS, MACH3, NATURELLA, OLD SPICE, ORAL B, SAFEGUARD, TAMPAX, TIDE, VENUS, ARIEL, AUSSIE, HEAD &amp; SHOULDERS, LENOR, MR. PROPER, MYTH, PANTENE, WELLA, AMBI PUR, DREFT, ORAL-B, SECRET</p>\r\n\r\n<p>&nbsp;</p>\r\n",
            "exclusive": false,
            "image": "http://cdn.admitad.com/campaign/images/2017/10/3/fffe92da99b37d46b356b84ea1a6b270.jpg",
            "actions": [
                {
                    "hold_time": 0,
                    "payment_size": "10%",
                    "type": "sale",
                    "name": "Оплаченные заказ маркетпплейс",
                    "id": 12392
                },
                {
                    "hold_time": 0,
                    "payment_size": "0.65%",
                    "type": "sale",
                    "name": "Оплаченный заказ купонный трафик",
                    "id": 12388
                },
                {
                    "hold_time": 0,
                    "payment_size": "3.25%",
                    "type": "sale",
                    "name": "Оплаченный заказ",
                    "id": 3923
                },
                {
                    "hold_time": 0,
                    "payment_size": "1.95%",
                    "type": "sale",
                    "name": "Оплаченный заказ cashback",
                    "id": 12387
                }
            ],
            "avg_money_transfer_time": 56,
            "currency": "RUB",
            "activation_date": "2011-11-23T00:11:08",
            "retag": false,
            "cr": 0.75,
            "ecpc": 1.62,
            "id": 153,
            "description": "&nbsp;\r\n\r\nОбновленный серверный трекинг - дополнительные выплаты сверх статистики admitad\r\n\r\nПартнерская программа KUPIVIP спешит сообщить, что KUPIVIP.RU совместно с admitad.com провел ряд мер по модернизации учета заказов.\r\n\r\nМы модернизировали серверный трекинг на сайте KUPIVIP.RU, благодаря чему мы фиксируем даже те заказы, которые по какой-то причине были пропущены трекингом admitad и не отображались в статистике CPA-сети.\r\n\r\nМы производим оплату этих заказов вебмастерам, благодаря чему, большинство вебмастеров получает дополнительные выплаты в конце месяца.\r\n\r\nВ среднем, доход вебмастеров по факту увеличивается на 10-35%, от дохода, отображаемого в статистике admitad.\r\n\r\nДля примера, величина доп.выплат для вебмастеров из ТОП-5 оффера:\r\n\r\n1 - 15%\r\n\r\n2 - 36%\r\n\r\n3 - 24%\r\n\r\n4 - 26%\r\n\r\n5 - 34%\r\n\r\nДополнительные выплаты начисляются через 10-15 дней после истечения календарного месяца и в этот же период менеджеры admitad вручную добавляют заказы в статистику личных кабинетов CPA-сети.\r\n\r\n&nbsp;\r\n\r\nЖелаем отличного заработка!\r\n\r\n&nbsp;\r\n\r\n\r\n\r\nKupiVip.ru &ndash; крупнейший онлайн-магазин, работающий на рынке России и стран СНГ.\r\n\r\nKupiVip.ru&nbsp;меняет стандарты покупки модных товаров. Это первый онлайн-аутлет, который предлагает скидки до 90% на престижные мировые бренды.\r\n\r\n&nbsp;\r\n\r\nО компании:\r\n\r\nБолее 100 акций распродаж со скидками до 90% ежедневно.\r\n\r\nБолее 2000 популярных и эксклюзивных брендов из десятка стран Европы, Азии, а также США.\r\n\r\nБолее 400&nbsp;000 товаров каждый день.\r\n\r\n6 офисов в Европе и Азии (Берлин, Милан, Париж, Москва, Санкт-Петербург, Алма-Ата).\r\n\r\nВходит в Топ-10 крупнейших онлайн-рекламодателей в России.\r\n\r\n&nbsp;\r\n\r\nИстория:\r\n\r\n2008 &ndash; Старт проекта\r\n\r\n2008-2012 &ndash; Рекордные для бизнеса онлайн-торговли инвестиции &ndash;$124 млн.&nbsp;от нескольких крупных мировых фондов.\r\n\r\n2012 &ndash; Запуск в Казахстане\r\n\r\n2012 &ndash; Запуск мобильного приложения\r\n\r\n2014 &ndash; Запуск в Беларуси\r\n\r\n2014 &ndash;&nbsp;EBITDA&nbsp;&ndash; безубыточность\r\n\r\n2014 &ndash;&nbsp;KupiVip.ru&nbsp;входит в топ-20 крупнейших онлайн-магазинов Рунета по версии&nbsp;Forbes\r\n\r\n2015 &ndash;&nbsp;KupiVip.ru&nbsp;входит в топ-50 самых быстрорастущих&nbsp;IT-проектов Европы по версии&nbsp;Tech&nbsp;Growth&nbsp;50\r\n\r\n&nbsp;\r\n\r\nАудитория:\r\n\r\nПокупатель&nbsp;KupiVip.ru&nbsp;&ndash; семейная, но самостоятельная женщина, которая любит модный шоппинг и разбирается в премиальных марках одежды, обуви и аксессуаров.\r\n\r\nГлавным мотиватором продажи выступает скидка и качество товара, который невозможно найти в другом интернет-магазине.\r\n\r\n&nbsp;\r\n\r\n90% Аудитории&nbsp;KupiVip.ru&nbsp;&ndash; женщины.\r\n\r\nПортрет покупательницы:\r\n\r\n30-45 лет\r\n\r\n60% покупок одежды совершает в Интернете\r\n\r\nДоход 50&nbsp;000 рублей +\r\n\r\nЛюбит скидки\r\n\r\nЖительница Москвы (45% покупателей), МО, Санкт-Петербурга, ЛО и других городов-миллионников.\r\n\r\n&nbsp;\r\n\r\nПреимущества для клиентов:\r\n\r\nДелая покупки в первом онлайн-аутлете&nbsp;KupiVip.ru, Ваши покупатели получают большое количество плюсов:\r\n\r\n- Постоянный ассортимент со скидками до 90% на ведущие мировые бренды;\r\n\r\n- Гарантии качества и подлинности товара;\r\n\r\n- Гарантия лучшей цены;\r\n\r\n- Примерка до оплаты заказа (Москва, МО, Санкт-Петербург);\r\n\r\n- Бесплатная доставка по всей России от 5999 рублей / Доставка по Москве 149 рублей;\r\n\r\n- 30 дней на возврат товара;\r\n\r\n- Клиентская служба 24 часа/7 дней в неделю;\r\n\r\n- Быстрая доставка (Москва &ndash; доставка на следующий день, СПб &ndash; 2-3 дня, регионы &ndash; 3-14 дней);\r\n\r\n- Широкий выбор способов оплаты (наличный расчет, пластиковые карты, Яндекс-деньги, Paypal, Qiwi).\r\n\r\n&nbsp;\r\n\r\nПреимущества работы с программой:\r\n\r\n- Высокая конверсия;\r\n\r\n- Широкий выбор рекламных форматов и акций (товарный фид, фид акций, Deep-link, баннеры, промо-коды);\r\n\r\n- Высокая маркетинговая активность, регулярные акции и промо-коды;\r\n\r\n- Широкая линейка разрешенных типов трафика;\r\n\r\n- Индивидуальные условия и бонусы для лучших вебмастеров;\r\n\r\n- Крупнейшая в России клиентская база лояльной аудитории к известному бренду.\r\n\r\nПо ссылке Вы сможете найти подробную информацию по средней стоимости товаров различных категорий на&nbsp;KUPIVIP, коэффиценты выкупаемости и остальная статистическая информация, которая может быть полезна.\r\n\r\nhttps://docs.google.com/spreadsheets/d/1fpINb_Tzi6zZsheASIyMva4z-YGaZJzfpGQYvOnCa_U/edit#gid=932367914\r\n\r\n&nbsp;\r\n\r\n&nbsp;",
            "modified_date": "2018-09-21T08:38:07",
            "cr_trend": "0.0000",
            "site_url": "http://kupivip.ru/",
            "regions": [
                {
                    "region": "RU"
                }
            ],
            "epc_trend": "0.0000",
            "geotargeting": false,
            "status": "active",
            "coupon_iframe_denied": false,
            "traffics": [
                {
                    "enabled": true,
                    "name": "Cashback",
                    "id": 1
                },
                {
                    "enabled": true,
                    "name": "PopUp / ClickUnder",
                    "id": 2
                },
                {
                    "enabled": false,
                    "name": "Контекстная реклама",
                    "id": 3
                },
                {
                    "enabled": false,
                    "name": "Дорвей - трафик",
                    "id": 4
                },
                {
                    "enabled": true,
                    "name": "Email - рассылка",
                    "id": 5
                },
                {
                    "enabled": false,
                    "name": "Контекстная реклама на Бренд",
                    "id": 6
                },
                {
                    "enabled": true,
                    "name": "Реклама в социальных сетях",
                    "id": 7
                },
                {
                    "enabled": true,
                    "name": "Мотивированный трафик",
                    "id": 8
                },
                {
                    "enabled": false,
                    "name": "Toolbar",
                    "id": 9
                },
                {
                    "enabled": false,
                    "name": "Adult - трафик",
                    "id": 14
                },
                {
                    "enabled": true,
                    "name": "Тизерные сети",
                    "id": 18
                },
                {
                    "enabled": true,
                    "name": "Youtube Канал",
                    "id": 19
                }
            ],
            "individual_terms": false,
            "avg_hold_time": 56,
            "raw_description": "<p style=\"text-align:center\">&nbsp;</p>\r\n\r\n<p><strong>Обновленный серверный трекинг - дополнительные выплаты сверх статистики admitad</strong></p>\r\n\r\n<p>Партнерская программа KUPIVIP спешит сообщить, что KUPIVIP.RU совместно с admitad.com провел ряд мер по модернизации учета заказов.</p>\r\n\r\n<p>Мы модернизировали серверный трекинг на сайте KUPIVIP.RU, благодаря чему мы фиксируем даже те заказы, которые по какой-то причине были пропущены трекингом admitad и не отображались в статистике CPA-сети.</p>\r\n\r\n<p>Мы производим оплату этих заказов вебмастерам, благодаря чему, большинство вебмастеров получает дополнительные выплаты в конце месяца.</p>\r\n\r\n<p>В среднем, доход вебмастеров по факту увеличивается на 10-35%, от дохода, отображаемого в статистике admitad.</p>\r\n\r\n<p>Для примера, величина доп.выплат для вебмастеров из ТОП-5 оффера:</p>\r\n\r\n<p>1 - 15%</p>\r\n\r\n<p>2 - 36%</p>\r\n\r\n<p>3 - 24%</p>\r\n\r\n<p>4 - 26%</p>\r\n\r\n<p>5 - 34%</p>\r\n\r\n<p>Дополнительные выплаты начисляются через 10-15 дней после истечения календарного месяца и в этот же период менеджеры admitad вручную добавляют заказы в статистику личных кабинетов CPA-сети.</p>\r\n\r\n<p>&nbsp;</p>\r\n\r\n<p>Желаем отличного заработка!</p>\r\n\r\n<p style=\"text-align:center\">&nbsp;</p>\r\n\r\n<p style=\"text-align:center\"><img alt=\"\" src=\"https://www.admitad.com/public/storage/2015/10/02/infografika_KupiVIP_RU.jpg\" style=\"font-size:12px; height:1637px; width:800px\" /></p>\r\n\r\n<p style=\"text-align:center\"><span style=\"color:#000000\"><strong>KupiVip.ru &ndash; крупнейший онлайн-магазин, работающий на рынке России и стран СНГ.</strong></span></p>\r\n\r\n<p><span style=\"color:#000000\"><strong>KupiVip.</strong><strong>ru</strong><strong>&nbsp;меняет стандарты покупки модных товаров. Это первый онлайн-аутлет, который предлагает скидки до 90% на престижные мировые бренды.</strong></span></p>\r\n\r\n<p>&nbsp;</p>\r\n\r\n<p><span style=\"color:#000000\"><strong>О компании:</strong></span></p>\r\n\r\n<p><span style=\"color:#000000\">Более 100 акций распродаж со скидками до 90% ежедневно.</span></p>\r\n\r\n<p><span style=\"color:#000000\">Более 2000 популярных и эксклюзивных брендов из десятка стран Европы, Азии, а также США.</span></p>\r\n\r\n<p><span style=\"color:#000000\">Более 400&nbsp;000 товаров каждый день.</span></p>\r\n\r\n<p><span style=\"color:#000000\">6 офисов в Европе и Азии (Берлин, Милан, Париж, Москва, Санкт-Петербург, Алма-Ата).</span></p>\r\n\r\n<p><span style=\"color:#000000\">Входит в Топ-10 крупнейших онлайн-рекламодателей в России.</span></p>\r\n\r\n<p>&nbsp;</p>\r\n\r\n<p><span style=\"color:#000000\"><strong>История:</strong></span></p>\r\n\r\n<p><span style=\"color:#000000\">2008 &ndash; Старт проекта</span></p>\r\n\r\n<p><span style=\"color:#000000\">2008-2012 &ndash; Рекордные для бизнеса онлайн-торговли инвестиции &ndash;$124 млн.&nbsp;от нескольких крупных мировых фондов.</span></p>\r\n\r\n<p><span style=\"color:#000000\">2012 &ndash; Запуск в Казахстане</span></p>\r\n\r\n<p><span style=\"color:#000000\">2012 &ndash; Запуск мобильного приложения</span></p>\r\n\r\n<p><span style=\"color:#000000\">2014 &ndash; Запуск в Беларуси</span></p>\r\n\r\n<p><span style=\"color:#000000\">2014 &ndash;&nbsp;EBITDA&nbsp;&ndash; безубыточность</span></p>\r\n\r\n<p><span style=\"color:#000000\">2014 &ndash;&nbsp;KupiVip.ru&nbsp;входит в топ-20 крупнейших онлайн-магазинов Рунета по версии&nbsp;Forbes</span></p>\r\n\r\n<p><span style=\"color:#000000\">2015 &ndash;&nbsp;KupiVip.ru&nbsp;входит в топ-50 самых быстрорастущих&nbsp;IT-проектов Европы по версии&nbsp;Tech&nbsp;Growth&nbsp;50</span></p>\r\n\r\n<p>&nbsp;</p>\r\n\r\n<p><span style=\"color:#000000\"><strong>Аудитория:</strong></span></p>\r\n\r\n<p><span style=\"color:#000000\">Покупатель&nbsp;KupiVip.ru&nbsp;&ndash; семейная, но самостоятельная женщина, которая любит модный шоппинг и разбирается в премиальных марках одежды, обуви и аксессуаров.</span></p>\r\n\r\n<p><span style=\"color:#000000\">Главным мотиватором продажи выступает скидка и качество товара, который невозможно найти в другом интернет-магазине.</span></p>\r\n\r\n<p>&nbsp;</p>\r\n\r\n<p><span style=\"color:#000000\">90% Аудитории&nbsp;KupiVip.ru&nbsp;&ndash; женщины.</span></p>\r\n\r\n<p><span style=\"color:#000000\">Портрет покупательницы:</span></p>\r\n\r\n<p><span style=\"color:#000000\">30-45 лет</span></p>\r\n\r\n<p><span style=\"color:#000000\">60% покупок одежды совершает в Интернете</span></p>\r\n\r\n<p><span style=\"color:#000000\">Доход 50&nbsp;000 рублей +</span></p>\r\n\r\n<p><span style=\"color:#000000\">Любит скидки</span></p>\r\n\r\n<p><span style=\"color:#000000\">Жительница Москвы (45% покупателей), МО, Санкт-Петербурга, ЛО и других городов-миллионников.</span></p>\r\n\r\n<p>&nbsp;</p>\r\n\r\n<p><span style=\"color:#000000\"><strong>Преимущества для клиентов:</strong></span></p>\r\n\r\n<p><span style=\"color:#000000\"><strong>Делая покупки в первом онлайн-аутлете</strong>&nbsp;<strong>KupiVip.ru, Ваши покупатели получают большое количество плюсов:</strong></span></p>\r\n\r\n<p><span style=\"color:#000000\">- Постоянный ассортимент со скидками до 90% на ведущие мировые бренды;</span></p>\r\n\r\n<p><span style=\"color:#000000\">- Гарантии качества и подлинности товара;</span></p>\r\n\r\n<p><span style=\"color:#000000\">- Гарантия лучшей цены;</span></p>\r\n\r\n<p><span style=\"color:#000000\">- Примерка до оплаты заказа (Москва, МО, Санкт-Петербург);</span></p>\r\n\r\n<p><span style=\"color:#000000\">- Бесплатная доставка по всей России от 5999 рублей / Доставка по Москве 149 рублей;</span></p>\r\n\r\n<p><span style=\"color:#000000\">- 30 дней на возврат товара;</span></p>\r\n\r\n<p><span style=\"color:#000000\">- Клиентская служба 24 часа/7 дней в неделю;</span></p>\r\n\r\n<p><span style=\"color:#000000\">- Быстрая доставка (Москва &ndash; доставка на следующий день, СПб &ndash; 2-3 дня, регионы &ndash; 3-14 дней);</span></p>\r\n\r\n<p><span style=\"color:#000000\">- Широкий выбор способов оплаты (наличный расчет, пластиковые карты, Яндекс-деньги, Paypal, Qiwi).</span></p>\r\n\r\n<p>&nbsp;</p>\r\n\r\n<p><span style=\"color:#000000\"><strong>Преимущества работы с программой:</strong></span></p>\r\n\r\n<p><span style=\"color:#000000\">- Высокая конверсия;</span></p>\r\n\r\n<p><span style=\"color:#000000\">- Широкий выбор рекламных форматов и акций (товарный фид, фид акций, Deep-link, баннеры, промо-коды);</span></p>\r\n\r\n<p><span style=\"color:#000000\">- Высокая маркетинговая активность, регулярные акции и промо-коды;</span></p>\r\n\r\n<p><span style=\"color:#000000\">- Широкая линейка разрешенных типов трафика;</span></p>\r\n\r\n<p><span style=\"color:#000000\">- Индивидуальные условия и бонусы для лучших вебмастеров;</span></p>\r\n\r\n<p><span style=\"color:#000000\">- Крупнейшая в России клиентская база лояльной аудитории к известному бренду.</span></p>\r\n\r\n<p>По ссылке Вы сможете найти подробную информацию по средней стоимости товаров различных категорий на&nbsp;KUPIVIP, коэффиценты выкупаемости и остальная статистическая информация, которая может быть полезна.</p>\r\n\r\n<p><span style=\"color:#000000\">https://docs.google.com/spreadsheets/d/1fpINb_Tzi6zZsheASIyMva4z-YGaZJzfpGQYvOnCa_U/edit#gid=932367914</span></p>\r\n\r\n<p>&nbsp;</p>\r\n\r\n<p>&nbsp;</p>\r\n",
            "denynewwms": false,
            "connected": false,
            "max_hold_time": null,
            "categories": [
                {
                    "language": "ru",
                    "id": 62,
                    "parent": null,
                    "name": "Интернет-магазины"
                },
                {
                    "language": "ru",
                    "id": 64,
                    "parent": {
                        "language": "ru",
                        "id": 62,
                        "parent": null,
                        "name": "Интернет-магазины"
                    },
                    "name": "Одежда & Обувь"
                },
                {
                    "language": "ru",
                    "id": 69,
                    "parent": {
                        "language": "ru",
                        "id": 62,
                        "parent": null,
                        "name": "Интернет-магазины"
                    },
                    "name": "Товары для детей"
                },
                {
                    "language": "ru",
                    "id": 71,
                    "parent": {
                        "language": "ru",
                        "id": 62,
                        "parent": null,
                        "name": "Интернет-магазины"
                    },
                    "name": "Аксессуары"
                },
                {
                    "language": "ru",
                    "id": 85,
                    "parent": {
                        "language": "ru",
                        "id": 62,
                        "parent": null,
                        "name": "Интернет-магазины"
                    },
                    "name": "Спорт"
                }
            ],
            "name_aliases": "KupiVIP, kupivipru, kupivip.ru, купивип, купи вип",
            "name": "Kupivip RU",
            "landing_code": null,
            "ecpc_trend": "0.0000",
            "landing_title": null,
            "action_type": "sale",
            "epc": 162.0,
            "allow_deeplink": true,
            "show_products_links": true
        }
    ],
    "_meta": {
        "count": 1665,
        "limit": 2,
        "offset": 0
    }
}`
	advCampaignsByWebsiteTestData = `{
    "results": [
        {
            "goto_cookie_lifetime": 30,
            "rating": "2.3",
            "exclusive": false,
            "image": "http://cdn.admitad.com/campaign/images/2011/12/12/8f6848b102dd82b7003e2b5957603d2c.jpg",
            "actions": [
                {
                    "hold_time": 0,
                    "payment_size": "5%",
                    "type": "sale",
                    "name": "Покупка: Электронные книги издательства Gardners (Новый пользователь)",
                    "id": 944
                },
                {
                    "hold_time": 0,
                    "payment_size": "10%",
                    "type": "sale",
                    "name": "Оплаченный заказ (Активный пользователь)",
                    "id": 11149
                },
                {
                    "hold_time": 0,
                    "payment_size": "3%",
                    "type": "sale",
                    "name": "Покупка: Электронные книги издательства Gardners (Активный пользователь)",
                    "id": 11150
                },
                {
                    "hold_time": 0,
                    "payment_size": "20%",
                    "type": "sale",
                    "name": "Оплата заказа (Новый пользователь)",
                    "id": 926
                }
            ],
            "avg_money_transfer_time": 30,
            "currency": "RUB",
            "activation_date": "2011-12-19T17:31:45",
            "retag": true,
            "cr": 2.35,
            "max_hold_time": null,
            "id": 1721,
            "avg_hold_time": 29,
            "ecpc": 0.71,
            "connection_status": "active",
            "gotolink": "https://ad.admitad.com/g/2d388421f4b72ea2020f3baa9723ff/",
            "site_url": "http://litres.ru/",
            "regions": [
                {
                    "region": "AM"
                },
                {
                    "region": "AZ"
                },
                {
                    "region": "BY"
                },
                {
                    "region": "KZ"
                },
                {
                    "region": "RU"
                },
                {
                    "region": "TJ"
                },
                {
                    "region": "TM"
                },
                {
                    "region": "UA"
                },
                {
                    "region": "UZ"
                }
            ],
            "actions_detail": [
                {
                    "hold_size": 0,
                    "tariffs": [
                        {
                            "action_id": 944,
                            "rates": [
                                {
                                    "price_s": "0.00",
                                    "tariff_id": 944,
                                    "country": null,
                                    "date_s": "2018-09-11",
                                    "is_percentage": true,
                                    "id": 152054,
                                    "size": "5.00"
                                }
                            ],
                            "id": 944,
                            "name": "Покупка: Электронные книги издательства Gardners - (Новый пользователь)"
                        }
                    ],
                    "type": "sale",
                    "name": "Покупка: Электронные книги издательства Gardners (Новый пользователь)",
                    "id": 944
                },
                {
                    "hold_size": 0,
                    "tariffs": [
                        {
                            "action_id": 11149,
                            "rates": [
                                {
                                    "price_s": "0.00",
                                    "tariff_id": 14001,
                                    "country": null,
                                    "date_s": "2018-09-11",
                                    "is_percentage": true,
                                    "id": 152029,
                                    "size": "10.00"
                                }
                            ],
                            "id": 14001,
                            "name": "Тариф по умолчанию"
                        }
                    ],
                    "type": "sale",
                    "name": "Оплаченный заказ (Активный пользователь)",
                    "id": 11149
                },
                {
                    "hold_size": 0,
                    "tariffs": [
                        {
                            "action_id": 11150,
                            "rates": [
                                {
                                    "price_s": "0.00",
                                    "tariff_id": 14002,
                                    "country": null,
                                    "date_s": "2018-09-11",
                                    "is_percentage": true,
                                    "id": 152040,
                                    "size": "3.00"
                                }
                            ],
                            "id": 14002,
                            "name": "Тариф по умолчанию"
                        }
                    ],
                    "type": "sale",
                    "name": "Покупка: Электронные книги издательства Gardners (Активный пользователь)",
                    "id": 11150
                },
                {
                    "hold_size": 0,
                    "tariffs": [
                        {
                            "action_id": 926,
                            "rates": [
                                {
                                    "price_s": "0.00",
                                    "tariff_id": 926,
                                    "country": null,
                                    "date_s": "2018-09-11",
                                    "is_percentage": true,
                                    "id": 152018,
                                    "size": "20.00"
                                }
                            ],
                            "id": 926,
                            "name": "Оплата заказа - default"
                        }
                    ],
                    "type": "sale",
                    "name": "Оплата заказа (Новый пользователь)",
                    "id": 926
                }
            ],
            "landing_title": null,
            "geotargeting": false,
            "status": "active",
            "coupon_iframe_denied": false,
            "traffics": [
                {
                    "enabled": true,
                    "name": "Cashback",
                    "id": 1
                },
                {
                    "enabled": true,
                    "name": "PopUp / ClickUnder",
                    "id": 2
                },
                {
                    "enabled": false,
                    "name": "Контекстная реклама",
                    "id": 3
                },
                {
                    "enabled": false,
                    "name": "Дорвей - трафик",
                    "id": 4
                },
                {
                    "enabled": true,
                    "name": "Email - рассылка",
                    "id": 5
                },
                {
                    "enabled": false,
                    "name": "Контекстная реклама на Бренд",
                    "id": 6
                },
                {
                    "enabled": true,
                    "name": "Реклама в социальных сетях",
                    "id": 7
                },
                {
                    "enabled": true,
                    "name": "Мотивированный трафик",
                    "id": 8
                },
                {
                    "enabled": false,
                    "name": "Toolbar",
                    "id": 9
                },
                {
                    "enabled": true,
                    "name": "Adult - трафик",
                    "id": 14
                },
                {
                    "enabled": true,
                    "name": "Тизерные сети",
                    "id": 18
                },
                {
                    "enabled": true,
                    "name": "Youtube Канал",
                    "id": 19
                },
                {
                    "enabled": false,
                    "name": "Брокерский трафик",
                    "id": 20
                }
            ],
            "description": "Список запрещенных слов по ссылке.\r\n\r\nПодключайтесь к партнерской программе Литрес!\r\n\r\nЗарабатывай с ЛитРес: Мегамаркет электронных книг №1 в России.\r\n\r\nКомпания &laquo;ЛитРес&raquo; основана в 2006 году. На сегодняшний день в каталоге представлено более 250 000 электронных книг на русском и иностранных языках. Каждый месяц ассортимент &laquo;ЛитРес&raquo; пополняется более, чем 2000 новых книг. Ежемесячно сайт компании&nbsp;litres.ru&nbsp;посещают более 2 000 000 человек.\r\n\r\nНаши приемущества:\r\n\r\n\r\n\tКомпания &laquo;ЛитРес&raquo; основана в 2005 году.\r\n\tВ каталоге более 850 000 электронных книг на русском и иностранных языках.&nbsp;\r\n\tКаждый месяц ассортимент &laquo;ЛитРес&raquo; пополняется более чем на 2000 новых книг.&nbsp;\r\n\tЕжемесячно сайт компании&nbsp;litres.ru&nbsp;посещают более&nbsp;5 000 000 человек.\r\n\tПрямые контракты с крупнейшими издательствами России (Эксмо, АСТ, Рипол-Классик, МИФ, Питер и др.) и со многими популярными авторами.&nbsp;\r\n\tАвторы-партнеры: Дарья Донцова, Борис Акунин, Вадим Панов, Александра Маринина, Евгений Гришковец и многие другие.&nbsp;\r\n\tПартнеры по продаже: Билайн, Мегафон, МТС, Softkey.ru, Google, VTB24, МВидео и множество других.\r\n\r\n\r\nВаша выгода при работе с партнерской программой &laquo;ЛитРес&raquo;:\r\n\r\n\r\n\t20% Ваше вознаграждение от каждой совершенной транзакции пользователя или с покупки всех электронных и аудиокниг на русском языке.\r\n\tВы получаете комиссию за покупку контента\r\n\t99,9% заказов одобряется\r\n\tВремя холда &ndash; всего 7 дней\r\n\tВремя Post Click Cookie составляет 60 дней\r\n\r\n\r\nПреимущества для пользователей:\r\n\r\n\r\n\tЗначительная экономия: электронные книги в разы дешевле бумажных\r\n\tУдобство использования: 14 различных форматов книг\r\n\tЛегкость и безопасность оплаты: самые популярные и удобные способы оплаты&nbsp;\r\n\tМобильность: приложения для iOS, Android, Windows 8 и Samsung Smart TV.\r\n\tАктивность на сайте: &laquo;ЛитРес&raquo; регулярно проводит акции на сайте с ценными. Акции широко освещаются на других площадках\r\n\tЗабота о клиентах: служба технической поддержки своевременно проконсультирует по всем вопросам и поможет решить возникшие трудности.\r\n\r\n\r\nЗапрещенные виды трафика:\r\n\r\nЗа нарушение &mdash; блокировка и лишение&nbsp;вознаграждения.\r\n\r\n1) Дорвеи;&nbsp;\r\n\r\n2) Контекст по брендовым запросам (ЛитРес и любые его написания, вариации и опечатки\r\n\r\nСписок запрещенных слов:\r\n\r\nлитрес\r\n\r\nлит рес\r\n\r\nлит.рес\r\n\r\nлитрес.ру\r\n\r\nkbnhtc.he\r\n\r\nдшекуыюкг\r\n\r\nkbnhtc\r\n\r\nдшекуы\r\n\r\nkbnh&#39;c\r\n\r\nлитрэс\r\n\r\nlitres.ru\r\n\r\nлитресс\r\n\r\nlitress\r\n\r\nlitres\r\n\r\nliters\r\n\r\nlitrec\r\n\r\nlit res;\r\n\r\n3) Фишинговые сайты, а также любые другие сайты, содержание которых противоречит дей ствующему законодательству РФ;\r\n\r\n4) Подсаживание cookie без открытия полноценной страницы сайта&nbsp;litres.ru",
            "cr_trend": "0.0000",
            "raw_description": "<p style=\"text-align:center\"><img alt=\"\" src=\"https://hq.admitad.com/public/storage/2017/04/11/Infografika_Litres.png\" /><br />\r\n<br />\r\n<strong>Список запрещенных слов по <a href=\"https://hq.admitad.com/public/storage/2017/04/11/Spisok_zapreshchennykh_slov_Litres.docx\" target=\"_blank\">ссылке</a>.</strong></p>\r\n\r\n<p style=\"text-align:center\"><span style=\"font-size:14px\"><strong>Подключайтесь к партнерской программе Литрес!</strong></span></p>\r\n\r\n<p>Зарабатывай с ЛитРес: Мегамаркет электронных книг №1 в России.</p>\r\n\r\n<p>Компания &laquo;ЛитРес&raquo; основана в 2006 году. На сегодняшний день в каталоге представлено более 250 000 электронных книг на русском и иностранных языках. Каждый месяц ассортимент &laquo;ЛитРес&raquo; пополняется более, чем 2000 новых книг. Ежемесячно сайт компании&nbsp;<strong>litres.ru</strong>&nbsp;посещают более 2 000 000 человек.</p>\r\n\r\n<p>Наши приемущества:</p>\r\n\r\n<ul>\r\n\t<li>Компания &laquo;ЛитРес&raquo; основана в 2005 году.</li>\r\n\t<li>В каталоге более 850 000 электронных книг на русском и иностранных языках.&nbsp;</li>\r\n\t<li>Каждый месяц ассортимент &laquo;ЛитРес&raquo; пополняется более чем на 2000 новых книг.&nbsp;</li>\r\n\t<li>Ежемесячно сайт компании&nbsp;<a href=\"http://litres.ru/\" target=\"_blank\">litres.ru</a>&nbsp;посещают более&nbsp;5 000 000 человек.</li>\r\n\t<li>Прямые контракты с крупнейшими издательствами России (Эксмо, АСТ, Рипол-Классик, МИФ, Питер и др.) и со многими популярными авторами.&nbsp;</li>\r\n\t<li>Авторы-партнеры: Дарья Донцова, Борис Акунин, Вадим Панов, Александра Маринина, Евгений Гришковец и многие другие.&nbsp;</li>\r\n\t<li>Партнеры по продаже: Билайн, Мегафон, МТС, Softkey.ru, Google, VTB24, МВидео и множество других.</li>\r\n</ul>\r\n\r\n<p><strong>Ваша выгода при работе с партнерской программой &laquo;ЛитРес&raquo;:</strong></p>\r\n\r\n<ul>\r\n\t<li>20% Ваше вознаграждение от каждой совершенной транзакции пользователя или с покупки всех электронных и аудиокниг на русском языке.</li>\r\n\t<li>Вы получаете комиссию за покупку контента</li>\r\n\t<li>99,9% заказов одобряется</li>\r\n\t<li>Время холда &ndash; всего 7 дней</li>\r\n\t<li>Время Post Click Cookie составляет 60 дней</li>\r\n</ul>\r\n\r\n<p><strong>Преимущества для пользователей:</strong></p>\r\n\r\n<ul>\r\n\t<li>Значительная экономия: электронные книги в разы дешевле бумажных</li>\r\n\t<li>Удобство использования: 14 различных форматов книг</li>\r\n\t<li>Легкость и безопасность оплаты: самые популярные и удобные способы оплаты&nbsp;</li>\r\n\t<li>Мобильность: приложения для iOS, Android, Windows 8 и Samsung Smart TV.</li>\r\n\t<li>Активность на сайте: &laquo;ЛитРес&raquo; регулярно проводит акции на сайте с ценными. Акции широко освещаются на других площадках</li>\r\n\t<li>Забота о клиентах: служба технической поддержки своевременно проконсультирует по всем вопросам и поможет решить возникшие трудности.</li>\r\n</ul>\r\n\r\n<p><strong>Запрещенные виды трафика:</strong></p>\r\n\r\n<p>За нарушение &mdash; блокировка и лишение&nbsp;вознаграждения.</p>\r\n\r\n<p>1) Дорвеи;&nbsp;</p>\r\n\r\n<p>2) Контекст по брендовым запросам (ЛитРес и любые его написания, вариации и опечатки</p>\r\n\r\n<p>Список запрещенных слов:</p>\r\n\r\n<p>литрес</p>\r\n\r\n<p>лит рес</p>\r\n\r\n<p>лит.рес</p>\r\n\r\n<p>литрес.ру</p>\r\n\r\n<p>kbnhtc.he</p>\r\n\r\n<p>дшекуыюкг</p>\r\n\r\n<p>kbnhtc</p>\r\n\r\n<p>дшекуы</p>\r\n\r\n<p>kbnh&#39;c</p>\r\n\r\n<p>литрэс</p>\r\n\r\n<p><a href=\"http://litres.ru/\" target=\"_blank\">litres.ru</a></p>\r\n\r\n<p>литресс</p>\r\n\r\n<p>litress</p>\r\n\r\n<p>litres</p>\r\n\r\n<p>liters</p>\r\n\r\n<p>litrec</p>\r\n\r\n<p>lit res;</p>\r\n\r\n<p>3) Фишинговые сайты, а также любые другие сайты, содержание которых противоречит дей ствующему законодательству РФ;</p>\r\n\r\n<p>4) Подсаживание cookie без открытия полноценной страницы сайта&nbsp;<a href=\"http://litres.ru/\" target=\"_blank\">litres.ru</a></p>\r\n",
            "modified_date": "2018-09-21T16:52:39",
            "denynewwms": false,
            "moderation": true,
            "categories": [
                {
                    "language": "ru",
                    "id": 62,
                    "parent": null,
                    "name": "Интернет-магазины"
                },
                {
                    "language": "ru",
                    "id": 70,
                    "parent": {
                        "language": "ru",
                        "id": 62,
                        "parent": null,
                        "name": "Интернет-магазины"
                    },
                    "name": "Книги"
                }
            ],
            "products_csv_link": "http://export.admitad.com/ru/webmaster/websites/710635/products/export_adv_products/?user=horechek&code=oajns6wqeo&feed_id=1721&format=csv",
            "products_xml_link": "http://export.admitad.com/ru/webmaster/websites/710635/products/export_adv_products/?user=horechek&code=oajns6wqeo&feed_id=1721&format=xml",
            "name": "ЛитРес",
            "feeds_info": [
                {
                    "advertiser_last_update": "2018-09-22 16:12:01",
                    "admitad_last_update": "2018-09-22 15:05:11",
                    "csv_link": "http://export.admitad.com/ru/webmaster/websites/710635/products/export_adv_products/?user=horechek&code=oajns6wqeo&feed_id=1721&format=csv",
                    "name": "ЛитРес",
                    "xml_link": "http://export.admitad.com/ru/webmaster/websites/710635/products/export_adv_products/?user=horechek&code=oajns6wqeo&feed_id=1721&format=xml"
                }
            ],
            "landing_code": null,
            "ecpc_trend": "0.0000",
            "epc_trend": "0.0000",
            "epc": 71.0,
            "allow_deeplink": true,
            "show_products_links": true
        },
        {
            "goto_cookie_lifetime": 7,
            "rating": "4.1",
            "exclusive": false,
            "image": "http://cdn.admitad.com/campaign/images/2016/06/27/32ab30eeb5192a42d7f8e69a64d5fa03.png",
            "actions": [
                {
                    "hold_time": 0,
                    "payment_size": "2%-10%",
                    "type": "sale",
                    "name": "Оплаченный заказ",
                    "id": 966
                }
            ],
            "avg_money_transfer_time": 10,
            "currency": "RUB",
            "activation_date": "2016-06-27T17:57:30",
            "retag": false,
            "cr": 1.8,
            "max_hold_time": null,
            "id": 1816,
            "avg_hold_time": 9,
            "ecpc": 1.16,
            "connection_status": "active",
            "gotolink": "https://ad.admitad.com/g/07d6913cf6b72ea2020f5ddd29e1bc/",
            "site_url": "http://labirint.ru/",
            "regions": [
                {
                    "region": "00"
                }
            ],
            "actions_detail": [
                {
                    "type": "sale",
                    "tariffs": [
                        {
                            "action_id": 966,
                            "rates": [
                                {
                                    "price_s": "0.00",
                                    "tariff_id": 4897,
                                    "country": null,
                                    "date_s": "2016-10-28",
                                    "is_percentage": true,
                                    "id": 52706,
                                    "size": "10.00"
                                }
                            ],
                            "id": 4897,
                            "name": "Оплаченный заказ - новый клиент"
                        },
                        {
                            "action_id": 966,
                            "rates": [
                                {
                                    "price_s": "0.00",
                                    "tariff_id": 966,
                                    "country": null,
                                    "date_s": "2016-10-28",
                                    "is_percentage": true,
                                    "id": 52705,
                                    "size": "2.00"
                                }
                            ],
                            "id": 966,
                            "name": "Оплаченный заказ - старый клиент"
                        }
                    ],
                    "hold_size": 0,
                    "name": "Оплаченный заказ",
                    "id": 966
                }
            ],
            "landing_title": null,
            "geotargeting": false,
            "status": "active",
            "coupon_iframe_denied": false,
            "traffics": [
                {
                    "enabled": true,
                    "name": "Cashback",
                    "id": 1
                },
                {
                    "enabled": false,
                    "name": "PopUp / ClickUnder",
                    "id": 2
                },
                {
                    "enabled": false,
                    "name": "Контекстная реклама",
                    "id": 3
                },
                {
                    "enabled": false,
                    "name": "Дорвей - трафик",
                    "id": 4
                },
                {
                    "enabled": true,
                    "name": "Email - рассылка",
                    "id": 5
                },
                {
                    "enabled": false,
                    "name": "Контекстная реклама на Бренд",
                    "id": 6
                },
                {
                    "enabled": true,
                    "name": "Реклама в социальных сетях",
                    "id": 7
                },
                {
                    "enabled": true,
                    "name": "Мотивированный трафик",
                    "id": 8
                },
                {
                    "enabled": false,
                    "name": "Toolbar",
                    "id": 9
                },
                {
                    "enabled": false,
                    "name": "Adult - трафик",
                    "id": 14
                },
                {
                    "enabled": true,
                    "name": "Тизерные сети",
                    "id": 18
                },
                {
                    "enabled": true,
                    "name": "Youtube Канал",
                    "id": 19
                },
                {
                    "enabled": false,
                    "name": "Брокерский трафик",
                    "id": 20
                }
            ],
            "description": "Подключайтесь с партнерской программе интрнет-магазина &laquo;Лабиринт&raquo;!\r\n&nbsp;\r\n\r\nВ оффере действует бонусная программа:\r\n\r\n1. Если за предыдущий календарный месяц Клиентами оплачено и получено&nbsp;от 101 до 500 единиц товаров:\r\nВознаграждение за старого клиента = 3%\r\nВознаграждение за нового клиента = 11%\r\n\r\n2. Если за предыдущий календарный месяц Клиентами оплачено и получено от 501 до 2000 единиц&nbsp;товаров:\r\nВознаграждение за старого клиента = 3,5%\r\nВознаграждение за нового клиента = 11,5%\r\n\r\n3. Если за предыдущий календарный месяц Клиентами оплачено и получено от 2001 до 5000 единиц товаров.\r\nВознаграждение за старого клиента = 4%\r\nВознаграждение за нового клиента = 12%\r\n\r\n4. Если за предыдущий календарный месяц Клиентами оплачено и получено&nbsp;от 5001 до 7000 единиц товаров:\r\nВознаграждение за старого клиента = 4,5%\r\nВознаграждение за нового клиента = 12,5%\r\n\r\n5. Если за предыдущий календарный месяц Клиентами оплачено и получено&nbsp;от 7001 единицы товаров:\r\nВознаграждение за старого клиента = 5%\r\nВознаграждение за нового клиента = 13%\r\n\r\nКнижный интернет-магазин &laquo;Лабиринт&raquo; &mdash; магазин книг, канцелярских товаров, фильмов, музыки, софта, игрушек.\r\n\r\n&laquo;Лабиринт&raquo;&nbsp;организован в рамках книготоргового и издательского холдинга &quot;Лабиринт&quot;. Онлайн &quot;Лабиринт&quot; входит в первую тройку самых больших книжных магазинов Москвы и России. Мы предлагаем более 200 000 товаров от 1000 производителей. У нас вы найдете художественную литературу, деловые издания, учебники, детские книги: новинки и бестселлеры; а также большой ассортимент игр, фильмов, музыки, программ, игрушек, канцтоваров по выгодным ценам.\r\n\r\nСпециально для наших покупателей мы проводим постоянные акции и конкурсы, устраиваем викторины, приглашаем авторов для общения с читателями, пишем обзоры книг.\r\nСреди других интернет-магазинов&nbsp;&laquo;Лабиринт&raquo;&nbsp;отличается удобным интерфейсом, простой навигацией, качеством обратной связи.\r\n\r\nПлюсы для ваших покупателей:\r\n\r\n+ Наша курьерская служба работает в 348 городах России!\r\n+ Заказ от 800 рублей мы привезем бесплатно!\r\n+ У нас 446 пунктов самовывоза по всей России, в городе Москва &mdash; 89!\r\n+ Множество скидок и акций для клиентов!\r\n&nbsp;\r\n\r\nПлюсы для партнеров:\r\n\r\n+ Известный Бренд - лояльность покупателей!\r\n+ Высокая конвесрия!\r\n+ Высокое вознаграждение за выкупленный заказ - до 13%!\r\n&nbsp;",
            "cr_trend": "0.0000",
            "raw_description": "<p style=\"text-align:center\"><span style=\"font-size:14px\"><strong>Подключайтесь с партнерской программе интрнет-магазина &laquo;Лабиринт&raquo;!</strong></span><br />\r\n&nbsp;</p>\r\n\r\n<p style=\"text-align:center\"><span style=\"font-size:14px\"><strong>В оффере действует бонусная программа:</strong></span></p>\r\n\r\n<p>1. Если за предыдущий календарный месяц Клиентами оплачено и получено&nbsp;<strong>от 101 до 500 </strong>единиц товаров:<br />\r\nВознаграждение за старого клиента = 3%<br />\r\nВознаграждение за нового клиента = 11%</p>\r\n\r\n<p>2. Если за предыдущий календарный месяц Клиентами оплачено и получено <strong>от 501 до 2000 </strong>единиц&nbsp;товаров<strong>:</strong><br />\r\nВознаграждение за старого клиента = 3,5%<br />\r\nВознаграждение за нового клиента = 11,5%</p>\r\n\r\n<p>3. Если за предыдущий календарный месяц Клиентами оплачено и получено <strong>от 2001 до 5000 </strong>единиц товаров.<br />\r\nВознаграждение за старого клиента = 4%<br />\r\nВознаграждение за нового клиента = 12%</p>\r\n\r\n<p>4. Если за предыдущий календарный месяц Клиентами оплачено и получено&nbsp;<strong>от 5001 до 7000 </strong>единиц товаров:<br />\r\nВознаграждение за старого клиента = 4,5%<br />\r\nВознаграждение за нового клиента = 12,5%</p>\r\n\r\n<p>5. Если за предыдущий календарный месяц Клиентами оплачено и получено&nbsp;<strong>от 7001 </strong>единицы товаров:<br />\r\nВознаграждение за старого клиента = 5%<br />\r\nВознаграждение за нового клиента = 13%</p>\r\n\r\n<p>Книжный интернет-магазин <strong>&laquo;Лабиринт&raquo;</strong> &mdash; магазин книг, канцелярских товаров, фильмов, музыки, софта, игрушек.</p>\r\n\r\n<p><strong>&laquo;Лабиринт&raquo;</strong>&nbsp;организован в рамках книготоргового и издательского холдинга &quot;Лабиринт&quot;. Онлайн &quot;Лабиринт&quot; входит в первую тройку самых больших книжных магазинов Москвы и России. Мы предлагаем <strong>более 200 000 товаров от 1000 производителей</strong>. У нас вы найдете художественную литературу, деловые издания, учебники, детские книги: новинки и бестселлеры; а также большой ассортимент игр, фильмов, музыки, программ, игрушек, канцтоваров по выгодным ценам.<br />\r\n<br />\r\nСпециально для наших покупателей мы проводим постоянные акции и конкурсы, устраиваем викторины, приглашаем авторов для общения с читателями, пишем обзоры книг.<br />\r\nСреди других интернет-магазинов&nbsp;<strong>&laquo;Лабиринт&raquo;</strong>&nbsp;отличается удобным интерфейсом, простой навигацией, качеством обратной связи.</p>\r\n\r\n<p><strong>Плюсы для ваших покупателей:</strong></p>\r\n\r\n<p>+ Наша курьерская служба работает в 348 городах России!<br />\r\n+ Заказ от 800 рублей мы привезем бесплатно!<br />\r\n+ У нас 446 пунктов самовывоза по всей России, в городе Москва &mdash; 89!<br />\r\n+ Множество скидок и акций для клиентов!<br />\r\n&nbsp;</p>\r\n\r\n<p><strong>Плюсы для партнеров:</strong><br />\r\n<br />\r\n+ Известный Бренд - лояльность покупателей!<br />\r\n+ Высокая конвесрия!<br />\r\n+ Высокое вознаграждение за выкупленный заказ - до 13%!<br />\r\n&nbsp;</p>\r\n",
            "modified_date": "2018-09-10T12:16:57",
            "denynewwms": false,
            "moderation": true,
            "categories": [
                {
                    "language": "ru",
                    "id": 62,
                    "parent": null,
                    "name": "Интернет-магазины"
                },
                {
                    "language": "ru",
                    "id": 69,
                    "parent": {
                        "language": "ru",
                        "id": 62,
                        "parent": null,
                        "name": "Интернет-магазины"
                    },
                    "name": "Товары для детей"
                },
                {
                    "language": "ru",
                    "id": 70,
                    "parent": {
                        "language": "ru",
                        "id": 62,
                        "parent": null,
                        "name": "Интернет-магазины"
                    },
                    "name": "Книги"
                },
                {
                    "language": "ru",
                    "id": 71,
                    "parent": {
                        "language": "ru",
                        "id": 62,
                        "parent": null,
                        "name": "Интернет-магазины"
                    },
                    "name": "Аксессуары"
                },
                {
                    "language": "ru",
                    "id": 72,
                    "parent": {
                        "language": "ru",
                        "id": 62,
                        "parent": null,
                        "name": "Интернет-магазины"
                    },
                    "name": "Подарки & Цветы"
                },
                {
                    "language": "ru",
                    "id": 89,
                    "parent": {
                        "language": "ru",
                        "id": 62,
                        "parent": null,
                        "name": "Интернет-магазины"
                    },
                    "name": "Товары для творчества"
                }
            ],
            "products_csv_link": "http://export.admitad.com/ru/webmaster/websites/710635/products/export_adv_products/?user=horechek&code=oajns6wqeo&feed_id=14514&format=csv",
            "products_xml_link": "http://export.admitad.com/ru/webmaster/websites/710635/products/export_adv_products/?user=horechek&code=oajns6wqeo&feed_id=14514&format=xml",
            "name": "Лабиринт",
            "feeds_info": [
                {
                    "advertiser_last_update": "2018-09-22 05:25:00",
                    "admitad_last_update": "2018-09-22 03:31:57",
                    "csv_link": "http://export.admitad.com/ru/webmaster/websites/710635/products/export_adv_products/?user=horechek&code=oajns6wqeo&feed_id=14514&format=csv",
                    "name": "Основной",
                    "xml_link": "http://export.admitad.com/ru/webmaster/websites/710635/products/export_adv_products/?user=horechek&code=oajns6wqeo&feed_id=14514&format=xml"
                }
            ],
            "landing_code": null,
            "ecpc_trend": "0.0000",
            "epc_trend": "0.0000",
            "epc": 116.0,
            "allow_deeplink": true,
            "show_products_links": false
        }
    ],
    "_meta": {
        "count": 10,
        "limit": 2,
        "offset": 0
    }
}`
	advCampaignTestData = `{
    "goto_cookie_lifetime": 45,
    "rating": "3.7",
    "rate_of_approve": "85",
    "more_rules": "<p><span style=\"color:#FF0000\"><strong>В оффере запрещено размещать неактуальные (просроченные, нерабочие) промо-коды и акции, все веб-мастера, нарушающие правила работы, будут отключены от программы!</strong></span></p>\r\n\r\n<p>Важно! Запрещена контекстная реклама на Бренд.&nbsp;</p>\r\n\r\n<p>Минус слова :</p>\r\n\r\n<p>Буквоед</p>\r\n\r\n<p>Bookvoed</p>\r\n\r\n<p>Промокод буквоед</p>\r\n\r\n<p>Скидка буквоед</p>\r\n\r\n<p>Акции буквоед</p>\r\n\r\n<p>Купон буквоед</p>\r\n\r\n<p>Промокод&nbsp;Bookvoed</p>\r\n\r\n<p>Скидка&nbsp;Bookvoed</p>\r\n\r\n<p>Акции&nbsp;Bookvoed</p>\r\n\r\n<p>Купон&nbsp;Bookvoed</p>\r\n\r\n<p>&nbsp;</p>\r\n\r\n<p>И все прочие запросы, которые содержат название бренда, включая его неправильное написание.</p>\r\n\r\n<p>&nbsp;</p>\r\n",
    "exclusive": false,
    "image": "http://cdn.admitad.com/campaign/images/2016/08/23/5e91e47f3c461414cd32813a3648b01e.jpg",
    "actions": [
        {
            "hold_time": 0,
            "payment_size": "9%",
            "type": "sale",
            "name": "Оплаченный  заказ",
            "id": 4928
        }
    ],
    "avg_money_transfer_time": 33,
    "currency": "RUB",
    "name_aliases": "Буквоед",
    "activation_date": "2013-05-08T15:30:09",
    "cr": 3.1,
    "ecpc": 2.25,
    "id": 3063,
    "individual_terms": true,
    "modified_date": "2018-08-20T13:06:09",
    "avg_hold_time": 32,
    "site_url": "http://bookvoed.ru/",
    "regions": [
        {
            "region": "RU"
        }
    ],
    "landing_title": null,
    "geotargeting": false,
    "status": "active",
    "coupon_iframe_denied": false,
    "traffics": [
        {
            "enabled": true,
            "name": "Cashback",
            "id": 1
        },
        {
            "enabled": false,
            "name": "PopUp / ClickUnder",
            "id": 2
        },
        {
            "enabled": false,
            "name": "Контекстная реклама",
            "id": 3
        },
        {
            "enabled": true,
            "name": "Дорвей - трафик",
            "id": 4
        },
        {
            "enabled": true,
            "name": "Email - рассылка",
            "id": 5
        },
        {
            "enabled": false,
            "name": "Контекстная реклама на Бренд",
            "id": 6
        },
        {
            "enabled": true,
            "name": "Реклама в социальных сетях",
            "id": 7
        },
        {
            "enabled": false,
            "name": "Мотивированный трафик",
            "id": 8
        },
        {
            "enabled": true,
            "name": "Toolbar",
            "id": 9
        },
        {
            "enabled": false,
            "name": "Adult - трафик",
            "id": 14
        },
        {
            "enabled": true,
            "name": "Тизерные сети",
            "id": 18
        },
        {
            "enabled": true,
            "name": "Youtube Канал",
            "id": 19
        }
    ],
    "description": "Подключайтесь к партнерской программе &laquo;Буквоед&raquo;!\r\n\r\n\r\nИзвестная Петербургская Книжная сеть &laquo;Буквоед&raquo; &ndash; это инновационная, социально ответственная, динамично развивающаяся сеть.\r\n\r\nВ настоящее время книготорговая сеть &laquo;Буквоед&raquo; состоит более чем из 100 магазинов и предлагает читателям более 125 000 наименований книг. &laquo;Буквоед&raquo; входит в пятерку крупнейших книготорговых предприятий России и занимает первое место по товарообороту на Северо-Западе.\r\n\r\nВ 2008 году в рамках сети &laquo;Буквоед&raquo; был создан Интернет-магазин, который предлагает весь ассортимент сети.\r\n\r\n&nbsp;\r\n\r\nРаботая с партнерской программой Буквоед, ваши покупатели получают огромное количество плюсов:\r\n&nbsp;\r\n\r\n- Круглосуточно работает call-центр\r\n\r\n- Почтовая доставка по всей России, в Санкт-Петербурге действует собственная курьерская служба\r\n\r\n- Разнообразные способы оплаты\r\n\r\n- Огромный каталог продукции\r\n&nbsp;\r\n\r\nМы будем рады видеть вас в числе наших партнеров!\r\n\r\n&nbsp;\r\n\r\nУсловия по работе с партнерской программой Буквоед:\r\n\r\nВремя жизни Post Click Cookie* - 45\r\n\r\nВремя холда** - 30\r\n\r\nСообщаем Вам, что с 1 сентября 2015 года все веб-мастера с купонными площадками будут получать 5% за оплаченный заказ.\r\n&nbsp;\r\n\r\nОсновные характеристики портрета целевой аудитории:\r\n\r\n1. ГЕО: Санкт-Петербург и весь СЗ регион (в первую очередь), вся РФ, с 2013 будут по всему миру доставлять.\r\n\r\n2. ДЕМОГРАФИЯ: возраст разнообразный, 20-50 лет, 60% женщины, в основном семейные, подавляющей частью - офисные работники и заказы делаются в основном с работы; образование 70% высшее или неоконченное высшее;\r\n\r\n3. ЭКОНОМИЧЕСКИЕ ФАКТОРЫ: уровень дохода - средний и выше;\r\n\r\n&nbsp;\r\n\r\nЕсли у Вас возникли вопросы или вы хотите получить дополнительную информацию -&nbsp; мы всегда готовы помочь и ответить на все ваши вопросы!",
    "cr_trend": "0.0000",
    "raw_description": "<p style=\"text-align: center;\"><span style=\"font-size:14px\"><strong><span style=\"color:rgb(128, 0, 128)\">Подключайтесь к партнерской программе &laquo;Буквоед&raquo;!</span></strong></span><br />\r\n<img alt=\"\" src=\"https://cdn.admitad-connect.com/public/storage/2017/06/30/1247x300.jpg\" style=\"height:300px; width:1247px\" /></p>\r\n\r\n<p><span style=\"font-size:14px\">Известная Петербургская Книжная сеть <strong>&laquo;Буквоед&raquo;</strong> &ndash; это инновационная, социально ответственная, динамично развивающаяся сеть.</span></p>\r\n\r\n<p><span style=\"font-size:14px\">В настоящее время книготорговая сеть <strong>&laquo;Буквоед&raquo;</strong> состоит более чем из 100 магазинов и предлагает читателям более 125 000 наименований книг. &laquo;Буквоед&raquo; входит в пятерку крупнейших книготорговых предприятий России и занимает первое место по товарообороту на Северо-Западе.</span></p>\r\n\r\n<p><span style=\"font-size:14px\">В 2008 году в рамках сети <strong>&laquo;Буквоед&raquo;</strong> был создан Интернет-магазин, который предлагает весь ассортимент сети.</span></p>\r\n\r\n<p>&nbsp;</p>\r\n\r\n<p><span style=\"font-size:14px\"><strong><span style=\"color:rgb(128, 0, 128)\">Работая с партнерской программой Буквоед, ваши покупатели получают огромное количество плюсов:</span></strong></span><br />\r\n&nbsp;</p>\r\n\r\n<p><span style=\"font-size:14px\">- Круглосуточно работает call-центр</span></p>\r\n\r\n<p><span style=\"font-size:14px\">- Почтовая доставка по всей России, в Санкт-Петербурге действует собственная курьерская служба</span></p>\r\n\r\n<p><span style=\"font-size:14px\">- Разнообразные способы оплаты</span></p>\r\n\r\n<p><span style=\"font-size:14px\">- Огромный каталог продукции</span><br />\r\n&nbsp;</p>\r\n\r\n<p><span style=\"font-size:14px\">Мы будем рады видеть вас в числе наших партнеров!</span></p>\r\n\r\n<p>&nbsp;</p>\r\n\r\n<p><span style=\"font-size:14px\"><span style=\"color:rgb(128, 0, 128)\"><strong>Условия по работе с партнерской программой Буквоед:</strong></span></span></p>\r\n\r\n<p><span style=\"font-size:14px\">Время жизни Post Click Cookie* - 45</span></p>\r\n\r\n<p><span style=\"font-size:14px\">Время холда** - 30<br />\r\n<br />\r\n<span style=\"color:#ff0000\"><strong>Сообщаем Вам, что с 1 сентября 2015 года все веб-мастера с купонными площадками будут получать 5% за оплаченный заказ.</strong></span></span><br />\r\n&nbsp;</p>\r\n\r\n<p><span style=\"font-size:14px\"><span style=\"color:rgb(128, 0, 128)\"><strong>Основные характеристики портрета целевой аудитории:</strong></span></span></p>\r\n\r\n<p><span style=\"font-size:14px\">1. ГЕО: Санкт-Петербург и весь СЗ регион (в первую очередь), вся РФ, с 2013 будут по всему миру доставлять.</span></p>\r\n\r\n<p><span style=\"font-size:14px\">2. ДЕМОГРАФИЯ: возраст разнообразный, 20-50 лет, 60% женщины, в основном семейные, подавляющей частью - офисные работники и заказы делаются в основном с работы; образование 70% высшее или неоконченное высшее;</span></p>\r\n\r\n<p><span style=\"font-size:14px\">3. ЭКОНОМИЧЕСКИЕ ФАКТОРЫ: уровень дохода - средний и выше;</span></p>\r\n\r\n<p>&nbsp;</p>\r\n\r\n<p><span style=\"font-size:14px\">Если у Вас возникли вопросы или вы хотите получить дополнительную информацию -&nbsp; мы всегда готовы помочь и ответить на все ваши вопросы!</span></p>\r\n",
    "denynewwms": false,
    "connected": true,
    "max_hold_time": null,
    "categories": [
        {
            "language": "ru",
            "id": 62,
            "parent": null,
            "name": "Интернет-магазины"
        },
        {
            "language": "ru",
            "id": 70,
            "parent": {
                "language": "ru",
                "id": 62,
                "parent": null,
                "name": "Интернет-магазины"
            },
            "name": "Книги"
        }
    ],
    "retag": true,
    "name": "Буквоед",
    "landing_code": null,
    "ecpc_trend": "0.0000",
    "epc_trend": "0.0000",
    "action_type": "sale",
    "epc": 225.0,
    "allow_deeplink": true,
    "show_products_links": true
}`
)
