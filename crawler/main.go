package main

import (
	"LearGoProject/crawler/engine"
	"LearGoProject/crawler/zhenai/parser"
)

func main() {

	engine.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})




}
	/*resp,err:=http.Get("http://www.zhenai.com/zhenghun")
	if err!=nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode!=http.StatusOK {
		fmt.Println("Error:status code",resp.StatusCode)
		return
	}

	e:=determineEncoding(resp.Body)

	utf8Reader:= transform.NewReader(resp.Body,e.NewDecoder())

	all,err:=ioutil.ReadAll(utf8Reader)
	if err!=nil {
		panic(err)
	}
	printCityList(all)
}

func printCityList(contents []byte)  {

	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`)
	matches := re.FindAllSubmatch(contents, -1)

	for _,m:=range matches{

			fmt.Printf("City:%s, URL:%s\n",m[2],m[1])
		}

	fmt.Printf("matches found:%d\n",len(matches))
}

func determineEncoding(r io.Reader) encoding.Encoding  {
	bytes,err:=bufio.NewReader(r).Peek(1024)
	if err!=nil {
		panic(err)
	}
	e, _,_ := charset.DetermineEncoding(bytes, "")

	return e
}*/