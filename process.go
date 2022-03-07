package main

import (
	"fmt"
	"log"
	"github.com/mft-labs/runb2bqueries/utils"
	"os"
	"strings"
	"sync"
	"time"
)

type Process struct {
	Util *utils.Util
}

func (proc *Process) RunQueries() (err error){
	queries := proc.Util.GetValue2("DEFAULT","queries",true)
	if len(queries) == 0 {
		return fmt.Errorf("No queries found")
	}
	command := proc.Util.GetValue2("DEFAULT","command",true)
	wg := sync.WaitGroup{}
	log.Printf("Running queries:%v",queries)
	qlist := strings.Split(queries,",")
	exectime := time.Now().Format("20060102150405")
	for idx, query:=range qlist {
		log.Printf("Running query: %d) %v",idx+1,query)
		//infile := proc.Util.GetValue2(query,"infile",true)
		//outfile := proc.Util.GetValue2(query,"outfile",true)
		infile := fmt.Sprintf("runb2bqueries_in_%s_%s.in",strings.ToLower(query),exectime)
		outfile := fmt.Sprintf("runb2bqueries_in_%s_%s.out",strings.ToLower(query),exectime)
		sql := proc.Util.GetValue2(query,"query",true)
		sql = fmt.Sprintf("%s\n/\n",sql)
		proc.Util.WriteFile(infile,[]byte(sql))
		//csvfile := proc.Util.GetValue2(query,"csvfile",true)
		csvfile := fmt.Sprintf("runb2bqueries_%s_%s.csv",strings.ToLower(query),exectime)
		columnHeaders := proc.Util.GetValue2(query,"column_header",true)
		cmd := fmt.Sprintf(command,infile,outfile)
		wg.Add(1)
		//var output string
		go func(cmd string) {
			args := strings.Split(cmd," ")
			log.Printf("Running command with:%#v",args)
			_, err = RunCommand(args[0],args[1:],&wg,true)
		}(cmd)
		wg.Wait()
		if err == nil {
			output, err := proc.Util.ReadFile(outfile)
			if err == nil {
				contents, err := proc.ParseOutput(string(output), columnHeaders,false)
				if err == nil {
					proc.Util.WriteFile(csvfile,[]byte(contents))
				} else {
					log.Printf("Error occurred while parsing output:%v",err)
				}
				os.Remove(infile)
				os.Remove(outfile)
			} else {
				log.Printf("Failed to read file:%v",outfile)
			}
		} else {
			log.Printf("Error occurred, please check %s",outfile)
		}
	}
	return nil
}

func (proc *Process) ParseOutput(output, colHeaders string, debug bool) (string, error) {
	//log.Printf("Returning output:\n%s",output)
	text  := strings.Split(output,"\n")
	result := colHeaders+"\n"
	curdate := time.Now().Format("01/02/2006")
	for idx, line := range text {
		if debug {
			log.Printf("%d) %s",idx+1,line)
		}
		if idx > 2 {
			fields := strings.Split(line,"\t")
			columns := strings.Join(fields,",")
			if len(columns)>0 {
				result += curdate+","+columns[0:len(columns)-1] +"\n"
			}

		}

	}
	return result, nil
}