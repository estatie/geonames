package main

import (
	"fmt"
	"github.com/gernest/wow"
	"github.com/gernest/wow/spin"
	"github.com/mkrou/geonames"
	"github.com/mkrou/geonames/models"
	"os"
	"runtime"
	"time"
)

func main() {
	p := geonames.NewParser()

	w := wow.New(os.Stdout, spin.Spinner{Frames: []string{"⚙️"}}, "  Parsing all cities with a population > 5000...")
	w.Persist()
	w.Text("").Spinner(spin.Get(spin.Earth)).Start()

	count := 0
	since := time.Now()
	var m runtime.MemStats
	var max uint64 = 0
	err := p.GetGeonames(geonames.Cities5000, func(geoname *models.Geoname) error {
		count++
		w.Text(fmt.Sprintf("%d: %s", count, geoname.Name))

		runtime.ReadMemStats(&m)
		if max < m.Alloc {
			max = m.Alloc
		}
		return nil
	})
	if err != nil {
		w.PersistWith(spin.Spinner{Frames: []string{"🔥"}}, fmt.Sprintf(" Error: %s", err.Error()))
		return
	}
	duration := time.Since(since)

	w.PersistWith(spin.Spinner{Frames: []string{"✅"}}, fmt.Sprintf(" Done!"))
	w.PersistWith(spin.Spinner{Frames: []string{"⛩"}}, fmt.Sprintf("  Cities: %d", count))
	w.PersistWith(spin.Spinner{Frames: []string{"⏱"}}, fmt.Sprintf("  Duration: %d sec", duration/time.Second))
	w.PersistWith(spin.Spinner{Frames: []string{"💾️‍"}}, fmt.Sprintf(" Memory: %d Mb", max/1024/1024))
}
