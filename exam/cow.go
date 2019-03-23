package main

import (
	"fmt"
	"runtime"
)

/**
 * 牛牛算法
 * jqk10，表示10牛，共5张牌，其中3张之和为10，,20,30，剩余之和与10取余就是多少牛
 * @author Administrator
 *
 */
public class Taurus {

 var temp,n,cardsTotal int//n表示jqk 10的个数，cow表示牛，cardsTotal表示牌总和
 int cow = -1;//默认没有牛
 card :=int[5]//5个牌
private int[] random = new int[5];//5个随机数
//模拟52张牌


/*//随机产生5张牌
@Test
public int[] getCard(){
Random rd = new Random();
for(int i=0;i<5;i++){
temp = random[i] = rd.nextInt(52);//随机生成0到51个int
for(int j=0;j<i;j++){
if(temp == random[j]) continue;//重复则重新随机
//				if(temp == 52) continue;//数组下标从0开始,只有51
}
}
System.out.println("随机数:"+Arrays.toString(random));*/
//取牌
for(int i=0;i<random.length;i++){
card[i] = cards [random[i]];
if(card[i]>=10) n++;
}
//给card排序
Arrays.sort(card);
//倒序
int[] card1 = new int[5];
for(int i=0;i<card.length ;i++){
card1[i] = card[4-i];
}
print(card1);
return card1;
}

//判断牛几方法
func getCow(card []int){
	var n,cardsTotal int
	cow := -1
	for i:=0;i<len(card);i++{
		cardsTotal +=card[i]
}
	fmt.Println("cardsTatoal=",cardsTotal)
//根据jqk10个数分情况讨论
	switch(n){
		case 0://一张都没用
			for i:=0;i<4;i++ {
				for j:=1;j<5;j++ {
					if((cardsTotal-card[i]-card[j])%10==0){
						cow = (card[i]+card[j])%10
							}
				}
			}
					return cow

		case 1://有一张jqk10的情况,剩余4张中有3张之和要被10整除
			for j:=0;j<5;j++ {
				if((cardsTotal-card[0]-card[j])%10==0){
					cow = (cardsTotal-card[0])%10
						}
				}
					return cow

		case 2://有2张jqk10的情况,剩余3张中有2张之和等于10或者剩余3张之和为10
			for i:=n;i<4;i++ {
				for j:=n+1;j<5;j++ {
					if((card[i]+card[j])==10 ||(cardsTotal-card[i]-card[j])%10 == 0){
						for k:=0;k<n;k++{
							cardsTotal-=card[k]
										}
								cow = cardsTotal%10
							}
					}

				}
					return cow
		default://有3,4,5张jqk10的情况

			for i:=0;i<n;i++{
				cardsTotal -= card[i]
					}
					cow = cardsTotal%10
					return cow

					runtime.GOMAXPROCS()
			}
	}


