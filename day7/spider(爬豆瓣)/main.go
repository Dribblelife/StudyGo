package main

import (
	"time"
	"fmt"
	"os"
	"strconv"
	"net/http"
	"io/ioutil"
	"regexp"
)

//定义 新的数据类型
type Spider struct {
	url string
	header map[string]string
}

//定义Spider get方法
func (keyword Spider) get_html_header() string  {
	client:=&http.Client{}
	println(keyword.url)
	req,err:=http.NewRequest("GET",keyword.url,nil)
	if err!=nil {
		fmt.Println(err)
	}
	for key,value:=range keyword.header {
		req.Header.Add(key,value)

	}
	respon,err:=client.Do(req)
	if err!=nil {
		fmt.Println(err)
	}
	defer respon.Body.Close()
	body,err:=ioutil.ReadAll(respon.Body)
	if err!=nil {
		fmt.Println(err)
	}
	return string(body)

}

func parse()  {
	header:=map[string]string{
		"Host":"movie.douban.com",
		"Connection":"keep-alive",
		"Cache-Control":"max-age=0",
		"Upgrade-Insecure-Requests":"1",
		"User-Agent":"Mozilla/5.0 (Windows NT 6.3; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.94 Safari/537.36",
		"Accept":"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8",
		"Referer":"https://movie.douban.com/cinema/nowplaying/xian/",
	}
	//创建excel文件
	f,err:=os.Create("F:/GoWorkSpace/src/LearGoProject/day7/spider(爬豆瓣)/豆瓣爬虫1.xlsx")

	if err!=nil {
		panic(err)
	}
	defer f.Close()
	//写入标题
	f.WriteString("电影名称"+"\t"+"简介"+"\t"+"评分"+"\t"+"评价人数"+"\t"+"\r\n")
	//循环每页解析并把结果写入excel
	for i:=0;i<2 ;i++  {
		fmt.Println("正在抓取第"+strconv.Itoa(i)+"页。。。")
		url:="https://movie.douban.com/top250?start="+strconv.Itoa(i*25)+"&filter="
		spider:=&Spider{url,header}
		html:=spider.get_html_header()

		//print HTML

		//电影名称
		pattern2:=`img width="100" alt="(.*?)" src=`
		rp2:=regexp.MustCompile(pattern2)
		find_txt2:=rp2.FindAllStringSubmatch(html,-1)
		//简介
		pattern3:=`<span class="inq">(.*?)</span>`
		rp3:=regexp.MustCompile(pattern3)
		find_txt3:=rp3.FindAllStringSubmatch(html,-1)
		//评分
		pattern4:=`property="v:average">(.*?)</span>`
		rp4:=regexp.MustCompile(pattern4)
		find_txt4:=rp4.FindAllStringSubmatch(html,-1)
		//评价人数
		pattern5:=`<span>(.*?)评价</span>`
		rp5:=regexp.MustCompile(pattern5)
		find_txt5:=rp5.FindAllStringSubmatch(html,-1)


		//写入utf-8 BOM
		f.WriteString("\xEF\xBB\xBF")		////写入UTF-8 BOM,此处如果不写入就会导致写入的汉字乱码
		for i:=0;i<len(find_txt2) ;i++  {
			fmt.Printf("%s %s %s %s\n",find_txt2[i][1],find_txt3[i][1],find_txt4[i][1],find_txt5[i][1])
			f.WriteString(find_txt2[i][1]+"\t"+find_txt3[i][1]+"\t"+find_txt4[i][1]+"\t"+find_txt5[i][1]+"\t"+"\r\n")
		}

	}



}

func main() {
	t1:=time.Now()	//get current time
	parse()
	elapsed:=time.Since(t1)
	fmt.Println("爬虫结束，总共耗时：",elapsed)
}
