package doushu

import (
	_ "embed"
	"encoding/csv"
	"log"
	"strings"
)

//go:embed table.csv
var data string

var p *Parser

func init() {
	p = &Parser{}

	var frags []string
	p.titles, p.headers, frags = splitData(data)

	buildIndex(p, frags)
}

type Parser struct {
	index   []Element
	titles  []string
	headers []string
	starts  []int
}

func (p *Parser) get(title string, left, top int) Element {
	var i int
	for ; i < len(p.titles); i++ {
		if p.titles[i] == title {
			break
		}
	}
	if i == len(p.titles) {
		log.Printf("titles not found: %s", title)
		return -1
	}
	start := p.starts[i]
	rows := int(p.index[start])
	if left > rows || top > int(p.index[start+1]) {
		log.Printf("%s out of range: %d %d", title, left, top)
		return -1
	}

	idx := start + 2 + top*rows + left

	return p.index[idx]
}

func (p *Parser) batchGet(title string, top int) []Element {
	var i int
	for ; i < len(p.titles); i++ {
		if p.titles[i] == title {
			break
		}
	}
	if i >= len(p.titles) {
		log.Printf("title not found: %s", title)
		return nil
	}
	start := p.starts[i]
	rows := int(p.index[start])
	if top > int(p.index[start+1]) {
		log.Printf("title %s, out of range: %d", title, top)
		return nil
	}

	idx := start + 2 + top*rows

	return p.index[idx : idx+rows]
}

func splitData(data string) (titles, headers, contents []string) {
	var title, header, value string
	for _, line := range strings.Split(data, "\n") {
		line := strings.TrimSpace(line)

		if strings.HasPrefix(line, "### ") {
			if value != "" {
				titles = append(titles, strings.TrimSpace(title))
				headers = append(headers, strings.TrimSpace(header))
				contents = append(contents, value)
			}
			title, header, _ = strings.Cut(line[4:], " ")
			value = ""
			continue
		}

		if len(line) == 0 || strings.HasPrefix(line, "# ") {
			continue
		}

		if value != "" {
			value += "\n" + line
		} else {
			value = line
		}
	}
	if value != "" {
		titles = append(titles, strings.TrimSpace(title))
		headers = append(headers, strings.TrimSpace(header))
		contents = append(contents, value)
	}

	return titles, headers, contents
}

func buildIndex(p *Parser, frags []string) {
	for i, frag := range frags {
		p.starts = append(p.starts, len(p.index))

		table := readCSV(frag)
		if len(table) < 2 {
			continue
		}

		if p.titles[i] == "诸星在十二宫庙旺利陷表" {
			buildLightIndex(p, table)
			continue
		}

		left, top, _ := strings.Cut(p.headers[i], "/")

		var multiRows = 1
		if strings.HasSuffix(left, "*") {
			multiRows = len(strings.Split(table[1][0], ""))
		}
		var multiCols = 1
		if strings.HasSuffix(top, "*") {
			multiCols = len(strings.Split(table[0][1], ""))
		}
		p.index = append(p.index, Element((len(table)-1)*multiRows), Element((len(table[0])-1)*multiCols))
		for col := 1; col < len(table[0]); col++ {
			for m := 0; m < multiCols; m++ { // 子丑,寅卯,辰巳,午未,申酉,戌亥
				for m := 0; m < multiRows; m++ { // 子辰申,丑巳酉,寅午戌,卯未亥
					for row := 1; row < len(table); row++ {
						name := table[row][col]
						code := ValueOf(name)
						p.index = append(p.index, code)
					}
				}
			}
		}
	}
}

func readCSV(fragment string) [][]string {
	table, err := csv.NewReader(strings.NewReader(fragment)).ReadAll()
	if err != nil {
		log.Printf("%s: %s", err, fragment)
	}
	return table
}

func buildLightIndex(p *Parser, table [][]string) {
	// /,庙,旺,得地,利益,平和,不得地,陷
	// 子,机府阴相梁破,武同贪巨杀,昌曲,,紫廉,,阳羊火铃
	// ==>
	// /,子,丑,寅,卯,辰,巳,午,未,申,酉,戌,亥
	// 紫微,平和,庙,。。。
	p.index = append(p.index, Mingzhu-Ziwei, 12)

	lightIndex := make([]Element, (Mingzhu-Ziwei)*12)

	for light := 1; light < len(table[0]); light++ {
		for zhi := 1; zhi < len(table); zhi++ {
			names := table[zhi][light]
			for _, name := range strings.Split(names, "") {
				var has bool
				for star := Ziwei; star <= Shenzhu; star++ {
					if strings.Contains(star.String(), name) {
						if name == "曲" && star == Wuqu {
							continue
						}
						// if has {
						// 	log.Printf("found duplicate %s: %s", name, star)
						// 	continue
						// }
						has = true
						idx := int(Mingzhu-Ziwei)*(zhi-1) + star.Value()
						lightIndex[idx] = Miao.Next(light - 1)
						break
					}
				}
				if !has {
					log.Printf("not found %s", name)
				}
			}
		}
	}
	p.index = append(p.index, lightIndex...)
}
