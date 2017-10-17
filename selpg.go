package main

import (
    "flag"
		"fmt"
		"os"
		"os/exec"
		"bufio"
)
/*===============================struct to store the args =====================*/
type selpg  struct {
		s int //存储开始页面
		e int //存储结束页面
		page_len int  //存储页面长度的行数
		page_type int //分页的类型 可以用-f默认一行一页
		in_filename string //存储读入的文件名
		print_dest string //print destination 输出选择 屏幕输出/写进文件输出
}
/*==============================usage()==========================================*/
func usage() {
	fmt.Fprintf(os.Stderr,
		`usage:         [program name:selpg]
		 [-s start page]
		 [-e end page(>=s)]
		 [-l length of page (default = 72)]
		 [-f type of file (default = 1)]
		 [-d dest]
		 [filename : input file]
`)
}
/*==================================main()=====================================*/
func main() {
		my_selpg := selpg {
			s:-1,
			e:-1,
			page_len:72,
			page_type:1,
			in_filename: "",
			print_dest:"",
		}
		flag.IntVar(&my_selpg.s,"s", -1, "specify start page")
		flag.IntVar(&my_selpg.e,"e", -1, "specify end page(>=s)")
		flag.IntVar(&my_selpg.page_len, "l", 72, "specify length of a page")
		page_type := flag.Bool("f", false, "-f=0 means default length")
		print_dest := flag.String("d", "", "specify print dest.")
		flag.Usage = usage
    flag.Parse()

		if my_selpg.s == -1 ||
		   my_selpg.e == -1 ||
			 my_selpg.s > my_selpg.e ||
			 my_selpg.s < 1 ||
			 my_selpg.e < 1{
			flag.Usage()
			return
		}

		if my_selpg.page_len != 72 && *page_type == true {
			flag.Usage()
			return
		}

		if *page_type == true {
			my_selpg.page_type = 2
		}

		if *print_dest  != "" {
			my_selpg.print_dest = *print_dest
		}

		if len(flag.Args()) > 1 {
			flag.Usage()
			return
		}

		if len(flag.Args()) == 1 {
			my_selpg.in_filename = flag.Args()[0]
		}

		if my_selpg.page_type == 1 {
			func_1(my_selpg, my_selpg.in_filename != "", my_selpg.print_dest != "");
		} else {
			func_2(my_selpg, my_selpg.in_filename != "", my_selpg.print_dest != "");
		}
}
/*==============================func_1()======================================*/
func func_1(my_selpg selpg, file bool, pipe bool) {
	cmd := exec.Command("cat", "-n")
	stdin, err:= cmd.StdinPipe()
	if err != nil {
		panic(err)
	}
	cur_page := 1
	cur_lines := 0
	if file {
		file_in, err := os.OpenFile(my_selpg.in_filename,os.O_RDWR,os.ModeType)
		defer file_in.Close()
		if err != nil {
			panic(err)
			return
		}
    line := bufio.NewScanner(file_in)
    for line.Scan() {
			if cur_page >= my_selpg.s && cur_page <= my_selpg.e {
				os.Stdout.Write([]byte(line.Text()+"\n"))
				stdin.Write([]byte(line.Text()+"\n"))
			}
			cur_lines++;
			if cur_lines %= my_selpg.page_len; cur_lines == 0 {
				cur_page++;
			}
    }
	} else {
		tmp_s := bufio.NewScanner(os.Stdin)
		for tmp_s.Scan() {
			if cur_page >= my_selpg.s && cur_page <= my_selpg.e {
				os.Stdout.Write([]byte(tmp_s.Text()+"\n"))
				stdin.Write([]byte(tmp_s.Text()+"\n"))
			}
			cur_lines++;
			if cur_lines %= my_selpg.page_len; cur_lines == 0 {
				cur_page++;
			}
		}
	}
	if cur_page < my_selpg.e {
		fmt.Fprintf(os.Stderr, "[error]  Can not reach end page!\n")
	}
	if pipe {
		stdin.Close()
		cmd.Stdout = os.Stdout;
		cmd.Start()
	}

}
/*==================================func_2()====================================*/
func func_2(sa selpg, file bool, pipe bool) {
	cmd := exec.Command("cat", "-n")
	stdin, err:= cmd.StdinPipe()
	if err != nil {
		panic(err)
	}
	cur_page := 1
	if file {
		file_in, err := os.OpenFile(sa.in_filename,os.O_RDWR,os.ModeType)
		defer file_in.Close()
		if err != nil {
			panic(err)
			return
		}
		line := bufio.NewScanner(file_in)
    for line.Scan() {
			flag := false
			for _,c := range line.Text() {
				if c == '\f' {
					if cur_page >= sa.s && cur_page <= sa.e {
						flag = true
						os.Stdout.Write([]byte("\n"))
						stdin.Write([]byte("\n"))
					}
					cur_page++;
				} else {
					if cur_page >= sa.s && cur_page <= sa.e {
						os.Stdout.Write([]byte(string(c)))
						stdin.Write([]byte(string(c)))
					}
				}
			}
			if flag != true && cur_page >= sa.s && cur_page <= sa.e {
				os.Stdout.Write([]byte("\n"))
				stdin.Write([]byte("\n"))
			}
			flag = false
    }
	} else {
		tmp_s := bufio.NewScanner(os.Stdin)
		for tmp_s.Scan() {
			flag := false
			for _,c := range tmp_s.Text() {
				if c == '\f' {
					if cur_page >= sa.s && cur_page <= sa.e {
						flag = true
						os.Stdout.Write([]byte("\n"))
						stdin.Write([]byte("\n"))
					}
					cur_page++;
				} else {
					if cur_page >= sa.s && cur_page <= sa.e {
						os.Stdout.Write([]byte(string(c)))
						stdin.Write([]byte(string(c)))
					}
				}
			}
			if flag != true && cur_page >= sa.s && cur_page <= sa.e {
				os.Stdout.Write([]byte("\n"))
				stdin.Write([]byte("\n"))
			}
			flag = false
		}
	}
	if cur_page < sa.e {
		fmt.Fprintf(os.Stderr, "[error]  Can not reach end page!\n")
	}
	if pipe {

		stdin.Close()
		cmd.Stdout = os.Stdout
		cmd.Start()
	}
}
