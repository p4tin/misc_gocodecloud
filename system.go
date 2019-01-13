package main

import (
	"fmt"
	"os"

	"github.com/cloudfoundry/gosigar"
)

const output_format = "%-15s %4s %4s %5s %4s %-15s\n"

func formatSize(size uint64) string {
	return sigar.FormatSize(size * 1024)
}

func format(val uint64) uint64 {
	return val / 1024
}

func main() {
	fmt.Println("\n")
	uptime()
	fmt.Println("\n\n")
	mem()
	fmt.Println("\n\n")
	df()
	//fmt.Println("\n\n")
	//ps()
}

//func ps() {
//	ps := sigar.ProcState{}
//	ps.Get()
//}

func uptime() {
	load := sigar.LoadAverage{}
	load.Get()
	fmt.Fprintf(os.Stdout, "Load Average: %f\n", load.One)
		uptime := sigar.Uptime{}
	uptime.Get()
	fmt.Fprintf(os.Stdout, "Uptime: %s\n", uptime.Format())
}

func mem() {
	mem := sigar.Mem{}
	swap := sigar.Swap{}

	mem.Get()
	swap.Get()

	fmt.Fprintf(os.Stdout, "%18s %10s %10s\n",
		"total", "used", "free")
	fmt.Fprintf(os.Stdout, "Mem:    %10d %10d %10d\n",
		format(mem.Total), format(mem.Used), format(mem.Free))

	fmt.Fprintf(os.Stdout, "-/+ buffers/cache: %10d %10d\n",
		format(mem.ActualUsed), format(mem.ActualFree))

	fmt.Fprintf(os.Stdout, "Swap:   %10d %10d %10d\n",
		format(swap.Total), format(swap.Used), format(swap.Free))
}

func df() {
	fslist := sigar.FileSystemList{}
	fslist.Get()

	fmt.Fprintf(os.Stdout, output_format,
		"Filesystem", "Size", "Used", "Avail", "Use%", "Mounted on")

	for _, fs := range fslist.List {
		dir_name := fs.DirName

		usage := sigar.FileSystemUsage{}

		usage.Get(dir_name)

		fmt.Fprintf(os.Stdout, output_format,
			fs.DevName,
			formatSize(usage.Total),
			formatSize(usage.Used),
			formatSize(usage.Avail),
			sigar.FormatPercent(usage.UsePercent()),
			dir_name)
	}
}
