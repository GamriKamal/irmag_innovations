package irmag.io.ProductService.Entity;

import java.util.Arrays;
import java.util.List;

public class InfoAboutDepartmentsAndCategories {
    private final List<Department> departments = Arrays.asList(
            new Department(2, "Компьютерная периферия", List.of(
                    new Category(1, "Периферия", "https://shop.kz/upload/iblock/74c/per-komp.jpg", List.of(
                            new Category(47, "Мониторы", "https://shop.kz/upload/iblock/7ca/1428409257_samsung_curved_monitor.jpg"),
                            new Category(48, "Крепления для мониторов", "https://shop.kz/upload/iblock/a0f/krepmon.jpg"),
                            new Category(49, "Мыши", "https://shop.kz/upload/iblock/30f/1mysh.jpg"),
                            new Category(50, "Коврики для мышей", "https://shop.kz/upload/iblock/ef7/1kovrik.jpg"),
                            new Category(51, "Клавиатуры", "https://shop.kz/upload/iblock/e43/1klaviaturf.jpg"),
                            new Category(52, "Аксессуары для клавиатур", "https://shop.kz/upload/iblock/be3/eduq05049l1w7frk4spz6d7tsuy4oc2l/aksessa_kl.jpg"),
                            new Category(53, "Наклейки для клавиатур", "https://shop.kz/upload/iblock/44e/nakl.jpg"),
                            new Category(54, "Акустические системы и колонки", "https://shop.kz/upload/iblock/130/1akkusticheskie_sistemy.jpg"),
                            new Category(55, "Наушники и гарнитуры", "https://shop.kz/upload/iblock/3cc/1garnitura.jpg"),
                            new Category(56, "Микрофоны", "https://shop.kz/upload/iblock/d36/1mikrofon.jpg"),
                            new Category(57, "Веб-камеры", "https://shop.kz/upload/iblock/975/1web.jpg"),
                            new Category(58, "ИБП (UPS)", "https://shop.kz/upload/iblock/038/1ibp.jpg"),
                            new Category(59, "Аккумуляторы для ИБП", "https://shop.kz/upload/iblock/62f/akb.jpg"),
                            new Category(60, "Стабилизаторы", "https://shop.kz/upload/iblock/7ce/1stabilizator.jpg"),
                            new Category(61, "Сетевые фильтры, адаптеры", "https://shop.kz/upload/iblock/026/1setevoy-filtr.jpg"),
                            new Category(62, "Джойстики, рули, геймпады", "https://shop.kz/upload/iblock/aa5/1dzhoystik.jpg"),
                            new Category(63, "Контроллеры для стриминга", "https://shop.kz/upload/iblock/ae8/gy4hsfa4lsfx8tvyk9fb3t72htxi5ze0/kontrollery.jpg"),
                            new Category(64, "Графические планшеты", "https://shop.kz/upload/iblock/fa0/1graf-planshet.jpg"),
                            new Category(65, "USB-hub", "https://shop.kz/upload/iblock/b25/1hub.jpg"),
                            new Category(66, "Картридеры", "https://shop.kz/upload/iblock/92d/1kartrider.jpg")
                    )),
                    new Category(2, "Носители информации", "https://shop.kz/upload/iblock/0dd/per-komp11.jpg", List.of(
                            new Category(67, "Флешки", "https://shop.kz/upload/iblock/9f5/1fleshki.jpg"),
                            new Category(68, "Карты памяти", "https://shop.kz/upload/iblock/ded/1karty-pamyati.jpg"),
                            new Category(69, "USB внешние накопители", "https://shop.kz/upload/iblock/a1c/1hdd.jpg"),
                            new Category(70, "Сетевые хранилища", "https://shop.kz/upload/iblock/153/1setevye-khranilishcha.jpg")
                    ))
            )),
            new Department(3, "Развлечения и отдых", List.of(
                    new Category(1, "Развлечения", "https://shop.kz/upload/iblock/8b0/raz-i-ot.jpg", List.of(
                            new Category(11, "Бинокли", "https://shop.kz/upload/iblock/3ca/1binokli.jpg"),
                            new Category(12, "Электронные книги", "https://shop.kz/upload/iblock/647/1elektronnye-knigi.jpg")
                    )),
                    new Category(2, "Путешествия и отдых", "https://shop.kz/upload/iblock/a88/put-i-otd.jpg", List.of(
                            new Category(29, "Чемоданы", "https://shop.kz/upload/iblock/bc9/1turizm.jpg"),
                            new Category(30, "Сумки и рюкзаки", "https://shop.kz/upload/iblock/3c8/dip1_1_.jpg"),
                            new Category(31, "Туризм и отдых на природе", "https://shop.kz/upload/iblock/bb1/turizm1.jpg"),
                            new Category(32, "Надувные матрасы", "https://shop.kz/upload/iblock/613/1naduvnye-matrasy.jpg"),
                            new Category(33, "Пляжный отдых", "https://shop.kz/upload/iblock/e69/1plyazhnyy-otdykh.jpg"),
                            new Category(34, "Бассейны", "https://shop.kz/upload/iblock/1b0/1basseyn.jpg"),
                            new Category(35, "Аксессуары для бассейнов", "https://shop.kz/upload/iblock/5d6/1aksessuary-dlya-basseynov.jpg"),
                            new Category(36, "Батуты", "https://shop.kz/upload/iblock/e9b/batut.jpg"),
                            new Category(37, "Надувные лодки", "https://shop.kz/upload/iblock/e90/bmz57bodfmrb4n1zawk3o66tl5m25wjz/boats.jpg")
                    )),
                    new Category(3, "Транспорт", "https://shop.kz/upload/iblock/c93/put-i-otd222.jpg", List.of(
                            new Category(38, "Гироскутеры", "https://shop.kz/upload/iblock/816/1giroskuter.jpg"),
                            new Category(39, "Сигвеи", "https://shop.kz/upload/iblock/70f/1sigvei.jpg"),
                            new Category(40, "Электросамокаты и велосипеды", "https://shop.kz/upload/iblock/bab/1elektrovelosiped.jpg")
                    )),
                    new Category(4, "Игрушки", "https://shop.kz/upload/iblock/9ba/igrus.jpg", List.of(
                            new Category(41, "Радиоуправляемые игрушки", "https://shop.kz/upload/iblock/700/1radioupravlyaemye-mashinki.jpg"),
                            new Category(42, "Игрушечные машинки", "https://shop.kz/upload/iblock/d47/1mashinki.jpg"),
                            new Category(43, "Конструкторы", "https://shop.kz/upload/iblock/77f/1konstruktory.jpg"),
                            new Category(44, "LEGO", "https://shop.kz/upload/iblock/ce4/w87w7antc8rmjmdloork8gazlk874ohf/lego.jpg"),
                            new Category(45, "Развивающие игрушки и робототехника", "https://shop.kz/upload/iblock/cae/robot150_120.jpg"),
                            new Category(46, "3D ручки", "https://shop.kz/upload/iblock/091/3druch.jpg")
                    ))
            )),
            new Department(4, "Бытовая техника", List.of(
                    new Category(1, "Техника для кухни", "https://shop.kz/upload/iblock/2d3/tekhn-dlya-kukh.jpg", List.of(
                            new Category(51, "Холодильники", "https://shop.kz/upload/iblock/0d4/1kholodilnik.jpg"),
                            new Category(52, "Морозильники", "https://shop.kz/upload/iblock/af6/moroz.jpg"),
                            new Category(53, "Льдогенераторы", "https://shop.kz/upload/iblock/a17/icegen.jpg"),
                            new Category(54, "Винные шкафы", "https://shop.kz/upload/iblock/962/winesk.jpg"),
                            new Category(55, "Микроволновые печи", "https://shop.kz/upload/iblock/118/1mikrovolnovka_2.jpg"),
                            new Category(56, "Варочные поверхности", "https://shop.kz/upload/iblock/01e/1varochnye-poverkhnosti.jpg"),
                            new Category(57, "Встраиваемые духовки", "https://shop.kz/upload/iblock/79b/1vstraivaemaya-dukhovka.jpg"),
                            new Category(58, "Электропечи", "https://shop.kz/upload/iblock/600/elektrop.jpg"),
                            new Category(59, "Плиты", "https://shop.kz/upload/iblock/f14/1plity.jpg"),
                            new Category(60, "Настольные плиты", "https://shop.kz/upload/iblock/37d/plitnast.jpg"),
                            new Category(61, "Вытяжки", "https://shop.kz/upload/iblock/feb/1vytyazhki.jpg"),
                            new Category(62, "Чайники, термопоты", "https://shop.kz/upload/iblock/0ef/1chaynik.jpg"),
                            new Category(63, "Блендеры, соковыжималки", "https://shop.kz/upload/iblock/1a7/1blendery-sokovyzhimalki.jpg"),
                            new Category(64, "Миксеры", "https://shop.kz/upload/iblock/d70/mixer.jpg"),
                            new Category(65, "Кухонные комбайны и машины", "https://shop.kz/upload/iblock/2a1/kitchmach.jpg"),
                            new Category(66, "Мультиварки, пароварки", "https://shop.kz/upload/iblock/f8a/1multivarka.jpg"),
                            new Category(67, "Мороженицы", "https://shop.kz/upload/iblock/b3d/2hglila2prwxo6998yak0x52wjp43bz2/morozhen.jpg"),
                            new Category(68, "Сушилки для продуктов", "https://shop.kz/upload/iblock/7ba/00rhq17vxkvpc2n71u3fpo1doxbrgdvc/sush.jpg"),
                            new Category(69, "Йогуртницы", "https://shop.kz/upload/iblock/456/j5q8lqn6b93qbmm27nkbsmp6rr5shgbu/yogurt.jpg"),
                            new Category(70, "Хлебопечи", "https://shop.kz/upload/iblock/90d/1khlebopech.jpg"),
                            new Category(71, "Кофеварки, кофемашины", "https://shop.kz/upload/iblock/0cd/1kofevarnki.jpg"),
                            new Category(72, "Кофемолки", "https://shop.kz/upload/iblock/7c0/kofem.jpg"),
                            new Category(73, "Турки", "https://shop.kz/upload/iblock/586/oq5e9f0386178squk9xz08179kwj2rdd/turki.jpg"),
                            new Category(74, "Весы кухонные", "https://shop.kz/upload/iblock/289/1vesy-kukhonnye.jpg"),
                            new Category(75, "Тостеры, гриль", "https://shop.kz/upload/iblock/acb/1toster.jpg"),
                            new Category(76, "Вакууматоры и аксессуары", "https://shop.kz/upload/iblock/265/hhpmpo06irpq51yf6v1n3vcbg362fsgu/vacuum.jpg"),
                            new Category(77, "Вафельницы, блинницы", "https://shop.kz/upload/iblock/7eb/waff.jpg"),
                            new Category(78, "Фритюрницы", "https://shop.kz/upload/iblock/aec/fri.jpg"),
                            new Category(79, "Сэндвичницы", "https://shop.kz/upload/iblock/76c/1amt7c3fk2x620m2l3v4u3twe2tjstkb/sandwich.jpg"),
                            new Category(80, "Электромясорубки", "https://shop.kz/upload/iblock/d22/1elektromyasorubka.jpg"),
                            new Category(81, "Электрошашлычницы", "https://shop.kz/upload/iblock/ac3/ted7cw8wiawej9rdxe4ctndg6fr6c6yz/shashl.jpg"),
                            new Category(82, "Гаджеты для кухни", "https://shop.kz/upload/iblock/8fa/gadjeti_kuhni_150_120.jpg")
                    )),
                    new Category(2, "Климатическая техника", "https://shop.kz/upload/iblock/b08/Klimaticheskaya-tekhnika.jpg", List.of(
                            new Category(21, "Кондиционеры", "https://shop.kz/upload/iblock/a8f/kondery.png"),
                            new Category(22, "Обогреватели", "https://shop.kz/upload/iblock/e63/2-obogrevateli.jpg"),
                            new Category(23, "Вентиляторы", "https://shop.kz/upload/iblock/6d5/158.jpg"),
                            new Category(24, "Увлажнители и очистители воздуха", "https://shop.kz/upload/iblock/1d5/uvlajniteli150_120.jpg")
                    )),
                    new Category(3, "Техника для дома", "https://shop.kz/upload/iblock/944/tekh-dlya-doma.jpg", List.of(
                            new Category(9, "Пылесосы", "https://shop.kz/upload/iblock/e7b/1pylesosy.jpg"),
                            new Category(10, "Аксессуары для пылесосов", "https://shop.kz/upload/iblock/df3/2zhhxndrcmzhaef826dcf35zo2zr5mfo/aksesspyles.jpg"),
                            new Category(11, "Стиральные машины", "https://shop.kz/upload/iblock/d8c/1stiralnye-mashiny.jpg"),
                            new Category(12, "Водонагреватели", "https://shop.kz/upload/iblock/b4d/1boylery.jpg"),
                            new Category(13, "Краны-водонагреватели", "https://shop.kz/upload/iblock/53e/kran.jpg"),
                            new Category(14, "Утюги, отпариватели", "https://shop.kz/upload/iblock/ee6/1utyug.jpg"),
                            new Category(15, "Уход за одеждой", "https://shop.kz/upload/iblock/928/ukhod.jpg"),
                            new Category(16, "Уборка дома", "https://shop.kz/upload/iblock/3c6/9et8dyw0fp6gmre9x1k8m5v7dzw7z4nv/shvabry150x120.jpg"),
                            new Category(17, "Пароочистители", "https://shop.kz/upload/iblock/7af/pzy0by3enht1jscza2d2mavf62shxwh8/paroochistiteli150x120.jpg"),
                            new Category(18, "Стеклоочистители", "https://shop.kz/upload/iblock/f7f/4bzhspc1ew8yf3uasggkzvyxiwbkvkvk/robotmoyshchik.jpg"),
                            new Category(19, "Телефоны", "https://shop.kz/upload/iblock/76f/1telefony.jpg"),
                            new Category(20, "Стабилизаторы", "https://shop.kz/upload/iblock/f1c/1stabilizator.jpg")
                    ))
            )),
            new Department(5, "Товары для дома", List.of(
                    new Category(1, "Инструменты и электроинструменты", "https://shop.kz/upload/iblock/c0b/qibeltotekbwf9uroofqhqipd9vtffnk/Instrumenty_i_elektroinstrumenty.jpg", List.of(
                            new Category(11, "Дрели, шуруповерты", "https://shop.kz/upload/iblock/4ff/105ah6uqqlhrrj3gk5xpczogdqovmvac/image_shurup.jpg"),
                            new Category(12, "Перфораторы", "https://shop.kz/upload/iblock/859/037y7myistvflgmknsm92osyy02qp4be/image_perf.jpg"),
                            new Category(13, "Шлифовальные машины", "https://shop.kz/upload/iblock/6b7/w0dhe7aaxtz32vri7d05iplgo6ew937b/image_SHlifovalnye.jpg"),
                            new Category(14, "Электролобзики, пилы", "https://shop.kz/upload/iblock/962/kzqodx39u29tykzjanngzqwch1ub0csl/image_Elektrolobziki.jpg"),
                            new Category(15, "Аккумуляторы для электроинструментов", "https://shop.kz/upload/iblock/08e/m0zsxcu6uyo085nwehftt6wrqmvapjfa/image_Akkumulyatory.jpg"),
                            new Category(16, "Мойки высокого давления", "https://shop.kz/upload/iblock/aaa/3nt3gkghwzova3vluhlqsd2tivrk1h4a/Moyka2.jpg"),
                            new Category(17, "Наборы инструментов", "https://shop.kz/upload/iblock/e08/icxvqk8c6np83p3m4njoivdy0lktw3m5/Nabor_instrumentov.jpg"),
                            new Category(18, "Обжимной инструмент", "https://shop.kz/upload/iblock/90a/4alkfc6qltz2dvtvre8vq0voa116ic1i/objim.jpg"),
                            new Category(19, "Лазерные дальномеры", "https://shop.kz/upload/iblock/bfc/grpbhnoewt541ebficl8jzk48he9f7t4/dalnomery.jpg"),
                            new Category(20, "Паяльное оборудование", "https://shop.kz/upload/iblock/77e/xeq8rj0qx62e2j72ok5ac6u0u3iu3pxp/Payalnoe_oborudovanie.jpg"),
                            new Category(21, "Сварочное оборудование", "https://shop.kz/upload/iblock/a8a/h4obbo81e1fifjkhzwb02rqxmb28ufka/svar.jpg"),
                            new Category(22, "Инструменты", "https://shop.kz/upload/iblock/360/nk36kf3nh227dj3wxtttc3qhi6pqday8/images.jpg"),
                            new Category(23, "Отвертки", "https://shop.kz/upload/iblock/f76/fc2137n8zosdpa6xzni3sw77hzcjjms6/otvertki.jpg"),
                            new Category(24, "Садовая техника", "https://shop.kz/upload/iblock/4eb/i2vvjfpda6pha8i9qoamq9w4utoqa68x/sadovaya_tehnica.jpg"),
                            new Category(25, "Прочие инструменты, оборудование", "https://shop.kz/upload/iblock/8cf/zzgdpsdpcq6cajyga0pcrdfx5px7udsr/ProchOborud.jpg")
                    )),
                    new Category(2, "Товары Karcher", "https://shop.kz/upload/iblock/613/lid20j3116n2yvdjoffc7tw4nhwlacni/Tovary-Karcher.jpg", List.of(
                            new Category(21, "Мойки высокого давления Karcher", "https://shop.kz/upload/iblock/abf/zuulsw7f6ah75zeb9pnfhtde531whqd3/kerkher-_-kopiya-123.jpg"),
                            new Category(22, "Пылесосы Karcher", "https://shop.kz/upload/iblock/5ac/y3n9ymc96vrxef5kgmf69cxdaekngzpn/Pylesosy-Karcher.jpg"),
                            new Category(23, "Бытовая химия", "https://shop.kz/upload/iblock/140/55dydd868m9y0j1aqq9himf2r75kmslc/Byt_khimiya.jpg"),
                            new Category(24, "Оборудование для полива", "https://shop.kz/upload/iblock/531/ran5nfu2w749vgmzsi1ir06ppqfk07k7/oborud_dlya_poliva.jpg")
                    )),
                    new Category(3, "Посуда", "https://shop.kz/upload/iblock/28b/Posuda.jpg", List.of(
                            new Category(31, "Сковороды, сотейники", "https://shop.kz/upload/iblock/618/pans.jpg"),
                            new Category(32, "Кастрюли", "https://shop.kzhttps://shop.kz/upload/iblock/e4e/pots.jpg"),
                            new Category(33, "Кухонные принадлежности", "/upload/iblock/19f/noji150_120.jpg"),
                            new Category(34, "Ножи", "https://shop.kz/upload/iblock/c30/1ukrusegrginnnj6ddtpt52hagj9xzea/knives.jpg"),
                            new Category(35, "Посуда для сервировки", "https://shop.kz/upload/iblock/ff3/stolovyj_nabor_150_120.jpg")
                    )),
                    new Category(4, "Умный дом", "https://shop.kz/upload/iblock/ea4/Byt-tekh.jpg", List.of(
                            new Category(41, "Системы видеонаблюдения", "https://shop.kz/upload/iblock/cd5/1kamery-videonablyudeniya.jpg"),
                            new Category(42, "Датчики, устройства и управление для умного дома", "https://shop.kz/upload/iblock/e0b/2signalizatsii.jpg"),
                            new Category(43, "Товары для животных", "https://shop.kz/upload/iblock/86e/5d1w4dg1bg5rg0dz6wirauhq06lhrwyy/zhivotnie-150x120.jpg"),
                            new Category(44, "Умное освещение", "https://shop.kz/upload/iblock/221/smartlamp.jpg"),
                            new Category(45, "Роботы-пылесосы", "https://shop.kz/upload/iblock/4a1/robot.jpg"),
                            new Category(46, "Колонки Яндекс Станция (Алиса)", "https://shop.kz/upload/iblock/cb8/yandexstation.jpg")
                    )),
                    new Category(5, "Красота и здоровье", "https://shop.kz/upload/iblock/ec2/krasota-i.jpg", List.of(
                            new Category(51, "Весы", "https://shop.kz/upload/iblock/d80/1vesy.jpg"),
                            new Category(52, "Электробритвы, эпиляторы", "https://shop.kz/upload/iblock/f09/1britvy.jpg"),
                            new Category(53, "Машинки для стрижки, триммеры", "https://shop.kz/upload/iblock/db0/trim.jpg"),
                            new Category(54, "Фены, электрощипцы", "https://shop.kz/upload/iblock/3ba/1feny.jpg"),
                            new Category(55, "Здоровье и гигиена", "https://shop.kz/upload/iblock/3bf/zdorovie_and_gigiena.jpg"),
                            new Category(56, "Приборы для чистки лица", "https://shop.kz/upload/iblock/1fa/m2nzz1i0ancro32yfqa6ot9pcd1tmcfw/Kak_vybrat_pribor_dlya_chistki_litsa_OKACHI_GLIYA.jpg"),
                            new Category(57, "Уход за полостью рта", "https://shop.kz/upload/iblock/aaa/cleanmouth.jpg"),
                            new Category(58, "Массажеры", "https://shop.kz/upload/iblock/387/o61ygwslkrr9ogrqyyjnonm4dkrtu75v/massage.jpg"),
                            new Category(59, "Беговые дорожки", "https://shop.kz/upload/iblock/172/b0ppvapdmj13cdldkneglaovpjo2cj2a/running.jpg")
                    )),
                    new Category(6, "Мебель", "https://shop.kz/upload/iblock/5fc/mebel.jpg", List.of(
                            new Category(61, "Столы", "https://shop.kz/upload/iblock/52e/1stoly.jpg"),
                            new Category(62, "Компьютерные кресла", "https://shop.kz/upload/iblock/f6f/1kresla.jpg")
                    )),
                    new Category(7, "Прочее", "https://shop.kz/upload/iblock/ccf/91463v5rywlnd4mvfi4wn1crgeocjfpy/prochee.jpg", List.of(
                            new Category(71, "Часы настенные", "https://shop.kz/upload/iblock/ab2/12prpidf3qm52v52oy2i35oq9bzwlkd1/chasyz1.jpg"),
                            new Category(72, "Швейные машинки, оверлоки", "https://shop.kz/upload/iblock/171/sm.jpg"),
                            new Category(73, "Лампы, светильники", "https://shop.kz/upload/iblock/1dc/1lampochki.jpg"),
                            new Category(74, "Фонари", "https://shop.kz/upload/iblock/b8c/1fonar.jpg"),
                            new Category(75, "Батарейки, зарядные устройства", "https://shop.kz/upload/iblock/774/1batareyki-zaryadki.jpg"),
                            new Category(76, "Чистящие средства", "https://shop.kz/upload/iblock/d74/1chistyashchie.jpg")
                    ))
            )),
            new Department(6, "Ноутбуки и компьютеры", List.of(
                    new Category(1, "Ноутбуки", "https://shop.kz/upload/iblock/dfc/Sotovye-telefony.jpg", List.of(
                            new Category(11, "Ноутбуки", "https://shop.kz/upload/iblock/fc3/noutbuk-_sdelat-oboi-s-interfeysom-vindy_.jpg"),
                            new Category(12, "Ультрабуки", "https://shop.kz/upload/iblock/aa5/ultrabuk.jpg"),
                            new Category(13, "Сумки для ноутбуков", "https://shop.kz/upload/iblock/275/1sumki.jpg"),
                            new Category(14, "Подставки для охлаждения", "https://shop.kz/upload/iblock/a4d/1podstavki.jpg"),
                            new Category(15, "Зарядные устройства для ноутбуков", "https://shop.kz/upload/iblock/c4e/1zaryadnye-ustroystva.jpg")
                    )),
                    new Category(2, "Компьютеры", "https://shop.kz/upload/iblock/50c/Gadzhety.jpg", List.of(
                            new Category(21, "Моноблоки", "https://shop.kz/upload/iblock/fc6/1monoblok.jpg"),
                            new Category(22, "Настольные компьютеры", "https://shop.kz/upload/iblock/739/1nastolnyy-pk.jpg"),
                            new Category(23, "Неттопы", "https://shop.kz/upload/iblock/09a/1nettop.jpg"),
                            new Category(24, "Мониторы", "https://shop.kz/upload/iblock/92c/1428409257_samsung_curved_monitor.jpg")
                    )),
                    new Category(3, "Программное обеспечение", "https://shop.kz/upload/iblock/80c/Aksessuary.jpg", List.of(
                            new Category(31, "Антивирусы", "https://shop.kz/upload/iblock/b79/1antivirus.png"),
                            new Category(32, "Операционные системы", "https://shop.kz/upload/iblock/dc4/1os.jpg"),
                            new Category(33, "Офисные программы", "https://shop.kz/upload/iblock/577/1ofisnye-prilozheniya.png"),
                            new Category(34, "Цифровые подписки", "https://shop.kz/upload/iblock/315/Digital.jpg")
                    )),
                    new Category(4, "Комплектующие для ноутбуков", "https://shop.kz/upload/iblock/80c/Aksessuary.jpg", List.of(
                            new Category(41, "SSD диски", "https://shop.kz/upload/iblock/e9c/1ssd.jpg"),
                            new Category(42, "Жесткие диски для ноутбуков", "https://shop.kz/upload/iblock/4ba/1zhestkie-diski.jpg"),
                            new Category(43, "Оперативная память для ноутбуков", "https://shop.kz/upload/iblock/371/1ozu.jpg"),
                            new Category(44, "Прочие комплектующие", "https://shop.kz/upload/iblock/a98/1prochie-komplektuyushchie.jpg")
                    )),
                    new Category(5, "Аксессуары", "https://shop.kz/upload/iblock/80c/Aksessuary.jpg", List.of(
                            new Category(51, "Чистящие средства", "https://shop.kz/upload/iblock/0b7/1chistyashchie.jpg"),
                            new Category(52, "Кабели, переходники", "https://shop.kz/upload/iblock/d1e/1kabeli-i-perekhodniki.jpg"),
                            new Category(53, "Док-станции, концентраторы", "https://shop.kz/upload/iblock/492/Dok_stanciya.jpg"),
                            new Category(54, "Прочие аксессуары", "https://shop.kz/upload/iblock/888/1aksessuary.jpg")
                    ))
            )),
            new Department(7, "Сетевое и серверное оборудование", List.of(
                    new Category(1, "Сетевое оборудование", "https://shop.kz/upload/iblock/dfc/Sotovye-telefony.jpg", List.of(
                            new Category(11, "Роутеры, модемы", "https://shop.kz/upload/iblock/900/1router.jpg"),
                            new Category(12, "Сетевые карты и адаптеры", "https://shop.kz/upload/iblock/306/1setevaya-karta.jpg"),
                            new Category(13, "Точки доступа", "https://shop.kz/upload/iblock/b92/1tochka-dostupa.jpg"),
                            new Category(14, "Коммутаторы (switch)", "https://shop.kz/upload/iblock/7a8/2instrumenty.jpg"),
                            new Category(15, "Антенны, усилители сигнала", "https://shop.kz/upload/iblock/bf9/1usilitel-signala.jpg"),
                            new Category(16, "Powerline-адаптеры", "https://shop.kz/upload/iblock/e11/1powerline_adaptery.jpg"),
                            new Category(17, "Сетевые хранилища", "https://shop.kz/upload/iblock/f32/1setevye-khranilishcha.jpg"),
                            new Category(18, "Медиаконвертеры", "https://shop.kz/upload/iblock/4a7/mediakonvektory.jpg"),
                            new Category(19, "Прочее оборудование", "https://shop.kz/upload/iblock/435/1prochee-oborudovanie.jpg")
                    )),
                    new Category(2, "Серверное оборудование", "https://shop.kz/upload/iblock/50c/Gadzhety.jpg", List.of(
                            new Category(21, "Серверы", "https://shop.kz/upload/iblock/fb7/1servery.jpg"),
                            new Category(22, "Комплектующие для серверов", "https://shop.kz/upload/iblock/282/1komlpektuyushchie-dlya-serverov.jpg"),
                            new Category(23, "Серверные шкафы", "https://shop.kz/upload/iblock/19e/1servernye-shkafy.png"),
                            new Category(24, "Комплектующие для шкафов", "https://shop.kz/upload/iblock/814/1komlpektuyushchie-dlya-servernykh-shkafov.jpg")
                    )),
                    new Category(3, "Монтажное оборудование", "https://shop.kz/upload/iblock/80c/Aksessuary.jpg", List.of(
                            new Category(31, "Кабели, коннекторы", "https://shop.kz/upload/iblock/69f/1kabeli-i-perekhodniki.jpg"),
                            new Category(32, "Кабель-каналы, фурнитура", "https://shop.kz/upload/iblock/cff/1kabel-kanaly.jpg"),
                            new Category(33, "Инструменты", "https://shop.kz/upload/iblock/12b/1instrumenty.jpg")
                    ))
            )),
            new Department(8, "Оргтехника и расходные материалы", List.of(
                    new Category(1, "Печатные устройства", "https://shop.kz/upload/iblock/c39/150_120_org-teh-pech.jpg", List.of(
                            new Category(11, "МФУ лазерные", "https://shop.kz/upload/iblock/67d/1mfu-lazer.jpg"),
                            new Category(12, "МФУ струйные", "https://shop.kz/upload/iblock/2cc/1struynyy-mfu.jpg"),
                            new Category(13, "Принтеры лазерные", "https://shop.kz/upload/iblock/5a2/1-printer-lazernyy.jpg"),
                            new Category(14, "Принтеры струйные", "https://shop.kz/upload/iblock/e47/1struynyy-printer.jpg"),
                            new Category(15, "Копировальные аппараты", "https://shop.kz/upload/iblock/08c/1kopir.jpg"),
                            new Category(16, "Плоттеры", "https://shop.kz/upload/iblock/fdd/1plotter.jpg"),
                            new Category(17, "3D принтеры", "https://shop.kz/upload/iblock/689/13D-printer.jpg"),
                            new Category(18, "3D ручки", "https://shop.kz/upload/iblock/f38/3druch.jpg")
                    )),
                    new Category(2, "Офисное и торговое оборудование", "https://shop.kz/upload/iblock/d20/150x120-of-i-torg.jpg", List.of(
                            new Category(21, "Телефоны", "https://shop.kz/upload/iblock/5fe/1telefony.jpg"),
                            new Category(22, "Резаки", "https://shop.kz/upload/iblock/f5b/1rezak.jpg"),
                            new Category(23, "Ламинаторы", "https://shop.kz/upload/iblock/fa6/1laminator.jpg"),
                            new Category(24, "Уничтожители документов", "https://shop.kz/upload/iblock/9ea/1unichtozhitel.jpg"),
                            new Category(25, "Переплетное оборудование", "https://shop.kz/upload/iblock/559/1perepletnaya-mashina.jpg"),
                            new Category(26, "Торговое оборудование", "https://shop.kz/upload/iblock/586/1torgovoe-oborudovanie.jpg")
                    )),
                    new Category(3, "Демонстрационное оборудование", "https://shop.kz/upload/iblock/9a5/150_120_dem-ob.jpg", List.of(
                            new Category(31, "Проекторы", "https://shop.kz/upload/iblock/366/1proektor.jpg"),
                            new Category(32, "Экраны проекционные", "https://shop.kz/upload/iblock/0b5/1proektsionnyy-ekran.jpg"),
                            new Category(33, "Интерактивные панели", "https://shop.kz/upload/iblock/a95/1qkl4usbvg3eo7yk8sqnqle1r7cghdxg/interpanel150_120.jpg"),
                            new Category(34, "Крепления для проекторов", "https://shop.kz/upload/iblock/24d/1kreplenie-dlya-proektorov.jpg"),
                            new Category(35, "Маркерные доски и аксессуары", "https://shop.kz/upload/iblock/881/1markernye-doski.jpg")
                    )),
                    new Category(4, "Расходные материалы и аксессуары", "https://shop.kz/upload/iblock/e34/150_120_rashod.jpg", List.of(
                            new Category(41, "Бумага, фотобумага", "https://shop.kz/upload/iblock/237/1bumaga.jpg"),
                            new Category(42, "Лазерные картриджи", "https://shop.kz/upload/iblock/fa4/1lazernyy-kartridzh.jpg"),
                            new Category(43, "Струйные картриджи и чернила", "https://shop.kz/upload/iblock/af4/1chernila.jpg"),
                            new Category(44, "Тонеры для лазерных принтеров", "https://shop.kz/upload/iblock/20c/1tonery-dlya-lazern.jpg"),
                            new Category(45, "Расходные материалы для лазерных принтеров", "https://shop.kz/upload/iblock/3b4/1raskhodnye-dlya-lazernykh.jpg"),
                            new Category(46, "Матричные картриджи", "https://shop.kz/upload/iblock/98e/1matrichnyy-kartridzh.jpg"),
                            new Category(47, "Расходные материалы для копиров", "https://shop.kz/upload/iblock/2b0/1tonery-dlya-kopirov.jpg"),
                            new Category(48, "Пленки для ламинаторов", "https://shop.kz/upload/iblock/f24/1plenki-dlya-laminatorov.jpg"),
                            new Category(49, "Аксессуары для переплетных машин", "https://shop.kz/upload/iblock/b87/1pruzhiny.jpg"),
                            new Category(50, "Пластик для 3D печати", "https://shop.kz/upload/iblock/c79/1plastik.jpg"),
                            new Category(51, "Кабели, переходники", "https://shop.kz/upload/iblock/228/1kabeli-i-perekhodniki.jpg"),
                            new Category(52, "Чистящие средства", "https://shop.kz/upload/iblock/497/1chistyashchie.jpg")
                    ))
            )),
            new Department(9, "Товары для геймеров", List.of(
                    new Category(1, "Ноутбуки, ПК и приставки", "https://shop.kz/upload/iblock/993/pc-i-prist.jpg", List.of(
                            new Category(11, "Игровые ноутбуки", "https://shop.kz/upload/iblock/0de/1noutbuki.jpg"),
                            new Category(12, "Игровые компьютеры", "https://shop.kz/upload/iblock/967/1nastolnye-pk.jpg")
                    )),
                    new Category(2, "Комплектующие", "https://shop.kz/upload/iblock/d7a/nout_kompl.jpg", List.of(
                            new Category(21, "Все для сборки компьютера", "https://shop.kz/upload/iblock/e2c/vse-dlya-sborki.jpg"),
                            new Category(22, "Дополнительные комплектующие", "https://shop.kz/upload/iblock/69c/dop-aks.jpg")
                    )),
                    new Category(3, "Периферия", "https://shop.kz/upload/iblock/e10/Bez-imeni_1-_7_.jpg", List.of(
                            new Category(31, "Игровые мыши", "https://shop.kz/upload/iblock/bfa/1mysh.jpg"),
                            new Category(32, "Игровые коврики для мышей", "https://shop.kz/upload/iblock/c76/1kovrik.jpg"),
                            new Category(33, "Игровые клавиатуры", "https://shop.kz/upload/iblock/f26/1klaviatura.jpg"),
                            new Category(34, "Игровые наушники", "https://shop.kz/upload/iblock/900/1garnitura.jpg"),
                            new Category(35, "Джойстики, рули, геймпады", "https://shop.kz/upload/iblock/6e4/1dzhoystik.jpg"),
                            new Category(36, "Контроллеры для стриминга", "https://shop.kz/upload/iblock/da5/p5hqgi7lm6pn3u82ft0ezv6kxtjmancy/kontrollery.jpg"),
                            new Category(37, "Игровые мониторы", "https://shop.kz/upload/iblock/7e5/1428409257_samsung_curved_monitor.jpg")
                    )),
                    new Category(4, "Мебель", "https://shop.kz/upload/iblock/1bb/Bez-imeni_1-_8_mebel.jpg", List.of(
                            new Category(41, "Игровые кресла", "https://shop.kz/upload/iblock/7d5/1kreslo.jpg"),
                            new Category(42, "Столы для геймеров", "https://shop.kz/upload/iblock/2f5/tablegamer.png")
                    ))
            )),
            new Department(10, "Смартфоны и гаджеты", List.of(
                    new Category(1, "Сотовые телефоны", "https://shop.kz/upload/iblock/dfc/Sotovye-telefony.jpg", List.of(
                            new Category(11, "Смартфоны", "https://shop.kz/upload/iblock/cdf/1smartfony2.jpg"),
                            new Category(12, "Кнопочные телефоны", "https://shop.kz/upload/iblock/dff/1knopochnye-telefony-1.jpg")
                    )),
                    new Category(2, "Гаджеты", "https://shop.kz/upload/iblock/50c/Gadzhety.jpg", List.of(
                            new Category(21, "Смарт-часы", "https://shop.kz/upload/iblock/2e8/1smartchasy-2.jpg"),
                            new Category(22, "Фитнес-браслеты", "https://shop.kz/upload/iblock/7b1/1fitnes-braslety-2.png"),
                            new Category(23, "Планшеты", "https://shop.kz/upload/iblock/a60/1planshety5.jpg"),
                            new Category(24, "Электронные книги", "https://shop.kz/upload/iblock/a06/1elektronnye-knigi.jpg")
                    )),
                    new Category(3, "Аксессуары", "https://shop.kz/upload/iblock/80c/Aksessuary.jpg", List.of(
                            new Category(31, "Бронепленки", "https://shop.kz/upload/iblock/0a7/f142rqw7ac8bf0r7296ta0zend6n1q01/tpu_2.jpg"),
                            new Category(32, "Защитные стекла и пленки", "https://shop.kz/upload/iblock/c3b/stekla150.jpg"),
                            new Category(33, "Наушники для телефонов", "https://shop.kz/upload/iblock/761/1naushniki.jpg"),
                            new Category(34, "Беспроводные наушники и гарнитуры", "https://shop.kz/upload/iblock/e97/wireless_earphones_2.jpg"),
                            new Category(35, "Портативные колонки", "https://shop.kz/upload/iblock/896/1portativnye-kolonki.jpg"),
                            new Category(36, "Чехлы для смартфонов", "https://shop.kz/upload/iblock/6f6/1chekhol_dlya_telefona.jpg"),
                            new Category(37, "Чехлы для электронных книг", "https://shop.kz/upload/iblock/050/g3vi4f8cr8uw88lq0rhnaugnt9p177so/cheink.jpg"),
                            new Category(38, "Крепления, подставки для смартфонов", "https://shop.kz/upload/iblock/4f9/q1orqvwkk0sg3g8hpzl3ultv95abx47g/forsmart.jpg"),
                            new Category(39, "Power bank (мобильные аккумуляторы)", "https://shop.kz/upload/iblock/7b2/1mobilnye-akkumulyatory.jpg"),
                            new Category(40, "Моноподы", "https://shop.kz/upload/iblock/a10/1monopod.jpg"),
                            new Category(41, "Зарядные устройства", "https://shop.kz/upload/iblock/aca/1zaryadnoe.jpg"),
                            new Category(42, "Кабели и переходники", "https://shop.kz/upload/iblock/f38/08e95da8_e77f_466d_b81f_b5a83de50a02.jpg"),
                            new Category(43, "Карты памяти", "https://shop.kz/upload/iblock/bcd/1karty-pamyati.jpg"),
                            new Category(43, "Ремешки для смарт-часов и фитнес-браслетов", "https://shop.kz/upload/iblock/58b/jph8pi8bm8qk4igm6gce5w6neitnfpf9/remeshok.jpg"),
                            new Category(43, "Чехлы для наушников", "https://shop.kz/upload/iblock/f7a/chehly.jpg"),
                            new Category(43, "Кольцевые лампы", "https://shop.kz/upload/iblock/e46/150x120led_lamps.jpg"),
                            new Category(43, "Чистящие средства", "https://shop.kz/upload/iblock/c35/1chistyashchie.jpg"),
                            new Category(43, "Прочие аксессуары", "https://shop.kz/upload/iblock/0c4/1prochie-aksessuary.jpg")

                    ))
            )),
            new Department(11, "Телевизоры, аудио и видео", List.of(
                    new Category(1, "TV", "https://shop.kz/upload/upload/iblock/444/TV.jpg", List.of(
                            new Category(11, "Телевизоры", "https://shop.kz/upload/iblock/ab2/1televizor.jpg"),
                            new Category(12, "Саундбары", "https://shop.kz/upload/iblock/148/soundbary.jpg"),
                            new Category(13, "Медиаплееры", "https://shop.kz/upload/iblock/5d0/1mediapleer.jpg"),
                            new Category(14, "ТВ-тюнеры", "https://shop.kz/upload/iblock/e3e/1tv-tyunery.jpg"),
                            new Category(15, "Кронштейны для телевизоров", "https://shop.kz/upload/iblock/579/1kronshteyny-dlya-televizorov.jpg"),
                            new Category(16, "Пульты", "https://shop.kz/upload/iblock/034/2kamery-videonablyudeniya.jpg"),
                            new Category(17, "Антенны для телевизора", "https://shop.kz/upload/iblock/454/AntennCatalog.jpg")
                    )),
                    new Category(2, "Видео", "https://shop.kz/upload/iblock/19d/video_TV.jpg", List.of(
                            new Category(21, "Экшн-камеры", "https://shop.kz/upload/iblock/313/1ekshn_kamery.jpg"),
                            new Category(22, "Аксессуары для экшн-камер", "https://shop.kz/upload/iblock/88d/1aksessuary-dlya-ekshn_kamer.jpg"),
                            new Category(23, "Аксессуары для квадрокоптеров", "https://shop.kz/upload/iblock/691/1aksessuary-dlya-kvadrokopterov.jpg"),
                            new Category(24, "Камеры видеонаблюдения", "https://shop.kz/upload/iblock/f44/1kamery-videonablyudeniya.jpg"),
                            new Category(25, "Веб-камеры", "https://shop.kz/upload/iblock/33f/1web.jpg")
                    )),
                    new Category(3, "Демонстрационное оборудование", "https://shop.kz/upload/iblock/61a/demo-ob.jpg", List.of(
                            new Category(31, "Проекторы", "https://shop.kz/upload/iblock/72d/1proektor.jpg"),
                            new Category(32, "Экраны проекционные", "https://shop.kz/upload/iblock/1ff/1proektsionnyy-ekran.jpg"),
                            new Category(33, "Крепления для проекторов", "https://shop.kz/upload/iblock/37e/1kreplenie-dlya-proektorov.jpg")
                    )),
                    new Category(4, "Аудио", "https://shop.kz/upload/iblock/dca/audio.jpg", List.of(
                            new Category(41, "Аудиотехника, Hi-Fi оборудование", "https://shop.kz/upload/iblock/ded/1kinoteatr.jpg"),
                            new Category(42, "Плееры", "https://shop.kz/upload/iblock/08b/813636.jpg"),
                            new Category(43, "Портативные колонки", "https://shop.kz/upload/iblock/15d/1portativnye-kolonki.jpg"),
                            new Category(44, "Акустические системы и колонки", "https://shop.kz/upload/iblock/5e5/1akkusticheskie-sistemy.jpg"),
                            new Category(45, "Наушники и гарнитуры", "https://shop.kz/upload/iblock/6e8/1garnitura.jpg"),
                            new Category(46, "Наушники для телефонов", "https://shop.kz/upload/iblock/a48/1naushniki.jpg"),
                            new Category(47, "Беспроводные наушники и гарнитуры", "https://shop.kz/upload/iblock/492/wireless_earphones_2.jpg"),
                            new Category(48, "Микрофоны", "https://shop.kz/upload/iblock/232/1mikrofon.jpg"),
                            new Category(49, "Аксессуары для микрофонов", "https://shop.kz/upload/iblock/417/zgsumjzk31uvuloidzdfv0sha0r4y0e6/acsessmic.jpg"),
                            new Category(50, "Звуковые карты и усилители", "https://shop.kz/upload/iblock/dc5/1zvukovaya-karta.jpg")

                    )),
                    new Category(5, "Фототехника", "https://shop.kz/upload/iblock/b1c/aotoap.jpg", List.of(
                            new Category(51, "Фотоаппараты", "https://shop.kz/upload/iblock/056/1fotoapparaty.jpg"),
                            new Category(52, "Аксессуары для фотоаппаратов", "https://shop.kz/upload/iblock/67b/1aksessuary-dlya-fotokamer.jpg"),
                            new Category(53, "Сумки для фото и видео", "https://shop.kz/upload/iblock/1ed/1sumki-dlya-foto-i-video.jpg")
                    )),
                    new Category(6, "Аксессуары", "https://shop.kz/upload/iblock/cc1/tv-acs.jpg", List.of(
                            new Category(61, "Карты памяти", "https://shop.kz/upload/iblock/37a/1karty-pamyati.jpg"),
                            new Category(62, "Штативы, моноподы", "https://shop.kz/upload/iblock/6cf/1shtativy_-monopody.jpg"),
                            new Category(63, "Зарядные устройства", "https://shop.kz/upload/iblock/efc/1zaryadnoe.jpg"),
                            new Category(64, "Чистящие средства", "https://shop.kz/upload/iblock/559/1chistyashchie.jpg"),
                            new Category(65, "Чехлы для наушников", "https://shop.kz/upload/iblock/e0f/chehly.jpg"),
                            new Category(66, "Кабели, переходники", "https://shop.kz/upload/iblock/a16/1kabeli-i-perekhodniki.jpg")

                    ))
            ))
    );

    private static final List<Category> categories = Arrays.asList(
            new Category(46, "Колонки Яндекс Станция (Алиса)", "https://shop.kz/upload/iblock/cb8/yandexstation.jpg"),
            new Category(11, "Ноутбуки", "https://shop.kz/upload/iblock/fc3/noutbuk-_sdelat-oboi-s-interfeysom-vindy_.jpg"),
            new Category(41, "SSD диски", "https://shop.kz/upload/iblock/e9c/1ssd.jpg"),
            new Category(47, "Мониторы", "https://shop.kz/upload/iblock/7ca/1428409257_samsung_curved_monitor.jpg"),
            new Category(11, "Смартфоны", "https://shop.kz/upload/iblock/cdf/1smartfony2.jpg"));


    public List<Department> getDepartments() {
        return departments;
    }

    public static List<Category> getCategories() {
            return categories;
    }
}
