package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"image/gif"
	"log"
	"os"
)


/*================================== 地图类 ==============================*/
type Map struct{
	SquareLen	int			//地图格子大小
	Entinexty	[5][4]int	//地图数据
	MapImage *image.Paletted
}

//绘制地图
func (m *Map) DrawTheMap(){
	for x := 0; x < 4; x++ {
		for y := 0; y < 5; y++ {
			if m.Entinexty[y][x] == 1 {
				rect := image.Rect(x * m.SquareLen,y * m.SquareLen,(x+1) * m.SquareLen,(y+1) * m.SquareLen)
				DrawRectangle(m.MapImage,rect,4)
			} else if m.Entinexty[y][x] == 2 {
				rect := image.Rect(x * m.SquareLen,y * m.SquareLen,(x+1) * m.SquareLen,(y+1) * m.SquareLen)
				DrawRectangle(m.MapImage,rect,2)
			} else if m.Entinexty[y][x] == 3 {
				rect := image.Rect(x * m.SquareLen,y * m.SquareLen,(x+1) * m.SquareLen,(y+1) * m.SquareLen)
				DrawRectangle(m.MapImage,rect,1)
			}
		}
	}
}
/*============================= end =======================================*/

/*=========================== gif图绘制所需 ============================*/
//创建调色板
var palette = []color.Color{
	color.RGBA{255,255,255,255},//底板色彩
	color.RGBA{0,255,0,255},//公主色彩
	color.RGBA{0,0,255,255},//已通过路径色彩
	color.RGBA{255,0,0,255},//回退点色彩
	color.RGBA{177,177,0,255},//障碍物色彩
	color.RGBA{0,255,255,255},//头部色彩
}

var anim = gif.GIF{LoopCount:510}	//创建gif图像，并设置动画帧数为510

var forwardImg  [255]Map//前进路径示意图临时存档数组
var rollbackImg	 [255]Map//回退路径示意图临时存放数组
var forwardCounter int = 0//前进路径示意图计数
var rollbackCounter int = 0	//回退路径示意图计数


//矩形绘制
func DrawRectangle(img *image.Paletted,rect image.Rectangle,color uint8){
	for x:= rect.Min.X; x < rect.Max.X; x++ {
		for y := rect.Min.Y; y < rect.Max.Y; y++ {
			img.SetColorIndex(x,y,color)
		}
	}
}

/*============================ end ======================================*/


/*========================= 骑士救公主 ^v^ ===============================*/
//地图内容
var MapArr = [5][4]int{	//0: 可通过；1： 障碍物； 2：骑士； 3：公主
	{0,0,1,0,},
	{0,0,0,0,},
	{0,0,1,0,},
	{0,1,3,0,},
	{0,0,0,1,},
}

//行走行为
var next = [4][2]int{
	{0,1},//向右走
	{1,0},//向下走
	{0,-1},//向左走
	{-1,0},//向上走
}

// 深度优先搜索算法实现
var min int = 99999		//最短路径
var book [5][4]int		//已行进坐标记录

func dfs(m *Map , x int, y int, step int, counter int) {
	var nextx,nexty int 
	
	if m.Entinexty[x][y] == 3 {
		if step < min {
			min = step
		} 
		return 
	}

	//枚举4种走法
	for k := 0; k <= 3; k++ {
		nextx = x+next[k][0]
		nexty = y+next[k][1]
		if nextx < 0 || nextx >= 5 || nexty < 0 || nexty >= 4 {	//越界
			continue
		}
		
		/*==== 分别为前进、后退绘制相应存图 =====*/
		forwardImg[forwardCounter].SquareLen = 60
		forwardImg[forwardCounter].MapImage = image.NewPaletted(image.Rect(0, 0, 240, 300),palette)
		forwardImg[forwardCounter].Entinexty = MapArr	
		forwardImg[forwardCounter].DrawTheMap()

		rollbackImg[forwardCounter].SquareLen = 60
		rollbackImg[forwardCounter].MapImage = image.NewPaletted(image.Rect(0, 0, 240, 300),palette)
		rollbackImg[forwardCounter].Entinexty = MapArr	
		rollbackImg[forwardCounter].DrawTheMap()


		for _x := 0; _x < 4; _x++ {
			for _y := 0; _y < 5; _y++ {
				if ( book[_y][_x] == 1 ){
					rect := image.Rect(_x * m.SquareLen, _y * m.SquareLen, (_x+1) * m.SquareLen, (_y+1) * m.SquareLen)
					DrawRectangle(forwardImg[forwardCounter].MapImage,rect,2)	
				}
			}
		}
		/*=====  end  =====*/

		if m.Entinexty[nextx][nexty] == 0 && book[nextx][nexty] == 0 || 
									m.Entinexty[nextx][nexty] == 3 {
			book[nextx][nexty] = 1	//记录当前路径

			/*==== 绘制前进图像 =====*/
			rect := image.Rect(nexty * m.SquareLen, nextx*m.SquareLen, (nexty+1) * m.SquareLen, (nextx+1) * m.SquareLen)
			DrawRectangle(forwardImg[forwardCounter].MapImage,rect,5)	
			anim.Image = append(anim.Image,forwardImg[forwardCounter].MapImage)
			forwardCounter += 1
			anim.Delay = append(anim.Delay,60)//加入gif图像,设置帧间间隔60ms
			/*===== end =====*/

			dfs(m,nextx,nexty,step+1,forwardCounter)
			
			/*======= 绘制回退图像 =======*/
			for _x := 0; _x < 4; _x++ {
				for _y := 0; _y < 5; _y++ {
					if ( book[_y][_x] == 1 ){
						rect := image.Rect(_x * m.SquareLen, _y * m.SquareLen, (_x+1) * m.SquareLen, (_y+1) * m.SquareLen)
						DrawRectangle(rollbackImg[rollbackCounter].MapImage,rect,2)		
					}
				}
			}
			
			rect2 := image.Rect(nexty * m.SquareLen, nextx * m.SquareLen, (nexty+1) * m.SquareLen, (nextx+1) * m.SquareLen)
			DrawRectangle(rollbackImg[rollbackCounter].MapImage,rect2,3)
			anim.Image = append(anim.Image,rollbackImg[rollbackCounter].MapImage)
			rollbackCounter += 1;
			anim.Delay = append(anim.Delay,60)//加入gif图像,设置帧间间隔60ms
			/*===== end =======*/

			book[nextx][nexty] = 0	//回退路径
		}
	}
	return 
}




func main(){
	imgfile, _ := os.Create(fmt.Sprintf("map.png"))
	giffile, _ := os.Create(fmt.Sprintf("map.gif"))


	
	var MazeMap Map
	//初始化地图数据
	MazeMap.Entinexty = MapArr
	MazeMap.SquareLen = 60
	MazeMap.MapImage = image.NewPaletted(image.Rect(0, 0, 240, 300),palette)
	MazeMap.DrawTheMap()


	book[0][0] = 1
	dfs(&MazeMap,0,0,0,forwardCounter)
	fmt.Println(min)

	err := png.Encode(imgfile, MazeMap.MapImage)
	if err != nil {
		log.Fatal(err)
		fmt.Printf("error is genrate\n")
	}

	err = gif.EncodeAll(giffile,&anim)
    if err != nil {
        log.Fatal(err)
    }
}
