package parser

import (
	"LearGoProject/crawler/engine"
	"regexp"
	"strconv"
	"LearGoProject/crawler/model"
)

var ageRe = regexp.MustCompile(
	`<td><span class="label">年龄：</span>([\d]+)岁</td>`)

/*var heightRe = regexp.MustCompile(
	`<div class="m-btn purple" data-v-bff6f798="">([\d]+)cm</div>`)*/

/*var incomeRe = regexp.MustCompile(
	`<div class="m-btn purple" data-v-bff6f798="">([^<]+)</div>`)*/

/*var weightRe = regexp.MustCompile(
	`<div class="m-btn purple" data-v-bff6f798="">([\d]+)kg</div>`)*/


/*var educationRe = regexp.MustCompile(
	`<div class="m-btn purple" data-v-bff6f798="">([^<]+)</div>`)

var occupationRe = regexp.MustCompile(
	`<div class="m-btn purple" data-v-bff6f798="">([^<]+)</div>`)

var hukouRe = regexp.MustCompile(
	`<div class="m-btn purple" data-v-bff6f798="">([^<]+)</div>`)

var houseRe =regexp.MustCompile(
	`<div class="m-btn pink" data-v-bff6f798="">([^<]+)</div>`)

var carRe = regexp.MustCompile(
	`<div class="m-btn pink" data-v-bff6f798="">([^<]+)</div>`)

var marrigeRe  = regexp.MustCompile(
	`<div class="m-btn purple" data-v-bff6f798="">([^<]+)</div>`)*/

func ParseProfile(contents []byte,
	name string) engine.ParseResult {

	profile:=model.Profile{}

	profile.Name=name

	age,err:=strconv.Atoi(
		extractString(contents,ageRe))
	if err!=nil {
		profile.Age=age
	}
	/*height,err:=strconv.Atoi(
		extractString(contents,heightRe))
	if err!=nil {
		profile.Height=height
	}
	weight,err:=strconv.Atoi(
		extractString(contents,weightRe))
	if err!=nil {
		profile.Weight=weight
	}*/

/*	profile.Income=extractString(contents,incomeRe)

	profile.Car=extractString(contents,carRe)
	profile.Education=extractString(contents,educationRe)
	profile.Marriage=extractString(contents,marrigeRe)
	profile.Occupation=extractString(contents,occupationRe)
	profile.Hukou=extractString(contents,hukouRe)*/

	result:=engine.ParseResult{
		Items:[]interface{} {profile,},
	}

	return result
	//re:=regexp.MustCompile(ageRe)
	/*match:= ageRe.FindSubmatch(contents)

	if match!=nil {
		age,err:=strconv.Atoi(string(match[1]))
		if err!=nil {
			profile.Age=age
		}

	}

	//re=regexp.MustCompile(marrigeRe)
	match=marrigerRe.FindSubmatch(contents)

	if match!=nil {
		profile.Marriage=string(match[1])*/

}

func extractString(contents []byte,re *regexp.Regexp) string {

	match:=re.FindSubmatch(contents)

	if len(match)>=2 {
		return string(match[1])
	}else {
		return ""
	}


}