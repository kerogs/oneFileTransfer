package base

import (
	"fmt"
	"os"
	"os/exec"
	"oft/config/colors"
	// "oft/config/colors"
)

func Help() {
	fmt.Println("Usage: oft <command> [options]")
	fmt.Println("Commands:")
	fmt.Println("  host -d <directory> : Start hosting files from specified directory")
	fmt.Println("  scan                 : Scan for hosts that are hosting files")
	fmt.Println("  kill                 : Terminate the program (ctrl + c -> work too)")
	fmt.Println("  clear                : Clear the console screen")
}

func ClearScreen() {
    cmd := exec.Command("cmd", "/c", "cls")
    cmd.Stdout = os.Stdout
    cmd.Run()
}

func AsciiStart(){
	fmt.Println(color.Red)
    fmt.Println("                       $o")
    fmt.Println("                       $                     .........")
    fmt.Println("                      $$$      .oo..     'oooo'oooo'ooooooooo....")
    fmt.Println("                       $       $$$$$$$")
    fmt.Println("                   .ooooooo.   $$!!!!!")
    fmt.Println("                 .'.........'. $$!!!!!      o$$oo.   ...oo,oooo,oooo'ooo''")
    fmt.Println("    $          .o'  oooooo   '.$$!!!!!      $$!!!!!       'oo''oooo''")
    fmt.Println(" ..o$ooo...    $                '!!''!.     $$!!!!!")
    fmt.Println(" $    ..  '''oo$$$$$$$$$$$$$.    '    'oo.  $$!!!!!")
    fmt.Println(" !.......      '''..$$ $$ $$$   ..        '.$$!!''!")
    fmt.Println(" !!$$$!!!!!!!!oooo......   '''  $$ $$ :o           'oo.")
    fmt.Println(" !!$$$!!!$$!$$!!!!!!!!!!oo.....     ' ''  o$$o .      ''oo..")
    fmt.Println(" !!!$$!!!!!!!!!!!!!!!!!!!!!!!!!!!!ooooo..      'o  oo..    $")
    fmt.Println("  '!!$$!!!!!!oneFileTransfer!!!!!!!!!!!!!!!oooooo..  ''   ,$")
    fmt.Println("   '!!$!!!!!!!!!!!!by kerogs!!!!!!!!!!!!!!!!!!!!!!!!oooo..$$")
    fmt.Println("    !!$!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!$'")
    fmt.Println("    '$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$!!!!!!!!!!!!!!!!!!,")
    fmt.Println(".....$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$.....")
	fmt.Println(color.Reset)
}