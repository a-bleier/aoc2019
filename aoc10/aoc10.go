package aoc10

import(
	"fmt"
	"github.com/a-bleier/aoc2019/fileio"
	"math"
	"sort"
)

type coord struct {
	x,y int
}

type coordAtan struct {
	c coord
	atan float64
}


func getSpace() [][]bool {
	lines := fileio.GetLinesFromFile("aoc10")
	var field [][]bool
	for i,line := range(lines) {
		field = append(field, make([]bool,0))
		for _,char := range(line) {
			field[i] = append(field[i],char == '#')
		}
	}

	return field
}

func abs(x int) int {
	return int(math.Abs(float64(x)))
}

func vaporize(field [][]bool, asteroidsToDestroy int, x, y int) coord {
	
	var coordList []coord
	cont := true
	i:=-1
	for cont {
		i++
		deltaDx, deltaDy := 1,1
		diffy, diffx := 0,0
		startwertx, startwerty := 0, 0
		fmt.Println("sector changed----")
		if i % 4 == 0 {
			deltaDx, deltaDy = 1,-1
			diffy = y
			diffx = len(field[0])-1-x
			startwertx, startwerty = 0,-1
		} else if i % 4 == 1{
			deltaDx, deltaDy = 1,1
			diffy = len(field)-1-y
			diffx = len(field[0])-1-x
			startwertx, startwerty = 1,0
		} else if i % 4 == 2{
			deltaDx, deltaDy = -1,1
			diffy = len(field)-1-y
			diffx = x
			startwertx, startwerty = 0,1
		} else if i % 4 == 3{
			deltaDx, deltaDy = -1,-1
			diffy = y
			diffx = x
			startwertx, startwerty = -1,0
		}

		coordList = make([]coord,0)
		count := 0
	
		for dy := startwerty; abs(dy) <= diffy; dy += deltaDy{ 
			for dx := startwertx; abs(dx) <= diffx; dx += deltaDx {
		
					//fmt.Println("me here")
				euc := euclid(abs(dx),abs(dy))
				if  !(euc == 1 || euc == 0) { //it is a scalar product
					continue
				}
				if  dx == 0 && dy == 0 {
					continue
				}
				
				xtemp, ytemp := x+dx,y+dy
						
					
				for xtemp >= 0 && xtemp < len(field[0]) && ytemp >= 0 && ytemp < len(field){
						
						
					if field[ytemp][xtemp] {
						fmt.Printf("vaporizing %d %d \n", xtemp, ytemp)
						count++
						field[ytemp][xtemp] = false						
						coordList = append(coordList, coord{xtemp, ytemp})	
						break			
					}
					xtemp, ytemp = xtemp+dx,ytemp+dy						
				}								
			}
		
		}


		if asteroidsToDestroy > count {
			asteroidsToDestroy -= count
		} else {
			cont = false
			break
		}	
		
	}

	
	var lastCoord coord
	var minRad float64 = 10
	list := make([]coordAtan, 0)
	
	for _,c := range(coordList) {
		dx, dy := c.x-x, y-c.y		
		rad :=  math.Atan2(float64(dy), float64(dx))
		list = append(list, coordAtan{c,rad})	
	}

	sort.Slice(list, func(i,j int) bool { return list[i].atan < list[j].atan })

	for i:= 1; i <=asteroidsToDestroy; i++ {
		ast := list[len(list)-i]
		if ast.atan < minRad {
			lastCoord = ast.c
			minRad = ast.atan
		}
	}

	
	return lastCoord
}

func countDetectableAteroids(field [][]bool, x ,y int ) int {
	counter := 0
	xfactor, yfactor := 1,1
	diffy, diffx := 0,0
	startwert := 0

	for i := 0; i < 4; i++ {

		if i == 0 {
			xfactor, yfactor = 1,1
			diffy = len(field)-1-y
			diffx = len(field[0])-1-x
			startwert = 1
		} else if i == 1{
			xfactor, yfactor = -1,1
			diffy = len(field)-1-y
			diffx = x
			startwert = 0
		} else if i == 2{
			xfactor, yfactor = 1,-1
			diffy = y
			diffx = len(field[0])-1-x
			startwert = 0
		} else if i == 3{
			xfactor, yfactor = -1,-1
			diffy = y
			diffx = x
			startwert = -1
		}

		for dy := startwert; abs(dy) <= diffy; dy += yfactor{ 
			for dx := startwert; abs(dx) <= diffx; dx += xfactor {
	
				//fmt.Println("me here")
				euc := euclid(abs(dx),abs(dy))
				if  !(euc == 1 || euc == 0) { //it is a scalar product
					continue
				}
				if  dx == 0 && dy == 0 {
					continue
				}
			

				xtemp, ytemp := x+dx,y+dy
	
				
				
				for xtemp >= 0 && xtemp < len(field[0]) && ytemp >= 0 && ytemp < len(field){
					//fmt.Printf("looking up %d %d ", ytemp, xtemp)
					if field[ytemp][xtemp] {
					//	fmt.Printf("found one\n")
						counter++
						break
					}
					//fmt.Printf("\n")
					xtemp, ytemp = xtemp+dx,ytemp+dy
				}
	
				
			
			}
	
		}

	}	
	return counter
}

func euclid(a, b int) int {
	for b != 0 {
		h := a % b
		a = b
		b = h
	}
	return a
}

func findMaxAsteroids(field [][]bool) int {
	var maxNum int
	var asteroids []coord

	for y,row := range(field) {
		for x,cell := range(row) {
			if cell {
				asteroids = append(asteroids,coord{x,y})
			}
		}
	}
	x,y := 0,0
	for _,ast := range(asteroids) {
		fmt.Println("checking", ast.y, ast.x)
		num := countDetectableAteroids(field, ast.x, ast.y)
		fmt.Println("num: ", num)
		if num > maxNum {
			maxNum = num
			x = ast.x
			y = ast.y
			
		}
	}
	fmt.Println("x: ", x,"y: ", y) // 17 22
	return maxNum
}

func findLastVaporizedAsteroid(field [][]bool){
	lastAsteroid := 200
	astCoord := vaporize(field, lastAsteroid, 11, 13)
	fmt.Println("Last vaporized asteroid is ", astCoord.x, astCoord.y)
}

func Aoc10Main() {
	field := getSpace()
	fmt.Println(field)
	fmt.Println("Max number of asteroids detected on the best location: ")
	fmt.Println(findMaxAsteroids(field))
	findLastVaporizedAsteroid(field)
}