# Gateway

*This project only for personal use*.

### Features

* Get bing daily background with CN and US region and can specific return type with json or redirect
* Get top list info with weibo,zhihu,36kr
* Get latest news with wall street cn

### Example

```shell
# Use [random] to alternately get Bing's background images in CN and US.
# You can replace [random] with [cn] or [us] to get specific region's background
# Use [json] to get json data format response.
# Use [redirect] replace [json] to redirect to the image.
curl --location 'https://noahamethyst.club/gateway/spider/bing/wallpaper/random/json'
```

Json Response like this

```json
{
  "data": {
    "imageUrl": "http://www.bing.com/th?id=OHR.HelloSeal_ZH-CN1064568368_1920x1080.jpg&rf=LaDigue_1920x1080.jpg&pid=hp"
  },
  "message": ""
}
```



