package evaluator

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/AvicennaJr/Nuru/object"
)

var builtins = map[string]*object.Builtin{
	"jaza": {
		Fn: func(args ...object.Object) object.Object {

			if len(args) > 1 {
				return newError("Samahani, hii function inapokea hoja 0 au 1, wewe umeweka %d", len(args))
			}

			if len(args) > 0 && args[0].Type() != object.STRING_OBJ {
				return newError(fmt.Sprintf(`Tafadhali tumia alama ya nukuu: "%s"`, args[0].Inspect()))
			}
			if len(args) == 1 {
				prompt := args[0].(*object.String).Value
				fmt.Fprint(os.Stdout, prompt)
			}

			buffer := bufio.NewReader(os.Stdin)

			line, _, err := buffer.ReadLine()
			if err != nil && err != io.EOF {
				return newError("Nimeshindwa kusoma uliyo yajaza")
			}

			return &object.String{Value: string(line)}
		},
	},
	"andika": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) == 0 {
				fmt.Println("")
			} else {
				var arr []string
				for _, arg := range args {
					if arg == nil {
						return newError("Hauwezi kufanya operesheni hii")
					}
					arr = append(arr, arg.Inspect())
				}
				str := strings.Join(arr, " ")
				print(str + "\n")
			}
			return nil
		},
	},
	"aina": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("Samahani, tunahitaji Hoja 1, wewe umeweka %d", len(args))
			}

			return &object.String{Value: string(args[0].Type())}
		},
	},
	"fungua": {
		Fn: func(args ...object.Object) object.Object {

			if len(args) > 2 {
				return newError("Samahani, Hatuhitaji hoja zaidi ya 2, wewe umeweka %d", len(args))
			}
			filename := args[0].(*object.String).Value
			mode := os.O_RDONLY
			if len(args) == 2 {
				fileMode := args[1].(*object.String).Value
				switch fileMode {
				case "soma":
					mode = os.O_RDONLY
				// still buggy, will work on this soon
				case "andika":
					mode = os.O_WRONLY
					os.Remove(filename)
				case "ongeza":
					mode = os.O_APPEND
				default:
					return newError("Tumeshindwa kufungua file na mode %s", fileMode)
				}
			}
			file, err := os.OpenFile(filename, os.O_CREATE|mode, 0644)
			if err != nil {
				return &object.Null{}
			}
			var reader *bufio.Reader
			var writer *bufio.Writer
			if mode == os.O_RDONLY {
				reader = bufio.NewReader(file)
			} else {
				writer = bufio.NewWriter(file)
			}
			return &object.File{Filename: filename, Reader: reader, Writer: writer, Handle: file}
		},
	},

	// "jumla": {
	// 	Fn: func(args ...object.Object) object.Object {
	// 		if len(args) != 1 {
	// 			return newError("Hoja hazilingani, tunahitaji=1, tumepewa=%d", len(args))
	// 		}

	// 		switch arg := args[0].(type) {
	// 		case *object.Array:

	// 			var sums float64
	// 			for _, num := range arg.Elements {

	// 				if num.Type() != object.INTEGER_OBJ && num.Type() != object.FLOAT_OBJ {
	// 					return newError("Samahani namba tu zinahitajika")
	// 				} else {
	// 					if num.Type() == object.INTEGER_OBJ {
	// 						no, _ := strconv.Atoi(num.Inspect())
	// 						floatnum := float64(no)
	// 						sums += floatnum
	// 					} else if num.Type() == object.FLOAT_OBJ {
	// 						no, _ := strconv.ParseFloat(num.Inspect(), 64)
	// 						sums += no
	// 					}

	// 				}
	// 			}

	// 			if math.Mod(sums, 1) == 0 {
	// 				return &object.Integer{Value: int64(sums)}
	// 			}

	// 			return &object.Float{Value: float64(sums)}

	// 		default:
	// 			return newError("Samahani, hii function haitumiki na %s", args[0].Type())
	// 		}
	// 	},
	// },
}
