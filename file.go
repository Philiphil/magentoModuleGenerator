package main
import (
	"os"
	"io/ioutil"
	"bufio"
	"strings"
	"time"
	"regexp"
)

func writeToFile(write string, file string){
	ioutil.WriteFile(file, []byte(write), 0644) 
}

func getFileType(file string)(typef string){
    r, _ := regexp.Compile("\\.(.+)*")
	typef = r.FindString(file)
	return strings.ToLower(typef[1:])
}

func getFileContent(file string)(string){
	content, _ := ioutil.ReadFile(file)
	return string(content)
}

func getFileContentLine(file string)(content []string){
	File, _ := os.Open(file)
    defer File.Close()

    scanner := bufio.NewScanner(File)
    for scanner.Scan() {
        content = append(content, scanner.Text())
	}
	return
}


func isDirectory(path string) (bool, error) {
    if  fileInfo, err := os.Stat(path); err == nil{
    	return fileInfo.IsDir(),nil
    }else{
    	 return false, err
    }
}
   
func directorySeparator(path string) (newpath string) {
	newpath =  path
	if   path[len(path)-1:] != string (os.PathSeparator){
		newpath += string (os.PathSeparator) 
	}
	return 
}

func explore(monitoredfilelocation string)(contents map[string]string){
	if boolean, r := isDirectory( monitoredfilelocation ); boolean && r == nil{
		monitoredfilelocation = directorySeparator(monitoredfilelocation)
	}
	
	contents = map[string]string{}
 	files, _ := ioutil.ReadDir(monitoredfilelocation)
   	for _, f := range files {
   		str := monitoredfilelocation + f.Name()
        if boolean, r := isDirectory( str ); (boolean && r == nil){
			str = directorySeparator(str)
			contents[str] =  f.ModTime().Format(time.RFC3339) 
			for k, v := range explore(str) {
				contents[k] = v
			}
        }else if (r == nil){
			contents[str] =  f.ModTime().Format(time.RFC3339) 
        }
	}
	return
}

func getDiff(oldstate  map[string]string, newstate  map[string]string) (diff map[string]string){
	diff = map[string]string{}
	for key, val := range newstate {
		if o_val, ok :=oldstate[key];ok {
			delete(oldstate, key)
			if o_val != val{
				diff[key] = val
			}
		}else{
			diff[key] = val
		}
	}
	for key, val := range oldstate {
		diff[key] = val	
	}
	return
}
