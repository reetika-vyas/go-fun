package tutorial_test

import (
	"fmt"
	"math"
	"math/rand"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var global, second_global = 5, 10

var _ = Describe("GoTour", func() {

	Context("Variables", func() {
		// var (
		// 	r, i rune = 8, 9
		// 	g         = 0.867 + 0.5i // complex128
		// )

		It("should have globals", func() {
			Expect(global).To(Not(BeNil()))
			Expect(second_global).To(Not(BeNil()))
		})

		It("should have locals", func() {
			var local string = "localvariable"
			shortHand := "Shorthand Variable"

			Expect(local).To(Not(BeNil()))
			Expect(shortHand).To(Not(BeNil()))
		})

		It("should have constants", func() {
			const constant_string = "Constant"
			Expect(constant_string).To(Not(BeNil()))

		})

		It("should have exported names", func() {
			Expect(math.Pi).To(Not(BeNil()))
		})

		It("should have enums", func() {
			const (
				MONDAY = 1 + iota
				TUESDAY
				WEDNESDAY
				THURSDAY
				FRIDAY
				SATURDAY
				SUNDAY
			)

			Expect(MONDAY).To(Equal(1))
			Expect(THURSDAY).To(Equal(4))
		})

		Context("Type Check", func() {
			var (
				string_var     = "hello"
				genric_var any = string_var
			)

			It("should cast valid string", func() {
				/** Empty Interface */
				casted_var, ok := genric_var.(string) //Type Casting
				Expect(ok).To(BeTrue())
				Expect(casted_var).To(Equal(string_var))
			})

			It("should not cast invalid float", func() {
				_, ok := genric_var.(float64) // Test Statement
				Expect(ok).To(BeFalse())
			})

			It("should fail string convert", func() {
				i, err := strconv.Atoi("XX")
				Expect(err).To(Not(BeNil()))
				Expect(i).To(Equal(0))
			})
		})

		Context("Struct", func() {
			var (
				vertex        Vertex
				pointerVertex *Vertex
			)

			BeforeEach(func() {
				vertex = Vertex{1, 2}
				pointerVertex = &vertex
			})

			It("should only mutate copy", func() {
				vertexCopy := vertex
				vertexCopy.Y = 7
				Expect(vertexCopy.Y).To(Equal(float64(7)))
				Expect(vertex.Y).To(Equal(float64(2)))
			})

			It("should mutate original on pointer", func() {
				pointerVertex.Y = 9
				Expect(vertex.Y).To(Equal(float64(9)))
			})

			Context("Interface", func() {

				var (
					a Abser
				)

				BeforeEach(func() {
					/** Interface */
					// a=vertex /** Gives Error as Abs takes only Pointer */
					a = pointerVertex
				})

				It("should compute absolute", func() {
					// /* While methods with pointer receivers take either a value or a pointer as the receiver when they are called: */
					Expect(vertex.Abs()).To(Equal(pointerVertex.Abs()))
					Expect(vertex.Abs()).To(Equal(a.Abs()))

					/** Null Handling pver.Abs() Would still work but when Abs will try to access X,Y Null Pointer would come. */
					// pver = nil
					// pver.Abs()
					/** Null Handling on Type would be error even if error is called on concrete type */
					// vertex = nil
					// vertex.Abs()
				})
			})
		})

		// fmt.Printf("Variables Type: %T Value: %v\n", r, r)
		// fmt.Printf("Variables Type: %T Value: %v\n", i, i)
		// fmt.Printf("Variables Type: %T Value: %v\n", g, g)

	})

	Context("Pointers", func() {
		var (
			i, j = 42, 2701
		)
		It("should resolve", func() {
			p := &i                    // point to i
			Expect(p).To(Not(BeNil())) // Address of i (Value of p)
			Expect(*p).To(Equal(i))    // read i through the pointer
			*p = 21                    // set i through the pointer
			Expect(i).To(Equal(21))    // see the new value of i
		})

		It("should overwrite", func() {
			p := &j                 // point to j
			*p = *p / 37            // divide j through the pointer
			Expect(j).To(Equal(73)) // see the new value of j
		})

	})

	Context("Collection", func() {
		var (
			a      [2]string
			primes = [6]int{2, 3, 5, 7, 11, 13}
		)

		It("should insert values", func() {
			a[0] = "Hello"
			a[1] = "World"
			Expect(a).To(HaveLen(2))
		})

		It("can be slices", func() {
			slice := primes[1:4]
			Expect(slice).To(Equal([]int{3, 5, 7}))

			// Len of slice is count of elements that have been sliced
			Expect(len(slice)).To(Equal(3))
			//The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
			Expect(cap(slice)).To(Equal(5))
		})

		It("can be two dimensional", func() {
			/** Two Dimensional */
			var twod [5][5]uint8 //Array 5x5
			twod[1][1] = 5
			Expect(twod[1][1]).To(Equal(uint8(5)))
			Expect(twod[3][4]).To(Equal(uint8(0)))

			Expect(len(twod)).To(Equal(5))
			Expect(cap(twod)).To(Equal(5))
		})

		It("can hold struct", func() {
			/** Slice referencing the Array as no Size is Specified for Struct Array */
			structSlice := []struct {
				i int
				b bool
			}{{2, true}, {3, false}}

			Expect(len(structSlice)).To(Equal(2))
			Expect(cap(structSlice)).To(Equal(2))
		})

		It("can be map", func() {
			hashMap := map[string]int{"One": 1, "Two": 2}
			v2, ok := hashMap["Two"]
			Expect(ok).To(BeTrue()) //Ok Holds if element is present or not.
			Expect(v2).To(Equal(2))
		})

		It("can be built via make and new", func() {
			make_slice := make([]int, 50, 100)
			Expect(len(make_slice)).To(Equal(50))
			Expect(cap(make_slice)).To(Equal(100))

			new_slice := new([100]int)[0:50]
			Expect(len(new_slice)).To(Equal(50))
			Expect(cap(new_slice)).To(Equal(100))
		})

		It("should count words", func() {
			word_count := WordCount("Hello World Hello Aman")
			Expect(word_count["Aman"]).To(Equal(1))
			Expect(word_count["Hello"]).To(Equal(2))
		})
	})

	Context("Switch", func() {
		It("should tell os", func() {
			switch os := runtime.GOOS; os {
			case "darwin":
				fmt.Println("OS X.")
			//fallthrough //implicit break if fallthrough not added
			case "linux":
				fmt.Println("Linux.")
			default:
				// freebsd, openbsd,
				// plan9, windows...
				fmt.Printf("%s.", os)
			}
		})

		It("should tell weekend", func() {
			/** Emulates long if/else chains */
			fmt.Print("When's Saturday? ")
			today := time.Now().Weekday()
			switch time.Saturday {
			case today + 0:
				fmt.Println("Today.")
			case today + 1:
				fmt.Println("Tomorrow.")
			case today + 2:
				fmt.Println("In two days.")
			default:
				fmt.Println("Too far away.")
			}

		})
	})

	Context("Math", func() {
		const (
			// Create a huge number by shifting a 1 bit left 100 places.
			// In other words, the binary number that is 1 followed by 100 zeroes.
			Big = 1 << 100
			// Shift it right again 99 places, so we end up with 1<<1, or 2.
			Small = Big >> 99
		)

		It("should work", func() {
			//Small is 2 and Big is 1^100
			Expect(needInt(Small)).To(Equal(21))
			Expect(needFloat(Small)).To(Equal(float64(0.2)))
			Expect(needFloat(Big)).To(BeNumerically(">", (float64(1.26765))))
		})

		It("should compute sqrt", func() {
			input := 8
			result, err := sqrt(input)
			Expect(err).To(BeNil())
			Expect(result).To(Equal(math.Sqrt((float64(input)))))
		})

		It("should not compute negative sqrt", func() {
			input := -2
			_, err := sqrt(input)
			Expect(err).To(Not(BeNil()))
		})

		It("should generate random", func() {
			rand.Seed(time.Now().UnixNano())
			Expect(rand.Intn(10)).To(Not(BeNil()))
		})

	})

	Context("Regex", func() {
		var (
			mysqlString  = "aman:aman@tcp(mysql:3306)/compute?charset=utf8&parseTime=True&loc=Local"
			mysqlMatcher = regexp.MustCompile(`^(.*)\((.*)\)(.*)$`)
		)

		It("should match", func() {
			matched := mysqlMatcher.FindAllStringSubmatch(mysqlString, 5)
			Expect(matched[0]).To(HaveLen(4))
			Expect(matched[0][2]).To(Equal("mysql:3306"))
			Expect(mysqlMatcher.ReplaceAllString(mysqlString, `$1#$2`)).To(Equal("aman:aman@tcp#mysql:3306"))
		})
	})

	Context("Loops", func() {
		It("should do iteration", func() {
			sum := 0
			count := 10
			for i := 0; i < count; i++ {
				sum += 1
			}
			Expect(sum).To(Equal(count))
		})

		It("should break infinite", func() {
			sum := 0
			for {
				/* Exit Condition */
				if sum += 50; sum > 300 {
					break
				}
			}
			Expect(sum).To(Equal(350))
		})
	})

	Context("Lamda", func() {
		It("should pass parameters", func() {
			Expect(lamdbaCompute(triple, 5)).To(Equal(15))
		})

		It("should demonstrate closure", func() {
			// Start with New Adders with Zero State
			pos, neg := adder(), adder()

			//Run Closures in opposite directions
			//with varied speeds
			for i := 0; i < 10; i++ {
				pos(i)
				neg(-2 * i)
			}

			//Match State stored in closure
			Expect(pos(0)).To(Equal(45))
			Expect(neg(0)).To(Equal(-90))
		})

		It("should have fibonacci", func() {
			f := fibonacciLambda()
			for i := 0; i < 10; i++ {
				f()
			}
			Expect(f()).To(Equal(55))
		})
	})

	Context("Defer", func() {

		It("should change message", func() {
			message := "Captured Argument"
			//Arguments Captured but will be executed at end.
			defer Expect(message).To(Equal("Captured Argument"))
			message = "Now Changed"
			Expect(message).To(Equal("Now Changed"))
		})

		It("should increment", func() {
			Expect(deferReturn()).To(Equal(5))
		})
	})
})

/* Structs */
type Abser interface {
	Abs() float64
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

/* Math */
func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func sqrt(x int) (float64, error) {
	if x < 0 {
		return math.NaN(), ErrNegativeSqrt(x)
	}
	fX := float64(x)
	z := float64(1)
	z = 1.0
	for i := 0; i < 10; i++ {
		z = z - ((math.Pow(z, 2) - fX) / (2 * z))
	}
	return z, nil
}

/* Error Handling */
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negativ number: %g", float64(e))
}

/* Lambda Fun */
type convert func(int) int

func triple(i int) int {
	return i * 3
}

func lamdbaCompute(convert_function convert, x int) (result int) {
	result = convert_function(x)
	return
}

/**
	Adds or Subtracts number passed to it
	from sum which starts from Zero.

	State is preserved for Sum due to closure
**/
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacciLambda() func() int {
	lastFibBeforeUpdate, lastFib := 0, 0
	fib := 1
	return func() int {
		lastFibBeforeUpdate, lastFib, fib = lastFib, fib, lastFib+fib // Simultaneous Assignment :D
		return lastFibBeforeUpdate
	}
}

/* Collections */
func WordCount(input string) map[string]int {
	countMap := make(map[string]int)
	fmt.Println(countMap)
	fields := strings.Fields(input)
	/** Ranges where i is optional can use _,v */
	for _, f := range fields {
		countMap[f] += 1 //No NPE :), No Init Required because entry value is primitive
	}
	return countMap
}

/* Defer */
func deferReturn() (i int) {
	defer func() {
		i++
	}()
	return 4
}
