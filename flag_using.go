package main

import (
	"flag"
	"fmt"
)

func main() {
	//flag简单使用方法
	backup_dir := flag.String("b", "/home/default_dir", "backup path")
	debug_mode := flag.Bool("d", false, "debug mode")
	flag.Parse()
	fmt.Println("backup_dir: ",*backup_dir)
	fmt.Println("backup_dir: ",backup_dir)
	fmt.Println("debug_mode: ",*debug_mode)
	fmt.Println("debug_mode: ",debug_mode)

	var ip = flag.Int("flagname",1234,"help message for flagname")
	fmt.Println("ip:",*ip)
	fmt.Println("ip:",ip)

	//s := "name o"
	//flag.Var(&s,"name","help message for flagname")

	data_path := flag.String("D","/home/manu/sample/","DB data path")

	log_file := flag.String("l","/home/manu/sample.log","log file")

	nowait_flag :=flag.Bool("W",false,"do not wait until operation completes")



	flag.Parse()



	var cmd string = flag.Arg(0);



	fmt.Printf("action : %s\n",cmd)

	fmt.Printf("data path: %s\n",*data_path)

	fmt.Printf("log file : %s\n",*log_file)

	fmt.Printf("nowait : %v\n",*nowait_flag)



	fmt.Printf("-------------------------------------------------------\n")



	fmt.Printf("there are %d non-flag input param\n",flag.NArg())
	fmt.Println("len:",len(flag.Args()))

	for i,param := range flag.Args() {

		fmt.Printf("#%d :%s\n", i, param)
	}
}