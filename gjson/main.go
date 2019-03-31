package main

import (
	"encoding/json"
	"fmt"
	"time"
)

var jsonStr string = `
{
    "includes": {
        "Asset": null,
        "Entry": [
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "colorFacets",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2016-10-13T17:53:08.889Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "10RSc6oVqEKMiEQiaQo0qk",
                    "locale": "en-US",
                    "revision": 1,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2016-10-13T17:53:08.889Z"
                },
                "Fields": {
                    "color": "multi",
                    "swatchUrl": "http://images.freepeople.com/is/image/FreePeople/MULTI_swatch"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "mobileReviewTemplateFieldModule",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2018-10-02T20:35:54.844Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "1Ebej9lTMkMoWaoII0CWY0",
                    "locale": "en-US",
                    "revision": 2,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2018-10-10T16:01:15.742Z"
                },
                "Fields": {
                    "fieldName": "sendemailalertwhenpublished",
                    "fieldType": "switch",
                    "title": "Email When Published"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "colorFacets",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2016-10-13T17:54:32.951Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "1M7RcVhYRi2YUE0eWcK4iy",
                    "locale": "en-US",
                    "revision": 1,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2016-10-13T17:54:32.951Z"
                },
                "Fields": {
                    "color": "assorted",
                    "displayName": "Assorted",
                    "hexCode": null,
                    "swatchUrl": "http://images.anthropologie.com/is/image/Anthropologie/assorted-facet-1"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "facetNameMapping",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2016-11-18T21:22:22.439Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "1NGk4d67OkmuWWi2eCWyGy",
                    "locale": "en-US",
                    "revision": 1,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2016-11-18T21:22:22.439Z"
                },
                "Fields": {
                    "groupByName": "tile.product.brand",
                    "name": "brand"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "settingsMobileApp",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2017-02-23T18:00:11.037Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "1Ra1cJWxaUeCGo4qm6kAiM",
                    "locale": "en-US",
                    "revision": 15,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2018-02-28T17:03:15.332Z"
                },
                "Fields": {
                    "environment": "dev",
                    "incompatibleOsMessage": "Please visit your phone settings & update your software to continue using the app. TEST MESSAGE UPDATE.",
                    "invalidAppMessage": "This app is out-of-date. Please visit the App Store to update it now. If an update is not available, please try again in a few hours. TEST MESSAGE UPDATE.",
                    "maintenanceMessage": "We're currently doing some housekeeping and the application should be available again by 6:00 AM EST. Please, check back in a couple hours.",
                    "maintenanceTitle": "Server Maintenance",
                    "maintenanceToggle": false,
                    "minAppVersion": "1.0",
                    "minOsVersion": "9"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "mobileCatalogContainer",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2019-02-14T18:34:47.743Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "1ZnccGS16oASut7WZ1aGNq",
                    "locale": "en-US",
                    "revision": 3,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2019-02-18T17:56:47.527Z"
                },
                "Fields": {
                    "blacklistTrendingTerms": [
                        "ping pong",
                        "coachella",
                        "cloth stone"
                    ],
                    "catalogFacetsDisplayName": [
                        {
                            "sys": {
                                "id": "6m26hRsQA86GsgQgCK2yqs",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "4hwlrOlius0MSqCwOkoAqA",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "5dw5dVLSSAcis2UmugKqsY",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "4kwl02GY7CIQQgWOMQiiyA",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "1tweOUrfWQ6oaMkae4cK0A",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "6ayJHXLT3ic8gS4egWkYm",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "2ysyPHFgucMo8aAkaMYu0w",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "4YEdyHtEzYYy2AWE60GgeC",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "66x4U9qZcA4SWAaCSG8ISg",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "5trVAe5sCQsIWeqQ2iuIeC",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "1KSMTYFJ0MSWi8wYOoOcwC",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "Gp2wNI0ciieaQe8c6eiqs",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "Y2c31AxQaGY2AMo6EU0y8",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        }
                    ],
                    "colorFacets": [
                        {
                            "sys": {
                                "id": "12fi9t3usqyaOcY8m8ewGS",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "5a2lznIMvYyQyIsUIcgqoS",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "6HEI2JdrmEw8IAgoc4gsco",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "MhHRuFZcAKkWwm6CMQkoY",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "5xWaNQxu6sU0yaIie0wqK4",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "4jL2AAHD2Eao2gcO8CSeyw",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "5eP2TVIeKkWWgUQ2A0mwGA",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "231lIJOzz6oYWAUsuquSsc",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "2LshNHGmPuAsqM8oYWcqeq",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "1muYaRWPVWEcEkgwWuYiOk",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "2XlwiKZmg8SKI2ue62QgcK",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "10RSc6oVqEKMiEQiaQo0qk",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "3cnDzqNPyoMYGYyCaauIqq",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "3Ge3MI5DIcaCSwY6ycSA6i",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "6oaf1Dh2MMQQCOm6Uagc4w",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "2GPMXdQjgQq4MCkWYSCoya",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "4EwmYFs0icMi0qaKAYKigK",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "5vMvozzMsMomOQIqyso68E",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "6RnURl0XcsMSkY8sgAugk2",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "6Mrp9RvgtiAcIw0iSyWYYk",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "1oTiIQhHiwYi4eC2i6UUOW",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "2K2TkOFH6MmA8oaeqw8ccQ",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        }
                    ],
                    "facetNameMappings": [
                        {
                            "sys": {
                                "id": "1NGk4d67OkmuWWi2eCWyGy",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "4CnXCZ3NfOssouCQqWaGaW",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "3JADy40qEE0sEGSMyMSi2y",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "35ASYaHK5qAyOqKgOc2AKi",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "4YNYblCJsQuO2gcwMCuWEE",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "3Dz7crZFUIeE6QscAKuuSo",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "3bCaTu6gEw6k4EGCqOSKCy",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "6KyCGGzjc408G6ACikmygy",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "7vnQciMIjCwomSimIAuIOK",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "baWGtlOeswcKsWGYCakAu",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "4mHZFzK4rmwQsWksqCmkGC",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "Pxrcb9I3Sg8YGkkywUeQI",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        }
                    ],
                    "giftWrapImageCode": "b",
                    "lowStockThreshold": 5,
                    "reviewTemplateContainer": [
                        {
                            "sys": {
                                "id": "52iz9NHTgsyiGw4oakoo2E",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "3R8QX1U6jeqcMIsAkAe0W8",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "ydAnrQ8jBuGOEwqw4ss0q",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "1ZrzOjS0pKW0KQYqakKQOK",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "62paV3Lb1Kyy0AQccCy4Kc",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        }
                    ],
                    "title": "Catalog",
                    "topRatedThreshold": 4.5
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "mobileReviewTemplateContainer",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2018-08-03T18:29:20.025Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "1ZrzOjS0pKW0KQYqakKQOK",
                    "locale": "en-US",
                    "revision": 3,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2019-01-24T16:25:36.482Z"
                },
                "Fields": {
                    "headerTitle": "Tell Others About Yourself",
                    "referenceModules": [
                        {
                            "sys": {
                                "id": "3hWrSaDWM8KMsgWwMWy2g6",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "2isYBattJyEucWoMisao4o",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "6tWumz5SVOaiceEcKyG6Yg",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "2stKsabuWMSW68SqGwqCgM",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "5miErEaVMcQ2AuaocumEyY",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "3313tdqFGUauiMKYMsueu2",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "3Nmi9D58re6sSC6EwWUYYy",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "3iF0p7TwwoC28IUAwC2Wqw",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "34O15Ci9jbHBkqrbm1PTGN",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "1qSF6MgPVei8YqC0YyaeuQ",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "5YVPIALjLaO4ymSmYu044G",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "3j7o9WdNTOY28CMGGgeEiE",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "4UJfqKW0E8ce00Mi0QYisa",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "6wO4NM12w0UYaCyiMse0AC",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "MWy1hHcx4ksQsmwqACw4U",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "4UhiS3uZ9KskCUKM6Oce8U",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "6suyexivTyyoC4Uke0Ke6A",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "F2ErF2XRx6aO4Kw6YCKMo",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "4NDCbQLBBYuguq8wowEYG0",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "5yi00B2Kic66os086seKsO",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "5cgIIFlMHCGiGqKY8CgEK0",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        }
                    ],
                    "title": "Group 4 Template"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "colorFacets",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2016-10-13T17:54:28.440Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "1muYaRWPVWEcEkgwWuYiOk",
                    "locale": "en-US",
                    "revision": 1,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2016-10-13T17:54:28.440Z"
                },
                "Fields": {
                    "color": "grey",
                    "displayName": "Gray",
                    "hexCode": "8c8c8c"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "colorFacets",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2016-10-13T17:54:26.782Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "1oTiIQhHiwYi4eC2i6UUOW",
                    "locale": "en-US",
                    "revision": 1,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2016-10-13T17:54:26.782Z"
                },
                "Fields": {
                    "color": "white",
                    "displayName": "White",
                    "hexCode": "ffffff"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "mobileReviewTemplateFieldModule",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2018-08-03T18:21:39.687Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "1qSF6MgPVei8YqC0YyaeuQ",
                    "locale": "en-US",
                    "revision": 1,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2018-08-03T18:21:39.687Z"
                },
                "Fields": {
                    "fieldName": "additionalfield_SizePurchased",
                    "fieldType": "floating",
                    "title": "Size Purchased"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "colorFacets",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2016-10-13T17:54:30.796Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "231lIJOzz6oYWAUsuquSsc",
                    "locale": "en-US",
                    "revision": 1,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2016-10-13T17:54:30.796Z"
                },
                "Fields": {
                    "color": "gold",
                    "hexCode": "eac975"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "markdownModule",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2016-10-13T18:27:20.077Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "25pMHCK1XaeIc8sM486oYQ",
                    "locale": "en-US",
                    "revision": 1,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2016-10-13T18:27:20.077Z"
                },
                "Fields": {
                    "bodyText": "### Flat Rate Returns\n\nIn-stock furniture, rugs and other oversized items can be returned within 30 days from the date of delivery for a full refund of the merchandise cost. Anthropologie will refund or replace items where a manufacturer defect was found within one year from the date of delivery. Visit [Returns & Exchanges](/help/returns-exchanges) for more information on completing a return.",
                    "slug": "product_returns-flat-rate",
                    "title": "PDP Flat Rate Returns Copy"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "markdownModule",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2017-06-14T15:06:17.643Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "2DHTKcBkViciU8EKcM80aM",
                    "locale": "en-US",
                    "revision": 9,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2017-08-01T13:21:30.256Z"
                },
                "Fields": {
                    "bodyText": "1. Order is placed\n2. Custom furniture is made\n3. Order is prepared\n4. Delivery is scheduled\n5. Schedule is confirmed\n6. Order is delivered",
                    "slug": "mto_shipping-next-steps",
                    "title": "MTO Shipping Next Steps list"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "colorFacets",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2016-10-13T17:54:31.310Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "2K2TkOFH6MmA8oaeqw8ccQ",
                    "locale": "en-US",
                    "revision": 1,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2016-10-13T17:54:31.310Z"
                },
                "Fields": {
                    "color": "yellow",
                    "displayName": "Yellow",
                    "hexCode": "feea00"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "mobileAppleMusicContainer",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2019-02-14T18:31:08.788Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "2KnKQZnaSFTT1WVaXwRH7D",
                    "locale": "en-US",
                    "revision": 1,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2019-02-14T18:31:08.788Z"
                },
                "Fields": {
                    "title": "Apple Music"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "colorFacets",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2016-10-13T17:54:29.687Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "2LshNHGmPuAsqM8oYWcqeq",
                    "locale": "en-US",
                    "revision": 1,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2016-10-13T17:54:29.687Z"
                },
                "Fields": {
                    "color": "green",
                    "hexCode": "457701"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "colorFacets",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2016-10-13T17:54:33.948Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "2QGjZek2VOmEyOEAAWKi0M",
                    "locale": "en-US",
                    "revision": 1,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2016-10-13T17:54:33.948Z"
                },
                "Fields": {
                    "color": "assorted",
                    "displayName": "Assorted",
                    "swatchUrl": "https://images.anthropologie.com/is/image/Anthropologie/assorted-facet-1"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "markdownModule",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2016-10-13T18:27:12.969Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "2Tp8q95upGQYAwOsgAUU28",
                    "locale": "en-US",
                    "revision": 3,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2017-07-31T15:52:20.187Z"
                },
                "Fields": {
                    "bodyText": "### Flat Rate Shipping\n\nNo matter how many Unlimited Furniture Delivery items you place in a single order, we provide white glove service for a $149 flat-rate fee. This includes inside delivery to your room of choice, unpacking, assembly (except for chandeliers), and debris removal.\n\nOrders containing more than one Unlimited Furniture Delivery item will be shipped together and delivered at the same time. If you prefer to receive your items as they become available, please contact our furniture specialists after submitting your order at [844.599.2557](tel:844.599.2557) from 7 a.m. to 8 p.m. ET, Monday through Saturday or email [furniture@anthropologie.com](mailto:furniture@anthropologie.com).\n\nOnce your items are ready to ship, we will contact you to schedule a delivery. Furniture will arrive 2-4 weeks after your order has shipped. While you may see a preauthorization on your selected method of payment prior to delivery, your account will not be charged until the items are delivered.\n\nPlease note Unlimited Furniture Delivery items are unable to ship to US territories, Alaska, Hawaii, islands within the continental US, PO boxes, APO/FPO addresses and locations outside of the US. Overnight and express delivery options are also unavailable.",
                    "slug": "product_shipping-flat-rate",
                    "title": "PDP Flat Rate Shipping Copy"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "mobileReviewTemplateFieldModule",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2018-08-03T18:20:08.160Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "2isYBattJyEucWoMisao4o",
                    "locale": "en-US",
                    "revision": 1,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2018-08-03T18:20:08.160Z"
                },
                "Fields": {
                    "fieldName": "userlocation",
                    "fieldType": "floating",
                    "title": "User Location"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "mobileReviewTemplateFieldModule",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2018-08-03T18:25:02.876Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "2stKsabuWMSW68SqGwqCgM",
                    "locale": "en-US",
                    "revision": 1,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2018-08-03T18:25:02.876Z"
                },
                "Fields": {
                    "fieldName": "contextdatavalue_heightFeet",
                    "fieldType": "floating",
                    "title": "Height Feet"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "mobileReviewTemplateFieldModule",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2018-08-03T18:20:40.713Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "3313tdqFGUauiMKYMsueu2",
                    "locale": "en-US",
                    "revision": 1,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2018-08-03T18:20:40.713Z"
                },
                "Fields": {
                    "fieldName": "contextdatavalue_bodyType",
                    "fieldType": "floating",
                    "title": "Body Type"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "mobileReviewTemplateFieldModule",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2019-01-24T16:25:24.093Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "34O15Ci9jbHBkqrbm1PTGN",
                    "locale": "en-US",
                    "revision": 1,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2019-01-24T16:25:24.093Z"
                },
                "Fields": {
                    "fieldName": "contextdatavalue_FitPurchased",
                    "fieldType": "floating",
                    "title": "Fit Purchased"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "facetNameMapping",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2016-11-28T18:42:13.235Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "35ASYaHK5qAyOqKgOc2AKi",
                    "locale": "en-US",
                    "revision": 2,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2016-11-28T21:48:42.195Z"
                },
                "Fields": {
                    "groupByName": "tile.product.parentCategory.displayName",
                    "name": "division"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "facetNameMapping",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2016-11-18T21:19:30.437Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "3Dz7crZFUIeE6QscAKuuSo",
                    "locale": "en-US",
                    "revision": 2,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2017-10-18T19:42:41.177Z"
                },
                "Fields": {
                    "groupByName": "tile.skuInfo.primarySlice.sliceItems.colorways.name",
                    "name": "oldcolor"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "colorFacets",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2016-10-13T17:54:30.524Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "3Ge3MI5DIcaCSwY6ycSA6i",
                    "locale": "en-US",
                    "revision": 1,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2016-10-13T17:54:30.524Z"
                },
                "Fields": {
                    "color": "orange",
                    "hexCode": "fe6418"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "facetNameMapping",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2016-11-18T21:22:55.813Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "3JADy40qEE0sEGSMyMSi2y",
                    "locale": "en-US",
                    "revision": 1,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2016-11-18T21:22:55.813Z"
                },
                "Fields": {
                    "groupByName": "tile.product.keywords",
                    "name": "keywords"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "mobileReviewTemplateFieldModule",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2018-08-03T17:20:31.388Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "3KAoJxyX60WCgOsWqqgq2o",
                    "locale": "en-US",
                    "revision": 1,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2018-08-03T17:20:31.388Z"
                },
                "Fields": {
                    "fieldName": "rating_fitSlide",
                    "fieldType": "slider",
                    "title": "Slider"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "mobileReviewTemplateFieldModule",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2018-08-03T17:22:25.976Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "3KgKMBYaLmoYS0Usk8sEQm",
                    "locale": "en-US",
                    "revision": 2,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2018-08-03T17:22:55.172Z"
                },
                "Fields": {
                    "fieldName": "title",
                    "fieldType": "floating",
                    "title": "Title"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "mobileReviewTemplateFieldModule",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2018-08-03T18:21:06.989Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "3Nmi9D58re6sSC6EwWUYYy",
                    "locale": "en-US",
                    "revision": 1,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2018-08-03T18:21:06.989Z"
                },
                "Fields": {
                    "fieldName": "contextdatavalue_style",
                    "fieldType": "floating",
                    "title": "Style"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "mobileReviewTemplateContainer",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2018-08-03T17:23:49.715Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "3R8QX1U6jeqcMIsAkAe0W8",
                    "locale": "en-US",
                    "revision": 2,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2018-08-21T20:40:45.324Z"
                },
                "Fields": {
                    "headerTitle": "Share Your Thoughts With Others",
                    "referenceModules": [
                        {
                            "sys": {
                                "id": "3KgKMBYaLmoYS0Usk8sEQm",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "7tHoQeYdIQwo8SWc4Iqe4u",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        }
                    ],
                    "title": "Group 2 Template"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "mobileReviewTemplateFieldModule",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2018-08-03T17:20:08.346Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "3Stp1vCEFywEi4EkwCUscq",
                    "locale": "en-US",
                    "revision": 2,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2018-08-03T17:20:40.206Z"
                },
                "Fields": {
                    "fieldName": "rating",
                    "fieldType": "starRating",
                    "title": "Rating"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "facetNameMapping",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2016-11-18T21:20:10.229Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "3bCaTu6gEw6k4EGCqOSKCy",
                    "locale": "en-US",
                    "revision": 1,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2016-11-18T21:20:10.229Z"
                },
                "Fields": {
                    "groupByName": "tile.skuInfo.salePriceLow",
                    "name": "price"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "mobileFeatureToggles",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2019-02-14T18:17:51.032Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "3dMkkCLM6KSQIV6UIvytVN",
                    "locale": "en-US",
                    "revision": 3,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2019-02-27T20:24:00.166Z"
                },
                "Fields": {
                    "afterPayAndroidToggle": true,
                    "afterpayToggle": true,
                    "akamaiToggle": true,
                    "appleMusicToggle": true,
                    "categoryL3Toggle": true,
                    "employeeIdToggle": false,
                    "iovationToggle": true,
                    "marketplaceReturnsToggle": true,
                    "siteSpectToggle": true,
                    "sizeClassToggle": true,
                    "title": "Mobile Feature Toggles"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "mobileReviewTemplateFieldModule",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2018-08-03T18:19:45.807Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "3hWrSaDWM8KMsgWwMWy2g6",
                    "locale": "en-US",
                    "revision": 1,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2018-08-03T18:19:45.807Z"
                },
                "Fields": {
                    "fieldName": "usernickname",
                    "fieldType": "floating",
                    "title": "User Nickname"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "mobileReviewTemplateFieldModule",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2018-08-03T18:21:23.177Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "3iF0p7TwwoC28IUAwC2Wqw",
                    "locale": "en-US",
                    "revision": 1,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2018-08-03T18:21:23.177Z"
                },
                "Fields": {
                    "fieldName": "additionalfield_SizeNormallyWorn",
                    "fieldType": "floating",
                    "title": "Size Normally Worn"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "mobileReviewTemplateFieldModule",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2018-08-03T18:22:34.802Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "3j7o9WdNTOY28CMGGgeEiE",
                    "locale": "en-US",
                    "revision": 1,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2018-08-03T18:22:34.802Z"
                },
                "Fields": {
                    "fieldName": "contextdatavalue_EyeColor",
                    "fieldType": "floating",
                    "title": "Eye Color"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "colorFacets",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2016-10-13T17:54:32.465Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "3mKv4Z2euQqUuAIweyukuS",
                    "locale": "en-US",
                    "revision": 1,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2016-10-13T17:54:32.465Z"
                },
                "Fields": {
                    "color": "mint",
                    "displayName": "Mint",
                    "hexCode": "e2ede7"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "colorFacets",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2016-10-13T17:54:27.641Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "3seDREQNDyQuIKSKWOYQ0W",
                    "locale": "en-US",
                    "revision": 1,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2016-10-13T17:54:27.641Z"
                },
                "Fields": {
                    "color": "beige",
                    "displayName": "Beige",
                    "hexCode": "dcc2aa"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "facetNameMapping",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2017-08-29T18:35:57.524Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "4CnXCZ3NfOssouCQqWaGaW",
                    "locale": "en-US",
                    "revision": 2,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2017-10-18T19:42:44.337Z"
                },
                "Fields": {
                    "groupByName": "tile.product.facets.colors.colorways.name",
                    "name": "color"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "colorFacets",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2016-10-13T17:53:08.428Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "4EwmYFs0icMi0qaKAYKigK",
                    "locale": "en-US",
                    "revision": 1,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2016-10-13T17:53:08.428Z"
                },
                "Fields": {
                    "color": "pink/purple",
                    "hexCode": "ffe4e1"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "featureConfigMto",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2017-05-18T18:45:15.162Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "4NB5lEnp0ACmcm4aUS00uW",
                    "locale": "en-US",
                    "revision": 23,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2017-08-04T15:15:30.848Z"
                },
                "Fields": {
                    "brandCode": "01",
                    "goodRatingThreshhold": 3,
                    "heroImageMaxZoomScaleLarge": 2,
                    "heroImageMaxZoomScaleSmall": 4,
                    "maxInitialMaterialSets": 6,
                    "maxItemQuantity": 10,
                    "shippingContent": [
                        {
                            "sys": {
                                "id": "2Tp8q95upGQYAwOsgAUU28",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "25pMHCK1XaeIc8sM486oYQ",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "2DHTKcBkViciU8EKcM80aM",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        }
                    ],
                    "slug": "mto",
                    "title": "MTO Feature Config"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "mobileReviewTemplateFieldModule",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2018-08-03T18:24:03.113Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "4NDCbQLBBYuguq8wowEYG0",
                    "locale": "en-US",
                    "revision": 1,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2018-08-03T18:24:03.113Z"
                },
                "Fields": {
                    "fieldName": "contextdatavalue_HomePersonalStyle",
                    "fieldType": "floating",
                    "title": "Home Personal Style"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "mobileReviewTemplateFieldModule",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2018-08-03T18:21:55.415Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "4UJfqKW0E8ce00Mi0QYisa",
                    "locale": "en-US",
                    "revision": 1,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2018-08-03T18:21:55.415Z"
                },
                "Fields": {
                    "fieldName": "contextdatavalue_SkinTone",
                    "fieldType": "floating",
                    "title": "Skin Tone"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "mobileReviewTemplateFieldModule",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2018-08-03T18:27:54.348Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "4UhiS3uZ9KskCUKM6Oce8U",
                    "locale": "en-US",
                    "revision": 1,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2018-08-03T18:27:54.348Z"
                },
                "Fields": {
                    "fieldName": "contextdatavalue_SkinConcern",
                    "fieldType": "floating",
                    "title": "Skin Concern"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "facetNameMapping",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2016-11-18T21:21:33.736Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "4YNYblCJsQuO2gcwMCuWEE",
                    "locale": "en-US",
                    "revision": 2,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2016-11-18T21:21:50.979Z"
                },
                "Fields": {
                    "groupByName": "tile.reviews.averageRating",
                    "name": "reviews"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "colorFacets",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2016-10-13T17:54:31.394Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "4jL2AAHD2Eao2gcO8CSeyw",
                    "locale": "en-US",
                    "revision": 1,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2016-10-13T17:54:31.394Z"
                },
                "Fields": {
                    "color": "clear",
                    "displayName": "Clear",
                    "hexCode": "ececec"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "colorFacets",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2016-10-13T17:54:28.169Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "4sM9rCyiLC08g6se8miaqA",
                    "locale": "en-US",
                    "revision": 1,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2016-10-13T17:54:28.169Z"
                },
                "Fields": {
                    "color": "purple",
                    "displayName": "Purple",
                    "hexCode": "530066"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "mobileReviewTemplateContainer",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2018-08-03T17:16:49.365Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "52iz9NHTgsyiGw4oakoo2E",
                    "locale": "en-US",
                    "revision": 8,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2019-01-09T19:11:56.724Z"
                },
                "Fields": {
                    "headerTitle": "Rate this product",
                    "referenceModules": [
                        {
                            "sys": {
                                "id": "3Stp1vCEFywEi4EkwCUscq",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "3KAoJxyX60WCgOsWqqgq2o",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "6gfgZlLJSwaiiiyEaWU0W6",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        }
                    ],
                    "title": "Group 1 Template"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "mobileReviewTemplateFieldModule",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2018-08-03T18:18:55.323Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "55WtUtTg0oqUceG4wiqC4W",
                    "locale": "en-US",
                    "revision": 1,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2018-08-03T18:18:55.323Z"
                },
                "Fields": {
                    "fieldName": "imageUploader",
                    "fieldType": "image",
                    "title": "Image Uploader"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "mobileReviewTemplateFieldModule",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2018-08-03T18:22:22.994Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "5YVPIALjLaO4ymSmYu044G",
                    "locale": "en-US",
                    "revision": 1,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2018-08-03T18:22:22.994Z"
                },
                "Fields": {
                    "fieldName": "contextdatavalue_BeautyStyle",
                    "fieldType": "floating",
                    "title": "Beauty Style"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "mobileReviewTemplateFieldModule",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2018-08-09T20:40:46.107Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "5cgIIFlMHCGiGqKY8CgEK0",
                    "locale": "en-US",
                    "revision": 1,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2018-08-09T20:40:46.107Z"
                },
                "Fields": {
                    "fieldName": "contextdatavalue_HairType",
                    "fieldType": "floating",
                    "title": "Hair Type"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "mobileReviewTemplateFieldModule",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2018-08-03T18:24:49.605Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "5miErEaVMcQ2AuaocumEyY",
                    "locale": "en-US",
                    "revision": 1,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2018-08-03T18:24:49.605Z"
                },
                "Fields": {
                    "fieldName": "contextdatavalue_heightInches",
                    "fieldType": "floating",
                    "title": "Height Inches"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "mobileFeatureCatalogContainer",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2018-01-15T21:34:57.357Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "5qNMCipb4QC6AiIQuMo8iC",
                    "locale": "en-US",
                    "revision": 40,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2019-02-14T18:34:58.702Z"
                },
                "Fields": {
                    "afterpayToggle": true,
                    "akamaiToggle": true,
                    "applePayToggle": true,
                    "blacklistTrendingTerms": [
                        "ping pong",
                        "coachella",
                        "cloth stone"
                    ],
                    "categoryL3Toggle": true,
                    "colorFacets": [
                        {
                            "sys": {
                                "id": "2QGjZek2VOmEyOEAAWKi0M",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "1M7RcVhYRi2YUE0eWcK4iy",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "3seDREQNDyQuIKSKWOYQ0W",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "6HEI2JdrmEw8IAgoc4gsco",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "MhHRuFZcAKkWwm6CMQkoY",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "4jL2AAHD2Eao2gcO8CSeyw",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "2LshNHGmPuAsqM8oYWcqeq",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "2K2TkOFH6MmA8oaeqw8ccQ",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "1oTiIQhHiwYi4eC2i6UUOW",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "6Mrp9RvgtiAcIw0iSyWYYk",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "6RnURl0XcsMSkY8sgAugk2",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "5vMvozzMsMomOQIqyso68E",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "4sM9rCyiLC08g6se8miaqA",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "4EwmYFs0icMi0qaKAYKigK",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "6oaf1Dh2MMQQCOm6Uagc4w",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "3Ge3MI5DIcaCSwY6ycSA6i",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "10RSc6oVqEKMiEQiaQo0qk",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "3mKv4Z2euQqUuAIweyukuS",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "1muYaRWPVWEcEkgwWuYiOk",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "231lIJOzz6oYWAUsuquSsc",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "5xWaNQxu6sU0yaIie0wqK4",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        }
                    ],
                    "curalateToggle": true,
                    "dynamicYieldProductIdentifier": "productId",
                    "facetNameMappings": [
                        {
                            "sys": {
                                "id": "1NGk4d67OkmuWWi2eCWyGy",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "4CnXCZ3NfOssouCQqWaGaW",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "3JADy40qEE0sEGSMyMSi2y",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "35ASYaHK5qAyOqKgOc2AKi",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "4YNYblCJsQuO2gcwMCuWEE",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "3Dz7crZFUIeE6QscAKuuSo",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "3bCaTu6gEw6k4EGCqOSKCy",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "6KyCGGzjc408G6ACikmygy",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        }
                    ],
                    "featureConfig": {
                        "sys": {
                            "id": "6Uaf1PTzr1sn5MF4hgccng",
                            "linkType": "Entry",
                            "type": "Link"
                        }
                    },
                    "featureToggles": {
                        "sys": {
                            "id": "3dMkkCLM6KSQIV6UIvytVN",
                            "linkType": "Entry",
                            "type": "Link"
                        }
                    },
                    "giftWrapImageCode": "b",
                    "imageCacheOptionToggle": false,
                    "iovationToggle": true,
                    "lowMarginToggle": true,
                    "lowStockThreshold": 5,
                    "mtoQuickShipSort": "stockLevel",
                    "reviewTemplateContainer": [
                        {
                            "sys": {
                                "id": "52iz9NHTgsyiGw4oakoo2E",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "3R8QX1U6jeqcMIsAkAe0W8",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "ydAnrQ8jBuGOEwqw4ss0q",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "1ZrzOjS0pKW0KQYqakKQOK",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        },
                        {
                            "sys": {
                                "id": "62paV3Lb1Kyy0AQccCy4Kc",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        }
                    ],
                    "siteSpectToggle": false,
                    "sizeClassToggle": false,
                    "slug": "mobileBrowseDisplayOptions",
                    "title": "Mobile Browse Display Options",
                    "topRatedThreshold": 4.5,
                    "ugcHashtags": [
                        "Anthropologie",
                        "AnthroStyle",
                        "TheWayIDress",
                        "AnthroHome",
                        "AnthroBeauty",
                        "50StatesOfDenim"
                    ]
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "colorFacets",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2016-10-13T17:54:30.802Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "5vMvozzMsMomOQIqyso68E",
                    "locale": "en-US",
                    "revision": 3,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2017-11-13T21:12:05.130Z"
                },
                "Fields": {
                    "color": "red",
                    "hexCode": "d00010",
                    "swatchUrl": null
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "colorFacets",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2016-10-13T17:54:29.138Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "5xWaNQxu6sU0yaIie0wqK4",
                    "locale": "en-US",
                    "revision": 1,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2016-10-13T17:54:29.138Z"
                },
                "Fields": {
                    "color": "brown",
                    "hexCode": "5b4334"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "mobileReviewTemplateFieldModule",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2018-08-09T20:40:10.293Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "5yi00B2Kic66os086seKsO",
                    "locale": "en-US",
                    "revision": 1,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2018-08-09T20:40:10.293Z"
                },
                "Fields": {
                    "fieldName": "contextdatavalue_HairConcern",
                    "fieldType": "floating",
                    "title": "Hair Concern"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "mobileReviewTemplateContainer",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2018-10-02T20:36:00.556Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "62paV3Lb1Kyy0AQccCy4Kc",
                    "locale": "en-US",
                    "revision": 1,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2018-10-02T20:36:00.556Z"
                },
                "Fields": {
                    "referenceModules": [
                        {
                            "sys": {
                                "id": "1Ebej9lTMkMoWaoII0CWY0",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        }
                    ],
                    "title": "Group 5 Template"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "colorFacets",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2016-10-13T17:54:05.125Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "6HEI2JdrmEw8IAgoc4gsco",
                    "locale": "en-US",
                    "revision": 1,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2016-10-13T17:54:05.125Z"
                },
                "Fields": {
                    "color": "black",
                    "displayName": "Black",
                    "hexCode": "000"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "facetNameMapping",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2016-11-18T21:19:03.131Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "6KyCGGzjc408G6ACikmygy",
                    "locale": "en-US",
                    "revision": 1,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2016-11-18T21:19:03.131Z"
                },
                "Fields": {
                    "groupByName": "tile.skuInfo.secondarySlice.sliceItems.includedSizes.displayName",
                    "name": "size"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "colorFacets",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2016-10-13T17:53:07.050Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "6Mrp9RvgtiAcIw0iSyWYYk",
                    "locale": "en-US",
                    "revision": 1,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2016-10-13T17:53:07.050Z"
                },
                "Fields": {
                    "color": "tan",
                    "hexCode": "dd7e2e"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "colorFacets",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2016-10-13T17:54:28.505Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "6RnURl0XcsMSkY8sgAugk2",
                    "locale": "en-US",
                    "revision": 1,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2016-10-13T17:54:28.505Z"
                },
                "Fields": {
                    "color": "silver",
                    "hexCode": "cfcfcf"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "mobileFeatureConfig",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2019-02-14T18:31:11.598Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "6Uaf1PTzr1sn5MF4hgccng",
                    "locale": "en-US",
                    "revision": 3,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2019-02-18T15:18:59.254Z"
                },
                "Fields": {
                    "appleMusic": {
                        "sys": {
                            "id": "2KnKQZnaSFTT1WVaXwRH7D",
                            "linkType": "Entry",
                            "type": "Link"
                        }
                    },
                    "catalog": {
                        "sys": {
                            "id": "1ZnccGS16oASut7WZ1aGNq",
                            "linkType": "Entry",
                            "type": "Link"
                        }
                    },
                    "mto": {
                        "sys": {
                            "id": "4NB5lEnp0ACmcm4aUS00uW",
                            "linkType": "Entry",
                            "type": "Link"
                        }
                    },
                    "title": "Mobile Feature Config"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "mobileReviewTemplateFieldModule",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2018-08-03T17:21:07.047Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "6gfgZlLJSwaiiiyEaWU0W6",
                    "locale": "en-US",
                    "revision": 3,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2018-08-09T14:23:25.279Z"
                },
                "Fields": {
                    "fieldName": "isrecommended",
                    "fieldType": "button",
                    "title": "Is Recommended"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "colorFacets",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2016-10-13T17:54:29.886Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "6oaf1Dh2MMQQCOm6Uagc4w",
                    "locale": "en-US",
                    "revision": 1,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2016-10-13T17:54:29.886Z"
                },
                "Fields": {
                    "color": "pink",
                    "hexCode": "ffa9b2"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "mobileReviewTemplateFieldModule",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2018-08-03T18:23:33.238Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "6suyexivTyyoC4Uke0Ke6A",
                    "locale": "en-US",
                    "revision": 1,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2018-08-03T18:23:33.238Z"
                },
                "Fields": {
                    "fieldName": "contextdatavalue_HomeType",
                    "fieldType": "floating",
                    "title": "Home Type"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "mobileReviewTemplateFieldModule",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2018-08-03T18:24:18.785Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "6tWumz5SVOaiceEcKyG6Yg",
                    "locale": "en-US",
                    "revision": 1,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2018-08-03T18:24:18.785Z"
                },
                "Fields": {
                    "fieldName": "contextdatavalue_age",
                    "fieldType": "floating",
                    "title": "age"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "mobileReviewTemplateFieldModule",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2018-08-03T18:22:10.445Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "6wO4NM12w0UYaCyiMse0AC",
                    "locale": "en-US",
                    "revision": 1,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2018-08-03T18:22:10.445Z"
                },
                "Fields": {
                    "fieldName": "contextdatavalue_SkinType",
                    "fieldType": "floating",
                    "title": "Skin Type"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "mobileReviewTemplateFieldModule",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2018-08-03T17:23:20.817Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "7tHoQeYdIQwo8SWc4Iqe4u",
                    "locale": "en-US",
                    "revision": 2,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2018-08-15T17:14:15.125Z"
                },
                "Fields": {
                    "fieldName": "reviewtext",
                    "fieldType": "multiline",
                    "title": "Review Text"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "mobileReviewTemplateFieldModule",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2018-08-03T18:23:49.657Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "F2ErF2XRx6aO4Kw6YCKMo",
                    "locale": "en-US",
                    "revision": 1,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2018-08-03T18:23:49.657Z"
                },
                "Fields": {
                    "fieldName": "contextdatavalue_RoomType",
                    "fieldType": "floating",
                    "title": "Room Type"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "mobileReviewTemplateFieldModule",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2018-08-03T18:23:00.479Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "MWy1hHcx4ksQsmwqACw4U",
                    "locale": "en-US",
                    "revision": 2,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2018-08-03T18:23:17.260Z"
                },
                "Fields": {
                    "fieldName": "contextdatavalue_HairTexture",
                    "fieldType": "floating",
                    "title": "Hair Texture"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "colorFacets",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2016-10-13T17:54:28.156Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "MhHRuFZcAKkWwm6CMQkoY",
                    "locale": "en-US",
                    "revision": 1,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2016-10-13T17:54:28.156Z"
                },
                "Fields": {
                    "color": "blue",
                    "displayName": "Blue",
                    "hexCode": "014d98"
                }
            },
            {
                "Sys": {
                    "contentType": {
                        "sys": {
                            "id": "mobileReviewTemplateContainer",
                            "linkType": "ContentType",
                            "type": "Link"
                        }
                    },
                    "createdAt": "2018-08-03T18:18:59.461Z",
                    "environment": {
                        "sys": {
                            "id": "master",
                            "linkType": "Environment",
                            "type": "Link"
                        }
                    },
                    "id": "ydAnrQ8jBuGOEwqw4ss0q",
                    "locale": "en-US",
                    "revision": 1,
                    "space": {
                        "sys": {
                            "id": "8zamwvziwka5",
                            "linkType": "Space",
                            "type": "Link"
                        }
                    },
                    "type": "Entry",
                    "updatedAt": "2018-08-03T18:18:59.461Z"
                },
                "Fields": {
                    "headerTitle": "Upload Photos",
                    "referenceModules": [
                        {
                            "sys": {
                                "id": "55WtUtTg0oqUceG4wiqC4W",
                                "linkType": "Entry",
                                "type": "Link"
                            }
                        }
                    ],
                    "title": "Group 3 Template"
                }
            }
        ]
    },
    "items": [
        {
            "Sys": {
                "contentType": {
                    "sys": {
                        "id": "appConfig",
                        "linkType": "ContentType",
                        "type": "Link"
                    }
                },
                "createdAt": "2017-11-13T18:56:53.820Z",
                "environment": {
                    "sys": {
                        "id": "master",
                        "linkType": "Environment",
                        "type": "Link"
                    }
                },
                "id": "4o8oXj8yHKqAOM8yCMWeUu",
                "locale": "en-US",
                "revision": 10,
                "space": {
                    "sys": {
                        "id": "8zamwvziwka5",
                        "linkType": "Space",
                        "type": "Link"
                    }
                },
                "type": "Entry",
                "updatedAt": "2018-03-08T15:27:37.698Z"
            },
            "Fields": {
                "features": {
                    "sys": {
                        "id": "5qNMCipb4QC6AiIQuMo8iC",
                        "linkType": "Entry",
                        "type": "Link"
                    }
                },
                "mtoFeatureConfig": {
                    "sys": {
                        "id": "4NB5lEnp0ACmcm4aUS00uW",
                        "linkType": "Entry",
                        "type": "Link"
                    }
                },
                "settings": {
                    "sys": {
                        "id": "1Ra1cJWxaUeCGo4qm6kAiM",
                        "linkType": "Entry",
                        "type": "Link"
                    }
                },
                "title": "UIOS Default App Config"
            }
        }
    ],
    "limit": 100,
    "skip": 0,
    "total": 1,
    "Errs": null
}
`

