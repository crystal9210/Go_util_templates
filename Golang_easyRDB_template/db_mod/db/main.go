package main

import (
	"fmt"
	"time"
)

const pageSize = 4096 // 4KB

type Cell struct {
	Data string
}

type Slot struct {
	Cell        *Cell
	Status      string
	LastUpdated time.Time
	Type        int // これはあらかじめどの種類のデータ内容を格納するスロットなのかのラベル情報をintのType値に対応させて分散処理側に保持しておき、アクセス時に参照し利用
}

type Page struct {
	Slots     []Slot
	MaxSlots  int
	UsedSlots int
}

func NewPage(maxSlots int) *Page {
	return &Page{
		Slots:     make([]Slot, maxSlots),
		MaxSlots:  maxSlots,
		UsedSlots: 0,
	}
}

type PageManager struct {
	Pages map[string]*Page // キーはページ名
}

func NewPageManager() *PageManager {
	return &PageManager{
		Pages: make(map[string]*Page),
	}
}

func (pm *PageManager) AddPage(name string, maxSlots int) {
	pm.Pages[name] = NewPage(maxSlots)
}

func (pm *PageManager) RemovePage(name string) {
	delete(pm.Pages, name) // 標準ライブラリからAPI提供されてるの便利
}

// データ走査機構
func (p *Page) ScanData() {
	for _, slot := range p.Slots {
		if slot.Cell != nil {
			fmt.Printf("データ: %s, タイプ: %d\n", slot.Cell.Data, slot.Type)
		}
	}
}

// どのようにセルやスロットを割り当てるか
// →用途による。今回はユーザデータの格納を想定；ページのインデックス番号、スロットのインデックス番号をそれぞれユーザIDに割り当てする
func (p *Page) AddCell(cell Cell) bool {
	if p.UsedSlots >= p.MaxSlots {
		return false
	}
	p.Slots[p.UsedSlots].Cell = &cell
	p.UsedSlots++
	// 実装仕様変更のため変更するが変更前は↓
	// for i := range p.Slots {
	// 	if p.Slots[i].Cell == nil {
	// 		p.Slots[i].Cell = &cell
	// 		return true
	// 	}
	// }
	return false
}

func main() {
	page := NewPage(10)

	cell1 := Cell{Data: "データ1"}
	success := page.AddCell(cell1)
	fmt.Println("セル追加成功:", success)

	for i, slot := range page.Slots {
		if slot.Cell != nil {
			fmt.Printf("スロット %d: %s\n", i, slot.Cell.Data)
		}
	}
}
