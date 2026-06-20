 *DeadLink*
## یک ابزار ساده برای چک کردن اینکه توی فایل لینک هاش زنده ان یا خیر 
---
### توی این پروسه چه اتفاقی میوفته ؟ 
- مرحله اول : فایلی که بهش گفتیمو میخونه 
- مرحله دوم : لینک ها رو استخراج میکنه
- مرحله سوم : ریکوئست میدیم ببینیم زندست یا نه
- مرحله چهارم : میگه کدوم زنده ان یا مردن
---
## چطور استفاده کنم ؟ 
```
go run main.go
```
---
## نمونه خروجی : 
```
example : example.txt
test.pp
---(test.pp)---
is alive : google.com
-----
error : timeout.com
-----
is alive : pypi.org
-----
Total: 3 | Alive: 2 | Dead: 1 | Forbidden: 0
```
---
نویسنده : [RealFinalTry](https://github.com/RealFinalTry/)