type Entry struct {
	Sys    map[string]interface{}
	Fields map[string]interface{}
}

type Response struct {
	Includes struct {
		Asset interface{} `json:"Asset"`
		Entry []struct {
			Sys struct {
				ContentType struct {
					Sys struct {
						ID       string `json:"id"`
						LinkType string `json:"linkType"`
						Type     string `json:"type"`
					} `json:"sys"`
				} `json:"contentType"`
				CreatedAt   time.Time `json:"createdAt"`
				Environment struct {
					Sys struct {
						ID       string `json:"id"`
						LinkType string `json:"linkType"`
						Type     string `json:"type"`
					} `json:"sys"`
				} `json:"environment"`
				ID       string `json:"id"`
				Locale   string `json:"locale"`
				Revision int    `json:"revision"`
				Space    struct {
					Sys struct {
						ID       string `json:"id"`
						LinkType string `json:"linkType"`
						Type     string `json:"type"`
					} `json:"sys"`
				} `json:"space"`
				Type      string    `json:"type"`
				UpdatedAt time.Time `json:"updatedAt"`
			} `json:"Sys"`
			Fields map[string]interface{} `json:"Fields"`
		} `json:"Entry"`
	} `json:"includes"`
	Items []struct {
		Sys struct {
			ContentType struct {
				Sys struct {
					ID       string `json:"id"`
					LinkType string `json:"linkType"`
					Type     string `json:"type"`
				} `json:"sys"`
			} `json:"contentType"`
			CreatedAt   time.Time `json:"createdAt"`
			Environment struct {
				Sys struct {
					ID       string `json:"id"`
					LinkType string `json:"linkType"`
					Type     string `json:"type"`
				} `json:"sys"`
			} `json:"environment"`
			ID       string `json:"id"`
			Locale   string `json:"locale"`
			Revision int    `json:"revision"`
			Space    struct {
				Sys struct {
					ID       string `json:"id"`
					LinkType string `json:"linkType"`
					Type     string `json:"type"`
				} `json:"sys"`
			} `json:"space"`
			Type      string    `json:"type"`
			UpdatedAt time.Time `json:"updatedAt"`
		} `json:"Sys"`
		Fields struct {
			Features struct {
				Sys struct {
					ID       string `json:"id"`
					LinkType string `json:"linkType"`
					Type     string `json:"type"`
				} `json:"sys"`
			} `json:"features"`
			MtoFeatureConfig struct {
				Sys struct {
					ID       string `json:"id"`
					LinkType string `json:"linkType"`
					Type     string `json:"type"`
				} `json:"sys"`
			} `json:"mtoFeatureConfig"`
			Settings struct {
				Sys struct {
					ID       string `json:"id"`
					LinkType string `json:"linkType"`
					Type     string `json:"type"`
				} `json:"sys"`
			} `json:"settings"`
			Title string `json:"title"`
		} `json:"Fields"`
	} `json:"items"`
	Limit int         `json:"limit"`
	Skip  int         `json:"skip"`
	Total int         `json:"total"`
	Errs  interface{} `json:"Errs"`
}

func main() {
	var resp Response
	json.Unmarshal([]byte(jsonStr), &resp)
	for x := 0; x < len(resp.Includes.Entry); x++ {
		fmt.Printf("%d: %+v -- %+v\n", x+1, resp.Includes.Entry[x].Sys.ContentType.Sys, resp.Includes.Entry[x].Fields)
	}
}
