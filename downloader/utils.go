package downloader

import (
	"fmt"
	"sort"
	"strings"

	"github.com/fatih/color"

	"github.com/iawia002/lux/extractors"
)

var (
	blue = color.New(color.FgBlue)
	cyan = color.New(color.FgCyan)
)

func genSortedStreams(streams map[string]*extractors.Stream) []*extractors.Stream {
	sortedStreams := make([]*extractors.Stream, 0, len(streams))
	for _, data := range streams {
		sortedStreams = append(sortedStreams, data)
	}
	if len(sortedStreams) > 1 {
		sort.SliceStable(
			sortedStreams, func(i, j int) bool { return sortedStreams[i].Size > sortedStreams[j].Size },
		)
	}
	return sortedStreams
}

func printHeader(data *extractors.Data, output func(result string, err string)) {
	fmt.Println()
	cyan.Printf(" Site:      ") // nolint
	fmt.Println(data.Site)
	cyan.Printf(" Title:     ") // nolint
	fmt.Println(data.Title)
	cyan.Printf(" Type:      ") // nolint
	fmt.Println(data.Type)

	var stringOpt []string
	stringOpt = append(stringOpt, fmt.Sprintln())
	stringOpt = append(stringOpt, " Site:      ")
	stringOpt = append(stringOpt, data.Site)
	stringOpt = append(stringOpt, fmt.Sprintln())
	stringOpt = append(stringOpt, " Title:      ")
	stringOpt = append(stringOpt, data.Title)
	stringOpt = append(stringOpt, fmt.Sprintln())
	stringOpt = append(stringOpt, " Type:      ")
	stringOpt = append(stringOpt, string(data.Type))
	output(strings.Join(stringOpt, ""), "0")
}

func printStream(stream *extractors.Stream, output func(result string, err string)) {
	blue.Println(fmt.Sprintf("     [%s]  -------------------", stream.ID)) // nolint
	if stream.Quality != "" {
		cyan.Printf("     Quality:         ") // nolint
		fmt.Println(stream.Quality)
	}
	cyan.Printf("     Size:            ") // nolint
	fmt.Printf("%.2f MiB (%d Bytes)\n", float64(stream.Size)/(1024*1024), stream.Size)
	cyan.Printf("     # download with: ") // nolint
	fmt.Printf("lux -f %s ...\n\n", stream.ID)

	var stringOpt []string
	stringOpt = append(stringOpt, fmt.Sprintf("     [%s]  -------------------", stream.ID))
	if stream.Quality != "" {
		stringOpt = append(stringOpt, fmt.Sprintln())
		stringOpt = append(stringOpt, fmt.Sprintf("     Quality:         "))
		stringOpt = append(stringOpt, fmt.Sprintln(stream.Quality))
	}
	stringOpt = append(stringOpt, "     Size:            ")
	stringOpt = append(stringOpt, fmt.Sprintf("%.2f MiB (%d Bytes)\n", float64(stream.Size)/(1024*1024), stream.Size))
	stringOpt = append(stringOpt, "     # download with: ")
	stringOpt = append(stringOpt, fmt.Sprintf("lux -f %s ...\n\n", stream.ID))
	output(strings.Join(stringOpt, ""), "0")
}

func printInfo(data *extractors.Data, sortedStreams []*extractors.Stream, output func(result string, err string)) {
	printHeader(data, output)
	cyan.Printf(" Streams:   ") // nolint
	fmt.Println("# All available quality")
	var stringOpt []string
	stringOpt = append(stringOpt, fmt.Sprintln())
	stringOpt = append(stringOpt, " Streams:   ")
	stringOpt = append(stringOpt, fmt.Sprint("# All available quality"))
	output(strings.Join(stringOpt, ""), "0")
	for _, stream := range sortedStreams {
		printStream(stream, output)
	}
}

func printStreamInfo(data *extractors.Data, stream *extractors.Stream, output func(result string, err string)) {
	printHeader(data, output)

	cyan.Printf(" Stream:   ") // nolint
	fmt.Println()
	var stringOpt []string
	stringOpt = append(stringOpt, " Stream:   ")
	stringOpt = append(stringOpt, fmt.Sprintln())
	output(strings.Join(stringOpt, ""), "0")
	printStream(stream, output)
}
