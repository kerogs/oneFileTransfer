package cli

import (
	"fmt"
	"os"
	"os/exec"
	"oft/config/colors"
	// "oft/config/colors"
)

func Help(version string) {
    fmt.Println("Version : " + version)
	fmt.Println("Usage: oft <command> [options] \n")
	fmt.Println("Commands:")
	fmt.Println("  host start -d <directory>  : Start hosting files from specified directory")
    fmt.Println("       start -p <port>       : Optional (default : 7000)")
    fmt.Println("       stop                  : Stop hosting")
	fmt.Println("  scan                 : Scan for hosts that are hosting files")
	fmt.Println("  kill                 : Terminate the program (ctrl + c -> work too)")
	fmt.Println("  clear                : Clear the console screen")
}

func ClearScreen() {
    cmd := exec.Command("cmd", "/c", "cls")
    cmd.Stdout = os.Stdout
    cmd.Run()
}

func AsciiStart(version string){
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
    fmt.Println("   '!!$!!!!!!!!v"+version+" kerogs!!!!!!!!!!!!!!!!!!!!!!!!oooo..$$")
    fmt.Println("    !!$!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!$'")
    fmt.Println("    '$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$!!!!!!!!!!!!!!!!!!,")
    fmt.Println(".....$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$$.....")
	fmt.Println(color.Reset)
}

func Status(cliPrefix, cliType, cliSuffix, cliEnd, folderPath, port string,) {
    // Prefix
    fmt.Print("Mode      ")
    if cliPrefix == "user" {
        fmt.Println(color.Green + cliPrefix + color.Reset)
    } else if cliPrefix == "host" {
        fmt.Println(color.Red + cliPrefix + color.Reset)
        fmt.Println("-d       "+ folderPath)
        fmt.Println("-p       "+ port)
        
    } else if cliPrefix == "lurker"{
        fmt.Println(color.Magenta + cliPrefix + color.Reset)
    }

    // Type
    fmt.Print("Type      ")
    if cliType == "@" {
        fmt.Println(color.Green + cliType + color.Reset)
    }

    // Suffix
    fmt.Print("Suffix    ")
    if cliSuffix != "" {
        fmt.Println(color.Orange + cliSuffix + color.Reset)
    }

}
