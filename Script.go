package main

import (
	"io/ioutil"
	"fmt"
	"os"
	"strings"
	"bufio"
)

func read_files_curr_dir() []os.FileInfo {
	files, _ := ioutil.ReadDir("./")
	return files

}

func read_file(path string) ([]string, error) {
	file, err := os.Open("./" + path + "/index.html");
	if err != nil {
		fmt.Println(err)
	}
	//close at return
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func add_html_tags(document_lines []string) []string {
	for i := 0; i < len(document_lines); i++ {
		line := document_lines[i]
		if strings.Contains(line, "href=\"/posts/") || strings.Contains(line, "href= \"/posts/") {
			divided_line := strings.Fields(line)
			href_element := divided_line[1]
			//replace
			if strings.Contains(href_element, ".html") {
				continue
			}
			href_element = href_element[:len(href_element)-1] + string(".html\"") + href_element[len(href_element):]
			divided_line[1] = href_element
			line = strings.Join(divided_line, " ")
			document_lines[i] = line
		}
	}
	return document_lines
}

func write_file(path string, document_lines []string) {
	file, err := os.Create("./" + path + "/index.html")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)

	for _, line := range document_lines {
		fmt.Fprintln(writer,line)
		writer.Flush()
	}

}

func main() {

	for _, fileInfo := range read_files_curr_dir() {
		dir_name := fileInfo.Name()
		if fileInfo.IsDir() && strings.Contains(dir_name, "page-") {
			file_in_lines, scanner_error := read_file(dir_name)
			if scanner_error != nil {
				fmt.Println(scanner_error)
			}
			file_in_lines = add_html_tags(file_in_lines)
			write_file(dir_name, file_in_lines)

		} else if dir_name == "index.html" {
			file_in_lines, scanner_error := read_file("")
			if scanner_error != nil {
				fmt.Println(scanner_error)
			}
			file_in_lines = add_html_tags(file_in_lines)
			write_file("",file_in_lines)
		}
	}

}
