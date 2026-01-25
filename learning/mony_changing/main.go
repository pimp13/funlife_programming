package main

import (
	"fmt"
	"log"
	"math/rand"
	"sort"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {

	var people []int

	for i := 0; i < 50; i++ {
		people = append(people, 100)
	}

	for i := 0; i < 10_000; i++ {
		for person1 := 0; person1 < 50; person1++ {
			person2 := rand.Intn(len(people))

			for people[person2] == 0 {
				person2 = rand.Intn(len(people))
			}
			if people[person1] != 0 {
				people[person1] -= 1
				people[person2] += 1
			}

		}
	}

	// ایجاد نمودار جدید
	p := plot.New()

	// عنوان نمودار و برچسب محور X و Y
	p.Title.Text = "نمودار ستونی"
	p.X.Label.Text = "شاخص"
	p.Y.Label.Text = "مقدار"

	// ایجاد داده‌های نمودار ستونی
	barWidth := vg.Points(3) // عرض ستون‌ها
	bars := make(plotter.Values, len(people))
	for i, v := range people {
		bars[i] = float64(v)
	}

	sort.Slice(bars, func(i, j int) bool {
		return int(bars[i]) > int(bars[j])
	})

	// ایجاد نمودار ستونی
	barChart, err := plotter.NewBarChart(bars, barWidth)
	if err != nil {
		log.Fatal(err)
	}

	// اضافه کردن نمودار ستونی به نمودار اصلی
	p.Add(barChart)

	// ذخیره نمودار به صورت یک فایل PNG
	if err := p.Save(20*vg.Inch, 10*vg.Inch, "barchart.png"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("نمودار ستونی با موفقیت ذخیره شد.")

}
